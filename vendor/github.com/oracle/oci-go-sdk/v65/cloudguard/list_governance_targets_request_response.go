// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListGovernanceTargetsRequest wrapper for the ListGovernanceTargets operation
type ListGovernanceTargetsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of the subject tenantId in which to list resources
	SubjectTenantId *string `mandatory:"false" contributesTo:"query" name:"subjectTenantId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListGovernanceTargetsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListGovernanceTargetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListGovernanceTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListGovernanceTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGovernanceTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGovernanceTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGovernanceTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGovernanceTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGovernanceTargetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGovernanceTargetsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListGovernanceTargetsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceTargetsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListGovernanceTargetsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceTargetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGovernanceTargetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceTargetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGovernanceTargetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGovernanceTargetsResponse wrapper for the ListGovernanceTargets operation
type ListGovernanceTargetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GovernanceTargetCollection instances
	GovernanceTargetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGovernanceTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGovernanceTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGovernanceTargetsAccessLevelEnum Enum with underlying type: string
type ListGovernanceTargetsAccessLevelEnum string

// Set of constants representing the allowable values for ListGovernanceTargetsAccessLevelEnum
const (
	ListGovernanceTargetsAccessLevelRestricted ListGovernanceTargetsAccessLevelEnum = "RESTRICTED"
	ListGovernanceTargetsAccessLevelAccessible ListGovernanceTargetsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListGovernanceTargetsAccessLevelEnum = map[string]ListGovernanceTargetsAccessLevelEnum{
	"RESTRICTED": ListGovernanceTargetsAccessLevelRestricted,
	"ACCESSIBLE": ListGovernanceTargetsAccessLevelAccessible,
}

var mappingListGovernanceTargetsAccessLevelEnumLowerCase = map[string]ListGovernanceTargetsAccessLevelEnum{
	"restricted": ListGovernanceTargetsAccessLevelRestricted,
	"accessible": ListGovernanceTargetsAccessLevelAccessible,
}

// GetListGovernanceTargetsAccessLevelEnumValues Enumerates the set of values for ListGovernanceTargetsAccessLevelEnum
func GetListGovernanceTargetsAccessLevelEnumValues() []ListGovernanceTargetsAccessLevelEnum {
	values := make([]ListGovernanceTargetsAccessLevelEnum, 0)
	for _, v := range mappingListGovernanceTargetsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceTargetsAccessLevelEnumStringValues Enumerates the set of values in String for ListGovernanceTargetsAccessLevelEnum
func GetListGovernanceTargetsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListGovernanceTargetsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceTargetsAccessLevelEnum(val string) (ListGovernanceTargetsAccessLevelEnum, bool) {
	enum, ok := mappingListGovernanceTargetsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceTargetsLifecycleStateEnum Enum with underlying type: string
type ListGovernanceTargetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListGovernanceTargetsLifecycleStateEnum
const (
	ListGovernanceTargetsLifecycleStateCreating ListGovernanceTargetsLifecycleStateEnum = "CREATING"
	ListGovernanceTargetsLifecycleStateUpdating ListGovernanceTargetsLifecycleStateEnum = "UPDATING"
	ListGovernanceTargetsLifecycleStateActive   ListGovernanceTargetsLifecycleStateEnum = "ACTIVE"
	ListGovernanceTargetsLifecycleStateInactive ListGovernanceTargetsLifecycleStateEnum = "INACTIVE"
	ListGovernanceTargetsLifecycleStateDeleting ListGovernanceTargetsLifecycleStateEnum = "DELETING"
	ListGovernanceTargetsLifecycleStateDeleted  ListGovernanceTargetsLifecycleStateEnum = "DELETED"
	ListGovernanceTargetsLifecycleStateFailed   ListGovernanceTargetsLifecycleStateEnum = "FAILED"
)

var mappingListGovernanceTargetsLifecycleStateEnum = map[string]ListGovernanceTargetsLifecycleStateEnum{
	"CREATING": ListGovernanceTargetsLifecycleStateCreating,
	"UPDATING": ListGovernanceTargetsLifecycleStateUpdating,
	"ACTIVE":   ListGovernanceTargetsLifecycleStateActive,
	"INACTIVE": ListGovernanceTargetsLifecycleStateInactive,
	"DELETING": ListGovernanceTargetsLifecycleStateDeleting,
	"DELETED":  ListGovernanceTargetsLifecycleStateDeleted,
	"FAILED":   ListGovernanceTargetsLifecycleStateFailed,
}

var mappingListGovernanceTargetsLifecycleStateEnumLowerCase = map[string]ListGovernanceTargetsLifecycleStateEnum{
	"creating": ListGovernanceTargetsLifecycleStateCreating,
	"updating": ListGovernanceTargetsLifecycleStateUpdating,
	"active":   ListGovernanceTargetsLifecycleStateActive,
	"inactive": ListGovernanceTargetsLifecycleStateInactive,
	"deleting": ListGovernanceTargetsLifecycleStateDeleting,
	"deleted":  ListGovernanceTargetsLifecycleStateDeleted,
	"failed":   ListGovernanceTargetsLifecycleStateFailed,
}

// GetListGovernanceTargetsLifecycleStateEnumValues Enumerates the set of values for ListGovernanceTargetsLifecycleStateEnum
func GetListGovernanceTargetsLifecycleStateEnumValues() []ListGovernanceTargetsLifecycleStateEnum {
	values := make([]ListGovernanceTargetsLifecycleStateEnum, 0)
	for _, v := range mappingListGovernanceTargetsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceTargetsLifecycleStateEnumStringValues Enumerates the set of values in String for ListGovernanceTargetsLifecycleStateEnum
func GetListGovernanceTargetsLifecycleStateEnumStringValues() []string {
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

// GetMappingListGovernanceTargetsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceTargetsLifecycleStateEnum(val string) (ListGovernanceTargetsLifecycleStateEnum, bool) {
	enum, ok := mappingListGovernanceTargetsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceTargetsSortOrderEnum Enum with underlying type: string
type ListGovernanceTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListGovernanceTargetsSortOrderEnum
const (
	ListGovernanceTargetsSortOrderAsc  ListGovernanceTargetsSortOrderEnum = "ASC"
	ListGovernanceTargetsSortOrderDesc ListGovernanceTargetsSortOrderEnum = "DESC"
)

var mappingListGovernanceTargetsSortOrderEnum = map[string]ListGovernanceTargetsSortOrderEnum{
	"ASC":  ListGovernanceTargetsSortOrderAsc,
	"DESC": ListGovernanceTargetsSortOrderDesc,
}

var mappingListGovernanceTargetsSortOrderEnumLowerCase = map[string]ListGovernanceTargetsSortOrderEnum{
	"asc":  ListGovernanceTargetsSortOrderAsc,
	"desc": ListGovernanceTargetsSortOrderDesc,
}

// GetListGovernanceTargetsSortOrderEnumValues Enumerates the set of values for ListGovernanceTargetsSortOrderEnum
func GetListGovernanceTargetsSortOrderEnumValues() []ListGovernanceTargetsSortOrderEnum {
	values := make([]ListGovernanceTargetsSortOrderEnum, 0)
	for _, v := range mappingListGovernanceTargetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceTargetsSortOrderEnumStringValues Enumerates the set of values in String for ListGovernanceTargetsSortOrderEnum
func GetListGovernanceTargetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGovernanceTargetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceTargetsSortOrderEnum(val string) (ListGovernanceTargetsSortOrderEnum, bool) {
	enum, ok := mappingListGovernanceTargetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceTargetsSortByEnum Enum with underlying type: string
type ListGovernanceTargetsSortByEnum string

// Set of constants representing the allowable values for ListGovernanceTargetsSortByEnum
const (
	ListGovernanceTargetsSortByTimecreated ListGovernanceTargetsSortByEnum = "timeCreated"
	ListGovernanceTargetsSortByDisplayname ListGovernanceTargetsSortByEnum = "displayName"
)

var mappingListGovernanceTargetsSortByEnum = map[string]ListGovernanceTargetsSortByEnum{
	"timeCreated": ListGovernanceTargetsSortByTimecreated,
	"displayName": ListGovernanceTargetsSortByDisplayname,
}

var mappingListGovernanceTargetsSortByEnumLowerCase = map[string]ListGovernanceTargetsSortByEnum{
	"timecreated": ListGovernanceTargetsSortByTimecreated,
	"displayname": ListGovernanceTargetsSortByDisplayname,
}

// GetListGovernanceTargetsSortByEnumValues Enumerates the set of values for ListGovernanceTargetsSortByEnum
func GetListGovernanceTargetsSortByEnumValues() []ListGovernanceTargetsSortByEnum {
	values := make([]ListGovernanceTargetsSortByEnum, 0)
	for _, v := range mappingListGovernanceTargetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceTargetsSortByEnumStringValues Enumerates the set of values in String for ListGovernanceTargetsSortByEnum
func GetListGovernanceTargetsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListGovernanceTargetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceTargetsSortByEnum(val string) (ListGovernanceTargetsSortByEnum, bool) {
	enum, ok := mappingListGovernanceTargetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
