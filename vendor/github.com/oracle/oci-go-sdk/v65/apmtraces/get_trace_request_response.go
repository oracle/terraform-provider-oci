// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetTraceRequest wrapper for the GetTrace operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/GetTrace.go.html to see an example of how to use GetTraceRequest.
type GetTraceRequest struct {

	// The APM Domain ID for the intended request.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// Unique Application Performance Monitoring trace identifier (traceId).
	TraceKey *string `mandatory:"true" contributesTo:"path" name:"traceKey"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Include traces that have a `minTraceStartTime` equal to or greater than this value.
	TimeTraceStartedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeTraceStartedGreaterThanOrEqualTo"`

	// Include traces that have a `minTraceStartTime` less than this value.
	TimeTraceStartedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeTraceStartedLessThan"`

	// Name space from which the trace details need to be retrieved.
	TraceNamespace GetTraceTraceNamespaceEnum `mandatory:"false" contributesTo:"query" name:"traceNamespace" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetTraceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetTraceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetTraceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetTraceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetTraceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetTraceTraceNamespaceEnum(string(request.TraceNamespace)); !ok && request.TraceNamespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TraceNamespace: %s. Supported values are: %s.", request.TraceNamespace, strings.Join(GetGetTraceTraceNamespaceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetTraceResponse wrapper for the GetTrace operation
type GetTraceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Trace instance
	Trace `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTraceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetTraceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetTraceTraceNamespaceEnum Enum with underlying type: string
type GetTraceTraceNamespaceEnum string

// Set of constants representing the allowable values for GetTraceTraceNamespaceEnum
const (
	GetTraceTraceNamespaceTraces    GetTraceTraceNamespaceEnum = "TRACES"
	GetTraceTraceNamespaceSynthetic GetTraceTraceNamespaceEnum = "SYNTHETIC"
)

var mappingGetTraceTraceNamespaceEnum = map[string]GetTraceTraceNamespaceEnum{
	"TRACES":    GetTraceTraceNamespaceTraces,
	"SYNTHETIC": GetTraceTraceNamespaceSynthetic,
}

var mappingGetTraceTraceNamespaceEnumLowerCase = map[string]GetTraceTraceNamespaceEnum{
	"traces":    GetTraceTraceNamespaceTraces,
	"synthetic": GetTraceTraceNamespaceSynthetic,
}

// GetGetTraceTraceNamespaceEnumValues Enumerates the set of values for GetTraceTraceNamespaceEnum
func GetGetTraceTraceNamespaceEnumValues() []GetTraceTraceNamespaceEnum {
	values := make([]GetTraceTraceNamespaceEnum, 0)
	for _, v := range mappingGetTraceTraceNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetGetTraceTraceNamespaceEnumStringValues Enumerates the set of values in String for GetTraceTraceNamespaceEnum
func GetGetTraceTraceNamespaceEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYNTHETIC",
	}
}

// GetMappingGetTraceTraceNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetTraceTraceNamespaceEnum(val string) (GetTraceTraceNamespaceEnum, bool) {
	enum, ok := mappingGetTraceTraceNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
