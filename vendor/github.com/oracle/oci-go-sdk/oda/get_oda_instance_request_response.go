// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetOdaInstanceRequest wrapper for the GetOdaInstance operation
type GetOdaInstanceRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetOdaInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetOdaInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetOdaInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetOdaInstanceResponse wrapper for the GetOdaInstance operation
type GetOdaInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OdaInstance instance
	OdaInstance `presentIn:"body"`

	// For use in a PUT or DELETE `if-match` query parameter for optimistic concurrency control.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you contact Oracle
	// about this request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetOdaInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetOdaInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
