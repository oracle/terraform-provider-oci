// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedInstanceGroupsRequest wrapper for the ListManagedInstanceGroups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListManagedInstanceGroups.go.html to see an example of how to use ListManagedInstanceGroupsRequest.
type ListManagedInstanceGroupsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagedInstanceGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListManagedInstanceGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The current lifecycle state for the object.
	LifecycleState ListManagedInstanceGroupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OS family for which to list resources.
	OsFamily ListManagedInstanceGroupsOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstanceGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstanceGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstanceGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstanceGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedInstanceGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedInstanceGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedInstanceGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedInstanceGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListManagedInstanceGroupsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedInstanceGroupsOsFamilyEnum(string(request.OsFamily)); !ok && request.OsFamily != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OsFamily: %s. Supported values are: %s.", request.OsFamily, strings.Join(GetListManagedInstanceGroupsOsFamilyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedInstanceGroupsResponse wrapper for the ListManagedInstanceGroups operation
type ListManagedInstanceGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagedInstanceGroupSummary instances
	Items []ManagedInstanceGroupSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstanceGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstanceGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstanceGroupsSortOrderEnum Enum with underlying type: string
type ListManagedInstanceGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsSortOrderEnum
const (
	ListManagedInstanceGroupsSortOrderAsc  ListManagedInstanceGroupsSortOrderEnum = "ASC"
	ListManagedInstanceGroupsSortOrderDesc ListManagedInstanceGroupsSortOrderEnum = "DESC"
)

var mappingListManagedInstanceGroupsSortOrderEnum = map[string]ListManagedInstanceGroupsSortOrderEnum{
	"ASC":  ListManagedInstanceGroupsSortOrderAsc,
	"DESC": ListManagedInstanceGroupsSortOrderDesc,
}

var mappingListManagedInstanceGroupsSortOrderEnumLowerCase = map[string]ListManagedInstanceGroupsSortOrderEnum{
	"asc":  ListManagedInstanceGroupsSortOrderAsc,
	"desc": ListManagedInstanceGroupsSortOrderDesc,
}

