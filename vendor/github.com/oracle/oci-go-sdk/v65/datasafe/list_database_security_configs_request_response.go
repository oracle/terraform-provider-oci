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

// ListDatabaseSecurityConfigsRequest wrapper for the ListDatabaseSecurityConfigs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListDatabaseSecurityConfigs.go.html to see an example of how to use ListDatabaseSecurityConfigsRequest.
type ListDatabaseSecurityConfigsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListDatabaseSecurityConfigsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the database security configuration.
	LifecycleState ListDatabaseSecurityConfigsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the database security configuration resource.
	DatabaseSecurityConfigId *string `mandatory:"false" contributesTo:"query" name:"databaseSecurityConfigId"`

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

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDatabaseSecurityConfigsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListDatabaseSecurityConfigsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseSecurityConfigsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseSecurityConfigsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseSecurityConfigsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseSecurityConfigsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDatabaseSecurityConfigsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDatabaseSecurityConfigsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListDatabaseSecurityConfigsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseSecurityConfigsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDatabaseSecurityConfigsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseSecurityConfigsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDatabaseSecurityConfigsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDatabaseSecurityConfigsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDatabaseSecurityConfigsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDatabaseSecurityConfigsResponse wrapper for the ListDatabaseSecurityConfigs operation
type ListDatabaseSecurityConfigsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseSecurityConfigCollection instances
	DatabaseSecurityConfigCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListDatabaseSecurityConfigsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseSecurityConfigsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseSecurityConfigsAccessLevelEnum Enum with underlying type: string
type ListDatabaseSecurityConfigsAccessLevelEnum string

