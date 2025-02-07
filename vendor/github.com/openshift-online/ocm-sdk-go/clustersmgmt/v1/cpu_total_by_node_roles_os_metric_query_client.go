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
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// CPUTotalByNodeRolesOSMetricQueryClient is the client of the 'CPU_total_by_node_roles_OS_metric_query' resource.
//
// Total cpu capacity in the cluster by node role and operating system.
type CPUTotalByNodeRolesOSMetricQueryClient struct {
	transport http.RoundTripper
	path      string
}

// NewCPUTotalByNodeRolesOSMetricQueryClient creates a new client for the 'CPU_total_by_node_roles_OS_metric_query'
// resource using the given transport to send the requests and receive the
// responses.
func NewCPUTotalByNodeRolesOSMetricQueryClient(transport http.RoundTripper, path string) *CPUTotalByNodeRolesOSMetricQueryClient {
	return &CPUTotalByNodeRolesOSMetricQueryClient{
		transport: transport,
		path:      path,
	}
}

// Get creates a request for the 'get' method.
//
// Retrieves the metrics.
func (c *CPUTotalByNodeRolesOSMetricQueryClient) Get() *CPUTotalByNodeRolesOSMetricQueryGetRequest {
	return &CPUTotalByNodeRolesOSMetricQueryGetRequest{
		transport: c.transport,
		path:      c.path,
	}
}

// CPUTotalByNodeRolesOSMetricQueryPollRequest is the request for the Poll method.
type CPUTotalByNodeRolesOSMetricQueryPollRequest struct {
	request    *CPUTotalByNodeRolesOSMetricQueryGetRequest
	interval   time.Duration
	statuses   []int
	predicates []func(interface{}) bool
}

// Parameter adds a query parameter to all the requests that will be used to retrieve the object.
func (r *CPUTotalByNodeRolesOSMetricQueryPollRequest) Parameter(name string, value interface{}) *CPUTotalByNodeRolesOSMetricQueryPollRequest {
	r.request.Parameter(name, value)
	return r
}

// Header adds a request header to all the requests that will be used to retrieve the object.
func (r *CPUTotalByNodeRolesOSMetricQueryPollRequest) Header(name string, value interface{}) *CPUTotalByNodeRolesOSMetricQueryPollRequest {
	r.request.Header(name, value)
	return r
}

// Interval sets the polling interval. This parameter is mandatory and must be greater than zero.
func (r *CPUTotalByNodeRolesOSMetricQueryPollRequest) Interval(value time.Duration) *CPUTotalByNodeRolesOSMetricQueryPollRequest {
	r.interval = value
	return r
}

// Status set the expected status of the response. Multiple values can be set calling this method
// multiple times. The response will be considered successful if the status is any of those values.
func (r *CPUTotalByNodeRolesOSMetricQueryPollRequest) Status(value int) *CPUTotalByNodeRolesOSMetricQueryPollRequest {
	r.statuses = append(r.statuses, value)
	return r
}

// Predicate adds a predicate that the response should satisfy be considered successful. Multiple
// predicates can be set calling this method multiple times. The response will be considered successful
// if all the predicates are satisfied.
func (r *CPUTotalByNodeRolesOSMetricQueryPollRequest) Predicate(value func(*CPUTotalByNodeRolesOSMetricQueryGetResponse) bool) *CPUTotalByNodeRolesOSMetricQueryPollRequest {
	r.predicates = append(r.predicates, func(response interface{}) bool {
		return value(response.(*CPUTotalByNodeRolesOSMetricQueryGetResponse))
	})
	return r
}

// StartContext starts the polling loop. Responses will be considered successful if the status is one of
// the values specified with the Status method and if all the predicates specified with the Predicate
// method return nil.
//
// The context must have a timeout or deadline, otherwise this method will immediately return an error.
func (r *CPUTotalByNodeRolesOSMetricQueryPollRequest) StartContext(ctx context.Context) (response *CPUTotalByNodeRolesOSMetricQueryPollResponse, err error) {
	result, err := helpers.PollContext(ctx, r.interval, r.statuses, r.predicates, r.task)
	if result != nil {
		response = &CPUTotalByNodeRolesOSMetricQueryPollResponse{
			response: result.(*CPUTotalByNodeRolesOSMetricQueryGetResponse),
		}
	}
	return
}

