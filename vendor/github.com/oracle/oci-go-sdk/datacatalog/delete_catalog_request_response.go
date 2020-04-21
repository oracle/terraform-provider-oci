// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// DeleteCatalogRequest wrapper for the DeleteCatalog operation
type DeleteCatalogRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

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

func (request DeleteCatalogRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteCatalogRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteCatalogRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteCatalogResponse wrapper for the DeleteCatalog operation
type DeleteCatalogResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID of the asynchronous request. Use GetWorkRequest (https://docs.cloud.oracle.com/api/#/en/workrequests/20160918/WorkRequest/GetWorkRequest) with this OCID to track the status of the asynchronous request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeleteCatalogResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteCatalogResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
