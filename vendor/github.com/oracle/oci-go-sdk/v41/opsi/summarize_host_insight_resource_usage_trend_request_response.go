// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// SummarizeHostInsightResourceUsageTrendRequest wrapper for the SummarizeHostInsightResourceUsageTrend operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightResourceUsageTrend.go.html to see an example of how to use SummarizeHostInsightResourceUsageTrendRequest.
type SummarizeHostInsightResourceUsageTrendRequest struct {

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
	PlatformType []SummarizeHostInsightResourceUsageTrendPlatformTypeEnum `contributesTo:"query" name:"platformType" omitEmpty:"true" collectionFormat:"multi"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeHostInsightResourceUsageTrendSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sorts using end timestamp, usage or capacity
	SortBy SummarizeHostInsightResourceUsageTrendSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeHostInsightResourceUsageTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightResourceUsageTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightResourceUsageTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightResourceUsageTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeHostInsightResourceUsageTrendResponse wrapper for the SummarizeHostInsightResourceUsageTrend operation
type SummarizeHostInsightResourceUsageTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeHostInsightResourceUsageTrendAggregationCollection instances
	SummarizeHostInsightResourceUsageTrendAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeHostInsightResourceUsageTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightResourceUsageTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightResourceUsageTrendPlatformTypeEnum Enum with underlying type: string
type SummarizeHostInsightResourceUsageTrendPlatformTypeEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUsageTrendPlatformTypeEnum
const (
	SummarizeHostInsightResourceUsageTrendPlatformTypeLinux SummarizeHostInsightResourceUsageTrendPlatformTypeEnum = "LINUX"
)

var mappingSummarizeHostInsightResourceUsageTrendPlatformType = map[string]SummarizeHostInsightResourceUsageTrendPlatformTypeEnum{
	"LINUX": SummarizeHostInsightResourceUsageTrendPlatformTypeLinux,
}

// GetSummarizeHostInsightResourceUsageTrendPlatformTypeEnumValues Enumerates the set of values for SummarizeHostInsightResourceUsageTrendPlatformTypeEnum
func GetSummarizeHostInsightResourceUsageTrendPlatformTypeEnumValues() []SummarizeHostInsightResourceUsageTrendPlatformTypeEnum {
	values := make([]SummarizeHostInsightResourceUsageTrendPlatformTypeEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUsageTrendPlatformType {
		values = append(values, v)
	}
	return values
}

// SummarizeHostInsightResourceUsageTrendSortOrderEnum Enum with underlying type: string
type SummarizeHostInsightResourceUsageTrendSortOrderEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUsageTrendSortOrderEnum
const (
	SummarizeHostInsightResourceUsageTrendSortOrderAsc  SummarizeHostInsightResourceUsageTrendSortOrderEnum = "ASC"
	SummarizeHostInsightResourceUsageTrendSortOrderDesc SummarizeHostInsightResourceUsageTrendSortOrderEnum = "DESC"
)

var mappingSummarizeHostInsightResourceUsageTrendSortOrder = map[string]SummarizeHostInsightResourceUsageTrendSortOrderEnum{
	"ASC":  SummarizeHostInsightResourceUsageTrendSortOrderAsc,
	"DESC": SummarizeHostInsightResourceUsageTrendSortOrderDesc,
}

// GetSummarizeHostInsightResourceUsageTrendSortOrderEnumValues Enumerates the set of values for SummarizeHostInsightResourceUsageTrendSortOrderEnum
func GetSummarizeHostInsightResourceUsageTrendSortOrderEnumValues() []SummarizeHostInsightResourceUsageTrendSortOrderEnum {
	values := make([]SummarizeHostInsightResourceUsageTrendSortOrderEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUsageTrendSortOrder {
		values = append(values, v)
	}
	return values
}

// SummarizeHostInsightResourceUsageTrendSortByEnum Enum with underlying type: string
type SummarizeHostInsightResourceUsageTrendSortByEnum string

// Set of constants representing the allowable values for SummarizeHostInsightResourceUsageTrendSortByEnum
const (
	SummarizeHostInsightResourceUsageTrendSortByEndtimestamp SummarizeHostInsightResourceUsageTrendSortByEnum = "endTimestamp"
	SummarizeHostInsightResourceUsageTrendSortByUsage        SummarizeHostInsightResourceUsageTrendSortByEnum = "usage"
	SummarizeHostInsightResourceUsageTrendSortByCapacity     SummarizeHostInsightResourceUsageTrendSortByEnum = "capacity"
)

var mappingSummarizeHostInsightResourceUsageTrendSortBy = map[string]SummarizeHostInsightResourceUsageTrendSortByEnum{
	"endTimestamp": SummarizeHostInsightResourceUsageTrendSortByEndtimestamp,
	"usage":        SummarizeHostInsightResourceUsageTrendSortByUsage,
	"capacity":     SummarizeHostInsightResourceUsageTrendSortByCapacity,
}

// GetSummarizeHostInsightResourceUsageTrendSortByEnumValues Enumerates the set of values for SummarizeHostInsightResourceUsageTrendSortByEnum
func GetSummarizeHostInsightResourceUsageTrendSortByEnumValues() []SummarizeHostInsightResourceUsageTrendSortByEnum {
	values := make([]SummarizeHostInsightResourceUsageTrendSortByEnum, 0)
	for _, v := range mappingSummarizeHostInsightResourceUsageTrendSortBy {
		values = append(values, v)
	}
	return values
}
