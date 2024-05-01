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

// ListMaskingPolicyHealthReportsRequest wrapper for the ListMaskingPolicyHealthReports operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingPolicyHealthReports.go.html to see an example of how to use ListMaskingPolicyHealthReportsRequest.
type ListMaskingPolicyHealthReportsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the specified masking policy health report OCID.
	MaskingPolicyHealthReportId *string `mandatory:"false" contributesTo:"query" name:"maskingPolicyHealthReportId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListMaskingPolicyHealthReportsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// sort by
	SortBy ListMaskingPolicyHealthReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskingPolicyHealthReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only the resources that match the specified masking policy OCID.
	MaskingPolicyId *string `mandatory:"false" contributesTo:"query" name:"maskingPolicyId"`

	// A filter to return only the resources that match the specified lifecycle states.
	LifecycleState ListMaskingPolicyHealthReportsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskingPolicyHealthReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingPolicyHealthReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingPolicyHealthReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingPolicyHealthReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingPolicyHealthReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskingPolicyHealthReportsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListMaskingPolicyHealthReportsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingPolicyHealthReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskingPolicyHealthReportsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingPolicyHealthReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskingPolicyHealthReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingPolicyHealthReportsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMaskingPolicyHealthReportsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingPolicyHealthReportsResponse wrapper for the ListMaskingPolicyHealthReports operation
type ListMaskingPolicyHealthReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingPolicyHealthReportCollection instances
	MaskingPolicyHealthReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingPolicyHealthReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingPolicyHealthReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingPolicyHealthReportsAccessLevelEnum Enum with underlying type: string
type ListMaskingPolicyHealthReportsAccessLevelEnum string

// Set of constants representing the allowable values for ListMaskingPolicyHealthReportsAccessLevelEnum
const (
	ListMaskingPolicyHealthReportsAccessLevelRestricted ListMaskingPolicyHealthReportsAccessLevelEnum = "RESTRICTED"
	ListMaskingPolicyHealthReportsAccessLevelAccessible ListMaskingPolicyHealthReportsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListMaskingPolicyHealthReportsAccessLevelEnum = map[string]ListMaskingPolicyHealthReportsAccessLevelEnum{
	"RESTRICTED": ListMaskingPolicyHealthReportsAccessLevelRestricted,
	"ACCESSIBLE": ListMaskingPolicyHealthReportsAccessLevelAccessible,
}

var mappingListMaskingPolicyHealthReportsAccessLevelEnumLowerCase = map[string]ListMaskingPolicyHealthReportsAccessLevelEnum{
	"restricted": ListMaskingPolicyHealthReportsAccessLevelRestricted,
	"accessible": ListMaskingPolicyHealthReportsAccessLevelAccessible,
}

// GetListMaskingPolicyHealthReportsAccessLevelEnumValues Enumerates the set of values for ListMaskingPolicyHealthReportsAccessLevelEnum
func GetListMaskingPolicyHealthReportsAccessLevelEnumValues() []ListMaskingPolicyHealthReportsAccessLevelEnum {
	values := make([]ListMaskingPolicyHealthReportsAccessLevelEnum, 0)
	for _, v := range mappingListMaskingPolicyHealthReportsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyHealthReportsAccessLevelEnumStringValues Enumerates the set of values in String for ListMaskingPolicyHealthReportsAccessLevelEnum
func GetListMaskingPolicyHealthReportsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListMaskingPolicyHealthReportsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyHealthReportsAccessLevelEnum(val string) (ListMaskingPolicyHealthReportsAccessLevelEnum, bool) {
	enum, ok := mappingListMaskingPolicyHealthReportsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPolicyHealthReportsSortByEnum Enum with underlying type: string
type ListMaskingPolicyHealthReportsSortByEnum string

// Set of constants representing the allowable values for ListMaskingPolicyHealthReportsSortByEnum
const (
	ListMaskingPolicyHealthReportsSortByDisplayname ListMaskingPolicyHealthReportsSortByEnum = "displayName"
	ListMaskingPolicyHealthReportsSortByTimecreated ListMaskingPolicyHealthReportsSortByEnum = "timeCreated"
)

var mappingListMaskingPolicyHealthReportsSortByEnum = map[string]ListMaskingPolicyHealthReportsSortByEnum{
	"displayName": ListMaskingPolicyHealthReportsSortByDisplayname,
	"timeCreated": ListMaskingPolicyHealthReportsSortByTimecreated,
}

var mappingListMaskingPolicyHealthReportsSortByEnumLowerCase = map[string]ListMaskingPolicyHealthReportsSortByEnum{
	"displayname": ListMaskingPolicyHealthReportsSortByDisplayname,
	"timecreated": ListMaskingPolicyHealthReportsSortByTimecreated,
}

// GetListMaskingPolicyHealthReportsSortByEnumValues Enumerates the set of values for ListMaskingPolicyHealthReportsSortByEnum
func GetListMaskingPolicyHealthReportsSortByEnumValues() []ListMaskingPolicyHealthReportsSortByEnum {
	values := make([]ListMaskingPolicyHealthReportsSortByEnum, 0)
	for _, v := range mappingListMaskingPolicyHealthReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyHealthReportsSortByEnumStringValues Enumerates the set of values in String for ListMaskingPolicyHealthReportsSortByEnum
func GetListMaskingPolicyHealthReportsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListMaskingPolicyHealthReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyHealthReportsSortByEnum(val string) (ListMaskingPolicyHealthReportsSortByEnum, bool) {
	enum, ok := mappingListMaskingPolicyHealthReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPolicyHealthReportsSortOrderEnum Enum with underlying type: string
type ListMaskingPolicyHealthReportsSortOrderEnum string

// Set of constants representing the allowable values for ListMaskingPolicyHealthReportsSortOrderEnum
const (
	ListMaskingPolicyHealthReportsSortOrderAsc  ListMaskingPolicyHealthReportsSortOrderEnum = "ASC"
	ListMaskingPolicyHealthReportsSortOrderDesc ListMaskingPolicyHealthReportsSortOrderEnum = "DESC"
)

var mappingListMaskingPolicyHealthReportsSortOrderEnum = map[string]ListMaskingPolicyHealthReportsSortOrderEnum{
	"ASC":  ListMaskingPolicyHealthReportsSortOrderAsc,
	"DESC": ListMaskingPolicyHealthReportsSortOrderDesc,
}

var mappingListMaskingPolicyHealthReportsSortOrderEnumLowerCase = map[string]ListMaskingPolicyHealthReportsSortOrderEnum{
	"asc":  ListMaskingPolicyHealthReportsSortOrderAsc,
	"desc": ListMaskingPolicyHealthReportsSortOrderDesc,
}

// GetListMaskingPolicyHealthReportsSortOrderEnumValues Enumerates the set of values for ListMaskingPolicyHealthReportsSortOrderEnum
func GetListMaskingPolicyHealthReportsSortOrderEnumValues() []ListMaskingPolicyHealthReportsSortOrderEnum {
	values := make([]ListMaskingPolicyHealthReportsSortOrderEnum, 0)
	for _, v := range mappingListMaskingPolicyHealthReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyHealthReportsSortOrderEnumStringValues Enumerates the set of values in String for ListMaskingPolicyHealthReportsSortOrderEnum
func GetListMaskingPolicyHealthReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskingPolicyHealthReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyHealthReportsSortOrderEnum(val string) (ListMaskingPolicyHealthReportsSortOrderEnum, bool) {
	enum, ok := mappingListMaskingPolicyHealthReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingPolicyHealthReportsLifecycleStateEnum Enum with underlying type: string
type ListMaskingPolicyHealthReportsLifecycleStateEnum string

// Set of constants representing the allowable values for ListMaskingPolicyHealthReportsLifecycleStateEnum
const (
	ListMaskingPolicyHealthReportsLifecycleStateCreating       ListMaskingPolicyHealthReportsLifecycleStateEnum = "CREATING"
	ListMaskingPolicyHealthReportsLifecycleStateActive         ListMaskingPolicyHealthReportsLifecycleStateEnum = "ACTIVE"
	ListMaskingPolicyHealthReportsLifecycleStateUpdating       ListMaskingPolicyHealthReportsLifecycleStateEnum = "UPDATING"
	ListMaskingPolicyHealthReportsLifecycleStateDeleting       ListMaskingPolicyHealthReportsLifecycleStateEnum = "DELETING"
	ListMaskingPolicyHealthReportsLifecycleStateDeleted        ListMaskingPolicyHealthReportsLifecycleStateEnum = "DELETED"
	ListMaskingPolicyHealthReportsLifecycleStateNeedsAttention ListMaskingPolicyHealthReportsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListMaskingPolicyHealthReportsLifecycleStateFailed         ListMaskingPolicyHealthReportsLifecycleStateEnum = "FAILED"
)

var mappingListMaskingPolicyHealthReportsLifecycleStateEnum = map[string]ListMaskingPolicyHealthReportsLifecycleStateEnum{
	"CREATING":        ListMaskingPolicyHealthReportsLifecycleStateCreating,
	"ACTIVE":          ListMaskingPolicyHealthReportsLifecycleStateActive,
	"UPDATING":        ListMaskingPolicyHealthReportsLifecycleStateUpdating,
	"DELETING":        ListMaskingPolicyHealthReportsLifecycleStateDeleting,
	"DELETED":         ListMaskingPolicyHealthReportsLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListMaskingPolicyHealthReportsLifecycleStateNeedsAttention,
	"FAILED":          ListMaskingPolicyHealthReportsLifecycleStateFailed,
}

var mappingListMaskingPolicyHealthReportsLifecycleStateEnumLowerCase = map[string]ListMaskingPolicyHealthReportsLifecycleStateEnum{
	"creating":        ListMaskingPolicyHealthReportsLifecycleStateCreating,
	"active":          ListMaskingPolicyHealthReportsLifecycleStateActive,
	"updating":        ListMaskingPolicyHealthReportsLifecycleStateUpdating,
	"deleting":        ListMaskingPolicyHealthReportsLifecycleStateDeleting,
	"deleted":         ListMaskingPolicyHealthReportsLifecycleStateDeleted,
	"needs_attention": ListMaskingPolicyHealthReportsLifecycleStateNeedsAttention,
	"failed":          ListMaskingPolicyHealthReportsLifecycleStateFailed,
}

// GetListMaskingPolicyHealthReportsLifecycleStateEnumValues Enumerates the set of values for ListMaskingPolicyHealthReportsLifecycleStateEnum
func GetListMaskingPolicyHealthReportsLifecycleStateEnumValues() []ListMaskingPolicyHealthReportsLifecycleStateEnum {
	values := make([]ListMaskingPolicyHealthReportsLifecycleStateEnum, 0)
	for _, v := range mappingListMaskingPolicyHealthReportsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingPolicyHealthReportsLifecycleStateEnumStringValues Enumerates the set of values in String for ListMaskingPolicyHealthReportsLifecycleStateEnum
func GetListMaskingPolicyHealthReportsLifecycleStateEnumStringValues() []string {
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

// GetMappingListMaskingPolicyHealthReportsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingPolicyHealthReportsLifecycleStateEnum(val string) (ListMaskingPolicyHealthReportsLifecycleStateEnum, bool) {
	enum, ok := mappingListMaskingPolicyHealthReportsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
