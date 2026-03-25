// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutonomousDatabasesInAutonomousContainerDatabaseBackupRequest wrapper for the ListAutonomousDatabasesInAutonomousContainerDatabaseBackup operation
type ListAutonomousDatabasesInAutonomousContainerDatabaseBackupRequest struct {

	// The Autonomous Container Database OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousContainerDatabaseId *string `mandatory:"true" contributesTo:"query" name:"autonomousContainerDatabaseId"`

	// The timestamp at which the list of autonomous databases present to be returned (ISO-8601 format, e.g., "2025-12-02T06:39:15Z"). The requested value must be after the end timestamp of the oldest available Autonomous Container Database backup.
	TimeStampRequested *common.SDKTime `mandatory:"true" contributesTo:"query" name:"timeStampRequested"`

	// The compartment OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm). If not provided, uses the Autonomous Container Database's compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousDatabasesInAutonomousContainerDatabaseBackupRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousDatabasesInAutonomousContainerDatabaseBackupRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousDatabasesInAutonomousContainerDatabaseBackupRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDatabasesInAutonomousContainerDatabaseBackupRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousDatabasesInAutonomousContainerDatabaseBackupRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousDatabasesInAutonomousContainerDatabaseBackupResponse wrapper for the ListAutonomousDatabasesInAutonomousContainerDatabaseBackup operation
type ListAutonomousDatabasesInAutonomousContainerDatabaseBackupResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AutonomousDatabaseInBackupCollection instances
	AutonomousDatabaseInBackupCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then there are additional items still to get. Include this value as the `page` parameter for the subsequent GET request. For information about pagination, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousDatabasesInAutonomousContainerDatabaseBackupResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousDatabasesInAutonomousContainerDatabaseBackupResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
