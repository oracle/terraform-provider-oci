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

// GetStatusAutoActivateRequest wrapper for the GetStatusAutoActivate operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/GetStatusAutoActivate.go.html to see an example of how to use GetStatusAutoActivateRequest.
type GetStatusAutoActivateRequest struct {

	// The APM Domain ID for the intended request.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// Data key type for which auto-activate needs to be turned on or off.
	DataKeyType GetStatusAutoActivateDataKeyTypeEnum `mandatory:"true" contributesTo:"query" name:"dataKeyType" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetStatusAutoActivateRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetStatusAutoActivateRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetStatusAutoActivateRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetStatusAutoActivateRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetStatusAutoActivateRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetStatusAutoActivateDataKeyTypeEnum(string(request.DataKeyType)); !ok && request.DataKeyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataKeyType: %s. Supported values are: %s.", request.DataKeyType, strings.Join(GetGetStatusAutoActivateDataKeyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetStatusAutoActivateResponse wrapper for the GetStatusAutoActivate operation
type GetStatusAutoActivateResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AutoActivateStatus instance
	AutoActivateStatus `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetStatusAutoActivateResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetStatusAutoActivateResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetStatusAutoActivateDataKeyTypeEnum Enum with underlying type: string
type GetStatusAutoActivateDataKeyTypeEnum string

// Set of constants representing the allowable values for GetStatusAutoActivateDataKeyTypeEnum
const (
	GetStatusAutoActivateDataKeyTypePrivateDataKey GetStatusAutoActivateDataKeyTypeEnum = "PRIVATE_DATA_KEY"
	GetStatusAutoActivateDataKeyTypePublicDataKey  GetStatusAutoActivateDataKeyTypeEnum = "PUBLIC_DATA_KEY"
)

var mappingGetStatusAutoActivateDataKeyTypeEnum = map[string]GetStatusAutoActivateDataKeyTypeEnum{
	"PRIVATE_DATA_KEY": GetStatusAutoActivateDataKeyTypePrivateDataKey,
	"PUBLIC_DATA_KEY":  GetStatusAutoActivateDataKeyTypePublicDataKey,
}

var mappingGetStatusAutoActivateDataKeyTypeEnumLowerCase = map[string]GetStatusAutoActivateDataKeyTypeEnum{
	"private_data_key": GetStatusAutoActivateDataKeyTypePrivateDataKey,
	"public_data_key":  GetStatusAutoActivateDataKeyTypePublicDataKey,
}

// GetGetStatusAutoActivateDataKeyTypeEnumValues Enumerates the set of values for GetStatusAutoActivateDataKeyTypeEnum
func GetGetStatusAutoActivateDataKeyTypeEnumValues() []GetStatusAutoActivateDataKeyTypeEnum {
	values := make([]GetStatusAutoActivateDataKeyTypeEnum, 0)
	for _, v := range mappingGetStatusAutoActivateDataKeyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetStatusAutoActivateDataKeyTypeEnumStringValues Enumerates the set of values in String for GetStatusAutoActivateDataKeyTypeEnum
func GetGetStatusAutoActivateDataKeyTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_DATA_KEY",
		"PUBLIC_DATA_KEY",
	}
}

// GetMappingGetStatusAutoActivateDataKeyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetStatusAutoActivateDataKeyTypeEnum(val string) (GetStatusAutoActivateDataKeyTypeEnum, bool) {
	enum, ok := mappingGetStatusAutoActivateDataKeyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
