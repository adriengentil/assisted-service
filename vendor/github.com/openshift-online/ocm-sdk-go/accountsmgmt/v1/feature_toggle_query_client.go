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
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// FeatureToggleQueryClient is the client of the 'feature_toggle_query' resource.
//
// Manages feature toggle query.
type FeatureToggleQueryClient struct {
	transport http.RoundTripper
	path      string
}

// NewFeatureToggleQueryClient creates a new client for the 'feature_toggle_query'
// resource using the given transport to send the requests and receive the
// responses.
func NewFeatureToggleQueryClient(transport http.RoundTripper, path string) *FeatureToggleQueryClient {
	return &FeatureToggleQueryClient{
		transport: transport,
		path:      path,
	}
}

// Post creates a request for the 'post' method.
//
// Retrieves the details of the feature toggle by providing query context
func (c *FeatureToggleQueryClient) Post() *FeatureToggleQueryPostRequest {
	return &FeatureToggleQueryPostRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// FeatureToggleQueryPostRequest is the request for the 'post' method.
type FeatureToggleQueryPostRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
	request   *FeatureToggleQueryRequest
}

// Parameter adds a query parameter.
func (r *FeatureToggleQueryPostRequest) Parameter(name string, value interface{}) *FeatureToggleQueryPostRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *FeatureToggleQueryPostRequest) Header(name string, value interface{}) *FeatureToggleQueryPostRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Request sets the value of the 'request' parameter.
func (r *FeatureToggleQueryPostRequest) Request(value *FeatureToggleQueryRequest) *FeatureToggleQueryPostRequest {
	r.request = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *FeatureToggleQueryPostRequest) Send() (result *FeatureToggleQueryPostResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *FeatureToggleQueryPostRequest) SendContext(ctx context.Context) (result *FeatureToggleQueryPostResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := &bytes.Buffer{}
	err = writeFeatureToggleQueryPostRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "POST",
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &FeatureToggleQueryPostResponse{}
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = readFeatureToggleQueryPostResponse(result, response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'post' method.
func (r *FeatureToggleQueryPostRequest) marshal(writer io.Writer) error {
	stream := helpers.NewStream(writer)
	r.stream(stream)
	return stream.Error
}
func (r *FeatureToggleQueryPostRequest) stream(stream *jsoniter.Stream) {
}

// FeatureToggleQueryPostResponse is the response for the 'post' method.
type FeatureToggleQueryPostResponse struct {
	status   int
	header   http.Header
	err      *errors.Error
	response *FeatureToggle
}

// Status returns the response status code.
func (r *FeatureToggleQueryPostResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *FeatureToggleQueryPostResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *FeatureToggleQueryPostResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Response returns the value of the 'response' parameter.
func (r *FeatureToggleQueryPostResponse) Response() *FeatureToggle {
	if r == nil {
		return nil
	}
	return r.response
}

// GetResponse returns the value of the 'response' parameter and
// a flag indicating if the parameter has a value.
func (r *FeatureToggleQueryPostResponse) GetResponse() (value *FeatureToggle, ok bool) {
	ok = r != nil && r.response != nil
	if ok {
		value = r.response
	}
	return
}
