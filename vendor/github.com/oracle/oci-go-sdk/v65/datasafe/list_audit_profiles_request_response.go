// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAuditProfilesRequest wrapper for the ListAuditProfiles operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditProfiles.go.html to see an example of how to use ListAuditProfilesRequest.
type ListAuditProfilesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAuditProfilesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A optional filter to return only resources that match the specified id.
	AuditProfileId *string `mandatory:"false" contributesTo:"query" name:"auditProfileId"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A optional filter to return only resources that match the specified lifecycle state.
	LifecycleState ListAuditProfilesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A optional filter to return only resources that match the specified retention configured value.
	IsOverrideGlobalRetentionSetting *bool `mandatory:"false" contributesTo:"query" name:"isOverrideGlobalRetentionSetting"`

	// Indicates if you want to continue audit record collection beyond the free limit
	// of one million audit records per month per target database, incurring additional charges.
	// The default value is inherited from the global settings. You can change at the global level
	// or at the target level.
	IsPaidUsageEnabled *bool `mandatory:"false" contributesTo:"query" name:"isPaidUsageEnabled"`

	// A filter to return only items that have count of audit records collected greater than or equal to the specified value.
	AuditCollectedVolumeGreaterThanOrEqualTo *int64 `mandatory:"false" contributesTo:"query" name:"auditCollectedVolumeGreaterThanOrEqualTo"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAuditProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListAuditProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuditProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuditProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuditProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuditProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuditProfilesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuditProfilesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAuditProfilesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditProfilesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAuditProfilesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditProfilesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAuditProfilesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditProfilesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAuditProfilesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuditProfilesResponse wrapper for the ListAuditProfiles operation
type ListAuditProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuditProfileCollection instances
	AuditProfileCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAuditProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuditProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuditProfilesAccessLevelEnum Enum with underlying type: string
type ListAuditProfilesAccessLevelEnum string

