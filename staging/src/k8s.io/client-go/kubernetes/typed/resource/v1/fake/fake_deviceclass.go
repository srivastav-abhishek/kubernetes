/*
Copyright The Kubernetes Authors.

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
	v1 "k8s.io/api/resource/v1"
	resourcev1 "k8s.io/client-go/applyconfigurations/resource/v1"
	gentype "k8s.io/client-go/gentype"
	typedresourcev1 "k8s.io/client-go/kubernetes/typed/resource/v1"
)

// fakeDeviceClasses implements DeviceClassInterface
type fakeDeviceClasses struct {
	*gentype.FakeClientWithListAndApply[*v1.DeviceClass, *v1.DeviceClassList, *resourcev1.DeviceClassApplyConfiguration]
	Fake *FakeResourceV1
}

func newFakeDeviceClasses(fake *FakeResourceV1) typedresourcev1.DeviceClassInterface {
	return &fakeDeviceClasses{
		gentype.NewFakeClientWithListAndApply[*v1.DeviceClass, *v1.DeviceClassList, *resourcev1.DeviceClassApplyConfiguration](
			fake.Fake,
			"",
			v1.SchemeGroupVersion.WithResource("deviceclasses"),
			v1.SchemeGroupVersion.WithKind("DeviceClass"),
			func() *v1.DeviceClass { return &v1.DeviceClass{} },
			func() *v1.DeviceClassList { return &v1.DeviceClassList{} },
			func(dst, src *v1.DeviceClassList) { dst.ListMeta = src.ListMeta },
			func(list *v1.DeviceClassList) []*v1.DeviceClass { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.DeviceClassList, items []*v1.DeviceClass) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
