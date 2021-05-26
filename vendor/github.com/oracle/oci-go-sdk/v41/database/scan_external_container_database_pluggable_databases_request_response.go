// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// ScanExternalContainerDatabasePluggableDatabasesRequest wrapper for the ScanExternalContainerDatabasePluggableDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ScanExternalContainerDatabasePluggableDatabases.go.html to see an example of how to use ScanExternalContainerDatabasePluggableDatabasesRequest.
type ScanExternalContainerDatabasePluggableDatabasesRequest struct {

	// The ExternalContainerDatabase OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ExternalContainerDatabaseId *string `mandatory:"true" contributesTo:"path" name:"externalContainerDatabaseId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// external database connector resource (`ExternalDatabaseConnectorId`).
	ExternalDatabaseConnectorId *string `mandatory:"true" contributesTo:"query" name:"externalDatabaseConnectorId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ScanExternalContainerDatabasePluggableDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ScanExternalContainerDatabasePluggableDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ScanExternalContainerDatabasePluggableDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ScanExternalContainerDatabasePluggableDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ScanExternalContainerDatabasePluggableDatabasesResponse wrapper for the ScanExternalContainerDatabasePluggableDatabases operation
type ScanExternalContainerDatabasePluggableDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the work request. Multiple OCID values are returned in a comma-separated list. Use GetWorkRequest with a work request OCID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ScanExternalContainerDatabasePluggableDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ScanExternalContainerDatabasePluggableDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
