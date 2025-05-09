/*
Copyright The KCP Authors.

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


// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	wildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/apis/wildwest/v1alpha1"
	wildwestv1alpha1listers "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/listers/wildwest/v1alpha1"
	clientset "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/cluster"
	scopedclientset "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned"
	"github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/informers/externalversions/internalinterfaces"
)

// CowboyClusterInformer provides access to a shared informer and lister for
// Cowboys.
type CowboyClusterInformer interface {
	Cluster(logicalcluster.Name) CowboyInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() wildwestv1alpha1listers.CowboyClusterLister
}

type cowboyClusterInformer struct {
	factory internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCowboyClusterInformer constructs a new informer for Cowboy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCowboyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredCowboyClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCowboyClusterInformer constructs a new informer for Cowboy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCowboyClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WildwestV1alpha1().Cowboys().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WildwestV1alpha1().Cowboys().Watch(context.TODO(), options)
			},
		},
		&wildwestv1alpha1.Cowboy{},
		resyncPeriod,
		indexers,
	)
}

func (f *cowboyClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredCowboyClusterInformer(client, resyncPeriod, cache.Indexers{
			kcpcache.ClusterIndexName: kcpcache.ClusterIndexFunc,
			kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc,}, 
		f.tweakListOptions,
	)
}

func (f *cowboyClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&wildwestv1alpha1.Cowboy{}, f.defaultInformer)
}

func (f *cowboyClusterInformer) Lister() wildwestv1alpha1listers.CowboyClusterLister {
	return wildwestv1alpha1listers.NewCowboyClusterLister(f.Informer().GetIndexer())
}


// CowboyInformer provides access to a shared informer and lister for
// Cowboys.
type CowboyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() wildwestv1alpha1listers.CowboyLister
}
func (f *cowboyClusterInformer) Cluster(clusterName logicalcluster.Name) CowboyInformer {
	return &cowboyInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type cowboyInformer struct {
	informer cache.SharedIndexInformer
	lister wildwestv1alpha1listers.CowboyLister
}

func (f *cowboyInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *cowboyInformer) Lister() wildwestv1alpha1listers.CowboyLister {
	return f.lister
}

type cowboyScopedInformer struct {
	factory internalinterfaces.SharedScopedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace string}

func (f *cowboyScopedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&wildwestv1alpha1.Cowboy{}, f.defaultInformer)
}

func (f *cowboyScopedInformer) Lister() wildwestv1alpha1listers.CowboyLister {
	return wildwestv1alpha1listers.NewCowboyLister(f.Informer().GetIndexer())
}

// NewCowboyInformer constructs a new informer for Cowboy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCowboyInformer(client scopedclientset.Interface, resyncPeriod time.Duration, namespace string,indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCowboyInformer(client, resyncPeriod,  namespace,indexers, nil)
}

// NewFilteredCowboyInformer constructs a new informer for Cowboy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCowboyInformer(client scopedclientset.Interface, resyncPeriod time.Duration,  namespace string,indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WildwestV1alpha1().Cowboys(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WildwestV1alpha1().Cowboys(namespace).Watch(context.TODO(), options)
			},
		},
		&wildwestv1alpha1.Cowboy{},
		resyncPeriod,
		indexers,
	)
}

func (f *cowboyScopedInformer) defaultInformer(client scopedclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCowboyInformer(client, resyncPeriod, f.namespace,cache.Indexers{ 
		cache.NamespaceIndex: cache.MetaNamespaceIndexFunc,
	}, f.tweakListOptions)
}

