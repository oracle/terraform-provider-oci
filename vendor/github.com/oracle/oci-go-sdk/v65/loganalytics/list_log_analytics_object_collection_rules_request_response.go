// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLogAnalyticsObjectCollectionRulesRequest wrapper for the ListLogAnalyticsObjectCollectionRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsObjectCollectionRules.go.html to see an example of how to use ListLogAnalyticsObjectCollectionRulesRequest.
type ListLogAnalyticsObjectCollectionRulesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return rules only matching with this name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Lifecycle state filter.
	LifecycleState ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogAnalyticsObjectCollectionRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeUpdated is descending.
	// Default order for name is ascending. If no value is specified timeUpdated is default.
	SortBy ListLogAnalyticsObjectCollectionRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogAnalyticsObjectCollectionRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogAnalyticsObjectCollectionRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogAnalyticsObjectCollectionRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogAnalyticsObjectCollectionRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogAnalyticsObjectCollectionRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLogAnalyticsObjectCollectionRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListLogAnalyticsObjectCollectionRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogAnalyticsObjectCollectionRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogAnalyticsObjectCollectionRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogAnalyticsObjectCollectionRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogAnalyticsObjectCollectionRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogAnalyticsObjectCollectionRulesResponse wrapper for the ListLogAnalyticsObjectCollectionRules operation
type ListLogAnalyticsObjectCollectionRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsObjectCollectionRuleCollection instances
	LogAnalyticsObjectCollectionRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogAnalyticsObjectCollectionRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogAnalyticsObjectCollectionRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum Enum with underlying type: string
type ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum
const (
	ListLogAnalyticsObjectCollectionRulesLifecycleStateActive  ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum = "ACTIVE"
	ListLogAnalyticsObjectCollectionRulesLifecycleStateDeleted ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum = "DELETED"
)

var mappingListLogAnalyticsObjectCollectionRulesLifecycleStateEnum = map[string]ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum{
	"ACTIVE":  ListLogAnalyticsObjectCollectionRulesLifecycleStateActive,
	"DELETED": ListLogAnalyticsObjectCollectionRulesLifecycleStateDeleted,
}

var mappingListLogAnalyticsObjectCollectionRulesLifecycleStateEnumLowerCase = map[string]ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum{
	"active":  ListLogAnalyticsObjectCollectionRulesLifecycleStateActive,
	"deleted": ListLogAnalyticsObjectCollectionRulesLifecycleStateDeleted,
}

