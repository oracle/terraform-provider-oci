// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeHostInsightResourceStatisticsRequest wrapper for the SummarizeHostInsightResourceStatistics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceStatistics.go.html to see an example of how to use SummarizeHostInsightResourceStatisticsRequest.
type SummarizeHostInsightResourceStatisticsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by host resource metric.
	// Supported values are CPU, MEMORY, LOGICAL_MEMORY, STORAGE and NETWORK.
	ResourceMetric *string `mandatory:"true" contributesTo:"query" name:"resourceMetric"`

	// Specify time period in ISO 8601 format with respect to current time.
	// Default is last 30 days represented by P30D.
	// If timeInterval is specified, then timeIntervalStart and timeIntervalEnd will be ignored.
	// Examples  P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months), . Maximum value allowed is 25 months prior to current time (P25M).
	AnalysisTimeInterval *string `mandatory:"false" contributesTo:"query" name:"analysisTimeInterval"`

	// Analysis start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// The minimum allowed value is 2 years prior to the current day.
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	// If analysisTimeInterval is specified, this parameter is ignored.
	TimeIntervalStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalStart"`

	// Analysis end time in UTC in ISO 8601 format(exclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd are used together.
	// If timeIntervalEnd is not specified, current time is used as timeIntervalEnd.
	TimeIntervalEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalEnd"`

	// Filter by one or more platform types.
	// Supported platformType(s) for MACS-managed external host insight: [LINUX, SOLARIS, WINDOWS].
	// Supported platformType(s) for MACS-managed cloud host insight: [LINUX].
	// Supported platformType(s) for EM-managed external host insight: [LINUX, SOLARIS, SUNOS, ZLINUX, WINDOWS, AIX, HP-UX].
	PlatformType []SummarizeHostInsightResourceStatisticsPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Percentile values of daily usage to be used for computing the aggregate resource usage.
	Percentile *int `mandatory:"false" contributesTo:"query" name:"percentile"`

	// Return data of a specific insight
	// Possible values are High Utilization, Low Utilization, Any ,High Utilization Forecast,
	// Low Utilization Forecast
	InsightBy *string `mandatory:"false" contributesTo:"query" name:"insightBy"`

	// Number of days used for utilization forecast analysis.
	ForecastDays *int `mandatory:"false" contributesTo:"query" name:"forecastDays"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeHostInsightResourceStatisticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The order in which resource statistics records are listed.
	SortBy SummarizeHostInsightResourceStatisticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Filter by one or more host types.
	// Possible values are CLOUD-HOST, EXTERNAL-HOST, COMANAGED-VM-HOST, COMANAGED-BM-HOST, COMANAGED-EXACS-HOST, COMANAGED-EXACC-HOST
	HostType []string `contributesTo:"query" name:"hostType" collectionFormat:"multi"`

	// Optional OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host (Compute Id)
	HostId *string `mandatory:"false" contributesTo:"query" name:"hostId"`

	// Optional list of Exadata Insight VM cluster name.
	VmclusterName []string `contributesTo:"query" name:"vmclusterName" collectionFormat:"multi"`

	// Percent value in which a resource metric is considered highly utilized.
	HighUtilizationThreshold *int `mandatory:"false" contributesTo:"query" name:"highUtilizationThreshold"`

	// Percent value in which a resource metric is considered low utilized.
	LowUtilizationThreshold *int `mandatory:"false" contributesTo:"query" name:"lowUtilizationThreshold"`

	// Resource Status
	Status []ResourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeHostInsightResourceStatisticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightResourceStatisticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightResourceStatisticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightResourceStatisticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeHostInsightResourceStatisticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.PlatformType {
		if _, ok := GetMappingSummarizeHostInsightResourceStatisticsPlatformTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlatformType: %s. Supported values are: %s.", val, strings.Join(GetSummarizeHostInsightResourceStatisticsPlatformTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeHostInsightResourceStatisticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeHostInsightResourceStatisticsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeHostInsightResourceStatisticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeHostInsightResourceStatisticsSortByEnumStringValues(), ",")))
	}
	for _, val := range request.Status {
		if _, ok := GetMappingResourceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetResourceStatusEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightResourceStatisticsResponse wrapper for the SummarizeHostInsightResourceStatistics operation
type SummarizeHostInsightResourceStatisticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeHostInsightResourceStatisticsAggregationCollection instances
	SummarizeHostInsightResourceStatisticsAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeHostInsightResourceStatisticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightResourceStatisticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightResourceStatisticsPlatformTypeEnum Enum with underlying type: string
type SummarizeHostInsightResourceStatisticsPlatformTypeEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceStatisticsPlatformTypeEnum
const (
	SummarizeHostInsightResourceStatisticsPlatformTypeLinux   SummarizeHostInsightResourceStatisticsPlatformTypeEnum = "LINUX"
	SummarizeHostInsightResourceStatisticsPlatformTypeSolaris SummarizeHostInsightResourceStatisticsPlatformTypeEnum = "SOLARIS"
	SummarizeHostInsightResourceStatisticsPlatformTypeSunos   SummarizeHostInsightResourceStatisticsPlatformTypeEnum = "SUNOS"
	SummarizeHostInsightResourceStatisticsPlatformTypeZlinux  SummarizeHostInsightResourceStatisticsPlatformTypeEnum = "ZLINUX"
	SummarizeHostInsightResourceStatisticsPlatformTypeWindows SummarizeHostInsightResourceStatisticsPlatformTypeEnum = "WINDOWS"
	SummarizeHostInsightResourceStatisticsPlatformTypeAix     SummarizeHostInsightResourceStatisticsPlatformTypeEnum = "AIX"
	SummarizeHostInsightResourceStatisticsPlatformTypeHpUx    SummarizeHostInsightResourceStatisticsPlatformTypeEnum = "HP_UX"
)

var mappingSummarizeHostInsightResourceStatisticsPlatformTypeEnum = map[string]SummarizeHostInsightResourceStatisticsPlatformTypeEnum{
	"LINUX":   SummarizeHostInsightResourceStatisticsPlatformTypeLinux,
	"SOLARIS": SummarizeHostInsightResourceStatisticsPlatformTypeSolaris,
	"SUNOS":   SummarizeHostInsightResourceStatisticsPlatformTypeSunos,
	"ZLINUX":  SummarizeHostInsightResourceStatisticsPlatformTypeZlinux,
	"WINDOWS": SummarizeHostInsightResourceStatisticsPlatformTypeWindows,
	"AIX":     SummarizeHostInsightResourceStatisticsPlatformTypeAix,
	"HP_UX":   SummarizeHostInsightResourceStatisticsPlatformTypeHpUx,
}

var mappingSummarizeHostInsightResourceStatisticsPlatformTypeEnumLowerCase = map[string]SummarizeHostInsightResourceStatisticsPlatformTypeEnum{
	"linux":   SummarizeHostInsightResourceStatisticsPlatformTypeLinux,
	"solaris": SummarizeHostInsightResourceStatisticsPlatformTypeSolaris,
	"sunos":   SummarizeHostInsightResourceStatisticsPlatformTypeSunos,
	"zlinux":  SummarizeHostInsightResourceStatisticsPlatformTypeZlinux,
	"windows": SummarizeHostInsightResourceStatisticsPlatformTypeWindows,
	"aix":     SummarizeHostInsightResourceStatisticsPlatformTypeAix,
	"hp_ux":   SummarizeHostInsightResourceStatisticsPlatformTypeHpUx,
}

// GetSummarizeHostInsightResourceStatisticsPlatformTypeEnumValues Enumerates the set of values for SummarizeHostInsightResourceStatisticsPlatformTypeEnum
func GetSummarizeHostInsightResourceStatisticsPlatformTypeEnumValues() []SummarizeHostInsightResourceStatisticsPlatformTypeEnum {
	values := make([]SummarizeHostInsightResourceStatisticsPlatformTypeEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceStatisticsPlatformTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceStatisticsPlatformTypeEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceStatisticsPlatformTypeEnum
func GetSummarizeHostInsightResourceStatisticsPlatformTypeEnumStringValues() []string {
	return []string{
		"LINUX",
		"SOLARIS",
		"SUNOS",
		"ZLINUX",
		"WINDOWS",
		"AIX",
		"HP_UX",
	}
}

// GetMappingSummarizeHostInsightResourceStatisticsPlatformTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceStatisticsPlatformTypeEnum(val string) (SummarizeHostInsightResourceStatisticsPlatformTypeEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceStatisticsPlatformTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightResourceStatisticsSortOrderEnum Enum with underlying type: string
type SummarizeHostInsightResourceStatisticsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceStatisticsSortOrderEnum
const (
	SummarizeHostInsightResourceStatisticsSortOrderAsc  SummarizeHostInsightResourceStatisticsSortOrderEnum = "ASC"
	SummarizeHostInsightResourceStatisticsSortOrderDesc SummarizeHostInsightResourceStatisticsSortOrderEnum = "DESC"
)

var mappingSummarizeHostInsightResourceStatisticsSortOrderEnum = map[string]SummarizeHostInsightResourceStatisticsSortOrderEnum{
	"ASC":  SummarizeHostInsightResourceStatisticsSortOrderAsc,
	"DESC": SummarizeHostInsightResourceStatisticsSortOrderDesc,
}

var mappingSummarizeHostInsightResourceStatisticsSortOrderEnumLowerCase = map[string]SummarizeHostInsightResourceStatisticsSortOrderEnum{
	"asc":  SummarizeHostInsightResourceStatisticsSortOrderAsc,
	"desc": SummarizeHostInsightResourceStatisticsSortOrderDesc,
}

// GetSummarizeHostInsightResourceStatisticsSortOrderEnumValues Enumerates the set of values for SummarizeHostInsightResourceStatisticsSortOrderEnum
func GetSummarizeHostInsightResourceStatisticsSortOrderEnumValues() []SummarizeHostInsightResourceStatisticsSortOrderEnum {
	values := make([]SummarizeHostInsightResourceStatisticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceStatisticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceStatisticsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceStatisticsSortOrderEnum
func GetSummarizeHostInsightResourceStatisticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeHostInsightResourceStatisticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceStatisticsSortOrderEnum(val string) (SummarizeHostInsightResourceStatisticsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceStatisticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeHostInsightResourceStatisticsSortByEnum Enum with underlying type: string
type SummarizeHostInsightResourceStatisticsSortByEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceStatisticsSortByEnum
const (
	SummarizeHostInsightResourceStatisticsSortByUtilizationpercent SummarizeHostInsightResourceStatisticsSortByEnum = "utilizationPercent"
	SummarizeHostInsightResourceStatisticsSortByUsage              SummarizeHostInsightResourceStatisticsSortByEnum = "usage"
	SummarizeHostInsightResourceStatisticsSortByUsagechangepercent SummarizeHostInsightResourceStatisticsSortByEnum = "usageChangePercent"
	SummarizeHostInsightResourceStatisticsSortByHostname           SummarizeHostInsightResourceStatisticsSortByEnum = "hostName"
	SummarizeHostInsightResourceStatisticsSortByPlatformtype       SummarizeHostInsightResourceStatisticsSortByEnum = "platformType"
)

var mappingSummarizeHostInsightResourceStatisticsSortByEnum = map[string]SummarizeHostInsightResourceStatisticsSortByEnum{
	"utilizationPercent": SummarizeHostInsightResourceStatisticsSortByUtilizationpercent,
	"usage":              SummarizeHostInsightResourceStatisticsSortByUsage,
	"usageChangePercent": SummarizeHostInsightResourceStatisticsSortByUsagechangepercent,
	"hostName":           SummarizeHostInsightResourceStatisticsSortByHostname,
	"platformType":       SummarizeHostInsightResourceStatisticsSortByPlatformtype,
}

var mappingSummarizeHostInsightResourceStatisticsSortByEnumLowerCase = map[string]SummarizeHostInsightResourceStatisticsSortByEnum{
	"utilizationpercent": SummarizeHostInsightResourceStatisticsSortByUtilizationpercent,
	"usage":              SummarizeHostInsightResourceStatisticsSortByUsage,
	"usagechangepercent": SummarizeHostInsightResourceStatisticsSortByUsagechangepercent,
	"hostname":           SummarizeHostInsightResourceStatisticsSortByHostname,
	"platformtype":       SummarizeHostInsightResourceStatisticsSortByPlatformtype,
}

// GetSummarizeHostInsightResourceStatisticsSortByEnumValues Enumerates the set of values for SummarizeHostInsightResourceStatisticsSortByEnum
func GetSummarizeHostInsightResourceStatisticsSortByEnumValues() []SummarizeHostInsightResourceStatisticsSortByEnum {
	values := make([]SummarizeHostInsightResourceStatisticsSortByEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceStatisticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightResourceStatisticsSortByEnumStringValues Enumerates the set of values in String for SummarizeHostInsightResourceStatisticsSortByEnum
func GetSummarizeHostInsightResourceStatisticsSortByEnumStringValues() []string {
	return []string{
		"utilizationPercent",
		"usage",
		"usageChangePercent",
		"hostName",
		"platformType",
	}
}

// GetMappingSummarizeHostInsightResourceStatisticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightResourceStatisticsSortByEnum(val string) (SummarizeHostInsightResourceStatisticsSortByEnum, bool) {
	enum, ok := mappingSummarizeHostInsightResourceStatisticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
