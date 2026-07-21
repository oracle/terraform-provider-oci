// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetRecallCountRequest wrapper for the GetRecallCount operation
type GetRecallCountRequest struct {

	// The Log Analytics namespace used for the request. The namespace can be obtained by running 'oci os ns get'
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The type of stored data.
	DataType GetRecallCountDataTypeEnum `mandatory:"false" contributesTo:"query" name:"dataType" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetRecallCountRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetRecallCountRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetRecallCountRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetRecallCountRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetRecallCountRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetRecallCountDataTypeEnum(string(request.DataType)); !ok && request.DataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataType: %s. Supported values are: %s.", request.DataType, strings.Join(GetGetRecallCountDataTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetRecallCountResponse wrapper for the GetRecallCount operation
type GetRecallCountResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The RecallCount instance
	RecallCount `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetRecallCountResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetRecallCountResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetRecallCountDataTypeEnum Enum with underlying type: string
type GetRecallCountDataTypeEnum string

// Set of constants representing the allowable values for GetRecallCountDataTypeEnum
const (
	GetRecallCountDataTypeLog    GetRecallCountDataTypeEnum = "LOG"
	GetRecallCountDataTypeLookup GetRecallCountDataTypeEnum = "LOOKUP"
	GetRecallCountDataTypeApm    GetRecallCountDataTypeEnum = "APM"
)

var mappingGetRecallCountDataTypeEnum = map[string]GetRecallCountDataTypeEnum{
	"LOG":    GetRecallCountDataTypeLog,
	"LOOKUP": GetRecallCountDataTypeLookup,
	"APM":    GetRecallCountDataTypeApm,
}

var mappingGetRecallCountDataTypeEnumLowerCase = map[string]GetRecallCountDataTypeEnum{
	"log":    GetRecallCountDataTypeLog,
	"lookup": GetRecallCountDataTypeLookup,
	"apm":    GetRecallCountDataTypeApm,
}

// GetGetRecallCountDataTypeEnumValues Enumerates the set of values for GetRecallCountDataTypeEnum
func GetGetRecallCountDataTypeEnumValues() []GetRecallCountDataTypeEnum {
	values := make([]GetRecallCountDataTypeEnum, 0)
	for _, v := range mappingGetRecallCountDataTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetRecallCountDataTypeEnumStringValues Enumerates the set of values in String for GetRecallCountDataTypeEnum
func GetGetRecallCountDataTypeEnumStringValues() []string {
	return []string{
		"LOG",
		"LOOKUP",
		"APM",
	}
}

// GetMappingGetRecallCountDataTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetRecallCountDataTypeEnum(val string) (GetRecallCountDataTypeEnum, bool) {
	enum, ok := mappingGetRecallCountDataTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
