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

// ListCryptoAssessmentFindingAnalyticsRequest wrapper for the ListCryptoAssessmentFindingAnalytics operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCryptoAssessmentFindingAnalytics.go.html to see an example of how to use ListCryptoAssessmentFindingAnalyticsRequest.
type ListCryptoAssessmentFindingAnalyticsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListCryptoAssessmentFindingAnalyticsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only findings in the specified category key.
	Category ListCryptoAssessmentFindingAnalyticsCategoryEnum `mandatory:"false" contributesTo:"query" name:"category" omitEmpty:"true"`

	// A filter to return only findings with any of the specified finding keys.
	FindingKey []string `contributesTo:"query" name:"findingKey" collectionFormat:"multi"`

	// A filter to return only findings that are or are not part of quantum-readiness checks.
	IsQuantumReadinessCheck *bool `mandatory:"false" contributesTo:"query" name:"isQuantumReadinessCheck"`

	// The field used to sort finding analytics results.
	SortBy ListCryptoAssessmentFindingAnalyticsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListCryptoAssessmentFindingAnalyticsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListCryptoAssessmentFindingAnalyticsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAssessmentFindingAnalyticsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAssessmentFindingAnalyticsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAssessmentFindingAnalyticsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAssessmentFindingAnalyticsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCryptoAssessmentFindingAnalyticsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCryptoAssessmentFindingAnalyticsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentFindingAnalyticsCategoryEnum(string(request.Category)); !ok && request.Category != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Category: %s. Supported values are: %s.", request.Category, strings.Join(GetListCryptoAssessmentFindingAnalyticsCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentFindingAnalyticsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCryptoAssessmentFindingAnalyticsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentFindingAnalyticsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCryptoAssessmentFindingAnalyticsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAssessmentFindingAnalyticsResponse wrapper for the ListCryptoAssessmentFindingAnalytics operation
type ListCryptoAssessmentFindingAnalyticsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentFindingAnalyticsCollection instances
	CryptoAssessmentFindingAnalyticsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCryptoAssessmentFindingAnalyticsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAssessmentFindingAnalyticsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAssessmentFindingAnalyticsAccessLevelEnum Enum with underlying type: string
