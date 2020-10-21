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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	openebsiov1alpha1 "github.com/openebs/api/pkg/apis/openebs.io/v1alpha1"
	versioned "github.com/openebs/api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openebs/api/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openebs/api/pkg/client/listers/openebs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// BlockDeviceClaimInformer provides access to a shared informer and lister for
// BlockDeviceClaims.
type BlockDeviceClaimInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.BlockDeviceClaimLister
}

type blockDeviceClaimInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewBlockDeviceClaimInformer constructs a new informer for BlockDeviceClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewBlockDeviceClaimInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredBlockDeviceClaimInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredBlockDeviceClaimInformer constructs a new informer for BlockDeviceClaim type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredBlockDeviceClaimInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().BlockDeviceClaims(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().BlockDeviceClaims(namespace).Watch(context.TODO(), options)
			},
		},
		&openebsiov1alpha1.BlockDeviceClaim{},
		resyncPeriod,
		indexers,
	)
}

func (f *blockDeviceClaimInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredBlockDeviceClaimInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *blockDeviceClaimInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&openebsiov1alpha1.BlockDeviceClaim{}, f.defaultInformer)
}

func (f *blockDeviceClaimInformer) Lister() v1alpha1.BlockDeviceClaimLister {
	return v1alpha1.NewBlockDeviceClaimLister(f.Informer().GetIndexer())
}
