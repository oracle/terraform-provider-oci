// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTaskSchedulesRequest wrapper for the ListTaskSchedules operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListTaskSchedules.go.html to see an example of how to use ListTaskSchedulesRequest.
type ListTaskSchedulesRequest struct {

	// The ID of the Fleet.
	FleetId *string `mandatory:"false" contributesTo:"query" name:"fleetId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) to identify this task schedule.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The Fleet-unique identifier of the related managed instance.
	ManagedInstanceId *string `mandatory:"false" contributesTo:"query" name:"managedInstanceId"`

	// The task name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filter the list with task schedule name contains the given value.
	TaskScheduleNameContains *string `mandatory:"false" contributesTo:"query" name:"taskScheduleNameContains"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListTaskSchedulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used to sort task schedule records. Only one sort order may be provided.
	// Default order for _timeCreated_ is **ascending**.
	// If no value is specified, _timeCreated_ is default.
	SortBy ListTaskSchedulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTaskSchedulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTaskSchedulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTaskSchedulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTaskSchedulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTaskSchedulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTaskSchedulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTaskSchedulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTaskSchedulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTaskSchedulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTaskSchedulesResponse wrapper for the ListTaskSchedules operation
type ListTaskSchedulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TaskScheduleCollection instances
	TaskScheduleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination, when this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTaskSchedulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTaskSchedulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTaskSchedulesSortOrderEnum Enum with underlying type: string
type ListTaskSchedulesSortOrderEnum string

// Set of constants representing the allowable values for ListTaskSchedulesSortOrderEnum
const (
	ListTaskSchedulesSortOrderAsc  ListTaskSchedulesSortOrderEnum = "ASC"
	ListTaskSchedulesSortOrderDesc ListTaskSchedulesSortOrderEnum = "DESC"
)

var mappingListTaskSchedulesSortOrderEnum = map[string]ListTaskSchedulesSortOrderEnum{
	"ASC":  ListTaskSchedulesSortOrderAsc,
	"DESC": ListTaskSchedulesSortOrderDesc,
}

var mappingListTaskSchedulesSortOrderEnumLowerCase = map[string]ListTaskSchedulesSortOrderEnum{
	"asc":  ListTaskSchedulesSortOrderAsc,
	"desc": ListTaskSchedulesSortOrderDesc,
}

// GetListTaskSchedulesSortOrderEnumValues Enumerates the set of values for ListTaskSchedulesSortOrderEnum
func GetListTaskSchedulesSortOrderEnumValues() []ListTaskSchedulesSortOrderEnum {
	values := make([]ListTaskSchedulesSortOrderEnum, 0)
	for _, v := range mappingListTaskSchedulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskSchedulesSortOrderEnumStringValues Enumerates the set of values in String for ListTaskSchedulesSortOrderEnum
func GetListTaskSchedulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTaskSchedulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskSchedulesSortOrderEnum(val string) (ListTaskSchedulesSortOrderEnum, bool) {
	enum, ok := mappingListTaskSchedulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTaskSchedulesSortByEnum Enum with underlying type: string
type ListTaskSchedulesSortByEnum string

// Set of constants representing the allowable values for ListTaskSchedulesSortByEnum
const (
	ListTaskSchedulesSortByName           ListTaskSchedulesSortByEnum = "name"
	ListTaskSchedulesSortByLifecyclestate ListTaskSchedulesSortByEnum = "lifecycleState"
	ListTaskSchedulesSortByTimecreated    ListTaskSchedulesSortByEnum = "timeCreated"
	ListTaskSchedulesSortByTimenextrun    ListTaskSchedulesSortByEnum = "timeNextRun"
	ListTaskSchedulesSortByTimelastrun    ListTaskSchedulesSortByEnum = "timeLastRun"
)

var mappingListTaskSchedulesSortByEnum = map[string]ListTaskSchedulesSortByEnum{
	"name":           ListTaskSchedulesSortByName,
	"lifecycleState": ListTaskSchedulesSortByLifecyclestate,
	"timeCreated":    ListTaskSchedulesSortByTimecreated,
	"timeNextRun":    ListTaskSchedulesSortByTimenextrun,
	"timeLastRun":    ListTaskSchedulesSortByTimelastrun,
}

var mappingListTaskSchedulesSortByEnumLowerCase = map[string]ListTaskSchedulesSortByEnum{
	"name":           ListTaskSchedulesSortByName,
	"lifecyclestate": ListTaskSchedulesSortByLifecyclestate,
	"timecreated":    ListTaskSchedulesSortByTimecreated,
	"timenextrun":    ListTaskSchedulesSortByTimenextrun,
	"timelastrun":    ListTaskSchedulesSortByTimelastrun,
}

// GetListTaskSchedulesSortByEnumValues Enumerates the set of values for ListTaskSchedulesSortByEnum
func GetListTaskSchedulesSortByEnumValues() []ListTaskSchedulesSortByEnum {
	values := make([]ListTaskSchedulesSortByEnum, 0)
	for _, v := range mappingListTaskSchedulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskSchedulesSortByEnumStringValues Enumerates the set of values in String for ListTaskSchedulesSortByEnum
func GetListTaskSchedulesSortByEnumStringValues() []string {
	return []string{
		"name",
		"lifecycleState",
		"timeCreated",
		"timeNextRun",
		"timeLastRun",
	}
}

// GetMappingListTaskSchedulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskSchedulesSortByEnum(val string) (ListTaskSchedulesSortByEnum, bool) {
	enum, ok := mappingListTaskSchedulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
