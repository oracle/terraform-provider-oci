// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListManagementAgentsRequest wrapper for the ListManagementAgents operation
type ListManagementAgentsRequest struct {

	// The ID of the compartment from which the Management Agents to be listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter to return only Management Agents having the particular Plugin installed.
	PluginName *string `mandatory:"false" contributesTo:"query" name:"pluginName"`

	// Filter to return only Management Agents having the particular agent version.
	Version *string `mandatory:"false" contributesTo:"query" name:"version"`

	// Filter to return only Management Agents having the particular display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter to return only Management Agents in the particular lifecycle state.
	LifecycleState ListManagementAgentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Filter to return only Management Agents having the particular platform type.
	PlatformType ListManagementAgentsPlatformTypeEnum `mandatory:"false" contributesTo:"query" name:"platformType" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementAgentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListManagementAgentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAgentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAgentsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementAgentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagementAgentsResponse wrapper for the ListManagementAgents operation
type ListManagementAgentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagementAgentSummary instances
	Items []ManagementAgentSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementAgentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementAgentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementAgentsLifecycleStateEnum Enum with underlying type: string
type ListManagementAgentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagementAgentsLifecycleStateEnum
const (
	ListManagementAgentsLifecycleStateCreating   ListManagementAgentsLifecycleStateEnum = "CREATING"
	ListManagementAgentsLifecycleStateUpdating   ListManagementAgentsLifecycleStateEnum = "UPDATING"
	ListManagementAgentsLifecycleStateActive     ListManagementAgentsLifecycleStateEnum = "ACTIVE"
	ListManagementAgentsLifecycleStateInactive   ListManagementAgentsLifecycleStateEnum = "INACTIVE"
	ListManagementAgentsLifecycleStateTerminated ListManagementAgentsLifecycleStateEnum = "TERMINATED"
	ListManagementAgentsLifecycleStateDeleting   ListManagementAgentsLifecycleStateEnum = "DELETING"
	ListManagementAgentsLifecycleStateDeleted    ListManagementAgentsLifecycleStateEnum = "DELETED"
	ListManagementAgentsLifecycleStateFailed     ListManagementAgentsLifecycleStateEnum = "FAILED"
)

var mappingListManagementAgentsLifecycleState = map[string]ListManagementAgentsLifecycleStateEnum{
	"CREATING":   ListManagementAgentsLifecycleStateCreating,
	"UPDATING":   ListManagementAgentsLifecycleStateUpdating,
	"ACTIVE":     ListManagementAgentsLifecycleStateActive,
	"INACTIVE":   ListManagementAgentsLifecycleStateInactive,
	"TERMINATED": ListManagementAgentsLifecycleStateTerminated,
	"DELETING":   ListManagementAgentsLifecycleStateDeleting,
	"DELETED":    ListManagementAgentsLifecycleStateDeleted,
	"FAILED":     ListManagementAgentsLifecycleStateFailed,
}

// GetListManagementAgentsLifecycleStateEnumValues Enumerates the set of values for ListManagementAgentsLifecycleStateEnum
func GetListManagementAgentsLifecycleStateEnumValues() []ListManagementAgentsLifecycleStateEnum {
	values := make([]ListManagementAgentsLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAgentsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentsPlatformTypeEnum Enum with underlying type: string
type ListManagementAgentsPlatformTypeEnum string

// Set of constants representing the allowable values for ListManagementAgentsPlatformTypeEnum
const (
	ListManagementAgentsPlatformTypeLinux   ListManagementAgentsPlatformTypeEnum = "LINUX"
	ListManagementAgentsPlatformTypeWindows ListManagementAgentsPlatformTypeEnum = "WINDOWS"
)

var mappingListManagementAgentsPlatformType = map[string]ListManagementAgentsPlatformTypeEnum{
	"LINUX":   ListManagementAgentsPlatformTypeLinux,
	"WINDOWS": ListManagementAgentsPlatformTypeWindows,
}

// GetListManagementAgentsPlatformTypeEnumValues Enumerates the set of values for ListManagementAgentsPlatformTypeEnum
func GetListManagementAgentsPlatformTypeEnumValues() []ListManagementAgentsPlatformTypeEnum {
	values := make([]ListManagementAgentsPlatformTypeEnum, 0)
	for _, v := range mappingListManagementAgentsPlatformType {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentsSortOrderEnum Enum with underlying type: string
type ListManagementAgentsSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAgentsSortOrderEnum
const (
	ListManagementAgentsSortOrderAsc  ListManagementAgentsSortOrderEnum = "ASC"
	ListManagementAgentsSortOrderDesc ListManagementAgentsSortOrderEnum = "DESC"
)

var mappingListManagementAgentsSortOrder = map[string]ListManagementAgentsSortOrderEnum{
	"ASC":  ListManagementAgentsSortOrderAsc,
	"DESC": ListManagementAgentsSortOrderDesc,
}

// GetListManagementAgentsSortOrderEnumValues Enumerates the set of values for ListManagementAgentsSortOrderEnum
func GetListManagementAgentsSortOrderEnumValues() []ListManagementAgentsSortOrderEnum {
	values := make([]ListManagementAgentsSortOrderEnum, 0)
	for _, v := range mappingListManagementAgentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentsSortByEnum Enum with underlying type: string
type ListManagementAgentsSortByEnum string

// Set of constants representing the allowable values for ListManagementAgentsSortByEnum
const (
	ListManagementAgentsSortByTimecreated ListManagementAgentsSortByEnum = "timeCreated"
	ListManagementAgentsSortByDisplayname ListManagementAgentsSortByEnum = "displayName"
)

var mappingListManagementAgentsSortBy = map[string]ListManagementAgentsSortByEnum{
	"timeCreated": ListManagementAgentsSortByTimecreated,
	"displayName": ListManagementAgentsSortByDisplayname,
}

// GetListManagementAgentsSortByEnumValues Enumerates the set of values for ListManagementAgentsSortByEnum
func GetListManagementAgentsSortByEnumValues() []ListManagementAgentsSortByEnum {
	values := make([]ListManagementAgentsSortByEnum, 0)
	for _, v := range mappingListManagementAgentsSortBy {
		values = append(values, v)
	}
	return values
}
