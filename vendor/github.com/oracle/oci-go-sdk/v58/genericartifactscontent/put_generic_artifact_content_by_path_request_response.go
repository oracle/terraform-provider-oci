// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package genericartifactscontent

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"io"
	"net/http"
	"strings"
)

// PutGenericArtifactContentByPathRequest wrapper for the PutGenericArtifactContentByPath operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/genericartifactscontent/PutGenericArtifactContentByPath.go.html to see an example of how to use PutGenericArtifactContentByPathRequest.
type PutGenericArtifactContentByPathRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.
	// Example: `ocid1.repository.oc1..exampleuniqueID`
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// A user-defined path to describe the location of an artifact. You can use slashes to organize the repository, but slashes do not create a directory structure. An artifact path does not include an artifact version.
	// Example: `project01/my-web-app/artifact-abc`
	ArtifactPath *string `mandatory:"true" contributesTo:"path" name:"artifactPath"`

	// A user-defined string to describe the artifact version.
	// Example: `1.1.2` or `1.2-beta-2`
	Version *string `mandatory:"true" contributesTo:"path" name:"version"`

	// Uploads an artifact. Provide artifact path, version and content. Avoid entering confidential information when you define the path and version.
	GenericArtifactContentBody io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match` parameter to the value of the `etag` from a previous GET or POST response for that resource.  The resource will be updated or deleted only if the `etag` you provide matches the resource's current `etag` value. When 'if-match' is provided and its value does not exactly match the 'etag' of the resource on the server, the request fails with the 412 response code.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned request ID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm)
	// Example: `bxxxxxxx-fxxx-4xxx-9xxx-bxxxxxxxxxxx`
	// If you contact Oracle about a request, provide this request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PutGenericArtifactContentByPathRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PutGenericArtifactContentByPathRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request PutGenericArtifactContentByPathRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.GenericArtifactContentBody)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PutGenericArtifactContentByPathRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request PutGenericArtifactContentByPathRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PutGenericArtifactContentByPathResponse wrapper for the PutGenericArtifactContentByPath operation
type PutGenericArtifactContentByPathResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The GenericArtifact instance
	GenericArtifact `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response PutGenericArtifactContentByPathResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PutGenericArtifactContentByPathResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
