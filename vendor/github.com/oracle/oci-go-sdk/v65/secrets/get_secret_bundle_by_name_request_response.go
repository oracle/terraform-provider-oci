// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package secrets

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetSecretBundleByNameRequest wrapper for the GetSecretBundleByName operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/secrets/GetSecretBundleByName.go.html to see an example of how to use GetSecretBundleByNameRequest.
type GetSecretBundleByNameRequest struct {

	// A user-friendly name for the secret. Secret names are unique within a vault. Secret names are case-sensitive.
	SecretName *string `mandatory:"true" contributesTo:"query" name:"secretName"`

	// The OCID of the vault that contains the secret.
	VaultId *string `mandatory:"true" contributesTo:"query" name:"vaultId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The version number of the secret.
	VersionNumber *int64 `mandatory:"false" contributesTo:"query" name:"versionNumber"`

	// The name of the secret. (This might be referred to as the name of the secret version. Names are unique across the different versions of a secret.)
	SecretVersionName *string `mandatory:"false" contributesTo:"query" name:"secretVersionName"`

	// The rotation state of the secret version.
	Stage GetSecretBundleByNameStageEnum `mandatory:"false" contributesTo:"query" name:"stage" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetSecretBundleByNameRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetSecretBundleByNameRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetSecretBundleByNameRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetSecretBundleByNameRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetSecretBundleByNameRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetSecretBundleByNameStageEnum(string(request.Stage)); !ok && request.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", request.Stage, strings.Join(GetGetSecretBundleByNameStageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetSecretBundleByNameResponse wrapper for the GetSecretBundleByName operation
type GetSecretBundleByNameResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SecretBundle instance
	SecretBundle `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetSecretBundleByNameResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetSecretBundleByNameResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetSecretBundleByNameStageEnum Enum with underlying type: string
type GetSecretBundleByNameStageEnum string

// Set of constants representing the allowable values for GetSecretBundleByNameStageEnum
const (
	GetSecretBundleByNameStageCurrent    GetSecretBundleByNameStageEnum = "CURRENT"
	GetSecretBundleByNameStagePending    GetSecretBundleByNameStageEnum = "PENDING"
	GetSecretBundleByNameStageLatest     GetSecretBundleByNameStageEnum = "LATEST"
	GetSecretBundleByNameStagePrevious   GetSecretBundleByNameStageEnum = "PREVIOUS"
	GetSecretBundleByNameStageDeprecated GetSecretBundleByNameStageEnum = "DEPRECATED"
)

var mappingGetSecretBundleByNameStageEnum = map[string]GetSecretBundleByNameStageEnum{
	"CURRENT":    GetSecretBundleByNameStageCurrent,
	"PENDING":    GetSecretBundleByNameStagePending,
	"LATEST":     GetSecretBundleByNameStageLatest,
	"PREVIOUS":   GetSecretBundleByNameStagePrevious,
	"DEPRECATED": GetSecretBundleByNameStageDeprecated,
}

var mappingGetSecretBundleByNameStageEnumLowerCase = map[string]GetSecretBundleByNameStageEnum{
	"current":    GetSecretBundleByNameStageCurrent,
	"pending":    GetSecretBundleByNameStagePending,
	"latest":     GetSecretBundleByNameStageLatest,
	"previous":   GetSecretBundleByNameStagePrevious,
	"deprecated": GetSecretBundleByNameStageDeprecated,
}

// GetGetSecretBundleByNameStageEnumValues Enumerates the set of values for GetSecretBundleByNameStageEnum
func GetGetSecretBundleByNameStageEnumValues() []GetSecretBundleByNameStageEnum {
	values := make([]GetSecretBundleByNameStageEnum, 0)
	for _, v := range mappingGetSecretBundleByNameStageEnum {
		values = append(values, v)
	}
	return values
}

// GetGetSecretBundleByNameStageEnumStringValues Enumerates the set of values in String for GetSecretBundleByNameStageEnum
func GetGetSecretBundleByNameStageEnumStringValues() []string {
	return []string{
		"CURRENT",
		"PENDING",
		"LATEST",
		"PREVIOUS",
		"DEPRECATED",
	}
}

// GetMappingGetSecretBundleByNameStageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetSecretBundleByNameStageEnum(val string) (GetSecretBundleByNameStageEnum, bool) {
	enum, ok := mappingGetSecretBundleByNameStageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
