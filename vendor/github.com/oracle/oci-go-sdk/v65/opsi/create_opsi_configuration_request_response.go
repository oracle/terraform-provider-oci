// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// CreateOpsiConfigurationRequest wrapper for the CreateOpsiConfiguration operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/CreateOpsiConfiguration.go.html to see an example of how to use CreateOpsiConfigurationRequest.
type CreateOpsiConfigurationRequest struct {

	// Information about OPSI configuration resource to be created.
	CreateOpsiConfigurationDetails `contributesTo:"body"`

	// A token that uniquely identifies a request that can be retried in case of a timeout or
	// server error without risk of executing the same action again. Retry tokens expire after 24
	// hours.
	// *Note:* Retry tokens can be invalidated before the 24 hour time limit due to conflicting
	// operations, such as a resource being deleted or purged from the system.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Optional fields to return as part of OpsiConfiguration object. Unless requested, these fields will not be returned by default.
	OpsiConfigField []CreateOpsiConfigurationOpsiConfigFieldEnum `contributesTo:"query" name:"opsiConfigField" omitEmpty:"true" collectionFormat:"multi"`

	// Specifies whether only customized configuration items or only non-customized configuration items or both have to be returned.
	// By default only customized configuration items are returned.
	ConfigItemCustomStatus []CreateOpsiConfigurationConfigItemCustomStatusEnum `contributesTo:"query" name:"configItemCustomStatus" omitEmpty:"true" collectionFormat:"multi"`

	// Returns the configuration items filtered by applicable contexts sent in this param. By default configuration items of all applicable contexts are returned.
	ConfigItemsApplicableContext []string `contributesTo:"query" name:"configItemsApplicableContext" collectionFormat:"multi"`

	// Specifies the fields to return in a config item summary.
	ConfigItemField []CreateOpsiConfigurationConfigItemFieldEnum `contributesTo:"query" name:"configItemField" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateOpsiConfigurationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateOpsiConfigurationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateOpsiConfigurationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateOpsiConfigurationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request CreateOpsiConfigurationRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.OpsiConfigField {
		if _, ok := GetMappingCreateOpsiConfigurationOpsiConfigFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpsiConfigField: %s. Supported values are: %s.", val, strings.Join(GetCreateOpsiConfigurationOpsiConfigFieldEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ConfigItemCustomStatus {
		if _, ok := GetMappingCreateOpsiConfigurationConfigItemCustomStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigItemCustomStatus: %s. Supported values are: %s.", val, strings.Join(GetCreateOpsiConfigurationConfigItemCustomStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ConfigItemField {
		if _, ok := GetMappingCreateOpsiConfigurationConfigItemFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigItemField: %s. Supported values are: %s.", val, strings.Join(GetCreateOpsiConfigurationConfigItemFieldEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateOpsiConfigurationResponse wrapper for the CreateOpsiConfiguration operation
type CreateOpsiConfigurationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OpsiConfiguration instance
	OpsiConfiguration `presentIn:"body"`

	// Unique Oracle-assigned identifier for the asynchronous request. You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response CreateOpsiConfigurationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateOpsiConfigurationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// CreateOpsiConfigurationOpsiConfigFieldEnum Enum with underlying type: string
type CreateOpsiConfigurationOpsiConfigFieldEnum string

// Set of constants representing the allowable values for CreateOpsiConfigurationOpsiConfigFieldEnum
const (
	CreateOpsiConfigurationOpsiConfigFieldConfigitems CreateOpsiConfigurationOpsiConfigFieldEnum = "configItems"
)

var mappingCreateOpsiConfigurationOpsiConfigFieldEnum = map[string]CreateOpsiConfigurationOpsiConfigFieldEnum{
	"configItems": CreateOpsiConfigurationOpsiConfigFieldConfigitems,
}

var mappingCreateOpsiConfigurationOpsiConfigFieldEnumLowerCase = map[string]CreateOpsiConfigurationOpsiConfigFieldEnum{
	"configitems": CreateOpsiConfigurationOpsiConfigFieldConfigitems,
}

// GetCreateOpsiConfigurationOpsiConfigFieldEnumValues Enumerates the set of values for CreateOpsiConfigurationOpsiConfigFieldEnum
func GetCreateOpsiConfigurationOpsiConfigFieldEnumValues() []CreateOpsiConfigurationOpsiConfigFieldEnum {
	values := make([]CreateOpsiConfigurationOpsiConfigFieldEnum, 0)
	for _, v := range mappingCreateOpsiConfigurationOpsiConfigFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOpsiConfigurationOpsiConfigFieldEnumStringValues Enumerates the set of values in String for CreateOpsiConfigurationOpsiConfigFieldEnum
func GetCreateOpsiConfigurationOpsiConfigFieldEnumStringValues() []string {
	return []string{
		"configItems",
	}
}

// GetMappingCreateOpsiConfigurationOpsiConfigFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOpsiConfigurationOpsiConfigFieldEnum(val string) (CreateOpsiConfigurationOpsiConfigFieldEnum, bool) {
	enum, ok := mappingCreateOpsiConfigurationOpsiConfigFieldEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateOpsiConfigurationConfigItemCustomStatusEnum Enum with underlying type: string
type CreateOpsiConfigurationConfigItemCustomStatusEnum string

// Set of constants representing the allowable values for CreateOpsiConfigurationConfigItemCustomStatusEnum
const (
	CreateOpsiConfigurationConfigItemCustomStatusCustomized    CreateOpsiConfigurationConfigItemCustomStatusEnum = "customized"
	CreateOpsiConfigurationConfigItemCustomStatusNoncustomized CreateOpsiConfigurationConfigItemCustomStatusEnum = "nonCustomized"
)

var mappingCreateOpsiConfigurationConfigItemCustomStatusEnum = map[string]CreateOpsiConfigurationConfigItemCustomStatusEnum{
	"customized":    CreateOpsiConfigurationConfigItemCustomStatusCustomized,
	"nonCustomized": CreateOpsiConfigurationConfigItemCustomStatusNoncustomized,
}

var mappingCreateOpsiConfigurationConfigItemCustomStatusEnumLowerCase = map[string]CreateOpsiConfigurationConfigItemCustomStatusEnum{
	"customized":    CreateOpsiConfigurationConfigItemCustomStatusCustomized,
	"noncustomized": CreateOpsiConfigurationConfigItemCustomStatusNoncustomized,
}

// GetCreateOpsiConfigurationConfigItemCustomStatusEnumValues Enumerates the set of values for CreateOpsiConfigurationConfigItemCustomStatusEnum
func GetCreateOpsiConfigurationConfigItemCustomStatusEnumValues() []CreateOpsiConfigurationConfigItemCustomStatusEnum {
	values := make([]CreateOpsiConfigurationConfigItemCustomStatusEnum, 0)
	for _, v := range mappingCreateOpsiConfigurationConfigItemCustomStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOpsiConfigurationConfigItemCustomStatusEnumStringValues Enumerates the set of values in String for CreateOpsiConfigurationConfigItemCustomStatusEnum
func GetCreateOpsiConfigurationConfigItemCustomStatusEnumStringValues() []string {
	return []string{
		"customized",
		"nonCustomized",
	}
}

// GetMappingCreateOpsiConfigurationConfigItemCustomStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOpsiConfigurationConfigItemCustomStatusEnum(val string) (CreateOpsiConfigurationConfigItemCustomStatusEnum, bool) {
	enum, ok := mappingCreateOpsiConfigurationConfigItemCustomStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateOpsiConfigurationConfigItemFieldEnum Enum with underlying type: string
type CreateOpsiConfigurationConfigItemFieldEnum string

// Set of constants representing the allowable values for CreateOpsiConfigurationConfigItemFieldEnum
const (
	CreateOpsiConfigurationConfigItemFieldName               CreateOpsiConfigurationConfigItemFieldEnum = "name"
	CreateOpsiConfigurationConfigItemFieldValue              CreateOpsiConfigurationConfigItemFieldEnum = "value"
	CreateOpsiConfigurationConfigItemFieldDefaultvalue       CreateOpsiConfigurationConfigItemFieldEnum = "defaultValue"
	CreateOpsiConfigurationConfigItemFieldMetadata           CreateOpsiConfigurationConfigItemFieldEnum = "metadata"
	CreateOpsiConfigurationConfigItemFieldApplicablecontexts CreateOpsiConfigurationConfigItemFieldEnum = "applicableContexts"
)

var mappingCreateOpsiConfigurationConfigItemFieldEnum = map[string]CreateOpsiConfigurationConfigItemFieldEnum{
	"name":               CreateOpsiConfigurationConfigItemFieldName,
	"value":              CreateOpsiConfigurationConfigItemFieldValue,
	"defaultValue":       CreateOpsiConfigurationConfigItemFieldDefaultvalue,
	"metadata":           CreateOpsiConfigurationConfigItemFieldMetadata,
	"applicableContexts": CreateOpsiConfigurationConfigItemFieldApplicablecontexts,
}

var mappingCreateOpsiConfigurationConfigItemFieldEnumLowerCase = map[string]CreateOpsiConfigurationConfigItemFieldEnum{
	"name":               CreateOpsiConfigurationConfigItemFieldName,
	"value":              CreateOpsiConfigurationConfigItemFieldValue,
	"defaultvalue":       CreateOpsiConfigurationConfigItemFieldDefaultvalue,
	"metadata":           CreateOpsiConfigurationConfigItemFieldMetadata,
	"applicablecontexts": CreateOpsiConfigurationConfigItemFieldApplicablecontexts,
}

// GetCreateOpsiConfigurationConfigItemFieldEnumValues Enumerates the set of values for CreateOpsiConfigurationConfigItemFieldEnum
func GetCreateOpsiConfigurationConfigItemFieldEnumValues() []CreateOpsiConfigurationConfigItemFieldEnum {
	values := make([]CreateOpsiConfigurationConfigItemFieldEnum, 0)
	for _, v := range mappingCreateOpsiConfigurationConfigItemFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateOpsiConfigurationConfigItemFieldEnumStringValues Enumerates the set of values in String for CreateOpsiConfigurationConfigItemFieldEnum
func GetCreateOpsiConfigurationConfigItemFieldEnumStringValues() []string {
	return []string{
		"name",
		"value",
		"defaultValue",
		"metadata",
		"applicableContexts",
	}
}

// GetMappingCreateOpsiConfigurationConfigItemFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateOpsiConfigurationConfigItemFieldEnum(val string) (CreateOpsiConfigurationConfigItemFieldEnum, bool) {
	enum, ok := mappingCreateOpsiConfigurationConfigItemFieldEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
