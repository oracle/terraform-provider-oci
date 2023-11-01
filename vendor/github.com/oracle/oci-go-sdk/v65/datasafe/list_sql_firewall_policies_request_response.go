// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSqlFirewallPoliciesRequest wrapper for the ListSqlFirewallPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSqlFirewallPolicies.go.html to see an example of how to use ListSqlFirewallPoliciesRequest.
type ListSqlFirewallPoliciesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSqlFirewallPoliciesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// An optional filter to return only resources that match the specified OCID of the security policy resource.
	SecurityPolicyId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyId"`

	// The current state of the SQL firewall policy.
	LifecycleState ListSqlFirewallPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the SQL firewall policy resource.
	SqlFirewallPolicyId *string `mandatory:"false" contributesTo:"query" name:"sqlFirewallPolicyId"`

	// A filter to return only items that match the specified user name.
	DbUserName *string `mandatory:"false" contributesTo:"query" name:"dbUserName"`

	// An optional filter to return only resources that match the specified violation action.
	ViolationAction ListSqlFirewallPoliciesViolationActionEnum `mandatory:"false" contributesTo:"query" name:"violationAction" omitEmpty:"true"`

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
	SortOrder ListSqlFirewallPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListSqlFirewallPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSqlFirewallPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSqlFirewallPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSqlFirewallPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSqlFirewallPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSqlFirewallPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSqlFirewallPoliciesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSqlFirewallPoliciesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSqlFirewallPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallPoliciesViolationActionEnum(string(request.ViolationAction)); !ok && request.ViolationAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ViolationAction: %s. Supported values are: %s.", request.ViolationAction, strings.Join(GetListSqlFirewallPoliciesViolationActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSqlFirewallPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSqlFirewallPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSqlFirewallPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSqlFirewallPoliciesResponse wrapper for the ListSqlFirewallPolicies operation
type ListSqlFirewallPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SqlFirewallPolicyCollection instances
	SqlFirewallPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSqlFirewallPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSqlFirewallPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSqlFirewallPoliciesAccessLevelEnum Enum with underlying type: string
type ListSqlFirewallPoliciesAccessLevelEnum string

