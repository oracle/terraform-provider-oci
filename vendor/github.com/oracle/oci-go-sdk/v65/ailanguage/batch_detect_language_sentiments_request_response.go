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

// BatchDetectLanguageSentimentsRequest wrapper for the BatchDetectLanguageSentiments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/ailanguage/BatchDetectLanguageSentiments.go.html to see an example of how to use BatchDetectLanguageSentimentsRequest.
type BatchDetectLanguageSentimentsRequest struct {

	// The details to make sentiment detect call.
	BatchDetectLanguageSentimentsDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Set this parameter for sentence and aspect level sentiment analysis.
	// Allowed values are:
	//    - ASPECT
	//    - SENTENCE
	Level []BatchDetectLanguageSentimentsLevelEnum `contributesTo:"query" name:"level" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request BatchDetectLanguageSentimentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request BatchDetectLanguageSentimentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request BatchDetectLanguageSentimentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request BatchDetectLanguageSentimentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request BatchDetectLanguageSentimentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Level {
		if _, ok := GetMappingBatchDetectLanguageSentimentsLevelEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Level: %s. Supported values are: %s.", val, strings.Join(GetBatchDetectLanguageSentimentsLevelEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BatchDetectLanguageSentimentsResponse wrapper for the BatchDetectLanguageSentiments operation
type BatchDetectLanguageSentimentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The BatchDetectLanguageSentimentsResult instance
	BatchDetectLanguageSentimentsResult `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response BatchDetectLanguageSentimentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response BatchDetectLanguageSentimentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// BatchDetectLanguageSentimentsLevelEnum Enum with underlying type: string
type BatchDetectLanguageSentimentsLevelEnum string

// Set of constants representing the allowable values for BatchDetectLanguageSentimentsLevelEnum
const (
	BatchDetectLanguageSentimentsLevelAspect   BatchDetectLanguageSentimentsLevelEnum = "ASPECT"
	BatchDetectLanguageSentimentsLevelSentence BatchDetectLanguageSentimentsLevelEnum = "SENTENCE"
)

var mappingBatchDetectLanguageSentimentsLevelEnum = map[string]BatchDetectLanguageSentimentsLevelEnum{
	"ASPECT":   BatchDetectLanguageSentimentsLevelAspect,
	"SENTENCE": BatchDetectLanguageSentimentsLevelSentence,
}

var mappingBatchDetectLanguageSentimentsLevelEnumLowerCase = map[string]BatchDetectLanguageSentimentsLevelEnum{
	"aspect":   BatchDetectLanguageSentimentsLevelAspect,
	"sentence": BatchDetectLanguageSentimentsLevelSentence,
}

// GetBatchDetectLanguageSentimentsLevelEnumValues Enumerates the set of values for BatchDetectLanguageSentimentsLevelEnum
func GetBatchDetectLanguageSentimentsLevelEnumValues() []BatchDetectLanguageSentimentsLevelEnum {
	values := make([]BatchDetectLanguageSentimentsLevelEnum, 0)
	for _, v := range mappingBatchDetectLanguageSentimentsLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchDetectLanguageSentimentsLevelEnumStringValues Enumerates the set of values in String for BatchDetectLanguageSentimentsLevelEnum
func GetBatchDetectLanguageSentimentsLevelEnumStringValues() []string {
	return []string{
		"ASPECT",
		"SENTENCE",
	}
}

// GetMappingBatchDetectLanguageSentimentsLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchDetectLanguageSentimentsLevelEnum(val string) (BatchDetectLanguageSentimentsLevelEnum, bool) {
	enum, ok := mappingBatchDetectLanguageSentimentsLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
