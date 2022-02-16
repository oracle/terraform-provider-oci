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

// ListParserFunctionsRequest wrapper for the ListParserFunctions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListParserFunctions.go.html to see an example of how to use ListParserFunctionsRequest.
type ListParserFunctionsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The parser name used for filtering.
	ParserName *string `mandatory:"false" contributesTo:"query" name:"parserName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The attribute used to sort the returned items
	SortBy ListParserFunctionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListParserFunctionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListParserFunctionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListParserFunctionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListParserFunctionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListParserFunctionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListParserFunctionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListParserFunctionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListParserFunctionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListParserFunctionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListParserFunctionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListParserFunctionsResponse wrapper for the ListParserFunctions operation
type ListParserFunctionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsParserFunctionCollection instances
	LogAnalyticsParserFunctionCollection `presentIn:"body"`

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

func (response ListParserFunctionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListParserFunctionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListParserFunctionsSortByEnum Enum with underlying type: string
type ListParserFunctionsSortByEnum string

// Set of constants representing the allowable values for ListParserFunctionsSortByEnum
const (
	ListParserFunctionsSortByName ListParserFunctionsSortByEnum = "name"
)

var mappingListParserFunctionsSortByEnum = map[string]ListParserFunctionsSortByEnum{
	"name": ListParserFunctionsSortByName,
}

// GetListParserFunctionsSortByEnumValues Enumerates the set of values for ListParserFunctionsSortByEnum
func GetListParserFunctionsSortByEnumValues() []ListParserFunctionsSortByEnum {
	values := make([]ListParserFunctionsSortByEnum, 0)
	for _, v := range mappingListParserFunctionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListParserFunctionsSortByEnumStringValues Enumerates the set of values in String for ListParserFunctionsSortByEnum
func GetListParserFunctionsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListParserFunctionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParserFunctionsSortByEnum(val string) (ListParserFunctionsSortByEnum, bool) {
	mappingListParserFunctionsSortByEnumIgnoreCase := make(map[string]ListParserFunctionsSortByEnum)
	for k, v := range mappingListParserFunctionsSortByEnum {
		mappingListParserFunctionsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListParserFunctionsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListParserFunctionsSortOrderEnum Enum with underlying type: string
type ListParserFunctionsSortOrderEnum string

// Set of constants representing the allowable values for ListParserFunctionsSortOrderEnum
const (
	ListParserFunctionsSortOrderAsc  ListParserFunctionsSortOrderEnum = "ASC"
	ListParserFunctionsSortOrderDesc ListParserFunctionsSortOrderEnum = "DESC"
)

var mappingListParserFunctionsSortOrderEnum = map[string]ListParserFunctionsSortOrderEnum{
	"ASC":  ListParserFunctionsSortOrderAsc,
	"DESC": ListParserFunctionsSortOrderDesc,
}

// GetListParserFunctionsSortOrderEnumValues Enumerates the set of values for ListParserFunctionsSortOrderEnum
func GetListParserFunctionsSortOrderEnumValues() []ListParserFunctionsSortOrderEnum {
	values := make([]ListParserFunctionsSortOrderEnum, 0)
	for _, v := range mappingListParserFunctionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListParserFunctionsSortOrderEnumStringValues Enumerates the set of values in String for ListParserFunctionsSortOrderEnum
func GetListParserFunctionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListParserFunctionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParserFunctionsSortOrderEnum(val string) (ListParserFunctionsSortOrderEnum, bool) {
	mappingListParserFunctionsSortOrderEnumIgnoreCase := make(map[string]ListParserFunctionsSortOrderEnum)
	for k, v := range mappingListParserFunctionsSortOrderEnum {
		mappingListParserFunctionsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListParserFunctionsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
