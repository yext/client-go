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

package v1beta1

import (
	v1beta1 "github.com/yext/api/authorization/v1beta1"
	"github.com/yext/apimachinery/pkg/api/errors"
	"github.com/yext/apimachinery/pkg/labels"
	"github.com/yext/client-go/tools/cache"
)

// SubjectAccessReviewLister helps list SubjectAccessReviews.
type SubjectAccessReviewLister interface {
	// List lists all SubjectAccessReviews in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.SubjectAccessReview, err error)
	// Get retrieves the SubjectAccessReview from the index for a given name.
	Get(name string) (*v1beta1.SubjectAccessReview, error)
	SubjectAccessReviewListerExpansion
}

// subjectAccessReviewLister implements the SubjectAccessReviewLister interface.
type subjectAccessReviewLister struct {
	indexer cache.Indexer
}

// NewSubjectAccessReviewLister returns a new SubjectAccessReviewLister.
func NewSubjectAccessReviewLister(indexer cache.Indexer) SubjectAccessReviewLister {
	return &subjectAccessReviewLister{indexer: indexer}
}

// List lists all SubjectAccessReviews in the indexer.
func (s *subjectAccessReviewLister) List(selector labels.Selector) (ret []*v1beta1.SubjectAccessReview, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.SubjectAccessReview))
	})
	return ret, err
}

// Get retrieves the SubjectAccessReview from the index for a given name.
func (s *subjectAccessReviewLister) Get(name string) (*v1beta1.SubjectAccessReview, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("subjectaccessreview"), name)
	}
	return obj.(*v1beta1.SubjectAccessReview), nil
}
