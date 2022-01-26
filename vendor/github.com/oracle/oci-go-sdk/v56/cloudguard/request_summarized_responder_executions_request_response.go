// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// RequestSummarizedResponderExecutionsRequest wrapper for the RequestSummarizedResponderExecutions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/RequestSummarizedResponderExecutions.go.html to see an example of how to use RequestSummarizedResponderExecutionsRequest.
type RequestSummarizedResponderExecutionsRequest struct {

	// The possible attributes based on which the responder executions can be distinguished
	ResponderExecutionsDimensions []ResponderDimensionEnum `contributesTo:"query" name:"responderExecutionsDimensions" omitEmpty:"true" collectionFormat:"multi"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The possible filters for Responder Type Dimension to distinguish Responder Executions.
	// If no values are passed, the metric for responder executions of all reponder types are returned
	ResponderTypeFilter []ResponderTypeEnum `contributesTo:"query" name:"responderTypeFilter" omitEmpty:"true" collectionFormat:"multi"`

	// The possible filters for Responder Type Dimension to distinguish Responder Executions.
	// If no values are passed, the metric for responder executions of all status are returned
	ResponderExecutionStatusFilter []ResponderExecutionStatusEnum `contributesTo:"query" name:"responderExecutionStatusFilter" omitEmpty:"true" collectionFormat:"multi"`

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
	AccessLevel RequestSummarizedResponderExecutionsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

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

func (request RequestSummarizedResponderExecutionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RequestSummarizedResponderExecutionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RequestSummarizedResponderExecutionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RequestSummarizedResponderExecutionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// RequestSummarizedResponderExecutionsResponse wrapper for the RequestSummarizedResponderExecutions operation
type RequestSummarizedResponderExecutionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResponderExecutionAggregationCollection instances
	ResponderExecutionAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response RequestSummarizedResponderExecutionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RequestSummarizedResponderExecutionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RequestSummarizedResponderExecutionsAccessLevelEnum Enum with underlying type: string
type RequestSummarizedResponderExecutionsAccessLevelEnum string

// Set of constants representing the allowable values for RequestSummarizedResponderExecutionsAccessLevelEnum
const (
	RequestSummarizedResponderExecutionsAccessLevelRestricted RequestSummarizedResponderExecutionsAccessLevelEnum = "RESTRICTED"
	RequestSummarizedResponderExecutionsAccessLevelAccessible RequestSummarizedResponderExecutionsAccessLevelEnum = "ACCESSIBLE"
)

var mappingRequestSummarizedResponderExecutionsAccessLevel = map[string]RequestSummarizedResponderExecutionsAccessLevelEnum{
	"RESTRICTED": RequestSummarizedResponderExecutionsAccessLevelRestricted,
	"ACCESSIBLE": RequestSummarizedResponderExecutionsAccessLevelAccessible,
}

// GetRequestSummarizedResponderExecutionsAccessLevelEnumValues Enumerates the set of values for RequestSummarizedResponderExecutionsAccessLevelEnum
func GetRequestSummarizedResponderExecutionsAccessLevelEnumValues() []RequestSummarizedResponderExecutionsAccessLevelEnum {
	values := make([]RequestSummarizedResponderExecutionsAccessLevelEnum, 0)
	for _, v := range mappingRequestSummarizedResponderExecutionsAccessLevel {
		values = append(values, v)
	}
	return values
}
