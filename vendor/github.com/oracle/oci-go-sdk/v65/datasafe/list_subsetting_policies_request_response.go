// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubsettingPoliciesRequest wrapper for the ListSubsettingPolicies operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingPolicies.go.html to see an example of how to use ListSubsettingPoliciesRequest.
type ListSubsettingPoliciesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the specified subsetting policy OCID.
	SubsettingPolicyId *string `mandatory:"false" contributesTo:"query" name:"subsettingPolicyId"`

	// A filter to return only the resources that match the specified masking policy OCID.
	MaskingPolicyId *string `mandatory:"false" contributesTo:"query" name:"maskingPolicyId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources that match the specified lifecycle states.
	LifecycleState ListSubsettingPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeCreated is descending.
	// The default order for displayName is ascending. The displayName sort order is case sensitive.
	SortBy ListSubsettingPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only the resources that match the specified sensitive data model OCID.
	SensitiveDataModelId *string `mandatory:"false" contributesTo:"query" name:"sensitiveDataModelId"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

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

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSubsettingPoliciesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSubsettingPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingPoliciesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingPoliciesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSubsettingPoliciesAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingPoliciesResponse wrapper for the ListSubsettingPolicies operation
type ListSubsettingPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingPolicyCollection instances
	SubsettingPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingPoliciesLifecycleStateEnum Enum with underlying type: string
type ListSubsettingPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSubsettingPoliciesLifecycleStateEnum
const (
	ListSubsettingPoliciesLifecycleStateCreating       ListSubsettingPoliciesLifecycleStateEnum = "CREATING"
	ListSubsettingPoliciesLifecycleStateActive         ListSubsettingPoliciesLifecycleStateEnum = "ACTIVE"
	ListSubsettingPoliciesLifecycleStateUpdating       ListSubsettingPoliciesLifecycleStateEnum = "UPDATING"
	ListSubsettingPoliciesLifecycleStateDeleting       ListSubsettingPoliciesLifecycleStateEnum = "DELETING"
	ListSubsettingPoliciesLifecycleStateDeleted        ListSubsettingPoliciesLifecycleStateEnum = "DELETED"
	ListSubsettingPoliciesLifecycleStateNeedsAttention ListSubsettingPoliciesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListSubsettingPoliciesLifecycleStateFailed         ListSubsettingPoliciesLifecycleStateEnum = "FAILED"
)

var mappingListSubsettingPoliciesLifecycleStateEnum = map[string]ListSubsettingPoliciesLifecycleStateEnum{
	"CREATING":        ListSubsettingPoliciesLifecycleStateCreating,
	"ACTIVE":          ListSubsettingPoliciesLifecycleStateActive,
	"UPDATING":        ListSubsettingPoliciesLifecycleStateUpdating,
	"DELETING":        ListSubsettingPoliciesLifecycleStateDeleting,
	"DELETED":         ListSubsettingPoliciesLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListSubsettingPoliciesLifecycleStateNeedsAttention,
	"FAILED":          ListSubsettingPoliciesLifecycleStateFailed,
}

var mappingListSubsettingPoliciesLifecycleStateEnumLowerCase = map[string]ListSubsettingPoliciesLifecycleStateEnum{
	"creating":        ListSubsettingPoliciesLifecycleStateCreating,
	"active":          ListSubsettingPoliciesLifecycleStateActive,
	"updating":        ListSubsettingPoliciesLifecycleStateUpdating,
	"deleting":        ListSubsettingPoliciesLifecycleStateDeleting,
	"deleted":         ListSubsettingPoliciesLifecycleStateDeleted,
	"needs_attention": ListSubsettingPoliciesLifecycleStateNeedsAttention,
	"failed":          ListSubsettingPoliciesLifecycleStateFailed,
}

