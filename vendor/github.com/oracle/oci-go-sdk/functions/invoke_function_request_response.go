// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package functions

import (
	"github.com/oracle/oci-go-sdk/common"
	"io"
	"net/http"
)

// InvokeFunctionRequest wrapper for the InvokeFunction operation
type InvokeFunctionRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this function.
	FunctionId *string `mandatory:"true" contributesTo:"path" name:"functionId"`

	// The body of the function invocation.
	// Note: The maximum size of the request is limited. This limit is currently 6MB and the endpoint will not accept requests that are bigger than this limit.
	InvokeFunctionBody io.ReadCloser `mandatory:"false" contributesTo:"body" encoding:"binary"`

	// An optional intent header that indicates to the FDK the way the event should be interpreted. E.g. 'httprequest', 'cloudevent'.
	FnIntent InvokeFunctionFnIntentEnum `mandatory:"false" contributesTo:"header" name:"fn-intent"`

	// Indicates whether the functions platform should execute the request directly and return the result ('sync') or
	// whether the platform should enqueue the request for later processing and acknowledge that it has been processed ('detached').
	FnInvokeType InvokeFunctionFnInvokeTypeEnum `mandatory:"false" contributesTo:"header" name:"fn-invoke-type"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request InvokeFunctionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request InvokeFunctionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request InvokeFunctionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// InvokeFunctionResponse wrapper for the InvokeFunction operation
type InvokeFunctionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response InvokeFunctionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response InvokeFunctionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// InvokeFunctionFnIntentEnum Enum with underlying type: string
type InvokeFunctionFnIntentEnum string

// Set of constants representing the allowable values for InvokeFunctionFnIntentEnum
const (
	InvokeFunctionFnIntentHttprequest InvokeFunctionFnIntentEnum = "httprequest"
	InvokeFunctionFnIntentCloudevent  InvokeFunctionFnIntentEnum = "cloudevent"
)

var mappingInvokeFunctionFnIntent = map[string]InvokeFunctionFnIntentEnum{
	"httprequest": InvokeFunctionFnIntentHttprequest,
	"cloudevent":  InvokeFunctionFnIntentCloudevent,
}

// GetInvokeFunctionFnIntentEnumValues Enumerates the set of values for InvokeFunctionFnIntentEnum
func GetInvokeFunctionFnIntentEnumValues() []InvokeFunctionFnIntentEnum {
	values := make([]InvokeFunctionFnIntentEnum, 0)
	for _, v := range mappingInvokeFunctionFnIntent {
		values = append(values, v)
	}
	return values
}

// InvokeFunctionFnInvokeTypeEnum Enum with underlying type: string
type InvokeFunctionFnInvokeTypeEnum string

// Set of constants representing the allowable values for InvokeFunctionFnInvokeTypeEnum
const (
	InvokeFunctionFnInvokeTypeDetached InvokeFunctionFnInvokeTypeEnum = "detached"
	InvokeFunctionFnInvokeTypeSync     InvokeFunctionFnInvokeTypeEnum = "sync"
)

var mappingInvokeFunctionFnInvokeType = map[string]InvokeFunctionFnInvokeTypeEnum{
	"detached": InvokeFunctionFnInvokeTypeDetached,
	"sync":     InvokeFunctionFnInvokeTypeSync,
}

// GetInvokeFunctionFnInvokeTypeEnumValues Enumerates the set of values for InvokeFunctionFnInvokeTypeEnum
func GetInvokeFunctionFnInvokeTypeEnumValues() []InvokeFunctionFnInvokeTypeEnum {
	values := make([]InvokeFunctionFnInvokeTypeEnum, 0)
	for _, v := range mappingInvokeFunctionFnInvokeType {
		values = append(values, v)
	}
	return values
}
