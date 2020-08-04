// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementagent

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// GetManagementAgentInstallKeyContentRequest wrapper for the GetManagementAgentInstallKeyContent operation
type GetManagementAgentInstallKeyContentRequest struct {

	// Unique Management Agent Install Key identifier
	ManagementAgentInstallKeyId *string `mandatory:"true" contributesTo:"path" name:"managementAgentInstallKeyId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Filter to return input plugin names uncommented in the output.
	PluginName []string `contributesTo:"query" name:"pluginName" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetManagementAgentInstallKeyContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetManagementAgentInstallKeyContentRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetManagementAgentInstallKeyContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetManagementAgentInstallKeyContentResponse wrapper for the GetManagementAgentInstallKeyContent operation
type GetManagementAgentInstallKeyContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The content size of the body in bytes.
	ContentLength *int64 `presentIn:"header" name:"content-length"`

	// The content type of the body.
	ContentType *string `presentIn:"header" name:"content-type"`
}

func (response GetManagementAgentInstallKeyContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetManagementAgentInstallKeyContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
