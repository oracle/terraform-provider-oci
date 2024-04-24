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

// RequestSummarizedProblemsRequest wrapper for the RequestSummarizedProblems operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedProblems.go.html to see an example of how to use RequestSummarizedProblemsRequest.
type RequestSummarizedProblemsRequest struct {

	// The possible attributes based on which the problems can be distinguished.
	ListDimensions []ProblemDimensionEnum `contributesTo:"query" name:"listDimensions" omitEmpty:"true" collectionFormat:"multi"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel RequestSummarizedProblemsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The maximum number of items to return
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RequestSummarizedProblemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RequestSummarizedProblemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RequestSummarizedProblemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RequestSummarizedProblemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request RequestSummarizedProblemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ListDimensions {
		if _, ok := GetMappingProblemDimensionEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListDimensions: %s. Supported values are: %s.", val, strings.Join(GetProblemDimensionEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingRequestSummarizedProblemsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetRequestSummarizedProblemsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RequestSummarizedProblemsResponse wrapper for the RequestSummarizedProblems operation
type RequestSummarizedProblemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProblemAggregationCollection instances
	ProblemAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response RequestSummarizedProblemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RequestSummarizedProblemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RequestSummarizedProblemsAccessLevelEnum Enum with underlying type: string
type RequestSummarizedProblemsAccessLevelEnum string

// Set of constants representing the allowable values for RequestSummarizedProblemsAccessLevelEnum
const (
	RequestSummarizedProblemsAccessLevelRestricted RequestSummarizedProblemsAccessLevelEnum = "RESTRICTED"
	RequestSummarizedProblemsAccessLevelAccessible RequestSummarizedProblemsAccessLevelEnum = "ACCESSIBLE"
)

var mappingRequestSummarizedProblemsAccessLevelEnum = map[string]RequestSummarizedProblemsAccessLevelEnum{
	"RESTRICTED": RequestSummarizedProblemsAccessLevelRestricted,
	"ACCESSIBLE": RequestSummarizedProblemsAccessLevelAccessible,
}

var mappingRequestSummarizedProblemsAccessLevelEnumLowerCase = map[string]RequestSummarizedProblemsAccessLevelEnum{
	"restricted": RequestSummarizedProblemsAccessLevelRestricted,
	"accessible": RequestSummarizedProblemsAccessLevelAccessible,
}

// GetRequestSummarizedProblemsAccessLevelEnumValues Enumerates the set of values for RequestSummarizedProblemsAccessLevelEnum
func GetRequestSummarizedProblemsAccessLevelEnumValues() []RequestSummarizedProblemsAccessLevelEnum {
	values := make([]RequestSummarizedProblemsAccessLevelEnum, 0)
	for _, v := range mappingRequestSummarizedProblemsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetRequestSummarizedProblemsAccessLevelEnumStringValues Enumerates the set of values in String for RequestSummarizedProblemsAccessLevelEnum
func GetRequestSummarizedProblemsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingRequestSummarizedProblemsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRequestSummarizedProblemsAccessLevelEnum(val string) (RequestSummarizedProblemsAccessLevelEnum, bool) {
	enum, ok := mappingRequestSummarizedProblemsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
