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

// ListQueryWorkRequestsRequest wrapper for the ListQueryWorkRequests operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListQueryWorkRequests.go.html to see an example of how to use ListQueryWorkRequestsRequest.
type ListQueryWorkRequestsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter based on job execution mode
	Mode ListQueryWorkRequestsModeEnum `mandatory:"false" contributesTo:"query" name:"mode" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListQueryWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeStarted is descending. If no value is specified timeStarted is default.
	SortBy ListQueryWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListQueryWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListQueryWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListQueryWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListQueryWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListQueryWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListQueryWorkRequestsModeEnum(string(request.Mode)); !ok && request.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", request.Mode, strings.Join(GetListQueryWorkRequestsModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListQueryWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListQueryWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListQueryWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListQueryWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListQueryWorkRequestsResponse wrapper for the ListQueryWorkRequests operation
type ListQueryWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of QueryWorkRequestCollection instances
	QueryWorkRequestCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListQueryWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListQueryWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListQueryWorkRequestsModeEnum Enum with underlying type: string
type ListQueryWorkRequestsModeEnum string

// Set of constants representing the allowable values for ListQueryWorkRequestsModeEnum
const (
	ListQueryWorkRequestsModeAll        ListQueryWorkRequestsModeEnum = "ALL"
	ListQueryWorkRequestsModeForeground ListQueryWorkRequestsModeEnum = "FOREGROUND"
	ListQueryWorkRequestsModeBackground ListQueryWorkRequestsModeEnum = "BACKGROUND"
)

var mappingListQueryWorkRequestsModeEnum = map[string]ListQueryWorkRequestsModeEnum{
	"ALL":        ListQueryWorkRequestsModeAll,
	"FOREGROUND": ListQueryWorkRequestsModeForeground,
	"BACKGROUND": ListQueryWorkRequestsModeBackground,
}

// GetListQueryWorkRequestsModeEnumValues Enumerates the set of values for ListQueryWorkRequestsModeEnum
func GetListQueryWorkRequestsModeEnumValues() []ListQueryWorkRequestsModeEnum {
	values := make([]ListQueryWorkRequestsModeEnum, 0)
	for _, v := range mappingListQueryWorkRequestsModeEnum {
		values = append(values, v)
	}
	return values
}

// GetListQueryWorkRequestsModeEnumStringValues Enumerates the set of values in String for ListQueryWorkRequestsModeEnum
func GetListQueryWorkRequestsModeEnumStringValues() []string {
	return []string{
		"ALL",
		"FOREGROUND",
		"BACKGROUND",
	}
}

// GetMappingListQueryWorkRequestsModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQueryWorkRequestsModeEnum(val string) (ListQueryWorkRequestsModeEnum, bool) {
	mappingListQueryWorkRequestsModeEnumIgnoreCase := make(map[string]ListQueryWorkRequestsModeEnum)
	for k, v := range mappingListQueryWorkRequestsModeEnum {
		mappingListQueryWorkRequestsModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListQueryWorkRequestsModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListQueryWorkRequestsSortOrderEnum Enum with underlying type: string
type ListQueryWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListQueryWorkRequestsSortOrderEnum
const (
	ListQueryWorkRequestsSortOrderAsc  ListQueryWorkRequestsSortOrderEnum = "ASC"
	ListQueryWorkRequestsSortOrderDesc ListQueryWorkRequestsSortOrderEnum = "DESC"
)

var mappingListQueryWorkRequestsSortOrderEnum = map[string]ListQueryWorkRequestsSortOrderEnum{
	"ASC":  ListQueryWorkRequestsSortOrderAsc,
	"DESC": ListQueryWorkRequestsSortOrderDesc,
}

// GetListQueryWorkRequestsSortOrderEnumValues Enumerates the set of values for ListQueryWorkRequestsSortOrderEnum
func GetListQueryWorkRequestsSortOrderEnumValues() []ListQueryWorkRequestsSortOrderEnum {
	values := make([]ListQueryWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListQueryWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListQueryWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListQueryWorkRequestsSortOrderEnum
func GetListQueryWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListQueryWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQueryWorkRequestsSortOrderEnum(val string) (ListQueryWorkRequestsSortOrderEnum, bool) {
	mappingListQueryWorkRequestsSortOrderEnumIgnoreCase := make(map[string]ListQueryWorkRequestsSortOrderEnum)
	for k, v := range mappingListQueryWorkRequestsSortOrderEnum {
		mappingListQueryWorkRequestsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListQueryWorkRequestsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListQueryWorkRequestsSortByEnum Enum with underlying type: string
type ListQueryWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListQueryWorkRequestsSortByEnum
const (
	ListQueryWorkRequestsSortByTimestarted ListQueryWorkRequestsSortByEnum = "timeStarted"
	ListQueryWorkRequestsSortByTimeexpires ListQueryWorkRequestsSortByEnum = "timeExpires"
)

var mappingListQueryWorkRequestsSortByEnum = map[string]ListQueryWorkRequestsSortByEnum{
	"timeStarted": ListQueryWorkRequestsSortByTimestarted,
	"timeExpires": ListQueryWorkRequestsSortByTimeexpires,
}

// GetListQueryWorkRequestsSortByEnumValues Enumerates the set of values for ListQueryWorkRequestsSortByEnum
func GetListQueryWorkRequestsSortByEnumValues() []ListQueryWorkRequestsSortByEnum {
	values := make([]ListQueryWorkRequestsSortByEnum, 0)
	for _, v := range mappingListQueryWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListQueryWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListQueryWorkRequestsSortByEnum
func GetListQueryWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeStarted",
		"timeExpires",
	}
}

// GetMappingListQueryWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQueryWorkRequestsSortByEnum(val string) (ListQueryWorkRequestsSortByEnum, bool) {
	mappingListQueryWorkRequestsSortByEnumIgnoreCase := make(map[string]ListQueryWorkRequestsSortByEnum)
	for k, v := range mappingListQueryWorkRequestsSortByEnum {
		mappingListQueryWorkRequestsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListQueryWorkRequestsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
