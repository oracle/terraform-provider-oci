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

// ListUnifiedAuditPoliciesRequest wrapper for the ListUnifiedAuditPolicies operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUnifiedAuditPolicies.go.html to see an example of how to use ListUnifiedAuditPoliciesRequest.
type ListUnifiedAuditPoliciesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// An optional filter to return only resources that match the specified OCID of the security policy resource.
	SecurityPolicyId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyId"`

	// The current state of the Unified Audit policy.
	LifecycleState ListUnifiedAuditPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListUnifiedAuditPoliciesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A boolean flag indicating to list seeded unified audit policies. Set this parameter to get list of seeded unified audit policies.
	IsSeeded *bool `mandatory:"false" contributesTo:"query" name:"isSeeded"`

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

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// An optional filter to return only resources that match the specified OCID of the unified audit policy definition resource.
	UnifiedAuditPolicyDefinitionId *string `mandatory:"false" contributesTo:"query" name:"unifiedAuditPolicyDefinitionId"`

	// An optional filter to return only resources that match the specified OCID of the Unified Audit policy resource.
	UnifiedAuditPolicyId *string `mandatory:"false" contributesTo:"query" name:"unifiedAuditPolicyId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListUnifiedAuditPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListUnifiedAuditPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUnifiedAuditPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUnifiedAuditPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUnifiedAuditPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUnifiedAuditPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUnifiedAuditPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUnifiedAuditPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListUnifiedAuditPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUnifiedAuditPoliciesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListUnifiedAuditPoliciesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUnifiedAuditPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUnifiedAuditPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUnifiedAuditPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUnifiedAuditPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUnifiedAuditPoliciesResponse wrapper for the ListUnifiedAuditPolicies operation
type ListUnifiedAuditPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UnifiedAuditPolicyCollection instances
	UnifiedAuditPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListUnifiedAuditPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUnifiedAuditPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUnifiedAuditPoliciesLifecycleStateEnum Enum with underlying type: string
type ListUnifiedAuditPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListUnifiedAuditPoliciesLifecycleStateEnum
const (
	ListUnifiedAuditPoliciesLifecycleStateCreating       ListUnifiedAuditPoliciesLifecycleStateEnum = "CREATING"
	ListUnifiedAuditPoliciesLifecycleStateUpdating       ListUnifiedAuditPoliciesLifecycleStateEnum = "UPDATING"
	ListUnifiedAuditPoliciesLifecycleStateActive         ListUnifiedAuditPoliciesLifecycleStateEnum = "ACTIVE"
	ListUnifiedAuditPoliciesLifecycleStateInactive       ListUnifiedAuditPoliciesLifecycleStateEnum = "INACTIVE"
	ListUnifiedAuditPoliciesLifecycleStateFailed         ListUnifiedAuditPoliciesLifecycleStateEnum = "FAILED"
	ListUnifiedAuditPoliciesLifecycleStateDeleting       ListUnifiedAuditPoliciesLifecycleStateEnum = "DELETING"
	ListUnifiedAuditPoliciesLifecycleStateNeedsAttention ListUnifiedAuditPoliciesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListUnifiedAuditPoliciesLifecycleStateDeleted        ListUnifiedAuditPoliciesLifecycleStateEnum = "DELETED"
)

var mappingListUnifiedAuditPoliciesLifecycleStateEnum = map[string]ListUnifiedAuditPoliciesLifecycleStateEnum{
	"CREATING":        ListUnifiedAuditPoliciesLifecycleStateCreating,
	"UPDATING":        ListUnifiedAuditPoliciesLifecycleStateUpdating,
	"ACTIVE":          ListUnifiedAuditPoliciesLifecycleStateActive,
	"INACTIVE":        ListUnifiedAuditPoliciesLifecycleStateInactive,
	"FAILED":          ListUnifiedAuditPoliciesLifecycleStateFailed,
	"DELETING":        ListUnifiedAuditPoliciesLifecycleStateDeleting,
	"NEEDS_ATTENTION": ListUnifiedAuditPoliciesLifecycleStateNeedsAttention,
	"DELETED":         ListUnifiedAuditPoliciesLifecycleStateDeleted,
}

var mappingListUnifiedAuditPoliciesLifecycleStateEnumLowerCase = map[string]ListUnifiedAuditPoliciesLifecycleStateEnum{
	"creating":        ListUnifiedAuditPoliciesLifecycleStateCreating,
	"updating":        ListUnifiedAuditPoliciesLifecycleStateUpdating,
	"active":          ListUnifiedAuditPoliciesLifecycleStateActive,
	"inactive":        ListUnifiedAuditPoliciesLifecycleStateInactive,
	"failed":          ListUnifiedAuditPoliciesLifecycleStateFailed,
	"deleting":        ListUnifiedAuditPoliciesLifecycleStateDeleting,
	"needs_attention": ListUnifiedAuditPoliciesLifecycleStateNeedsAttention,
	"deleted":         ListUnifiedAuditPoliciesLifecycleStateDeleted,
}

// GetListUnifiedAuditPoliciesLifecycleStateEnumValues Enumerates the set of values for ListUnifiedAuditPoliciesLifecycleStateEnum
func GetListUnifiedAuditPoliciesLifecycleStateEnumValues() []ListUnifiedAuditPoliciesLifecycleStateEnum {
	values := make([]ListUnifiedAuditPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListUnifiedAuditPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAuditPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListUnifiedAuditPoliciesLifecycleStateEnum
func GetListUnifiedAuditPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"NEEDS_ATTENTION",
		"DELETED",
	}
}

// GetMappingListUnifiedAuditPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAuditPoliciesLifecycleStateEnum(val string) (ListUnifiedAuditPoliciesLifecycleStateEnum, bool) {
	enum, ok := mappingListUnifiedAuditPoliciesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUnifiedAuditPoliciesAccessLevelEnum Enum with underlying type: string
type ListUnifiedAuditPoliciesAccessLevelEnum string

// Set of constants representing the allowable values for ListUnifiedAuditPoliciesAccessLevelEnum
const (
	ListUnifiedAuditPoliciesAccessLevelRestricted ListUnifiedAuditPoliciesAccessLevelEnum = "RESTRICTED"
	ListUnifiedAuditPoliciesAccessLevelAccessible ListUnifiedAuditPoliciesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListUnifiedAuditPoliciesAccessLevelEnum = map[string]ListUnifiedAuditPoliciesAccessLevelEnum{
	"RESTRICTED": ListUnifiedAuditPoliciesAccessLevelRestricted,
	"ACCESSIBLE": ListUnifiedAuditPoliciesAccessLevelAccessible,
}

var mappingListUnifiedAuditPoliciesAccessLevelEnumLowerCase = map[string]ListUnifiedAuditPoliciesAccessLevelEnum{
	"restricted": ListUnifiedAuditPoliciesAccessLevelRestricted,
	"accessible": ListUnifiedAuditPoliciesAccessLevelAccessible,
}

// GetListUnifiedAuditPoliciesAccessLevelEnumValues Enumerates the set of values for ListUnifiedAuditPoliciesAccessLevelEnum
func GetListUnifiedAuditPoliciesAccessLevelEnumValues() []ListUnifiedAuditPoliciesAccessLevelEnum {
	values := make([]ListUnifiedAuditPoliciesAccessLevelEnum, 0)
	for _, v := range mappingListUnifiedAuditPoliciesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAuditPoliciesAccessLevelEnumStringValues Enumerates the set of values in String for ListUnifiedAuditPoliciesAccessLevelEnum
func GetListUnifiedAuditPoliciesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListUnifiedAuditPoliciesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAuditPoliciesAccessLevelEnum(val string) (ListUnifiedAuditPoliciesAccessLevelEnum, bool) {
	enum, ok := mappingListUnifiedAuditPoliciesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUnifiedAuditPoliciesSortOrderEnum Enum with underlying type: string
type ListUnifiedAuditPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListUnifiedAuditPoliciesSortOrderEnum
const (
	ListUnifiedAuditPoliciesSortOrderAsc  ListUnifiedAuditPoliciesSortOrderEnum = "ASC"
	ListUnifiedAuditPoliciesSortOrderDesc ListUnifiedAuditPoliciesSortOrderEnum = "DESC"
)

var mappingListUnifiedAuditPoliciesSortOrderEnum = map[string]ListUnifiedAuditPoliciesSortOrderEnum{
	"ASC":  ListUnifiedAuditPoliciesSortOrderAsc,
	"DESC": ListUnifiedAuditPoliciesSortOrderDesc,
}

var mappingListUnifiedAuditPoliciesSortOrderEnumLowerCase = map[string]ListUnifiedAuditPoliciesSortOrderEnum{
	"asc":  ListUnifiedAuditPoliciesSortOrderAsc,
	"desc": ListUnifiedAuditPoliciesSortOrderDesc,
}

// GetListUnifiedAuditPoliciesSortOrderEnumValues Enumerates the set of values for ListUnifiedAuditPoliciesSortOrderEnum
func GetListUnifiedAuditPoliciesSortOrderEnumValues() []ListUnifiedAuditPoliciesSortOrderEnum {
	values := make([]ListUnifiedAuditPoliciesSortOrderEnum, 0)
	for _, v := range mappingListUnifiedAuditPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAuditPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListUnifiedAuditPoliciesSortOrderEnum
func GetListUnifiedAuditPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUnifiedAuditPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAuditPoliciesSortOrderEnum(val string) (ListUnifiedAuditPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListUnifiedAuditPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUnifiedAuditPoliciesSortByEnum Enum with underlying type: string
type ListUnifiedAuditPoliciesSortByEnum string

// Set of constants representing the allowable values for ListUnifiedAuditPoliciesSortByEnum
const (
	ListUnifiedAuditPoliciesSortByTimecreated ListUnifiedAuditPoliciesSortByEnum = "TIMECREATED"
	ListUnifiedAuditPoliciesSortByDisplayname ListUnifiedAuditPoliciesSortByEnum = "DISPLAYNAME"
)

var mappingListUnifiedAuditPoliciesSortByEnum = map[string]ListUnifiedAuditPoliciesSortByEnum{
	"TIMECREATED": ListUnifiedAuditPoliciesSortByTimecreated,
	"DISPLAYNAME": ListUnifiedAuditPoliciesSortByDisplayname,
}

var mappingListUnifiedAuditPoliciesSortByEnumLowerCase = map[string]ListUnifiedAuditPoliciesSortByEnum{
	"timecreated": ListUnifiedAuditPoliciesSortByTimecreated,
	"displayname": ListUnifiedAuditPoliciesSortByDisplayname,
}

// GetListUnifiedAuditPoliciesSortByEnumValues Enumerates the set of values for ListUnifiedAuditPoliciesSortByEnum
func GetListUnifiedAuditPoliciesSortByEnumValues() []ListUnifiedAuditPoliciesSortByEnum {
	values := make([]ListUnifiedAuditPoliciesSortByEnum, 0)
	for _, v := range mappingListUnifiedAuditPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAuditPoliciesSortByEnumStringValues Enumerates the set of values in String for ListUnifiedAuditPoliciesSortByEnum
func GetListUnifiedAuditPoliciesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListUnifiedAuditPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAuditPoliciesSortByEnum(val string) (ListUnifiedAuditPoliciesSortByEnum, bool) {
	enum, ok := mappingListUnifiedAuditPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
