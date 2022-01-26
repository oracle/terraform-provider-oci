// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListTargetResponderRecipesRequest wrapper for the ListTargetResponderRecipes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargetResponderRecipes.go.html to see an example of how to use ListTargetResponderRecipesRequest.
type ListTargetResponderRecipesRequest struct {

	// OCID of target
	TargetId *string `mandatory:"true" contributesTo:"path" name:"targetId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListTargetResponderRecipesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTargetResponderRecipesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListTargetResponderRecipesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetResponderRecipesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetResponderRecipesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetResponderRecipesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetResponderRecipesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTargetResponderRecipesResponse wrapper for the ListTargetResponderRecipes operation
type ListTargetResponderRecipesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetResponderRecipeCollection instances
	TargetResponderRecipeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTargetResponderRecipesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetResponderRecipesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetResponderRecipesLifecycleStateEnum Enum with underlying type: string
type ListTargetResponderRecipesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTargetResponderRecipesLifecycleStateEnum
const (
	ListTargetResponderRecipesLifecycleStateCreating ListTargetResponderRecipesLifecycleStateEnum = "CREATING"
	ListTargetResponderRecipesLifecycleStateUpdating ListTargetResponderRecipesLifecycleStateEnum = "UPDATING"
	ListTargetResponderRecipesLifecycleStateActive   ListTargetResponderRecipesLifecycleStateEnum = "ACTIVE"
	ListTargetResponderRecipesLifecycleStateInactive ListTargetResponderRecipesLifecycleStateEnum = "INACTIVE"
	ListTargetResponderRecipesLifecycleStateDeleting ListTargetResponderRecipesLifecycleStateEnum = "DELETING"
	ListTargetResponderRecipesLifecycleStateDeleted  ListTargetResponderRecipesLifecycleStateEnum = "DELETED"
	ListTargetResponderRecipesLifecycleStateFailed   ListTargetResponderRecipesLifecycleStateEnum = "FAILED"
)

var mappingListTargetResponderRecipesLifecycleState = map[string]ListTargetResponderRecipesLifecycleStateEnum{
	"CREATING": ListTargetResponderRecipesLifecycleStateCreating,
	"UPDATING": ListTargetResponderRecipesLifecycleStateUpdating,
	"ACTIVE":   ListTargetResponderRecipesLifecycleStateActive,
	"INACTIVE": ListTargetResponderRecipesLifecycleStateInactive,
	"DELETING": ListTargetResponderRecipesLifecycleStateDeleting,
	"DELETED":  ListTargetResponderRecipesLifecycleStateDeleted,
	"FAILED":   ListTargetResponderRecipesLifecycleStateFailed,
}

// GetListTargetResponderRecipesLifecycleStateEnumValues Enumerates the set of values for ListTargetResponderRecipesLifecycleStateEnum
func GetListTargetResponderRecipesLifecycleStateEnumValues() []ListTargetResponderRecipesLifecycleStateEnum {
	values := make([]ListTargetResponderRecipesLifecycleStateEnum, 0)
	for _, v := range mappingListTargetResponderRecipesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListTargetResponderRecipesSortOrderEnum Enum with underlying type: string
type ListTargetResponderRecipesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetResponderRecipesSortOrderEnum
const (
	ListTargetResponderRecipesSortOrderAsc  ListTargetResponderRecipesSortOrderEnum = "ASC"
	ListTargetResponderRecipesSortOrderDesc ListTargetResponderRecipesSortOrderEnum = "DESC"
)

var mappingListTargetResponderRecipesSortOrder = map[string]ListTargetResponderRecipesSortOrderEnum{
	"ASC":  ListTargetResponderRecipesSortOrderAsc,
	"DESC": ListTargetResponderRecipesSortOrderDesc,
}

// GetListTargetResponderRecipesSortOrderEnumValues Enumerates the set of values for ListTargetResponderRecipesSortOrderEnum
func GetListTargetResponderRecipesSortOrderEnumValues() []ListTargetResponderRecipesSortOrderEnum {
	values := make([]ListTargetResponderRecipesSortOrderEnum, 0)
	for _, v := range mappingListTargetResponderRecipesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListTargetResponderRecipesSortByEnum Enum with underlying type: string
type ListTargetResponderRecipesSortByEnum string

// Set of constants representing the allowable values for ListTargetResponderRecipesSortByEnum
const (
	ListTargetResponderRecipesSortByTimecreated ListTargetResponderRecipesSortByEnum = "timeCreated"
	ListTargetResponderRecipesSortByDisplayname ListTargetResponderRecipesSortByEnum = "displayName"
)

var mappingListTargetResponderRecipesSortBy = map[string]ListTargetResponderRecipesSortByEnum{
	"timeCreated": ListTargetResponderRecipesSortByTimecreated,
	"displayName": ListTargetResponderRecipesSortByDisplayname,
}

// GetListTargetResponderRecipesSortByEnumValues Enumerates the set of values for ListTargetResponderRecipesSortByEnum
func GetListTargetResponderRecipesSortByEnumValues() []ListTargetResponderRecipesSortByEnum {
	values := make([]ListTargetResponderRecipesSortByEnum, 0)
	for _, v := range mappingListTargetResponderRecipesSortBy {
		values = append(values, v)
	}
	return values
}
