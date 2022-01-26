// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListManagedDatabaseGroupsRequest wrapper for the ListManagedDatabaseGroups operation
//
// See also
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

var mappingListManagedDatabaseGroupsLifecycleState = map[string]ListManagedDatabaseGroupsLifecycleStateEnum{
	"CREATING": ListManagedDatabaseGroupsLifecycleStateCreating,
	"UPDATING": ListManagedDatabaseGroupsLifecycleStateUpdating,
	"ACTIVE":   ListManagedDatabaseGroupsLifecycleStateActive,
	"DELETING": ListManagedDatabaseGroupsLifecycleStateDeleting,
	"DELETED":  ListManagedDatabaseGroupsLifecycleStateDeleted,
	"FAILED":   ListManagedDatabaseGroupsLifecycleStateFailed,
}

// GetListManagedDatabaseGroupsLifecycleStateEnumValues Enumerates the set of values for ListManagedDatabaseGroupsLifecycleStateEnum
func GetListManagedDatabaseGroupsLifecycleStateEnumValues() []ListManagedDatabaseGroupsLifecycleStateEnum {
	values := make([]ListManagedDatabaseGroupsLifecycleStateEnum, 0)
	for _, v := range mappingListManagedDatabaseGroupsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListManagedDatabaseGroupsSortByEnum Enum with underlying type: string
type ListManagedDatabaseGroupsSortByEnum string

// Set of constants representing the allowable values for ListManagedDatabaseGroupsSortByEnum
const (
	ListManagedDatabaseGroupsSortByTimecreated ListManagedDatabaseGroupsSortByEnum = "TIMECREATED"
	ListManagedDatabaseGroupsSortByName        ListManagedDatabaseGroupsSortByEnum = "NAME"
)

var mappingListManagedDatabaseGroupsSortBy = map[string]ListManagedDatabaseGroupsSortByEnum{
	"TIMECREATED": ListManagedDatabaseGroupsSortByTimecreated,
	"NAME":        ListManagedDatabaseGroupsSortByName,
}

// GetListManagedDatabaseGroupsSortByEnumValues Enumerates the set of values for ListManagedDatabaseGroupsSortByEnum
func GetListManagedDatabaseGroupsSortByEnumValues() []ListManagedDatabaseGroupsSortByEnum {
	values := make([]ListManagedDatabaseGroupsSortByEnum, 0)
	for _, v := range mappingListManagedDatabaseGroupsSortBy {
		values = append(values, v)
	}
	return values
}

// ListManagedDatabaseGroupsSortOrderEnum Enum with underlying type: string
type ListManagedDatabaseGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListManagedDatabaseGroupsSortOrderEnum
const (
	ListManagedDatabaseGroupsSortOrderAsc  ListManagedDatabaseGroupsSortOrderEnum = "ASC"
	ListManagedDatabaseGroupsSortOrderDesc ListManagedDatabaseGroupsSortOrderEnum = "DESC"
)

var mappingListManagedDatabaseGroupsSortOrder = map[string]ListManagedDatabaseGroupsSortOrderEnum{
	"ASC":  ListManagedDatabaseGroupsSortOrderAsc,
	"DESC": ListManagedDatabaseGroupsSortOrderDesc,
}

// GetListManagedDatabaseGroupsSortOrderEnumValues Enumerates the set of values for ListManagedDatabaseGroupsSortOrderEnum
func GetListManagedDatabaseGroupsSortOrderEnumValues() []ListManagedDatabaseGroupsSortOrderEnum {
	values := make([]ListManagedDatabaseGroupsSortOrderEnum, 0)
	for _, v := range mappingListManagedDatabaseGroupsSortOrder {
		values = append(values, v)
	}
	return values
}
