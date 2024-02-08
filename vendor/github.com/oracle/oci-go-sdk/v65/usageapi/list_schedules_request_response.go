// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSchedulesRequest wrapper for the ListSchedules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListSchedules.go.html to see an example of how to use ListSchedulesRequest.
type ListSchedulesRequest struct {

	// The compartment ID in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page at which to start retrieving results.
	// This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximumimum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. If not specified, the default is timeCreated.
	SortBy ListSchedulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListSchedulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Query parameter for filtering by name
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSchedulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSchedulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSchedulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSchedulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSchedulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSchedulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSchedulesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSchedulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSchedulesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSchedulesResponse wrapper for the ListSchedules operation
type ListSchedulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ScheduleCollection instances
	ScheduleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSchedulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSchedulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSchedulesSortByEnum Enum with underlying type: string
type ListSchedulesSortByEnum string

// Set of constants representing the allowable values for ListSchedulesSortByEnum
const (
	ListSchedulesSortByName        ListSchedulesSortByEnum = "name"
	ListSchedulesSortByTimecreated ListSchedulesSortByEnum = "timeCreated"
)

var mappingListSchedulesSortByEnum = map[string]ListSchedulesSortByEnum{
	"name":        ListSchedulesSortByName,
	"timeCreated": ListSchedulesSortByTimecreated,
}

var mappingListSchedulesSortByEnumLowerCase = map[string]ListSchedulesSortByEnum{
	"name":        ListSchedulesSortByName,
	"timecreated": ListSchedulesSortByTimecreated,
}

// GetListSchedulesSortByEnumValues Enumerates the set of values for ListSchedulesSortByEnum
func GetListSchedulesSortByEnumValues() []ListSchedulesSortByEnum {
	values := make([]ListSchedulesSortByEnum, 0)
	for _, v := range mappingListSchedulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulesSortByEnumStringValues Enumerates the set of values in String for ListSchedulesSortByEnum
func GetListSchedulesSortByEnumStringValues() []string {
	return []string{
		"name",
		"timeCreated",
	}
}

// GetMappingListSchedulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulesSortByEnum(val string) (ListSchedulesSortByEnum, bool) {
	enum, ok := mappingListSchedulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSchedulesSortOrderEnum Enum with underlying type: string
type ListSchedulesSortOrderEnum string

// Set of constants representing the allowable values for ListSchedulesSortOrderEnum
const (
	ListSchedulesSortOrderAsc  ListSchedulesSortOrderEnum = "ASC"
	ListSchedulesSortOrderDesc ListSchedulesSortOrderEnum = "DESC"
)

var mappingListSchedulesSortOrderEnum = map[string]ListSchedulesSortOrderEnum{
	"ASC":  ListSchedulesSortOrderAsc,
	"DESC": ListSchedulesSortOrderDesc,
}

var mappingListSchedulesSortOrderEnumLowerCase = map[string]ListSchedulesSortOrderEnum{
	"asc":  ListSchedulesSortOrderAsc,
	"desc": ListSchedulesSortOrderDesc,
}

// GetListSchedulesSortOrderEnumValues Enumerates the set of values for ListSchedulesSortOrderEnum
func GetListSchedulesSortOrderEnumValues() []ListSchedulesSortOrderEnum {
	values := make([]ListSchedulesSortOrderEnum, 0)
	for _, v := range mappingListSchedulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchedulesSortOrderEnumStringValues Enumerates the set of values in String for ListSchedulesSortOrderEnum
func GetListSchedulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSchedulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchedulesSortOrderEnum(val string) (ListSchedulesSortOrderEnum, bool) {
	enum, ok := mappingListSchedulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
