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

package v1beta1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	versioned "kubevirt.io/client-go/generated/containerized-data-importer/clientset/versioned"
	internalinterfaces "kubevirt.io/client-go/generated/containerized-data-importer/informers/externalversions/internalinterfaces"
	v1beta1 "kubevirt.io/client-go/generated/containerized-data-importer/listers/core/v1beta1"
	corev1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// CDIConfigInformer provides access to a shared informer and lister for
// CDIConfigs.
type CDIConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.CDIConfigLister
}

type cDIConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCDIConfigInformer constructs a new informer for CDIConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCDIConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCDIConfigInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCDIConfigInformer constructs a new informer for CDIConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCDIConfigInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CdiV1beta1().CDIConfigs().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CdiV1beta1().CDIConfigs().Watch(context.TODO(), options)
			},
		},
		&corev1beta1.CDIConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *cDIConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCDIConfigInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cDIConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&corev1beta1.CDIConfig{}, f.defaultInformer)
}

func (f *cDIConfigInformer) Lister() v1beta1.CDIConfigLister {
	return v1beta1.NewCDIConfigLister(f.Informer().GetIndexer())
}
