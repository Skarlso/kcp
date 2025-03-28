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
	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	topologyv1alpha1 "github.com/kcp-dev/kcp/sdk/apis/topology/v1alpha1"


	topologyv1alpha1client "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/topology/v1alpha1"
)

// PartitionSetsClusterGetter has a method to return a PartitionSetClusterInterface.
// A group's cluster client should implement this interface.
type PartitionSetsClusterGetter interface {
	PartitionSets() PartitionSetClusterInterface
}

// PartitionSetClusterInterface can operate on PartitionSets across all clusters,
// or scope down to one cluster and return a topologyv1alpha1client.PartitionSetInterface.
type PartitionSetClusterInterface interface {
	Cluster(logicalcluster.Path) topologyv1alpha1client.PartitionSetInterface
	List(ctx context.Context, opts metav1.ListOptions) (*topologyv1alpha1.PartitionSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type partitionSetsClusterInterface struct {
	clientCache kcpclient.Cache[*topologyv1alpha1client.TopologyV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *partitionSetsClusterInterface) Cluster(clusterPath logicalcluster.Path) topologyv1alpha1client.PartitionSetInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).PartitionSets()
}


// List returns the entire collection of all PartitionSets across all clusters. 
func (c *partitionSetsClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*topologyv1alpha1.PartitionSetList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PartitionSets().List(ctx, opts)
}

// Watch begins to watch all PartitionSets across all clusters.
func (c *partitionSetsClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).PartitionSets().Watch(ctx, opts)
}
