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

// ListDataMaskRulesRequest wrapper for the ListDataMaskRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDataMaskRules.go.html to see an example of how to use ListDataMaskRulesRequest.
type ListDataMaskRulesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListDataMaskRulesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListDataMaskRulesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataMaskRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDataMaskRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The status of the dataMaskRule.
	DataMaskRuleStatus ListDataMaskRulesDataMaskRuleStatusEnum `mandatory:"false" contributesTo:"query" name:"dataMaskRuleStatus" omitEmpty:"true"`

	// OCID of target
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// OCID of iamGroup
	IamGroupId *string `mandatory:"false" contributesTo:"query" name:"iamGroupId"`

	// Type of target
	TargetType *string `mandatory:"false" contributesTo:"query" name:"targetType"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataMaskRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataMaskRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataMaskRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataMaskRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataMaskRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataMaskRulesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDataMaskRulesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataMaskRulesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListDataMaskRulesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataMaskRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataMaskRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataMaskRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataMaskRulesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataMaskRulesDataMaskRuleStatusEnum(string(request.DataMaskRuleStatus)); !ok && request.DataMaskRuleStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataMaskRuleStatus: %s. Supported values are: %s.", request.DataMaskRuleStatus, strings.Join(GetListDataMaskRulesDataMaskRuleStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataMaskRulesResponse wrapper for the ListDataMaskRules operation
type ListDataMaskRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataMaskRuleCollection instances
	DataMaskRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataMaskRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataMaskRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataMaskRulesLifecycleStateEnum Enum with underlying type: string
type ListDataMaskRulesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataMaskRulesLifecycleStateEnum
const (
	ListDataMaskRulesLifecycleStateCreating ListDataMaskRulesLifecycleStateEnum = "CREATING"
	ListDataMaskRulesLifecycleStateUpdating ListDataMaskRulesLifecycleStateEnum = "UPDATING"
	ListDataMaskRulesLifecycleStateActive   ListDataMaskRulesLifecycleStateEnum = "ACTIVE"
	ListDataMaskRulesLifecycleStateInactive ListDataMaskRulesLifecycleStateEnum = "INACTIVE"
	ListDataMaskRulesLifecycleStateDeleting ListDataMaskRulesLifecycleStateEnum = "DELETING"
	ListDataMaskRulesLifecycleStateDeleted  ListDataMaskRulesLifecycleStateEnum = "DELETED"
	ListDataMaskRulesLifecycleStateFailed   ListDataMaskRulesLifecycleStateEnum = "FAILED"
)

var mappingListDataMaskRulesLifecycleStateEnum = map[string]ListDataMaskRulesLifecycleStateEnum{
	"CREATING": ListDataMaskRulesLifecycleStateCreating,
	"UPDATING": ListDataMaskRulesLifecycleStateUpdating,
	"ACTIVE":   ListDataMaskRulesLifecycleStateActive,
	"INACTIVE": ListDataMaskRulesLifecycleStateInactive,
	"DELETING": ListDataMaskRulesLifecycleStateDeleting,
	"DELETED":  ListDataMaskRulesLifecycleStateDeleted,
	"FAILED":   ListDataMaskRulesLifecycleStateFailed,
}

var mappingListDataMaskRulesLifecycleStateEnumLowerCase = map[string]ListDataMaskRulesLifecycleStateEnum{
	"creating": ListDataMaskRulesLifecycleStateCreating,
	"updating": ListDataMaskRulesLifecycleStateUpdating,
	"active":   ListDataMaskRulesLifecycleStateActive,
	"inactive": ListDataMaskRulesLifecycleStateInactive,
	"deleting": ListDataMaskRulesLifecycleStateDeleting,
	"deleted":  ListDataMaskRulesLifecycleStateDeleted,
	"failed":   ListDataMaskRulesLifecycleStateFailed,
}

// GetListDataMaskRulesLifecycleStateEnumValues Enumerates the set of values for ListDataMaskRulesLifecycleStateEnum
func GetListDataMaskRulesLifecycleStateEnumValues() []ListDataMaskRulesLifecycleStateEnum {
	values := make([]ListDataMaskRulesLifecycleStateEnum, 0)
	for _, v := range mappingListDataMaskRulesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataMaskRulesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDataMaskRulesLifecycleStateEnum
func GetListDataMaskRulesLifecycleStateEnumStringValues() []string {
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

// GetMappingListDataMaskRulesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataMaskRulesLifecycleStateEnum(val string) (ListDataMaskRulesLifecycleStateEnum, bool) {
	enum, ok := mappingListDataMaskRulesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataMaskRulesAccessLevelEnum Enum with underlying type: string
type ListDataMaskRulesAccessLevelEnum string

// Set of constants representing the allowable values for ListDataMaskRulesAccessLevelEnum
const (
	ListDataMaskRulesAccessLevelRestricted ListDataMaskRulesAccessLevelEnum = "RESTRICTED"
	ListDataMaskRulesAccessLevelAccessible ListDataMaskRulesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListDataMaskRulesAccessLevelEnum = map[string]ListDataMaskRulesAccessLevelEnum{
	"RESTRICTED": ListDataMaskRulesAccessLevelRestricted,
	"ACCESSIBLE": ListDataMaskRulesAccessLevelAccessible,
}

var mappingListDataMaskRulesAccessLevelEnumLowerCase = map[string]ListDataMaskRulesAccessLevelEnum{
	"restricted": ListDataMaskRulesAccessLevelRestricted,
	"accessible": ListDataMaskRulesAccessLevelAccessible,
}

// GetListDataMaskRulesAccessLevelEnumValues Enumerates the set of values for ListDataMaskRulesAccessLevelEnum
func GetListDataMaskRulesAccessLevelEnumValues() []ListDataMaskRulesAccessLevelEnum {
	values := make([]ListDataMaskRulesAccessLevelEnum, 0)
	for _, v := range mappingListDataMaskRulesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataMaskRulesAccessLevelEnumStringValues Enumerates the set of values in String for ListDataMaskRulesAccessLevelEnum
func GetListDataMaskRulesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListDataMaskRulesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataMaskRulesAccessLevelEnum(val string) (ListDataMaskRulesAccessLevelEnum, bool) {
	enum, ok := mappingListDataMaskRulesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataMaskRulesSortOrderEnum Enum with underlying type: string
type ListDataMaskRulesSortOrderEnum string

// Set of constants representing the allowable values for ListDataMaskRulesSortOrderEnum
const (
	ListDataMaskRulesSortOrderAsc  ListDataMaskRulesSortOrderEnum = "ASC"
	ListDataMaskRulesSortOrderDesc ListDataMaskRulesSortOrderEnum = "DESC"
)

var mappingListDataMaskRulesSortOrderEnum = map[string]ListDataMaskRulesSortOrderEnum{
	"ASC":  ListDataMaskRulesSortOrderAsc,
	"DESC": ListDataMaskRulesSortOrderDesc,
}

var mappingListDataMaskRulesSortOrderEnumLowerCase = map[string]ListDataMaskRulesSortOrderEnum{
	"asc":  ListDataMaskRulesSortOrderAsc,
	"desc": ListDataMaskRulesSortOrderDesc,
}

// GetListDataMaskRulesSortOrderEnumValues Enumerates the set of values for ListDataMaskRulesSortOrderEnum
func GetListDataMaskRulesSortOrderEnumValues() []ListDataMaskRulesSortOrderEnum {
	values := make([]ListDataMaskRulesSortOrderEnum, 0)
	for _, v := range mappingListDataMaskRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataMaskRulesSortOrderEnumStringValues Enumerates the set of values in String for ListDataMaskRulesSortOrderEnum
func GetListDataMaskRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataMaskRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataMaskRulesSortOrderEnum(val string) (ListDataMaskRulesSortOrderEnum, bool) {
	enum, ok := mappingListDataMaskRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataMaskRulesSortByEnum Enum with underlying type: string
type ListDataMaskRulesSortByEnum string

// Set of constants representing the allowable values for ListDataMaskRulesSortByEnum
const (
	ListDataMaskRulesSortByTimecreated ListDataMaskRulesSortByEnum = "timeCreated"
	ListDataMaskRulesSortByDisplayname ListDataMaskRulesSortByEnum = "displayName"
)

var mappingListDataMaskRulesSortByEnum = map[string]ListDataMaskRulesSortByEnum{
	"timeCreated": ListDataMaskRulesSortByTimecreated,
	"displayName": ListDataMaskRulesSortByDisplayname,
}

var mappingListDataMaskRulesSortByEnumLowerCase = map[string]ListDataMaskRulesSortByEnum{
	"timecreated": ListDataMaskRulesSortByTimecreated,
	"displayname": ListDataMaskRulesSortByDisplayname,
}

// GetListDataMaskRulesSortByEnumValues Enumerates the set of values for ListDataMaskRulesSortByEnum
func GetListDataMaskRulesSortByEnumValues() []ListDataMaskRulesSortByEnum {
	values := make([]ListDataMaskRulesSortByEnum, 0)
	for _, v := range mappingListDataMaskRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataMaskRulesSortByEnumStringValues Enumerates the set of values in String for ListDataMaskRulesSortByEnum
func GetListDataMaskRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDataMaskRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataMaskRulesSortByEnum(val string) (ListDataMaskRulesSortByEnum, bool) {
	enum, ok := mappingListDataMaskRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataMaskRulesDataMaskRuleStatusEnum Enum with underlying type: string
type ListDataMaskRulesDataMaskRuleStatusEnum string

// Set of constants representing the allowable values for ListDataMaskRulesDataMaskRuleStatusEnum
const (
	ListDataMaskRulesDataMaskRuleStatusEnabled  ListDataMaskRulesDataMaskRuleStatusEnum = "ENABLED"
	ListDataMaskRulesDataMaskRuleStatusDisabled ListDataMaskRulesDataMaskRuleStatusEnum = "DISABLED"
)

var mappingListDataMaskRulesDataMaskRuleStatusEnum = map[string]ListDataMaskRulesDataMaskRuleStatusEnum{
	"ENABLED":  ListDataMaskRulesDataMaskRuleStatusEnabled,
	"DISABLED": ListDataMaskRulesDataMaskRuleStatusDisabled,
}

var mappingListDataMaskRulesDataMaskRuleStatusEnumLowerCase = map[string]ListDataMaskRulesDataMaskRuleStatusEnum{
	"enabled":  ListDataMaskRulesDataMaskRuleStatusEnabled,
	"disabled": ListDataMaskRulesDataMaskRuleStatusDisabled,
}

// GetListDataMaskRulesDataMaskRuleStatusEnumValues Enumerates the set of values for ListDataMaskRulesDataMaskRuleStatusEnum
func GetListDataMaskRulesDataMaskRuleStatusEnumValues() []ListDataMaskRulesDataMaskRuleStatusEnum {
	values := make([]ListDataMaskRulesDataMaskRuleStatusEnum, 0)
	for _, v := range mappingListDataMaskRulesDataMaskRuleStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataMaskRulesDataMaskRuleStatusEnumStringValues Enumerates the set of values in String for ListDataMaskRulesDataMaskRuleStatusEnum
func GetListDataMaskRulesDataMaskRuleStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingListDataMaskRulesDataMaskRuleStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataMaskRulesDataMaskRuleStatusEnum(val string) (ListDataMaskRulesDataMaskRuleStatusEnum, bool) {
	enum, ok := mappingListDataMaskRulesDataMaskRuleStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
