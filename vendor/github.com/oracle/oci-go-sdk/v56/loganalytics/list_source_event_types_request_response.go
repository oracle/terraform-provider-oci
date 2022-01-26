// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListSourceEventTypesRequest wrapper for the ListSourceEventTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourceEventTypes.go.html to see an example of how to use ListSourceEventTypesRequest.
type ListSourceEventTypesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The source name.
	SourceName *string `mandatory:"true" contributesTo:"path" name:"sourceName"`

	// The text used for filtering event types by name or description.
	DisplayText *string `mandatory:"false" contributesTo:"query" name:"displayText"`

	// The system value used for filtering.  Only items with the specified system value
	// will be returned.  Valid values are built in, custom (for user defined items), or
	// all (for all items, regardless of system value).
	IsSystem ListSourceEventTypesIsSystemEnum `mandatory:"false" contributesTo:"query" name:"isSystem" omitEmpty:"true"`

	// The enabled flag used for filtering.  Only items with the specified enabled value
	// will be returned.
	IsEnabled *bool `mandatory:"false" contributesTo:"query" name:"isEnabled"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The attribute used to sort the returned source event types.
	SortBy ListSourceEventTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSourceEventTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSourceEventTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSourceEventTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSourceEventTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSourceEventTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSourceEventTypesResponse wrapper for the ListSourceEventTypes operation
type ListSourceEventTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EventTypeCollection instances
	EventTypeCollection `presentIn:"body"`

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

func (response ListSourceEventTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSourceEventTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSourceEventTypesIsSystemEnum Enum with underlying type: string
type ListSourceEventTypesIsSystemEnum string

// Set of constants representing the allowable values for ListSourceEventTypesIsSystemEnum
const (
	ListSourceEventTypesIsSystemAll     ListSourceEventTypesIsSystemEnum = "ALL"
	ListSourceEventTypesIsSystemCustom  ListSourceEventTypesIsSystemEnum = "CUSTOM"
	ListSourceEventTypesIsSystemBuiltIn ListSourceEventTypesIsSystemEnum = "BUILT_IN"
)

var mappingListSourceEventTypesIsSystem = map[string]ListSourceEventTypesIsSystemEnum{
	"ALL":      ListSourceEventTypesIsSystemAll,
	"CUSTOM":   ListSourceEventTypesIsSystemCustom,
	"BUILT_IN": ListSourceEventTypesIsSystemBuiltIn,
}

// GetListSourceEventTypesIsSystemEnumValues Enumerates the set of values for ListSourceEventTypesIsSystemEnum
func GetListSourceEventTypesIsSystemEnumValues() []ListSourceEventTypesIsSystemEnum {
	values := make([]ListSourceEventTypesIsSystemEnum, 0)
	for _, v := range mappingListSourceEventTypesIsSystem {
		values = append(values, v)
	}
	return values
}

// ListSourceEventTypesSortByEnum Enum with underlying type: string
type ListSourceEventTypesSortByEnum string

// Set of constants representing the allowable values for ListSourceEventTypesSortByEnum
const (
	ListSourceEventTypesSortByEventtype   ListSourceEventTypesSortByEnum = "eventType"
	ListSourceEventTypesSortByTimeupdated ListSourceEventTypesSortByEnum = "timeUpdated"
)

var mappingListSourceEventTypesSortBy = map[string]ListSourceEventTypesSortByEnum{
	"eventType":   ListSourceEventTypesSortByEventtype,
	"timeUpdated": ListSourceEventTypesSortByTimeupdated,
}

// GetListSourceEventTypesSortByEnumValues Enumerates the set of values for ListSourceEventTypesSortByEnum
func GetListSourceEventTypesSortByEnumValues() []ListSourceEventTypesSortByEnum {
	values := make([]ListSourceEventTypesSortByEnum, 0)
	for _, v := range mappingListSourceEventTypesSortBy {
		values = append(values, v)
	}
	return values
}

// ListSourceEventTypesSortOrderEnum Enum with underlying type: string
type ListSourceEventTypesSortOrderEnum string

// Set of constants representing the allowable values for ListSourceEventTypesSortOrderEnum
const (
	ListSourceEventTypesSortOrderAsc  ListSourceEventTypesSortOrderEnum = "ASC"
	ListSourceEventTypesSortOrderDesc ListSourceEventTypesSortOrderEnum = "DESC"
)

var mappingListSourceEventTypesSortOrder = map[string]ListSourceEventTypesSortOrderEnum{
	"ASC":  ListSourceEventTypesSortOrderAsc,
	"DESC": ListSourceEventTypesSortOrderDesc,
}

// GetListSourceEventTypesSortOrderEnumValues Enumerates the set of values for ListSourceEventTypesSortOrderEnum
func GetListSourceEventTypesSortOrderEnumValues() []ListSourceEventTypesSortOrderEnum {
	values := make([]ListSourceEventTypesSortOrderEnum, 0)
	for _, v := range mappingListSourceEventTypesSortOrder {
		values = append(values, v)
	}
	return values
}
