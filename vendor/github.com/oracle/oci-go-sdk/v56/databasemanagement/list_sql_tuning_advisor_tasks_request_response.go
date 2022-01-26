// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListSqlTuningAdvisorTasksRequest wrapper for the ListSqlTuningAdvisorTasks operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListSqlTuningAdvisorTasks.go.html to see an example of how to use ListSqlTuningAdvisorTasksRequest.
type ListSqlTuningAdvisorTasksRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The optional query parameter to filter the SQL Tuning Advisor task list by name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The optional query parameter to filter the SQL Tuning Advisor task list by status.
	Status ListSqlTuningAdvisorTasksStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The optional greater than or equal to query parameter to filter the timestamp.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the SQL Tuning advisor task summary data.
	SortBy ListSqlTuningAdvisorTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Descending order is the default order.
	SortOrder ListSqlTuningAdvisorTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlTuningAdvisorTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlTuningAdvisorTasksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlTuningAdvisorTasksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlTuningAdvisorTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSqlTuningAdvisorTasksResponse wrapper for the ListSqlTuningAdvisorTasks operation
type ListSqlTuningAdvisorTasksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlTuningAdvisorTaskCollection instances
	SqlTuningAdvisorTaskCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSqlTuningAdvisorTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlTuningAdvisorTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlTuningAdvisorTasksStatusEnum Enum with underlying type: string
type ListSqlTuningAdvisorTasksStatusEnum string

// Set of constants representing the allowable values for ListSqlTuningAdvisorTasksStatusEnum
const (
	ListSqlTuningAdvisorTasksStatusInitial     ListSqlTuningAdvisorTasksStatusEnum = "INITIAL"
	ListSqlTuningAdvisorTasksStatusExecuting   ListSqlTuningAdvisorTasksStatusEnum = "EXECUTING"
	ListSqlTuningAdvisorTasksStatusInterrupted ListSqlTuningAdvisorTasksStatusEnum = "INTERRUPTED"
	ListSqlTuningAdvisorTasksStatusCompleted   ListSqlTuningAdvisorTasksStatusEnum = "COMPLETED"
	ListSqlTuningAdvisorTasksStatusError       ListSqlTuningAdvisorTasksStatusEnum = "ERROR"
)

var mappingListSqlTuningAdvisorTasksStatus = map[string]ListSqlTuningAdvisorTasksStatusEnum{
	"INITIAL":     ListSqlTuningAdvisorTasksStatusInitial,
	"EXECUTING":   ListSqlTuningAdvisorTasksStatusExecuting,
	"INTERRUPTED": ListSqlTuningAdvisorTasksStatusInterrupted,
	"COMPLETED":   ListSqlTuningAdvisorTasksStatusCompleted,
	"ERROR":       ListSqlTuningAdvisorTasksStatusError,
}

// GetListSqlTuningAdvisorTasksStatusEnumValues Enumerates the set of values for ListSqlTuningAdvisorTasksStatusEnum
func GetListSqlTuningAdvisorTasksStatusEnumValues() []ListSqlTuningAdvisorTasksStatusEnum {
	values := make([]ListSqlTuningAdvisorTasksStatusEnum, 0)
	for _, v := range mappingListSqlTuningAdvisorTasksStatus {
		values = append(values, v)
	}
	return values
}

// ListSqlTuningAdvisorTasksSortByEnum Enum with underlying type: string
type ListSqlTuningAdvisorTasksSortByEnum string

// Set of constants representing the allowable values for ListSqlTuningAdvisorTasksSortByEnum
const (
	ListSqlTuningAdvisorTasksSortByName      ListSqlTuningAdvisorTasksSortByEnum = "NAME"
	ListSqlTuningAdvisorTasksSortByStartTime ListSqlTuningAdvisorTasksSortByEnum = "START_TIME"
)

var mappingListSqlTuningAdvisorTasksSortBy = map[string]ListSqlTuningAdvisorTasksSortByEnum{
	"NAME":       ListSqlTuningAdvisorTasksSortByName,
	"START_TIME": ListSqlTuningAdvisorTasksSortByStartTime,
}

// GetListSqlTuningAdvisorTasksSortByEnumValues Enumerates the set of values for ListSqlTuningAdvisorTasksSortByEnum
func GetListSqlTuningAdvisorTasksSortByEnumValues() []ListSqlTuningAdvisorTasksSortByEnum {
	values := make([]ListSqlTuningAdvisorTasksSortByEnum, 0)
	for _, v := range mappingListSqlTuningAdvisorTasksSortBy {
		values = append(values, v)
	}
	return values
}

// ListSqlTuningAdvisorTasksSortOrderEnum Enum with underlying type: string
type ListSqlTuningAdvisorTasksSortOrderEnum string

// Set of constants representing the allowable values for ListSqlTuningAdvisorTasksSortOrderEnum
const (
	ListSqlTuningAdvisorTasksSortOrderAsc  ListSqlTuningAdvisorTasksSortOrderEnum = "ASC"
	ListSqlTuningAdvisorTasksSortOrderDesc ListSqlTuningAdvisorTasksSortOrderEnum = "DESC"
)

var mappingListSqlTuningAdvisorTasksSortOrder = map[string]ListSqlTuningAdvisorTasksSortOrderEnum{
	"ASC":  ListSqlTuningAdvisorTasksSortOrderAsc,
	"DESC": ListSqlTuningAdvisorTasksSortOrderDesc,
}

// GetListSqlTuningAdvisorTasksSortOrderEnumValues Enumerates the set of values for ListSqlTuningAdvisorTasksSortOrderEnum
func GetListSqlTuningAdvisorTasksSortOrderEnumValues() []ListSqlTuningAdvisorTasksSortOrderEnum {
	values := make([]ListSqlTuningAdvisorTasksSortOrderEnum, 0)
	for _, v := range mappingListSqlTuningAdvisorTasksSortOrder {
		values = append(values, v)
	}
	return values
}
