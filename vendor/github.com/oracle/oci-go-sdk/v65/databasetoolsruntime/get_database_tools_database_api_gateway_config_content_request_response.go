// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// GetDatabaseToolsDatabaseApiGatewayConfigContentRequest wrapper for the GetDatabaseToolsDatabaseApiGatewayConfigContent operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetDatabaseToolsDatabaseApiGatewayConfigContent.go.html to see an example of how to use GetDatabaseToolsDatabaseApiGatewayConfigContentRequest.
type GetDatabaseToolsDatabaseApiGatewayConfigContentRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
	DatabaseToolsDatabaseApiGatewayConfigId *string `mandatory:"true" contributesTo:"path" name:"databaseToolsDatabaseApiGatewayConfigId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If-Match is most often used with state-changing methods (e.g., POST, PUT, DELETE) to prevent
	// accidental overwrites when multiple user agentss might be acting in parallel on the same
	// resource (i.e., to prevent the "lost update" problem). In general, it can be used with any
	// method that involves the selection or modification of a representation to abort the request
	// if the selected representation's current entity tag is not a member within the If-Match field value.
	// When specified on an action-specific subresource, the ETag value of the resource on which the
	// action is requested should be provided.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// For conditional requests. In the GET call for a resource, set the
	// `If-None-Match` header to the value of the ETag from a previous GET (or
	// POST or PUT) response for that resource. The server will return with
	// either a 304 Not Modified response if the resource has not changed, or a
	// 200 OK response with the updated representation.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDatabaseToolsDatabaseApiGatewayConfigContentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDatabaseToolsDatabaseApiGatewayConfigContentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDatabaseToolsDatabaseApiGatewayConfigContentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDatabaseToolsDatabaseApiGatewayConfigContentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetDatabaseToolsDatabaseApiGatewayConfigContentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetDatabaseToolsDatabaseApiGatewayConfigContentResponse wrapper for the GetDatabaseToolsDatabaseApiGatewayConfigContent operation
type GetDatabaseToolsDatabaseApiGatewayConfigContentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDatabaseToolsDatabaseApiGatewayConfigContentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDatabaseToolsDatabaseApiGatewayConfigContentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
