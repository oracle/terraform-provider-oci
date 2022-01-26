// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListCategoriesRequest wrapper for the ListCategories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListCategories.go.html to see an example of how to use ListCategoriesRequest.
type ListCategoriesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// A comma-separated list of category types used for filtering. Only categories of the
	// specified types will be returned.
	CategoryType *string `mandatory:"false" contributesTo:"query" name:"categoryType"`

	// The category display text used for filtering. Only categories matching the specified display
	// name or description will be returned.
	CategoryDisplayText *string `mandatory:"false" contributesTo:"query" name:"categoryDisplayText"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCategoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned categories
	SortBy ListCategoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only log analytics entities whose name matches the entire name given. The match
	// is case-insensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCategoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCategoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCategoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCategoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCategoriesResponse wrapper for the ListCategories operation
type ListCategoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsCategoryCollection instances
	LogAnalyticsCategoryCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListCategoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCategoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCategoriesSortOrderEnum Enum with underlying type: string
type ListCategoriesSortOrderEnum string

// Set of constants representing the allowable values for ListCategoriesSortOrderEnum
const (
	ListCategoriesSortOrderAsc  ListCategoriesSortOrderEnum = "ASC"
	ListCategoriesSortOrderDesc ListCategoriesSortOrderEnum = "DESC"
)

var mappingListCategoriesSortOrder = map[string]ListCategoriesSortOrderEnum{
	"ASC":  ListCategoriesSortOrderAsc,
	"DESC": ListCategoriesSortOrderDesc,
}

// GetListCategoriesSortOrderEnumValues Enumerates the set of values for ListCategoriesSortOrderEnum
func GetListCategoriesSortOrderEnumValues() []ListCategoriesSortOrderEnum {
	values := make([]ListCategoriesSortOrderEnum, 0)
	for _, v := range mappingListCategoriesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListCategoriesSortByEnum Enum with underlying type: string
type ListCategoriesSortByEnum string

// Set of constants representing the allowable values for ListCategoriesSortByEnum
const (
	ListCategoriesSortByDisplayname ListCategoriesSortByEnum = "displayName"
	ListCategoriesSortByType        ListCategoriesSortByEnum = "type"
)

var mappingListCategoriesSortBy = map[string]ListCategoriesSortByEnum{
	"displayName": ListCategoriesSortByDisplayname,
	"type":        ListCategoriesSortByType,
}

// GetListCategoriesSortByEnumValues Enumerates the set of values for ListCategoriesSortByEnum
func GetListCategoriesSortByEnumValues() []ListCategoriesSortByEnum {
	values := make([]ListCategoriesSortByEnum, 0)
	for _, v := range mappingListCategoriesSortBy {
		values = append(values, v)
	}
	return values
}
