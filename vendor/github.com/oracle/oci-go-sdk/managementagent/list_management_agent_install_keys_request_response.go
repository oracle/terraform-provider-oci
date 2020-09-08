// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListManagementAgentInstallKeysRequest wrapper for the ListManagementAgentInstallKeys operation
type ListManagementAgentInstallKeysRequest struct {

	// The ID of the compartment from which the Management Agents to be listed.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// if set to true then it fetches install key for all compartments where user has access to else only on the compartment specified.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Value of this is always "ACCESSIBLE" and any other value is not supported.
	AccessLevel *string `mandatory:"false" contributesTo:"query" name:"accessLevel"`

	// Filter to return only Management Agents in the particular lifecycle state.
	LifecycleState ListManagementAgentInstallKeysLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The display name for which the Key needs to be listed.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagementAgentInstallKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListManagementAgentInstallKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagementAgentInstallKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagementAgentInstallKeysRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagementAgentInstallKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagementAgentInstallKeysResponse wrapper for the ListManagementAgentInstallKeys operation
type ListManagementAgentInstallKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagementAgentInstallKeySummary instances
	Items []ManagementAgentInstallKeySummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListManagementAgentInstallKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagementAgentInstallKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagementAgentInstallKeysLifecycleStateEnum Enum with underlying type: string
type ListManagementAgentInstallKeysLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagementAgentInstallKeysLifecycleStateEnum
const (
	ListManagementAgentInstallKeysLifecycleStateCreating   ListManagementAgentInstallKeysLifecycleStateEnum = "CREATING"
	ListManagementAgentInstallKeysLifecycleStateUpdating   ListManagementAgentInstallKeysLifecycleStateEnum = "UPDATING"
	ListManagementAgentInstallKeysLifecycleStateActive     ListManagementAgentInstallKeysLifecycleStateEnum = "ACTIVE"
	ListManagementAgentInstallKeysLifecycleStateInactive   ListManagementAgentInstallKeysLifecycleStateEnum = "INACTIVE"
	ListManagementAgentInstallKeysLifecycleStateTerminated ListManagementAgentInstallKeysLifecycleStateEnum = "TERMINATED"
	ListManagementAgentInstallKeysLifecycleStateDeleting   ListManagementAgentInstallKeysLifecycleStateEnum = "DELETING"
	ListManagementAgentInstallKeysLifecycleStateDeleted    ListManagementAgentInstallKeysLifecycleStateEnum = "DELETED"
	ListManagementAgentInstallKeysLifecycleStateFailed     ListManagementAgentInstallKeysLifecycleStateEnum = "FAILED"
)

var mappingListManagementAgentInstallKeysLifecycleState = map[string]ListManagementAgentInstallKeysLifecycleStateEnum{
	"CREATING":   ListManagementAgentInstallKeysLifecycleStateCreating,
	"UPDATING":   ListManagementAgentInstallKeysLifecycleStateUpdating,
	"ACTIVE":     ListManagementAgentInstallKeysLifecycleStateActive,
	"INACTIVE":   ListManagementAgentInstallKeysLifecycleStateInactive,
	"TERMINATED": ListManagementAgentInstallKeysLifecycleStateTerminated,
	"DELETING":   ListManagementAgentInstallKeysLifecycleStateDeleting,
	"DELETED":    ListManagementAgentInstallKeysLifecycleStateDeleted,
	"FAILED":     ListManagementAgentInstallKeysLifecycleStateFailed,
}

// GetListManagementAgentInstallKeysLifecycleStateEnumValues Enumerates the set of values for ListManagementAgentInstallKeysLifecycleStateEnum
func GetListManagementAgentInstallKeysLifecycleStateEnumValues() []ListManagementAgentInstallKeysLifecycleStateEnum {
	values := make([]ListManagementAgentInstallKeysLifecycleStateEnum, 0)
	for _, v := range mappingListManagementAgentInstallKeysLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentInstallKeysSortOrderEnum Enum with underlying type: string
type ListManagementAgentInstallKeysSortOrderEnum string

// Set of constants representing the allowable values for ListManagementAgentInstallKeysSortOrderEnum
const (
	ListManagementAgentInstallKeysSortOrderAsc  ListManagementAgentInstallKeysSortOrderEnum = "ASC"
	ListManagementAgentInstallKeysSortOrderDesc ListManagementAgentInstallKeysSortOrderEnum = "DESC"
)

var mappingListManagementAgentInstallKeysSortOrder = map[string]ListManagementAgentInstallKeysSortOrderEnum{
	"ASC":  ListManagementAgentInstallKeysSortOrderAsc,
	"DESC": ListManagementAgentInstallKeysSortOrderDesc,
}

// GetListManagementAgentInstallKeysSortOrderEnumValues Enumerates the set of values for ListManagementAgentInstallKeysSortOrderEnum
func GetListManagementAgentInstallKeysSortOrderEnumValues() []ListManagementAgentInstallKeysSortOrderEnum {
	values := make([]ListManagementAgentInstallKeysSortOrderEnum, 0)
	for _, v := range mappingListManagementAgentInstallKeysSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagementAgentInstallKeysSortByEnum Enum with underlying type: string
type ListManagementAgentInstallKeysSortByEnum string

// Set of constants representing the allowable values for ListManagementAgentInstallKeysSortByEnum
const (
	ListManagementAgentInstallKeysSortByTimecreated ListManagementAgentInstallKeysSortByEnum = "timeCreated"
	ListManagementAgentInstallKeysSortByDisplayname ListManagementAgentInstallKeysSortByEnum = "displayName"
)

var mappingListManagementAgentInstallKeysSortBy = map[string]ListManagementAgentInstallKeysSortByEnum{
	"timeCreated": ListManagementAgentInstallKeysSortByTimecreated,
	"displayName": ListManagementAgentInstallKeysSortByDisplayname,
}

// GetListManagementAgentInstallKeysSortByEnumValues Enumerates the set of values for ListManagementAgentInstallKeysSortByEnum
func GetListManagementAgentInstallKeysSortByEnumValues() []ListManagementAgentInstallKeysSortByEnum {
	values := make([]ListManagementAgentInstallKeysSortByEnum, 0)
	for _, v := range mappingListManagementAgentInstallKeysSortBy {
		values = append(values, v)
	}
	return values
}
