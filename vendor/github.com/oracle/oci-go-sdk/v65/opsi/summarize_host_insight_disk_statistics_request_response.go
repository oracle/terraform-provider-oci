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

// SummarizeHostInsightDiskStatisticsRequest wrapper for the SummarizeHostInsightDiskStatistics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightDiskStatistics.go.html to see an example of how to use SummarizeHostInsightDiskStatisticsRequest.
type SummarizeHostInsightDiskStatisticsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Required OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id *string `mandatory:"true" contributesTo:"query" name:"id"`

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

	// Optional OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host (Compute Id)
	HostId *string `mandatory:"false" contributesTo:"query" name:"hostId"`

	// Choose the type of statistic metric data to be used for forecasting.
	Statistic SummarizeHostInsightDiskStatisticsStatisticEnum `mandatory:"false" contributesTo:"query" name:"statistic" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeHostInsightDiskStatisticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightDiskStatisticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightDiskStatisticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightDiskStatisticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeHostInsightDiskStatisticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightDiskStatisticsStatisticEnum(string(request.Statistic)); !ok && request.Statistic != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Statistic: %s. Supported values are: %s.", request.Statistic, strings.Join(GetSummarizeHostInsightDiskStatisticsStatisticEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightDiskStatisticsResponse wrapper for the SummarizeHostInsightDiskStatistics operation
type SummarizeHostInsightDiskStatisticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SummarizeHostInsightsDiskStatisticsCollection instance
	SummarizeHostInsightsDiskStatisticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeHostInsightDiskStatisticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightDiskStatisticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightDiskStatisticsStatisticEnum Enum with underlying type: string
type SummarizeHostInsightDiskStatisticsStatisticEnum string

// Set of constants representing the allowable values for SummarizeHostInsightDiskStatisticsStatisticEnum
const (
	SummarizeHostInsightDiskStatisticsStatisticAvg SummarizeHostInsightDiskStatisticsStatisticEnum = "AVG"
	SummarizeHostInsightDiskStatisticsStatisticMax SummarizeHostInsightDiskStatisticsStatisticEnum = "MAX"
)

var mappingSummarizeHostInsightDiskStatisticsStatisticEnum = map[string]SummarizeHostInsightDiskStatisticsStatisticEnum{
	"AVG": SummarizeHostInsightDiskStatisticsStatisticAvg,
	"MAX": SummarizeHostInsightDiskStatisticsStatisticMax,
}

var mappingSummarizeHostInsightDiskStatisticsStatisticEnumLowerCase = map[string]SummarizeHostInsightDiskStatisticsStatisticEnum{
	"avg": SummarizeHostInsightDiskStatisticsStatisticAvg,
	"max": SummarizeHostInsightDiskStatisticsStatisticMax,
}

// GetSummarizeHostInsightDiskStatisticsStatisticEnumValues Enumerates the set of values for SummarizeHostInsightDiskStatisticsStatisticEnum
func GetSummarizeHostInsightDiskStatisticsStatisticEnumValues() []SummarizeHostInsightDiskStatisticsStatisticEnum {
	values := make([]SummarizeHostInsightDiskStatisticsStatisticEnum, 0)
	for _, v := range mappingSummarizeHostInsightDiskStatisticsStatisticEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightDiskStatisticsStatisticEnumStringValues Enumerates the set of values in String for SummarizeHostInsightDiskStatisticsStatisticEnum
func GetSummarizeHostInsightDiskStatisticsStatisticEnumStringValues() []string {
	return []string{
		"AVG",
		"MAX",
	}
}

// GetMappingSummarizeHostInsightDiskStatisticsStatisticEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightDiskStatisticsStatisticEnum(val string) (SummarizeHostInsightDiskStatisticsStatisticEnum, bool) {
	enum, ok := mappingSummarizeHostInsightDiskStatisticsStatisticEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
