// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// SummarizeExadataInsightResourceStatisticsRequest wrapper for the SummarizeExadataInsightResourceStatistics operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceStatistics.go.html to see an example of how to use SummarizeExadataInsightResourceStatisticsRequest.
type SummarizeExadataInsightResourceStatisticsRequest struct {

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource.
	ExadataInsightId *string `mandatory:"true" contributesTo:"query" name:"exadataInsightId"`

	// Filter by resource.
	// Supported values are HOST , STORAGE_SERVER and DATABASE
	ResourceType *string `mandatory:"true" contributesTo:"query" name:"resourceType"`

	// Filter by resource metric.
	// Supported values are CPU , STORAGE, MEMORY, IO, IOPS, THROUGHPUT
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

	// Filter by one or more Exadata types.
	// Possible value are DBMACHINE, EXACS, and EXACC.
	ExadataType []string `contributesTo:"query" name:"exadataType" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// Percentile values of daily usage to be used for computing the aggregate resource usage.
	Percentile *int `mandatory:"false" contributesTo:"query" name:"percentile"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeExadataInsightResourceStatisticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The order in which resource statistics records are listed
	SortBy SummarizeExadataInsightResourceStatisticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeExadataInsightResourceStatisticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeExadataInsightResourceStatisticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeExadataInsightResourceStatisticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeExadataInsightResourceStatisticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeExadataInsightResourceStatisticsResponse wrapper for the SummarizeExadataInsightResourceStatistics operation
type SummarizeExadataInsightResourceStatisticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeExadataInsightResourceStatisticsAggregationCollection instances
	SummarizeExadataInsightResourceStatisticsAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. The total number of items in the result.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeExadataInsightResourceStatisticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeExadataInsightResourceStatisticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeExadataInsightResourceStatisticsSortOrderEnum Enum with underlying type: string
type SummarizeExadataInsightResourceStatisticsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceStatisticsSortOrderEnum
const (
	SummarizeExadataInsightResourceStatisticsSortOrderAsc  SummarizeExadataInsightResourceStatisticsSortOrderEnum = "ASC"
	SummarizeExadataInsightResourceStatisticsSortOrderDesc SummarizeExadataInsightResourceStatisticsSortOrderEnum = "DESC"
)

var mappingSummarizeExadataInsightResourceStatisticsSortOrder = map[string]SummarizeExadataInsightResourceStatisticsSortOrderEnum{
	"ASC":  SummarizeExadataInsightResourceStatisticsSortOrderAsc,
	"DESC": SummarizeExadataInsightResourceStatisticsSortOrderDesc,
}

// GetSummarizeExadataInsightResourceStatisticsSortOrderEnumValues Enumerates the set of values for SummarizeExadataInsightResourceStatisticsSortOrderEnum
func GetSummarizeExadataInsightResourceStatisticsSortOrderEnumValues() []SummarizeExadataInsightResourceStatisticsSortOrderEnum {
	values := make([]SummarizeExadataInsightResourceStatisticsSortOrderEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceStatisticsSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeExadataInsightResourceStatisticsSortByEnum Enum with underlying type: string
type SummarizeExadataInsightResourceStatisticsSortByEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceStatisticsSortByEnum
const (
	SummarizeExadataInsightResourceStatisticsSortByUtilizationpercent SummarizeExadataInsightResourceStatisticsSortByEnum = "utilizationPercent"
	SummarizeExadataInsightResourceStatisticsSortByUsage              SummarizeExadataInsightResourceStatisticsSortByEnum = "usage"
	SummarizeExadataInsightResourceStatisticsSortByUsagechangepercent SummarizeExadataInsightResourceStatisticsSortByEnum = "usageChangePercent"
)

var mappingSummarizeExadataInsightResourceStatisticsSortBy = map[string]SummarizeExadataInsightResourceStatisticsSortByEnum{
	"utilizationPercent": SummarizeExadataInsightResourceStatisticsSortByUtilizationpercent,
	"usage":              SummarizeExadataInsightResourceStatisticsSortByUsage,
	"usageChangePercent": SummarizeExadataInsightResourceStatisticsSortByUsagechangepercent,
}

// GetSummarizeExadataInsightResourceStatisticsSortByEnumValues Enumerates the set of values for SummarizeExadataInsightResourceStatisticsSortByEnum
func GetSummarizeExadataInsightResourceStatisticsSortByEnumValues() []SummarizeExadataInsightResourceStatisticsSortByEnum {
	values := make([]SummarizeExadataInsightResourceStatisticsSortByEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceStatisticsSortBy {
		values = append(values, v)
	}
	return values
}
