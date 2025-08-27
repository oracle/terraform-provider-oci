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

// UpdateWlsDomainCredentialRequest wrapper for the UpdateWlsDomainCredential operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/wlms/UpdateWlsDomainCredential.go.html to see an example of how to use UpdateWlsDomainCredentialRequest.
type UpdateWlsDomainCredentialRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the WebLogic domain.
	WlsDomainId *string `mandatory:"true" contributesTo:"path" name:"wlsDomainId"`

	// The type of the credentials.
	CredentialType UpdateWlsDomainCredentialCredentialTypeEnum `mandatory:"true" contributesTo:"path" name:"credentialType"`

	// The WebLogic domain credentials.
	UpdateWlsDomainCredentialDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// ETag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the ETag you
	// provide matches the resource's current ETag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateWlsDomainCredentialRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateWlsDomainCredentialRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateWlsDomainCredentialRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateWlsDomainCredentialRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateWlsDomainCredentialRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateWlsDomainCredentialCredentialTypeEnum(string(request.CredentialType)); !ok && request.CredentialType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CredentialType: %s. Supported values are: %s.", request.CredentialType, strings.Join(GetUpdateWlsDomainCredentialCredentialTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateWlsDomainCredentialResponse wrapper for the UpdateWlsDomainCredential operation
type UpdateWlsDomainCredentialResponse struct {

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

func (response UpdateWlsDomainCredentialResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateWlsDomainCredentialResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UpdateWlsDomainCredentialCredentialTypeEnum Enum with underlying type: string
type UpdateWlsDomainCredentialCredentialTypeEnum string

// Set of constants representing the allowable values for UpdateWlsDomainCredentialCredentialTypeEnum
const (
	UpdateWlsDomainCredentialCredentialTypeWeblogicadminuser UpdateWlsDomainCredentialCredentialTypeEnum = "weblogicAdminUser"
	UpdateWlsDomainCredentialCredentialTypeNodemanageruser   UpdateWlsDomainCredentialCredentialTypeEnum = "nodemanagerUser"
)

var mappingUpdateWlsDomainCredentialCredentialTypeEnum = map[string]UpdateWlsDomainCredentialCredentialTypeEnum{
	"weblogicAdminUser": UpdateWlsDomainCredentialCredentialTypeWeblogicadminuser,
	"nodemanagerUser":   UpdateWlsDomainCredentialCredentialTypeNodemanageruser,
}

var mappingUpdateWlsDomainCredentialCredentialTypeEnumLowerCase = map[string]UpdateWlsDomainCredentialCredentialTypeEnum{
	"weblogicadminuser": UpdateWlsDomainCredentialCredentialTypeWeblogicadminuser,
	"nodemanageruser":   UpdateWlsDomainCredentialCredentialTypeNodemanageruser,
}

// GetUpdateWlsDomainCredentialCredentialTypeEnumValues Enumerates the set of values for UpdateWlsDomainCredentialCredentialTypeEnum
func GetUpdateWlsDomainCredentialCredentialTypeEnumValues() []UpdateWlsDomainCredentialCredentialTypeEnum {
	values := make([]UpdateWlsDomainCredentialCredentialTypeEnum, 0)
	for _, v := range mappingUpdateWlsDomainCredentialCredentialTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateWlsDomainCredentialCredentialTypeEnumStringValues Enumerates the set of values in String for UpdateWlsDomainCredentialCredentialTypeEnum
func GetUpdateWlsDomainCredentialCredentialTypeEnumStringValues() []string {
	return []string{
		"weblogicAdminUser",
		"nodemanagerUser",
	}
}

// GetMappingUpdateWlsDomainCredentialCredentialTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateWlsDomainCredentialCredentialTypeEnum(val string) (UpdateWlsDomainCredentialCredentialTypeEnum, bool) {
	enum, ok := mappingUpdateWlsDomainCredentialCredentialTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
