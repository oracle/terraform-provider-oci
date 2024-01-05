// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeAddmDbParametersRequest wrapper for the SummarizeAddmDbParameters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAddmDbParameters.go.html to see an example of how to use SummarizeAddmDbParametersRequest.
type SummarizeAddmDbParametersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Optional list of database OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
	DatabaseId []string `contributesTo:"query" name:"databaseId" collectionFormat:"multi"`

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	Id []string `contributesTo:"query" name:"id" collectionFormat:"multi"`

	// The optional single value query parameter to filter by database instance number.
	InstanceNumber *string `mandatory:"false" contributesTo:"query" name:"instanceNumber"`

	// Analysis start time in UTC in ISO 8601 format(inclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// The minimum allowed value is 2 years prior to the current day.
	// timeIntervalStart and timeIntervalEnd parameters are used together.
	// If analysisTimeInterval is specified, this parameter is ignored.
	TimeIntervalStart *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalStart"`

	// Analysis end time in UTC in ISO 8601 format(exclusive).
	// Example 2019-10-30T00:00:00Z (yyyy-MM-ddThh:mm:ssZ).
	// timeIntervalStart and timeIntervalEnd are used together.
	// If timeIntervalEnd is not specified, current time is used as timeIntervalEnd.
	TimeIntervalEnd *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeIntervalEnd"`

	// Optional value filter to match the parameter category exactly. Note the list of possible
	// category names can be retrieved from the following endpoint:
	// /databases/{databaseId}/addmDbParameterCategories.
	CategoryName *string `mandatory:"false" contributesTo:"query" name:"categoryName"`

	// Optional filter to return only resources whose name or value contains the substring given. The
	// match is not case sensitive.
	NameOrValueContains *string `mandatory:"false" contributesTo:"query" name:"nameOrValueContains"`

	// Optional filter to return only parameters whose value changed in the specified time period.
	// Valid values include: TRUE, FALSE
	IsChanged SummarizeAddmDbParametersIsChangedEnum `mandatory:"false" contributesTo:"query" name:"isChanged" omitEmpty:"true"`

	// Optional filter to return only parameters whose end value was set to the default value (TRUE)
	// or was specified in the parameter file (FALSE). Valid values include: TRUE, FALSE
	IsDefault SummarizeAddmDbParametersIsDefaultEnum `mandatory:"false" contributesTo:"query" name:"isDefault" omitEmpty:"true"`

	// Optional filter to return only parameters which have recommendations in the specified time period.
	// Valid values include: TRUE, FALSE
	HasRecommendations SummarizeAddmDbParametersHasRecommendationsEnum `mandatory:"false" contributesTo:"query" name:"hasRecommendations" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeAddmDbParametersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Field name for sorting the database parameter data
	SortBy SummarizeAddmDbParametersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A list of tag filters to apply.  Only resources with a defined tag matching the value will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagEquals []string `contributesTo:"query" name:"definedTagEquals" collectionFormat:"multi"`

	// A list of tag filters to apply.  Only resources with a freeform tag matching the value will be returned.
	// The key for each tag is "{tagName}.{value}".  All inputs are case-insensitive.
	// Multiple values for the same tag name are interpreted as "OR".  Values for different tag names are interpreted as "AND".
	FreeformTagEquals []string `contributesTo:"query" name:"freeformTagEquals" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified defined tags exist will be returned.
	// Each item in the list has the format "{namespace}.{tagName}.true" (for checking existence of a defined tag)
	// or "{namespace}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for the same key (i.e. same namespace and tag name) are interpreted as "OR".
	// Values for different keys (i.e. different namespaces, different tag names, or both) are interpreted as "AND".
	DefinedTagExists []string `contributesTo:"query" name:"definedTagExists" collectionFormat:"multi"`

	// A list of tag existence filters to apply.  Only resources for which the specified freeform tags exist the value will be returned.
	// The key for each tag is "{tagName}.true".  All inputs are case-insensitive.
	// Currently, only existence ("true" at the end) is supported. Absence ("false" at the end) is not supported.
	// Multiple values for different tag names are interpreted as "AND".
	FreeformTagExists []string `contributesTo:"query" name:"freeformTagExists" collectionFormat:"multi"`

	// A flag to search all resources within a given compartment and all sub-compartments.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeAddmDbParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAddmDbParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAddmDbParametersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAddmDbParametersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAddmDbParametersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAddmDbParametersIsChangedEnum(string(request.IsChanged)); !ok && request.IsChanged != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsChanged: %s. Supported values are: %s.", request.IsChanged, strings.Join(GetSummarizeAddmDbParametersIsChangedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAddmDbParametersIsDefaultEnum(string(request.IsDefault)); !ok && request.IsDefault != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IsDefault: %s. Supported values are: %s.", request.IsDefault, strings.Join(GetSummarizeAddmDbParametersIsDefaultEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAddmDbParametersHasRecommendationsEnum(string(request.HasRecommendations)); !ok && request.HasRecommendations != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HasRecommendations: %s. Supported values are: %s.", request.HasRecommendations, strings.Join(GetSummarizeAddmDbParametersHasRecommendationsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAddmDbParametersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAddmDbParametersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAddmDbParametersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAddmDbParametersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAddmDbParametersResponse wrapper for the SummarizeAddmDbParameters operation
type SummarizeAddmDbParametersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AddmDbParameterAggregationCollection instances
	AddmDbParameterAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAddmDbParametersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAddmDbParametersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAddmDbParametersIsChangedEnum Enum with underlying type: string
type SummarizeAddmDbParametersIsChangedEnum string

// Set of constants representing the allowable values for SummarizeAddmDbParametersIsChangedEnum
const (
	SummarizeAddmDbParametersIsChangedTrue  SummarizeAddmDbParametersIsChangedEnum = "true"
	SummarizeAddmDbParametersIsChangedFalse SummarizeAddmDbParametersIsChangedEnum = "false"
)

var mappingSummarizeAddmDbParametersIsChangedEnum = map[string]SummarizeAddmDbParametersIsChangedEnum{
	"true":  SummarizeAddmDbParametersIsChangedTrue,
	"false": SummarizeAddmDbParametersIsChangedFalse,
}

var mappingSummarizeAddmDbParametersIsChangedEnumLowerCase = map[string]SummarizeAddmDbParametersIsChangedEnum{
	"true":  SummarizeAddmDbParametersIsChangedTrue,
	"false": SummarizeAddmDbParametersIsChangedFalse,
}

// GetSummarizeAddmDbParametersIsChangedEnumValues Enumerates the set of values for SummarizeAddmDbParametersIsChangedEnum
func GetSummarizeAddmDbParametersIsChangedEnumValues() []SummarizeAddmDbParametersIsChangedEnum {
	values := make([]SummarizeAddmDbParametersIsChangedEnum, 0)
	for _, v := range mappingSummarizeAddmDbParametersIsChangedEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAddmDbParametersIsChangedEnumStringValues Enumerates the set of values in String for SummarizeAddmDbParametersIsChangedEnum
func GetSummarizeAddmDbParametersIsChangedEnumStringValues() []string {
	return []string{
		"true",
		"false",
	}
}

// GetMappingSummarizeAddmDbParametersIsChangedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAddmDbParametersIsChangedEnum(val string) (SummarizeAddmDbParametersIsChangedEnum, bool) {
	enum, ok := mappingSummarizeAddmDbParametersIsChangedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAddmDbParametersIsDefaultEnum Enum with underlying type: string
type SummarizeAddmDbParametersIsDefaultEnum string

// Set of constants representing the allowable values for SummarizeAddmDbParametersIsDefaultEnum
const (
	SummarizeAddmDbParametersIsDefaultTrue  SummarizeAddmDbParametersIsDefaultEnum = "true"
	SummarizeAddmDbParametersIsDefaultFalse SummarizeAddmDbParametersIsDefaultEnum = "false"
)

var mappingSummarizeAddmDbParametersIsDefaultEnum = map[string]SummarizeAddmDbParametersIsDefaultEnum{
	"true":  SummarizeAddmDbParametersIsDefaultTrue,
	"false": SummarizeAddmDbParametersIsDefaultFalse,
}

var mappingSummarizeAddmDbParametersIsDefaultEnumLowerCase = map[string]SummarizeAddmDbParametersIsDefaultEnum{
	"true":  SummarizeAddmDbParametersIsDefaultTrue,
	"false": SummarizeAddmDbParametersIsDefaultFalse,
}

// GetSummarizeAddmDbParametersIsDefaultEnumValues Enumerates the set of values for SummarizeAddmDbParametersIsDefaultEnum
func GetSummarizeAddmDbParametersIsDefaultEnumValues() []SummarizeAddmDbParametersIsDefaultEnum {
	values := make([]SummarizeAddmDbParametersIsDefaultEnum, 0)
	for _, v := range mappingSummarizeAddmDbParametersIsDefaultEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAddmDbParametersIsDefaultEnumStringValues Enumerates the set of values in String for SummarizeAddmDbParametersIsDefaultEnum
func GetSummarizeAddmDbParametersIsDefaultEnumStringValues() []string {
	return []string{
		"true",
		"false",
	}
}

// GetMappingSummarizeAddmDbParametersIsDefaultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAddmDbParametersIsDefaultEnum(val string) (SummarizeAddmDbParametersIsDefaultEnum, bool) {
	enum, ok := mappingSummarizeAddmDbParametersIsDefaultEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAddmDbParametersHasRecommendationsEnum Enum with underlying type: string
type SummarizeAddmDbParametersHasRecommendationsEnum string

// Set of constants representing the allowable values for SummarizeAddmDbParametersHasRecommendationsEnum
const (
	SummarizeAddmDbParametersHasRecommendationsTrue  SummarizeAddmDbParametersHasRecommendationsEnum = "true"
	SummarizeAddmDbParametersHasRecommendationsFalse SummarizeAddmDbParametersHasRecommendationsEnum = "false"
)

var mappingSummarizeAddmDbParametersHasRecommendationsEnum = map[string]SummarizeAddmDbParametersHasRecommendationsEnum{
	"true":  SummarizeAddmDbParametersHasRecommendationsTrue,
	"false": SummarizeAddmDbParametersHasRecommendationsFalse,
}

var mappingSummarizeAddmDbParametersHasRecommendationsEnumLowerCase = map[string]SummarizeAddmDbParametersHasRecommendationsEnum{
	"true":  SummarizeAddmDbParametersHasRecommendationsTrue,
	"false": SummarizeAddmDbParametersHasRecommendationsFalse,
}

// GetSummarizeAddmDbParametersHasRecommendationsEnumValues Enumerates the set of values for SummarizeAddmDbParametersHasRecommendationsEnum
func GetSummarizeAddmDbParametersHasRecommendationsEnumValues() []SummarizeAddmDbParametersHasRecommendationsEnum {
	values := make([]SummarizeAddmDbParametersHasRecommendationsEnum, 0)
	for _, v := range mappingSummarizeAddmDbParametersHasRecommendationsEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAddmDbParametersHasRecommendationsEnumStringValues Enumerates the set of values in String for SummarizeAddmDbParametersHasRecommendationsEnum
func GetSummarizeAddmDbParametersHasRecommendationsEnumStringValues() []string {
	return []string{
		"true",
		"false",
	}
}

// GetMappingSummarizeAddmDbParametersHasRecommendationsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAddmDbParametersHasRecommendationsEnum(val string) (SummarizeAddmDbParametersHasRecommendationsEnum, bool) {
	enum, ok := mappingSummarizeAddmDbParametersHasRecommendationsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAddmDbParametersSortOrderEnum Enum with underlying type: string
type SummarizeAddmDbParametersSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAddmDbParametersSortOrderEnum
const (
	SummarizeAddmDbParametersSortOrderAsc  SummarizeAddmDbParametersSortOrderEnum = "ASC"
	SummarizeAddmDbParametersSortOrderDesc SummarizeAddmDbParametersSortOrderEnum = "DESC"
)

var mappingSummarizeAddmDbParametersSortOrderEnum = map[string]SummarizeAddmDbParametersSortOrderEnum{
	"ASC":  SummarizeAddmDbParametersSortOrderAsc,
	"DESC": SummarizeAddmDbParametersSortOrderDesc,
}

var mappingSummarizeAddmDbParametersSortOrderEnumLowerCase = map[string]SummarizeAddmDbParametersSortOrderEnum{
	"asc":  SummarizeAddmDbParametersSortOrderAsc,
	"desc": SummarizeAddmDbParametersSortOrderDesc,
}

// GetSummarizeAddmDbParametersSortOrderEnumValues Enumerates the set of values for SummarizeAddmDbParametersSortOrderEnum
func GetSummarizeAddmDbParametersSortOrderEnumValues() []SummarizeAddmDbParametersSortOrderEnum {
	values := make([]SummarizeAddmDbParametersSortOrderEnum, 0)
	for _, v := range mappingSummarizeAddmDbParametersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAddmDbParametersSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAddmDbParametersSortOrderEnum
func GetSummarizeAddmDbParametersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAddmDbParametersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAddmDbParametersSortOrderEnum(val string) (SummarizeAddmDbParametersSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAddmDbParametersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAddmDbParametersSortByEnum Enum with underlying type: string
type SummarizeAddmDbParametersSortByEnum string

// Set of constants representing the allowable values for SummarizeAddmDbParametersSortByEnum
const (
	SummarizeAddmDbParametersSortByIschanged SummarizeAddmDbParametersSortByEnum = "isChanged"
	SummarizeAddmDbParametersSortByName      SummarizeAddmDbParametersSortByEnum = "name"
)

var mappingSummarizeAddmDbParametersSortByEnum = map[string]SummarizeAddmDbParametersSortByEnum{
	"isChanged": SummarizeAddmDbParametersSortByIschanged,
	"name":      SummarizeAddmDbParametersSortByName,
}

var mappingSummarizeAddmDbParametersSortByEnumLowerCase = map[string]SummarizeAddmDbParametersSortByEnum{
	"ischanged": SummarizeAddmDbParametersSortByIschanged,
	"name":      SummarizeAddmDbParametersSortByName,
}

// GetSummarizeAddmDbParametersSortByEnumValues Enumerates the set of values for SummarizeAddmDbParametersSortByEnum
func GetSummarizeAddmDbParametersSortByEnumValues() []SummarizeAddmDbParametersSortByEnum {
	values := make([]SummarizeAddmDbParametersSortByEnum, 0)
	for _, v := range mappingSummarizeAddmDbParametersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAddmDbParametersSortByEnumStringValues Enumerates the set of values in String for SummarizeAddmDbParametersSortByEnum
func GetSummarizeAddmDbParametersSortByEnumStringValues() []string {
	return []string{
		"isChanged",
		"name",
	}
}

// GetMappingSummarizeAddmDbParametersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAddmDbParametersSortByEnum(val string) (SummarizeAddmDbParametersSortByEnum, bool) {
	enum, ok := mappingSummarizeAddmDbParametersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
