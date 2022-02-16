// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetResponderRecipesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetResponderRecipesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTargetResponderRecipesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetResponderRecipesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetResponderRecipesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetResponderRecipesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetResponderRecipesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListTargetResponderRecipesLifecycleStateEnum = map[string]ListTargetResponderRecipesLifecycleStateEnum{
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
	for _, v := range mappingListTargetResponderRecipesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetResponderRecipesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTargetResponderRecipesLifecycleStateEnum
func GetListTargetResponderRecipesLifecycleStateEnumStringValues() []string {
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

// GetMappingListTargetResponderRecipesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetResponderRecipesLifecycleStateEnum(val string) (ListTargetResponderRecipesLifecycleStateEnum, bool) {
	mappingListTargetResponderRecipesLifecycleStateEnumIgnoreCase := make(map[string]ListTargetResponderRecipesLifecycleStateEnum)
	for k, v := range mappingListTargetResponderRecipesLifecycleStateEnum {
		mappingListTargetResponderRecipesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTargetResponderRecipesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetResponderRecipesSortOrderEnum Enum with underlying type: string
type ListTargetResponderRecipesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetResponderRecipesSortOrderEnum
const (
	ListTargetResponderRecipesSortOrderAsc  ListTargetResponderRecipesSortOrderEnum = "ASC"
	ListTargetResponderRecipesSortOrderDesc ListTargetResponderRecipesSortOrderEnum = "DESC"
)

var mappingListTargetResponderRecipesSortOrderEnum = map[string]ListTargetResponderRecipesSortOrderEnum{
	"ASC":  ListTargetResponderRecipesSortOrderAsc,
	"DESC": ListTargetResponderRecipesSortOrderDesc,
}

// GetListTargetResponderRecipesSortOrderEnumValues Enumerates the set of values for ListTargetResponderRecipesSortOrderEnum
func GetListTargetResponderRecipesSortOrderEnumValues() []ListTargetResponderRecipesSortOrderEnum {
	values := make([]ListTargetResponderRecipesSortOrderEnum, 0)
	for _, v := range mappingListTargetResponderRecipesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetResponderRecipesSortOrderEnumStringValues Enumerates the set of values in String for ListTargetResponderRecipesSortOrderEnum
func GetListTargetResponderRecipesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetResponderRecipesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetResponderRecipesSortOrderEnum(val string) (ListTargetResponderRecipesSortOrderEnum, bool) {
	mappingListTargetResponderRecipesSortOrderEnumIgnoreCase := make(map[string]ListTargetResponderRecipesSortOrderEnum)
	for k, v := range mappingListTargetResponderRecipesSortOrderEnum {
		mappingListTargetResponderRecipesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTargetResponderRecipesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetResponderRecipesSortByEnum Enum with underlying type: string
type ListTargetResponderRecipesSortByEnum string

// Set of constants representing the allowable values for ListTargetResponderRecipesSortByEnum
const (
	ListTargetResponderRecipesSortByTimecreated ListTargetResponderRecipesSortByEnum = "timeCreated"
	ListTargetResponderRecipesSortByDisplayname ListTargetResponderRecipesSortByEnum = "displayName"
)

var mappingListTargetResponderRecipesSortByEnum = map[string]ListTargetResponderRecipesSortByEnum{
	"timeCreated": ListTargetResponderRecipesSortByTimecreated,
	"displayName": ListTargetResponderRecipesSortByDisplayname,
}

// GetListTargetResponderRecipesSortByEnumValues Enumerates the set of values for ListTargetResponderRecipesSortByEnum
func GetListTargetResponderRecipesSortByEnumValues() []ListTargetResponderRecipesSortByEnum {
	values := make([]ListTargetResponderRecipesSortByEnum, 0)
	for _, v := range mappingListTargetResponderRecipesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetResponderRecipesSortByEnumStringValues Enumerates the set of values in String for ListTargetResponderRecipesSortByEnum
func GetListTargetResponderRecipesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTargetResponderRecipesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetResponderRecipesSortByEnum(val string) (ListTargetResponderRecipesSortByEnum, bool) {
	mappingListTargetResponderRecipesSortByEnumIgnoreCase := make(map[string]ListTargetResponderRecipesSortByEnum)
	for k, v := range mappingListTargetResponderRecipesSortByEnum {
		mappingListTargetResponderRecipesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTargetResponderRecipesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
