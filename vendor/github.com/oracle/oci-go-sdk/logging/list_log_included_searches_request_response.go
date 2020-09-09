// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLogIncludedSearchesRequest wrapper for the ListLogIncludedSearches operation
type ListLogIncludedSearchesRequest struct {

	// Compartment OCID to list resources in. Please see compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// OCID of the LogIncludedSearch
	LogIncludedSearchId *string `mandatory:"false" contributesTo:"query" name:"logIncludedSearchId"`

	// Resource name
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogIncludedSearchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListLogIncludedSearchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogIncludedSearchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogIncludedSearchesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogIncludedSearchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogIncludedSearchesResponse wrapper for the ListLogIncludedSearches operation
type ListLogIncludedSearchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogIncludedSearchSummaryCollection instances
	LogIncludedSearchSummaryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLogIncludedSearchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogIncludedSearchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogIncludedSearchesSortByEnum Enum with underlying type: string
type ListLogIncludedSearchesSortByEnum string

// Set of constants representing the allowable values for ListLogIncludedSearchesSortByEnum
const (
	ListLogIncludedSearchesSortByTimecreated ListLogIncludedSearchesSortByEnum = "timeCreated"
	ListLogIncludedSearchesSortByDisplayname ListLogIncludedSearchesSortByEnum = "displayName"
)

var mappingListLogIncludedSearchesSortBy = map[string]ListLogIncludedSearchesSortByEnum{
	"timeCreated": ListLogIncludedSearchesSortByTimecreated,
	"displayName": ListLogIncludedSearchesSortByDisplayname,
}

// GetListLogIncludedSearchesSortByEnumValues Enumerates the set of values for ListLogIncludedSearchesSortByEnum
func GetListLogIncludedSearchesSortByEnumValues() []ListLogIncludedSearchesSortByEnum {
	values := make([]ListLogIncludedSearchesSortByEnum, 0)
	for _, v := range mappingListLogIncludedSearchesSortBy {
		values = append(values, v)
	}
	return values
}

// ListLogIncludedSearchesSortOrderEnum Enum with underlying type: string
type ListLogIncludedSearchesSortOrderEnum string

// Set of constants representing the allowable values for ListLogIncludedSearchesSortOrderEnum
const (
	ListLogIncludedSearchesSortOrderAsc  ListLogIncludedSearchesSortOrderEnum = "ASC"
	ListLogIncludedSearchesSortOrderDesc ListLogIncludedSearchesSortOrderEnum = "DESC"
)

var mappingListLogIncludedSearchesSortOrder = map[string]ListLogIncludedSearchesSortOrderEnum{
	"ASC":  ListLogIncludedSearchesSortOrderAsc,
	"DESC": ListLogIncludedSearchesSortOrderDesc,
}

// GetListLogIncludedSearchesSortOrderEnumValues Enumerates the set of values for ListLogIncludedSearchesSortOrderEnum
func GetListLogIncludedSearchesSortOrderEnumValues() []ListLogIncludedSearchesSortOrderEnum {
	values := make([]ListLogIncludedSearchesSortOrderEnum, 0)
	for _, v := range mappingListLogIncludedSearchesSortOrder {
		values = append(values, v)
	}
	return values
}
