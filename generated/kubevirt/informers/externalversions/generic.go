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

package externalversions

import (
	"fmt"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
	v1alpha1 "kubevirt.io/api/clone/v1alpha1"
	v1 "kubevirt.io/api/core/v1"
	exportv1alpha1 "kubevirt.io/api/export/v1alpha1"
	migrationsv1alpha1 "kubevirt.io/api/migrations/v1alpha1"
	poolv1alpha1 "kubevirt.io/api/pool/v1alpha1"
	snapshotv1alpha1 "kubevirt.io/api/snapshot/v1alpha1"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=clone.kubevirt.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("virtualmachineclones"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clone().V1alpha1().VirtualMachineClones().Informer()}, nil

		// Group=export.kubevirt.io, Version=v1alpha1
	case exportv1alpha1.SchemeGroupVersion.WithResource("virtualmachineexports"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Export().V1alpha1().VirtualMachineExports().Informer()}, nil

		// Group=kubevirt.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("kubevirts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubevirt().V1().KubeVirts().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualmachines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubevirt().V1().VirtualMachines().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualmachineinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubevirt().V1().VirtualMachineInstances().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualmachineinstancemigrations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubevirt().V1().VirtualMachineInstanceMigrations().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualmachineinstancepresets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubevirt().V1().VirtualMachineInstancePresets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("virtualmachineinstancereplicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubevirt().V1().VirtualMachineInstanceReplicaSets().Informer()}, nil

		// Group=migrations.kubevirt.io, Version=v1alpha1
	case migrationsv1alpha1.SchemeGroupVersion.WithResource("migrationpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Migrations().V1alpha1().MigrationPolicies().Informer()}, nil

		// Group=pool.kubevirt.io, Version=v1alpha1
	case poolv1alpha1.SchemeGroupVersion.WithResource("virtualmachinepools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Pool().V1alpha1().VirtualMachinePools().Informer()}, nil

		// Group=snapshot.kubevirt.io, Version=v1alpha1
	case snapshotv1alpha1.SchemeGroupVersion.WithResource("virtualmachinerestores"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Snapshot().V1alpha1().VirtualMachineRestores().Informer()}, nil
	case snapshotv1alpha1.SchemeGroupVersion.WithResource("virtualmachinesnapshots"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Snapshot().V1alpha1().VirtualMachineSnapshots().Informer()}, nil
	case snapshotv1alpha1.SchemeGroupVersion.WithResource("virtualmachinesnapshotcontents"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Snapshot().V1alpha1().VirtualMachineSnapshotContents().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
