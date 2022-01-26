// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// DeleteObjectRequest wrapper for the DeleteObject operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/objectstorage/DeleteObject.go.html to see an example of how to use DeleteObjectRequest.
type DeleteObjectRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The name of the object. Avoid entering confidential information.
	// Example: `test/object1.log`
	ObjectName *string `mandatory:"true" contributesTo:"path" name:"objectName"`

	// The entity tag (ETag) to match with the ETag of an existing resource. If the specified ETag matches the ETag of
	// the existing resource, GET and HEAD requests will return the resource and PUT and POST requests will upload
	// the resource.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// VersionId used to identify a particular version of the object
	VersionId *string `mandatory:"false" contributesTo:"query" name:"versionId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteObjectRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteObjectRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteObjectRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteObjectRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteObjectResponse wrapper for the DeleteObject operation
type DeleteObjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The time the object was deleted, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29).
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`

	// The `versionId` of the delete marker created as a result of the DELETE Object.
	// If the request contains a specific `versionId`, then this response header will be the same as the requested `versionId` of the object that was deleted.
	VersionId *string `presentIn:"header" name:"version-id"`

	// This is `true` if the deleted object is a delete marker, otherwise `false`
	IsDeleteMarker *bool `presentIn:"header" name:"is-delete-marker"`
}

func (response DeleteObjectResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteObjectResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
