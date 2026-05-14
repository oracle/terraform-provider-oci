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

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest wrapper for the CreateDatabaseToolsDatabaseApiGatewayConfigPool operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/CreateDatabaseToolsDatabaseApiGatewayConfigPool.go.html to see an example of how to use CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest.
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
	DatabaseToolsDatabaseApiGatewayConfigId *string `mandatory:"true" contributesTo:"path" name:"databaseToolsDatabaseApiGatewayConfigId"`

	// Details for the new Database Tools database API gateway config pool resource.
	CreateDatabaseToolsDatabaseApiGatewayConfigPoolDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	// Accepted characters: ASCII alphanumerics plus underscore (U+005F LOW LINE "_") and dash (U+002D HYPHEN-MINUS "-")
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateDatabaseToolsDatabaseApiGatewayConfigPoolRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse wrapper for the CreateDatabaseToolsDatabaseApiGatewayConfigPool operation
type CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DatabaseToolsDatabaseApiGatewayConfigPool instance
	DatabaseToolsDatabaseApiGatewayConfigPool `presentIn:"body"`

	// URI of the new resource which was created by the request.
	Location *string `presentIn:"header" name:"location"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateDatabaseToolsDatabaseApiGatewayConfigPoolResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
