// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// GetEventReportRequest wrapper for the GetEventReport operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetEventReport.go.html to see an example of how to use GetEventReportRequest.
type GetEventReportRequest struct {

	// Instance Oracle Cloud identifier (ocid)
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// filter event occurrence. Selecting only those last occurred before given date in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	LatestTimestampLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"latestTimestampLessThan"`

	// filter event occurrence. Selecting only those last occurred on or after given date in ISO 8601 format
	// Example: 2017-07-14T02:40:00.000Z
	LatestTimestampGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"latestTimestampGreaterThanOrEqualTo"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetEventReportRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetEventReportRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetEventReportRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetEventReportRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetEventReportResponse wrapper for the GetEventReport operation
type GetEventReportResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The EventReport instance
	EventReport `presentIn:"body"`

	// identifier for the request
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetEventReportResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetEventReportResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
