// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSensitiveTypeGroupsRequest wrapper for the ListSensitiveTypeGroups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveTypeGroups.go.html to see an example of how to use ListSensitiveTypeGroupsRequest.
type ListSensitiveTypeGroupsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSensitiveTypeGroupsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources that match the specified lifecycle state.
	LifecycleState ListSensitiveTypeGroupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the sensitive type group resource.
	SensitiveTypeGroupId *string `mandatory:"false" contributesTo:"query" name:"sensitiveTypeGroupId"`

	// A filter to return only the resources that were created after the specified date and time, as defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Using TimeCreatedGreaterThanOrEqualToQueryParam parameter retrieves all resources created after that date.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// Search for resources that were created before a specific date.
	// Specifying this parameter corresponding `timeCreatedLessThan`
	// parameter will retrieve all resources created before the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSensitiveTypeGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeCreated is descending.
	// The default order for displayName is ascending.
	SortBy ListSensitiveTypeGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSensitiveTypeGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSensitiveTypeGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSensitiveTypeGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSensitiveTypeGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSensitiveTypeGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSensitiveTypeGroupsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSensitiveTypeGroupsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypeGroupsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSensitiveTypeGroupsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypeGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSensitiveTypeGroupsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypeGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSensitiveTypeGroupsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSensitiveTypeGroupsResponse wrapper for the ListSensitiveTypeGroups operation
type ListSensitiveTypeGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SensitiveTypeGroupCollection instances
	SensitiveTypeGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSensitiveTypeGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSensitiveTypeGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSensitiveTypeGroupsAccessLevelEnum Enum with underlying type: string
type ListSensitiveTypeGroupsAccessLevelEnum string

