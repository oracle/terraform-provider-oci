// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmconfig

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// HeadDataFileRequest wrapper for the HeadDataFile operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmconfig/HeadDataFile.go.html to see an example of how to use HeadDataFileRequest.
type HeadDataFileRequest struct {

	// The name of the data file.
	DataFileName *string `mandatory:"true" contributesTo:"path" name:"dataFileName"`

	// The APM Domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// The type of the data file.
	ApmType *string `mandatory:"true" contributesTo:"query" name:"apmType"`

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

func (request HeadDataFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request HeadDataFileRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request HeadDataFileRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request HeadDataFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request HeadDataFileRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HeadDataFileResponse wrapper for the HeadDataFile operation
type HeadDataFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The client request ID.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Content-Type header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.17).
	ContentType *string `presentIn:"header" name:"content-type"`

	// The object size in bytes.
	ContentLength *int64 `presentIn:"header" name:"content-length"`

	// The base-64 encoded MD5 hash of the request body as computed by the server.
	ContentMd5 *string `presentIn:"header" name:"content-md5"`

	// Content-Encoding header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.11).
	ContentEncoding *string `presentIn:"header" name:"content-encoding"`

	// Content-Language header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.12).
	ContentLanguage *string `presentIn:"header" name:"content-language"`

	// Content-Disposition header, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-19.5.1).
	ContentDisposition *string `presentIn:"header" name:"content-disposition"`

	// The last time the object was modified, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29).
	// Expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2020-02-19T22:47:12.613Z`
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`

	// A string containing a JSON-encoded object with metadata related to the uploaded file or resource.
	// Example:
	//   {"fileName":"report.pdf","uploader":"jane.doe","category":"financial"}
	Metadata *string `presentIn:"header" name:"metadata"`
}

func (response HeadDataFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response HeadDataFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
