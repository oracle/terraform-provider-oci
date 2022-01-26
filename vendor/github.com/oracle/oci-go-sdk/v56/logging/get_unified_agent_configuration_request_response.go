// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// GetUnifiedAgentConfigurationRequest wrapper for the GetUnifiedAgentConfiguration operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/logging/GetUnifiedAgentConfiguration.go.html to see an example of how to use GetUnifiedAgentConfigurationRequest.
type GetUnifiedAgentConfigurationRequest struct {

	// The OCID of the Unified Agent configuration.
	UnifiedAgentConfigurationId *string `mandatory:"true" contributesTo:"path" name:"unifiedAgentConfigurationId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetUnifiedAgentConfigurationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetUnifiedAgentConfigurationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetUnifiedAgentConfigurationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetUnifiedAgentConfigurationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetUnifiedAgentConfigurationResponse wrapper for the GetUnifiedAgentConfiguration operation
type GetUnifiedAgentConfigurationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The UnifiedAgentConfiguration instance
	UnifiedAgentConfiguration `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetUnifiedAgentConfigurationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetUnifiedAgentConfigurationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
