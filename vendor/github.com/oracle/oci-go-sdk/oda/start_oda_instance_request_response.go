// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// StartOdaInstanceRequest wrapper for the StartOdaInstance operation
type StartOdaInstanceRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control in a PUT or DELETE call for
	// a Digital Assistant instance, set the `if-match` query parameter
	// to the value of the `ETAG` header from a previous GET or POST
	// response for that instance. The service updates or deletes the
	// instance only if the etag that you provide matches the instance's
	// current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token that uniquely identifies a request so that you can retry the request if there's
	// a timeout or server error without the risk of executing that same action again.
	// Retry tokens expire after 24 hours, but they can become invalid before then if there are
	// conflicting operations. For example, if an instance was deleted and purged from the system,
	// then the service might reject a retry of the original creation request.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request StartOdaInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request StartOdaInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request StartOdaInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// StartOdaInstanceResponse wrapper for the StartOdaInstance operation
type StartOdaInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the asynchronous request. You can use this to query status of
	// the operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. The request id consists of 3 components.
	// The first is the request ID provided by the client in the request headers and the remaining 2 are
	// generated for tracing requests.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response StartOdaInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response StartOdaInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
