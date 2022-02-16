// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package healthchecks

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListHttpMonitorsRequest wrapper for the ListHttpMonitors operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListHttpMonitors.go.html to see an example of how to use ListHttpMonitorsRequest.
type ListHttpMonitorsRequest struct {

	// Filters results by compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by when listing monitors.
	SortBy ListHttpMonitorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Controls the sort order of results.
	SortOrder ListHttpMonitorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filters results that exactly match the `displayName` field.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filters results that match the `homeRegion`.
	HomeRegion *string `mandatory:"false" contributesTo:"query" name:"homeRegion"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHttpMonitorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHttpMonitorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHttpMonitorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHttpMonitorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHttpMonitorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListHttpMonitorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHttpMonitorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHttpMonitorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHttpMonitorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHttpMonitorsResponse wrapper for the ListHttpMonitors operation
type ListHttpMonitorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []HttpMonitorSummary instances
	Items []HttpMonitorSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHttpMonitorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHttpMonitorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHttpMonitorsSortByEnum Enum with underlying type: string
type ListHttpMonitorsSortByEnum string

// Set of constants representing the allowable values for ListHttpMonitorsSortByEnum
const (
	ListHttpMonitorsSortById          ListHttpMonitorsSortByEnum = "id"
	ListHttpMonitorsSortByDisplayname ListHttpMonitorsSortByEnum = "displayName"
	ListHttpMonitorsSortByTimecreated ListHttpMonitorsSortByEnum = "timeCreated"
)

var mappingListHttpMonitorsSortByEnum = map[string]ListHttpMonitorsSortByEnum{
	"id":          ListHttpMonitorsSortById,
	"displayName": ListHttpMonitorsSortByDisplayname,
	"timeCreated": ListHttpMonitorsSortByTimecreated,
}

// GetListHttpMonitorsSortByEnumValues Enumerates the set of values for ListHttpMonitorsSortByEnum
func GetListHttpMonitorsSortByEnumValues() []ListHttpMonitorsSortByEnum {
	values := make([]ListHttpMonitorsSortByEnum, 0)
	for _, v := range mappingListHttpMonitorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHttpMonitorsSortByEnumStringValues Enumerates the set of values in String for ListHttpMonitorsSortByEnum
func GetListHttpMonitorsSortByEnumStringValues() []string {
	return []string{
		"id",
		"displayName",
		"timeCreated",
	}
}

// GetMappingListHttpMonitorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHttpMonitorsSortByEnum(val string) (ListHttpMonitorsSortByEnum, bool) {
	mappingListHttpMonitorsSortByEnumIgnoreCase := make(map[string]ListHttpMonitorsSortByEnum)
	for k, v := range mappingListHttpMonitorsSortByEnum {
		mappingListHttpMonitorsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListHttpMonitorsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListHttpMonitorsSortOrderEnum Enum with underlying type: string
type ListHttpMonitorsSortOrderEnum string

// Set of constants representing the allowable values for ListHttpMonitorsSortOrderEnum
const (
	ListHttpMonitorsSortOrderAsc  ListHttpMonitorsSortOrderEnum = "ASC"
	ListHttpMonitorsSortOrderDesc ListHttpMonitorsSortOrderEnum = "DESC"
)

var mappingListHttpMonitorsSortOrderEnum = map[string]ListHttpMonitorsSortOrderEnum{
	"ASC":  ListHttpMonitorsSortOrderAsc,
	"DESC": ListHttpMonitorsSortOrderDesc,
}

// GetListHttpMonitorsSortOrderEnumValues Enumerates the set of values for ListHttpMonitorsSortOrderEnum
func GetListHttpMonitorsSortOrderEnumValues() []ListHttpMonitorsSortOrderEnum {
	values := make([]ListHttpMonitorsSortOrderEnum, 0)
	for _, v := range mappingListHttpMonitorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHttpMonitorsSortOrderEnumStringValues Enumerates the set of values in String for ListHttpMonitorsSortOrderEnum
func GetListHttpMonitorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHttpMonitorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHttpMonitorsSortOrderEnum(val string) (ListHttpMonitorsSortOrderEnum, bool) {
	mappingListHttpMonitorsSortOrderEnumIgnoreCase := make(map[string]ListHttpMonitorsSortOrderEnum)
	for k, v := range mappingListHttpMonitorsSortOrderEnum {
		mappingListHttpMonitorsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListHttpMonitorsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
