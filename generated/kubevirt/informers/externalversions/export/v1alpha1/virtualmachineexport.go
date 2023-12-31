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

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	exportv1alpha1 "kubevirt.io/api/export/v1alpha1"
	versioned "kubevirt.io/client-go/generated/kubevirt/clientset/versioned"
	internalinterfaces "kubevirt.io/client-go/generated/kubevirt/informers/externalversions/internalinterfaces"
	v1alpha1 "kubevirt.io/client-go/generated/kubevirt/listers/export/v1alpha1"
)

// VirtualMachineExportInformer provides access to a shared informer and lister for
// VirtualMachineExports.
type VirtualMachineExportInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.VirtualMachineExportLister
}

type virtualMachineExportInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewVirtualMachineExportInformer constructs a new informer for VirtualMachineExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewVirtualMachineExportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredVirtualMachineExportInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredVirtualMachineExportInformer constructs a new informer for VirtualMachineExport type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredVirtualMachineExportInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExportV1alpha1().VirtualMachineExports(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExportV1alpha1().VirtualMachineExports(namespace).Watch(context.TODO(), options)
			},
		},
		&exportv1alpha1.VirtualMachineExport{},
		resyncPeriod,
		indexers,
	)
}

func (f *virtualMachineExportInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredVirtualMachineExportInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *virtualMachineExportInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&exportv1alpha1.VirtualMachineExport{}, f.defaultInformer)
}

func (f *virtualMachineExportInformer) Lister() v1alpha1.VirtualMachineExportLister {
	return v1alpha1.NewVirtualMachineExportLister(f.Informer().GetIndexer())
}