// task adapts the types of the request/response types so that they can be used with the generic
// polling function from the helpers package.
func (r *CPUTotalByNodeRolesOSMetricQueryPollRequest) task(ctx context.Context) (status int, result interface{}, err error) {
	response, err := r.request.SendContext(ctx)
	if response != nil {
		status = response.Status()
		result = response
	}
	return
}

// CPUTotalByNodeRolesOSMetricQueryPollResponse is the response for the Poll method.
type CPUTotalByNodeRolesOSMetricQueryPollResponse struct {
	response *CPUTotalByNodeRolesOSMetricQueryGetResponse
}

// Status returns the response status code.
func (r *CPUTotalByNodeRolesOSMetricQueryPollResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.response.Status()
}

// Header returns header of the response.
func (r *CPUTotalByNodeRolesOSMetricQueryPollResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.response.Header()
}

// Error returns the response error.
func (r *CPUTotalByNodeRolesOSMetricQueryPollResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.response.Error()
}

// Body returns the value of the 'body' parameter.
func (r *CPUTotalByNodeRolesOSMetricQueryPollResponse) Body() *CPUTotalsNodeRoleOSMetricNode {
	return r.response.Body()
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
func (r *CPUTotalByNodeRolesOSMetricQueryPollResponse) GetBody() (value *CPUTotalsNodeRoleOSMetricNode, ok bool) {
	return r.response.GetBody()
}

// Poll creates a request to repeatedly retrieve the object till the response has one of a given set
// of states and satisfies a set of predicates.
func (c *CPUTotalByNodeRolesOSMetricQueryClient) Poll() *CPUTotalByNodeRolesOSMetricQueryPollRequest {
	return &CPUTotalByNodeRolesOSMetricQueryPollRequest{
		request: c.Get(),
	}
}

// CPUTotalByNodeRolesOSMetricQueryGetRequest is the request for the 'get' method.
type CPUTotalByNodeRolesOSMetricQueryGetRequest struct {
	transport http.RoundTripper
	path      string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *CPUTotalByNodeRolesOSMetricQueryGetRequest) Parameter(name string, value interface{}) *CPUTotalByNodeRolesOSMetricQueryGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *CPUTotalByNodeRolesOSMetricQueryGetRequest) Header(name string, value interface{}) *CPUTotalByNodeRolesOSMetricQueryGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *CPUTotalByNodeRolesOSMetricQueryGetRequest) Send() (result *CPUTotalByNodeRolesOSMetricQueryGetResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *CPUTotalByNodeRolesOSMetricQueryGetRequest) SendContext(ctx context.Context) (result *CPUTotalByNodeRolesOSMetricQueryGetResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "GET",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &CPUTotalByNodeRolesOSMetricQueryGetResponse{}
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
	err = readCPUTotalByNodeRolesOSMetricQueryGetResponse(result, response.Body)
	if err != nil {
		return
	}
	return
}

// CPUTotalByNodeRolesOSMetricQueryGetResponse is the response for the 'get' method.
type CPUTotalByNodeRolesOSMetricQueryGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *CPUTotalsNodeRoleOSMetricNode
}

// Status returns the response status code.
func (r *CPUTotalByNodeRolesOSMetricQueryGetResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *CPUTotalByNodeRolesOSMetricQueryGetResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *CPUTotalByNodeRolesOSMetricQueryGetResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
func (r *CPUTotalByNodeRolesOSMetricQueryGetResponse) Body() *CPUTotalsNodeRoleOSMetricNode {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
func (r *CPUTotalByNodeRolesOSMetricQueryGetResponse) GetBody() (value *CPUTotalsNodeRoleOSMetricNode, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}
