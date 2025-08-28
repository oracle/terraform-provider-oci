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

// ListUnifiedAuditPolicyDefinitionsRequest wrapper for the ListUnifiedAuditPolicyDefinitions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUnifiedAuditPolicyDefinitions.go.html to see an example of how to use ListUnifiedAuditPolicyDefinitionsRequest.
type ListUnifiedAuditPolicyDefinitionsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The current state of the unified audit policy definition.
	LifecycleState ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the unified audit policy definition resource.
	UnifiedAuditPolicyDefinitionId *string `mandatory:"false" contributesTo:"query" name:"unifiedAuditPolicyDefinitionId"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListUnifiedAuditPolicyDefinitionsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A boolean flag indicating to list seeded unified audit policy definitions. Set this parameter to get list of seeded unified audit policy definitions.
	IsSeeded *bool `mandatory:"false" contributesTo:"query" name:"isSeeded"`

	// The category to which the unified audit policy definition belongs to.
	UnifiedAuditPolicyCategory UnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum `mandatory:"false" contributesTo:"query" name:"unifiedAuditPolicyCategory" omitEmpty:"true"`

	// The name of the unified audit policy.
	UnifiedAuditPolicyName *string `mandatory:"false" contributesTo:"query" name:"unifiedAuditPolicyName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListUnifiedAuditPolicyDefinitionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListUnifiedAuditPolicyDefinitionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUnifiedAuditPolicyDefinitionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUnifiedAuditPolicyDefinitionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUnifiedAuditPolicyDefinitionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUnifiedAuditPolicyDefinitionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUnifiedAuditPolicyDefinitionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUnifiedAuditPolicyDefinitionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListUnifiedAuditPolicyDefinitionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUnifiedAuditPolicyDefinitionsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListUnifiedAuditPolicyDefinitionsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnum(string(request.UnifiedAuditPolicyCategory)); !ok && request.UnifiedAuditPolicyCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UnifiedAuditPolicyCategory: %s. Supported values are: %s.", request.UnifiedAuditPolicyCategory, strings.Join(GetUnifiedAuditPolicyDefinitionAuditPolicyCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUnifiedAuditPolicyDefinitionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUnifiedAuditPolicyDefinitionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUnifiedAuditPolicyDefinitionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUnifiedAuditPolicyDefinitionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUnifiedAuditPolicyDefinitionsResponse wrapper for the ListUnifiedAuditPolicyDefinitions operation
type ListUnifiedAuditPolicyDefinitionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UnifiedAuditPolicyDefinitionCollection instances
	UnifiedAuditPolicyDefinitionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListUnifiedAuditPolicyDefinitionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUnifiedAuditPolicyDefinitionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum Enum with underlying type: string
type ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum
const (
	ListUnifiedAuditPolicyDefinitionsLifecycleStateCreating       ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum = "CREATING"
	ListUnifiedAuditPolicyDefinitionsLifecycleStateUpdating       ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum = "UPDATING"
	ListUnifiedAuditPolicyDefinitionsLifecycleStateActive         ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum = "ACTIVE"
	ListUnifiedAuditPolicyDefinitionsLifecycleStateInactive       ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum = "INACTIVE"
	ListUnifiedAuditPolicyDefinitionsLifecycleStateFailed         ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum = "FAILED"
	ListUnifiedAuditPolicyDefinitionsLifecycleStateDeleting       ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum = "DELETING"
	ListUnifiedAuditPolicyDefinitionsLifecycleStateNeedsAttention ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListUnifiedAuditPolicyDefinitionsLifecycleStateEnum = map[string]ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum{
	"CREATING":        ListUnifiedAuditPolicyDefinitionsLifecycleStateCreating,
	"UPDATING":        ListUnifiedAuditPolicyDefinitionsLifecycleStateUpdating,
	"ACTIVE":          ListUnifiedAuditPolicyDefinitionsLifecycleStateActive,
	"INACTIVE":        ListUnifiedAuditPolicyDefinitionsLifecycleStateInactive,
	"FAILED":          ListUnifiedAuditPolicyDefinitionsLifecycleStateFailed,
	"DELETING":        ListUnifiedAuditPolicyDefinitionsLifecycleStateDeleting,
	"NEEDS_ATTENTION": ListUnifiedAuditPolicyDefinitionsLifecycleStateNeedsAttention,
}

var mappingListUnifiedAuditPolicyDefinitionsLifecycleStateEnumLowerCase = map[string]ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum{
	"creating":        ListUnifiedAuditPolicyDefinitionsLifecycleStateCreating,
	"updating":        ListUnifiedAuditPolicyDefinitionsLifecycleStateUpdating,
	"active":          ListUnifiedAuditPolicyDefinitionsLifecycleStateActive,
	"inactive":        ListUnifiedAuditPolicyDefinitionsLifecycleStateInactive,
	"failed":          ListUnifiedAuditPolicyDefinitionsLifecycleStateFailed,
	"deleting":        ListUnifiedAuditPolicyDefinitionsLifecycleStateDeleting,
	"needs_attention": ListUnifiedAuditPolicyDefinitionsLifecycleStateNeedsAttention,
}

// GetListUnifiedAuditPolicyDefinitionsLifecycleStateEnumValues Enumerates the set of values for ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum
func GetListUnifiedAuditPolicyDefinitionsLifecycleStateEnumValues() []ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum {
	values := make([]ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum, 0)
	for _, v := range mappingListUnifiedAuditPolicyDefinitionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAuditPolicyDefinitionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum
func GetListUnifiedAuditPolicyDefinitionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"FAILED",
		"DELETING",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListUnifiedAuditPolicyDefinitionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAuditPolicyDefinitionsLifecycleStateEnum(val string) (ListUnifiedAuditPolicyDefinitionsLifecycleStateEnum, bool) {
	enum, ok := mappingListUnifiedAuditPolicyDefinitionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUnifiedAuditPolicyDefinitionsAccessLevelEnum Enum with underlying type: string
type ListUnifiedAuditPolicyDefinitionsAccessLevelEnum string

// Set of constants representing the allowable values for ListUnifiedAuditPolicyDefinitionsAccessLevelEnum
const (
	ListUnifiedAuditPolicyDefinitionsAccessLevelRestricted ListUnifiedAuditPolicyDefinitionsAccessLevelEnum = "RESTRICTED"
	ListUnifiedAuditPolicyDefinitionsAccessLevelAccessible ListUnifiedAuditPolicyDefinitionsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListUnifiedAuditPolicyDefinitionsAccessLevelEnum = map[string]ListUnifiedAuditPolicyDefinitionsAccessLevelEnum{
	"RESTRICTED": ListUnifiedAuditPolicyDefinitionsAccessLevelRestricted,
	"ACCESSIBLE": ListUnifiedAuditPolicyDefinitionsAccessLevelAccessible,
}

var mappingListUnifiedAuditPolicyDefinitionsAccessLevelEnumLowerCase = map[string]ListUnifiedAuditPolicyDefinitionsAccessLevelEnum{
	"restricted": ListUnifiedAuditPolicyDefinitionsAccessLevelRestricted,
	"accessible": ListUnifiedAuditPolicyDefinitionsAccessLevelAccessible,
}

// GetListUnifiedAuditPolicyDefinitionsAccessLevelEnumValues Enumerates the set of values for ListUnifiedAuditPolicyDefinitionsAccessLevelEnum
func GetListUnifiedAuditPolicyDefinitionsAccessLevelEnumValues() []ListUnifiedAuditPolicyDefinitionsAccessLevelEnum {
	values := make([]ListUnifiedAuditPolicyDefinitionsAccessLevelEnum, 0)
	for _, v := range mappingListUnifiedAuditPolicyDefinitionsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAuditPolicyDefinitionsAccessLevelEnumStringValues Enumerates the set of values in String for ListUnifiedAuditPolicyDefinitionsAccessLevelEnum
func GetListUnifiedAuditPolicyDefinitionsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListUnifiedAuditPolicyDefinitionsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAuditPolicyDefinitionsAccessLevelEnum(val string) (ListUnifiedAuditPolicyDefinitionsAccessLevelEnum, bool) {
	enum, ok := mappingListUnifiedAuditPolicyDefinitionsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUnifiedAuditPolicyDefinitionsSortOrderEnum Enum with underlying type: string
type ListUnifiedAuditPolicyDefinitionsSortOrderEnum string

// Set of constants representing the allowable values for ListUnifiedAuditPolicyDefinitionsSortOrderEnum
const (
	ListUnifiedAuditPolicyDefinitionsSortOrderAsc  ListUnifiedAuditPolicyDefinitionsSortOrderEnum = "ASC"
	ListUnifiedAuditPolicyDefinitionsSortOrderDesc ListUnifiedAuditPolicyDefinitionsSortOrderEnum = "DESC"
)

var mappingListUnifiedAuditPolicyDefinitionsSortOrderEnum = map[string]ListUnifiedAuditPolicyDefinitionsSortOrderEnum{
	"ASC":  ListUnifiedAuditPolicyDefinitionsSortOrderAsc,
	"DESC": ListUnifiedAuditPolicyDefinitionsSortOrderDesc,
}

var mappingListUnifiedAuditPolicyDefinitionsSortOrderEnumLowerCase = map[string]ListUnifiedAuditPolicyDefinitionsSortOrderEnum{
	"asc":  ListUnifiedAuditPolicyDefinitionsSortOrderAsc,
	"desc": ListUnifiedAuditPolicyDefinitionsSortOrderDesc,
}

// GetListUnifiedAuditPolicyDefinitionsSortOrderEnumValues Enumerates the set of values for ListUnifiedAuditPolicyDefinitionsSortOrderEnum
func GetListUnifiedAuditPolicyDefinitionsSortOrderEnumValues() []ListUnifiedAuditPolicyDefinitionsSortOrderEnum {
	values := make([]ListUnifiedAuditPolicyDefinitionsSortOrderEnum, 0)
	for _, v := range mappingListUnifiedAuditPolicyDefinitionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAuditPolicyDefinitionsSortOrderEnumStringValues Enumerates the set of values in String for ListUnifiedAuditPolicyDefinitionsSortOrderEnum
func GetListUnifiedAuditPolicyDefinitionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUnifiedAuditPolicyDefinitionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAuditPolicyDefinitionsSortOrderEnum(val string) (ListUnifiedAuditPolicyDefinitionsSortOrderEnum, bool) {
	enum, ok := mappingListUnifiedAuditPolicyDefinitionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUnifiedAuditPolicyDefinitionsSortByEnum Enum with underlying type: string
type ListUnifiedAuditPolicyDefinitionsSortByEnum string

// Set of constants representing the allowable values for ListUnifiedAuditPolicyDefinitionsSortByEnum
const (
	ListUnifiedAuditPolicyDefinitionsSortByTimecreated ListUnifiedAuditPolicyDefinitionsSortByEnum = "TIMECREATED"
	ListUnifiedAuditPolicyDefinitionsSortByDisplayname ListUnifiedAuditPolicyDefinitionsSortByEnum = "DISPLAYNAME"
)

var mappingListUnifiedAuditPolicyDefinitionsSortByEnum = map[string]ListUnifiedAuditPolicyDefinitionsSortByEnum{
	"TIMECREATED": ListUnifiedAuditPolicyDefinitionsSortByTimecreated,
	"DISPLAYNAME": ListUnifiedAuditPolicyDefinitionsSortByDisplayname,
}

var mappingListUnifiedAuditPolicyDefinitionsSortByEnumLowerCase = map[string]ListUnifiedAuditPolicyDefinitionsSortByEnum{
	"timecreated": ListUnifiedAuditPolicyDefinitionsSortByTimecreated,
	"displayname": ListUnifiedAuditPolicyDefinitionsSortByDisplayname,
}

// GetListUnifiedAuditPolicyDefinitionsSortByEnumValues Enumerates the set of values for ListUnifiedAuditPolicyDefinitionsSortByEnum
func GetListUnifiedAuditPolicyDefinitionsSortByEnumValues() []ListUnifiedAuditPolicyDefinitionsSortByEnum {
	values := make([]ListUnifiedAuditPolicyDefinitionsSortByEnum, 0)
	for _, v := range mappingListUnifiedAuditPolicyDefinitionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUnifiedAuditPolicyDefinitionsSortByEnumStringValues Enumerates the set of values in String for ListUnifiedAuditPolicyDefinitionsSortByEnum
func GetListUnifiedAuditPolicyDefinitionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListUnifiedAuditPolicyDefinitionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUnifiedAuditPolicyDefinitionsSortByEnum(val string) (ListUnifiedAuditPolicyDefinitionsSortByEnum, bool) {
	enum, ok := mappingListUnifiedAuditPolicyDefinitionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
