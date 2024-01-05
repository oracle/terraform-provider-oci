// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSqlPlanBaselineJobsRequest wrapper for the ListSqlPlanBaselineJobs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlPlanBaselineJobs.go.html to see an example of how to use ListSqlPlanBaselineJobsRequest.
type ListSqlPlanBaselineJobsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// A filter to return the SQL plan baseline jobs that match the name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListSqlPlanBaselineJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListSqlPlanBaselineJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlPlanBaselineJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlPlanBaselineJobsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlPlanBaselineJobsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlPlanBaselineJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlPlanBaselineJobsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlPlanBaselineJobsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlPlanBaselineJobsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlPlanBaselineJobsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlPlanBaselineJobsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlPlanBaselineJobsResponse wrapper for the ListSqlPlanBaselineJobs operation
type ListSqlPlanBaselineJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlPlanBaselineJobCollection instances
	SqlPlanBaselineJobCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSqlPlanBaselineJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlPlanBaselineJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlPlanBaselineJobsSortByEnum Enum with underlying type: string
type ListSqlPlanBaselineJobsSortByEnum string

// Set of constants representing the allowable values for ListSqlPlanBaselineJobsSortByEnum
const (
	ListSqlPlanBaselineJobsSortByTimecreated ListSqlPlanBaselineJobsSortByEnum = "TIMECREATED"
	ListSqlPlanBaselineJobsSortByName        ListSqlPlanBaselineJobsSortByEnum = "NAME"
)

var mappingListSqlPlanBaselineJobsSortByEnum = map[string]ListSqlPlanBaselineJobsSortByEnum{
	"TIMECREATED": ListSqlPlanBaselineJobsSortByTimecreated,
	"NAME":        ListSqlPlanBaselineJobsSortByName,
}

var mappingListSqlPlanBaselineJobsSortByEnumLowerCase = map[string]ListSqlPlanBaselineJobsSortByEnum{
	"timecreated": ListSqlPlanBaselineJobsSortByTimecreated,
	"name":        ListSqlPlanBaselineJobsSortByName,
}

// GetListSqlPlanBaselineJobsSortByEnumValues Enumerates the set of values for ListSqlPlanBaselineJobsSortByEnum
func GetListSqlPlanBaselineJobsSortByEnumValues() []ListSqlPlanBaselineJobsSortByEnum {
	values := make([]ListSqlPlanBaselineJobsSortByEnum, 0)
	for _, v := range mappingListSqlPlanBaselineJobsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlPlanBaselineJobsSortByEnumStringValues Enumerates the set of values in String for ListSqlPlanBaselineJobsSortByEnum
func GetListSqlPlanBaselineJobsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListSqlPlanBaselineJobsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlPlanBaselineJobsSortByEnum(val string) (ListSqlPlanBaselineJobsSortByEnum, bool) {
	enum, ok := mappingListSqlPlanBaselineJobsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlPlanBaselineJobsSortOrderEnum Enum with underlying type: string
type ListSqlPlanBaselineJobsSortOrderEnum string

// Set of constants representing the allowable values for ListSqlPlanBaselineJobsSortOrderEnum
const (
	ListSqlPlanBaselineJobsSortOrderAsc  ListSqlPlanBaselineJobsSortOrderEnum = "ASC"
	ListSqlPlanBaselineJobsSortOrderDesc ListSqlPlanBaselineJobsSortOrderEnum = "DESC"
)

var mappingListSqlPlanBaselineJobsSortOrderEnum = map[string]ListSqlPlanBaselineJobsSortOrderEnum{
	"ASC":  ListSqlPlanBaselineJobsSortOrderAsc,
	"DESC": ListSqlPlanBaselineJobsSortOrderDesc,
}

var mappingListSqlPlanBaselineJobsSortOrderEnumLowerCase = map[string]ListSqlPlanBaselineJobsSortOrderEnum{
	"asc":  ListSqlPlanBaselineJobsSortOrderAsc,
	"desc": ListSqlPlanBaselineJobsSortOrderDesc,
}

// GetListSqlPlanBaselineJobsSortOrderEnumValues Enumerates the set of values for ListSqlPlanBaselineJobsSortOrderEnum
func GetListSqlPlanBaselineJobsSortOrderEnumValues() []ListSqlPlanBaselineJobsSortOrderEnum {
	values := make([]ListSqlPlanBaselineJobsSortOrderEnum, 0)
	for _, v := range mappingListSqlPlanBaselineJobsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlPlanBaselineJobsSortOrderEnumStringValues Enumerates the set of values in String for ListSqlPlanBaselineJobsSortOrderEnum
func GetListSqlPlanBaselineJobsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlPlanBaselineJobsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlPlanBaselineJobsSortOrderEnum(val string) (ListSqlPlanBaselineJobsSortOrderEnum, bool) {
	enum, ok := mappingListSqlPlanBaselineJobsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
