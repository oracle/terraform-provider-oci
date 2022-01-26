// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package genericartifactscontent

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"io"
	"net/http"
)

// GetGenericArtifactContentRequest wrapper for the GetGenericArtifactContent operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/genericartifactscontent/GetGenericArtifactContent.go.html to see an example of how to use GetGenericArtifactContentRequest.
type GetGenericArtifactContentRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the artifact.
	// Example: `ocid1.genericartifact.oc1..exampleuniqueID`
	ArtifactId *string `mandatory:"true" contributesTo:"path" name:"artifactId"`

	// Unique Oracle-assigned request ID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm)
	// Example: `bxxxxxxx-fxxx-4xxx-9xxx-bxxxxxxxxxxx`
	// If you contact Oracle about a request, provide this request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetGenericArtifactContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetGenericArtifactContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetGenericArtifactContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetGenericArtifactContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetGenericArtifactContentResponse wrapper for the GetGenericArtifactContent operation
type GetGenericArtifactContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetGenericArtifactContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetGenericArtifactContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
