// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v43/common"
	"net/http"
)

// SummarizeHostInsightResourceCapacityTrendRequest wrapper for the SummarizeHostInsightResourceCapacityTrend operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeHostInsightResourceCapacityTrendRequest.
type SummarizeHostInsightResourceCapacityTrendRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by host resource metric.
	// Supported values are CPU, MEMORY, and LOGICAL_MEMORY.
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
	// Possible value is LINUX.
	PlatformType []SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// Filter by utilization level by the following buckets:
	//   - HIGH_UTILIZATION: DBs with utilization greater or equal than 75.
	//   - LOW_UTILIZATION: DBs with utilization lower than 25.
	//   - MEDIUM_HIGH_UTILIZATION: DBs with utilization greater or equal than 50 but lower than 75.
	//   - MEDIUM_LOW_UTILIZATION: DBs with utilization greater or equal than 25 but lower than 50.
	UtilizationLevel SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum `mandatory:"false" contributesTo:"query" name:"utilizationLevel" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeHostInsightResourceCapacityTrendSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sorts using end timestamp or capacity
	SortBy SummarizeHostInsightResourceCapacityTrendSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeHostInsightResourceCapacityTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightResourceCapacityTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightResourceCapacityTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightResourceCapacityTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeHostInsightResourceCapacityTrendResponse wrapper for the SummarizeHostInsightResourceCapacityTrend operation
type SummarizeHostInsightResourceCapacityTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeHostInsightResourceCapacityTrendAggregationCollection instances
	SummarizeHostInsightResourceCapacityTrendAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeHostInsightResourceCapacityTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightResourceCapacityTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum
const (
	SummarizeHostInsightResourceCapacityTrendPlatformTypeLinux SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum = "LINUX"
)

var mappingSummarizeHostInsightResourceCapacityTrendPlatformType = map[string]SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum{
	"LINUX": SummarizeHostInsightResourceCapacityTrendPlatformTypeLinux,
}

// GetSummarizeHostInsightResourceCapacityTrendPlatformTypeEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum
func GetSummarizeHostInsightResourceCapacityTrendPlatformTypeEnumValues() []SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendPlatformTypeEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendPlatformType {
		values = append(values, v)
	}
	return values
}

// SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum
const (
	SummarizeHostInsightResourceCapacityTrendUtilizationLevelHighUtilization       SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum = "HIGH_UTILIZATION"
	SummarizeHostInsightResourceCapacityTrendUtilizationLevelLowUtilization        SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum = "LOW_UTILIZATION"
	SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum = "MEDIUM_HIGH_UTILIZATION"
	SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization  SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum = "MEDIUM_LOW_UTILIZATION"
)

var mappingSummarizeHostInsightResourceCapacityTrendUtilizationLevel = map[string]SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum{
	"HIGH_UTILIZATION":        SummarizeHostInsightResourceCapacityTrendUtilizationLevelHighUtilization,
	"LOW_UTILIZATION":         SummarizeHostInsightResourceCapacityTrendUtilizationLevelLowUtilization,
	"MEDIUM_HIGH_UTILIZATION": SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumHighUtilization,
	"MEDIUM_LOW_UTILIZATION":  SummarizeHostInsightResourceCapacityTrendUtilizationLevelMediumLowUtilization,
}

// GetSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum
func GetSummarizeHostInsightResourceCapacityTrendUtilizationLevelEnumValues() []SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendUtilizationLevelEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendUtilizationLevel {
		values = append(values, v)
	}
	return values
}

// SummarizeHostInsightResourceCapacityTrendSortOrderEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendSortOrderEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendSortOrderEnum
const (
	SummarizeHostInsightResourceCapacityTrendSortOrderAsc  SummarizeHostInsightResourceCapacityTrendSortOrderEnum = "ASC"
	SummarizeHostInsightResourceCapacityTrendSortOrderDesc SummarizeHostInsightResourceCapacityTrendSortOrderEnum = "DESC"
)

var mappingSummarizeHostInsightResourceCapacityTrendSortOrder = map[string]SummarizeHostInsightResourceCapacityTrendSortOrderEnum{
	"ASC":  SummarizeHostInsightResourceCapacityTrendSortOrderAsc,
	"DESC": SummarizeHostInsightResourceCapacityTrendSortOrderDesc,
}

// GetSummarizeHostInsightResourceCapacityTrendSortOrderEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendSortOrderEnum
func GetSummarizeHostInsightResourceCapacityTrendSortOrderEnumValues() []SummarizeHostInsightResourceCapacityTrendSortOrderEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendSortOrderEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeHostInsightResourceCapacityTrendSortByEnum Enum with underlying type: string
type SummarizeHostInsightResourceCapacityTrendSortByEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceCapacityTrendSortByEnum
const (
	SummarizeHostInsightResourceCapacityTrendSortByEndtimestamp SummarizeHostInsightResourceCapacityTrendSortByEnum = "endTimestamp"
	SummarizeHostInsightResourceCapacityTrendSortByCapacity     SummarizeHostInsightResourceCapacityTrendSortByEnum = "capacity"
)

var mappingSummarizeHostInsightResourceCapacityTrendSortBy = map[string]SummarizeHostInsightResourceCapacityTrendSortByEnum{
	"endTimestamp": SummarizeHostInsightResourceCapacityTrendSortByEndtimestamp,
	"capacity":     SummarizeHostInsightResourceCapacityTrendSortByCapacity,
}

// GetSummarizeHostInsightResourceCapacityTrendSortByEnumValues Enumerates the set of values for SummarizeHostInsightResourceCapacityTrendSortByEnum
func GetSummarizeHostInsightResourceCapacityTrendSortByEnumValues() []SummarizeHostInsightResourceCapacityTrendSortByEnum {
	values := make([]SummarizeHostInsightResourceCapacityTrendSortByEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceCapacityTrendSortBy {
		values = append(values, v)
	}
	return values
}
