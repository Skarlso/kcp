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
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"	
	"github.com/kcp-dev/logicalcluster/v3"
	
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/api/errors"

	apisv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"
	)

// APIExportClusterLister can list APIExports across all workspaces, or scope down to a APIExportLister for one workspace.
// All objects returned here must be treated as read-only.
type APIExportClusterLister interface {
	// List lists all APIExports in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apisv1alpha1.APIExport, err error)
	// Cluster returns a lister that can list and get APIExports in one workspace.
Cluster(clusterName logicalcluster.Name) APIExportLister
APIExportClusterListerExpansion
}

type aPIExportClusterLister struct {
	indexer cache.Indexer
}

// NewAPIExportClusterLister returns a new APIExportClusterLister.
// We assume that the indexer:
// - is fed by a cross-workspace LIST+WATCH
// - uses kcpcache.MetaClusterNamespaceKeyFunc as the key function
// - has the kcpcache.ClusterIndex as an index
func NewAPIExportClusterLister(indexer cache.Indexer) *aPIExportClusterLister {
	return &aPIExportClusterLister{indexer: indexer}
}

// List lists all APIExports in the indexer across all workspaces.
func (s *aPIExportClusterLister) List(selector labels.Selector) (ret []*apisv1alpha1.APIExport, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*apisv1alpha1.APIExport))
	})
	return ret, err
}

// Cluster scopes the lister to one workspace, allowing users to list and get APIExports.
func (s *aPIExportClusterLister) Cluster(clusterName logicalcluster.Name) APIExportLister {
return &aPIExportLister{indexer: s.indexer, clusterName: clusterName}
}

// APIExportLister can list all APIExports, or get one in particular.
// All objects returned here must be treated as read-only.
type APIExportLister interface {
	// List lists all APIExports in the workspace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*apisv1alpha1.APIExport, err error)
// Get retrieves the APIExport from the indexer for a given workspace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*apisv1alpha1.APIExport, error)
APIExportListerExpansion
}
// aPIExportLister can list all APIExports inside a workspace.
type aPIExportLister struct {
	indexer cache.Indexer
	clusterName logicalcluster.Name
}

// List lists all APIExports in the indexer for a workspace.
func (s *aPIExportLister) List(selector labels.Selector) (ret []*apisv1alpha1.APIExport, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*apisv1alpha1.APIExport))
	})
	return ret, err
}

// Get retrieves the APIExport from the indexer for a given workspace and name.
func (s *aPIExportLister) Get(name string) (*apisv1alpha1.APIExport, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(apisv1alpha1.Resource("apiexports"), name)
	}
	return obj.(*apisv1alpha1.APIExport), nil
}
// NewAPIExportLister returns a new APIExportLister.
// We assume that the indexer:
// - is fed by a workspace-scoped LIST+WATCH
// - uses cache.MetaNamespaceKeyFunc as the key function
func NewAPIExportLister(indexer cache.Indexer) *aPIExportScopedLister {
	return &aPIExportScopedLister{indexer: indexer}
}

// aPIExportScopedLister can list all APIExports inside a workspace.
type aPIExportScopedLister struct {
	indexer cache.Indexer
}

// List lists all APIExports in the indexer for a workspace.
func (s *aPIExportScopedLister) List(selector labels.Selector) (ret []*apisv1alpha1.APIExport, err error) {
	err = cache.ListAll(s.indexer, selector, func(i interface{}) {
		ret = append(ret, i.(*apisv1alpha1.APIExport))
	})
	return ret, err
}

// Get retrieves the APIExport from the indexer for a given workspace and name.
func (s *aPIExportScopedLister) Get(name string) (*apisv1alpha1.APIExport, error) {
	key := name
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(apisv1alpha1.Resource("apiexports"), name)
	}
	return obj.(*apisv1alpha1.APIExport), nil
}
