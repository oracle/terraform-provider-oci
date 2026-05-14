// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetPropertySetRequest wrapper for the GetPropertySet operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/GetPropertySet.go.html to see an example of how to use GetPropertySetRequest.
type GetPropertySetRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.
	DatabaseToolsConnectionId *string `mandatory:"true" contributesTo:"path" name:"databaseToolsConnectionId"`

	// The name of the property set
	PropertySetKey GetPropertySetPropertySetKeyEnum `mandatory:"true" contributesTo:"path" name:"propertySetKey"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// If-Match is most often used with state-changing methods (e.g., POST, PUT, DELETE) to prevent
	// accidental overwrites when multiple user agentss might be acting in parallel on the same
	// resource (i.e., to prevent the "lost update" problem). In general, it can be used with any
	// method that involves the selection or modification of a representation to abort the request
	// if the selected representation's current entity tag is not a member within the If-Match field value.
	// When specified on an action-specific subresource, the ETag value of the resource on which the
	// action is requested should be provided.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPropertySetRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPropertySetRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetPropertySetRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPropertySetRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetPropertySetRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetPropertySetPropertySetKeyEnum(string(request.PropertySetKey)); !ok && request.PropertySetKey != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PropertySetKey: %s. Supported values are: %s.", request.PropertySetKey, strings.Join(GetGetPropertySetPropertySetKeyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetPropertySetResponse wrapper for the GetPropertySet operation
type GetPropertySetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The PropertySet instance
	PropertySet `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPropertySetResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPropertySetResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetPropertySetPropertySetKeyEnum Enum with underlying type: string
type GetPropertySetPropertySetKeyEnum string

// Set of constants representing the allowable values for GetPropertySetPropertySetKeyEnum
const (
	GetPropertySetPropertySetKeyApexDocumentGenerator                GetPropertySetPropertySetKeyEnum = "APEX_DOCUMENT_GENERATOR"
	GetPropertySetPropertySetKeyApex                                 GetPropertySetPropertySetKeyEnum = "APEX"
	GetPropertySetPropertySetKeyApexFaIntegration                    GetPropertySetPropertySetKeyEnum = "APEX_FA_INTEGRATION"
	GetPropertySetPropertySetKeyOracleDatabaseExternalAuthentication GetPropertySetPropertySetKeyEnum = "ORACLE_DATABASE_EXTERNAL_AUTHENTICATION"
)

var mappingGetPropertySetPropertySetKeyEnum = map[string]GetPropertySetPropertySetKeyEnum{
	"APEX_DOCUMENT_GENERATOR": GetPropertySetPropertySetKeyApexDocumentGenerator,
	"APEX":                    GetPropertySetPropertySetKeyApex,
	"APEX_FA_INTEGRATION":     GetPropertySetPropertySetKeyApexFaIntegration,
	"ORACLE_DATABASE_EXTERNAL_AUTHENTICATION": GetPropertySetPropertySetKeyOracleDatabaseExternalAuthentication,
}

var mappingGetPropertySetPropertySetKeyEnumLowerCase = map[string]GetPropertySetPropertySetKeyEnum{
	"apex_document_generator": GetPropertySetPropertySetKeyApexDocumentGenerator,
	"apex":                    GetPropertySetPropertySetKeyApex,
	"apex_fa_integration":     GetPropertySetPropertySetKeyApexFaIntegration,
	"oracle_database_external_authentication": GetPropertySetPropertySetKeyOracleDatabaseExternalAuthentication,
}

// GetGetPropertySetPropertySetKeyEnumValues Enumerates the set of values for GetPropertySetPropertySetKeyEnum
func GetGetPropertySetPropertySetKeyEnumValues() []GetPropertySetPropertySetKeyEnum {
	values := make([]GetPropertySetPropertySetKeyEnum, 0)
	for _, v := range mappingGetPropertySetPropertySetKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetGetPropertySetPropertySetKeyEnumStringValues Enumerates the set of values in String for GetPropertySetPropertySetKeyEnum
func GetGetPropertySetPropertySetKeyEnumStringValues() []string {
	return []string{
		"APEX_DOCUMENT_GENERATOR",
		"APEX",
		"APEX_FA_INTEGRATION",
		"ORACLE_DATABASE_EXTERNAL_AUTHENTICATION",
	}
}

// GetMappingGetPropertySetPropertySetKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetPropertySetPropertySetKeyEnum(val string) (GetPropertySetPropertySetKeyEnum, bool) {
	enum, ok := mappingGetPropertySetPropertySetKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
