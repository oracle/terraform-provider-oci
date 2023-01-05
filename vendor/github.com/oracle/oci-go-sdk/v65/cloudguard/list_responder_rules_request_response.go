// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListResponderRulesRequest wrapper for the ListResponderRules operation
type ListResponderRulesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListResponderRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResponderRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListResponderRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResponderRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResponderRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResponderRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResponderRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResponderRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResponderRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListResponderRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResponderRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResponderRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResponderRulesResponse wrapper for the ListResponderRules operation
type ListResponderRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResponderRuleCollection instances
	ResponderRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResponderRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResponderRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResponderRulesLifecycleStateEnum Enum with underlying type: string
type ListResponderRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListResponderRulesLifecycleStateEnum
const (
	ListResponderRulesLifecycleStateCreating ListResponderRulesLifecycleStateEnum = "CREATING"
	ListResponderRulesLifecycleStateUpdating ListResponderRulesLifecycleStateEnum = "UPDATING"
	ListResponderRulesLifecycleStateActive   ListResponderRulesLifecycleStateEnum = "ACTIVE"
	ListResponderRulesLifecycleStateInactive ListResponderRulesLifecycleStateEnum = "INACTIVE"
	ListResponderRulesLifecycleStateDeleting ListResponderRulesLifecycleStateEnum = "DELETING"
	ListResponderRulesLifecycleStateDeleted  ListResponderRulesLifecycleStateEnum = "DELETED"
	ListResponderRulesLifecycleStateFailed   ListResponderRulesLifecycleStateEnum = "FAILED"
)

var mappingListResponderRulesLifecycleStateEnum = map[string]ListResponderRulesLifecycleStateEnum{
	"CREATING": ListResponderRulesLifecycleStateCreating,
	"UPDATING": ListResponderRulesLifecycleStateUpdating,
	"ACTIVE":   ListResponderRulesLifecycleStateActive,
	"INACTIVE": ListResponderRulesLifecycleStateInactive,
	"DELETING": ListResponderRulesLifecycleStateDeleting,
	"DELETED":  ListResponderRulesLifecycleStateDeleted,
	"FAILED":   ListResponderRulesLifecycleStateFailed,
}

var mappingListResponderRulesLifecycleStateEnumLowerCase = map[string]ListResponderRulesLifecycleStateEnum{
	"creating": ListResponderRulesLifecycleStateCreating,
	"updating": ListResponderRulesLifecycleStateUpdating,
	"active":   ListResponderRulesLifecycleStateActive,
	"inactive": ListResponderRulesLifecycleStateInactive,
	"deleting": ListResponderRulesLifecycleStateDeleting,
	"deleted":  ListResponderRulesLifecycleStateDeleted,
	"failed":   ListResponderRulesLifecycleStateFailed,
}

// GetListResponderRulesLifecycleStateEnumValues Enumerates the set of values for ListResponderRulesLifecycleStateEnum
func GetListResponderRulesLifecycleStateEnumValues() []ListResponderRulesLifecycleStateEnum {
	values := make([]ListResponderRulesLifecycleStateEnum, 0)
	for _, v := range mappingListResponderRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListResponderRulesLifecycleStateEnum
func GetListResponderRulesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListResponderRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderRulesLifecycleStateEnum(val string) (ListResponderRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListResponderRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderRulesSortOrderEnum Enum with underlying type: string
type ListResponderRulesSortOrderEnum string

// Set of constants representing the allowable values for ListResponderRulesSortOrderEnum
const (
	ListResponderRulesSortOrderAsc  ListResponderRulesSortOrderEnum = "ASC"
	ListResponderRulesSortOrderDesc ListResponderRulesSortOrderEnum = "DESC"
)

var mappingListResponderRulesSortOrderEnum = map[string]ListResponderRulesSortOrderEnum{
	"ASC":  ListResponderRulesSortOrderAsc,
	"DESC": ListResponderRulesSortOrderDesc,
}

var mappingListResponderRulesSortOrderEnumLowerCase = map[string]ListResponderRulesSortOrderEnum{
	"asc":  ListResponderRulesSortOrderAsc,
	"desc": ListResponderRulesSortOrderDesc,
}

// GetListResponderRulesSortOrderEnumValues Enumerates the set of values for ListResponderRulesSortOrderEnum
func GetListResponderRulesSortOrderEnumValues() []ListResponderRulesSortOrderEnum {
	values := make([]ListResponderRulesSortOrderEnum, 0)
	for _, v := range mappingListResponderRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderRulesSortOrderEnumStringValues Enumerates the set of values in String for ListResponderRulesSortOrderEnum
func GetListResponderRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResponderRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderRulesSortOrderEnum(val string) (ListResponderRulesSortOrderEnum, bool) {
	enum, ok := mappingListResponderRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderRulesSortByEnum Enum with underlying type: string
type ListResponderRulesSortByEnum string

// Set of constants representing the allowable values for ListResponderRulesSortByEnum
const (
	ListResponderRulesSortByTimecreated ListResponderRulesSortByEnum = "timeCreated"
	ListResponderRulesSortByDisplayname ListResponderRulesSortByEnum = "displayName"
)

var mappingListResponderRulesSortByEnum = map[string]ListResponderRulesSortByEnum{
	"timeCreated": ListResponderRulesSortByTimecreated,
	"displayName": ListResponderRulesSortByDisplayname,
}

var mappingListResponderRulesSortByEnumLowerCase = map[string]ListResponderRulesSortByEnum{
	"timecreated": ListResponderRulesSortByTimecreated,
	"displayname": ListResponderRulesSortByDisplayname,
}

// GetListResponderRulesSortByEnumValues Enumerates the set of values for ListResponderRulesSortByEnum
func GetListResponderRulesSortByEnumValues() []ListResponderRulesSortByEnum {
	values := make([]ListResponderRulesSortByEnum, 0)
	for _, v := range mappingListResponderRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderRulesSortByEnumStringValues Enumerates the set of values in String for ListResponderRulesSortByEnum
func GetListResponderRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListResponderRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderRulesSortByEnum(val string) (ListResponderRulesSortByEnum, bool) {
	enum, ok := mappingListResponderRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
