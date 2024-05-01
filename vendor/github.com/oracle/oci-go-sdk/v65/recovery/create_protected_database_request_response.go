// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// CreateProtectedDatabaseRequest wrapper for the CreateProtectedDatabase operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/recovery/CreateProtectedDatabase.go.html to see an example of how to use CreateProtectedDatabaseRequest.
type CreateProtectedDatabaseRequest struct {

	// Describes the parameters required to create a protected database.
	CreateProtectedDatabaseDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicates if the request is to test the preparedness for creating a protected database, without actually creating a protected database.
	// If you set the `opcDryRun` option as `true`, then Recovery Service only performs a test run to check for any missing prerequisites or configurations required to create a protected database, and then returns error messages to warn you about any missing requirements.
	// If an error occurs, you can review, correct, and repeat the dry run until the `createProtectedDatabase` request does not return any errors.
	// These are the common issues that you can identify by performing a dry run of the `createProtectedDatabase` request:
	// * The Recovery Service subnet has insufficient free IP addresses to support the required number of private endpoints. See, troubleshooting (https://docs.oracle.com/en/cloud/paas/recovery-service/dbrsu/troubleshoot-backup-failures-recovery-service.html#GUID-05FA08B8-421D-4E52-B84B-7AFB84ADECF9) information
	// * Recovery Service does not have permissions to manage the network resources in a chosen compartment
	// * Recovery Service is out of capacity. See, Service Limits (https://docs.oracle.com/en-us/iaas/Content/General/Concepts/servicelimits.htm) for more information
	// * Recovery Service resources exceed quota limits
	// * A protected database, having the same database ID, already exists
	// * The specified protection policy does not exist, or it is not in an Active state
	// * The specified Recovery Service subnet does not exist, or it is not in an Active state
	// See, Prerequisites for Using Recovery Service (https://docs.oracle.com/en/cloud/paas/recovery-service/dbrsu/backup-recover-recovery-service.html#GUID-B2ABF281-DFF8-4A4E-AC85-629801AAF36A) for more information.
	OpcDryRun *bool `mandatory:"false" contributesTo:"header" name:"opc-dry-run"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateProtectedDatabaseRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateProtectedDatabaseRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateProtectedDatabaseRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateProtectedDatabaseRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateProtectedDatabaseRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateProtectedDatabaseResponse wrapper for the CreateProtectedDatabase operation
type CreateProtectedDatabaseResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ProtectedDatabase instance
	ProtectedDatabase `presentIn:"body"`

	// Relative URL of the newly created resource.
	Location *string `presentIn:"header" name:"location"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier of the work request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response CreateProtectedDatabaseResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateProtectedDatabaseResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
