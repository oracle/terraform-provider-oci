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

// ListPingMonitorsRequest wrapper for the ListPingMonitors operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListPingMonitors.go.html to see an example of how to use ListPingMonitorsRequest.
type ListPingMonitorsRequest struct {

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
	SortBy ListPingMonitorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Controls the sort order of results.
	SortOrder ListPingMonitorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filters results that exactly match the `displayName` field.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filters results that match the `homeRegion`.
	HomeRegion *string `mandatory:"false" contributesTo:"query" name:"homeRegion"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPingMonitorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPingMonitorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPingMonitorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPingMonitorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPingMonitorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPingMonitorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPingMonitorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPingMonitorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPingMonitorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPingMonitorsResponse wrapper for the ListPingMonitors operation
type ListPingMonitorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PingMonitorSummary instances
	Items []PingMonitorSummary `presentIn:"body"`

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

func (response ListPingMonitorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPingMonitorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPingMonitorsSortByEnum Enum with underlying type: string
type ListPingMonitorsSortByEnum string

// Set of constants representing the allowable values for ListPingMonitorsSortByEnum
const (
	ListPingMonitorsSortById          ListPingMonitorsSortByEnum = "id"
	ListPingMonitorsSortByDisplayname ListPingMonitorsSortByEnum = "displayName"
	ListPingMonitorsSortByTimecreated ListPingMonitorsSortByEnum = "timeCreated"
)

var mappingListPingMonitorsSortByEnum = map[string]ListPingMonitorsSortByEnum{
	"id":          ListPingMonitorsSortById,
	"displayName": ListPingMonitorsSortByDisplayname,
	"timeCreated": ListPingMonitorsSortByTimecreated,
}

// GetListPingMonitorsSortByEnumValues Enumerates the set of values for ListPingMonitorsSortByEnum
func GetListPingMonitorsSortByEnumValues() []ListPingMonitorsSortByEnum {
	values := make([]ListPingMonitorsSortByEnum, 0)
	for _, v := range mappingListPingMonitorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPingMonitorsSortByEnumStringValues Enumerates the set of values in String for ListPingMonitorsSortByEnum
func GetListPingMonitorsSortByEnumStringValues() []string {
	return []string{
		"id",
		"displayName",
		"timeCreated",
	}
}

// GetMappingListPingMonitorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPingMonitorsSortByEnum(val string) (ListPingMonitorsSortByEnum, bool) {
	mappingListPingMonitorsSortByEnumIgnoreCase := make(map[string]ListPingMonitorsSortByEnum)
	for k, v := range mappingListPingMonitorsSortByEnum {
		mappingListPingMonitorsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPingMonitorsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListPingMonitorsSortOrderEnum Enum with underlying type: string
type ListPingMonitorsSortOrderEnum string

// Set of constants representing the allowable values for ListPingMonitorsSortOrderEnum
const (
	ListPingMonitorsSortOrderAsc  ListPingMonitorsSortOrderEnum = "ASC"
	ListPingMonitorsSortOrderDesc ListPingMonitorsSortOrderEnum = "DESC"
)

var mappingListPingMonitorsSortOrderEnum = map[string]ListPingMonitorsSortOrderEnum{
	"ASC":  ListPingMonitorsSortOrderAsc,
	"DESC": ListPingMonitorsSortOrderDesc,
}

// GetListPingMonitorsSortOrderEnumValues Enumerates the set of values for ListPingMonitorsSortOrderEnum
func GetListPingMonitorsSortOrderEnumValues() []ListPingMonitorsSortOrderEnum {
	values := make([]ListPingMonitorsSortOrderEnum, 0)
	for _, v := range mappingListPingMonitorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPingMonitorsSortOrderEnumStringValues Enumerates the set of values in String for ListPingMonitorsSortOrderEnum
func GetListPingMonitorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPingMonitorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPingMonitorsSortOrderEnum(val string) (ListPingMonitorsSortOrderEnum, bool) {
	mappingListPingMonitorsSortOrderEnumIgnoreCase := make(map[string]ListPingMonitorsSortOrderEnum)
	for k, v := range mappingListPingMonitorsSortOrderEnum {
		mappingListPingMonitorsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPingMonitorsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
