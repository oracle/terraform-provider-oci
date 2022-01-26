// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListUserAssessmentsRequest wrapper for the ListUserAssessments operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListUserAssessments.go.html to see an example of how to use ListUserAssessmentsRequest.
type ListUserAssessmentsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListUserAssessmentsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of the user assessment of type SAVE_SCHEDULE.
	ScheduleUserAssessmentId *string `mandatory:"false" contributesTo:"query" name:"scheduleUserAssessmentId"`

	// A filter to return only user assessments of type SAVE_SCHEDULE.
	IsScheduleAssessment *bool `mandatory:"false" contributesTo:"query" name:"isScheduleAssessment"`

	// A filter to return only user assessments that are set as baseline.
	IsBaseline *bool `mandatory:"false" contributesTo:"query" name:"isBaseline"`

	// A filter to return only items that match the specified target.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only items that match the specified assessment type.
	Type ListUserAssessmentsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// A filter to return user assessments that were created by either the system or by a user only.
	TriggeredBy ListUserAssessmentsTriggeredByEnum `mandatory:"false" contributesTo:"query" name:"triggeredBy" omitEmpty:"true"`

	// A filter to return only user assessments that were created after the specified date and time, as defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Using timeCreatedGreaterThanOrEqualTo parameter retrieves all assessments created after that date.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// Search for items that were created before a specific date.
	// Specifying this parameter corresponding `timeCreatedLessThan`
	// parameter will retrieve all items created before the
	// specified created date, in "YYYY-MM-ddThh:mmZ" format with a Z offset, as
	// defined by RFC 3339.
	// **Example:** 2016-12-19T16:39:57.600Z
	TimeCreatedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThan"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The current state of the user assessment.
	LifecycleState ListUserAssessmentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListUserAssessmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order for timeCreated is descending.
	SortBy ListUserAssessmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUserAssessmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUserAssessmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUserAssessmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUserAssessmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListUserAssessmentsResponse wrapper for the ListUserAssessments operation
type ListUserAssessmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []UserAssessmentSummary instances
	Items []UserAssessmentSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListUserAssessmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUserAssessmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUserAssessmentsAccessLevelEnum Enum with underlying type: string
type ListUserAssessmentsAccessLevelEnum string

