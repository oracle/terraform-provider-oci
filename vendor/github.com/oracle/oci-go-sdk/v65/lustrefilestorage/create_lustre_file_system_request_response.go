// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// CreateLustreFileSystemRequest wrapper for the CreateLustreFileSystem operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/lustrefilestorage/CreateLustreFileSystem.go.html to see an example of how to use CreateLustreFileSystemRequest.
type CreateLustreFileSystemRequest struct {

	// Details for the new Lustre file system.
	CreateLustreFileSystemDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of running that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and removed from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateLustreFileSystemRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateLustreFileSystemRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateLustreFileSystemRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateLustreFileSystemRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateLustreFileSystemRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateLustreFileSystemResponse wrapper for the CreateLustreFileSystem operation
type CreateLustreFileSystemResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LustreFileSystem instance
	LustreFileSystem `presentIn:"body"`

	// URL for the created Lustre file system. The lustreFileSystem OCID is generated after this request is sent.
	Location *string `presentIn:"header" name:"location"`

	// Same as location.
	ContentLocation *string `presentIn:"header" name:"content-location"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	// Use GetWorkRequest with this ID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateLustreFileSystemResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateLustreFileSystemResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