// Set of constants representing the allowable values for ListAuditProfilesAccessLevelEnum
const (
	ListAuditProfilesAccessLevelRestricted ListAuditProfilesAccessLevelEnum = "RESTRICTED"
	ListAuditProfilesAccessLevelAccessible ListAuditProfilesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAuditProfilesAccessLevelEnum = map[string]ListAuditProfilesAccessLevelEnum{
	"RESTRICTED": ListAuditProfilesAccessLevelRestricted,
	"ACCESSIBLE": ListAuditProfilesAccessLevelAccessible,
}

var mappingListAuditProfilesAccessLevelEnumLowerCase = map[string]ListAuditProfilesAccessLevelEnum{
	"restricted": ListAuditProfilesAccessLevelRestricted,
	"accessible": ListAuditProfilesAccessLevelAccessible,
}

// GetListAuditProfilesAccessLevelEnumValues Enumerates the set of values for ListAuditProfilesAccessLevelEnum
func GetListAuditProfilesAccessLevelEnumValues() []ListAuditProfilesAccessLevelEnum {
	values := make([]ListAuditProfilesAccessLevelEnum, 0)
	for _, v := range mappingListAuditProfilesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditProfilesAccessLevelEnumStringValues Enumerates the set of values in String for ListAuditProfilesAccessLevelEnum
func GetListAuditProfilesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAuditProfilesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditProfilesAccessLevelEnum(val string) (ListAuditProfilesAccessLevelEnum, bool) {
	enum, ok := mappingListAuditProfilesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditProfilesLifecycleStateEnum Enum with underlying type: string
type ListAuditProfilesLifecycleStateEnum string

// Set of constants representing the allowable values for ListAuditProfilesLifecycleStateEnum
const (
	ListAuditProfilesLifecycleStateCreating       ListAuditProfilesLifecycleStateEnum = "CREATING"
	ListAuditProfilesLifecycleStateUpdating       ListAuditProfilesLifecycleStateEnum = "UPDATING"
	ListAuditProfilesLifecycleStateActive         ListAuditProfilesLifecycleStateEnum = "ACTIVE"
	ListAuditProfilesLifecycleStateDeleting       ListAuditProfilesLifecycleStateEnum = "DELETING"
	ListAuditProfilesLifecycleStateFailed         ListAuditProfilesLifecycleStateEnum = "FAILED"
	ListAuditProfilesLifecycleStateNeedsAttention ListAuditProfilesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListAuditProfilesLifecycleStateDeleted        ListAuditProfilesLifecycleStateEnum = "DELETED"
)

var mappingListAuditProfilesLifecycleStateEnum = map[string]ListAuditProfilesLifecycleStateEnum{
	"CREATING":        ListAuditProfilesLifecycleStateCreating,
	"UPDATING":        ListAuditProfilesLifecycleStateUpdating,
	"ACTIVE":          ListAuditProfilesLifecycleStateActive,
	"DELETING":        ListAuditProfilesLifecycleStateDeleting,
	"FAILED":          ListAuditProfilesLifecycleStateFailed,
	"NEEDS_ATTENTION": ListAuditProfilesLifecycleStateNeedsAttention,
	"DELETED":         ListAuditProfilesLifecycleStateDeleted,
}

var mappingListAuditProfilesLifecycleStateEnumLowerCase = map[string]ListAuditProfilesLifecycleStateEnum{
	"creating":        ListAuditProfilesLifecycleStateCreating,
	"updating":        ListAuditProfilesLifecycleStateUpdating,
	"active":          ListAuditProfilesLifecycleStateActive,
	"deleting":        ListAuditProfilesLifecycleStateDeleting,
	"failed":          ListAuditProfilesLifecycleStateFailed,
	"needs_attention": ListAuditProfilesLifecycleStateNeedsAttention,
	"deleted":         ListAuditProfilesLifecycleStateDeleted,
}

// GetListAuditProfilesLifecycleStateEnumValues Enumerates the set of values for ListAuditProfilesLifecycleStateEnum
func GetListAuditProfilesLifecycleStateEnumValues() []ListAuditProfilesLifecycleStateEnum {
	values := make([]ListAuditProfilesLifecycleStateEnum, 0)
	for _, v := range mappingListAuditProfilesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditProfilesLifecycleStateEnumStringValues Enumerates the set of values in String for ListAuditProfilesLifecycleStateEnum
func GetListAuditProfilesLifecycleStateEnumStringValues() []string {
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

// GetMappingListAuditProfilesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditProfilesLifecycleStateEnum(val string) (ListAuditProfilesLifecycleStateEnum, bool) {
	enum, ok := mappingListAuditProfilesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditProfilesSortOrderEnum Enum with underlying type: string
type ListAuditProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListAuditProfilesSortOrderEnum
const (
	ListAuditProfilesSortOrderAsc  ListAuditProfilesSortOrderEnum = "ASC"
	ListAuditProfilesSortOrderDesc ListAuditProfilesSortOrderEnum = "DESC"
)

var mappingListAuditProfilesSortOrderEnum = map[string]ListAuditProfilesSortOrderEnum{
	"ASC":  ListAuditProfilesSortOrderAsc,
	"DESC": ListAuditProfilesSortOrderDesc,
}

var mappingListAuditProfilesSortOrderEnumLowerCase = map[string]ListAuditProfilesSortOrderEnum{
	"asc":  ListAuditProfilesSortOrderAsc,
	"desc": ListAuditProfilesSortOrderDesc,
}

// GetListAuditProfilesSortOrderEnumValues Enumerates the set of values for ListAuditProfilesSortOrderEnum
func GetListAuditProfilesSortOrderEnumValues() []ListAuditProfilesSortOrderEnum {
	values := make([]ListAuditProfilesSortOrderEnum, 0)
	for _, v := range mappingListAuditProfilesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditProfilesSortOrderEnumStringValues Enumerates the set of values in String for ListAuditProfilesSortOrderEnum
func GetListAuditProfilesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAuditProfilesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditProfilesSortOrderEnum(val string) (ListAuditProfilesSortOrderEnum, bool) {
	enum, ok := mappingListAuditProfilesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditProfilesSortByEnum Enum with underlying type: string
type ListAuditProfilesSortByEnum string

// Set of constants representing the allowable values for ListAuditProfilesSortByEnum
const (
	ListAuditProfilesSortByTimecreated ListAuditProfilesSortByEnum = "TIMECREATED"
	ListAuditProfilesSortByDisplayname ListAuditProfilesSortByEnum = "DISPLAYNAME"
)

var mappingListAuditProfilesSortByEnum = map[string]ListAuditProfilesSortByEnum{
	"TIMECREATED": ListAuditProfilesSortByTimecreated,
	"DISPLAYNAME": ListAuditProfilesSortByDisplayname,
}

var mappingListAuditProfilesSortByEnumLowerCase = map[string]ListAuditProfilesSortByEnum{
	"timecreated": ListAuditProfilesSortByTimecreated,
	"displayname": ListAuditProfilesSortByDisplayname,
}

// GetListAuditProfilesSortByEnumValues Enumerates the set of values for ListAuditProfilesSortByEnum
func GetListAuditProfilesSortByEnumValues() []ListAuditProfilesSortByEnum {
	values := make([]ListAuditProfilesSortByEnum, 0)
	for _, v := range mappingListAuditProfilesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditProfilesSortByEnumStringValues Enumerates the set of values in String for ListAuditProfilesSortByEnum
func GetListAuditProfilesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAuditProfilesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditProfilesSortByEnum(val string) (ListAuditProfilesSortByEnum, bool) {
	enum, ok := mappingListAuditProfilesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
