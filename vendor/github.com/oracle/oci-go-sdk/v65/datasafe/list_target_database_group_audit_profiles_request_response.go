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

// ListTargetDatabaseGroupAuditProfilesRequest wrapper for the ListTargetDatabaseGroupAuditProfiles operation
type ListTargetDatabaseGroupAuditProfilesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListTargetDatabaseGroupAuditProfilesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A optional filter to return only resources that match the specified id.
	TargetDatabaseGroupAuditProfileId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseGroupAuditProfileId"`

	// A filter to return the target database group that matches the specified OCID.
	TargetDatabaseGroupId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseGroupId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A optional filter to return only resources that match the specified lifecycle state.
	LifecycleState ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Indicates if you want to continue audit record collection beyond the free limit
	// of one million audit records per month per target database, incurring additional charges.
	// The default value is inherited from the global settings. You can change at the global level
	// or at the target level.
	IsPaidUsageEnabled *bool `mandatory:"false" contributesTo:"query" name:"isPaidUsageEnabled"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTargetDatabaseGroupAuditProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListTargetDatabaseGroupAuditProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetDatabaseGroupAuditProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetDatabaseGroupAuditProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetDatabaseGroupAuditProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetDatabaseGroupAuditProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetDatabaseGroupAuditProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetDatabaseGroupAuditProfilesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListTargetDatabaseGroupAuditProfilesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabaseGroupAuditProfilesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTargetDatabaseGroupAuditProfilesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabaseGroupAuditProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetDatabaseGroupAuditProfilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabaseGroupAuditProfilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetDatabaseGroupAuditProfilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTargetDatabaseGroupAuditProfilesResponse wrapper for the ListTargetDatabaseGroupAuditProfiles operation
type ListTargetDatabaseGroupAuditProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetDatabaseGroupAuditProfileCollection instances
	TargetDatabaseGroupAuditProfileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTargetDatabaseGroupAuditProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetDatabaseGroupAuditProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetDatabaseGroupAuditProfilesAccessLevelEnum Enum with underlying type: string
type ListTargetDatabaseGroupAuditProfilesAccessLevelEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupAuditProfilesAccessLevelEnum
const (
	ListTargetDatabaseGroupAuditProfilesAccessLevelRestricted ListTargetDatabaseGroupAuditProfilesAccessLevelEnum = "RESTRICTED"
	ListTargetDatabaseGroupAuditProfilesAccessLevelAccessible ListTargetDatabaseGroupAuditProfilesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListTargetDatabaseGroupAuditProfilesAccessLevelEnum = map[string]ListTargetDatabaseGroupAuditProfilesAccessLevelEnum{
	"RESTRICTED": ListTargetDatabaseGroupAuditProfilesAccessLevelRestricted,
	"ACCESSIBLE": ListTargetDatabaseGroupAuditProfilesAccessLevelAccessible,
}

var mappingListTargetDatabaseGroupAuditProfilesAccessLevelEnumLowerCase = map[string]ListTargetDatabaseGroupAuditProfilesAccessLevelEnum{
	"restricted": ListTargetDatabaseGroupAuditProfilesAccessLevelRestricted,
	"accessible": ListTargetDatabaseGroupAuditProfilesAccessLevelAccessible,
}

// GetListTargetDatabaseGroupAuditProfilesAccessLevelEnumValues Enumerates the set of values for ListTargetDatabaseGroupAuditProfilesAccessLevelEnum
func GetListTargetDatabaseGroupAuditProfilesAccessLevelEnumValues() []ListTargetDatabaseGroupAuditProfilesAccessLevelEnum {
	values := make([]ListTargetDatabaseGroupAuditProfilesAccessLevelEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupAuditProfilesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupAuditProfilesAccessLevelEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupAuditProfilesAccessLevelEnum
func GetListTargetDatabaseGroupAuditProfilesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListTargetDatabaseGroupAuditProfilesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupAuditProfilesAccessLevelEnum(val string) (ListTargetDatabaseGroupAuditProfilesAccessLevelEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupAuditProfilesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum Enum with underlying type: string
type ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum
const (
	ListTargetDatabaseGroupAuditProfilesLifecycleStateCreating       ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum = "CREATING"
	ListTargetDatabaseGroupAuditProfilesLifecycleStateUpdating       ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum = "UPDATING"
	ListTargetDatabaseGroupAuditProfilesLifecycleStateActive         ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum = "ACTIVE"
	ListTargetDatabaseGroupAuditProfilesLifecycleStateDeleting       ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum = "DELETING"
	ListTargetDatabaseGroupAuditProfilesLifecycleStateFailed         ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum = "FAILED"
	ListTargetDatabaseGroupAuditProfilesLifecycleStateNeedsAttention ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListTargetDatabaseGroupAuditProfilesLifecycleStateDeleted        ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum = "DELETED"
)

var mappingListTargetDatabaseGroupAuditProfilesLifecycleStateEnum = map[string]ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum{
	"CREATING":        ListTargetDatabaseGroupAuditProfilesLifecycleStateCreating,
	"UPDATING":        ListTargetDatabaseGroupAuditProfilesLifecycleStateUpdating,
	"ACTIVE":          ListTargetDatabaseGroupAuditProfilesLifecycleStateActive,
	"DELETING":        ListTargetDatabaseGroupAuditProfilesLifecycleStateDeleting,
	"FAILED":          ListTargetDatabaseGroupAuditProfilesLifecycleStateFailed,
	"NEEDS_ATTENTION": ListTargetDatabaseGroupAuditProfilesLifecycleStateNeedsAttention,
	"DELETED":         ListTargetDatabaseGroupAuditProfilesLifecycleStateDeleted,
}

var mappingListTargetDatabaseGroupAuditProfilesLifecycleStateEnumLowerCase = map[string]ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum{
	"creating":        ListTargetDatabaseGroupAuditProfilesLifecycleStateCreating,
	"updating":        ListTargetDatabaseGroupAuditProfilesLifecycleStateUpdating,
	"active":          ListTargetDatabaseGroupAuditProfilesLifecycleStateActive,
	"deleting":        ListTargetDatabaseGroupAuditProfilesLifecycleStateDeleting,
	"failed":          ListTargetDatabaseGroupAuditProfilesLifecycleStateFailed,
	"needs_attention": ListTargetDatabaseGroupAuditProfilesLifecycleStateNeedsAttention,
	"deleted":         ListTargetDatabaseGroupAuditProfilesLifecycleStateDeleted,
}

// GetListTargetDatabaseGroupAuditProfilesLifecycleStateEnumValues Enumerates the set of values for ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum
func GetListTargetDatabaseGroupAuditProfilesLifecycleStateEnumValues() []ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum {
	values := make([]ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupAuditProfilesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupAuditProfilesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum
func GetListTargetDatabaseGroupAuditProfilesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingListTargetDatabaseGroupAuditProfilesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupAuditProfilesLifecycleStateEnum(val string) (ListTargetDatabaseGroupAuditProfilesLifecycleStateEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupAuditProfilesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabaseGroupAuditProfilesSortOrderEnum Enum with underlying type: string
type ListTargetDatabaseGroupAuditProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupAuditProfilesSortOrderEnum
const (
	ListTargetDatabaseGroupAuditProfilesSortOrderAsc  ListTargetDatabaseGroupAuditProfilesSortOrderEnum = "ASC"
	ListTargetDatabaseGroupAuditProfilesSortOrderDesc ListTargetDatabaseGroupAuditProfilesSortOrderEnum = "DESC"
)

var mappingListTargetDatabaseGroupAuditProfilesSortOrderEnum = map[string]ListTargetDatabaseGroupAuditProfilesSortOrderEnum{
	"ASC":  ListTargetDatabaseGroupAuditProfilesSortOrderAsc,
	"DESC": ListTargetDatabaseGroupAuditProfilesSortOrderDesc,
}

var mappingListTargetDatabaseGroupAuditProfilesSortOrderEnumLowerCase = map[string]ListTargetDatabaseGroupAuditProfilesSortOrderEnum{
	"asc":  ListTargetDatabaseGroupAuditProfilesSortOrderAsc,
	"desc": ListTargetDatabaseGroupAuditProfilesSortOrderDesc,
}

// GetListTargetDatabaseGroupAuditProfilesSortOrderEnumValues Enumerates the set of values for ListTargetDatabaseGroupAuditProfilesSortOrderEnum
func GetListTargetDatabaseGroupAuditProfilesSortOrderEnumValues() []ListTargetDatabaseGroupAuditProfilesSortOrderEnum {
	values := make([]ListTargetDatabaseGroupAuditProfilesSortOrderEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupAuditProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupAuditProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupAuditProfilesSortOrderEnum
func GetListTargetDatabaseGroupAuditProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetDatabaseGroupAuditProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupAuditProfilesSortOrderEnum(val string) (ListTargetDatabaseGroupAuditProfilesSortOrderEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupAuditProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabaseGroupAuditProfilesSortByEnum Enum with underlying type: string
type ListTargetDatabaseGroupAuditProfilesSortByEnum string

// Set of constants representing the allowable values for ListTargetDatabaseGroupAuditProfilesSortByEnum
const (
	ListTargetDatabaseGroupAuditProfilesSortByTimecreated ListTargetDatabaseGroupAuditProfilesSortByEnum = "TIMECREATED"
	ListTargetDatabaseGroupAuditProfilesSortByDisplayname ListTargetDatabaseGroupAuditProfilesSortByEnum = "DISPLAYNAME"
)

var mappingListTargetDatabaseGroupAuditProfilesSortByEnum = map[string]ListTargetDatabaseGroupAuditProfilesSortByEnum{
	"TIMECREATED": ListTargetDatabaseGroupAuditProfilesSortByTimecreated,
	"DISPLAYNAME": ListTargetDatabaseGroupAuditProfilesSortByDisplayname,
}

var mappingListTargetDatabaseGroupAuditProfilesSortByEnumLowerCase = map[string]ListTargetDatabaseGroupAuditProfilesSortByEnum{
	"timecreated": ListTargetDatabaseGroupAuditProfilesSortByTimecreated,
	"displayname": ListTargetDatabaseGroupAuditProfilesSortByDisplayname,
}

// GetListTargetDatabaseGroupAuditProfilesSortByEnumValues Enumerates the set of values for ListTargetDatabaseGroupAuditProfilesSortByEnum
func GetListTargetDatabaseGroupAuditProfilesSortByEnumValues() []ListTargetDatabaseGroupAuditProfilesSortByEnum {
	values := make([]ListTargetDatabaseGroupAuditProfilesSortByEnum, 0)
	for _, v := range mappingListTargetDatabaseGroupAuditProfilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabaseGroupAuditProfilesSortByEnumStringValues Enumerates the set of values in String for ListTargetDatabaseGroupAuditProfilesSortByEnum
func GetListTargetDatabaseGroupAuditProfilesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListTargetDatabaseGroupAuditProfilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabaseGroupAuditProfilesSortByEnum(val string) (ListTargetDatabaseGroupAuditProfilesSortByEnum, bool) {
	enum, ok := mappingListTargetDatabaseGroupAuditProfilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
