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

// ListCryptoAssessmentFindingTargetsRequest wrapper for the ListCryptoAssessmentFindingTargets operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListCryptoAssessmentFindingTargets.go.html to see an example of how to use ListCryptoAssessmentFindingTargetsRequest.
type ListCryptoAssessmentFindingTargetsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The finding keys for which target occurrences are listed.
	FindingKey []string `contributesTo:"query" name:"findingKey" collectionFormat:"multi"`

	// A filter to return targets from assessments of the specified type.
	AssessmentType CryptoAssessmentTypeEnum `mandatory:"false" contributesTo:"query" name:"assessmentType" omitEmpty:"true"`

	// Filters results to targets with an exact matching target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only resources associated with any of the specified target OCIDs.
	TargetIds []string `contributesTo:"query" name:"targetIds" collectionFormat:"multi"`

	// A filter to return only finding target rows with the specified status.
	Status ListCryptoAssessmentFindingTargetsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return only findings that are or are not part of quantum-readiness checks.
	IsQuantumReadinessCheck *bool `mandatory:"false" contributesTo:"query" name:"isQuantumReadinessCheck"`

	// The field used to sort finding target results.
	SortBy ListCryptoAssessmentFindingTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListCryptoAssessmentFindingTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListCryptoAssessmentFindingTargetsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

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

func (request ListCryptoAssessmentFindingTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCryptoAssessmentFindingTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCryptoAssessmentFindingTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCryptoAssessmentFindingTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCryptoAssessmentFindingTargetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(request.AssessmentType)); !ok && request.AssessmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessmentType: %s. Supported values are: %s.", request.AssessmentType, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentFindingTargetsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListCryptoAssessmentFindingTargetsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentFindingTargetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCryptoAssessmentFindingTargetsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentFindingTargetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCryptoAssessmentFindingTargetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCryptoAssessmentFindingTargetsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListCryptoAssessmentFindingTargetsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCryptoAssessmentFindingTargetsResponse wrapper for the ListCryptoAssessmentFindingTargets operation
type ListCryptoAssessmentFindingTargetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CryptoAssessmentFindingTargetCollection instances
	CryptoAssessmentFindingTargetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCryptoAssessmentFindingTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCryptoAssessmentFindingTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCryptoAssessmentFindingTargetsStatusEnum Enum with underlying type: string
type ListCryptoAssessmentFindingTargetsStatusEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingTargetsStatusEnum
const (
	ListCryptoAssessmentFindingTargetsStatusPass          ListCryptoAssessmentFindingTargetsStatusEnum = "PASS"
	ListCryptoAssessmentFindingTargetsStatusFail          ListCryptoAssessmentFindingTargetsStatusEnum = "FAIL"
	ListCryptoAssessmentFindingTargetsStatusError         ListCryptoAssessmentFindingTargetsStatusEnum = "ERROR"
	ListCryptoAssessmentFindingTargetsStatusEvaluate      ListCryptoAssessmentFindingTargetsStatusEnum = "EVALUATE"
	ListCryptoAssessmentFindingTargetsStatusNotAvailable  ListCryptoAssessmentFindingTargetsStatusEnum = "NOT_AVAILABLE"
	ListCryptoAssessmentFindingTargetsStatusNotApplicable ListCryptoAssessmentFindingTargetsStatusEnum = "NOT_APPLICABLE"
	ListCryptoAssessmentFindingTargetsStatusNotSupported  ListCryptoAssessmentFindingTargetsStatusEnum = "NOT_SUPPORTED"
)

var mappingListCryptoAssessmentFindingTargetsStatusEnum = map[string]ListCryptoAssessmentFindingTargetsStatusEnum{
	"PASS":           ListCryptoAssessmentFindingTargetsStatusPass,
	"FAIL":           ListCryptoAssessmentFindingTargetsStatusFail,
	"ERROR":          ListCryptoAssessmentFindingTargetsStatusError,
	"EVALUATE":       ListCryptoAssessmentFindingTargetsStatusEvaluate,
	"NOT_AVAILABLE":  ListCryptoAssessmentFindingTargetsStatusNotAvailable,
	"NOT_APPLICABLE": ListCryptoAssessmentFindingTargetsStatusNotApplicable,
	"NOT_SUPPORTED":  ListCryptoAssessmentFindingTargetsStatusNotSupported,
}

