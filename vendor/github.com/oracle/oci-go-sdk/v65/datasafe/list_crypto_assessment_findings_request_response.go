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

// ListCryptoAssessmentFindingsRequest wrapper for the ListCryptoAssessmentFindings operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCryptoAssessmentFindings.go.html to see an example of how to use ListCryptoAssessmentFindingsRequest.
type ListCryptoAssessmentFindingsRequest struct {

	// The OCID of the crypto assessment.
	CryptoAssessmentId *string `mandatory:"true" contributesTo:"path" name:"cryptoAssessmentId"`

	// A filter to return only findings with the specified finding key.
	FindingKey *string `mandatory:"false" contributesTo:"query" name:"findingKey"`

	// A filter to return only findings with the specified title.
	Title *string `mandatory:"false" contributesTo:"query" name:"title"`

	// A filter to return only findings in the specified category key.
	Category ListCryptoAssessmentFindingsCategoryEnum `mandatory:"false" contributesTo:"query" name:"category" omitEmpty:"true"`

	// A filter to return only findings with the specified status.
	Status ListCryptoAssessmentFindingsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return only findings that are or are not part of quantum-readiness checks.
	IsQuantumReadinessCheck *bool `mandatory:"false" contributesTo:"query" name:"isQuantumReadinessCheck"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCryptoAssessmentFindingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAssessmentFindingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAssessmentFindingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAssessmentFindingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAssessmentFindingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCryptoAssessmentFindingsCategoryEnum(string(request.Category)); !ok && request.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", request.Category, strings.Join(GetListCryptoAssessmentFindingsCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentFindingsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListCryptoAssessmentFindingsStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAssessmentFindingsResponse wrapper for the ListCryptoAssessmentFindings operation
type ListCryptoAssessmentFindingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentFindingCollection instances
	CryptoAssessmentFindingCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCryptoAssessmentFindingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAssessmentFindingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAssessmentFindingsCategoryEnum Enum with underlying type: string
type ListCryptoAssessmentFindingsCategoryEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingsCategoryEnum
const (
	ListCryptoAssessmentFindingsCategoryNetworkEncryption            ListCryptoAssessmentFindingsCategoryEnum = "NETWORK_ENCRYPTION"
	ListCryptoAssessmentFindingsCategoryDataEncryption               ListCryptoAssessmentFindingsCategoryEnum = "DATA_ENCRYPTION"
	ListCryptoAssessmentFindingsCategoryCertificatesAndKeyManagement ListCryptoAssessmentFindingsCategoryEnum = "CERTIFICATES_AND_KEY_MANAGEMENT"
	ListCryptoAssessmentFindingsCategoryBackupAndExportEncryption    ListCryptoAssessmentFindingsCategoryEnum = "BACKUP_AND_EXPORT_ENCRYPTION"
	ListCryptoAssessmentFindingsCategoryPostQuantumReadiness         ListCryptoAssessmentFindingsCategoryEnum = "POST_QUANTUM_READINESS"
	ListCryptoAssessmentFindingsCategoryNotSupported                 ListCryptoAssessmentFindingsCategoryEnum = "NOT_SUPPORTED"
)

var mappingListCryptoAssessmentFindingsCategoryEnum = map[string]ListCryptoAssessmentFindingsCategoryEnum{
	"NETWORK_ENCRYPTION":              ListCryptoAssessmentFindingsCategoryNetworkEncryption,
	"DATA_ENCRYPTION":                 ListCryptoAssessmentFindingsCategoryDataEncryption,
	"CERTIFICATES_AND_KEY_MANAGEMENT": ListCryptoAssessmentFindingsCategoryCertificatesAndKeyManagement,
	"BACKUP_AND_EXPORT_ENCRYPTION":    ListCryptoAssessmentFindingsCategoryBackupAndExportEncryption,
	"POST_QUANTUM_READINESS":          ListCryptoAssessmentFindingsCategoryPostQuantumReadiness,
	"NOT_SUPPORTED":                   ListCryptoAssessmentFindingsCategoryNotSupported,
}

var mappingListCryptoAssessmentFindingsCategoryEnumLowerCase = map[string]ListCryptoAssessmentFindingsCategoryEnum{
	"network_encryption":              ListCryptoAssessmentFindingsCategoryNetworkEncryption,
	"data_encryption":                 ListCryptoAssessmentFindingsCategoryDataEncryption,
	"certificates_and_key_management": ListCryptoAssessmentFindingsCategoryCertificatesAndKeyManagement,
	"backup_and_export_encryption":    ListCryptoAssessmentFindingsCategoryBackupAndExportEncryption,
	"post_quantum_readiness":          ListCryptoAssessmentFindingsCategoryPostQuantumReadiness,
	"not_supported":                   ListCryptoAssessmentFindingsCategoryNotSupported,
}

// GetListCryptoAssessmentFindingsCategoryEnumValues Enumerates the set of values for ListCryptoAssessmentFindingsCategoryEnum
func GetListCryptoAssessmentFindingsCategoryEnumValues() []ListCryptoAssessmentFindingsCategoryEnum {
	values := make([]ListCryptoAssessmentFindingsCategoryEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingsCategoryEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingsCategoryEnum
func GetListCryptoAssessmentFindingsCategoryEnumStringValues() []string {
	return []string{
		"NETWORK_ENCRYPTION",
		"DATA_ENCRYPTION",
		"CERTIFICATES_AND_KEY_MANAGEMENT",
		"BACKUP_AND_EXPORT_ENCRYPTION",
		"POST_QUANTUM_READINESS",
		"NOT_SUPPORTED",
	}
}

// GetMappingListCryptoAssessmentFindingsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingsCategoryEnum(val string) (ListCryptoAssessmentFindingsCategoryEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentFindingsStatusEnum Enum with underlying type: string
type ListCryptoAssessmentFindingsStatusEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingsStatusEnum
const (
	ListCryptoAssessmentFindingsStatusPass          ListCryptoAssessmentFindingsStatusEnum = "PASS"
	ListCryptoAssessmentFindingsStatusFail          ListCryptoAssessmentFindingsStatusEnum = "FAIL"
	ListCryptoAssessmentFindingsStatusError         ListCryptoAssessmentFindingsStatusEnum = "ERROR"
	ListCryptoAssessmentFindingsStatusEvaluate      ListCryptoAssessmentFindingsStatusEnum = "EVALUATE"
	ListCryptoAssessmentFindingsStatusNotAvailable  ListCryptoAssessmentFindingsStatusEnum = "NOT_AVAILABLE"
	ListCryptoAssessmentFindingsStatusNotApplicable ListCryptoAssessmentFindingsStatusEnum = "NOT_APPLICABLE"
	ListCryptoAssessmentFindingsStatusNotSupported  ListCryptoAssessmentFindingsStatusEnum = "NOT_SUPPORTED"
)

var mappingListCryptoAssessmentFindingsStatusEnum = map[string]ListCryptoAssessmentFindingsStatusEnum{
	"PASS":           ListCryptoAssessmentFindingsStatusPass,
	"FAIL":           ListCryptoAssessmentFindingsStatusFail,
	"ERROR":          ListCryptoAssessmentFindingsStatusError,
	"EVALUATE":       ListCryptoAssessmentFindingsStatusEvaluate,
	"NOT_AVAILABLE":  ListCryptoAssessmentFindingsStatusNotAvailable,
	"NOT_APPLICABLE": ListCryptoAssessmentFindingsStatusNotApplicable,
	"NOT_SUPPORTED":  ListCryptoAssessmentFindingsStatusNotSupported,
}

var mappingListCryptoAssessmentFindingsStatusEnumLowerCase = map[string]ListCryptoAssessmentFindingsStatusEnum{
	"pass":           ListCryptoAssessmentFindingsStatusPass,
	"fail":           ListCryptoAssessmentFindingsStatusFail,
	"error":          ListCryptoAssessmentFindingsStatusError,
	"evaluate":       ListCryptoAssessmentFindingsStatusEvaluate,
	"not_available":  ListCryptoAssessmentFindingsStatusNotAvailable,
	"not_applicable": ListCryptoAssessmentFindingsStatusNotApplicable,
	"not_supported":  ListCryptoAssessmentFindingsStatusNotSupported,
}

// GetListCryptoAssessmentFindingsStatusEnumValues Enumerates the set of values for ListCryptoAssessmentFindingsStatusEnum
func GetListCryptoAssessmentFindingsStatusEnumValues() []ListCryptoAssessmentFindingsStatusEnum {
	values := make([]ListCryptoAssessmentFindingsStatusEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingsStatusEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingsStatusEnum
func GetListCryptoAssessmentFindingsStatusEnumStringValues() []string {
	return []string{
		"PASS",
		"FAIL",
		"ERROR",
		"EVALUATE",
		"NOT_AVAILABLE",
		"NOT_APPLICABLE",
		"NOT_SUPPORTED",
	}
}

// GetMappingListCryptoAssessmentFindingsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingsStatusEnum(val string) (ListCryptoAssessmentFindingsStatusEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
