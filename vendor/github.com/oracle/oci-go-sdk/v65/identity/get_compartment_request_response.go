// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetCompartmentRequest wrapper for the GetCompartment operation
type GetCompartmentRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"path" name:"compartmentId"`

	// This parameter is required to retrieve securityZoneId associated with the compartment.
	VerboseLevel GetCompartmentVerboseLevelEnum `mandatory:"false" contributesTo:"query" name:"verboseLevel" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCompartmentRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetCompartmentRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetCompartmentRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetCompartmentVerboseLevelEnum(string(request.VerboseLevel)); !ok && request.VerboseLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VerboseLevel: %s. Supported values are: %s.", request.VerboseLevel, strings.Join(GetGetCompartmentVerboseLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetCompartmentResponse wrapper for the GetCompartment operation
type GetCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Compartment instance
	Compartment `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetCompartmentVerboseLevelEnum Enum with underlying type: string
type GetCompartmentVerboseLevelEnum string

// Set of constants representing the allowable values for GetCompartmentVerboseLevelEnum
const (
	GetCompartmentVerboseLevelSecurityzone GetCompartmentVerboseLevelEnum = "securityZone"
)

var mappingGetCompartmentVerboseLevelEnum = map[string]GetCompartmentVerboseLevelEnum{
	"securityZone": GetCompartmentVerboseLevelSecurityzone,
}

var mappingGetCompartmentVerboseLevelEnumLowerCase = map[string]GetCompartmentVerboseLevelEnum{
	"securityzone": GetCompartmentVerboseLevelSecurityzone,
}

// GetGetCompartmentVerboseLevelEnumValues Enumerates the set of values for GetCompartmentVerboseLevelEnum
func GetGetCompartmentVerboseLevelEnumValues() []GetCompartmentVerboseLevelEnum {
	values := make([]GetCompartmentVerboseLevelEnum, 0)
	for _, v := range mappingGetCompartmentVerboseLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetGetCompartmentVerboseLevelEnumStringValues Enumerates the set of values in String for GetCompartmentVerboseLevelEnum
func GetGetCompartmentVerboseLevelEnumStringValues() []string {
	return []string{
		"securityZone",
	}
}

// GetMappingGetCompartmentVerboseLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetCompartmentVerboseLevelEnum(val string) (GetCompartmentVerboseLevelEnum, bool) {
	enum, ok := mappingGetCompartmentVerboseLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
