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

// ListCryptoAssessmentBackupSetsRequest wrapper for the ListCryptoAssessmentBackupSets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCryptoAssessmentBackupSets.go.html to see an example of how to use ListCryptoAssessmentBackupSetsRequest.
type ListCryptoAssessmentBackupSetsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListCryptoAssessmentBackupSetsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only resources associated with the specified crypto assessment OCID.
	AssessmentId *string `mandatory:"false" contributesTo:"query" name:"assessmentId"`

	// A filter to return targets from assessments of the specified type.
	AssessmentType CryptoAssessmentTypeEnum `mandatory:"false" contributesTo:"query" name:"assessmentType" omitEmpty:"true"`

	// A filter to return only inventory rows associated with the specified target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only resources associated with any of the specified target OCIDs.
	TargetIds []string `contributesTo:"query" name:"targetIds" collectionFormat:"multi"`

	// Filters backup set summary rows to an exact matching backupSetKey.
	BackupSetKey *string `mandatory:"false" contributesTo:"query" name:"backupSetKey"`

	// Filters backup set summary rows by whether the backup set is encrypted.
	IsEncrypted *bool `mandatory:"false" contributesTo:"query" name:"isEncrypted"`

	// The field used to sort backup set results.
	SortBy ListCryptoAssessmentBackupSetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListCryptoAssessmentBackupSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListCryptoAssessmentBackupSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAssessmentBackupSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAssessmentBackupSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAssessmentBackupSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAssessmentBackupSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCryptoAssessmentBackupSetsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCryptoAssessmentBackupSetsAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(request.AssessmentType)); !ok && request.AssessmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessmentType: %s. Supported values are: %s.", request.AssessmentType, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentBackupSetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCryptoAssessmentBackupSetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentBackupSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCryptoAssessmentBackupSetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAssessmentBackupSetsResponse wrapper for the ListCryptoAssessmentBackupSets operation
type ListCryptoAssessmentBackupSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentBackupSetCollection instances
	CryptoAssessmentBackupSetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCryptoAssessmentBackupSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAssessmentBackupSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAssessmentBackupSetsAccessLevelEnum Enum with underlying type: string
