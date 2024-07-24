// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetShardedDatabaseRequest wrapper for the GetShardedDatabase operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/globallydistributeddatabase/GetShardedDatabase.go.html to see an example of how to use GetShardedDatabaseRequest.
type GetShardedDatabaseRequest struct {

	// Sharded Database identifier
	ShardedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"shardedDatabaseId"`

	// Comma separated names of argument corresponding to which metadata need to be retrived, namely VM_CLUSTER_INFO, ADDITIONAL_RESOURCE_INFO.
	// An example is metadata=VM_CLUSTER_INFO,ADDITIONAL_RESOURCE_INFO.
	Metadata *string `mandatory:"false" contributesTo:"query" name:"metadata"`

	// For conditional requests. In the GET call for a resource, set the
	// `If-None-Match` header to the value of the ETag from a previous GET (or
	// POST or PUT) response for that resource. The server will return with
	// either a 304 Not Modified response if the resource has not changed, or a
	// 200 OK response with the updated representation.
	IfNoneMatch *string `mandatory:"false" contributesTo:"header" name:"if-none-match"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetShardedDatabaseRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetShardedDatabaseRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetShardedDatabaseRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetShardedDatabaseRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetShardedDatabaseRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetShardedDatabaseResponse wrapper for the GetShardedDatabase operation
type GetShardedDatabaseResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ShardedDatabase instance
	ShardedDatabase `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetShardedDatabaseResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetShardedDatabaseResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
