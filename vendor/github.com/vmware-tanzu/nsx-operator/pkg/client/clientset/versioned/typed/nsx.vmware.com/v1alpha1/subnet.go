/* Copyright © 2023 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: Apache-2.0 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/vmware-tanzu/nsx-operator/pkg/apis/nsx.vmware.com/v1alpha1"
	scheme "github.com/vmware-tanzu/nsx-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SubnetsGetter has a method to return a SubnetInterface.
// A group's client should implement this interface.
type SubnetsGetter interface {
	Subnets(namespace string) SubnetInterface
}

// SubnetInterface has methods to work with Subnet resources.
type SubnetInterface interface {
	Create(ctx context.Context, subnet *v1alpha1.Subnet, opts v1.CreateOptions) (*v1alpha1.Subnet, error)
	Update(ctx context.Context, subnet *v1alpha1.Subnet, opts v1.UpdateOptions) (*v1alpha1.Subnet, error)
	UpdateStatus(ctx context.Context, subnet *v1alpha1.Subnet, opts v1.UpdateOptions) (*v1alpha1.Subnet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Subnet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SubnetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Subnet, err error)
	SubnetExpansion
}

// subnets implements SubnetInterface
type subnets struct {
	client rest.Interface
	ns     string
}

// newSubnets returns a Subnets
func newSubnets(c *NsxV1alpha1Client, namespace string) *subnets {
	return &subnets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the subnet, and returns the corresponding subnet object, and an error if there is any.
func (c *subnets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Subnets that match those selectors.
func (c *subnets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubnetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SubnetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("subnets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested subnets.
func (c *subnets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("subnets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a subnet and creates it.  Returns the server's representation of the subnet, and an error, if there is any.
func (c *subnets) Create(ctx context.Context, subnet *v1alpha1.Subnet, opts v1.CreateOptions) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("subnets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subnet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a subnet and updates it. Returns the server's representation of the subnet, and an error, if there is any.
func (c *subnets) Update(ctx context.Context, subnet *v1alpha1.Subnet, opts v1.UpdateOptions) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnets").
		Name(subnet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subnet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *subnets) UpdateStatus(ctx context.Context, subnet *v1alpha1.Subnet, opts v1.UpdateOptions) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("subnets").
		Name(subnet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(subnet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the subnet and deletes it. Returns an error if one occurs.
func (c *subnets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *subnets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("subnets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched subnet.
func (c *subnets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Subnet, err error) {
	result = &v1alpha1.Subnet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("subnets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
