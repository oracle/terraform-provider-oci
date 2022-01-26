// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetools

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListWorkRequestErrorsRequest wrapper for the ListWorkRequestErrors operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrorsRequest.
type ListWorkRequestErrorsRequest struct {

	// The ID of the asynchronous request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListWorkRequestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListWorkRequestErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkRequestErrorsResponse wrapper for the ListWorkRequestErrors operation
type ListWorkRequestErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestErrorCollection instances
	WorkRequestErrorCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListWorkRequestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestErrorsSortOrderEnum Enum with underlying type: string
type ListWorkRequestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestErrorsSortOrderEnum
const (
	ListWorkRequestErrorsSortOrderAsc  ListWorkRequestErrorsSortOrderEnum = "ASC"
	ListWorkRequestErrorsSortOrderDesc ListWorkRequestErrorsSortOrderEnum = "DESC"
)

var mappingListWorkRequestErrorsSortOrder = map[string]ListWorkRequestErrorsSortOrderEnum{
	"ASC":  ListWorkRequestErrorsSortOrderAsc,
	"DESC": ListWorkRequestErrorsSortOrderDesc,
}

// GetListWorkRequestErrorsSortOrderEnumValues Enumerates the set of values for ListWorkRequestErrorsSortOrderEnum
func GetListWorkRequestErrorsSortOrderEnumValues() []ListWorkRequestErrorsSortOrderEnum {
	values := make([]ListWorkRequestErrorsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestErrorsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestErrorsSortByEnum Enum with underlying type: string
type ListWorkRequestErrorsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestErrorsSortByEnum
const (
	ListWorkRequestErrorsSortByTimecreated ListWorkRequestErrorsSortByEnum = "timeCreated"
	ListWorkRequestErrorsSortByDisplayname ListWorkRequestErrorsSortByEnum = "displayName"
)

var mappingListWorkRequestErrorsSortBy = map[string]ListWorkRequestErrorsSortByEnum{
	"timeCreated": ListWorkRequestErrorsSortByTimecreated,
	"displayName": ListWorkRequestErrorsSortByDisplayname,
}

// GetListWorkRequestErrorsSortByEnumValues Enumerates the set of values for ListWorkRequestErrorsSortByEnum
func GetListWorkRequestErrorsSortByEnumValues() []ListWorkRequestErrorsSortByEnum {
	values := make([]ListWorkRequestErrorsSortByEnum, 0)
	for _, v := range mappingListWorkRequestErrorsSortBy {
		values = append(values, v)
	}
	return values
}