// GetListSubsettingPoliciesLifecycleStateEnumValues Enumerates the set of values for ListSubsettingPoliciesLifecycleStateEnum
func GetListSubsettingPoliciesLifecycleStateEnumValues() []ListSubsettingPoliciesLifecycleStateEnum {
	values := make([]ListSubsettingPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListSubsettingPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListSubsettingPoliciesLifecycleStateEnum
func GetListSubsettingPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingListSubsettingPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPoliciesLifecycleStateEnum(val string) (ListSubsettingPoliciesLifecycleStateEnum, bool) {
	enum, ok := mappingListSubsettingPoliciesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingPoliciesSortOrderEnum Enum with underlying type: string
type ListSubsettingPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingPoliciesSortOrderEnum
const (
	ListSubsettingPoliciesSortOrderAsc  ListSubsettingPoliciesSortOrderEnum = "ASC"
	ListSubsettingPoliciesSortOrderDesc ListSubsettingPoliciesSortOrderEnum = "DESC"
)

var mappingListSubsettingPoliciesSortOrderEnum = map[string]ListSubsettingPoliciesSortOrderEnum{
	"ASC":  ListSubsettingPoliciesSortOrderAsc,
	"DESC": ListSubsettingPoliciesSortOrderDesc,
}

var mappingListSubsettingPoliciesSortOrderEnumLowerCase = map[string]ListSubsettingPoliciesSortOrderEnum{
	"asc":  ListSubsettingPoliciesSortOrderAsc,
	"desc": ListSubsettingPoliciesSortOrderDesc,
}

// GetListSubsettingPoliciesSortOrderEnumValues Enumerates the set of values for ListSubsettingPoliciesSortOrderEnum
func GetListSubsettingPoliciesSortOrderEnumValues() []ListSubsettingPoliciesSortOrderEnum {
	values := make([]ListSubsettingPoliciesSortOrderEnum, 0)
	for _, v := range mappingListSubsettingPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingPoliciesSortOrderEnum
func GetListSubsettingPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPoliciesSortOrderEnum(val string) (ListSubsettingPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingPoliciesSortByEnum Enum with underlying type: string
type ListSubsettingPoliciesSortByEnum string

// Set of constants representing the allowable values for ListSubsettingPoliciesSortByEnum
const (
	ListSubsettingPoliciesSortByDisplayname ListSubsettingPoliciesSortByEnum = "displayName"
	ListSubsettingPoliciesSortByTimecreated ListSubsettingPoliciesSortByEnum = "timeCreated"
)

var mappingListSubsettingPoliciesSortByEnum = map[string]ListSubsettingPoliciesSortByEnum{
	"displayName": ListSubsettingPoliciesSortByDisplayname,
	"timeCreated": ListSubsettingPoliciesSortByTimecreated,
}

var mappingListSubsettingPoliciesSortByEnumLowerCase = map[string]ListSubsettingPoliciesSortByEnum{
	"displayname": ListSubsettingPoliciesSortByDisplayname,
	"timecreated": ListSubsettingPoliciesSortByTimecreated,
}

// GetListSubsettingPoliciesSortByEnumValues Enumerates the set of values for ListSubsettingPoliciesSortByEnum
func GetListSubsettingPoliciesSortByEnumValues() []ListSubsettingPoliciesSortByEnum {
	values := make([]ListSubsettingPoliciesSortByEnum, 0)
	for _, v := range mappingListSubsettingPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPoliciesSortByEnumStringValues Enumerates the set of values in String for ListSubsettingPoliciesSortByEnum
func GetListSubsettingPoliciesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListSubsettingPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPoliciesSortByEnum(val string) (ListSubsettingPoliciesSortByEnum, bool) {
	enum, ok := mappingListSubsettingPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingPoliciesAccessLevelEnum Enum with underlying type: string
type ListSubsettingPoliciesAccessLevelEnum string

// Set of constants representing the allowable values for ListSubsettingPoliciesAccessLevelEnum
const (
	ListSubsettingPoliciesAccessLevelRestricted ListSubsettingPoliciesAccessLevelEnum = "RESTRICTED"
	ListSubsettingPoliciesAccessLevelAccessible ListSubsettingPoliciesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSubsettingPoliciesAccessLevelEnum = map[string]ListSubsettingPoliciesAccessLevelEnum{
	"RESTRICTED": ListSubsettingPoliciesAccessLevelRestricted,
	"ACCESSIBLE": ListSubsettingPoliciesAccessLevelAccessible,
}

var mappingListSubsettingPoliciesAccessLevelEnumLowerCase = map[string]ListSubsettingPoliciesAccessLevelEnum{
	"restricted": ListSubsettingPoliciesAccessLevelRestricted,
	"accessible": ListSubsettingPoliciesAccessLevelAccessible,
}

// GetListSubsettingPoliciesAccessLevelEnumValues Enumerates the set of values for ListSubsettingPoliciesAccessLevelEnum
func GetListSubsettingPoliciesAccessLevelEnumValues() []ListSubsettingPoliciesAccessLevelEnum {
	values := make([]ListSubsettingPoliciesAccessLevelEnum, 0)
	for _, v := range mappingListSubsettingPoliciesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPoliciesAccessLevelEnumStringValues Enumerates the set of values in String for ListSubsettingPoliciesAccessLevelEnum
func GetListSubsettingPoliciesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSubsettingPoliciesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPoliciesAccessLevelEnum(val string) (ListSubsettingPoliciesAccessLevelEnum, bool) {
	enum, ok := mappingListSubsettingPoliciesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
