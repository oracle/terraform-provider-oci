// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAlarmSuppressionsRequest wrapper for the ListAlarmSuppressions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/monitoring/ListAlarmSuppressions.go.html to see an example of how to use ListAlarmSuppressionsRequest.
type ListAlarmSuppressionsRequest struct {

	// Customer part of the request identifier token. If you need to contact Oracle about a particular
	// request, please provide the complete request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm that is the target of the alarm suppression.
	AlarmId *string `mandatory:"false" contributesTo:"query" name:"alarmId"`

	// A filter to return only resources that match the given display name exactly.
	// Use this filter to list an alarm suppression by name.
	// Alternatively, when you know the alarm suppression OCID, use the GetAlarmSuppression operation.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state exactly. When not specified, only resources in the ACTIVE lifecycle state are listed.
	LifecycleState AlarmSuppressionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The level of this alarm suppression.
	// `ALARM` indicates a suppression of the entire alarm, regardless of dimension.
	// `DIMENSION` indicates a suppression configured for specified dimensions.
	Level AlarmSuppressionLevelEnum `mandatory:"false" contributesTo:"query" name:"level" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment for searching.
	// Use the tenancy OCID to search in the root compartment.
	// If targetType is not specified, searches all suppressions defined under the compartment.
	// If targetType is `COMPARTMENT`, searches suppressions in the specified compartment only.
	// Example: `ocid1.compartment.oc1..exampleuniqueID`
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// When true, returns resources from all compartments and subcompartments. The parameter can
	// only be set to true when compartmentId is the tenancy OCID (the tenancy is the root compartment).
	// A true value requires the user to have tenancy-level permissions. If this requirement is not met,
	// then the call is rejected. When false, returns resources from only the compartment specified in
	// compartmentId. Default is false.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The target type to use when listing alarm suppressions.
	// `ALARM` lists all suppression records for the specified alarm.
	// `COMPARTMENT` lists all suppression records for the specified compartment or tenancy.
	TargetType ListAlarmSuppressionsTargetTypeEnum `mandatory:"false" contributesTo:"query" name:"targetType" omitEmpty:"true"`

	// Setting this parameter to true requires the query to specify the alarm (`alarmId`).
	// When true, lists all alarm suppressions that affect the specified alarm,
	// including suppressions that target the corresponding compartment or tenancy.
	// When false, lists only the alarm suppressions that target the specified alarm.
	// Default is false.
	IsAllSuppressions *bool `mandatory:"false" contributesTo:"query" name:"isAllSuppressions"`

	// The field to use when sorting returned alarm suppressions. Only one sorting level is provided.
	// Example: `timeCreated`
	SortBy ListAlarmSuppressionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use when sorting returned alarm suppressions. Ascending (ASC) or descending (DESC).
	// Example: `ASC`
	SortOrder ListAlarmSuppressionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Default: 1000
	// Example: 500
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAlarmSuppressionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAlarmSuppressionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAlarmSuppressionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAlarmSuppressionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAlarmSuppressionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlarmSuppressionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetAlarmSuppressionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlarmSuppressionLevelEnum(string(request.Level)); !ok && request.Level != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Level: %s. Supported values are: %s.", request.Level, strings.Join(GetAlarmSuppressionLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlarmSuppressionsTargetTypeEnum(string(request.TargetType)); !ok && request.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", request.TargetType, strings.Join(GetListAlarmSuppressionsTargetTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlarmSuppressionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAlarmSuppressionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlarmSuppressionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAlarmSuppressionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAlarmSuppressionsResponse wrapper for the ListAlarmSuppressions operation
type ListAlarmSuppressionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AlarmSuppressionCollection instances
	AlarmSuppressionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, next page of results remains.
	// For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results remains.
	// For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`
}

