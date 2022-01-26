// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// GetQueryResultRequest wrapper for the GetQueryResult operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/GetQueryResult.go.html to see an example of how to use GetQueryResultRequest.
type GetQueryResultRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Work Request Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the asynchronous request.
	WorkRequestId *string `mandatory:"true" contributesTo:"query" name:"workRequestId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Maximum number of results to return in this request.  Note a limit=-1 returns all results from pageId onwards up to maxtotalCount.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Include columns in response
	ShouldIncludeColumns *bool `mandatory:"false" contributesTo:"query" name:"shouldIncludeColumns"`

	// Include fields in response
	ShouldIncludeFields *bool `mandatory:"false" contributesTo:"query" name:"shouldIncludeFields"`

	// Specifies the format for the returned results.
	OutputMode GetQueryResultOutputModeEnum `mandatory:"false" contributesTo:"query" name:"outputMode" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetQueryResultRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetQueryResultRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetQueryResultRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetQueryResultRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetQueryResultResponse wrapper for the GetQueryResult operation
type GetQueryResultResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of QueryAggregation instances
	QueryAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// A decimal number representing the number of seconds the client should wait before polling this endpoint again.
	RetryAfter *float32 `presentIn:"header" name:"retry-after"`
}

func (response GetQueryResultResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetQueryResultResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetQueryResultOutputModeEnum Enum with underlying type: string
type GetQueryResultOutputModeEnum string

// Set of constants representing the allowable values for GetQueryResultOutputModeEnum
const (
	GetQueryResultOutputModeJsonRows GetQueryResultOutputModeEnum = "JSON_ROWS"
)

var mappingGetQueryResultOutputMode = map[string]GetQueryResultOutputModeEnum{
	"JSON_ROWS": GetQueryResultOutputModeJsonRows,
}

// GetGetQueryResultOutputModeEnumValues Enumerates the set of values for GetQueryResultOutputModeEnum
func GetGetQueryResultOutputModeEnumValues() []GetQueryResultOutputModeEnum {
	values := make([]GetQueryResultOutputModeEnum, 0)
	for _, v := range mappingGetQueryResultOutputMode {
		values = append(values, v)
	}
	return values
}
