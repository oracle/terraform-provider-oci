// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package nosql

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// SummarizeStatementRequest wrapper for the SummarizeStatement operation
type SummarizeStatementRequest struct {

	// The ID of a table's compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A NoSQL SQL statement.
	Statement *string `mandatory:"true" contributesTo:"query" name:"statement"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeStatementRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeStatementRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeStatementRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// SummarizeStatementResponse wrapper for the SummarizeStatement operation
type SummarizeStatementResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The StatementSummary instance
	StatementSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SummarizeStatementResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeStatementResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
