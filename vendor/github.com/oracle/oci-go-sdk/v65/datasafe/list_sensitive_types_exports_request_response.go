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

// ListSensitiveTypesExportsRequest wrapper for the ListSensitiveTypesExports operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveTypesExports.go.html to see an example of how to use ListSensitiveTypesExportsRequest.
type ListSensitiveTypesExportsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSensitiveTypesExportsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only the resources that match the specified lifecycle state.
	LifecycleState ListSensitiveTypesExportsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the sensitive types export resource.
	SensitiveTypesExportId *string `mandatory:"false" contributesTo:"query" name:"sensitiveTypesExportId"`

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
	SortOrder ListSensitiveTypesExportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListSensitiveTypesExportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSensitiveTypesExportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSensitiveTypesExportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSensitiveTypesExportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSensitiveTypesExportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSensitiveTypesExportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSensitiveTypesExportsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSensitiveTypesExportsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypesExportsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSensitiveTypesExportsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypesExportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSensitiveTypesExportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveTypesExportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSensitiveTypesExportsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSensitiveTypesExportsResponse wrapper for the ListSensitiveTypesExports operation
type ListSensitiveTypesExportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SensitiveTypesExportCollection instances
	SensitiveTypesExportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSensitiveTypesExportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSensitiveTypesExportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSensitiveTypesExportsAccessLevelEnum Enum with underlying type: string
type ListSensitiveTypesExportsAccessLevelEnum string

