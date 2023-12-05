// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// HeadAwrHubObjectRequest wrapper for the HeadAwrHubObject operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/HeadAwrHubObject.go.html to see an example of how to use HeadAwrHubObjectRequest.
type HeadAwrHubObjectRequest struct {

	// Unique Awr Hub Source identifier
	AwrHubSourceId *string `mandatory:"true" contributesTo:"path" name:"awrHubSourceId"`

	// Unique Awr Hub Object identifier
	ObjectName *string `mandatory:"true" contributesTo:"path" name:"objectName"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request HeadAwrHubObjectRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request HeadAwrHubObjectRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request HeadAwrHubObjectRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request HeadAwrHubObjectRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request HeadAwrHubObjectRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HeadAwrHubObjectResponse wrapper for the HeadAwrHubObject operation
type HeadAwrHubObjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// The user-defined metadata for the Awr Hub object.
	OpcMeta map[string]string `presentIn:"header-collection" prefix:"opc-meta-"`

	// The Awr Hub object size in bytes.
	ContentLength *int64 `presentIn:"header" name:"content-length"`

	// Content-MD5 header.
	ContentMd5 *string `presentIn:"header" name:"content-md5"`

	OpcMultipartMd5 *string `presentIn:"header" name:"opc-multipart-md5"`

	// Content-Type header.
	ContentType *string `presentIn:"header" name:"content-type"`

	// Content-Language header.
	ContentLanguage *string `presentIn:"header" name:"content-language"`

	// Content-Encoding header.
	ContentEncoding *string `presentIn:"header" name:"content-encoding"`

	// Cache-Control header.
	CacheControl *string `presentIn:"header" name:"cache-control"`

	// Content-Disposition header.
	ContentDisposition *string `presentIn:"header" name:"content-disposition"`

	// The Awr Hub object modification time.
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`

	// The storage tier that the Awr Hub object is stored in.
	StorageTier HeadAwrHubObjectStorageTierEnum `presentIn:"header" name:"storage-tier"`

	// Archival state of an Awr Hub object. This field is set only for Awr Hub objects in Archive tier.
	ArchivalState HeadAwrHubObjectArchivalStateEnum `presentIn:"header" name:"archival-state"`

	// Time that the Awr Hub object is returned to the archived state.
	TimeOfArchival *common.SDKTime `presentIn:"header" name:"time-of-archival"`

	// VersionId of the requested Awr Hub object.
	VersionId *string `presentIn:"header" name:"version-id"`

	// Flag to indicate whether or not the object was modified.  If this is true,
	// the getter for the object itself will return null.  Callers should check this
	// if they specified one of the request params that might result in a conditional
	// response (like 'if-match'/'if-none-match').
	IsNotModified bool
}

func (response HeadAwrHubObjectResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response HeadAwrHubObjectResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// HeadAwrHubObjectStorageTierEnum Enum with underlying type: string
type HeadAwrHubObjectStorageTierEnum string

// Set of constants representing the allowable values for HeadAwrHubObjectStorageTierEnum
const (
	HeadAwrHubObjectStorageTierStandard         HeadAwrHubObjectStorageTierEnum = "STANDARD"
	HeadAwrHubObjectStorageTierInfrequentaccess HeadAwrHubObjectStorageTierEnum = "INFREQUENTACCESS"
	HeadAwrHubObjectStorageTierArchive          HeadAwrHubObjectStorageTierEnum = "ARCHIVE"
)

var mappingHeadAwrHubObjectStorageTierEnum = map[string]HeadAwrHubObjectStorageTierEnum{
	"STANDARD":         HeadAwrHubObjectStorageTierStandard,
	"INFREQUENTACCESS": HeadAwrHubObjectStorageTierInfrequentaccess,
	"ARCHIVE":          HeadAwrHubObjectStorageTierArchive,
}

var mappingHeadAwrHubObjectStorageTierEnumLowerCase = map[string]HeadAwrHubObjectStorageTierEnum{
	"standard":         HeadAwrHubObjectStorageTierStandard,
	"infrequentaccess": HeadAwrHubObjectStorageTierInfrequentaccess,
	"archive":          HeadAwrHubObjectStorageTierArchive,
}

// GetHeadAwrHubObjectStorageTierEnumValues Enumerates the set of values for HeadAwrHubObjectStorageTierEnum
func GetHeadAwrHubObjectStorageTierEnumValues() []HeadAwrHubObjectStorageTierEnum {
	values := make([]HeadAwrHubObjectStorageTierEnum, 0)
	for _, v := range mappingHeadAwrHubObjectStorageTierEnum {
		values = append(values, v)
	}
	return values
}

// GetHeadAwrHubObjectStorageTierEnumStringValues Enumerates the set of values in String for HeadAwrHubObjectStorageTierEnum
func GetHeadAwrHubObjectStorageTierEnumStringValues() []string {
	return []string{
		"STANDARD",
		"INFREQUENTACCESS",
		"ARCHIVE",
	}
}

// GetMappingHeadAwrHubObjectStorageTierEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHeadAwrHubObjectStorageTierEnum(val string) (HeadAwrHubObjectStorageTierEnum, bool) {
	enum, ok := mappingHeadAwrHubObjectStorageTierEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// HeadAwrHubObjectArchivalStateEnum Enum with underlying type: string
type HeadAwrHubObjectArchivalStateEnum string

// Set of constants representing the allowable values for HeadAwrHubObjectArchivalStateEnum
const (
	HeadAwrHubObjectArchivalStateArchived  HeadAwrHubObjectArchivalStateEnum = "ARCHIVED"
	HeadAwrHubObjectArchivalStateRestoring HeadAwrHubObjectArchivalStateEnum = "RESTORING"
	HeadAwrHubObjectArchivalStateRestored  HeadAwrHubObjectArchivalStateEnum = "RESTORED"
)

var mappingHeadAwrHubObjectArchivalStateEnum = map[string]HeadAwrHubObjectArchivalStateEnum{
	"ARCHIVED":  HeadAwrHubObjectArchivalStateArchived,
	"RESTORING": HeadAwrHubObjectArchivalStateRestoring,
	"RESTORED":  HeadAwrHubObjectArchivalStateRestored,
}

var mappingHeadAwrHubObjectArchivalStateEnumLowerCase = map[string]HeadAwrHubObjectArchivalStateEnum{
	"archived":  HeadAwrHubObjectArchivalStateArchived,
	"restoring": HeadAwrHubObjectArchivalStateRestoring,
	"restored":  HeadAwrHubObjectArchivalStateRestored,
}

// GetHeadAwrHubObjectArchivalStateEnumValues Enumerates the set of values for HeadAwrHubObjectArchivalStateEnum
func GetHeadAwrHubObjectArchivalStateEnumValues() []HeadAwrHubObjectArchivalStateEnum {
	values := make([]HeadAwrHubObjectArchivalStateEnum, 0)
	for _, v := range mappingHeadAwrHubObjectArchivalStateEnum {
		values = append(values, v)
	}
	return values
}

// GetHeadAwrHubObjectArchivalStateEnumStringValues Enumerates the set of values in String for HeadAwrHubObjectArchivalStateEnum
func GetHeadAwrHubObjectArchivalStateEnumStringValues() []string {
	return []string{
		"ARCHIVED",
		"RESTORING",
		"RESTORED",
	}
}

// GetMappingHeadAwrHubObjectArchivalStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHeadAwrHubObjectArchivalStateEnum(val string) (HeadAwrHubObjectArchivalStateEnum, bool) {
	enum, ok := mappingHeadAwrHubObjectArchivalStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
