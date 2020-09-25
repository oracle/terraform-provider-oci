// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"net/http"
)

// AcceptLocalPeeringTokenRequest wrapper for the AcceptLocalPeeringToken operation
type AcceptLocalPeeringTokenRequest struct {

	// The OCID of the local peering connection. This feature is currently in preview and may change before public release. Do not use it for production workloads.
	LocalPeeringConnectionId *string `mandatory:"true" contributesTo:"path" name:"localPeeringConnectionId"`

	// Details regarding the token to accept.
	AcceptLocalPeeringTokenDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request AcceptLocalPeeringTokenRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request AcceptLocalPeeringTokenRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request AcceptLocalPeeringTokenRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// AcceptLocalPeeringTokenResponse wrapper for the AcceptLocalPeeringToken operation
type AcceptLocalPeeringTokenResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response AcceptLocalPeeringTokenResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response AcceptLocalPeeringTokenResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
