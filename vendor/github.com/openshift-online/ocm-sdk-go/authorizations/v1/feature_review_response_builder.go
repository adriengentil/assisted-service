/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

// FeatureReviewResponseBuilder contains the data and logic needed to build 'feature_review_response' objects.
//
// Representation of a feature review response
type FeatureReviewResponseBuilder struct {
	bitmap_   uint32
	featureID string
	enabled   bool
}

// NewFeatureReviewResponse creates a new builder of 'feature_review_response' objects.
func NewFeatureReviewResponse() *FeatureReviewResponseBuilder {
	return &FeatureReviewResponseBuilder{}
}

// Enabled sets the value of the 'enabled' attribute to the given value.
func (b *FeatureReviewResponseBuilder) Enabled(value bool) *FeatureReviewResponseBuilder {
	b.enabled = value
	b.bitmap_ |= 1
	return b
}

// FeatureID sets the value of the 'feature_ID' attribute to the given value.
func (b *FeatureReviewResponseBuilder) FeatureID(value string) *FeatureReviewResponseBuilder {
	b.featureID = value
	b.bitmap_ |= 2
	return b
}

// Copy copies the attributes of the given object into this builder, discarding any previous values.
func (b *FeatureReviewResponseBuilder) Copy(object *FeatureReviewResponse) *FeatureReviewResponseBuilder {
	if object == nil {
		return b
	}
	b.bitmap_ = object.bitmap_
	b.enabled = object.enabled
	b.featureID = object.featureID
	return b
}

// Build creates a 'feature_review_response' object using the configuration stored in the builder.
func (b *FeatureReviewResponseBuilder) Build() (object *FeatureReviewResponse, err error) {
	object = new(FeatureReviewResponse)
	object.bitmap_ = b.bitmap_
	object.enabled = b.enabled
	object.featureID = b.featureID
	return
}
