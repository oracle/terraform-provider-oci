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

// SummarizeAlertLogCountsRequest wrapper for the SummarizeAlertLogCounts operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/SummarizeAlertLogCounts.go.html to see an example of how to use SummarizeAlertLogCountsRequest.
type SummarizeAlertLogCountsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The optional greater than or equal to timestamp to filter the logs.
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to timestamp to filter the logs.
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

	// The optional parameter to filter the alert logs by log level.
	LevelFilter SummarizeAlertLogCountsLevelFilterEnum `mandatory:"false" contributesTo:"query" name:"levelFilter" omitEmpty:"true"`

	// The optional parameter used to group different alert logs.
	GroupBy SummarizeAlertLogCountsGroupByEnum `mandatory:"false" contributesTo:"query" name:"groupBy" omitEmpty:"true"`

	// The optional parameter to filter the attention or alert logs by type.
	TypeFilter SummarizeAlertLogCountsTypeFilterEnum `mandatory:"false" contributesTo:"query" name:"typeFilter" omitEmpty:"true"`

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

func (request SummarizeAlertLogCountsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAlertLogCountsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAlertLogCountsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAlertLogCountsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAlertLogCountsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAlertLogCountsLevelFilterEnum(string(request.LevelFilter)); !ok && request.LevelFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LevelFilter: %s. Supported values are: %s.", request.LevelFilter, strings.Join(GetSummarizeAlertLogCountsLevelFilterEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAlertLogCountsGroupByEnum(string(request.GroupBy)); !ok && request.GroupBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GroupBy: %s. Supported values are: %s.", request.GroupBy, strings.Join(GetSummarizeAlertLogCountsGroupByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAlertLogCountsTypeFilterEnum(string(request.TypeFilter)); !ok && request.TypeFilter != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TypeFilter: %s. Supported values are: %s.", request.TypeFilter, strings.Join(GetSummarizeAlertLogCountsTypeFilterEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAlertLogCountsResponse wrapper for the SummarizeAlertLogCounts operation
type SummarizeAlertLogCountsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AlertLogCountsCollection instances
	AlertLogCountsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAlertLogCountsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAlertLogCountsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAlertLogCountsLevelFilterEnum Enum with underlying type: string
type SummarizeAlertLogCountsLevelFilterEnum string

// Set of constants representing the allowable values for SummarizeAlertLogCountsLevelFilterEnum
const (
	SummarizeAlertLogCountsLevelFilterCritical  SummarizeAlertLogCountsLevelFilterEnum = "CRITICAL"
	SummarizeAlertLogCountsLevelFilterSevere    SummarizeAlertLogCountsLevelFilterEnum = "SEVERE"
	SummarizeAlertLogCountsLevelFilterImportant SummarizeAlertLogCountsLevelFilterEnum = "IMPORTANT"
	SummarizeAlertLogCountsLevelFilterNormal    SummarizeAlertLogCountsLevelFilterEnum = "NORMAL"
	SummarizeAlertLogCountsLevelFilterAll       SummarizeAlertLogCountsLevelFilterEnum = "ALL"
)

var mappingSummarizeAlertLogCountsLevelFilterEnum = map[string]SummarizeAlertLogCountsLevelFilterEnum{
	"CRITICAL":  SummarizeAlertLogCountsLevelFilterCritical,
	"SEVERE":    SummarizeAlertLogCountsLevelFilterSevere,
	"IMPORTANT": SummarizeAlertLogCountsLevelFilterImportant,
	"NORMAL":    SummarizeAlertLogCountsLevelFilterNormal,
	"ALL":       SummarizeAlertLogCountsLevelFilterAll,
}

var mappingSummarizeAlertLogCountsLevelFilterEnumLowerCase = map[string]SummarizeAlertLogCountsLevelFilterEnum{
	"critical":  SummarizeAlertLogCountsLevelFilterCritical,
	"severe":    SummarizeAlertLogCountsLevelFilterSevere,
	"important": SummarizeAlertLogCountsLevelFilterImportant,
	"normal":    SummarizeAlertLogCountsLevelFilterNormal,
	"all":       SummarizeAlertLogCountsLevelFilterAll,
}

// GetSummarizeAlertLogCountsLevelFilterEnumValues Enumerates the set of values for SummarizeAlertLogCountsLevelFilterEnum
func GetSummarizeAlertLogCountsLevelFilterEnumValues() []SummarizeAlertLogCountsLevelFilterEnum {
	values := make([]SummarizeAlertLogCountsLevelFilterEnum, 0)
	for _, v := range mappingSummarizeAlertLogCountsLevelFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAlertLogCountsLevelFilterEnumStringValues Enumerates the set of values in String for SummarizeAlertLogCountsLevelFilterEnum
func GetSummarizeAlertLogCountsLevelFilterEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"SEVERE",
		"IMPORTANT",
		"NORMAL",
		"ALL",
	}
}

// GetMappingSummarizeAlertLogCountsLevelFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAlertLogCountsLevelFilterEnum(val string) (SummarizeAlertLogCountsLevelFilterEnum, bool) {
	enum, ok := mappingSummarizeAlertLogCountsLevelFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAlertLogCountsGroupByEnum Enum with underlying type: string
type SummarizeAlertLogCountsGroupByEnum string

// Set of constants representing the allowable values for SummarizeAlertLogCountsGroupByEnum
const (
	SummarizeAlertLogCountsGroupByLevel SummarizeAlertLogCountsGroupByEnum = "LEVEL"
	SummarizeAlertLogCountsGroupByType  SummarizeAlertLogCountsGroupByEnum = "TYPE"
)

var mappingSummarizeAlertLogCountsGroupByEnum = map[string]SummarizeAlertLogCountsGroupByEnum{
	"LEVEL": SummarizeAlertLogCountsGroupByLevel,
	"TYPE":  SummarizeAlertLogCountsGroupByType,
}

var mappingSummarizeAlertLogCountsGroupByEnumLowerCase = map[string]SummarizeAlertLogCountsGroupByEnum{
	"level": SummarizeAlertLogCountsGroupByLevel,
	"type":  SummarizeAlertLogCountsGroupByType,
}

// GetSummarizeAlertLogCountsGroupByEnumValues Enumerates the set of values for SummarizeAlertLogCountsGroupByEnum
func GetSummarizeAlertLogCountsGroupByEnumValues() []SummarizeAlertLogCountsGroupByEnum {
	values := make([]SummarizeAlertLogCountsGroupByEnum, 0)
	for _, v := range mappingSummarizeAlertLogCountsGroupByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAlertLogCountsGroupByEnumStringValues Enumerates the set of values in String for SummarizeAlertLogCountsGroupByEnum
func GetSummarizeAlertLogCountsGroupByEnumStringValues() []string {
	return []string{
		"LEVEL",
		"TYPE",
	}
}

// GetMappingSummarizeAlertLogCountsGroupByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAlertLogCountsGroupByEnum(val string) (SummarizeAlertLogCountsGroupByEnum, bool) {
	enum, ok := mappingSummarizeAlertLogCountsGroupByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAlertLogCountsTypeFilterEnum Enum with underlying type: string
type SummarizeAlertLogCountsTypeFilterEnum string

// Set of constants representing the allowable values for SummarizeAlertLogCountsTypeFilterEnum
const (
	SummarizeAlertLogCountsTypeFilterUnknown       SummarizeAlertLogCountsTypeFilterEnum = "UNKNOWN"
	SummarizeAlertLogCountsTypeFilterIncidentError SummarizeAlertLogCountsTypeFilterEnum = "INCIDENT_ERROR"
	SummarizeAlertLogCountsTypeFilterError         SummarizeAlertLogCountsTypeFilterEnum = "ERROR"
	SummarizeAlertLogCountsTypeFilterWarning       SummarizeAlertLogCountsTypeFilterEnum = "WARNING"
	SummarizeAlertLogCountsTypeFilterNotification  SummarizeAlertLogCountsTypeFilterEnum = "NOTIFICATION"
	SummarizeAlertLogCountsTypeFilterTrace         SummarizeAlertLogCountsTypeFilterEnum = "TRACE"
	SummarizeAlertLogCountsTypeFilterAll           SummarizeAlertLogCountsTypeFilterEnum = "ALL"
)

var mappingSummarizeAlertLogCountsTypeFilterEnum = map[string]SummarizeAlertLogCountsTypeFilterEnum{
	"UNKNOWN":        SummarizeAlertLogCountsTypeFilterUnknown,
	"INCIDENT_ERROR": SummarizeAlertLogCountsTypeFilterIncidentError,
	"ERROR":          SummarizeAlertLogCountsTypeFilterError,
	"WARNING":        SummarizeAlertLogCountsTypeFilterWarning,
	"NOTIFICATION":   SummarizeAlertLogCountsTypeFilterNotification,
	"TRACE":          SummarizeAlertLogCountsTypeFilterTrace,
	"ALL":            SummarizeAlertLogCountsTypeFilterAll,
}

var mappingSummarizeAlertLogCountsTypeFilterEnumLowerCase = map[string]SummarizeAlertLogCountsTypeFilterEnum{
	"unknown":        SummarizeAlertLogCountsTypeFilterUnknown,
	"incident_error": SummarizeAlertLogCountsTypeFilterIncidentError,
	"error":          SummarizeAlertLogCountsTypeFilterError,
	"warning":        SummarizeAlertLogCountsTypeFilterWarning,
	"notification":   SummarizeAlertLogCountsTypeFilterNotification,
	"trace":          SummarizeAlertLogCountsTypeFilterTrace,
	"all":            SummarizeAlertLogCountsTypeFilterAll,
}

// GetSummarizeAlertLogCountsTypeFilterEnumValues Enumerates the set of values for SummarizeAlertLogCountsTypeFilterEnum
func GetSummarizeAlertLogCountsTypeFilterEnumValues() []SummarizeAlertLogCountsTypeFilterEnum {
	values := make([]SummarizeAlertLogCountsTypeFilterEnum, 0)
	for _, v := range mappingSummarizeAlertLogCountsTypeFilterEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAlertLogCountsTypeFilterEnumStringValues Enumerates the set of values in String for SummarizeAlertLogCountsTypeFilterEnum
func GetSummarizeAlertLogCountsTypeFilterEnumStringValues() []string {
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

// GetMappingSummarizeAlertLogCountsTypeFilterEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAlertLogCountsTypeFilterEnum(val string) (SummarizeAlertLogCountsTypeFilterEnum, bool) {
	enum, ok := mappingSummarizeAlertLogCountsTypeFilterEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
