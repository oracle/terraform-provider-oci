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

// ListTargetDetectorRecipeDetectorRulesRequest wrapper for the ListTargetDetectorRecipeDetectorRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargetDetectorRecipeDetectorRules.go.html to see an example of how to use ListTargetDetectorRecipeDetectorRulesRequest.
type ListTargetDetectorRecipeDetectorRulesRequest struct {

	// OCID of the target
	TargetId *string `mandatory:"true" contributesTo:"path" name:"targetId"`

	// OCID of the target detector recipe.
	TargetDetectorRecipeId *string `mandatory:"true" contributesTo:"path" name:"targetDetectorRecipeId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListTargetDetectorRecipeDetectorRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListTargetDetectorRecipeDetectorRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetDetectorRecipeDetectorRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetDetectorRecipeDetectorRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetDetectorRecipeDetectorRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetDetectorRecipeDetectorRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetDetectorRecipeDetectorRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetDetectorRecipeDetectorRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTargetDetectorRecipeDetectorRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDetectorRecipeDetectorRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetDetectorRecipeDetectorRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDetectorRecipeDetectorRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetDetectorRecipeDetectorRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetDetectorRecipeDetectorRulesResponse wrapper for the ListTargetDetectorRecipeDetectorRules operation
type ListTargetDetectorRecipeDetectorRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetDetectorRecipeDetectorRuleCollection instances
	TargetDetectorRecipeDetectorRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTargetDetectorRecipeDetectorRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetDetectorRecipeDetectorRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum Enum with underlying type: string
type ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum
const (
	ListTargetDetectorRecipeDetectorRulesLifecycleStateCreating ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum = "CREATING"
	ListTargetDetectorRecipeDetectorRulesLifecycleStateUpdating ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum = "UPDATING"
	ListTargetDetectorRecipeDetectorRulesLifecycleStateActive   ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum = "ACTIVE"
	ListTargetDetectorRecipeDetectorRulesLifecycleStateInactive ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum = "INACTIVE"
	ListTargetDetectorRecipeDetectorRulesLifecycleStateDeleting ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum = "DELETING"
	ListTargetDetectorRecipeDetectorRulesLifecycleStateDeleted  ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum = "DELETED"
	ListTargetDetectorRecipeDetectorRulesLifecycleStateFailed   ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum = "FAILED"
)

var mappingListTargetDetectorRecipeDetectorRulesLifecycleStateEnum = map[string]ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum{
	"CREATING": ListTargetDetectorRecipeDetectorRulesLifecycleStateCreating,
	"UPDATING": ListTargetDetectorRecipeDetectorRulesLifecycleStateUpdating,
	"ACTIVE":   ListTargetDetectorRecipeDetectorRulesLifecycleStateActive,
	"INACTIVE": ListTargetDetectorRecipeDetectorRulesLifecycleStateInactive,
	"DELETING": ListTargetDetectorRecipeDetectorRulesLifecycleStateDeleting,
	"DELETED":  ListTargetDetectorRecipeDetectorRulesLifecycleStateDeleted,
	"FAILED":   ListTargetDetectorRecipeDetectorRulesLifecycleStateFailed,
}

var mappingListTargetDetectorRecipeDetectorRulesLifecycleStateEnumLowerCase = map[string]ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum{
	"creating": ListTargetDetectorRecipeDetectorRulesLifecycleStateCreating,
	"updating": ListTargetDetectorRecipeDetectorRulesLifecycleStateUpdating,
	"active":   ListTargetDetectorRecipeDetectorRulesLifecycleStateActive,
	"inactive": ListTargetDetectorRecipeDetectorRulesLifecycleStateInactive,
	"deleting": ListTargetDetectorRecipeDetectorRulesLifecycleStateDeleting,
	"deleted":  ListTargetDetectorRecipeDetectorRulesLifecycleStateDeleted,
	"failed":   ListTargetDetectorRecipeDetectorRulesLifecycleStateFailed,
}

// GetListTargetDetectorRecipeDetectorRulesLifecycleStateEnumValues Enumerates the set of values for ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum
func GetListTargetDetectorRecipeDetectorRulesLifecycleStateEnumValues() []ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum {
	values := make([]ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum, 0)
	for _, v := range mappingListTargetDetectorRecipeDetectorRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDetectorRecipeDetectorRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum
func GetListTargetDetectorRecipeDetectorRulesLifecycleStateEnumStringValues() []string {
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

// GetMappingListTargetDetectorRecipeDetectorRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDetectorRecipeDetectorRulesLifecycleStateEnum(val string) (ListTargetDetectorRecipeDetectorRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListTargetDetectorRecipeDetectorRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDetectorRecipeDetectorRulesSortOrderEnum Enum with underlying type: string
type ListTargetDetectorRecipeDetectorRulesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetDetectorRecipeDetectorRulesSortOrderEnum
const (
	ListTargetDetectorRecipeDetectorRulesSortOrderAsc  ListTargetDetectorRecipeDetectorRulesSortOrderEnum = "ASC"
	ListTargetDetectorRecipeDetectorRulesSortOrderDesc ListTargetDetectorRecipeDetectorRulesSortOrderEnum = "DESC"
)

var mappingListTargetDetectorRecipeDetectorRulesSortOrderEnum = map[string]ListTargetDetectorRecipeDetectorRulesSortOrderEnum{
	"ASC":  ListTargetDetectorRecipeDetectorRulesSortOrderAsc,
	"DESC": ListTargetDetectorRecipeDetectorRulesSortOrderDesc,
}

var mappingListTargetDetectorRecipeDetectorRulesSortOrderEnumLowerCase = map[string]ListTargetDetectorRecipeDetectorRulesSortOrderEnum{
	"asc":  ListTargetDetectorRecipeDetectorRulesSortOrderAsc,
	"desc": ListTargetDetectorRecipeDetectorRulesSortOrderDesc,
}

// GetListTargetDetectorRecipeDetectorRulesSortOrderEnumValues Enumerates the set of values for ListTargetDetectorRecipeDetectorRulesSortOrderEnum
func GetListTargetDetectorRecipeDetectorRulesSortOrderEnumValues() []ListTargetDetectorRecipeDetectorRulesSortOrderEnum {
	values := make([]ListTargetDetectorRecipeDetectorRulesSortOrderEnum, 0)
	for _, v := range mappingListTargetDetectorRecipeDetectorRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDetectorRecipeDetectorRulesSortOrderEnumStringValues Enumerates the set of values in String for ListTargetDetectorRecipeDetectorRulesSortOrderEnum
func GetListTargetDetectorRecipeDetectorRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetDetectorRecipeDetectorRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDetectorRecipeDetectorRulesSortOrderEnum(val string) (ListTargetDetectorRecipeDetectorRulesSortOrderEnum, bool) {
	enum, ok := mappingListTargetDetectorRecipeDetectorRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDetectorRecipeDetectorRulesSortByEnum Enum with underlying type: string
type ListTargetDetectorRecipeDetectorRulesSortByEnum string

// Set of constants representing the allowable values for ListTargetDetectorRecipeDetectorRulesSortByEnum
const (
	ListTargetDetectorRecipeDetectorRulesSortByDisplayname ListTargetDetectorRecipeDetectorRulesSortByEnum = "displayName"
	ListTargetDetectorRecipeDetectorRulesSortByRisklevel   ListTargetDetectorRecipeDetectorRulesSortByEnum = "riskLevel"
)

var mappingListTargetDetectorRecipeDetectorRulesSortByEnum = map[string]ListTargetDetectorRecipeDetectorRulesSortByEnum{
	"displayName": ListTargetDetectorRecipeDetectorRulesSortByDisplayname,
	"riskLevel":   ListTargetDetectorRecipeDetectorRulesSortByRisklevel,
}

var mappingListTargetDetectorRecipeDetectorRulesSortByEnumLowerCase = map[string]ListTargetDetectorRecipeDetectorRulesSortByEnum{
	"displayname": ListTargetDetectorRecipeDetectorRulesSortByDisplayname,
	"risklevel":   ListTargetDetectorRecipeDetectorRulesSortByRisklevel,
}

// GetListTargetDetectorRecipeDetectorRulesSortByEnumValues Enumerates the set of values for ListTargetDetectorRecipeDetectorRulesSortByEnum
func GetListTargetDetectorRecipeDetectorRulesSortByEnumValues() []ListTargetDetectorRecipeDetectorRulesSortByEnum {
	values := make([]ListTargetDetectorRecipeDetectorRulesSortByEnum, 0)
	for _, v := range mappingListTargetDetectorRecipeDetectorRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDetectorRecipeDetectorRulesSortByEnumStringValues Enumerates the set of values in String for ListTargetDetectorRecipeDetectorRulesSortByEnum
func GetListTargetDetectorRecipeDetectorRulesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"riskLevel",
	}
}

// GetMappingListTargetDetectorRecipeDetectorRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDetectorRecipeDetectorRulesSortByEnum(val string) (ListTargetDetectorRecipeDetectorRulesSortByEnum, bool) {
	enum, ok := mappingListTargetDetectorRecipeDetectorRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
