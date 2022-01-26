// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListWorkRequestLogsRequest wrapper for the ListWorkRequestLogs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogsRequest.
type ListWorkRequestLogsRequest struct {

	// The ID of the asynchronous work request to retrieve.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListWorkRequestLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListWorkRequestLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkRequestLogsResponse wrapper for the ListWorkRequestLogs operation
type ListWorkRequestLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkRequestLogEntry instances
	Items []WorkRequestLogEntry `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestLogsSortOrderEnum Enum with underlying type: string
type ListWorkRequestLogsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestLogsSortOrderEnum
const (
	ListWorkRequestLogsSortOrderAsc  ListWorkRequestLogsSortOrderEnum = "ASC"
	ListWorkRequestLogsSortOrderDesc ListWorkRequestLogsSortOrderEnum = "DESC"
)

var mappingListWorkRequestLogsSortOrder = map[string]ListWorkRequestLogsSortOrderEnum{
	"ASC":  ListWorkRequestLogsSortOrderAsc,
	"DESC": ListWorkRequestLogsSortOrderDesc,
}

// GetListWorkRequestLogsSortOrderEnumValues Enumerates the set of values for ListWorkRequestLogsSortOrderEnum
func GetListWorkRequestLogsSortOrderEnumValues() []ListWorkRequestLogsSortOrderEnum {
	values := make([]ListWorkRequestLogsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestLogsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestLogsSortByEnum Enum with underlying type: string
type ListWorkRequestLogsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestLogsSortByEnum
const (
	ListWorkRequestLogsSortByTimeCreated ListWorkRequestLogsSortByEnum = "TIME_CREATED"
	ListWorkRequestLogsSortByDisplayName ListWorkRequestLogsSortByEnum = "DISPLAY_NAME"
)

var mappingListWorkRequestLogsSortBy = map[string]ListWorkRequestLogsSortByEnum{
	"TIME_CREATED": ListWorkRequestLogsSortByTimeCreated,
	"DISPLAY_NAME": ListWorkRequestLogsSortByDisplayName,
}

// GetListWorkRequestLogsSortByEnumValues Enumerates the set of values for ListWorkRequestLogsSortByEnum
func GetListWorkRequestLogsSortByEnumValues() []ListWorkRequestLogsSortByEnum {
	values := make([]ListWorkRequestLogsSortByEnum, 0)
	for _, v := range mappingListWorkRequestLogsSortBy {
		values = append(values, v)
	}
	return values
}
