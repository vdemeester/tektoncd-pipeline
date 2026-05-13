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

package pod

import (
	"encoding/json"
	"testing"

	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/pkg/entrypoint"
)

func TestArtifactEntrypointArgs(t *testing.T) {
	taskSpec := &v1.TaskSpec{
		Artifacts: &v1.ArtifactDeclarations{
			Outputs: []v1.ArtifactDeclaration{
				{Name: "sbom", MediaType: "application/vnd.cyclonedx+json"},
			},
		},
	}

	repository := "registry:5000/artifacts"
	insecure := true

	args := artifactEntrypointArgs(taskSpec, repository, insecure)

	// Should contain -artifact_outputs flag
	found := false
	for i, arg := range args {
		if arg == "-artifact_outputs" {
			found = true
			if i+1 >= len(args) {
				t.Fatal("-artifact_outputs flag has no value")
			}
			var outputs []entrypoint.ArtifactOutput
			if err := json.Unmarshal([]byte(args[i+1]), &outputs); err != nil {
				t.Fatalf("failed to parse artifact outputs JSON: %v", err)
			}
			if len(outputs) != 1 {
				t.Fatalf("expected 1 output, got %d", len(outputs))
			}
			if outputs[0].Name != "sbom" {
				t.Errorf("expected name 'sbom', got %q", outputs[0].Name)
			}
			if outputs[0].MediaType != "application/vnd.cyclonedx+json" {
				t.Errorf("expected mediaType 'application/vnd.cyclonedx+json', got %q", outputs[0].MediaType)
			}
			if outputs[0].Repository != "registry:5000/artifacts/sbom" {
				t.Errorf("expected repository 'registry:5000/artifacts/sbom', got %q", outputs[0].Repository)
			}
		}
	}
	if !found {
		t.Error("expected -artifact_outputs flag in args")
	}

	// Should contain -artifact_insecure flag
	foundInsecure := false
	for _, arg := range args {
		if arg == "-artifact_insecure" {
			foundInsecure = true
		}
	}
	if !foundInsecure {
		t.Error("expected -artifact_insecure flag in args")
	}
}

func TestArtifactEntrypointArgs_NoArtifacts(t *testing.T) {
	taskSpec := &v1.TaskSpec{}
	args := artifactEntrypointArgs(taskSpec, "registry:5000/artifacts", false)
	if len(args) != 0 {
		t.Errorf("expected empty args for task without artifacts, got %v", args)
	}
}
