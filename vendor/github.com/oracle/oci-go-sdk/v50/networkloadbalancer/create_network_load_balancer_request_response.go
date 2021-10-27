// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkloadbalancer

import (
	"github.com/oracle/oci-go-sdk/v50/common"
	"net/http"
)

// CreateNetworkLoadBalancerRequest wrapper for the CreateNetworkLoadBalancer operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkloadbalancer/CreateNetworkLoadBalancer.go.html to see an example of how to use CreateNetworkLoadBalancerRequest.
type CreateNetworkLoadBalancerRequest struct {

	// Details for the new network load balancer.
	CreateNetworkLoadBalancerDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so that it can be retried in case of a timeout or
	// server error without risk of rerunning that same action. Retry tokens expire after 24
	// hours but they can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The unique Oracle-assigned identifier for the request. If you must contact Oracle about a
	// particular request, then provide the request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateNetworkLoadBalancerRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateNetworkLoadBalancerRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateNetworkLoadBalancerRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateNetworkLoadBalancerRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateNetworkLoadBalancerResponse wrapper for the CreateNetworkLoadBalancer operation
type CreateNetworkLoadBalancerResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The NetworkLoadBalancer instance
	NetworkLoadBalancer `presentIn:"body"`

	// Unique Oracle-assigned identifier for the asynchronous request. You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response CreateNetworkLoadBalancerResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateNetworkLoadBalancerResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
