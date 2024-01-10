// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAlertLogsRequest wrapper for the ListAlertLogs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAlertLogs.go.html to see an example of how to use ListAlertLogsRequest.
type ListAlertLogsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The optional greater than or equal to timestamp to filter the logs.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to timestamp to filter the logs.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The optional parameter to filter the alert logs by log level.
	LevelFilter ListAlertLogsLevelFilterEnum `mandatory:"false" contributesTo:"query" name:"levelFilter" omitEmpty:"true"`

	// The optional parameter to filter the attention or alert logs by type.
	TypeFilter ListAlertLogsTypeFilterEnum `mandatory:"false" contributesTo:"query" name:"typeFilter" omitEmpty:"true"`

	// The optional query parameter to filter the attention or alert logs by search text.
	LogSearchText *string `mandatory:"false" contributesTo:"query" name:"logSearchText"`

	// The flag to indicate whether the search text is regular expression or not.
	IsRegularExpression *bool `mandatory:"false" contributesTo:"query" name:"isRegularExpression"`

	// The possible sortBy values of attention logs.
	SortBy ListAlertLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListAlertLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAlertLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAlertLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAlertLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAlertLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAlertLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAlertLogsLevelFilterEnum(string(request.LevelFilter)); !ok && request.LevelFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LevelFilter: %s. Supported values are: %s.", request.LevelFilter, strings.Join(GetListAlertLogsLevelFilterEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertLogsTypeFilterEnum(string(request.TypeFilter)); !ok && request.TypeFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TypeFilter: %s. Supported values are: %s.", request.TypeFilter, strings.Join(GetListAlertLogsTypeFilterEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAlertLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAlertLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAlertLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAlertLogsResponse wrapper for the ListAlertLogs operation
type ListAlertLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AlertLogCollection instances
	AlertLogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAlertLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAlertLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAlertLogsLevelFilterEnum Enum with underlying type: string
type ListAlertLogsLevelFilterEnum string

// Set of constants representing the allowable values for ListAlertLogsLevelFilterEnum
const (
	ListAlertLogsLevelFilterCritical  ListAlertLogsLevelFilterEnum = "CRITICAL"
	ListAlertLogsLevelFilterSevere    ListAlertLogsLevelFilterEnum = "SEVERE"
	ListAlertLogsLevelFilterImportant ListAlertLogsLevelFilterEnum = "IMPORTANT"
	ListAlertLogsLevelFilterNormal    ListAlertLogsLevelFilterEnum = "NORMAL"
	ListAlertLogsLevelFilterAll       ListAlertLogsLevelFilterEnum = "ALL"
)

var mappingListAlertLogsLevelFilterEnum = map[string]ListAlertLogsLevelFilterEnum{
	"CRITICAL":  ListAlertLogsLevelFilterCritical,
	"SEVERE":    ListAlertLogsLevelFilterSevere,
	"IMPORTANT": ListAlertLogsLevelFilterImportant,
	"NORMAL":    ListAlertLogsLevelFilterNormal,
	"ALL":       ListAlertLogsLevelFilterAll,
}

var mappingListAlertLogsLevelFilterEnumLowerCase = map[string]ListAlertLogsLevelFilterEnum{
	"critical":  ListAlertLogsLevelFilterCritical,
	"severe":    ListAlertLogsLevelFilterSevere,
	"important": ListAlertLogsLevelFilterImportant,
	"normal":    ListAlertLogsLevelFilterNormal,
	"all":       ListAlertLogsLevelFilterAll,
}

// GetListAlertLogsLevelFilterEnumValues Enumerates the set of values for ListAlertLogsLevelFilterEnum
func GetListAlertLogsLevelFilterEnumValues() []ListAlertLogsLevelFilterEnum {
	values := make([]ListAlertLogsLevelFilterEnum, 0)
	for _, v := range mappingListAlertLogsLevelFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertLogsLevelFilterEnumStringValues Enumerates the set of values in String for ListAlertLogsLevelFilterEnum
func GetListAlertLogsLevelFilterEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"SEVERE",
		"IMPORTANT",
		"NORMAL",
		"ALL",
	}
}

// GetMappingListAlertLogsLevelFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertLogsLevelFilterEnum(val string) (ListAlertLogsLevelFilterEnum, bool) {
	enum, ok := mappingListAlertLogsLevelFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertLogsTypeFilterEnum Enum with underlying type: string
type ListAlertLogsTypeFilterEnum string

// Set of constants representing the allowable values for ListAlertLogsTypeFilterEnum
const (
	ListAlertLogsTypeFilterUnknown       ListAlertLogsTypeFilterEnum = "UNKNOWN"
	ListAlertLogsTypeFilterIncidentError ListAlertLogsTypeFilterEnum = "INCIDENT_ERROR"
	ListAlertLogsTypeFilterError         ListAlertLogsTypeFilterEnum = "ERROR"
	ListAlertLogsTypeFilterWarning       ListAlertLogsTypeFilterEnum = "WARNING"
	ListAlertLogsTypeFilterNotification  ListAlertLogsTypeFilterEnum = "NOTIFICATION"
	ListAlertLogsTypeFilterTrace         ListAlertLogsTypeFilterEnum = "TRACE"
	ListAlertLogsTypeFilterAll           ListAlertLogsTypeFilterEnum = "ALL"
)

