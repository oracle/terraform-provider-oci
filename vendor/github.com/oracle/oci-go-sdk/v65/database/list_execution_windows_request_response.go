// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExecutionWindowsRequest wrapper for the ListExecutionWindows operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExecutionWindows.go.html to see an example of how to use ListExecutionWindowsRequest.
type ListExecutionWindowsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListExecutionWindowsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the given resource id exactly.
	ExecutionResourceId *string `mandatory:"false" contributesTo:"query" name:"executionResourceId"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExecutionWindowsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ExecutionWindowSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExecutionWindowsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExecutionWindowsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExecutionWindowsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExecutionWindowsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExecutionWindowsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExecutionWindowsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExecutionWindowsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExecutionWindowsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExecutionWindowsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExecutionWindowSummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetExecutionWindowSummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExecutionWindowsResponse wrapper for the ListExecutionWindows operation
type ListExecutionWindowsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExecutionWindowSummary instances
	Items []ExecutionWindowSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExecutionWindowsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExecutionWindowsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExecutionWindowsSortByEnum Enum with underlying type: string
type ListExecutionWindowsSortByEnum string

// Set of constants representing the allowable values for ListExecutionWindowsSortByEnum
const (
	ListExecutionWindowsSortByTimecreated ListExecutionWindowsSortByEnum = "TIMECREATED"
	ListExecutionWindowsSortByDisplayname ListExecutionWindowsSortByEnum = "DISPLAYNAME"
)

var mappingListExecutionWindowsSortByEnum = map[string]ListExecutionWindowsSortByEnum{
	"TIMECREATED": ListExecutionWindowsSortByTimecreated,
	"DISPLAYNAME": ListExecutionWindowsSortByDisplayname,
}

var mappingListExecutionWindowsSortByEnumLowerCase = map[string]ListExecutionWindowsSortByEnum{
	"timecreated": ListExecutionWindowsSortByTimecreated,
	"displayname": ListExecutionWindowsSortByDisplayname,
}

// GetListExecutionWindowsSortByEnumValues Enumerates the set of values for ListExecutionWindowsSortByEnum
func GetListExecutionWindowsSortByEnumValues() []ListExecutionWindowsSortByEnum {
	values := make([]ListExecutionWindowsSortByEnum, 0)
	for _, v := range mappingListExecutionWindowsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExecutionWindowsSortByEnumStringValues Enumerates the set of values in String for ListExecutionWindowsSortByEnum
func GetListExecutionWindowsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListExecutionWindowsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExecutionWindowsSortByEnum(val string) (ListExecutionWindowsSortByEnum, bool) {
	enum, ok := mappingListExecutionWindowsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExecutionWindowsSortOrderEnum Enum with underlying type: string
type ListExecutionWindowsSortOrderEnum string

// Set of constants representing the allowable values for ListExecutionWindowsSortOrderEnum
const (
	ListExecutionWindowsSortOrderAsc  ListExecutionWindowsSortOrderEnum = "ASC"
	ListExecutionWindowsSortOrderDesc ListExecutionWindowsSortOrderEnum = "DESC"
)

var mappingListExecutionWindowsSortOrderEnum = map[string]ListExecutionWindowsSortOrderEnum{
	"ASC":  ListExecutionWindowsSortOrderAsc,
	"DESC": ListExecutionWindowsSortOrderDesc,
}

var mappingListExecutionWindowsSortOrderEnumLowerCase = map[string]ListExecutionWindowsSortOrderEnum{
	"asc":  ListExecutionWindowsSortOrderAsc,
	"desc": ListExecutionWindowsSortOrderDesc,
}

// GetListExecutionWindowsSortOrderEnumValues Enumerates the set of values for ListExecutionWindowsSortOrderEnum
func GetListExecutionWindowsSortOrderEnumValues() []ListExecutionWindowsSortOrderEnum {
	values := make([]ListExecutionWindowsSortOrderEnum, 0)
	for _, v := range mappingListExecutionWindowsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExecutionWindowsSortOrderEnumStringValues Enumerates the set of values in String for ListExecutionWindowsSortOrderEnum
func GetListExecutionWindowsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExecutionWindowsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExecutionWindowsSortOrderEnum(val string) (ListExecutionWindowsSortOrderEnum, bool) {
	enum, ok := mappingListExecutionWindowsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
