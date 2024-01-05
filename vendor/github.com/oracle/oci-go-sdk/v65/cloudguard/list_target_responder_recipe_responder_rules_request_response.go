// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTargetResponderRecipeResponderRulesRequest wrapper for the ListTargetResponderRecipeResponderRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargetResponderRecipeResponderRules.go.html to see an example of how to use ListTargetResponderRecipeResponderRulesRequest.
type ListTargetResponderRecipeResponderRulesRequest struct {

	// OCID of target
	TargetId *string `mandatory:"true" contributesTo:"path" name:"targetId"`

	// OCID of TargetResponderRecipe
	TargetResponderRecipeId *string `mandatory:"true" contributesTo:"path" name:"targetResponderRecipeId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListTargetResponderRecipeResponderRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTargetResponderRecipeResponderRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListTargetResponderRecipeResponderRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetResponderRecipeResponderRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetResponderRecipeResponderRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetResponderRecipeResponderRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetResponderRecipeResponderRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetResponderRecipeResponderRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetResponderRecipeResponderRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTargetResponderRecipeResponderRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetResponderRecipeResponderRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetResponderRecipeResponderRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetResponderRecipeResponderRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetResponderRecipeResponderRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetResponderRecipeResponderRulesResponse wrapper for the ListTargetResponderRecipeResponderRules operation
type ListTargetResponderRecipeResponderRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetResponderRecipeResponderRuleCollection instances
	TargetResponderRecipeResponderRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTargetResponderRecipeResponderRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetResponderRecipeResponderRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetResponderRecipeResponderRulesLifecycleStateEnum Enum with underlying type: string
type ListTargetResponderRecipeResponderRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTargetResponderRecipeResponderRulesLifecycleStateEnum
const (
	ListTargetResponderRecipeResponderRulesLifecycleStateCreating ListTargetResponderRecipeResponderRulesLifecycleStateEnum = "CREATING"
	ListTargetResponderRecipeResponderRulesLifecycleStateUpdating ListTargetResponderRecipeResponderRulesLifecycleStateEnum = "UPDATING"
	ListTargetResponderRecipeResponderRulesLifecycleStateActive   ListTargetResponderRecipeResponderRulesLifecycleStateEnum = "ACTIVE"
	ListTargetResponderRecipeResponderRulesLifecycleStateInactive ListTargetResponderRecipeResponderRulesLifecycleStateEnum = "INACTIVE"
	ListTargetResponderRecipeResponderRulesLifecycleStateDeleting ListTargetResponderRecipeResponderRulesLifecycleStateEnum = "DELETING"
	ListTargetResponderRecipeResponderRulesLifecycleStateDeleted  ListTargetResponderRecipeResponderRulesLifecycleStateEnum = "DELETED"
	ListTargetResponderRecipeResponderRulesLifecycleStateFailed   ListTargetResponderRecipeResponderRulesLifecycleStateEnum = "FAILED"
)

var mappingListTargetResponderRecipeResponderRulesLifecycleStateEnum = map[string]ListTargetResponderRecipeResponderRulesLifecycleStateEnum{
	"CREATING": ListTargetResponderRecipeResponderRulesLifecycleStateCreating,
	"UPDATING": ListTargetResponderRecipeResponderRulesLifecycleStateUpdating,
	"ACTIVE":   ListTargetResponderRecipeResponderRulesLifecycleStateActive,
	"INACTIVE": ListTargetResponderRecipeResponderRulesLifecycleStateInactive,
	"DELETING": ListTargetResponderRecipeResponderRulesLifecycleStateDeleting,
	"DELETED":  ListTargetResponderRecipeResponderRulesLifecycleStateDeleted,
	"FAILED":   ListTargetResponderRecipeResponderRulesLifecycleStateFailed,
}

var mappingListTargetResponderRecipeResponderRulesLifecycleStateEnumLowerCase = map[string]ListTargetResponderRecipeResponderRulesLifecycleStateEnum{
	"creating": ListTargetResponderRecipeResponderRulesLifecycleStateCreating,
	"updating": ListTargetResponderRecipeResponderRulesLifecycleStateUpdating,
	"active":   ListTargetResponderRecipeResponderRulesLifecycleStateActive,
	"inactive": ListTargetResponderRecipeResponderRulesLifecycleStateInactive,
	"deleting": ListTargetResponderRecipeResponderRulesLifecycleStateDeleting,
	"deleted":  ListTargetResponderRecipeResponderRulesLifecycleStateDeleted,
	"failed":   ListTargetResponderRecipeResponderRulesLifecycleStateFailed,
}

// GetListTargetResponderRecipeResponderRulesLifecycleStateEnumValues Enumerates the set of values for ListTargetResponderRecipeResponderRulesLifecycleStateEnum
func GetListTargetResponderRecipeResponderRulesLifecycleStateEnumValues() []ListTargetResponderRecipeResponderRulesLifecycleStateEnum {
	values := make([]ListTargetResponderRecipeResponderRulesLifecycleStateEnum, 0)
	for _, v := range mappingListTargetResponderRecipeResponderRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetResponderRecipeResponderRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTargetResponderRecipeResponderRulesLifecycleStateEnum
func GetListTargetResponderRecipeResponderRulesLifecycleStateEnumStringValues() []string {
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

// GetMappingListTargetResponderRecipeResponderRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetResponderRecipeResponderRulesLifecycleStateEnum(val string) (ListTargetResponderRecipeResponderRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListTargetResponderRecipeResponderRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetResponderRecipeResponderRulesSortOrderEnum Enum with underlying type: string
type ListTargetResponderRecipeResponderRulesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetResponderRecipeResponderRulesSortOrderEnum
const (
	ListTargetResponderRecipeResponderRulesSortOrderAsc  ListTargetResponderRecipeResponderRulesSortOrderEnum = "ASC"
	ListTargetResponderRecipeResponderRulesSortOrderDesc ListTargetResponderRecipeResponderRulesSortOrderEnum = "DESC"
)

var mappingListTargetResponderRecipeResponderRulesSortOrderEnum = map[string]ListTargetResponderRecipeResponderRulesSortOrderEnum{
	"ASC":  ListTargetResponderRecipeResponderRulesSortOrderAsc,
	"DESC": ListTargetResponderRecipeResponderRulesSortOrderDesc,
}

var mappingListTargetResponderRecipeResponderRulesSortOrderEnumLowerCase = map[string]ListTargetResponderRecipeResponderRulesSortOrderEnum{
	"asc":  ListTargetResponderRecipeResponderRulesSortOrderAsc,
	"desc": ListTargetResponderRecipeResponderRulesSortOrderDesc,
}

// GetListTargetResponderRecipeResponderRulesSortOrderEnumValues Enumerates the set of values for ListTargetResponderRecipeResponderRulesSortOrderEnum
func GetListTargetResponderRecipeResponderRulesSortOrderEnumValues() []ListTargetResponderRecipeResponderRulesSortOrderEnum {
	values := make([]ListTargetResponderRecipeResponderRulesSortOrderEnum, 0)
	for _, v := range mappingListTargetResponderRecipeResponderRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetResponderRecipeResponderRulesSortOrderEnumStringValues Enumerates the set of values in String for ListTargetResponderRecipeResponderRulesSortOrderEnum
func GetListTargetResponderRecipeResponderRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetResponderRecipeResponderRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetResponderRecipeResponderRulesSortOrderEnum(val string) (ListTargetResponderRecipeResponderRulesSortOrderEnum, bool) {
	enum, ok := mappingListTargetResponderRecipeResponderRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetResponderRecipeResponderRulesSortByEnum Enum with underlying type: string
type ListTargetResponderRecipeResponderRulesSortByEnum string

// Set of constants representing the allowable values for ListTargetResponderRecipeResponderRulesSortByEnum
const (
	ListTargetResponderRecipeResponderRulesSortByDisplayname ListTargetResponderRecipeResponderRulesSortByEnum = "displayName"
	ListTargetResponderRecipeResponderRulesSortByRisklevel   ListTargetResponderRecipeResponderRulesSortByEnum = "riskLevel"
)

var mappingListTargetResponderRecipeResponderRulesSortByEnum = map[string]ListTargetResponderRecipeResponderRulesSortByEnum{
	"displayName": ListTargetResponderRecipeResponderRulesSortByDisplayname,
	"riskLevel":   ListTargetResponderRecipeResponderRulesSortByRisklevel,
}

var mappingListTargetResponderRecipeResponderRulesSortByEnumLowerCase = map[string]ListTargetResponderRecipeResponderRulesSortByEnum{
	"displayname": ListTargetResponderRecipeResponderRulesSortByDisplayname,
	"risklevel":   ListTargetResponderRecipeResponderRulesSortByRisklevel,
}

// GetListTargetResponderRecipeResponderRulesSortByEnumValues Enumerates the set of values for ListTargetResponderRecipeResponderRulesSortByEnum
func GetListTargetResponderRecipeResponderRulesSortByEnumValues() []ListTargetResponderRecipeResponderRulesSortByEnum {
	values := make([]ListTargetResponderRecipeResponderRulesSortByEnum, 0)
	for _, v := range mappingListTargetResponderRecipeResponderRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetResponderRecipeResponderRulesSortByEnumStringValues Enumerates the set of values in String for ListTargetResponderRecipeResponderRulesSortByEnum
func GetListTargetResponderRecipeResponderRulesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"riskLevel",
	}
}

// GetMappingListTargetResponderRecipeResponderRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetResponderRecipeResponderRulesSortByEnum(val string) (ListTargetResponderRecipeResponderRulesSortByEnum, bool) {
	enum, ok := mappingListTargetResponderRecipeResponderRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
