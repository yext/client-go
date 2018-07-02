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

package v1

import (
	time "time"

	networkingv1 "github.com/yext/api/networking/v1"
	metav1 "github.com/yext/apimachinery/pkg/apis/meta/v1"
	runtime "github.com/yext/apimachinery/pkg/runtime"
	watch "github.com/yext/apimachinery/pkg/watch"
	internalinterfaces "github.com/yext/client-go/informers/internalinterfaces"
	kubernetes "github.com/yext/client-go/kubernetes"
	v1 "github.com/yext/client-go/listers/networking/v1"
	cache "github.com/yext/client-go/tools/cache"
)

// NetworkPolicyInformer provides access to a shared informer and lister for
// NetworkPolicies.
type NetworkPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.NetworkPolicyLister
}

type networkPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewNetworkPolicyInformer constructs a new informer for NetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewNetworkPolicyInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredNetworkPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredNetworkPolicyInformer constructs a new informer for NetworkPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredNetworkPolicyInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().NetworkPolicies(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().NetworkPolicies(namespace).Watch(options)
			},
		},
		&networkingv1.NetworkPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *networkPolicyInformer) defaultInformer(client kubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredNetworkPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *networkPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1.NetworkPolicy{}, f.defaultInformer)
}

func (f *networkPolicyInformer) Lister() v1.NetworkPolicyLister {
	return v1.NewNetworkPolicyLister(f.Informer().GetIndexer())
}
