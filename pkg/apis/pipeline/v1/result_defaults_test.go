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

package v1_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	v1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	"github.com/tektoncd/pipeline/test/diff"
)

func TestTaskResult_SetDefaults(t *testing.T) {
	tests := []struct {
		name   string
		before *v1.TaskResult
		after  *v1.TaskResult
	}{{
		name:   "empty taskresult",
		before: nil,
		after:  nil,
	}, {
		name: "inferred string type",
		before: &v1.TaskResult{
			Name: "resultname",
		},
		after: &v1.TaskResult{
			Name: "resultname",
			Type: v1.ResultsTypeString,
		},
	}, {
		name: "string type specified not changed",
		before: &v1.TaskResult{
			Name: "resultname",
			Type: v1.ResultsTypeString,
		},
		after: &v1.TaskResult{
			Name: "resultname",
			Type: v1.ResultsTypeString,
		},
	}, {
		name: "array type specified not changed",
		before: &v1.TaskResult{
			Name: "resultname",
			Type: v1.ResultsTypeArray,
		},
		after: &v1.TaskResult{
			Name: "resultname",
			Type: v1.ResultsTypeArray,
		},
	}, {
		name: "inferred object type from properties - PropertySpec type is provided",
		before: &v1.TaskResult{
			Name:       "resultname",
			Properties: map[string]v1.PropertySpec{"key1": {v1.ParamTypeString}},
		},
		after: &v1.TaskResult{
			Name:       "resultname",
			Type:       v1.ResultsTypeObject,
			Properties: map[string]v1.PropertySpec{"key1": {v1.ParamTypeString}},
		},
	}, {
		name: "inferred type from properties - PropertySpec type is not provided",
		before: &v1.TaskResult{
			Name:       "resultname",
			Properties: map[string]v1.PropertySpec{"key1": {}},
		},
		after: &v1.TaskResult{
			Name:       "resultname",
			Type:       v1.ResultsTypeObject,
			Properties: map[string]v1.PropertySpec{"key1": {v1.ParamTypeString}},
		},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := t.Context()
			tc.before.SetDefaults(ctx)
			if d := cmp.Diff(tc.after, tc.before); d != "" {
				t.Error(diff.PrintWantGot(d))
			}
		})
	}
}

func TestStepResult_SetDefaults(t *testing.T) {
	tests := []struct {
		name   string
		before *v1.StepResult
		after  *v1.StepResult
	}{{
		name:   "empty taskresult",
		before: nil,
		after:  nil,
	}, {
		name: "inferred string type",
		before: &v1.StepResult{
			Name: "resultname",
		},
		after: &v1.StepResult{
			Name: "resultname",
			Type: v1.ResultsTypeString,
		},
	}, {
		name: "string type specified not changed",
		before: &v1.StepResult{
			Name: "resultname",
			Type: v1.ResultsTypeString,
		},
		after: &v1.StepResult{
			Name: "resultname",
			Type: v1.ResultsTypeString,
		},
	}, {
		name: "array type specified not changed",
		before: &v1.StepResult{
			Name: "resultname",
			Type: v1.ResultsTypeArray,
		},
		after: &v1.StepResult{
			Name: "resultname",
			Type: v1.ResultsTypeArray,
		},
	}, {
		name: "inferred object type from properties - PropertySpec type is provided",
		before: &v1.StepResult{
			Name:       "resultname",
			Properties: map[string]v1.PropertySpec{"key1": {Type: v1.ParamTypeString}},
		},
		after: &v1.StepResult{
			Name:       "resultname",
			Type:       v1.ResultsTypeObject,
			Properties: map[string]v1.PropertySpec{"key1": {Type: v1.ParamTypeString}},
		},
	}, {
		name: "inferred type from properties - PropertySpec type is not provided",
		before: &v1.StepResult{
			Name:       "resultname",
			Properties: map[string]v1.PropertySpec{"key1": {}},
		},
		after: &v1.StepResult{
			Name:       "resultname",
			Type:       v1.ResultsTypeObject,
			Properties: map[string]v1.PropertySpec{"key1": {Type: v1.ParamTypeString}},
		},
	}}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ctx := t.Context()
			tc.before.SetDefaults(ctx)
			if d := cmp.Diff(tc.after, tc.before); d != "" {
				t.Error(diff.PrintWantGot(d))
			}
		})
	}
}
