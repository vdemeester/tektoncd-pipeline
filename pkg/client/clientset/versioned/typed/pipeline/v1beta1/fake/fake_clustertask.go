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

package fake

import (
	v1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	pipelinev1beta1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1beta1"
	gentype "k8s.io/client-go/gentype"
)

// fakeClusterTasks implements ClusterTaskInterface
type fakeClusterTasks struct {
	*gentype.FakeClientWithList[*v1beta1.ClusterTask, *v1beta1.ClusterTaskList]
	Fake *FakeTektonV1beta1
}

func newFakeClusterTasks(fake *FakeTektonV1beta1) pipelinev1beta1.ClusterTaskInterface {
	return &fakeClusterTasks{
		gentype.NewFakeClientWithList[*v1beta1.ClusterTask, *v1beta1.ClusterTaskList](
			fake.Fake,
			"",
			v1beta1.SchemeGroupVersion.WithResource("clustertasks"),
			v1beta1.SchemeGroupVersion.WithKind("ClusterTask"),
			func() *v1beta1.ClusterTask { return &v1beta1.ClusterTask{} },
			func() *v1beta1.ClusterTaskList { return &v1beta1.ClusterTaskList{} },
			func(dst, src *v1beta1.ClusterTaskList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.ClusterTaskList) []*v1beta1.ClusterTask { return gentype.ToPointerSlice(list.Items) },
			func(list *v1beta1.ClusterTaskList, items []*v1beta1.ClusterTask) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
