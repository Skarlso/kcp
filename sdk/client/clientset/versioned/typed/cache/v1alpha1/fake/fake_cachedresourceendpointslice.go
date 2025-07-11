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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	gentype "k8s.io/client-go/gentype"

	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/cache/v1alpha1"
	cachev1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/cache/v1alpha1"
	typedcachev1alpha1 "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/cache/v1alpha1"
)

// fakeCachedResourceEndpointSlices implements CachedResourceEndpointSliceInterface
type fakeCachedResourceEndpointSlices struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.CachedResourceEndpointSlice, *v1alpha1.CachedResourceEndpointSliceList, *cachev1alpha1.CachedResourceEndpointSliceApplyConfiguration]
	Fake *FakeCacheV1alpha1
}

func newFakeCachedResourceEndpointSlices(fake *FakeCacheV1alpha1) typedcachev1alpha1.CachedResourceEndpointSliceInterface {
	return &fakeCachedResourceEndpointSlices{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.CachedResourceEndpointSlice, *v1alpha1.CachedResourceEndpointSliceList, *cachev1alpha1.CachedResourceEndpointSliceApplyConfiguration](
			fake.Fake,
			"",
			v1alpha1.SchemeGroupVersion.WithResource("cachedresourceendpointslices"),
			v1alpha1.SchemeGroupVersion.WithKind("CachedResourceEndpointSlice"),
			func() *v1alpha1.CachedResourceEndpointSlice { return &v1alpha1.CachedResourceEndpointSlice{} },
			func() *v1alpha1.CachedResourceEndpointSliceList { return &v1alpha1.CachedResourceEndpointSliceList{} },
			func(dst, src *v1alpha1.CachedResourceEndpointSliceList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.CachedResourceEndpointSliceList) []*v1alpha1.CachedResourceEndpointSlice {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.CachedResourceEndpointSliceList, items []*v1alpha1.CachedResourceEndpointSlice) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
