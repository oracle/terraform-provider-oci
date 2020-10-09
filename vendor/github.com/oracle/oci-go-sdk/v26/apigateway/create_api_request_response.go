// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apigateway

import (
	"github.com/oracle/oci-go-sdk/v26/common"
	"net/http"
)

// CreateApiRequest wrapper for the CreateApi operation
type CreateApiRequest struct {

	// Details for the new API.
	CreateApiDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request id for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateApiRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateApiRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateApiRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateApiResponse wrapper for the CreateApi operation
type CreateApiResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Api instance
	Api `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// The OCID of the work request. Use
	// GetWorkRequest with
	// this id to track the status
	// of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// id.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Location of the resource.
	Location *string `presentIn:"header" name:"location"`
}

func (response CreateApiResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateApiResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
