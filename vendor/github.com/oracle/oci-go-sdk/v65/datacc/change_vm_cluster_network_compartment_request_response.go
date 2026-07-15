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

// ChangeVmClusterNetworkCompartmentRequest wrapper for the ChangeVmClusterNetworkCompartment operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacc/ChangeVmClusterNetworkCompartment.go.html to see an example of how to use ChangeVmClusterNetworkCompartmentRequest.
type ChangeVmClusterNetworkCompartmentRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.
	VmClusterNetworkId *string `mandatory:"true" contributesTo:"path" name:"vmClusterNetworkId"`

	// Updated Compartment details.
	ChangeVmClusterNetworkCompartmentDetails `contributesTo:"body"`

	// For Optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the 'if-match' parameter to the value of the
	// Etag from a previous GET or POST response for that resource.
	// The resource is updated or deleted only if the Etag that you
	// provide matches the current Etag value for the resource.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token that uniquely identifies a request, so the request can be retried without risk of processing that
	// same action again, in the event of a timeout or server error. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// was deleted and purged from the system, then a retry of the original creation request
	// can be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ChangeVmClusterNetworkCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ChangeVmClusterNetworkCompartmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ChangeVmClusterNetworkCompartmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ChangeVmClusterNetworkCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ChangeVmClusterNetworkCompartmentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChangeVmClusterNetworkCompartmentResponse wrapper for the ChangeVmClusterNetworkCompartment operation
type ChangeVmClusterNetworkCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique identifier created by Oracle for the asynchronous request. You can use this identifier to query the status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique identifier assigned by Oracle for the request. If you need to contact
	// Oracle about a particular request, then please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ChangeVmClusterNetworkCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ChangeVmClusterNetworkCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
