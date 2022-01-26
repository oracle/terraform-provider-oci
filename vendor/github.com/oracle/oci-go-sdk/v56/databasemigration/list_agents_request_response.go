// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListAgentsRequest wrapper for the ListAgents operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListAgents.go.html to see an example of how to use ListAgentsRequest.
type ListAgentsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAgentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAgentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The current state of the Database Migration Deployment.
	LifecycleState ListAgentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAgentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAgentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAgentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAgentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAgentsResponse wrapper for the ListAgents operation
type ListAgentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AgentCollection instances
	AgentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAgentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAgentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAgentsSortByEnum Enum with underlying type: string
type ListAgentsSortByEnum string

// Set of constants representing the allowable values for ListAgentsSortByEnum
const (
	ListAgentsSortByTimecreated ListAgentsSortByEnum = "timeCreated"
	ListAgentsSortByDisplayname ListAgentsSortByEnum = "displayName"
)

var mappingListAgentsSortBy = map[string]ListAgentsSortByEnum{
	"timeCreated": ListAgentsSortByTimecreated,
	"displayName": ListAgentsSortByDisplayname,
}

// GetListAgentsSortByEnumValues Enumerates the set of values for ListAgentsSortByEnum
func GetListAgentsSortByEnumValues() []ListAgentsSortByEnum {
	values := make([]ListAgentsSortByEnum, 0)
	for _, v := range mappingListAgentsSortBy {
		values = append(values, v)
	}
	return values
}

// ListAgentsSortOrderEnum Enum with underlying type: string
type ListAgentsSortOrderEnum string

// Set of constants representing the allowable values for ListAgentsSortOrderEnum
const (
	ListAgentsSortOrderAsc  ListAgentsSortOrderEnum = "ASC"
	ListAgentsSortOrderDesc ListAgentsSortOrderEnum = "DESC"
)

var mappingListAgentsSortOrder = map[string]ListAgentsSortOrderEnum{
	"ASC":  ListAgentsSortOrderAsc,
	"DESC": ListAgentsSortOrderDesc,
}

// GetListAgentsSortOrderEnumValues Enumerates the set of values for ListAgentsSortOrderEnum
func GetListAgentsSortOrderEnumValues() []ListAgentsSortOrderEnum {
	values := make([]ListAgentsSortOrderEnum, 0)
	for _, v := range mappingListAgentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAgentsLifecycleStateEnum Enum with underlying type: string
type ListAgentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAgentsLifecycleStateEnum
const (
	ListAgentsLifecycleStateCreating ListAgentsLifecycleStateEnum = "CREATING"
	ListAgentsLifecycleStateUpdating ListAgentsLifecycleStateEnum = "UPDATING"
	ListAgentsLifecycleStateActive   ListAgentsLifecycleStateEnum = "ACTIVE"
	ListAgentsLifecycleStateInactive ListAgentsLifecycleStateEnum = "INACTIVE"
	ListAgentsLifecycleStateDeleting ListAgentsLifecycleStateEnum = "DELETING"
	ListAgentsLifecycleStateDeleted  ListAgentsLifecycleStateEnum = "DELETED"
	ListAgentsLifecycleStateFailed   ListAgentsLifecycleStateEnum = "FAILED"
)

var mappingListAgentsLifecycleState = map[string]ListAgentsLifecycleStateEnum{
	"CREATING": ListAgentsLifecycleStateCreating,
	"UPDATING": ListAgentsLifecycleStateUpdating,
	"ACTIVE":   ListAgentsLifecycleStateActive,
	"INACTIVE": ListAgentsLifecycleStateInactive,
	"DELETING": ListAgentsLifecycleStateDeleting,
	"DELETED":  ListAgentsLifecycleStateDeleted,
	"FAILED":   ListAgentsLifecycleStateFailed,
}

// GetListAgentsLifecycleStateEnumValues Enumerates the set of values for ListAgentsLifecycleStateEnum
func GetListAgentsLifecycleStateEnumValues() []ListAgentsLifecycleStateEnum {
	values := make([]ListAgentsLifecycleStateEnum, 0)
	for _, v := range mappingListAgentsLifecycleState {
		values = append(values, v)
	}
	return values
}
