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

// UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest wrapper for the UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasetoolsruntime/UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal.go.html to see an example of how to use UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest.
type UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools database API gateway config.
	DatabaseToolsDatabaseApiGatewayConfigId *string `mandatory:"true" contributesTo:"path" name:"databaseToolsDatabaseApiGatewayConfigId"`

	// The key of the global config.
	GlobalKey UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum `mandatory:"true" contributesTo:"path" name:"globalKey"`

	// The information to be updated.
	UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalDetails `contributesTo:"body"`

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

func (request UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum(string(request.GlobalKey)); !ok && request.GlobalKey != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GlobalKey: %s. Supported values are: %s.", request.GlobalKey, strings.Join(GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse wrapper for the UpdateDatabaseToolsDatabaseApiGatewayConfigGlobal operation
type UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DatabaseToolsDatabaseApiGatewayConfigGlobal instance
	DatabaseToolsDatabaseApiGatewayConfigGlobal `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum Enum with underlying type: string
type UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum string

// Set of constants representing the allowable values for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum
const (
	UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeySettings UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum = "SETTINGS"
)

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum{
	"SETTINGS": UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeySettings,
}

var mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnumLowerCase = map[string]UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum{
	"settings": UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeySettings,
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnumValues Enumerates the set of values for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnumValues() []UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum {
	values := make([]UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum, 0)
	for _, v := range mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnumStringValues Enumerates the set of values in String for UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum
func GetUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnumStringValues() []string {
	return []string{
		"SETTINGS",
	}
}

// GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum(val string) (UpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnum, bool) {
	enum, ok := mappingUpdateDatabaseToolsDatabaseApiGatewayConfigGlobalGlobalKeyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
