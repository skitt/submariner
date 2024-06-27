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

// FakeGlobalEgressIPs implements GlobalEgressIPInterface
type FakeGlobalEgressIPs struct {
	Fake *FakeSubmarinerV1
	ns   string
}

var globalegressipsResource = v1.SchemeGroupVersion.WithResource("globalegressips")

var globalegressipsKind = v1.SchemeGroupVersion.WithKind("GlobalEgressIP")

// Get takes name of the globalEgressIP, and returns the corresponding globalEgressIP object, and an error if there is any.
func (c *FakeGlobalEgressIPs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GlobalEgressIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(globalegressipsResource, c.ns, name), &v1.GlobalEgressIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GlobalEgressIP), err
}

// List takes label and field selectors, and returns the list of GlobalEgressIPs that match those selectors.
func (c *FakeGlobalEgressIPs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GlobalEgressIPList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(globalegressipsResource, globalegressipsKind, c.ns, opts), &v1.GlobalEgressIPList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.GlobalEgressIPList{ListMeta: obj.(*v1.GlobalEgressIPList).ListMeta}
	for _, item := range obj.(*v1.GlobalEgressIPList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested globalEgressIPs.
func (c *FakeGlobalEgressIPs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(globalegressipsResource, c.ns, opts))

}

// Create takes the representation of a globalEgressIP and creates it.  Returns the server's representation of the globalEgressIP, and an error, if there is any.
func (c *FakeGlobalEgressIPs) Create(ctx context.Context, globalEgressIP *v1.GlobalEgressIP, opts metav1.CreateOptions) (result *v1.GlobalEgressIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(globalegressipsResource, c.ns, globalEgressIP), &v1.GlobalEgressIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GlobalEgressIP), err
}

// Update takes the representation of a globalEgressIP and updates it. Returns the server's representation of the globalEgressIP, and an error, if there is any.
func (c *FakeGlobalEgressIPs) Update(ctx context.Context, globalEgressIP *v1.GlobalEgressIP, opts metav1.UpdateOptions) (result *v1.GlobalEgressIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(globalegressipsResource, c.ns, globalEgressIP), &v1.GlobalEgressIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GlobalEgressIP), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeGlobalEgressIPs) UpdateStatus(ctx context.Context, globalEgressIP *v1.GlobalEgressIP, opts metav1.UpdateOptions) (*v1.GlobalEgressIP, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(globalegressipsResource, "status", c.ns, globalEgressIP), &v1.GlobalEgressIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GlobalEgressIP), err
}

// Delete takes name of the globalEgressIP and deletes it. Returns an error if one occurs.
func (c *FakeGlobalEgressIPs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(globalegressipsResource, c.ns, name, opts), &v1.GlobalEgressIP{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGlobalEgressIPs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(globalegressipsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.GlobalEgressIPList{})
	return err
}

// Patch applies the patch and returns the patched globalEgressIP.
func (c *FakeGlobalEgressIPs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GlobalEgressIP, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(globalegressipsResource, c.ns, name, pt, data, subresources...), &v1.GlobalEgressIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GlobalEgressIP), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied globalEgressIP.
func (c *FakeGlobalEgressIPs) Apply(ctx context.Context, globalEgressIP *submarineriov1.GlobalEgressIPApplyConfiguration, opts metav1.ApplyOptions) (result *v1.GlobalEgressIP, err error) {
	if globalEgressIP == nil {
		return nil, fmt.Errorf("globalEgressIP provided to Apply must not be nil")
	}
	data, err := json.Marshal(globalEgressIP)
	if err != nil {
		return nil, err
	}
	name := globalEgressIP.Name
	if name == nil {
		return nil, fmt.Errorf("globalEgressIP.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(globalegressipsResource, c.ns, *name, types.ApplyPatchType, data), &v1.GlobalEgressIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GlobalEgressIP), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeGlobalEgressIPs) ApplyStatus(ctx context.Context, globalEgressIP *submarineriov1.GlobalEgressIPApplyConfiguration, opts metav1.ApplyOptions) (result *v1.GlobalEgressIP, err error) {
	if globalEgressIP == nil {
		return nil, fmt.Errorf("globalEgressIP provided to Apply must not be nil")
	}
	data, err := json.Marshal(globalEgressIP)
	if err != nil {
		return nil, err
	}
	name := globalEgressIP.Name
	if name == nil {
		return nil, fmt.Errorf("globalEgressIP.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(globalegressipsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1.GlobalEgressIP{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.GlobalEgressIP), err
}
