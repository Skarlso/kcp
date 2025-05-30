/*
Copyright 2023 The KCP Authors.

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

package replicatelogicalcluster

import (
	"fmt"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"

	"github.com/kcp-dev/kcp/pkg/reconciler/cache/labellogicalcluster"
	"github.com/kcp-dev/kcp/pkg/reconciler/cache/replication"
	"github.com/kcp-dev/kcp/sdk/apis/apis"
	apisv1alpha2 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha2"
	corev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/core/v1alpha1"
	kcpclientset "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/cluster"
	apisv1alpha2informers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/apis/v1alpha2"
	corev1alpha1informers "github.com/kcp-dev/kcp/sdk/client/informers/externalversions/core/v1alpha1"
)

const (
	ControllerName = "kcp-apis-replicate-logicalcluster"
)

// NewController returns a new controller for labelling LogicalClusters that should be replicated.

func NewController(
	kcpClusterClient kcpclientset.ClusterInterface,
	logicalClusterInformer corev1alpha1informers.LogicalClusterClusterInformer,
	apiExportInformer apisv1alpha2informers.APIExportClusterInformer,
) labellogicalcluster.Controller {
	logicalClusterLister := logicalClusterInformer.Lister()
	apiExportIndexer := apiExportInformer.Informer().GetIndexer()

	c := labellogicalcluster.NewController(
		ControllerName,
		apis.GroupName,
		func(cluster *corev1alpha1.LogicalCluster) bool {
			// If there are any APIExports for this logical cluster, then the LogicalCluster object should be replicated.
			keys, err := apiExportIndexer.IndexKeys(kcpcache.ClusterIndexName, kcpcache.ClusterIndexKey(logicalcluster.From(cluster)))
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("failed to list APIExports: %v", err))
				return false
			}
			return len(keys) > 0
		},
		kcpClusterClient,
		logicalClusterInformer,
	)

	// enqueue the logical cluster every time the APIExport changes
	enqueueAPIExport := func(obj interface{}) {
		if tombstone, ok := obj.(cache.DeletedFinalStateUnknown); ok {
			obj = tombstone.Obj
		}

		export, ok := obj.(*apisv1alpha2.APIExport)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("unexpected object type: %T", obj))
			return
		}

		cluster, err := logicalClusterLister.Cluster(logicalcluster.From(export)).Get(corev1alpha1.LogicalClusterName)
		if err != nil && !apierrors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("failed to get logical cluster: %v", err))
			return
		} else if apierrors.IsNotFound(err) {
			return
		}

		c.EnqueueLogicalCluster(cluster, "reason", "APIExport changed", "apiexport", export.Name)
	}

	_, _ = apiExportInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: replication.IsNoSystemClusterName,
		Handler: cache.ResourceEventHandlerFuncs{
			AddFunc: func(obj interface{}) {
				enqueueAPIExport(obj)
			},
			DeleteFunc: func(obj interface{}) {
				enqueueAPIExport(obj)
			},
		},
	})

	return c
}
