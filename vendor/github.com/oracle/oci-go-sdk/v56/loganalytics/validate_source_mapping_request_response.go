// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ValidateSourceMappingRequest wrapper for the ValidateSourceMapping operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ValidateSourceMapping.go.html to see an example of how to use ValidateSourceMappingRequest.
type ValidateSourceMappingRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Location of the log file.
	ObjectLocation *string `mandatory:"true" contributesTo:"query" name:"objectLocation"`

	// The name of the file being uploaded. The extension of the filename part will be used to detect the type of file (like zip, tar).
	Filename *string `mandatory:"true" contributesTo:"query" name:"filename"`

	// Name of the log source that will be used to process the files being uploaded.
	LogSourceName *string `mandatory:"true" contributesTo:"query" name:"logSourceName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ValidateSourceMappingRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ValidateSourceMappingRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ValidateSourceMappingRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ValidateSourceMappingRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateSourceMappingResponse wrapper for the ValidateSourceMapping operation
type ValidateSourceMappingResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SourceMappingResponse instance
	SourceMappingResponse `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ValidateSourceMappingResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ValidateSourceMappingResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
