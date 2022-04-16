// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDbSystemUpgradeHistoryEntriesRequest wrapper for the ListDbSystemUpgradeHistoryEntries operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDbSystemUpgradeHistoryEntries.go.html to see an example of how to use ListDbSystemUpgradeHistoryEntriesRequest.
type ListDbSystemUpgradeHistoryEntriesRequest struct {

	// The DB system OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DbSystemId *string `mandatory:"true" contributesTo:"path" name:"dbSystemId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDbSystemUpgradeHistoryEntriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is ascending.
	SortBy ListDbSystemUpgradeHistoryEntriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only upgradeHistoryEntries that match the specified Upgrade Action.
	UpgradeAction DbSystemUpgradeHistoryEntrySummaryActionEnum `mandatory:"false" contributesTo:"query" name:"upgradeAction" omitEmpty:"true"`

	// A filter to return only upgrade history entries that match the given lifecycle state exactly.
	LifecycleState DbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbSystemUpgradeHistoryEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbSystemUpgradeHistoryEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbSystemUpgradeHistoryEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbSystemUpgradeHistoryEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDbSystemUpgradeHistoryEntriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDbSystemUpgradeHistoryEntriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDbSystemUpgradeHistoryEntriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDbSystemUpgradeHistoryEntriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDbSystemUpgradeHistoryEntriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemUpgradeHistoryEntrySummaryActionEnum(string(request.UpgradeAction)); !ok && request.UpgradeAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpgradeAction: %s. Supported values are: %s.", request.UpgradeAction, strings.Join(GetDbSystemUpgradeHistoryEntrySummaryActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDbSystemUpgradeHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDbSystemUpgradeHistoryEntriesResponse wrapper for the ListDbSystemUpgradeHistoryEntries operation
type ListDbSystemUpgradeHistoryEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DbSystemUpgradeHistoryEntrySummary instances
	Items []DbSystemUpgradeHistoryEntrySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbSystemUpgradeHistoryEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbSystemUpgradeHistoryEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbSystemUpgradeHistoryEntriesSortOrderEnum Enum with underlying type: string
type ListDbSystemUpgradeHistoryEntriesSortOrderEnum string

// Set of constants representing the allowable values for ListDbSystemUpgradeHistoryEntriesSortOrderEnum
const (
	ListDbSystemUpgradeHistoryEntriesSortOrderAsc  ListDbSystemUpgradeHistoryEntriesSortOrderEnum = "ASC"
	ListDbSystemUpgradeHistoryEntriesSortOrderDesc ListDbSystemUpgradeHistoryEntriesSortOrderEnum = "DESC"
)

var mappingListDbSystemUpgradeHistoryEntriesSortOrderEnum = map[string]ListDbSystemUpgradeHistoryEntriesSortOrderEnum{
	"ASC":  ListDbSystemUpgradeHistoryEntriesSortOrderAsc,
	"DESC": ListDbSystemUpgradeHistoryEntriesSortOrderDesc,
}

var mappingListDbSystemUpgradeHistoryEntriesSortOrderEnumLowerCase = map[string]ListDbSystemUpgradeHistoryEntriesSortOrderEnum{
	"asc":  ListDbSystemUpgradeHistoryEntriesSortOrderAsc,
	"desc": ListDbSystemUpgradeHistoryEntriesSortOrderDesc,
}

// GetListDbSystemUpgradeHistoryEntriesSortOrderEnumValues Enumerates the set of values for ListDbSystemUpgradeHistoryEntriesSortOrderEnum
func GetListDbSystemUpgradeHistoryEntriesSortOrderEnumValues() []ListDbSystemUpgradeHistoryEntriesSortOrderEnum {
	values := make([]ListDbSystemUpgradeHistoryEntriesSortOrderEnum, 0)
	for _, v := range mappingListDbSystemUpgradeHistoryEntriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemUpgradeHistoryEntriesSortOrderEnumStringValues Enumerates the set of values in String for ListDbSystemUpgradeHistoryEntriesSortOrderEnum
func GetListDbSystemUpgradeHistoryEntriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDbSystemUpgradeHistoryEntriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemUpgradeHistoryEntriesSortOrderEnum(val string) (ListDbSystemUpgradeHistoryEntriesSortOrderEnum, bool) {
	enum, ok := mappingListDbSystemUpgradeHistoryEntriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDbSystemUpgradeHistoryEntriesSortByEnum Enum with underlying type: string
type ListDbSystemUpgradeHistoryEntriesSortByEnum string

// Set of constants representing the allowable values for ListDbSystemUpgradeHistoryEntriesSortByEnum
const (
	ListDbSystemUpgradeHistoryEntriesSortByTimestarted ListDbSystemUpgradeHistoryEntriesSortByEnum = "TIMESTARTED"
)

var mappingListDbSystemUpgradeHistoryEntriesSortByEnum = map[string]ListDbSystemUpgradeHistoryEntriesSortByEnum{
	"TIMESTARTED": ListDbSystemUpgradeHistoryEntriesSortByTimestarted,
}

var mappingListDbSystemUpgradeHistoryEntriesSortByEnumLowerCase = map[string]ListDbSystemUpgradeHistoryEntriesSortByEnum{
	"timestarted": ListDbSystemUpgradeHistoryEntriesSortByTimestarted,
}

// GetListDbSystemUpgradeHistoryEntriesSortByEnumValues Enumerates the set of values for ListDbSystemUpgradeHistoryEntriesSortByEnum
func GetListDbSystemUpgradeHistoryEntriesSortByEnumValues() []ListDbSystemUpgradeHistoryEntriesSortByEnum {
	values := make([]ListDbSystemUpgradeHistoryEntriesSortByEnum, 0)
	for _, v := range mappingListDbSystemUpgradeHistoryEntriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDbSystemUpgradeHistoryEntriesSortByEnumStringValues Enumerates the set of values in String for ListDbSystemUpgradeHistoryEntriesSortByEnum
func GetListDbSystemUpgradeHistoryEntriesSortByEnumStringValues() []string {
	return []string{
		"TIMESTARTED",
	}
}

// GetMappingListDbSystemUpgradeHistoryEntriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDbSystemUpgradeHistoryEntriesSortByEnum(val string) (ListDbSystemUpgradeHistoryEntriesSortByEnum, bool) {
	enum, ok := mappingListDbSystemUpgradeHistoryEntriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
