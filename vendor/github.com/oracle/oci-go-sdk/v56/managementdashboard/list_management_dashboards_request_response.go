// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementdashboard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListManagementDashboardsRequest wrapper for the ListManagementDashboards operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ListManagementDashboards.go.html to see an example of how to use ListManagementDashboardsRequest.
type ListManagementDashboardsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page on which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementDashboardsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is the default.
	SortBy ListManagementDashboardsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementDashboardsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementDashboardsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementDashboardsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementDashboardsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagementDashboardsResponse wrapper for the ListManagementDashboards operation
type ListManagementDashboardsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagementDashboardCollection instances
	ManagementDashboardCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementDashboardsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementDashboardsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementDashboardsSortOrderEnum Enum with underlying type: string
type ListManagementDashboardsSortOrderEnum string

// Set of constants representing the allowable values for ListManagementDashboardsSortOrderEnum
const (
	ListManagementDashboardsSortOrderAsc  ListManagementDashboardsSortOrderEnum = "ASC"
	ListManagementDashboardsSortOrderDesc ListManagementDashboardsSortOrderEnum = "DESC"
)

var mappingListManagementDashboardsSortOrder = map[string]ListManagementDashboardsSortOrderEnum{
	"ASC":  ListManagementDashboardsSortOrderAsc,
	"DESC": ListManagementDashboardsSortOrderDesc,
}

// GetListManagementDashboardsSortOrderEnumValues Enumerates the set of values for ListManagementDashboardsSortOrderEnum
func GetListManagementDashboardsSortOrderEnumValues() []ListManagementDashboardsSortOrderEnum {
	values := make([]ListManagementDashboardsSortOrderEnum, 0)
	for _, v := range mappingListManagementDashboardsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagementDashboardsSortByEnum Enum with underlying type: string
type ListManagementDashboardsSortByEnum string

// Set of constants representing the allowable values for ListManagementDashboardsSortByEnum
const (
	ListManagementDashboardsSortByTimecreated ListManagementDashboardsSortByEnum = "timeCreated"
	ListManagementDashboardsSortByDisplayname ListManagementDashboardsSortByEnum = "displayName"
)

var mappingListManagementDashboardsSortBy = map[string]ListManagementDashboardsSortByEnum{
	"timeCreated": ListManagementDashboardsSortByTimecreated,
	"displayName": ListManagementDashboardsSortByDisplayname,
}

// GetListManagementDashboardsSortByEnumValues Enumerates the set of values for ListManagementDashboardsSortByEnum
func GetListManagementDashboardsSortByEnumValues() []ListManagementDashboardsSortByEnum {
	values := make([]ListManagementDashboardsSortByEnum, 0)
	for _, v := range mappingListManagementDashboardsSortBy {
		values = append(values, v)
	}
	return values
}
