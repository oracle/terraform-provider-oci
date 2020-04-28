// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package integration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListIntegrationInstancesRequest wrapper for the ListIntegrationInstances operation
type ListIntegrationInstancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Life cycle state to query on.
	LifecycleState ListIntegrationInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListIntegrationInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order
	// for TIMECREATED is descending. Default order for DISPLAYNAME is
	// ascending. If no value is specified TIMECREATED is default.
	SortBy ListIntegrationInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIntegrationInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIntegrationInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIntegrationInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListIntegrationInstancesResponse wrapper for the ListIntegrationInstances operation
type ListIntegrationInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []IntegrationInstanceSummary instances
	Items []IntegrationInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, additional pages of results have been previously returned
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`
}

func (response ListIntegrationInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIntegrationInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIntegrationInstancesLifecycleStateEnum Enum with underlying type: string
type ListIntegrationInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListIntegrationInstancesLifecycleStateEnum
const (
	ListIntegrationInstancesLifecycleStateCreating ListIntegrationInstancesLifecycleStateEnum = "CREATING"
	ListIntegrationInstancesLifecycleStateUpdating ListIntegrationInstancesLifecycleStateEnum = "UPDATING"
	ListIntegrationInstancesLifecycleStateActive   ListIntegrationInstancesLifecycleStateEnum = "ACTIVE"
	ListIntegrationInstancesLifecycleStateInactive ListIntegrationInstancesLifecycleStateEnum = "INACTIVE"
	ListIntegrationInstancesLifecycleStateDeleting ListIntegrationInstancesLifecycleStateEnum = "DELETING"
	ListIntegrationInstancesLifecycleStateDeleted  ListIntegrationInstancesLifecycleStateEnum = "DELETED"
	ListIntegrationInstancesLifecycleStateFailed   ListIntegrationInstancesLifecycleStateEnum = "FAILED"
)

var mappingListIntegrationInstancesLifecycleState = map[string]ListIntegrationInstancesLifecycleStateEnum{
	"CREATING": ListIntegrationInstancesLifecycleStateCreating,
	"UPDATING": ListIntegrationInstancesLifecycleStateUpdating,
	"ACTIVE":   ListIntegrationInstancesLifecycleStateActive,
	"INACTIVE": ListIntegrationInstancesLifecycleStateInactive,
	"DELETING": ListIntegrationInstancesLifecycleStateDeleting,
	"DELETED":  ListIntegrationInstancesLifecycleStateDeleted,
	"FAILED":   ListIntegrationInstancesLifecycleStateFailed,
}

// GetListIntegrationInstancesLifecycleStateEnumValues Enumerates the set of values for ListIntegrationInstancesLifecycleStateEnum
func GetListIntegrationInstancesLifecycleStateEnumValues() []ListIntegrationInstancesLifecycleStateEnum {
	values := make([]ListIntegrationInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListIntegrationInstancesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListIntegrationInstancesSortOrderEnum Enum with underlying type: string
type ListIntegrationInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListIntegrationInstancesSortOrderEnum
const (
	ListIntegrationInstancesSortOrderAsc  ListIntegrationInstancesSortOrderEnum = "ASC"
	ListIntegrationInstancesSortOrderDesc ListIntegrationInstancesSortOrderEnum = "DESC"
)

var mappingListIntegrationInstancesSortOrder = map[string]ListIntegrationInstancesSortOrderEnum{
	"ASC":  ListIntegrationInstancesSortOrderAsc,
	"DESC": ListIntegrationInstancesSortOrderDesc,
}

// GetListIntegrationInstancesSortOrderEnumValues Enumerates the set of values for ListIntegrationInstancesSortOrderEnum
func GetListIntegrationInstancesSortOrderEnumValues() []ListIntegrationInstancesSortOrderEnum {
	values := make([]ListIntegrationInstancesSortOrderEnum, 0)
	for _, v := range mappingListIntegrationInstancesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListIntegrationInstancesSortByEnum Enum with underlying type: string
type ListIntegrationInstancesSortByEnum string

// Set of constants representing the allowable values for ListIntegrationInstancesSortByEnum
const (
	ListIntegrationInstancesSortByTimecreated ListIntegrationInstancesSortByEnum = "TIMECREATED"
	ListIntegrationInstancesSortByDisplayname ListIntegrationInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListIntegrationInstancesSortBy = map[string]ListIntegrationInstancesSortByEnum{
	"TIMECREATED": ListIntegrationInstancesSortByTimecreated,
	"DISPLAYNAME": ListIntegrationInstancesSortByDisplayname,
}

// GetListIntegrationInstancesSortByEnumValues Enumerates the set of values for ListIntegrationInstancesSortByEnum
func GetListIntegrationInstancesSortByEnumValues() []ListIntegrationInstancesSortByEnum {
	values := make([]ListIntegrationInstancesSortByEnum, 0)
	for _, v := range mappingListIntegrationInstancesSortBy {
		values = append(values, v)
	}
	return values
}
