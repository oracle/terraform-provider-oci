// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListIndexesRequest wrapper for the ListIndexes operation
type ListIndexesRequest struct {

	// A table name within the compartment, or a table OCID.
	TableNameOrId *string `mandatory:"true" contributesTo:"path" name:"tableNameOrId"`

	// The ID of a table's compartment. When a table is identified
	// by name, the compartmentId is often needed to provide
	// context for interpreting the name.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A shell-globbing-style (*?[]) filter for names.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filter list by the lifecycle state of the item.
	LifecycleState ListIndexesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start
	// retrieving results. This is usually retrieved from a previous
	// list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListIndexesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be
	// provided. Default order for timeCreated is descending. Default
	// order for name is ascending. If no value is specified
	// timeCreated is default.
	SortBy ListIndexesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListIndexesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListIndexesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListIndexesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListIndexesResponse wrapper for the ListIndexes operation
type ListIndexesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IndexCollection instances
	IndexCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list,
	// if this header appears in the response, then a partial list
	// might have been returned. Include this value as the `page`
	// parameter for the subsequent GET request to get the next batch
	// of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListIndexesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListIndexesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListIndexesLifecycleStateEnum Enum with underlying type: string
type ListIndexesLifecycleStateEnum string

// Set of constants representing the allowable values for ListIndexesLifecycleStateEnum
const (
	ListIndexesLifecycleStateAll      ListIndexesLifecycleStateEnum = "ALL"
	ListIndexesLifecycleStateCreating ListIndexesLifecycleStateEnum = "CREATING"
	ListIndexesLifecycleStateUpdating ListIndexesLifecycleStateEnum = "UPDATING"
	ListIndexesLifecycleStateActive   ListIndexesLifecycleStateEnum = "ACTIVE"
	ListIndexesLifecycleStateDeleting ListIndexesLifecycleStateEnum = "DELETING"
	ListIndexesLifecycleStateDeleted  ListIndexesLifecycleStateEnum = "DELETED"
	ListIndexesLifecycleStateFailed   ListIndexesLifecycleStateEnum = "FAILED"
)

var mappingListIndexesLifecycleState = map[string]ListIndexesLifecycleStateEnum{
	"ALL":      ListIndexesLifecycleStateAll,
	"CREATING": ListIndexesLifecycleStateCreating,
	"UPDATING": ListIndexesLifecycleStateUpdating,
	"ACTIVE":   ListIndexesLifecycleStateActive,
	"DELETING": ListIndexesLifecycleStateDeleting,
	"DELETED":  ListIndexesLifecycleStateDeleted,
	"FAILED":   ListIndexesLifecycleStateFailed,
}

// GetListIndexesLifecycleStateEnumValues Enumerates the set of values for ListIndexesLifecycleStateEnum
func GetListIndexesLifecycleStateEnumValues() []ListIndexesLifecycleStateEnum {
	values := make([]ListIndexesLifecycleStateEnum, 0)
	for _, v := range mappingListIndexesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListIndexesSortOrderEnum Enum with underlying type: string
type ListIndexesSortOrderEnum string

// Set of constants representing the allowable values for ListIndexesSortOrderEnum
const (
	ListIndexesSortOrderAsc  ListIndexesSortOrderEnum = "ASC"
	ListIndexesSortOrderDesc ListIndexesSortOrderEnum = "DESC"
)

var mappingListIndexesSortOrder = map[string]ListIndexesSortOrderEnum{
	"ASC":  ListIndexesSortOrderAsc,
	"DESC": ListIndexesSortOrderDesc,
}

// GetListIndexesSortOrderEnumValues Enumerates the set of values for ListIndexesSortOrderEnum
func GetListIndexesSortOrderEnumValues() []ListIndexesSortOrderEnum {
	values := make([]ListIndexesSortOrderEnum, 0)
	for _, v := range mappingListIndexesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListIndexesSortByEnum Enum with underlying type: string
type ListIndexesSortByEnum string

// Set of constants representing the allowable values for ListIndexesSortByEnum
const (
	ListIndexesSortByTimecreated ListIndexesSortByEnum = "timeCreated"
	ListIndexesSortByName        ListIndexesSortByEnum = "name"
)

var mappingListIndexesSortBy = map[string]ListIndexesSortByEnum{
	"timeCreated": ListIndexesSortByTimecreated,
	"name":        ListIndexesSortByName,
}

// GetListIndexesSortByEnumValues Enumerates the set of values for ListIndexesSortByEnum
func GetListIndexesSortByEnumValues() []ListIndexesSortByEnum {
	values := make([]ListIndexesSortByEnum, 0)
	for _, v := range mappingListIndexesSortBy {
		values = append(values, v)
	}
	return values
}
