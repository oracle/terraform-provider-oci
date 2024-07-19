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

// SummarizeDatabaseInsightResourceForecastTrendRequest wrapper for the SummarizeDatabaseInsightResourceForecastTrend operation
//
// # See also
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

	// Optional list of Exadata Insight VM cluster name.
	VmclusterName []string `contributesTo:"query" name:"vmclusterName" collectionFormat:"multi"`

	// Percent value in which a resource metric is considered highly utilized.
	HighUtilizationThreshold *int `mandatory:"false" contributesTo:"query" name:"highUtilizationThreshold"`

	// Percent value in which a resource metric is considered low utilized.
	LowUtilizationThreshold *int `mandatory:"false" contributesTo:"query" name:"lowUtilizationThreshold"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeDatabaseInsightResourceForecastTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeDatabaseInsightResourceForecastTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeDatabaseInsightResourceForecastTrendRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.DatabaseType {
		if _, ok := GetMappingSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", val, strings.Join(GetSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingSummarizeDatabaseInsightResourceForecastTrendStatisticEnum(string(request.Statistic)); !ok && request.Statistic != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Statistic: %s. Supported values are: %s.", request.Statistic, strings.Join(GetSummarizeDatabaseInsightResourceForecastTrendStatisticEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceForecastTrendForecastModelEnum(string(request.ForecastModel)); !ok && request.ForecastModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ForecastModel: %s. Supported values are: %s.", request.ForecastModel, strings.Join(GetSummarizeDatabaseInsightResourceForecastTrendForecastModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum(string(request.UtilizationLevel)); !ok && request.UtilizationLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UtilizationLevel: %s. Supported values are: %s.", request.UtilizationLevel, strings.Join(GetSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwS                 SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "ADW-S"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpS                 SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "ATP-S"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwD                 SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "ADW-D"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpD                 SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "ATP-D"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalPdb          SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "EXTERNAL-PDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalNoncdb       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "EXTERNAL-NONCDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedVmCdb       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "COMANAGED-VM-CDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedVmPdb       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "COMANAGED-VM-PDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedVmNoncdb    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "COMANAGED-VM-NONCDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedBmCdb       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "COMANAGED-BM-CDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedBmPdb       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "COMANAGED-BM-PDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedBmNoncdb    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "COMANAGED-BM-NONCDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedExacsCdb    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "COMANAGED-EXACS-CDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedExacsPdb    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "COMANAGED-EXACS-PDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedExacsNoncdb SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "COMANAGED-EXACS-NONCDB"
	SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeMdsMysql             SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = "MDS-MYSQL"
)

var mappingSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum = map[string]SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum{
	"ADW-S":                  SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwS,
	"ATP-S":                  SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpS,
	"ADW-D":                  SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwD,
	"ATP-D":                  SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpD,
	"EXTERNAL-PDB":           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalPdb,
	"EXTERNAL-NONCDB":        SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalNoncdb,
	"COMANAGED-VM-CDB":       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedVmCdb,
	"COMANAGED-VM-PDB":       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedVmPdb,
	"COMANAGED-VM-NONCDB":    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedVmNoncdb,
	"COMANAGED-BM-CDB":       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedBmCdb,
	"COMANAGED-BM-PDB":       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedBmPdb,
	"COMANAGED-BM-NONCDB":    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedBmNoncdb,
	"COMANAGED-EXACS-CDB":    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedExacsCdb,
	"COMANAGED-EXACS-PDB":    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedExacsPdb,
	"COMANAGED-EXACS-NONCDB": SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedExacsNoncdb,
	"MDS-MYSQL":              SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeMdsMysql,
}

var mappingSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnumLowerCase = map[string]SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum{
	"adw-s":                  SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwS,
	"atp-s":                  SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpS,
	"adw-d":                  SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAdwD,
	"atp-d":                  SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeAtpD,
	"external-pdb":           SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalPdb,
	"external-noncdb":        SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeExternalNoncdb,
	"comanaged-vm-cdb":       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedVmCdb,
	"comanaged-vm-pdb":       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedVmPdb,
	"comanaged-vm-noncdb":    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedVmNoncdb,
	"comanaged-bm-cdb":       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedBmCdb,
	"comanaged-bm-pdb":       SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedBmPdb,
	"comanaged-bm-noncdb":    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedBmNoncdb,
	"comanaged-exacs-cdb":    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedExacsCdb,
	"comanaged-exacs-pdb":    SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedExacsPdb,
	"comanaged-exacs-noncdb": SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeComanagedExacsNoncdb,
	"mds-mysql":              SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeMdsMysql,
}

// GetSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnumValues() []SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum
func GetSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnumStringValues() []string {
	return []string{
		"ADW-S",
		"ATP-S",
		"ADW-D",
		"ATP-D",
		"EXTERNAL-PDB",
		"EXTERNAL-NONCDB",
		"COMANAGED-VM-CDB",
		"COMANAGED-VM-PDB",
		"COMANAGED-VM-NONCDB",
		"COMANAGED-BM-CDB",
		"COMANAGED-BM-PDB",
		"COMANAGED-BM-NONCDB",
		"COMANAGED-EXACS-CDB",
		"COMANAGED-EXACS-PDB",
		"COMANAGED-EXACS-NONCDB",
		"MDS-MYSQL",
	}
}

// GetMappingSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum(val string) (SummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceForecastTrendDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceForecastTrendStatisticEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceForecastTrendStatisticEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceForecastTrendStatisticEnum
const (
	SummarizeDatabaseInsightResourceForecastTrendStatisticAvg SummarizeDatabaseInsightResourceForecastTrendStatisticEnum = "AVG"
	SummarizeDatabaseInsightResourceForecastTrendStatisticMax SummarizeDatabaseInsightResourceForecastTrendStatisticEnum = "MAX"
)

var mappingSummarizeDatabaseInsightResourceForecastTrendStatisticEnum = map[string]SummarizeDatabaseInsightResourceForecastTrendStatisticEnum{
	"AVG": SummarizeDatabaseInsightResourceForecastTrendStatisticAvg,
	"MAX": SummarizeDatabaseInsightResourceForecastTrendStatisticMax,
}

var mappingSummarizeDatabaseInsightResourceForecastTrendStatisticEnumLowerCase = map[string]SummarizeDatabaseInsightResourceForecastTrendStatisticEnum{
	"avg": SummarizeDatabaseInsightResourceForecastTrendStatisticAvg,
	"max": SummarizeDatabaseInsightResourceForecastTrendStatisticMax,
}

// GetSummarizeDatabaseInsightResourceForecastTrendStatisticEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendStatisticEnum
func GetSummarizeDatabaseInsightResourceForecastTrendStatisticEnumValues() []SummarizeDatabaseInsightResourceForecastTrendStatisticEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendStatisticEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendStatisticEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceForecastTrendStatisticEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceForecastTrendStatisticEnum
func GetSummarizeDatabaseInsightResourceForecastTrendStatisticEnumStringValues() []string {
	return []string{
		"AVG",
		"MAX",
	}
}

// GetMappingSummarizeDatabaseInsightResourceForecastTrendStatisticEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceForecastTrendStatisticEnum(val string) (SummarizeDatabaseInsightResourceForecastTrendStatisticEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceForecastTrendStatisticEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum Enum with underlying type: string
type SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum string

// Set of constants representing the allowable values for SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum
const (
	SummarizeDatabaseInsightResourceForecastTrendForecastModelLinear   SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum = "LINEAR"
	SummarizeDatabaseInsightResourceForecastTrendForecastModelMlAuto   SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum = "ML_AUTO"
	SummarizeDatabaseInsightResourceForecastTrendForecastModelMlNoAuto SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum = "ML_NO_AUTO"
)

var mappingSummarizeDatabaseInsightResourceForecastTrendForecastModelEnum = map[string]SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum{
	"LINEAR":     SummarizeDatabaseInsightResourceForecastTrendForecastModelLinear,
	"ML_AUTO":    SummarizeDatabaseInsightResourceForecastTrendForecastModelMlAuto,
	"ML_NO_AUTO": SummarizeDatabaseInsightResourceForecastTrendForecastModelMlNoAuto,
}

var mappingSummarizeDatabaseInsightResourceForecastTrendForecastModelEnumLowerCase = map[string]SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum{
	"linear":     SummarizeDatabaseInsightResourceForecastTrendForecastModelLinear,
	"ml_auto":    SummarizeDatabaseInsightResourceForecastTrendForecastModelMlAuto,
	"ml_no_auto": SummarizeDatabaseInsightResourceForecastTrendForecastModelMlNoAuto,
}

// GetSummarizeDatabaseInsightResourceForecastTrendForecastModelEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum
func GetSummarizeDatabaseInsightResourceForecastTrendForecastModelEnumValues() []SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendForecastModelEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceForecastTrendForecastModelEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum
func GetSummarizeDatabaseInsightResourceForecastTrendForecastModelEnumStringValues() []string {
	return []string{
		"LINEAR",
		"ML_AUTO",
		"ML_NO_AUTO",
	}
}

// GetMappingSummarizeDatabaseInsightResourceForecastTrendForecastModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceForecastTrendForecastModelEnum(val string) (SummarizeDatabaseInsightResourceForecastTrendForecastModelEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceForecastTrendForecastModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum = map[string]SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum{
	"HIGH_UTILIZATION":        SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelHighUtilization,
	"LOW_UTILIZATION":         SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelLowUtilization,
	"MEDIUM_HIGH_UTILIZATION": SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelMediumHighUtilization,
	"MEDIUM_LOW_UTILIZATION":  SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelMediumLowUtilization,
}

var mappingSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnumLowerCase = map[string]SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum{
	"high_utilization":        SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelHighUtilization,
	"low_utilization":         SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelLowUtilization,
	"medium_high_utilization": SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelMediumHighUtilization,
	"medium_low_utilization":  SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelMediumLowUtilization,
}

// GetSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnumValues Enumerates the set of values for SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum
func GetSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnumValues() []SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum {
	values := make([]SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum, 0)
	for _, v := range mappingSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnumStringValues Enumerates the set of values in String for SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum
func GetSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnumStringValues() []string {
	return []string{
		"HIGH_UTILIZATION",
		"LOW_UTILIZATION",
		"MEDIUM_HIGH_UTILIZATION",
		"MEDIUM_LOW_UTILIZATION",
	}
}

// GetMappingSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum(val string) (SummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnum, bool) {
	enum, ok := mappingSummarizeDatabaseInsightResourceForecastTrendUtilizationLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
