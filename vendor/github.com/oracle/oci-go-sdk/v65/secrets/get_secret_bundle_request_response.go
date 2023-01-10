// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package secrets

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetSecretBundleRequest wrapper for the GetSecretBundle operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/secrets/GetSecretBundle.go.html to see an example of how to use GetSecretBundleRequest.
type GetSecretBundleRequest struct {

	// The OCID of the secret.
	SecretId *string `mandatory:"true" contributesTo:"path" name:"secretId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The version number of the secret.
	VersionNumber *int64 `mandatory:"false" contributesTo:"query" name:"versionNumber"`

	// The name of the secret. (This might be referred to as the name of the secret version. Names are unique across the different versions of a secret.)
	SecretVersionName *string `mandatory:"false" contributesTo:"query" name:"secretVersionName"`

	// The rotation state of the secret version.
	Stage GetSecretBundleStageEnum `mandatory:"false" contributesTo:"query" name:"stage" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetSecretBundleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetSecretBundleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetSecretBundleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetSecretBundleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetSecretBundleRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetSecretBundleStageEnum(string(request.Stage)); !ok && request.Stage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stage: %s. Supported values are: %s.", request.Stage, strings.Join(GetGetSecretBundleStageEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetSecretBundleResponse wrapper for the GetSecretBundle operation
type GetSecretBundleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The SecretBundle instance
	SecretBundle `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetSecretBundleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetSecretBundleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetSecretBundleStageEnum Enum with underlying type: string
type GetSecretBundleStageEnum string

// Set of constants representing the allowable values for GetSecretBundleStageEnum
const (
	GetSecretBundleStageCurrent    GetSecretBundleStageEnum = "CURRENT"
	GetSecretBundleStagePending    GetSecretBundleStageEnum = "PENDING"
	GetSecretBundleStageLatest     GetSecretBundleStageEnum = "LATEST"
	GetSecretBundleStagePrevious   GetSecretBundleStageEnum = "PREVIOUS"
	GetSecretBundleStageDeprecated GetSecretBundleStageEnum = "DEPRECATED"
)

var mappingGetSecretBundleStageEnum = map[string]GetSecretBundleStageEnum{
	"CURRENT":    GetSecretBundleStageCurrent,
	"PENDING":    GetSecretBundleStagePending,
	"LATEST":     GetSecretBundleStageLatest,
	"PREVIOUS":   GetSecretBundleStagePrevious,
	"DEPRECATED": GetSecretBundleStageDeprecated,
}

var mappingGetSecretBundleStageEnumLowerCase = map[string]GetSecretBundleStageEnum{
	"current":    GetSecretBundleStageCurrent,
	"pending":    GetSecretBundleStagePending,
	"latest":     GetSecretBundleStageLatest,
	"previous":   GetSecretBundleStagePrevious,
	"deprecated": GetSecretBundleStageDeprecated,
}

// GetGetSecretBundleStageEnumValues Enumerates the set of values for GetSecretBundleStageEnum
func GetGetSecretBundleStageEnumValues() []GetSecretBundleStageEnum {
	values := make([]GetSecretBundleStageEnum, 0)
	for _, v := range mappingGetSecretBundleStageEnum {
		values = append(values, v)
	}
	return values
}

// GetGetSecretBundleStageEnumStringValues Enumerates the set of values in String for GetSecretBundleStageEnum
func GetGetSecretBundleStageEnumStringValues() []string {
	return []string{
		"CURRENT",
		"PENDING",
		"LATEST",
		"PREVIOUS",
		"DEPRECATED",
	}
}

// GetMappingGetSecretBundleStageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetSecretBundleStageEnum(val string) (GetSecretBundleStageEnum, bool) {
	enum, ok := mappingGetSecretBundleStageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
