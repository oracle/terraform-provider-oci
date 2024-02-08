// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedDatabaseGroupsRequest wrapper for the ListManagedDatabaseGroups operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedDatabaseGroups.go.html to see an example of how to use ListManagedDatabaseGroupsRequest.
type ListManagedDatabaseGroupsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The identifier of the resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The lifecycle state of a resource.
	LifecycleState ListManagedDatabaseGroupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListManagedDatabaseGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListManagedDatabaseGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedDatabaseGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedDatabaseGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedDatabaseGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedDatabaseGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedDatabaseGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedDatabaseGroupsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListManagedDatabaseGroupsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedDatabaseGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedDatabaseGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedDatabaseGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedDatabaseGroupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedDatabaseGroupsResponse wrapper for the ListManagedDatabaseGroups operation
type ListManagedDatabaseGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedDatabaseGroupCollection instances
	ManagedDatabaseGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedDatabaseGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedDatabaseGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedDatabaseGroupsLifecycleStateEnum Enum with underlying type: string
type ListManagedDatabaseGroupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagedDatabaseGroupsLifecycleStateEnum
const (
	ListManagedDatabaseGroupsLifecycleStateCreating ListManagedDatabaseGroupsLifecycleStateEnum = "CREATING"
	ListManagedDatabaseGroupsLifecycleStateUpdating ListManagedDatabaseGroupsLifecycleStateEnum = "UPDATING"
	ListManagedDatabaseGroupsLifecycleStateActive   ListManagedDatabaseGroupsLifecycleStateEnum = "ACTIVE"
	ListManagedDatabaseGroupsLifecycleStateDeleting ListManagedDatabaseGroupsLifecycleStateEnum = "DELETING"
	ListManagedDatabaseGroupsLifecycleStateDeleted  ListManagedDatabaseGroupsLifecycleStateEnum = "DELETED"
	ListManagedDatabaseGroupsLifecycleStateFailed   ListManagedDatabaseGroupsLifecycleStateEnum = "FAILED"
)

var mappingListManagedDatabaseGroupsLifecycleStateEnum = map[string]ListManagedDatabaseGroupsLifecycleStateEnum{
	"CREATING": ListManagedDatabaseGroupsLifecycleStateCreating,
	"UPDATING": ListManagedDatabaseGroupsLifecycleStateUpdating,
	"ACTIVE":   ListManagedDatabaseGroupsLifecycleStateActive,
	"DELETING": ListManagedDatabaseGroupsLifecycleStateDeleting,
	"DELETED":  ListManagedDatabaseGroupsLifecycleStateDeleted,
	"FAILED":   ListManagedDatabaseGroupsLifecycleStateFailed,
}

var mappingListManagedDatabaseGroupsLifecycleStateEnumLowerCase = map[string]ListManagedDatabaseGroupsLifecycleStateEnum{
	"creating": ListManagedDatabaseGroupsLifecycleStateCreating,
	"updating": ListManagedDatabaseGroupsLifecycleStateUpdating,
	"active":   ListManagedDatabaseGroupsLifecycleStateActive,
	"deleting": ListManagedDatabaseGroupsLifecycleStateDeleting,
	"deleted":  ListManagedDatabaseGroupsLifecycleStateDeleted,
	"failed":   ListManagedDatabaseGroupsLifecycleStateFailed,
}

