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

// ListLibraryMaskingFormatsRequest wrapper for the ListLibraryMaskingFormats operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListLibraryMaskingFormats.go.html to see an example of how to use ListLibraryMaskingFormatsRequest.
type ListLibraryMaskingFormatsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the specified library masking format OCID.
	LibraryMaskingFormatId *string `mandatory:"false" contributesTo:"query" name:"libraryMaskingFormatId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListLibraryMaskingFormatsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources that match the specified lifecycle states.
	LifecycleState ListLibraryMaskingFormatsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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

	// A filter to return the library masking format resources based on the value of their source attribute.
	LibraryMaskingFormatSource ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum `mandatory:"false" contributesTo:"query" name:"libraryMaskingFormatSource" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListLibraryMaskingFormatsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder). The default order for timeCreated is descending.
	// The default order for displayName is ascending. The displayName sort order is case sensitive.
	SortBy ListLibraryMaskingFormatsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLibraryMaskingFormatsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLibraryMaskingFormatsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLibraryMaskingFormatsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLibraryMaskingFormatsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLibraryMaskingFormatsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLibraryMaskingFormatsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListLibraryMaskingFormatsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLibraryMaskingFormatsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListLibraryMaskingFormatsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum(string(request.LibraryMaskingFormatSource)); !ok && request.LibraryMaskingFormatSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LibraryMaskingFormatSource: %s. Supported values are: %s.", request.LibraryMaskingFormatSource, strings.Join(GetListLibraryMaskingFormatsLibraryMaskingFormatSourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLibraryMaskingFormatsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLibraryMaskingFormatsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLibraryMaskingFormatsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLibraryMaskingFormatsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLibraryMaskingFormatsResponse wrapper for the ListLibraryMaskingFormats operation
type ListLibraryMaskingFormatsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LibraryMaskingFormatCollection instances
	LibraryMaskingFormatCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListLibraryMaskingFormatsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLibraryMaskingFormatsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLibraryMaskingFormatsAccessLevelEnum Enum with underlying type: string
type ListLibraryMaskingFormatsAccessLevelEnum string

