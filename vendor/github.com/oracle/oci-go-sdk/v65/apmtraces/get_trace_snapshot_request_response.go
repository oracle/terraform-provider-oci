// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetTraceSnapshotRequest wrapper for the GetTraceSnapshot operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/GetTraceSnapshot.go.html to see an example of how to use GetTraceSnapshotRequest.
type GetTraceSnapshotRequest struct {

	// The APM Domain ID for the intended request.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// Unique Application Performance Monitoring trace identifier (traceId).
	TraceKey *string `mandatory:"true" contributesTo:"path" name:"traceKey"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If enabled, only span level details are sent.
	IsSummarized *bool `mandatory:"false" contributesTo:"query" name:"isSummarized"`

	// Thread ID for which snapshots need to be retrieved. This identifier of a thread is a long positive number generated when a thread is created.
	ThreadId *string `mandatory:"false" contributesTo:"query" name:"threadId"`

	// Epoch time of snapshot.
	SnapshotTime *string `mandatory:"false" contributesTo:"query" name:"snapshotTime"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetTraceSnapshotRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetTraceSnapshotRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetTraceSnapshotRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetTraceSnapshotRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetTraceSnapshotRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetTraceSnapshotResponse wrapper for the GetTraceSnapshot operation
type GetTraceSnapshotResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TraceSnapshot instance
	TraceSnapshot `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTraceSnapshotResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetTraceSnapshotResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
