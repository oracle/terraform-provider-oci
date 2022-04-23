// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAddmReportsRequest wrapper for the ListAddmReports operation
type ListAddmReportsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The beginning of the time range to search for ADDM reports as defined by date-time RFC3339 format.
	TimeStart *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeStart"`

	// The end of the time range to search for ADDM reports as defined by date-time RFC3339 format.
	TimeEnd *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeEnd"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the list of ADDM tasks.
	SortBy ListAddmReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder ListAddmReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAddmReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAddmReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAddmReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAddmReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAddmReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAddmReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAddmReportsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAddmReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAddmReportsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAddmReportsResponse wrapper for the ListAddmReports operation
type ListAddmReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ListAddmReportsCollection instances
	ListAddmReportsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAddmReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAddmReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAddmReportsSortByEnum Enum with underlying type: string
type ListAddmReportsSortByEnum string

// Set of constants representing the allowable values for ListAddmReportsSortByEnum
const (
	ListAddmReportsSortByTaskName    ListAddmReportsSortByEnum = "TASK_NAME"
	ListAddmReportsSortByTaskId      ListAddmReportsSortByEnum = "TASK_ID"
	ListAddmReportsSortByDescription ListAddmReportsSortByEnum = "DESCRIPTION"
	ListAddmReportsSortByDbUser      ListAddmReportsSortByEnum = "DB_USER"
	ListAddmReportsSortByStatus      ListAddmReportsSortByEnum = "STATUS"
	ListAddmReportsSortByTimeCreated ListAddmReportsSortByEnum = "TIME_CREATED"
	ListAddmReportsSortByBeginTime   ListAddmReportsSortByEnum = "BEGIN_TIME"
	ListAddmReportsSortByEndTime     ListAddmReportsSortByEnum = "END_TIME"
)

var mappingListAddmReportsSortByEnum = map[string]ListAddmReportsSortByEnum{
	"TASK_NAME":    ListAddmReportsSortByTaskName,
	"TASK_ID":      ListAddmReportsSortByTaskId,
	"DESCRIPTION":  ListAddmReportsSortByDescription,
	"DB_USER":      ListAddmReportsSortByDbUser,
	"STATUS":       ListAddmReportsSortByStatus,
	"TIME_CREATED": ListAddmReportsSortByTimeCreated,
	"BEGIN_TIME":   ListAddmReportsSortByBeginTime,
	"END_TIME":     ListAddmReportsSortByEndTime,
}

var mappingListAddmReportsSortByEnumLowerCase = map[string]ListAddmReportsSortByEnum{
	"task_name":    ListAddmReportsSortByTaskName,
	"task_id":      ListAddmReportsSortByTaskId,
	"description":  ListAddmReportsSortByDescription,
	"db_user":      ListAddmReportsSortByDbUser,
	"status":       ListAddmReportsSortByStatus,
	"time_created": ListAddmReportsSortByTimeCreated,
	"begin_time":   ListAddmReportsSortByBeginTime,
	"end_time":     ListAddmReportsSortByEndTime,
}

// GetListAddmReportsSortByEnumValues Enumerates the set of values for ListAddmReportsSortByEnum
func GetListAddmReportsSortByEnumValues() []ListAddmReportsSortByEnum {
	values := make([]ListAddmReportsSortByEnum, 0)
	for _, v := range mappingListAddmReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddmReportsSortByEnumStringValues Enumerates the set of values in String for ListAddmReportsSortByEnum
func GetListAddmReportsSortByEnumStringValues() []string {
	return []string{
		"TASK_NAME",
		"TASK_ID",
		"DESCRIPTION",
		"DB_USER",
		"STATUS",
		"TIME_CREATED",
		"BEGIN_TIME",
		"END_TIME",
	}
}

// GetMappingListAddmReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddmReportsSortByEnum(val string) (ListAddmReportsSortByEnum, bool) {
	enum, ok := mappingListAddmReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAddmReportsSortOrderEnum Enum with underlying type: string
type ListAddmReportsSortOrderEnum string

// Set of constants representing the allowable values for ListAddmReportsSortOrderEnum
const (
	ListAddmReportsSortOrderAsc  ListAddmReportsSortOrderEnum = "ASC"
	ListAddmReportsSortOrderDesc ListAddmReportsSortOrderEnum = "DESC"
)

var mappingListAddmReportsSortOrderEnum = map[string]ListAddmReportsSortOrderEnum{
	"ASC":  ListAddmReportsSortOrderAsc,
	"DESC": ListAddmReportsSortOrderDesc,
}

var mappingListAddmReportsSortOrderEnumLowerCase = map[string]ListAddmReportsSortOrderEnum{
	"asc":  ListAddmReportsSortOrderAsc,
	"desc": ListAddmReportsSortOrderDesc,
}

// GetListAddmReportsSortOrderEnumValues Enumerates the set of values for ListAddmReportsSortOrderEnum
func GetListAddmReportsSortOrderEnumValues() []ListAddmReportsSortOrderEnum {
	values := make([]ListAddmReportsSortOrderEnum, 0)
	for _, v := range mappingListAddmReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAddmReportsSortOrderEnumStringValues Enumerates the set of values in String for ListAddmReportsSortOrderEnum
func GetListAddmReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAddmReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAddmReportsSortOrderEnum(val string) (ListAddmReportsSortOrderEnum, bool) {
	enum, ok := mappingListAddmReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
