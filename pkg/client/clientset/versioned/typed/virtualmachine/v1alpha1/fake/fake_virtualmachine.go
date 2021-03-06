/*
Copyright 2018 Rancher Labs, Inc.

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
package fake

import (
	v1alpha1 "github.com/rancher/vm/pkg/apis/ranchervm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVirtualMachines implements VirtualMachineInterface
type FakeVirtualMachines struct {
	Fake *FakeVirtualmachineV1alpha1
}

var virtualmachinesResource = schema.GroupVersionResource{Group: "virtualmachine.rancher.com", Version: "v1alpha1", Resource: "virtualmachines"}

var virtualmachinesKind = schema.GroupVersionKind{Group: "virtualmachine.rancher.com", Version: "v1alpha1", Kind: "VirtualMachine"}

// Get takes name of the virtualMachine, and returns the corresponding virtualMachine object, and an error if there is any.
func (c *FakeVirtualMachines) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(virtualmachinesResource, name), &v1alpha1.VirtualMachine{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachine), err
}

// List takes label and field selectors, and returns the list of VirtualMachines that match those selectors.
func (c *FakeVirtualMachines) List(opts v1.ListOptions) (result *v1alpha1.VirtualMachineList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(virtualmachinesResource, virtualmachinesKind, opts), &v1alpha1.VirtualMachineList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VirtualMachineList{}
	for _, item := range obj.(*v1alpha1.VirtualMachineList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested virtualMachines.
func (c *FakeVirtualMachines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(virtualmachinesResource, opts))
}

// Create takes the representation of a virtualMachine and creates it.  Returns the server's representation of the virtualMachine, and an error, if there is any.
func (c *FakeVirtualMachines) Create(virtualMachine *v1alpha1.VirtualMachine) (result *v1alpha1.VirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(virtualmachinesResource, virtualMachine), &v1alpha1.VirtualMachine{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachine), err
}

// Update takes the representation of a virtualMachine and updates it. Returns the server's representation of the virtualMachine, and an error, if there is any.
func (c *FakeVirtualMachines) Update(virtualMachine *v1alpha1.VirtualMachine) (result *v1alpha1.VirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(virtualmachinesResource, virtualMachine), &v1alpha1.VirtualMachine{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachine), err
}

// Delete takes name of the virtualMachine and deletes it. Returns an error if one occurs.
func (c *FakeVirtualMachines) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(virtualmachinesResource, name), &v1alpha1.VirtualMachine{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVirtualMachines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(virtualmachinesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VirtualMachineList{})
	return err
}

// Patch applies the patch and returns the patched virtualMachine.
func (c *FakeVirtualMachines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualMachine, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(virtualmachinesResource, name, data, subresources...), &v1alpha1.VirtualMachine{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VirtualMachine), err
}
