// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeDatabaseInsightResourceForecastTrendRequest wrapper for the SummarizeDatabaseInsightResourceForecastTrend operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeDatabaseInsightResourceForecastTrend.go.html to see an example of how to use SummarizeDatabaseInsightResourceForecastTrendRequest.
type SummarizeDatabaseInsightResourceForecastTrendRequest struct {

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
	DatabaseType []SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum `contributesTo:"query" name:"databaseType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Choose the type of statistic metric data to be used for forecasting.
	Statistic SummarizeDatabaseInsightResourceForecastTrendStatisticEnum `mandatory:"false" contributesTo:"query" name:"statistic" omitEmpty:"true"`

	// Number of days used for utilization forecast analysis.
	ForecastDays *int `mandatory:"false" contributesTo:"query" name:"forecastDays"`

	// Choose algorithm model for the forecasting.
	// Possible values:
	//   - LINEAR: Uses linear regression algorithm for forecasting.
	//   - ML_AUTO: Automatically detects best algorithm to use for forecasting.
	//   - ML_NO_AUTO: Automatically detects seasonality of the data for forecasting using linear or seasonal algorithm.
	ForecastModel SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum `mandatory:"false" contributesTo:"query" name:"forecastModel" omitEmpty:"true"`

	// Filter by utilization level by the following buckets:
	//   - HIGH_UTILIZATION: DBs with utilization greater or equal than 75.
	//   - LOW_UTILIZATION: DBs with utilization lower than 25.
	//   - MEDIUM_HIGH_UTILIZATION: DBs with utilization greater or equal than 50 but lower than 75.
	//   - MEDIUM_LOW_UTILIZATION: DBs with utilization greater or equal than 25 but lower than 50.
	UtilizationLevel SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum `mandatory:"false" contributesTo:"query" name:"utilizationLevel" omitEmpty:"true"`

	// This parameter is used to change data's confidence level, this data is ingested by the
	// forecast algorithm.
	// Confidence is the probability of an interval to contain the expected population parameter.
	// Manipulation of this value will lead to different results.
	// If not set, default confidence value is 95%.
	Confidence *int `mandatory:"false" contributesTo:"query" name:"confidence"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Filter by one or more hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Tablespace name for a database
	TablespaceName *string `mandatory:"false" contributesTo:"query" name:"tablespaceName"`

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

func (request SummarizeDatabaseInsightResourceForecastTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceForecastTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceForecastTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeDatabaseInsightResourceForecastTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeDatabaseInsightResourceForecastTrendResponse wrapper for the SummarizeDatabaseInsightResourceForecastTrend operation
type SummarizeDatabaseInsightResourceForecastTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeDatabaseInsightResourceForecastTrendAggregation instances
	SummarizeDatabaseInsightResourceForecastTrendAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeDatabaseInsightResourceForecastTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeDatabaseInsightResourceForecastTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum
const (
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwS           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "ADW-S"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpS           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "ATP-S"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwD           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "ADW-D"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpD           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "ATP-D"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalPdb    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "EXTERNAL-PDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalNoncdb SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "EXTERNAL-NONCDB"
)

var mappingSummarizeDatabaseInsightResourceForecastTrendDatabaseType = map[string]SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum{
	"ADW-S":           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwS,
	"ATP-S":           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpS,
	"ADW-D":           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwD,
	"ATP-D":           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpD,
	"EXTERNAL-PDB":    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB": SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalNoncdb,
}

// GetSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnumValues() []SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendDatabaseType {
		values = append(values, v)
	}
	return values
}

// SummarizeDatabaseInsightResourceForecastTrendStatisticEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceForecastTrendStatisticEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceForecastTrendStatisticEnum
const (
	SummarizeDatabaseInsightResourceForecastTrendStatisticAvg SummarizeDatabaseInsightResourceForecastTrendStatisticEnum = "AVG"
	SummarizeDatabaseInsightResourceForecastTrendStatisticMax SummarizeDatabaseInsightResourceForecastTrendStatisticEnum = "MAX"
)

var mappingSummarizeDatabaseInsightResourceForecastTrendStatistic = map[string]SummarizeDatabaseInsightResourceForecastTrendStatisticEnum{
	"AVG": SummarizeDatabaseInsightResourceForecastTrendStatisticAvg,
	"MAX": SummarizeDatabaseInsightResourceForecastTrendStatisticMax,
}

// GetSummarizeDatabaseInsightResourceForecastTrendStatisticEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendStatisticEnum
func GetSummarizeDatabaseInsightResourceForecastTrendStatisticEnumValues() []SummarizeDatabaseInsightResourceForecastTrendStatisticEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendStatisticEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendStatistic {
		values = append(values, v)
	}
	return values
}

// SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum
const (
	SummarizeDatabaseInsightResourceForecastTrendForecastModelLinear   SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum = "LINEAR"
	SummarizeDatabaseInsightResourceForecastTrendForecastModelMlAuto   SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum = "ML_AUTO"
	SummarizeDatabaseInsightResourceForecastTrendForecastModelMlNoAuto SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum = "ML_NO_AUTO"
)

var mappingSummarizeDatabaseInsightResourceForecastTrendForecastModel = map[string]SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum{
	"LINEAR":     SummarizeDatabaseInsightResourceForecastTrendForecastModelLinear,
	"ML_AUTO":    SummarizeDatabaseInsightResourceForecastTrendForecastModelMlAuto,
	"ML_NO_AUTO": SummarizeDatabaseInsightResourceForecastTrendForecastModelMlNoAuto,
}

// GetSummarizeDatabaseInsightResourceForecastTrendForecastModelEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum
func GetSummarizeDatabaseInsightResourceForecastTrendForecastModelEnumValues() []SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendForecastModel {
		values = append(values, v)
	}
	return values
}

// SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum
const (
	SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelHighUtilization       SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum = "HIGH_UTILIZATION"
	SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelLowUtilization        SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum = "LOW_UTILIZATION"
	SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelMediumHighUtilization SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum = "MEDIUM_HIGH_UTILIZATION"
	SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelMediumLowUtilization  SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum = "MEDIUM_LOW_UTILIZATION"
)

var mappingSummarizeDatabaseInsightResourceForecastTrendUtilizationLevel = map[string]SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum{
	"HIGH_UTILIZATION":        SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelHighUtilization,
	"LOW_UTILIZATION":         SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelLowUtilization,
	"MEDIUM_HIGH_UTILIZATION": SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelMediumHighUtilization,
	"MEDIUM_LOW_UTILIZATION":  SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelMediumLowUtilization,
}

// GetSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum
func GetSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnumValues() []SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendUtilizationLevel {
		values = append(values, v)
	}
	return values
}
