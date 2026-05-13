/*
Copyright 2024 The Tekton Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package entrypoint

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/go-containerregistry/pkg/name"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	"github.com/google/go-containerregistry/pkg/v1/empty"
	"github.com/google/go-containerregistry/pkg/v1/mutate"
	"github.com/google/go-containerregistry/pkg/v1/remote"
	"github.com/google/go-containerregistry/pkg/v1/tarball"
	"github.com/google/go-containerregistry/pkg/v1/types"

	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

// ArtifactInput describes an artifact to download before step execution.
type ArtifactInput struct {
	Name string `json:"name"`       // artifact name
	URI  string `json:"uri"`        // OCI URI (e.g., "registry:5000/artifacts/test@sha256:...")
	Path string `json:"path"`       // local path to extract to
}

// ArtifactOutput describes an artifact to upload after step execution.
type ArtifactOutput struct {
	Name        string `json:"name"`                  // artifact name
	Path        string `json:"path"`                  // local path to archive from
	Repository  string `json:"repository"`            // OCI repository to push to
	MediaType   string `json:"mediaType"`             // artifact media type
	BuildOutput bool   `json:"buildOutput,omitempty"` // whether this is a primary build output
}

// UploadArtifact archives the contents at output.Path and pushes them as an OCI
// image with a single tar.gz layer to output.Repository. It returns an
// ArtifactValue with the pushed digest-based URI.
func UploadArtifact(ctx context.Context, output ArtifactOutput, insecure bool, opts ...remote.Option) (*pipelinev1.ArtifactValue, error) {
	// Create a tar.gz buffer from the output path
	buf, err := createTarGzBuffer(output.Path)
	if err != nil {
		return nil, fmt.Errorf("creating tar.gz archive: %w", err)
	}
	layer, err := tarball.LayerFromReader(buf, tarball.WithMediaType(types.MediaType(output.MediaType)))
	if err != nil {
		return nil, fmt.Errorf("creating layer: %w", err)
	}

	img, err := mutate.AppendLayers(empty.Image, layer)
	if err != nil {
		return nil, fmt.Errorf("creating image with layer: %w", err)
	}

	// Parse the repository reference
	nameOpts := []name.Option{}
	if insecure {
		nameOpts = append(nameOpts, name.Insecure)
	}
	repo, err := name.NewRepository(output.Repository, nameOpts...)
	if err != nil {
		return nil, fmt.Errorf("parsing repository %q: %w", output.Repository, err)
	}

	// Get the digest before pushing (this materializes the layer)
	digest, err := img.Digest()
	if err != nil {
		return nil, fmt.Errorf("computing digest: %w", err)
	}

	ref := repo.Digest(digest.String())

	// Push
	remoteOpts := append([]remote.Option{remote.WithContext(ctx)}, opts...)
	if err := remote.Write(ref, img, remoteOpts...); err != nil {
		return nil, fmt.Errorf("pushing artifact to %s: %w", ref.String(), err)
	}

	return &pipelinev1.ArtifactValue{
		Uri: ref.String(),
		Digest: map[pipelinev1.Algorithm]string{
			"sha256": strings.TrimPrefix(digest.String(), "sha256:"),
		},
	}, nil
}

// DownloadArtifact pulls an OCI artifact and extracts its layer contents to input.Path.
func DownloadArtifact(ctx context.Context, input ArtifactInput, insecure bool, opts ...remote.Option) error {
	nameOpts := []name.Option{}
	if insecure {
		nameOpts = append(nameOpts, name.Insecure)
	}
	ref, err := name.ParseReference(input.URI, nameOpts...)
	if err != nil {
		return fmt.Errorf("parsing artifact reference %q: %w", input.URI, err)
	}

	remoteOpts := append([]remote.Option{remote.WithContext(ctx)}, opts...)
	img, err := remote.Image(ref, remoteOpts...)
	if err != nil {
		return fmt.Errorf("pulling artifact from %s: %w", ref.String(), err)
	}

	layers, err := img.Layers()
	if err != nil {
		return fmt.Errorf("getting layers: %w", err)
	}

	for _, layer := range layers {
		if err := extractLayer(layer, input.Path); err != nil {
			return fmt.Errorf("extracting layer: %w", err)
		}
	}

	return nil
}

// createTarGzBuffer creates a tar.gz archive of the directory at path into a buffer.
func createTarGzBuffer(path string) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)
	tw := tar.NewWriter(gw)

	err := filepath.Walk(path, func(file string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(path, file)
		if err != nil {
			return err
		}
		if rel == "." {
			return nil
		}

		header, err := tar.FileInfoHeader(fi, "")
		if err != nil {
			return err
		}
		header.Name = rel

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		f, err := os.Open(file)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = io.Copy(tw, f)
		return err
	})
	if err != nil {
		return nil, err
	}

	if err := tw.Close(); err != nil {
		return nil, err
	}
	if err := gw.Close(); err != nil {
		return nil, err
	}

	return &buf, nil
}

// extractLayer extracts a tar.gz layer to the destination directory.
func extractLayer(layer v1.Layer, dst string) error {
	rc, err := layer.Compressed()
	if err != nil {
		return err
	}
	defer rc.Close()

	gr, err := gzip.NewReader(rc)
	if err != nil {
		return err
	}
	defer gr.Close()

	tr := tar.NewReader(gr)
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(dst, header.Name) //nolint:gosec // controlled input from our own uploads
		if !strings.HasPrefix(target, filepath.Clean(dst)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid tar entry: %s", header.Name)
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
				return err
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			h := sha256.New()
			if _, err := io.Copy(f, io.TeeReader(tr, h)); err != nil { //nolint:gosec // size bounded by OCI layer
				f.Close()
				return err
			}
			f.Close()
		}
	}

	return nil
}
