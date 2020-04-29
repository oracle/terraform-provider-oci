// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListUpcomingScheduledJobsRequest wrapper for the ListUpcomingScheduledJobs operation
type ListUpcomingScheduledJobsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The cut-off time before which to list all upcoming schedules, in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	TimeEnd *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeEnd"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListUpcomingScheduledJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListUpcomingScheduledJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The name of the tag.
	TagName *string `mandatory:"false" contributesTo:"query" name:"tagName"`

	// The value for the tag.
	TagValue *string `mandatory:"false" contributesTo:"query" name:"tagValue"`

	// The current lifecycle state for the object.
	LifecycleState ListUpcomingScheduledJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OS family for which to list resources.
	OsFamily ListUpcomingScheduledJobsOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUpcomingScheduledJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUpcomingScheduledJobsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUpcomingScheduledJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListUpcomingScheduledJobsResponse wrapper for the ListUpcomingScheduledJobs operation
type ListUpcomingScheduledJobsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ScheduledJobSummary instances
	Items []ScheduledJobSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListUpcomingScheduledJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUpcomingScheduledJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUpcomingScheduledJobsSortOrderEnum Enum with underlying type: string
type ListUpcomingScheduledJobsSortOrderEnum string

// Set of constants representing the allowable values for ListUpcomingScheduledJobsSortOrderEnum
const (
	ListUpcomingScheduledJobsSortOrderAsc  ListUpcomingScheduledJobsSortOrderEnum = "ASC"
	ListUpcomingScheduledJobsSortOrderDesc ListUpcomingScheduledJobsSortOrderEnum = "DESC"
)

var mappingListUpcomingScheduledJobsSortOrder = map[string]ListUpcomingScheduledJobsSortOrderEnum{
	"ASC":  ListUpcomingScheduledJobsSortOrderAsc,
	"DESC": ListUpcomingScheduledJobsSortOrderDesc,
}

// GetListUpcomingScheduledJobsSortOrderEnumValues Enumerates the set of values for ListUpcomingScheduledJobsSortOrderEnum
func GetListUpcomingScheduledJobsSortOrderEnumValues() []ListUpcomingScheduledJobsSortOrderEnum {
	values := make([]ListUpcomingScheduledJobsSortOrderEnum, 0)
	for _, v := range mappingListUpcomingScheduledJobsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListUpcomingScheduledJobsSortByEnum Enum with underlying type: string
type ListUpcomingScheduledJobsSortByEnum string

// Set of constants representing the allowable values for ListUpcomingScheduledJobsSortByEnum
const (
	ListUpcomingScheduledJobsSortByTimecreated ListUpcomingScheduledJobsSortByEnum = "TIMECREATED"
	ListUpcomingScheduledJobsSortByDisplayname ListUpcomingScheduledJobsSortByEnum = "DISPLAYNAME"
)

var mappingListUpcomingScheduledJobsSortBy = map[string]ListUpcomingScheduledJobsSortByEnum{
	"TIMECREATED": ListUpcomingScheduledJobsSortByTimecreated,
	"DISPLAYNAME": ListUpcomingScheduledJobsSortByDisplayname,
}

// GetListUpcomingScheduledJobsSortByEnumValues Enumerates the set of values for ListUpcomingScheduledJobsSortByEnum
func GetListUpcomingScheduledJobsSortByEnumValues() []ListUpcomingScheduledJobsSortByEnum {
	values := make([]ListUpcomingScheduledJobsSortByEnum, 0)
	for _, v := range mappingListUpcomingScheduledJobsSortBy {
		values = append(values, v)
	}
	return values
}

// ListUpcomingScheduledJobsLifecycleStateEnum Enum with underlying type: string
type ListUpcomingScheduledJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListUpcomingScheduledJobsLifecycleStateEnum
const (
	ListUpcomingScheduledJobsLifecycleStateCreating ListUpcomingScheduledJobsLifecycleStateEnum = "CREATING"
	ListUpcomingScheduledJobsLifecycleStateUpdating ListUpcomingScheduledJobsLifecycleStateEnum = "UPDATING"
	ListUpcomingScheduledJobsLifecycleStateActive   ListUpcomingScheduledJobsLifecycleStateEnum = "ACTIVE"
	ListUpcomingScheduledJobsLifecycleStateDeleting ListUpcomingScheduledJobsLifecycleStateEnum = "DELETING"
	ListUpcomingScheduledJobsLifecycleStateDeleted  ListUpcomingScheduledJobsLifecycleStateEnum = "DELETED"
	ListUpcomingScheduledJobsLifecycleStateFailed   ListUpcomingScheduledJobsLifecycleStateEnum = "FAILED"
)

var mappingListUpcomingScheduledJobsLifecycleState = map[string]ListUpcomingScheduledJobsLifecycleStateEnum{
	"CREATING": ListUpcomingScheduledJobsLifecycleStateCreating,
	"UPDATING": ListUpcomingScheduledJobsLifecycleStateUpdating,
	"ACTIVE":   ListUpcomingScheduledJobsLifecycleStateActive,
	"DELETING": ListUpcomingScheduledJobsLifecycleStateDeleting,
	"DELETED":  ListUpcomingScheduledJobsLifecycleStateDeleted,
	"FAILED":   ListUpcomingScheduledJobsLifecycleStateFailed,
}

// GetListUpcomingScheduledJobsLifecycleStateEnumValues Enumerates the set of values for ListUpcomingScheduledJobsLifecycleStateEnum
func GetListUpcomingScheduledJobsLifecycleStateEnumValues() []ListUpcomingScheduledJobsLifecycleStateEnum {
	values := make([]ListUpcomingScheduledJobsLifecycleStateEnum, 0)
	for _, v := range mappingListUpcomingScheduledJobsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListUpcomingScheduledJobsOsFamilyEnum Enum with underlying type: string
type ListUpcomingScheduledJobsOsFamilyEnum string

// Set of constants representing the allowable values for ListUpcomingScheduledJobsOsFamilyEnum
const (
	ListUpcomingScheduledJobsOsFamilyLinux   ListUpcomingScheduledJobsOsFamilyEnum = "LINUX"
	ListUpcomingScheduledJobsOsFamilyWindows ListUpcomingScheduledJobsOsFamilyEnum = "WINDOWS"
	ListUpcomingScheduledJobsOsFamilyAll     ListUpcomingScheduledJobsOsFamilyEnum = "ALL"
)

var mappingListUpcomingScheduledJobsOsFamily = map[string]ListUpcomingScheduledJobsOsFamilyEnum{
	"LINUX":   ListUpcomingScheduledJobsOsFamilyLinux,
	"WINDOWS": ListUpcomingScheduledJobsOsFamilyWindows,
	"ALL":     ListUpcomingScheduledJobsOsFamilyAll,
}

// GetListUpcomingScheduledJobsOsFamilyEnumValues Enumerates the set of values for ListUpcomingScheduledJobsOsFamilyEnum
func GetListUpcomingScheduledJobsOsFamilyEnumValues() []ListUpcomingScheduledJobsOsFamilyEnum {
	values := make([]ListUpcomingScheduledJobsOsFamilyEnum, 0)
	for _, v := range mappingListUpcomingScheduledJobsOsFamily {
		values = append(values, v)
	}
	return values
}
