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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	pipelinev1beta1 "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// CustomRunLister helps list CustomRuns.
// All objects returned here must be treated as read-only.
type CustomRunLister interface {
	// List lists all CustomRuns in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*pipelinev1beta1.CustomRun, err error)
	// CustomRuns returns an object that can list and get CustomRuns.
	CustomRuns(namespace string) CustomRunNamespaceLister
	CustomRunListerExpansion
}

// customRunLister implements the CustomRunLister interface.
type customRunLister struct {
	listers.ResourceIndexer[*pipelinev1beta1.CustomRun]
}

// NewCustomRunLister returns a new CustomRunLister.
func NewCustomRunLister(indexer cache.Indexer) CustomRunLister {
	return &customRunLister{listers.New[*pipelinev1beta1.CustomRun](indexer, pipelinev1beta1.Resource("customrun"))}
}

// CustomRuns returns an object that can list and get CustomRuns.
func (s *customRunLister) CustomRuns(namespace string) CustomRunNamespaceLister {
	return customRunNamespaceLister{listers.NewNamespaced[*pipelinev1beta1.CustomRun](s.ResourceIndexer, namespace)}
}

// CustomRunNamespaceLister helps list and get CustomRuns.
// All objects returned here must be treated as read-only.
type CustomRunNamespaceLister interface {
	// List lists all CustomRuns in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*pipelinev1beta1.CustomRun, err error)
	// Get retrieves the CustomRun from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*pipelinev1beta1.CustomRun, error)
	CustomRunNamespaceListerExpansion
}

// customRunNamespaceLister implements the CustomRunNamespaceLister
// interface.
type customRunNamespaceLister struct {
	listers.ResourceIndexer[*pipelinev1beta1.CustomRun]
}
