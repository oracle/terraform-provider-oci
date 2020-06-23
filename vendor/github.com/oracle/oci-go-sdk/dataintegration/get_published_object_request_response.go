// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetPublishedObjectRequest wrapper for the GetPublishedObject operation
type GetPublishedObjectRequest struct {

	// DIS workspace id
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// DIS application key
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// DIS published object key
	PublishedObjectKey *string `mandatory:"true" contributesTo:"path" name:"publishedObjectKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This is used to expand references of the object. If value is true, then all referenced objects will be expanded. If value is false, then shallow objects will be returned in place of references. Default is false. <br><br><B>Examples:-</B><br> <ul> <li><B>?expandReferences=true</B> returns all objects of type data loader task</li> </ul>
	ExpandReferences *string `mandatory:"false" contributesTo:"query" name:"expandReferences"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPublishedObjectRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPublishedObjectRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPublishedObjectRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetPublishedObjectResponse wrapper for the GetPublishedObject operation
type GetPublishedObjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PublishedObject instance
	PublishedObject `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPublishedObjectResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPublishedObjectResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
