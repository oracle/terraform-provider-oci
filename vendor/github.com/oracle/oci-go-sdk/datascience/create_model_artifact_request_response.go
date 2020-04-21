// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// CreateModelArtifactRequest wrapper for the CreateModelArtifact operation
type CreateModelArtifactRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the model.
	ModelId *string `mandatory:"true" contributesTo:"path" name:"modelId"`

	// The content length of the body.
	ContentLength *int64 `mandatory:"true" contributesTo:"header" name:"content-length"`

	// The model artifact to upload.
	ModelArtifact io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or server error without risk of executing that same action again. Retry tokens expire after 24 hours, but can be invalidated before then due to conflicting operations. For example, if a resource has been deleted and purged from the system, then a retry of the original creation request might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The content disposition of the body.
	ContentDisposition *string `mandatory:"false" contributesTo:"header" name:"content-disposition"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateModelArtifactRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateModelArtifactRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateModelArtifactRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateModelArtifactResponse wrapper for the CreateModelArtifact operation
type CreateModelArtifactResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateModelArtifactResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateModelArtifactResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
