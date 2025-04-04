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

package fake

import (
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	kcpfakediscovery "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/discovery/fake"
	"github.com/kcp-dev/logicalcluster/v3"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"

	client "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned"
	clientscheme "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/scheme"

	kcpclient "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/cluster"
	wildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/typed/wildwest/v1alpha1"
	kcpwildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/cluster/typed/wildwest/v1alpha1"
	fakewildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/cluster/typed/wildwest/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *ClusterClientset {
	o := kcptesting.NewObjectTracker(clientscheme.Scheme, clientscheme.Codecs.UniversalDecoder())
	o.AddAll(objects...)

	cs := &ClusterClientset{Fake: &kcptesting.Fake{}, tracker: o}
	cs.discovery = &kcpfakediscovery.FakeDiscovery{Fake: cs.Fake, ClusterPath: logicalcluster.Wildcard}
	cs.AddReactor("*", "*", kcptesting.ObjectReaction(o))
	cs.AddWatchReactor("*", kcptesting.WatchReaction(o))

	return cs
}

var _ kcpclient.ClusterInterface = (*ClusterClientset)(nil)

// ClusterClientset contains the clients for groups.
type ClusterClientset struct {
	*kcptesting.Fake
	discovery *kcpfakediscovery.FakeDiscovery
	tracker   kcptesting.ObjectTracker
}

// Discovery retrieves the DiscoveryClient
func (c *ClusterClientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *ClusterClientset) Tracker() kcptesting.ObjectTracker {
	return c.tracker
}


// WildwestV1alpha1 retrieves the WildwestV1alpha1ClusterClient.  
func (c *ClusterClientset) WildwestV1alpha1() kcpwildwestv1alpha1.WildwestV1alpha1ClusterInterface {
	return &fakewildwestv1alpha1.WildwestV1alpha1ClusterClient{Fake: c.Fake}
}
// Cluster scopes this clientset to one cluster.
func (c *ClusterClientset) Cluster(clusterPath logicalcluster.Path) client.Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &Clientset{
		Fake: c.Fake,
		discovery: &kcpfakediscovery.FakeDiscovery{Fake: c.Fake, ClusterPath: clusterPath},
		tracker: c.tracker.Cluster(clusterPath),
		clusterPath: clusterPath,
	}
}

var _ client.Interface = (*Clientset)(nil)

// Clientset contains the clients for groups.
type Clientset struct {
	*kcptesting.Fake
	discovery *kcpfakediscovery.FakeDiscovery
	tracker   kcptesting.ScopedObjectTracker
	clusterPath logicalcluster.Path
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() kcptesting.ScopedObjectTracker {
	return c.tracker
}


// WildwestV1alpha1 retrieves the WildwestV1alpha1Client.  
func (c *Clientset) WildwestV1alpha1() wildwestv1alpha1.WildwestV1alpha1Interface {
	return &fakewildwestv1alpha1.WildwestV1alpha1Client{Fake: c.Fake, ClusterPath: c.clusterPath}
}
