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
	"fmt"
	"path/filepath"

	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
)

// artifactEntrypointArgs generates entrypoint CLI args for artifact transport
// based on the TaskSpec's artifact declarations and the artifact storage config.
func artifactEntrypointArgs(taskSpec *v1.TaskSpec, ociRepository string, insecure bool) []string {
	if taskSpec == nil || taskSpec.Artifacts == nil {
		return nil
	}

	var args []string

	if len(taskSpec.Artifacts.Inputs) > 0 {
		type artifactInput struct {
			Name string `json:"name"`
			URI  string `json:"uri"`
			Path string `json:"path"`
		}
		var inputs []artifactInput
		for _, decl := range taskSpec.Artifacts.Inputs {
			inputs = append(inputs, artifactInput{
				Name: decl.Name,
				// URI will be resolved at runtime from pipeline bindings or params
				Path: filepath.Join("/workspace/artifacts/inputs", decl.Name),
			})
		}
		data, _ := json.Marshal(inputs)
		args = append(args, "-artifact_inputs", string(data))
	}

	if len(taskSpec.Artifacts.Outputs) > 0 {
		type artifactOutput struct {
			Name        string `json:"name"`
			Path        string `json:"path"`
			Repository  string `json:"repository"`
			MediaType   string `json:"mediaType"`
			BuildOutput bool   `json:"buildOutput,omitempty"`
		}
		var outputs []artifactOutput
		for _, decl := range taskSpec.Artifacts.Outputs {
			outputs = append(outputs, artifactOutput{
				Name:        decl.Name,
				Path:        filepath.Join("/workspace/artifacts/outputs", decl.Name),
				Repository:  fmt.Sprintf("%s/%s", ociRepository, decl.Name),
				MediaType:   decl.MediaType,
				BuildOutput: decl.BuildOutput,
			})
		}
		data, _ := json.Marshal(outputs)
		args = append(args, "-artifact_outputs", string(data))
	}

	if insecure {
		args = append(args, "-artifact_insecure")
	}

	return args
}
