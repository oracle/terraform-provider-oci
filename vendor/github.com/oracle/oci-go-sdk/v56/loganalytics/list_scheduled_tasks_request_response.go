// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListScheduledTasksRequest wrapper for the ListScheduledTasks operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListScheduledTasks.go.html to see an example of how to use ListScheduledTasksRequest.
type ListScheduledTasksRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Required parameter to specify schedule task type.
	TaskType ListScheduledTasksTaskTypeEnum `mandatory:"true" contributesTo:"query" name:"taskType" omitEmpty:"true"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListScheduledTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListScheduledTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only scheduled tasks whose stream action savedSearchId matches the given
	// ManagementSavedSearch id [OCID] exactly.
	SavedSearchId *string `mandatory:"false" contributesTo:"query" name:"savedSearchId"`

	// A filter to return only resources whose display name contains the substring.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScheduledTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScheduledTasksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListScheduledTasksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScheduledTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListScheduledTasksResponse wrapper for the ListScheduledTasks operation
type ListScheduledTasksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScheduledTaskCollection instances
	ScheduledTaskCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListScheduledTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScheduledTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScheduledTasksTaskTypeEnum Enum with underlying type: string
type ListScheduledTasksTaskTypeEnum string

// Set of constants representing the allowable values for ListScheduledTasksTaskTypeEnum
const (
	ListScheduledTasksTaskTypeSavedSearch             ListScheduledTasksTaskTypeEnum = "SAVED_SEARCH"
	ListScheduledTasksTaskTypeAcceleration            ListScheduledTasksTaskTypeEnum = "ACCELERATION"
	ListScheduledTasksTaskTypePurge                   ListScheduledTasksTaskTypeEnum = "PURGE"
	ListScheduledTasksTaskTypeAccelerationMaintenance ListScheduledTasksTaskTypeEnum = "ACCELERATION_MAINTENANCE"
)

var mappingListScheduledTasksTaskType = map[string]ListScheduledTasksTaskTypeEnum{
	"SAVED_SEARCH":             ListScheduledTasksTaskTypeSavedSearch,
	"ACCELERATION":             ListScheduledTasksTaskTypeAcceleration,
	"PURGE":                    ListScheduledTasksTaskTypePurge,
	"ACCELERATION_MAINTENANCE": ListScheduledTasksTaskTypeAccelerationMaintenance,
}

// GetListScheduledTasksTaskTypeEnumValues Enumerates the set of values for ListScheduledTasksTaskTypeEnum
func GetListScheduledTasksTaskTypeEnumValues() []ListScheduledTasksTaskTypeEnum {
	values := make([]ListScheduledTasksTaskTypeEnum, 0)
	for _, v := range mappingListScheduledTasksTaskType {
		values = append(values, v)
	}
	return values
}

// ListScheduledTasksSortOrderEnum Enum with underlying type: string
type ListScheduledTasksSortOrderEnum string

// Set of constants representing the allowable values for ListScheduledTasksSortOrderEnum
const (
	ListScheduledTasksSortOrderAsc  ListScheduledTasksSortOrderEnum = "ASC"
	ListScheduledTasksSortOrderDesc ListScheduledTasksSortOrderEnum = "DESC"
)

var mappingListScheduledTasksSortOrder = map[string]ListScheduledTasksSortOrderEnum{
	"ASC":  ListScheduledTasksSortOrderAsc,
	"DESC": ListScheduledTasksSortOrderDesc,
}

// GetListScheduledTasksSortOrderEnumValues Enumerates the set of values for ListScheduledTasksSortOrderEnum
func GetListScheduledTasksSortOrderEnumValues() []ListScheduledTasksSortOrderEnum {
	values := make([]ListScheduledTasksSortOrderEnum, 0)
	for _, v := range mappingListScheduledTasksSortOrder {
		values = append(values, v)
	}
	return values
}

// ListScheduledTasksSortByEnum Enum with underlying type: string
type ListScheduledTasksSortByEnum string

// Set of constants representing the allowable values for ListScheduledTasksSortByEnum
const (
	ListScheduledTasksSortByTimecreated ListScheduledTasksSortByEnum = "timeCreated"
	ListScheduledTasksSortByTimeupdated ListScheduledTasksSortByEnum = "timeUpdated"
	ListScheduledTasksSortByDisplayname ListScheduledTasksSortByEnum = "displayName"
)

var mappingListScheduledTasksSortBy = map[string]ListScheduledTasksSortByEnum{
	"timeCreated": ListScheduledTasksSortByTimecreated,
	"timeUpdated": ListScheduledTasksSortByTimeupdated,
	"displayName": ListScheduledTasksSortByDisplayname,
}

// GetListScheduledTasksSortByEnumValues Enumerates the set of values for ListScheduledTasksSortByEnum
func GetListScheduledTasksSortByEnumValues() []ListScheduledTasksSortByEnum {
	values := make([]ListScheduledTasksSortByEnum, 0)
	for _, v := range mappingListScheduledTasksSortBy {
		values = append(values, v)
	}
	return values
}
