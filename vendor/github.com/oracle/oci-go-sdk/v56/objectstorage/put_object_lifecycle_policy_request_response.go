// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// PutObjectLifecyclePolicyRequest wrapper for the PutObjectLifecyclePolicy operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/objectstorage/PutObjectLifecyclePolicy.go.html to see an example of how to use PutObjectLifecyclePolicyRequest.
type PutObjectLifecyclePolicyRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The lifecycle policy to apply to the bucket.
	PutObjectLifecyclePolicyDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// The entity tag (ETag) to match with the ETag of an existing resource. If the specified ETag matches the ETag of
	// the existing resource, GET and HEAD requests will return the resource and PUT and POST requests will upload
	// the resource.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The entity tag (ETag) to avoid matching. The only valid value is '*', which indicates that the request should
	// fail if the resource already exists.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PutObjectLifecyclePolicyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PutObjectLifecyclePolicyRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request PutObjectLifecyclePolicyRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PutObjectLifecyclePolicyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// PutObjectLifecyclePolicyResponse wrapper for the PutObjectLifecyclePolicy operation
type PutObjectLifecyclePolicyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ObjectLifecyclePolicy instance
	ObjectLifecyclePolicy `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// The entity tag (ETag) for the object lifecycle policy.
	ETag *string `presentIn:"header" name:"etag"`
}

func (response PutObjectLifecyclePolicyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PutObjectLifecyclePolicyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
