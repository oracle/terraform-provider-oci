// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"net/http"
)

// UpdateInternalPublicIpRequest wrapper for the UpdateInternalPublicIp operation
type UpdateInternalPublicIpRequest struct {

	// The OCID of the internal public IP.
	InternalPublicIpId *string `mandatory:"true" contributesTo:"path" name:"internalPublicIpId"`

	// Internal Public IP details.
	UpdateInternalPublicIpDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// This is the operation name used for authorization. This is only used when the API is called by a service as part of another API.
	InternalAuthzOperationName *string `mandatory:"false" contributesTo:"header" name:"internal-authz-operation-name"`

	// This is the resource kind used for authorization. This is only used when the API is called by a service as part of another API.
	InternalAuthzResourceKind *string `mandatory:"false" contributesTo:"header" name:"internal-authz-resource-kind"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateInternalPublicIpRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateInternalPublicIpRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateInternalPublicIpRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateInternalPublicIpResponse wrapper for the UpdateInternalPublicIp operation
type UpdateInternalPublicIpResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InternalPublicIp instance
	InternalPublicIp `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateInternalPublicIpResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateInternalPublicIpResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
