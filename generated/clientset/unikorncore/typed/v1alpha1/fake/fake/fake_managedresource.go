/*
Copyright 2022-2024 EscherCloud.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	fake "github.com/eschercloudai/unikorn-core/pkg/apis/unikorn/v1alpha1/fake"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeManagedResources implements ManagedResourceInterface
type FakeManagedResources struct {
	Fake *FakeUnikornFake
	ns   string
}

var managedresourcesResource = fake.SchemeGroupVersion.WithResource("managedresources")

var managedresourcesKind = fake.SchemeGroupVersion.WithKind("ManagedResource")

// Get takes name of the managedResource, and returns the corresponding managedResource object, and an error if there is any.
func (c *FakeManagedResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *fake.ManagedResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(managedresourcesResource, c.ns, name), &fake.ManagedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fake.ManagedResource), err
}

// List takes label and field selectors, and returns the list of ManagedResources that match those selectors.
func (c *FakeManagedResources) List(ctx context.Context, opts v1.ListOptions) (result *fake.ManagedResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(managedresourcesResource, managedresourcesKind, c.ns, opts), &fake.ManagedResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &fake.ManagedResourceList{ListMeta: obj.(*fake.ManagedResourceList).ListMeta}
	for _, item := range obj.(*fake.ManagedResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested managedResources.
func (c *FakeManagedResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(managedresourcesResource, c.ns, opts))

}

// Create takes the representation of a managedResource and creates it.  Returns the server's representation of the managedResource, and an error, if there is any.
func (c *FakeManagedResources) Create(ctx context.Context, managedResource *fake.ManagedResource, opts v1.CreateOptions) (result *fake.ManagedResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(managedresourcesResource, c.ns, managedResource), &fake.ManagedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fake.ManagedResource), err
}

// Update takes the representation of a managedResource and updates it. Returns the server's representation of the managedResource, and an error, if there is any.
func (c *FakeManagedResources) Update(ctx context.Context, managedResource *fake.ManagedResource, opts v1.UpdateOptions) (result *fake.ManagedResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(managedresourcesResource, c.ns, managedResource), &fake.ManagedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fake.ManagedResource), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManagedResources) UpdateStatus(ctx context.Context, managedResource *fake.ManagedResource, opts v1.UpdateOptions) (*fake.ManagedResource, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(managedresourcesResource, "status", c.ns, managedResource), &fake.ManagedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fake.ManagedResource), err
}

// Delete takes name of the managedResource and deletes it. Returns an error if one occurs.
func (c *FakeManagedResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(managedresourcesResource, c.ns, name, opts), &fake.ManagedResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManagedResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(managedresourcesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &fake.ManagedResourceList{})
	return err
}

// Patch applies the patch and returns the patched managedResource.
func (c *FakeManagedResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *fake.ManagedResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(managedresourcesResource, c.ns, name, pt, data, subresources...), &fake.ManagedResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*fake.ManagedResource), err
}
