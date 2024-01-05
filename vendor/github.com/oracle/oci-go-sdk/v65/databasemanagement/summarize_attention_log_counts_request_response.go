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

// SummarizeAttentionLogCountsRequest wrapper for the SummarizeAttentionLogCounts operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAttentionLogCounts.go.html to see an example of how to use SummarizeAttentionLogCountsRequest.
type SummarizeAttentionLogCountsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The optional greater than or equal to timestamp to filter the logs.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to timestamp to filter the logs.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The optional parameter to filter the attention logs by urgency.
	UrgencyFilter SummarizeAttentionLogCountsUrgencyFilterEnum `mandatory:"false" contributesTo:"query" name:"urgencyFilter" omitEmpty:"true"`

	// The optional parameter used to group different attention logs.
	GroupBy SummarizeAttentionLogCountsGroupByEnum `mandatory:"false" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// The optional parameter to filter the attention or alert logs by type.
	TypeFilter SummarizeAttentionLogCountsTypeFilterEnum `mandatory:"false" contributesTo:"query" name:"typeFilter" omitEmpty:"true"`

	// The optional query parameter to filter the attention or alert logs by search text.
	LogSearchText *string `mandatory:"false" contributesTo:"query" name:"logSearchText"`

	// The flag to indicate whether the search text is regular expression or not.
	IsRegularExpression *bool `mandatory:"false" contributesTo:"query" name:"isRegularExpression"`

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

func (request SummarizeAttentionLogCountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAttentionLogCountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAttentionLogCountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAttentionLogCountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAttentionLogCountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAttentionLogCountsUrgencyFilterEnum(string(request.UrgencyFilter)); !ok && request.UrgencyFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UrgencyFilter: %s. Supported values are: %s.", request.UrgencyFilter, strings.Join(GetSummarizeAttentionLogCountsUrgencyFilterEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAttentionLogCountsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetSummarizeAttentionLogCountsGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAttentionLogCountsTypeFilterEnum(string(request.TypeFilter)); !ok && request.TypeFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TypeFilter: %s. Supported values are: %s.", request.TypeFilter, strings.Join(GetSummarizeAttentionLogCountsTypeFilterEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAttentionLogCountsResponse wrapper for the SummarizeAttentionLogCounts operation
type SummarizeAttentionLogCountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AttentionLogCountsCollection instances
	AttentionLogCountsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAttentionLogCountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAttentionLogCountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAttentionLogCountsUrgencyFilterEnum Enum with underlying type: string
type SummarizeAttentionLogCountsUrgencyFilterEnum string

// Set of constants representing the allowable values for SummarizeAttentionLogCountsUrgencyFilterEnum
const (
	SummarizeAttentionLogCountsUrgencyFilterImmediate  SummarizeAttentionLogCountsUrgencyFilterEnum = "IMMEDIATE"
	SummarizeAttentionLogCountsUrgencyFilterSoon       SummarizeAttentionLogCountsUrgencyFilterEnum = "SOON"
	SummarizeAttentionLogCountsUrgencyFilterDeferrable SummarizeAttentionLogCountsUrgencyFilterEnum = "DEFERRABLE"
	SummarizeAttentionLogCountsUrgencyFilterInfo       SummarizeAttentionLogCountsUrgencyFilterEnum = "INFO"
	SummarizeAttentionLogCountsUrgencyFilterAll        SummarizeAttentionLogCountsUrgencyFilterEnum = "ALL"
)

var mappingSummarizeAttentionLogCountsUrgencyFilterEnum = map[string]SummarizeAttentionLogCountsUrgencyFilterEnum{
	"IMMEDIATE":  SummarizeAttentionLogCountsUrgencyFilterImmediate,
	"SOON":       SummarizeAttentionLogCountsUrgencyFilterSoon,
	"DEFERRABLE": SummarizeAttentionLogCountsUrgencyFilterDeferrable,
	"INFO":       SummarizeAttentionLogCountsUrgencyFilterInfo,
	"ALL":        SummarizeAttentionLogCountsUrgencyFilterAll,
}

var mappingSummarizeAttentionLogCountsUrgencyFilterEnumLowerCase = map[string]SummarizeAttentionLogCountsUrgencyFilterEnum{
	"immediate":  SummarizeAttentionLogCountsUrgencyFilterImmediate,
	"soon":       SummarizeAttentionLogCountsUrgencyFilterSoon,
	"deferrable": SummarizeAttentionLogCountsUrgencyFilterDeferrable,
	"info":       SummarizeAttentionLogCountsUrgencyFilterInfo,
	"all":        SummarizeAttentionLogCountsUrgencyFilterAll,
}

// GetSummarizeAttentionLogCountsUrgencyFilterEnumValues Enumerates the set of values for SummarizeAttentionLogCountsUrgencyFilterEnum
func GetSummarizeAttentionLogCountsUrgencyFilterEnumValues() []SummarizeAttentionLogCountsUrgencyFilterEnum {
	values := make([]SummarizeAttentionLogCountsUrgencyFilterEnum, 0)
	for _, v := range mappingSummarizeAttentionLogCountsUrgencyFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAttentionLogCountsUrgencyFilterEnumStringValues Enumerates the set of values in String for SummarizeAttentionLogCountsUrgencyFilterEnum
func GetSummarizeAttentionLogCountsUrgencyFilterEnumStringValues() []string {
	return []string{
		"IMMEDIATE",
		"SOON",
		"DEFERRABLE",
		"INFO",
		"ALL",
	}
}

// GetMappingSummarizeAttentionLogCountsUrgencyFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAttentionLogCountsUrgencyFilterEnum(val string) (SummarizeAttentionLogCountsUrgencyFilterEnum, bool) {
	enum, ok := mappingSummarizeAttentionLogCountsUrgencyFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAttentionLogCountsGroupByEnum Enum with underlying type: string
type SummarizeAttentionLogCountsGroupByEnum string

// Set of constants representing the allowable values for SummarizeAttentionLogCountsGroupByEnum
const (
	SummarizeAttentionLogCountsGroupByUrgency SummarizeAttentionLogCountsGroupByEnum = "URGENCY"
	SummarizeAttentionLogCountsGroupByType    SummarizeAttentionLogCountsGroupByEnum = "TYPE"
)

var mappingSummarizeAttentionLogCountsGroupByEnum = map[string]SummarizeAttentionLogCountsGroupByEnum{
	"URGENCY": SummarizeAttentionLogCountsGroupByUrgency,
	"TYPE":    SummarizeAttentionLogCountsGroupByType,
}

var mappingSummarizeAttentionLogCountsGroupByEnumLowerCase = map[string]SummarizeAttentionLogCountsGroupByEnum{
	"urgency": SummarizeAttentionLogCountsGroupByUrgency,
	"type":    SummarizeAttentionLogCountsGroupByType,
}

// GetSummarizeAttentionLogCountsGroupByEnumValues Enumerates the set of values for SummarizeAttentionLogCountsGroupByEnum
func GetSummarizeAttentionLogCountsGroupByEnumValues() []SummarizeAttentionLogCountsGroupByEnum {
	values := make([]SummarizeAttentionLogCountsGroupByEnum, 0)
	for _, v := range mappingSummarizeAttentionLogCountsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAttentionLogCountsGroupByEnumStringValues Enumerates the set of values in String for SummarizeAttentionLogCountsGroupByEnum
func GetSummarizeAttentionLogCountsGroupByEnumStringValues() []string {
	return []string{
		"URGENCY",
		"TYPE",
	}
}

// GetMappingSummarizeAttentionLogCountsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAttentionLogCountsGroupByEnum(val string) (SummarizeAttentionLogCountsGroupByEnum, bool) {
	enum, ok := mappingSummarizeAttentionLogCountsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAttentionLogCountsTypeFilterEnum Enum with underlying type: string
type SummarizeAttentionLogCountsTypeFilterEnum string

// Set of constants representing the allowable values for SummarizeAttentionLogCountsTypeFilterEnum
const (
	SummarizeAttentionLogCountsTypeFilterUnknown       SummarizeAttentionLogCountsTypeFilterEnum = "UNKNOWN"
	SummarizeAttentionLogCountsTypeFilterIncidentError SummarizeAttentionLogCountsTypeFilterEnum = "INCIDENT_ERROR"
	SummarizeAttentionLogCountsTypeFilterError         SummarizeAttentionLogCountsTypeFilterEnum = "ERROR"
	SummarizeAttentionLogCountsTypeFilterWarning       SummarizeAttentionLogCountsTypeFilterEnum = "WARNING"
	SummarizeAttentionLogCountsTypeFilterNotification  SummarizeAttentionLogCountsTypeFilterEnum = "NOTIFICATION"
	SummarizeAttentionLogCountsTypeFilterTrace         SummarizeAttentionLogCountsTypeFilterEnum = "TRACE"
	SummarizeAttentionLogCountsTypeFilterAll           SummarizeAttentionLogCountsTypeFilterEnum = "ALL"
)

var mappingSummarizeAttentionLogCountsTypeFilterEnum = map[string]SummarizeAttentionLogCountsTypeFilterEnum{
	"UNKNOWN":        SummarizeAttentionLogCountsTypeFilterUnknown,
	"INCIDENT_ERROR": SummarizeAttentionLogCountsTypeFilterIncidentError,
	"ERROR":          SummarizeAttentionLogCountsTypeFilterError,
	"WARNING":        SummarizeAttentionLogCountsTypeFilterWarning,
	"NOTIFICATION":   SummarizeAttentionLogCountsTypeFilterNotification,
	"TRACE":          SummarizeAttentionLogCountsTypeFilterTrace,
	"ALL":            SummarizeAttentionLogCountsTypeFilterAll,
}

var mappingSummarizeAttentionLogCountsTypeFilterEnumLowerCase = map[string]SummarizeAttentionLogCountsTypeFilterEnum{
	"unknown":        SummarizeAttentionLogCountsTypeFilterUnknown,
	"incident_error": SummarizeAttentionLogCountsTypeFilterIncidentError,
	"error":          SummarizeAttentionLogCountsTypeFilterError,
	"warning":        SummarizeAttentionLogCountsTypeFilterWarning,
	"notification":   SummarizeAttentionLogCountsTypeFilterNotification,
	"trace":          SummarizeAttentionLogCountsTypeFilterTrace,
	"all":            SummarizeAttentionLogCountsTypeFilterAll,
}

// GetSummarizeAttentionLogCountsTypeFilterEnumValues Enumerates the set of values for SummarizeAttentionLogCountsTypeFilterEnum
func GetSummarizeAttentionLogCountsTypeFilterEnumValues() []SummarizeAttentionLogCountsTypeFilterEnum {
	values := make([]SummarizeAttentionLogCountsTypeFilterEnum, 0)
	for _, v := range mappingSummarizeAttentionLogCountsTypeFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAttentionLogCountsTypeFilterEnumStringValues Enumerates the set of values in String for SummarizeAttentionLogCountsTypeFilterEnum
func GetSummarizeAttentionLogCountsTypeFilterEnumStringValues() []string {
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

// GetMappingSummarizeAttentionLogCountsTypeFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAttentionLogCountsTypeFilterEnum(val string) (SummarizeAttentionLogCountsTypeFilterEnum, bool) {
	enum, ok := mappingSummarizeAttentionLogCountsTypeFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
