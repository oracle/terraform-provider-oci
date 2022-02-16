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

// ListMetaSourceTypesRequest wrapper for the ListMetaSourceTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListMetaSourceTypes.go.html to see an example of how to use ListMetaSourceTypesRequest.
type ListMetaSourceTypesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The attribute used to sort the returned items
	SortBy ListMetaSourceTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListMetaSourceTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMetaSourceTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMetaSourceTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMetaSourceTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMetaSourceTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMetaSourceTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMetaSourceTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMetaSourceTypesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMetaSourceTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMetaSourceTypesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMetaSourceTypesResponse wrapper for the ListMetaSourceTypes operation
type ListMetaSourceTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsMetaSourceTypeCollection instances
	LogAnalyticsMetaSourceTypeCollection `presentIn:"body"`

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

func (response ListMetaSourceTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMetaSourceTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMetaSourceTypesSortByEnum Enum with underlying type: string
type ListMetaSourceTypesSortByEnum string

// Set of constants representing the allowable values for ListMetaSourceTypesSortByEnum
const (
	ListMetaSourceTypesSortByName ListMetaSourceTypesSortByEnum = "name"
)

var mappingListMetaSourceTypesSortByEnum = map[string]ListMetaSourceTypesSortByEnum{
	"name": ListMetaSourceTypesSortByName,
}

// GetListMetaSourceTypesSortByEnumValues Enumerates the set of values for ListMetaSourceTypesSortByEnum
func GetListMetaSourceTypesSortByEnumValues() []ListMetaSourceTypesSortByEnum {
	values := make([]ListMetaSourceTypesSortByEnum, 0)
	for _, v := range mappingListMetaSourceTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetaSourceTypesSortByEnumStringValues Enumerates the set of values in String for ListMetaSourceTypesSortByEnum
func GetListMetaSourceTypesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListMetaSourceTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetaSourceTypesSortByEnum(val string) (ListMetaSourceTypesSortByEnum, bool) {
	mappingListMetaSourceTypesSortByEnumIgnoreCase := make(map[string]ListMetaSourceTypesSortByEnum)
	for k, v := range mappingListMetaSourceTypesSortByEnum {
		mappingListMetaSourceTypesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMetaSourceTypesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListMetaSourceTypesSortOrderEnum Enum with underlying type: string
type ListMetaSourceTypesSortOrderEnum string

// Set of constants representing the allowable values for ListMetaSourceTypesSortOrderEnum
const (
	ListMetaSourceTypesSortOrderAsc  ListMetaSourceTypesSortOrderEnum = "ASC"
	ListMetaSourceTypesSortOrderDesc ListMetaSourceTypesSortOrderEnum = "DESC"
)

var mappingListMetaSourceTypesSortOrderEnum = map[string]ListMetaSourceTypesSortOrderEnum{
	"ASC":  ListMetaSourceTypesSortOrderAsc,
	"DESC": ListMetaSourceTypesSortOrderDesc,
}

// GetListMetaSourceTypesSortOrderEnumValues Enumerates the set of values for ListMetaSourceTypesSortOrderEnum
func GetListMetaSourceTypesSortOrderEnumValues() []ListMetaSourceTypesSortOrderEnum {
	values := make([]ListMetaSourceTypesSortOrderEnum, 0)
	for _, v := range mappingListMetaSourceTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMetaSourceTypesSortOrderEnumStringValues Enumerates the set of values in String for ListMetaSourceTypesSortOrderEnum
func GetListMetaSourceTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMetaSourceTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMetaSourceTypesSortOrderEnum(val string) (ListMetaSourceTypesSortOrderEnum, bool) {
	mappingListMetaSourceTypesSortOrderEnumIgnoreCase := make(map[string]ListMetaSourceTypesSortOrderEnum)
	for k, v := range mappingListMetaSourceTypesSortOrderEnum {
		mappingListMetaSourceTypesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMetaSourceTypesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
