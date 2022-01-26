// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeDatabaseInsightResourceStatisticsRequest wrapper for the SummarizeDatabaseInsightResourceStatistics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceStatistics.go.html to see an example of how to use SummarizeDatabaseInsightResourceStatisticsRequest.
type SummarizeDatabaseInsightResourceStatisticsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by resource metric.
	// Supported values are CPU , STORAGE, MEMORY and IO.
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

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	DatabaseType []SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

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
	SortOrder SummarizeDatabaseInsightResourceStatisticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The order in which resource statistics records are listed
	SortBy SummarizeDatabaseInsightResourceStatisticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter by one or more hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Flag to indicate if database instance level metrics should be returned. The flag is ignored when a host name filter is not applied.
	// When a hostname filter is applied this flag will determine whether to return metrics for the instances located on the specified host or for the
	// whole database which contains an instance on this host.
	IsDatabaseInstanceLevelMetrics *bool `mandatory:"false" contributesTo:"query" name:"isDatabaseInstanceLevelMetrics"`

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

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeDatabaseInsightResourceStatisticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceStatisticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceStatisticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeDatabaseInsightResourceStatisticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeDatabaseInsightResourceStatisticsResponse wrapper for the SummarizeDatabaseInsightResourceStatistics operation
type SummarizeDatabaseInsightResourceStatisticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeDatabaseInsightResourceStatisticsAggregationCollection instances
	SummarizeDatabaseInsightResourceStatisticsAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeDatabaseInsightResourceStatisticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeDatabaseInsightResourceStatisticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum
const (
	SummarizeDatabaseInsightResourceStatisticsDatabaseTypeAdwS           SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum = "ADW-S"
	SummarizeDatabaseInsightResourceStatisticsDatabaseTypeAtpS           SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum = "ATP-S"
	SummarizeDatabaseInsightResourceStatisticsDatabaseTypeAdwD           SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum = "ADW-D"
	SummarizeDatabaseInsightResourceStatisticsDatabaseTypeAtpD           SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum = "ATP-D"
	SummarizeDatabaseInsightResourceStatisticsDatabaseTypeExternalPdb    SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum = "EXTERNAL-PDB"
	SummarizeDatabaseInsightResourceStatisticsDatabaseTypeExternalNoncdb SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum = "EXTERNAL-NONCDB"
)

var mappingSummarizeDatabaseInsightResourceStatisticsDatabaseType = map[string]SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum{
	"ADW-S":           SummarizeDatabaseInsightResourceStatisticsDatabaseTypeAdwS,
	"ATP-S":           SummarizeDatabaseInsightResourceStatisticsDatabaseTypeAtpS,
	"ADW-D":           SummarizeDatabaseInsightResourceStatisticsDatabaseTypeAdwD,
	"ATP-D":           SummarizeDatabaseInsightResourceStatisticsDatabaseTypeAtpD,
	"EXTERNAL-PDB":    SummarizeDatabaseInsightResourceStatisticsDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB": SummarizeDatabaseInsightResourceStatisticsDatabaseTypeExternalNoncdb,
}

// GetSummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnumValues() []SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum {
	values := make([]SummarizeDatabaseInsightResourceStatisticsDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceStatisticsDatabaseType {
		values = append(values, v)
	}
	return values
}

// SummarizeDatabaseInsightResourceStatisticsSortOrderEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceStatisticsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceStatisticsSortOrderEnum
const (
	SummarizeDatabaseInsightResourceStatisticsSortOrderAsc  SummarizeDatabaseInsightResourceStatisticsSortOrderEnum = "ASC"
	SummarizeDatabaseInsightResourceStatisticsSortOrderDesc SummarizeDatabaseInsightResourceStatisticsSortOrderEnum = "DESC"
)

var mappingSummarizeDatabaseInsightResourceStatisticsSortOrder = map[string]SummarizeDatabaseInsightResourceStatisticsSortOrderEnum{
	"ASC":  SummarizeDatabaseInsightResourceStatisticsSortOrderAsc,
	"DESC": SummarizeDatabaseInsightResourceStatisticsSortOrderDesc,
}

// GetSummarizeDatabaseInsightResourceStatisticsSortOrderEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceStatisticsSortOrderEnum
func GetSummarizeDatabaseInsightResourceStatisticsSortOrderEnumValues() []SummarizeDatabaseInsightResourceStatisticsSortOrderEnum {
	values := make([]SummarizeDatabaseInsightResourceStatisticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceStatisticsSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeDatabaseInsightResourceStatisticsSortByEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceStatisticsSortByEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceStatisticsSortByEnum
const (
	SummarizeDatabaseInsightResourceStatisticsSortByUtilizationpercent SummarizeDatabaseInsightResourceStatisticsSortByEnum = "utilizationPercent"
	SummarizeDatabaseInsightResourceStatisticsSortByUsage              SummarizeDatabaseInsightResourceStatisticsSortByEnum = "usage"
	SummarizeDatabaseInsightResourceStatisticsSortByUsagechangepercent SummarizeDatabaseInsightResourceStatisticsSortByEnum = "usageChangePercent"
	SummarizeDatabaseInsightResourceStatisticsSortByDatabasename       SummarizeDatabaseInsightResourceStatisticsSortByEnum = "databaseName"
	SummarizeDatabaseInsightResourceStatisticsSortByDatabasetype       SummarizeDatabaseInsightResourceStatisticsSortByEnum = "databaseType"
)

var mappingSummarizeDatabaseInsightResourceStatisticsSortBy = map[string]SummarizeDatabaseInsightResourceStatisticsSortByEnum{
	"utilizationPercent": SummarizeDatabaseInsightResourceStatisticsSortByUtilizationpercent,
	"usage":              SummarizeDatabaseInsightResourceStatisticsSortByUsage,
	"usageChangePercent": SummarizeDatabaseInsightResourceStatisticsSortByUsagechangepercent,
	"databaseName":       SummarizeDatabaseInsightResourceStatisticsSortByDatabasename,
	"databaseType":       SummarizeDatabaseInsightResourceStatisticsSortByDatabasetype,
}

// GetSummarizeDatabaseInsightResourceStatisticsSortByEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceStatisticsSortByEnum
func GetSummarizeDatabaseInsightResourceStatisticsSortByEnumValues() []SummarizeDatabaseInsightResourceStatisticsSortByEnum {
	values := make([]SummarizeDatabaseInsightResourceStatisticsSortByEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceStatisticsSortBy {
		values = append(values, v)
	}
	return values
}
