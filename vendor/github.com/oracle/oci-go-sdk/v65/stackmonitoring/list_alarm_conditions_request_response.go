// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAlarmConditionsRequest wrapper for the ListAlarmConditions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListAlarmConditions.go.html to see an example of how to use ListAlarmConditionsRequest.
type ListAlarmConditionsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitoring template.
	MonitoringTemplateId *string `mandatory:"true" contributesTo:"path" name:"monitoringTemplateId"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeUpdated is descending.
	SortBy ListAlarmConditionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAlarmConditionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return alarm condition based on input status.
	Status ListAlarmConditionsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return alarm condition based on Lifecycle State.
	LifecycleState ListAlarmConditionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Multiple resource types filter.
	ResourceTypes []string `contributesTo:"query" name:"resourceTypes" collectionFormat:"multi"`

	// metricName filter.
	MetricName []string `contributesTo:"query" name:"metricName" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAlarmConditionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAlarmConditionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAlarmConditionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAlarmConditionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAlarmConditionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAlarmConditionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAlarmConditionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlarmConditionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAlarmConditionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlarmConditionsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListAlarmConditionsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlarmConditionsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListAlarmConditionsLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAlarmConditionsResponse wrapper for the ListAlarmConditions operation
type ListAlarmConditionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AlarmConditionCollection instances
	AlarmConditionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAlarmConditionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAlarmConditionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAlarmConditionsSortByEnum Enum with underlying type: string
type ListAlarmConditionsSortByEnum string

// Set of constants representing the allowable values for ListAlarmConditionsSortByEnum
const (
	ListAlarmConditionsSortByMetricname     ListAlarmConditionsSortByEnum = "metricName"
	ListAlarmConditionsSortByLifecyclestate ListAlarmConditionsSortByEnum = "lifeCycleState"
	ListAlarmConditionsSortByResourcetype   ListAlarmConditionsSortByEnum = "resourceType"
	ListAlarmConditionsSortByStatus         ListAlarmConditionsSortByEnum = "status"
	ListAlarmConditionsSortByTimeupdated    ListAlarmConditionsSortByEnum = "timeUpdated"
	ListAlarmConditionsSortByTimecreated    ListAlarmConditionsSortByEnum = "timeCreated"
)

var mappingListAlarmConditionsSortByEnum = map[string]ListAlarmConditionsSortByEnum{
	"metricName":     ListAlarmConditionsSortByMetricname,
	"lifeCycleState": ListAlarmConditionsSortByLifecyclestate,
	"resourceType":   ListAlarmConditionsSortByResourcetype,
	"status":         ListAlarmConditionsSortByStatus,
	"timeUpdated":    ListAlarmConditionsSortByTimeupdated,
	"timeCreated":    ListAlarmConditionsSortByTimecreated,
}

var mappingListAlarmConditionsSortByEnumLowerCase = map[string]ListAlarmConditionsSortByEnum{
	"metricname":     ListAlarmConditionsSortByMetricname,
	"lifecyclestate": ListAlarmConditionsSortByLifecyclestate,
	"resourcetype":   ListAlarmConditionsSortByResourcetype,
	"status":         ListAlarmConditionsSortByStatus,
	"timeupdated":    ListAlarmConditionsSortByTimeupdated,
	"timecreated":    ListAlarmConditionsSortByTimecreated,
}

