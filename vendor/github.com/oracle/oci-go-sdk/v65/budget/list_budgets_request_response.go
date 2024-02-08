// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package budget

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBudgetsRequest wrapper for the ListBudgets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListBudgets.go.html to see an example of how to use ListBudgetsRequest.
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

	// A user-friendly name. This does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The type of target to filter by:
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
func (request ListBudgetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBudgetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBudgetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBudgetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBudgetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBudgetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBudgetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBudgetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBudgetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListBudgetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBudgetsTargetTypeEnum(string(request.TargetType)); !ok && request.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", request.TargetType, strings.Join(GetListBudgetsTargetTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	// is a partial list of budgets. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of budgets.
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

var mappingListBudgetsSortOrderEnum = map[string]ListBudgetsSortOrderEnum{
	"ASC":  ListBudgetsSortOrderAsc,
	"DESC": ListBudgetsSortOrderDesc,
}

var mappingListBudgetsSortOrderEnumLowerCase = map[string]ListBudgetsSortOrderEnum{
	"asc":  ListBudgetsSortOrderAsc,
	"desc": ListBudgetsSortOrderDesc,
}

// GetListBudgetsSortOrderEnumValues Enumerates the set of values for ListBudgetsSortOrderEnum
func GetListBudgetsSortOrderEnumValues() []ListBudgetsSortOrderEnum {
	values := make([]ListBudgetsSortOrderEnum, 0)
	for _, v := range mappingListBudgetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBudgetsSortOrderEnumStringValues Enumerates the set of values in String for ListBudgetsSortOrderEnum
func GetListBudgetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBudgetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBudgetsSortOrderEnum(val string) (ListBudgetsSortOrderEnum, bool) {
	enum, ok := mappingListBudgetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBudgetsSortByEnum Enum with underlying type: string
type ListBudgetsSortByEnum string

// Set of constants representing the allowable values for ListBudgetsSortByEnum
const (
	ListBudgetsSortByTimecreated ListBudgetsSortByEnum = "timeCreated"
	ListBudgetsSortByDisplayname ListBudgetsSortByEnum = "displayName"
)

var mappingListBudgetsSortByEnum = map[string]ListBudgetsSortByEnum{
	"timeCreated": ListBudgetsSortByTimecreated,
	"displayName": ListBudgetsSortByDisplayname,
}

var mappingListBudgetsSortByEnumLowerCase = map[string]ListBudgetsSortByEnum{
	"timecreated": ListBudgetsSortByTimecreated,
	"displayname": ListBudgetsSortByDisplayname,
}

// GetListBudgetsSortByEnumValues Enumerates the set of values for ListBudgetsSortByEnum
func GetListBudgetsSortByEnumValues() []ListBudgetsSortByEnum {
	values := make([]ListBudgetsSortByEnum, 0)
	for _, v := range mappingListBudgetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBudgetsSortByEnumStringValues Enumerates the set of values in String for ListBudgetsSortByEnum
func GetListBudgetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListBudgetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBudgetsSortByEnum(val string) (ListBudgetsSortByEnum, bool) {
	enum, ok := mappingListBudgetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBudgetsLifecycleStateEnum Enum with underlying type: string
type ListBudgetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListBudgetsLifecycleStateEnum
const (
	ListBudgetsLifecycleStateActive   ListBudgetsLifecycleStateEnum = "ACTIVE"
	ListBudgetsLifecycleStateInactive ListBudgetsLifecycleStateEnum = "INACTIVE"
)

var mappingListBudgetsLifecycleStateEnum = map[string]ListBudgetsLifecycleStateEnum{
	"ACTIVE":   ListBudgetsLifecycleStateActive,
	"INACTIVE": ListBudgetsLifecycleStateInactive,
}

var mappingListBudgetsLifecycleStateEnumLowerCase = map[string]ListBudgetsLifecycleStateEnum{
	"active":   ListBudgetsLifecycleStateActive,
	"inactive": ListBudgetsLifecycleStateInactive,
}

// GetListBudgetsLifecycleStateEnumValues Enumerates the set of values for ListBudgetsLifecycleStateEnum
func GetListBudgetsLifecycleStateEnumValues() []ListBudgetsLifecycleStateEnum {
	values := make([]ListBudgetsLifecycleStateEnum, 0)
	for _, v := range mappingListBudgetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListBudgetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListBudgetsLifecycleStateEnum
func GetListBudgetsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListBudgetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBudgetsLifecycleStateEnum(val string) (ListBudgetsLifecycleStateEnum, bool) {
	enum, ok := mappingListBudgetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBudgetsTargetTypeEnum Enum with underlying type: string
type ListBudgetsTargetTypeEnum string

// Set of constants representing the allowable values for ListBudgetsTargetTypeEnum
const (
	ListBudgetsTargetTypeAll         ListBudgetsTargetTypeEnum = "ALL"
	ListBudgetsTargetTypeCompartment ListBudgetsTargetTypeEnum = "COMPARTMENT"
	ListBudgetsTargetTypeTag         ListBudgetsTargetTypeEnum = "TAG"
)

var mappingListBudgetsTargetTypeEnum = map[string]ListBudgetsTargetTypeEnum{
	"ALL":         ListBudgetsTargetTypeAll,
	"COMPARTMENT": ListBudgetsTargetTypeCompartment,
	"TAG":         ListBudgetsTargetTypeTag,
}

var mappingListBudgetsTargetTypeEnumLowerCase = map[string]ListBudgetsTargetTypeEnum{
	"all":         ListBudgetsTargetTypeAll,
	"compartment": ListBudgetsTargetTypeCompartment,
	"tag":         ListBudgetsTargetTypeTag,
}

// GetListBudgetsTargetTypeEnumValues Enumerates the set of values for ListBudgetsTargetTypeEnum
func GetListBudgetsTargetTypeEnumValues() []ListBudgetsTargetTypeEnum {
	values := make([]ListBudgetsTargetTypeEnum, 0)
	for _, v := range mappingListBudgetsTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListBudgetsTargetTypeEnumStringValues Enumerates the set of values in String for ListBudgetsTargetTypeEnum
func GetListBudgetsTargetTypeEnumStringValues() []string {
	return []string{
		"ALL",
		"COMPARTMENT",
		"TAG",
	}
}

// GetMappingListBudgetsTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBudgetsTargetTypeEnum(val string) (ListBudgetsTargetTypeEnum, bool) {
	enum, ok := mappingListBudgetsTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
