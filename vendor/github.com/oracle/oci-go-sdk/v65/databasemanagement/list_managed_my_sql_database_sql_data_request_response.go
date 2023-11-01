// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedMySqlDatabaseSqlDataRequest wrapper for the ListManagedMySqlDatabaseSqlData operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedMySqlDatabaseSqlData.go.html to see an example of how to use ListManagedMySqlDatabaseSqlDataRequest.
type ListManagedMySqlDatabaseSqlDataRequest struct {

	// The OCID of the Managed MySQL Database.
	ManagedMySqlDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedMySqlDatabaseId"`

	// The start time of the time range to retrieve the health metrics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	StartTime *string `mandatory:"true" contributesTo:"query" name:"startTime"`

	// The end time of the time range to retrieve the health metrics of a Managed Database
	// in UTC in ISO-8601 format, which is "yyyy-MM-dd'T'hh:mm:ss.sss'Z'".
	EndTime *string `mandatory:"true" contributesTo:"query" name:"endTime"`

	// The parameter to filter results by key criteria which include :
	// - SUM_TIMER_WAIT
	// - COUNT_STAR
	// - SUM_ERRORS
	// - SUM_ROWS_AFFECTED
	// - SUM_ROWS_SENT
	// - SUM_ROWS_EXAMINED
	// - SUM_CREATED_TMP_TABLES
	// - SUM_NO_INDEX_USED
	// - SUM_NO_GOOD_INDEX_USED
	// - FIRST_SEEN
	// - LAST_SEEN
	FilterColumn *string `mandatory:"false" contributesTo:"query" name:"filterColumn"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListManagedMySqlDatabaseSqlDataSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListManagedMySqlDatabaseSqlDataSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedMySqlDatabaseSqlDataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedMySqlDatabaseSqlDataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedMySqlDatabaseSqlDataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedMySqlDatabaseSqlDataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedMySqlDatabaseSqlDataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedMySqlDatabaseSqlDataSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedMySqlDatabaseSqlDataSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedMySqlDatabaseSqlDataSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedMySqlDatabaseSqlDataSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedMySqlDatabaseSqlDataResponse wrapper for the ListManagedMySqlDatabaseSqlData operation
type ListManagedMySqlDatabaseSqlDataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MySqlDataCollection instances
	MySqlDataCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedMySqlDatabaseSqlDataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedMySqlDatabaseSqlDataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedMySqlDatabaseSqlDataSortByEnum Enum with underlying type: string
type ListManagedMySqlDatabaseSqlDataSortByEnum string

// Set of constants representing the allowable values for ListManagedMySqlDatabaseSqlDataSortByEnum
const (
	ListManagedMySqlDatabaseSqlDataSortByTimecreated ListManagedMySqlDatabaseSqlDataSortByEnum = "TIMECREATED"
	ListManagedMySqlDatabaseSqlDataSortByName        ListManagedMySqlDatabaseSqlDataSortByEnum = "NAME"
)

var mappingListManagedMySqlDatabaseSqlDataSortByEnum = map[string]ListManagedMySqlDatabaseSqlDataSortByEnum{
	"TIMECREATED": ListManagedMySqlDatabaseSqlDataSortByTimecreated,
	"NAME":        ListManagedMySqlDatabaseSqlDataSortByName,
}

var mappingListManagedMySqlDatabaseSqlDataSortByEnumLowerCase = map[string]ListManagedMySqlDatabaseSqlDataSortByEnum{
	"timecreated": ListManagedMySqlDatabaseSqlDataSortByTimecreated,
	"name":        ListManagedMySqlDatabaseSqlDataSortByName,
}

// GetListManagedMySqlDatabaseSqlDataSortByEnumValues Enumerates the set of values for ListManagedMySqlDatabaseSqlDataSortByEnum
func GetListManagedMySqlDatabaseSqlDataSortByEnumValues() []ListManagedMySqlDatabaseSqlDataSortByEnum {
	values := make([]ListManagedMySqlDatabaseSqlDataSortByEnum, 0)
	for _, v := range mappingListManagedMySqlDatabaseSqlDataSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedMySqlDatabaseSqlDataSortByEnumStringValues Enumerates the set of values in String for ListManagedMySqlDatabaseSqlDataSortByEnum
func GetListManagedMySqlDatabaseSqlDataSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListManagedMySqlDatabaseSqlDataSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedMySqlDatabaseSqlDataSortByEnum(val string) (ListManagedMySqlDatabaseSqlDataSortByEnum, bool) {
	enum, ok := mappingListManagedMySqlDatabaseSqlDataSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedMySqlDatabaseSqlDataSortOrderEnum Enum with underlying type: string
type ListManagedMySqlDatabaseSqlDataSortOrderEnum string

// Set of constants representing the allowable values for ListManagedMySqlDatabaseSqlDataSortOrderEnum
const (
	ListManagedMySqlDatabaseSqlDataSortOrderAsc  ListManagedMySqlDatabaseSqlDataSortOrderEnum = "ASC"
	ListManagedMySqlDatabaseSqlDataSortOrderDesc ListManagedMySqlDatabaseSqlDataSortOrderEnum = "DESC"
)

var mappingListManagedMySqlDatabaseSqlDataSortOrderEnum = map[string]ListManagedMySqlDatabaseSqlDataSortOrderEnum{
	"ASC":  ListManagedMySqlDatabaseSqlDataSortOrderAsc,
	"DESC": ListManagedMySqlDatabaseSqlDataSortOrderDesc,
}

var mappingListManagedMySqlDatabaseSqlDataSortOrderEnumLowerCase = map[string]ListManagedMySqlDatabaseSqlDataSortOrderEnum{
	"asc":  ListManagedMySqlDatabaseSqlDataSortOrderAsc,
	"desc": ListManagedMySqlDatabaseSqlDataSortOrderDesc,
}

// GetListManagedMySqlDatabaseSqlDataSortOrderEnumValues Enumerates the set of values for ListManagedMySqlDatabaseSqlDataSortOrderEnum
func GetListManagedMySqlDatabaseSqlDataSortOrderEnumValues() []ListManagedMySqlDatabaseSqlDataSortOrderEnum {
	values := make([]ListManagedMySqlDatabaseSqlDataSortOrderEnum, 0)
	for _, v := range mappingListManagedMySqlDatabaseSqlDataSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedMySqlDatabaseSqlDataSortOrderEnumStringValues Enumerates the set of values in String for ListManagedMySqlDatabaseSqlDataSortOrderEnum
func GetListManagedMySqlDatabaseSqlDataSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedMySqlDatabaseSqlDataSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedMySqlDatabaseSqlDataSortOrderEnum(val string) (ListManagedMySqlDatabaseSqlDataSortOrderEnum, bool) {
	enum, ok := mappingListManagedMySqlDatabaseSqlDataSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