var mappingListCryptoAssessmentFindingTargetsStatusEnumLowerCase = map[string]ListCryptoAssessmentFindingTargetsStatusEnum{
	"pass":           ListCryptoAssessmentFindingTargetsStatusPass,
	"fail":           ListCryptoAssessmentFindingTargetsStatusFail,
	"error":          ListCryptoAssessmentFindingTargetsStatusError,
	"evaluate":       ListCryptoAssessmentFindingTargetsStatusEvaluate,
	"not_available":  ListCryptoAssessmentFindingTargetsStatusNotAvailable,
	"not_applicable": ListCryptoAssessmentFindingTargetsStatusNotApplicable,
	"not_supported":  ListCryptoAssessmentFindingTargetsStatusNotSupported,
}

// GetListCryptoAssessmentFindingTargetsStatusEnumValues Enumerates the set of values for ListCryptoAssessmentFindingTargetsStatusEnum
func GetListCryptoAssessmentFindingTargetsStatusEnumValues() []ListCryptoAssessmentFindingTargetsStatusEnum {
	values := make([]ListCryptoAssessmentFindingTargetsStatusEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingTargetsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingTargetsStatusEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingTargetsStatusEnum
func GetListCryptoAssessmentFindingTargetsStatusEnumStringValues() []string {
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

// GetMappingListCryptoAssessmentFindingTargetsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingTargetsStatusEnum(val string) (ListCryptoAssessmentFindingTargetsStatusEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingTargetsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentFindingTargetsSortByEnum Enum with underlying type: string
type ListCryptoAssessmentFindingTargetsSortByEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingTargetsSortByEnum
const (
	ListCryptoAssessmentFindingTargetsSortByTargetid        ListCryptoAssessmentFindingTargetsSortByEnum = "targetId"
	ListCryptoAssessmentFindingTargetsSortByDatabaseversion ListCryptoAssessmentFindingTargetsSortByEnum = "databaseVersion"
)

var mappingListCryptoAssessmentFindingTargetsSortByEnum = map[string]ListCryptoAssessmentFindingTargetsSortByEnum{
	"targetId":        ListCryptoAssessmentFindingTargetsSortByTargetid,
	"databaseVersion": ListCryptoAssessmentFindingTargetsSortByDatabaseversion,
}

var mappingListCryptoAssessmentFindingTargetsSortByEnumLowerCase = map[string]ListCryptoAssessmentFindingTargetsSortByEnum{
	"targetid":        ListCryptoAssessmentFindingTargetsSortByTargetid,
	"databaseversion": ListCryptoAssessmentFindingTargetsSortByDatabaseversion,
}

// GetListCryptoAssessmentFindingTargetsSortByEnumValues Enumerates the set of values for ListCryptoAssessmentFindingTargetsSortByEnum
func GetListCryptoAssessmentFindingTargetsSortByEnumValues() []ListCryptoAssessmentFindingTargetsSortByEnum {
	values := make([]ListCryptoAssessmentFindingTargetsSortByEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingTargetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingTargetsSortByEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingTargetsSortByEnum
func GetListCryptoAssessmentFindingTargetsSortByEnumStringValues() []string {
	return []string{
		"targetId",
		"databaseVersion",
	}
}

// GetMappingListCryptoAssessmentFindingTargetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingTargetsSortByEnum(val string) (ListCryptoAssessmentFindingTargetsSortByEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingTargetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentFindingTargetsSortOrderEnum Enum with underlying type: string
type ListCryptoAssessmentFindingTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingTargetsSortOrderEnum
const (
	ListCryptoAssessmentFindingTargetsSortOrderAsc  ListCryptoAssessmentFindingTargetsSortOrderEnum = "ASC"
	ListCryptoAssessmentFindingTargetsSortOrderDesc ListCryptoAssessmentFindingTargetsSortOrderEnum = "DESC"
)

var mappingListCryptoAssessmentFindingTargetsSortOrderEnum = map[string]ListCryptoAssessmentFindingTargetsSortOrderEnum{
	"ASC":  ListCryptoAssessmentFindingTargetsSortOrderAsc,
	"DESC": ListCryptoAssessmentFindingTargetsSortOrderDesc,
}

var mappingListCryptoAssessmentFindingTargetsSortOrderEnumLowerCase = map[string]ListCryptoAssessmentFindingTargetsSortOrderEnum{
	"asc":  ListCryptoAssessmentFindingTargetsSortOrderAsc,
	"desc": ListCryptoAssessmentFindingTargetsSortOrderDesc,
}

// GetListCryptoAssessmentFindingTargetsSortOrderEnumValues Enumerates the set of values for ListCryptoAssessmentFindingTargetsSortOrderEnum
func GetListCryptoAssessmentFindingTargetsSortOrderEnumValues() []ListCryptoAssessmentFindingTargetsSortOrderEnum {
	values := make([]ListCryptoAssessmentFindingTargetsSortOrderEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingTargetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingTargetsSortOrderEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingTargetsSortOrderEnum
func GetListCryptoAssessmentFindingTargetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCryptoAssessmentFindingTargetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingTargetsSortOrderEnum(val string) (ListCryptoAssessmentFindingTargetsSortOrderEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingTargetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCryptoAssessmentFindingTargetsAccessLevelEnum Enum with underlying type: string
type ListCryptoAssessmentFindingTargetsAccessLevelEnum string

// Set of constants representing the allowable values for ListCryptoAssessmentFindingTargetsAccessLevelEnum
const (
	ListCryptoAssessmentFindingTargetsAccessLevelRestricted ListCryptoAssessmentFindingTargetsAccessLevelEnum = "RESTRICTED"
	ListCryptoAssessmentFindingTargetsAccessLevelAccessible ListCryptoAssessmentFindingTargetsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListCryptoAssessmentFindingTargetsAccessLevelEnum = map[string]ListCryptoAssessmentFindingTargetsAccessLevelEnum{
	"RESTRICTED": ListCryptoAssessmentFindingTargetsAccessLevelRestricted,
	"ACCESSIBLE": ListCryptoAssessmentFindingTargetsAccessLevelAccessible,
}

var mappingListCryptoAssessmentFindingTargetsAccessLevelEnumLowerCase = map[string]ListCryptoAssessmentFindingTargetsAccessLevelEnum{
	"restricted": ListCryptoAssessmentFindingTargetsAccessLevelRestricted,
	"accessible": ListCryptoAssessmentFindingTargetsAccessLevelAccessible,
}

// GetListCryptoAssessmentFindingTargetsAccessLevelEnumValues Enumerates the set of values for ListCryptoAssessmentFindingTargetsAccessLevelEnum
func GetListCryptoAssessmentFindingTargetsAccessLevelEnumValues() []ListCryptoAssessmentFindingTargetsAccessLevelEnum {
	values := make([]ListCryptoAssessmentFindingTargetsAccessLevelEnum, 0)
	for _, v := range mappingListCryptoAssessmentFindingTargetsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListCryptoAssessmentFindingTargetsAccessLevelEnumStringValues Enumerates the set of values in String for ListCryptoAssessmentFindingTargetsAccessLevelEnum
func GetListCryptoAssessmentFindingTargetsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListCryptoAssessmentFindingTargetsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCryptoAssessmentFindingTargetsAccessLevelEnum(val string) (ListCryptoAssessmentFindingTargetsAccessLevelEnum, bool) {
	enum, ok := mappingListCryptoAssessmentFindingTargetsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
