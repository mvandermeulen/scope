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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openebs/maya/pkg/apis/openebs.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StoragePoolClaimLister helps list StoragePoolClaims.
type StoragePoolClaimLister interface {
	// List lists all StoragePoolClaims in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.StoragePoolClaim, err error)
	// Get retrieves the StoragePoolClaim from the index for a given name.
	Get(name string) (*v1alpha1.StoragePoolClaim, error)
	StoragePoolClaimListerExpansion
}

// storagePoolClaimLister implements the StoragePoolClaimLister interface.
type storagePoolClaimLister struct {
	indexer cache.Indexer
}

// NewStoragePoolClaimLister returns a new StoragePoolClaimLister.
func NewStoragePoolClaimLister(indexer cache.Indexer) StoragePoolClaimLister {
	return &storagePoolClaimLister{indexer: indexer}
}

// List lists all StoragePoolClaims in the indexer.
func (s *storagePoolClaimLister) List(selector labels.Selector) (ret []*v1alpha1.StoragePoolClaim, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StoragePoolClaim))
	})
	return ret, err
}

// Get retrieves the StoragePoolClaim from the index for a given name.
func (s *storagePoolClaimLister) Get(name string) (*v1alpha1.StoragePoolClaim, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storagepoolclaim"), name)
	}
	return obj.(*v1alpha1.StoragePoolClaim), nil
}
