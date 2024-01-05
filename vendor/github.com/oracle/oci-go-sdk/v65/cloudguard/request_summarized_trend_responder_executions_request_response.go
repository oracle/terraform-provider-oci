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

// RequestSummarizedTrendResponderExecutionsRequest wrapper for the RequestSummarizedTrendResponderExecutions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedTrendResponderExecutions.go.html to see an example of how to use RequestSummarizedTrendResponderExecutionsRequest.
type RequestSummarizedTrendResponderExecutionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Completion End Time
	TimeCompletedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCompletedGreaterThanOrEqualTo"`

	// Completion Start Time
	TimeCompletedLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCompletedLessThanOrEqualTo"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel RequestSummarizedTrendResponderExecutionsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

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

func (request RequestSummarizedTrendResponderExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RequestSummarizedTrendResponderExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RequestSummarizedTrendResponderExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RequestSummarizedTrendResponderExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request RequestSummarizedTrendResponderExecutionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRequestSummarizedTrendResponderExecutionsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetRequestSummarizedTrendResponderExecutionsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestSummarizedTrendResponderExecutionsResponse wrapper for the RequestSummarizedTrendResponderExecutions operation
type RequestSummarizedTrendResponderExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResponderExecutionTrendAggregationCollection instances
	ResponderExecutionTrendAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response RequestSummarizedTrendResponderExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RequestSummarizedTrendResponderExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RequestSummarizedTrendResponderExecutionsAccessLevelEnum Enum with underlying type: string
type RequestSummarizedTrendResponderExecutionsAccessLevelEnum string

// Set of constants representing the allowable values for RequestSummarizedTrendResponderExecutionsAccessLevelEnum
const (
	RequestSummarizedTrendResponderExecutionsAccessLevelRestricted RequestSummarizedTrendResponderExecutionsAccessLevelEnum = "RESTRICTED"
	RequestSummarizedTrendResponderExecutionsAccessLevelAccessible RequestSummarizedTrendResponderExecutionsAccessLevelEnum = "ACCESSIBLE"
)

var mappingRequestSummarizedTrendResponderExecutionsAccessLevelEnum = map[string]RequestSummarizedTrendResponderExecutionsAccessLevelEnum{
	"RESTRICTED": RequestSummarizedTrendResponderExecutionsAccessLevelRestricted,
	"ACCESSIBLE": RequestSummarizedTrendResponderExecutionsAccessLevelAccessible,
}

var mappingRequestSummarizedTrendResponderExecutionsAccessLevelEnumLowerCase = map[string]RequestSummarizedTrendResponderExecutionsAccessLevelEnum{
	"restricted": RequestSummarizedTrendResponderExecutionsAccessLevelRestricted,
	"accessible": RequestSummarizedTrendResponderExecutionsAccessLevelAccessible,
}

// GetRequestSummarizedTrendResponderExecutionsAccessLevelEnumValues Enumerates the set of values for RequestSummarizedTrendResponderExecutionsAccessLevelEnum
func GetRequestSummarizedTrendResponderExecutionsAccessLevelEnumValues() []RequestSummarizedTrendResponderExecutionsAccessLevelEnum {
	values := make([]RequestSummarizedTrendResponderExecutionsAccessLevelEnum, 0)
	for _, v := range mappingRequestSummarizedTrendResponderExecutionsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedTrendResponderExecutionsAccessLevelEnumStringValues Enumerates the set of values in String for RequestSummarizedTrendResponderExecutionsAccessLevelEnum
func GetRequestSummarizedTrendResponderExecutionsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingRequestSummarizedTrendResponderExecutionsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedTrendResponderExecutionsAccessLevelEnum(val string) (RequestSummarizedTrendResponderExecutionsAccessLevelEnum, bool) {
	enum, ok := mappingRequestSummarizedTrendResponderExecutionsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