// GetListManagedDatabaseGroupsLifecycleStateEnumValues Enumerates the set of values for ListManagedDatabaseGroupsLifecycleStateEnum
func GetListManagedDatabaseGroupsLifecycleStateEnumValues() []ListManagedDatabaseGroupsLifecycleStateEnum {
	values := make([]ListManagedDatabaseGroupsLifecycleStateEnum, 0)
	for _, v := range mappingListManagedDatabaseGroupsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedDatabaseGroupsLifecycleStateEnumStringValues Enumerates the set of values in String for ListManagedDatabaseGroupsLifecycleStateEnum
func GetListManagedDatabaseGroupsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListManagedDatabaseGroupsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedDatabaseGroupsLifecycleStateEnum(val string) (ListManagedDatabaseGroupsLifecycleStateEnum, bool) {
	enum, ok := mappingListManagedDatabaseGroupsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedDatabaseGroupsSortByEnum Enum with underlying type: string
type ListManagedDatabaseGroupsSortByEnum string

// Set of constants representing the allowable values for ListManagedDatabaseGroupsSortByEnum
const (
	ListManagedDatabaseGroupsSortByTimecreated ListManagedDatabaseGroupsSortByEnum = "TIMECREATED"
	ListManagedDatabaseGroupsSortByName        ListManagedDatabaseGroupsSortByEnum = "NAME"
)

var mappingListManagedDatabaseGroupsSortByEnum = map[string]ListManagedDatabaseGroupsSortByEnum{
	"TIMECREATED": ListManagedDatabaseGroupsSortByTimecreated,
	"NAME":        ListManagedDatabaseGroupsSortByName,
}

var mappingListManagedDatabaseGroupsSortByEnumLowerCase = map[string]ListManagedDatabaseGroupsSortByEnum{
	"timecreated": ListManagedDatabaseGroupsSortByTimecreated,
	"name":        ListManagedDatabaseGroupsSortByName,
}

// GetListManagedDatabaseGroupsSortByEnumValues Enumerates the set of values for ListManagedDatabaseGroupsSortByEnum
func GetListManagedDatabaseGroupsSortByEnumValues() []ListManagedDatabaseGroupsSortByEnum {
	values := make([]ListManagedDatabaseGroupsSortByEnum, 0)
	for _, v := range mappingListManagedDatabaseGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedDatabaseGroupsSortByEnumStringValues Enumerates the set of values in String for ListManagedDatabaseGroupsSortByEnum
func GetListManagedDatabaseGroupsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListManagedDatabaseGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedDatabaseGroupsSortByEnum(val string) (ListManagedDatabaseGroupsSortByEnum, bool) {
	enum, ok := mappingListManagedDatabaseGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedDatabaseGroupsSortOrderEnum Enum with underlying type: string
type ListManagedDatabaseGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListManagedDatabaseGroupsSortOrderEnum
const (
	ListManagedDatabaseGroupsSortOrderAsc  ListManagedDatabaseGroupsSortOrderEnum = "ASC"
	ListManagedDatabaseGroupsSortOrderDesc ListManagedDatabaseGroupsSortOrderEnum = "DESC"
)

var mappingListManagedDatabaseGroupsSortOrderEnum = map[string]ListManagedDatabaseGroupsSortOrderEnum{
	"ASC":  ListManagedDatabaseGroupsSortOrderAsc,
	"DESC": ListManagedDatabaseGroupsSortOrderDesc,
}

var mappingListManagedDatabaseGroupsSortOrderEnumLowerCase = map[string]ListManagedDatabaseGroupsSortOrderEnum{
	"asc":  ListManagedDatabaseGroupsSortOrderAsc,
	"desc": ListManagedDatabaseGroupsSortOrderDesc,
}

// GetListManagedDatabaseGroupsSortOrderEnumValues Enumerates the set of values for ListManagedDatabaseGroupsSortOrderEnum
func GetListManagedDatabaseGroupsSortOrderEnumValues() []ListManagedDatabaseGroupsSortOrderEnum {
	values := make([]ListManagedDatabaseGroupsSortOrderEnum, 0)
	for _, v := range mappingListManagedDatabaseGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedDatabaseGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListManagedDatabaseGroupsSortOrderEnum
func GetListManagedDatabaseGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedDatabaseGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedDatabaseGroupsSortOrderEnum(val string) (ListManagedDatabaseGroupsSortOrderEnum, bool) {
	enum, ok := mappingListManagedDatabaseGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
