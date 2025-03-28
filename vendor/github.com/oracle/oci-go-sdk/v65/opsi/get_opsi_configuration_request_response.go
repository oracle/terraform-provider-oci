// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetOpsiConfigurationRequest wrapper for the GetOpsiConfiguration operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/GetOpsiConfiguration.go.html to see an example of how to use GetOpsiConfigurationRequest.
type GetOpsiConfigurationRequest struct {

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OPSI configuration resource.
	OpsiConfigurationId *string `mandatory:"true" contributesTo:"path" name:"opsiConfigurationId"`

	// Optional fields to return as part of OpsiConfiguration object. Unless requested, these fields will not be returned by default.
	OpsiConfigField []GetOpsiConfigurationOpsiConfigFieldEnum `contributesTo:"query" name:"opsiConfigField" omitEmpty:"true" collectionFormat:"multi"`

	// Specifies whether only customized configuration items or only non-customized configuration items or both have to be returned.
	// By default only customized configuration items are returned.
	ConfigItemCustomStatus []GetOpsiConfigurationConfigItemCustomStatusEnum `contributesTo:"query" name:"configItemCustomStatus" omitEmpty:"true" collectionFormat:"multi"`

	// Returns the configuration items filtered by applicable contexts sent in this param. By default configuration items of all applicable contexts are returned.
	ConfigItemsApplicableContext []string `contributesTo:"query" name:"configItemsApplicableContext" collectionFormat:"multi"`

	// Specifies the fields to return in a config item summary.
	ConfigItemField []GetOpsiConfigurationConfigItemFieldEnum `contributesTo:"query" name:"configItemField" omitEmpty:"true" collectionFormat:"multi"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetOpsiConfigurationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetOpsiConfigurationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetOpsiConfigurationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetOpsiConfigurationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetOpsiConfigurationRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.OpsiConfigField {
		if _, ok := GetMappingGetOpsiConfigurationOpsiConfigFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpsiConfigField: %s. Supported values are: %s.", val, strings.Join(GetGetOpsiConfigurationOpsiConfigFieldEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ConfigItemCustomStatus {
		if _, ok := GetMappingGetOpsiConfigurationConfigItemCustomStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigItemCustomStatus: %s. Supported values are: %s.", val, strings.Join(GetGetOpsiConfigurationConfigItemCustomStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ConfigItemField {
		if _, ok := GetMappingGetOpsiConfigurationConfigItemFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigItemField: %s. Supported values are: %s.", val, strings.Join(GetGetOpsiConfigurationConfigItemFieldEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetOpsiConfigurationResponse wrapper for the GetOpsiConfiguration operation
type GetOpsiConfigurationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OpsiConfiguration instance
	OpsiConfiguration `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetOpsiConfigurationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetOpsiConfigurationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetOpsiConfigurationOpsiConfigFieldEnum Enum with underlying type: string
type GetOpsiConfigurationOpsiConfigFieldEnum string

// Set of constants representing the allowable values for GetOpsiConfigurationOpsiConfigFieldEnum
const (
	GetOpsiConfigurationOpsiConfigFieldConfigitems GetOpsiConfigurationOpsiConfigFieldEnum = "configItems"
)

var mappingGetOpsiConfigurationOpsiConfigFieldEnum = map[string]GetOpsiConfigurationOpsiConfigFieldEnum{
	"configItems": GetOpsiConfigurationOpsiConfigFieldConfigitems,
}

var mappingGetOpsiConfigurationOpsiConfigFieldEnumLowerCase = map[string]GetOpsiConfigurationOpsiConfigFieldEnum{
	"configitems": GetOpsiConfigurationOpsiConfigFieldConfigitems,
}

// GetGetOpsiConfigurationOpsiConfigFieldEnumValues Enumerates the set of values for GetOpsiConfigurationOpsiConfigFieldEnum
func GetGetOpsiConfigurationOpsiConfigFieldEnumValues() []GetOpsiConfigurationOpsiConfigFieldEnum {
	values := make([]GetOpsiConfigurationOpsiConfigFieldEnum, 0)
	for _, v := range mappingGetOpsiConfigurationOpsiConfigFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetGetOpsiConfigurationOpsiConfigFieldEnumStringValues Enumerates the set of values in String for GetOpsiConfigurationOpsiConfigFieldEnum
func GetGetOpsiConfigurationOpsiConfigFieldEnumStringValues() []string {
	return []string{
		"configItems",
	}
}

// GetMappingGetOpsiConfigurationOpsiConfigFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetOpsiConfigurationOpsiConfigFieldEnum(val string) (GetOpsiConfigurationOpsiConfigFieldEnum, bool) {
	enum, ok := mappingGetOpsiConfigurationOpsiConfigFieldEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetOpsiConfigurationConfigItemCustomStatusEnum Enum with underlying type: string
type GetOpsiConfigurationConfigItemCustomStatusEnum string

// Set of constants representing the allowable values for GetOpsiConfigurationConfigItemCustomStatusEnum
const (
	GetOpsiConfigurationConfigItemCustomStatusCustomized    GetOpsiConfigurationConfigItemCustomStatusEnum = "customized"
	GetOpsiConfigurationConfigItemCustomStatusNoncustomized GetOpsiConfigurationConfigItemCustomStatusEnum = "nonCustomized"
)

var mappingGetOpsiConfigurationConfigItemCustomStatusEnum = map[string]GetOpsiConfigurationConfigItemCustomStatusEnum{
	"customized":    GetOpsiConfigurationConfigItemCustomStatusCustomized,
	"nonCustomized": GetOpsiConfigurationConfigItemCustomStatusNoncustomized,
}

var mappingGetOpsiConfigurationConfigItemCustomStatusEnumLowerCase = map[string]GetOpsiConfigurationConfigItemCustomStatusEnum{
	"customized":    GetOpsiConfigurationConfigItemCustomStatusCustomized,
	"noncustomized": GetOpsiConfigurationConfigItemCustomStatusNoncustomized,
}

// GetGetOpsiConfigurationConfigItemCustomStatusEnumValues Enumerates the set of values for GetOpsiConfigurationConfigItemCustomStatusEnum
func GetGetOpsiConfigurationConfigItemCustomStatusEnumValues() []GetOpsiConfigurationConfigItemCustomStatusEnum {
	values := make([]GetOpsiConfigurationConfigItemCustomStatusEnum, 0)
	for _, v := range mappingGetOpsiConfigurationConfigItemCustomStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetGetOpsiConfigurationConfigItemCustomStatusEnumStringValues Enumerates the set of values in String for GetOpsiConfigurationConfigItemCustomStatusEnum
func GetGetOpsiConfigurationConfigItemCustomStatusEnumStringValues() []string {
	return []string{
		"customized",
		"nonCustomized",
	}
}

// GetMappingGetOpsiConfigurationConfigItemCustomStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetOpsiConfigurationConfigItemCustomStatusEnum(val string) (GetOpsiConfigurationConfigItemCustomStatusEnum, bool) {
	enum, ok := mappingGetOpsiConfigurationConfigItemCustomStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetOpsiConfigurationConfigItemFieldEnum Enum with underlying type: string
type GetOpsiConfigurationConfigItemFieldEnum string

// Set of constants representing the allowable values for GetOpsiConfigurationConfigItemFieldEnum
const (
	GetOpsiConfigurationConfigItemFieldName               GetOpsiConfigurationConfigItemFieldEnum = "name"
	GetOpsiConfigurationConfigItemFieldValue              GetOpsiConfigurationConfigItemFieldEnum = "value"
	GetOpsiConfigurationConfigItemFieldDefaultvalue       GetOpsiConfigurationConfigItemFieldEnum = "defaultValue"
	GetOpsiConfigurationConfigItemFieldMetadata           GetOpsiConfigurationConfigItemFieldEnum = "metadata"
	GetOpsiConfigurationConfigItemFieldApplicablecontexts GetOpsiConfigurationConfigItemFieldEnum = "applicableContexts"
)

var mappingGetOpsiConfigurationConfigItemFieldEnum = map[string]GetOpsiConfigurationConfigItemFieldEnum{
	"name":               GetOpsiConfigurationConfigItemFieldName,
	"value":              GetOpsiConfigurationConfigItemFieldValue,
	"defaultValue":       GetOpsiConfigurationConfigItemFieldDefaultvalue,
	"metadata":           GetOpsiConfigurationConfigItemFieldMetadata,
	"applicableContexts": GetOpsiConfigurationConfigItemFieldApplicablecontexts,
}

var mappingGetOpsiConfigurationConfigItemFieldEnumLowerCase = map[string]GetOpsiConfigurationConfigItemFieldEnum{
	"name":               GetOpsiConfigurationConfigItemFieldName,
	"value":              GetOpsiConfigurationConfigItemFieldValue,
	"defaultvalue":       GetOpsiConfigurationConfigItemFieldDefaultvalue,
	"metadata":           GetOpsiConfigurationConfigItemFieldMetadata,
	"applicablecontexts": GetOpsiConfigurationConfigItemFieldApplicablecontexts,
}

// GetGetOpsiConfigurationConfigItemFieldEnumValues Enumerates the set of values for GetOpsiConfigurationConfigItemFieldEnum
func GetGetOpsiConfigurationConfigItemFieldEnumValues() []GetOpsiConfigurationConfigItemFieldEnum {
	values := make([]GetOpsiConfigurationConfigItemFieldEnum, 0)
	for _, v := range mappingGetOpsiConfigurationConfigItemFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetGetOpsiConfigurationConfigItemFieldEnumStringValues Enumerates the set of values in String for GetOpsiConfigurationConfigItemFieldEnum
func GetGetOpsiConfigurationConfigItemFieldEnumStringValues() []string {
	return []string{
		"name",
		"value",
		"defaultValue",
		"metadata",
		"applicableContexts",
	}
}

// GetMappingGetOpsiConfigurationConfigItemFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetOpsiConfigurationConfigItemFieldEnum(val string) (GetOpsiConfigurationConfigItemFieldEnum, bool) {
	enum, ok := mappingGetOpsiConfigurationConfigItemFieldEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
