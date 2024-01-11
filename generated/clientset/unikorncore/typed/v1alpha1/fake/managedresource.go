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
	"time"

	scheme "github.com/eschercloudai/unikorn-core/generated/clientset/unikorncore/scheme"
	fake "github.com/eschercloudai/unikorn-core/pkg/apis/unikorn/v1alpha1/fake"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ManagedResourcesGetter has a method to return a ManagedResourceInterface.
// A group's client should implement this interface.
type ManagedResourcesGetter interface {
	ManagedResources(namespace string) ManagedResourceInterface
}

// ManagedResourceInterface has methods to work with ManagedResource resources.
type ManagedResourceInterface interface {
	Create(ctx context.Context, managedResource *fake.ManagedResource, opts v1.CreateOptions) (*fake.ManagedResource, error)
	Update(ctx context.Context, managedResource *fake.ManagedResource, opts v1.UpdateOptions) (*fake.ManagedResource, error)
	UpdateStatus(ctx context.Context, managedResource *fake.ManagedResource, opts v1.UpdateOptions) (*fake.ManagedResource, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*fake.ManagedResource, error)
	List(ctx context.Context, opts v1.ListOptions) (*fake.ManagedResourceList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *fake.ManagedResource, err error)
	ManagedResourceExpansion
}

// managedResources implements ManagedResourceInterface
type managedResources struct {
	client rest.Interface
	ns     string
}

// newManagedResources returns a ManagedResources
func newManagedResources(c *UnikornFakeClient, namespace string) *managedResources {
	return &managedResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the managedResource, and returns the corresponding managedResource object, and an error if there is any.
func (c *managedResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *fake.ManagedResource, err error) {
	result = &fake.ManagedResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managedresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ManagedResources that match those selectors.
func (c *managedResources) List(ctx context.Context, opts v1.ListOptions) (result *fake.ManagedResourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &fake.ManagedResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("managedresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested managedResources.
func (c *managedResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("managedresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a managedResource and creates it.  Returns the server's representation of the managedResource, and an error, if there is any.
func (c *managedResources) Create(ctx context.Context, managedResource *fake.ManagedResource, opts v1.CreateOptions) (result *fake.ManagedResource, err error) {
	result = &fake.ManagedResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("managedresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managedResource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a managedResource and updates it. Returns the server's representation of the managedResource, and an error, if there is any.
func (c *managedResources) Update(ctx context.Context, managedResource *fake.ManagedResource, opts v1.UpdateOptions) (result *fake.ManagedResource, err error) {
	result = &fake.ManagedResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managedresources").
		Name(managedResource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managedResource).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *managedResources) UpdateStatus(ctx context.Context, managedResource *fake.ManagedResource, opts v1.UpdateOptions) (result *fake.ManagedResource, err error) {
	result = &fake.ManagedResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("managedresources").
		Name(managedResource.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(managedResource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the managedResource and deletes it. Returns an error if one occurs.
func (c *managedResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managedresources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *managedResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("managedresources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched managedResource.
func (c *managedResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *fake.ManagedResource, err error) {
	result = &fake.ManagedResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("managedresources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
