// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetBdsInstanceRequest wrapper for the GetBdsInstance operation
type GetBdsInstanceRequest struct {

	// The OCID of the BDS instance
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetBdsInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetBdsInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetBdsInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetBdsInstanceResponse wrapper for the GetBdsInstance operation
type GetBdsInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The BdsInstance instance
	BdsInstance `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetBdsInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetBdsInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
