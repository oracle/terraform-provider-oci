// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListLogAnalyticsLogGroupsRequest wrapper for the ListLogAnalyticsLogGroups operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsLogGroups.go.html to see an example of how to use ListLogAnalyticsLogGroupsRequest.
type ListLogAnalyticsLogGroupsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only log analytics log groups whose displayName matches the entire display name given.
	// The match is case-insensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogAnalyticsLogGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListLogAnalyticsLogGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogAnalyticsLogGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogAnalyticsLogGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogAnalyticsLogGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogAnalyticsLogGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogAnalyticsLogGroupsResponse wrapper for the ListLogAnalyticsLogGroups operation
type ListLogAnalyticsLogGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsLogGroupSummaryCollection instances
	LogAnalyticsLogGroupSummaryCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogAnalyticsLogGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogAnalyticsLogGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogAnalyticsLogGroupsSortOrderEnum Enum with underlying type: string
type ListLogAnalyticsLogGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListLogAnalyticsLogGroupsSortOrderEnum
const (
	ListLogAnalyticsLogGroupsSortOrderAsc  ListLogAnalyticsLogGroupsSortOrderEnum = "ASC"
	ListLogAnalyticsLogGroupsSortOrderDesc ListLogAnalyticsLogGroupsSortOrderEnum = "DESC"
)

var mappingListLogAnalyticsLogGroupsSortOrder = map[string]ListLogAnalyticsLogGroupsSortOrderEnum{
	"ASC":  ListLogAnalyticsLogGroupsSortOrderAsc,
	"DESC": ListLogAnalyticsLogGroupsSortOrderDesc,
}

// GetListLogAnalyticsLogGroupsSortOrderEnumValues Enumerates the set of values for ListLogAnalyticsLogGroupsSortOrderEnum
func GetListLogAnalyticsLogGroupsSortOrderEnumValues() []ListLogAnalyticsLogGroupsSortOrderEnum {
	values := make([]ListLogAnalyticsLogGroupsSortOrderEnum, 0)
	for _, v := range mappingListLogAnalyticsLogGroupsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsLogGroupsSortByEnum Enum with underlying type: string
type ListLogAnalyticsLogGroupsSortByEnum string

// Set of constants representing the allowable values for ListLogAnalyticsLogGroupsSortByEnum
const (
	ListLogAnalyticsLogGroupsSortByTimecreated ListLogAnalyticsLogGroupsSortByEnum = "timeCreated"
	ListLogAnalyticsLogGroupsSortByTimeupdated ListLogAnalyticsLogGroupsSortByEnum = "timeUpdated"
	ListLogAnalyticsLogGroupsSortByDisplayname ListLogAnalyticsLogGroupsSortByEnum = "displayName"
)

var mappingListLogAnalyticsLogGroupsSortBy = map[string]ListLogAnalyticsLogGroupsSortByEnum{
	"timeCreated": ListLogAnalyticsLogGroupsSortByTimecreated,
	"timeUpdated": ListLogAnalyticsLogGroupsSortByTimeupdated,
	"displayName": ListLogAnalyticsLogGroupsSortByDisplayname,
}

// GetListLogAnalyticsLogGroupsSortByEnumValues Enumerates the set of values for ListLogAnalyticsLogGroupsSortByEnum
func GetListLogAnalyticsLogGroupsSortByEnumValues() []ListLogAnalyticsLogGroupsSortByEnum {
	values := make([]ListLogAnalyticsLogGroupsSortByEnum, 0)
	for _, v := range mappingListLogAnalyticsLogGroupsSortBy {
		values = append(values, v)
	}
	return values
}
