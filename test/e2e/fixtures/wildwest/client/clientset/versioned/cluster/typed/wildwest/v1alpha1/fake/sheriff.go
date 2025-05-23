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

// Code generated by cluster-client-gen. DO NOT EDIT.

package fake

import (
	kcpgentype "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/gentype"
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v3"

	wildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/apis/wildwest/v1alpha1"
	kcpv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/applyconfiguration/wildwest/v1alpha1"
	typedkcpwildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/cluster/typed/wildwest/v1alpha1"
	typedwildwestv1alpha1 "github.com/kcp-dev/kcp/test/e2e/fixtures/wildwest/client/clientset/versioned/typed/wildwest/v1alpha1"
)

// sheriffClusterClient implements SheriffClusterInterface
type sheriffClusterClient struct {
	*kcpgentype.FakeClusterClientWithList[*wildwestv1alpha1.Sheriff, *wildwestv1alpha1.SheriffList]
	Fake *kcptesting.Fake
}

func newFakeSheriffClusterClient(fake *WildwestV1alpha1ClusterClient) typedkcpwildwestv1alpha1.SheriffClusterInterface {
	return &sheriffClusterClient{
		kcpgentype.NewFakeClusterClientWithList[*wildwestv1alpha1.Sheriff, *wildwestv1alpha1.SheriffList](
			fake.Fake,
			wildwestv1alpha1.SchemeGroupVersion.WithResource("sherifves"),
			wildwestv1alpha1.SchemeGroupVersion.WithKind("Sheriff"),
			func() *wildwestv1alpha1.Sheriff { return &wildwestv1alpha1.Sheriff{} },
			func() *wildwestv1alpha1.SheriffList { return &wildwestv1alpha1.SheriffList{} },
			func(dst, src *wildwestv1alpha1.SheriffList) { dst.ListMeta = src.ListMeta },
			func(list *wildwestv1alpha1.SheriffList) []*wildwestv1alpha1.Sheriff {
				return kcpgentype.ToPointerSlice(list.Items)
			},
			func(list *wildwestv1alpha1.SheriffList, items []*wildwestv1alpha1.Sheriff) {
				list.Items = kcpgentype.FromPointerSlice(items)
			},
		),
		fake.Fake,
	}
}

func (c *sheriffClusterClient) Cluster(cluster logicalcluster.Path) typedwildwestv1alpha1.SheriffInterface {
	return newFakeSheriffClient(c.Fake, cluster)
}

// sheriffScopedClient implements SheriffInterface
type sheriffScopedClient struct {
	*kcpgentype.FakeClientWithListAndApply[*wildwestv1alpha1.Sheriff, *wildwestv1alpha1.SheriffList, *kcpv1alpha1.SheriffApplyConfiguration]
	Fake        *kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func newFakeSheriffClient(fake *kcptesting.Fake, clusterPath logicalcluster.Path) typedwildwestv1alpha1.SheriffInterface {
	return &sheriffScopedClient{
		kcpgentype.NewFakeClientWithListAndApply[*wildwestv1alpha1.Sheriff, *wildwestv1alpha1.SheriffList, *kcpv1alpha1.SheriffApplyConfiguration](
			fake,
			clusterPath,
			"",
			wildwestv1alpha1.SchemeGroupVersion.WithResource("sherifves"),
			wildwestv1alpha1.SchemeGroupVersion.WithKind("Sheriff"),
			func() *wildwestv1alpha1.Sheriff { return &wildwestv1alpha1.Sheriff{} },
			func() *wildwestv1alpha1.SheriffList { return &wildwestv1alpha1.SheriffList{} },
			func(dst, src *wildwestv1alpha1.SheriffList) { dst.ListMeta = src.ListMeta },
			func(list *wildwestv1alpha1.SheriffList) []*wildwestv1alpha1.Sheriff {
				return kcpgentype.ToPointerSlice(list.Items)
			},
			func(list *wildwestv1alpha1.SheriffList, items []*wildwestv1alpha1.Sheriff) {
				list.Items = kcpgentype.FromPointerSlice(items)
			},
		),
		fake,
		clusterPath,
	}
}
