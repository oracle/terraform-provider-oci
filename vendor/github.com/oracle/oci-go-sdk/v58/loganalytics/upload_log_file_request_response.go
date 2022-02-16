// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"io"
	"net/http"
	"strings"
)

// UploadLogFileRequest wrapper for the UploadLogFile operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UploadLogFile.go.html to see an example of how to use UploadLogFileRequest.
type UploadLogFileRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the upload. This can be considered as a container name where different kind of logs will be collected and searched together. This upload name/id can further be used for retrieving the details of the upload, including its status or deleting the upload.
	UploadName *string `mandatory:"true" contributesTo:"query" name:"uploadName"`

	// Name of the log source that will be used to process the files being uploaded.
	LogSourceName *string `mandatory:"true" contributesTo:"query" name:"logSourceName"`

	// The name of the file being uploaded. The extension of the filename part will be used to detect the type of file (like zip, tar).
	Filename *string `mandatory:"true" contributesTo:"query" name:"filename"`

	// The log group OCID to which the log data in this upload will be mapped to.
	OpcMetaLoggrpid *string `mandatory:"true" contributesTo:"header" name:"opc-meta-loggrpid"`

	// Log data
	UploadLogFileBody io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// The entity OCID.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// Timezone to be used when processing log entries whose timestamps do not include an explicit timezone. When this property is not specified, the timezone of the entity specified is used. If the entity is also not specified or do not have a valid timezone then UTC is used
	Timezone *string `mandatory:"false" contributesTo:"query" name:"timezone"`

	// Character encoding to be used to detect the encoding type of file(s) being uploaded.
	// When this property is not specified, system detected character encoding will be used.
	CharEncoding *string `mandatory:"false" contributesTo:"query" name:"charEncoding"`

	// This property is used to specify the format of the date. This is to be used for ambiguous dates like 12/11/10. This property can take any of the following values -  MONTH_DAY_YEAR, DAY_MONTH_YEAR, YEAR_MONTH_DAY, MONTH_DAY, DAY_MONTH.
	DateFormat *string `mandatory:"false" contributesTo:"query" name:"dateFormat"`

	// Used to indicate the year where the log entries timestamp do not mention year (Ex: Nov  8 20:45:56).
	DateYear *string `mandatory:"false" contributesTo:"query" name:"dateYear"`

	// This property can be used to reset configuration cache in case of an issue with the upload.
	InvalidateCache *bool `mandatory:"false" contributesTo:"query" name:"invalidateCache"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The base-64 encoded MD5 hash of the body. If the Content-MD5 header is present, Logging Analytics performs an integrity check
	// on the body of the HTTP request by computing the MD5 hash for the body and comparing it to the MD5 hash supplied in the header.
	// If the two hashes do not match, the log data is rejected and an HTTP-400 Unmatched Content MD5 error is returned with the message:
	// "The computed MD5 of the request body (ACTUAL_MD5) does not match the Content-MD5 header (HEADER_MD5)"
	ContentMd5 *string `mandatory:"false" contributesTo:"header" name:"content-md5"`

	// The content type of the log data.
	ContentType *string `mandatory:"false" contributesTo:"header" name:"content-type"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The log set that gets associated with the uploaded logs.
	LogSet *string `mandatory:"false" contributesTo:"query" name:"logSet"`

	// A value of `100-continue` requests preliminary verification of the request method, path, and headers before the request body is sent.
	// If no error results from such verification, the server will send a 100 (Continue) interim response to indicate readiness for the request body.
	// The only allowed value for this parameter is "100-Continue" (case-insensitive).
	Expect *string `mandatory:"false" contributesTo:"header" name:"expect"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UploadLogFileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UploadLogFileRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request UploadLogFileRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.UploadLogFileBody)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UploadLogFileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UploadLogFileRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UploadLogFileResponse wrapper for the UploadLogFile operation
type UploadLogFileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Upload instance
	Upload `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The base-64 encoded MD5 hash of the request body as computed by the server.
	OpcContentMd5 *string `presentIn:"header" name:"opc-content-md5"`

	// Unique Oracle-assigned identifier for log data.
	OpcObjectId *string `presentIn:"header" name:"opc-object-id"`
}

func (response UploadLogFileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UploadLogFileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
