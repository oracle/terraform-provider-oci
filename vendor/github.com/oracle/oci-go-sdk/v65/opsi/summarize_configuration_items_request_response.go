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

// SummarizeConfigurationItemsRequest wrapper for the SummarizeConfigurationItems operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeConfigurationItems.go.html to see an example of how to use SummarizeConfigurationItemsRequest.
type SummarizeConfigurationItemsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Filter to return configuration items based on configuration type of OPSI configuration.
	OpsiConfigType SummarizeConfigurationItemsOpsiConfigTypeEnum `mandatory:"false" contributesTo:"query" name:"opsiConfigType" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Returns the configuration items filtered by applicable contexts sent in this param. By default configuration items of all applicable contexts are returned.
	ConfigItemsApplicableContext []string `contributesTo:"query" name:"configItemsApplicableContext" collectionFormat:"multi"`

	// Specifies the fields to return in a config item summary.
	ConfigItemField []SummarizeConfigurationItemsConfigItemFieldEnum `contributesTo:"query" name:"configItemField" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only configuration items that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeConfigurationItemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeConfigurationItemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeConfigurationItemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeConfigurationItemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeConfigurationItemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeConfigurationItemsOpsiConfigTypeEnum(string(request.OpsiConfigType)); !ok && request.OpsiConfigType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpsiConfigType: %s. Supported values are: %s.", request.OpsiConfigType, strings.Join(GetSummarizeConfigurationItemsOpsiConfigTypeEnumStringValues(), ",")))
	}
	for _, val := range request.ConfigItemField {
		if _, ok := GetMappingSummarizeConfigurationItemsConfigItemFieldEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConfigItemField: %s. Supported values are: %s.", val, strings.Join(GetSummarizeConfigurationItemsConfigItemFieldEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeConfigurationItemsResponse wrapper for the SummarizeConfigurationItems operation
type SummarizeConfigurationItemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConfigurationItemsCollection instances
	ConfigurationItemsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeConfigurationItemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeConfigurationItemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeConfigurationItemsOpsiConfigTypeEnum Enum with underlying type: string
type SummarizeConfigurationItemsOpsiConfigTypeEnum string

// Set of constants representing the allowable values for SummarizeConfigurationItemsOpsiConfigTypeEnum
const (
	SummarizeConfigurationItemsOpsiConfigTypeUxConfiguration SummarizeConfigurationItemsOpsiConfigTypeEnum = "UX_CONFIGURATION"
)

var mappingSummarizeConfigurationItemsOpsiConfigTypeEnum = map[string]SummarizeConfigurationItemsOpsiConfigTypeEnum{
	"UX_CONFIGURATION": SummarizeConfigurationItemsOpsiConfigTypeUxConfiguration,
}

var mappingSummarizeConfigurationItemsOpsiConfigTypeEnumLowerCase = map[string]SummarizeConfigurationItemsOpsiConfigTypeEnum{
	"ux_configuration": SummarizeConfigurationItemsOpsiConfigTypeUxConfiguration,
}

// GetSummarizeConfigurationItemsOpsiConfigTypeEnumValues Enumerates the set of values for SummarizeConfigurationItemsOpsiConfigTypeEnum
func GetSummarizeConfigurationItemsOpsiConfigTypeEnumValues() []SummarizeConfigurationItemsOpsiConfigTypeEnum {
	values := make([]SummarizeConfigurationItemsOpsiConfigTypeEnum, 0)
	for _, v := range mappingSummarizeConfigurationItemsOpsiConfigTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeConfigurationItemsOpsiConfigTypeEnumStringValues Enumerates the set of values in String for SummarizeConfigurationItemsOpsiConfigTypeEnum
func GetSummarizeConfigurationItemsOpsiConfigTypeEnumStringValues() []string {
	return []string{
		"UX_CONFIGURATION",
	}
}

// GetMappingSummarizeConfigurationItemsOpsiConfigTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeConfigurationItemsOpsiConfigTypeEnum(val string) (SummarizeConfigurationItemsOpsiConfigTypeEnum, bool) {
	enum, ok := mappingSummarizeConfigurationItemsOpsiConfigTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeConfigurationItemsConfigItemFieldEnum Enum with underlying type: string
type SummarizeConfigurationItemsConfigItemFieldEnum string

// Set of constants representing the allowable values for SummarizeConfigurationItemsConfigItemFieldEnum
const (
	SummarizeConfigurationItemsConfigItemFieldName               SummarizeConfigurationItemsConfigItemFieldEnum = "name"
	SummarizeConfigurationItemsConfigItemFieldValue              SummarizeConfigurationItemsConfigItemFieldEnum = "value"
	SummarizeConfigurationItemsConfigItemFieldDefaultvalue       SummarizeConfigurationItemsConfigItemFieldEnum = "defaultValue"
	SummarizeConfigurationItemsConfigItemFieldValuesourceconfig  SummarizeConfigurationItemsConfigItemFieldEnum = "valueSourceConfig"
	SummarizeConfigurationItemsConfigItemFieldMetadata           SummarizeConfigurationItemsConfigItemFieldEnum = "metadata"
	SummarizeConfigurationItemsConfigItemFieldApplicablecontexts SummarizeConfigurationItemsConfigItemFieldEnum = "applicableContexts"
)

var mappingSummarizeConfigurationItemsConfigItemFieldEnum = map[string]SummarizeConfigurationItemsConfigItemFieldEnum{
	"name":               SummarizeConfigurationItemsConfigItemFieldName,
	"value":              SummarizeConfigurationItemsConfigItemFieldValue,
	"defaultValue":       SummarizeConfigurationItemsConfigItemFieldDefaultvalue,
	"valueSourceConfig":  SummarizeConfigurationItemsConfigItemFieldValuesourceconfig,
	"metadata":           SummarizeConfigurationItemsConfigItemFieldMetadata,
	"applicableContexts": SummarizeConfigurationItemsConfigItemFieldApplicablecontexts,
}

var mappingSummarizeConfigurationItemsConfigItemFieldEnumLowerCase = map[string]SummarizeConfigurationItemsConfigItemFieldEnum{
	"name":               SummarizeConfigurationItemsConfigItemFieldName,
	"value":              SummarizeConfigurationItemsConfigItemFieldValue,
	"defaultvalue":       SummarizeConfigurationItemsConfigItemFieldDefaultvalue,
	"valuesourceconfig":  SummarizeConfigurationItemsConfigItemFieldValuesourceconfig,
	"metadata":           SummarizeConfigurationItemsConfigItemFieldMetadata,
	"applicablecontexts": SummarizeConfigurationItemsConfigItemFieldApplicablecontexts,
}

// GetSummarizeConfigurationItemsConfigItemFieldEnumValues Enumerates the set of values for SummarizeConfigurationItemsConfigItemFieldEnum
func GetSummarizeConfigurationItemsConfigItemFieldEnumValues() []SummarizeConfigurationItemsConfigItemFieldEnum {
	values := make([]SummarizeConfigurationItemsConfigItemFieldEnum, 0)
	for _, v := range mappingSummarizeConfigurationItemsConfigItemFieldEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeConfigurationItemsConfigItemFieldEnumStringValues Enumerates the set of values in String for SummarizeConfigurationItemsConfigItemFieldEnum
func GetSummarizeConfigurationItemsConfigItemFieldEnumStringValues() []string {
	return []string{
		"name",
		"value",
		"defaultValue",
		"valueSourceConfig",
		"metadata",
		"applicableContexts",
	}
}

// GetMappingSummarizeConfigurationItemsConfigItemFieldEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeConfigurationItemsConfigItemFieldEnum(val string) (SummarizeConfigurationItemsConfigItemFieldEnum, bool) {
	enum, ok := mappingSummarizeConfigurationItemsConfigItemFieldEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
