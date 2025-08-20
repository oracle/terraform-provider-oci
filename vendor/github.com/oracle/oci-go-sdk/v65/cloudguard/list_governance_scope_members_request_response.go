// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGovernanceScopeMembersRequest wrapper for the ListGovernanceScopeMembers operation
type ListGovernanceScopeMembersRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the detection id.
	DetectionId *string `mandatory:"false" contributesTo:"query" name:"detectionId"`

	// The OCID of the subject tenantId in which to list resources
	SubjectTenantId *string `mandatory:"false" contributesTo:"query" name:"subjectTenantId"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListGovernanceScopeMembersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListGovernanceScopeMembersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListGovernanceScopeMembersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGovernanceScopeMembersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGovernanceScopeMembersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGovernanceScopeMembersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGovernanceScopeMembersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGovernanceScopeMembersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGovernanceScopeMembersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListGovernanceScopeMembersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceScopeMembersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGovernanceScopeMembersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceScopeMembersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGovernanceScopeMembersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGovernanceScopeMembersResponse wrapper for the ListGovernanceScopeMembers operation
type ListGovernanceScopeMembersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GovernanceScopeMemberCollection instances
	GovernanceScopeMemberCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGovernanceScopeMembersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGovernanceScopeMembersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGovernanceScopeMembersLifecycleStateEnum Enum with underlying type: string
type ListGovernanceScopeMembersLifecycleStateEnum string

// Set of constants representing the allowable values for ListGovernanceScopeMembersLifecycleStateEnum
const (
	ListGovernanceScopeMembersLifecycleStateCreating ListGovernanceScopeMembersLifecycleStateEnum = "CREATING"
	ListGovernanceScopeMembersLifecycleStateUpdating ListGovernanceScopeMembersLifecycleStateEnum = "UPDATING"
	ListGovernanceScopeMembersLifecycleStateActive   ListGovernanceScopeMembersLifecycleStateEnum = "ACTIVE"
	ListGovernanceScopeMembersLifecycleStateInactive ListGovernanceScopeMembersLifecycleStateEnum = "INACTIVE"
	ListGovernanceScopeMembersLifecycleStateDeleting ListGovernanceScopeMembersLifecycleStateEnum = "DELETING"
	ListGovernanceScopeMembersLifecycleStateDeleted  ListGovernanceScopeMembersLifecycleStateEnum = "DELETED"
	ListGovernanceScopeMembersLifecycleStateFailed   ListGovernanceScopeMembersLifecycleStateEnum = "FAILED"
)

var mappingListGovernanceScopeMembersLifecycleStateEnum = map[string]ListGovernanceScopeMembersLifecycleStateEnum{
	"CREATING": ListGovernanceScopeMembersLifecycleStateCreating,
	"UPDATING": ListGovernanceScopeMembersLifecycleStateUpdating,
	"ACTIVE":   ListGovernanceScopeMembersLifecycleStateActive,
	"INACTIVE": ListGovernanceScopeMembersLifecycleStateInactive,
	"DELETING": ListGovernanceScopeMembersLifecycleStateDeleting,
	"DELETED":  ListGovernanceScopeMembersLifecycleStateDeleted,
	"FAILED":   ListGovernanceScopeMembersLifecycleStateFailed,
}

var mappingListGovernanceScopeMembersLifecycleStateEnumLowerCase = map[string]ListGovernanceScopeMembersLifecycleStateEnum{
	"creating": ListGovernanceScopeMembersLifecycleStateCreating,
	"updating": ListGovernanceScopeMembersLifecycleStateUpdating,
	"active":   ListGovernanceScopeMembersLifecycleStateActive,
	"inactive": ListGovernanceScopeMembersLifecycleStateInactive,
	"deleting": ListGovernanceScopeMembersLifecycleStateDeleting,
	"deleted":  ListGovernanceScopeMembersLifecycleStateDeleted,
	"failed":   ListGovernanceScopeMembersLifecycleStateFailed,
}

// GetListGovernanceScopeMembersLifecycleStateEnumValues Enumerates the set of values for ListGovernanceScopeMembersLifecycleStateEnum
func GetListGovernanceScopeMembersLifecycleStateEnumValues() []ListGovernanceScopeMembersLifecycleStateEnum {
	values := make([]ListGovernanceScopeMembersLifecycleStateEnum, 0)
	for _, v := range mappingListGovernanceScopeMembersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceScopeMembersLifecycleStateEnumStringValues Enumerates the set of values in String for ListGovernanceScopeMembersLifecycleStateEnum
func GetListGovernanceScopeMembersLifecycleStateEnumStringValues() []string {
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

// GetMappingListGovernanceScopeMembersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceScopeMembersLifecycleStateEnum(val string) (ListGovernanceScopeMembersLifecycleStateEnum, bool) {
	enum, ok := mappingListGovernanceScopeMembersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceScopeMembersSortOrderEnum Enum with underlying type: string
type ListGovernanceScopeMembersSortOrderEnum string

// Set of constants representing the allowable values for ListGovernanceScopeMembersSortOrderEnum
const (
	ListGovernanceScopeMembersSortOrderAsc  ListGovernanceScopeMembersSortOrderEnum = "ASC"
	ListGovernanceScopeMembersSortOrderDesc ListGovernanceScopeMembersSortOrderEnum = "DESC"
)

var mappingListGovernanceScopeMembersSortOrderEnum = map[string]ListGovernanceScopeMembersSortOrderEnum{
	"ASC":  ListGovernanceScopeMembersSortOrderAsc,
	"DESC": ListGovernanceScopeMembersSortOrderDesc,
}

var mappingListGovernanceScopeMembersSortOrderEnumLowerCase = map[string]ListGovernanceScopeMembersSortOrderEnum{
	"asc":  ListGovernanceScopeMembersSortOrderAsc,
	"desc": ListGovernanceScopeMembersSortOrderDesc,
}

// GetListGovernanceScopeMembersSortOrderEnumValues Enumerates the set of values for ListGovernanceScopeMembersSortOrderEnum
func GetListGovernanceScopeMembersSortOrderEnumValues() []ListGovernanceScopeMembersSortOrderEnum {
	values := make([]ListGovernanceScopeMembersSortOrderEnum, 0)
	for _, v := range mappingListGovernanceScopeMembersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceScopeMembersSortOrderEnumStringValues Enumerates the set of values in String for ListGovernanceScopeMembersSortOrderEnum
func GetListGovernanceScopeMembersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGovernanceScopeMembersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceScopeMembersSortOrderEnum(val string) (ListGovernanceScopeMembersSortOrderEnum, bool) {
	enum, ok := mappingListGovernanceScopeMembersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceScopeMembersSortByEnum Enum with underlying type: string
type ListGovernanceScopeMembersSortByEnum string

// Set of constants representing the allowable values for ListGovernanceScopeMembersSortByEnum
const (
	ListGovernanceScopeMembersSortByTimecreated ListGovernanceScopeMembersSortByEnum = "timeCreated"
	ListGovernanceScopeMembersSortByDisplayname ListGovernanceScopeMembersSortByEnum = "displayName"
)

var mappingListGovernanceScopeMembersSortByEnum = map[string]ListGovernanceScopeMembersSortByEnum{
	"timeCreated": ListGovernanceScopeMembersSortByTimecreated,
	"displayName": ListGovernanceScopeMembersSortByDisplayname,
}

var mappingListGovernanceScopeMembersSortByEnumLowerCase = map[string]ListGovernanceScopeMembersSortByEnum{
	"timecreated": ListGovernanceScopeMembersSortByTimecreated,
	"displayname": ListGovernanceScopeMembersSortByDisplayname,
}

// GetListGovernanceScopeMembersSortByEnumValues Enumerates the set of values for ListGovernanceScopeMembersSortByEnum
func GetListGovernanceScopeMembersSortByEnumValues() []ListGovernanceScopeMembersSortByEnum {
	values := make([]ListGovernanceScopeMembersSortByEnum, 0)
	for _, v := range mappingListGovernanceScopeMembersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceScopeMembersSortByEnumStringValues Enumerates the set of values in String for ListGovernanceScopeMembersSortByEnum
func GetListGovernanceScopeMembersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListGovernanceScopeMembersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceScopeMembersSortByEnum(val string) (ListGovernanceScopeMembersSortByEnum, bool) {
	enum, ok := mappingListGovernanceScopeMembersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
