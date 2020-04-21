// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetScheduledJobRequest wrapper for the GetScheduledJob operation
type GetScheduledJobRequest struct {

	// The ID of the scheduled job.
	ScheduledJobId *string `mandatory:"true" contributesTo:"path" name:"scheduledJobId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetScheduledJobRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetScheduledJobRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetScheduledJobRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetScheduledJobResponse wrapper for the GetScheduledJob operation
type GetScheduledJobResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ScheduledJob instance
	ScheduledJob `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// some decimal number representing the number of seconds
	// the client should wait before polling this endpoint again
	RetryAfter *float32 `presentIn:"header" name:"retry-after"`
}

func (response GetScheduledJobResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetScheduledJobResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