// Set of constants representing the allowable values for ListLibraryMaskingFormatsAccessLevelEnum
const (
	ListLibraryMaskingFormatsAccessLevelRestricted ListLibraryMaskingFormatsAccessLevelEnum = "RESTRICTED"
	ListLibraryMaskingFormatsAccessLevelAccessible ListLibraryMaskingFormatsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListLibraryMaskingFormatsAccessLevelEnum = map[string]ListLibraryMaskingFormatsAccessLevelEnum{
	"RESTRICTED": ListLibraryMaskingFormatsAccessLevelRestricted,
	"ACCESSIBLE": ListLibraryMaskingFormatsAccessLevelAccessible,
}

var mappingListLibraryMaskingFormatsAccessLevelEnumLowerCase = map[string]ListLibraryMaskingFormatsAccessLevelEnum{
	"restricted": ListLibraryMaskingFormatsAccessLevelRestricted,
	"accessible": ListLibraryMaskingFormatsAccessLevelAccessible,
}

// GetListLibraryMaskingFormatsAccessLevelEnumValues Enumerates the set of values for ListLibraryMaskingFormatsAccessLevelEnum
func GetListLibraryMaskingFormatsAccessLevelEnumValues() []ListLibraryMaskingFormatsAccessLevelEnum {
	values := make([]ListLibraryMaskingFormatsAccessLevelEnum, 0)
	for _, v := range mappingListLibraryMaskingFormatsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListLibraryMaskingFormatsAccessLevelEnumStringValues Enumerates the set of values in String for ListLibraryMaskingFormatsAccessLevelEnum
func GetListLibraryMaskingFormatsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListLibraryMaskingFormatsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLibraryMaskingFormatsAccessLevelEnum(val string) (ListLibraryMaskingFormatsAccessLevelEnum, bool) {
	enum, ok := mappingListLibraryMaskingFormatsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLibraryMaskingFormatsLifecycleStateEnum Enum with underlying type: string
type ListLibraryMaskingFormatsLifecycleStateEnum string

// Set of constants representing the allowable values for ListLibraryMaskingFormatsLifecycleStateEnum
const (
	ListLibraryMaskingFormatsLifecycleStateCreating       ListLibraryMaskingFormatsLifecycleStateEnum = "CREATING"
	ListLibraryMaskingFormatsLifecycleStateActive         ListLibraryMaskingFormatsLifecycleStateEnum = "ACTIVE"
	ListLibraryMaskingFormatsLifecycleStateUpdating       ListLibraryMaskingFormatsLifecycleStateEnum = "UPDATING"
	ListLibraryMaskingFormatsLifecycleStateDeleting       ListLibraryMaskingFormatsLifecycleStateEnum = "DELETING"
	ListLibraryMaskingFormatsLifecycleStateDeleted        ListLibraryMaskingFormatsLifecycleStateEnum = "DELETED"
	ListLibraryMaskingFormatsLifecycleStateNeedsAttention ListLibraryMaskingFormatsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListLibraryMaskingFormatsLifecycleStateFailed         ListLibraryMaskingFormatsLifecycleStateEnum = "FAILED"
)

var mappingListLibraryMaskingFormatsLifecycleStateEnum = map[string]ListLibraryMaskingFormatsLifecycleStateEnum{
	"CREATING":        ListLibraryMaskingFormatsLifecycleStateCreating,
	"ACTIVE":          ListLibraryMaskingFormatsLifecycleStateActive,
	"UPDATING":        ListLibraryMaskingFormatsLifecycleStateUpdating,
	"DELETING":        ListLibraryMaskingFormatsLifecycleStateDeleting,
	"DELETED":         ListLibraryMaskingFormatsLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListLibraryMaskingFormatsLifecycleStateNeedsAttention,
	"FAILED":          ListLibraryMaskingFormatsLifecycleStateFailed,
}

var mappingListLibraryMaskingFormatsLifecycleStateEnumLowerCase = map[string]ListLibraryMaskingFormatsLifecycleStateEnum{
	"creating":        ListLibraryMaskingFormatsLifecycleStateCreating,
	"active":          ListLibraryMaskingFormatsLifecycleStateActive,
	"updating":        ListLibraryMaskingFormatsLifecycleStateUpdating,
	"deleting":        ListLibraryMaskingFormatsLifecycleStateDeleting,
	"deleted":         ListLibraryMaskingFormatsLifecycleStateDeleted,
	"needs_attention": ListLibraryMaskingFormatsLifecycleStateNeedsAttention,
	"failed":          ListLibraryMaskingFormatsLifecycleStateFailed,
}

// GetListLibraryMaskingFormatsLifecycleStateEnumValues Enumerates the set of values for ListLibraryMaskingFormatsLifecycleStateEnum
func GetListLibraryMaskingFormatsLifecycleStateEnumValues() []ListLibraryMaskingFormatsLifecycleStateEnum {
	values := make([]ListLibraryMaskingFormatsLifecycleStateEnum, 0)
	for _, v := range mappingListLibraryMaskingFormatsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListLibraryMaskingFormatsLifecycleStateEnumStringValues Enumerates the set of values in String for ListLibraryMaskingFormatsLifecycleStateEnum
func GetListLibraryMaskingFormatsLifecycleStateEnumStringValues() []string {
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

// GetMappingListLibraryMaskingFormatsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLibraryMaskingFormatsLifecycleStateEnum(val string) (ListLibraryMaskingFormatsLifecycleStateEnum, bool) {
	enum, ok := mappingListLibraryMaskingFormatsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum Enum with underlying type: string
type ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum string

// Set of constants representing the allowable values for ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum
const (
	ListLibraryMaskingFormatsLibraryMaskingFormatSourceOracle ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum = "ORACLE"
	ListLibraryMaskingFormatsLibraryMaskingFormatSourceUser   ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum = "USER"
)

var mappingListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum = map[string]ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum{
	"ORACLE": ListLibraryMaskingFormatsLibraryMaskingFormatSourceOracle,
	"USER":   ListLibraryMaskingFormatsLibraryMaskingFormatSourceUser,
}

var mappingListLibraryMaskingFormatsLibraryMaskingFormatSourceEnumLowerCase = map[string]ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum{
	"oracle": ListLibraryMaskingFormatsLibraryMaskingFormatSourceOracle,
	"user":   ListLibraryMaskingFormatsLibraryMaskingFormatSourceUser,
}

// GetListLibraryMaskingFormatsLibraryMaskingFormatSourceEnumValues Enumerates the set of values for ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum
func GetListLibraryMaskingFormatsLibraryMaskingFormatSourceEnumValues() []ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum {
	values := make([]ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum, 0)
	for _, v := range mappingListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetListLibraryMaskingFormatsLibraryMaskingFormatSourceEnumStringValues Enumerates the set of values in String for ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum
func GetListLibraryMaskingFormatsLibraryMaskingFormatSourceEnumStringValues() []string {
	return []string{
		"ORACLE",
		"USER",
	}
}

// GetMappingListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum(val string) (ListLibraryMaskingFormatsLibraryMaskingFormatSourceEnum, bool) {
	enum, ok := mappingListLibraryMaskingFormatsLibraryMaskingFormatSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLibraryMaskingFormatsSortOrderEnum Enum with underlying type: string
type ListLibraryMaskingFormatsSortOrderEnum string

// Set of constants representing the allowable values for ListLibraryMaskingFormatsSortOrderEnum
const (
	ListLibraryMaskingFormatsSortOrderAsc  ListLibraryMaskingFormatsSortOrderEnum = "ASC"
	ListLibraryMaskingFormatsSortOrderDesc ListLibraryMaskingFormatsSortOrderEnum = "DESC"
)

var mappingListLibraryMaskingFormatsSortOrderEnum = map[string]ListLibraryMaskingFormatsSortOrderEnum{
	"ASC":  ListLibraryMaskingFormatsSortOrderAsc,
	"DESC": ListLibraryMaskingFormatsSortOrderDesc,
}

var mappingListLibraryMaskingFormatsSortOrderEnumLowerCase = map[string]ListLibraryMaskingFormatsSortOrderEnum{
	"asc":  ListLibraryMaskingFormatsSortOrderAsc,
	"desc": ListLibraryMaskingFormatsSortOrderDesc,
}

// GetListLibraryMaskingFormatsSortOrderEnumValues Enumerates the set of values for ListLibraryMaskingFormatsSortOrderEnum
func GetListLibraryMaskingFormatsSortOrderEnumValues() []ListLibraryMaskingFormatsSortOrderEnum {
	values := make([]ListLibraryMaskingFormatsSortOrderEnum, 0)
	for _, v := range mappingListLibraryMaskingFormatsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLibraryMaskingFormatsSortOrderEnumStringValues Enumerates the set of values in String for ListLibraryMaskingFormatsSortOrderEnum
func GetListLibraryMaskingFormatsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLibraryMaskingFormatsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLibraryMaskingFormatsSortOrderEnum(val string) (ListLibraryMaskingFormatsSortOrderEnum, bool) {
	enum, ok := mappingListLibraryMaskingFormatsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLibraryMaskingFormatsSortByEnum Enum with underlying type: string
type ListLibraryMaskingFormatsSortByEnum string

// Set of constants representing the allowable values for ListLibraryMaskingFormatsSortByEnum
const (
	ListLibraryMaskingFormatsSortByDisplayname ListLibraryMaskingFormatsSortByEnum = "displayName"
	ListLibraryMaskingFormatsSortByTimecreated ListLibraryMaskingFormatsSortByEnum = "timeCreated"
)

var mappingListLibraryMaskingFormatsSortByEnum = map[string]ListLibraryMaskingFormatsSortByEnum{
	"displayName": ListLibraryMaskingFormatsSortByDisplayname,
	"timeCreated": ListLibraryMaskingFormatsSortByTimecreated,
}

var mappingListLibraryMaskingFormatsSortByEnumLowerCase = map[string]ListLibraryMaskingFormatsSortByEnum{
	"displayname": ListLibraryMaskingFormatsSortByDisplayname,
	"timecreated": ListLibraryMaskingFormatsSortByTimecreated,
}

// GetListLibraryMaskingFormatsSortByEnumValues Enumerates the set of values for ListLibraryMaskingFormatsSortByEnum
func GetListLibraryMaskingFormatsSortByEnumValues() []ListLibraryMaskingFormatsSortByEnum {
	values := make([]ListLibraryMaskingFormatsSortByEnum, 0)
	for _, v := range mappingListLibraryMaskingFormatsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLibraryMaskingFormatsSortByEnumStringValues Enumerates the set of values in String for ListLibraryMaskingFormatsSortByEnum
func GetListLibraryMaskingFormatsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListLibraryMaskingFormatsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLibraryMaskingFormatsSortByEnum(val string) (ListLibraryMaskingFormatsSortByEnum, bool) {
	enum, ok := mappingListLibraryMaskingFormatsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