// Set of constants representing the allowable values for ListUserAssessmentsAccessLevelEnum
const (
	ListUserAssessmentsAccessLevelRestricted ListUserAssessmentsAccessLevelEnum = "RESTRICTED"
	ListUserAssessmentsAccessLevelAccessible ListUserAssessmentsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListUserAssessmentsAccessLevel = map[string]ListUserAssessmentsAccessLevelEnum{
	"RESTRICTED": ListUserAssessmentsAccessLevelRestricted,
	"ACCESSIBLE": ListUserAssessmentsAccessLevelAccessible,
}

// GetListUserAssessmentsAccessLevelEnumValues Enumerates the set of values for ListUserAssessmentsAccessLevelEnum
func GetListUserAssessmentsAccessLevelEnumValues() []ListUserAssessmentsAccessLevelEnum {
	values := make([]ListUserAssessmentsAccessLevelEnum, 0)
	for _, v := range mappingListUserAssessmentsAccessLevel {
		values = append(values, v)
	}
	return values
}

// ListUserAssessmentsTypeEnum Enum with underlying type: string
type ListUserAssessmentsTypeEnum string

// Set of constants representing the allowable values for ListUserAssessmentsTypeEnum
const (
	ListUserAssessmentsTypeLatest       ListUserAssessmentsTypeEnum = "LATEST"
	ListUserAssessmentsTypeSaved        ListUserAssessmentsTypeEnum = "SAVED"
	ListUserAssessmentsTypeCompartment  ListUserAssessmentsTypeEnum = "COMPARTMENT"
	ListUserAssessmentsTypeSaveSchedule ListUserAssessmentsTypeEnum = "SAVE_SCHEDULE"
)

var mappingListUserAssessmentsType = map[string]ListUserAssessmentsTypeEnum{
	"LATEST":        ListUserAssessmentsTypeLatest,
	"SAVED":         ListUserAssessmentsTypeSaved,
	"COMPARTMENT":   ListUserAssessmentsTypeCompartment,
	"SAVE_SCHEDULE": ListUserAssessmentsTypeSaveSchedule,
}

// GetListUserAssessmentsTypeEnumValues Enumerates the set of values for ListUserAssessmentsTypeEnum
func GetListUserAssessmentsTypeEnumValues() []ListUserAssessmentsTypeEnum {
	values := make([]ListUserAssessmentsTypeEnum, 0)
	for _, v := range mappingListUserAssessmentsType {
		values = append(values, v)
	}
	return values
}

// ListUserAssessmentsTriggeredByEnum Enum with underlying type: string
type ListUserAssessmentsTriggeredByEnum string

// Set of constants representing the allowable values for ListUserAssessmentsTriggeredByEnum
const (
	ListUserAssessmentsTriggeredByUser   ListUserAssessmentsTriggeredByEnum = "USER"
	ListUserAssessmentsTriggeredBySystem ListUserAssessmentsTriggeredByEnum = "SYSTEM"
)

var mappingListUserAssessmentsTriggeredBy = map[string]ListUserAssessmentsTriggeredByEnum{
	"USER":   ListUserAssessmentsTriggeredByUser,
	"SYSTEM": ListUserAssessmentsTriggeredBySystem,
}

// GetListUserAssessmentsTriggeredByEnumValues Enumerates the set of values for ListUserAssessmentsTriggeredByEnum
func GetListUserAssessmentsTriggeredByEnumValues() []ListUserAssessmentsTriggeredByEnum {
	values := make([]ListUserAssessmentsTriggeredByEnum, 0)
	for _, v := range mappingListUserAssessmentsTriggeredBy {
		values = append(values, v)
	}
	return values
}

// ListUserAssessmentsLifecycleStateEnum Enum with underlying type: string
type ListUserAssessmentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListUserAssessmentsLifecycleStateEnum
const (
	ListUserAssessmentsLifecycleStateCreating  ListUserAssessmentsLifecycleStateEnum = "CREATING"
	ListUserAssessmentsLifecycleStateSucceeded ListUserAssessmentsLifecycleStateEnum = "SUCCEEDED"
	ListUserAssessmentsLifecycleStateUpdating  ListUserAssessmentsLifecycleStateEnum = "UPDATING"
	ListUserAssessmentsLifecycleStateDeleting  ListUserAssessmentsLifecycleStateEnum = "DELETING"
	ListUserAssessmentsLifecycleStateFailed    ListUserAssessmentsLifecycleStateEnum = "FAILED"
)

var mappingListUserAssessmentsLifecycleState = map[string]ListUserAssessmentsLifecycleStateEnum{
	"CREATING":  ListUserAssessmentsLifecycleStateCreating,
	"SUCCEEDED": ListUserAssessmentsLifecycleStateSucceeded,
	"UPDATING":  ListUserAssessmentsLifecycleStateUpdating,
	"DELETING":  ListUserAssessmentsLifecycleStateDeleting,
	"FAILED":    ListUserAssessmentsLifecycleStateFailed,
}

// GetListUserAssessmentsLifecycleStateEnumValues Enumerates the set of values for ListUserAssessmentsLifecycleStateEnum
func GetListUserAssessmentsLifecycleStateEnumValues() []ListUserAssessmentsLifecycleStateEnum {
	values := make([]ListUserAssessmentsLifecycleStateEnum, 0)
	for _, v := range mappingListUserAssessmentsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListUserAssessmentsSortOrderEnum Enum with underlying type: string
type ListUserAssessmentsSortOrderEnum string

// Set of constants representing the allowable values for ListUserAssessmentsSortOrderEnum
const (
	ListUserAssessmentsSortOrderAsc  ListUserAssessmentsSortOrderEnum = "ASC"
	ListUserAssessmentsSortOrderDesc ListUserAssessmentsSortOrderEnum = "DESC"
)

var mappingListUserAssessmentsSortOrder = map[string]ListUserAssessmentsSortOrderEnum{
	"ASC":  ListUserAssessmentsSortOrderAsc,
	"DESC": ListUserAssessmentsSortOrderDesc,
}

// GetListUserAssessmentsSortOrderEnumValues Enumerates the set of values for ListUserAssessmentsSortOrderEnum
func GetListUserAssessmentsSortOrderEnumValues() []ListUserAssessmentsSortOrderEnum {
	values := make([]ListUserAssessmentsSortOrderEnum, 0)
	for _, v := range mappingListUserAssessmentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListUserAssessmentsSortByEnum Enum with underlying type: string
type ListUserAssessmentsSortByEnum string

// Set of constants representing the allowable values for ListUserAssessmentsSortByEnum
const (
	ListUserAssessmentsSortByTimecreated ListUserAssessmentsSortByEnum = "timeCreated"
	ListUserAssessmentsSortByDisplayname ListUserAssessmentsSortByEnum = "displayName"
)

var mappingListUserAssessmentsSortBy = map[string]ListUserAssessmentsSortByEnum{
	"timeCreated": ListUserAssessmentsSortByTimecreated,
	"displayName": ListUserAssessmentsSortByDisplayname,
}

// GetListUserAssessmentsSortByEnumValues Enumerates the set of values for ListUserAssessmentsSortByEnum
func GetListUserAssessmentsSortByEnumValues() []ListUserAssessmentsSortByEnum {
	values := make([]ListUserAssessmentsSortByEnum, 0)
	for _, v := range mappingListUserAssessmentsSortBy {
		values = append(values, v)
	}
	return values
}
