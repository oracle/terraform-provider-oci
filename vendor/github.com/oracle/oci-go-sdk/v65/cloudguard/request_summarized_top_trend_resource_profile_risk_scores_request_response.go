// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// RequestSummarizedTopTrendResourceProfileRiskScoresRequest wrapper for the RequestSummarizedTopTrendResourceProfileRiskScores operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedTopTrendResourceProfileRiskScores.go.html to see an example of how to use RequestSummarizedTopTrendResourceProfileRiskScoresRequest.
type RequestSummarizedTopTrendResourceProfileRiskScoresRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Start time for a filter. If start time is not specified, start time will be set to today's current time - 30 days.
	TimeScoreComputedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScoreComputedGreaterThanOrEqualTo"`

	// End time for a filter. If end time is not specified, end time will be set to today's current time.
	TimeScoreComputedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeScoreComputedLessThanOrEqualTo"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The field to sort trendlines for resource profiles. Only one sort order may be provided. If no value is specified riskScore is default.
	SortBy RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Number of resource profile risk score trend-lines to be displayed. Default value is 10.
	Count *int `mandatory:"false" contributesTo:"query" name:"count"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RequestSummarizedTopTrendResourceProfileRiskScoresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RequestSummarizedTopTrendResourceProfileRiskScoresRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RequestSummarizedTopTrendResourceProfileRiskScoresRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RequestSummarizedTopTrendResourceProfileRiskScoresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request RequestSummarizedTopTrendResourceProfileRiskScoresRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestSummarizedTopTrendResourceProfileRiskScoresResponse wrapper for the RequestSummarizedTopTrendResourceProfileRiskScores operation
type RequestSummarizedTopTrendResourceProfileRiskScoresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceProfileRiskScoreAggregationSummaryCollection instances
	ResourceProfileRiskScoreAggregationSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response RequestSummarizedTopTrendResourceProfileRiskScoresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RequestSummarizedTopTrendResourceProfileRiskScoresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum Enum with underlying type: string
type RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum string

// Set of constants representing the allowable values for RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum
const (
	RequestSummarizedTopTrendResourceProfileRiskScoresSortByRiskscore         RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum = "riskScore"
	RequestSummarizedTopTrendResourceProfileRiskScoresSortByRiskscoregrowth   RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum = "riskScoreGrowth"
	RequestSummarizedTopTrendResourceProfileRiskScoresSortByTimefirstdetected RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum = "timeFirstDetected"
	RequestSummarizedTopTrendResourceProfileRiskScoresSortByTimelastdetected  RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum = "timeLastDetected"
)

var mappingRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum = map[string]RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum{
	"riskScore":         RequestSummarizedTopTrendResourceProfileRiskScoresSortByRiskscore,
	"riskScoreGrowth":   RequestSummarizedTopTrendResourceProfileRiskScoresSortByRiskscoregrowth,
	"timeFirstDetected": RequestSummarizedTopTrendResourceProfileRiskScoresSortByTimefirstdetected,
	"timeLastDetected":  RequestSummarizedTopTrendResourceProfileRiskScoresSortByTimelastdetected,
}

var mappingRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnumLowerCase = map[string]RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum{
	"riskscore":         RequestSummarizedTopTrendResourceProfileRiskScoresSortByRiskscore,
	"riskscoregrowth":   RequestSummarizedTopTrendResourceProfileRiskScoresSortByRiskscoregrowth,
	"timefirstdetected": RequestSummarizedTopTrendResourceProfileRiskScoresSortByTimefirstdetected,
	"timelastdetected":  RequestSummarizedTopTrendResourceProfileRiskScoresSortByTimelastdetected,
}

// GetRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnumValues Enumerates the set of values for RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum
func GetRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnumValues() []RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum {
	values := make([]RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum, 0)
	for _, v := range mappingRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnumStringValues Enumerates the set of values in String for RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum
func GetRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnumStringValues() []string {
	return []string{
		"riskScore",
		"riskScoreGrowth",
		"timeFirstDetected",
		"timeLastDetected",
	}
}

// GetMappingRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum(val string) (RequestSummarizedTopTrendResourceProfileRiskScoresSortByEnum, bool) {
	enum, ok := mappingRequestSummarizedTopTrendResourceProfileRiskScoresSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum Enum with underlying type: string
type RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum string

// Set of constants representing the allowable values for RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum
const (
	RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelRestricted RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum = "RESTRICTED"
	RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelAccessible RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum = "ACCESSIBLE"
)

var mappingRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum = map[string]RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum{
	"RESTRICTED": RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelRestricted,
	"ACCESSIBLE": RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelAccessible,
}

var mappingRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnumLowerCase = map[string]RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum{
	"restricted": RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelRestricted,
	"accessible": RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelAccessible,
}

// GetRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnumValues Enumerates the set of values for RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum
func GetRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnumValues() []RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum {
	values := make([]RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum, 0)
	for _, v := range mappingRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnumStringValues Enumerates the set of values in String for RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum
func GetRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum(val string) (RequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnum, bool) {
	enum, ok := mappingRequestSummarizedTopTrendResourceProfileRiskScoresAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