type ListCryptoAssessmentBackupSetsAccessLevelEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentBackupSetsAccessLevelEnum
const (
	ListCryptoAssessmentBackupSetsAccessLevelRestricted ListCryptoAssessmentBackupSetsAccessLevelEnum = "RESTRICTED"
	ListCryptoAssessmentBackupSetsAccessLevelAccessible ListCryptoAssessmentBackupSetsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCryptoAssessmentBackupSetsAccessLevelEnum = map[string]ListCryptoAssessmentBackupSetsAccessLevelEnum{
	"RESTRICTED": ListCryptoAssessmentBackupSetsAccessLevelRestricted,
	"ACCESSIBLE": ListCryptoAssessmentBackupSetsAccessLevelAccessible,
}

var mappingListCryptoAssessmentBackupSetsAccessLevelEnumLowerCase = map[string]ListCryptoAssessmentBackupSetsAccessLevelEnum{
	"restricted": ListCryptoAssessmentBackupSetsAccessLevelRestricted,
	"accessible": ListCryptoAssessmentBackupSetsAccessLevelAccessible,
}

// GetListCryptoAssessmentBackupSetsAccessLevelEnumValues Enumerates the set of values for ListCryptoAssessmentBackupSetsAccessLevelEnum
func GetListCryptoAssessmentBackupSetsAccessLevelEnumValues() []ListCryptoAssessmentBackupSetsAccessLevelEnum {
	values := make([]ListCryptoAssessmentBackupSetsAccessLevelEnum, 0)
	for _, v := range mappingListCryptoAssessmentBackupSetsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentBackupSetsAccessLevelEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentBackupSetsAccessLevelEnum
func GetListCryptoAssessmentBackupSetsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCryptoAssessmentBackupSetsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentBackupSetsAccessLevelEnum(val string) (ListCryptoAssessmentBackupSetsAccessLevelEnum, bool) {
	enum, ok := mappingListCryptoAssessmentBackupSetsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentBackupSetsSortByEnum Enum with underlying type: string
type ListCryptoAssessmentBackupSetsSortByEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentBackupSetsSortByEnum
const (
	ListCryptoAssessmentBackupSetsSortBySetstamp     ListCryptoAssessmentBackupSetsSortByEnum = "setStamp"
	ListCryptoAssessmentBackupSetsSortByBackupsetkey ListCryptoAssessmentBackupSetsSortByEnum = "backupSetKey"
	ListCryptoAssessmentBackupSetsSortBySizeingbs    ListCryptoAssessmentBackupSetsSortByEnum = "sizeInGBs"
	ListCryptoAssessmentBackupSetsSortByTimecreated  ListCryptoAssessmentBackupSetsSortByEnum = "timeCreated"
)

var mappingListCryptoAssessmentBackupSetsSortByEnum = map[string]ListCryptoAssessmentBackupSetsSortByEnum{
	"setStamp":     ListCryptoAssessmentBackupSetsSortBySetstamp,
	"backupSetKey": ListCryptoAssessmentBackupSetsSortByBackupsetkey,
	"sizeInGBs":    ListCryptoAssessmentBackupSetsSortBySizeingbs,
	"timeCreated":  ListCryptoAssessmentBackupSetsSortByTimecreated,
}

var mappingListCryptoAssessmentBackupSetsSortByEnumLowerCase = map[string]ListCryptoAssessmentBackupSetsSortByEnum{
	"setstamp":     ListCryptoAssessmentBackupSetsSortBySetstamp,
	"backupsetkey": ListCryptoAssessmentBackupSetsSortByBackupsetkey,
	"sizeingbs":    ListCryptoAssessmentBackupSetsSortBySizeingbs,
	"timecreated":  ListCryptoAssessmentBackupSetsSortByTimecreated,
}

// GetListCryptoAssessmentBackupSetsSortByEnumValues Enumerates the set of values for ListCryptoAssessmentBackupSetsSortByEnum
func GetListCryptoAssessmentBackupSetsSortByEnumValues() []ListCryptoAssessmentBackupSetsSortByEnum {
	values := make([]ListCryptoAssessmentBackupSetsSortByEnum, 0)
	for _, v := range mappingListCryptoAssessmentBackupSetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentBackupSetsSortByEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentBackupSetsSortByEnum
func GetListCryptoAssessmentBackupSetsSortByEnumStringValues() []string {
	return []string{
		"setStamp",
		"backupSetKey",
		"sizeInGBs",
		"timeCreated",
	}
}

// GetMappingListCryptoAssessmentBackupSetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentBackupSetsSortByEnum(val string) (ListCryptoAssessmentBackupSetsSortByEnum, bool) {
	enum, ok := mappingListCryptoAssessmentBackupSetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentBackupSetsSortOrderEnum Enum with underlying type: string
type ListCryptoAssessmentBackupSetsSortOrderEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentBackupSetsSortOrderEnum
const (
	ListCryptoAssessmentBackupSetsSortOrderAsc  ListCryptoAssessmentBackupSetsSortOrderEnum = "ASC"
	ListCryptoAssessmentBackupSetsSortOrderDesc ListCryptoAssessmentBackupSetsSortOrderEnum = "DESC"
)

var mappingListCryptoAssessmentBackupSetsSortOrderEnum = map[string]ListCryptoAssessmentBackupSetsSortOrderEnum{
	"ASC":  ListCryptoAssessmentBackupSetsSortOrderAsc,
	"DESC": ListCryptoAssessmentBackupSetsSortOrderDesc,
}

var mappingListCryptoAssessmentBackupSetsSortOrderEnumLowerCase = map[string]ListCryptoAssessmentBackupSetsSortOrderEnum{
	"asc":  ListCryptoAssessmentBackupSetsSortOrderAsc,
	"desc": ListCryptoAssessmentBackupSetsSortOrderDesc,
}

// GetListCryptoAssessmentBackupSetsSortOrderEnumValues Enumerates the set of values for ListCryptoAssessmentBackupSetsSortOrderEnum
func GetListCryptoAssessmentBackupSetsSortOrderEnumValues() []ListCryptoAssessmentBackupSetsSortOrderEnum {
	values := make([]ListCryptoAssessmentBackupSetsSortOrderEnum, 0)
	for _, v := range mappingListCryptoAssessmentBackupSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentBackupSetsSortOrderEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentBackupSetsSortOrderEnum
func GetListCryptoAssessmentBackupSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCryptoAssessmentBackupSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentBackupSetsSortOrderEnum(val string) (ListCryptoAssessmentBackupSetsSortOrderEnum, bool) {
	enum, ok := mappingListCryptoAssessmentBackupSetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
