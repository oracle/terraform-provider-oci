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

// ListSecurityZonesRequest wrapper for the ListSecurityZones operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListSecurityZones.go.html to see an example of how to use ListSecurityZonesRequest.
type ListSecurityZonesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field lifecycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListSecurityZonesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The unique identifier of the security zone (`SecurityZone` resource).
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The unique identifier of the security zone recipe. (`SecurityRecipe` resource).
	SecurityRecipeId *string `mandatory:"false" contributesTo:"query" name:"securityRecipeId"`

	// Is security zones in the subtree?
	IsRequiredSecurityZonesInSubtree *bool `mandatory:"false" contributesTo:"query" name:"isRequiredSecurityZonesInSubtree"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use
	SortOrder ListSecurityZonesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListSecurityZonesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityZonesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityZonesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityZonesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityZonesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityZonesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityZonesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSecurityZonesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityZonesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecurityZonesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityZonesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecurityZonesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityZonesResponse wrapper for the ListSecurityZones operation
type ListSecurityZonesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SecurityZoneCollection instances
	SecurityZoneCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSecurityZonesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityZonesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityZonesLifecycleStateEnum Enum with underlying type: string
type ListSecurityZonesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSecurityZonesLifecycleStateEnum
const (
	ListSecurityZonesLifecycleStateCreating ListSecurityZonesLifecycleStateEnum = "CREATING"
	ListSecurityZonesLifecycleStateUpdating ListSecurityZonesLifecycleStateEnum = "UPDATING"
	ListSecurityZonesLifecycleStateActive   ListSecurityZonesLifecycleStateEnum = "ACTIVE"
	ListSecurityZonesLifecycleStateInactive ListSecurityZonesLifecycleStateEnum = "INACTIVE"
	ListSecurityZonesLifecycleStateDeleting ListSecurityZonesLifecycleStateEnum = "DELETING"
	ListSecurityZonesLifecycleStateDeleted  ListSecurityZonesLifecycleStateEnum = "DELETED"
	ListSecurityZonesLifecycleStateFailed   ListSecurityZonesLifecycleStateEnum = "FAILED"
)

var mappingListSecurityZonesLifecycleStateEnum = map[string]ListSecurityZonesLifecycleStateEnum{
	"CREATING": ListSecurityZonesLifecycleStateCreating,
	"UPDATING": ListSecurityZonesLifecycleStateUpdating,
	"ACTIVE":   ListSecurityZonesLifecycleStateActive,
	"INACTIVE": ListSecurityZonesLifecycleStateInactive,
	"DELETING": ListSecurityZonesLifecycleStateDeleting,
	"DELETED":  ListSecurityZonesLifecycleStateDeleted,
	"FAILED":   ListSecurityZonesLifecycleStateFailed,
}

var mappingListSecurityZonesLifecycleStateEnumLowerCase = map[string]ListSecurityZonesLifecycleStateEnum{
	"creating": ListSecurityZonesLifecycleStateCreating,
	"updating": ListSecurityZonesLifecycleStateUpdating,
	"active":   ListSecurityZonesLifecycleStateActive,
	"inactive": ListSecurityZonesLifecycleStateInactive,
	"deleting": ListSecurityZonesLifecycleStateDeleting,
	"deleted":  ListSecurityZonesLifecycleStateDeleted,
	"failed":   ListSecurityZonesLifecycleStateFailed,
}

// GetListSecurityZonesLifecycleStateEnumValues Enumerates the set of values for ListSecurityZonesLifecycleStateEnum
func GetListSecurityZonesLifecycleStateEnumValues() []ListSecurityZonesLifecycleStateEnum {
	values := make([]ListSecurityZonesLifecycleStateEnum, 0)
	for _, v := range mappingListSecurityZonesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityZonesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSecurityZonesLifecycleStateEnum
func GetListSecurityZonesLifecycleStateEnumStringValues() []string {
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

// GetMappingListSecurityZonesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityZonesLifecycleStateEnum(val string) (ListSecurityZonesLifecycleStateEnum, bool) {
	enum, ok := mappingListSecurityZonesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityZonesSortOrderEnum Enum with underlying type: string
type ListSecurityZonesSortOrderEnum string

// Set of constants representing the allowable values for ListSecurityZonesSortOrderEnum
const (
	ListSecurityZonesSortOrderAsc  ListSecurityZonesSortOrderEnum = "ASC"
	ListSecurityZonesSortOrderDesc ListSecurityZonesSortOrderEnum = "DESC"
)

var mappingListSecurityZonesSortOrderEnum = map[string]ListSecurityZonesSortOrderEnum{
	"ASC":  ListSecurityZonesSortOrderAsc,
	"DESC": ListSecurityZonesSortOrderDesc,
}

var mappingListSecurityZonesSortOrderEnumLowerCase = map[string]ListSecurityZonesSortOrderEnum{
	"asc":  ListSecurityZonesSortOrderAsc,
	"desc": ListSecurityZonesSortOrderDesc,
}

// GetListSecurityZonesSortOrderEnumValues Enumerates the set of values for ListSecurityZonesSortOrderEnum
func GetListSecurityZonesSortOrderEnumValues() []ListSecurityZonesSortOrderEnum {
	values := make([]ListSecurityZonesSortOrderEnum, 0)
	for _, v := range mappingListSecurityZonesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityZonesSortOrderEnumStringValues Enumerates the set of values in String for ListSecurityZonesSortOrderEnum
func GetListSecurityZonesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSecurityZonesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityZonesSortOrderEnum(val string) (ListSecurityZonesSortOrderEnum, bool) {
	enum, ok := mappingListSecurityZonesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityZonesSortByEnum Enum with underlying type: string
type ListSecurityZonesSortByEnum string

// Set of constants representing the allowable values for ListSecurityZonesSortByEnum
const (
	ListSecurityZonesSortByTimecreated ListSecurityZonesSortByEnum = "timeCreated"
	ListSecurityZonesSortByDisplayname ListSecurityZonesSortByEnum = "displayName"
)

var mappingListSecurityZonesSortByEnum = map[string]ListSecurityZonesSortByEnum{
	"timeCreated": ListSecurityZonesSortByTimecreated,
	"displayName": ListSecurityZonesSortByDisplayname,
}

var mappingListSecurityZonesSortByEnumLowerCase = map[string]ListSecurityZonesSortByEnum{
	"timecreated": ListSecurityZonesSortByTimecreated,
	"displayname": ListSecurityZonesSortByDisplayname,
}

// GetListSecurityZonesSortByEnumValues Enumerates the set of values for ListSecurityZonesSortByEnum
func GetListSecurityZonesSortByEnumValues() []ListSecurityZonesSortByEnum {
	values := make([]ListSecurityZonesSortByEnum, 0)
	for _, v := range mappingListSecurityZonesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityZonesSortByEnumStringValues Enumerates the set of values in String for ListSecurityZonesSortByEnum
func GetListSecurityZonesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSecurityZonesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityZonesSortByEnum(val string) (ListSecurityZonesSortByEnum, bool) {
	enum, ok := mappingListSecurityZonesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
