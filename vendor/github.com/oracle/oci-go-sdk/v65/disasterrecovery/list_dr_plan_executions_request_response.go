// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package disasterrecovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDrPlanExecutionsRequest wrapper for the ListDrPlanExecutions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/disasterrecovery/ListDrPlanExecutions.go.html to see an example of how to use ListDrPlanExecutionsRequest.
type ListDrPlanExecutionsRequest struct {

	// The OCID of the DR protection group. Mandatory query param.
	// Example: `ocid1.drprotectiongroup.oc1..uniqueID`
	DrProtectionGroupId *string `mandatory:"true" contributesTo:"query" name:"drProtectionGroupId"`

	// A filter to return only DR plan executions that match the given lifecycle state.
	LifecycleState ListDrPlanExecutionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID of the DR plan execution.
	// Example: `ocid1.drplanexecution.oc1..uniqueID`
	DrPlanExecutionId *string `mandatory:"false" contributesTo:"query" name:"drPlanExecutionId"`

	// The DR plan execution type.
	DrPlanExecutionType ListDrPlanExecutionsDrPlanExecutionTypeEnum `mandatory:"false" contributesTo:"query" name:"drPlanExecutionType" omitEmpty:"true"`

	// A filter to return only resources that match the given display name.
	// Example: `MyResourceDisplayName`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `100`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDrPlanExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	// Example: `MyResourceDisplayName`
	SortBy ListDrPlanExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDrPlanExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDrPlanExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDrPlanExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDrPlanExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDrPlanExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDrPlanExecutionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListDrPlanExecutionsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrPlanExecutionsDrPlanExecutionTypeEnum(string(request.DrPlanExecutionType)); !ok && request.DrPlanExecutionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DrPlanExecutionType: %s. Supported values are: %s.", request.DrPlanExecutionType, strings.Join(GetListDrPlanExecutionsDrPlanExecutionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrPlanExecutionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDrPlanExecutionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDrPlanExecutionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDrPlanExecutionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDrPlanExecutionsResponse wrapper for the ListDrPlanExecutions operation
type ListDrPlanExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DrPlanExecutionCollection instances
	DrPlanExecutionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDrPlanExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDrPlanExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDrPlanExecutionsLifecycleStateEnum Enum with underlying type: string
type ListDrPlanExecutionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDrPlanExecutionsLifecycleStateEnum
const (
	ListDrPlanExecutionsLifecycleStateAccepted   ListDrPlanExecutionsLifecycleStateEnum = "ACCEPTED"
	ListDrPlanExecutionsLifecycleStateInProgress ListDrPlanExecutionsLifecycleStateEnum = "IN_PROGRESS"
	ListDrPlanExecutionsLifecycleStateWaiting    ListDrPlanExecutionsLifecycleStateEnum = "WAITING"
	ListDrPlanExecutionsLifecycleStateCanceling  ListDrPlanExecutionsLifecycleStateEnum = "CANCELING"
	ListDrPlanExecutionsLifecycleStateCanceled   ListDrPlanExecutionsLifecycleStateEnum = "CANCELED"
	ListDrPlanExecutionsLifecycleStateSucceeded  ListDrPlanExecutionsLifecycleStateEnum = "SUCCEEDED"
	ListDrPlanExecutionsLifecycleStateFailed     ListDrPlanExecutionsLifecycleStateEnum = "FAILED"
	ListDrPlanExecutionsLifecycleStateDeleting   ListDrPlanExecutionsLifecycleStateEnum = "DELETING"
	ListDrPlanExecutionsLifecycleStateDeleted    ListDrPlanExecutionsLifecycleStateEnum = "DELETED"
	ListDrPlanExecutionsLifecycleStatePausing    ListDrPlanExecutionsLifecycleStateEnum = "PAUSING"
	ListDrPlanExecutionsLifecycleStatePaused     ListDrPlanExecutionsLifecycleStateEnum = "PAUSED"
	ListDrPlanExecutionsLifecycleStateResuming   ListDrPlanExecutionsLifecycleStateEnum = "RESUMING"
)

var mappingListDrPlanExecutionsLifecycleStateEnum = map[string]ListDrPlanExecutionsLifecycleStateEnum{
	"ACCEPTED":    ListDrPlanExecutionsLifecycleStateAccepted,
	"IN_PROGRESS": ListDrPlanExecutionsLifecycleStateInProgress,
	"WAITING":     ListDrPlanExecutionsLifecycleStateWaiting,
	"CANCELING":   ListDrPlanExecutionsLifecycleStateCanceling,
	"CANCELED":    ListDrPlanExecutionsLifecycleStateCanceled,
	"SUCCEEDED":   ListDrPlanExecutionsLifecycleStateSucceeded,
	"FAILED":      ListDrPlanExecutionsLifecycleStateFailed,
	"DELETING":    ListDrPlanExecutionsLifecycleStateDeleting,
	"DELETED":     ListDrPlanExecutionsLifecycleStateDeleted,
	"PAUSING":     ListDrPlanExecutionsLifecycleStatePausing,
	"PAUSED":      ListDrPlanExecutionsLifecycleStatePaused,
	"RESUMING":    ListDrPlanExecutionsLifecycleStateResuming,
}

var mappingListDrPlanExecutionsLifecycleStateEnumLowerCase = map[string]ListDrPlanExecutionsLifecycleStateEnum{
	"accepted":    ListDrPlanExecutionsLifecycleStateAccepted,
	"in_progress": ListDrPlanExecutionsLifecycleStateInProgress,
	"waiting":     ListDrPlanExecutionsLifecycleStateWaiting,
	"canceling":   ListDrPlanExecutionsLifecycleStateCanceling,
	"canceled":    ListDrPlanExecutionsLifecycleStateCanceled,
	"succeeded":   ListDrPlanExecutionsLifecycleStateSucceeded,
	"failed":      ListDrPlanExecutionsLifecycleStateFailed,
	"deleting":    ListDrPlanExecutionsLifecycleStateDeleting,
	"deleted":     ListDrPlanExecutionsLifecycleStateDeleted,
	"pausing":     ListDrPlanExecutionsLifecycleStatePausing,
	"paused":      ListDrPlanExecutionsLifecycleStatePaused,
	"resuming":    ListDrPlanExecutionsLifecycleStateResuming,
}

// GetListDrPlanExecutionsLifecycleStateEnumValues Enumerates the set of values for ListDrPlanExecutionsLifecycleStateEnum
func GetListDrPlanExecutionsLifecycleStateEnumValues() []ListDrPlanExecutionsLifecycleStateEnum {
	values := make([]ListDrPlanExecutionsLifecycleStateEnum, 0)
	for _, v := range mappingListDrPlanExecutionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrPlanExecutionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListDrPlanExecutionsLifecycleStateEnum
func GetListDrPlanExecutionsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"CANCELING",
		"CANCELED",
		"SUCCEEDED",
		"FAILED",
		"DELETING",
		"DELETED",
		"PAUSING",
		"PAUSED",
		"RESUMING",
	}
}

// GetMappingListDrPlanExecutionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrPlanExecutionsLifecycleStateEnum(val string) (ListDrPlanExecutionsLifecycleStateEnum, bool) {
	enum, ok := mappingListDrPlanExecutionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrPlanExecutionsDrPlanExecutionTypeEnum Enum with underlying type: string
type ListDrPlanExecutionsDrPlanExecutionTypeEnum string

// Set of constants representing the allowable values for ListDrPlanExecutionsDrPlanExecutionTypeEnum
const (
	ListDrPlanExecutionsDrPlanExecutionTypeSwitchover         ListDrPlanExecutionsDrPlanExecutionTypeEnum = "SWITCHOVER"
	ListDrPlanExecutionsDrPlanExecutionTypeSwitchoverPrecheck ListDrPlanExecutionsDrPlanExecutionTypeEnum = "SWITCHOVER_PRECHECK"
	ListDrPlanExecutionsDrPlanExecutionTypeFailover           ListDrPlanExecutionsDrPlanExecutionTypeEnum = "FAILOVER"
	ListDrPlanExecutionsDrPlanExecutionTypeFailoverPrecheck   ListDrPlanExecutionsDrPlanExecutionTypeEnum = "FAILOVER_PRECHECK"
	ListDrPlanExecutionsDrPlanExecutionTypeStartDrill         ListDrPlanExecutionsDrPlanExecutionTypeEnum = "START_DRILL"
	ListDrPlanExecutionsDrPlanExecutionTypeStartDrillPrecheck ListDrPlanExecutionsDrPlanExecutionTypeEnum = "START_DRILL_PRECHECK"
	ListDrPlanExecutionsDrPlanExecutionTypeStopDrill          ListDrPlanExecutionsDrPlanExecutionTypeEnum = "STOP_DRILL"
	ListDrPlanExecutionsDrPlanExecutionTypeStopDrillPrecheck  ListDrPlanExecutionsDrPlanExecutionTypeEnum = "STOP_DRILL_PRECHECK"
)

