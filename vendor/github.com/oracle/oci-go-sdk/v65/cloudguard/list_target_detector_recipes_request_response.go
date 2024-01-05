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

// ListTargetDetectorRecipesRequest wrapper for the ListTargetDetectorRecipes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargetDetectorRecipes.go.html to see an example of how to use ListTargetDetectorRecipesRequest.
type ListTargetDetectorRecipesRequest struct {

	// OCID of target
	TargetId *string `mandatory:"true" contributesTo:"path" name:"targetId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListTargetDetectorRecipesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTargetDetectorRecipesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListTargetDetectorRecipesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetDetectorRecipesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetDetectorRecipesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetDetectorRecipesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetDetectorRecipesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetDetectorRecipesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetDetectorRecipesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTargetDetectorRecipesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDetectorRecipesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetDetectorRecipesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDetectorRecipesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetDetectorRecipesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetDetectorRecipesResponse wrapper for the ListTargetDetectorRecipes operation
type ListTargetDetectorRecipesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetDetectorRecipeCollection instances
	TargetDetectorRecipeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTargetDetectorRecipesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetDetectorRecipesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetDetectorRecipesLifecycleStateEnum Enum with underlying type: string
type ListTargetDetectorRecipesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTargetDetectorRecipesLifecycleStateEnum
const (
	ListTargetDetectorRecipesLifecycleStateCreating ListTargetDetectorRecipesLifecycleStateEnum = "CREATING"
	ListTargetDetectorRecipesLifecycleStateUpdating ListTargetDetectorRecipesLifecycleStateEnum = "UPDATING"
	ListTargetDetectorRecipesLifecycleStateActive   ListTargetDetectorRecipesLifecycleStateEnum = "ACTIVE"
	ListTargetDetectorRecipesLifecycleStateInactive ListTargetDetectorRecipesLifecycleStateEnum = "INACTIVE"
	ListTargetDetectorRecipesLifecycleStateDeleting ListTargetDetectorRecipesLifecycleStateEnum = "DELETING"
	ListTargetDetectorRecipesLifecycleStateDeleted  ListTargetDetectorRecipesLifecycleStateEnum = "DELETED"
	ListTargetDetectorRecipesLifecycleStateFailed   ListTargetDetectorRecipesLifecycleStateEnum = "FAILED"
)

var mappingListTargetDetectorRecipesLifecycleStateEnum = map[string]ListTargetDetectorRecipesLifecycleStateEnum{
	"CREATING": ListTargetDetectorRecipesLifecycleStateCreating,
	"UPDATING": ListTargetDetectorRecipesLifecycleStateUpdating,
	"ACTIVE":   ListTargetDetectorRecipesLifecycleStateActive,
	"INACTIVE": ListTargetDetectorRecipesLifecycleStateInactive,
	"DELETING": ListTargetDetectorRecipesLifecycleStateDeleting,
	"DELETED":  ListTargetDetectorRecipesLifecycleStateDeleted,
	"FAILED":   ListTargetDetectorRecipesLifecycleStateFailed,
}

var mappingListTargetDetectorRecipesLifecycleStateEnumLowerCase = map[string]ListTargetDetectorRecipesLifecycleStateEnum{
	"creating": ListTargetDetectorRecipesLifecycleStateCreating,
	"updating": ListTargetDetectorRecipesLifecycleStateUpdating,
	"active":   ListTargetDetectorRecipesLifecycleStateActive,
	"inactive": ListTargetDetectorRecipesLifecycleStateInactive,
	"deleting": ListTargetDetectorRecipesLifecycleStateDeleting,
	"deleted":  ListTargetDetectorRecipesLifecycleStateDeleted,
	"failed":   ListTargetDetectorRecipesLifecycleStateFailed,
}

// GetListTargetDetectorRecipesLifecycleStateEnumValues Enumerates the set of values for ListTargetDetectorRecipesLifecycleStateEnum
func GetListTargetDetectorRecipesLifecycleStateEnumValues() []ListTargetDetectorRecipesLifecycleStateEnum {
	values := make([]ListTargetDetectorRecipesLifecycleStateEnum, 0)
	for _, v := range mappingListTargetDetectorRecipesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDetectorRecipesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTargetDetectorRecipesLifecycleStateEnum
func GetListTargetDetectorRecipesLifecycleStateEnumStringValues() []string {
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

// GetMappingListTargetDetectorRecipesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDetectorRecipesLifecycleStateEnum(val string) (ListTargetDetectorRecipesLifecycleStateEnum, bool) {
	enum, ok := mappingListTargetDetectorRecipesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDetectorRecipesSortOrderEnum Enum with underlying type: string
type ListTargetDetectorRecipesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetDetectorRecipesSortOrderEnum
const (
	ListTargetDetectorRecipesSortOrderAsc  ListTargetDetectorRecipesSortOrderEnum = "ASC"
	ListTargetDetectorRecipesSortOrderDesc ListTargetDetectorRecipesSortOrderEnum = "DESC"
)

var mappingListTargetDetectorRecipesSortOrderEnum = map[string]ListTargetDetectorRecipesSortOrderEnum{
	"ASC":  ListTargetDetectorRecipesSortOrderAsc,
	"DESC": ListTargetDetectorRecipesSortOrderDesc,
}

var mappingListTargetDetectorRecipesSortOrderEnumLowerCase = map[string]ListTargetDetectorRecipesSortOrderEnum{
	"asc":  ListTargetDetectorRecipesSortOrderAsc,
	"desc": ListTargetDetectorRecipesSortOrderDesc,
}

// GetListTargetDetectorRecipesSortOrderEnumValues Enumerates the set of values for ListTargetDetectorRecipesSortOrderEnum
func GetListTargetDetectorRecipesSortOrderEnumValues() []ListTargetDetectorRecipesSortOrderEnum {
	values := make([]ListTargetDetectorRecipesSortOrderEnum, 0)
	for _, v := range mappingListTargetDetectorRecipesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDetectorRecipesSortOrderEnumStringValues Enumerates the set of values in String for ListTargetDetectorRecipesSortOrderEnum
func GetListTargetDetectorRecipesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetDetectorRecipesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDetectorRecipesSortOrderEnum(val string) (ListTargetDetectorRecipesSortOrderEnum, bool) {
	enum, ok := mappingListTargetDetectorRecipesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDetectorRecipesSortByEnum Enum with underlying type: string
type ListTargetDetectorRecipesSortByEnum string

// Set of constants representing the allowable values for ListTargetDetectorRecipesSortByEnum
const (
	ListTargetDetectorRecipesSortByTimecreated ListTargetDetectorRecipesSortByEnum = "timeCreated"
	ListTargetDetectorRecipesSortByDisplayname ListTargetDetectorRecipesSortByEnum = "displayName"
)

var mappingListTargetDetectorRecipesSortByEnum = map[string]ListTargetDetectorRecipesSortByEnum{
	"timeCreated": ListTargetDetectorRecipesSortByTimecreated,
	"displayName": ListTargetDetectorRecipesSortByDisplayname,
}

var mappingListTargetDetectorRecipesSortByEnumLowerCase = map[string]ListTargetDetectorRecipesSortByEnum{
	"timecreated": ListTargetDetectorRecipesSortByTimecreated,
	"displayname": ListTargetDetectorRecipesSortByDisplayname,
}

// GetListTargetDetectorRecipesSortByEnumValues Enumerates the set of values for ListTargetDetectorRecipesSortByEnum
func GetListTargetDetectorRecipesSortByEnumValues() []ListTargetDetectorRecipesSortByEnum {
	values := make([]ListTargetDetectorRecipesSortByEnum, 0)
	for _, v := range mappingListTargetDetectorRecipesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDetectorRecipesSortByEnumStringValues Enumerates the set of values in String for ListTargetDetectorRecipesSortByEnum
func GetListTargetDetectorRecipesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTargetDetectorRecipesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDetectorRecipesSortByEnum(val string) (ListTargetDetectorRecipesSortByEnum, bool) {
	enum, ok := mappingListTargetDetectorRecipesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
