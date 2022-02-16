// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListWorkRequests.go.html to see an example of how to use ListWorkRequestsRequest.
type ListWorkRequestsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the managed instance for which to list resources.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OS family for which to list resources.
	OsFamily ListWorkRequestsOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkRequestsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListWorkRequestsOsFamilyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkRequestsResponse wrapper for the ListWorkRequests operation
type ListWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkRequestSummary instances
	Items []WorkRequestSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsSortOrderEnum Enum with underlying type: string
type ListWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortOrderEnum
const (
	ListWorkRequestsSortOrderAsc  ListWorkRequestsSortOrderEnum = "ASC"
	ListWorkRequestsSortOrderDesc ListWorkRequestsSortOrderEnum = "DESC"
)

var mappingListWorkRequestsSortOrderEnum = map[string]ListWorkRequestsSortOrderEnum{
	"ASC":  ListWorkRequestsSortOrderAsc,
	"DESC": ListWorkRequestsSortOrderDesc,
}

// GetListWorkRequestsSortOrderEnumValues Enumerates the set of values for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumValues() []ListWorkRequestsSortOrderEnum {
	values := make([]ListWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortOrderEnum(val string) (ListWorkRequestsSortOrderEnum, bool) {
	mappingListWorkRequestsSortOrderEnumIgnoreCase := make(map[string]ListWorkRequestsSortOrderEnum)
	for k, v := range mappingListWorkRequestsSortOrderEnum {
		mappingListWorkRequestsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWorkRequestsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortByEnum Enum with underlying type: string
type ListWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortByEnum
const (
	ListWorkRequestsSortByTimecreated ListWorkRequestsSortByEnum = "TIMECREATED"
	ListWorkRequestsSortByDisplayname ListWorkRequestsSortByEnum = "DISPLAYNAME"
)

var mappingListWorkRequestsSortByEnum = map[string]ListWorkRequestsSortByEnum{
	"TIMECREATED": ListWorkRequestsSortByTimecreated,
	"DISPLAYNAME": ListWorkRequestsSortByDisplayname,
}

// GetListWorkRequestsSortByEnumValues Enumerates the set of values for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumValues() []ListWorkRequestsSortByEnum {
	values := make([]ListWorkRequestsSortByEnum, 0)
	for _, v := range mappingListWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortByEnum(val string) (ListWorkRequestsSortByEnum, bool) {
	mappingListWorkRequestsSortByEnumIgnoreCase := make(map[string]ListWorkRequestsSortByEnum)
	for k, v := range mappingListWorkRequestsSortByEnum {
		mappingListWorkRequestsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWorkRequestsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsOsFamilyEnum Enum with underlying type: string
type ListWorkRequestsOsFamilyEnum string

// Set of constants representing the allowable values for ListWorkRequestsOsFamilyEnum
const (
	ListWorkRequestsOsFamilyLinux   ListWorkRequestsOsFamilyEnum = "LINUX"
	ListWorkRequestsOsFamilyWindows ListWorkRequestsOsFamilyEnum = "WINDOWS"
	ListWorkRequestsOsFamilyAll     ListWorkRequestsOsFamilyEnum = "ALL"
)

var mappingListWorkRequestsOsFamilyEnum = map[string]ListWorkRequestsOsFamilyEnum{
	"LINUX":   ListWorkRequestsOsFamilyLinux,
	"WINDOWS": ListWorkRequestsOsFamilyWindows,
	"ALL":     ListWorkRequestsOsFamilyAll,
}

// GetListWorkRequestsOsFamilyEnumValues Enumerates the set of values for ListWorkRequestsOsFamilyEnum
func GetListWorkRequestsOsFamilyEnumValues() []ListWorkRequestsOsFamilyEnum {
	values := make([]ListWorkRequestsOsFamilyEnum, 0)
	for _, v := range mappingListWorkRequestsOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsOsFamilyEnumStringValues Enumerates the set of values in String for ListWorkRequestsOsFamilyEnum
func GetListWorkRequestsOsFamilyEnumStringValues() []string {
	return []string{
		"LINUX",
		"WINDOWS",
		"ALL",
	}
}

// GetMappingListWorkRequestsOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsOsFamilyEnum(val string) (ListWorkRequestsOsFamilyEnum, bool) {
	mappingListWorkRequestsOsFamilyEnumIgnoreCase := make(map[string]ListWorkRequestsOsFamilyEnum)
	for k, v := range mappingListWorkRequestsOsFamilyEnum {
		mappingListWorkRequestsOsFamilyEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListWorkRequestsOsFamilyEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
