// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDrProtectionGroupsRequest wrapper for the ListDrProtectionGroups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListDrProtectionGroups.go.html to see an example of how to use ListDrProtectionGroupsRequest.
type ListDrProtectionGroupsRequest struct {

	// The ID (OCID) of the compartment in which to list resources.
	// Example: `ocid1.compartment.oc1..uniqueID`
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only DR protection groups that match the given lifecycle state.
	LifecycleState ListDrProtectionGroupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID of the DR protection group. Optional query param.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	DrProtectionGroupId *string `mandatory:"false" contributesTo:"query" name:"drProtectionGroupId"`

	// A filter to return only resources that match the given display name.
	// Example: `MyResourceDisplayName`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `100`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDrProtectionGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	// Example: `MyResourceDisplayName`
	SortBy ListDrProtectionGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The DR protection group Role.
	Role ListDrProtectionGroupsRoleEnum `mandatory:"false" contributesTo:"query" name:"role" omitEmpty:"true"`

	// A filter to return only DR protection groups that match the given lifecycle sub-state.
	LifecycleSubState ListDrProtectionGroupsLifecycleSubStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleSubState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDrProtectionGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDrProtectionGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDrProtectionGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDrProtectionGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDrProtectionGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDrProtectionGroupsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDrProtectionGroupsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrProtectionGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDrProtectionGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrProtectionGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDrProtectionGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrProtectionGroupsRoleEnum(string(request.Role)); !ok && request.Role != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Role: %s. Supported values are: %s.", request.Role, strings.Join(GetListDrProtectionGroupsRoleEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrProtectionGroupsLifecycleSubStateEnum(string(request.LifecycleSubState)); !ok && request.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", request.LifecycleSubState, strings.Join(GetListDrProtectionGroupsLifecycleSubStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDrProtectionGroupsResponse wrapper for the ListDrProtectionGroups operation
type ListDrProtectionGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DrProtectionGroupCollection instances
	DrProtectionGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDrProtectionGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDrProtectionGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDrProtectionGroupsLifecycleStateEnum Enum with underlying type: string
type ListDrProtectionGroupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDrProtectionGroupsLifecycleStateEnum
const (
	ListDrProtectionGroupsLifecycleStateCreating       ListDrProtectionGroupsLifecycleStateEnum = "CREATING"
	ListDrProtectionGroupsLifecycleStateActive         ListDrProtectionGroupsLifecycleStateEnum = "ACTIVE"
	ListDrProtectionGroupsLifecycleStateUpdating       ListDrProtectionGroupsLifecycleStateEnum = "UPDATING"
	ListDrProtectionGroupsLifecycleStateInactive       ListDrProtectionGroupsLifecycleStateEnum = "INACTIVE"
	ListDrProtectionGroupsLifecycleStateNeedsAttention ListDrProtectionGroupsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListDrProtectionGroupsLifecycleStateDeleting       ListDrProtectionGroupsLifecycleStateEnum = "DELETING"
	ListDrProtectionGroupsLifecycleStateDeleted        ListDrProtectionGroupsLifecycleStateEnum = "DELETED"
	ListDrProtectionGroupsLifecycleStateFailed         ListDrProtectionGroupsLifecycleStateEnum = "FAILED"
)

var mappingListDrProtectionGroupsLifecycleStateEnum = map[string]ListDrProtectionGroupsLifecycleStateEnum{
	"CREATING":        ListDrProtectionGroupsLifecycleStateCreating,
	"ACTIVE":          ListDrProtectionGroupsLifecycleStateActive,
	"UPDATING":        ListDrProtectionGroupsLifecycleStateUpdating,
	"INACTIVE":        ListDrProtectionGroupsLifecycleStateInactive,
	"NEEDS_ATTENTION": ListDrProtectionGroupsLifecycleStateNeedsAttention,
	"DELETING":        ListDrProtectionGroupsLifecycleStateDeleting,
	"DELETED":         ListDrProtectionGroupsLifecycleStateDeleted,
	"FAILED":          ListDrProtectionGroupsLifecycleStateFailed,
}

var mappingListDrProtectionGroupsLifecycleStateEnumLowerCase = map[string]ListDrProtectionGroupsLifecycleStateEnum{
	"creating":        ListDrProtectionGroupsLifecycleStateCreating,
	"active":          ListDrProtectionGroupsLifecycleStateActive,
	"updating":        ListDrProtectionGroupsLifecycleStateUpdating,
	"inactive":        ListDrProtectionGroupsLifecycleStateInactive,
	"needs_attention": ListDrProtectionGroupsLifecycleStateNeedsAttention,
	"deleting":        ListDrProtectionGroupsLifecycleStateDeleting,
	"deleted":         ListDrProtectionGroupsLifecycleStateDeleted,
	"failed":          ListDrProtectionGroupsLifecycleStateFailed,
}

// GetListDrProtectionGroupsLifecycleStateEnumValues Enumerates the set of values for ListDrProtectionGroupsLifecycleStateEnum
func GetListDrProtectionGroupsLifecycleStateEnumValues() []ListDrProtectionGroupsLifecycleStateEnum {
	values := make([]ListDrProtectionGroupsLifecycleStateEnum, 0)
	for _, v := range mappingListDrProtectionGroupsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrProtectionGroupsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDrProtectionGroupsLifecycleStateEnum
func GetListDrProtectionGroupsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"INACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListDrProtectionGroupsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrProtectionGroupsLifecycleStateEnum(val string) (ListDrProtectionGroupsLifecycleStateEnum, bool) {
	enum, ok := mappingListDrProtectionGroupsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrProtectionGroupsSortOrderEnum Enum with underlying type: string
type ListDrProtectionGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListDrProtectionGroupsSortOrderEnum
const (
	ListDrProtectionGroupsSortOrderAsc  ListDrProtectionGroupsSortOrderEnum = "ASC"
	ListDrProtectionGroupsSortOrderDesc ListDrProtectionGroupsSortOrderEnum = "DESC"
)

var mappingListDrProtectionGroupsSortOrderEnum = map[string]ListDrProtectionGroupsSortOrderEnum{
	"ASC":  ListDrProtectionGroupsSortOrderAsc,
	"DESC": ListDrProtectionGroupsSortOrderDesc,
}

var mappingListDrProtectionGroupsSortOrderEnumLowerCase = map[string]ListDrProtectionGroupsSortOrderEnum{
	"asc":  ListDrProtectionGroupsSortOrderAsc,
	"desc": ListDrProtectionGroupsSortOrderDesc,
}

// GetListDrProtectionGroupsSortOrderEnumValues Enumerates the set of values for ListDrProtectionGroupsSortOrderEnum
func GetListDrProtectionGroupsSortOrderEnumValues() []ListDrProtectionGroupsSortOrderEnum {
	values := make([]ListDrProtectionGroupsSortOrderEnum, 0)
	for _, v := range mappingListDrProtectionGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrProtectionGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListDrProtectionGroupsSortOrderEnum
func GetListDrProtectionGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDrProtectionGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrProtectionGroupsSortOrderEnum(val string) (ListDrProtectionGroupsSortOrderEnum, bool) {
	enum, ok := mappingListDrProtectionGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrProtectionGroupsSortByEnum Enum with underlying type: string
type ListDrProtectionGroupsSortByEnum string

// Set of constants representing the allowable values for ListDrProtectionGroupsSortByEnum
const (
	ListDrProtectionGroupsSortByTimecreated ListDrProtectionGroupsSortByEnum = "timeCreated"
	ListDrProtectionGroupsSortByDisplayname ListDrProtectionGroupsSortByEnum = "displayName"
)

var mappingListDrProtectionGroupsSortByEnum = map[string]ListDrProtectionGroupsSortByEnum{
	"timeCreated": ListDrProtectionGroupsSortByTimecreated,
	"displayName": ListDrProtectionGroupsSortByDisplayname,
}

var mappingListDrProtectionGroupsSortByEnumLowerCase = map[string]ListDrProtectionGroupsSortByEnum{
	"timecreated": ListDrProtectionGroupsSortByTimecreated,
	"displayname": ListDrProtectionGroupsSortByDisplayname,
}

// GetListDrProtectionGroupsSortByEnumValues Enumerates the set of values for ListDrProtectionGroupsSortByEnum
func GetListDrProtectionGroupsSortByEnumValues() []ListDrProtectionGroupsSortByEnum {
	values := make([]ListDrProtectionGroupsSortByEnum, 0)
	for _, v := range mappingListDrProtectionGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrProtectionGroupsSortByEnumStringValues Enumerates the set of values in String for ListDrProtectionGroupsSortByEnum
func GetListDrProtectionGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDrProtectionGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrProtectionGroupsSortByEnum(val string) (ListDrProtectionGroupsSortByEnum, bool) {
	enum, ok := mappingListDrProtectionGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrProtectionGroupsRoleEnum Enum with underlying type: string
type ListDrProtectionGroupsRoleEnum string

// Set of constants representing the allowable values for ListDrProtectionGroupsRoleEnum
const (
	ListDrProtectionGroupsRolePrimary      ListDrProtectionGroupsRoleEnum = "PRIMARY"
	ListDrProtectionGroupsRoleStandby      ListDrProtectionGroupsRoleEnum = "STANDBY"
	ListDrProtectionGroupsRoleUnconfigured ListDrProtectionGroupsRoleEnum = "UNCONFIGURED"
)

var mappingListDrProtectionGroupsRoleEnum = map[string]ListDrProtectionGroupsRoleEnum{
	"PRIMARY":      ListDrProtectionGroupsRolePrimary,
	"STANDBY":      ListDrProtectionGroupsRoleStandby,
	"UNCONFIGURED": ListDrProtectionGroupsRoleUnconfigured,
}

var mappingListDrProtectionGroupsRoleEnumLowerCase = map[string]ListDrProtectionGroupsRoleEnum{
	"primary":      ListDrProtectionGroupsRolePrimary,
	"standby":      ListDrProtectionGroupsRoleStandby,
	"unconfigured": ListDrProtectionGroupsRoleUnconfigured,
}

// GetListDrProtectionGroupsRoleEnumValues Enumerates the set of values for ListDrProtectionGroupsRoleEnum
func GetListDrProtectionGroupsRoleEnumValues() []ListDrProtectionGroupsRoleEnum {
	values := make([]ListDrProtectionGroupsRoleEnum, 0)
	for _, v := range mappingListDrProtectionGroupsRoleEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrProtectionGroupsRoleEnumStringValues Enumerates the set of values in String for ListDrProtectionGroupsRoleEnum
func GetListDrProtectionGroupsRoleEnumStringValues() []string {
	return []string{
		"PRIMARY",
		"STANDBY",
		"UNCONFIGURED",
	}
}

// GetMappingListDrProtectionGroupsRoleEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrProtectionGroupsRoleEnum(val string) (ListDrProtectionGroupsRoleEnum, bool) {
	enum, ok := mappingListDrProtectionGroupsRoleEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrProtectionGroupsLifecycleSubStateEnum Enum with underlying type: string
type ListDrProtectionGroupsLifecycleSubStateEnum string

// Set of constants representing the allowable values for ListDrProtectionGroupsLifecycleSubStateEnum
const (
	ListDrProtectionGroupsLifecycleSubStateDrDrillInProgress ListDrProtectionGroupsLifecycleSubStateEnum = "DR_DRILL_IN_PROGRESS"
)

var mappingListDrProtectionGroupsLifecycleSubStateEnum = map[string]ListDrProtectionGroupsLifecycleSubStateEnum{
	"DR_DRILL_IN_PROGRESS": ListDrProtectionGroupsLifecycleSubStateDrDrillInProgress,
}

var mappingListDrProtectionGroupsLifecycleSubStateEnumLowerCase = map[string]ListDrProtectionGroupsLifecycleSubStateEnum{
	"dr_drill_in_progress": ListDrProtectionGroupsLifecycleSubStateDrDrillInProgress,
}

// GetListDrProtectionGroupsLifecycleSubStateEnumValues Enumerates the set of values for ListDrProtectionGroupsLifecycleSubStateEnum
func GetListDrProtectionGroupsLifecycleSubStateEnumValues() []ListDrProtectionGroupsLifecycleSubStateEnum {
	values := make([]ListDrProtectionGroupsLifecycleSubStateEnum, 0)
	for _, v := range mappingListDrProtectionGroupsLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrProtectionGroupsLifecycleSubStateEnumStringValues Enumerates the set of values in String for ListDrProtectionGroupsLifecycleSubStateEnum
func GetListDrProtectionGroupsLifecycleSubStateEnumStringValues() []string {
	return []string{
		"DR_DRILL_IN_PROGRESS",
	}
}

// GetMappingListDrProtectionGroupsLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrProtectionGroupsLifecycleSubStateEnum(val string) (ListDrProtectionGroupsLifecycleSubStateEnum, bool) {
	enum, ok := mappingListDrProtectionGroupsLifecycleSubStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
