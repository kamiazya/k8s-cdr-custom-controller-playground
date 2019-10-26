/*
Copyright The Kubernetes Authors.

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
	hogev1 "github.com/kamiazya/k8s-cdr-custom-controller-playground/pkg/apis/hoge/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHoges implements HogeInterface
type FakeHoges struct {
	Fake *FakeSampleV1
	ns   string
}

var hogesResource = schema.GroupVersionResource{Group: "sample.kamiazya.tech", Version: "v1", Resource: "hoges"}

var hogesKind = schema.GroupVersionKind{Group: "sample.kamiazya.tech", Version: "v1", Kind: "Hoge"}

// Get takes name of the hoge, and returns the corresponding hoge object, and an error if there is any.
func (c *FakeHoges) Get(name string, options v1.GetOptions) (result *hogev1.Hoge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hogesResource, c.ns, name), &hogev1.Hoge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hogev1.Hoge), err
}

// List takes label and field selectors, and returns the list of Hoges that match those selectors.
func (c *FakeHoges) List(opts v1.ListOptions) (result *hogev1.HogeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hogesResource, hogesKind, c.ns, opts), &hogev1.HogeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &hogev1.HogeList{ListMeta: obj.(*hogev1.HogeList).ListMeta}
	for _, item := range obj.(*hogev1.HogeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hoges.
func (c *FakeHoges) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hogesResource, c.ns, opts))

}

// Create takes the representation of a hoge and creates it.  Returns the server's representation of the hoge, and an error, if there is any.
func (c *FakeHoges) Create(hoge *hogev1.Hoge) (result *hogev1.Hoge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hogesResource, c.ns, hoge), &hogev1.Hoge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hogev1.Hoge), err
}

// Update takes the representation of a hoge and updates it. Returns the server's representation of the hoge, and an error, if there is any.
func (c *FakeHoges) Update(hoge *hogev1.Hoge) (result *hogev1.Hoge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hogesResource, c.ns, hoge), &hogev1.Hoge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hogev1.Hoge), err
}

// Delete takes name of the hoge and deletes it. Returns an error if one occurs.
func (c *FakeHoges) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(hogesResource, c.ns, name), &hogev1.Hoge{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHoges) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hogesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &hogev1.HogeList{})
	return err
}

// Patch applies the patch and returns the patched hoge.
func (c *FakeHoges) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *hogev1.Hoge, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hogesResource, c.ns, name, pt, data, subresources...), &hogev1.Hoge{})

	if obj == nil {
		return nil, err
	}
	return obj.(*hogev1.Hoge), err
}