var mappingListDrPlanExecutionsDrPlanExecutionTypeEnum = map[string]ListDrPlanExecutionsDrPlanExecutionTypeEnum{
	"SWITCHOVER":           ListDrPlanExecutionsDrPlanExecutionTypeSwitchover,
	"SWITCHOVER_PRECHECK":  ListDrPlanExecutionsDrPlanExecutionTypeSwitchoverPrecheck,
	"FAILOVER":             ListDrPlanExecutionsDrPlanExecutionTypeFailover,
	"FAILOVER_PRECHECK":    ListDrPlanExecutionsDrPlanExecutionTypeFailoverPrecheck,
	"START_DRILL":          ListDrPlanExecutionsDrPlanExecutionTypeStartDrill,
	"START_DRILL_PRECHECK": ListDrPlanExecutionsDrPlanExecutionTypeStartDrillPrecheck,
	"STOP_DRILL":           ListDrPlanExecutionsDrPlanExecutionTypeStopDrill,
	"STOP_DRILL_PRECHECK":  ListDrPlanExecutionsDrPlanExecutionTypeStopDrillPrecheck,
}

var mappingListDrPlanExecutionsDrPlanExecutionTypeEnumLowerCase = map[string]ListDrPlanExecutionsDrPlanExecutionTypeEnum{
	"switchover":           ListDrPlanExecutionsDrPlanExecutionTypeSwitchover,
	"switchover_precheck":  ListDrPlanExecutionsDrPlanExecutionTypeSwitchoverPrecheck,
	"failover":             ListDrPlanExecutionsDrPlanExecutionTypeFailover,
	"failover_precheck":    ListDrPlanExecutionsDrPlanExecutionTypeFailoverPrecheck,
	"start_drill":          ListDrPlanExecutionsDrPlanExecutionTypeStartDrill,
	"start_drill_precheck": ListDrPlanExecutionsDrPlanExecutionTypeStartDrillPrecheck,
	"stop_drill":           ListDrPlanExecutionsDrPlanExecutionTypeStopDrill,
	"stop_drill_precheck":  ListDrPlanExecutionsDrPlanExecutionTypeStopDrillPrecheck,
}

