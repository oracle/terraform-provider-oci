// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// TestParserRequest wrapper for the TestParser operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/TestParser.go.html to see an example of how to use TestParserRequest.
type TestParserRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Details for test payload
	TestParserPayloadDetails `contributesTo:"body"`

	// The scope used when testing a parser.
	Scope TestParserScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// The module to test.  A value of 'ParserFunctionTest' will result in testing of
	// the parser functions.
	ReqOriginModule *string `mandatory:"false" contributesTo:"query" name:"reqOriginModule"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request TestParserRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request TestParserRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request TestParserRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request TestParserRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// TestParserResponse wrapper for the TestParser operation
type TestParserResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ParserTestResult instance
	ParserTestResult `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response TestParserResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response TestParserResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// TestParserScopeEnum Enum with underlying type: string
type TestParserScopeEnum string

// Set of constants representing the allowable values for TestParserScopeEnum
const (
	TestParserScopeLines           TestParserScopeEnum = "LOG_LINES"
	TestParserScopeEntries         TestParserScopeEnum = "LOG_ENTRIES"
	TestParserScopeLinesLogEntries TestParserScopeEnum = "LOG_LINES_LOG_ENTRIES"
)

var mappingTestParserScope = map[string]TestParserScopeEnum{
	"LOG_LINES":             TestParserScopeLines,
	"LOG_ENTRIES":           TestParserScopeEntries,
	"LOG_LINES_LOG_ENTRIES": TestParserScopeLinesLogEntries,
}

// GetTestParserScopeEnumValues Enumerates the set of values for TestParserScopeEnum
func GetTestParserScopeEnumValues() []TestParserScopeEnum {
	values := make([]TestParserScopeEnum, 0)
	for _, v := range mappingTestParserScope {
		values = append(values, v)
	}
	return values
}
