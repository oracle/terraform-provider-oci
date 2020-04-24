// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListScheduledJobsRequest wrapper for the ListScheduledJobs operation
type ListScheduledJobsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the managed instance for which to list resources.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The ID of the managed instace group for which to list resources.
	ManagedInstanceGroupId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceGroupId"`

	// The operation type for which to list resources
	OperationType ListScheduledJobsOperationTypeEnum `mandatory:"false" contributesTo:"query" name:"operationType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListScheduledJobsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListScheduledJobsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The current lifecycle state for the object.
	LifecycleState ListScheduledJobsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OS family for which to list resources.
	OsFamily ListScheduledJobsOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListScheduledJobsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListScheduledJobsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListScheduledJobsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListScheduledJobsResponse wrapper for the ListScheduledJobs operation
type ListScheduledJobsResponse struct {

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

func (response ListScheduledJobsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListScheduledJobsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListScheduledJobsOperationTypeEnum Enum with underlying type: string
type ListScheduledJobsOperationTypeEnum string

// Set of constants representing the allowable values for ListScheduledJobsOperationTypeEnum
const (
	ListScheduledJobsOperationTypeInstall   ListScheduledJobsOperationTypeEnum = "INSTALL"
	ListScheduledJobsOperationTypeUpdate    ListScheduledJobsOperationTypeEnum = "UPDATE"
	ListScheduledJobsOperationTypeRemove    ListScheduledJobsOperationTypeEnum = "REMOVE"
	ListScheduledJobsOperationTypeUpdateall ListScheduledJobsOperationTypeEnum = "UPDATEALL"
)

var mappingListScheduledJobsOperationType = map[string]ListScheduledJobsOperationTypeEnum{
	"INSTALL":   ListScheduledJobsOperationTypeInstall,
	"UPDATE":    ListScheduledJobsOperationTypeUpdate,
	"REMOVE":    ListScheduledJobsOperationTypeRemove,
	"UPDATEALL": ListScheduledJobsOperationTypeUpdateall,
}

// GetListScheduledJobsOperationTypeEnumValues Enumerates the set of values for ListScheduledJobsOperationTypeEnum
func GetListScheduledJobsOperationTypeEnumValues() []ListScheduledJobsOperationTypeEnum {
	values := make([]ListScheduledJobsOperationTypeEnum, 0)
	for _, v := range mappingListScheduledJobsOperationType {
		values = append(values, v)
	}
	return values
}

// ListScheduledJobsSortOrderEnum Enum with underlying type: string
type ListScheduledJobsSortOrderEnum string

// Set of constants representing the allowable values for ListScheduledJobsSortOrderEnum
const (
	ListScheduledJobsSortOrderAsc  ListScheduledJobsSortOrderEnum = "ASC"
	ListScheduledJobsSortOrderDesc ListScheduledJobsSortOrderEnum = "DESC"
)

var mappingListScheduledJobsSortOrder = map[string]ListScheduledJobsSortOrderEnum{
	"ASC":  ListScheduledJobsSortOrderAsc,
	"DESC": ListScheduledJobsSortOrderDesc,
}

// GetListScheduledJobsSortOrderEnumValues Enumerates the set of values for ListScheduledJobsSortOrderEnum
func GetListScheduledJobsSortOrderEnumValues() []ListScheduledJobsSortOrderEnum {
	values := make([]ListScheduledJobsSortOrderEnum, 0)
	for _, v := range mappingListScheduledJobsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListScheduledJobsSortByEnum Enum with underlying type: string
type ListScheduledJobsSortByEnum string

// Set of constants representing the allowable values for ListScheduledJobsSortByEnum
const (
	ListScheduledJobsSortByTimecreated ListScheduledJobsSortByEnum = "TIMECREATED"
	ListScheduledJobsSortByDisplayname ListScheduledJobsSortByEnum = "DISPLAYNAME"
)

var mappingListScheduledJobsSortBy = map[string]ListScheduledJobsSortByEnum{
	"TIMECREATED": ListScheduledJobsSortByTimecreated,
	"DISPLAYNAME": ListScheduledJobsSortByDisplayname,
}

// GetListScheduledJobsSortByEnumValues Enumerates the set of values for ListScheduledJobsSortByEnum
func GetListScheduledJobsSortByEnumValues() []ListScheduledJobsSortByEnum {
	values := make([]ListScheduledJobsSortByEnum, 0)
	for _, v := range mappingListScheduledJobsSortBy {
		values = append(values, v)
	}
	return values
}

// ListScheduledJobsLifecycleStateEnum Enum with underlying type: string
type ListScheduledJobsLifecycleStateEnum string

// Set of constants representing the allowable values for ListScheduledJobsLifecycleStateEnum
const (
	ListScheduledJobsLifecycleStateCreating ListScheduledJobsLifecycleStateEnum = "CREATING"
	ListScheduledJobsLifecycleStateUpdating ListScheduledJobsLifecycleStateEnum = "UPDATING"
	ListScheduledJobsLifecycleStateActive   ListScheduledJobsLifecycleStateEnum = "ACTIVE"
	ListScheduledJobsLifecycleStateDeleting ListScheduledJobsLifecycleStateEnum = "DELETING"
	ListScheduledJobsLifecycleStateDeleted  ListScheduledJobsLifecycleStateEnum = "DELETED"
	ListScheduledJobsLifecycleStateFailed   ListScheduledJobsLifecycleStateEnum = "FAILED"
)

var mappingListScheduledJobsLifecycleState = map[string]ListScheduledJobsLifecycleStateEnum{
	"CREATING": ListScheduledJobsLifecycleStateCreating,
	"UPDATING": ListScheduledJobsLifecycleStateUpdating,
	"ACTIVE":   ListScheduledJobsLifecycleStateActive,
	"DELETING": ListScheduledJobsLifecycleStateDeleting,
	"DELETED":  ListScheduledJobsLifecycleStateDeleted,
	"FAILED":   ListScheduledJobsLifecycleStateFailed,
}

// GetListScheduledJobsLifecycleStateEnumValues Enumerates the set of values for ListScheduledJobsLifecycleStateEnum
func GetListScheduledJobsLifecycleStateEnumValues() []ListScheduledJobsLifecycleStateEnum {
	values := make([]ListScheduledJobsLifecycleStateEnum, 0)
	for _, v := range mappingListScheduledJobsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListScheduledJobsOsFamilyEnum Enum with underlying type: string
type ListScheduledJobsOsFamilyEnum string

// Set of constants representing the allowable values for ListScheduledJobsOsFamilyEnum
const (
	ListScheduledJobsOsFamilyLinux   ListScheduledJobsOsFamilyEnum = "LINUX"
	ListScheduledJobsOsFamilyWindows ListScheduledJobsOsFamilyEnum = "WINDOWS"
	ListScheduledJobsOsFamilyAll     ListScheduledJobsOsFamilyEnum = "ALL"
)

var mappingListScheduledJobsOsFamily = map[string]ListScheduledJobsOsFamilyEnum{
	"LINUX":   ListScheduledJobsOsFamilyLinux,
	"WINDOWS": ListScheduledJobsOsFamilyWindows,
	"ALL":     ListScheduledJobsOsFamilyAll,
}

// GetListScheduledJobsOsFamilyEnumValues Enumerates the set of values for ListScheduledJobsOsFamilyEnum
func GetListScheduledJobsOsFamilyEnumValues() []ListScheduledJobsOsFamilyEnum {
	values := make([]ListScheduledJobsOsFamilyEnum, 0)
	for _, v := range mappingListScheduledJobsOsFamily {
		values = append(values, v)
	}
	return values
}
