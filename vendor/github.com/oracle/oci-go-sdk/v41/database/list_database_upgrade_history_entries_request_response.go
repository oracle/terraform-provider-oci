// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// ListDatabaseUpgradeHistoryEntriesRequest wrapper for the ListDatabaseUpgradeHistoryEntries operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListDatabaseUpgradeHistoryEntries.go.html to see an example of how to use ListDatabaseUpgradeHistoryEntriesRequest.
type ListDatabaseUpgradeHistoryEntriesRequest struct {

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DatabaseId *string `mandatory:"true" contributesTo:"path" name:"databaseId"`

	// A filter to return only upgradeHistoryEntries that match the specified Upgrade Action.
	UpgradeAction DatabaseUpgradeHistoryEntrySummaryActionEnum `mandatory:"false" contributesTo:"query" name:"upgradeAction" omitEmpty:"true"`

	// A filter to return only upgradeHistoryEntries that match the given lifecycle state exactly.
	LifecycleState DatabaseUpgradeHistoryEntrySummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is ascending.
	SortBy ListDatabaseUpgradeHistoryEntriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListDatabaseUpgradeHistoryEntriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseUpgradeHistoryEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseUpgradeHistoryEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseUpgradeHistoryEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseUpgradeHistoryEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDatabaseUpgradeHistoryEntriesResponse wrapper for the ListDatabaseUpgradeHistoryEntries operation
type ListDatabaseUpgradeHistoryEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DatabaseUpgradeHistoryEntrySummary instances
	Items []DatabaseUpgradeHistoryEntrySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseUpgradeHistoryEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseUpgradeHistoryEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseUpgradeHistoryEntriesSortByEnum Enum with underlying type: string
type ListDatabaseUpgradeHistoryEntriesSortByEnum string

// Set of constants representing the allowable values for ListDatabaseUpgradeHistoryEntriesSortByEnum
const (
	ListDatabaseUpgradeHistoryEntriesSortByTimestarted ListDatabaseUpgradeHistoryEntriesSortByEnum = "TIMESTARTED"
)

var mappingListDatabaseUpgradeHistoryEntriesSortBy = map[string]ListDatabaseUpgradeHistoryEntriesSortByEnum{
	"TIMESTARTED": ListDatabaseUpgradeHistoryEntriesSortByTimestarted,
}

// GetListDatabaseUpgradeHistoryEntriesSortByEnumValues Enumerates the set of values for ListDatabaseUpgradeHistoryEntriesSortByEnum
func GetListDatabaseUpgradeHistoryEntriesSortByEnumValues() []ListDatabaseUpgradeHistoryEntriesSortByEnum {
	values := make([]ListDatabaseUpgradeHistoryEntriesSortByEnum, 0)
	for _, v := range mappingListDatabaseUpgradeHistoryEntriesSortBy {
		values = append(values, v)
	}
	return values
}

// ListDatabaseUpgradeHistoryEntriesSortOrderEnum Enum with underlying type: string
type ListDatabaseUpgradeHistoryEntriesSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseUpgradeHistoryEntriesSortOrderEnum
const (
	ListDatabaseUpgradeHistoryEntriesSortOrderAsc  ListDatabaseUpgradeHistoryEntriesSortOrderEnum = "ASC"
	ListDatabaseUpgradeHistoryEntriesSortOrderDesc ListDatabaseUpgradeHistoryEntriesSortOrderEnum = "DESC"
)

var mappingListDatabaseUpgradeHistoryEntriesSortOrder = map[string]ListDatabaseUpgradeHistoryEntriesSortOrderEnum{
	"ASC":  ListDatabaseUpgradeHistoryEntriesSortOrderAsc,
	"DESC": ListDatabaseUpgradeHistoryEntriesSortOrderDesc,
}

// GetListDatabaseUpgradeHistoryEntriesSortOrderEnumValues Enumerates the set of values for ListDatabaseUpgradeHistoryEntriesSortOrderEnum
func GetListDatabaseUpgradeHistoryEntriesSortOrderEnumValues() []ListDatabaseUpgradeHistoryEntriesSortOrderEnum {
	values := make([]ListDatabaseUpgradeHistoryEntriesSortOrderEnum, 0)
	for _, v := range mappingListDatabaseUpgradeHistoryEntriesSortOrder {
		values = append(values, v)
	}
	return values
}
