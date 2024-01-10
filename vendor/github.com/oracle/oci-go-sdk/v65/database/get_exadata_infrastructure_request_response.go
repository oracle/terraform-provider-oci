// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetExadataInfrastructureRequest wrapper for the GetExadataInfrastructure operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/GetExadataInfrastructure.go.html to see an example of how to use GetExadataInfrastructureRequest.
type GetExadataInfrastructureRequest struct {

	// The Exadata infrastructure OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ExadataInfrastructureId *string `mandatory:"true" contributesTo:"path" name:"exadataInfrastructureId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If provided, the specified fields will be excluded in the response.
	ExcludedFields []GetExadataInfrastructureExcludedFieldsEnum `contributesTo:"query" name:"excludedFields" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetExadataInfrastructureRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetExadataInfrastructureRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetExadataInfrastructureRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetExadataInfrastructureRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetExadataInfrastructureRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ExcludedFields {
		if _, ok := GetMappingGetExadataInfrastructureExcludedFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExcludedFields: %s. Supported values are: %s.", val, strings.Join(GetGetExadataInfrastructureExcludedFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetExadataInfrastructureResponse wrapper for the GetExadataInfrastructure operation
type GetExadataInfrastructureResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The ExadataInfrastructure instance
	ExadataInfrastructure `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetExadataInfrastructureResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetExadataInfrastructureResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetExadataInfrastructureExcludedFieldsEnum Enum with underlying type: string
type GetExadataInfrastructureExcludedFieldsEnum string

// Set of constants representing the allowable values for GetExadataInfrastructureExcludedFieldsEnum
const (
	GetExadataInfrastructureExcludedFieldsMultirackconfigurationfile GetExadataInfrastructureExcludedFieldsEnum = "multiRackConfigurationFile"
)

var mappingGetExadataInfrastructureExcludedFieldsEnum = map[string]GetExadataInfrastructureExcludedFieldsEnum{
	"multiRackConfigurationFile": GetExadataInfrastructureExcludedFieldsMultirackconfigurationfile,
}

var mappingGetExadataInfrastructureExcludedFieldsEnumLowerCase = map[string]GetExadataInfrastructureExcludedFieldsEnum{
	"multirackconfigurationfile": GetExadataInfrastructureExcludedFieldsMultirackconfigurationfile,
}

// GetGetExadataInfrastructureExcludedFieldsEnumValues Enumerates the set of values for GetExadataInfrastructureExcludedFieldsEnum
func GetGetExadataInfrastructureExcludedFieldsEnumValues() []GetExadataInfrastructureExcludedFieldsEnum {
	values := make([]GetExadataInfrastructureExcludedFieldsEnum, 0)
	for _, v := range mappingGetExadataInfrastructureExcludedFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetExadataInfrastructureExcludedFieldsEnumStringValues Enumerates the set of values in String for GetExadataInfrastructureExcludedFieldsEnum
func GetGetExadataInfrastructureExcludedFieldsEnumStringValues() []string {
	return []string{
		"multiRackConfigurationFile",
	}
}

// GetMappingGetExadataInfrastructureExcludedFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetExadataInfrastructureExcludedFieldsEnum(val string) (GetExadataInfrastructureExcludedFieldsEnum, bool) {
	enum, ok := mappingGetExadataInfrastructureExcludedFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
