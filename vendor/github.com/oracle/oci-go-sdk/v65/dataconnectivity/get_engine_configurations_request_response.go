// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetEngineConfigurationsRequest wrapper for the GetEngineConfigurations operation
type GetEngineConfigurationsRequest struct {

	// The registry OCID.
	RegistryId *string `mandatory:"true" contributesTo:"path" name:"registryId"`

	// The connection key.
	ConnectionKey *string `mandatory:"true" contributesTo:"path" name:"connectionKey"`

	// Specifies the runtime engine for the bulk read/write operation. Default is SPARK.
	EngineTypeQueryParam GetEngineConfigurationsEngineTypeQueryParamEnum `mandatory:"false" contributesTo:"query" name:"engineTypeQueryParam" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetEngineConfigurationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetEngineConfigurationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetEngineConfigurationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetEngineConfigurationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetEngineConfigurationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetEngineConfigurationsEngineTypeQueryParamEnum(string(request.EngineTypeQueryParam)); !ok && request.EngineTypeQueryParam != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EngineTypeQueryParam: %s. Supported values are: %s.", request.EngineTypeQueryParam, strings.Join(GetGetEngineConfigurationsEngineTypeQueryParamEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetEngineConfigurationsResponse wrapper for the GetEngineConfigurations operation
type GetEngineConfigurationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ConfigDetails instance
	ConfigDetails `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetEngineConfigurationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetEngineConfigurationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetEngineConfigurationsEngineTypeQueryParamEnum Enum with underlying type: string
type GetEngineConfigurationsEngineTypeQueryParamEnum string

// Set of constants representing the allowable values for GetEngineConfigurationsEngineTypeQueryParamEnum
const (
	GetEngineConfigurationsEngineTypeQueryParamSpark GetEngineConfigurationsEngineTypeQueryParamEnum = "SPARK"
)

var mappingGetEngineConfigurationsEngineTypeQueryParamEnum = map[string]GetEngineConfigurationsEngineTypeQueryParamEnum{
	"SPARK": GetEngineConfigurationsEngineTypeQueryParamSpark,
}

var mappingGetEngineConfigurationsEngineTypeQueryParamEnumLowerCase = map[string]GetEngineConfigurationsEngineTypeQueryParamEnum{
	"spark": GetEngineConfigurationsEngineTypeQueryParamSpark,
}

// GetGetEngineConfigurationsEngineTypeQueryParamEnumValues Enumerates the set of values for GetEngineConfigurationsEngineTypeQueryParamEnum
func GetGetEngineConfigurationsEngineTypeQueryParamEnumValues() []GetEngineConfigurationsEngineTypeQueryParamEnum {
	values := make([]GetEngineConfigurationsEngineTypeQueryParamEnum, 0)
	for _, v := range mappingGetEngineConfigurationsEngineTypeQueryParamEnum {
		values = append(values, v)
	}
	return values
}

// GetGetEngineConfigurationsEngineTypeQueryParamEnumStringValues Enumerates the set of values in String for GetEngineConfigurationsEngineTypeQueryParamEnum
func GetGetEngineConfigurationsEngineTypeQueryParamEnumStringValues() []string {
	return []string{
		"SPARK",
	}
}

// GetMappingGetEngineConfigurationsEngineTypeQueryParamEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetEngineConfigurationsEngineTypeQueryParamEnum(val string) (GetEngineConfigurationsEngineTypeQueryParamEnum, bool) {
	enum, ok := mappingGetEngineConfigurationsEngineTypeQueryParamEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