// Set of constants representing the allowable values for ListSensitiveTypesExportsAccessLevelEnum
const (
	ListSensitiveTypesExportsAccessLevelRestricted ListSensitiveTypesExportsAccessLevelEnum = "RESTRICTED"
	ListSensitiveTypesExportsAccessLevelAccessible ListSensitiveTypesExportsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSensitiveTypesExportsAccessLevelEnum = map[string]ListSensitiveTypesExportsAccessLevelEnum{
	"RESTRICTED": ListSensitiveTypesExportsAccessLevelRestricted,
	"ACCESSIBLE": ListSensitiveTypesExportsAccessLevelAccessible,
}

var mappingListSensitiveTypesExportsAccessLevelEnumLowerCase = map[string]ListSensitiveTypesExportsAccessLevelEnum{
	"restricted": ListSensitiveTypesExportsAccessLevelRestricted,
	"accessible": ListSensitiveTypesExportsAccessLevelAccessible,
}

// GetListSensitiveTypesExportsAccessLevelEnumValues Enumerates the set of values for ListSensitiveTypesExportsAccessLevelEnum
func GetListSensitiveTypesExportsAccessLevelEnumValues() []ListSensitiveTypesExportsAccessLevelEnum {
	values := make([]ListSensitiveTypesExportsAccessLevelEnum, 0)
	for _, v := range mappingListSensitiveTypesExportsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesExportsAccessLevelEnumStringValues Enumerates the set of values in String for ListSensitiveTypesExportsAccessLevelEnum
func GetListSensitiveTypesExportsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSensitiveTypesExportsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesExportsAccessLevelEnum(val string) (ListSensitiveTypesExportsAccessLevelEnum, bool) {
	enum, ok := mappingListSensitiveTypesExportsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypesExportsLifecycleStateEnum Enum with underlying type: string
type ListSensitiveTypesExportsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSensitiveTypesExportsLifecycleStateEnum
const (
	ListSensitiveTypesExportsLifecycleStateCreating ListSensitiveTypesExportsLifecycleStateEnum = "CREATING"
	ListSensitiveTypesExportsLifecycleStateActive   ListSensitiveTypesExportsLifecycleStateEnum = "ACTIVE"
	ListSensitiveTypesExportsLifecycleStateUpdating ListSensitiveTypesExportsLifecycleStateEnum = "UPDATING"
	ListSensitiveTypesExportsLifecycleStateDeleting ListSensitiveTypesExportsLifecycleStateEnum = "DELETING"
	ListSensitiveTypesExportsLifecycleStateDeleted  ListSensitiveTypesExportsLifecycleStateEnum = "DELETED"
	ListSensitiveTypesExportsLifecycleStateFailed   ListSensitiveTypesExportsLifecycleStateEnum = "FAILED"
)

var mappingListSensitiveTypesExportsLifecycleStateEnum = map[string]ListSensitiveTypesExportsLifecycleStateEnum{
	"CREATING": ListSensitiveTypesExportsLifecycleStateCreating,
	"ACTIVE":   ListSensitiveTypesExportsLifecycleStateActive,
	"UPDATING": ListSensitiveTypesExportsLifecycleStateUpdating,
	"DELETING": ListSensitiveTypesExportsLifecycleStateDeleting,
	"DELETED":  ListSensitiveTypesExportsLifecycleStateDeleted,
	"FAILED":   ListSensitiveTypesExportsLifecycleStateFailed,
}

var mappingListSensitiveTypesExportsLifecycleStateEnumLowerCase = map[string]ListSensitiveTypesExportsLifecycleStateEnum{
	"creating": ListSensitiveTypesExportsLifecycleStateCreating,
	"active":   ListSensitiveTypesExportsLifecycleStateActive,
	"updating": ListSensitiveTypesExportsLifecycleStateUpdating,
	"deleting": ListSensitiveTypesExportsLifecycleStateDeleting,
	"deleted":  ListSensitiveTypesExportsLifecycleStateDeleted,
	"failed":   ListSensitiveTypesExportsLifecycleStateFailed,
}

// GetListSensitiveTypesExportsLifecycleStateEnumValues Enumerates the set of values for ListSensitiveTypesExportsLifecycleStateEnum
func GetListSensitiveTypesExportsLifecycleStateEnumValues() []ListSensitiveTypesExportsLifecycleStateEnum {
	values := make([]ListSensitiveTypesExportsLifecycleStateEnum, 0)
	for _, v := range mappingListSensitiveTypesExportsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesExportsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSensitiveTypesExportsLifecycleStateEnum
func GetListSensitiveTypesExportsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSensitiveTypesExportsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesExportsLifecycleStateEnum(val string) (ListSensitiveTypesExportsLifecycleStateEnum, bool) {
	enum, ok := mappingListSensitiveTypesExportsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypesExportsSortOrderEnum Enum with underlying type: string
type ListSensitiveTypesExportsSortOrderEnum string

// Set of constants representing the allowable values for ListSensitiveTypesExportsSortOrderEnum
const (
	ListSensitiveTypesExportsSortOrderAsc  ListSensitiveTypesExportsSortOrderEnum = "ASC"
	ListSensitiveTypesExportsSortOrderDesc ListSensitiveTypesExportsSortOrderEnum = "DESC"
)

var mappingListSensitiveTypesExportsSortOrderEnum = map[string]ListSensitiveTypesExportsSortOrderEnum{
	"ASC":  ListSensitiveTypesExportsSortOrderAsc,
	"DESC": ListSensitiveTypesExportsSortOrderDesc,
}

var mappingListSensitiveTypesExportsSortOrderEnumLowerCase = map[string]ListSensitiveTypesExportsSortOrderEnum{
	"asc":  ListSensitiveTypesExportsSortOrderAsc,
	"desc": ListSensitiveTypesExportsSortOrderDesc,
}

// GetListSensitiveTypesExportsSortOrderEnumValues Enumerates the set of values for ListSensitiveTypesExportsSortOrderEnum
func GetListSensitiveTypesExportsSortOrderEnumValues() []ListSensitiveTypesExportsSortOrderEnum {
	values := make([]ListSensitiveTypesExportsSortOrderEnum, 0)
	for _, v := range mappingListSensitiveTypesExportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesExportsSortOrderEnumStringValues Enumerates the set of values in String for ListSensitiveTypesExportsSortOrderEnum
func GetListSensitiveTypesExportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSensitiveTypesExportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesExportsSortOrderEnum(val string) (ListSensitiveTypesExportsSortOrderEnum, bool) {
	enum, ok := mappingListSensitiveTypesExportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveTypesExportsSortByEnum Enum with underlying type: string
type ListSensitiveTypesExportsSortByEnum string

// Set of constants representing the allowable values for ListSensitiveTypesExportsSortByEnum
const (
	ListSensitiveTypesExportsSortByTimecreated ListSensitiveTypesExportsSortByEnum = "TIMECREATED"
	ListSensitiveTypesExportsSortByDisplayname ListSensitiveTypesExportsSortByEnum = "DISPLAYNAME"
)

var mappingListSensitiveTypesExportsSortByEnum = map[string]ListSensitiveTypesExportsSortByEnum{
	"TIMECREATED": ListSensitiveTypesExportsSortByTimecreated,
	"DISPLAYNAME": ListSensitiveTypesExportsSortByDisplayname,
}

var mappingListSensitiveTypesExportsSortByEnumLowerCase = map[string]ListSensitiveTypesExportsSortByEnum{
	"timecreated": ListSensitiveTypesExportsSortByTimecreated,
	"displayname": ListSensitiveTypesExportsSortByDisplayname,
}

// GetListSensitiveTypesExportsSortByEnumValues Enumerates the set of values for ListSensitiveTypesExportsSortByEnum
func GetListSensitiveTypesExportsSortByEnumValues() []ListSensitiveTypesExportsSortByEnum {
	values := make([]ListSensitiveTypesExportsSortByEnum, 0)
	for _, v := range mappingListSensitiveTypesExportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveTypesExportsSortByEnumStringValues Enumerates the set of values in String for ListSensitiveTypesExportsSortByEnum
func GetListSensitiveTypesExportsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSensitiveTypesExportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveTypesExportsSortByEnum(val string) (ListSensitiveTypesExportsSortByEnum, bool) {
	enum, ok := mappingListSensitiveTypesExportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
