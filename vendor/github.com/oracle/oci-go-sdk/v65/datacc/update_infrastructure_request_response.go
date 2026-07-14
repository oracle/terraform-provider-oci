// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacc

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// UpdateInfrastructureRequest wrapper for the UpdateInfrastructure operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/UpdateInfrastructure.go.html to see an example of how to use UpdateInfrastructureRequest.
type UpdateInfrastructureRequest struct {

	// The updated Database Infrastructure details.
	UpdateInfrastructureDetails `contributesTo:"body"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Infrastructure.
	InfrastructureId *string `mandatory:"true" contributesTo:"path" name:"infrastructureId"`

	// For Optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the 'if-match' parameter to the value of the
	// Etag from a previous GET or POST response for that resource.
	// The resource is updated or deleted only if the Etag that you
	// provide matches the current Etag value for the resource.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateInfrastructureRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateInfrastructureRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateInfrastructureRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateInfrastructureRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateInfrastructureRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateInfrastructureResponse wrapper for the UpdateInfrastructure operation
type UpdateInfrastructureResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Infrastructure instance
	Infrastructure `presentIn:"body"`

	// Unique identifier created by Oracle for the asynchronous request. You can use this identifier to query the status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique identifier assigned by Oracle for the request. If you need to contact
	// Oracle about a particular request, then please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For Optimistic concurrency control. See the `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// The full URI of the resource.
	Location *string `presentIn:"header" name:"location"`
}

func (response UpdateInfrastructureResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateInfrastructureResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
