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

// ListTargetDatabaseGroupsRequest wrapper for the ListTargetDatabaseGroups operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetDatabaseGroups.go.html to see an example of how to use ListTargetDatabaseGroupsRequest.
type ListTargetDatabaseGroupsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The scim query filter parameter accepts filter expressions that use the syntax described in Section 3.2.2.2
	// of the System for Cross-Domain Identity Management (SCIM) specification, which is available
	// at RFC3339 (https://tools.ietf.org/html/draft-ietf-scim-api-12). In SCIM filtering expressions,
	// text, date, and time values must be enclosed in quotation marks, with date and time values using ISO-8601 format.
	// (Numeric and boolean values should not be quoted.)
	// Ex:** filter=(targetDatabaseId eq 'ocid1.datasafetargetdatabase.oc1.iad.abuwcljr3u2va4ba5wek53idpe5qq5kkbigzclscc6mysfecxzjt5dgmxqza')
	Filter *string `mandatory:"false" contributesTo:"query" name:"filter"`

	// The sorting field for your request. You can only specify a single sorting order (sortOrder). The default order for timeCreated is descending.
	SortBy ListTargetDatabaseGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to retrieve resources that exclusively align with the designated lifecycle state.
	LifecycleState ListTargetDatabaseGroupsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListTargetDatabaseGroupsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTargetDatabaseGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return the target database group that matches the specified OCID.
	TargetDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseGroupId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetDatabaseGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetDatabaseGroupsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetDatabaseGroupsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetDatabaseGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetDatabaseGroupsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetDatabaseGroupsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetDatabaseGroupsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabaseGroupsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTargetDatabaseGroupsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabaseGroupsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListTargetDatabaseGroupsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabaseGroupsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetDatabaseGroupsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetDatabaseGroupsResponse wrapper for the ListTargetDatabaseGroups operation
type ListTargetDatabaseGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetDatabaseGroupCollection instances
	TargetDatabaseGroupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTargetDatabaseGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetDatabaseGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetDatabaseGroupsSortByEnum Enum with underlying type: string
type ListTargetDatabaseGroupsSortByEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupsSortByEnum
const (
	ListTargetDatabaseGroupsSortByTimecreated          ListTargetDatabaseGroupsSortByEnum = "timeCreated"
	ListTargetDatabaseGroupsSortByDisplayname          ListTargetDatabaseGroupsSortByEnum = "displayName"
	ListTargetDatabaseGroupsSortByLifecyclestate       ListTargetDatabaseGroupsSortByEnum = "lifecycleState"
	ListTargetDatabaseGroupsSortByMembershipupdatetime ListTargetDatabaseGroupsSortByEnum = "membershipUpdateTime"
)

var mappingListTargetDatabaseGroupsSortByEnum = map[string]ListTargetDatabaseGroupsSortByEnum{
	"timeCreated":          ListTargetDatabaseGroupsSortByTimecreated,
	"displayName":          ListTargetDatabaseGroupsSortByDisplayname,
	"lifecycleState":       ListTargetDatabaseGroupsSortByLifecyclestate,
	"membershipUpdateTime": ListTargetDatabaseGroupsSortByMembershipupdatetime,
}

var mappingListTargetDatabaseGroupsSortByEnumLowerCase = map[string]ListTargetDatabaseGroupsSortByEnum{
	"timecreated":          ListTargetDatabaseGroupsSortByTimecreated,
	"displayname":          ListTargetDatabaseGroupsSortByDisplayname,
	"lifecyclestate":       ListTargetDatabaseGroupsSortByLifecyclestate,
	"membershipupdatetime": ListTargetDatabaseGroupsSortByMembershipupdatetime,
}

// GetListTargetDatabaseGroupsSortByEnumValues Enumerates the set of values for ListTargetDatabaseGroupsSortByEnum
func GetListTargetDatabaseGroupsSortByEnumValues() []ListTargetDatabaseGroupsSortByEnum {
	values := make([]ListTargetDatabaseGroupsSortByEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupsSortByEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupsSortByEnum
func GetListTargetDatabaseGroupsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"lifecycleState",
		"membershipUpdateTime",
	}
}

// GetMappingListTargetDatabaseGroupsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupsSortByEnum(val string) (ListTargetDatabaseGroupsSortByEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabaseGroupsLifecycleStateEnum Enum with underlying type: string
type ListTargetDatabaseGroupsLifecycleStateEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupsLifecycleStateEnum
const (
	ListTargetDatabaseGroupsLifecycleStateCreating       ListTargetDatabaseGroupsLifecycleStateEnum = "CREATING"
	ListTargetDatabaseGroupsLifecycleStateUpdating       ListTargetDatabaseGroupsLifecycleStateEnum = "UPDATING"
	ListTargetDatabaseGroupsLifecycleStateActive         ListTargetDatabaseGroupsLifecycleStateEnum = "ACTIVE"
	ListTargetDatabaseGroupsLifecycleStateDeleting       ListTargetDatabaseGroupsLifecycleStateEnum = "DELETING"
	ListTargetDatabaseGroupsLifecycleStateDeleted        ListTargetDatabaseGroupsLifecycleStateEnum = "DELETED"
	ListTargetDatabaseGroupsLifecycleStateFailed         ListTargetDatabaseGroupsLifecycleStateEnum = "FAILED"
	ListTargetDatabaseGroupsLifecycleStateNeedsAttention ListTargetDatabaseGroupsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListTargetDatabaseGroupsLifecycleStateEnum = map[string]ListTargetDatabaseGroupsLifecycleStateEnum{
	"CREATING":        ListTargetDatabaseGroupsLifecycleStateCreating,
	"UPDATING":        ListTargetDatabaseGroupsLifecycleStateUpdating,
	"ACTIVE":          ListTargetDatabaseGroupsLifecycleStateActive,
	"DELETING":        ListTargetDatabaseGroupsLifecycleStateDeleting,
	"DELETED":         ListTargetDatabaseGroupsLifecycleStateDeleted,
	"FAILED":          ListTargetDatabaseGroupsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListTargetDatabaseGroupsLifecycleStateNeedsAttention,
}

var mappingListTargetDatabaseGroupsLifecycleStateEnumLowerCase = map[string]ListTargetDatabaseGroupsLifecycleStateEnum{
	"creating":        ListTargetDatabaseGroupsLifecycleStateCreating,
	"updating":        ListTargetDatabaseGroupsLifecycleStateUpdating,
	"active":          ListTargetDatabaseGroupsLifecycleStateActive,
	"deleting":        ListTargetDatabaseGroupsLifecycleStateDeleting,
	"deleted":         ListTargetDatabaseGroupsLifecycleStateDeleted,
	"failed":          ListTargetDatabaseGroupsLifecycleStateFailed,
	"needs_attention": ListTargetDatabaseGroupsLifecycleStateNeedsAttention,
}

// GetListTargetDatabaseGroupsLifecycleStateEnumValues Enumerates the set of values for ListTargetDatabaseGroupsLifecycleStateEnum
func GetListTargetDatabaseGroupsLifecycleStateEnumValues() []ListTargetDatabaseGroupsLifecycleStateEnum {
	values := make([]ListTargetDatabaseGroupsLifecycleStateEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupsLifecycleStateEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupsLifecycleStateEnum
func GetListTargetDatabaseGroupsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListTargetDatabaseGroupsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupsLifecycleStateEnum(val string) (ListTargetDatabaseGroupsLifecycleStateEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabaseGroupsAccessLevelEnum Enum with underlying type: string
type ListTargetDatabaseGroupsAccessLevelEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupsAccessLevelEnum
const (
	ListTargetDatabaseGroupsAccessLevelRestricted ListTargetDatabaseGroupsAccessLevelEnum = "RESTRICTED"
	ListTargetDatabaseGroupsAccessLevelAccessible ListTargetDatabaseGroupsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListTargetDatabaseGroupsAccessLevelEnum = map[string]ListTargetDatabaseGroupsAccessLevelEnum{
	"RESTRICTED": ListTargetDatabaseGroupsAccessLevelRestricted,
	"ACCESSIBLE": ListTargetDatabaseGroupsAccessLevelAccessible,
}

var mappingListTargetDatabaseGroupsAccessLevelEnumLowerCase = map[string]ListTargetDatabaseGroupsAccessLevelEnum{
	"restricted": ListTargetDatabaseGroupsAccessLevelRestricted,
	"accessible": ListTargetDatabaseGroupsAccessLevelAccessible,
}

// GetListTargetDatabaseGroupsAccessLevelEnumValues Enumerates the set of values for ListTargetDatabaseGroupsAccessLevelEnum
func GetListTargetDatabaseGroupsAccessLevelEnumValues() []ListTargetDatabaseGroupsAccessLevelEnum {
	values := make([]ListTargetDatabaseGroupsAccessLevelEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupsAccessLevelEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupsAccessLevelEnum
func GetListTargetDatabaseGroupsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListTargetDatabaseGroupsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupsAccessLevelEnum(val string) (ListTargetDatabaseGroupsAccessLevelEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabaseGroupsSortOrderEnum Enum with underlying type: string
type ListTargetDatabaseGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupsSortOrderEnum
const (
	ListTargetDatabaseGroupsSortOrderAsc  ListTargetDatabaseGroupsSortOrderEnum = "ASC"
	ListTargetDatabaseGroupsSortOrderDesc ListTargetDatabaseGroupsSortOrderEnum = "DESC"
)

var mappingListTargetDatabaseGroupsSortOrderEnum = map[string]ListTargetDatabaseGroupsSortOrderEnum{
	"ASC":  ListTargetDatabaseGroupsSortOrderAsc,
	"DESC": ListTargetDatabaseGroupsSortOrderDesc,
}

var mappingListTargetDatabaseGroupsSortOrderEnumLowerCase = map[string]ListTargetDatabaseGroupsSortOrderEnum{
	"asc":  ListTargetDatabaseGroupsSortOrderAsc,
	"desc": ListTargetDatabaseGroupsSortOrderDesc,
}

// GetListTargetDatabaseGroupsSortOrderEnumValues Enumerates the set of values for ListTargetDatabaseGroupsSortOrderEnum
func GetListTargetDatabaseGroupsSortOrderEnumValues() []ListTargetDatabaseGroupsSortOrderEnum {
	values := make([]ListTargetDatabaseGroupsSortOrderEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupsSortOrderEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupsSortOrderEnum
func GetListTargetDatabaseGroupsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetDatabaseGroupsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupsSortOrderEnum(val string) (ListTargetDatabaseGroupsSortOrderEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
