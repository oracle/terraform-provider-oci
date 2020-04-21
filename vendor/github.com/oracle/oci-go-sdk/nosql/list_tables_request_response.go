// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTablesRequest wrapper for the ListTables operation
type ListTablesRequest struct {

	// The ID of a table's compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A shell-globbing-style (*?[]) filter for names.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start
	// retrieving results. This is usually retrieved from a previous
	// list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListTablesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be
	// provided. Default order for timeCreated is descending. Default
	// order for name is ascending. If no value is specified
	// timeCreated is default.
	SortBy ListTablesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter list by the lifecycle state of the item.
	LifecycleState ListTablesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTablesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTablesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTablesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTablesResponse wrapper for the ListTables operation
type ListTablesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TableCollection instances
	TableCollection `presentIn:"body"`

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

func (response ListTablesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTablesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTablesSortOrderEnum Enum with underlying type: string
type ListTablesSortOrderEnum string

// Set of constants representing the allowable values for ListTablesSortOrderEnum
const (
	ListTablesSortOrderAsc  ListTablesSortOrderEnum = "ASC"
	ListTablesSortOrderDesc ListTablesSortOrderEnum = "DESC"
)

var mappingListTablesSortOrder = map[string]ListTablesSortOrderEnum{
	"ASC":  ListTablesSortOrderAsc,
	"DESC": ListTablesSortOrderDesc,
}

// GetListTablesSortOrderEnumValues Enumerates the set of values for ListTablesSortOrderEnum
func GetListTablesSortOrderEnumValues() []ListTablesSortOrderEnum {
	values := make([]ListTablesSortOrderEnum, 0)
	for _, v := range mappingListTablesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListTablesSortByEnum Enum with underlying type: string
type ListTablesSortByEnum string

// Set of constants representing the allowable values for ListTablesSortByEnum
const (
	ListTablesSortByTimecreated ListTablesSortByEnum = "timeCreated"
	ListTablesSortByName        ListTablesSortByEnum = "name"
)

var mappingListTablesSortBy = map[string]ListTablesSortByEnum{
	"timeCreated": ListTablesSortByTimecreated,
	"name":        ListTablesSortByName,
}

// GetListTablesSortByEnumValues Enumerates the set of values for ListTablesSortByEnum
func GetListTablesSortByEnumValues() []ListTablesSortByEnum {
	values := make([]ListTablesSortByEnum, 0)
	for _, v := range mappingListTablesSortBy {
		values = append(values, v)
	}
	return values
}

// ListTablesLifecycleStateEnum Enum with underlying type: string
type ListTablesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTablesLifecycleStateEnum
const (
	ListTablesLifecycleStateAll      ListTablesLifecycleStateEnum = "ALL"
	ListTablesLifecycleStateCreating ListTablesLifecycleStateEnum = "CREATING"
	ListTablesLifecycleStateUpdating ListTablesLifecycleStateEnum = "UPDATING"
	ListTablesLifecycleStateActive   ListTablesLifecycleStateEnum = "ACTIVE"
	ListTablesLifecycleStateDeleting ListTablesLifecycleStateEnum = "DELETING"
	ListTablesLifecycleStateDeleted  ListTablesLifecycleStateEnum = "DELETED"
	ListTablesLifecycleStateFailed   ListTablesLifecycleStateEnum = "FAILED"
)

var mappingListTablesLifecycleState = map[string]ListTablesLifecycleStateEnum{
	"ALL":      ListTablesLifecycleStateAll,
	"CREATING": ListTablesLifecycleStateCreating,
	"UPDATING": ListTablesLifecycleStateUpdating,
	"ACTIVE":   ListTablesLifecycleStateActive,
	"DELETING": ListTablesLifecycleStateDeleting,
	"DELETED":  ListTablesLifecycleStateDeleted,
	"FAILED":   ListTablesLifecycleStateFailed,
}

// GetListTablesLifecycleStateEnumValues Enumerates the set of values for ListTablesLifecycleStateEnum
func GetListTablesLifecycleStateEnumValues() []ListTablesLifecycleStateEnum {
	values := make([]ListTablesLifecycleStateEnum, 0)
	for _, v := range mappingListTablesLifecycleState {
		values = append(values, v)
	}
	return values
}
