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

// AddmTasksRequest wrapper for the AddmTasks operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/AddmTasks.go.html to see an example of how to use AddmTasksRequest.
type AddmTasksRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The beginning of the time range to search for ADDM tasks as defined by date-time RFC3339 format.
	TimeStart *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeStart"`

	// The end of the time range to search for ADDM tasks as defined by date-time RFC3339 format.
	TimeEnd *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeEnd"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the list of ADDM tasks.
	SortBy AddmTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder AddmTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request AddmTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request AddmTasksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request AddmTasksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request AddmTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request AddmTasksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddmTasksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetAddmTasksSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAddmTasksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetAddmTasksSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddmTasksResponse wrapper for the AddmTasks operation
type AddmTasksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AddmTasksCollection instances
	AddmTasksCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response AddmTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response AddmTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// AddmTasksSortByEnum Enum with underlying type: string
type AddmTasksSortByEnum string

// Set of constants representing the allowable values for AddmTasksSortByEnum
const (
	AddmTasksSortByTaskName    AddmTasksSortByEnum = "TASK_NAME"
	AddmTasksSortByTaskId      AddmTasksSortByEnum = "TASK_ID"
	AddmTasksSortByDescription AddmTasksSortByEnum = "DESCRIPTION"
	AddmTasksSortByDbUser      AddmTasksSortByEnum = "DB_USER"
	AddmTasksSortByStatus      AddmTasksSortByEnum = "STATUS"
	AddmTasksSortByTimeCreated AddmTasksSortByEnum = "TIME_CREATED"
	AddmTasksSortByBeginTime   AddmTasksSortByEnum = "BEGIN_TIME"
	AddmTasksSortByEndTime     AddmTasksSortByEnum = "END_TIME"
)

var mappingAddmTasksSortByEnum = map[string]AddmTasksSortByEnum{
	"TASK_NAME":    AddmTasksSortByTaskName,
	"TASK_ID":      AddmTasksSortByTaskId,
	"DESCRIPTION":  AddmTasksSortByDescription,
	"DB_USER":      AddmTasksSortByDbUser,
	"STATUS":       AddmTasksSortByStatus,
	"TIME_CREATED": AddmTasksSortByTimeCreated,
	"BEGIN_TIME":   AddmTasksSortByBeginTime,
	"END_TIME":     AddmTasksSortByEndTime,
}

var mappingAddmTasksSortByEnumLowerCase = map[string]AddmTasksSortByEnum{
	"task_name":    AddmTasksSortByTaskName,
	"task_id":      AddmTasksSortByTaskId,
	"description":  AddmTasksSortByDescription,
	"db_user":      AddmTasksSortByDbUser,
	"status":       AddmTasksSortByStatus,
	"time_created": AddmTasksSortByTimeCreated,
	"begin_time":   AddmTasksSortByBeginTime,
	"end_time":     AddmTasksSortByEndTime,
}

// GetAddmTasksSortByEnumValues Enumerates the set of values for AddmTasksSortByEnum
func GetAddmTasksSortByEnumValues() []AddmTasksSortByEnum {
	values := make([]AddmTasksSortByEnum, 0)
	for _, v := range mappingAddmTasksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetAddmTasksSortByEnumStringValues Enumerates the set of values in String for AddmTasksSortByEnum
func GetAddmTasksSortByEnumStringValues() []string {
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

// GetMappingAddmTasksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddmTasksSortByEnum(val string) (AddmTasksSortByEnum, bool) {
	enum, ok := mappingAddmTasksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AddmTasksSortOrderEnum Enum with underlying type: string
type AddmTasksSortOrderEnum string

// Set of constants representing the allowable values for AddmTasksSortOrderEnum
const (
	AddmTasksSortOrderAsc  AddmTasksSortOrderEnum = "ASC"
	AddmTasksSortOrderDesc AddmTasksSortOrderEnum = "DESC"
)

var mappingAddmTasksSortOrderEnum = map[string]AddmTasksSortOrderEnum{
	"ASC":  AddmTasksSortOrderAsc,
	"DESC": AddmTasksSortOrderDesc,
}

var mappingAddmTasksSortOrderEnumLowerCase = map[string]AddmTasksSortOrderEnum{
	"asc":  AddmTasksSortOrderAsc,
	"desc": AddmTasksSortOrderDesc,
}

// GetAddmTasksSortOrderEnumValues Enumerates the set of values for AddmTasksSortOrderEnum
func GetAddmTasksSortOrderEnumValues() []AddmTasksSortOrderEnum {
	values := make([]AddmTasksSortOrderEnum, 0)
	for _, v := range mappingAddmTasksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetAddmTasksSortOrderEnumStringValues Enumerates the set of values in String for AddmTasksSortOrderEnum
func GetAddmTasksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingAddmTasksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddmTasksSortOrderEnum(val string) (AddmTasksSortOrderEnum, bool) {
	enum, ok := mappingAddmTasksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
