// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package desktops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// StopDesktopRequest wrapper for the StopDesktop operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/desktops/StopDesktop.go.html to see an example of how to use StopDesktopRequest.
type StopDesktopRequest struct {

	// The OCID of the desktop.
	DesktopId *string `mandatory:"true" contributesTo:"path" name:"desktopId"`

	// The unique identifier of the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token that uniquely identifies a request.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Force a STOP(power off) of the desktop if set to false
	IsSoftStop *bool `mandatory:"false" contributesTo:"query" name:"isSoftStop"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request StopDesktopRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request StopDesktopRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request StopDesktopRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request StopDesktopRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request StopDesktopRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StopDesktopResponse wrapper for the StopDesktop operation
type StopDesktopResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The unique identifier of the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// A unique identifier for an asynchronous request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response StopDesktopResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response StopDesktopResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
