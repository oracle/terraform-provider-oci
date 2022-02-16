// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCategoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCategoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCategoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCategoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCategoriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListCategoriesSortOrderEnum = map[string]ListCategoriesSortOrderEnum{
	"ASC":  ListCategoriesSortOrderAsc,
	"DESC": ListCategoriesSortOrderDesc,
}

// GetListCategoriesSortOrderEnumValues Enumerates the set of values for ListCategoriesSortOrderEnum
func GetListCategoriesSortOrderEnumValues() []ListCategoriesSortOrderEnum {
	values := make([]ListCategoriesSortOrderEnum, 0)
	for _, v := range mappingListCategoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCategoriesSortOrderEnumStringValues Enumerates the set of values in String for ListCategoriesSortOrderEnum
func GetListCategoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCategoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCategoriesSortOrderEnum(val string) (ListCategoriesSortOrderEnum, bool) {
	mappingListCategoriesSortOrderEnumIgnoreCase := make(map[string]ListCategoriesSortOrderEnum)
	for k, v := range mappingListCategoriesSortOrderEnum {
		mappingListCategoriesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCategoriesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListCategoriesSortByEnum Enum with underlying type: string
type ListCategoriesSortByEnum string

// Set of constants representing the allowable values for ListCategoriesSortByEnum
const (
	ListCategoriesSortByDisplayname ListCategoriesSortByEnum = "displayName"
	ListCategoriesSortByType        ListCategoriesSortByEnum = "type"
)

var mappingListCategoriesSortByEnum = map[string]ListCategoriesSortByEnum{
	"displayName": ListCategoriesSortByDisplayname,
	"type":        ListCategoriesSortByType,
}

// GetListCategoriesSortByEnumValues Enumerates the set of values for ListCategoriesSortByEnum
func GetListCategoriesSortByEnumValues() []ListCategoriesSortByEnum {
	values := make([]ListCategoriesSortByEnum, 0)
	for _, v := range mappingListCategoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCategoriesSortByEnumStringValues Enumerates the set of values in String for ListCategoriesSortByEnum
func GetListCategoriesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"type",
	}
}

// GetMappingListCategoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCategoriesSortByEnum(val string) (ListCategoriesSortByEnum, bool) {
	mappingListCategoriesSortByEnumIgnoreCase := make(map[string]ListCategoriesSortByEnum)
	for k, v := range mappingListCategoriesSortByEnum {
		mappingListCategoriesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListCategoriesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
