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

// ListResourceCategoriesRequest wrapper for the ListResourceCategories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListResourceCategories.go.html to see an example of how to use ListResourceCategoriesRequest.
type ListResourceCategoriesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// A comma-separated list of categories used for filtering
	Categories *string `mandatory:"false" contributesTo:"query" name:"categories"`

	// A comma-separated list of resource types used for filtering. Only resources of the types
	// specified will be returned. Examples include SOURCE, PARSER, LOOKUP, etc.
	ResourceTypes *string `mandatory:"false" contributesTo:"query" name:"resourceTypes"`

	// A comma-separated list of resource unique identifiers used for filtering. Only resources
	// with matching unique identifiers will be returned.
	ResourceIds *string `mandatory:"false" contributesTo:"query" name:"resourceIds"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListResourceCategoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned category resources.
	SortBy ListResourceCategoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceCategoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceCategoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceCategoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceCategoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResourceCategoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResourceCategoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResourceCategoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResourceCategoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResourceCategoriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResourceCategoriesResponse wrapper for the ListResourceCategories operation
type ListResourceCategoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsResourceCategoryCollection instances
	LogAnalyticsResourceCategoryCollection `presentIn:"body"`

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

func (response ListResourceCategoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceCategoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceCategoriesSortOrderEnum Enum with underlying type: string
type ListResourceCategoriesSortOrderEnum string

// Set of constants representing the allowable values for ListResourceCategoriesSortOrderEnum
const (
	ListResourceCategoriesSortOrderAsc  ListResourceCategoriesSortOrderEnum = "ASC"
	ListResourceCategoriesSortOrderDesc ListResourceCategoriesSortOrderEnum = "DESC"
)

var mappingListResourceCategoriesSortOrderEnum = map[string]ListResourceCategoriesSortOrderEnum{
	"ASC":  ListResourceCategoriesSortOrderAsc,
	"DESC": ListResourceCategoriesSortOrderDesc,
}

// GetListResourceCategoriesSortOrderEnumValues Enumerates the set of values for ListResourceCategoriesSortOrderEnum
func GetListResourceCategoriesSortOrderEnumValues() []ListResourceCategoriesSortOrderEnum {
	values := make([]ListResourceCategoriesSortOrderEnum, 0)
	for _, v := range mappingListResourceCategoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceCategoriesSortOrderEnumStringValues Enumerates the set of values in String for ListResourceCategoriesSortOrderEnum
func GetListResourceCategoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResourceCategoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceCategoriesSortOrderEnum(val string) (ListResourceCategoriesSortOrderEnum, bool) {
	mappingListResourceCategoriesSortOrderEnumIgnoreCase := make(map[string]ListResourceCategoriesSortOrderEnum)
	for k, v := range mappingListResourceCategoriesSortOrderEnum {
		mappingListResourceCategoriesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListResourceCategoriesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListResourceCategoriesSortByEnum Enum with underlying type: string
type ListResourceCategoriesSortByEnum string

// Set of constants representing the allowable values for ListResourceCategoriesSortByEnum
const (
	ListResourceCategoriesSortByResourcetype ListResourceCategoriesSortByEnum = "resourceType"
	ListResourceCategoriesSortByCategoryname ListResourceCategoriesSortByEnum = "categoryName"
	ListResourceCategoriesSortByResourceid   ListResourceCategoriesSortByEnum = "resourceId"
)

var mappingListResourceCategoriesSortByEnum = map[string]ListResourceCategoriesSortByEnum{
	"resourceType": ListResourceCategoriesSortByResourcetype,
	"categoryName": ListResourceCategoriesSortByCategoryname,
	"resourceId":   ListResourceCategoriesSortByResourceid,
}

// GetListResourceCategoriesSortByEnumValues Enumerates the set of values for ListResourceCategoriesSortByEnum
func GetListResourceCategoriesSortByEnumValues() []ListResourceCategoriesSortByEnum {
	values := make([]ListResourceCategoriesSortByEnum, 0)
	for _, v := range mappingListResourceCategoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResourceCategoriesSortByEnumStringValues Enumerates the set of values in String for ListResourceCategoriesSortByEnum
func GetListResourceCategoriesSortByEnumStringValues() []string {
	return []string{
		"resourceType",
		"categoryName",
		"resourceId",
	}
}

// GetMappingListResourceCategoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResourceCategoriesSortByEnum(val string) (ListResourceCategoriesSortByEnum, bool) {
	mappingListResourceCategoriesSortByEnumIgnoreCase := make(map[string]ListResourceCategoriesSortByEnum)
	for k, v := range mappingListResourceCategoriesSortByEnum {
		mappingListResourceCategoriesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListResourceCategoriesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
