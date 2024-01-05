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

// ListSecurityAssessmentsRequest wrapper for the ListSecurityAssessments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityAssessments.go.html to see an example of how to use ListSecurityAssessmentsRequest.
type ListSecurityAssessmentsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSecurityAssessmentsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only items that match the specified security assessment type.
	Type ListSecurityAssessmentsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The OCID of the security assessment of type SAVE_SCHEDULE.
	ScheduleAssessmentId *string `mandatory:"false" contributesTo:"query" name:"scheduleAssessmentId"`

	// A filter to return only security assessments of type save schedule.
	IsScheduleAssessment *bool `mandatory:"false" contributesTo:"query" name:"isScheduleAssessment"`

	// A filter to return only security asessments that were created by either user or system.
	TriggeredBy ListSecurityAssessmentsTriggeredByEnum `mandatory:"false" contributesTo:"query" name:"triggeredBy" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSecurityAssessmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only the security assessments that are set as a baseline.
	IsBaseline *bool `mandatory:"false" contributesTo:"query" name:"isBaseline"`

	// The field to sort by. You can specify only one sort order(sortOrder). The default order for timeCreated is descending.
	SortBy ListSecurityAssessmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified lifecycle state.
	LifecycleState ListSecurityAssessmentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityAssessmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityAssessmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityAssessmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityAssessmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityAssessmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityAssessmentsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSecurityAssessmentsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityAssessmentsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListSecurityAssessmentsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityAssessmentsTriggeredByEnum(string(request.TriggeredBy)); !ok && request.TriggeredBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TriggeredBy: %s. Supported values are: %s.", request.TriggeredBy, strings.Join(GetListSecurityAssessmentsTriggeredByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityAssessmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSecurityAssessmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityAssessmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSecurityAssessmentsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityAssessmentsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListSecurityAssessmentsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityAssessmentsResponse wrapper for the ListSecurityAssessments operation
type ListSecurityAssessmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SecurityAssessmentSummary instances
	Items []SecurityAssessmentSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSecurityAssessmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityAssessmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityAssessmentsAccessLevelEnum Enum with underlying type: string
type ListSecurityAssessmentsAccessLevelEnum string

// Set of constants representing the allowable values for ListSecurityAssessmentsAccessLevelEnum
const (
	ListSecurityAssessmentsAccessLevelRestricted ListSecurityAssessmentsAccessLevelEnum = "RESTRICTED"
	ListSecurityAssessmentsAccessLevelAccessible ListSecurityAssessmentsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSecurityAssessmentsAccessLevelEnum = map[string]ListSecurityAssessmentsAccessLevelEnum{
	"RESTRICTED": ListSecurityAssessmentsAccessLevelRestricted,
	"ACCESSIBLE": ListSecurityAssessmentsAccessLevelAccessible,
}

var mappingListSecurityAssessmentsAccessLevelEnumLowerCase = map[string]ListSecurityAssessmentsAccessLevelEnum{
	"restricted": ListSecurityAssessmentsAccessLevelRestricted,
	"accessible": ListSecurityAssessmentsAccessLevelAccessible,
}

// GetListSecurityAssessmentsAccessLevelEnumValues Enumerates the set of values for ListSecurityAssessmentsAccessLevelEnum
func GetListSecurityAssessmentsAccessLevelEnumValues() []ListSecurityAssessmentsAccessLevelEnum {
	values := make([]ListSecurityAssessmentsAccessLevelEnum, 0)
	for _, v := range mappingListSecurityAssessmentsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityAssessmentsAccessLevelEnumStringValues Enumerates the set of values in String for ListSecurityAssessmentsAccessLevelEnum
func GetListSecurityAssessmentsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSecurityAssessmentsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityAssessmentsAccessLevelEnum(val string) (ListSecurityAssessmentsAccessLevelEnum, bool) {
	enum, ok := mappingListSecurityAssessmentsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityAssessmentsTypeEnum Enum with underlying type: string
type ListSecurityAssessmentsTypeEnum string

// Set of constants representing the allowable values for ListSecurityAssessmentsTypeEnum
const (
	ListSecurityAssessmentsTypeLatest       ListSecurityAssessmentsTypeEnum = "LATEST"
	ListSecurityAssessmentsTypeSaved        ListSecurityAssessmentsTypeEnum = "SAVED"
	ListSecurityAssessmentsTypeSaveSchedule ListSecurityAssessmentsTypeEnum = "SAVE_SCHEDULE"
	ListSecurityAssessmentsTypeCompartment  ListSecurityAssessmentsTypeEnum = "COMPARTMENT"
)

var mappingListSecurityAssessmentsTypeEnum = map[string]ListSecurityAssessmentsTypeEnum{
	"LATEST":        ListSecurityAssessmentsTypeLatest,
	"SAVED":         ListSecurityAssessmentsTypeSaved,
	"SAVE_SCHEDULE": ListSecurityAssessmentsTypeSaveSchedule,
	"COMPARTMENT":   ListSecurityAssessmentsTypeCompartment,
}

var mappingListSecurityAssessmentsTypeEnumLowerCase = map[string]ListSecurityAssessmentsTypeEnum{
	"latest":        ListSecurityAssessmentsTypeLatest,
	"saved":         ListSecurityAssessmentsTypeSaved,
	"save_schedule": ListSecurityAssessmentsTypeSaveSchedule,
	"compartment":   ListSecurityAssessmentsTypeCompartment,
}

// GetListSecurityAssessmentsTypeEnumValues Enumerates the set of values for ListSecurityAssessmentsTypeEnum
func GetListSecurityAssessmentsTypeEnumValues() []ListSecurityAssessmentsTypeEnum {
	values := make([]ListSecurityAssessmentsTypeEnum, 0)
	for _, v := range mappingListSecurityAssessmentsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityAssessmentsTypeEnumStringValues Enumerates the set of values in String for ListSecurityAssessmentsTypeEnum
func GetListSecurityAssessmentsTypeEnumStringValues() []string {
	return []string{
		"LATEST",
		"SAVED",
		"SAVE_SCHEDULE",
		"COMPARTMENT",
	}
}

// GetMappingListSecurityAssessmentsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityAssessmentsTypeEnum(val string) (ListSecurityAssessmentsTypeEnum, bool) {
	enum, ok := mappingListSecurityAssessmentsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityAssessmentsTriggeredByEnum Enum with underlying type: string
type ListSecurityAssessmentsTriggeredByEnum string

// Set of constants representing the allowable values for ListSecurityAssessmentsTriggeredByEnum
const (
	ListSecurityAssessmentsTriggeredByUser   ListSecurityAssessmentsTriggeredByEnum = "USER"
	ListSecurityAssessmentsTriggeredBySystem ListSecurityAssessmentsTriggeredByEnum = "SYSTEM"
)

var mappingListSecurityAssessmentsTriggeredByEnum = map[string]ListSecurityAssessmentsTriggeredByEnum{
	"USER":   ListSecurityAssessmentsTriggeredByUser,
	"SYSTEM": ListSecurityAssessmentsTriggeredBySystem,
}

var mappingListSecurityAssessmentsTriggeredByEnumLowerCase = map[string]ListSecurityAssessmentsTriggeredByEnum{
	"user":   ListSecurityAssessmentsTriggeredByUser,
	"system": ListSecurityAssessmentsTriggeredBySystem,
}

// GetListSecurityAssessmentsTriggeredByEnumValues Enumerates the set of values for ListSecurityAssessmentsTriggeredByEnum
func GetListSecurityAssessmentsTriggeredByEnumValues() []ListSecurityAssessmentsTriggeredByEnum {
	values := make([]ListSecurityAssessmentsTriggeredByEnum, 0)
	for _, v := range mappingListSecurityAssessmentsTriggeredByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityAssessmentsTriggeredByEnumStringValues Enumerates the set of values in String for ListSecurityAssessmentsTriggeredByEnum
func GetListSecurityAssessmentsTriggeredByEnumStringValues() []string {
	return []string{
		"USER",
		"SYSTEM",
	}
}

// GetMappingListSecurityAssessmentsTriggeredByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityAssessmentsTriggeredByEnum(val string) (ListSecurityAssessmentsTriggeredByEnum, bool) {
	enum, ok := mappingListSecurityAssessmentsTriggeredByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityAssessmentsSortOrderEnum Enum with underlying type: string
type ListSecurityAssessmentsSortOrderEnum string

// Set of constants representing the allowable values for ListSecurityAssessmentsSortOrderEnum
const (
	ListSecurityAssessmentsSortOrderAsc  ListSecurityAssessmentsSortOrderEnum = "ASC"
	ListSecurityAssessmentsSortOrderDesc ListSecurityAssessmentsSortOrderEnum = "DESC"
)

var mappingListSecurityAssessmentsSortOrderEnum = map[string]ListSecurityAssessmentsSortOrderEnum{
	"ASC":  ListSecurityAssessmentsSortOrderAsc,
	"DESC": ListSecurityAssessmentsSortOrderDesc,
}

var mappingListSecurityAssessmentsSortOrderEnumLowerCase = map[string]ListSecurityAssessmentsSortOrderEnum{
	"asc":  ListSecurityAssessmentsSortOrderAsc,
	"desc": ListSecurityAssessmentsSortOrderDesc,
}

// GetListSecurityAssessmentsSortOrderEnumValues Enumerates the set of values for ListSecurityAssessmentsSortOrderEnum
func GetListSecurityAssessmentsSortOrderEnumValues() []ListSecurityAssessmentsSortOrderEnum {
	values := make([]ListSecurityAssessmentsSortOrderEnum, 0)
	for _, v := range mappingListSecurityAssessmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityAssessmentsSortOrderEnumStringValues Enumerates the set of values in String for ListSecurityAssessmentsSortOrderEnum
func GetListSecurityAssessmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSecurityAssessmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityAssessmentsSortOrderEnum(val string) (ListSecurityAssessmentsSortOrderEnum, bool) {
	enum, ok := mappingListSecurityAssessmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityAssessmentsSortByEnum Enum with underlying type: string
type ListSecurityAssessmentsSortByEnum string

// Set of constants representing the allowable values for ListSecurityAssessmentsSortByEnum
const (
	ListSecurityAssessmentsSortByTimecreated ListSecurityAssessmentsSortByEnum = "timeCreated"
	ListSecurityAssessmentsSortByDisplayname ListSecurityAssessmentsSortByEnum = "displayName"
)

var mappingListSecurityAssessmentsSortByEnum = map[string]ListSecurityAssessmentsSortByEnum{
	"timeCreated": ListSecurityAssessmentsSortByTimecreated,
	"displayName": ListSecurityAssessmentsSortByDisplayname,
}

var mappingListSecurityAssessmentsSortByEnumLowerCase = map[string]ListSecurityAssessmentsSortByEnum{
	"timecreated": ListSecurityAssessmentsSortByTimecreated,
	"displayname": ListSecurityAssessmentsSortByDisplayname,
}

// GetListSecurityAssessmentsSortByEnumValues Enumerates the set of values for ListSecurityAssessmentsSortByEnum
func GetListSecurityAssessmentsSortByEnumValues() []ListSecurityAssessmentsSortByEnum {
	values := make([]ListSecurityAssessmentsSortByEnum, 0)
	for _, v := range mappingListSecurityAssessmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityAssessmentsSortByEnumStringValues Enumerates the set of values in String for ListSecurityAssessmentsSortByEnum
func GetListSecurityAssessmentsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListSecurityAssessmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityAssessmentsSortByEnum(val string) (ListSecurityAssessmentsSortByEnum, bool) {
	enum, ok := mappingListSecurityAssessmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityAssessmentsLifecycleStateEnum Enum with underlying type: string
type ListSecurityAssessmentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSecurityAssessmentsLifecycleStateEnum
const (
	ListSecurityAssessmentsLifecycleStateCreating  ListSecurityAssessmentsLifecycleStateEnum = "CREATING"
	ListSecurityAssessmentsLifecycleStateSucceeded ListSecurityAssessmentsLifecycleStateEnum = "SUCCEEDED"
	ListSecurityAssessmentsLifecycleStateUpdating  ListSecurityAssessmentsLifecycleStateEnum = "UPDATING"
	ListSecurityAssessmentsLifecycleStateDeleting  ListSecurityAssessmentsLifecycleStateEnum = "DELETING"
	ListSecurityAssessmentsLifecycleStateDeleted   ListSecurityAssessmentsLifecycleStateEnum = "DELETED"
	ListSecurityAssessmentsLifecycleStateFailed    ListSecurityAssessmentsLifecycleStateEnum = "FAILED"
)

var mappingListSecurityAssessmentsLifecycleStateEnum = map[string]ListSecurityAssessmentsLifecycleStateEnum{
	"CREATING":  ListSecurityAssessmentsLifecycleStateCreating,
	"SUCCEEDED": ListSecurityAssessmentsLifecycleStateSucceeded,
	"UPDATING":  ListSecurityAssessmentsLifecycleStateUpdating,
	"DELETING":  ListSecurityAssessmentsLifecycleStateDeleting,
	"DELETED":   ListSecurityAssessmentsLifecycleStateDeleted,
	"FAILED":    ListSecurityAssessmentsLifecycleStateFailed,
}

var mappingListSecurityAssessmentsLifecycleStateEnumLowerCase = map[string]ListSecurityAssessmentsLifecycleStateEnum{
	"creating":  ListSecurityAssessmentsLifecycleStateCreating,
	"succeeded": ListSecurityAssessmentsLifecycleStateSucceeded,
	"updating":  ListSecurityAssessmentsLifecycleStateUpdating,
	"deleting":  ListSecurityAssessmentsLifecycleStateDeleting,
	"deleted":   ListSecurityAssessmentsLifecycleStateDeleted,
	"failed":    ListSecurityAssessmentsLifecycleStateFailed,
}

// GetListSecurityAssessmentsLifecycleStateEnumValues Enumerates the set of values for ListSecurityAssessmentsLifecycleStateEnum
func GetListSecurityAssessmentsLifecycleStateEnumValues() []ListSecurityAssessmentsLifecycleStateEnum {
	values := make([]ListSecurityAssessmentsLifecycleStateEnum, 0)
	for _, v := range mappingListSecurityAssessmentsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityAssessmentsLifecycleStateEnumStringValues Enumerates the set of values in String for ListSecurityAssessmentsLifecycleStateEnum
func GetListSecurityAssessmentsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"SUCCEEDED",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListSecurityAssessmentsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityAssessmentsLifecycleStateEnum(val string) (ListSecurityAssessmentsLifecycleStateEnum, bool) {
	enum, ok := mappingListSecurityAssessmentsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
