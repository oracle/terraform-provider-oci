// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEventsRequest wrapper for the ListEvents operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListEvents.go.html to see an example of how to use ListEventsRequest.
type ListEventsRequest struct {

	// Instance Oracle Cloud identifier (ocid)
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique event identifier (OCID)
	EventId *string `mandatory:"false" contributesTo:"query" name:"eventId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only event of given type.
	EventType ListEventsEventTypeEnum `mandatory:"false" contributesTo:"query" name:"eventType" omitEmpty:"true"`

	// filter event occurrence. Selecting only those last occurred before given date in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	LatestTimestampLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"latestTimestampLessThan"`

	// filter event occurrence. Selecting only those last occurred on or after given date in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	LatestTimestampGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"latestTimestampGreaterThanOrEqualTo"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEventsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEventsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEventsEventTypeEnum(string(request.EventType)); !ok && request.EventType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EventType: %s. Supported values are: %s.", request.EventType, strings.Join(GetListEventsEventTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEventsResponse wrapper for the ListEvents operation
type ListEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EventCollection instances
	EventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEventsSortOrderEnum Enum with underlying type: string
type ListEventsSortOrderEnum string

// Set of constants representing the allowable values for ListEventsSortOrderEnum
const (
	ListEventsSortOrderAsc  ListEventsSortOrderEnum = "ASC"
	ListEventsSortOrderDesc ListEventsSortOrderEnum = "DESC"
)

var mappingListEventsSortOrderEnum = map[string]ListEventsSortOrderEnum{
	"ASC":  ListEventsSortOrderAsc,
	"DESC": ListEventsSortOrderDesc,
}

var mappingListEventsSortOrderEnumLowerCase = map[string]ListEventsSortOrderEnum{
	"asc":  ListEventsSortOrderAsc,
	"desc": ListEventsSortOrderDesc,
}

// GetListEventsSortOrderEnumValues Enumerates the set of values for ListEventsSortOrderEnum
func GetListEventsSortOrderEnumValues() []ListEventsSortOrderEnum {
	values := make([]ListEventsSortOrderEnum, 0)
	for _, v := range mappingListEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEventsSortOrderEnumStringValues Enumerates the set of values in String for ListEventsSortOrderEnum
func GetListEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEventsSortOrderEnum(val string) (ListEventsSortOrderEnum, bool) {
	enum, ok := mappingListEventsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEventsSortByEnum Enum with underlying type: string
type ListEventsSortByEnum string

// Set of constants representing the allowable values for ListEventsSortByEnum
const (
	ListEventsSortByTimecreated ListEventsSortByEnum = "TIMECREATED"
	ListEventsSortByDisplayname ListEventsSortByEnum = "DISPLAYNAME"
)

var mappingListEventsSortByEnum = map[string]ListEventsSortByEnum{
	"TIMECREATED": ListEventsSortByTimecreated,
	"DISPLAYNAME": ListEventsSortByDisplayname,
}

var mappingListEventsSortByEnumLowerCase = map[string]ListEventsSortByEnum{
	"timecreated": ListEventsSortByTimecreated,
	"displayname": ListEventsSortByDisplayname,
}

// GetListEventsSortByEnumValues Enumerates the set of values for ListEventsSortByEnum
func GetListEventsSortByEnumValues() []ListEventsSortByEnum {
	values := make([]ListEventsSortByEnum, 0)
	for _, v := range mappingListEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEventsSortByEnumStringValues Enumerates the set of values in String for ListEventsSortByEnum
func GetListEventsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEventsSortByEnum(val string) (ListEventsSortByEnum, bool) {
	enum, ok := mappingListEventsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEventsEventTypeEnum Enum with underlying type: string
type ListEventsEventTypeEnum string

// Set of constants representing the allowable values for ListEventsEventTypeEnum
const (
	ListEventsEventTypeKernelOops       ListEventsEventTypeEnum = "KERNEL_OOPS"
	ListEventsEventTypeKernelCrash      ListEventsEventTypeEnum = "KERNEL_CRASH"
	ListEventsEventTypeCrash            ListEventsEventTypeEnum = "CRASH"
	ListEventsEventTypeExploitAttempt   ListEventsEventTypeEnum = "EXPLOIT_ATTEMPT"
	ListEventsEventTypeCompliance       ListEventsEventTypeEnum = "COMPLIANCE"
	ListEventsEventTypeTuningSuggestion ListEventsEventTypeEnum = "TUNING_SUGGESTION"
	ListEventsEventTypeTuningApplied    ListEventsEventTypeEnum = "TUNING_APPLIED"
	ListEventsEventTypeSecurity         ListEventsEventTypeEnum = "SECURITY"
	ListEventsEventTypeError            ListEventsEventTypeEnum = "ERROR"
	ListEventsEventTypeWarning          ListEventsEventTypeEnum = "WARNING"
)

var mappingListEventsEventTypeEnum = map[string]ListEventsEventTypeEnum{
	"KERNEL_OOPS":       ListEventsEventTypeKernelOops,
	"KERNEL_CRASH":      ListEventsEventTypeKernelCrash,
	"CRASH":             ListEventsEventTypeCrash,
	"EXPLOIT_ATTEMPT":   ListEventsEventTypeExploitAttempt,
	"COMPLIANCE":        ListEventsEventTypeCompliance,
	"TUNING_SUGGESTION": ListEventsEventTypeTuningSuggestion,
	"TUNING_APPLIED":    ListEventsEventTypeTuningApplied,
	"SECURITY":          ListEventsEventTypeSecurity,
	"ERROR":             ListEventsEventTypeError,
	"WARNING":           ListEventsEventTypeWarning,
}

var mappingListEventsEventTypeEnumLowerCase = map[string]ListEventsEventTypeEnum{
	"kernel_oops":       ListEventsEventTypeKernelOops,
	"kernel_crash":      ListEventsEventTypeKernelCrash,
	"crash":             ListEventsEventTypeCrash,
	"exploit_attempt":   ListEventsEventTypeExploitAttempt,
	"compliance":        ListEventsEventTypeCompliance,
	"tuning_suggestion": ListEventsEventTypeTuningSuggestion,
	"tuning_applied":    ListEventsEventTypeTuningApplied,
	"security":          ListEventsEventTypeSecurity,
	"error":             ListEventsEventTypeError,
	"warning":           ListEventsEventTypeWarning,
}

// GetListEventsEventTypeEnumValues Enumerates the set of values for ListEventsEventTypeEnum
func GetListEventsEventTypeEnumValues() []ListEventsEventTypeEnum {
	values := make([]ListEventsEventTypeEnum, 0)
	for _, v := range mappingListEventsEventTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListEventsEventTypeEnumStringValues Enumerates the set of values in String for ListEventsEventTypeEnum
func GetListEventsEventTypeEnumStringValues() []string {
	return []string{
		"KERNEL_OOPS",
		"KERNEL_CRASH",
		"CRASH",
		"EXPLOIT_ATTEMPT",
		"COMPLIANCE",
		"TUNING_SUGGESTION",
		"TUNING_APPLIED",
		"SECURITY",
		"ERROR",
		"WARNING",
	}
}

// GetMappingListEventsEventTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEventsEventTypeEnum(val string) (ListEventsEventTypeEnum, bool) {
	enum, ok := mappingListEventsEventTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
