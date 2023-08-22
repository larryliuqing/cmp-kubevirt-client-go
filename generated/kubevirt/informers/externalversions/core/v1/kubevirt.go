/*
Copyright 2023 The KubeVirt Authors.

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
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	corev1 "kubevirt.io/api/core/v1"
	versioned "kubevirt.io/client-go/generated/kubevirt/clientset/versioned"
	internalinterfaces "kubevirt.io/client-go/generated/kubevirt/informers/externalversions/internalinterfaces"
	v1 "kubevirt.io/client-go/generated/kubevirt/listers/core/v1"
)

// KubeVirtInformer provides access to a shared informer and lister for
// KubeVirts.
type KubeVirtInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.KubeVirtLister
}

type kubeVirtInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKubeVirtInformer constructs a new informer for KubeVirt type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubeVirtInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubeVirtInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKubeVirtInformer constructs a new informer for KubeVirt type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubeVirtInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubevirtV1().KubeVirts(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubevirtV1().KubeVirts(namespace).Watch(context.TODO(), options)
			},
		},
		&corev1.KubeVirt{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubeVirtInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubeVirtInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubeVirtInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1.KubeVirt{}, f.defaultInformer)
}

func (f *kubeVirtInformer) Lister() v1.KubeVirtLister {
	return v1.NewKubeVirtLister(f.Informer().GetIndexer())
}
