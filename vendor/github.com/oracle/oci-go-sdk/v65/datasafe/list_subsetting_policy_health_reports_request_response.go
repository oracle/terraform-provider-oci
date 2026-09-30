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

// ListSubsettingPolicyHealthReportsRequest wrapper for the ListSubsettingPolicyHealthReports operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingPolicyHealthReports.go.html to see an example of how to use ListSubsettingPolicyHealthReportsRequest.
type ListSubsettingPolicyHealthReportsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the specified subsetting policy health report OCID.
	SubsettingPolicyHealthReportId *string `mandatory:"false" contributesTo:"query" name:"subsettingPolicyHealthReportId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
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
	AccessLevel ListSubsettingPolicyHealthReportsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// sort by
	SortBy ListSubsettingPolicyHealthReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingPolicyHealthReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only the resources that match the specified subsetting policy OCID.
	SubsettingPolicyId *string `mandatory:"false" contributesTo:"query" name:"subsettingPolicyId"`

	// A filter to return only the resources that match the specified lifecycle states.
	LifecycleState ListSubsettingPolicyHealthReportsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingPolicyHealthReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingPolicyHealthReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingPolicyHealthReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingPolicyHealthReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingPolicyHealthReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingPolicyHealthReportsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSubsettingPolicyHealthReportsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingPolicyHealthReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingPolicyHealthReportsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingPolicyHealthReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingPolicyHealthReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingPolicyHealthReportsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSubsettingPolicyHealthReportsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingPolicyHealthReportsResponse wrapper for the ListSubsettingPolicyHealthReports operation
type ListSubsettingPolicyHealthReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingPolicyHealthReportCollection instances
	SubsettingPolicyHealthReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingPolicyHealthReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingPolicyHealthReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingPolicyHealthReportsAccessLevelEnum Enum with underlying type: string
type ListSubsettingPolicyHealthReportsAccessLevelEnum string

// Set of constants representing the allowable values for ListSubsettingPolicyHealthReportsAccessLevelEnum
const (
	ListSubsettingPolicyHealthReportsAccessLevelRestricted ListSubsettingPolicyHealthReportsAccessLevelEnum = "RESTRICTED"
	ListSubsettingPolicyHealthReportsAccessLevelAccessible ListSubsettingPolicyHealthReportsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSubsettingPolicyHealthReportsAccessLevelEnum = map[string]ListSubsettingPolicyHealthReportsAccessLevelEnum{
	"RESTRICTED": ListSubsettingPolicyHealthReportsAccessLevelRestricted,
	"ACCESSIBLE": ListSubsettingPolicyHealthReportsAccessLevelAccessible,
}

var mappingListSubsettingPolicyHealthReportsAccessLevelEnumLowerCase = map[string]ListSubsettingPolicyHealthReportsAccessLevelEnum{
	"restricted": ListSubsettingPolicyHealthReportsAccessLevelRestricted,
	"accessible": ListSubsettingPolicyHealthReportsAccessLevelAccessible,
}

// GetListSubsettingPolicyHealthReportsAccessLevelEnumValues Enumerates the set of values for ListSubsettingPolicyHealthReportsAccessLevelEnum
func GetListSubsettingPolicyHealthReportsAccessLevelEnumValues() []ListSubsettingPolicyHealthReportsAccessLevelEnum {
	values := make([]ListSubsettingPolicyHealthReportsAccessLevelEnum, 0)
	for _, v := range mappingListSubsettingPolicyHealthReportsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPolicyHealthReportsAccessLevelEnumStringValues Enumerates the set of values in String for ListSubsettingPolicyHealthReportsAccessLevelEnum
func GetListSubsettingPolicyHealthReportsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSubsettingPolicyHealthReportsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPolicyHealthReportsAccessLevelEnum(val string) (ListSubsettingPolicyHealthReportsAccessLevelEnum, bool) {
	enum, ok := mappingListSubsettingPolicyHealthReportsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingPolicyHealthReportsSortByEnum Enum with underlying type: string
type ListSubsettingPolicyHealthReportsSortByEnum string

// Set of constants representing the allowable values for ListSubsettingPolicyHealthReportsSortByEnum
const (
	ListSubsettingPolicyHealthReportsSortByDisplayname ListSubsettingPolicyHealthReportsSortByEnum = "displayName"
	ListSubsettingPolicyHealthReportsSortByTimecreated ListSubsettingPolicyHealthReportsSortByEnum = "timeCreated"
)

var mappingListSubsettingPolicyHealthReportsSortByEnum = map[string]ListSubsettingPolicyHealthReportsSortByEnum{
	"displayName": ListSubsettingPolicyHealthReportsSortByDisplayname,
	"timeCreated": ListSubsettingPolicyHealthReportsSortByTimecreated,
}

var mappingListSubsettingPolicyHealthReportsSortByEnumLowerCase = map[string]ListSubsettingPolicyHealthReportsSortByEnum{
	"displayname": ListSubsettingPolicyHealthReportsSortByDisplayname,
	"timecreated": ListSubsettingPolicyHealthReportsSortByTimecreated,
}

// GetListSubsettingPolicyHealthReportsSortByEnumValues Enumerates the set of values for ListSubsettingPolicyHealthReportsSortByEnum
func GetListSubsettingPolicyHealthReportsSortByEnumValues() []ListSubsettingPolicyHealthReportsSortByEnum {
	values := make([]ListSubsettingPolicyHealthReportsSortByEnum, 0)
	for _, v := range mappingListSubsettingPolicyHealthReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPolicyHealthReportsSortByEnumStringValues Enumerates the set of values in String for ListSubsettingPolicyHealthReportsSortByEnum
func GetListSubsettingPolicyHealthReportsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListSubsettingPolicyHealthReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPolicyHealthReportsSortByEnum(val string) (ListSubsettingPolicyHealthReportsSortByEnum, bool) {
	enum, ok := mappingListSubsettingPolicyHealthReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingPolicyHealthReportsSortOrderEnum Enum with underlying type: string
type ListSubsettingPolicyHealthReportsSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingPolicyHealthReportsSortOrderEnum
const (
	ListSubsettingPolicyHealthReportsSortOrderAsc  ListSubsettingPolicyHealthReportsSortOrderEnum = "ASC"
	ListSubsettingPolicyHealthReportsSortOrderDesc ListSubsettingPolicyHealthReportsSortOrderEnum = "DESC"
)

var mappingListSubsettingPolicyHealthReportsSortOrderEnum = map[string]ListSubsettingPolicyHealthReportsSortOrderEnum{
	"ASC":  ListSubsettingPolicyHealthReportsSortOrderAsc,
	"DESC": ListSubsettingPolicyHealthReportsSortOrderDesc,
}

var mappingListSubsettingPolicyHealthReportsSortOrderEnumLowerCase = map[string]ListSubsettingPolicyHealthReportsSortOrderEnum{
	"asc":  ListSubsettingPolicyHealthReportsSortOrderAsc,
	"desc": ListSubsettingPolicyHealthReportsSortOrderDesc,
}

// GetListSubsettingPolicyHealthReportsSortOrderEnumValues Enumerates the set of values for ListSubsettingPolicyHealthReportsSortOrderEnum
func GetListSubsettingPolicyHealthReportsSortOrderEnumValues() []ListSubsettingPolicyHealthReportsSortOrderEnum {
	values := make([]ListSubsettingPolicyHealthReportsSortOrderEnum, 0)
	for _, v := range mappingListSubsettingPolicyHealthReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPolicyHealthReportsSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingPolicyHealthReportsSortOrderEnum
func GetListSubsettingPolicyHealthReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingPolicyHealthReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPolicyHealthReportsSortOrderEnum(val string) (ListSubsettingPolicyHealthReportsSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingPolicyHealthReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingPolicyHealthReportsLifecycleStateEnum Enum with underlying type: string
type ListSubsettingPolicyHealthReportsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSubsettingPolicyHealthReportsLifecycleStateEnum
const (
	ListSubsettingPolicyHealthReportsLifecycleStateCreating       ListSubsettingPolicyHealthReportsLifecycleStateEnum = "CREATING"
	ListSubsettingPolicyHealthReportsLifecycleStateActive         ListSubsettingPolicyHealthReportsLifecycleStateEnum = "ACTIVE"
	ListSubsettingPolicyHealthReportsLifecycleStateUpdating       ListSubsettingPolicyHealthReportsLifecycleStateEnum = "UPDATING"
	ListSubsettingPolicyHealthReportsLifecycleStateDeleting       ListSubsettingPolicyHealthReportsLifecycleStateEnum = "DELETING"
	ListSubsettingPolicyHealthReportsLifecycleStateDeleted        ListSubsettingPolicyHealthReportsLifecycleStateEnum = "DELETED"
	ListSubsettingPolicyHealthReportsLifecycleStateNeedsAttention ListSubsettingPolicyHealthReportsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListSubsettingPolicyHealthReportsLifecycleStateFailed         ListSubsettingPolicyHealthReportsLifecycleStateEnum = "FAILED"
)

var mappingListSubsettingPolicyHealthReportsLifecycleStateEnum = map[string]ListSubsettingPolicyHealthReportsLifecycleStateEnum{
	"CREATING":        ListSubsettingPolicyHealthReportsLifecycleStateCreating,
	"ACTIVE":          ListSubsettingPolicyHealthReportsLifecycleStateActive,
	"UPDATING":        ListSubsettingPolicyHealthReportsLifecycleStateUpdating,
	"DELETING":        ListSubsettingPolicyHealthReportsLifecycleStateDeleting,
	"DELETED":         ListSubsettingPolicyHealthReportsLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListSubsettingPolicyHealthReportsLifecycleStateNeedsAttention,
	"FAILED":          ListSubsettingPolicyHealthReportsLifecycleStateFailed,
}

var mappingListSubsettingPolicyHealthReportsLifecycleStateEnumLowerCase = map[string]ListSubsettingPolicyHealthReportsLifecycleStateEnum{
	"creating":        ListSubsettingPolicyHealthReportsLifecycleStateCreating,
	"active":          ListSubsettingPolicyHealthReportsLifecycleStateActive,
	"updating":        ListSubsettingPolicyHealthReportsLifecycleStateUpdating,
	"deleting":        ListSubsettingPolicyHealthReportsLifecycleStateDeleting,
	"deleted":         ListSubsettingPolicyHealthReportsLifecycleStateDeleted,
	"needs_attention": ListSubsettingPolicyHealthReportsLifecycleStateNeedsAttention,
	"failed":          ListSubsettingPolicyHealthReportsLifecycleStateFailed,
}

// GetListSubsettingPolicyHealthReportsLifecycleStateEnumValues Enumerates the set of values for ListSubsettingPolicyHealthReportsLifecycleStateEnum
func GetListSubsettingPolicyHealthReportsLifecycleStateEnumValues() []ListSubsettingPolicyHealthReportsLifecycleStateEnum {
	values := make([]ListSubsettingPolicyHealthReportsLifecycleStateEnum, 0)
	for _, v := range mappingListSubsettingPolicyHealthReportsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingPolicyHealthReportsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSubsettingPolicyHealthReportsLifecycleStateEnum
func GetListSubsettingPolicyHealthReportsLifecycleStateEnumStringValues() []string {
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

// GetMappingListSubsettingPolicyHealthReportsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingPolicyHealthReportsLifecycleStateEnum(val string) (ListSubsettingPolicyHealthReportsLifecycleStateEnum, bool) {
	enum, ok := mappingListSubsettingPolicyHealthReportsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
