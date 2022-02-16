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

// UpdateLookupDataRequest wrapper for the UpdateLookupData operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLookupData.go.html to see an example of how to use UpdateLookupDataRequest.
type UpdateLookupDataRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the lookup to operate on.
	LookupName *string `mandatory:"true" contributesTo:"path" name:"lookupName"`

	// The file to use for the lookup update.
	UpdateLookupFileBody io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// is force
	IsForce *bool `mandatory:"false" contributesTo:"query" name:"isForce"`

	// The character encoding of the uploaded file.
	CharEncoding *string `mandatory:"false" contributesTo:"query" name:"charEncoding"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A value of `100-continue` requests preliminary verification of the request method, path, and headers before the request body is sent.
	// If no error results from such verification, the server will send a 100 (Continue) interim response to indicate readiness for the request body.
	// The only allowed value for this parameter is "100-Continue" (case-insensitive).
	Expect *string `mandatory:"false" contributesTo:"header" name:"expect"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateLookupDataRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateLookupDataRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateLookupDataRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.UpdateLookupFileBody)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateLookupDataRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateLookupDataRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateLookupDataResponse wrapper for the UpdateLookupData operation
type UpdateLookupDataResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the asynchronous request. You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateLookupDataResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateLookupDataResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
