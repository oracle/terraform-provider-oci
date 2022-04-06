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

// ListPdbConversionHistoryEntriesRequest wrapper for the ListPdbConversionHistoryEntries operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListPdbConversionHistoryEntries.go.html to see an example of how to use ListPdbConversionHistoryEntriesRequest.
type ListPdbConversionHistoryEntriesRequest struct {

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DatabaseId *string `mandatory:"true" contributesTo:"path" name:"databaseId"`

	// A filter to return only the pluggable database conversion history entries that match the specified conversion action. For example, you can use this filter to return only entries for the precheck operation.
	PdbConversionAction PdbConversionHistoryEntrySummaryActionEnum `mandatory:"false" contributesTo:"query" name:"pdbConversionAction" omitEmpty:"true"`

	// A filter to return only the pluggable database conversion history entries that match the specified lifecycle state. For example, you can use this filter to return only entries in the "failed" lifecycle state.
	LifecycleState PdbConversionHistoryEntrySummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). The default order for `TIMECREATED` is ascending.
	SortBy ListPdbConversionHistoryEntriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPdbConversionHistoryEntriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListPdbConversionHistoryEntriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPdbConversionHistoryEntriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPdbConversionHistoryEntriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPdbConversionHistoryEntriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPdbConversionHistoryEntriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPdbConversionHistoryEntrySummaryActionEnum(string(request.PdbConversionAction)); !ok && request.PdbConversionAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PdbConversionAction: %s. Supported values are: %s.", request.PdbConversionAction, strings.Join(GetPdbConversionHistoryEntrySummaryActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPdbConversionHistoryEntrySummaryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPdbConversionHistoryEntrySummaryLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPdbConversionHistoryEntriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPdbConversionHistoryEntriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPdbConversionHistoryEntriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPdbConversionHistoryEntriesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPdbConversionHistoryEntriesResponse wrapper for the ListPdbConversionHistoryEntries operation
type ListPdbConversionHistoryEntriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PdbConversionHistoryEntrySummary instances
	Items []PdbConversionHistoryEntrySummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPdbConversionHistoryEntriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPdbConversionHistoryEntriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPdbConversionHistoryEntriesSortByEnum Enum with underlying type: string
type ListPdbConversionHistoryEntriesSortByEnum string

// Set of constants representing the allowable values for ListPdbConversionHistoryEntriesSortByEnum
const (
	ListPdbConversionHistoryEntriesSortByTimestarted ListPdbConversionHistoryEntriesSortByEnum = "TIMESTARTED"
)

var mappingListPdbConversionHistoryEntriesSortByEnum = map[string]ListPdbConversionHistoryEntriesSortByEnum{
	"TIMESTARTED": ListPdbConversionHistoryEntriesSortByTimestarted,
}

var mappingListPdbConversionHistoryEntriesSortByEnumLowerCase = map[string]ListPdbConversionHistoryEntriesSortByEnum{
	"timestarted": ListPdbConversionHistoryEntriesSortByTimestarted,
}

// GetListPdbConversionHistoryEntriesSortByEnumValues Enumerates the set of values for ListPdbConversionHistoryEntriesSortByEnum
func GetListPdbConversionHistoryEntriesSortByEnumValues() []ListPdbConversionHistoryEntriesSortByEnum {
	values := make([]ListPdbConversionHistoryEntriesSortByEnum, 0)
	for _, v := range mappingListPdbConversionHistoryEntriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPdbConversionHistoryEntriesSortByEnumStringValues Enumerates the set of values in String for ListPdbConversionHistoryEntriesSortByEnum
func GetListPdbConversionHistoryEntriesSortByEnumStringValues() []string {
	return []string{
		"TIMESTARTED",
	}
}

// GetMappingListPdbConversionHistoryEntriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPdbConversionHistoryEntriesSortByEnum(val string) (ListPdbConversionHistoryEntriesSortByEnum, bool) {
	enum, ok := mappingListPdbConversionHistoryEntriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPdbConversionHistoryEntriesSortOrderEnum Enum with underlying type: string
type ListPdbConversionHistoryEntriesSortOrderEnum string

// Set of constants representing the allowable values for ListPdbConversionHistoryEntriesSortOrderEnum
const (
	ListPdbConversionHistoryEntriesSortOrderAsc  ListPdbConversionHistoryEntriesSortOrderEnum = "ASC"
	ListPdbConversionHistoryEntriesSortOrderDesc ListPdbConversionHistoryEntriesSortOrderEnum = "DESC"
)

var mappingListPdbConversionHistoryEntriesSortOrderEnum = map[string]ListPdbConversionHistoryEntriesSortOrderEnum{
	"ASC":  ListPdbConversionHistoryEntriesSortOrderAsc,
	"DESC": ListPdbConversionHistoryEntriesSortOrderDesc,
}

var mappingListPdbConversionHistoryEntriesSortOrderEnumLowerCase = map[string]ListPdbConversionHistoryEntriesSortOrderEnum{
	"asc":  ListPdbConversionHistoryEntriesSortOrderAsc,
	"desc": ListPdbConversionHistoryEntriesSortOrderDesc,
}

// GetListPdbConversionHistoryEntriesSortOrderEnumValues Enumerates the set of values for ListPdbConversionHistoryEntriesSortOrderEnum
func GetListPdbConversionHistoryEntriesSortOrderEnumValues() []ListPdbConversionHistoryEntriesSortOrderEnum {
	values := make([]ListPdbConversionHistoryEntriesSortOrderEnum, 0)
	for _, v := range mappingListPdbConversionHistoryEntriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPdbConversionHistoryEntriesSortOrderEnumStringValues Enumerates the set of values in String for ListPdbConversionHistoryEntriesSortOrderEnum
func GetListPdbConversionHistoryEntriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPdbConversionHistoryEntriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPdbConversionHistoryEntriesSortOrderEnum(val string) (ListPdbConversionHistoryEntriesSortOrderEnum, bool) {
	enum, ok := mappingListPdbConversionHistoryEntriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
