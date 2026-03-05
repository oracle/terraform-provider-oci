// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbSystemOsPatchHistoryEntriesRequest wrapper for the ListDbSystemOsPatchHistoryEntries operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemOsPatchHistoryEntries.go.html to see an example of how to use ListDbSystemOsPatchHistoryEntriesRequest.
type ListDbSystemOsPatchHistoryEntriesRequest struct {

	// The DB system OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"true" contributesTo:"path" name:"dbSystemId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDbSystemOsPatchHistoryEntriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMESTARTED is descending.
	SortBy ListDbSystemOsPatchHistoryEntriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only OS patch history entries that match the given lifecycle state exactly.
	LifecycleState DbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only OS patch history entries that match the specified OS patch action.
	Action DbSystemOsPatchHistoryEntrySummaryActionEnum `mandatory:"false" contributesTo:"query" name:"action" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbSystemOsPatchHistoryEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbSystemOsPatchHistoryEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbSystemOsPatchHistoryEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbSystemOsPatchHistoryEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbSystemOsPatchHistoryEntriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDbSystemOsPatchHistoryEntriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbSystemOsPatchHistoryEntriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbSystemOsPatchHistoryEntriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbSystemOsPatchHistoryEntriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbSystemOsPatchHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemOsPatchHistoryEntrySummaryActionEnum(string(request.Action)); !ok && request.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", request.Action, strings.Join(GetDbSystemOsPatchHistoryEntrySummaryActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbSystemOsPatchHistoryEntriesResponse wrapper for the ListDbSystemOsPatchHistoryEntries operation
type ListDbSystemOsPatchHistoryEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DbSystemOsPatchHistoryEntryCollection instances
	DbSystemOsPatchHistoryEntryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbSystemOsPatchHistoryEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbSystemOsPatchHistoryEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbSystemOsPatchHistoryEntriesSortOrderEnum Enum with underlying type: string
type ListDbSystemOsPatchHistoryEntriesSortOrderEnum string

// Set of constants representing the allowable values for ListDbSystemOsPatchHistoryEntriesSortOrderEnum
const (
	ListDbSystemOsPatchHistoryEntriesSortOrderAsc  ListDbSystemOsPatchHistoryEntriesSortOrderEnum = "ASC"
	ListDbSystemOsPatchHistoryEntriesSortOrderDesc ListDbSystemOsPatchHistoryEntriesSortOrderEnum = "DESC"
)

var mappingListDbSystemOsPatchHistoryEntriesSortOrderEnum = map[string]ListDbSystemOsPatchHistoryEntriesSortOrderEnum{
	"ASC":  ListDbSystemOsPatchHistoryEntriesSortOrderAsc,
	"DESC": ListDbSystemOsPatchHistoryEntriesSortOrderDesc,
}

var mappingListDbSystemOsPatchHistoryEntriesSortOrderEnumLowerCase = map[string]ListDbSystemOsPatchHistoryEntriesSortOrderEnum{
	"asc":  ListDbSystemOsPatchHistoryEntriesSortOrderAsc,
	"desc": ListDbSystemOsPatchHistoryEntriesSortOrderDesc,
}

// GetListDbSystemOsPatchHistoryEntriesSortOrderEnumValues Enumerates the set of values for ListDbSystemOsPatchHistoryEntriesSortOrderEnum
func GetListDbSystemOsPatchHistoryEntriesSortOrderEnumValues() []ListDbSystemOsPatchHistoryEntriesSortOrderEnum {
	values := make([]ListDbSystemOsPatchHistoryEntriesSortOrderEnum, 0)
	for _, v := range mappingListDbSystemOsPatchHistoryEntriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemOsPatchHistoryEntriesSortOrderEnumStringValues Enumerates the set of values in String for ListDbSystemOsPatchHistoryEntriesSortOrderEnum
func GetListDbSystemOsPatchHistoryEntriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbSystemOsPatchHistoryEntriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemOsPatchHistoryEntriesSortOrderEnum(val string) (ListDbSystemOsPatchHistoryEntriesSortOrderEnum, bool) {
	enum, ok := mappingListDbSystemOsPatchHistoryEntriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbSystemOsPatchHistoryEntriesSortByEnum Enum with underlying type: string
type ListDbSystemOsPatchHistoryEntriesSortByEnum string

// Set of constants representing the allowable values for ListDbSystemOsPatchHistoryEntriesSortByEnum
const (
	ListDbSystemOsPatchHistoryEntriesSortByTimestarted ListDbSystemOsPatchHistoryEntriesSortByEnum = "TIMESTARTED"
	ListDbSystemOsPatchHistoryEntriesSortByAction      ListDbSystemOsPatchHistoryEntriesSortByEnum = "ACTION"
)

var mappingListDbSystemOsPatchHistoryEntriesSortByEnum = map[string]ListDbSystemOsPatchHistoryEntriesSortByEnum{
	"TIMESTARTED": ListDbSystemOsPatchHistoryEntriesSortByTimestarted,
	"ACTION":      ListDbSystemOsPatchHistoryEntriesSortByAction,
}

var mappingListDbSystemOsPatchHistoryEntriesSortByEnumLowerCase = map[string]ListDbSystemOsPatchHistoryEntriesSortByEnum{
	"timestarted": ListDbSystemOsPatchHistoryEntriesSortByTimestarted,
	"action":      ListDbSystemOsPatchHistoryEntriesSortByAction,
}

// GetListDbSystemOsPatchHistoryEntriesSortByEnumValues Enumerates the set of values for ListDbSystemOsPatchHistoryEntriesSortByEnum
func GetListDbSystemOsPatchHistoryEntriesSortByEnumValues() []ListDbSystemOsPatchHistoryEntriesSortByEnum {
	values := make([]ListDbSystemOsPatchHistoryEntriesSortByEnum, 0)
	for _, v := range mappingListDbSystemOsPatchHistoryEntriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemOsPatchHistoryEntriesSortByEnumStringValues Enumerates the set of values in String for ListDbSystemOsPatchHistoryEntriesSortByEnum
func GetListDbSystemOsPatchHistoryEntriesSortByEnumStringValues() []string {
	return []string{
		"TIMESTARTED",
		"ACTION",
	}
}

// GetMappingListDbSystemOsPatchHistoryEntriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemOsPatchHistoryEntriesSortByEnum(val string) (ListDbSystemOsPatchHistoryEntriesSortByEnum, bool) {
	enum, ok := mappingListDbSystemOsPatchHistoryEntriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
