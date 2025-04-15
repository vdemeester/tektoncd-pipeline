/*
Copyright 2020 The Tekton Authors

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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	context "context"

	pipelinev1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1"
	scheme "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// TaskRunsGetter has a method to return a TaskRunInterface.
// A group's client should implement this interface.
type TaskRunsGetter interface {
	TaskRuns(namespace string) TaskRunInterface
}

// TaskRunInterface has methods to work with TaskRun resources.
type TaskRunInterface interface {
	Create(ctx context.Context, taskRun *pipelinev1.TaskRun, opts metav1.CreateOptions) (*pipelinev1.TaskRun, error)
	Update(ctx context.Context, taskRun *pipelinev1.TaskRun, opts metav1.UpdateOptions) (*pipelinev1.TaskRun, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, taskRun *pipelinev1.TaskRun, opts metav1.UpdateOptions) (*pipelinev1.TaskRun, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*pipelinev1.TaskRun, error)
	List(ctx context.Context, opts metav1.ListOptions) (*pipelinev1.TaskRunList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *pipelinev1.TaskRun, err error)
	TaskRunExpansion
}

// taskRuns implements TaskRunInterface
type taskRuns struct {
	*gentype.ClientWithList[*pipelinev1.TaskRun, *pipelinev1.TaskRunList]
}

// newTaskRuns returns a TaskRuns
func newTaskRuns(c *TektonV1Client, namespace string) *taskRuns {
	return &taskRuns{
		gentype.NewClientWithList[*pipelinev1.TaskRun, *pipelinev1.TaskRunList](
			"taskruns",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *pipelinev1.TaskRun { return &pipelinev1.TaskRun{} },
			func() *pipelinev1.TaskRunList { return &pipelinev1.TaskRunList{} },
		),
	}
}
