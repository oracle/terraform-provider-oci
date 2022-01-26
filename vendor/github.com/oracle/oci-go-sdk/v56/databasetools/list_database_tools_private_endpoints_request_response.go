// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetools

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDatabaseToolsPrivateEndpointsRequest wrapper for the ListDatabaseToolsPrivateEndpoints operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetools/ListDatabaseToolsPrivateEndpoints.go.html to see an example of how to use ListDatabaseToolsPrivateEndpointsRequest.
type ListDatabaseToolsPrivateEndpointsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources their subnetId matches the given subnetId.
	SubnetId *string `mandatory:"false" contributesTo:"query" name:"subnetId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseToolsPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDatabaseToolsPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources their type matches the given type.
	EndpointServiceId *string `mandatory:"false" contributesTo:"query" name:"endpointServiceId"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState ListDatabaseToolsPrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseToolsPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseToolsPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseToolsPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseToolsPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDatabaseToolsPrivateEndpointsResponse wrapper for the ListDatabaseToolsPrivateEndpoints operation
type ListDatabaseToolsPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseToolsPrivateEndpointCollection instances
	DatabaseToolsPrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseToolsPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseToolsPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseToolsPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListDatabaseToolsPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseToolsPrivateEndpointsSortOrderEnum
const (
	ListDatabaseToolsPrivateEndpointsSortOrderAsc  ListDatabaseToolsPrivateEndpointsSortOrderEnum = "ASC"
	ListDatabaseToolsPrivateEndpointsSortOrderDesc ListDatabaseToolsPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListDatabaseToolsPrivateEndpointsSortOrder = map[string]ListDatabaseToolsPrivateEndpointsSortOrderEnum{
	"ASC":  ListDatabaseToolsPrivateEndpointsSortOrderAsc,
	"DESC": ListDatabaseToolsPrivateEndpointsSortOrderDesc,
}

// GetListDatabaseToolsPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListDatabaseToolsPrivateEndpointsSortOrderEnum
func GetListDatabaseToolsPrivateEndpointsSortOrderEnumValues() []ListDatabaseToolsPrivateEndpointsSortOrderEnum {
	values := make([]ListDatabaseToolsPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseToolsPrivateEndpointsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDatabaseToolsPrivateEndpointsSortByEnum Enum with underlying type: string
type ListDatabaseToolsPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseToolsPrivateEndpointsSortByEnum
const (
	ListDatabaseToolsPrivateEndpointsSortByTimecreated ListDatabaseToolsPrivateEndpointsSortByEnum = "timeCreated"
	ListDatabaseToolsPrivateEndpointsSortByDisplayname ListDatabaseToolsPrivateEndpointsSortByEnum = "displayName"
)

var mappingListDatabaseToolsPrivateEndpointsSortBy = map[string]ListDatabaseToolsPrivateEndpointsSortByEnum{
	"timeCreated": ListDatabaseToolsPrivateEndpointsSortByTimecreated,
	"displayName": ListDatabaseToolsPrivateEndpointsSortByDisplayname,
}

// GetListDatabaseToolsPrivateEndpointsSortByEnumValues Enumerates the set of values for ListDatabaseToolsPrivateEndpointsSortByEnum
func GetListDatabaseToolsPrivateEndpointsSortByEnumValues() []ListDatabaseToolsPrivateEndpointsSortByEnum {
	values := make([]ListDatabaseToolsPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListDatabaseToolsPrivateEndpointsSortBy {
		values = append(values, v)
	}
	return values
}

// ListDatabaseToolsPrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListDatabaseToolsPrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseToolsPrivateEndpointsLifecycleStateEnum
const (
	ListDatabaseToolsPrivateEndpointsLifecycleStateCreating ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "CREATING"
	ListDatabaseToolsPrivateEndpointsLifecycleStateUpdating ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListDatabaseToolsPrivateEndpointsLifecycleStateActive   ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListDatabaseToolsPrivateEndpointsLifecycleStateDeleting ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "DELETING"
	ListDatabaseToolsPrivateEndpointsLifecycleStateDeleted  ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "DELETED"
	ListDatabaseToolsPrivateEndpointsLifecycleStateFailed   ListDatabaseToolsPrivateEndpointsLifecycleStateEnum = "FAILED"
)

var mappingListDatabaseToolsPrivateEndpointsLifecycleState = map[string]ListDatabaseToolsPrivateEndpointsLifecycleStateEnum{
	"CREATING": ListDatabaseToolsPrivateEndpointsLifecycleStateCreating,
	"UPDATING": ListDatabaseToolsPrivateEndpointsLifecycleStateUpdating,
	"ACTIVE":   ListDatabaseToolsPrivateEndpointsLifecycleStateActive,
	"DELETING": ListDatabaseToolsPrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListDatabaseToolsPrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListDatabaseToolsPrivateEndpointsLifecycleStateFailed,
}

// GetListDatabaseToolsPrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListDatabaseToolsPrivateEndpointsLifecycleStateEnum
func GetListDatabaseToolsPrivateEndpointsLifecycleStateEnumValues() []ListDatabaseToolsPrivateEndpointsLifecycleStateEnum {
	values := make([]ListDatabaseToolsPrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseToolsPrivateEndpointsLifecycleState {
		values = append(values, v)
	}
	return values
}
