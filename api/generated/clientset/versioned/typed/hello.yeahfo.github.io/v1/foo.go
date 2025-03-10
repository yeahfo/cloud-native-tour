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

package v1

import (
	context "context"

	applyconfigurationhelloyeahfogithubiov1 "github.com/yeahfo/cloud-native-tour/api/generated/applyconfiguration/hello.yeahfo.github.io/v1"
	scheme "github.com/yeahfo/cloud-native-tour/api/generated/clientset/versioned/scheme"
	helloyeahfogithubiov1 "github.com/yeahfo/cloud-native-tour/api/hello.yeahfo.github.io/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// FoosGetter has a method to return a FooInterface.
// A group's client should implement this interface.
type FoosGetter interface {
	Foos(namespace string) FooInterface
}

// FooInterface has methods to work with Foo resources.
type FooInterface interface {
	Create(ctx context.Context, foo *helloyeahfogithubiov1.Foo, opts metav1.CreateOptions) (*helloyeahfogithubiov1.Foo, error)
	Update(ctx context.Context, foo *helloyeahfogithubiov1.Foo, opts metav1.UpdateOptions) (*helloyeahfogithubiov1.Foo, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*helloyeahfogithubiov1.Foo, error)
	List(ctx context.Context, opts metav1.ListOptions) (*helloyeahfogithubiov1.FooList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *helloyeahfogithubiov1.Foo, err error)
	Apply(ctx context.Context, foo *applyconfigurationhelloyeahfogithubiov1.FooApplyConfiguration, opts metav1.ApplyOptions) (result *helloyeahfogithubiov1.Foo, err error)
	FooExpansion
}

// foos implements FooInterface
type foos struct {
	*gentype.ClientWithListAndApply[*helloyeahfogithubiov1.Foo, *helloyeahfogithubiov1.FooList, *applyconfigurationhelloyeahfogithubiov1.FooApplyConfiguration]
}

// newFoos returns a Foos
func newFoos(c *HelloV1Client, namespace string) *foos {
	return &foos{
		gentype.NewClientWithListAndApply[*helloyeahfogithubiov1.Foo, *helloyeahfogithubiov1.FooList, *applyconfigurationhelloyeahfogithubiov1.FooApplyConfiguration](
			"foos",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *helloyeahfogithubiov1.Foo { return &helloyeahfogithubiov1.Foo{} },
			func() *helloyeahfogithubiov1.FooList { return &helloyeahfogithubiov1.FooList{} },
			gentype.PrefersProtobuf[*helloyeahfogithubiov1.Foo](),
		),
	}
}
