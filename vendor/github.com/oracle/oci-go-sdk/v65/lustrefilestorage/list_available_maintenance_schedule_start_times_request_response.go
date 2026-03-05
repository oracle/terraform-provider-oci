// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAvailableMaintenanceScheduleStartTimesRequest wrapper for the ListAvailableMaintenanceScheduleStartTimes operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/ListAvailableMaintenanceScheduleStartTimes.go.html to see an example of how to use ListAvailableMaintenanceScheduleStartTimesRequest.
type ListAvailableMaintenanceScheduleStartTimesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// Day of the week filter
	DayOfWeek ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum `mandatory:"false" contributesTo:"query" name:"dayOfWeek" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide only one sort order.
	SortBy ListAvailableMaintenanceScheduleStartTimesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAvailableMaintenanceScheduleStartTimesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableMaintenanceScheduleStartTimesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableMaintenanceScheduleStartTimesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAvailableMaintenanceScheduleStartTimesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableMaintenanceScheduleStartTimesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAvailableMaintenanceScheduleStartTimesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum(string(request.DayOfWeek)); !ok && request.DayOfWeek != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DayOfWeek: %s. Supported values are: %s.", request.DayOfWeek, strings.Join(GetListAvailableMaintenanceScheduleStartTimesDayOfWeekEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableMaintenanceScheduleStartTimesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAvailableMaintenanceScheduleStartTimesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAvailableMaintenanceScheduleStartTimesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAvailableMaintenanceScheduleStartTimesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAvailableMaintenanceScheduleStartTimesResponse wrapper for the ListAvailableMaintenanceScheduleStartTimes operation
type ListAvailableMaintenanceScheduleStartTimesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AvailableMaintenanceScheduleStartTimeCollection instances
	AvailableMaintenanceScheduleStartTimeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableMaintenanceScheduleStartTimesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableMaintenanceScheduleStartTimesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum Enum with underlying type: string
type ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum string

// Set of constants representing the allowable values for ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum
const (
	ListAvailableMaintenanceScheduleStartTimesDayOfWeekMonday    ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum = "MONDAY"
	ListAvailableMaintenanceScheduleStartTimesDayOfWeekTuesday   ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum = "TUESDAY"
	ListAvailableMaintenanceScheduleStartTimesDayOfWeekWednesday ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum = "WEDNESDAY"
	ListAvailableMaintenanceScheduleStartTimesDayOfWeekThursday  ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum = "THURSDAY"
	ListAvailableMaintenanceScheduleStartTimesDayOfWeekFriday    ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum = "FRIDAY"
	ListAvailableMaintenanceScheduleStartTimesDayOfWeekSaturday  ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum = "SATURDAY"
	ListAvailableMaintenanceScheduleStartTimesDayOfWeekSunday    ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum = "SUNDAY"
)

var mappingListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum = map[string]ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum{
	"MONDAY":    ListAvailableMaintenanceScheduleStartTimesDayOfWeekMonday,
	"TUESDAY":   ListAvailableMaintenanceScheduleStartTimesDayOfWeekTuesday,
	"WEDNESDAY": ListAvailableMaintenanceScheduleStartTimesDayOfWeekWednesday,
	"THURSDAY":  ListAvailableMaintenanceScheduleStartTimesDayOfWeekThursday,
	"FRIDAY":    ListAvailableMaintenanceScheduleStartTimesDayOfWeekFriday,
	"SATURDAY":  ListAvailableMaintenanceScheduleStartTimesDayOfWeekSaturday,
	"SUNDAY":    ListAvailableMaintenanceScheduleStartTimesDayOfWeekSunday,
}

var mappingListAvailableMaintenanceScheduleStartTimesDayOfWeekEnumLowerCase = map[string]ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum{
	"monday":    ListAvailableMaintenanceScheduleStartTimesDayOfWeekMonday,
	"tuesday":   ListAvailableMaintenanceScheduleStartTimesDayOfWeekTuesday,
	"wednesday": ListAvailableMaintenanceScheduleStartTimesDayOfWeekWednesday,
	"thursday":  ListAvailableMaintenanceScheduleStartTimesDayOfWeekThursday,
	"friday":    ListAvailableMaintenanceScheduleStartTimesDayOfWeekFriday,
	"saturday":  ListAvailableMaintenanceScheduleStartTimesDayOfWeekSaturday,
	"sunday":    ListAvailableMaintenanceScheduleStartTimesDayOfWeekSunday,
}

// GetListAvailableMaintenanceScheduleStartTimesDayOfWeekEnumValues Enumerates the set of values for ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum
func GetListAvailableMaintenanceScheduleStartTimesDayOfWeekEnumValues() []ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum {
	values := make([]ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum, 0)
	for _, v := range mappingListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableMaintenanceScheduleStartTimesDayOfWeekEnumStringValues Enumerates the set of values in String for ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum
func GetListAvailableMaintenanceScheduleStartTimesDayOfWeekEnumStringValues() []string {
	return []string{
		"MONDAY",
		"TUESDAY",
		"WEDNESDAY",
		"THURSDAY",
		"FRIDAY",
		"SATURDAY",
		"SUNDAY",
	}
}

// GetMappingListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum(val string) (ListAvailableMaintenanceScheduleStartTimesDayOfWeekEnum, bool) {
	enum, ok := mappingListAvailableMaintenanceScheduleStartTimesDayOfWeekEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableMaintenanceScheduleStartTimesSortByEnum Enum with underlying type: string
type ListAvailableMaintenanceScheduleStartTimesSortByEnum string

// Set of constants representing the allowable values for ListAvailableMaintenanceScheduleStartTimesSortByEnum
const (
	ListAvailableMaintenanceScheduleStartTimesSortByDayofweek ListAvailableMaintenanceScheduleStartTimesSortByEnum = "dayOfWeek"
)

var mappingListAvailableMaintenanceScheduleStartTimesSortByEnum = map[string]ListAvailableMaintenanceScheduleStartTimesSortByEnum{
	"dayOfWeek": ListAvailableMaintenanceScheduleStartTimesSortByDayofweek,
}

var mappingListAvailableMaintenanceScheduleStartTimesSortByEnumLowerCase = map[string]ListAvailableMaintenanceScheduleStartTimesSortByEnum{
	"dayofweek": ListAvailableMaintenanceScheduleStartTimesSortByDayofweek,
}

// GetListAvailableMaintenanceScheduleStartTimesSortByEnumValues Enumerates the set of values for ListAvailableMaintenanceScheduleStartTimesSortByEnum
func GetListAvailableMaintenanceScheduleStartTimesSortByEnumValues() []ListAvailableMaintenanceScheduleStartTimesSortByEnum {
	values := make([]ListAvailableMaintenanceScheduleStartTimesSortByEnum, 0)
	for _, v := range mappingListAvailableMaintenanceScheduleStartTimesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableMaintenanceScheduleStartTimesSortByEnumStringValues Enumerates the set of values in String for ListAvailableMaintenanceScheduleStartTimesSortByEnum
func GetListAvailableMaintenanceScheduleStartTimesSortByEnumStringValues() []string {
	return []string{
		"dayOfWeek",
	}
}

// GetMappingListAvailableMaintenanceScheduleStartTimesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableMaintenanceScheduleStartTimesSortByEnum(val string) (ListAvailableMaintenanceScheduleStartTimesSortByEnum, bool) {
	enum, ok := mappingListAvailableMaintenanceScheduleStartTimesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAvailableMaintenanceScheduleStartTimesSortOrderEnum Enum with underlying type: string
type ListAvailableMaintenanceScheduleStartTimesSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableMaintenanceScheduleStartTimesSortOrderEnum
const (
	ListAvailableMaintenanceScheduleStartTimesSortOrderAsc  ListAvailableMaintenanceScheduleStartTimesSortOrderEnum = "ASC"
	ListAvailableMaintenanceScheduleStartTimesSortOrderDesc ListAvailableMaintenanceScheduleStartTimesSortOrderEnum = "DESC"
)

var mappingListAvailableMaintenanceScheduleStartTimesSortOrderEnum = map[string]ListAvailableMaintenanceScheduleStartTimesSortOrderEnum{
	"ASC":  ListAvailableMaintenanceScheduleStartTimesSortOrderAsc,
	"DESC": ListAvailableMaintenanceScheduleStartTimesSortOrderDesc,
}

var mappingListAvailableMaintenanceScheduleStartTimesSortOrderEnumLowerCase = map[string]ListAvailableMaintenanceScheduleStartTimesSortOrderEnum{
	"asc":  ListAvailableMaintenanceScheduleStartTimesSortOrderAsc,
	"desc": ListAvailableMaintenanceScheduleStartTimesSortOrderDesc,
}

// GetListAvailableMaintenanceScheduleStartTimesSortOrderEnumValues Enumerates the set of values for ListAvailableMaintenanceScheduleStartTimesSortOrderEnum
func GetListAvailableMaintenanceScheduleStartTimesSortOrderEnumValues() []ListAvailableMaintenanceScheduleStartTimesSortOrderEnum {
	values := make([]ListAvailableMaintenanceScheduleStartTimesSortOrderEnum, 0)
	for _, v := range mappingListAvailableMaintenanceScheduleStartTimesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAvailableMaintenanceScheduleStartTimesSortOrderEnumStringValues Enumerates the set of values in String for ListAvailableMaintenanceScheduleStartTimesSortOrderEnum
func GetListAvailableMaintenanceScheduleStartTimesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAvailableMaintenanceScheduleStartTimesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAvailableMaintenanceScheduleStartTimesSortOrderEnum(val string) (ListAvailableMaintenanceScheduleStartTimesSortOrderEnum, bool) {
	enum, ok := mappingListAvailableMaintenanceScheduleStartTimesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
