// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// SummarizeExadataInsightResourceForecastTrendAggregatedRequest wrapper for the SummarizeExadataInsightResourceForecastTrendAggregated operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceForecastTrendAggregated.go.html to see an example of how to use SummarizeExadataInsightResourceForecastTrendAggregatedRequest.
type SummarizeExadataInsightResourceForecastTrendAggregatedRequest struct {

	// Filter by resource.
	// Supported values are HOST , STORAGE_SERVER and DATABASE
	ResourceType *string `mandatory:"true" contributesTo:"query" name:"resourceType"`

	// Filter by resource metric.
	// Supported values are CPU , STORAGE, MEMORY, IO, IOPS, THROUGHPUT
	ResourceMetric *string `mandatory:"true" contributesTo:"query" name:"resourceMetric"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

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

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more Exadata types.
	// Possible value are DBMACHINE, EXACS, and EXACC.
	ExadataType []string `contributesTo:"query" name:"exadataType" collectionFormat:"multi"`

	// Choose the type of statistic metric data to be used for forecasting.
	Statistic SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum `mandatory:"false" contributesTo:"query" name:"statistic" omitEmpty:"true"`

	// Number of days used for utilization forecast analysis.
	ForecastStartDay *int `mandatory:"false" contributesTo:"query" name:"forecastStartDay"`

	// Number of days used for utilization forecast analysis.
	ForecastDays *int `mandatory:"false" contributesTo:"query" name:"forecastDays"`

	// Choose algorithm model for the forecasting.
	// Possible values:
	//   - LINEAR: Uses linear regression algorithm for forecasting.
	//   - ML_AUTO: Automatically detects best algorithm to use for forecasting.
	//   - ML_NO_AUTO: Automatically detects seasonality of the data for forecasting using linear or seasonal algorithm.
	ForecastModel SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum `mandatory:"false" contributesTo:"query" name:"forecastModel" omitEmpty:"true"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

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

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeExadataInsightResourceForecastTrendAggregatedRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeExadataInsightResourceForecastTrendAggregatedRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeExadataInsightResourceForecastTrendAggregatedRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeExadataInsightResourceForecastTrendAggregatedRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeExadataInsightResourceForecastTrendAggregatedRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum(string(request.Statistic)); !ok && request.Statistic != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Statistic: %s. Supported values are: %s.", request.Statistic, strings.Join(GetSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum(string(request.ForecastModel)); !ok && request.ForecastModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ForecastModel: %s. Supported values are: %s.", request.ForecastModel, strings.Join(GetSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceForecastTrendAggregatedResponse wrapper for the SummarizeExadataInsightResourceForecastTrendAggregated operation
type SummarizeExadataInsightResourceForecastTrendAggregatedResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeExadataInsightResourceForecastTrendAggregation instances
	SummarizeExadataInsightResourceForecastTrendAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeExadataInsightResourceForecastTrendAggregatedResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeExadataInsightResourceForecastTrendAggregatedResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum
const (
	SummarizeExadataInsightResourceForecastTrendAggregatedStatisticAvg SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum = "AVG"
	SummarizeExadataInsightResourceForecastTrendAggregatedStatisticMax SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum = "MAX"
)

var mappingSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum = map[string]SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum{
	"AVG": SummarizeExadataInsightResourceForecastTrendAggregatedStatisticAvg,
	"MAX": SummarizeExadataInsightResourceForecastTrendAggregatedStatisticMax,
}

// GetSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnumValues() []SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnumStringValues() []string {
	return []string{
		"AVG",
		"MAX",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum(val string) (SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum, bool) {
	mappingSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum)
	for k, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnum {
		mappingSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendAggregatedStatisticEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum
const (
	SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelLinear   SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum = "LINEAR"
	SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelMlAuto   SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum = "ML_AUTO"
	SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelMlNoAuto SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum = "ML_NO_AUTO"
)

var mappingSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum = map[string]SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum{
	"LINEAR":     SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelLinear,
	"ML_AUTO":    SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelMlAuto,
	"ML_NO_AUTO": SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelMlNoAuto,
}

// GetSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnumValues() []SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum
func GetSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnumStringValues() []string {
	return []string{
		"LINEAR",
		"ML_AUTO",
		"ML_NO_AUTO",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum(val string) (SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum, bool) {
	mappingSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum)
	for k, v := range mappingSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnum {
		mappingSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendAggregatedForecastModelEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
