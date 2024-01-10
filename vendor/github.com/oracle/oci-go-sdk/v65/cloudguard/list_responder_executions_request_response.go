// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListResponderExecutionsRequest wrapper for the ListResponderExecutions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResponderExecutions.go.html to see an example of how to use ListResponderExecutionsRequest.
type ListResponderExecutionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListResponderExecutionsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Responder Rule Ids filter for the Responder Executions.
	ResponderRuleIds []string `contributesTo:"query" name:"responderRuleIds" collectionFormat:"multi"`

	// Creation Start time for filtering
	TimeCreatedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedGreaterThanOrEqualTo"`

	// Creation End time for filtering
	TimeCreatedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreatedLessThanOrEqualTo"`

	// Completion End Time
	TimeCompletedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCompletedGreaterThanOrEqualTo"`

	// Completion Start Time
	TimeCompletedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCompletedLessThanOrEqualTo"`

	// The ID of the target in which to list resources.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// Resource Type associated with the resource.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// The field to list the Responder Executions by Responder Type. Valid values are REMEDIATION and NOTIFICATION
	ResponderType ListResponderExecutionsResponderTypeEnum `mandatory:"false" contributesTo:"query" name:"responderType" omitEmpty:"true"`

	// The status of the responder execution in which to list responders.
	ResponderExecutionStatus ListResponderExecutionsResponderExecutionStatusEnum `mandatory:"false" contributesTo:"query" name:"responderExecutionStatus" omitEmpty:"true"`

	// The mode of the responder execution in which to list responders.
	ResponderExecutionMode ListResponderExecutionsResponderExecutionModeEnum `mandatory:"false" contributesTo:"query" name:"responderExecutionMode" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResponderExecutionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for responderRuleName and resourceName is ascending. If no value is specified timeCreated is default.
	SortBy ListResponderExecutionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResponderExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResponderExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResponderExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResponderExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListResponderExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListResponderExecutionsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListResponderExecutionsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderExecutionsResponderTypeEnum(string(request.ResponderType)); !ok && request.ResponderType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponderType: %s. Supported values are: %s.", request.ResponderType, strings.Join(GetListResponderExecutionsResponderTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderExecutionsResponderExecutionStatusEnum(string(request.ResponderExecutionStatus)); !ok && request.ResponderExecutionStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponderExecutionStatus: %s. Supported values are: %s.", request.ResponderExecutionStatus, strings.Join(GetListResponderExecutionsResponderExecutionStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderExecutionsResponderExecutionModeEnum(string(request.ResponderExecutionMode)); !ok && request.ResponderExecutionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResponderExecutionMode: %s. Supported values are: %s.", request.ResponderExecutionMode, strings.Join(GetListResponderExecutionsResponderExecutionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderExecutionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListResponderExecutionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListResponderExecutionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListResponderExecutionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListResponderExecutionsResponse wrapper for the ListResponderExecutions operation
type ListResponderExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResponderExecutionCollection instances
	ResponderExecutionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResponderExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResponderExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResponderExecutionsAccessLevelEnum Enum with underlying type: string
type ListResponderExecutionsAccessLevelEnum string

// Set of constants representing the allowable values for ListResponderExecutionsAccessLevelEnum
const (
	ListResponderExecutionsAccessLevelRestricted ListResponderExecutionsAccessLevelEnum = "RESTRICTED"
	ListResponderExecutionsAccessLevelAccessible ListResponderExecutionsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListResponderExecutionsAccessLevelEnum = map[string]ListResponderExecutionsAccessLevelEnum{
	"RESTRICTED": ListResponderExecutionsAccessLevelRestricted,
	"ACCESSIBLE": ListResponderExecutionsAccessLevelAccessible,
}

var mappingListResponderExecutionsAccessLevelEnumLowerCase = map[string]ListResponderExecutionsAccessLevelEnum{
	"restricted": ListResponderExecutionsAccessLevelRestricted,
	"accessible": ListResponderExecutionsAccessLevelAccessible,
}

// GetListResponderExecutionsAccessLevelEnumValues Enumerates the set of values for ListResponderExecutionsAccessLevelEnum
func GetListResponderExecutionsAccessLevelEnumValues() []ListResponderExecutionsAccessLevelEnum {
	values := make([]ListResponderExecutionsAccessLevelEnum, 0)
	for _, v := range mappingListResponderExecutionsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderExecutionsAccessLevelEnumStringValues Enumerates the set of values in String for ListResponderExecutionsAccessLevelEnum
func GetListResponderExecutionsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListResponderExecutionsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderExecutionsAccessLevelEnum(val string) (ListResponderExecutionsAccessLevelEnum, bool) {
	enum, ok := mappingListResponderExecutionsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderExecutionsResponderTypeEnum Enum with underlying type: string
type ListResponderExecutionsResponderTypeEnum string

// Set of constants representing the allowable values for ListResponderExecutionsResponderTypeEnum
const (
	ListResponderExecutionsResponderTypeRemediation  ListResponderExecutionsResponderTypeEnum = "REMEDIATION"
	ListResponderExecutionsResponderTypeNotification ListResponderExecutionsResponderTypeEnum = "NOTIFICATION"
)

var mappingListResponderExecutionsResponderTypeEnum = map[string]ListResponderExecutionsResponderTypeEnum{
	"REMEDIATION":  ListResponderExecutionsResponderTypeRemediation,
	"NOTIFICATION": ListResponderExecutionsResponderTypeNotification,
}

var mappingListResponderExecutionsResponderTypeEnumLowerCase = map[string]ListResponderExecutionsResponderTypeEnum{
	"remediation":  ListResponderExecutionsResponderTypeRemediation,
	"notification": ListResponderExecutionsResponderTypeNotification,
}

// GetListResponderExecutionsResponderTypeEnumValues Enumerates the set of values for ListResponderExecutionsResponderTypeEnum
func GetListResponderExecutionsResponderTypeEnumValues() []ListResponderExecutionsResponderTypeEnum {
	values := make([]ListResponderExecutionsResponderTypeEnum, 0)
	for _, v := range mappingListResponderExecutionsResponderTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderExecutionsResponderTypeEnumStringValues Enumerates the set of values in String for ListResponderExecutionsResponderTypeEnum
func GetListResponderExecutionsResponderTypeEnumStringValues() []string {
	return []string{
		"REMEDIATION",
		"NOTIFICATION",
	}
}

// GetMappingListResponderExecutionsResponderTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderExecutionsResponderTypeEnum(val string) (ListResponderExecutionsResponderTypeEnum, bool) {
	enum, ok := mappingListResponderExecutionsResponderTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderExecutionsResponderExecutionStatusEnum Enum with underlying type: string
type ListResponderExecutionsResponderExecutionStatusEnum string

// Set of constants representing the allowable values for ListResponderExecutionsResponderExecutionStatusEnum
const (
	ListResponderExecutionsResponderExecutionStatusStarted              ListResponderExecutionsResponderExecutionStatusEnum = "STARTED"
	ListResponderExecutionsResponderExecutionStatusAwaitingConfirmation ListResponderExecutionsResponderExecutionStatusEnum = "AWAITING_CONFIRMATION"
	ListResponderExecutionsResponderExecutionStatusAwaitingInput        ListResponderExecutionsResponderExecutionStatusEnum = "AWAITING_INPUT"
	ListResponderExecutionsResponderExecutionStatusSucceeded            ListResponderExecutionsResponderExecutionStatusEnum = "SUCCEEDED"
	ListResponderExecutionsResponderExecutionStatusFailed               ListResponderExecutionsResponderExecutionStatusEnum = "FAILED"
	ListResponderExecutionsResponderExecutionStatusSkipped              ListResponderExecutionsResponderExecutionStatusEnum = "SKIPPED"
	ListResponderExecutionsResponderExecutionStatusAll                  ListResponderExecutionsResponderExecutionStatusEnum = "ALL"
)

var mappingListResponderExecutionsResponderExecutionStatusEnum = map[string]ListResponderExecutionsResponderExecutionStatusEnum{
	"STARTED":               ListResponderExecutionsResponderExecutionStatusStarted,
	"AWAITING_CONFIRMATION": ListResponderExecutionsResponderExecutionStatusAwaitingConfirmation,
	"AWAITING_INPUT":        ListResponderExecutionsResponderExecutionStatusAwaitingInput,
	"SUCCEEDED":             ListResponderExecutionsResponderExecutionStatusSucceeded,
	"FAILED":                ListResponderExecutionsResponderExecutionStatusFailed,
	"SKIPPED":               ListResponderExecutionsResponderExecutionStatusSkipped,
	"ALL":                   ListResponderExecutionsResponderExecutionStatusAll,
}

var mappingListResponderExecutionsResponderExecutionStatusEnumLowerCase = map[string]ListResponderExecutionsResponderExecutionStatusEnum{
	"started":               ListResponderExecutionsResponderExecutionStatusStarted,
	"awaiting_confirmation": ListResponderExecutionsResponderExecutionStatusAwaitingConfirmation,
	"awaiting_input":        ListResponderExecutionsResponderExecutionStatusAwaitingInput,
	"succeeded":             ListResponderExecutionsResponderExecutionStatusSucceeded,
	"failed":                ListResponderExecutionsResponderExecutionStatusFailed,
	"skipped":               ListResponderExecutionsResponderExecutionStatusSkipped,
	"all":                   ListResponderExecutionsResponderExecutionStatusAll,
}

// GetListResponderExecutionsResponderExecutionStatusEnumValues Enumerates the set of values for ListResponderExecutionsResponderExecutionStatusEnum
func GetListResponderExecutionsResponderExecutionStatusEnumValues() []ListResponderExecutionsResponderExecutionStatusEnum {
	values := make([]ListResponderExecutionsResponderExecutionStatusEnum, 0)
	for _, v := range mappingListResponderExecutionsResponderExecutionStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderExecutionsResponderExecutionStatusEnumStringValues Enumerates the set of values in String for ListResponderExecutionsResponderExecutionStatusEnum
func GetListResponderExecutionsResponderExecutionStatusEnumStringValues() []string {
	return []string{
		"STARTED",
		"AWAITING_CONFIRMATION",
		"AWAITING_INPUT",
		"SUCCEEDED",
		"FAILED",
		"SKIPPED",
		"ALL",
	}
}

// GetMappingListResponderExecutionsResponderExecutionStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderExecutionsResponderExecutionStatusEnum(val string) (ListResponderExecutionsResponderExecutionStatusEnum, bool) {
	enum, ok := mappingListResponderExecutionsResponderExecutionStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderExecutionsResponderExecutionModeEnum Enum with underlying type: string
type ListResponderExecutionsResponderExecutionModeEnum string

// Set of constants representing the allowable values for ListResponderExecutionsResponderExecutionModeEnum
const (
	ListResponderExecutionsResponderExecutionModeManual    ListResponderExecutionsResponderExecutionModeEnum = "MANUAL"
	ListResponderExecutionsResponderExecutionModeAutomated ListResponderExecutionsResponderExecutionModeEnum = "AUTOMATED"
	ListResponderExecutionsResponderExecutionModeAll       ListResponderExecutionsResponderExecutionModeEnum = "ALL"
)

var mappingListResponderExecutionsResponderExecutionModeEnum = map[string]ListResponderExecutionsResponderExecutionModeEnum{
	"MANUAL":    ListResponderExecutionsResponderExecutionModeManual,
	"AUTOMATED": ListResponderExecutionsResponderExecutionModeAutomated,
	"ALL":       ListResponderExecutionsResponderExecutionModeAll,
}

var mappingListResponderExecutionsResponderExecutionModeEnumLowerCase = map[string]ListResponderExecutionsResponderExecutionModeEnum{
	"manual":    ListResponderExecutionsResponderExecutionModeManual,
	"automated": ListResponderExecutionsResponderExecutionModeAutomated,
	"all":       ListResponderExecutionsResponderExecutionModeAll,
}

// GetListResponderExecutionsResponderExecutionModeEnumValues Enumerates the set of values for ListResponderExecutionsResponderExecutionModeEnum
func GetListResponderExecutionsResponderExecutionModeEnumValues() []ListResponderExecutionsResponderExecutionModeEnum {
	values := make([]ListResponderExecutionsResponderExecutionModeEnum, 0)
	for _, v := range mappingListResponderExecutionsResponderExecutionModeEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderExecutionsResponderExecutionModeEnumStringValues Enumerates the set of values in String for ListResponderExecutionsResponderExecutionModeEnum
func GetListResponderExecutionsResponderExecutionModeEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTOMATED",
		"ALL",
	}
}

// GetMappingListResponderExecutionsResponderExecutionModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderExecutionsResponderExecutionModeEnum(val string) (ListResponderExecutionsResponderExecutionModeEnum, bool) {
	enum, ok := mappingListResponderExecutionsResponderExecutionModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderExecutionsSortOrderEnum Enum with underlying type: string
type ListResponderExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListResponderExecutionsSortOrderEnum
const (
	ListResponderExecutionsSortOrderAsc  ListResponderExecutionsSortOrderEnum = "ASC"
	ListResponderExecutionsSortOrderDesc ListResponderExecutionsSortOrderEnum = "DESC"
)

var mappingListResponderExecutionsSortOrderEnum = map[string]ListResponderExecutionsSortOrderEnum{
	"ASC":  ListResponderExecutionsSortOrderAsc,
	"DESC": ListResponderExecutionsSortOrderDesc,
}

var mappingListResponderExecutionsSortOrderEnumLowerCase = map[string]ListResponderExecutionsSortOrderEnum{
	"asc":  ListResponderExecutionsSortOrderAsc,
	"desc": ListResponderExecutionsSortOrderDesc,
}

// GetListResponderExecutionsSortOrderEnumValues Enumerates the set of values for ListResponderExecutionsSortOrderEnum
func GetListResponderExecutionsSortOrderEnumValues() []ListResponderExecutionsSortOrderEnum {
	values := make([]ListResponderExecutionsSortOrderEnum, 0)
	for _, v := range mappingListResponderExecutionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderExecutionsSortOrderEnumStringValues Enumerates the set of values in String for ListResponderExecutionsSortOrderEnum
func GetListResponderExecutionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListResponderExecutionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderExecutionsSortOrderEnum(val string) (ListResponderExecutionsSortOrderEnum, bool) {
	enum, ok := mappingListResponderExecutionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListResponderExecutionsSortByEnum Enum with underlying type: string
type ListResponderExecutionsSortByEnum string

// Set of constants representing the allowable values for ListResponderExecutionsSortByEnum
const (
	ListResponderExecutionsSortByTimecreated       ListResponderExecutionsSortByEnum = "timeCreated"
	ListResponderExecutionsSortByResponderrulename ListResponderExecutionsSortByEnum = "responderRuleName"
	ListResponderExecutionsSortByResourcename      ListResponderExecutionsSortByEnum = "resourceName"
	ListResponderExecutionsSortByTimecompleted     ListResponderExecutionsSortByEnum = "timeCompleted"
)

var mappingListResponderExecutionsSortByEnum = map[string]ListResponderExecutionsSortByEnum{
	"timeCreated":       ListResponderExecutionsSortByTimecreated,
	"responderRuleName": ListResponderExecutionsSortByResponderrulename,
	"resourceName":      ListResponderExecutionsSortByResourcename,
	"timeCompleted":     ListResponderExecutionsSortByTimecompleted,
}

var mappingListResponderExecutionsSortByEnumLowerCase = map[string]ListResponderExecutionsSortByEnum{
	"timecreated":       ListResponderExecutionsSortByTimecreated,
	"responderrulename": ListResponderExecutionsSortByResponderrulename,
	"resourcename":      ListResponderExecutionsSortByResourcename,
	"timecompleted":     ListResponderExecutionsSortByTimecompleted,
}

// GetListResponderExecutionsSortByEnumValues Enumerates the set of values for ListResponderExecutionsSortByEnum
func GetListResponderExecutionsSortByEnumValues() []ListResponderExecutionsSortByEnum {
	values := make([]ListResponderExecutionsSortByEnum, 0)
	for _, v := range mappingListResponderExecutionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListResponderExecutionsSortByEnumStringValues Enumerates the set of values in String for ListResponderExecutionsSortByEnum
func GetListResponderExecutionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"responderRuleName",
		"resourceName",
		"timeCompleted",
	}
}

// GetMappingListResponderExecutionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListResponderExecutionsSortByEnum(val string) (ListResponderExecutionsSortByEnum, bool) {
	enum, ok := mappingListResponderExecutionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
