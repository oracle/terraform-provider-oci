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

// ListAuditTrailsRequest wrapper for the ListAuditTrails operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListAuditTrails.go.html to see an example of how to use ListAuditTrailsRequest.
type ListAuditTrailsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListAuditTrailsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A optional filter to return only resources that match the specified id.
	AuditTrailId *string `mandatory:"false" contributesTo:"query" name:"auditTrailId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A optional filter to return only resources that match the specified lifecycle state.
	LifecycleState ListAuditTrailsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A optional filter to return only resources that match the specified sub-state of audit trail.
	Status ListAuditTrailsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListAuditTrailsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListAuditTrailsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAuditTrailsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAuditTrailsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAuditTrailsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAuditTrailsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAuditTrailsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAuditTrailsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListAuditTrailsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditTrailsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAuditTrailsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditTrailsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListAuditTrailsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditTrailsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAuditTrailsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAuditTrailsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAuditTrailsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAuditTrailsResponse wrapper for the ListAuditTrails operation
type ListAuditTrailsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AuditTrailCollection instances
	AuditTrailCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListAuditTrailsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAuditTrailsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAuditTrailsAccessLevelEnum Enum with underlying type: string
type ListAuditTrailsAccessLevelEnum string

// Set of constants representing the allowable values for ListAuditTrailsAccessLevelEnum
const (
	ListAuditTrailsAccessLevelRestricted ListAuditTrailsAccessLevelEnum = "RESTRICTED"
	ListAuditTrailsAccessLevelAccessible ListAuditTrailsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListAuditTrailsAccessLevelEnum = map[string]ListAuditTrailsAccessLevelEnum{
	"RESTRICTED": ListAuditTrailsAccessLevelRestricted,
	"ACCESSIBLE": ListAuditTrailsAccessLevelAccessible,
}

var mappingListAuditTrailsAccessLevelEnumLowerCase = map[string]ListAuditTrailsAccessLevelEnum{
	"restricted": ListAuditTrailsAccessLevelRestricted,
	"accessible": ListAuditTrailsAccessLevelAccessible,
}

// GetListAuditTrailsAccessLevelEnumValues Enumerates the set of values for ListAuditTrailsAccessLevelEnum
func GetListAuditTrailsAccessLevelEnumValues() []ListAuditTrailsAccessLevelEnum {
	values := make([]ListAuditTrailsAccessLevelEnum, 0)
	for _, v := range mappingListAuditTrailsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditTrailsAccessLevelEnumStringValues Enumerates the set of values in String for ListAuditTrailsAccessLevelEnum
func GetListAuditTrailsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListAuditTrailsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditTrailsAccessLevelEnum(val string) (ListAuditTrailsAccessLevelEnum, bool) {
	enum, ok := mappingListAuditTrailsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditTrailsLifecycleStateEnum Enum with underlying type: string
type ListAuditTrailsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAuditTrailsLifecycleStateEnum
const (
	ListAuditTrailsLifecycleStateInactive       ListAuditTrailsLifecycleStateEnum = "INACTIVE"
	ListAuditTrailsLifecycleStateUpdating       ListAuditTrailsLifecycleStateEnum = "UPDATING"
	ListAuditTrailsLifecycleStateActive         ListAuditTrailsLifecycleStateEnum = "ACTIVE"
	ListAuditTrailsLifecycleStateDeleting       ListAuditTrailsLifecycleStateEnum = "DELETING"
	ListAuditTrailsLifecycleStateFailed         ListAuditTrailsLifecycleStateEnum = "FAILED"
	ListAuditTrailsLifecycleStateNeedsAttention ListAuditTrailsLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingListAuditTrailsLifecycleStateEnum = map[string]ListAuditTrailsLifecycleStateEnum{
	"INACTIVE":        ListAuditTrailsLifecycleStateInactive,
	"UPDATING":        ListAuditTrailsLifecycleStateUpdating,
	"ACTIVE":          ListAuditTrailsLifecycleStateActive,
	"DELETING":        ListAuditTrailsLifecycleStateDeleting,
	"FAILED":          ListAuditTrailsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListAuditTrailsLifecycleStateNeedsAttention,
}

var mappingListAuditTrailsLifecycleStateEnumLowerCase = map[string]ListAuditTrailsLifecycleStateEnum{
	"inactive":        ListAuditTrailsLifecycleStateInactive,
	"updating":        ListAuditTrailsLifecycleStateUpdating,
	"active":          ListAuditTrailsLifecycleStateActive,
	"deleting":        ListAuditTrailsLifecycleStateDeleting,
	"failed":          ListAuditTrailsLifecycleStateFailed,
	"needs_attention": ListAuditTrailsLifecycleStateNeedsAttention,
}

// GetListAuditTrailsLifecycleStateEnumValues Enumerates the set of values for ListAuditTrailsLifecycleStateEnum
func GetListAuditTrailsLifecycleStateEnumValues() []ListAuditTrailsLifecycleStateEnum {
	values := make([]ListAuditTrailsLifecycleStateEnum, 0)
	for _, v := range mappingListAuditTrailsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditTrailsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAuditTrailsLifecycleStateEnum
func GetListAuditTrailsLifecycleStateEnumStringValues() []string {
	return []string{
		"INACTIVE",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListAuditTrailsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditTrailsLifecycleStateEnum(val string) (ListAuditTrailsLifecycleStateEnum, bool) {
	enum, ok := mappingListAuditTrailsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditTrailsStatusEnum Enum with underlying type: string
type ListAuditTrailsStatusEnum string

// Set of constants representing the allowable values for ListAuditTrailsStatusEnum
const (
	ListAuditTrailsStatusStarting         ListAuditTrailsStatusEnum = "STARTING"
	ListAuditTrailsStatusCollecting       ListAuditTrailsStatusEnum = "COLLECTING"
	ListAuditTrailsStatusRecovering       ListAuditTrailsStatusEnum = "RECOVERING"
	ListAuditTrailsStatusIdle             ListAuditTrailsStatusEnum = "IDLE"
	ListAuditTrailsStatusStopping         ListAuditTrailsStatusEnum = "STOPPING"
	ListAuditTrailsStatusStopped          ListAuditTrailsStatusEnum = "STOPPED"
	ListAuditTrailsStatusResuming         ListAuditTrailsStatusEnum = "RESUMING"
	ListAuditTrailsStatusRetrying         ListAuditTrailsStatusEnum = "RETRYING"
	ListAuditTrailsStatusNotStarted       ListAuditTrailsStatusEnum = "NOT_STARTED"
	ListAuditTrailsStatusStoppedNeedsAttn ListAuditTrailsStatusEnum = "STOPPED_NEEDS_ATTN"
	ListAuditTrailsStatusStoppedFailed    ListAuditTrailsStatusEnum = "STOPPED_FAILED"
)

var mappingListAuditTrailsStatusEnum = map[string]ListAuditTrailsStatusEnum{
	"STARTING":           ListAuditTrailsStatusStarting,
	"COLLECTING":         ListAuditTrailsStatusCollecting,
	"RECOVERING":         ListAuditTrailsStatusRecovering,
	"IDLE":               ListAuditTrailsStatusIdle,
	"STOPPING":           ListAuditTrailsStatusStopping,
	"STOPPED":            ListAuditTrailsStatusStopped,
	"RESUMING":           ListAuditTrailsStatusResuming,
	"RETRYING":           ListAuditTrailsStatusRetrying,
	"NOT_STARTED":        ListAuditTrailsStatusNotStarted,
	"STOPPED_NEEDS_ATTN": ListAuditTrailsStatusStoppedNeedsAttn,
	"STOPPED_FAILED":     ListAuditTrailsStatusStoppedFailed,
}

var mappingListAuditTrailsStatusEnumLowerCase = map[string]ListAuditTrailsStatusEnum{
	"starting":           ListAuditTrailsStatusStarting,
	"collecting":         ListAuditTrailsStatusCollecting,
	"recovering":         ListAuditTrailsStatusRecovering,
	"idle":               ListAuditTrailsStatusIdle,
	"stopping":           ListAuditTrailsStatusStopping,
	"stopped":            ListAuditTrailsStatusStopped,
	"resuming":           ListAuditTrailsStatusResuming,
	"retrying":           ListAuditTrailsStatusRetrying,
	"not_started":        ListAuditTrailsStatusNotStarted,
	"stopped_needs_attn": ListAuditTrailsStatusStoppedNeedsAttn,
	"stopped_failed":     ListAuditTrailsStatusStoppedFailed,
}

// GetListAuditTrailsStatusEnumValues Enumerates the set of values for ListAuditTrailsStatusEnum
func GetListAuditTrailsStatusEnumValues() []ListAuditTrailsStatusEnum {
	values := make([]ListAuditTrailsStatusEnum, 0)
	for _, v := range mappingListAuditTrailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditTrailsStatusEnumStringValues Enumerates the set of values in String for ListAuditTrailsStatusEnum
func GetListAuditTrailsStatusEnumStringValues() []string {
	return []string{
		"STARTING",
		"COLLECTING",
		"RECOVERING",
		"IDLE",
		"STOPPING",
		"STOPPED",
		"RESUMING",
		"RETRYING",
		"NOT_STARTED",
		"STOPPED_NEEDS_ATTN",
		"STOPPED_FAILED",
	}
}

// GetMappingListAuditTrailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditTrailsStatusEnum(val string) (ListAuditTrailsStatusEnum, bool) {
	enum, ok := mappingListAuditTrailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditTrailsSortOrderEnum Enum with underlying type: string
type ListAuditTrailsSortOrderEnum string

// Set of constants representing the allowable values for ListAuditTrailsSortOrderEnum
const (
	ListAuditTrailsSortOrderAsc  ListAuditTrailsSortOrderEnum = "ASC"
	ListAuditTrailsSortOrderDesc ListAuditTrailsSortOrderEnum = "DESC"
)

var mappingListAuditTrailsSortOrderEnum = map[string]ListAuditTrailsSortOrderEnum{
	"ASC":  ListAuditTrailsSortOrderAsc,
	"DESC": ListAuditTrailsSortOrderDesc,
}

var mappingListAuditTrailsSortOrderEnumLowerCase = map[string]ListAuditTrailsSortOrderEnum{
	"asc":  ListAuditTrailsSortOrderAsc,
	"desc": ListAuditTrailsSortOrderDesc,
}

// GetListAuditTrailsSortOrderEnumValues Enumerates the set of values for ListAuditTrailsSortOrderEnum
func GetListAuditTrailsSortOrderEnumValues() []ListAuditTrailsSortOrderEnum {
	values := make([]ListAuditTrailsSortOrderEnum, 0)
	for _, v := range mappingListAuditTrailsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditTrailsSortOrderEnumStringValues Enumerates the set of values in String for ListAuditTrailsSortOrderEnum
func GetListAuditTrailsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAuditTrailsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditTrailsSortOrderEnum(val string) (ListAuditTrailsSortOrderEnum, bool) {
	enum, ok := mappingListAuditTrailsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAuditTrailsSortByEnum Enum with underlying type: string
type ListAuditTrailsSortByEnum string

// Set of constants representing the allowable values for ListAuditTrailsSortByEnum
const (
	ListAuditTrailsSortByTimecreated ListAuditTrailsSortByEnum = "TIMECREATED"
	ListAuditTrailsSortByDisplayname ListAuditTrailsSortByEnum = "DISPLAYNAME"
)

var mappingListAuditTrailsSortByEnum = map[string]ListAuditTrailsSortByEnum{
	"TIMECREATED": ListAuditTrailsSortByTimecreated,
	"DISPLAYNAME": ListAuditTrailsSortByDisplayname,
}

var mappingListAuditTrailsSortByEnumLowerCase = map[string]ListAuditTrailsSortByEnum{
	"timecreated": ListAuditTrailsSortByTimecreated,
	"displayname": ListAuditTrailsSortByDisplayname,
}

// GetListAuditTrailsSortByEnumValues Enumerates the set of values for ListAuditTrailsSortByEnum
func GetListAuditTrailsSortByEnumValues() []ListAuditTrailsSortByEnum {
	values := make([]ListAuditTrailsSortByEnum, 0)
	for _, v := range mappingListAuditTrailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAuditTrailsSortByEnumStringValues Enumerates the set of values in String for ListAuditTrailsSortByEnum
func GetListAuditTrailsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAuditTrailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAuditTrailsSortByEnum(val string) (ListAuditTrailsSortByEnum, bool) {
	enum, ok := mappingListAuditTrailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
