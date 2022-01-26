// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListResponderExecutionsRequest wrapper for the ListResponderExecutions operation
//
// See also
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

var mappingListResponderExecutionsAccessLevel = map[string]ListResponderExecutionsAccessLevelEnum{
	"RESTRICTED": ListResponderExecutionsAccessLevelRestricted,
	"ACCESSIBLE": ListResponderExecutionsAccessLevelAccessible,
}

// GetListResponderExecutionsAccessLevelEnumValues Enumerates the set of values for ListResponderExecutionsAccessLevelEnum
func GetListResponderExecutionsAccessLevelEnumValues() []ListResponderExecutionsAccessLevelEnum {
	values := make([]ListResponderExecutionsAccessLevelEnum, 0)
	for _, v := range mappingListResponderExecutionsAccessLevel {
		values = append(values, v)
	}
	return values
}

// ListResponderExecutionsResponderTypeEnum Enum with underlying type: string
type ListResponderExecutionsResponderTypeEnum string

// Set of constants representing the allowable values for ListResponderExecutionsResponderTypeEnum
const (
	ListResponderExecutionsResponderTypeRemediation  ListResponderExecutionsResponderTypeEnum = "REMEDIATION"
	ListResponderExecutionsResponderTypeNotification ListResponderExecutionsResponderTypeEnum = "NOTIFICATION"
)

var mappingListResponderExecutionsResponderType = map[string]ListResponderExecutionsResponderTypeEnum{
	"REMEDIATION":  ListResponderExecutionsResponderTypeRemediation,
	"NOTIFICATION": ListResponderExecutionsResponderTypeNotification,
}

// GetListResponderExecutionsResponderTypeEnumValues Enumerates the set of values for ListResponderExecutionsResponderTypeEnum
func GetListResponderExecutionsResponderTypeEnumValues() []ListResponderExecutionsResponderTypeEnum {
	values := make([]ListResponderExecutionsResponderTypeEnum, 0)
	for _, v := range mappingListResponderExecutionsResponderType {
		values = append(values, v)
	}
	return values
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

var mappingListResponderExecutionsResponderExecutionStatus = map[string]ListResponderExecutionsResponderExecutionStatusEnum{
	"STARTED":               ListResponderExecutionsResponderExecutionStatusStarted,
	"AWAITING_CONFIRMATION": ListResponderExecutionsResponderExecutionStatusAwaitingConfirmation,
	"AWAITING_INPUT":        ListResponderExecutionsResponderExecutionStatusAwaitingInput,
	"SUCCEEDED":             ListResponderExecutionsResponderExecutionStatusSucceeded,
	"FAILED":                ListResponderExecutionsResponderExecutionStatusFailed,
	"SKIPPED":               ListResponderExecutionsResponderExecutionStatusSkipped,
	"ALL":                   ListResponderExecutionsResponderExecutionStatusAll,
}

// GetListResponderExecutionsResponderExecutionStatusEnumValues Enumerates the set of values for ListResponderExecutionsResponderExecutionStatusEnum
func GetListResponderExecutionsResponderExecutionStatusEnumValues() []ListResponderExecutionsResponderExecutionStatusEnum {
	values := make([]ListResponderExecutionsResponderExecutionStatusEnum, 0)
	for _, v := range mappingListResponderExecutionsResponderExecutionStatus {
		values = append(values, v)
	}
	return values
}

// ListResponderExecutionsResponderExecutionModeEnum Enum with underlying type: string
type ListResponderExecutionsResponderExecutionModeEnum string

// Set of constants representing the allowable values for ListResponderExecutionsResponderExecutionModeEnum
const (
	ListResponderExecutionsResponderExecutionModeManual    ListResponderExecutionsResponderExecutionModeEnum = "MANUAL"
	ListResponderExecutionsResponderExecutionModeAutomated ListResponderExecutionsResponderExecutionModeEnum = "AUTOMATED"
	ListResponderExecutionsResponderExecutionModeAll       ListResponderExecutionsResponderExecutionModeEnum = "ALL"
)

var mappingListResponderExecutionsResponderExecutionMode = map[string]ListResponderExecutionsResponderExecutionModeEnum{
	"MANUAL":    ListResponderExecutionsResponderExecutionModeManual,
	"AUTOMATED": ListResponderExecutionsResponderExecutionModeAutomated,
	"ALL":       ListResponderExecutionsResponderExecutionModeAll,
}

// GetListResponderExecutionsResponderExecutionModeEnumValues Enumerates the set of values for ListResponderExecutionsResponderExecutionModeEnum
func GetListResponderExecutionsResponderExecutionModeEnumValues() []ListResponderExecutionsResponderExecutionModeEnum {
	values := make([]ListResponderExecutionsResponderExecutionModeEnum, 0)
	for _, v := range mappingListResponderExecutionsResponderExecutionMode {
		values = append(values, v)
	}
	return values
}

// ListResponderExecutionsSortOrderEnum Enum with underlying type: string
type ListResponderExecutionsSortOrderEnum string

// Set of constants representing the allowable values for ListResponderExecutionsSortOrderEnum
const (
	ListResponderExecutionsSortOrderAsc  ListResponderExecutionsSortOrderEnum = "ASC"
	ListResponderExecutionsSortOrderDesc ListResponderExecutionsSortOrderEnum = "DESC"
)

var mappingListResponderExecutionsSortOrder = map[string]ListResponderExecutionsSortOrderEnum{
	"ASC":  ListResponderExecutionsSortOrderAsc,
	"DESC": ListResponderExecutionsSortOrderDesc,
}

// GetListResponderExecutionsSortOrderEnumValues Enumerates the set of values for ListResponderExecutionsSortOrderEnum
func GetListResponderExecutionsSortOrderEnumValues() []ListResponderExecutionsSortOrderEnum {
	values := make([]ListResponderExecutionsSortOrderEnum, 0)
	for _, v := range mappingListResponderExecutionsSortOrder {
		values = append(values, v)
	}
	return values
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

var mappingListResponderExecutionsSortBy = map[string]ListResponderExecutionsSortByEnum{
	"timeCreated":       ListResponderExecutionsSortByTimecreated,
	"responderRuleName": ListResponderExecutionsSortByResponderrulename,
	"resourceName":      ListResponderExecutionsSortByResourcename,
	"timeCompleted":     ListResponderExecutionsSortByTimecompleted,
}

// GetListResponderExecutionsSortByEnumValues Enumerates the set of values for ListResponderExecutionsSortByEnum
func GetListResponderExecutionsSortByEnumValues() []ListResponderExecutionsSortByEnum {
	values := make([]ListResponderExecutionsSortByEnum, 0)
	for _, v := range mappingListResponderExecutionsSortBy {
		values = append(values, v)
	}
	return values
}
