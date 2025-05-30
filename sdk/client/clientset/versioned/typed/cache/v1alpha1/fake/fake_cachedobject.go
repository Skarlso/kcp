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

// fakeCachedObjects implements CachedObjectInterface
type fakeCachedObjects struct {
	*gentype.FakeClientWithListAndApply[*v1alpha1.CachedObject, *v1alpha1.CachedObjectList, *cachev1alpha1.CachedObjectApplyConfiguration]
	Fake *FakeCacheV1alpha1
}

func newFakeCachedObjects(fake *FakeCacheV1alpha1) typedcachev1alpha1.CachedObjectInterface {
	return &fakeCachedObjects{
		gentype.NewFakeClientWithListAndApply[*v1alpha1.CachedObject, *v1alpha1.CachedObjectList, *cachev1alpha1.CachedObjectApplyConfiguration](
			fake.Fake,
			"",
			v1alpha1.SchemeGroupVersion.WithResource("cachedobjects"),
			v1alpha1.SchemeGroupVersion.WithKind("CachedObject"),
			func() *v1alpha1.CachedObject { return &v1alpha1.CachedObject{} },
			func() *v1alpha1.CachedObjectList { return &v1alpha1.CachedObjectList{} },
			func(dst, src *v1alpha1.CachedObjectList) { dst.ListMeta = src.ListMeta },
			func(list *v1alpha1.CachedObjectList) []*v1alpha1.CachedObject {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1alpha1.CachedObjectList, items []*v1alpha1.CachedObject) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
