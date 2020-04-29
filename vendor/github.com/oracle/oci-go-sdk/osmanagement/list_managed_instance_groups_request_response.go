// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListManagedInstanceGroupsRequest wrapper for the ListManagedInstanceGroups operation
type ListManagedInstanceGroupsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagedInstanceGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListManagedInstanceGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The current lifecycle state for the object.
	LifecycleState ListManagedInstanceGroupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OS family for which to list resources.
	OsFamily ListManagedInstanceGroupsOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceGroupsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagedInstanceGroupsResponse wrapper for the ListManagedInstanceGroups operation
type ListManagedInstanceGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagedInstanceGroupSummary instances
	Items []ManagedInstanceGroupSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `ManagedInstanceGroup`s. If this
	// header appears in the response, then this is a partial list of
	// managed instance groups. Include this value as the `page`
	// parameter in a subsequent GET request to get the next batch of
	// managed instance groups.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceGroupsSortOrderEnum Enum with underlying type: string
type ListManagedInstanceGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsSortOrderEnum
const (
	ListManagedInstanceGroupsSortOrderAsc  ListManagedInstanceGroupsSortOrderEnum = "ASC"
	ListManagedInstanceGroupsSortOrderDesc ListManagedInstanceGroupsSortOrderEnum = "DESC"
)

var mappingListManagedInstanceGroupsSortOrder = map[string]ListManagedInstanceGroupsSortOrderEnum{
	"ASC":  ListManagedInstanceGroupsSortOrderAsc,
	"DESC": ListManagedInstanceGroupsSortOrderDesc,
}

// GetListManagedInstanceGroupsSortOrderEnumValues Enumerates the set of values for ListManagedInstanceGroupsSortOrderEnum
func GetListManagedInstanceGroupsSortOrderEnumValues() []ListManagedInstanceGroupsSortOrderEnum {
	values := make([]ListManagedInstanceGroupsSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagedInstanceGroupsSortByEnum Enum with underlying type: string
type ListManagedInstanceGroupsSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsSortByEnum
const (
	ListManagedInstanceGroupsSortByTimecreated ListManagedInstanceGroupsSortByEnum = "TIMECREATED"
	ListManagedInstanceGroupsSortByDisplayname ListManagedInstanceGroupsSortByEnum = "DISPLAYNAME"
)

var mappingListManagedInstanceGroupsSortBy = map[string]ListManagedInstanceGroupsSortByEnum{
	"TIMECREATED": ListManagedInstanceGroupsSortByTimecreated,
	"DISPLAYNAME": ListManagedInstanceGroupsSortByDisplayname,
}

// GetListManagedInstanceGroupsSortByEnumValues Enumerates the set of values for ListManagedInstanceGroupsSortByEnum
func GetListManagedInstanceGroupsSortByEnumValues() []ListManagedInstanceGroupsSortByEnum {
	values := make([]ListManagedInstanceGroupsSortByEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsSortBy {
		values = append(values, v)
	}
	return values
}

// ListManagedInstanceGroupsLifecycleStateEnum Enum with underlying type: string
type ListManagedInstanceGroupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsLifecycleStateEnum
const (
	ListManagedInstanceGroupsLifecycleStateCreating ListManagedInstanceGroupsLifecycleStateEnum = "CREATING"
	ListManagedInstanceGroupsLifecycleStateUpdating ListManagedInstanceGroupsLifecycleStateEnum = "UPDATING"
	ListManagedInstanceGroupsLifecycleStateActive   ListManagedInstanceGroupsLifecycleStateEnum = "ACTIVE"
	ListManagedInstanceGroupsLifecycleStateDeleting ListManagedInstanceGroupsLifecycleStateEnum = "DELETING"
	ListManagedInstanceGroupsLifecycleStateDeleted  ListManagedInstanceGroupsLifecycleStateEnum = "DELETED"
	ListManagedInstanceGroupsLifecycleStateFailed   ListManagedInstanceGroupsLifecycleStateEnum = "FAILED"
)

var mappingListManagedInstanceGroupsLifecycleState = map[string]ListManagedInstanceGroupsLifecycleStateEnum{
	"CREATING": ListManagedInstanceGroupsLifecycleStateCreating,
	"UPDATING": ListManagedInstanceGroupsLifecycleStateUpdating,
	"ACTIVE":   ListManagedInstanceGroupsLifecycleStateActive,
	"DELETING": ListManagedInstanceGroupsLifecycleStateDeleting,
	"DELETED":  ListManagedInstanceGroupsLifecycleStateDeleted,
	"FAILED":   ListManagedInstanceGroupsLifecycleStateFailed,
}

// GetListManagedInstanceGroupsLifecycleStateEnumValues Enumerates the set of values for ListManagedInstanceGroupsLifecycleStateEnum
func GetListManagedInstanceGroupsLifecycleStateEnumValues() []ListManagedInstanceGroupsLifecycleStateEnum {
	values := make([]ListManagedInstanceGroupsLifecycleStateEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListManagedInstanceGroupsOsFamilyEnum Enum with underlying type: string
type ListManagedInstanceGroupsOsFamilyEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsOsFamilyEnum
const (
	ListManagedInstanceGroupsOsFamilyLinux   ListManagedInstanceGroupsOsFamilyEnum = "LINUX"
	ListManagedInstanceGroupsOsFamilyWindows ListManagedInstanceGroupsOsFamilyEnum = "WINDOWS"
	ListManagedInstanceGroupsOsFamilyAll     ListManagedInstanceGroupsOsFamilyEnum = "ALL"
)

var mappingListManagedInstanceGroupsOsFamily = map[string]ListManagedInstanceGroupsOsFamilyEnum{
	"LINUX":   ListManagedInstanceGroupsOsFamilyLinux,
	"WINDOWS": ListManagedInstanceGroupsOsFamilyWindows,
	"ALL":     ListManagedInstanceGroupsOsFamilyAll,
}

// GetListManagedInstanceGroupsOsFamilyEnumValues Enumerates the set of values for ListManagedInstanceGroupsOsFamilyEnum
func GetListManagedInstanceGroupsOsFamilyEnumValues() []ListManagedInstanceGroupsOsFamilyEnum {
	values := make([]ListManagedInstanceGroupsOsFamilyEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsOsFamily {
		values = append(values, v)
	}
	return values
}
