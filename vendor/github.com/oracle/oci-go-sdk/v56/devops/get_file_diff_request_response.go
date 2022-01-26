// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// GetFileDiffRequest wrapper for the GetFileDiff operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetFileDiff.go.html to see an example of how to use GetFileDiffRequest.
type GetFileDiffRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// Path to a file within a repository.
	FilePath *string `mandatory:"true" contributesTo:"path" name:"filePath"`

	// The branch to compare changes against.
	BaseVersion *string `mandatory:"true" contributesTo:"query" name:"baseVersion"`

	// The branch where changes are coming from.
	TargetVersion *string `mandatory:"true" contributesTo:"query" name:"targetVersion"`

	// Boolean to indicate whether to use merge base or most recent revision.
	IsComparisonFromMergeBase *bool `mandatory:"false" contributesTo:"query" name:"isComparisonFromMergeBase"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetFileDiffRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetFileDiffRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetFileDiffRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetFileDiffRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetFileDiffResponse wrapper for the GetFileDiff operation
type GetFileDiffResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The FileDiffResponse instance
	FileDiffResponse `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetFileDiffResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetFileDiffResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
