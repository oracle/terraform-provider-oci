// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ExtractStructuredLogFieldPathsRequest wrapper for the ExtractStructuredLogFieldPaths operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ExtractStructuredLogFieldPaths.go.html to see an example of how to use ExtractStructuredLogFieldPathsRequest.
type ExtractStructuredLogFieldPathsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// parser definition
	LoganParserDetails LogAnalyticsParser `contributesTo:"body"`

	// The parser type - possible values are XML, JSON or DELIMITED.
	ParserType ExtractStructuredLogFieldPathsParserTypeEnum `mandatory:"false" contributesTo:"query" name:"parserType" omitEmpty:"true"`

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

func (request ExtractStructuredLogFieldPathsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ExtractStructuredLogFieldPathsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ExtractStructuredLogFieldPathsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ExtractStructuredLogFieldPathsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ExtractStructuredLogFieldPathsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExtractStructuredLogFieldPathsParserTypeEnum(string(request.ParserType)); !ok && request.ParserType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ParserType: %s. Supported values are: %s.", request.ParserType, strings.Join(GetExtractStructuredLogFieldPathsParserTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtractStructuredLogFieldPathsResponse wrapper for the ExtractStructuredLogFieldPaths operation
type ExtractStructuredLogFieldPathsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ExtractLogFieldResults instance
	ExtractLogFieldResults `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ExtractStructuredLogFieldPathsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ExtractStructuredLogFieldPathsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ExtractStructuredLogFieldPathsParserTypeEnum Enum with underlying type: string
type ExtractStructuredLogFieldPathsParserTypeEnum string

// Set of constants representing the allowable values for ExtractStructuredLogFieldPathsParserTypeEnum
const (
	ExtractStructuredLogFieldPathsParserTypeXml       ExtractStructuredLogFieldPathsParserTypeEnum = "XML"
	ExtractStructuredLogFieldPathsParserTypeJson      ExtractStructuredLogFieldPathsParserTypeEnum = "JSON"
	ExtractStructuredLogFieldPathsParserTypeDelimited ExtractStructuredLogFieldPathsParserTypeEnum = "DELIMITED"
)

var mappingExtractStructuredLogFieldPathsParserTypeEnum = map[string]ExtractStructuredLogFieldPathsParserTypeEnum{
	"XML":       ExtractStructuredLogFieldPathsParserTypeXml,
	"JSON":      ExtractStructuredLogFieldPathsParserTypeJson,
	"DELIMITED": ExtractStructuredLogFieldPathsParserTypeDelimited,
}

// GetExtractStructuredLogFieldPathsParserTypeEnumValues Enumerates the set of values for ExtractStructuredLogFieldPathsParserTypeEnum
func GetExtractStructuredLogFieldPathsParserTypeEnumValues() []ExtractStructuredLogFieldPathsParserTypeEnum {
	values := make([]ExtractStructuredLogFieldPathsParserTypeEnum, 0)
	for _, v := range mappingExtractStructuredLogFieldPathsParserTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetExtractStructuredLogFieldPathsParserTypeEnumStringValues Enumerates the set of values in String for ExtractStructuredLogFieldPathsParserTypeEnum
func GetExtractStructuredLogFieldPathsParserTypeEnumStringValues() []string {
	return []string{
		"XML",
		"JSON",
		"DELIMITED",
	}
}

// GetMappingExtractStructuredLogFieldPathsParserTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtractStructuredLogFieldPathsParserTypeEnum(val string) (ExtractStructuredLogFieldPathsParserTypeEnum, bool) {
	mappingExtractStructuredLogFieldPathsParserTypeEnumIgnoreCase := make(map[string]ExtractStructuredLogFieldPathsParserTypeEnum)
	for k, v := range mappingExtractStructuredLogFieldPathsParserTypeEnum {
		mappingExtractStructuredLogFieldPathsParserTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingExtractStructuredLogFieldPathsParserTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
