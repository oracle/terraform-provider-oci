// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// InvokeFunctionRequest wrapper for the InvokeFunction operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/InvokeFunction.go.html to see an example of how to use InvokeFunctionRequest.
type InvokeFunctionRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this function.
	FunctionId *string `mandatory:"true" contributesTo:"path" name:"functionId"`

	// The body of the function invocation.
	// Note: The maximum size of the request is limited. This limit is currently 6MB and the endpoint will not accept requests that are bigger than this limit.
	InvokeFunctionBody io.ReadCloser `mandatory:"false" contributesTo:"body" encoding:"binary"`

	// An optional intent header that indicates to the FDK the way the event should be interpreted. E.g. 'httprequest', 'cloudevent'.
	FnIntent InvokeFunctionFnIntentEnum `mandatory:"false" contributesTo:"header" name:"fn-intent"`

	// Indicates whether Oracle Functions should execute the request and return the result ('sync') of the execution,
	// or whether Oracle Functions should return as soon as processing has begun ('detached') and leave result handling to the function.
	FnInvokeType InvokeFunctionFnInvokeTypeEnum `mandatory:"false" contributesTo:"header" name:"fn-invoke-type"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Indicates that the request is a dry run, if set to "true". A dry run request does not execute the function.
	IsDryRun *bool `mandatory:"false" contributesTo:"header" name:"is-dry-run"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request InvokeFunctionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request InvokeFunctionRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request InvokeFunctionRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.InvokeFunctionBody)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request InvokeFunctionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request InvokeFunctionRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInvokeFunctionFnIntentEnum(string(request.FnIntent)); !ok && request.FnIntent != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FnIntent: %s. Supported values are: %s.", request.FnIntent, strings.Join(GetInvokeFunctionFnIntentEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInvokeFunctionFnInvokeTypeEnum(string(request.FnInvokeType)); !ok && request.FnInvokeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FnInvokeType: %s. Supported values are: %s.", request.FnInvokeType, strings.Join(GetInvokeFunctionFnInvokeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingInvokeFunctionFnIntentEnum = map[string]InvokeFunctionFnIntentEnum{
	"httprequest": InvokeFunctionFnIntentHttprequest,
	"cloudevent":  InvokeFunctionFnIntentCloudevent,
}

var mappingInvokeFunctionFnIntentEnumLowerCase = map[string]InvokeFunctionFnIntentEnum{
	"httprequest": InvokeFunctionFnIntentHttprequest,
	"cloudevent":  InvokeFunctionFnIntentCloudevent,
}

// GetInvokeFunctionFnIntentEnumValues Enumerates the set of values for InvokeFunctionFnIntentEnum
func GetInvokeFunctionFnIntentEnumValues() []InvokeFunctionFnIntentEnum {
	values := make([]InvokeFunctionFnIntentEnum, 0)
	for _, v := range mappingInvokeFunctionFnIntentEnum {
		values = append(values, v)
	}
	return values
}

// GetInvokeFunctionFnIntentEnumStringValues Enumerates the set of values in String for InvokeFunctionFnIntentEnum
func GetInvokeFunctionFnIntentEnumStringValues() []string {
	return []string{
		"httprequest",
		"cloudevent",
	}
}

// GetMappingInvokeFunctionFnIntentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInvokeFunctionFnIntentEnum(val string) (InvokeFunctionFnIntentEnum, bool) {
	enum, ok := mappingInvokeFunctionFnIntentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InvokeFunctionFnInvokeTypeEnum Enum with underlying type: string
type InvokeFunctionFnInvokeTypeEnum string

// Set of constants representing the allowable values for InvokeFunctionFnInvokeTypeEnum
const (
	InvokeFunctionFnInvokeTypeDetached InvokeFunctionFnInvokeTypeEnum = "detached"
	InvokeFunctionFnInvokeTypeSync     InvokeFunctionFnInvokeTypeEnum = "sync"
)

var mappingInvokeFunctionFnInvokeTypeEnum = map[string]InvokeFunctionFnInvokeTypeEnum{
	"detached": InvokeFunctionFnInvokeTypeDetached,
	"sync":     InvokeFunctionFnInvokeTypeSync,
}

var mappingInvokeFunctionFnInvokeTypeEnumLowerCase = map[string]InvokeFunctionFnInvokeTypeEnum{
	"detached": InvokeFunctionFnInvokeTypeDetached,
	"sync":     InvokeFunctionFnInvokeTypeSync,
}

// GetInvokeFunctionFnInvokeTypeEnumValues Enumerates the set of values for InvokeFunctionFnInvokeTypeEnum
func GetInvokeFunctionFnInvokeTypeEnumValues() []InvokeFunctionFnInvokeTypeEnum {
	values := make([]InvokeFunctionFnInvokeTypeEnum, 0)
	for _, v := range mappingInvokeFunctionFnInvokeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInvokeFunctionFnInvokeTypeEnumStringValues Enumerates the set of values in String for InvokeFunctionFnInvokeTypeEnum
func GetInvokeFunctionFnInvokeTypeEnumStringValues() []string {
	return []string{
		"detached",
		"sync",
	}
}

// GetMappingInvokeFunctionFnInvokeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInvokeFunctionFnInvokeTypeEnum(val string) (InvokeFunctionFnInvokeTypeEnum, bool) {
	enum, ok := mappingInvokeFunctionFnInvokeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
