// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// UpdateJavaLicenseAcceptanceRecordRequest wrapper for the UpdateJavaLicenseAcceptanceRecord operation
type UpdateJavaLicenseAcceptanceRecordRequest struct {

	// Unique Java license acceptance record identifier.
	JavaLicenseAcceptanceRecordId *string `mandatory:"true" contributesTo:"path" name:"javaLicenseAcceptanceRecordId"`

	// Attributes for updating the JavaLicenseAcceptanceRecord.
	UpdateJavaLicenseAcceptanceRecordDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// ETag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the ETag you
	// provide matches the resource's current ETag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateJavaLicenseAcceptanceRecordRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateJavaLicenseAcceptanceRecordRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateJavaLicenseAcceptanceRecordRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateJavaLicenseAcceptanceRecordRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateJavaLicenseAcceptanceRecordRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateJavaLicenseAcceptanceRecordResponse wrapper for the UpdateJavaLicenseAcceptanceRecord operation
type UpdateJavaLicenseAcceptanceRecordResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The JavaLicenseAcceptanceRecord instance
	JavaLicenseAcceptanceRecord `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response UpdateJavaLicenseAcceptanceRecordResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateJavaLicenseAcceptanceRecordResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
