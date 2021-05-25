// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// DisableAutonomousDatabaseOperationsInsightsRequest wrapper for the DisableAutonomousDatabaseOperationsInsights operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/DisableAutonomousDatabaseOperationsInsights.go.html to see an example of how to use DisableAutonomousDatabaseOperationsInsightsRequest.
type DisableAutonomousDatabaseOperationsInsightsRequest struct {

	// The database OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	AutonomousDatabaseId *string `mandatory:"true" contributesTo:"path" name:"autonomousDatabaseId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DisableAutonomousDatabaseOperationsInsightsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DisableAutonomousDatabaseOperationsInsightsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DisableAutonomousDatabaseOperationsInsightsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DisableAutonomousDatabaseOperationsInsightsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DisableAutonomousDatabaseOperationsInsightsResponse wrapper for the DisableAutonomousDatabaseOperationsInsights operation
type DisableAutonomousDatabaseOperationsInsightsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request. Multiple OCID values are returned in a comma-separated list. Use GetWorkRequest with a work request OCID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you must contact Oracle about
	// a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DisableAutonomousDatabaseOperationsInsightsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DisableAutonomousDatabaseOperationsInsightsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
