// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oce

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListOceInstancesRequest wrapper for the ListOceInstances operation
type ListOceInstancesRequest struct {

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
	SortOrder ListOceInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOceInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter results on lifecycleState.
	LifecycleState ListOceInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOceInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOceInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOceInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOceInstancesResponse wrapper for the ListOceInstances operation
type ListOceInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []OceInstanceSummary instances
	Items []OceInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `OceInstance`s. If this header appears in the response, then this
	// is a partial list of OceInstances. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of OceInstances.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOceInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOceInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOceInstancesSortOrderEnum Enum with underlying type: string
type ListOceInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListOceInstancesSortOrderEnum
const (
	ListOceInstancesSortOrderAsc  ListOceInstancesSortOrderEnum = "ASC"
	ListOceInstancesSortOrderDesc ListOceInstancesSortOrderEnum = "DESC"
)

var mappingListOceInstancesSortOrder = map[string]ListOceInstancesSortOrderEnum{
	"ASC":  ListOceInstancesSortOrderAsc,
	"DESC": ListOceInstancesSortOrderDesc,
}

// GetListOceInstancesSortOrderEnumValues Enumerates the set of values for ListOceInstancesSortOrderEnum
func GetListOceInstancesSortOrderEnumValues() []ListOceInstancesSortOrderEnum {
	values := make([]ListOceInstancesSortOrderEnum, 0)
	for _, v := range mappingListOceInstancesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListOceInstancesSortByEnum Enum with underlying type: string
type ListOceInstancesSortByEnum string

// Set of constants representing the allowable values for ListOceInstancesSortByEnum
const (
	ListOceInstancesSortByTimecreated ListOceInstancesSortByEnum = "timeCreated"
	ListOceInstancesSortByDisplayname ListOceInstancesSortByEnum = "displayName"
)

var mappingListOceInstancesSortBy = map[string]ListOceInstancesSortByEnum{
	"timeCreated": ListOceInstancesSortByTimecreated,
	"displayName": ListOceInstancesSortByDisplayname,
}

// GetListOceInstancesSortByEnumValues Enumerates the set of values for ListOceInstancesSortByEnum
func GetListOceInstancesSortByEnumValues() []ListOceInstancesSortByEnum {
	values := make([]ListOceInstancesSortByEnum, 0)
	for _, v := range mappingListOceInstancesSortBy {
		values = append(values, v)
	}
	return values
}

// ListOceInstancesLifecycleStateEnum Enum with underlying type: string
type ListOceInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListOceInstancesLifecycleStateEnum
const (
	ListOceInstancesLifecycleStateCreating ListOceInstancesLifecycleStateEnum = "CREATING"
	ListOceInstancesLifecycleStateUpdating ListOceInstancesLifecycleStateEnum = "UPDATING"
	ListOceInstancesLifecycleStateActive   ListOceInstancesLifecycleStateEnum = "ACTIVE"
	ListOceInstancesLifecycleStateDeleting ListOceInstancesLifecycleStateEnum = "DELETING"
	ListOceInstancesLifecycleStateDeleted  ListOceInstancesLifecycleStateEnum = "DELETED"
	ListOceInstancesLifecycleStateFailed   ListOceInstancesLifecycleStateEnum = "FAILED"
)

var mappingListOceInstancesLifecycleState = map[string]ListOceInstancesLifecycleStateEnum{
	"CREATING": ListOceInstancesLifecycleStateCreating,
	"UPDATING": ListOceInstancesLifecycleStateUpdating,
	"ACTIVE":   ListOceInstancesLifecycleStateActive,
	"DELETING": ListOceInstancesLifecycleStateDeleting,
	"DELETED":  ListOceInstancesLifecycleStateDeleted,
	"FAILED":   ListOceInstancesLifecycleStateFailed,
}

// GetListOceInstancesLifecycleStateEnumValues Enumerates the set of values for ListOceInstancesLifecycleStateEnum
func GetListOceInstancesLifecycleStateEnumValues() []ListOceInstancesLifecycleStateEnum {
	values := make([]ListOceInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListOceInstancesLifecycleState {
		values = append(values, v)
	}
	return values
}
