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

package v1

import (
	v1 "github.com/TheDeng/my_k8s_controller/pkg/apis/whu/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StudentLister helps list Students.
// All objects returned here must be treated as read-only.
type StudentLister interface {
	// List lists all Students in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Student, err error)
	// Students returns an object that can list and get Students.
	Students(namespace string) StudentNamespaceLister
	StudentListerExpansion
}

// studentLister implements the StudentLister interface.
type studentLister struct {
	indexer cache.Indexer
}

// NewStudentLister returns a new StudentLister.
func NewStudentLister(indexer cache.Indexer) StudentLister {
	return &studentLister{indexer: indexer}
}

// List lists all Students in the indexer.
func (s *studentLister) List(selector labels.Selector) (ret []*v1.Student, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Student))
	})
	return ret, err
}

// Students returns an object that can list and get Students.
func (s *studentLister) Students(namespace string) StudentNamespaceLister {
	return studentNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StudentNamespaceLister helps list and get Students.
// All objects returned here must be treated as read-only.
type StudentNamespaceLister interface {
	// List lists all Students in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Student, err error)
	// Get retrieves the Student from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Student, error)
	StudentNamespaceListerExpansion
}

// studentNamespaceLister implements the StudentNamespaceLister
// interface.
type studentNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Students in the indexer for a given namespace.
func (s studentNamespaceLister) List(selector labels.Selector) (ret []*v1.Student, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Student))
	})
	return ret, err
}

// Get retrieves the Student from the indexer for a given namespace and name.
func (s studentNamespaceLister) Get(name string) (*v1.Student, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("student"), name)
	}
	return obj.(*v1.Student), nil
}
