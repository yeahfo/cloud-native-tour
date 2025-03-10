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
// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	helloyeahfogithubiov2 "github.com/yeahfo/cloud-native-tour/api/hello.yeahfo.github.io/v2"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// FooLister helps list Foos.
// All objects returned here must be treated as read-only.
type FooLister interface {
	// List lists all Foos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*helloyeahfogithubiov2.Foo, err error)
	// Foos returns an object that can list and get Foos.
	Foos(namespace string) FooNamespaceLister
	FooListerExpansion
}

// fooLister implements the FooLister interface.
type fooLister struct {
	listers.ResourceIndexer[*helloyeahfogithubiov2.Foo]
}

// NewFooLister returns a new FooLister.
func NewFooLister(indexer cache.Indexer) FooLister {
	return &fooLister{listers.New[*helloyeahfogithubiov2.Foo](indexer, helloyeahfogithubiov2.Resource("foo"))}
}

// Foos returns an object that can list and get Foos.
func (s *fooLister) Foos(namespace string) FooNamespaceLister {
	return fooNamespaceLister{listers.NewNamespaced[*helloyeahfogithubiov2.Foo](s.ResourceIndexer, namespace)}
}

// FooNamespaceLister helps list and get Foos.
// All objects returned here must be treated as read-only.
type FooNamespaceLister interface {
	// List lists all Foos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*helloyeahfogithubiov2.Foo, err error)
	// Get retrieves the Foo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*helloyeahfogithubiov2.Foo, error)
	FooNamespaceListerExpansion
}

// fooNamespaceLister implements the FooNamespaceLister
// interface.
type fooNamespaceLister struct {
	listers.ResourceIndexer[*helloyeahfogithubiov2.Foo]
}
