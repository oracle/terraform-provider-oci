// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest wrapper for the UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec.go.html to see an example of how to use UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest.
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
	DatabaseToolsDatabaseApiGatewayConfigId *string `mandatory:"true" contributesTo:"path" name:"databaseToolsDatabaseApiGatewayConfigId"`

	// The key of the pool config.
	PoolKey *string `mandatory:"true" contributesTo:"path" name:"poolKey"`

	// The key of the auto API spec config.
	AutoApiSpecKey *string `mandatory:"true" contributesTo:"path" name:"autoApiSpecKey"`

	// The information to be updated.
	UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecDetails `contributesTo:"body"`

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

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse wrapper for the UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec operation
type UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec instance
	DatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpec `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateDatabaseToolsDatabaseApiGatewayConfigPoolAutoApiSpecResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
