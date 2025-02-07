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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalFeatureToggleQueryRequestList writes a list of values of the 'feature_toggle_query_request' type to
// the given writer.
func MarshalFeatureToggleQueryRequestList(list []*FeatureToggleQueryRequest, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeFeatureToggleQueryRequestList(list, stream)
	stream.Flush()
	return stream.Error
}

// writeFeatureToggleQueryRequestList writes a list of value of the 'feature_toggle_query_request' type to
// the given stream.
func writeFeatureToggleQueryRequestList(list []*FeatureToggleQueryRequest, stream *jsoniter.Stream) {
	stream.WriteArrayStart()
	for i, value := range list {
		if i > 0 {
			stream.WriteMore()
		}
		writeFeatureToggleQueryRequest(value, stream)
	}
	stream.WriteArrayEnd()
}

// UnmarshalFeatureToggleQueryRequestList reads a list of values of the 'feature_toggle_query_request' type
// from the given source, which can be a slice of bytes, a string or a reader.
func UnmarshalFeatureToggleQueryRequestList(source interface{}) (items []*FeatureToggleQueryRequest, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	items = readFeatureToggleQueryRequestList(iterator)
	err = iterator.Error
	return
}

// readFeatureToggleQueryRequestList reads list of values of the ”feature_toggle_query_request' type from
// the given iterator.
func readFeatureToggleQueryRequestList(iterator *jsoniter.Iterator) []*FeatureToggleQueryRequest {
	list := []*FeatureToggleQueryRequest{}
	for iterator.ReadArray() {
		item := readFeatureToggleQueryRequest(iterator)
		list = append(list, item)
	}
	return list
}
