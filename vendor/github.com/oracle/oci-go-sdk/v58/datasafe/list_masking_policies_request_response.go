// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListMaskingPoliciesRequest wrapper for the ListMaskingPolicies operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingPolicies.go.html to see an example of how to use ListMaskingPoliciesRequest.
type ListMaskingPoliciesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the specified masking policy OCID.
	MaskingPolicyId *string `mandatory:"false" contributesTo:"query" name:"maskingPolicyId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources that match the specified lifecycle states.
	LifecycleState ListMaskingPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskingPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order for timeCreated is descending.
	// The default order for displayName is ascending. The displayName sort order is case sensitive.
	SortBy ListMaskingPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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
	AccessLevel ListMaskingPoliciesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskingPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskingPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMaskingPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskingPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskingPoliciesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingPoliciesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListMaskingPoliciesAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingPoliciesResponse wrapper for the ListMaskingPolicies operation
type ListMaskingPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingPolicyCollection instances
	MaskingPolicyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingPoliciesLifecycleStateEnum Enum with underlying type: string
type ListMaskingPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListMaskingPoliciesLifecycleStateEnum
const (
	ListMaskingPoliciesLifecycleStateCreating       ListMaskingPoliciesLifecycleStateEnum = "CREATING"
	ListMaskingPoliciesLifecycleStateActive         ListMaskingPoliciesLifecycleStateEnum = "ACTIVE"
	ListMaskingPoliciesLifecycleStateUpdating       ListMaskingPoliciesLifecycleStateEnum = "UPDATING"
	ListMaskingPoliciesLifecycleStateDeleting       ListMaskingPoliciesLifecycleStateEnum = "DELETING"
	ListMaskingPoliciesLifecycleStateDeleted        ListMaskingPoliciesLifecycleStateEnum = "DELETED"
	ListMaskingPoliciesLifecycleStateNeedsAttention ListMaskingPoliciesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListMaskingPoliciesLifecycleStateFailed         ListMaskingPoliciesLifecycleStateEnum = "FAILED"
)

var mappingListMaskingPoliciesLifecycleStateEnum = map[string]ListMaskingPoliciesLifecycleStateEnum{
	"CREATING":        ListMaskingPoliciesLifecycleStateCreating,
	"ACTIVE":          ListMaskingPoliciesLifecycleStateActive,
	"UPDATING":        ListMaskingPoliciesLifecycleStateUpdating,
	"DELETING":        ListMaskingPoliciesLifecycleStateDeleting,
	"DELETED":         ListMaskingPoliciesLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListMaskingPoliciesLifecycleStateNeedsAttention,
	"FAILED":          ListMaskingPoliciesLifecycleStateFailed,
}

// GetListMaskingPoliciesLifecycleStateEnumValues Enumerates the set of values for ListMaskingPoliciesLifecycleStateEnum
func GetListMaskingPoliciesLifecycleStateEnumValues() []ListMaskingPoliciesLifecycleStateEnum {
	values := make([]ListMaskingPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListMaskingPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListMaskingPoliciesLifecycleStateEnum
func GetListMaskingPoliciesLifecycleStateEnumStringValues() []string {
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

// GetMappingListMaskingPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPoliciesLifecycleStateEnum(val string) (ListMaskingPoliciesLifecycleStateEnum, bool) {
	mappingListMaskingPoliciesLifecycleStateEnumIgnoreCase := make(map[string]ListMaskingPoliciesLifecycleStateEnum)
	for k, v := range mappingListMaskingPoliciesLifecycleStateEnum {
		mappingListMaskingPoliciesLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMaskingPoliciesLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPoliciesSortOrderEnum Enum with underlying type: string
type ListMaskingPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListMaskingPoliciesSortOrderEnum
const (
	ListMaskingPoliciesSortOrderAsc  ListMaskingPoliciesSortOrderEnum = "ASC"
	ListMaskingPoliciesSortOrderDesc ListMaskingPoliciesSortOrderEnum = "DESC"
)

var mappingListMaskingPoliciesSortOrderEnum = map[string]ListMaskingPoliciesSortOrderEnum{
	"ASC":  ListMaskingPoliciesSortOrderAsc,
	"DESC": ListMaskingPoliciesSortOrderDesc,
}

// GetListMaskingPoliciesSortOrderEnumValues Enumerates the set of values for ListMaskingPoliciesSortOrderEnum
func GetListMaskingPoliciesSortOrderEnumValues() []ListMaskingPoliciesSortOrderEnum {
	values := make([]ListMaskingPoliciesSortOrderEnum, 0)
	for _, v := range mappingListMaskingPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListMaskingPoliciesSortOrderEnum
func GetListMaskingPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskingPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPoliciesSortOrderEnum(val string) (ListMaskingPoliciesSortOrderEnum, bool) {
	mappingListMaskingPoliciesSortOrderEnumIgnoreCase := make(map[string]ListMaskingPoliciesSortOrderEnum)
	for k, v := range mappingListMaskingPoliciesSortOrderEnum {
		mappingListMaskingPoliciesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMaskingPoliciesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPoliciesSortByEnum Enum with underlying type: string
type ListMaskingPoliciesSortByEnum string

// Set of constants representing the allowable values for ListMaskingPoliciesSortByEnum
const (
	ListMaskingPoliciesSortByDisplayname ListMaskingPoliciesSortByEnum = "displayName"
	ListMaskingPoliciesSortByTimecreated ListMaskingPoliciesSortByEnum = "timeCreated"
)

var mappingListMaskingPoliciesSortByEnum = map[string]ListMaskingPoliciesSortByEnum{
	"displayName": ListMaskingPoliciesSortByDisplayname,
	"timeCreated": ListMaskingPoliciesSortByTimecreated,
}

// GetListMaskingPoliciesSortByEnumValues Enumerates the set of values for ListMaskingPoliciesSortByEnum
func GetListMaskingPoliciesSortByEnumValues() []ListMaskingPoliciesSortByEnum {
	values := make([]ListMaskingPoliciesSortByEnum, 0)
	for _, v := range mappingListMaskingPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPoliciesSortByEnumStringValues Enumerates the set of values in String for ListMaskingPoliciesSortByEnum
func GetListMaskingPoliciesSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListMaskingPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPoliciesSortByEnum(val string) (ListMaskingPoliciesSortByEnum, bool) {
	mappingListMaskingPoliciesSortByEnumIgnoreCase := make(map[string]ListMaskingPoliciesSortByEnum)
	for k, v := range mappingListMaskingPoliciesSortByEnum {
		mappingListMaskingPoliciesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMaskingPoliciesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPoliciesAccessLevelEnum Enum with underlying type: string
type ListMaskingPoliciesAccessLevelEnum string

// Set of constants representing the allowable values for ListMaskingPoliciesAccessLevelEnum
const (
	ListMaskingPoliciesAccessLevelRestricted ListMaskingPoliciesAccessLevelEnum = "RESTRICTED"
	ListMaskingPoliciesAccessLevelAccessible ListMaskingPoliciesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListMaskingPoliciesAccessLevelEnum = map[string]ListMaskingPoliciesAccessLevelEnum{
	"RESTRICTED": ListMaskingPoliciesAccessLevelRestricted,
	"ACCESSIBLE": ListMaskingPoliciesAccessLevelAccessible,
}

// GetListMaskingPoliciesAccessLevelEnumValues Enumerates the set of values for ListMaskingPoliciesAccessLevelEnum
func GetListMaskingPoliciesAccessLevelEnumValues() []ListMaskingPoliciesAccessLevelEnum {
	values := make([]ListMaskingPoliciesAccessLevelEnum, 0)
	for _, v := range mappingListMaskingPoliciesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPoliciesAccessLevelEnumStringValues Enumerates the set of values in String for ListMaskingPoliciesAccessLevelEnum
func GetListMaskingPoliciesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListMaskingPoliciesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPoliciesAccessLevelEnum(val string) (ListMaskingPoliciesAccessLevelEnum, bool) {
	mappingListMaskingPoliciesAccessLevelEnumIgnoreCase := make(map[string]ListMaskingPoliciesAccessLevelEnum)
	for k, v := range mappingListMaskingPoliciesAccessLevelEnum {
		mappingListMaskingPoliciesAccessLevelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMaskingPoliciesAccessLevelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
