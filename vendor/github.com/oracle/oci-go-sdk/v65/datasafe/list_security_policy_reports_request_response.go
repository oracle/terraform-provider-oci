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

// ListSecurityPolicyReportsRequest wrapper for the ListSecurityPolicyReports operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityPolicyReports.go.html to see an example of how to use ListSecurityPolicyReportsRequest.
type ListSecurityPolicyReportsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSecurityPolicyReportsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the security policy report.
	LifecycleState ListSecurityPolicyReportsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only resources that match the specified OCID of the security policy report resource.
	SecurityPolicyReportId *string `mandatory:"false" contributesTo:"query" name:"securityPolicyReportId"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSecurityPolicyReportsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListSecurityPolicyReportsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityPolicyReportsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityPolicyReportsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityPolicyReportsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityPolicyReportsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityPolicyReportsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityPolicyReportsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSecurityPolicyReportsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPolicyReportsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSecurityPolicyReportsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPolicyReportsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecurityPolicyReportsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityPolicyReportsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecurityPolicyReportsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityPolicyReportsResponse wrapper for the ListSecurityPolicyReports operation
type ListSecurityPolicyReportsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SecurityPolicyReportCollection instances
	SecurityPolicyReportCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSecurityPolicyReportsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityPolicyReportsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityPolicyReportsAccessLevelEnum Enum with underlying type: string
type ListSecurityPolicyReportsAccessLevelEnum string

// Set of constants representing the allowable values for ListSecurityPolicyReportsAccessLevelEnum
const (
	ListSecurityPolicyReportsAccessLevelRestricted ListSecurityPolicyReportsAccessLevelEnum = "RESTRICTED"
	ListSecurityPolicyReportsAccessLevelAccessible ListSecurityPolicyReportsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSecurityPolicyReportsAccessLevelEnum = map[string]ListSecurityPolicyReportsAccessLevelEnum{
	"RESTRICTED": ListSecurityPolicyReportsAccessLevelRestricted,
	"ACCESSIBLE": ListSecurityPolicyReportsAccessLevelAccessible,
}

var mappingListSecurityPolicyReportsAccessLevelEnumLowerCase = map[string]ListSecurityPolicyReportsAccessLevelEnum{
	"restricted": ListSecurityPolicyReportsAccessLevelRestricted,
	"accessible": ListSecurityPolicyReportsAccessLevelAccessible,
}

// GetListSecurityPolicyReportsAccessLevelEnumValues Enumerates the set of values for ListSecurityPolicyReportsAccessLevelEnum
func GetListSecurityPolicyReportsAccessLevelEnumValues() []ListSecurityPolicyReportsAccessLevelEnum {
	values := make([]ListSecurityPolicyReportsAccessLevelEnum, 0)
	for _, v := range mappingListSecurityPolicyReportsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyReportsAccessLevelEnumStringValues Enumerates the set of values in String for ListSecurityPolicyReportsAccessLevelEnum
func GetListSecurityPolicyReportsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSecurityPolicyReportsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyReportsAccessLevelEnum(val string) (ListSecurityPolicyReportsAccessLevelEnum, bool) {
	enum, ok := mappingListSecurityPolicyReportsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPolicyReportsLifecycleStateEnum Enum with underlying type: string
type ListSecurityPolicyReportsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSecurityPolicyReportsLifecycleStateEnum
const (
	ListSecurityPolicyReportsLifecycleStateCreating       ListSecurityPolicyReportsLifecycleStateEnum = "CREATING"
	ListSecurityPolicyReportsLifecycleStateSucceeded      ListSecurityPolicyReportsLifecycleStateEnum = "SUCCEEDED"
	ListSecurityPolicyReportsLifecycleStateUpdating       ListSecurityPolicyReportsLifecycleStateEnum = "UPDATING"
	ListSecurityPolicyReportsLifecycleStateDeleting       ListSecurityPolicyReportsLifecycleStateEnum = "DELETING"
	ListSecurityPolicyReportsLifecycleStateDeleted        ListSecurityPolicyReportsLifecycleStateEnum = "DELETED"
	ListSecurityPolicyReportsLifecycleStateFailed         ListSecurityPolicyReportsLifecycleStateEnum = "FAILED"
	ListSecurityPolicyReportsLifecycleStateNeedsAttention ListSecurityPolicyReportsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListSecurityPolicyReportsLifecycleStateEnum = map[string]ListSecurityPolicyReportsLifecycleStateEnum{
	"CREATING":        ListSecurityPolicyReportsLifecycleStateCreating,
	"SUCCEEDED":       ListSecurityPolicyReportsLifecycleStateSucceeded,
	"UPDATING":        ListSecurityPolicyReportsLifecycleStateUpdating,
	"DELETING":        ListSecurityPolicyReportsLifecycleStateDeleting,
	"DELETED":         ListSecurityPolicyReportsLifecycleStateDeleted,
	"FAILED":          ListSecurityPolicyReportsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListSecurityPolicyReportsLifecycleStateNeedsAttention,
}

var mappingListSecurityPolicyReportsLifecycleStateEnumLowerCase = map[string]ListSecurityPolicyReportsLifecycleStateEnum{
	"creating":        ListSecurityPolicyReportsLifecycleStateCreating,
	"succeeded":       ListSecurityPolicyReportsLifecycleStateSucceeded,
	"updating":        ListSecurityPolicyReportsLifecycleStateUpdating,
	"deleting":        ListSecurityPolicyReportsLifecycleStateDeleting,
	"deleted":         ListSecurityPolicyReportsLifecycleStateDeleted,
	"failed":          ListSecurityPolicyReportsLifecycleStateFailed,
	"needs_attention": ListSecurityPolicyReportsLifecycleStateNeedsAttention,
}

// GetListSecurityPolicyReportsLifecycleStateEnumValues Enumerates the set of values for ListSecurityPolicyReportsLifecycleStateEnum
func GetListSecurityPolicyReportsLifecycleStateEnumValues() []ListSecurityPolicyReportsLifecycleStateEnum {
	values := make([]ListSecurityPolicyReportsLifecycleStateEnum, 0)
	for _, v := range mappingListSecurityPolicyReportsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyReportsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSecurityPolicyReportsLifecycleStateEnum
func GetListSecurityPolicyReportsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"SUCCEEDED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListSecurityPolicyReportsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyReportsLifecycleStateEnum(val string) (ListSecurityPolicyReportsLifecycleStateEnum, bool) {
	enum, ok := mappingListSecurityPolicyReportsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPolicyReportsSortOrderEnum Enum with underlying type: string
type ListSecurityPolicyReportsSortOrderEnum string

// Set of constants representing the allowable values for ListSecurityPolicyReportsSortOrderEnum
const (
	ListSecurityPolicyReportsSortOrderAsc  ListSecurityPolicyReportsSortOrderEnum = "ASC"
	ListSecurityPolicyReportsSortOrderDesc ListSecurityPolicyReportsSortOrderEnum = "DESC"
)

var mappingListSecurityPolicyReportsSortOrderEnum = map[string]ListSecurityPolicyReportsSortOrderEnum{
	"ASC":  ListSecurityPolicyReportsSortOrderAsc,
	"DESC": ListSecurityPolicyReportsSortOrderDesc,
}

var mappingListSecurityPolicyReportsSortOrderEnumLowerCase = map[string]ListSecurityPolicyReportsSortOrderEnum{
	"asc":  ListSecurityPolicyReportsSortOrderAsc,
	"desc": ListSecurityPolicyReportsSortOrderDesc,
}

// GetListSecurityPolicyReportsSortOrderEnumValues Enumerates the set of values for ListSecurityPolicyReportsSortOrderEnum
func GetListSecurityPolicyReportsSortOrderEnumValues() []ListSecurityPolicyReportsSortOrderEnum {
	values := make([]ListSecurityPolicyReportsSortOrderEnum, 0)
	for _, v := range mappingListSecurityPolicyReportsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyReportsSortOrderEnumStringValues Enumerates the set of values in String for ListSecurityPolicyReportsSortOrderEnum
func GetListSecurityPolicyReportsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSecurityPolicyReportsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyReportsSortOrderEnum(val string) (ListSecurityPolicyReportsSortOrderEnum, bool) {
	enum, ok := mappingListSecurityPolicyReportsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityPolicyReportsSortByEnum Enum with underlying type: string
type ListSecurityPolicyReportsSortByEnum string

// Set of constants representing the allowable values for ListSecurityPolicyReportsSortByEnum
const (
	ListSecurityPolicyReportsSortByTimecreated ListSecurityPolicyReportsSortByEnum = "TIMECREATED"
	ListSecurityPolicyReportsSortByDisplayname ListSecurityPolicyReportsSortByEnum = "DISPLAYNAME"
)

var mappingListSecurityPolicyReportsSortByEnum = map[string]ListSecurityPolicyReportsSortByEnum{
	"TIMECREATED": ListSecurityPolicyReportsSortByTimecreated,
	"DISPLAYNAME": ListSecurityPolicyReportsSortByDisplayname,
}

var mappingListSecurityPolicyReportsSortByEnumLowerCase = map[string]ListSecurityPolicyReportsSortByEnum{
	"timecreated": ListSecurityPolicyReportsSortByTimecreated,
	"displayname": ListSecurityPolicyReportsSortByDisplayname,
}

// GetListSecurityPolicyReportsSortByEnumValues Enumerates the set of values for ListSecurityPolicyReportsSortByEnum
func GetListSecurityPolicyReportsSortByEnumValues() []ListSecurityPolicyReportsSortByEnum {
	values := make([]ListSecurityPolicyReportsSortByEnum, 0)
	for _, v := range mappingListSecurityPolicyReportsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityPolicyReportsSortByEnumStringValues Enumerates the set of values in String for ListSecurityPolicyReportsSortByEnum
func GetListSecurityPolicyReportsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListSecurityPolicyReportsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityPolicyReportsSortByEnum(val string) (ListSecurityPolicyReportsSortByEnum, bool) {
	enum, ok := mappingListSecurityPolicyReportsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
