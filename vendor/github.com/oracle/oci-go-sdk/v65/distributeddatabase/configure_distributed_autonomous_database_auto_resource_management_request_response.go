// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ConfigureDistributedAutonomousDatabaseAutoResourceManagementRequest wrapper for the ConfigureDistributedAutonomousDatabaseAutoResourceManagement operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/distributeddatabase/ConfigureDistributedAutonomousDatabaseAutoResourceManagement.go.html to see an example of how to use ConfigureDistributedAutonomousDatabaseAutoResourceManagementRequest.
type ConfigureDistributedAutonomousDatabaseAutoResourceManagementRequest struct {

	// Globally distributed autonomous database identifier.
	DistributedAutonomousDatabaseId *string `mandatory:"true" contributesTo:"path" name:"distributedAutonomousDatabaseId"`

	// The payload for configuring autoResourceManagement for autonomous distributed database.
	ConfigureDistributedAutonomousDatabaseAutoResourceManagementDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ConfigureDistributedAutonomousDatabaseAutoResourceManagementRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ConfigureDistributedAutonomousDatabaseAutoResourceManagementRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ConfigureDistributedAutonomousDatabaseAutoResourceManagementRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ConfigureDistributedAutonomousDatabaseAutoResourceManagementRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ConfigureDistributedAutonomousDatabaseAutoResourceManagementRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConfigureDistributedAutonomousDatabaseAutoResourceManagementResponse wrapper for the ConfigureDistributedAutonomousDatabaseAutoResourceManagement operation
type ConfigureDistributedAutonomousDatabaseAutoResourceManagementResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DistributedAutonomousDatabaseAutoResourceManagementDetails instance
	DistributedAutonomousDatabaseAutoResourceManagementDetails `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response ConfigureDistributedAutonomousDatabaseAutoResourceManagementResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ConfigureDistributedAutonomousDatabaseAutoResourceManagementResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
