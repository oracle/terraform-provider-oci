// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListManagementAgentPluginsRequest wrapper for the ListManagementAgentPlugins operation
type ListManagementAgentPluginsRequest struct {

	// The ID of the compartment from which the Management Agents to be listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter to return only Management Agent Plugins having the particular display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementAgentPluginsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListManagementAgentPluginsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter to return only Management Agents in the particular lifecycle state.
	LifecycleState ListManagementAgentPluginsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAgentPluginsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAgentPluginsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementAgentPluginsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagementAgentPluginsResponse wrapper for the ListManagementAgentPlugins operation
type ListManagementAgentPluginsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagementAgentPluginSummary instances
	Items []ManagementAgentPluginSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagementAgentPluginsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementAgentPluginsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementAgentPluginsSortOrderEnum Enum with underlying type: string
type ListManagementAgentPluginsSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAgentPluginsSortOrderEnum
const (
	ListManagementAgentPluginsSortOrderAsc  ListManagementAgentPluginsSortOrderEnum = "ASC"
	ListManagementAgentPluginsSortOrderDesc ListManagementAgentPluginsSortOrderEnum = "DESC"
)

var mappingListManagementAgentPluginsSortOrder = map[string]ListManagementAgentPluginsSortOrderEnum{
	"ASC":  ListManagementAgentPluginsSortOrderAsc,
	"DESC": ListManagementAgentPluginsSortOrderDesc,
}

// GetListManagementAgentPluginsSortOrderEnumValues Enumerates the set of values for ListManagementAgentPluginsSortOrderEnum
func GetListManagementAgentPluginsSortOrderEnumValues() []ListManagementAgentPluginsSortOrderEnum {
	values := make([]ListManagementAgentPluginsSortOrderEnum, 0)
	for _, v := range mappingListManagementAgentPluginsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentPluginsSortByEnum Enum with underlying type: string
type ListManagementAgentPluginsSortByEnum string

// Set of constants representing the allowable values for ListManagementAgentPluginsSortByEnum
const (
	ListManagementAgentPluginsSortByDisplayname ListManagementAgentPluginsSortByEnum = "displayName"
)

var mappingListManagementAgentPluginsSortBy = map[string]ListManagementAgentPluginsSortByEnum{
	"displayName": ListManagementAgentPluginsSortByDisplayname,
}

// GetListManagementAgentPluginsSortByEnumValues Enumerates the set of values for ListManagementAgentPluginsSortByEnum
func GetListManagementAgentPluginsSortByEnumValues() []ListManagementAgentPluginsSortByEnum {
	values := make([]ListManagementAgentPluginsSortByEnum, 0)
	for _, v := range mappingListManagementAgentPluginsSortBy {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentPluginsLifecycleStateEnum Enum with underlying type: string
type ListManagementAgentPluginsLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagementAgentPluginsLifecycleStateEnum
const (
	ListManagementAgentPluginsLifecycleStateCreating   ListManagementAgentPluginsLifecycleStateEnum = "CREATING"
	ListManagementAgentPluginsLifecycleStateUpdating   ListManagementAgentPluginsLifecycleStateEnum = "UPDATING"
	ListManagementAgentPluginsLifecycleStateActive     ListManagementAgentPluginsLifecycleStateEnum = "ACTIVE"
	ListManagementAgentPluginsLifecycleStateInactive   ListManagementAgentPluginsLifecycleStateEnum = "INACTIVE"
	ListManagementAgentPluginsLifecycleStateTerminated ListManagementAgentPluginsLifecycleStateEnum = "TERMINATED"
	ListManagementAgentPluginsLifecycleStateDeleting   ListManagementAgentPluginsLifecycleStateEnum = "DELETING"
	ListManagementAgentPluginsLifecycleStateDeleted    ListManagementAgentPluginsLifecycleStateEnum = "DELETED"
	ListManagementAgentPluginsLifecycleStateFailed     ListManagementAgentPluginsLifecycleStateEnum = "FAILED"
)

var mappingListManagementAgentPluginsLifecycleState = map[string]ListManagementAgentPluginsLifecycleStateEnum{
	"CREATING":   ListManagementAgentPluginsLifecycleStateCreating,
	"UPDATING":   ListManagementAgentPluginsLifecycleStateUpdating,
	"ACTIVE":     ListManagementAgentPluginsLifecycleStateActive,
	"INACTIVE":   ListManagementAgentPluginsLifecycleStateInactive,
	"TERMINATED": ListManagementAgentPluginsLifecycleStateTerminated,
	"DELETING":   ListManagementAgentPluginsLifecycleStateDeleting,
	"DELETED":    ListManagementAgentPluginsLifecycleStateDeleted,
	"FAILED":     ListManagementAgentPluginsLifecycleStateFailed,
}

// GetListManagementAgentPluginsLifecycleStateEnumValues Enumerates the set of values for ListManagementAgentPluginsLifecycleStateEnum
func GetListManagementAgentPluginsLifecycleStateEnumValues() []ListManagementAgentPluginsLifecycleStateEnum {
	values := make([]ListManagementAgentPluginsLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAgentPluginsLifecycleState {
		values = append(values, v)
	}
	return values
}
