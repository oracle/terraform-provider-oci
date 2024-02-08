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

// SummarizeExadataInsightResourceForecastTrendRequest wrapper for the SummarizeExadataInsightResourceForecastTrend operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceForecastTrend.go.html to see an example of how to use SummarizeExadataInsightResourceForecastTrendRequest.
type SummarizeExadataInsightResourceForecastTrendRequest struct {

	// Filter by resource.
	// Supported values are HOST , STORAGE_SERVER and DATABASE
	ResourceType *string `mandatory:"true" contributesTo:"query" name:"resourceType"`

	// Filter by resource metric.
	// Supported values are CPU , STORAGE, MEMORY, IO, IOPS, THROUGHPUT
	ResourceMetric *string `mandatory:"true" contributesTo:"query" name:"resourceMetric"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource.
	ExadataInsightId *string `mandatory:"true" contributesTo:"query" name:"exadataInsightId"`

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

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DatabaseInsightId []string `contributesTo:"query" name:"databaseInsightId" collectionFormat:"multi"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	HostInsightId []string `contributesTo:"query" name:"hostInsightId" collectionFormat:"multi"`

	// Optional storage server name on an exadata system.
	StorageServerName []string `contributesTo:"query" name:"storageServerName" collectionFormat:"multi"`

	// Filter by one or more Exadata types.
	// Possible value are DBMACHINE, EXACS, and EXACC.
	ExadataType []string `contributesTo:"query" name:"exadataType" collectionFormat:"multi"`

	// Choose the type of statistic metric data to be used for forecasting.
	Statistic SummarizeExadataInsightResourceForecastTrendStatisticEnum `mandatory:"false" contributesTo:"query" name:"statistic" omitEmpty:"true"`

	// Number of days used for utilization forecast analysis.
	ForecastStartDay *int `mandatory:"false" contributesTo:"query" name:"forecastStartDay"`

	// Number of days used for utilization forecast analysis.
	ForecastDays *int `mandatory:"false" contributesTo:"query" name:"forecastDays"`

	// Choose algorithm model for the forecasting.
	// Possible values:
	//   - LINEAR: Uses linear regression algorithm for forecasting.
	//   - ML_AUTO: Automatically detects best algorithm to use for forecasting.
	//   - ML_NO_AUTO: Automatically detects seasonality of the data for forecasting using linear or seasonal algorithm.
	ForecastModel SummarizeExadataInsightResourceForecastTrendForecastModelEnum `mandatory:"false" contributesTo:"query" name:"forecastModel" omitEmpty:"true"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// This parameter is used to change data's confidence level, this data is ingested by the
	// forecast algorithm.
	// Confidence is the probability of an interval to contain the expected population parameter.
	// Manipulation of this value will lead to different results.
	// If not set, default confidence value is 95%.
	Confidence *int `mandatory:"false" contributesTo:"query" name:"confidence"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeExadataInsightResourceForecastTrendSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The order in which resource Forecast trend records are listed
	SortBy SummarizeExadataInsightResourceForecastTrendSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeExadataInsightResourceForecastTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeExadataInsightResourceForecastTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeExadataInsightResourceForecastTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeExadataInsightResourceForecastTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeExadataInsightResourceForecastTrendRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendStatisticEnum(string(request.Statistic)); !ok && request.Statistic != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Statistic: %s. Supported values are: %s.", request.Statistic, strings.Join(GetSummarizeExadataInsightResourceForecastTrendStatisticEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendForecastModelEnum(string(request.ForecastModel)); !ok && request.ForecastModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ForecastModel: %s. Supported values are: %s.", request.ForecastModel, strings.Join(GetSummarizeExadataInsightResourceForecastTrendForecastModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeExadataInsightResourceForecastTrendSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceForecastTrendSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeExadataInsightResourceForecastTrendSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceForecastTrendResponse wrapper for the SummarizeExadataInsightResourceForecastTrend operation
type SummarizeExadataInsightResourceForecastTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeExadataInsightResourceForecastTrendCollection instances
	SummarizeExadataInsightResourceForecastTrendCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeExadataInsightResourceForecastTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeExadataInsightResourceForecastTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeExadataInsightResourceForecastTrendStatisticEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendStatisticEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendStatisticEnum
const (
	SummarizeExadataInsightResourceForecastTrendStatisticAvg SummarizeExadataInsightResourceForecastTrendStatisticEnum = "AVG"
	SummarizeExadataInsightResourceForecastTrendStatisticMax SummarizeExadataInsightResourceForecastTrendStatisticEnum = "MAX"
)

var mappingSummarizeExadataInsightResourceForecastTrendStatisticEnum = map[string]SummarizeExadataInsightResourceForecastTrendStatisticEnum{
	"AVG": SummarizeExadataInsightResourceForecastTrendStatisticAvg,
	"MAX": SummarizeExadataInsightResourceForecastTrendStatisticMax,
}

var mappingSummarizeExadataInsightResourceForecastTrendStatisticEnumLowerCase = map[string]SummarizeExadataInsightResourceForecastTrendStatisticEnum{
	"avg": SummarizeExadataInsightResourceForecastTrendStatisticAvg,
	"max": SummarizeExadataInsightResourceForecastTrendStatisticMax,
}

// GetSummarizeExadataInsightResourceForecastTrendStatisticEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendStatisticEnum
func GetSummarizeExadataInsightResourceForecastTrendStatisticEnumValues() []SummarizeExadataInsightResourceForecastTrendStatisticEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendStatisticEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendStatisticEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendStatisticEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendStatisticEnum
func GetSummarizeExadataInsightResourceForecastTrendStatisticEnumStringValues() []string {
	return []string{
		"AVG",
		"MAX",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendStatisticEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendStatisticEnum(val string) (SummarizeExadataInsightResourceForecastTrendStatisticEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendStatisticEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceForecastTrendForecastModelEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendForecastModelEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendForecastModelEnum
const (
	SummarizeExadataInsightResourceForecastTrendForecastModelLinear   SummarizeExadataInsightResourceForecastTrendForecastModelEnum = "LINEAR"
	SummarizeExadataInsightResourceForecastTrendForecastModelMlAuto   SummarizeExadataInsightResourceForecastTrendForecastModelEnum = "ML_AUTO"
	SummarizeExadataInsightResourceForecastTrendForecastModelMlNoAuto SummarizeExadataInsightResourceForecastTrendForecastModelEnum = "ML_NO_AUTO"
)

var mappingSummarizeExadataInsightResourceForecastTrendForecastModelEnum = map[string]SummarizeExadataInsightResourceForecastTrendForecastModelEnum{
	"LINEAR":     SummarizeExadataInsightResourceForecastTrendForecastModelLinear,
	"ML_AUTO":    SummarizeExadataInsightResourceForecastTrendForecastModelMlAuto,
	"ML_NO_AUTO": SummarizeExadataInsightResourceForecastTrendForecastModelMlNoAuto,
}

var mappingSummarizeExadataInsightResourceForecastTrendForecastModelEnumLowerCase = map[string]SummarizeExadataInsightResourceForecastTrendForecastModelEnum{
	"linear":     SummarizeExadataInsightResourceForecastTrendForecastModelLinear,
	"ml_auto":    SummarizeExadataInsightResourceForecastTrendForecastModelMlAuto,
	"ml_no_auto": SummarizeExadataInsightResourceForecastTrendForecastModelMlNoAuto,
}

// GetSummarizeExadataInsightResourceForecastTrendForecastModelEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendForecastModelEnum
func GetSummarizeExadataInsightResourceForecastTrendForecastModelEnumValues() []SummarizeExadataInsightResourceForecastTrendForecastModelEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendForecastModelEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendForecastModelEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendForecastModelEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendForecastModelEnum
func GetSummarizeExadataInsightResourceForecastTrendForecastModelEnumStringValues() []string {
	return []string{
		"LINEAR",
		"ML_AUTO",
		"ML_NO_AUTO",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendForecastModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendForecastModelEnum(val string) (SummarizeExadataInsightResourceForecastTrendForecastModelEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendForecastModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceForecastTrendSortOrderEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendSortOrderEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendSortOrderEnum
const (
	SummarizeExadataInsightResourceForecastTrendSortOrderAsc  SummarizeExadataInsightResourceForecastTrendSortOrderEnum = "ASC"
	SummarizeExadataInsightResourceForecastTrendSortOrderDesc SummarizeExadataInsightResourceForecastTrendSortOrderEnum = "DESC"
)

var mappingSummarizeExadataInsightResourceForecastTrendSortOrderEnum = map[string]SummarizeExadataInsightResourceForecastTrendSortOrderEnum{
	"ASC":  SummarizeExadataInsightResourceForecastTrendSortOrderAsc,
	"DESC": SummarizeExadataInsightResourceForecastTrendSortOrderDesc,
}

var mappingSummarizeExadataInsightResourceForecastTrendSortOrderEnumLowerCase = map[string]SummarizeExadataInsightResourceForecastTrendSortOrderEnum{
	"asc":  SummarizeExadataInsightResourceForecastTrendSortOrderAsc,
	"desc": SummarizeExadataInsightResourceForecastTrendSortOrderDesc,
}

// GetSummarizeExadataInsightResourceForecastTrendSortOrderEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendSortOrderEnum
func GetSummarizeExadataInsightResourceForecastTrendSortOrderEnumValues() []SummarizeExadataInsightResourceForecastTrendSortOrderEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendSortOrderEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendSortOrderEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendSortOrderEnum
func GetSummarizeExadataInsightResourceForecastTrendSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendSortOrderEnum(val string) (SummarizeExadataInsightResourceForecastTrendSortOrderEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceForecastTrendSortByEnum Enum with underlying type: string
type SummarizeExadataInsightResourceForecastTrendSortByEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceForecastTrendSortByEnum
const (
	SummarizeExadataInsightResourceForecastTrendSortById                  SummarizeExadataInsightResourceForecastTrendSortByEnum = "id"
	SummarizeExadataInsightResourceForecastTrendSortByName                SummarizeExadataInsightResourceForecastTrendSortByEnum = "name"
	SummarizeExadataInsightResourceForecastTrendSortByDaystoreachcapacity SummarizeExadataInsightResourceForecastTrendSortByEnum = "daysToReachCapacity"
)

var mappingSummarizeExadataInsightResourceForecastTrendSortByEnum = map[string]SummarizeExadataInsightResourceForecastTrendSortByEnum{
	"id":                  SummarizeExadataInsightResourceForecastTrendSortById,
	"name":                SummarizeExadataInsightResourceForecastTrendSortByName,
	"daysToReachCapacity": SummarizeExadataInsightResourceForecastTrendSortByDaystoreachcapacity,
}

var mappingSummarizeExadataInsightResourceForecastTrendSortByEnumLowerCase = map[string]SummarizeExadataInsightResourceForecastTrendSortByEnum{
	"id":                  SummarizeExadataInsightResourceForecastTrendSortById,
	"name":                SummarizeExadataInsightResourceForecastTrendSortByName,
	"daystoreachcapacity": SummarizeExadataInsightResourceForecastTrendSortByDaystoreachcapacity,
}

// GetSummarizeExadataInsightResourceForecastTrendSortByEnumValues Enumerates the set of values for SummarizeExadataInsightResourceForecastTrendSortByEnum
func GetSummarizeExadataInsightResourceForecastTrendSortByEnumValues() []SummarizeExadataInsightResourceForecastTrendSortByEnum {
	values := make([]SummarizeExadataInsightResourceForecastTrendSortByEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceForecastTrendSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceForecastTrendSortByEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceForecastTrendSortByEnum
func GetSummarizeExadataInsightResourceForecastTrendSortByEnumStringValues() []string {
	return []string{
		"id",
		"name",
		"daysToReachCapacity",
	}
}

// GetMappingSummarizeExadataInsightResourceForecastTrendSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceForecastTrendSortByEnum(val string) (SummarizeExadataInsightResourceForecastTrendSortByEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceForecastTrendSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
