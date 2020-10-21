/*
Copyright 2020 The OpenEBS Authors

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

package v1

import (
	v1 "github.com/openebs/api/pkg/apis/cstor/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CStorVolumeReplicaLister helps list CStorVolumeReplicas.
// All objects returned here must be treated as read-only.
type CStorVolumeReplicaLister interface {
	// List lists all CStorVolumeReplicas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CStorVolumeReplica, err error)
	// CStorVolumeReplicas returns an object that can list and get CStorVolumeReplicas.
	CStorVolumeReplicas(namespace string) CStorVolumeReplicaNamespaceLister
	CStorVolumeReplicaListerExpansion
}

// cStorVolumeReplicaLister implements the CStorVolumeReplicaLister interface.
type cStorVolumeReplicaLister struct {
	indexer cache.Indexer
}

// NewCStorVolumeReplicaLister returns a new CStorVolumeReplicaLister.
func NewCStorVolumeReplicaLister(indexer cache.Indexer) CStorVolumeReplicaLister {
	return &cStorVolumeReplicaLister{indexer: indexer}
}

// List lists all CStorVolumeReplicas in the indexer.
func (s *cStorVolumeReplicaLister) List(selector labels.Selector) (ret []*v1.CStorVolumeReplica, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CStorVolumeReplica))
	})
	return ret, err
}

// CStorVolumeReplicas returns an object that can list and get CStorVolumeReplicas.
func (s *cStorVolumeReplicaLister) CStorVolumeReplicas(namespace string) CStorVolumeReplicaNamespaceLister {
	return cStorVolumeReplicaNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CStorVolumeReplicaNamespaceLister helps list and get CStorVolumeReplicas.
// All objects returned here must be treated as read-only.
type CStorVolumeReplicaNamespaceLister interface {
	// List lists all CStorVolumeReplicas in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CStorVolumeReplica, err error)
	// Get retrieves the CStorVolumeReplica from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CStorVolumeReplica, error)
	CStorVolumeReplicaNamespaceListerExpansion
}

// cStorVolumeReplicaNamespaceLister implements the CStorVolumeReplicaNamespaceLister
// interface.
type cStorVolumeReplicaNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CStorVolumeReplicas in the indexer for a given namespace.
func (s cStorVolumeReplicaNamespaceLister) List(selector labels.Selector) (ret []*v1.CStorVolumeReplica, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CStorVolumeReplica))
	})
	return ret, err
}

// Get retrieves the CStorVolumeReplica from the indexer for a given namespace and name.
func (s cStorVolumeReplicaNamespaceLister) Get(name string) (*v1.CStorVolumeReplica, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cstorvolumereplica"), name)
	}
	return obj.(*v1.CStorVolumeReplica), nil
}
