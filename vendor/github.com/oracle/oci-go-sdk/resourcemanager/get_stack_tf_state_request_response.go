// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// GetStackTfStateRequest wrapper for the GetStackTfState operation
type GetStackTfStateRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the stack.
	StackId *string `mandatory:"true" contributesTo:"path" name:"stackId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetStackTfStateRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetStackTfStateRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetStackTfStateRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetStackTfStateResponse wrapper for the GetStackTfState operation
type GetStackTfStateResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetStackTfStateResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetStackTfStateResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
