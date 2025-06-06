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

package v1alpha1

import (
	resourcev1alpha1 "github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// PipelineResourceLister helps list PipelineResources.
// All objects returned here must be treated as read-only.
type PipelineResourceLister interface {
	// List lists all PipelineResources in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*resourcev1alpha1.PipelineResource, err error)
	// PipelineResources returns an object that can list and get PipelineResources.
	PipelineResources(namespace string) PipelineResourceNamespaceLister
	PipelineResourceListerExpansion
}

// pipelineResourceLister implements the PipelineResourceLister interface.
type pipelineResourceLister struct {
	listers.ResourceIndexer[*resourcev1alpha1.PipelineResource]
}

// NewPipelineResourceLister returns a new PipelineResourceLister.
func NewPipelineResourceLister(indexer cache.Indexer) PipelineResourceLister {
	return &pipelineResourceLister{listers.New[*resourcev1alpha1.PipelineResource](indexer, resourcev1alpha1.Resource("pipelineresource"))}
}

// PipelineResources returns an object that can list and get PipelineResources.
func (s *pipelineResourceLister) PipelineResources(namespace string) PipelineResourceNamespaceLister {
	return pipelineResourceNamespaceLister{listers.NewNamespaced[*resourcev1alpha1.PipelineResource](s.ResourceIndexer, namespace)}
}

// PipelineResourceNamespaceLister helps list and get PipelineResources.
// All objects returned here must be treated as read-only.
type PipelineResourceNamespaceLister interface {
	// List lists all PipelineResources in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*resourcev1alpha1.PipelineResource, err error)
	// Get retrieves the PipelineResource from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*resourcev1alpha1.PipelineResource, error)
	PipelineResourceNamespaceListerExpansion
}

// pipelineResourceNamespaceLister implements the PipelineResourceNamespaceLister
// interface.
type pipelineResourceNamespaceLister struct {
	listers.ResourceIndexer[*resourcev1alpha1.PipelineResource]
}
