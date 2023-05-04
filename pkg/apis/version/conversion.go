/*
Copyright 2022 The Tekton Authors

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

package version

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SerializeToMetadata serializes the input field and adds it as an annotation to
// the metadata under the input key.
func SerializeToMetadata(meta *metav1.ObjectMeta, field interface{}, key string) error {
	data, err := json.Marshal(field)
	if err != nil {
		return fmt.Errorf("error serializing field: %w", err)
	}
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)
	if _, err := gz.Write(data); err != nil {
		return err
	}
	if err := gz.Close(); err != nil {
		return err
	}
	compressedAndEncoded := base64.StdEncoding.EncodeToString(b.Bytes())
	if meta.Annotations == nil {
		meta.Annotations = make(map[string]string)
	}
	meta.Annotations[key] = string(compressedAndEncoded)
	return nil
}

// DeserializeFromMetadata takes the value of the input key from the metadata's annotations,
// deserializes it into "to", and removes the key from the metadata's annotations.
// Returns nil if the key is not present in the annotations.
func DeserializeFromMetadata(meta *metav1.ObjectMeta, to interface{}, key string) error {
	if meta.Annotations == nil {
		return nil
	}
	if str, ok := meta.Annotations[key]; ok {
		decoded, err := base64.StdEncoding.DecodeString(str)
		if err != nil {
			return err
		}
		gz, err := gzip.NewReader(bytes.NewReader([]byte(decoded)))
		if err != nil {
			return err
		}
		data, err := ioutil.ReadAll(gz)
		if err != nil {
			return err
		}
		if err := gz.Close(); err != nil {
			return err
		}
		if err != nil {
			return fmt.Errorf("error decoding string from encoded marshalled bytes %w", err)
		}
		if err := json.Unmarshal(data, to); err != nil {
			return fmt.Errorf("error deserializing key %s from metadata: %w", key, err)
		}
		delete(meta.Annotations, key)
		if len(meta.Annotations) == 0 {
			meta.Annotations = nil
		}
	}
	return nil
}
