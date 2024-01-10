// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jmsjavadownloads

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetJavaLicenseRequest wrapper for the GetJavaLicense operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jmsjavadownloads/GetJavaLicense.go.html to see an example of how to use GetJavaLicenseRequest.
type GetJavaLicenseRequest struct {

	// Unique Java license type.
	LicenseType GetJavaLicenseLicenseTypeEnum `mandatory:"true" contributesTo:"path" name:"licenseType"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetJavaLicenseRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetJavaLicenseRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetJavaLicenseRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetJavaLicenseRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetJavaLicenseRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetJavaLicenseLicenseTypeEnum(string(request.LicenseType)); !ok && request.LicenseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseType: %s. Supported values are: %s.", request.LicenseType, strings.Join(GetGetJavaLicenseLicenseTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetJavaLicenseResponse wrapper for the GetJavaLicense operation
type GetJavaLicenseResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The JavaLicense instance
	JavaLicense `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetJavaLicenseResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetJavaLicenseResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetJavaLicenseLicenseTypeEnum Enum with underlying type: string
type GetJavaLicenseLicenseTypeEnum string

// Set of constants representing the allowable values for GetJavaLicenseLicenseTypeEnum
const (
	GetJavaLicenseLicenseTypeOtn        GetJavaLicenseLicenseTypeEnum = "OTN"
	GetJavaLicenseLicenseTypeNftc       GetJavaLicenseLicenseTypeEnum = "NFTC"
	GetJavaLicenseLicenseTypeRestricted GetJavaLicenseLicenseTypeEnum = "RESTRICTED"
)

var mappingGetJavaLicenseLicenseTypeEnum = map[string]GetJavaLicenseLicenseTypeEnum{
	"OTN":        GetJavaLicenseLicenseTypeOtn,
	"NFTC":       GetJavaLicenseLicenseTypeNftc,
	"RESTRICTED": GetJavaLicenseLicenseTypeRestricted,
}

var mappingGetJavaLicenseLicenseTypeEnumLowerCase = map[string]GetJavaLicenseLicenseTypeEnum{
	"otn":        GetJavaLicenseLicenseTypeOtn,
	"nftc":       GetJavaLicenseLicenseTypeNftc,
	"restricted": GetJavaLicenseLicenseTypeRestricted,
}

// GetGetJavaLicenseLicenseTypeEnumValues Enumerates the set of values for GetJavaLicenseLicenseTypeEnum
func GetGetJavaLicenseLicenseTypeEnumValues() []GetJavaLicenseLicenseTypeEnum {
	values := make([]GetJavaLicenseLicenseTypeEnum, 0)
	for _, v := range mappingGetJavaLicenseLicenseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetJavaLicenseLicenseTypeEnumStringValues Enumerates the set of values in String for GetJavaLicenseLicenseTypeEnum
func GetGetJavaLicenseLicenseTypeEnumStringValues() []string {
	return []string{
		"OTN",
		"NFTC",
		"RESTRICTED",
	}
}

// GetMappingGetJavaLicenseLicenseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetJavaLicenseLicenseTypeEnum(val string) (GetJavaLicenseLicenseTypeEnum, bool) {
	enum, ok := mappingGetJavaLicenseLicenseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
