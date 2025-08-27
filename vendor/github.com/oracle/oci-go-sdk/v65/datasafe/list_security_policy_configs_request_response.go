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

// ListSecurityPolicyConfigsRequest wrapper for the ListSecurityPolicyConfigs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityPolicyConfigs.go.html to see an example of how to use ListSecurityPolicyConfigsRequest.
type ListSecurityPolicyConfigsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// An optional filter to return only resources that match the specified OCID of the security policy configuration resource.
	SecurityPolicyConfigId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyConfigId"`

	// An optional filter to return only resources that match the specified OCID of the security policy resource.
	SecurityPolicyId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSecurityPolicyConfigsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The current state of the security policy configuration resource.
	LifecycleState ListSecurityPolicyConfigsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSecurityPolicyConfigsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListSecurityPolicyConfigsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityPolicyConfigsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityPolicyConfigsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityPolicyConfigsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityPolicyConfigsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityPolicyConfigsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityPolicyConfigsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSecurityPolicyConfigsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPolicyConfigsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSecurityPolicyConfigsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPolicyConfigsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecurityPolicyConfigsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPolicyConfigsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecurityPolicyConfigsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityPolicyConfigsResponse wrapper for the ListSecurityPolicyConfigs operation
type ListSecurityPolicyConfigsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SecurityPolicyConfigCollection instances
	SecurityPolicyConfigCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSecurityPolicyConfigsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityPolicyConfigsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityPolicyConfigsAccessLevelEnum Enum with underlying type: string
type ListSecurityPolicyConfigsAccessLevelEnum string

