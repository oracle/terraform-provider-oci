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

// ListCryptoAssessmentWalletsRequest wrapper for the ListCryptoAssessmentWallets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCryptoAssessmentWallets.go.html to see an example of how to use ListCryptoAssessmentWalletsRequest.
type ListCryptoAssessmentWalletsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListCryptoAssessmentWalletsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources associated with the specified crypto assessment OCID.
	AssessmentId *string `mandatory:"false" contributesTo:"query" name:"assessmentId"`

	// A filter to return targets from assessments of the specified type.
	AssessmentType CryptoAssessmentTypeEnum `mandatory:"false" contributesTo:"query" name:"assessmentType" omitEmpty:"true"`

	// A filter to return only inventory rows associated with the specified target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only resources associated with any of the specified target OCIDs.
	TargetIds []string `contributesTo:"query" name:"targetIds" collectionFormat:"multi"`

	// A filter to return only wallets for the specified feature.
	Feature CryptoAssessmentWalletSummaryFeatureEnum `mandatory:"false" contributesTo:"query" name:"feature" omitEmpty:"true"`

	// A filter to return only wallets whose encryption algorithm exactly matches any of the specified values, case-insensitively.
	WalletEncryptionAlgorithm []string `contributesTo:"query" name:"walletEncryptionAlgorithm" collectionFormat:"multi"`

	// The field used to sort wallet results.
	SortBy ListCryptoAssessmentWalletsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListCryptoAssessmentWalletsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListCryptoAssessmentWalletsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAssessmentWalletsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAssessmentWalletsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAssessmentWalletsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAssessmentWalletsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCryptoAssessmentWalletsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCryptoAssessmentWalletsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(request.AssessmentType)); !ok && request.AssessmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessmentType: %s. Supported values are: %s.", request.AssessmentType, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentWalletSummaryFeatureEnum(string(request.Feature)); !ok && request.Feature != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Feature: %s. Supported values are: %s.", request.Feature, strings.Join(GetCryptoAssessmentWalletSummaryFeatureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentWalletsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCryptoAssessmentWalletsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentWalletsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCryptoAssessmentWalletsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAssessmentWalletsResponse wrapper for the ListCryptoAssessmentWallets operation
type ListCryptoAssessmentWalletsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentWalletCollection instances
	CryptoAssessmentWalletCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCryptoAssessmentWalletsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAssessmentWalletsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAssessmentWalletsAccessLevelEnum Enum with underlying type: string
type ListCryptoAssessmentWalletsAccessLevelEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentWalletsAccessLevelEnum
const (
	ListCryptoAssessmentWalletsAccessLevelRestricted ListCryptoAssessmentWalletsAccessLevelEnum = "RESTRICTED"
	ListCryptoAssessmentWalletsAccessLevelAccessible ListCryptoAssessmentWalletsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCryptoAssessmentWalletsAccessLevelEnum = map[string]ListCryptoAssessmentWalletsAccessLevelEnum{
	"RESTRICTED": ListCryptoAssessmentWalletsAccessLevelRestricted,
	"ACCESSIBLE": ListCryptoAssessmentWalletsAccessLevelAccessible,
}

var mappingListCryptoAssessmentWalletsAccessLevelEnumLowerCase = map[string]ListCryptoAssessmentWalletsAccessLevelEnum{
	"restricted": ListCryptoAssessmentWalletsAccessLevelRestricted,
	"accessible": ListCryptoAssessmentWalletsAccessLevelAccessible,
}

// GetListCryptoAssessmentWalletsAccessLevelEnumValues Enumerates the set of values for ListCryptoAssessmentWalletsAccessLevelEnum
func GetListCryptoAssessmentWalletsAccessLevelEnumValues() []ListCryptoAssessmentWalletsAccessLevelEnum {
	values := make([]ListCryptoAssessmentWalletsAccessLevelEnum, 0)
	for _, v := range mappingListCryptoAssessmentWalletsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentWalletsAccessLevelEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentWalletsAccessLevelEnum
func GetListCryptoAssessmentWalletsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCryptoAssessmentWalletsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentWalletsAccessLevelEnum(val string) (ListCryptoAssessmentWalletsAccessLevelEnum, bool) {
	enum, ok := mappingListCryptoAssessmentWalletsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentWalletsSortByEnum Enum with underlying type: string
type ListCryptoAssessmentWalletsSortByEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentWalletsSortByEnum
const (
	ListCryptoAssessmentWalletsSortByFeature     ListCryptoAssessmentWalletsSortByEnum = "feature"
	ListCryptoAssessmentWalletsSortByTimecreated ListCryptoAssessmentWalletsSortByEnum = "timeCreated"
)

var mappingListCryptoAssessmentWalletsSortByEnum = map[string]ListCryptoAssessmentWalletsSortByEnum{
	"feature":     ListCryptoAssessmentWalletsSortByFeature,
	"timeCreated": ListCryptoAssessmentWalletsSortByTimecreated,
}

var mappingListCryptoAssessmentWalletsSortByEnumLowerCase = map[string]ListCryptoAssessmentWalletsSortByEnum{
	"feature":     ListCryptoAssessmentWalletsSortByFeature,
	"timecreated": ListCryptoAssessmentWalletsSortByTimecreated,
}

// GetListCryptoAssessmentWalletsSortByEnumValues Enumerates the set of values for ListCryptoAssessmentWalletsSortByEnum
func GetListCryptoAssessmentWalletsSortByEnumValues() []ListCryptoAssessmentWalletsSortByEnum {
	values := make([]ListCryptoAssessmentWalletsSortByEnum, 0)
	for _, v := range mappingListCryptoAssessmentWalletsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentWalletsSortByEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentWalletsSortByEnum
func GetListCryptoAssessmentWalletsSortByEnumStringValues() []string {
	return []string{
		"feature",
		"timeCreated",
	}
}

// GetMappingListCryptoAssessmentWalletsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentWalletsSortByEnum(val string) (ListCryptoAssessmentWalletsSortByEnum, bool) {
	enum, ok := mappingListCryptoAssessmentWalletsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentWalletsSortOrderEnum Enum with underlying type: string
type ListCryptoAssessmentWalletsSortOrderEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentWalletsSortOrderEnum
const (
	ListCryptoAssessmentWalletsSortOrderAsc  ListCryptoAssessmentWalletsSortOrderEnum = "ASC"
	ListCryptoAssessmentWalletsSortOrderDesc ListCryptoAssessmentWalletsSortOrderEnum = "DESC"
)

var mappingListCryptoAssessmentWalletsSortOrderEnum = map[string]ListCryptoAssessmentWalletsSortOrderEnum{
	"ASC":  ListCryptoAssessmentWalletsSortOrderAsc,
	"DESC": ListCryptoAssessmentWalletsSortOrderDesc,
}

var mappingListCryptoAssessmentWalletsSortOrderEnumLowerCase = map[string]ListCryptoAssessmentWalletsSortOrderEnum{
	"asc":  ListCryptoAssessmentWalletsSortOrderAsc,
	"desc": ListCryptoAssessmentWalletsSortOrderDesc,
}

// GetListCryptoAssessmentWalletsSortOrderEnumValues Enumerates the set of values for ListCryptoAssessmentWalletsSortOrderEnum
func GetListCryptoAssessmentWalletsSortOrderEnumValues() []ListCryptoAssessmentWalletsSortOrderEnum {
	values := make([]ListCryptoAssessmentWalletsSortOrderEnum, 0)
	for _, v := range mappingListCryptoAssessmentWalletsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentWalletsSortOrderEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentWalletsSortOrderEnum
func GetListCryptoAssessmentWalletsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCryptoAssessmentWalletsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentWalletsSortOrderEnum(val string) (ListCryptoAssessmentWalletsSortOrderEnum, bool) {
	enum, ok := mappingListCryptoAssessmentWalletsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
