// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package budget

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListBudgetsRequest wrapper for the ListBudgets operation
type ListBudgetsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBudgetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If not specified, the default is timeCreated.
	// The default sort order for timeCreated is DESC.
	// The default sort order for displayName is ASC in alphanumeric order.
	SortBy ListBudgetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The current state of the resource to filter by.
	LifecycleState ListBudgetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The type of target to filter by.
	//   * ALL - List all budgets
	//   * COMPARTMENT - List all budgets with targetType == "COMPARTMENT"
	//   * TAG - List all budgets with targetType == "TAG"
	TargetType ListBudgetsTargetTypeEnum `mandatory:"false" contributesTo:"query" name:"targetType" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBudgetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBudgetsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBudgetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBudgetsResponse wrapper for the ListBudgets operation
type ListBudgetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BudgetSummary instances
	Items []BudgetSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `Budget`s. If this header appears in the response, then this
	// is a partial list of Budgets. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of Budgets.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBudgetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBudgetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBudgetsSortOrderEnum Enum with underlying type: string
type ListBudgetsSortOrderEnum string

// Set of constants representing the allowable values for ListBudgetsSortOrderEnum
const (
	ListBudgetsSortOrderAsc  ListBudgetsSortOrderEnum = "ASC"
	ListBudgetsSortOrderDesc ListBudgetsSortOrderEnum = "DESC"
)

var mappingListBudgetsSortOrder = map[string]ListBudgetsSortOrderEnum{
	"ASC":  ListBudgetsSortOrderAsc,
	"DESC": ListBudgetsSortOrderDesc,
}

// GetListBudgetsSortOrderEnumValues Enumerates the set of values for ListBudgetsSortOrderEnum
func GetListBudgetsSortOrderEnumValues() []ListBudgetsSortOrderEnum {
	values := make([]ListBudgetsSortOrderEnum, 0)
	for _, v := range mappingListBudgetsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListBudgetsSortByEnum Enum with underlying type: string
type ListBudgetsSortByEnum string

// Set of constants representing the allowable values for ListBudgetsSortByEnum
const (
	ListBudgetsSortByTimecreated ListBudgetsSortByEnum = "timeCreated"
	ListBudgetsSortByDisplayname ListBudgetsSortByEnum = "displayName"
)

var mappingListBudgetsSortBy = map[string]ListBudgetsSortByEnum{
	"timeCreated": ListBudgetsSortByTimecreated,
	"displayName": ListBudgetsSortByDisplayname,
}

// GetListBudgetsSortByEnumValues Enumerates the set of values for ListBudgetsSortByEnum
func GetListBudgetsSortByEnumValues() []ListBudgetsSortByEnum {
	values := make([]ListBudgetsSortByEnum, 0)
	for _, v := range mappingListBudgetsSortBy {
		values = append(values, v)
	}
	return values
}

// ListBudgetsLifecycleStateEnum Enum with underlying type: string
type ListBudgetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListBudgetsLifecycleStateEnum
const (
	ListBudgetsLifecycleStateActive   ListBudgetsLifecycleStateEnum = "ACTIVE"
	ListBudgetsLifecycleStateInactive ListBudgetsLifecycleStateEnum = "INACTIVE"
)

var mappingListBudgetsLifecycleState = map[string]ListBudgetsLifecycleStateEnum{
	"ACTIVE":   ListBudgetsLifecycleStateActive,
	"INACTIVE": ListBudgetsLifecycleStateInactive,
}

// GetListBudgetsLifecycleStateEnumValues Enumerates the set of values for ListBudgetsLifecycleStateEnum
func GetListBudgetsLifecycleStateEnumValues() []ListBudgetsLifecycleStateEnum {
	values := make([]ListBudgetsLifecycleStateEnum, 0)
	for _, v := range mappingListBudgetsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListBudgetsTargetTypeEnum Enum with underlying type: string
type ListBudgetsTargetTypeEnum string

// Set of constants representing the allowable values for ListBudgetsTargetTypeEnum
const (
	ListBudgetsTargetTypeAll         ListBudgetsTargetTypeEnum = "ALL"
	ListBudgetsTargetTypeCompartment ListBudgetsTargetTypeEnum = "COMPARTMENT"
	ListBudgetsTargetTypeTag         ListBudgetsTargetTypeEnum = "TAG"
)

var mappingListBudgetsTargetType = map[string]ListBudgetsTargetTypeEnum{
	"ALL":         ListBudgetsTargetTypeAll,
	"COMPARTMENT": ListBudgetsTargetTypeCompartment,
	"TAG":         ListBudgetsTargetTypeTag,
}

// GetListBudgetsTargetTypeEnumValues Enumerates the set of values for ListBudgetsTargetTypeEnum
func GetListBudgetsTargetTypeEnumValues() []ListBudgetsTargetTypeEnum {
	values := make([]ListBudgetsTargetTypeEnum, 0)
	for _, v := range mappingListBudgetsTargetType {
		values = append(values, v)
	}
	return values
}
