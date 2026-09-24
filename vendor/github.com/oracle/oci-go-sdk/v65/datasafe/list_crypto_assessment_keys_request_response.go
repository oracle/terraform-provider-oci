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

// ListCryptoAssessmentKeysRequest wrapper for the ListCryptoAssessmentKeys operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCryptoAssessmentKeys.go.html to see an example of how to use ListCryptoAssessmentKeysRequest.
type ListCryptoAssessmentKeysRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListCryptoAssessmentKeysAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources associated with the specified crypto assessment OCID.
	AssessmentId *string `mandatory:"false" contributesTo:"query" name:"assessmentId"`

	// A filter to return targets from assessments of the specified type.
	AssessmentType CryptoAssessmentTypeEnum `mandatory:"false" contributesTo:"query" name:"assessmentType" omitEmpty:"true"`

	// A filter to return only inventory rows associated with the specified target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only resources associated with any of the specified target OCIDs.
	TargetIds []string `contributesTo:"query" name:"targetIds" collectionFormat:"multi"`

	// A filter to return only records for the specified feature.
	Feature ListCryptoAssessmentKeysFeatureEnum `mandatory:"false" contributesTo:"query" name:"feature" omitEmpty:"true"`

	// Filters key results to rows with an exact matching keyId.
	KeyId *string `mandatory:"false" contributesTo:"query" name:"keyId"`

	// Filters key results to rows with the specified key type.
	KeyType CryptoAssessmentKeySummaryKeyTypeEnum `mandatory:"false" contributesTo:"query" name:"keyType" omitEmpty:"true"`

	// Filters key results to rows whose primary or secondary keystore type matches any of the specified key manager types.
	KeyManagerType []CryptoKeystoreTypeEnum `contributesTo:"query" name:"keyManagerType" omitEmpty:"true" collectionFormat:"multi"`

	// The field used to sort key results.
	SortBy ListCryptoAssessmentKeysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListCryptoAssessmentKeysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListCryptoAssessmentKeysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAssessmentKeysRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAssessmentKeysRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAssessmentKeysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAssessmentKeysRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCryptoAssessmentKeysAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCryptoAssessmentKeysAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(request.AssessmentType)); !ok && request.AssessmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessmentType: %s. Supported values are: %s.", request.AssessmentType, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentKeysFeatureEnum(string(request.Feature)); !ok && request.Feature != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Feature: %s. Supported values are: %s.", request.Feature, strings.Join(GetListCryptoAssessmentKeysFeatureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentKeySummaryKeyTypeEnum(string(request.KeyType)); !ok && request.KeyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyType: %s. Supported values are: %s.", request.KeyType, strings.Join(GetCryptoAssessmentKeySummaryKeyTypeEnumStringValues(), ",")))
	}
	for _, val := range request.KeyManagerType {
		if _, ok := GetMappingCryptoKeystoreTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyManagerType: %s. Supported values are: %s.", val, strings.Join(GetCryptoKeystoreTypeEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListCryptoAssessmentKeysSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCryptoAssessmentKeysSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentKeysSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCryptoAssessmentKeysSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAssessmentKeysResponse wrapper for the ListCryptoAssessmentKeys operation
type ListCryptoAssessmentKeysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentKeyCollection instances
	CryptoAssessmentKeyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCryptoAssessmentKeysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAssessmentKeysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAssessmentKeysAccessLevelEnum Enum with underlying type: string
type ListCryptoAssessmentKeysAccessLevelEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentKeysAccessLevelEnum
const (
	ListCryptoAssessmentKeysAccessLevelRestricted ListCryptoAssessmentKeysAccessLevelEnum = "RESTRICTED"
	ListCryptoAssessmentKeysAccessLevelAccessible ListCryptoAssessmentKeysAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCryptoAssessmentKeysAccessLevelEnum = map[string]ListCryptoAssessmentKeysAccessLevelEnum{
	"RESTRICTED": ListCryptoAssessmentKeysAccessLevelRestricted,
	"ACCESSIBLE": ListCryptoAssessmentKeysAccessLevelAccessible,
}

var mappingListCryptoAssessmentKeysAccessLevelEnumLowerCase = map[string]ListCryptoAssessmentKeysAccessLevelEnum{
	"restricted": ListCryptoAssessmentKeysAccessLevelRestricted,
	"accessible": ListCryptoAssessmentKeysAccessLevelAccessible,
}

// GetListCryptoAssessmentKeysAccessLevelEnumValues Enumerates the set of values for ListCryptoAssessmentKeysAccessLevelEnum
func GetListCryptoAssessmentKeysAccessLevelEnumValues() []ListCryptoAssessmentKeysAccessLevelEnum {
	values := make([]ListCryptoAssessmentKeysAccessLevelEnum, 0)
	for _, v := range mappingListCryptoAssessmentKeysAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentKeysAccessLevelEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentKeysAccessLevelEnum
func GetListCryptoAssessmentKeysAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCryptoAssessmentKeysAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentKeysAccessLevelEnum(val string) (ListCryptoAssessmentKeysAccessLevelEnum, bool) {
	enum, ok := mappingListCryptoAssessmentKeysAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentKeysFeatureEnum Enum with underlying type: string
type ListCryptoAssessmentKeysFeatureEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentKeysFeatureEnum
const (
	ListCryptoAssessmentKeysFeatureTde ListCryptoAssessmentKeysFeatureEnum = "TDE"
	ListCryptoAssessmentKeysFeatureTls ListCryptoAssessmentKeysFeatureEnum = "TLS"
	ListCryptoAssessmentKeysFeatureNne ListCryptoAssessmentKeysFeatureEnum = "NNE"
)

var mappingListCryptoAssessmentKeysFeatureEnum = map[string]ListCryptoAssessmentKeysFeatureEnum{
	"TDE": ListCryptoAssessmentKeysFeatureTde,
	"TLS": ListCryptoAssessmentKeysFeatureTls,
	"NNE": ListCryptoAssessmentKeysFeatureNne,
}

var mappingListCryptoAssessmentKeysFeatureEnumLowerCase = map[string]ListCryptoAssessmentKeysFeatureEnum{
	"tde": ListCryptoAssessmentKeysFeatureTde,
	"tls": ListCryptoAssessmentKeysFeatureTls,
	"nne": ListCryptoAssessmentKeysFeatureNne,
}

// GetListCryptoAssessmentKeysFeatureEnumValues Enumerates the set of values for ListCryptoAssessmentKeysFeatureEnum
func GetListCryptoAssessmentKeysFeatureEnumValues() []ListCryptoAssessmentKeysFeatureEnum {
	values := make([]ListCryptoAssessmentKeysFeatureEnum, 0)
	for _, v := range mappingListCryptoAssessmentKeysFeatureEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentKeysFeatureEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentKeysFeatureEnum
func GetListCryptoAssessmentKeysFeatureEnumStringValues() []string {
	return []string{
		"TDE",
		"TLS",
		"NNE",
	}
}

// GetMappingListCryptoAssessmentKeysFeatureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentKeysFeatureEnum(val string) (ListCryptoAssessmentKeysFeatureEnum, bool) {
	enum, ok := mappingListCryptoAssessmentKeysFeatureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentKeysSortByEnum Enum with underlying type: string
type ListCryptoAssessmentKeysSortByEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentKeysSortByEnum
const (
	ListCryptoAssessmentKeysSortByTimecreated ListCryptoAssessmentKeysSortByEnum = "timeCreated"
)

var mappingListCryptoAssessmentKeysSortByEnum = map[string]ListCryptoAssessmentKeysSortByEnum{
	"timeCreated": ListCryptoAssessmentKeysSortByTimecreated,
}

var mappingListCryptoAssessmentKeysSortByEnumLowerCase = map[string]ListCryptoAssessmentKeysSortByEnum{
	"timecreated": ListCryptoAssessmentKeysSortByTimecreated,
}

// GetListCryptoAssessmentKeysSortByEnumValues Enumerates the set of values for ListCryptoAssessmentKeysSortByEnum
func GetListCryptoAssessmentKeysSortByEnumValues() []ListCryptoAssessmentKeysSortByEnum {
	values := make([]ListCryptoAssessmentKeysSortByEnum, 0)
	for _, v := range mappingListCryptoAssessmentKeysSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentKeysSortByEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentKeysSortByEnum
func GetListCryptoAssessmentKeysSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListCryptoAssessmentKeysSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentKeysSortByEnum(val string) (ListCryptoAssessmentKeysSortByEnum, bool) {
	enum, ok := mappingListCryptoAssessmentKeysSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentKeysSortOrderEnum Enum with underlying type: string
type ListCryptoAssessmentKeysSortOrderEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentKeysSortOrderEnum
const (
	ListCryptoAssessmentKeysSortOrderAsc  ListCryptoAssessmentKeysSortOrderEnum = "ASC"
	ListCryptoAssessmentKeysSortOrderDesc ListCryptoAssessmentKeysSortOrderEnum = "DESC"
)

var mappingListCryptoAssessmentKeysSortOrderEnum = map[string]ListCryptoAssessmentKeysSortOrderEnum{
	"ASC":  ListCryptoAssessmentKeysSortOrderAsc,
	"DESC": ListCryptoAssessmentKeysSortOrderDesc,
}

var mappingListCryptoAssessmentKeysSortOrderEnumLowerCase = map[string]ListCryptoAssessmentKeysSortOrderEnum{
	"asc":  ListCryptoAssessmentKeysSortOrderAsc,
	"desc": ListCryptoAssessmentKeysSortOrderDesc,
}

// GetListCryptoAssessmentKeysSortOrderEnumValues Enumerates the set of values for ListCryptoAssessmentKeysSortOrderEnum
func GetListCryptoAssessmentKeysSortOrderEnumValues() []ListCryptoAssessmentKeysSortOrderEnum {
	values := make([]ListCryptoAssessmentKeysSortOrderEnum, 0)
	for _, v := range mappingListCryptoAssessmentKeysSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentKeysSortOrderEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentKeysSortOrderEnum
func GetListCryptoAssessmentKeysSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCryptoAssessmentKeysSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentKeysSortOrderEnum(val string) (ListCryptoAssessmentKeysSortOrderEnum, bool) {
	enum, ok := mappingListCryptoAssessmentKeysSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
