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
	v1 "github.com/yeahfo/cloud-native-tour/code-generator-examples/apiserver/apis/core/v1"
	corev1 "github.com/yeahfo/cloud-native-tour/code-generator-examples/apiserver/applyconfiguration/core/v1"
	typedcorev1 "github.com/yeahfo/cloud-native-tour/code-generator-examples/apiserver/clientset/versioned/typed/core/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeTestTypes implements TestTypeInterface
type fakeTestTypes struct {
	*gentype.FakeClientWithListAndApply[*v1.TestType, *v1.TestTypeList, *corev1.TestTypeApplyConfiguration]
	Fake *FakeCoreV1
}

func newFakeTestTypes(fake *FakeCoreV1, namespace string) typedcorev1.TestTypeInterface {
	return &fakeTestTypes{
		gentype.NewFakeClientWithListAndApply[*v1.TestType, *v1.TestTypeList, *corev1.TestTypeApplyConfiguration](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("testtypes"),
			v1.SchemeGroupVersion.WithKind("TestType"),
			func() *v1.TestType { return &v1.TestType{} },
			func() *v1.TestTypeList { return &v1.TestTypeList{} },
			func(dst, src *v1.TestTypeList) { dst.ListMeta = src.ListMeta },
			func(list *v1.TestTypeList) []*v1.TestType { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.TestTypeList, items []*v1.TestType) { list.Items = gentype.FromPointerSlice(items) },
		),
		fake,
	}
}
