/*
Copyright The Kubernetes Authors.

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
	time "time"

	appsv1beta1 "github.com/yext/api/apps/v1beta1"
	v1 "github.com/yext/apimachinery/pkg/apis/meta/v1"
	runtime "github.com/yext/apimachinery/pkg/runtime"
	watch "github.com/yext/apimachinery/pkg/watch"
	internalinterfaces "github.com/yext/client-go/informers/internalinterfaces"
	kubernetes "github.com/yext/client-go/kubernetes"
	v1beta1 "github.com/yext/client-go/listers/apps/v1beta1"
	cache "github.com/yext/client-go/tools/cache"
)

// ControllerRevisionInformer provides access to a shared informer and lister for
// ControllerRevisions.
type ControllerRevisionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ControllerRevisionLister
}

type controllerRevisionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewControllerRevisionInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewControllerRevisionInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredControllerRevisionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredControllerRevisionInformer constructs a new informer for ControllerRevision type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredControllerRevisionInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1beta1().ControllerRevisions(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1beta1().ControllerRevisions(namespace).Watch(options)
			},
		},
		&appsv1beta1.ControllerRevision{},
		resyncPeriod,
		indexers,
	)
}

func (f *controllerRevisionInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredControllerRevisionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *controllerRevisionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appsv1beta1.ControllerRevision{}, f.defaultInformer)
}

func (f *controllerRevisionInformer) Lister() v1beta1.ControllerRevisionLister {
	return v1beta1.NewControllerRevisionLister(f.Informer().GetIndexer())
}
