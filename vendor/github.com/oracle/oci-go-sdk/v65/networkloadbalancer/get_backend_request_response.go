// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetBackendRequest wrapper for the GetBackend operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkloadbalancer/GetBackend.go.html to see an example of how to use GetBackendRequest.
type GetBackendRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network load balancer to update.
	NetworkLoadBalancerId *string `mandatory:"true" contributesTo:"path" name:"networkLoadBalancerId"`

	// The name of the backend set that includes the backend server.
	// Example: `example_backend_set`
	BackendSetName *string `mandatory:"true" contributesTo:"path" name:"backendSetName"`

	// The name of the backend server to retrieve.
	// If the backend was created with an explicitly specified name, that name should be used here.
	// If the backend was created without explicitly specifying the name, but was created using ipAddress, this is specified as <ipAddress>:<port>.
	// If the backend was created without explicitly specifying the name, but was created using targetId, this is specified as <targetId>:<port>.
	// Example: `10.0.0.3:8080` or `ocid1.privateip..oc1.<var>&lt;unique_ID&gt;</var>:8080`
	BackendName *string `mandatory:"true" contributesTo:"path" name:"backendName"`

	// The unique Oracle-assigned identifier for the request. If you must contact Oracle about a
	// particular request, then provide the request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The system returns the requested resource, with a 200 status, only if the resource has no etag
	// matching the one specified. If the condition fails for the GET and HEAD methods, then the system returns the
	// HTTP status code `304 (Not Modified)`.
	// Example: `example-etag`
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetBackendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetBackendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetBackendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetBackendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetBackendRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetBackendResponse wrapper for the GetBackend operation
type GetBackendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Backend instance
	Backend `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetBackendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetBackendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
