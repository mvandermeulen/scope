/*
Copyright 2018 The OpenEBS Authors

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

	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCStorPoolClusters implements CStorPoolClusterInterface
type FakeCStorPoolClusters struct {
	Fake *FakeOpenebsV1alpha1
	ns   string
}

var cstorpoolclustersResource = schema.GroupVersionResource{Group: "openebs.io", Version: "v1alpha1", Resource: "cstorpoolclusters"}

var cstorpoolclustersKind = schema.GroupVersionKind{Group: "openebs.io", Version: "v1alpha1", Kind: "CStorPoolCluster"}

// Get takes name of the cStorPoolCluster, and returns the corresponding cStorPoolCluster object, and an error if there is any.
func (c *FakeCStorPoolClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CStorPoolCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(cstorpoolclustersResource, c.ns, name), &v1alpha1.CStorPoolCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorPoolCluster), err
}

// List takes label and field selectors, and returns the list of CStorPoolClusters that match those selectors.
func (c *FakeCStorPoolClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CStorPoolClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(cstorpoolclustersResource, cstorpoolclustersKind, c.ns, opts), &v1alpha1.CStorPoolClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.CStorPoolClusterList{ListMeta: obj.(*v1alpha1.CStorPoolClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.CStorPoolClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cStorPoolClusters.
func (c *FakeCStorPoolClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(cstorpoolclustersResource, c.ns, opts))

}

// Create takes the representation of a cStorPoolCluster and creates it.  Returns the server's representation of the cStorPoolCluster, and an error, if there is any.
func (c *FakeCStorPoolClusters) Create(ctx context.Context, cStorPoolCluster *v1alpha1.CStorPoolCluster, opts v1.CreateOptions) (result *v1alpha1.CStorPoolCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(cstorpoolclustersResource, c.ns, cStorPoolCluster), &v1alpha1.CStorPoolCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorPoolCluster), err
}

// Update takes the representation of a cStorPoolCluster and updates it. Returns the server's representation of the cStorPoolCluster, and an error, if there is any.
func (c *FakeCStorPoolClusters) Update(ctx context.Context, cStorPoolCluster *v1alpha1.CStorPoolCluster, opts v1.UpdateOptions) (result *v1alpha1.CStorPoolCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(cstorpoolclustersResource, c.ns, cStorPoolCluster), &v1alpha1.CStorPoolCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorPoolCluster), err
}

// Delete takes name of the cStorPoolCluster and deletes it. Returns an error if one occurs.
func (c *FakeCStorPoolClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(cstorpoolclustersResource, c.ns, name), &v1alpha1.CStorPoolCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCStorPoolClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(cstorpoolclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.CStorPoolClusterList{})
	return err
}

// Patch applies the patch and returns the patched cStorPoolCluster.
func (c *FakeCStorPoolClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CStorPoolCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(cstorpoolclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.CStorPoolCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.CStorPoolCluster), err
}