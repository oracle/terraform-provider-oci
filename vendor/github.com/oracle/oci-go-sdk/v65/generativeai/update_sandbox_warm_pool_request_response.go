// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// UpdateSandboxWarmPoolRequest wrapper for the UpdateSandboxWarmPool operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/UpdateSandboxWarmPool.go.html to see an example of how to use UpdateSandboxWarmPoolRequest.
type UpdateSandboxWarmPoolRequest struct {

	// Owning Generative AI Project OCID. Resolve parent compartment for authorization.
	ProjectId *string `mandatory:"true" contributesTo:"path" name:"projectId"`

	// Immutable name within the Project.
	Name *string `mandatory:"true" contributesTo:"path" name:"name"`

	// Replacement desired capacity to persist.
	UpdateSandboxWarmPoolDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the
	// `if-match` parameter to the value of the etag from a previous GET or POST response for
	// that resource. The resource will be updated or deleted only if the etag you provide
	// matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateSandboxWarmPoolRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateSandboxWarmPoolRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateSandboxWarmPoolRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateSandboxWarmPoolRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateSandboxWarmPoolRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateSandboxWarmPoolResponse wrapper for the UpdateSandboxWarmPool operation
type UpdateSandboxWarmPoolResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SandboxWarmPool instance
	SandboxWarmPool `presentIn:"body"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	// Use GetWorkRequest with this ID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Opaque ETag of the persisted configuration.
	Etag *string `presentIn:"header" name:"etag"`

	// Request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateSandboxWarmPoolResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateSandboxWarmPoolResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
