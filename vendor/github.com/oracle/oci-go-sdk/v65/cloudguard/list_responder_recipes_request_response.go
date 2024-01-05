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

// ListResponderRecipesRequest wrapper for the ListResponderRecipes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResponderRecipes.go.html to see an example of how to use ListResponderRecipesRequest.
type ListResponderRecipesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the list of all Oracle Managed Resources
	// Metadata supported by Cloud Guard are returned.
	ResourceMetadataOnly *bool `mandatory:"false" contributesTo:"query" name:"resourceMetadataOnly"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListResponderRecipesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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
	AccessLevel ListResponderRecipesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResponderRecipesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListResponderRecipesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResponderRecipesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResponderRecipesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResponderRecipesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResponderRecipesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResponderRecipesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResponderRecipesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListResponderRecipesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderRecipesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListResponderRecipesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderRecipesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResponderRecipesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderRecipesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResponderRecipesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResponderRecipesResponse wrapper for the ListResponderRecipes operation
type ListResponderRecipesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResponderRecipeCollection instances
	ResponderRecipeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResponderRecipesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResponderRecipesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResponderRecipesLifecycleStateEnum Enum with underlying type: string
type ListResponderRecipesLifecycleStateEnum string

// Set of constants representing the allowable values for ListResponderRecipesLifecycleStateEnum
const (
	ListResponderRecipesLifecycleStateCreating ListResponderRecipesLifecycleStateEnum = "CREATING"
	ListResponderRecipesLifecycleStateUpdating ListResponderRecipesLifecycleStateEnum = "UPDATING"
	ListResponderRecipesLifecycleStateActive   ListResponderRecipesLifecycleStateEnum = "ACTIVE"
	ListResponderRecipesLifecycleStateInactive ListResponderRecipesLifecycleStateEnum = "INACTIVE"
	ListResponderRecipesLifecycleStateDeleting ListResponderRecipesLifecycleStateEnum = "DELETING"
	ListResponderRecipesLifecycleStateDeleted  ListResponderRecipesLifecycleStateEnum = "DELETED"
	ListResponderRecipesLifecycleStateFailed   ListResponderRecipesLifecycleStateEnum = "FAILED"
)

var mappingListResponderRecipesLifecycleStateEnum = map[string]ListResponderRecipesLifecycleStateEnum{
	"CREATING": ListResponderRecipesLifecycleStateCreating,
	"UPDATING": ListResponderRecipesLifecycleStateUpdating,
	"ACTIVE":   ListResponderRecipesLifecycleStateActive,
	"INACTIVE": ListResponderRecipesLifecycleStateInactive,
	"DELETING": ListResponderRecipesLifecycleStateDeleting,
	"DELETED":  ListResponderRecipesLifecycleStateDeleted,
	"FAILED":   ListResponderRecipesLifecycleStateFailed,
}

var mappingListResponderRecipesLifecycleStateEnumLowerCase = map[string]ListResponderRecipesLifecycleStateEnum{
	"creating": ListResponderRecipesLifecycleStateCreating,
	"updating": ListResponderRecipesLifecycleStateUpdating,
	"active":   ListResponderRecipesLifecycleStateActive,
	"inactive": ListResponderRecipesLifecycleStateInactive,
	"deleting": ListResponderRecipesLifecycleStateDeleting,
	"deleted":  ListResponderRecipesLifecycleStateDeleted,
	"failed":   ListResponderRecipesLifecycleStateFailed,
}

// GetListResponderRecipesLifecycleStateEnumValues Enumerates the set of values for ListResponderRecipesLifecycleStateEnum
func GetListResponderRecipesLifecycleStateEnumValues() []ListResponderRecipesLifecycleStateEnum {
	values := make([]ListResponderRecipesLifecycleStateEnum, 0)
	for _, v := range mappingListResponderRecipesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderRecipesLifecycleStateEnumStringValues Enumerates the set of values in String for ListResponderRecipesLifecycleStateEnum
func GetListResponderRecipesLifecycleStateEnumStringValues() []string {
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

// GetMappingListResponderRecipesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderRecipesLifecycleStateEnum(val string) (ListResponderRecipesLifecycleStateEnum, bool) {
	enum, ok := mappingListResponderRecipesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderRecipesAccessLevelEnum Enum with underlying type: string
type ListResponderRecipesAccessLevelEnum string

// Set of constants representing the allowable values for ListResponderRecipesAccessLevelEnum
const (
	ListResponderRecipesAccessLevelRestricted ListResponderRecipesAccessLevelEnum = "RESTRICTED"
	ListResponderRecipesAccessLevelAccessible ListResponderRecipesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListResponderRecipesAccessLevelEnum = map[string]ListResponderRecipesAccessLevelEnum{
	"RESTRICTED": ListResponderRecipesAccessLevelRestricted,
	"ACCESSIBLE": ListResponderRecipesAccessLevelAccessible,
}

var mappingListResponderRecipesAccessLevelEnumLowerCase = map[string]ListResponderRecipesAccessLevelEnum{
	"restricted": ListResponderRecipesAccessLevelRestricted,
	"accessible": ListResponderRecipesAccessLevelAccessible,
}

// GetListResponderRecipesAccessLevelEnumValues Enumerates the set of values for ListResponderRecipesAccessLevelEnum
func GetListResponderRecipesAccessLevelEnumValues() []ListResponderRecipesAccessLevelEnum {
	values := make([]ListResponderRecipesAccessLevelEnum, 0)
	for _, v := range mappingListResponderRecipesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderRecipesAccessLevelEnumStringValues Enumerates the set of values in String for ListResponderRecipesAccessLevelEnum
func GetListResponderRecipesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListResponderRecipesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderRecipesAccessLevelEnum(val string) (ListResponderRecipesAccessLevelEnum, bool) {
	enum, ok := mappingListResponderRecipesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderRecipesSortOrderEnum Enum with underlying type: string
type ListResponderRecipesSortOrderEnum string

// Set of constants representing the allowable values for ListResponderRecipesSortOrderEnum
const (
	ListResponderRecipesSortOrderAsc  ListResponderRecipesSortOrderEnum = "ASC"
	ListResponderRecipesSortOrderDesc ListResponderRecipesSortOrderEnum = "DESC"
)

var mappingListResponderRecipesSortOrderEnum = map[string]ListResponderRecipesSortOrderEnum{
	"ASC":  ListResponderRecipesSortOrderAsc,
	"DESC": ListResponderRecipesSortOrderDesc,
}

var mappingListResponderRecipesSortOrderEnumLowerCase = map[string]ListResponderRecipesSortOrderEnum{
	"asc":  ListResponderRecipesSortOrderAsc,
	"desc": ListResponderRecipesSortOrderDesc,
}

// GetListResponderRecipesSortOrderEnumValues Enumerates the set of values for ListResponderRecipesSortOrderEnum
func GetListResponderRecipesSortOrderEnumValues() []ListResponderRecipesSortOrderEnum {
	values := make([]ListResponderRecipesSortOrderEnum, 0)
	for _, v := range mappingListResponderRecipesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderRecipesSortOrderEnumStringValues Enumerates the set of values in String for ListResponderRecipesSortOrderEnum
func GetListResponderRecipesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResponderRecipesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderRecipesSortOrderEnum(val string) (ListResponderRecipesSortOrderEnum, bool) {
	enum, ok := mappingListResponderRecipesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderRecipesSortByEnum Enum with underlying type: string
type ListResponderRecipesSortByEnum string

// Set of constants representing the allowable values for ListResponderRecipesSortByEnum
const (
	ListResponderRecipesSortByTimecreated ListResponderRecipesSortByEnum = "timeCreated"
	ListResponderRecipesSortByDisplayname ListResponderRecipesSortByEnum = "displayName"
)

var mappingListResponderRecipesSortByEnum = map[string]ListResponderRecipesSortByEnum{
	"timeCreated": ListResponderRecipesSortByTimecreated,
	"displayName": ListResponderRecipesSortByDisplayname,
}

var mappingListResponderRecipesSortByEnumLowerCase = map[string]ListResponderRecipesSortByEnum{
	"timecreated": ListResponderRecipesSortByTimecreated,
	"displayname": ListResponderRecipesSortByDisplayname,
}

// GetListResponderRecipesSortByEnumValues Enumerates the set of values for ListResponderRecipesSortByEnum
func GetListResponderRecipesSortByEnumValues() []ListResponderRecipesSortByEnum {
	values := make([]ListResponderRecipesSortByEnum, 0)
	for _, v := range mappingListResponderRecipesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderRecipesSortByEnumStringValues Enumerates the set of values in String for ListResponderRecipesSortByEnum
func GetListResponderRecipesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListResponderRecipesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderRecipesSortByEnum(val string) (ListResponderRecipesSortByEnum, bool) {
	enum, ok := mappingListResponderRecipesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
