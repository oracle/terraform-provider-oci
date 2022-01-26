// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ChangeCatalogPrivateEndpointCompartmentRequest wrapper for the ChangeCatalogPrivateEndpointCompartment operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ChangeCatalogPrivateEndpointCompartment.go.html to see an example of how to use ChangeCatalogPrivateEndpointCompartmentRequest.
type ChangeCatalogPrivateEndpointCompartmentRequest struct {

	// Details for the target compartment.
	ChangeCatalogPrivateEndpointCompartmentDetails `contributesTo:"body"`

	// Unique private reverse connection identifier.
	CatalogPrivateEndpointId *string `mandatory:"true" contributesTo:"path" name:"catalogPrivateEndpointId"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ChangeCatalogPrivateEndpointCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ChangeCatalogPrivateEndpointCompartmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ChangeCatalogPrivateEndpointCompartmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ChangeCatalogPrivateEndpointCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ChangeCatalogPrivateEndpointCompartmentResponse wrapper for the ChangeCatalogPrivateEndpointCompartment operation
type ChangeCatalogPrivateEndpointCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID of the asynchronous request. Use GetWorkRequest (https://docs.cloud.oracle.com/api/#/en/workrequests/20160918/WorkRequest/GetWorkRequest) with this OCID to track the status of the asynchronous request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ChangeCatalogPrivateEndpointCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ChangeCatalogPrivateEndpointCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
