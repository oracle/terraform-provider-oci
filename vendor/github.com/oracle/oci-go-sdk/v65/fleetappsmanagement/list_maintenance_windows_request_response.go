// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMaintenanceWindowsRequest wrapper for the ListMaintenanceWindows operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListMaintenanceWindows.go.html to see an example of how to use ListMaintenanceWindowsRequest.
type ListMaintenanceWindowsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources whose lifecycleState matches the given lifecycleState.
	LifecycleState MaintenanceWindowLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose timeScheduleStart is greater than or equal to the provided date and time.
	TimeScheduleStartGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScheduleStartGreaterThanOrEqualTo"`

	// A filter to return only the Maintenance Windows whose identifier matches the given identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListMaintenanceWindowsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListMaintenanceWindowsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaintenanceWindowsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaintenanceWindowsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaintenanceWindowsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaintenanceWindowsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaintenanceWindowsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMaintenanceWindowLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMaintenanceWindowLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceWindowsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaintenanceWindowsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaintenanceWindowsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaintenanceWindowsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaintenanceWindowsResponse wrapper for the ListMaintenanceWindows operation
type ListMaintenanceWindowsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaintenanceWindowCollection instances
	MaintenanceWindowCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A  number representing the the total number of results available.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListMaintenanceWindowsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaintenanceWindowsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaintenanceWindowsSortOrderEnum Enum with underlying type: string
type ListMaintenanceWindowsSortOrderEnum string

// Set of constants representing the allowable values for ListMaintenanceWindowsSortOrderEnum
const (
	ListMaintenanceWindowsSortOrderAsc  ListMaintenanceWindowsSortOrderEnum = "ASC"
	ListMaintenanceWindowsSortOrderDesc ListMaintenanceWindowsSortOrderEnum = "DESC"
)

var mappingListMaintenanceWindowsSortOrderEnum = map[string]ListMaintenanceWindowsSortOrderEnum{
	"ASC":  ListMaintenanceWindowsSortOrderAsc,
	"DESC": ListMaintenanceWindowsSortOrderDesc,
}

var mappingListMaintenanceWindowsSortOrderEnumLowerCase = map[string]ListMaintenanceWindowsSortOrderEnum{
	"asc":  ListMaintenanceWindowsSortOrderAsc,
	"desc": ListMaintenanceWindowsSortOrderDesc,
}

// GetListMaintenanceWindowsSortOrderEnumValues Enumerates the set of values for ListMaintenanceWindowsSortOrderEnum
func GetListMaintenanceWindowsSortOrderEnumValues() []ListMaintenanceWindowsSortOrderEnum {
	values := make([]ListMaintenanceWindowsSortOrderEnum, 0)
	for _, v := range mappingListMaintenanceWindowsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceWindowsSortOrderEnumStringValues Enumerates the set of values in String for ListMaintenanceWindowsSortOrderEnum
func GetListMaintenanceWindowsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaintenanceWindowsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceWindowsSortOrderEnum(val string) (ListMaintenanceWindowsSortOrderEnum, bool) {
	enum, ok := mappingListMaintenanceWindowsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaintenanceWindowsSortByEnum Enum with underlying type: string
type ListMaintenanceWindowsSortByEnum string

// Set of constants representing the allowable values for ListMaintenanceWindowsSortByEnum
const (
	ListMaintenanceWindowsSortByTimecreated ListMaintenanceWindowsSortByEnum = "timeCreated"
	ListMaintenanceWindowsSortByDisplayname ListMaintenanceWindowsSortByEnum = "displayName"
)

var mappingListMaintenanceWindowsSortByEnum = map[string]ListMaintenanceWindowsSortByEnum{
	"timeCreated": ListMaintenanceWindowsSortByTimecreated,
	"displayName": ListMaintenanceWindowsSortByDisplayname,
}

var mappingListMaintenanceWindowsSortByEnumLowerCase = map[string]ListMaintenanceWindowsSortByEnum{
	"timecreated": ListMaintenanceWindowsSortByTimecreated,
	"displayname": ListMaintenanceWindowsSortByDisplayname,
}

// GetListMaintenanceWindowsSortByEnumValues Enumerates the set of values for ListMaintenanceWindowsSortByEnum
func GetListMaintenanceWindowsSortByEnumValues() []ListMaintenanceWindowsSortByEnum {
	values := make([]ListMaintenanceWindowsSortByEnum, 0)
	for _, v := range mappingListMaintenanceWindowsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaintenanceWindowsSortByEnumStringValues Enumerates the set of values in String for ListMaintenanceWindowsSortByEnum
func GetListMaintenanceWindowsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMaintenanceWindowsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaintenanceWindowsSortByEnum(val string) (ListMaintenanceWindowsSortByEnum, bool) {
	enum, ok := mappingListMaintenanceWindowsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
