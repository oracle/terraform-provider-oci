// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operatoraccesscontrol

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListOperatorActionsRequest wrapper for the ListOperatorActions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListOperatorActions.go.html to see an example of how to use ListOperatorActionsRequest.
type ListOperatorActionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only lists of resources that match the entire given service type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// A filter to return only resources whose lifecycleState matches the given OperatorAction lifecycleState.
	LifecycleState ListOperatorActionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOperatorActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOperatorActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOperatorActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOperatorActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOperatorActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOperatorActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOperatorActionsResponse wrapper for the ListOperatorActions operation
type ListOperatorActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OperatorActionCollection instances
	OperatorActionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOperatorActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOperatorActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOperatorActionsLifecycleStateEnum Enum with underlying type: string
type ListOperatorActionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListOperatorActionsLifecycleStateEnum
const (
	ListOperatorActionsLifecycleStateActive   ListOperatorActionsLifecycleStateEnum = "ACTIVE"
	ListOperatorActionsLifecycleStateInactive ListOperatorActionsLifecycleStateEnum = "INACTIVE"
)

var mappingListOperatorActionsLifecycleState = map[string]ListOperatorActionsLifecycleStateEnum{
	"ACTIVE":   ListOperatorActionsLifecycleStateActive,
	"INACTIVE": ListOperatorActionsLifecycleStateInactive,
}

// GetListOperatorActionsLifecycleStateEnumValues Enumerates the set of values for ListOperatorActionsLifecycleStateEnum
func GetListOperatorActionsLifecycleStateEnumValues() []ListOperatorActionsLifecycleStateEnum {
	values := make([]ListOperatorActionsLifecycleStateEnum, 0)
	for _, v := range mappingListOperatorActionsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListOperatorActionsSortOrderEnum Enum with underlying type: string
type ListOperatorActionsSortOrderEnum string

// Set of constants representing the allowable values for ListOperatorActionsSortOrderEnum
const (
	ListOperatorActionsSortOrderAsc  ListOperatorActionsSortOrderEnum = "ASC"
	ListOperatorActionsSortOrderDesc ListOperatorActionsSortOrderEnum = "DESC"
)

var mappingListOperatorActionsSortOrder = map[string]ListOperatorActionsSortOrderEnum{
	"ASC":  ListOperatorActionsSortOrderAsc,
	"DESC": ListOperatorActionsSortOrderDesc,
}

// GetListOperatorActionsSortOrderEnumValues Enumerates the set of values for ListOperatorActionsSortOrderEnum
func GetListOperatorActionsSortOrderEnumValues() []ListOperatorActionsSortOrderEnum {
	values := make([]ListOperatorActionsSortOrderEnum, 0)
	for _, v := range mappingListOperatorActionsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListOperatorActionsSortByEnum Enum with underlying type: string
type ListOperatorActionsSortByEnum string

// Set of constants representing the allowable values for ListOperatorActionsSortByEnum
const (
	ListOperatorActionsSortByTimecreated ListOperatorActionsSortByEnum = "timeCreated"
	ListOperatorActionsSortByDisplayname ListOperatorActionsSortByEnum = "displayName"
)

var mappingListOperatorActionsSortBy = map[string]ListOperatorActionsSortByEnum{
	"timeCreated": ListOperatorActionsSortByTimecreated,
	"displayName": ListOperatorActionsSortByDisplayname,
}

// GetListOperatorActionsSortByEnumValues Enumerates the set of values for ListOperatorActionsSortByEnum
func GetListOperatorActionsSortByEnumValues() []ListOperatorActionsSortByEnum {
	values := make([]ListOperatorActionsSortByEnum, 0)
	for _, v := range mappingListOperatorActionsSortBy {
		values = append(values, v)
	}
	return values
}
