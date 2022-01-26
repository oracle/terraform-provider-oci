// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package visualbuilder

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListVbInstancesRequest wrapper for the ListVbInstances operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/visualbuilder/ListVbInstances.go.html to see an example of how to use ListVbInstancesRequest.
type ListVbInstancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Life cycle state to query on.
	LifecycleState ListVbInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListVbInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order
	// for timeCreated is descending. Default order for displayName is
	// ascending. If no value is specified timeCreated is default.
	SortBy ListVbInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVbInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVbInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVbInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVbInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListVbInstancesResponse wrapper for the ListVbInstances operation
type ListVbInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VbInstanceSummaryCollection instances
	VbInstanceSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, additional pages of results have been previously returned
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`
}

func (response ListVbInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVbInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVbInstancesLifecycleStateEnum Enum with underlying type: string
type ListVbInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListVbInstancesLifecycleStateEnum
const (
	ListVbInstancesLifecycleStateCreating ListVbInstancesLifecycleStateEnum = "CREATING"
	ListVbInstancesLifecycleStateUpdating ListVbInstancesLifecycleStateEnum = "UPDATING"
	ListVbInstancesLifecycleStateActive   ListVbInstancesLifecycleStateEnum = "ACTIVE"
	ListVbInstancesLifecycleStateInactive ListVbInstancesLifecycleStateEnum = "INACTIVE"
	ListVbInstancesLifecycleStateDeleting ListVbInstancesLifecycleStateEnum = "DELETING"
	ListVbInstancesLifecycleStateDeleted  ListVbInstancesLifecycleStateEnum = "DELETED"
	ListVbInstancesLifecycleStateFailed   ListVbInstancesLifecycleStateEnum = "FAILED"
)

var mappingListVbInstancesLifecycleState = map[string]ListVbInstancesLifecycleStateEnum{
	"CREATING": ListVbInstancesLifecycleStateCreating,
	"UPDATING": ListVbInstancesLifecycleStateUpdating,
	"ACTIVE":   ListVbInstancesLifecycleStateActive,
	"INACTIVE": ListVbInstancesLifecycleStateInactive,
	"DELETING": ListVbInstancesLifecycleStateDeleting,
	"DELETED":  ListVbInstancesLifecycleStateDeleted,
	"FAILED":   ListVbInstancesLifecycleStateFailed,
}

// GetListVbInstancesLifecycleStateEnumValues Enumerates the set of values for ListVbInstancesLifecycleStateEnum
func GetListVbInstancesLifecycleStateEnumValues() []ListVbInstancesLifecycleStateEnum {
	values := make([]ListVbInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListVbInstancesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListVbInstancesSortOrderEnum Enum with underlying type: string
type ListVbInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListVbInstancesSortOrderEnum
const (
	ListVbInstancesSortOrderAsc  ListVbInstancesSortOrderEnum = "ASC"
	ListVbInstancesSortOrderDesc ListVbInstancesSortOrderEnum = "DESC"
)

var mappingListVbInstancesSortOrder = map[string]ListVbInstancesSortOrderEnum{
	"ASC":  ListVbInstancesSortOrderAsc,
	"DESC": ListVbInstancesSortOrderDesc,
}

// GetListVbInstancesSortOrderEnumValues Enumerates the set of values for ListVbInstancesSortOrderEnum
func GetListVbInstancesSortOrderEnumValues() []ListVbInstancesSortOrderEnum {
	values := make([]ListVbInstancesSortOrderEnum, 0)
	for _, v := range mappingListVbInstancesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListVbInstancesSortByEnum Enum with underlying type: string
type ListVbInstancesSortByEnum string

// Set of constants representing the allowable values for ListVbInstancesSortByEnum
const (
	ListVbInstancesSortByTimecreated ListVbInstancesSortByEnum = "timeCreated"
	ListVbInstancesSortByDisplayname ListVbInstancesSortByEnum = "displayName"
)

var mappingListVbInstancesSortBy = map[string]ListVbInstancesSortByEnum{
	"timeCreated": ListVbInstancesSortByTimecreated,
	"displayName": ListVbInstancesSortByDisplayname,
}

// GetListVbInstancesSortByEnumValues Enumerates the set of values for ListVbInstancesSortByEnum
func GetListVbInstancesSortByEnumValues() []ListVbInstancesSortByEnum {
	values := make([]ListVbInstancesSortByEnum, 0)
	for _, v := range mappingListVbInstancesSortBy {
		values = append(values, v)
	}
	return values
}
