// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// HeadObjectRequest wrapper for the HeadObject operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/objectstorage/HeadObject.go.html to see an example of how to use HeadObjectRequest.
type HeadObjectRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the bucket. Avoid entering confidential information.
	// Example: `my-new-bucket1`
	BucketName *string `mandatory:"true" contributesTo:"path" name:"bucketName"`

	// The name of the object. Avoid entering confidential information.
	// Example: `test/object1.log`
	ObjectName *string `mandatory:"true" contributesTo:"path" name:"objectName"`

	// VersionId used to identify a particular version of the object
	VersionId *string `mandatory:"false" contributesTo:"query" name:"versionId"`

	// The entity tag (ETag) to match with the ETag of an existing resource. If the specified ETag matches the ETag of
	// the existing resource, GET and HEAD requests will return the resource and PUT and POST requests will upload
	// the resource.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The entity tag (ETag) to avoid matching. Wildcards ('*') are not allowed. If the specified ETag does not
	// match the ETag of the existing resource, the request returns the expected response. If the ETag matches
	// the ETag of the existing resource, the request returns an HTTP 304 status without a response body.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// The optional header that specifies "AES256" as the encryption algorithm. For more information, see
	// Using Your Own Keys for Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).
	OpcSseCustomerAlgorithm *string `mandatory:"false" contributesTo:"header" name:"opc-sse-customer-algorithm"`

	// The optional header that specifies the base64-encoded 256-bit encryption key to use to encrypt or
	// decrypt the data. For more information, see
	// Using Your Own Keys for Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).
	OpcSseCustomerKey *string `mandatory:"false" contributesTo:"header" name:"opc-sse-customer-key"`

	// The optional header that specifies the base64-encoded SHA256 hash of the encryption key. This
	// value is used to check the integrity of the encryption key. For more information, see
	// Using Your Own Keys for Server-Side Encryption (https://docs.cloud.oracle.com/Content/Object/Tasks/usingyourencryptionkeys.htm).
	OpcSseCustomerKeySha256 *string `mandatory:"false" contributesTo:"header" name:"opc-sse-customer-key-sha256"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request HeadObjectRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request HeadObjectRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request HeadObjectRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request HeadObjectRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// HeadObjectResponse wrapper for the HeadObject operation
type HeadObjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The entity tag (ETag) for the object.
	ETag *string `presentIn:"header" name:"etag"`

	// The user-defined metadata for the object.
	OpcMeta map[string]string `presentIn:"header-collection" prefix:"opc-meta-"`

	// The object size in bytes.
	ContentLength *int64 `presentIn:"header" name:"content-length"`

	// Content-MD5 header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.15).
	// Unavailable for objects uploaded using multipart upload.
	ContentMd5 *string `presentIn:"header" name:"content-md5"`

	// Only applicable to objects uploaded using multipart upload.
	// Base-64 representation of the multipart object hash.
	// The multipart object hash is calculated by taking the MD5 hashes of the parts,
	// concatenating the binary representation of those hashes in order of their part numbers,
	// and then calculating the MD5 hash of the concatenated values.
	OpcMultipartMd5 *string `presentIn:"header" name:"opc-multipart-md5"`

	// Content-Type header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.17).
	ContentType *string `presentIn:"header" name:"content-type"`

	// Content-Language header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.12).
	ContentLanguage *string `presentIn:"header" name:"content-language"`

	// Content-Encoding header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.11).
	ContentEncoding *string `presentIn:"header" name:"content-encoding"`

	// Cache-Control header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.9).
	CacheControl *string `presentIn:"header" name:"cache-control"`

	// Content-Disposition header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-19.5.1).
	ContentDisposition *string `presentIn:"header" name:"content-disposition"`

	// The object modification time, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29).
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`

	// The storage tier that the object is stored in.
	StorageTier HeadObjectStorageTierEnum `presentIn:"header" name:"storage-tier"`

	// Archival state of an object. This field is set only for objects in Archive tier.
	ArchivalState HeadObjectArchivalStateEnum `presentIn:"header" name:"archival-state"`

	// Time that the object is returned to the archived state. This field is only present for restored objects.
	TimeOfArchival *common.SDKTime `presentIn:"header" name:"time-of-archival"`

	// VersionId of the object requested
	VersionId *string `presentIn:"header" name:"version-id"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response HeadObjectResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response HeadObjectResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// HeadObjectStorageTierEnum Enum with underlying type: string
type HeadObjectStorageTierEnum string

// Set of constants representing the allowable values for HeadObjectStorageTierEnum
const (
	HeadObjectStorageTierStandard         HeadObjectStorageTierEnum = "Standard"
	HeadObjectStorageTierInfrequentaccess HeadObjectStorageTierEnum = "InfrequentAccess"
	HeadObjectStorageTierArchive          HeadObjectStorageTierEnum = "Archive"
)

var mappingHeadObjectStorageTier = map[string]HeadObjectStorageTierEnum{
	"Standard":         HeadObjectStorageTierStandard,
	"InfrequentAccess": HeadObjectStorageTierInfrequentaccess,
	"Archive":          HeadObjectStorageTierArchive,
}

// GetHeadObjectStorageTierEnumValues Enumerates the set of values for HeadObjectStorageTierEnum
func GetHeadObjectStorageTierEnumValues() []HeadObjectStorageTierEnum {
	values := make([]HeadObjectStorageTierEnum, 0)
	for _, v := range mappingHeadObjectStorageTier {
		values = append(values, v)
	}
	return values
}

// HeadObjectArchivalStateEnum Enum with underlying type: string
type HeadObjectArchivalStateEnum string

// Set of constants representing the allowable values for HeadObjectArchivalStateEnum
const (
	HeadObjectArchivalStateArchived  HeadObjectArchivalStateEnum = "Archived"
	HeadObjectArchivalStateRestoring HeadObjectArchivalStateEnum = "Restoring"
	HeadObjectArchivalStateRestored  HeadObjectArchivalStateEnum = "Restored"
)

var mappingHeadObjectArchivalState = map[string]HeadObjectArchivalStateEnum{
	"Archived":  HeadObjectArchivalStateArchived,
	"Restoring": HeadObjectArchivalStateRestoring,
	"Restored":  HeadObjectArchivalStateRestored,
}

// GetHeadObjectArchivalStateEnumValues Enumerates the set of values for HeadObjectArchivalStateEnum
func GetHeadObjectArchivalStateEnumValues() []HeadObjectArchivalStateEnum {
	values := make([]HeadObjectArchivalStateEnum, 0)
	for _, v := range mappingHeadObjectArchivalState {
		values = append(values, v)
	}
	return values
}
