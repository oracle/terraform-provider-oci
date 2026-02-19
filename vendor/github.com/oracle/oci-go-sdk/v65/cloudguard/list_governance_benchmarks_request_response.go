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

// ListGovernanceBenchmarksRequest wrapper for the ListGovernanceBenchmarks operation
type ListGovernanceBenchmarksRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

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
	AccessLevel ListGovernanceBenchmarksAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListGovernanceBenchmarksLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListGovernanceBenchmarksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListGovernanceBenchmarksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGovernanceBenchmarksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGovernanceBenchmarksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListGovernanceBenchmarksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGovernanceBenchmarksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListGovernanceBenchmarksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListGovernanceBenchmarksAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListGovernanceBenchmarksAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceBenchmarksLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListGovernanceBenchmarksLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceBenchmarksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListGovernanceBenchmarksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListGovernanceBenchmarksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListGovernanceBenchmarksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListGovernanceBenchmarksResponse wrapper for the ListGovernanceBenchmarks operation
type ListGovernanceBenchmarksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of GovernanceBenchmarkCollection instances
	GovernanceBenchmarkCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGovernanceBenchmarksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGovernanceBenchmarksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGovernanceBenchmarksAccessLevelEnum Enum with underlying type: string
type ListGovernanceBenchmarksAccessLevelEnum string

// Set of constants representing the allowable values for ListGovernanceBenchmarksAccessLevelEnum
const (
	ListGovernanceBenchmarksAccessLevelRestricted ListGovernanceBenchmarksAccessLevelEnum = "RESTRICTED"
	ListGovernanceBenchmarksAccessLevelAccessible ListGovernanceBenchmarksAccessLevelEnum = "ACCESSIBLE"
)

var mappingListGovernanceBenchmarksAccessLevelEnum = map[string]ListGovernanceBenchmarksAccessLevelEnum{
	"RESTRICTED": ListGovernanceBenchmarksAccessLevelRestricted,
	"ACCESSIBLE": ListGovernanceBenchmarksAccessLevelAccessible,
}

var mappingListGovernanceBenchmarksAccessLevelEnumLowerCase = map[string]ListGovernanceBenchmarksAccessLevelEnum{
	"restricted": ListGovernanceBenchmarksAccessLevelRestricted,
	"accessible": ListGovernanceBenchmarksAccessLevelAccessible,
}