// Set of constants representing the allowable values for ListSqlFirewallPoliciesAccessLevelEnum
const (
	ListSqlFirewallPoliciesAccessLevelRestricted ListSqlFirewallPoliciesAccessLevelEnum = "RESTRICTED"
	ListSqlFirewallPoliciesAccessLevelAccessible ListSqlFirewallPoliciesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSqlFirewallPoliciesAccessLevelEnum = map[string]ListSqlFirewallPoliciesAccessLevelEnum{
	"RESTRICTED": ListSqlFirewallPoliciesAccessLevelRestricted,
	"ACCESSIBLE": ListSqlFirewallPoliciesAccessLevelAccessible,
}

var mappingListSqlFirewallPoliciesAccessLevelEnumLowerCase = map[string]ListSqlFirewallPoliciesAccessLevelEnum{
	"restricted": ListSqlFirewallPoliciesAccessLevelRestricted,
	"accessible": ListSqlFirewallPoliciesAccessLevelAccessible,
}

// GetListSqlFirewallPoliciesAccessLevelEnumValues Enumerates the set of values for ListSqlFirewallPoliciesAccessLevelEnum
func GetListSqlFirewallPoliciesAccessLevelEnumValues() []ListSqlFirewallPoliciesAccessLevelEnum {
	values := make([]ListSqlFirewallPoliciesAccessLevelEnum, 0)
	for _, v := range mappingListSqlFirewallPoliciesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallPoliciesAccessLevelEnumStringValues Enumerates the set of values in String for ListSqlFirewallPoliciesAccessLevelEnum
func GetListSqlFirewallPoliciesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSqlFirewallPoliciesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallPoliciesAccessLevelEnum(val string) (ListSqlFirewallPoliciesAccessLevelEnum, bool) {
	enum, ok := mappingListSqlFirewallPoliciesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallPoliciesLifecycleStateEnum Enum with underlying type: string
type ListSqlFirewallPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSqlFirewallPoliciesLifecycleStateEnum
const (
	ListSqlFirewallPoliciesLifecycleStateCreating       ListSqlFirewallPoliciesLifecycleStateEnum = "CREATING"
	ListSqlFirewallPoliciesLifecycleStateUpdating       ListSqlFirewallPoliciesLifecycleStateEnum = "UPDATING"
	ListSqlFirewallPoliciesLifecycleStateActive         ListSqlFirewallPoliciesLifecycleStateEnum = "ACTIVE"
	ListSqlFirewallPoliciesLifecycleStateInactive       ListSqlFirewallPoliciesLifecycleStateEnum = "INACTIVE"
	ListSqlFirewallPoliciesLifecycleStateFailed         ListSqlFirewallPoliciesLifecycleStateEnum = "FAILED"
	ListSqlFirewallPoliciesLifecycleStateDeleting       ListSqlFirewallPoliciesLifecycleStateEnum = "DELETING"
	ListSqlFirewallPoliciesLifecycleStateDeleted        ListSqlFirewallPoliciesLifecycleStateEnum = "DELETED"
	ListSqlFirewallPoliciesLifecycleStateNeedsAttention ListSqlFirewallPoliciesLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListSqlFirewallPoliciesLifecycleStateEnum = map[string]ListSqlFirewallPoliciesLifecycleStateEnum{
	"CREATING":        ListSqlFirewallPoliciesLifecycleStateCreating,
	"UPDATING":        ListSqlFirewallPoliciesLifecycleStateUpdating,
	"ACTIVE":          ListSqlFirewallPoliciesLifecycleStateActive,
	"INACTIVE":        ListSqlFirewallPoliciesLifecycleStateInactive,
	"FAILED":          ListSqlFirewallPoliciesLifecycleStateFailed,
	"DELETING":        ListSqlFirewallPoliciesLifecycleStateDeleting,
	"DELETED":         ListSqlFirewallPoliciesLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListSqlFirewallPoliciesLifecycleStateNeedsAttention,
}

var mappingListSqlFirewallPoliciesLifecycleStateEnumLowerCase = map[string]ListSqlFirewallPoliciesLifecycleStateEnum{
	"creating":        ListSqlFirewallPoliciesLifecycleStateCreating,
	"updating":        ListSqlFirewallPoliciesLifecycleStateUpdating,
	"active":          ListSqlFirewallPoliciesLifecycleStateActive,
	"inactive":        ListSqlFirewallPoliciesLifecycleStateInactive,
	"failed":          ListSqlFirewallPoliciesLifecycleStateFailed,
	"deleting":        ListSqlFirewallPoliciesLifecycleStateDeleting,
	"deleted":         ListSqlFirewallPoliciesLifecycleStateDeleted,
	"needs_attention": ListSqlFirewallPoliciesLifecycleStateNeedsAttention,
}

// GetListSqlFirewallPoliciesLifecycleStateEnumValues Enumerates the set of values for ListSqlFirewallPoliciesLifecycleStateEnum
func GetListSqlFirewallPoliciesLifecycleStateEnumValues() []ListSqlFirewallPoliciesLifecycleStateEnum {
	values := make([]ListSqlFirewallPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListSqlFirewallPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSqlFirewallPoliciesLifecycleStateEnum
func GetListSqlFirewallPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListSqlFirewallPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallPoliciesLifecycleStateEnum(val string) (ListSqlFirewallPoliciesLifecycleStateEnum, bool) {
	enum, ok := mappingListSqlFirewallPoliciesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallPoliciesViolationActionEnum Enum with underlying type: string
type ListSqlFirewallPoliciesViolationActionEnum string

// Set of constants representing the allowable values for ListSqlFirewallPoliciesViolationActionEnum
const (
	ListSqlFirewallPoliciesViolationActionBlock   ListSqlFirewallPoliciesViolationActionEnum = "block"
	ListSqlFirewallPoliciesViolationActionObserve ListSqlFirewallPoliciesViolationActionEnum = "observe"
)

var mappingListSqlFirewallPoliciesViolationActionEnum = map[string]ListSqlFirewallPoliciesViolationActionEnum{
	"block":   ListSqlFirewallPoliciesViolationActionBlock,
	"observe": ListSqlFirewallPoliciesViolationActionObserve,
}

var mappingListSqlFirewallPoliciesViolationActionEnumLowerCase = map[string]ListSqlFirewallPoliciesViolationActionEnum{
	"block":   ListSqlFirewallPoliciesViolationActionBlock,
	"observe": ListSqlFirewallPoliciesViolationActionObserve,
}

// GetListSqlFirewallPoliciesViolationActionEnumValues Enumerates the set of values for ListSqlFirewallPoliciesViolationActionEnum
func GetListSqlFirewallPoliciesViolationActionEnumValues() []ListSqlFirewallPoliciesViolationActionEnum {
	values := make([]ListSqlFirewallPoliciesViolationActionEnum, 0)
	for _, v := range mappingListSqlFirewallPoliciesViolationActionEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallPoliciesViolationActionEnumStringValues Enumerates the set of values in String for ListSqlFirewallPoliciesViolationActionEnum
func GetListSqlFirewallPoliciesViolationActionEnumStringValues() []string {
	return []string{
		"block",
		"observe",
	}
}

// GetMappingListSqlFirewallPoliciesViolationActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallPoliciesViolationActionEnum(val string) (ListSqlFirewallPoliciesViolationActionEnum, bool) {
	enum, ok := mappingListSqlFirewallPoliciesViolationActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallPoliciesSortOrderEnum Enum with underlying type: string
type ListSqlFirewallPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListSqlFirewallPoliciesSortOrderEnum
const (
	ListSqlFirewallPoliciesSortOrderAsc  ListSqlFirewallPoliciesSortOrderEnum = "ASC"
	ListSqlFirewallPoliciesSortOrderDesc ListSqlFirewallPoliciesSortOrderEnum = "DESC"
)

var mappingListSqlFirewallPoliciesSortOrderEnum = map[string]ListSqlFirewallPoliciesSortOrderEnum{
	"ASC":  ListSqlFirewallPoliciesSortOrderAsc,
	"DESC": ListSqlFirewallPoliciesSortOrderDesc,
}

var mappingListSqlFirewallPoliciesSortOrderEnumLowerCase = map[string]ListSqlFirewallPoliciesSortOrderEnum{
	"asc":  ListSqlFirewallPoliciesSortOrderAsc,
	"desc": ListSqlFirewallPoliciesSortOrderDesc,
}

// GetListSqlFirewallPoliciesSortOrderEnumValues Enumerates the set of values for ListSqlFirewallPoliciesSortOrderEnum
func GetListSqlFirewallPoliciesSortOrderEnumValues() []ListSqlFirewallPoliciesSortOrderEnum {
	values := make([]ListSqlFirewallPoliciesSortOrderEnum, 0)
	for _, v := range mappingListSqlFirewallPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListSqlFirewallPoliciesSortOrderEnum
func GetListSqlFirewallPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSqlFirewallPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallPoliciesSortOrderEnum(val string) (ListSqlFirewallPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListSqlFirewallPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSqlFirewallPoliciesSortByEnum Enum with underlying type: string
type ListSqlFirewallPoliciesSortByEnum string

// Set of constants representing the allowable values for ListSqlFirewallPoliciesSortByEnum
const (
	ListSqlFirewallPoliciesSortByTimecreated ListSqlFirewallPoliciesSortByEnum = "TIMECREATED"
	ListSqlFirewallPoliciesSortByDisplayname ListSqlFirewallPoliciesSortByEnum = "DISPLAYNAME"
)

var mappingListSqlFirewallPoliciesSortByEnum = map[string]ListSqlFirewallPoliciesSortByEnum{
	"TIMECREATED": ListSqlFirewallPoliciesSortByTimecreated,
	"DISPLAYNAME": ListSqlFirewallPoliciesSortByDisplayname,
}

var mappingListSqlFirewallPoliciesSortByEnumLowerCase = map[string]ListSqlFirewallPoliciesSortByEnum{
	"timecreated": ListSqlFirewallPoliciesSortByTimecreated,
	"displayname": ListSqlFirewallPoliciesSortByDisplayname,
}

// GetListSqlFirewallPoliciesSortByEnumValues Enumerates the set of values for ListSqlFirewallPoliciesSortByEnum
func GetListSqlFirewallPoliciesSortByEnumValues() []ListSqlFirewallPoliciesSortByEnum {
	values := make([]ListSqlFirewallPoliciesSortByEnum, 0)
	for _, v := range mappingListSqlFirewallPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSqlFirewallPoliciesSortByEnumStringValues Enumerates the set of values in String for ListSqlFirewallPoliciesSortByEnum
func GetListSqlFirewallPoliciesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSqlFirewallPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSqlFirewallPoliciesSortByEnum(val string) (ListSqlFirewallPoliciesSortByEnum, bool) {
	enum, ok := mappingListSqlFirewallPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