func (response ListAlarmSuppressionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAlarmSuppressionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAlarmSuppressionsTargetTypeEnum Enum with underlying type: string
type ListAlarmSuppressionsTargetTypeEnum string

// Set of constants representing the allowable values for ListAlarmSuppressionsTargetTypeEnum
const (
	ListAlarmSuppressionsTargetTypeAlarm       ListAlarmSuppressionsTargetTypeEnum = "ALARM"
	ListAlarmSuppressionsTargetTypeCompartment ListAlarmSuppressionsTargetTypeEnum = "COMPARTMENT"
)

var mappingListAlarmSuppressionsTargetTypeEnum = map[string]ListAlarmSuppressionsTargetTypeEnum{
	"ALARM":       ListAlarmSuppressionsTargetTypeAlarm,
	"COMPARTMENT": ListAlarmSuppressionsTargetTypeCompartment,
}

var mappingListAlarmSuppressionsTargetTypeEnumLowerCase = map[string]ListAlarmSuppressionsTargetTypeEnum{
	"alarm":       ListAlarmSuppressionsTargetTypeAlarm,
	"compartment": ListAlarmSuppressionsTargetTypeCompartment,
}

// GetListAlarmSuppressionsTargetTypeEnumValues Enumerates the set of values for ListAlarmSuppressionsTargetTypeEnum
func GetListAlarmSuppressionsTargetTypeEnumValues() []ListAlarmSuppressionsTargetTypeEnum {
	values := make([]ListAlarmSuppressionsTargetTypeEnum, 0)
	for _, v := range mappingListAlarmSuppressionsTargetTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmSuppressionsTargetTypeEnumStringValues Enumerates the set of values in String for ListAlarmSuppressionsTargetTypeEnum
func GetListAlarmSuppressionsTargetTypeEnumStringValues() []string {
	return []string{
		"ALARM",
		"COMPARTMENT",
	}
}

// GetMappingListAlarmSuppressionsTargetTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmSuppressionsTargetTypeEnum(val string) (ListAlarmSuppressionsTargetTypeEnum, bool) {
	enum, ok := mappingListAlarmSuppressionsTargetTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlarmSuppressionsSortByEnum Enum with underlying type: string
type ListAlarmSuppressionsSortByEnum string

// Set of constants representing the allowable values for ListAlarmSuppressionsSortByEnum
const (
	ListAlarmSuppressionsSortByDisplayname      ListAlarmSuppressionsSortByEnum = "displayName"
	ListAlarmSuppressionsSortByTimecreated      ListAlarmSuppressionsSortByEnum = "timeCreated"
	ListAlarmSuppressionsSortByTimesuppressfrom ListAlarmSuppressionsSortByEnum = "timeSuppressFrom"
)

var mappingListAlarmSuppressionsSortByEnum = map[string]ListAlarmSuppressionsSortByEnum{
	"displayName":      ListAlarmSuppressionsSortByDisplayname,
	"timeCreated":      ListAlarmSuppressionsSortByTimecreated,
	"timeSuppressFrom": ListAlarmSuppressionsSortByTimesuppressfrom,
}

var mappingListAlarmSuppressionsSortByEnumLowerCase = map[string]ListAlarmSuppressionsSortByEnum{
	"displayname":      ListAlarmSuppressionsSortByDisplayname,
	"timecreated":      ListAlarmSuppressionsSortByTimecreated,
	"timesuppressfrom": ListAlarmSuppressionsSortByTimesuppressfrom,
}

// GetListAlarmSuppressionsSortByEnumValues Enumerates the set of values for ListAlarmSuppressionsSortByEnum
func GetListAlarmSuppressionsSortByEnumValues() []ListAlarmSuppressionsSortByEnum {
	values := make([]ListAlarmSuppressionsSortByEnum, 0)
	for _, v := range mappingListAlarmSuppressionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmSuppressionsSortByEnumStringValues Enumerates the set of values in String for ListAlarmSuppressionsSortByEnum
func GetListAlarmSuppressionsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
		"timeSuppressFrom",
	}
}

// GetMappingListAlarmSuppressionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmSuppressionsSortByEnum(val string) (ListAlarmSuppressionsSortByEnum, bool) {
	enum, ok := mappingListAlarmSuppressionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlarmSuppressionsSortOrderEnum Enum with underlying type: string
type ListAlarmSuppressionsSortOrderEnum string

// Set of constants representing the allowable values for ListAlarmSuppressionsSortOrderEnum
const (
	ListAlarmSuppressionsSortOrderAsc  ListAlarmSuppressionsSortOrderEnum = "ASC"
	ListAlarmSuppressionsSortOrderDesc ListAlarmSuppressionsSortOrderEnum = "DESC"
)

var mappingListAlarmSuppressionsSortOrderEnum = map[string]ListAlarmSuppressionsSortOrderEnum{
	"ASC":  ListAlarmSuppressionsSortOrderAsc,
	"DESC": ListAlarmSuppressionsSortOrderDesc,
}

var mappingListAlarmSuppressionsSortOrderEnumLowerCase = map[string]ListAlarmSuppressionsSortOrderEnum{
	"asc":  ListAlarmSuppressionsSortOrderAsc,
	"desc": ListAlarmSuppressionsSortOrderDesc,
}

// GetListAlarmSuppressionsSortOrderEnumValues Enumerates the set of values for ListAlarmSuppressionsSortOrderEnum
func GetListAlarmSuppressionsSortOrderEnumValues() []ListAlarmSuppressionsSortOrderEnum {
	values := make([]ListAlarmSuppressionsSortOrderEnum, 0)
	for _, v := range mappingListAlarmSuppressionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlarmSuppressionsSortOrderEnumStringValues Enumerates the set of values in String for ListAlarmSuppressionsSortOrderEnum
func GetListAlarmSuppressionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAlarmSuppressionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlarmSuppressionsSortOrderEnum(val string) (ListAlarmSuppressionsSortOrderEnum, bool) {
	enum, ok := mappingListAlarmSuppressionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