// GetListGovernanceBenchmarksAccessLevelEnumValues Enumerates the set of values for ListGovernanceBenchmarksAccessLevelEnum
func GetListGovernanceBenchmarksAccessLevelEnumValues() []ListGovernanceBenchmarksAccessLevelEnum {
	values := make([]ListGovernanceBenchmarksAccessLevelEnum, 0)
	for _, v := range mappingListGovernanceBenchmarksAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceBenchmarksAccessLevelEnumStringValues Enumerates the set of values in String for ListGovernanceBenchmarksAccessLevelEnum
func GetListGovernanceBenchmarksAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListGovernanceBenchmarksAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceBenchmarksAccessLevelEnum(val string) (ListGovernanceBenchmarksAccessLevelEnum, bool) {
	enum, ok := mappingListGovernanceBenchmarksAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceBenchmarksLifecycleStateEnum Enum with underlying type: string
type ListGovernanceBenchmarksLifecycleStateEnum string

// Set of constants representing the allowable values for ListGovernanceBenchmarksLifecycleStateEnum
const (
	ListGovernanceBenchmarksLifecycleStateCreating ListGovernanceBenchmarksLifecycleStateEnum = "CREATING"
	ListGovernanceBenchmarksLifecycleStateUpdating ListGovernanceBenchmarksLifecycleStateEnum = "UPDATING"
	ListGovernanceBenchmarksLifecycleStateActive   ListGovernanceBenchmarksLifecycleStateEnum = "ACTIVE"
	ListGovernanceBenchmarksLifecycleStateInactive ListGovernanceBenchmarksLifecycleStateEnum = "INACTIVE"
	ListGovernanceBenchmarksLifecycleStateDeleting ListGovernanceBenchmarksLifecycleStateEnum = "DELETING"
	ListGovernanceBenchmarksLifecycleStateDeleted  ListGovernanceBenchmarksLifecycleStateEnum = "DELETED"
	ListGovernanceBenchmarksLifecycleStateFailed   ListGovernanceBenchmarksLifecycleStateEnum = "FAILED"
)

var mappingListGovernanceBenchmarksLifecycleStateEnum = map[string]ListGovernanceBenchmarksLifecycleStateEnum{
	"CREATING": ListGovernanceBenchmarksLifecycleStateCreating,
	"UPDATING": ListGovernanceBenchmarksLifecycleStateUpdating,
	"ACTIVE":   ListGovernanceBenchmarksLifecycleStateActive,
	"INACTIVE": ListGovernanceBenchmarksLifecycleStateInactive,
	"DELETING": ListGovernanceBenchmarksLifecycleStateDeleting,
	"DELETED":  ListGovernanceBenchmarksLifecycleStateDeleted,
	"FAILED":   ListGovernanceBenchmarksLifecycleStateFailed,
}

var mappingListGovernanceBenchmarksLifecycleStateEnumLowerCase = map[string]ListGovernanceBenchmarksLifecycleStateEnum{
	"creating": ListGovernanceBenchmarksLifecycleStateCreating,
	"updating": ListGovernanceBenchmarksLifecycleStateUpdating,
	"active":   ListGovernanceBenchmarksLifecycleStateActive,
	"inactive": ListGovernanceBenchmarksLifecycleStateInactive,
	"deleting": ListGovernanceBenchmarksLifecycleStateDeleting,
	"deleted":  ListGovernanceBenchmarksLifecycleStateDeleted,
	"failed":   ListGovernanceBenchmarksLifecycleStateFailed,
}

// GetListGovernanceBenchmarksLifecycleStateEnumValues Enumerates the set of values for ListGovernanceBenchmarksLifecycleStateEnum
func GetListGovernanceBenchmarksLifecycleStateEnumValues() []ListGovernanceBenchmarksLifecycleStateEnum {
	values := make([]ListGovernanceBenchmarksLifecycleStateEnum, 0)
	for _, v := range mappingListGovernanceBenchmarksLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceBenchmarksLifecycleStateEnumStringValues Enumerates the set of values in String for ListGovernanceBenchmarksLifecycleStateEnum
func GetListGovernanceBenchmarksLifecycleStateEnumStringValues() []string {
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

// GetMappingListGovernanceBenchmarksLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceBenchmarksLifecycleStateEnum(val string) (ListGovernanceBenchmarksLifecycleStateEnum, bool) {
	enum, ok := mappingListGovernanceBenchmarksLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceBenchmarksSortOrderEnum Enum with underlying type: string
type ListGovernanceBenchmarksSortOrderEnum string

// Set of constants representing the allowable values for ListGovernanceBenchmarksSortOrderEnum
const (
	ListGovernanceBenchmarksSortOrderAsc  ListGovernanceBenchmarksSortOrderEnum = "ASC"
	ListGovernanceBenchmarksSortOrderDesc ListGovernanceBenchmarksSortOrderEnum = "DESC"
)

var mappingListGovernanceBenchmarksSortOrderEnum = map[string]ListGovernanceBenchmarksSortOrderEnum{
	"ASC":  ListGovernanceBenchmarksSortOrderAsc,
	"DESC": ListGovernanceBenchmarksSortOrderDesc,
}

var mappingListGovernanceBenchmarksSortOrderEnumLowerCase = map[string]ListGovernanceBenchmarksSortOrderEnum{
	"asc":  ListGovernanceBenchmarksSortOrderAsc,
	"desc": ListGovernanceBenchmarksSortOrderDesc,
}

// GetListGovernanceBenchmarksSortOrderEnumValues Enumerates the set of values for ListGovernanceBenchmarksSortOrderEnum
func GetListGovernanceBenchmarksSortOrderEnumValues() []ListGovernanceBenchmarksSortOrderEnum {
	values := make([]ListGovernanceBenchmarksSortOrderEnum, 0)
	for _, v := range mappingListGovernanceBenchmarksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceBenchmarksSortOrderEnumStringValues Enumerates the set of values in String for ListGovernanceBenchmarksSortOrderEnum
func GetListGovernanceBenchmarksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListGovernanceBenchmarksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceBenchmarksSortOrderEnum(val string) (ListGovernanceBenchmarksSortOrderEnum, bool) {
	enum, ok := mappingListGovernanceBenchmarksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListGovernanceBenchmarksSortByEnum Enum with underlying type: string
type ListGovernanceBenchmarksSortByEnum string

// Set of constants representing the allowable values for ListGovernanceBenchmarksSortByEnum
const (
	ListGovernanceBenchmarksSortByTimecreated ListGovernanceBenchmarksSortByEnum = "timeCreated"
	ListGovernanceBenchmarksSortByDisplayname ListGovernanceBenchmarksSortByEnum = "displayName"
)

var mappingListGovernanceBenchmarksSortByEnum = map[string]ListGovernanceBenchmarksSortByEnum{
	"timeCreated": ListGovernanceBenchmarksSortByTimecreated,
	"displayName": ListGovernanceBenchmarksSortByDisplayname,
}

var mappingListGovernanceBenchmarksSortByEnumLowerCase = map[string]ListGovernanceBenchmarksSortByEnum{
	"timecreated": ListGovernanceBenchmarksSortByTimecreated,
	"displayname": ListGovernanceBenchmarksSortByDisplayname,
}

// GetListGovernanceBenchmarksSortByEnumValues Enumerates the set of values for ListGovernanceBenchmarksSortByEnum
func GetListGovernanceBenchmarksSortByEnumValues() []ListGovernanceBenchmarksSortByEnum {
	values := make([]ListGovernanceBenchmarksSortByEnum, 0)
	for _, v := range mappingListGovernanceBenchmarksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListGovernanceBenchmarksSortByEnumStringValues Enumerates the set of values in String for ListGovernanceBenchmarksSortByEnum
func GetListGovernanceBenchmarksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListGovernanceBenchmarksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListGovernanceBenchmarksSortByEnum(val string) (ListGovernanceBenchmarksSortByEnum, bool) {
	enum, ok := mappingListGovernanceBenchmarksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