type ListCryptoAssessmentFindingAnalyticsAccessLevelEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingAnalyticsAccessLevelEnum
const (
	ListCryptoAssessmentFindingAnalyticsAccessLevelRestricted ListCryptoAssessmentFindingAnalyticsAccessLevelEnum = "RESTRICTED"
	ListCryptoAssessmentFindingAnalyticsAccessLevelAccessible ListCryptoAssessmentFindingAnalyticsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCryptoAssessmentFindingAnalyticsAccessLevelEnum = map[string]ListCryptoAssessmentFindingAnalyticsAccessLevelEnum{
	"RESTRICTED": ListCryptoAssessmentFindingAnalyticsAccessLevelRestricted,
	"ACCESSIBLE": ListCryptoAssessmentFindingAnalyticsAccessLevelAccessible,
}

var mappingListCryptoAssessmentFindingAnalyticsAccessLevelEnumLowerCase = map[string]ListCryptoAssessmentFindingAnalyticsAccessLevelEnum{
	"restricted": ListCryptoAssessmentFindingAnalyticsAccessLevelRestricted,
	"accessible": ListCryptoAssessmentFindingAnalyticsAccessLevelAccessible,
}

// GetListCryptoAssessmentFindingAnalyticsAccessLevelEnumValues Enumerates the set of values for ListCryptoAssessmentFindingAnalyticsAccessLevelEnum
func GetListCryptoAssessmentFindingAnalyticsAccessLevelEnumValues() []ListCryptoAssessmentFindingAnalyticsAccessLevelEnum {
	values := make([]ListCryptoAssessmentFindingAnalyticsAccessLevelEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingAnalyticsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingAnalyticsAccessLevelEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingAnalyticsAccessLevelEnum
func GetListCryptoAssessmentFindingAnalyticsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCryptoAssessmentFindingAnalyticsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingAnalyticsAccessLevelEnum(val string) (ListCryptoAssessmentFindingAnalyticsAccessLevelEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingAnalyticsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentFindingAnalyticsCategoryEnum Enum with underlying type: string
type ListCryptoAssessmentFindingAnalyticsCategoryEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingAnalyticsCategoryEnum
const (
	ListCryptoAssessmentFindingAnalyticsCategoryNetworkEncryption            ListCryptoAssessmentFindingAnalyticsCategoryEnum = "NETWORK_ENCRYPTION"
	ListCryptoAssessmentFindingAnalyticsCategoryDataEncryption               ListCryptoAssessmentFindingAnalyticsCategoryEnum = "DATA_ENCRYPTION"
	ListCryptoAssessmentFindingAnalyticsCategoryCertificatesAndKeyManagement ListCryptoAssessmentFindingAnalyticsCategoryEnum = "CERTIFICATES_AND_KEY_MANAGEMENT"
	ListCryptoAssessmentFindingAnalyticsCategoryBackupAndExportEncryption    ListCryptoAssessmentFindingAnalyticsCategoryEnum = "BACKUP_AND_EXPORT_ENCRYPTION"
	ListCryptoAssessmentFindingAnalyticsCategoryPostQuantumReadiness         ListCryptoAssessmentFindingAnalyticsCategoryEnum = "POST_QUANTUM_READINESS"
	ListCryptoAssessmentFindingAnalyticsCategoryNotSupported                 ListCryptoAssessmentFindingAnalyticsCategoryEnum = "NOT_SUPPORTED"
)

var mappingListCryptoAssessmentFindingAnalyticsCategoryEnum = map[string]ListCryptoAssessmentFindingAnalyticsCategoryEnum{
	"NETWORK_ENCRYPTION":              ListCryptoAssessmentFindingAnalyticsCategoryNetworkEncryption,
	"DATA_ENCRYPTION":                 ListCryptoAssessmentFindingAnalyticsCategoryDataEncryption,
	"CERTIFICATES_AND_KEY_MANAGEMENT": ListCryptoAssessmentFindingAnalyticsCategoryCertificatesAndKeyManagement,
	"BACKUP_AND_EXPORT_ENCRYPTION":    ListCryptoAssessmentFindingAnalyticsCategoryBackupAndExportEncryption,
	"POST_QUANTUM_READINESS":          ListCryptoAssessmentFindingAnalyticsCategoryPostQuantumReadiness,
	"NOT_SUPPORTED":                   ListCryptoAssessmentFindingAnalyticsCategoryNotSupported,
}

var mappingListCryptoAssessmentFindingAnalyticsCategoryEnumLowerCase = map[string]ListCryptoAssessmentFindingAnalyticsCategoryEnum{
	"network_encryption":              ListCryptoAssessmentFindingAnalyticsCategoryNetworkEncryption,
	"data_encryption":                 ListCryptoAssessmentFindingAnalyticsCategoryDataEncryption,
	"certificates_and_key_management": ListCryptoAssessmentFindingAnalyticsCategoryCertificatesAndKeyManagement,
	"backup_and_export_encryption":    ListCryptoAssessmentFindingAnalyticsCategoryBackupAndExportEncryption,
	"post_quantum_readiness":          ListCryptoAssessmentFindingAnalyticsCategoryPostQuantumReadiness,
	"not_supported":                   ListCryptoAssessmentFindingAnalyticsCategoryNotSupported,
}

// GetListCryptoAssessmentFindingAnalyticsCategoryEnumValues Enumerates the set of values for ListCryptoAssessmentFindingAnalyticsCategoryEnum
func GetListCryptoAssessmentFindingAnalyticsCategoryEnumValues() []ListCryptoAssessmentFindingAnalyticsCategoryEnum {
	values := make([]ListCryptoAssessmentFindingAnalyticsCategoryEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingAnalyticsCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingAnalyticsCategoryEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingAnalyticsCategoryEnum
func GetListCryptoAssessmentFindingAnalyticsCategoryEnumStringValues() []string {
	return []string{
		"NETWORK_ENCRYPTION",
		"DATA_ENCRYPTION",
		"CERTIFICATES_AND_KEY_MANAGEMENT",
		"BACKUP_AND_EXPORT_ENCRYPTION",
		"POST_QUANTUM_READINESS",
		"NOT_SUPPORTED",
	}
}

// GetMappingListCryptoAssessmentFindingAnalyticsCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingAnalyticsCategoryEnum(val string) (ListCryptoAssessmentFindingAnalyticsCategoryEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingAnalyticsCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentFindingAnalyticsSortByEnum Enum with underlying type: string
type ListCryptoAssessmentFindingAnalyticsSortByEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingAnalyticsSortByEnum
const (
	ListCryptoAssessmentFindingAnalyticsSortByTargetcount ListCryptoAssessmentFindingAnalyticsSortByEnum = "targetCount"
	ListCryptoAssessmentFindingAnalyticsSortByPriority    ListCryptoAssessmentFindingAnalyticsSortByEnum = "priority"
)

var mappingListCryptoAssessmentFindingAnalyticsSortByEnum = map[string]ListCryptoAssessmentFindingAnalyticsSortByEnum{
	"targetCount": ListCryptoAssessmentFindingAnalyticsSortByTargetcount,
	"priority":    ListCryptoAssessmentFindingAnalyticsSortByPriority,
}

var mappingListCryptoAssessmentFindingAnalyticsSortByEnumLowerCase = map[string]ListCryptoAssessmentFindingAnalyticsSortByEnum{
	"targetcount": ListCryptoAssessmentFindingAnalyticsSortByTargetcount,
	"priority":    ListCryptoAssessmentFindingAnalyticsSortByPriority,
}

// GetListCryptoAssessmentFindingAnalyticsSortByEnumValues Enumerates the set of values for ListCryptoAssessmentFindingAnalyticsSortByEnum
func GetListCryptoAssessmentFindingAnalyticsSortByEnumValues() []ListCryptoAssessmentFindingAnalyticsSortByEnum {
	values := make([]ListCryptoAssessmentFindingAnalyticsSortByEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingAnalyticsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingAnalyticsSortByEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingAnalyticsSortByEnum
func GetListCryptoAssessmentFindingAnalyticsSortByEnumStringValues() []string {
	return []string{
		"targetCount",
		"priority",
	}
}

// GetMappingListCryptoAssessmentFindingAnalyticsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingAnalyticsSortByEnum(val string) (ListCryptoAssessmentFindingAnalyticsSortByEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingAnalyticsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentFindingAnalyticsSortOrderEnum Enum with underlying type: string
type ListCryptoAssessmentFindingAnalyticsSortOrderEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingAnalyticsSortOrderEnum
const (
	ListCryptoAssessmentFindingAnalyticsSortOrderAsc  ListCryptoAssessmentFindingAnalyticsSortOrderEnum = "ASC"
	ListCryptoAssessmentFindingAnalyticsSortOrderDesc ListCryptoAssessmentFindingAnalyticsSortOrderEnum = "DESC"
)

var mappingListCryptoAssessmentFindingAnalyticsSortOrderEnum = map[string]ListCryptoAssessmentFindingAnalyticsSortOrderEnum{
	"ASC":  ListCryptoAssessmentFindingAnalyticsSortOrderAsc,
	"DESC": ListCryptoAssessmentFindingAnalyticsSortOrderDesc,
}

var mappingListCryptoAssessmentFindingAnalyticsSortOrderEnumLowerCase = map[string]ListCryptoAssessmentFindingAnalyticsSortOrderEnum{
	"asc":  ListCryptoAssessmentFindingAnalyticsSortOrderAsc,
	"desc": ListCryptoAssessmentFindingAnalyticsSortOrderDesc,
}

// GetListCryptoAssessmentFindingAnalyticsSortOrderEnumValues Enumerates the set of values for ListCryptoAssessmentFindingAnalyticsSortOrderEnum
func GetListCryptoAssessmentFindingAnalyticsSortOrderEnumValues() []ListCryptoAssessmentFindingAnalyticsSortOrderEnum {
	values := make([]ListCryptoAssessmentFindingAnalyticsSortOrderEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingAnalyticsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingAnalyticsSortOrderEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingAnalyticsSortOrderEnum
func GetListCryptoAssessmentFindingAnalyticsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCryptoAssessmentFindingAnalyticsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingAnalyticsSortOrderEnum(val string) (ListCryptoAssessmentFindingAnalyticsSortOrderEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingAnalyticsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
