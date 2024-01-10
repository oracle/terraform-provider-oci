// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// GetJobTfPlanRequest wrapper for the GetJobTfPlan operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourcemanager/GetJobTfPlan.go.html to see an example of how to use GetJobTfPlanRequest.
type GetJobTfPlanRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the job.
	JobId *string `mandatory:"true" contributesTo:"path" name:"jobId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The output format of the Terraform plan.
	TfPlanFormat GetJobTfPlanTfPlanFormatEnum `mandatory:"false" contributesTo:"query" name:"tfPlanFormat" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetJobTfPlanRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJobTfPlanRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetJobTfPlanRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJobTfPlanRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetJobTfPlanRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetJobTfPlanTfPlanFormatEnum(string(request.TfPlanFormat)); !ok && request.TfPlanFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TfPlanFormat: %s. Supported values are: %s.", request.TfPlanFormat, strings.Join(GetGetJobTfPlanTfPlanFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetJobTfPlanResponse wrapper for the GetJobTfPlan operation
type GetJobTfPlanResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The io.ReadCloser instance
	Content io.ReadCloser `presentIn:"body" encoding:"binary"`

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetJobTfPlanResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJobTfPlanResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetJobTfPlanTfPlanFormatEnum Enum with underlying type: string
type GetJobTfPlanTfPlanFormatEnum string

// Set of constants representing the allowable values for GetJobTfPlanTfPlanFormatEnum
const (
	GetJobTfPlanTfPlanFormatBinary GetJobTfPlanTfPlanFormatEnum = "BINARY"
	GetJobTfPlanTfPlanFormatJson   GetJobTfPlanTfPlanFormatEnum = "JSON"
)

var mappingGetJobTfPlanTfPlanFormatEnum = map[string]GetJobTfPlanTfPlanFormatEnum{
	"BINARY": GetJobTfPlanTfPlanFormatBinary,
	"JSON":   GetJobTfPlanTfPlanFormatJson,
}

var mappingGetJobTfPlanTfPlanFormatEnumLowerCase = map[string]GetJobTfPlanTfPlanFormatEnum{
	"binary": GetJobTfPlanTfPlanFormatBinary,
	"json":   GetJobTfPlanTfPlanFormatJson,
}

// GetGetJobTfPlanTfPlanFormatEnumValues Enumerates the set of values for GetJobTfPlanTfPlanFormatEnum
func GetGetJobTfPlanTfPlanFormatEnumValues() []GetJobTfPlanTfPlanFormatEnum {
	values := make([]GetJobTfPlanTfPlanFormatEnum, 0)
	for _, v := range mappingGetJobTfPlanTfPlanFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetGetJobTfPlanTfPlanFormatEnumStringValues Enumerates the set of values in String for GetJobTfPlanTfPlanFormatEnum
func GetGetJobTfPlanTfPlanFormatEnumStringValues() []string {
	return []string{
		"BINARY",
		"JSON",
	}
}

// GetMappingGetJobTfPlanTfPlanFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetJobTfPlanTfPlanFormatEnum(val string) (GetJobTfPlanTfPlanFormatEnum, bool) {
	enum, ok := mappingGetJobTfPlanTfPlanFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
