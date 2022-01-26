// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetools

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDatabaseToolsConnectionsRequest wrapper for the ListDatabaseToolsConnections operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsConnections.go.html to see an example of how to use ListDatabaseToolsConnectionsRequest.
type ListDatabaseToolsConnectionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListDatabaseToolsConnectionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources their endpointServiceId matches the given endpointServiceId.
	Type []ConnectionTypeEnum `contributesTo:"query" name:"type" omitEmpty:"true" collectionFormat:"multi"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsConnectionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsConnectionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsConnectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsConnectionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsConnectionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsConnectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDatabaseToolsConnectionsResponse wrapper for the ListDatabaseToolsConnections operation
type ListDatabaseToolsConnectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsConnectionCollection instances
	DatabaseToolsConnectionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsConnectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsConnectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsConnectionsLifecycleStateEnum Enum with underlying type: string
type ListDatabaseToolsConnectionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseToolsConnectionsLifecycleStateEnum
const (
	ListDatabaseToolsConnectionsLifecycleStateCreating ListDatabaseToolsConnectionsLifecycleStateEnum = "CREATING"
	ListDatabaseToolsConnectionsLifecycleStateUpdating ListDatabaseToolsConnectionsLifecycleStateEnum = "UPDATING"
	ListDatabaseToolsConnectionsLifecycleStateActive   ListDatabaseToolsConnectionsLifecycleStateEnum = "ACTIVE"
	ListDatabaseToolsConnectionsLifecycleStateDeleting ListDatabaseToolsConnectionsLifecycleStateEnum = "DELETING"
	ListDatabaseToolsConnectionsLifecycleStateDeleted  ListDatabaseToolsConnectionsLifecycleStateEnum = "DELETED"
	ListDatabaseToolsConnectionsLifecycleStateFailed   ListDatabaseToolsConnectionsLifecycleStateEnum = "FAILED"
)

var mappingListDatabaseToolsConnectionsLifecycleState = map[string]ListDatabaseToolsConnectionsLifecycleStateEnum{
	"CREATING": ListDatabaseToolsConnectionsLifecycleStateCreating,
	"UPDATING": ListDatabaseToolsConnectionsLifecycleStateUpdating,
	"ACTIVE":   ListDatabaseToolsConnectionsLifecycleStateActive,
	"DELETING": ListDatabaseToolsConnectionsLifecycleStateDeleting,
	"DELETED":  ListDatabaseToolsConnectionsLifecycleStateDeleted,
	"FAILED":   ListDatabaseToolsConnectionsLifecycleStateFailed,
}

// GetListDatabaseToolsConnectionsLifecycleStateEnumValues Enumerates the set of values for ListDatabaseToolsConnectionsLifecycleStateEnum
func GetListDatabaseToolsConnectionsLifecycleStateEnumValues() []ListDatabaseToolsConnectionsLifecycleStateEnum {
	values := make([]ListDatabaseToolsConnectionsLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseToolsConnectionsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDatabaseToolsConnectionsSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsConnectionsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsConnectionsSortOrderEnum
const (
	ListDatabaseToolsConnectionsSortOrderAsc  ListDatabaseToolsConnectionsSortOrderEnum = "ASC"
	ListDatabaseToolsConnectionsSortOrderDesc ListDatabaseToolsConnectionsSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsConnectionsSortOrder = map[string]ListDatabaseToolsConnectionsSortOrderEnum{
	"ASC":  ListDatabaseToolsConnectionsSortOrderAsc,
	"DESC": ListDatabaseToolsConnectionsSortOrderDesc,
}

// GetListDatabaseToolsConnectionsSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsConnectionsSortOrderEnum
func GetListDatabaseToolsConnectionsSortOrderEnumValues() []ListDatabaseToolsConnectionsSortOrderEnum {
	values := make([]ListDatabaseToolsConnectionsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsConnectionsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDatabaseToolsConnectionsSortByEnum Enum with underlying type: string
type ListDatabaseToolsConnectionsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsConnectionsSortByEnum
const (
	ListDatabaseToolsConnectionsSortByTimecreated ListDatabaseToolsConnectionsSortByEnum = "timeCreated"
	ListDatabaseToolsConnectionsSortByDisplayname ListDatabaseToolsConnectionsSortByEnum = "displayName"
)

var mappingListDatabaseToolsConnectionsSortBy = map[string]ListDatabaseToolsConnectionsSortByEnum{
	"timeCreated": ListDatabaseToolsConnectionsSortByTimecreated,
	"displayName": ListDatabaseToolsConnectionsSortByDisplayname,
}

// GetListDatabaseToolsConnectionsSortByEnumValues Enumerates the set of values for ListDatabaseToolsConnectionsSortByEnum
func GetListDatabaseToolsConnectionsSortByEnumValues() []ListDatabaseToolsConnectionsSortByEnum {
	values := make([]ListDatabaseToolsConnectionsSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsConnectionsSortBy {
		values = append(values, v)
	}
	return values
}
