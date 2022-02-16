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

// ListParserMetaPluginsRequest wrapper for the ListParserMetaPlugins operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListParserMetaPlugins.go.html to see an example of how to use ListParserMetaPluginsRequest.
type ListParserMetaPluginsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The attribute used to sort the returned items
	SortBy ListParserMetaPluginsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListParserMetaPluginsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListParserMetaPluginsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListParserMetaPluginsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListParserMetaPluginsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListParserMetaPluginsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListParserMetaPluginsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListParserMetaPluginsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListParserMetaPluginsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListParserMetaPluginsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListParserMetaPluginsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListParserMetaPluginsResponse wrapper for the ListParserMetaPlugins operation
type ListParserMetaPluginsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsParserMetaPluginCollection instances
	LogAnalyticsParserMetaPluginCollection `presentIn:"body"`

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

func (response ListParserMetaPluginsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListParserMetaPluginsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListParserMetaPluginsSortByEnum Enum with underlying type: string
type ListParserMetaPluginsSortByEnum string

// Set of constants representing the allowable values for ListParserMetaPluginsSortByEnum
const (
	ListParserMetaPluginsSortByName ListParserMetaPluginsSortByEnum = "name"
)

var mappingListParserMetaPluginsSortByEnum = map[string]ListParserMetaPluginsSortByEnum{
	"name": ListParserMetaPluginsSortByName,
}

// GetListParserMetaPluginsSortByEnumValues Enumerates the set of values for ListParserMetaPluginsSortByEnum
func GetListParserMetaPluginsSortByEnumValues() []ListParserMetaPluginsSortByEnum {
	values := make([]ListParserMetaPluginsSortByEnum, 0)
	for _, v := range mappingListParserMetaPluginsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListParserMetaPluginsSortByEnumStringValues Enumerates the set of values in String for ListParserMetaPluginsSortByEnum
func GetListParserMetaPluginsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListParserMetaPluginsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParserMetaPluginsSortByEnum(val string) (ListParserMetaPluginsSortByEnum, bool) {
	mappingListParserMetaPluginsSortByEnumIgnoreCase := make(map[string]ListParserMetaPluginsSortByEnum)
	for k, v := range mappingListParserMetaPluginsSortByEnum {
		mappingListParserMetaPluginsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListParserMetaPluginsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListParserMetaPluginsSortOrderEnum Enum with underlying type: string
type ListParserMetaPluginsSortOrderEnum string

// Set of constants representing the allowable values for ListParserMetaPluginsSortOrderEnum
const (
	ListParserMetaPluginsSortOrderAsc  ListParserMetaPluginsSortOrderEnum = "ASC"
	ListParserMetaPluginsSortOrderDesc ListParserMetaPluginsSortOrderEnum = "DESC"
)

var mappingListParserMetaPluginsSortOrderEnum = map[string]ListParserMetaPluginsSortOrderEnum{
	"ASC":  ListParserMetaPluginsSortOrderAsc,
	"DESC": ListParserMetaPluginsSortOrderDesc,
}

// GetListParserMetaPluginsSortOrderEnumValues Enumerates the set of values for ListParserMetaPluginsSortOrderEnum
func GetListParserMetaPluginsSortOrderEnumValues() []ListParserMetaPluginsSortOrderEnum {
	values := make([]ListParserMetaPluginsSortOrderEnum, 0)
	for _, v := range mappingListParserMetaPluginsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListParserMetaPluginsSortOrderEnumStringValues Enumerates the set of values in String for ListParserMetaPluginsSortOrderEnum
func GetListParserMetaPluginsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListParserMetaPluginsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListParserMetaPluginsSortOrderEnum(val string) (ListParserMetaPluginsSortOrderEnum, bool) {
	mappingListParserMetaPluginsSortOrderEnumIgnoreCase := make(map[string]ListParserMetaPluginsSortOrderEnum)
	for k, v := range mappingListParserMetaPluginsSortOrderEnum {
		mappingListParserMetaPluginsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListParserMetaPluginsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