// GetListLogAnalyticsObjectCollectionRulesLifecycleStateEnumValues Enumerates the set of values for ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum
func GetListLogAnalyticsObjectCollectionRulesLifecycleStateEnumValues() []ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum {
	values := make([]ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum, 0)
	for _, v := range mappingListLogAnalyticsObjectCollectionRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogAnalyticsObjectCollectionRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum
func GetListLogAnalyticsObjectCollectionRulesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingListLogAnalyticsObjectCollectionRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogAnalyticsObjectCollectionRulesLifecycleStateEnum(val string) (ListLogAnalyticsObjectCollectionRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListLogAnalyticsObjectCollectionRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogAnalyticsObjectCollectionRulesSortOrderEnum Enum with underlying type: string
type ListLogAnalyticsObjectCollectionRulesSortOrderEnum string

// Set of constants representing the allowable values for ListLogAnalyticsObjectCollectionRulesSortOrderEnum
const (
	ListLogAnalyticsObjectCollectionRulesSortOrderAsc  ListLogAnalyticsObjectCollectionRulesSortOrderEnum = "ASC"
	ListLogAnalyticsObjectCollectionRulesSortOrderDesc ListLogAnalyticsObjectCollectionRulesSortOrderEnum = "DESC"
)

var mappingListLogAnalyticsObjectCollectionRulesSortOrderEnum = map[string]ListLogAnalyticsObjectCollectionRulesSortOrderEnum{
	"ASC":  ListLogAnalyticsObjectCollectionRulesSortOrderAsc,
	"DESC": ListLogAnalyticsObjectCollectionRulesSortOrderDesc,
}

var mappingListLogAnalyticsObjectCollectionRulesSortOrderEnumLowerCase = map[string]ListLogAnalyticsObjectCollectionRulesSortOrderEnum{
	"asc":  ListLogAnalyticsObjectCollectionRulesSortOrderAsc,
	"desc": ListLogAnalyticsObjectCollectionRulesSortOrderDesc,
}

// GetListLogAnalyticsObjectCollectionRulesSortOrderEnumValues Enumerates the set of values for ListLogAnalyticsObjectCollectionRulesSortOrderEnum
func GetListLogAnalyticsObjectCollectionRulesSortOrderEnumValues() []ListLogAnalyticsObjectCollectionRulesSortOrderEnum {
	values := make([]ListLogAnalyticsObjectCollectionRulesSortOrderEnum, 0)
	for _, v := range mappingListLogAnalyticsObjectCollectionRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogAnalyticsObjectCollectionRulesSortOrderEnumStringValues Enumerates the set of values in String for ListLogAnalyticsObjectCollectionRulesSortOrderEnum
func GetListLogAnalyticsObjectCollectionRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogAnalyticsObjectCollectionRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogAnalyticsObjectCollectionRulesSortOrderEnum(val string) (ListLogAnalyticsObjectCollectionRulesSortOrderEnum, bool) {
	enum, ok := mappingListLogAnalyticsObjectCollectionRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogAnalyticsObjectCollectionRulesSortByEnum Enum with underlying type: string
type ListLogAnalyticsObjectCollectionRulesSortByEnum string

// Set of constants representing the allowable values for ListLogAnalyticsObjectCollectionRulesSortByEnum
const (
	ListLogAnalyticsObjectCollectionRulesSortByTimeupdated ListLogAnalyticsObjectCollectionRulesSortByEnum = "timeUpdated"
	ListLogAnalyticsObjectCollectionRulesSortByTimecreated ListLogAnalyticsObjectCollectionRulesSortByEnum = "timeCreated"
	ListLogAnalyticsObjectCollectionRulesSortByName        ListLogAnalyticsObjectCollectionRulesSortByEnum = "name"
)

var mappingListLogAnalyticsObjectCollectionRulesSortByEnum = map[string]ListLogAnalyticsObjectCollectionRulesSortByEnum{
	"timeUpdated": ListLogAnalyticsObjectCollectionRulesSortByTimeupdated,
	"timeCreated": ListLogAnalyticsObjectCollectionRulesSortByTimecreated,
	"name":        ListLogAnalyticsObjectCollectionRulesSortByName,
}

var mappingListLogAnalyticsObjectCollectionRulesSortByEnumLowerCase = map[string]ListLogAnalyticsObjectCollectionRulesSortByEnum{
	"timeupdated": ListLogAnalyticsObjectCollectionRulesSortByTimeupdated,
	"timecreated": ListLogAnalyticsObjectCollectionRulesSortByTimecreated,
	"name":        ListLogAnalyticsObjectCollectionRulesSortByName,
}

// GetListLogAnalyticsObjectCollectionRulesSortByEnumValues Enumerates the set of values for ListLogAnalyticsObjectCollectionRulesSortByEnum
func GetListLogAnalyticsObjectCollectionRulesSortByEnumValues() []ListLogAnalyticsObjectCollectionRulesSortByEnum {
	values := make([]ListLogAnalyticsObjectCollectionRulesSortByEnum, 0)
	for _, v := range mappingListLogAnalyticsObjectCollectionRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogAnalyticsObjectCollectionRulesSortByEnumStringValues Enumerates the set of values in String for ListLogAnalyticsObjectCollectionRulesSortByEnum
func GetListLogAnalyticsObjectCollectionRulesSortByEnumStringValues() []string {
	return []string{
		"timeUpdated",
		"timeCreated",
		"name",
	}
}

// GetMappingListLogAnalyticsObjectCollectionRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogAnalyticsObjectCollectionRulesSortByEnum(val string) (ListLogAnalyticsObjectCollectionRulesSortByEnum, bool) {
	enum, ok := mappingListLogAnalyticsObjectCollectionRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
