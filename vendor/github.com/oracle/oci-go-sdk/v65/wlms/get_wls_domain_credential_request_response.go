// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package wlms

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetWlsDomainCredentialRequest wrapper for the GetWlsDomainCredential operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/GetWlsDomainCredential.go.html to see an example of how to use GetWlsDomainCredentialRequest.
type GetWlsDomainCredentialRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"true" contributesTo:"path" name:"wlsDomainId"`

	// The type of the credentials.
	CredentialType GetWlsDomainCredentialCredentialTypeEnum `mandatory:"true" contributesTo:"path" name:"credentialType"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetWlsDomainCredentialRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetWlsDomainCredentialRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetWlsDomainCredentialRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetWlsDomainCredentialRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetWlsDomainCredentialRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetWlsDomainCredentialCredentialTypeEnum(string(request.CredentialType)); !ok && request.CredentialType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CredentialType: %s. Supported values are: %s.", request.CredentialType, strings.Join(GetGetWlsDomainCredentialCredentialTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetWlsDomainCredentialResponse wrapper for the GetWlsDomainCredential operation
type GetWlsDomainCredentialResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The WlsDomainCredential instance
	WlsDomainCredential `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetWlsDomainCredentialResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetWlsDomainCredentialResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetWlsDomainCredentialCredentialTypeEnum Enum with underlying type: string
type GetWlsDomainCredentialCredentialTypeEnum string

// Set of constants representing the allowable values for GetWlsDomainCredentialCredentialTypeEnum
const (
	GetWlsDomainCredentialCredentialTypeWeblogicadminuser GetWlsDomainCredentialCredentialTypeEnum = "weblogicAdminUser"
	GetWlsDomainCredentialCredentialTypeNodemanageruser   GetWlsDomainCredentialCredentialTypeEnum = "nodemanagerUser"
)

var mappingGetWlsDomainCredentialCredentialTypeEnum = map[string]GetWlsDomainCredentialCredentialTypeEnum{
	"weblogicAdminUser": GetWlsDomainCredentialCredentialTypeWeblogicadminuser,
	"nodemanagerUser":   GetWlsDomainCredentialCredentialTypeNodemanageruser,
}

var mappingGetWlsDomainCredentialCredentialTypeEnumLowerCase = map[string]GetWlsDomainCredentialCredentialTypeEnum{
	"weblogicadminuser": GetWlsDomainCredentialCredentialTypeWeblogicadminuser,
	"nodemanageruser":   GetWlsDomainCredentialCredentialTypeNodemanageruser,
}

// GetGetWlsDomainCredentialCredentialTypeEnumValues Enumerates the set of values for GetWlsDomainCredentialCredentialTypeEnum
func GetGetWlsDomainCredentialCredentialTypeEnumValues() []GetWlsDomainCredentialCredentialTypeEnum {
	values := make([]GetWlsDomainCredentialCredentialTypeEnum, 0)
	for _, v := range mappingGetWlsDomainCredentialCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetWlsDomainCredentialCredentialTypeEnumStringValues Enumerates the set of values in String for GetWlsDomainCredentialCredentialTypeEnum
func GetGetWlsDomainCredentialCredentialTypeEnumStringValues() []string {
	return []string{
		"weblogicAdminUser",
		"nodemanagerUser",
	}
}

// GetMappingGetWlsDomainCredentialCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetWlsDomainCredentialCredentialTypeEnum(val string) (GetWlsDomainCredentialCredentialTypeEnum, bool) {
	enum, ok := mappingGetWlsDomainCredentialCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
