// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// GetStepArtifactContentRequest wrapper for the GetStepArtifactContent operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/GetStepArtifactContent.go.html to see an example of how to use GetStepArtifactContentRequest.
type GetStepArtifactContentRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pipeline.
	PipelineId *string `mandatory:"true" contributesTo:"path" name:"pipelineId"`

	// Unique Step identifier in a pipeline.
	StepName *string `mandatory:"true" contributesTo:"path" name:"stepName"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Optional byte range to fetch, as described in RFC 7233 (https://tools.ietf.org/html/rfc7232#section-2.1), section 2.1.
	// Note that only a single range of bytes is supported.
	Range *string `mandatory:"false" contributesTo:"header" name:"range"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetStepArtifactContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetStepArtifactContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetStepArtifactContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetStepArtifactContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetStepArtifactContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetStepArtifactContentResponse wrapper for the GetStepArtifactContent operation
type GetStepArtifactContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The content size of the body in bytes.
	ContentLength *int64 `presentIn:"header" name:"content-length"`

	// The base-64 encoded MD5 hash of the body, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.15.
	// Unavailable for objects uploaded using multipart upload.
	// If the `content-md5` header is present, Object Storage performs an integrity check on the body of the HTTP request by computing the MD5 hash for the body and comparing it to the MD5 hash supplied in the header.
	// If the two hashes do not match, the object is rejected and an HTTP-400 Unmatched Content MD5 error is returned with the message:
	// "The computed MD5 of the request body (ACTUAL_MD5) does not match the Content-MD5 header (HEADER_MD5)"
	ContentMd5 *string `presentIn:"header" name:"content-md5"`

	// The artifact modification time, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 14.29.
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`

	// The content disposition of the body, as described in RFC 2616 (https://tools.ietf.org/rfc/rfc2616), section 19.5.1.
	ContentDisposition *string `presentIn:"header" name:"content-disposition"`
}

func (response GetStepArtifactContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetStepArtifactContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