var mappingListAlertLogsTypeFilterEnum = map[string]ListAlertLogsTypeFilterEnum{
	"UNKNOWN":        ListAlertLogsTypeFilterUnknown,
	"INCIDENT_ERROR": ListAlertLogsTypeFilterIncidentError,
	"ERROR":          ListAlertLogsTypeFilterError,
	"WARNING":        ListAlertLogsTypeFilterWarning,
	"NOTIFICATION":   ListAlertLogsTypeFilterNotification,
	"TRACE":          ListAlertLogsTypeFilterTrace,
	"ALL":            ListAlertLogsTypeFilterAll,
}

var mappingListAlertLogsTypeFilterEnumLowerCase = map[string]ListAlertLogsTypeFilterEnum{
	"unknown":        ListAlertLogsTypeFilterUnknown,
	"incident_error": ListAlertLogsTypeFilterIncidentError,
	"error":          ListAlertLogsTypeFilterError,
	"warning":        ListAlertLogsTypeFilterWarning,
	"notification":   ListAlertLogsTypeFilterNotification,
	"trace":          ListAlertLogsTypeFilterTrace,
	"all":            ListAlertLogsTypeFilterAll,
}

// GetListAlertLogsTypeFilterEnumValues Enumerates the set of values for ListAlertLogsTypeFilterEnum
func GetListAlertLogsTypeFilterEnumValues() []ListAlertLogsTypeFilterEnum {
	values := make([]ListAlertLogsTypeFilterEnum, 0)
	for _, v := range mappingListAlertLogsTypeFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertLogsTypeFilterEnumStringValues Enumerates the set of values in String for ListAlertLogsTypeFilterEnum
func GetListAlertLogsTypeFilterEnumStringValues() []string {
	return []string{
		"UNKNOWN",
		"INCIDENT_ERROR",
		"ERROR",
		"WARNING",
		"NOTIFICATION",
		"TRACE",
		"ALL",
	}
}

// GetMappingListAlertLogsTypeFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertLogsTypeFilterEnum(val string) (ListAlertLogsTypeFilterEnum, bool) {
	enum, ok := mappingListAlertLogsTypeFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertLogsSortByEnum Enum with underlying type: string
type ListAlertLogsSortByEnum string

// Set of constants representing the allowable values for ListAlertLogsSortByEnum
const (
	ListAlertLogsSortByLevel     ListAlertLogsSortByEnum = "LEVEL"
	ListAlertLogsSortByType      ListAlertLogsSortByEnum = "TYPE"
	ListAlertLogsSortByMessage   ListAlertLogsSortByEnum = "MESSAGE"
	ListAlertLogsSortByTimestamp ListAlertLogsSortByEnum = "TIMESTAMP"
)

var mappingListAlertLogsSortByEnum = map[string]ListAlertLogsSortByEnum{
	"LEVEL":     ListAlertLogsSortByLevel,
	"TYPE":      ListAlertLogsSortByType,
	"MESSAGE":   ListAlertLogsSortByMessage,
	"TIMESTAMP": ListAlertLogsSortByTimestamp,
}

var mappingListAlertLogsSortByEnumLowerCase = map[string]ListAlertLogsSortByEnum{
	"level":     ListAlertLogsSortByLevel,
	"type":      ListAlertLogsSortByType,
	"message":   ListAlertLogsSortByMessage,
	"timestamp": ListAlertLogsSortByTimestamp,
}

// GetListAlertLogsSortByEnumValues Enumerates the set of values for ListAlertLogsSortByEnum
func GetListAlertLogsSortByEnumValues() []ListAlertLogsSortByEnum {
	values := make([]ListAlertLogsSortByEnum, 0)
	for _, v := range mappingListAlertLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertLogsSortByEnumStringValues Enumerates the set of values in String for ListAlertLogsSortByEnum
func GetListAlertLogsSortByEnumStringValues() []string {
	return []string{
		"LEVEL",
		"TYPE",
		"MESSAGE",
		"TIMESTAMP",
	}
}

// GetMappingListAlertLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertLogsSortByEnum(val string) (ListAlertLogsSortByEnum, bool) {
	enum, ok := mappingListAlertLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAlertLogsSortOrderEnum Enum with underlying type: string
type ListAlertLogsSortOrderEnum string

// Set of constants representing the allowable values for ListAlertLogsSortOrderEnum
const (
	ListAlertLogsSortOrderAsc  ListAlertLogsSortOrderEnum = "ASC"
	ListAlertLogsSortOrderDesc ListAlertLogsSortOrderEnum = "DESC"
)

var mappingListAlertLogsSortOrderEnum = map[string]ListAlertLogsSortOrderEnum{
	"ASC":  ListAlertLogsSortOrderAsc,
	"DESC": ListAlertLogsSortOrderDesc,
}

var mappingListAlertLogsSortOrderEnumLowerCase = map[string]ListAlertLogsSortOrderEnum{
	"asc":  ListAlertLogsSortOrderAsc,
	"desc": ListAlertLogsSortOrderDesc,
}

// GetListAlertLogsSortOrderEnumValues Enumerates the set of values for ListAlertLogsSortOrderEnum
func GetListAlertLogsSortOrderEnumValues() []ListAlertLogsSortOrderEnum {
	values := make([]ListAlertLogsSortOrderEnum, 0)
	for _, v := range mappingListAlertLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAlertLogsSortOrderEnumStringValues Enumerates the set of values in String for ListAlertLogsSortOrderEnum
func GetListAlertLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAlertLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAlertLogsSortOrderEnum(val string) (ListAlertLogsSortOrderEnum, bool) {
	enum, ok := mappingListAlertLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
