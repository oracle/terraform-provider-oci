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

// SummarizeHostInsightTopProcessesUsageTrendRequest wrapper for the SummarizeHostInsightTopProcessesUsageTrend operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightTopProcessesUsageTrend.go.html to see an example of how to use SummarizeHostInsightTopProcessesUsageTrendRequest.
type SummarizeHostInsightTopProcessesUsageTrendRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Required OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
	Id *string `mandatory:"true" contributesTo:"query" name:"id"`

	// Host top processes resource metric sort options.
	// Supported values are CPU, MEMORY, VIIRTUAL_MEMORY.
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

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter by one or more host types.
	// Possible values are CLOUD-HOST, EXTERNAL-HOST, COMANAGED-VM-HOST, COMANAGED-BM-HOST, COMANAGED-EXACS-HOST
	HostType []string `contributesTo:"query" name:"hostType" collectionFormat:"multi"`

	// Optional OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host (Compute Id)
	HostId *string `mandatory:"false" contributesTo:"query" name:"hostId"`

	// Unique identifier for a process.
	ProcessHash *string `mandatory:"false" contributesTo:"query" name:"processHash"`

	// Choose the type of statistic metric data to be used for forecasting.
	Statistic SummarizeHostInsightTopProcessesUsageTrendStatisticEnum `mandatory:"false" contributesTo:"query" name:"statistic" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeHostInsightTopProcessesUsageTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightTopProcessesUsageTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightTopProcessesUsageTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightTopProcessesUsageTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeHostInsightTopProcessesUsageTrendRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightTopProcessesUsageTrendStatisticEnum(string(request.Statistic)); !ok && request.Statistic != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Statistic: %s. Supported values are: %s.", request.Statistic, strings.Join(GetSummarizeHostInsightTopProcessesUsageTrendStatisticEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightTopProcessesUsageTrendResponse wrapper for the SummarizeHostInsightTopProcessesUsageTrend operation
type SummarizeHostInsightTopProcessesUsageTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeHostInsightsTopProcessesUsageTrendCollection instances
	SummarizeHostInsightsTopProcessesUsageTrendCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeHostInsightTopProcessesUsageTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightTopProcessesUsageTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightTopProcessesUsageTrendStatisticEnum Enum with underlying type: string
type SummarizeHostInsightTopProcessesUsageTrendStatisticEnum string

// Set of constants representing the allowable values for SummarizeHostInsightTopProcessesUsageTrendStatisticEnum
const (
	SummarizeHostInsightTopProcessesUsageTrendStatisticAvg SummarizeHostInsightTopProcessesUsageTrendStatisticEnum = "AVG"
	SummarizeHostInsightTopProcessesUsageTrendStatisticMax SummarizeHostInsightTopProcessesUsageTrendStatisticEnum = "MAX"
)

var mappingSummarizeHostInsightTopProcessesUsageTrendStatisticEnum = map[string]SummarizeHostInsightTopProcessesUsageTrendStatisticEnum{
	"AVG": SummarizeHostInsightTopProcessesUsageTrendStatisticAvg,
	"MAX": SummarizeHostInsightTopProcessesUsageTrendStatisticMax,
}

var mappingSummarizeHostInsightTopProcessesUsageTrendStatisticEnumLowerCase = map[string]SummarizeHostInsightTopProcessesUsageTrendStatisticEnum{
	"avg": SummarizeHostInsightTopProcessesUsageTrendStatisticAvg,
	"max": SummarizeHostInsightTopProcessesUsageTrendStatisticMax,
}

// GetSummarizeHostInsightTopProcessesUsageTrendStatisticEnumValues Enumerates the set of values for SummarizeHostInsightTopProcessesUsageTrendStatisticEnum
func GetSummarizeHostInsightTopProcessesUsageTrendStatisticEnumValues() []SummarizeHostInsightTopProcessesUsageTrendStatisticEnum {
	values := make([]SummarizeHostInsightTopProcessesUsageTrendStatisticEnum, 0)
	for _, v := range mappingSummarizeHostInsightTopProcessesUsageTrendStatisticEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightTopProcessesUsageTrendStatisticEnumStringValues Enumerates the set of values in String for SummarizeHostInsightTopProcessesUsageTrendStatisticEnum
func GetSummarizeHostInsightTopProcessesUsageTrendStatisticEnumStringValues() []string {
	return []string{
		"AVG",
		"MAX",
	}
}

// GetMappingSummarizeHostInsightTopProcessesUsageTrendStatisticEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightTopProcessesUsageTrendStatisticEnum(val string) (SummarizeHostInsightTopProcessesUsageTrendStatisticEnum, bool) {
	enum, ok := mappingSummarizeHostInsightTopProcessesUsageTrendStatisticEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
