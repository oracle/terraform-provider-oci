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

// ListCryptoAssessmentCertificatesRequest wrapper for the ListCryptoAssessmentCertificates operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCryptoAssessmentCertificates.go.html to see an example of how to use ListCryptoAssessmentCertificatesRequest.
type ListCryptoAssessmentCertificatesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListCryptoAssessmentCertificatesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources associated with the specified crypto assessment OCID.
	AssessmentId *string `mandatory:"false" contributesTo:"query" name:"assessmentId"`

	// A filter to return targets from assessments of the specified type.
	AssessmentType CryptoAssessmentTypeEnum `mandatory:"false" contributesTo:"query" name:"assessmentType" omitEmpty:"true"`

	// A filter to return only inventory rows associated with the specified target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only resources associated with any of the specified target OCIDs.
	TargetIds []string `contributesTo:"query" name:"targetIds" collectionFormat:"multi"`

	// A filter to return only certificates of any of the specified types.
	CertificateType []CryptoAssessmentCertificateSummaryCertificateTypeEnum `contributesTo:"query" name:"certificateType" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only certificates with any of the specified statuses.
	Status []CryptoAssessmentCertificateSummaryStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only certificates with any of the specified public key types. Stored values are normalized forms such as RSA2048, RSA4096, or EC256.
	PublicKeyType []string `contributesTo:"query" name:"publicKeyType" collectionFormat:"multi"`

	// A filter to return only certificates whose signature algorithm contains any of the specified values. For example, use SHA1 to match SHA1-based certificate signatures.
	SignatureAlgorithm []string `contributesTo:"query" name:"signatureAlgorithm" collectionFormat:"multi"`

	// A filter to return only certificates in the specified expiry bucket. Supported values are 0_15, 15_30, 30_60, 60_90, and 90_PLUS.
	ExpiryBucket *string `mandatory:"false" contributesTo:"query" name:"expiryBucket"`

	// A filter to return certificates whose validTill timestamp is on or before the current time plus the specified number of days. Negative values are allowed and filter certificates that expired on or before that many days ago.
	DaysToExpiry *int `mandatory:"false" contributesTo:"query" name:"daysToExpiry"`

	// The field used to sort certificate results.
	SortBy ListCryptoAssessmentCertificatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListCryptoAssessmentCertificatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListCryptoAssessmentCertificatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAssessmentCertificatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAssessmentCertificatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAssessmentCertificatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAssessmentCertificatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCryptoAssessmentCertificatesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCryptoAssessmentCertificatesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(request.AssessmentType)); !ok && request.AssessmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessmentType: %s. Supported values are: %s.", request.AssessmentType, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	for _, val := range request.CertificateType {
		if _, ok := GetMappingCryptoAssessmentCertificateSummaryCertificateTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateType: %s. Supported values are: %s.", val, strings.Join(GetCryptoAssessmentCertificateSummaryCertificateTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.Status {
		if _, ok := GetMappingCryptoAssessmentCertificateSummaryStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetCryptoAssessmentCertificateSummaryStatusEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListCryptoAssessmentCertificatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCryptoAssessmentCertificatesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentCertificatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCryptoAssessmentCertificatesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAssessmentCertificatesResponse wrapper for the ListCryptoAssessmentCertificates operation
type ListCryptoAssessmentCertificatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentCertificateCollection instances
	CryptoAssessmentCertificateCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCryptoAssessmentCertificatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAssessmentCertificatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAssessmentCertificatesAccessLevelEnum Enum with underlying type: string
type ListCryptoAssessmentCertificatesAccessLevelEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentCertificatesAccessLevelEnum
const (
	ListCryptoAssessmentCertificatesAccessLevelRestricted ListCryptoAssessmentCertificatesAccessLevelEnum = "RESTRICTED"
	ListCryptoAssessmentCertificatesAccessLevelAccessible ListCryptoAssessmentCertificatesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCryptoAssessmentCertificatesAccessLevelEnum = map[string]ListCryptoAssessmentCertificatesAccessLevelEnum{
	"RESTRICTED": ListCryptoAssessmentCertificatesAccessLevelRestricted,
	"ACCESSIBLE": ListCryptoAssessmentCertificatesAccessLevelAccessible,
}

var mappingListCryptoAssessmentCertificatesAccessLevelEnumLowerCase = map[string]ListCryptoAssessmentCertificatesAccessLevelEnum{
	"restricted": ListCryptoAssessmentCertificatesAccessLevelRestricted,
	"accessible": ListCryptoAssessmentCertificatesAccessLevelAccessible,
}

// GetListCryptoAssessmentCertificatesAccessLevelEnumValues Enumerates the set of values for ListCryptoAssessmentCertificatesAccessLevelEnum
func GetListCryptoAssessmentCertificatesAccessLevelEnumValues() []ListCryptoAssessmentCertificatesAccessLevelEnum {
	values := make([]ListCryptoAssessmentCertificatesAccessLevelEnum, 0)
	for _, v := range mappingListCryptoAssessmentCertificatesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentCertificatesAccessLevelEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentCertificatesAccessLevelEnum
func GetListCryptoAssessmentCertificatesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCryptoAssessmentCertificatesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentCertificatesAccessLevelEnum(val string) (ListCryptoAssessmentCertificatesAccessLevelEnum, bool) {
	enum, ok := mappingListCryptoAssessmentCertificatesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentCertificatesSortByEnum Enum with underlying type: string
type ListCryptoAssessmentCertificatesSortByEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentCertificatesSortByEnum
const (
	ListCryptoAssessmentCertificatesSortByTimevaliduntil ListCryptoAssessmentCertificatesSortByEnum = "timeValidUntil"
	ListCryptoAssessmentCertificatesSortBySubject        ListCryptoAssessmentCertificatesSortByEnum = "subject"
)

var mappingListCryptoAssessmentCertificatesSortByEnum = map[string]ListCryptoAssessmentCertificatesSortByEnum{
	"timeValidUntil": ListCryptoAssessmentCertificatesSortByTimevaliduntil,
	"subject":        ListCryptoAssessmentCertificatesSortBySubject,
}

var mappingListCryptoAssessmentCertificatesSortByEnumLowerCase = map[string]ListCryptoAssessmentCertificatesSortByEnum{
	"timevaliduntil": ListCryptoAssessmentCertificatesSortByTimevaliduntil,
	"subject":        ListCryptoAssessmentCertificatesSortBySubject,
}

// GetListCryptoAssessmentCertificatesSortByEnumValues Enumerates the set of values for ListCryptoAssessmentCertificatesSortByEnum
func GetListCryptoAssessmentCertificatesSortByEnumValues() []ListCryptoAssessmentCertificatesSortByEnum {
	values := make([]ListCryptoAssessmentCertificatesSortByEnum, 0)
	for _, v := range mappingListCryptoAssessmentCertificatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentCertificatesSortByEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentCertificatesSortByEnum
func GetListCryptoAssessmentCertificatesSortByEnumStringValues() []string {
	return []string{
		"timeValidUntil",
		"subject",
	}
}

// GetMappingListCryptoAssessmentCertificatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentCertificatesSortByEnum(val string) (ListCryptoAssessmentCertificatesSortByEnum, bool) {
	enum, ok := mappingListCryptoAssessmentCertificatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentCertificatesSortOrderEnum Enum with underlying type: string
type ListCryptoAssessmentCertificatesSortOrderEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentCertificatesSortOrderEnum
const (
	ListCryptoAssessmentCertificatesSortOrderAsc  ListCryptoAssessmentCertificatesSortOrderEnum = "ASC"
	ListCryptoAssessmentCertificatesSortOrderDesc ListCryptoAssessmentCertificatesSortOrderEnum = "DESC"
)

var mappingListCryptoAssessmentCertificatesSortOrderEnum = map[string]ListCryptoAssessmentCertificatesSortOrderEnum{
	"ASC":  ListCryptoAssessmentCertificatesSortOrderAsc,
	"DESC": ListCryptoAssessmentCertificatesSortOrderDesc,
}

var mappingListCryptoAssessmentCertificatesSortOrderEnumLowerCase = map[string]ListCryptoAssessmentCertificatesSortOrderEnum{
	"asc":  ListCryptoAssessmentCertificatesSortOrderAsc,
	"desc": ListCryptoAssessmentCertificatesSortOrderDesc,
}

// GetListCryptoAssessmentCertificatesSortOrderEnumValues Enumerates the set of values for ListCryptoAssessmentCertificatesSortOrderEnum
func GetListCryptoAssessmentCertificatesSortOrderEnumValues() []ListCryptoAssessmentCertificatesSortOrderEnum {
	values := make([]ListCryptoAssessmentCertificatesSortOrderEnum, 0)
	for _, v := range mappingListCryptoAssessmentCertificatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentCertificatesSortOrderEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentCertificatesSortOrderEnum
func GetListCryptoAssessmentCertificatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCryptoAssessmentCertificatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentCertificatesSortOrderEnum(val string) (ListCryptoAssessmentCertificatesSortOrderEnum, bool) {
	enum, ok := mappingListCryptoAssessmentCertificatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
