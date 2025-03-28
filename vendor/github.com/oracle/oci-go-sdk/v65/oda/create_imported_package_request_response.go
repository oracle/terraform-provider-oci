// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// CreateImportedPackageRequest wrapper for the CreateImportedPackage operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/CreateImportedPackage.go.html to see an example of how to use CreateImportedPackageRequest.
type CreateImportedPackageRequest struct {

	// Parameter values required to import the package.
	CreateImportedPackageDetails `contributesTo:"body"`

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// A token that uniquely identifies a request so that you can retry the request if there's
	// a timeout or server error without the risk of executing that same action again.
	// Retry tokens expire after 24 hours, but they can become invalid before then if there are
	// conflicting operations. For example, if an instance was deleted and purged from the system,
	// then the service might reject a retry of the original creation request.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateImportedPackageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateImportedPackageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateImportedPackageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateImportedPackageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateImportedPackageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateImportedPackageResponse wrapper for the CreateImportedPackage operation
type CreateImportedPackageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ImportedPackage instance
	ImportedPackage `presentIn:"body"`

	// Fully qualified URL for the newly created resource.
	Location *string `presentIn:"header" name:"location"`

	// For use in a PUT or DELETE `if-match` query parameter for optimistic concurrency control.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the asynchronous request. You can use this to query status
	// of the operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateImportedPackageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateImportedPackageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
