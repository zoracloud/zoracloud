// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/lyft/flinkk8soperator/pkg/apis/app/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFlinkApplications implements FlinkApplicationInterface
type FakeFlinkApplications struct {
	Fake *FakeFlinkV1beta1
	ns   string
}

var flinkapplicationsResource = schema.GroupVersionResource{Group: "flink.k8s.io", Version: "v1beta1", Resource: "flinkapplications"}

var flinkapplicationsKind = schema.GroupVersionKind{Group: "flink.k8s.io", Version: "v1beta1", Kind: "FlinkApplication"}

// Get takes name of the flinkApplication, and returns the corresponding flinkApplication object, and an error if there is any.
func (c *FakeFlinkApplications) Get(name string, options v1.GetOptions) (result *v1beta1.FlinkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(flinkapplicationsResource, c.ns, name), &v1beta1.FlinkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FlinkApplication), err
}

// List takes label and field selectors, and returns the list of FlinkApplications that match those selectors.
func (c *FakeFlinkApplications) List(opts v1.ListOptions) (result *v1beta1.FlinkApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(flinkapplicationsResource, flinkapplicationsKind, c.ns, opts), &v1beta1.FlinkApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FlinkApplicationList{ListMeta: obj.(*v1beta1.FlinkApplicationList).ListMeta}
	for _, item := range obj.(*v1beta1.FlinkApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested flinkApplications.
func (c *FakeFlinkApplications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(flinkapplicationsResource, c.ns, opts))

}

// Create takes the representation of a flinkApplication and creates it.  Returns the server's representation of the flinkApplication, and an error, if there is any.
func (c *FakeFlinkApplications) Create(flinkApplication *v1beta1.FlinkApplication) (result *v1beta1.FlinkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(flinkapplicationsResource, c.ns, flinkApplication), &v1beta1.FlinkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FlinkApplication), err
}

// Update takes the representation of a flinkApplication and updates it. Returns the server's representation of the flinkApplication, and an error, if there is any.
func (c *FakeFlinkApplications) Update(flinkApplication *v1beta1.FlinkApplication) (result *v1beta1.FlinkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(flinkapplicationsResource, c.ns, flinkApplication), &v1beta1.FlinkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FlinkApplication), err
}

// Delete takes name of the flinkApplication and deletes it. Returns an error if one occurs.
func (c *FakeFlinkApplications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(flinkapplicationsResource, c.ns, name), &v1beta1.FlinkApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFlinkApplications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(flinkapplicationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.FlinkApplicationList{})
	return err
}

// Patch applies the patch and returns the patched flinkApplication.
func (c *FakeFlinkApplications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FlinkApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(flinkapplicationsResource, c.ns, name, pt, data, subresources...), &v1beta1.FlinkApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FlinkApplication), err
}