// GetListDrPlanExecutionsDrPlanExecutionTypeEnumValues Enumerates the set of values for ListDrPlanExecutionsDrPlanExecutionTypeEnum
func GetListDrPlanExecutionsDrPlanExecutionTypeEnumValues() []ListDrPlanExecutionsDrPlanExecutionTypeEnum {
	values := make([]ListDrPlanExecutionsDrPlanExecutionTypeEnum, 0)
	for _, v := range mappingListDrPlanExecutionsDrPlanExecutionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrPlanExecutionsDrPlanExecutionTypeEnumStringValues Enumerates the set of values in String for ListDrPlanExecutionsDrPlanExecutionTypeEnum
func GetListDrPlanExecutionsDrPlanExecutionTypeEnumStringValues() []string {
	return []string{
		"SWITCHOVER",
		"SWITCHOVER_PRECHECK",
		"FAILOVER",
		"FAILOVER_PRECHECK",
		"START_DRILL",
		"START_DRILL_PRECHECK",
		"STOP_DRILL",
		"STOP_DRILL_PRECHECK",
	}
}

// GetMappingListDrPlanExecutionsDrPlanExecutionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrPlanExecutionsDrPlanExecutionTypeEnum(val string) (ListDrPlanExecutionsDrPlanExecutionTypeEnum, bool) {
	enum, ok := mappingListDrPlanExecutionsDrPlanExecutionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrPlanExecutionsSortOrderEnum Enum with underlying type: string
type ListDrPlanExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListDrPlanExecutionsSortOrderEnum
const (
	ListDrPlanExecutionsSortOrderAsc  ListDrPlanExecutionsSortOrderEnum = "ASC"
	ListDrPlanExecutionsSortOrderDesc ListDrPlanExecutionsSortOrderEnum = "DESC"
)

var mappingListDrPlanExecutionsSortOrderEnum = map[string]ListDrPlanExecutionsSortOrderEnum{
	"ASC":  ListDrPlanExecutionsSortOrderAsc,
	"DESC": ListDrPlanExecutionsSortOrderDesc,
}

var mappingListDrPlanExecutionsSortOrderEnumLowerCase = map[string]ListDrPlanExecutionsSortOrderEnum{
	"asc":  ListDrPlanExecutionsSortOrderAsc,
	"desc": ListDrPlanExecutionsSortOrderDesc,
}

// GetListDrPlanExecutionsSortOrderEnumValues Enumerates the set of values for ListDrPlanExecutionsSortOrderEnum
func GetListDrPlanExecutionsSortOrderEnumValues() []ListDrPlanExecutionsSortOrderEnum {
	values := make([]ListDrPlanExecutionsSortOrderEnum, 0)
	for _, v := range mappingListDrPlanExecutionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrPlanExecutionsSortOrderEnumStringValues Enumerates the set of values in String for ListDrPlanExecutionsSortOrderEnum
func GetListDrPlanExecutionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDrPlanExecutionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrPlanExecutionsSortOrderEnum(val string) (ListDrPlanExecutionsSortOrderEnum, bool) {
	enum, ok := mappingListDrPlanExecutionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDrPlanExecutionsSortByEnum Enum with underlying type: string
type ListDrPlanExecutionsSortByEnum string

// Set of constants representing the allowable values for ListDrPlanExecutionsSortByEnum
const (
	ListDrPlanExecutionsSortByTimecreated ListDrPlanExecutionsSortByEnum = "timeCreated"
	ListDrPlanExecutionsSortByDisplayname ListDrPlanExecutionsSortByEnum = "displayName"
)

var mappingListDrPlanExecutionsSortByEnum = map[string]ListDrPlanExecutionsSortByEnum{
	"timeCreated": ListDrPlanExecutionsSortByTimecreated,
	"displayName": ListDrPlanExecutionsSortByDisplayname,
}

var mappingListDrPlanExecutionsSortByEnumLowerCase = map[string]ListDrPlanExecutionsSortByEnum{
	"timecreated": ListDrPlanExecutionsSortByTimecreated,
	"displayname": ListDrPlanExecutionsSortByDisplayname,
}

// GetListDrPlanExecutionsSortByEnumValues Enumerates the set of values for ListDrPlanExecutionsSortByEnum
func GetListDrPlanExecutionsSortByEnumValues() []ListDrPlanExecutionsSortByEnum {
	values := make([]ListDrPlanExecutionsSortByEnum, 0)
	for _, v := range mappingListDrPlanExecutionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDrPlanExecutionsSortByEnumStringValues Enumerates the set of values in String for ListDrPlanExecutionsSortByEnum
func GetListDrPlanExecutionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDrPlanExecutionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDrPlanExecutionsSortByEnum(val string) (ListDrPlanExecutionsSortByEnum, bool) {
	enum, ok := mappingListDrPlanExecutionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
