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

// ListAttentionLogsRequest wrapper for the ListAttentionLogs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAttentionLogs.go.html to see an example of how to use ListAttentionLogsRequest.
type ListAttentionLogsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The optional greater than or equal to timestamp to filter the logs.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to timestamp to filter the logs.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The optional parameter to filter the attention logs by urgency.
	UrgencyFilter ListAttentionLogsUrgencyFilterEnum `mandatory:"false" contributesTo:"query" name:"urgencyFilter" omitEmpty:"true"`

	// The optional parameter to filter the attention or alert logs by type.
	TypeFilter ListAttentionLogsTypeFilterEnum `mandatory:"false" contributesTo:"query" name:"typeFilter" omitEmpty:"true"`

	// The optional query parameter to filter the attention or alert logs by search text.
	LogSearchText *string `mandatory:"false" contributesTo:"query" name:"logSearchText"`

	// The flag to indicate whether the search text is regular expression or not.
	IsRegularExpression *bool `mandatory:"false" contributesTo:"query" name:"isRegularExpression"`

	// The possible sortBy values of attention logs.
	SortBy ListAttentionLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListAttentionLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListAttentionLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAttentionLogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAttentionLogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAttentionLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAttentionLogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAttentionLogsUrgencyFilterEnum(string(request.UrgencyFilter)); !ok && request.UrgencyFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UrgencyFilter: %s. Supported values are: %s.", request.UrgencyFilter, strings.Join(GetListAttentionLogsUrgencyFilterEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttentionLogsTypeFilterEnum(string(request.TypeFilter)); !ok && request.TypeFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TypeFilter: %s. Supported values are: %s.", request.TypeFilter, strings.Join(GetListAttentionLogsTypeFilterEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttentionLogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAttentionLogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAttentionLogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAttentionLogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAttentionLogsResponse wrapper for the ListAttentionLogs operation
type ListAttentionLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AttentionLogCollection instances
	AttentionLogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAttentionLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAttentionLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAttentionLogsUrgencyFilterEnum Enum with underlying type: string
type ListAttentionLogsUrgencyFilterEnum string

// Set of constants representing the allowable values for ListAttentionLogsUrgencyFilterEnum
const (
	ListAttentionLogsUrgencyFilterImmediate  ListAttentionLogsUrgencyFilterEnum = "IMMEDIATE"
	ListAttentionLogsUrgencyFilterSoon       ListAttentionLogsUrgencyFilterEnum = "SOON"
	ListAttentionLogsUrgencyFilterDeferrable ListAttentionLogsUrgencyFilterEnum = "DEFERRABLE"
	ListAttentionLogsUrgencyFilterInfo       ListAttentionLogsUrgencyFilterEnum = "INFO"
	ListAttentionLogsUrgencyFilterAll        ListAttentionLogsUrgencyFilterEnum = "ALL"
)

var mappingListAttentionLogsUrgencyFilterEnum = map[string]ListAttentionLogsUrgencyFilterEnum{
	"IMMEDIATE":  ListAttentionLogsUrgencyFilterImmediate,
	"SOON":       ListAttentionLogsUrgencyFilterSoon,
	"DEFERRABLE": ListAttentionLogsUrgencyFilterDeferrable,
	"INFO":       ListAttentionLogsUrgencyFilterInfo,
	"ALL":        ListAttentionLogsUrgencyFilterAll,
}

var mappingListAttentionLogsUrgencyFilterEnumLowerCase = map[string]ListAttentionLogsUrgencyFilterEnum{
	"immediate":  ListAttentionLogsUrgencyFilterImmediate,
	"soon":       ListAttentionLogsUrgencyFilterSoon,
	"deferrable": ListAttentionLogsUrgencyFilterDeferrable,
	"info":       ListAttentionLogsUrgencyFilterInfo,
	"all":        ListAttentionLogsUrgencyFilterAll,
}

// GetListAttentionLogsUrgencyFilterEnumValues Enumerates the set of values for ListAttentionLogsUrgencyFilterEnum
func GetListAttentionLogsUrgencyFilterEnumValues() []ListAttentionLogsUrgencyFilterEnum {
	values := make([]ListAttentionLogsUrgencyFilterEnum, 0)
	for _, v := range mappingListAttentionLogsUrgencyFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttentionLogsUrgencyFilterEnumStringValues Enumerates the set of values in String for ListAttentionLogsUrgencyFilterEnum
func GetListAttentionLogsUrgencyFilterEnumStringValues() []string {
	return []string{
		"IMMEDIATE",
		"SOON",
		"DEFERRABLE",
		"INFO",
		"ALL",
	}
}

// GetMappingListAttentionLogsUrgencyFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttentionLogsUrgencyFilterEnum(val string) (ListAttentionLogsUrgencyFilterEnum, bool) {
	enum, ok := mappingListAttentionLogsUrgencyFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttentionLogsTypeFilterEnum Enum with underlying type: string
type ListAttentionLogsTypeFilterEnum string

// Set of constants representing the allowable values for ListAttentionLogsTypeFilterEnum
const (
	ListAttentionLogsTypeFilterUnknown       ListAttentionLogsTypeFilterEnum = "UNKNOWN"
	ListAttentionLogsTypeFilterIncidentError ListAttentionLogsTypeFilterEnum = "INCIDENT_ERROR"
	ListAttentionLogsTypeFilterError         ListAttentionLogsTypeFilterEnum = "ERROR"
	ListAttentionLogsTypeFilterWarning       ListAttentionLogsTypeFilterEnum = "WARNING"
	ListAttentionLogsTypeFilterNotification  ListAttentionLogsTypeFilterEnum = "NOTIFICATION"
	ListAttentionLogsTypeFilterTrace         ListAttentionLogsTypeFilterEnum = "TRACE"
	ListAttentionLogsTypeFilterAll           ListAttentionLogsTypeFilterEnum = "ALL"
)

var mappingListAttentionLogsTypeFilterEnum = map[string]ListAttentionLogsTypeFilterEnum{
	"UNKNOWN":        ListAttentionLogsTypeFilterUnknown,
	"INCIDENT_ERROR": ListAttentionLogsTypeFilterIncidentError,
	"ERROR":          ListAttentionLogsTypeFilterError,
	"WARNING":        ListAttentionLogsTypeFilterWarning,
	"NOTIFICATION":   ListAttentionLogsTypeFilterNotification,
	"TRACE":          ListAttentionLogsTypeFilterTrace,
	"ALL":            ListAttentionLogsTypeFilterAll,
}

var mappingListAttentionLogsTypeFilterEnumLowerCase = map[string]ListAttentionLogsTypeFilterEnum{
	"unknown":        ListAttentionLogsTypeFilterUnknown,
	"incident_error": ListAttentionLogsTypeFilterIncidentError,
	"error":          ListAttentionLogsTypeFilterError,
	"warning":        ListAttentionLogsTypeFilterWarning,
	"notification":   ListAttentionLogsTypeFilterNotification,
	"trace":          ListAttentionLogsTypeFilterTrace,
	"all":            ListAttentionLogsTypeFilterAll,
}

// GetListAttentionLogsTypeFilterEnumValues Enumerates the set of values for ListAttentionLogsTypeFilterEnum
func GetListAttentionLogsTypeFilterEnumValues() []ListAttentionLogsTypeFilterEnum {
	values := make([]ListAttentionLogsTypeFilterEnum, 0)
	for _, v := range mappingListAttentionLogsTypeFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttentionLogsTypeFilterEnumStringValues Enumerates the set of values in String for ListAttentionLogsTypeFilterEnum
func GetListAttentionLogsTypeFilterEnumStringValues() []string {
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

// GetMappingListAttentionLogsTypeFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttentionLogsTypeFilterEnum(val string) (ListAttentionLogsTypeFilterEnum, bool) {
	enum, ok := mappingListAttentionLogsTypeFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttentionLogsSortByEnum Enum with underlying type: string
type ListAttentionLogsSortByEnum string

// Set of constants representing the allowable values for ListAttentionLogsSortByEnum
const (
	ListAttentionLogsSortByUrgency    ListAttentionLogsSortByEnum = "URGENCY"
	ListAttentionLogsSortByType       ListAttentionLogsSortByEnum = "TYPE"
	ListAttentionLogsSortByMessage    ListAttentionLogsSortByEnum = "MESSAGE"
	ListAttentionLogsSortByTimestamp  ListAttentionLogsSortByEnum = "TIMESTAMP"
	ListAttentionLogsSortByScope      ListAttentionLogsSortByEnum = "SCOPE"
	ListAttentionLogsSortByTargetUser ListAttentionLogsSortByEnum = "TARGET_USER"
)

var mappingListAttentionLogsSortByEnum = map[string]ListAttentionLogsSortByEnum{
	"URGENCY":     ListAttentionLogsSortByUrgency,
	"TYPE":        ListAttentionLogsSortByType,
	"MESSAGE":     ListAttentionLogsSortByMessage,
	"TIMESTAMP":   ListAttentionLogsSortByTimestamp,
	"SCOPE":       ListAttentionLogsSortByScope,
	"TARGET_USER": ListAttentionLogsSortByTargetUser,
}

var mappingListAttentionLogsSortByEnumLowerCase = map[string]ListAttentionLogsSortByEnum{
	"urgency":     ListAttentionLogsSortByUrgency,
	"type":        ListAttentionLogsSortByType,
	"message":     ListAttentionLogsSortByMessage,
	"timestamp":   ListAttentionLogsSortByTimestamp,
	"scope":       ListAttentionLogsSortByScope,
	"target_user": ListAttentionLogsSortByTargetUser,
}

// GetListAttentionLogsSortByEnumValues Enumerates the set of values for ListAttentionLogsSortByEnum
func GetListAttentionLogsSortByEnumValues() []ListAttentionLogsSortByEnum {
	values := make([]ListAttentionLogsSortByEnum, 0)
	for _, v := range mappingListAttentionLogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttentionLogsSortByEnumStringValues Enumerates the set of values in String for ListAttentionLogsSortByEnum
func GetListAttentionLogsSortByEnumStringValues() []string {
	return []string{
		"URGENCY",
		"TYPE",
		"MESSAGE",
		"TIMESTAMP",
		"SCOPE",
		"TARGET_USER",
	}
}

// GetMappingListAttentionLogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttentionLogsSortByEnum(val string) (ListAttentionLogsSortByEnum, bool) {
	enum, ok := mappingListAttentionLogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAttentionLogsSortOrderEnum Enum with underlying type: string
type ListAttentionLogsSortOrderEnum string

// Set of constants representing the allowable values for ListAttentionLogsSortOrderEnum
const (
	ListAttentionLogsSortOrderAsc  ListAttentionLogsSortOrderEnum = "ASC"
	ListAttentionLogsSortOrderDesc ListAttentionLogsSortOrderEnum = "DESC"
)

var mappingListAttentionLogsSortOrderEnum = map[string]ListAttentionLogsSortOrderEnum{
	"ASC":  ListAttentionLogsSortOrderAsc,
	"DESC": ListAttentionLogsSortOrderDesc,
}

var mappingListAttentionLogsSortOrderEnumLowerCase = map[string]ListAttentionLogsSortOrderEnum{
	"asc":  ListAttentionLogsSortOrderAsc,
	"desc": ListAttentionLogsSortOrderDesc,
}

// GetListAttentionLogsSortOrderEnumValues Enumerates the set of values for ListAttentionLogsSortOrderEnum
func GetListAttentionLogsSortOrderEnumValues() []ListAttentionLogsSortOrderEnum {
	values := make([]ListAttentionLogsSortOrderEnum, 0)
	for _, v := range mappingListAttentionLogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAttentionLogsSortOrderEnumStringValues Enumerates the set of values in String for ListAttentionLogsSortOrderEnum
func GetListAttentionLogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAttentionLogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAttentionLogsSortOrderEnum(val string) (ListAttentionLogsSortOrderEnum, bool) {
	enum, ok := mappingListAttentionLogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
