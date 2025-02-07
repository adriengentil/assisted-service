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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalAlertInfoList writes a list of values of the 'alert_info' type to
// the given writer.
func MarshalAlertInfoList(list []*AlertInfo, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeAlertInfoList(list, stream)
	stream.Flush()
	return stream.Error
}

// writeAlertInfoList writes a list of value of the 'alert_info' type to
// the given stream.
func writeAlertInfoList(list []*AlertInfo, stream *jsoniter.Stream) {
	stream.WriteArrayStart()
	for i, value := range list {
		if i > 0 {
			stream.WriteMore()
		}
		writeAlertInfo(value, stream)
	}
	stream.WriteArrayEnd()
}

// UnmarshalAlertInfoList reads a list of values of the 'alert_info' type
// from the given source, which can be a slice of bytes, a string or a reader.
func UnmarshalAlertInfoList(source interface{}) (items []*AlertInfo, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	items = readAlertInfoList(iterator)
	err = iterator.Error
	return
}

// readAlertInfoList reads list of values of the ”alert_info' type from
// the given iterator.
func readAlertInfoList(iterator *jsoniter.Iterator) []*AlertInfo {
	list := []*AlertInfo{}
	for iterator.ReadArray() {
		item := readAlertInfo(iterator)
		list = append(list, item)
	}
	return list
}