// GetListManagedInstanceGroupsSortOrderEnumValues Enumerates the set of values for ListManagedInstanceGroupsSortOrderEnum
func GetListManagedInstanceGroupsSortOrderEnumValues() []ListManagedInstanceGroupsSortOrderEnum {
	values := make([]ListManagedInstanceGroupsSortOrderEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupsSortOrderEnum
func GetListManagedInstanceGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedInstanceGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupsSortOrderEnum(val string) (ListManagedInstanceGroupsSortOrderEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupsSortByEnum Enum with underlying type: string
type ListManagedInstanceGroupsSortByEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsSortByEnum
const (
	ListManagedInstanceGroupsSortByTimecreated ListManagedInstanceGroupsSortByEnum = "TIMECREATED"
	ListManagedInstanceGroupsSortByDisplayname ListManagedInstanceGroupsSortByEnum = "DISPLAYNAME"
)

var mappingListManagedInstanceGroupsSortByEnum = map[string]ListManagedInstanceGroupsSortByEnum{
	"TIMECREATED": ListManagedInstanceGroupsSortByTimecreated,
	"DISPLAYNAME": ListManagedInstanceGroupsSortByDisplayname,
}

var mappingListManagedInstanceGroupsSortByEnumLowerCase = map[string]ListManagedInstanceGroupsSortByEnum{
	"timecreated": ListManagedInstanceGroupsSortByTimecreated,
	"displayname": ListManagedInstanceGroupsSortByDisplayname,
}

// GetListManagedInstanceGroupsSortByEnumValues Enumerates the set of values for ListManagedInstanceGroupsSortByEnum
func GetListManagedInstanceGroupsSortByEnumValues() []ListManagedInstanceGroupsSortByEnum {
	values := make([]ListManagedInstanceGroupsSortByEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupsSortByEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupsSortByEnum
func GetListManagedInstanceGroupsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListManagedInstanceGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupsSortByEnum(val string) (ListManagedInstanceGroupsSortByEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupsLifecycleStateEnum Enum with underlying type: string
type ListManagedInstanceGroupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsLifecycleStateEnum
const (
	ListManagedInstanceGroupsLifecycleStateCreating ListManagedInstanceGroupsLifecycleStateEnum = "CREATING"
	ListManagedInstanceGroupsLifecycleStateUpdating ListManagedInstanceGroupsLifecycleStateEnum = "UPDATING"
	ListManagedInstanceGroupsLifecycleStateActive   ListManagedInstanceGroupsLifecycleStateEnum = "ACTIVE"
	ListManagedInstanceGroupsLifecycleStateDeleting ListManagedInstanceGroupsLifecycleStateEnum = "DELETING"
	ListManagedInstanceGroupsLifecycleStateDeleted  ListManagedInstanceGroupsLifecycleStateEnum = "DELETED"
	ListManagedInstanceGroupsLifecycleStateFailed   ListManagedInstanceGroupsLifecycleStateEnum = "FAILED"
)

var mappingListManagedInstanceGroupsLifecycleStateEnum = map[string]ListManagedInstanceGroupsLifecycleStateEnum{
	"CREATING": ListManagedInstanceGroupsLifecycleStateCreating,
	"UPDATING": ListManagedInstanceGroupsLifecycleStateUpdating,
	"ACTIVE":   ListManagedInstanceGroupsLifecycleStateActive,
	"DELETING": ListManagedInstanceGroupsLifecycleStateDeleting,
	"DELETED":  ListManagedInstanceGroupsLifecycleStateDeleted,
	"FAILED":   ListManagedInstanceGroupsLifecycleStateFailed,
}

var mappingListManagedInstanceGroupsLifecycleStateEnumLowerCase = map[string]ListManagedInstanceGroupsLifecycleStateEnum{
	"creating": ListManagedInstanceGroupsLifecycleStateCreating,
	"updating": ListManagedInstanceGroupsLifecycleStateUpdating,
	"active":   ListManagedInstanceGroupsLifecycleStateActive,
	"deleting": ListManagedInstanceGroupsLifecycleStateDeleting,
	"deleted":  ListManagedInstanceGroupsLifecycleStateDeleted,
	"failed":   ListManagedInstanceGroupsLifecycleStateFailed,
}

// GetListManagedInstanceGroupsLifecycleStateEnumValues Enumerates the set of values for ListManagedInstanceGroupsLifecycleStateEnum
func GetListManagedInstanceGroupsLifecycleStateEnumValues() []ListManagedInstanceGroupsLifecycleStateEnum {
	values := make([]ListManagedInstanceGroupsLifecycleStateEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupsLifecycleStateEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupsLifecycleStateEnum
func GetListManagedInstanceGroupsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListManagedInstanceGroupsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupsLifecycleStateEnum(val string) (ListManagedInstanceGroupsLifecycleStateEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedInstanceGroupsOsFamilyEnum Enum with underlying type: string
type ListManagedInstanceGroupsOsFamilyEnum string

// Set of constants representing the allowable values for ListManagedInstanceGroupsOsFamilyEnum
const (
	ListManagedInstanceGroupsOsFamilyLinux   ListManagedInstanceGroupsOsFamilyEnum = "LINUX"
	ListManagedInstanceGroupsOsFamilyWindows ListManagedInstanceGroupsOsFamilyEnum = "WINDOWS"
	ListManagedInstanceGroupsOsFamilyAll     ListManagedInstanceGroupsOsFamilyEnum = "ALL"
)

var mappingListManagedInstanceGroupsOsFamilyEnum = map[string]ListManagedInstanceGroupsOsFamilyEnum{
	"LINUX":   ListManagedInstanceGroupsOsFamilyLinux,
	"WINDOWS": ListManagedInstanceGroupsOsFamilyWindows,
	"ALL":     ListManagedInstanceGroupsOsFamilyAll,
}

var mappingListManagedInstanceGroupsOsFamilyEnumLowerCase = map[string]ListManagedInstanceGroupsOsFamilyEnum{
	"linux":   ListManagedInstanceGroupsOsFamilyLinux,
	"windows": ListManagedInstanceGroupsOsFamilyWindows,
	"all":     ListManagedInstanceGroupsOsFamilyAll,
}

// GetListManagedInstanceGroupsOsFamilyEnumValues Enumerates the set of values for ListManagedInstanceGroupsOsFamilyEnum
func GetListManagedInstanceGroupsOsFamilyEnumValues() []ListManagedInstanceGroupsOsFamilyEnum {
	values := make([]ListManagedInstanceGroupsOsFamilyEnum, 0)
	for _, v := range mappingListManagedInstanceGroupsOsFamilyEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedInstanceGroupsOsFamilyEnumStringValues Enumerates the set of values in String for ListManagedInstanceGroupsOsFamilyEnum
func GetListManagedInstanceGroupsOsFamilyEnumStringValues() []string {
	return []string{
		"LINUX",
		"WINDOWS",
		"ALL",
	}
}

// GetMappingListManagedInstanceGroupsOsFamilyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedInstanceGroupsOsFamilyEnum(val string) (ListManagedInstanceGroupsOsFamilyEnum, bool) {
	enum, ok := mappingListManagedInstanceGroupsOsFamilyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
