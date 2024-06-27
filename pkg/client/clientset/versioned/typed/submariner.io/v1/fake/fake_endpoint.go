/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

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
	json "encoding/json"
	"fmt"

	v1 "github.com/submariner-io/submariner/pkg/apis/submariner.io/v1"
	submarineriov1 "github.com/submariner-io/submariner/pkg/client/applyconfiguration/submariner.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEndpoints implements EndpointInterface
type FakeEndpoints struct {
	Fake *FakeSubmarinerV1
	ns   string
}

var endpointsResource = v1.SchemeGroupVersion.WithResource("endpoints")

var endpointsKind = v1.SchemeGroupVersion.WithKind("Endpoint")

// Get takes name of the endpoint, and returns the corresponding endpoint object, and an error if there is any.
func (c *FakeEndpoints) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Endpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(endpointsResource, c.ns, name), &v1.Endpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoint), err
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *FakeEndpoints) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EndpointList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(endpointsResource, endpointsKind, c.ns, opts), &v1.EndpointList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.EndpointList{ListMeta: obj.(*v1.EndpointList).ListMeta}
	for _, item := range obj.(*v1.EndpointList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *FakeEndpoints) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(endpointsResource, c.ns, opts))

}

// Create takes the representation of a endpoint and creates it.  Returns the server's representation of the endpoint, and an error, if there is any.
func (c *FakeEndpoints) Create(ctx context.Context, endpoint *v1.Endpoint, opts metav1.CreateOptions) (result *v1.Endpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(endpointsResource, c.ns, endpoint), &v1.Endpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoint), err
}

// Update takes the representation of a endpoint and updates it. Returns the server's representation of the endpoint, and an error, if there is any.
func (c *FakeEndpoints) Update(ctx context.Context, endpoint *v1.Endpoint, opts metav1.UpdateOptions) (result *v1.Endpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(endpointsResource, c.ns, endpoint), &v1.Endpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoint), err
}

// Delete takes name of the endpoint and deletes it. Returns an error if one occurs.
func (c *FakeEndpoints) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(endpointsResource, c.ns, name, opts), &v1.Endpoint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEndpoints) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(endpointsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.EndpointList{})
	return err
}

// Patch applies the patch and returns the patched endpoint.
func (c *FakeEndpoints) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Endpoint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, name, pt, data, subresources...), &v1.Endpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoint), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied endpoint.
func (c *FakeEndpoints) Apply(ctx context.Context, endpoint *submarineriov1.EndpointApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Endpoint, err error) {
	if endpoint == nil {
		return nil, fmt.Errorf("endpoint provided to Apply must not be nil")
	}
	data, err := json.Marshal(endpoint)
	if err != nil {
		return nil, err
	}
	name := endpoint.Name
	if name == nil {
		return nil, fmt.Errorf("endpoint.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(endpointsResource, c.ns, *name, types.ApplyPatchType, data), &v1.Endpoint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Endpoint), err
}
