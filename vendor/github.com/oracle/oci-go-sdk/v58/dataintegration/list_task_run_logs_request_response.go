// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListTaskRunLogsRequest wrapper for the ListTaskRunLogs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTaskRunLogs.go.html to see an example of how to use ListTaskRunLogsRequest.
type ListTaskRunLogsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// The task run key.
	TaskRunKey *string `mandatory:"true" contributesTo:"path" name:"taskRunKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListTaskRunLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListTaskRunLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTaskRunLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTaskRunLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTaskRunLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTaskRunLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTaskRunLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTaskRunLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTaskRunLogsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTaskRunLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTaskRunLogsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTaskRunLogsResponse wrapper for the ListTaskRunLogs operation
type ListTaskRunLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []TaskRunLogSummary instances
	Items []TaskRunLogSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListTaskRunLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTaskRunLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTaskRunLogsSortOrderEnum Enum with underlying type: string
type ListTaskRunLogsSortOrderEnum string

// Set of constants representing the allowable values for ListTaskRunLogsSortOrderEnum
const (
	ListTaskRunLogsSortOrderAsc  ListTaskRunLogsSortOrderEnum = "ASC"
	ListTaskRunLogsSortOrderDesc ListTaskRunLogsSortOrderEnum = "DESC"
)

var mappingListTaskRunLogsSortOrderEnum = map[string]ListTaskRunLogsSortOrderEnum{
	"ASC":  ListTaskRunLogsSortOrderAsc,
	"DESC": ListTaskRunLogsSortOrderDesc,
}

// GetListTaskRunLogsSortOrderEnumValues Enumerates the set of values for ListTaskRunLogsSortOrderEnum
func GetListTaskRunLogsSortOrderEnumValues() []ListTaskRunLogsSortOrderEnum {
	values := make([]ListTaskRunLogsSortOrderEnum, 0)
	for _, v := range mappingListTaskRunLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskRunLogsSortOrderEnumStringValues Enumerates the set of values in String for ListTaskRunLogsSortOrderEnum
func GetListTaskRunLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTaskRunLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskRunLogsSortOrderEnum(val string) (ListTaskRunLogsSortOrderEnum, bool) {
	mappingListTaskRunLogsSortOrderEnumIgnoreCase := make(map[string]ListTaskRunLogsSortOrderEnum)
	for k, v := range mappingListTaskRunLogsSortOrderEnum {
		mappingListTaskRunLogsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTaskRunLogsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListTaskRunLogsSortByEnum Enum with underlying type: string
type ListTaskRunLogsSortByEnum string

// Set of constants representing the allowable values for ListTaskRunLogsSortByEnum
const (
	ListTaskRunLogsSortByTimeCreated ListTaskRunLogsSortByEnum = "TIME_CREATED"
	ListTaskRunLogsSortByDisplayName ListTaskRunLogsSortByEnum = "DISPLAY_NAME"
)

var mappingListTaskRunLogsSortByEnum = map[string]ListTaskRunLogsSortByEnum{
	"TIME_CREATED": ListTaskRunLogsSortByTimeCreated,
	"DISPLAY_NAME": ListTaskRunLogsSortByDisplayName,
}

// GetListTaskRunLogsSortByEnumValues Enumerates the set of values for ListTaskRunLogsSortByEnum
func GetListTaskRunLogsSortByEnumValues() []ListTaskRunLogsSortByEnum {
	values := make([]ListTaskRunLogsSortByEnum, 0)
	for _, v := range mappingListTaskRunLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskRunLogsSortByEnumStringValues Enumerates the set of values in String for ListTaskRunLogsSortByEnum
func GetListTaskRunLogsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListTaskRunLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskRunLogsSortByEnum(val string) (ListTaskRunLogsSortByEnum, bool) {
	mappingListTaskRunLogsSortByEnumIgnoreCase := make(map[string]ListTaskRunLogsSortByEnum)
	for k, v := range mappingListTaskRunLogsSortByEnum {
		mappingListTaskRunLogsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTaskRunLogsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
