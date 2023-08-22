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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "kubevirt.io/api/clone/v1alpha1"
)

// VirtualMachineCloneLister helps list VirtualMachineClones.
// All objects returned here must be treated as read-only.
type VirtualMachineCloneLister interface {
	// List lists all VirtualMachineClones in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineClone, err error)
	// VirtualMachineClones returns an object that can list and get VirtualMachineClones.
	VirtualMachineClones(namespace string) VirtualMachineCloneNamespaceLister
	VirtualMachineCloneListerExpansion
}

// virtualMachineCloneLister implements the VirtualMachineCloneLister interface.
type virtualMachineCloneLister struct {
	indexer cache.Indexer
}

// NewVirtualMachineCloneLister returns a new VirtualMachineCloneLister.
func NewVirtualMachineCloneLister(indexer cache.Indexer) VirtualMachineCloneLister {
	return &virtualMachineCloneLister{indexer: indexer}
}

// List lists all VirtualMachineClones in the indexer.
func (s *virtualMachineCloneLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineClone, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineClone))
	})
	return ret, err
}

// VirtualMachineClones returns an object that can list and get VirtualMachineClones.
func (s *virtualMachineCloneLister) VirtualMachineClones(namespace string) VirtualMachineCloneNamespaceLister {
	return virtualMachineCloneNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualMachineCloneNamespaceLister helps list and get VirtualMachineClones.
// All objects returned here must be treated as read-only.
type VirtualMachineCloneNamespaceLister interface {
	// List lists all VirtualMachineClones in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineClone, err error)
	// Get retrieves the VirtualMachineClone from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VirtualMachineClone, error)
	VirtualMachineCloneNamespaceListerExpansion
}

// virtualMachineCloneNamespaceLister implements the VirtualMachineCloneNamespaceLister
// interface.
type virtualMachineCloneNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualMachineClones in the indexer for a given namespace.
func (s virtualMachineCloneNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualMachineClone, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualMachineClone))
	})
	return ret, err
}

// Get retrieves the VirtualMachineClone from the indexer for a given namespace and name.
func (s virtualMachineCloneNamespaceLister) Get(name string) (*v1alpha1.VirtualMachineClone, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualmachineclone"), name)
	}
	return obj.(*v1alpha1.VirtualMachineClone), nil
}
