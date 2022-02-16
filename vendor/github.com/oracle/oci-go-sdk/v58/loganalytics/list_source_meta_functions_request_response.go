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

// ListSourceMetaFunctionsRequest wrapper for the ListSourceMetaFunctions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourceMetaFunctions.go.html to see an example of how to use ListSourceMetaFunctionsRequest.
type ListSourceMetaFunctionsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The attribute used to sort the returned items
	SortBy ListSourceMetaFunctionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSourceMetaFunctionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSourceMetaFunctionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSourceMetaFunctionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSourceMetaFunctionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSourceMetaFunctionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSourceMetaFunctionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSourceMetaFunctionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSourceMetaFunctionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSourceMetaFunctionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSourceMetaFunctionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSourceMetaFunctionsResponse wrapper for the ListSourceMetaFunctions operation
type ListSourceMetaFunctionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsMetaFunctionCollection instances
	LogAnalyticsMetaFunctionCollection `presentIn:"body"`

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

func (response ListSourceMetaFunctionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSourceMetaFunctionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSourceMetaFunctionsSortByEnum Enum with underlying type: string
type ListSourceMetaFunctionsSortByEnum string

// Set of constants representing the allowable values for ListSourceMetaFunctionsSortByEnum
const (
	ListSourceMetaFunctionsSortByName ListSourceMetaFunctionsSortByEnum = "name"
)

var mappingListSourceMetaFunctionsSortByEnum = map[string]ListSourceMetaFunctionsSortByEnum{
	"name": ListSourceMetaFunctionsSortByName,
}

// GetListSourceMetaFunctionsSortByEnumValues Enumerates the set of values for ListSourceMetaFunctionsSortByEnum
func GetListSourceMetaFunctionsSortByEnumValues() []ListSourceMetaFunctionsSortByEnum {
	values := make([]ListSourceMetaFunctionsSortByEnum, 0)
	for _, v := range mappingListSourceMetaFunctionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourceMetaFunctionsSortByEnumStringValues Enumerates the set of values in String for ListSourceMetaFunctionsSortByEnum
func GetListSourceMetaFunctionsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListSourceMetaFunctionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourceMetaFunctionsSortByEnum(val string) (ListSourceMetaFunctionsSortByEnum, bool) {
	mappingListSourceMetaFunctionsSortByEnumIgnoreCase := make(map[string]ListSourceMetaFunctionsSortByEnum)
	for k, v := range mappingListSourceMetaFunctionsSortByEnum {
		mappingListSourceMetaFunctionsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSourceMetaFunctionsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSourceMetaFunctionsSortOrderEnum Enum with underlying type: string
type ListSourceMetaFunctionsSortOrderEnum string

// Set of constants representing the allowable values for ListSourceMetaFunctionsSortOrderEnum
const (
	ListSourceMetaFunctionsSortOrderAsc  ListSourceMetaFunctionsSortOrderEnum = "ASC"
	ListSourceMetaFunctionsSortOrderDesc ListSourceMetaFunctionsSortOrderEnum = "DESC"
)

var mappingListSourceMetaFunctionsSortOrderEnum = map[string]ListSourceMetaFunctionsSortOrderEnum{
	"ASC":  ListSourceMetaFunctionsSortOrderAsc,
	"DESC": ListSourceMetaFunctionsSortOrderDesc,
}

// GetListSourceMetaFunctionsSortOrderEnumValues Enumerates the set of values for ListSourceMetaFunctionsSortOrderEnum
func GetListSourceMetaFunctionsSortOrderEnumValues() []ListSourceMetaFunctionsSortOrderEnum {
	values := make([]ListSourceMetaFunctionsSortOrderEnum, 0)
	for _, v := range mappingListSourceMetaFunctionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSourceMetaFunctionsSortOrderEnumStringValues Enumerates the set of values in String for ListSourceMetaFunctionsSortOrderEnum
func GetListSourceMetaFunctionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSourceMetaFunctionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSourceMetaFunctionsSortOrderEnum(val string) (ListSourceMetaFunctionsSortOrderEnum, bool) {
	mappingListSourceMetaFunctionsSortOrderEnumIgnoreCase := make(map[string]ListSourceMetaFunctionsSortOrderEnum)
	for k, v := range mappingListSourceMetaFunctionsSortOrderEnum {
		mappingListSourceMetaFunctionsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSourceMetaFunctionsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
