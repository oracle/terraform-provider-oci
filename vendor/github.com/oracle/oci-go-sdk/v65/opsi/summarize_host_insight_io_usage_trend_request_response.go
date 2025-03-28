// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeHostInsightIoUsageTrendRequest wrapper for the SummarizeHostInsightIoUsageTrend operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeHostInsightIoUsageTrend.go.html to see an example of how to use SummarizeHostInsightIoUsageTrendRequest.
type SummarizeHostInsightIoUsageTrendRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Required OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
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

	// Optional OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host (Compute Id)
	HostId *string `mandatory:"false" contributesTo:"query" name:"hostId"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Choose the type of statistic metric data to be used for forecasting.
	Statistic SummarizeHostInsightIoUsageTrendStatisticEnum `mandatory:"false" contributesTo:"query" name:"statistic" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Resource Status
	Status []ResourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeHostInsightIoUsageTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeHostInsightIoUsageTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeHostInsightIoUsageTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeHostInsightIoUsageTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeHostInsightIoUsageTrendRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeHostInsightIoUsageTrendStatisticEnum(string(request.Statistic)); !ok && request.Statistic != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Statistic: %s. Supported values are: %s.", request.Statistic, strings.Join(GetSummarizeHostInsightIoUsageTrendStatisticEnumStringValues(), ",")))
	}
	for _, val := range request.Status {
		if _, ok := GetMappingResourceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetResourceStatusEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeHostInsightIoUsageTrendResponse wrapper for the SummarizeHostInsightIoUsageTrend operation
type SummarizeHostInsightIoUsageTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeHostInsightIoUsageTrendAggregationCollection instances
	SummarizeHostInsightIoUsageTrendAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeHostInsightIoUsageTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeHostInsightIoUsageTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeHostInsightIoUsageTrendStatisticEnum Enum with underlying type: string
type SummarizeHostInsightIoUsageTrendStatisticEnum string

// Set of constants representing the allowable values for SummarizeHostInsightIoUsageTrendStatisticEnum
const (
	SummarizeHostInsightIoUsageTrendStatisticAvg SummarizeHostInsightIoUsageTrendStatisticEnum = "AVG"
	SummarizeHostInsightIoUsageTrendStatisticMax SummarizeHostInsightIoUsageTrendStatisticEnum = "MAX"
)

var mappingSummarizeHostInsightIoUsageTrendStatisticEnum = map[string]SummarizeHostInsightIoUsageTrendStatisticEnum{
	"AVG": SummarizeHostInsightIoUsageTrendStatisticAvg,
	"MAX": SummarizeHostInsightIoUsageTrendStatisticMax,
}

var mappingSummarizeHostInsightIoUsageTrendStatisticEnumLowerCase = map[string]SummarizeHostInsightIoUsageTrendStatisticEnum{
	"avg": SummarizeHostInsightIoUsageTrendStatisticAvg,
	"max": SummarizeHostInsightIoUsageTrendStatisticMax,
}

// GetSummarizeHostInsightIoUsageTrendStatisticEnumValues Enumerates the set of values for SummarizeHostInsightIoUsageTrendStatisticEnum
func GetSummarizeHostInsightIoUsageTrendStatisticEnumValues() []SummarizeHostInsightIoUsageTrendStatisticEnum {
	values := make([]SummarizeHostInsightIoUsageTrendStatisticEnum, 0)
	for _, v := range mappingSummarizeHostInsightIoUsageTrendStatisticEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeHostInsightIoUsageTrendStatisticEnumStringValues Enumerates the set of values in String for SummarizeHostInsightIoUsageTrendStatisticEnum
func GetSummarizeHostInsightIoUsageTrendStatisticEnumStringValues() []string {
	return []string{
		"AVG",
		"MAX",
	}
}

// GetMappingSummarizeHostInsightIoUsageTrendStatisticEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeHostInsightIoUsageTrendStatisticEnum(val string) (SummarizeHostInsightIoUsageTrendStatisticEnum, bool) {
	enum, ok := mappingSummarizeHostInsightIoUsageTrendStatisticEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
