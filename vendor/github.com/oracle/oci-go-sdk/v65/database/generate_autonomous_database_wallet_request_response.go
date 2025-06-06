// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// GenerateAutonomousDatabaseWalletRequest wrapper for the GenerateAutonomousDatabaseWallet operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GenerateAutonomousDatabaseWallet.go.html to see an example of how to use GenerateAutonomousDatabaseWalletRequest.
type GenerateAutonomousDatabaseWalletRequest struct {

	// The database OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	AutonomousDatabaseId *string `mandatory:"true" contributesTo:"path" name:"autonomousDatabaseId"`

	// Request to create a new Autonomous Database wallet.
	GenerateAutonomousDatabaseWalletDetails `contributesTo:"body"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations (for example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Indicates that the request is a dry run, if set to "true". A dry run request does not actually
	// creating or updating a resource and is used only to perform validation on the submitted data.
	OpcDryRun *bool `mandatory:"false" contributesTo:"header" name:"opc-dry-run"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GenerateAutonomousDatabaseWalletRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GenerateAutonomousDatabaseWalletRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GenerateAutonomousDatabaseWalletRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GenerateAutonomousDatabaseWalletRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GenerateAutonomousDatabaseWalletRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GenerateAutonomousDatabaseWalletResponse wrapper for the GenerateAutonomousDatabaseWallet operation
type GenerateAutonomousDatabaseWalletResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Size of the file.
	ContentLength *int64 `presentIn:"header" name:"content-length"`

	// The date and time the wallet was created, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`
}

func (response GenerateAutonomousDatabaseWalletResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GenerateAutonomousDatabaseWalletResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
