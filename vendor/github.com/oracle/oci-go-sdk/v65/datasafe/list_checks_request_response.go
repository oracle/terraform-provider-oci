// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListChecksRequest wrapper for the ListChecks operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListChecks.go.html to see an example of how to use ListChecksRequest.
type ListChecksRequest struct {

	// The OCID of the security assessment.
	SecurityAssessmentId *string `mandatory:"true" contributesTo:"path" name:"securityAssessmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListChecksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order(sortOrder). The default order for title is ascending.
	SortBy ListChecksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only checks of a particular risk level.
	SuggestedSeverity ListChecksSuggestedSeverityEnum `mandatory:"false" contributesTo:"query" name:"suggestedSeverity" omitEmpty:"true"`

	// A filter to return only findings that match the specified risk level(s). Use containsSeverity parameter if need to filter by multiple risk levels.
	ContainsSeverity []ListChecksContainsSeverityEnum `contributesTo:"query" name:"containsSeverity" omitEmpty:"true" collectionFormat:"multi"`

	// An optional filter to return only findings that match the specified references. Use containsReferences param if need to filter by multiple references.
	ContainsReferences []SecurityAssessmentReferencesEnum `contributesTo:"query" name:"containsReferences" omitEmpty:"true" collectionFormat:"multi"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListChecksAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Each check in security assessment has an associated key (think of key as a check's name).
	// For a given check, the key will be the same across targets. The user can use these keys to filter the checks.
	Key *string `mandatory:"false" contributesTo:"query" name:"key"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListChecksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListChecksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListChecksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListChecksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListChecksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListChecksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListChecksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChecksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListChecksSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListChecksSuggestedSeverityEnum(string(request.SuggestedSeverity)); !ok && request.SuggestedSeverity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SuggestedSeverity: %s. Supported values are: %s.", request.SuggestedSeverity, strings.Join(GetListChecksSuggestedSeverityEnumStringValues(), ",")))
	}
	for _, val := range request.ContainsSeverity {
		if _, ok := GetMappingListChecksContainsSeverityEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContainsSeverity: %s. Supported values are: %s.", val, strings.Join(GetListChecksContainsSeverityEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ContainsReferences {
		if _, ok := GetMappingSecurityAssessmentReferencesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ContainsReferences: %s. Supported values are: %s.", val, strings.Join(GetSecurityAssessmentReferencesEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListChecksAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListChecksAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListChecksResponse wrapper for the ListChecks operation
type ListChecksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CheckSummary instances
	Items []CheckSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListChecksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListChecksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListChecksSortOrderEnum Enum with underlying type: string
type ListChecksSortOrderEnum string

// Set of constants representing the allowable values for ListChecksSortOrderEnum
const (
	ListChecksSortOrderAsc  ListChecksSortOrderEnum = "ASC"
	ListChecksSortOrderDesc ListChecksSortOrderEnum = "DESC"
)

var mappingListChecksSortOrderEnum = map[string]ListChecksSortOrderEnum{
	"ASC":  ListChecksSortOrderAsc,
	"DESC": ListChecksSortOrderDesc,
}

var mappingListChecksSortOrderEnumLowerCase = map[string]ListChecksSortOrderEnum{
	"asc":  ListChecksSortOrderAsc,
	"desc": ListChecksSortOrderDesc,
}

// GetListChecksSortOrderEnumValues Enumerates the set of values for ListChecksSortOrderEnum
func GetListChecksSortOrderEnumValues() []ListChecksSortOrderEnum {
	values := make([]ListChecksSortOrderEnum, 0)
	for _, v := range mappingListChecksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListChecksSortOrderEnumStringValues Enumerates the set of values in String for ListChecksSortOrderEnum
func GetListChecksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListChecksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChecksSortOrderEnum(val string) (ListChecksSortOrderEnum, bool) {
	enum, ok := mappingListChecksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChecksSortByEnum Enum with underlying type: string
type ListChecksSortByEnum string

// Set of constants representing the allowable values for ListChecksSortByEnum
const (
	ListChecksSortByTitle    ListChecksSortByEnum = "title"
	ListChecksSortByCategory ListChecksSortByEnum = "category"
)

var mappingListChecksSortByEnum = map[string]ListChecksSortByEnum{
	"title":    ListChecksSortByTitle,
	"category": ListChecksSortByCategory,
}

var mappingListChecksSortByEnumLowerCase = map[string]ListChecksSortByEnum{
	"title":    ListChecksSortByTitle,
	"category": ListChecksSortByCategory,
}

// GetListChecksSortByEnumValues Enumerates the set of values for ListChecksSortByEnum
func GetListChecksSortByEnumValues() []ListChecksSortByEnum {
	values := make([]ListChecksSortByEnum, 0)
	for _, v := range mappingListChecksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListChecksSortByEnumStringValues Enumerates the set of values in String for ListChecksSortByEnum
func GetListChecksSortByEnumStringValues() []string {
	return []string{
		"title",
		"category",
	}
}

// GetMappingListChecksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChecksSortByEnum(val string) (ListChecksSortByEnum, bool) {
	enum, ok := mappingListChecksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChecksSuggestedSeverityEnum Enum with underlying type: string
type ListChecksSuggestedSeverityEnum string

// Set of constants representing the allowable values for ListChecksSuggestedSeverityEnum
const (
	ListChecksSuggestedSeverityHigh     ListChecksSuggestedSeverityEnum = "HIGH"
	ListChecksSuggestedSeverityMedium   ListChecksSuggestedSeverityEnum = "MEDIUM"
	ListChecksSuggestedSeverityLow      ListChecksSuggestedSeverityEnum = "LOW"
	ListChecksSuggestedSeverityEvaluate ListChecksSuggestedSeverityEnum = "EVALUATE"
	ListChecksSuggestedSeverityAdvisory ListChecksSuggestedSeverityEnum = "ADVISORY"
	ListChecksSuggestedSeverityPass     ListChecksSuggestedSeverityEnum = "PASS"
	ListChecksSuggestedSeverityDeferred ListChecksSuggestedSeverityEnum = "DEFERRED"
)

var mappingListChecksSuggestedSeverityEnum = map[string]ListChecksSuggestedSeverityEnum{
	"HIGH":     ListChecksSuggestedSeverityHigh,
	"MEDIUM":   ListChecksSuggestedSeverityMedium,
	"LOW":      ListChecksSuggestedSeverityLow,
	"EVALUATE": ListChecksSuggestedSeverityEvaluate,
	"ADVISORY": ListChecksSuggestedSeverityAdvisory,
	"PASS":     ListChecksSuggestedSeverityPass,
	"DEFERRED": ListChecksSuggestedSeverityDeferred,
}

var mappingListChecksSuggestedSeverityEnumLowerCase = map[string]ListChecksSuggestedSeverityEnum{
	"high":     ListChecksSuggestedSeverityHigh,
	"medium":   ListChecksSuggestedSeverityMedium,
	"low":      ListChecksSuggestedSeverityLow,
	"evaluate": ListChecksSuggestedSeverityEvaluate,
	"advisory": ListChecksSuggestedSeverityAdvisory,
	"pass":     ListChecksSuggestedSeverityPass,
	"deferred": ListChecksSuggestedSeverityDeferred,
}

// GetListChecksSuggestedSeverityEnumValues Enumerates the set of values for ListChecksSuggestedSeverityEnum
func GetListChecksSuggestedSeverityEnumValues() []ListChecksSuggestedSeverityEnum {
	values := make([]ListChecksSuggestedSeverityEnum, 0)
	for _, v := range mappingListChecksSuggestedSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetListChecksSuggestedSeverityEnumStringValues Enumerates the set of values in String for ListChecksSuggestedSeverityEnum
func GetListChecksSuggestedSeverityEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
		"EVALUATE",
		"ADVISORY",
		"PASS",
		"DEFERRED",
	}
}

// GetMappingListChecksSuggestedSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChecksSuggestedSeverityEnum(val string) (ListChecksSuggestedSeverityEnum, bool) {
	enum, ok := mappingListChecksSuggestedSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChecksContainsSeverityEnum Enum with underlying type: string
type ListChecksContainsSeverityEnum string

// Set of constants representing the allowable values for ListChecksContainsSeverityEnum
const (
	ListChecksContainsSeverityHigh     ListChecksContainsSeverityEnum = "HIGH"
	ListChecksContainsSeverityMedium   ListChecksContainsSeverityEnum = "MEDIUM"
	ListChecksContainsSeverityLow      ListChecksContainsSeverityEnum = "LOW"
	ListChecksContainsSeverityEvaluate ListChecksContainsSeverityEnum = "EVALUATE"
	ListChecksContainsSeverityAdvisory ListChecksContainsSeverityEnum = "ADVISORY"
	ListChecksContainsSeverityPass     ListChecksContainsSeverityEnum = "PASS"
	ListChecksContainsSeverityDeferred ListChecksContainsSeverityEnum = "DEFERRED"
)

var mappingListChecksContainsSeverityEnum = map[string]ListChecksContainsSeverityEnum{
	"HIGH":     ListChecksContainsSeverityHigh,
	"MEDIUM":   ListChecksContainsSeverityMedium,
	"LOW":      ListChecksContainsSeverityLow,
	"EVALUATE": ListChecksContainsSeverityEvaluate,
	"ADVISORY": ListChecksContainsSeverityAdvisory,
	"PASS":     ListChecksContainsSeverityPass,
	"DEFERRED": ListChecksContainsSeverityDeferred,
}

var mappingListChecksContainsSeverityEnumLowerCase = map[string]ListChecksContainsSeverityEnum{
	"high":     ListChecksContainsSeverityHigh,
	"medium":   ListChecksContainsSeverityMedium,
	"low":      ListChecksContainsSeverityLow,
	"evaluate": ListChecksContainsSeverityEvaluate,
	"advisory": ListChecksContainsSeverityAdvisory,
	"pass":     ListChecksContainsSeverityPass,
	"deferred": ListChecksContainsSeverityDeferred,
}

// GetListChecksContainsSeverityEnumValues Enumerates the set of values for ListChecksContainsSeverityEnum
func GetListChecksContainsSeverityEnumValues() []ListChecksContainsSeverityEnum {
	values := make([]ListChecksContainsSeverityEnum, 0)
	for _, v := range mappingListChecksContainsSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetListChecksContainsSeverityEnumStringValues Enumerates the set of values in String for ListChecksContainsSeverityEnum
func GetListChecksContainsSeverityEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
		"EVALUATE",
		"ADVISORY",
		"PASS",
		"DEFERRED",
	}
}

// GetMappingListChecksContainsSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChecksContainsSeverityEnum(val string) (ListChecksContainsSeverityEnum, bool) {
	enum, ok := mappingListChecksContainsSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListChecksAccessLevelEnum Enum with underlying type: string
type ListChecksAccessLevelEnum string

// Set of constants representing the allowable values for ListChecksAccessLevelEnum
const (
	ListChecksAccessLevelRestricted ListChecksAccessLevelEnum = "RESTRICTED"
	ListChecksAccessLevelAccessible ListChecksAccessLevelEnum = "ACCESSIBLE"
)

var mappingListChecksAccessLevelEnum = map[string]ListChecksAccessLevelEnum{
	"RESTRICTED": ListChecksAccessLevelRestricted,
	"ACCESSIBLE": ListChecksAccessLevelAccessible,
}

var mappingListChecksAccessLevelEnumLowerCase = map[string]ListChecksAccessLevelEnum{
	"restricted": ListChecksAccessLevelRestricted,
	"accessible": ListChecksAccessLevelAccessible,
}

// GetListChecksAccessLevelEnumValues Enumerates the set of values for ListChecksAccessLevelEnum
func GetListChecksAccessLevelEnumValues() []ListChecksAccessLevelEnum {
	values := make([]ListChecksAccessLevelEnum, 0)
	for _, v := range mappingListChecksAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListChecksAccessLevelEnumStringValues Enumerates the set of values in String for ListChecksAccessLevelEnum
func GetListChecksAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListChecksAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListChecksAccessLevelEnum(val string) (ListChecksAccessLevelEnum, bool) {
	enum, ok := mappingListChecksAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
