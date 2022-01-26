// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDetectorRecipesRequest wrapper for the ListDetectorRecipes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDetectorRecipes.go.html to see an example of how to use ListDetectorRecipesRequest.
type ListDetectorRecipesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Default is false.
	// When set to true, the list of all Oracle Managed Resources
	// Metadata supported by Cloud Guard are returned.
	ResourceMetadataOnly *bool `mandatory:"false" contributesTo:"query" name:"resourceMetadataOnly"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListDetectorRecipesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListDetectorRecipesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDetectorRecipesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListDetectorRecipesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDetectorRecipesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDetectorRecipesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDetectorRecipesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDetectorRecipesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDetectorRecipesResponse wrapper for the ListDetectorRecipes operation
type ListDetectorRecipesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DetectorRecipeCollection instances
	DetectorRecipeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDetectorRecipesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDetectorRecipesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDetectorRecipesLifecycleStateEnum Enum with underlying type: string
type ListDetectorRecipesLifecycleStateEnum string

// Set of constants representing the allowable values for ListDetectorRecipesLifecycleStateEnum
const (
	ListDetectorRecipesLifecycleStateCreating ListDetectorRecipesLifecycleStateEnum = "CREATING"
	ListDetectorRecipesLifecycleStateUpdating ListDetectorRecipesLifecycleStateEnum = "UPDATING"
	ListDetectorRecipesLifecycleStateActive   ListDetectorRecipesLifecycleStateEnum = "ACTIVE"
	ListDetectorRecipesLifecycleStateInactive ListDetectorRecipesLifecycleStateEnum = "INACTIVE"
	ListDetectorRecipesLifecycleStateDeleting ListDetectorRecipesLifecycleStateEnum = "DELETING"
	ListDetectorRecipesLifecycleStateDeleted  ListDetectorRecipesLifecycleStateEnum = "DELETED"
	ListDetectorRecipesLifecycleStateFailed   ListDetectorRecipesLifecycleStateEnum = "FAILED"
)

var mappingListDetectorRecipesLifecycleState = map[string]ListDetectorRecipesLifecycleStateEnum{
	"CREATING": ListDetectorRecipesLifecycleStateCreating,
	"UPDATING": ListDetectorRecipesLifecycleStateUpdating,
	"ACTIVE":   ListDetectorRecipesLifecycleStateActive,
	"INACTIVE": ListDetectorRecipesLifecycleStateInactive,
	"DELETING": ListDetectorRecipesLifecycleStateDeleting,
	"DELETED":  ListDetectorRecipesLifecycleStateDeleted,
	"FAILED":   ListDetectorRecipesLifecycleStateFailed,
}

// GetListDetectorRecipesLifecycleStateEnumValues Enumerates the set of values for ListDetectorRecipesLifecycleStateEnum
func GetListDetectorRecipesLifecycleStateEnumValues() []ListDetectorRecipesLifecycleStateEnum {
	values := make([]ListDetectorRecipesLifecycleStateEnum, 0)
	for _, v := range mappingListDetectorRecipesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDetectorRecipesAccessLevelEnum Enum with underlying type: string
type ListDetectorRecipesAccessLevelEnum string

// Set of constants representing the allowable values for ListDetectorRecipesAccessLevelEnum
const (
	ListDetectorRecipesAccessLevelRestricted ListDetectorRecipesAccessLevelEnum = "RESTRICTED"
	ListDetectorRecipesAccessLevelAccessible ListDetectorRecipesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListDetectorRecipesAccessLevel = map[string]ListDetectorRecipesAccessLevelEnum{
	"RESTRICTED": ListDetectorRecipesAccessLevelRestricted,
	"ACCESSIBLE": ListDetectorRecipesAccessLevelAccessible,
}

// GetListDetectorRecipesAccessLevelEnumValues Enumerates the set of values for ListDetectorRecipesAccessLevelEnum
func GetListDetectorRecipesAccessLevelEnumValues() []ListDetectorRecipesAccessLevelEnum {
	values := make([]ListDetectorRecipesAccessLevelEnum, 0)
	for _, v := range mappingListDetectorRecipesAccessLevel {
		values = append(values, v)
	}
	return values
}

// ListDetectorRecipesSortOrderEnum Enum with underlying type: string
type ListDetectorRecipesSortOrderEnum string

// Set of constants representing the allowable values for ListDetectorRecipesSortOrderEnum
const (
	ListDetectorRecipesSortOrderAsc  ListDetectorRecipesSortOrderEnum = "ASC"
	ListDetectorRecipesSortOrderDesc ListDetectorRecipesSortOrderEnum = "DESC"
)

var mappingListDetectorRecipesSortOrder = map[string]ListDetectorRecipesSortOrderEnum{
	"ASC":  ListDetectorRecipesSortOrderAsc,
	"DESC": ListDetectorRecipesSortOrderDesc,
}

// GetListDetectorRecipesSortOrderEnumValues Enumerates the set of values for ListDetectorRecipesSortOrderEnum
func GetListDetectorRecipesSortOrderEnumValues() []ListDetectorRecipesSortOrderEnum {
	values := make([]ListDetectorRecipesSortOrderEnum, 0)
	for _, v := range mappingListDetectorRecipesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDetectorRecipesSortByEnum Enum with underlying type: string
type ListDetectorRecipesSortByEnum string

// Set of constants representing the allowable values for ListDetectorRecipesSortByEnum
const (
	ListDetectorRecipesSortByTimecreated ListDetectorRecipesSortByEnum = "timeCreated"
	ListDetectorRecipesSortByDisplayname ListDetectorRecipesSortByEnum = "displayName"
)

var mappingListDetectorRecipesSortBy = map[string]ListDetectorRecipesSortByEnum{
	"timeCreated": ListDetectorRecipesSortByTimecreated,
	"displayName": ListDetectorRecipesSortByDisplayname,
}

// GetListDetectorRecipesSortByEnumValues Enumerates the set of values for ListDetectorRecipesSortByEnum
func GetListDetectorRecipesSortByEnumValues() []ListDetectorRecipesSortByEnum {
	values := make([]ListDetectorRecipesSortByEnum, 0)
	for _, v := range mappingListDetectorRecipesSortBy {
		values = append(values, v)
	}
	return values
}
