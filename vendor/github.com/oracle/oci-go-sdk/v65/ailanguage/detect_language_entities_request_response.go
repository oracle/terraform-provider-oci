// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ailanguage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// DetectLanguageEntitiesRequest wrapper for the DetectLanguageEntities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/DetectLanguageEntities.go.html to see an example of how to use DetectLanguageEntitiesRequest.
type DetectLanguageEntitiesRequest struct {

	// The details to make a Entity detect call.
	DetectLanguageEntitiesDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Named Entity Recognition model versions. By default user will get output from V2.1 implementation.
	ModelVersion DetectLanguageEntitiesModelVersionEnum `mandatory:"false" contributesTo:"query" name:"modelVersion" omitEmpty:"true"`

	// If this parameter is set to true, you only get PII (Personally identifiable information) entities
	// like PhoneNumber, Email, Person, and so on. Default value is false.
	IsPii *bool `mandatory:"false" contributesTo:"query" name:"isPii"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DetectLanguageEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DetectLanguageEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DetectLanguageEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DetectLanguageEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DetectLanguageEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDetectLanguageEntitiesModelVersionEnum(string(request.ModelVersion)); !ok && request.ModelVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelVersion: %s. Supported values are: %s.", request.ModelVersion, strings.Join(GetDetectLanguageEntitiesModelVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DetectLanguageEntitiesResponse wrapper for the DetectLanguageEntities operation
type DetectLanguageEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DetectLanguageEntitiesResult instance
	DetectLanguageEntitiesResult `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// This API will be retired on Monday, 10 Oct 2023 00:00:00 GMT
	Sunset *string `presentIn:"header" name:"sunset"`
}

func (response DetectLanguageEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DetectLanguageEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// DetectLanguageEntitiesModelVersionEnum Enum with underlying type: string
type DetectLanguageEntitiesModelVersionEnum string

// Set of constants representing the allowable values for DetectLanguageEntitiesModelVersionEnum
const (
	DetectLanguageEntitiesModelVersionV21 DetectLanguageEntitiesModelVersionEnum = "V2_1"
	DetectLanguageEntitiesModelVersionV11 DetectLanguageEntitiesModelVersionEnum = "V1_1"
)

var mappingDetectLanguageEntitiesModelVersionEnum = map[string]DetectLanguageEntitiesModelVersionEnum{
	"V2_1": DetectLanguageEntitiesModelVersionV21,
	"V1_1": DetectLanguageEntitiesModelVersionV11,
}

var mappingDetectLanguageEntitiesModelVersionEnumLowerCase = map[string]DetectLanguageEntitiesModelVersionEnum{
	"v2_1": DetectLanguageEntitiesModelVersionV21,
	"v1_1": DetectLanguageEntitiesModelVersionV11,
}

// GetDetectLanguageEntitiesModelVersionEnumValues Enumerates the set of values for DetectLanguageEntitiesModelVersionEnum
func GetDetectLanguageEntitiesModelVersionEnumValues() []DetectLanguageEntitiesModelVersionEnum {
	values := make([]DetectLanguageEntitiesModelVersionEnum, 0)
	for _, v := range mappingDetectLanguageEntitiesModelVersionEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectLanguageEntitiesModelVersionEnumStringValues Enumerates the set of values in String for DetectLanguageEntitiesModelVersionEnum
func GetDetectLanguageEntitiesModelVersionEnumStringValues() []string {
	return []string{
		"V2_1",
		"V1_1",
	}
}

// GetMappingDetectLanguageEntitiesModelVersionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectLanguageEntitiesModelVersionEnum(val string) (DetectLanguageEntitiesModelVersionEnum, bool) {
	enum, ok := mappingDetectLanguageEntitiesModelVersionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
