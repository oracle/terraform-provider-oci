// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMonitorsRequest wrapper for the ListMonitors operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmsynthetics/ListMonitors.go.html to see an example of how to use ListMonitorsRequest.
type ListMonitorsRequest struct {

	// The APM domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// A filter to return only the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only monitors using scriptId.
	ScriptId *string `mandatory:"false" contributesTo:"query" name:"scriptId"`

	// The name of the public or dedicated vantage point.
	VantagePoint *string `mandatory:"false" contributesTo:"query" name:"vantagePoint"`

	// A filter to return only monitors that match the given monitor type.
	// Supported values are SCRIPTED_BROWSER, BROWSER, SCRIPTED_REST, REST, NETWORK, DNS, FTP and SQL.
	MonitorType *string `mandatory:"false" contributesTo:"query" name:"monitorType"`

	// A filter to return only monitors that match the status given.
	Status ListMonitorsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The maximum number of results per page, or items to return in a paginated
	// "List" call. For information on how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return the monitors whose maintenance window is currently active.
	IsMaintenanceWindowActive *bool `mandatory:"false" contributesTo:"query" name:"isMaintenanceWindowActive"`

	// A filter to return the monitors whose maintenance window is set.
	IsMaintenanceWindowSet *bool `mandatory:"false" contributesTo:"query" name:"isMaintenanceWindowSet"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). Default sort order is ascending.
	SortOrder ListMonitorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order of displayName is ascending.
	// Default order of timeCreated and timeUpdated is descending.
	// The displayName sort by is case insensitive.
	SortBy ListMonitorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMonitorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMonitorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMonitorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMonitorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMonitorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMonitorsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListMonitorsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMonitorsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMonitorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMonitorsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMonitorsResponse wrapper for the ListMonitors operation
type ListMonitorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitorCollection instances
	MonitorCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMonitorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMonitorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMonitorsStatusEnum Enum with underlying type: string
type ListMonitorsStatusEnum string

// Set of constants representing the allowable values for ListMonitorsStatusEnum
const (
	ListMonitorsStatusEnabled  ListMonitorsStatusEnum = "ENABLED"
	ListMonitorsStatusDisabled ListMonitorsStatusEnum = "DISABLED"
	ListMonitorsStatusInvalid  ListMonitorsStatusEnum = "INVALID"
)

var mappingListMonitorsStatusEnum = map[string]ListMonitorsStatusEnum{
	"ENABLED":  ListMonitorsStatusEnabled,
	"DISABLED": ListMonitorsStatusDisabled,
	"INVALID":  ListMonitorsStatusInvalid,
}

var mappingListMonitorsStatusEnumLowerCase = map[string]ListMonitorsStatusEnum{
	"enabled":  ListMonitorsStatusEnabled,
	"disabled": ListMonitorsStatusDisabled,
	"invalid":  ListMonitorsStatusInvalid,
}

// GetListMonitorsStatusEnumValues Enumerates the set of values for ListMonitorsStatusEnum
func GetListMonitorsStatusEnumValues() []ListMonitorsStatusEnum {
	values := make([]ListMonitorsStatusEnum, 0)
	for _, v := range mappingListMonitorsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitorsStatusEnumStringValues Enumerates the set of values in String for ListMonitorsStatusEnum
func GetListMonitorsStatusEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"INVALID",
	}
}

// GetMappingListMonitorsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitorsStatusEnum(val string) (ListMonitorsStatusEnum, bool) {
	enum, ok := mappingListMonitorsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitorsSortOrderEnum Enum with underlying type: string
type ListMonitorsSortOrderEnum string

// Set of constants representing the allowable values for ListMonitorsSortOrderEnum
const (
	ListMonitorsSortOrderAsc  ListMonitorsSortOrderEnum = "ASC"
	ListMonitorsSortOrderDesc ListMonitorsSortOrderEnum = "DESC"
)

var mappingListMonitorsSortOrderEnum = map[string]ListMonitorsSortOrderEnum{
	"ASC":  ListMonitorsSortOrderAsc,
	"DESC": ListMonitorsSortOrderDesc,
}

var mappingListMonitorsSortOrderEnumLowerCase = map[string]ListMonitorsSortOrderEnum{
	"asc":  ListMonitorsSortOrderAsc,
	"desc": ListMonitorsSortOrderDesc,
}

// GetListMonitorsSortOrderEnumValues Enumerates the set of values for ListMonitorsSortOrderEnum
func GetListMonitorsSortOrderEnumValues() []ListMonitorsSortOrderEnum {
	values := make([]ListMonitorsSortOrderEnum, 0)
	for _, v := range mappingListMonitorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitorsSortOrderEnumStringValues Enumerates the set of values in String for ListMonitorsSortOrderEnum
func GetListMonitorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMonitorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitorsSortOrderEnum(val string) (ListMonitorsSortOrderEnum, bool) {
	enum, ok := mappingListMonitorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMonitorsSortByEnum Enum with underlying type: string
type ListMonitorsSortByEnum string

// Set of constants representing the allowable values for ListMonitorsSortByEnum
const (
	ListMonitorsSortByDisplayname                  ListMonitorsSortByEnum = "displayName"
	ListMonitorsSortByTimecreated                  ListMonitorsSortByEnum = "timeCreated"
	ListMonitorsSortByTimeupdated                  ListMonitorsSortByEnum = "timeUpdated"
	ListMonitorsSortByStatus                       ListMonitorsSortByEnum = "status"
	ListMonitorsSortByMonitortype                  ListMonitorsSortByEnum = "monitorType"
	ListMonitorsSortByMaintenancewindowtimestarted ListMonitorsSortByEnum = "maintenanceWindowTimeStarted"
)

var mappingListMonitorsSortByEnum = map[string]ListMonitorsSortByEnum{
	"displayName":                  ListMonitorsSortByDisplayname,
	"timeCreated":                  ListMonitorsSortByTimecreated,
	"timeUpdated":                  ListMonitorsSortByTimeupdated,
	"status":                       ListMonitorsSortByStatus,
	"monitorType":                  ListMonitorsSortByMonitortype,
	"maintenanceWindowTimeStarted": ListMonitorsSortByMaintenancewindowtimestarted,
}

var mappingListMonitorsSortByEnumLowerCase = map[string]ListMonitorsSortByEnum{
	"displayname":                  ListMonitorsSortByDisplayname,
	"timecreated":                  ListMonitorsSortByTimecreated,
	"timeupdated":                  ListMonitorsSortByTimeupdated,
	"status":                       ListMonitorsSortByStatus,
	"monitortype":                  ListMonitorsSortByMonitortype,
	"maintenancewindowtimestarted": ListMonitorsSortByMaintenancewindowtimestarted,
}

// GetListMonitorsSortByEnumValues Enumerates the set of values for ListMonitorsSortByEnum
func GetListMonitorsSortByEnumValues() []ListMonitorsSortByEnum {
	values := make([]ListMonitorsSortByEnum, 0)
	for _, v := range mappingListMonitorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMonitorsSortByEnumStringValues Enumerates the set of values in String for ListMonitorsSortByEnum
func GetListMonitorsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
		"timeUpdated",
		"status",
		"monitorType",
		"maintenanceWindowTimeStarted",
	}
}

// GetMappingListMonitorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMonitorsSortByEnum(val string) (ListMonitorsSortByEnum, bool) {
	enum, ok := mappingListMonitorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
