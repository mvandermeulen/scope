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

// CStorCompletedBackupLister helps list CStorCompletedBackups.
// All objects returned here must be treated as read-only.
type CStorCompletedBackupLister interface {
	// List lists all CStorCompletedBackups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CStorCompletedBackup, err error)
	// CStorCompletedBackups returns an object that can list and get CStorCompletedBackups.
	CStorCompletedBackups(namespace string) CStorCompletedBackupNamespaceLister
	CStorCompletedBackupListerExpansion
}

// cStorCompletedBackupLister implements the CStorCompletedBackupLister interface.
type cStorCompletedBackupLister struct {
	indexer cache.Indexer
}

// NewCStorCompletedBackupLister returns a new CStorCompletedBackupLister.
func NewCStorCompletedBackupLister(indexer cache.Indexer) CStorCompletedBackupLister {
	return &cStorCompletedBackupLister{indexer: indexer}
}

// List lists all CStorCompletedBackups in the indexer.
func (s *cStorCompletedBackupLister) List(selector labels.Selector) (ret []*v1.CStorCompletedBackup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CStorCompletedBackup))
	})
	return ret, err
}

// CStorCompletedBackups returns an object that can list and get CStorCompletedBackups.
func (s *cStorCompletedBackupLister) CStorCompletedBackups(namespace string) CStorCompletedBackupNamespaceLister {
	return cStorCompletedBackupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CStorCompletedBackupNamespaceLister helps list and get CStorCompletedBackups.
// All objects returned here must be treated as read-only.
type CStorCompletedBackupNamespaceLister interface {
	// List lists all CStorCompletedBackups in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CStorCompletedBackup, err error)
	// Get retrieves the CStorCompletedBackup from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CStorCompletedBackup, error)
	CStorCompletedBackupNamespaceListerExpansion
}

// cStorCompletedBackupNamespaceLister implements the CStorCompletedBackupNamespaceLister
// interface.
type cStorCompletedBackupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CStorCompletedBackups in the indexer for a given namespace.
func (s cStorCompletedBackupNamespaceLister) List(selector labels.Selector) (ret []*v1.CStorCompletedBackup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CStorCompletedBackup))
	})
	return ret, err
}

// Get retrieves the CStorCompletedBackup from the indexer for a given namespace and name.
func (s cStorCompletedBackupNamespaceLister) Get(name string) (*v1.CStorCompletedBackup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("cstorcompletedbackup"), name)
	}
	return obj.(*v1.CStorCompletedBackup), nil
}