// GetListAlarmConditionsSortByEnumValues Enumerates the set of values for ListAlarmConditionsSortByEnum
func GetListAlarmConditionsSortByEnumValues() []ListAlarmConditionsSortByEnum {
	values := make([]ListAlarmConditionsSortByEnum, 0)
	for _, v := range mappingListAlarmConditionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmConditionsSortByEnumStringValues Enumerates the set of values in String for ListAlarmConditionsSortByEnum
func GetListAlarmConditionsSortByEnumStringValues() []string {
	return []string{
		"metricName",
		"lifeCycleState",
		"resourceType",
		"status",
		"timeUpdated",
		"timeCreated",
	}
}

// GetMappingListAlarmConditionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmConditionsSortByEnum(val string) (ListAlarmConditionsSortByEnum, bool) {
	enum, ok := mappingListAlarmConditionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlarmConditionsSortOrderEnum Enum with underlying type: string
type ListAlarmConditionsSortOrderEnum string

// Set of constants representing the allowable values for ListAlarmConditionsSortOrderEnum
const (
	ListAlarmConditionsSortOrderAsc  ListAlarmConditionsSortOrderEnum = "ASC"
	ListAlarmConditionsSortOrderDesc ListAlarmConditionsSortOrderEnum = "DESC"
)

var mappingListAlarmConditionsSortOrderEnum = map[string]ListAlarmConditionsSortOrderEnum{
	"ASC":  ListAlarmConditionsSortOrderAsc,
	"DESC": ListAlarmConditionsSortOrderDesc,
}

var mappingListAlarmConditionsSortOrderEnumLowerCase = map[string]ListAlarmConditionsSortOrderEnum{
	"asc":  ListAlarmConditionsSortOrderAsc,
	"desc": ListAlarmConditionsSortOrderDesc,
}

// GetListAlarmConditionsSortOrderEnumValues Enumerates the set of values for ListAlarmConditionsSortOrderEnum
func GetListAlarmConditionsSortOrderEnumValues() []ListAlarmConditionsSortOrderEnum {
	values := make([]ListAlarmConditionsSortOrderEnum, 0)
	for _, v := range mappingListAlarmConditionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmConditionsSortOrderEnumStringValues Enumerates the set of values in String for ListAlarmConditionsSortOrderEnum
func GetListAlarmConditionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAlarmConditionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmConditionsSortOrderEnum(val string) (ListAlarmConditionsSortOrderEnum, bool) {
	enum, ok := mappingListAlarmConditionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlarmConditionsStatusEnum Enum with underlying type: string
type ListAlarmConditionsStatusEnum string

// Set of constants representing the allowable values for ListAlarmConditionsStatusEnum
const (
	ListAlarmConditionsStatusNotApplied     ListAlarmConditionsStatusEnum = "NOT_APPLIED"
	ListAlarmConditionsStatusApplied        ListAlarmConditionsStatusEnum = "APPLIED"
	ListAlarmConditionsStatusPartialApplied ListAlarmConditionsStatusEnum = "PARTIAL_APPLIED"
	ListAlarmConditionsStatusError          ListAlarmConditionsStatusEnum = "ERROR"
)

var mappingListAlarmConditionsStatusEnum = map[string]ListAlarmConditionsStatusEnum{
	"NOT_APPLIED":     ListAlarmConditionsStatusNotApplied,
	"APPLIED":         ListAlarmConditionsStatusApplied,
	"PARTIAL_APPLIED": ListAlarmConditionsStatusPartialApplied,
	"ERROR":           ListAlarmConditionsStatusError,
}

var mappingListAlarmConditionsStatusEnumLowerCase = map[string]ListAlarmConditionsStatusEnum{
	"not_applied":     ListAlarmConditionsStatusNotApplied,
	"applied":         ListAlarmConditionsStatusApplied,
	"partial_applied": ListAlarmConditionsStatusPartialApplied,
	"error":           ListAlarmConditionsStatusError,
}

// GetListAlarmConditionsStatusEnumValues Enumerates the set of values for ListAlarmConditionsStatusEnum
func GetListAlarmConditionsStatusEnumValues() []ListAlarmConditionsStatusEnum {
	values := make([]ListAlarmConditionsStatusEnum, 0)
	for _, v := range mappingListAlarmConditionsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmConditionsStatusEnumStringValues Enumerates the set of values in String for ListAlarmConditionsStatusEnum
func GetListAlarmConditionsStatusEnumStringValues() []string {
	return []string{
		"NOT_APPLIED",
		"APPLIED",
		"PARTIAL_APPLIED",
		"ERROR",
	}
}

// GetMappingListAlarmConditionsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmConditionsStatusEnum(val string) (ListAlarmConditionsStatusEnum, bool) {
	enum, ok := mappingListAlarmConditionsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlarmConditionsLifecycleStateEnum Enum with underlying type: string
type ListAlarmConditionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListAlarmConditionsLifecycleStateEnum
const (
	ListAlarmConditionsLifecycleStateCreating ListAlarmConditionsLifecycleStateEnum = "CREATING"
	ListAlarmConditionsLifecycleStateActive   ListAlarmConditionsLifecycleStateEnum = "ACTIVE"
	ListAlarmConditionsLifecycleStateInactive ListAlarmConditionsLifecycleStateEnum = "INACTIVE"
	ListAlarmConditionsLifecycleStateUpdating ListAlarmConditionsLifecycleStateEnum = "UPDATING"
	ListAlarmConditionsLifecycleStateDeleted  ListAlarmConditionsLifecycleStateEnum = "DELETED"
)

var mappingListAlarmConditionsLifecycleStateEnum = map[string]ListAlarmConditionsLifecycleStateEnum{
	"CREATING": ListAlarmConditionsLifecycleStateCreating,
	"ACTIVE":   ListAlarmConditionsLifecycleStateActive,
	"INACTIVE": ListAlarmConditionsLifecycleStateInactive,
	"UPDATING": ListAlarmConditionsLifecycleStateUpdating,
	"DELETED":  ListAlarmConditionsLifecycleStateDeleted,
}

var mappingListAlarmConditionsLifecycleStateEnumLowerCase = map[string]ListAlarmConditionsLifecycleStateEnum{
	"creating": ListAlarmConditionsLifecycleStateCreating,
	"active":   ListAlarmConditionsLifecycleStateActive,
	"inactive": ListAlarmConditionsLifecycleStateInactive,
	"updating": ListAlarmConditionsLifecycleStateUpdating,
	"deleted":  ListAlarmConditionsLifecycleStateDeleted,
}

// GetListAlarmConditionsLifecycleStateEnumValues Enumerates the set of values for ListAlarmConditionsLifecycleStateEnum
func GetListAlarmConditionsLifecycleStateEnumValues() []ListAlarmConditionsLifecycleStateEnum {
	values := make([]ListAlarmConditionsLifecycleStateEnum, 0)
	for _, v := range mappingListAlarmConditionsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmConditionsLifecycleStateEnumStringValues Enumerates the set of values in String for ListAlarmConditionsLifecycleStateEnum
func GetListAlarmConditionsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETED",
	}
}

// GetMappingListAlarmConditionsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmConditionsLifecycleStateEnum(val string) (ListAlarmConditionsLifecycleStateEnum, bool) {
	enum, ok := mappingListAlarmConditionsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
