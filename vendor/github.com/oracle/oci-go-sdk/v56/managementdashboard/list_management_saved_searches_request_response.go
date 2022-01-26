// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementdashboard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListManagementSavedSearchesRequest wrapper for the ListManagementSavedSearches operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ListManagementSavedSearches.go.html to see an example of how to use ListManagementSavedSearchesRequest.
type ListManagementSavedSearchesRequest struct {

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
	SortOrder ListManagementSavedSearchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is the default.
	SortBy ListManagementSavedSearchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementSavedSearchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementSavedSearchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagementSavedSearchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementSavedSearchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagementSavedSearchesResponse wrapper for the ListManagementSavedSearches operation
type ListManagementSavedSearchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagementSavedSearchCollection instances
	ManagementSavedSearchCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementSavedSearchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementSavedSearchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementSavedSearchesSortOrderEnum Enum with underlying type: string
type ListManagementSavedSearchesSortOrderEnum string

// Set of constants representing the allowable values for ListManagementSavedSearchesSortOrderEnum
const (
	ListManagementSavedSearchesSortOrderAsc  ListManagementSavedSearchesSortOrderEnum = "ASC"
	ListManagementSavedSearchesSortOrderDesc ListManagementSavedSearchesSortOrderEnum = "DESC"
)

var mappingListManagementSavedSearchesSortOrder = map[string]ListManagementSavedSearchesSortOrderEnum{
	"ASC":  ListManagementSavedSearchesSortOrderAsc,
	"DESC": ListManagementSavedSearchesSortOrderDesc,
}

// GetListManagementSavedSearchesSortOrderEnumValues Enumerates the set of values for ListManagementSavedSearchesSortOrderEnum
func GetListManagementSavedSearchesSortOrderEnumValues() []ListManagementSavedSearchesSortOrderEnum {
	values := make([]ListManagementSavedSearchesSortOrderEnum, 0)
	for _, v := range mappingListManagementSavedSearchesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagementSavedSearchesSortByEnum Enum with underlying type: string
type ListManagementSavedSearchesSortByEnum string

// Set of constants representing the allowable values for ListManagementSavedSearchesSortByEnum
const (
	ListManagementSavedSearchesSortByTimecreated ListManagementSavedSearchesSortByEnum = "timeCreated"
	ListManagementSavedSearchesSortByDisplayname ListManagementSavedSearchesSortByEnum = "displayName"
)

var mappingListManagementSavedSearchesSortBy = map[string]ListManagementSavedSearchesSortByEnum{
	"timeCreated": ListManagementSavedSearchesSortByTimecreated,
	"displayName": ListManagementSavedSearchesSortByDisplayname,
}

// GetListManagementSavedSearchesSortByEnumValues Enumerates the set of values for ListManagementSavedSearchesSortByEnum
func GetListManagementSavedSearchesSortByEnumValues() []ListManagementSavedSearchesSortByEnum {
	values := make([]ListManagementSavedSearchesSortByEnum, 0)
	for _, v := range mappingListManagementSavedSearchesSortBy {
		values = append(values, v)
	}
	return values
}
