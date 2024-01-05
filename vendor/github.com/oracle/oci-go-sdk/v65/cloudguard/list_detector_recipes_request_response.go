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

// ListDetectorRecipesRequest wrapper for the ListDetectorRecipes operation
//
// # See also
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDetectorRecipesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDetectorRecipesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDetectorRecipesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectorRecipesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListDetectorRecipesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectorRecipesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDetectorRecipesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDetectorRecipesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDetectorRecipesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListDetectorRecipesLifecycleStateEnum = map[string]ListDetectorRecipesLifecycleStateEnum{
	"CREATING": ListDetectorRecipesLifecycleStateCreating,
	"UPDATING": ListDetectorRecipesLifecycleStateUpdating,
	"ACTIVE":   ListDetectorRecipesLifecycleStateActive,
	"INACTIVE": ListDetectorRecipesLifecycleStateInactive,
	"DELETING": ListDetectorRecipesLifecycleStateDeleting,
	"DELETED":  ListDetectorRecipesLifecycleStateDeleted,
	"FAILED":   ListDetectorRecipesLifecycleStateFailed,
}

var mappingListDetectorRecipesLifecycleStateEnumLowerCase = map[string]ListDetectorRecipesLifecycleStateEnum{
	"creating": ListDetectorRecipesLifecycleStateCreating,
	"updating": ListDetectorRecipesLifecycleStateUpdating,
	"active":   ListDetectorRecipesLifecycleStateActive,
	"inactive": ListDetectorRecipesLifecycleStateInactive,
	"deleting": ListDetectorRecipesLifecycleStateDeleting,
	"deleted":  ListDetectorRecipesLifecycleStateDeleted,
	"failed":   ListDetectorRecipesLifecycleStateFailed,
}

// GetListDetectorRecipesLifecycleStateEnumValues Enumerates the set of values for ListDetectorRecipesLifecycleStateEnum
func GetListDetectorRecipesLifecycleStateEnumValues() []ListDetectorRecipesLifecycleStateEnum {
	values := make([]ListDetectorRecipesLifecycleStateEnum, 0)
	for _, v := range mappingListDetectorRecipesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorRecipesLifecycleStateEnumStringValues Enumerates the set of values in String for ListDetectorRecipesLifecycleStateEnum
func GetListDetectorRecipesLifecycleStateEnumStringValues() []string {
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

// GetMappingListDetectorRecipesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorRecipesLifecycleStateEnum(val string) (ListDetectorRecipesLifecycleStateEnum, bool) {
	enum, ok := mappingListDetectorRecipesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDetectorRecipesAccessLevelEnum Enum with underlying type: string
type ListDetectorRecipesAccessLevelEnum string

// Set of constants representing the allowable values for ListDetectorRecipesAccessLevelEnum
const (
	ListDetectorRecipesAccessLevelRestricted ListDetectorRecipesAccessLevelEnum = "RESTRICTED"
	ListDetectorRecipesAccessLevelAccessible ListDetectorRecipesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListDetectorRecipesAccessLevelEnum = map[string]ListDetectorRecipesAccessLevelEnum{
	"RESTRICTED": ListDetectorRecipesAccessLevelRestricted,
	"ACCESSIBLE": ListDetectorRecipesAccessLevelAccessible,
}

var mappingListDetectorRecipesAccessLevelEnumLowerCase = map[string]ListDetectorRecipesAccessLevelEnum{
	"restricted": ListDetectorRecipesAccessLevelRestricted,
	"accessible": ListDetectorRecipesAccessLevelAccessible,
}

// GetListDetectorRecipesAccessLevelEnumValues Enumerates the set of values for ListDetectorRecipesAccessLevelEnum
func GetListDetectorRecipesAccessLevelEnumValues() []ListDetectorRecipesAccessLevelEnum {
	values := make([]ListDetectorRecipesAccessLevelEnum, 0)
	for _, v := range mappingListDetectorRecipesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorRecipesAccessLevelEnumStringValues Enumerates the set of values in String for ListDetectorRecipesAccessLevelEnum
func GetListDetectorRecipesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListDetectorRecipesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorRecipesAccessLevelEnum(val string) (ListDetectorRecipesAccessLevelEnum, bool) {
	enum, ok := mappingListDetectorRecipesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDetectorRecipesSortOrderEnum Enum with underlying type: string
type ListDetectorRecipesSortOrderEnum string

// Set of constants representing the allowable values for ListDetectorRecipesSortOrderEnum
const (
	ListDetectorRecipesSortOrderAsc  ListDetectorRecipesSortOrderEnum = "ASC"
	ListDetectorRecipesSortOrderDesc ListDetectorRecipesSortOrderEnum = "DESC"
)

var mappingListDetectorRecipesSortOrderEnum = map[string]ListDetectorRecipesSortOrderEnum{
	"ASC":  ListDetectorRecipesSortOrderAsc,
	"DESC": ListDetectorRecipesSortOrderDesc,
}

var mappingListDetectorRecipesSortOrderEnumLowerCase = map[string]ListDetectorRecipesSortOrderEnum{
	"asc":  ListDetectorRecipesSortOrderAsc,
	"desc": ListDetectorRecipesSortOrderDesc,
}

// GetListDetectorRecipesSortOrderEnumValues Enumerates the set of values for ListDetectorRecipesSortOrderEnum
func GetListDetectorRecipesSortOrderEnumValues() []ListDetectorRecipesSortOrderEnum {
	values := make([]ListDetectorRecipesSortOrderEnum, 0)
	for _, v := range mappingListDetectorRecipesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorRecipesSortOrderEnumStringValues Enumerates the set of values in String for ListDetectorRecipesSortOrderEnum
func GetListDetectorRecipesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDetectorRecipesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorRecipesSortOrderEnum(val string) (ListDetectorRecipesSortOrderEnum, bool) {
	enum, ok := mappingListDetectorRecipesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDetectorRecipesSortByEnum Enum with underlying type: string
type ListDetectorRecipesSortByEnum string

// Set of constants representing the allowable values for ListDetectorRecipesSortByEnum
const (
	ListDetectorRecipesSortByTimecreated ListDetectorRecipesSortByEnum = "timeCreated"
	ListDetectorRecipesSortByDisplayname ListDetectorRecipesSortByEnum = "displayName"
)

var mappingListDetectorRecipesSortByEnum = map[string]ListDetectorRecipesSortByEnum{
	"timeCreated": ListDetectorRecipesSortByTimecreated,
	"displayName": ListDetectorRecipesSortByDisplayname,
}

var mappingListDetectorRecipesSortByEnumLowerCase = map[string]ListDetectorRecipesSortByEnum{
	"timecreated": ListDetectorRecipesSortByTimecreated,
	"displayname": ListDetectorRecipesSortByDisplayname,
}

// GetListDetectorRecipesSortByEnumValues Enumerates the set of values for ListDetectorRecipesSortByEnum
func GetListDetectorRecipesSortByEnumValues() []ListDetectorRecipesSortByEnum {
	values := make([]ListDetectorRecipesSortByEnum, 0)
	for _, v := range mappingListDetectorRecipesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDetectorRecipesSortByEnumStringValues Enumerates the set of values in String for ListDetectorRecipesSortByEnum
func GetListDetectorRecipesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDetectorRecipesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDetectorRecipesSortByEnum(val string) (ListDetectorRecipesSortByEnum, bool) {
	enum, ok := mappingListDetectorRecipesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
