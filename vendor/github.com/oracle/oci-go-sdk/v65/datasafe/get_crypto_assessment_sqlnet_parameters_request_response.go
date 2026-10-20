// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetCryptoAssessmentSqlnetParametersRequest wrapper for the GetCryptoAssessmentSqlnetParameters operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/GetCryptoAssessmentSqlnetParameters.go.html to see an example of how to use GetCryptoAssessmentSqlnetParametersRequest.
type GetCryptoAssessmentSqlnetParametersRequest struct {

	// The OCID of the crypto assessment.
	CryptoAssessmentId *string `mandatory:"true" contributesTo:"path" name:"cryptoAssessmentId"`

	// A filter to return only the SQLNET parameter with the specified name.
	Parameter *string `mandatory:"false" contributesTo:"query" name:"parameter"`

	// Filters SQLNET parameters by quantum-readiness category.
	QuantumReadiness GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum `mandatory:"false" contributesTo:"query" name:"quantumReadiness" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCryptoAssessmentSqlnetParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCryptoAssessmentSqlnetParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetCryptoAssessmentSqlnetParametersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCryptoAssessmentSqlnetParametersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetCryptoAssessmentSqlnetParametersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetCryptoAssessmentSqlnetParametersQuantumReadinessEnum(string(request.QuantumReadiness)); !ok && request.QuantumReadiness != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QuantumReadiness: %s. Supported values are: %s.", request.QuantumReadiness, strings.Join(GetGetCryptoAssessmentSqlnetParametersQuantumReadinessEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetCryptoAssessmentSqlnetParametersResponse wrapper for the GetCryptoAssessmentSqlnetParameters operation
type GetCryptoAssessmentSqlnetParametersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentSqlnetParameters instances
	CryptoAssessmentSqlnetParameters `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCryptoAssessmentSqlnetParametersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCryptoAssessmentSqlnetParametersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum Enum with underlying type: string
type GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum string

// Set of constants representing the allowable values for GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum
const (
	GetCryptoAssessmentSqlnetParametersQuantumReadinessResistant     GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum = "RESISTANT"
	GetCryptoAssessmentSqlnetParametersQuantumReadinessNotResistant  GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum = "NOT_RESISTANT"
	GetCryptoAssessmentSqlnetParametersQuantumReadinessNotAvailable  GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum = "NOT_AVAILABLE"
	GetCryptoAssessmentSqlnetParametersQuantumReadinessNotApplicable GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum = "NOT_APPLICABLE"
	GetCryptoAssessmentSqlnetParametersQuantumReadinessNotSupported  GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum = "NOT_SUPPORTED"
)

var mappingGetCryptoAssessmentSqlnetParametersQuantumReadinessEnum = map[string]GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum{
	"RESISTANT":      GetCryptoAssessmentSqlnetParametersQuantumReadinessResistant,
	"NOT_RESISTANT":  GetCryptoAssessmentSqlnetParametersQuantumReadinessNotResistant,
	"NOT_AVAILABLE":  GetCryptoAssessmentSqlnetParametersQuantumReadinessNotAvailable,
	"NOT_APPLICABLE": GetCryptoAssessmentSqlnetParametersQuantumReadinessNotApplicable,
	"NOT_SUPPORTED":  GetCryptoAssessmentSqlnetParametersQuantumReadinessNotSupported,
}

var mappingGetCryptoAssessmentSqlnetParametersQuantumReadinessEnumLowerCase = map[string]GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum{
	"resistant":      GetCryptoAssessmentSqlnetParametersQuantumReadinessResistant,
	"not_resistant":  GetCryptoAssessmentSqlnetParametersQuantumReadinessNotResistant,
	"not_available":  GetCryptoAssessmentSqlnetParametersQuantumReadinessNotAvailable,
	"not_applicable": GetCryptoAssessmentSqlnetParametersQuantumReadinessNotApplicable,
	"not_supported":  GetCryptoAssessmentSqlnetParametersQuantumReadinessNotSupported,
}

// GetGetCryptoAssessmentSqlnetParametersQuantumReadinessEnumValues Enumerates the set of values for GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum
func GetGetCryptoAssessmentSqlnetParametersQuantumReadinessEnumValues() []GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum {
	values := make([]GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum, 0)
	for _, v := range mappingGetCryptoAssessmentSqlnetParametersQuantumReadinessEnum {
		values = append(values, v)
	}
	return values
}

// GetGetCryptoAssessmentSqlnetParametersQuantumReadinessEnumStringValues Enumerates the set of values in String for GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum
func GetGetCryptoAssessmentSqlnetParametersQuantumReadinessEnumStringValues() []string {
	return []string{
		"RESISTANT",
		"NOT_RESISTANT",
		"NOT_AVAILABLE",
		"NOT_APPLICABLE",
		"NOT_SUPPORTED",
	}
}

// GetMappingGetCryptoAssessmentSqlnetParametersQuantumReadinessEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetCryptoAssessmentSqlnetParametersQuantumReadinessEnum(val string) (GetCryptoAssessmentSqlnetParametersQuantumReadinessEnum, bool) {
	enum, ok := mappingGetCryptoAssessmentSqlnetParametersQuantumReadinessEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
