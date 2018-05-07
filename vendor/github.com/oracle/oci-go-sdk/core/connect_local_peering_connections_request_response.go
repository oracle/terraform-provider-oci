// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ConnectLocalPeeringConnectionsRequest wrapper for the ConnectLocalPeeringConnections operation
type ConnectLocalPeeringConnectionsRequest struct {

	// The OCID of the local peering connection. This feature is currently in preview and may change before public release. Do not use it for production workloads.
	LocalPeeringConnectionId *string `mandatory:"true" contributesTo:"path" name:"localPeeringConnectionId"`

	// Details regarding the local peering connection to connect.
	ConnectLocalPeeringConnectionsDetails `contributesTo:"body"`

	// A comma separated list of tenancy OCIDs that might be accessed by this request. Only required
	// for cross tenancy requests. May be `null` for requests that do not cross tenancy boundaries.
	XCrossTenancyRequest *string `mandatory:"false" contributesTo:"header" name:"x-cross-tenancy-request"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ConnectLocalPeeringConnectionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ConnectLocalPeeringConnectionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ConnectLocalPeeringConnectionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ConnectLocalPeeringConnectionsResponse wrapper for the ConnectLocalPeeringConnections operation
type ConnectLocalPeeringConnectionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ConnectLocalPeeringConnectionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ConnectLocalPeeringConnectionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