// Set of constants representing the allowable values for ListSensitiveTypeGroupsAccessLevelEnum
const (
	ListSensitiveTypeGroupsAccessLevelRestricted ListSensitiveTypeGroupsAccessLevelEnum = "RESTRICTED"
	ListSensitiveTypeGroupsAccessLevelAccessible ListSensitiveTypeGroupsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSensitiveTypeGroupsAccessLevelEnum = map[string]ListSensitiveTypeGroupsAccessLevelEnum{
	"RESTRICTED": ListSensitiveTypeGroupsAccessLevelRestricted,
	"ACCESSIBLE": ListSensitiveTypeGroupsAccessLevelAccessible,
}

var mappingListSensitiveTypeGroupsAccessLevelEnumLowerCase = map[string]ListSensitiveTypeGroupsAccessLevelEnum{
	"restricted": ListSensitiveTypeGroupsAccessLevelRestricted,
	"accessible": ListSensitiveTypeGroupsAccessLevelAccessible,
}

// GetListSensitiveTypeGroupsAccessLevelEnumValues Enumerates the set of values for ListSensitiveTypeGroupsAccessLevelEnum
func GetListSensitiveTypeGroupsAccessLevelEnumValues() []ListSensitiveTypeGroupsAccessLevelEnum {
	values := make([]ListSensitiveTypeGroupsAccessLevelEnum, 0)
	for _, v := range mappingListSensitiveTypeGroupsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypeGroupsAccessLevelEnumStringValues Enumerates the set of values in String for ListSensitiveTypeGroupsAccessLevelEnum
func GetListSensitiveTypeGroupsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSensitiveTypeGroupsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypeGroupsAccessLevelEnum(val string) (ListSensitiveTypeGroupsAccessLevelEnum, bool) {
	enum, ok := mappingListSensitiveTypeGroupsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypeGroupsLifecycleStateEnum Enum with underlying type: string
type ListSensitiveTypeGroupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSensitiveTypeGroupsLifecycleStateEnum
const (
	ListSensitiveTypeGroupsLifecycleStateCreating ListSensitiveTypeGroupsLifecycleStateEnum = "CREATING"
	ListSensitiveTypeGroupsLifecycleStateActive   ListSensitiveTypeGroupsLifecycleStateEnum = "ACTIVE"
	ListSensitiveTypeGroupsLifecycleStateUpdating ListSensitiveTypeGroupsLifecycleStateEnum = "UPDATING"
	ListSensitiveTypeGroupsLifecycleStateDeleting ListSensitiveTypeGroupsLifecycleStateEnum = "DELETING"
	ListSensitiveTypeGroupsLifecycleStateDeleted  ListSensitiveTypeGroupsLifecycleStateEnum = "DELETED"
	ListSensitiveTypeGroupsLifecycleStateFailed   ListSensitiveTypeGroupsLifecycleStateEnum = "FAILED"
)

var mappingListSensitiveTypeGroupsLifecycleStateEnum = map[string]ListSensitiveTypeGroupsLifecycleStateEnum{
	"CREATING": ListSensitiveTypeGroupsLifecycleStateCreating,
	"ACTIVE":   ListSensitiveTypeGroupsLifecycleStateActive,
	"UPDATING": ListSensitiveTypeGroupsLifecycleStateUpdating,
	"DELETING": ListSensitiveTypeGroupsLifecycleStateDeleting,
	"DELETED":  ListSensitiveTypeGroupsLifecycleStateDeleted,
	"FAILED":   ListSensitiveTypeGroupsLifecycleStateFailed,
}

var mappingListSensitiveTypeGroupsLifecycleStateEnumLowerCase = map[string]ListSensitiveTypeGroupsLifecycleStateEnum{
	"creating": ListSensitiveTypeGroupsLifecycleStateCreating,
	"active":   ListSensitiveTypeGroupsLifecycleStateActive,
	"updating": ListSensitiveTypeGroupsLifecycleStateUpdating,
	"deleting": ListSensitiveTypeGroupsLifecycleStateDeleting,
	"deleted":  ListSensitiveTypeGroupsLifecycleStateDeleted,
	"failed":   ListSensitiveTypeGroupsLifecycleStateFailed,
}

// GetListSensitiveTypeGroupsLifecycleStateEnumValues Enumerates the set of values for ListSensitiveTypeGroupsLifecycleStateEnum
func GetListSensitiveTypeGroupsLifecycleStateEnumValues() []ListSensitiveTypeGroupsLifecycleStateEnum {
	values := make([]ListSensitiveTypeGroupsLifecycleStateEnum, 0)
	for _, v := range mappingListSensitiveTypeGroupsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypeGroupsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSensitiveTypeGroupsLifecycleStateEnum
func GetListSensitiveTypeGroupsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSensitiveTypeGroupsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypeGroupsLifecycleStateEnum(val string) (ListSensitiveTypeGroupsLifecycleStateEnum, bool) {
	enum, ok := mappingListSensitiveTypeGroupsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypeGroupsSortOrderEnum Enum with underlying type: string
type ListSensitiveTypeGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListSensitiveTypeGroupsSortOrderEnum
const (
	ListSensitiveTypeGroupsSortOrderAsc  ListSensitiveTypeGroupsSortOrderEnum = "ASC"
	ListSensitiveTypeGroupsSortOrderDesc ListSensitiveTypeGroupsSortOrderEnum = "DESC"
)

var mappingListSensitiveTypeGroupsSortOrderEnum = map[string]ListSensitiveTypeGroupsSortOrderEnum{
	"ASC":  ListSensitiveTypeGroupsSortOrderAsc,
	"DESC": ListSensitiveTypeGroupsSortOrderDesc,
}

var mappingListSensitiveTypeGroupsSortOrderEnumLowerCase = map[string]ListSensitiveTypeGroupsSortOrderEnum{
	"asc":  ListSensitiveTypeGroupsSortOrderAsc,
	"desc": ListSensitiveTypeGroupsSortOrderDesc,
}

// GetListSensitiveTypeGroupsSortOrderEnumValues Enumerates the set of values for ListSensitiveTypeGroupsSortOrderEnum
func GetListSensitiveTypeGroupsSortOrderEnumValues() []ListSensitiveTypeGroupsSortOrderEnum {
	values := make([]ListSensitiveTypeGroupsSortOrderEnum, 0)
	for _, v := range mappingListSensitiveTypeGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypeGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListSensitiveTypeGroupsSortOrderEnum
func GetListSensitiveTypeGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSensitiveTypeGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypeGroupsSortOrderEnum(val string) (ListSensitiveTypeGroupsSortOrderEnum, bool) {
	enum, ok := mappingListSensitiveTypeGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypeGroupsSortByEnum Enum with underlying type: string
type ListSensitiveTypeGroupsSortByEnum string

// Set of constants representing the allowable values for ListSensitiveTypeGroupsSortByEnum
const (
	ListSensitiveTypeGroupsSortByTimecreated ListSensitiveTypeGroupsSortByEnum = "timeCreated"
	ListSensitiveTypeGroupsSortByDisplayname ListSensitiveTypeGroupsSortByEnum = "displayName"
)

var mappingListSensitiveTypeGroupsSortByEnum = map[string]ListSensitiveTypeGroupsSortByEnum{
	"timeCreated": ListSensitiveTypeGroupsSortByTimecreated,
	"displayName": ListSensitiveTypeGroupsSortByDisplayname,
}

var mappingListSensitiveTypeGroupsSortByEnumLowerCase = map[string]ListSensitiveTypeGroupsSortByEnum{
	"timecreated": ListSensitiveTypeGroupsSortByTimecreated,
	"displayname": ListSensitiveTypeGroupsSortByDisplayname,
}

// GetListSensitiveTypeGroupsSortByEnumValues Enumerates the set of values for ListSensitiveTypeGroupsSortByEnum
func GetListSensitiveTypeGroupsSortByEnumValues() []ListSensitiveTypeGroupsSortByEnum {
	values := make([]ListSensitiveTypeGroupsSortByEnum, 0)
	for _, v := range mappingListSensitiveTypeGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypeGroupsSortByEnumStringValues Enumerates the set of values in String for ListSensitiveTypeGroupsSortByEnum
func GetListSensitiveTypeGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSensitiveTypeGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypeGroupsSortByEnum(val string) (ListSensitiveTypeGroupsSortByEnum, bool) {
	enum, ok := mappingListSensitiveTypeGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
