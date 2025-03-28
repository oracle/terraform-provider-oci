// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// PutToggleAutoActivateRequest wrapper for the PutToggleAutoActivate operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmtraces/PutToggleAutoActivate.go.html to see an example of how to use PutToggleAutoActivateRequest.
type PutToggleAutoActivateRequest struct {

	// The APM Domain ID for the intended request.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// Autoactivate toggle switch.  Set to true to turn on auto-activate.  Set to false to turn off auto-activate.
	IsAutoActivateOn *bool `mandatory:"true" contributesTo:"query" name:"isAutoActivateOn"`

	// Data key type for which auto-activate needs to be turned on or off.
	DataKeyType PutToggleAutoActivateDataKeyTypeEnum `mandatory:"true" contributesTo:"query" name:"dataKeyType" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PutToggleAutoActivateRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PutToggleAutoActivateRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request PutToggleAutoActivateRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PutToggleAutoActivateRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request PutToggleAutoActivateRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPutToggleAutoActivateDataKeyTypeEnum(string(request.DataKeyType)); !ok && request.DataKeyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataKeyType: %s. Supported values are: %s.", request.DataKeyType, strings.Join(GetPutToggleAutoActivateDataKeyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PutToggleAutoActivateResponse wrapper for the PutToggleAutoActivate operation
type PutToggleAutoActivateResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AutoActivateToggleStatus instance
	AutoActivateToggleStatus `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response PutToggleAutoActivateResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PutToggleAutoActivateResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// PutToggleAutoActivateDataKeyTypeEnum Enum with underlying type: string
type PutToggleAutoActivateDataKeyTypeEnum string

// Set of constants representing the allowable values for PutToggleAutoActivateDataKeyTypeEnum
const (
	PutToggleAutoActivateDataKeyTypePrivateDataKey PutToggleAutoActivateDataKeyTypeEnum = "PRIVATE_DATA_KEY"
	PutToggleAutoActivateDataKeyTypePublicDataKey  PutToggleAutoActivateDataKeyTypeEnum = "PUBLIC_DATA_KEY"
)

var mappingPutToggleAutoActivateDataKeyTypeEnum = map[string]PutToggleAutoActivateDataKeyTypeEnum{
	"PRIVATE_DATA_KEY": PutToggleAutoActivateDataKeyTypePrivateDataKey,
	"PUBLIC_DATA_KEY":  PutToggleAutoActivateDataKeyTypePublicDataKey,
}

var mappingPutToggleAutoActivateDataKeyTypeEnumLowerCase = map[string]PutToggleAutoActivateDataKeyTypeEnum{
	"private_data_key": PutToggleAutoActivateDataKeyTypePrivateDataKey,
	"public_data_key":  PutToggleAutoActivateDataKeyTypePublicDataKey,
}

// GetPutToggleAutoActivateDataKeyTypeEnumValues Enumerates the set of values for PutToggleAutoActivateDataKeyTypeEnum
func GetPutToggleAutoActivateDataKeyTypeEnumValues() []PutToggleAutoActivateDataKeyTypeEnum {
	values := make([]PutToggleAutoActivateDataKeyTypeEnum, 0)
	for _, v := range mappingPutToggleAutoActivateDataKeyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPutToggleAutoActivateDataKeyTypeEnumStringValues Enumerates the set of values in String for PutToggleAutoActivateDataKeyTypeEnum
func GetPutToggleAutoActivateDataKeyTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_DATA_KEY",
		"PUBLIC_DATA_KEY",
	}
}

// GetMappingPutToggleAutoActivateDataKeyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPutToggleAutoActivateDataKeyTypeEnum(val string) (PutToggleAutoActivateDataKeyTypeEnum, bool) {
	enum, ok := mappingPutToggleAutoActivateDataKeyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