// Set of constants representing the allowable values for ListDatabaseSecurityConfigsAccessLevelEnum
const (
	ListDatabaseSecurityConfigsAccessLevelRestricted ListDatabaseSecurityConfigsAccessLevelEnum = "RESTRICTED"
	ListDatabaseSecurityConfigsAccessLevelAccessible ListDatabaseSecurityConfigsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListDatabaseSecurityConfigsAccessLevelEnum = map[string]ListDatabaseSecurityConfigsAccessLevelEnum{
	"RESTRICTED": ListDatabaseSecurityConfigsAccessLevelRestricted,
	"ACCESSIBLE": ListDatabaseSecurityConfigsAccessLevelAccessible,
}

var mappingListDatabaseSecurityConfigsAccessLevelEnumLowerCase = map[string]ListDatabaseSecurityConfigsAccessLevelEnum{
	"restricted": ListDatabaseSecurityConfigsAccessLevelRestricted,
	"accessible": ListDatabaseSecurityConfigsAccessLevelAccessible,
}

// GetListDatabaseSecurityConfigsAccessLevelEnumValues Enumerates the set of values for ListDatabaseSecurityConfigsAccessLevelEnum
func GetListDatabaseSecurityConfigsAccessLevelEnumValues() []ListDatabaseSecurityConfigsAccessLevelEnum {
	values := make([]ListDatabaseSecurityConfigsAccessLevelEnum, 0)
	for _, v := range mappingListDatabaseSecurityConfigsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseSecurityConfigsAccessLevelEnumStringValues Enumerates the set of values in String for ListDatabaseSecurityConfigsAccessLevelEnum
func GetListDatabaseSecurityConfigsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListDatabaseSecurityConfigsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseSecurityConfigsAccessLevelEnum(val string) (ListDatabaseSecurityConfigsAccessLevelEnum, bool) {
	enum, ok := mappingListDatabaseSecurityConfigsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseSecurityConfigsLifecycleStateEnum Enum with underlying type: string
type ListDatabaseSecurityConfigsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseSecurityConfigsLifecycleStateEnum
const (
	ListDatabaseSecurityConfigsLifecycleStateCreating       ListDatabaseSecurityConfigsLifecycleStateEnum = "CREATING"
	ListDatabaseSecurityConfigsLifecycleStateUpdating       ListDatabaseSecurityConfigsLifecycleStateEnum = "UPDATING"
	ListDatabaseSecurityConfigsLifecycleStateActive         ListDatabaseSecurityConfigsLifecycleStateEnum = "ACTIVE"
	ListDatabaseSecurityConfigsLifecycleStateFailed         ListDatabaseSecurityConfigsLifecycleStateEnum = "FAILED"
	ListDatabaseSecurityConfigsLifecycleStateNeedsAttention ListDatabaseSecurityConfigsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListDatabaseSecurityConfigsLifecycleStateDeleting       ListDatabaseSecurityConfigsLifecycleStateEnum = "DELETING"
	ListDatabaseSecurityConfigsLifecycleStateDeleted        ListDatabaseSecurityConfigsLifecycleStateEnum = "DELETED"
)

var mappingListDatabaseSecurityConfigsLifecycleStateEnum = map[string]ListDatabaseSecurityConfigsLifecycleStateEnum{
	"CREATING":        ListDatabaseSecurityConfigsLifecycleStateCreating,
	"UPDATING":        ListDatabaseSecurityConfigsLifecycleStateUpdating,
	"ACTIVE":          ListDatabaseSecurityConfigsLifecycleStateActive,
	"FAILED":          ListDatabaseSecurityConfigsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDatabaseSecurityConfigsLifecycleStateNeedsAttention,
	"DELETING":        ListDatabaseSecurityConfigsLifecycleStateDeleting,
	"DELETED":         ListDatabaseSecurityConfigsLifecycleStateDeleted,
}

var mappingListDatabaseSecurityConfigsLifecycleStateEnumLowerCase = map[string]ListDatabaseSecurityConfigsLifecycleStateEnum{
	"creating":        ListDatabaseSecurityConfigsLifecycleStateCreating,
	"updating":        ListDatabaseSecurityConfigsLifecycleStateUpdating,
	"active":          ListDatabaseSecurityConfigsLifecycleStateActive,
	"failed":          ListDatabaseSecurityConfigsLifecycleStateFailed,
	"needs_attention": ListDatabaseSecurityConfigsLifecycleStateNeedsAttention,
	"deleting":        ListDatabaseSecurityConfigsLifecycleStateDeleting,
	"deleted":         ListDatabaseSecurityConfigsLifecycleStateDeleted,
}

// GetListDatabaseSecurityConfigsLifecycleStateEnumValues Enumerates the set of values for ListDatabaseSecurityConfigsLifecycleStateEnum
func GetListDatabaseSecurityConfigsLifecycleStateEnumValues() []ListDatabaseSecurityConfigsLifecycleStateEnum {
	values := make([]ListDatabaseSecurityConfigsLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseSecurityConfigsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseSecurityConfigsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDatabaseSecurityConfigsLifecycleStateEnum
func GetListDatabaseSecurityConfigsLifecycleStateEnumStringValues() []string {
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

// GetMappingListDatabaseSecurityConfigsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseSecurityConfigsLifecycleStateEnum(val string) (ListDatabaseSecurityConfigsLifecycleStateEnum, bool) {
	enum, ok := mappingListDatabaseSecurityConfigsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseSecurityConfigsSortOrderEnum Enum with underlying type: string
type ListDatabaseSecurityConfigsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseSecurityConfigsSortOrderEnum
const (
	ListDatabaseSecurityConfigsSortOrderAsc  ListDatabaseSecurityConfigsSortOrderEnum = "ASC"
	ListDatabaseSecurityConfigsSortOrderDesc ListDatabaseSecurityConfigsSortOrderEnum = "DESC"
)

var mappingListDatabaseSecurityConfigsSortOrderEnum = map[string]ListDatabaseSecurityConfigsSortOrderEnum{
	"ASC":  ListDatabaseSecurityConfigsSortOrderAsc,
	"DESC": ListDatabaseSecurityConfigsSortOrderDesc,
}

var mappingListDatabaseSecurityConfigsSortOrderEnumLowerCase = map[string]ListDatabaseSecurityConfigsSortOrderEnum{
	"asc":  ListDatabaseSecurityConfigsSortOrderAsc,
	"desc": ListDatabaseSecurityConfigsSortOrderDesc,
}

// GetListDatabaseSecurityConfigsSortOrderEnumValues Enumerates the set of values for ListDatabaseSecurityConfigsSortOrderEnum
func GetListDatabaseSecurityConfigsSortOrderEnumValues() []ListDatabaseSecurityConfigsSortOrderEnum {
	values := make([]ListDatabaseSecurityConfigsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseSecurityConfigsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseSecurityConfigsSortOrderEnumStringValues Enumerates the set of values in String for ListDatabaseSecurityConfigsSortOrderEnum
func GetListDatabaseSecurityConfigsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDatabaseSecurityConfigsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseSecurityConfigsSortOrderEnum(val string) (ListDatabaseSecurityConfigsSortOrderEnum, bool) {
	enum, ok := mappingListDatabaseSecurityConfigsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDatabaseSecurityConfigsSortByEnum Enum with underlying type: string
type ListDatabaseSecurityConfigsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseSecurityConfigsSortByEnum
const (
	ListDatabaseSecurityConfigsSortByTimecreated ListDatabaseSecurityConfigsSortByEnum = "TIMECREATED"
	ListDatabaseSecurityConfigsSortByDisplayname ListDatabaseSecurityConfigsSortByEnum = "DISPLAYNAME"
)

var mappingListDatabaseSecurityConfigsSortByEnum = map[string]ListDatabaseSecurityConfigsSortByEnum{
	"TIMECREATED": ListDatabaseSecurityConfigsSortByTimecreated,
	"DISPLAYNAME": ListDatabaseSecurityConfigsSortByDisplayname,
}

var mappingListDatabaseSecurityConfigsSortByEnumLowerCase = map[string]ListDatabaseSecurityConfigsSortByEnum{
	"timecreated": ListDatabaseSecurityConfigsSortByTimecreated,
	"displayname": ListDatabaseSecurityConfigsSortByDisplayname,
}

// GetListDatabaseSecurityConfigsSortByEnumValues Enumerates the set of values for ListDatabaseSecurityConfigsSortByEnum
func GetListDatabaseSecurityConfigsSortByEnumValues() []ListDatabaseSecurityConfigsSortByEnum {
	values := make([]ListDatabaseSecurityConfigsSortByEnum, 0)
	for _, v := range mappingListDatabaseSecurityConfigsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDatabaseSecurityConfigsSortByEnumStringValues Enumerates the set of values in String for ListDatabaseSecurityConfigsSortByEnum
func GetListDatabaseSecurityConfigsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListDatabaseSecurityConfigsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDatabaseSecurityConfigsSortByEnum(val string) (ListDatabaseSecurityConfigsSortByEnum, bool) {
	enum, ok := mappingListDatabaseSecurityConfigsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
