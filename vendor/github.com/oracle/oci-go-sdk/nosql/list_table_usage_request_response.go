// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTableUsageRequest wrapper for the ListTableUsage operation
type ListTableUsageRequest struct {

	// A table name within the compartment, or a table OCID.
	TableNameOrId *string `mandatory:"true" contributesTo:"path" name:"tableNameOrId"`

	// The ID of a table's compartment. When a table is identified
	// by name, the compartmentId is often needed to provide
	// context for interpreting the name.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The start time to use for the request. If no time range
	// is set for this request, the most recent complete usage
	// record is returned.
	TimeStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStart"`

	// The end time to use for the request.
	TimeEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeEnd"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start
	// retrieving results. This is usually retrieved from a previous
	// list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTableUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTableUsageRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTableUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTableUsageResponse wrapper for the ListTableUsage operation
type ListTableUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TableUsageCollection instances
	TableUsageCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list,
	// if this header appears in the response, then a partial list
	// might have been returned. Include this value as the `page`
	// parameter for the subsequent GET request to get the next batch
	// of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListTableUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTableUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
