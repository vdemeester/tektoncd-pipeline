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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"
	time "time"

	apisresolutionv1beta1 "github.com/tektoncd/pipeline/pkg/apis/resolution/v1beta1"
	versioned "github.com/tektoncd/pipeline/pkg/client/resolution/clientset/versioned"
	internalinterfaces "github.com/tektoncd/pipeline/pkg/client/resolution/informers/externalversions/internalinterfaces"
	resolutionv1beta1 "github.com/tektoncd/pipeline/pkg/client/resolution/listers/resolution/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ResolutionRequestInformer provides access to a shared informer and lister for
// ResolutionRequests.
type ResolutionRequestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() resolutionv1beta1.ResolutionRequestLister
}

type resolutionRequestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResolutionRequestInformer constructs a new informer for ResolutionRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResolutionRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResolutionRequestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResolutionRequestInformer constructs a new informer for ResolutionRequest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResolutionRequestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResolutionV1beta1().ResolutionRequests(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ResolutionV1beta1().ResolutionRequests(namespace).Watch(context.TODO(), options)
			},
		},
		&apisresolutionv1beta1.ResolutionRequest{},
		resyncPeriod,
		indexers,
	)
}

func (f *resolutionRequestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResolutionRequestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resolutionRequestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisresolutionv1beta1.ResolutionRequest{}, f.defaultInformer)
}

func (f *resolutionRequestInformer) Lister() resolutionv1beta1.ResolutionRequestLister {
	return resolutionv1beta1.NewResolutionRequestLister(f.Informer().GetIndexer())
}
