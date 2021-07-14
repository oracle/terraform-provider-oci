// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"github.com/oracle/oci-go-sdk/v44/common"
	"net/http"
)

// RequestSummarizedManagedInstanceUsageRequest wrapper for the RequestSummarizedManagedInstanceUsage operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/RequestSummarizedManagedInstanceUsage.go.html to see an example of how to use RequestSummarizedManagedInstanceUsageRequest.
type RequestSummarizedManagedInstanceUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Fleet.
	FleetId *string `mandatory:"true" contributesTo:"path" name:"fleetId"`

	// Parameters for filtering managed instance usage.
	RequestSummarizedManagedInstanceUsageDetails `contributesTo:"body"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RequestSummarizedManagedInstanceUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RequestSummarizedManagedInstanceUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RequestSummarizedManagedInstanceUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RequestSummarizedManagedInstanceUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// RequestSummarizedManagedInstanceUsageResponse wrapper for the RequestSummarizedManagedInstanceUsage operation
type RequestSummarizedManagedInstanceUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedInstanceUsageCollection instances
	ManagedInstanceUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response RequestSummarizedManagedInstanceUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RequestSummarizedManagedInstanceUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
