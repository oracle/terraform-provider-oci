// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeleteRowRequest wrapper for the DeleteRow operation
type DeleteRowRequest struct {

	// A table name within the compartment, or a table OCID.
	TableNameOrId *string `mandatory:"true" contributesTo:"path" name:"tableNameOrId"`

	// An array of strings, each of the format "column-name:value",
	// representing the primary key of the row.
	Key []string `contributesTo:"query" name:"key" collectionFormat:"multi"`

	// The ID of a table's compartment. When a table is identified
	// by name, the compartmentId is often needed to provide
	// context for interpreting the name.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// If true, and the operation fails due to an option setting
	// (ifVersion et al), then the existing row will be returned.
	IsGetReturnRow *bool `mandatory:"false" contributesTo:"query" name:"isGetReturnRow"`

	// Timeout setting for this operation.
	TimeoutInMs *int `mandatory:"false" contributesTo:"query" name:"timeoutInMs"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteRowRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteRowRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteRowRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteRowResponse wrapper for the DeleteRow operation
type DeleteRowResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DeleteRowResult instance
	DeleteRowResult `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteRowResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteRowResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
