// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// GetExternalPluggableDatabaseRequest wrapper for the GetExternalPluggableDatabase operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExternalPluggableDatabase.go.html to see an example of how to use GetExternalPluggableDatabaseRequest.
type GetExternalPluggableDatabaseRequest struct {

	// The ExternalPluggableDatabaseId OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ExternalPluggableDatabaseId *string `mandatory:"true" contributesTo:"path" name:"externalPluggableDatabaseId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetExternalPluggableDatabaseRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetExternalPluggableDatabaseRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetExternalPluggableDatabaseRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetExternalPluggableDatabaseRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetExternalPluggableDatabaseResponse wrapper for the GetExternalPluggableDatabase operation
type GetExternalPluggableDatabaseResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ExternalPluggableDatabase instance
	ExternalPluggableDatabase `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetExternalPluggableDatabaseResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetExternalPluggableDatabaseResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
