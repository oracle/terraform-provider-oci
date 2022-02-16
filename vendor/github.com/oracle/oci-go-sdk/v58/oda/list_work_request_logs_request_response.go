// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListWorkRequestLogsRequest wrapper for the ListWorkRequestLogs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListWorkRequestLogs.go.html to see an example of how to use ListWorkRequestLogsRequest.
type ListWorkRequestLogsRequest struct {

	// The identifier of the asynchronous work request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. You can specify only one sort order. If no value is specified, then the default is `TIMESTAMP`.
	// The default sort order for both `TIMESTAMP` and `MESSAGE` is ascending.
	SortBy ListWorkRequestLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListWorkRequestLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkRequestLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkRequestLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkRequestLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkRequestLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkRequestLogsResponse wrapper for the ListWorkRequestLogs operation
type ListWorkRequestLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkRequestLogEntry instances
	Items []WorkRequestLogEntry `presentIn:"body"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you contact Oracle
	// about this request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListWorkRequestLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestLogsSortByEnum Enum with underlying type: string
type ListWorkRequestLogsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestLogsSortByEnum
const (
	ListWorkRequestLogsSortByMessage   ListWorkRequestLogsSortByEnum = "MESSAGE"
	ListWorkRequestLogsSortByTimestamp ListWorkRequestLogsSortByEnum = "TIMESTAMP"
)

var mappingListWorkRequestLogsSortByEnum = map[string]ListWorkRequestLogsSortByEnum{
	"MESSAGE":   ListWorkRequestLogsSortByMessage,
	"TIMESTAMP": ListWorkRequestLogsSortByTimestamp,
}

// GetListWorkRequestLogsSortByEnumValues Enumerates the set of values for ListWorkRequestLogsSortByEnum
func GetListWorkRequestLogsSortByEnumValues() []ListWorkRequestLogsSortByEnum {
	values := make([]ListWorkRequestLogsSortByEnum, 0)
	for _, v := range mappingListWorkRequestLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestLogsSortByEnumStringValues Enumerates the set of values in String for ListWorkRequestLogsSortByEnum
func GetListWorkRequestLogsSortByEnumStringValues() []string {
	return []string{
		"MESSAGE",
		"TIMESTAMP",
	}
}

// GetMappingListWorkRequestLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestLogsSortByEnum(val string) (ListWorkRequestLogsSortByEnum, bool) {
	mappingListWorkRequestLogsSortByEnumIgnoreCase := make(map[string]ListWorkRequestLogsSortByEnum)
	for k, v := range mappingListWorkRequestLogsSortByEnum {
		mappingListWorkRequestLogsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWorkRequestLogsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestLogsSortOrderEnum Enum with underlying type: string
type ListWorkRequestLogsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestLogsSortOrderEnum
const (
	ListWorkRequestLogsSortOrderAsc  ListWorkRequestLogsSortOrderEnum = "ASC"
	ListWorkRequestLogsSortOrderDesc ListWorkRequestLogsSortOrderEnum = "DESC"
)

var mappingListWorkRequestLogsSortOrderEnum = map[string]ListWorkRequestLogsSortOrderEnum{
	"ASC":  ListWorkRequestLogsSortOrderAsc,
	"DESC": ListWorkRequestLogsSortOrderDesc,
}

// GetListWorkRequestLogsSortOrderEnumValues Enumerates the set of values for ListWorkRequestLogsSortOrderEnum
func GetListWorkRequestLogsSortOrderEnumValues() []ListWorkRequestLogsSortOrderEnum {
	values := make([]ListWorkRequestLogsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestLogsSortOrderEnumStringValues Enumerates the set of values in String for ListWorkRequestLogsSortOrderEnum
func GetListWorkRequestLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkRequestLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestLogsSortOrderEnum(val string) (ListWorkRequestLogsSortOrderEnum, bool) {
	mappingListWorkRequestLogsSortOrderEnumIgnoreCase := make(map[string]ListWorkRequestLogsSortOrderEnum)
	for k, v := range mappingListWorkRequestLogsSortOrderEnum {
		mappingListWorkRequestLogsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWorkRequestLogsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
