// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// FailoverAutonomousContainerDatabaseDataguardRequest wrapper for the FailoverAutonomousContainerDatabaseDataguard operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/FailoverAutonomousContainerDatabaseDataguard.go.html to see an example of how to use FailoverAutonomousContainerDatabaseDataguardRequest.
type FailoverAutonomousContainerDatabaseDataguardRequest struct {

	// The Autonomous Container Database OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"true" contributesTo:"path" name:"autonomousContainerDatabaseId"`

	// For optimistic concurrency control. In the PUT or DELETE call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous GET or POST response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request FailoverAutonomousContainerDatabaseDataguardRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request FailoverAutonomousContainerDatabaseDataguardRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request FailoverAutonomousContainerDatabaseDataguardRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request FailoverAutonomousContainerDatabaseDataguardRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request FailoverAutonomousContainerDatabaseDataguardRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FailoverAutonomousContainerDatabaseDataguardResponse wrapper for the FailoverAutonomousContainerDatabaseDataguard operation
type FailoverAutonomousContainerDatabaseDataguardResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AutonomousContainerDatabase instance
	AutonomousContainerDatabase `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you must contact Oracle about
	// a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the work request. Multiple OCID values are returned in a comma-separated list. Use GetWorkRequest with a work request OCID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response FailoverAutonomousContainerDatabaseDataguardResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response FailoverAutonomousContainerDatabaseDataguardResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