// Set of constants representing the allowable values for ListSecurityPolicyConfigsAccessLevelEnum
const (
	ListSecurityPolicyConfigsAccessLevelRestricted ListSecurityPolicyConfigsAccessLevelEnum = "RESTRICTED"
	ListSecurityPolicyConfigsAccessLevelAccessible ListSecurityPolicyConfigsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSecurityPolicyConfigsAccessLevelEnum = map[string]ListSecurityPolicyConfigsAccessLevelEnum{
	"RESTRICTED": ListSecurityPolicyConfigsAccessLevelRestricted,
	"ACCESSIBLE": ListSecurityPolicyConfigsAccessLevelAccessible,
}

var mappingListSecurityPolicyConfigsAccessLevelEnumLowerCase = map[string]ListSecurityPolicyConfigsAccessLevelEnum{
	"restricted": ListSecurityPolicyConfigsAccessLevelRestricted,
	"accessible": ListSecurityPolicyConfigsAccessLevelAccessible,
}

// GetListSecurityPolicyConfigsAccessLevelEnumValues Enumerates the set of values for ListSecurityPolicyConfigsAccessLevelEnum
func GetListSecurityPolicyConfigsAccessLevelEnumValues() []ListSecurityPolicyConfigsAccessLevelEnum {
	values := make([]ListSecurityPolicyConfigsAccessLevelEnum, 0)
	for _, v := range mappingListSecurityPolicyConfigsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyConfigsAccessLevelEnumStringValues Enumerates the set of values in String for ListSecurityPolicyConfigsAccessLevelEnum
func GetListSecurityPolicyConfigsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSecurityPolicyConfigsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyConfigsAccessLevelEnum(val string) (ListSecurityPolicyConfigsAccessLevelEnum, bool) {
	enum, ok := mappingListSecurityPolicyConfigsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPolicyConfigsLifecycleStateEnum Enum with underlying type: string
type ListSecurityPolicyConfigsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSecurityPolicyConfigsLifecycleStateEnum
const (
	ListSecurityPolicyConfigsLifecycleStateCreating       ListSecurityPolicyConfigsLifecycleStateEnum = "CREATING"
	ListSecurityPolicyConfigsLifecycleStateUpdating       ListSecurityPolicyConfigsLifecycleStateEnum = "UPDATING"
	ListSecurityPolicyConfigsLifecycleStateActive         ListSecurityPolicyConfigsLifecycleStateEnum = "ACTIVE"
	ListSecurityPolicyConfigsLifecycleStateFailed         ListSecurityPolicyConfigsLifecycleStateEnum = "FAILED"
	ListSecurityPolicyConfigsLifecycleStateNeedsAttention ListSecurityPolicyConfigsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListSecurityPolicyConfigsLifecycleStateDeleting       ListSecurityPolicyConfigsLifecycleStateEnum = "DELETING"
	ListSecurityPolicyConfigsLifecycleStateDeleted        ListSecurityPolicyConfigsLifecycleStateEnum = "DELETED"
)

var mappingListSecurityPolicyConfigsLifecycleStateEnum = map[string]ListSecurityPolicyConfigsLifecycleStateEnum{
	"CREATING":        ListSecurityPolicyConfigsLifecycleStateCreating,
	"UPDATING":        ListSecurityPolicyConfigsLifecycleStateUpdating,
	"ACTIVE":          ListSecurityPolicyConfigsLifecycleStateActive,
	"FAILED":          ListSecurityPolicyConfigsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListSecurityPolicyConfigsLifecycleStateNeedsAttention,
	"DELETING":        ListSecurityPolicyConfigsLifecycleStateDeleting,
	"DELETED":         ListSecurityPolicyConfigsLifecycleStateDeleted,
}

var mappingListSecurityPolicyConfigsLifecycleStateEnumLowerCase = map[string]ListSecurityPolicyConfigsLifecycleStateEnum{
	"creating":        ListSecurityPolicyConfigsLifecycleStateCreating,
	"updating":        ListSecurityPolicyConfigsLifecycleStateUpdating,
	"active":          ListSecurityPolicyConfigsLifecycleStateActive,
	"failed":          ListSecurityPolicyConfigsLifecycleStateFailed,
	"needs_attention": ListSecurityPolicyConfigsLifecycleStateNeedsAttention,
	"deleting":        ListSecurityPolicyConfigsLifecycleStateDeleting,
	"deleted":         ListSecurityPolicyConfigsLifecycleStateDeleted,
}

// GetListSecurityPolicyConfigsLifecycleStateEnumValues Enumerates the set of values for ListSecurityPolicyConfigsLifecycleStateEnum
func GetListSecurityPolicyConfigsLifecycleStateEnumValues() []ListSecurityPolicyConfigsLifecycleStateEnum {
	values := make([]ListSecurityPolicyConfigsLifecycleStateEnum, 0)
	for _, v := range mappingListSecurityPolicyConfigsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyConfigsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSecurityPolicyConfigsLifecycleStateEnum
func GetListSecurityPolicyConfigsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"FAILED",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
	}
}

// GetMappingListSecurityPolicyConfigsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyConfigsLifecycleStateEnum(val string) (ListSecurityPolicyConfigsLifecycleStateEnum, bool) {
	enum, ok := mappingListSecurityPolicyConfigsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPolicyConfigsSortOrderEnum Enum with underlying type: string
type ListSecurityPolicyConfigsSortOrderEnum string

// Set of constants representing the allowable values for ListSecurityPolicyConfigsSortOrderEnum
const (
	ListSecurityPolicyConfigsSortOrderAsc  ListSecurityPolicyConfigsSortOrderEnum = "ASC"
	ListSecurityPolicyConfigsSortOrderDesc ListSecurityPolicyConfigsSortOrderEnum = "DESC"
)

var mappingListSecurityPolicyConfigsSortOrderEnum = map[string]ListSecurityPolicyConfigsSortOrderEnum{
	"ASC":  ListSecurityPolicyConfigsSortOrderAsc,
	"DESC": ListSecurityPolicyConfigsSortOrderDesc,
}

var mappingListSecurityPolicyConfigsSortOrderEnumLowerCase = map[string]ListSecurityPolicyConfigsSortOrderEnum{
	"asc":  ListSecurityPolicyConfigsSortOrderAsc,
	"desc": ListSecurityPolicyConfigsSortOrderDesc,
}

// GetListSecurityPolicyConfigsSortOrderEnumValues Enumerates the set of values for ListSecurityPolicyConfigsSortOrderEnum
func GetListSecurityPolicyConfigsSortOrderEnumValues() []ListSecurityPolicyConfigsSortOrderEnum {
	values := make([]ListSecurityPolicyConfigsSortOrderEnum, 0)
	for _, v := range mappingListSecurityPolicyConfigsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyConfigsSortOrderEnumStringValues Enumerates the set of values in String for ListSecurityPolicyConfigsSortOrderEnum
func GetListSecurityPolicyConfigsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSecurityPolicyConfigsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyConfigsSortOrderEnum(val string) (ListSecurityPolicyConfigsSortOrderEnum, bool) {
	enum, ok := mappingListSecurityPolicyConfigsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPolicyConfigsSortByEnum Enum with underlying type: string
type ListSecurityPolicyConfigsSortByEnum string

// Set of constants representing the allowable values for ListSecurityPolicyConfigsSortByEnum
const (
	ListSecurityPolicyConfigsSortByTimecreated ListSecurityPolicyConfigsSortByEnum = "TIMECREATED"
	ListSecurityPolicyConfigsSortByDisplayname ListSecurityPolicyConfigsSortByEnum = "DISPLAYNAME"
)

var mappingListSecurityPolicyConfigsSortByEnum = map[string]ListSecurityPolicyConfigsSortByEnum{
	"TIMECREATED": ListSecurityPolicyConfigsSortByTimecreated,
	"DISPLAYNAME": ListSecurityPolicyConfigsSortByDisplayname,
}

var mappingListSecurityPolicyConfigsSortByEnumLowerCase = map[string]ListSecurityPolicyConfigsSortByEnum{
	"timecreated": ListSecurityPolicyConfigsSortByTimecreated,
	"displayname": ListSecurityPolicyConfigsSortByDisplayname,
}

// GetListSecurityPolicyConfigsSortByEnumValues Enumerates the set of values for ListSecurityPolicyConfigsSortByEnum
func GetListSecurityPolicyConfigsSortByEnumValues() []ListSecurityPolicyConfigsSortByEnum {
	values := make([]ListSecurityPolicyConfigsSortByEnum, 0)
	for _, v := range mappingListSecurityPolicyConfigsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyConfigsSortByEnumStringValues Enumerates the set of values in String for ListSecurityPolicyConfigsSortByEnum
func GetListSecurityPolicyConfigsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSecurityPolicyConfigsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyConfigsSortByEnum(val string) (ListSecurityPolicyConfigsSortByEnum, bool) {
	enum, ok := mappingListSecurityPolicyConfigsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
