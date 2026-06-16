// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmconfig

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// PutDataFileRequest wrapper for the PutDataFile operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/PutDataFile.go.html to see an example of how to use PutDataFileRequest.
type PutDataFileRequest struct {

	// The data file to be uploaded.
	PutDataFileBody io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// The name of the data file.
	DataFileName *string `mandatory:"true" contributesTo:"path" name:"dataFileName"`

	// The APM Domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// The type of the data file.
	ApmType *string `mandatory:"true" contributesTo:"query" name:"apmType"`

	// Optional base64-encoded MD5 hash of the request body. If provided, the server will perform
	// a data integrity check by computing the MD5 of the received content and comparing it to the
	// supplied value.
	// If the values do not match, the request will be rejected with an HTTP 400 error and a message such as:
	// "The computed MD5 of the request body (ACTUAL_MD5) does not match the Content-MD5 header (HEADER_MD5)"
	ContentMD5 *string `mandatory:"false" contributesTo:"header" name:"Content-MD5"`

	// Optional parameter specifying the media type (MIME type) of the request or response body.
	// If not specified, the default is `application/octet-stream`.
	// This value can be used by recipients to determine how to interpret or render the content.
	ContentType *string `mandatory:"false" contributesTo:"header" name:"Content-Type"`

	// Optional parameter that indicates the natural language of the content.
	// This value can be used by clients or intermediaries to select or display content based on language preferences.
	ContentLanguage *string `mandatory:"false" contributesTo:"header" name:"Content-Language"`

	// Optional parameter indicating the content encodings applied to the request body (e.g., gzip, deflate).
	// This value can be used by recipients to determine how to decode the content.
	ContentEncoding *string `mandatory:"false" contributesTo:"header" name:"Content-Encoding"`

	// Optional parameter that provides presentation information for how the content should be displayed or handled by the recipient.
	// For example, to prompt a file download with a custom filename:
	// `attachment; filename="example.txt"`
	ContentDisposition *string `mandatory:"false" contributesTo:"header" name:"Content-Disposition"`

	// A string containing a JSON-encoded object with metadata related to the uploaded file or resource.
	// Example:
	//   {"fileName":"report.pdf","uploader":"jane.doe","category":"financial"}
	Metadata *string `mandatory:"false" contributesTo:"header" name:"metadata"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PutDataFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PutDataFileRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request PutDataFileRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.PutDataFileBody)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PutDataFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request PutDataFileRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PutDataFileResponse wrapper for the PutDataFile operation
type PutDataFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The client request ID.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// The base-64 encoded MD5 hash of the request body as computed by the server.
	ContentMd5 *string `presentIn:"header" name:"content-md5"`

	// The last time the object was modified, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29).
	// Expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2020-02-19T22:47:12.613Z`
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`
}

func (response PutDataFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PutDataFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
