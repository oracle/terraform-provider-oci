// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetRowRequest wrapper for the GetRow operation
type GetRowRequest struct {

	// A table name within the compartment, or a table OCID.
	TableNameOrId *string `mandatory:"true" contributesTo:"path" name:"tableNameOrId"`

	// An array of strings, each of the format "column-name:value",
	// representing the primary key of the row.
	Key []string `contributesTo:"query" name:"key" collectionFormat:"multi"`

	// The ID of a table's compartment. When a table is identified
	// by name, the compartmentId is often needed to provide
	// context for interpreting the name.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Consistency requirement for a read operation.
	Consistency GetRowConsistencyEnum `mandatory:"false" contributesTo:"query" name:"consistency" omitEmpty:"true"`

	// Timeout setting for this operation.
	TimeoutInMs *int `mandatory:"false" contributesTo:"query" name:"timeoutInMs"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetRowRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetRowRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetRowRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetRowResponse wrapper for the GetRow operation
type GetRowResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Row instance
	Row `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetRowResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetRowResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetRowConsistencyEnum Enum with underlying type: string
type GetRowConsistencyEnum string

// Set of constants representing the allowable values for GetRowConsistencyEnum
const (
	GetRowConsistencyEventual GetRowConsistencyEnum = "EVENTUAL"
	GetRowConsistencyAbsolute GetRowConsistencyEnum = "ABSOLUTE"
)

var mappingGetRowConsistency = map[string]GetRowConsistencyEnum{
	"EVENTUAL": GetRowConsistencyEventual,
	"ABSOLUTE": GetRowConsistencyAbsolute,
}

// GetGetRowConsistencyEnumValues Enumerates the set of values for GetRowConsistencyEnum
func GetGetRowConsistencyEnumValues() []GetRowConsistencyEnum {
	values := make([]GetRowConsistencyEnum, 0)
	for _, v := range mappingGetRowConsistency {
		values = append(values, v)
	}
	return values
}
