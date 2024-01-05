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

// ListAlertRulesRequest wrapper for the ListAlertRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListAlertRules.go.html to see an example of how to use ListAlertRulesRequest.
type ListAlertRulesRequest struct {

	// The unique budget OCID.
	BudgetId *string `mandatory:"true" contributesTo:"path" name:"budgetId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAlertRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If not specified, the default is timeCreated.
	// The default sort order for timeCreated is DESC.
	// The default sort order for displayName is ASC in alphanumeric order.
	SortBy ListAlertRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The current state of the resource to filter by.
	LifecycleState ListAlertRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A user-friendly name. This does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAlertRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAlertRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAlertRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAlertRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAlertRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAlertRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAlertRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAlertRulesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAlertRulesLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAlertRulesResponse wrapper for the ListAlertRules operation
type ListAlertRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AlertRuleSummary instances
	Items []AlertRuleSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `AlertRuleSummary`. If this header appears in the response, then this
	// is a partial list of AlertRuleSummaries. Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of AlertRuleSummaries.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAlertRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAlertRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAlertRulesSortOrderEnum Enum with underlying type: string
type ListAlertRulesSortOrderEnum string

// Set of constants representing the allowable values for ListAlertRulesSortOrderEnum
const (
	ListAlertRulesSortOrderAsc  ListAlertRulesSortOrderEnum = "ASC"
	ListAlertRulesSortOrderDesc ListAlertRulesSortOrderEnum = "DESC"
)

var mappingListAlertRulesSortOrderEnum = map[string]ListAlertRulesSortOrderEnum{
	"ASC":  ListAlertRulesSortOrderAsc,
	"DESC": ListAlertRulesSortOrderDesc,
}

var mappingListAlertRulesSortOrderEnumLowerCase = map[string]ListAlertRulesSortOrderEnum{
	"asc":  ListAlertRulesSortOrderAsc,
	"desc": ListAlertRulesSortOrderDesc,
}

// GetListAlertRulesSortOrderEnumValues Enumerates the set of values for ListAlertRulesSortOrderEnum
func GetListAlertRulesSortOrderEnumValues() []ListAlertRulesSortOrderEnum {
	values := make([]ListAlertRulesSortOrderEnum, 0)
	for _, v := range mappingListAlertRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertRulesSortOrderEnumStringValues Enumerates the set of values in String for ListAlertRulesSortOrderEnum
func GetListAlertRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAlertRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertRulesSortOrderEnum(val string) (ListAlertRulesSortOrderEnum, bool) {
	enum, ok := mappingListAlertRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertRulesSortByEnum Enum with underlying type: string
type ListAlertRulesSortByEnum string

// Set of constants representing the allowable values for ListAlertRulesSortByEnum
const (
	ListAlertRulesSortByTimecreated ListAlertRulesSortByEnum = "timeCreated"
	ListAlertRulesSortByDisplayname ListAlertRulesSortByEnum = "displayName"
)

var mappingListAlertRulesSortByEnum = map[string]ListAlertRulesSortByEnum{
	"timeCreated": ListAlertRulesSortByTimecreated,
	"displayName": ListAlertRulesSortByDisplayname,
}

var mappingListAlertRulesSortByEnumLowerCase = map[string]ListAlertRulesSortByEnum{
	"timecreated": ListAlertRulesSortByTimecreated,
	"displayname": ListAlertRulesSortByDisplayname,
}

// GetListAlertRulesSortByEnumValues Enumerates the set of values for ListAlertRulesSortByEnum
func GetListAlertRulesSortByEnumValues() []ListAlertRulesSortByEnum {
	values := make([]ListAlertRulesSortByEnum, 0)
	for _, v := range mappingListAlertRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertRulesSortByEnumStringValues Enumerates the set of values in String for ListAlertRulesSortByEnum
func GetListAlertRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAlertRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertRulesSortByEnum(val string) (ListAlertRulesSortByEnum, bool) {
	enum, ok := mappingListAlertRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertRulesLifecycleStateEnum Enum with underlying type: string
type ListAlertRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListAlertRulesLifecycleStateEnum
const (
	ListAlertRulesLifecycleStateActive   ListAlertRulesLifecycleStateEnum = "ACTIVE"
	ListAlertRulesLifecycleStateInactive ListAlertRulesLifecycleStateEnum = "INACTIVE"
)

var mappingListAlertRulesLifecycleStateEnum = map[string]ListAlertRulesLifecycleStateEnum{
	"ACTIVE":   ListAlertRulesLifecycleStateActive,
	"INACTIVE": ListAlertRulesLifecycleStateInactive,
}

var mappingListAlertRulesLifecycleStateEnumLowerCase = map[string]ListAlertRulesLifecycleStateEnum{
	"active":   ListAlertRulesLifecycleStateActive,
	"inactive": ListAlertRulesLifecycleStateInactive,
}

// GetListAlertRulesLifecycleStateEnumValues Enumerates the set of values for ListAlertRulesLifecycleStateEnum
func GetListAlertRulesLifecycleStateEnumValues() []ListAlertRulesLifecycleStateEnum {
	values := make([]ListAlertRulesLifecycleStateEnum, 0)
	for _, v := range mappingListAlertRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListAlertRulesLifecycleStateEnum
func GetListAlertRulesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingListAlertRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertRulesLifecycleStateEnum(val string) (ListAlertRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListAlertRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
