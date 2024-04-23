// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSecurityRecipesRequest wrapper for the ListSecurityRecipes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSecurityRecipes.go.html to see an example of how to use ListSecurityRecipesRequest.
type ListSecurityRecipesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListSecurityRecipesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The unique identifier of the security zone recipe. (`SecurityRecipe`)
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListSecurityRecipesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListSecurityRecipesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityRecipesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityRecipesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityRecipesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityRecipesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityRecipesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityRecipesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSecurityRecipesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityRecipesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecurityRecipesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityRecipesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecurityRecipesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityRecipesResponse wrapper for the ListSecurityRecipes operation
type ListSecurityRecipesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SecurityRecipeCollection instances
	SecurityRecipeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSecurityRecipesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityRecipesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityRecipesLifecycleStateEnum Enum with underlying type: string
type ListSecurityRecipesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSecurityRecipesLifecycleStateEnum
const (
	ListSecurityRecipesLifecycleStateCreating ListSecurityRecipesLifecycleStateEnum = "CREATING"
	ListSecurityRecipesLifecycleStateUpdating ListSecurityRecipesLifecycleStateEnum = "UPDATING"
	ListSecurityRecipesLifecycleStateActive   ListSecurityRecipesLifecycleStateEnum = "ACTIVE"
	ListSecurityRecipesLifecycleStateInactive ListSecurityRecipesLifecycleStateEnum = "INACTIVE"
	ListSecurityRecipesLifecycleStateDeleting ListSecurityRecipesLifecycleStateEnum = "DELETING"
	ListSecurityRecipesLifecycleStateDeleted  ListSecurityRecipesLifecycleStateEnum = "DELETED"
	ListSecurityRecipesLifecycleStateFailed   ListSecurityRecipesLifecycleStateEnum = "FAILED"
)

var mappingListSecurityRecipesLifecycleStateEnum = map[string]ListSecurityRecipesLifecycleStateEnum{
	"CREATING": ListSecurityRecipesLifecycleStateCreating,
	"UPDATING": ListSecurityRecipesLifecycleStateUpdating,
	"ACTIVE":   ListSecurityRecipesLifecycleStateActive,
	"INACTIVE": ListSecurityRecipesLifecycleStateInactive,
	"DELETING": ListSecurityRecipesLifecycleStateDeleting,
	"DELETED":  ListSecurityRecipesLifecycleStateDeleted,
	"FAILED":   ListSecurityRecipesLifecycleStateFailed,
}

var mappingListSecurityRecipesLifecycleStateEnumLowerCase = map[string]ListSecurityRecipesLifecycleStateEnum{
	"creating": ListSecurityRecipesLifecycleStateCreating,
	"updating": ListSecurityRecipesLifecycleStateUpdating,
	"active":   ListSecurityRecipesLifecycleStateActive,
	"inactive": ListSecurityRecipesLifecycleStateInactive,
	"deleting": ListSecurityRecipesLifecycleStateDeleting,
	"deleted":  ListSecurityRecipesLifecycleStateDeleted,
	"failed":   ListSecurityRecipesLifecycleStateFailed,
}

// GetListSecurityRecipesLifecycleStateEnumValues Enumerates the set of values for ListSecurityRecipesLifecycleStateEnum
func GetListSecurityRecipesLifecycleStateEnumValues() []ListSecurityRecipesLifecycleStateEnum {
	values := make([]ListSecurityRecipesLifecycleStateEnum, 0)
	for _, v := range mappingListSecurityRecipesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityRecipesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSecurityRecipesLifecycleStateEnum
func GetListSecurityRecipesLifecycleStateEnumStringValues() []string {
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

// GetMappingListSecurityRecipesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityRecipesLifecycleStateEnum(val string) (ListSecurityRecipesLifecycleStateEnum, bool) {
	enum, ok := mappingListSecurityRecipesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityRecipesSortOrderEnum Enum with underlying type: string
type ListSecurityRecipesSortOrderEnum string

// Set of constants representing the allowable values for ListSecurityRecipesSortOrderEnum
const (
	ListSecurityRecipesSortOrderAsc  ListSecurityRecipesSortOrderEnum = "ASC"
	ListSecurityRecipesSortOrderDesc ListSecurityRecipesSortOrderEnum = "DESC"
)

var mappingListSecurityRecipesSortOrderEnum = map[string]ListSecurityRecipesSortOrderEnum{
	"ASC":  ListSecurityRecipesSortOrderAsc,
	"DESC": ListSecurityRecipesSortOrderDesc,
}

var mappingListSecurityRecipesSortOrderEnumLowerCase = map[string]ListSecurityRecipesSortOrderEnum{
	"asc":  ListSecurityRecipesSortOrderAsc,
	"desc": ListSecurityRecipesSortOrderDesc,
}

// GetListSecurityRecipesSortOrderEnumValues Enumerates the set of values for ListSecurityRecipesSortOrderEnum
func GetListSecurityRecipesSortOrderEnumValues() []ListSecurityRecipesSortOrderEnum {
	values := make([]ListSecurityRecipesSortOrderEnum, 0)
	for _, v := range mappingListSecurityRecipesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityRecipesSortOrderEnumStringValues Enumerates the set of values in String for ListSecurityRecipesSortOrderEnum
func GetListSecurityRecipesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSecurityRecipesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityRecipesSortOrderEnum(val string) (ListSecurityRecipesSortOrderEnum, bool) {
	enum, ok := mappingListSecurityRecipesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityRecipesSortByEnum Enum with underlying type: string
type ListSecurityRecipesSortByEnum string

// Set of constants representing the allowable values for ListSecurityRecipesSortByEnum
const (
	ListSecurityRecipesSortByTimecreated ListSecurityRecipesSortByEnum = "timeCreated"
	ListSecurityRecipesSortByDisplayname ListSecurityRecipesSortByEnum = "displayName"
)

var mappingListSecurityRecipesSortByEnum = map[string]ListSecurityRecipesSortByEnum{
	"timeCreated": ListSecurityRecipesSortByTimecreated,
	"displayName": ListSecurityRecipesSortByDisplayname,
}

var mappingListSecurityRecipesSortByEnumLowerCase = map[string]ListSecurityRecipesSortByEnum{
	"timecreated": ListSecurityRecipesSortByTimecreated,
	"displayname": ListSecurityRecipesSortByDisplayname,
}

// GetListSecurityRecipesSortByEnumValues Enumerates the set of values for ListSecurityRecipesSortByEnum
func GetListSecurityRecipesSortByEnumValues() []ListSecurityRecipesSortByEnum {
	values := make([]ListSecurityRecipesSortByEnum, 0)
	for _, v := range mappingListSecurityRecipesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityRecipesSortByEnumStringValues Enumerates the set of values in String for ListSecurityRecipesSortByEnum
func GetListSecurityRecipesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSecurityRecipesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityRecipesSortByEnum(val string) (ListSecurityRecipesSortByEnum, bool) {
	enum, ok := mappingListSecurityRecipesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
