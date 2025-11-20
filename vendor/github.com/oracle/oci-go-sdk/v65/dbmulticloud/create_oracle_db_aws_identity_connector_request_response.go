// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dbmulticloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// CreateOracleDbAwsIdentityConnectorRequest wrapper for the CreateOracleDbAwsIdentityConnector operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dbmulticloud/CreateOracleDbAwsIdentityConnector.go.html to see an example of how to use CreateOracleDbAwsIdentityConnectorRequest.
type CreateOracleDbAwsIdentityConnectorRequest struct {

	// Details for to Create Oracle DB AWS Identity Connector resource.
	CreateOracleDbAwsIdentityConnectorDetails `contributesTo:"body"`

	// A token that uniquely identifies a request, allowing it to be safely retried in the event of a timeout or server error without the risk of the action being executed more than once.
	// Retry tokens expire after 24 hours but can be invalidated sooner if conflicting operations occur.
	// For example, if a resource has been deleted and permanently purged from the system, a retry of the original creation request may be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateOracleDbAwsIdentityConnectorRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateOracleDbAwsIdentityConnectorRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateOracleDbAwsIdentityConnectorRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateOracleDbAwsIdentityConnectorRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateOracleDbAwsIdentityConnectorRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOracleDbAwsIdentityConnectorResponse wrapper for the CreateOracleDbAwsIdentityConnector operation
type CreateOracleDbAwsIdentityConnectorResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OracleDbAwsIdentityConnector instance
	OracleDbAwsIdentityConnector `presentIn:"body"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	// Use GetWorkRequest with this ID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// URI of the resource
	Location *string `presentIn:"header" name:"location"`

	// URI of the resource
	ContentLocation *string `presentIn:"header" name:"content-location"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response CreateOracleDbAwsIdentityConnectorResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateOracleDbAwsIdentityConnectorResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
