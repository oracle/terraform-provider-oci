// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetModuleStreamProfileRequest wrapper for the GetModuleStreamProfile operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/GetModuleStreamProfile.go.html to see an example of how to use GetModuleStreamProfileRequest.
type GetModuleStreamProfileRequest struct {

	// The OCID of the software source.
	SoftwareSourceId *string `mandatory:"true" contributesTo:"path" name:"softwareSourceId"`

	// The name of the module
	ModuleName *string `mandatory:"true" contributesTo:"path" name:"moduleName"`

	// The name of the stream of the containing module
	StreamName *string `mandatory:"true" contributesTo:"path" name:"streamName"`

	// The name of the profile of the containing module stream
	ProfileName *string `mandatory:"true" contributesTo:"path" name:"profileName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetModuleStreamProfileRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetModuleStreamProfileRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetModuleStreamProfileRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetModuleStreamProfileRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetModuleStreamProfileRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetModuleStreamProfileResponse wrapper for the GetModuleStreamProfile operation
type GetModuleStreamProfileResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ModuleStreamProfile instance
	ModuleStreamProfile `presentIn:"body"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query the status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetModuleStreamProfileResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetModuleStreamProfileResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
