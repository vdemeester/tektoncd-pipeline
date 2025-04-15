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
	v1alpha1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1alpha1"
	pipelinev1alpha1 "github.com/tektoncd/pipeline/pkg/client/clientset/versioned/typed/pipeline/v1alpha1"
	gentype "k8s.io/client-go/gentype"
)

// fakeStepActions implements StepActionInterface
type fakeStepActions struct {
	*gentype.FakeClientWithList[*v1alpha1.StepAction, *v1alpha1.StepActionList]
	Fake *FakeTektonV1alpha1
}

func newFakeStepActions(fake *FakeTektonV1alpha1, namespace string) pipelinev1alpha1.StepActionInterface {
	return &fakeStepActions{
		gentype.NewFakeClientWithList[*v1alpha1.StepAction, *v1alpha1.StepActionList](
			fake.Fake,
			namespace,
			v1alpha1.SchemeGroupVersion.WithResource("stepactions"),
			v1alpha1.SchemeGroupVersion.WithKind("StepAction"),
			func() *v1alpha1.StepAction { return &v1alpha1.StepAction{} },
			func() *v1alpha1.StepActionList { return &v1alpha1.StepActionList{} },
			func(dst, src *v1alpha1.StepActionList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.StepActionList) []*v1alpha1.StepAction { return gentype.ToPointerSlice(list.Items) },
			func(list *v1alpha1.StepActionList, items []*v1alpha1.StepAction) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
