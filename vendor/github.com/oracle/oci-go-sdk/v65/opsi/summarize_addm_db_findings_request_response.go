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

// SummarizeAddmDbFindingsRequest wrapper for the SummarizeAddmDbFindings operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeAddmDbFindings.go.html to see an example of how to use SummarizeAddmDbFindingsRequest.
type SummarizeAddmDbFindingsRequest struct {

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

	// Optional value filter to match the finding category exactly.
	CategoryName *string `mandatory:"false" contributesTo:"query" name:"categoryName"`

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
	SortOrder SummarizeAddmDbFindingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Field name for sorting the ADDM finding summary data
	SortBy SummarizeAddmDbFindingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

func (request SummarizeAddmDbFindingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeAddmDbFindingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeAddmDbFindingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeAddmDbFindingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeAddmDbFindingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeAddmDbFindingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeAddmDbFindingsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeAddmDbFindingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeAddmDbFindingsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeAddmDbFindingsResponse wrapper for the SummarizeAddmDbFindings operation
type SummarizeAddmDbFindingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AddmDbFindingAggregationCollection instances
	AddmDbFindingAggregationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeAddmDbFindingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeAddmDbFindingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeAddmDbFindingsSortOrderEnum Enum with underlying type: string
type SummarizeAddmDbFindingsSortOrderEnum string

// Set of constants representing the allowable values for SummarizeAddmDbFindingsSortOrderEnum
const (
	SummarizeAddmDbFindingsSortOrderAsc  SummarizeAddmDbFindingsSortOrderEnum = "ASC"
	SummarizeAddmDbFindingsSortOrderDesc SummarizeAddmDbFindingsSortOrderEnum = "DESC"
)

var mappingSummarizeAddmDbFindingsSortOrderEnum = map[string]SummarizeAddmDbFindingsSortOrderEnum{
	"ASC":  SummarizeAddmDbFindingsSortOrderAsc,
	"DESC": SummarizeAddmDbFindingsSortOrderDesc,
}

var mappingSummarizeAddmDbFindingsSortOrderEnumLowerCase = map[string]SummarizeAddmDbFindingsSortOrderEnum{
	"asc":  SummarizeAddmDbFindingsSortOrderAsc,
	"desc": SummarizeAddmDbFindingsSortOrderDesc,
}

// GetSummarizeAddmDbFindingsSortOrderEnumValues Enumerates the set of values for SummarizeAddmDbFindingsSortOrderEnum
func GetSummarizeAddmDbFindingsSortOrderEnumValues() []SummarizeAddmDbFindingsSortOrderEnum {
	values := make([]SummarizeAddmDbFindingsSortOrderEnum, 0)
	for _, v := range mappingSummarizeAddmDbFindingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAddmDbFindingsSortOrderEnumStringValues Enumerates the set of values in String for SummarizeAddmDbFindingsSortOrderEnum
func GetSummarizeAddmDbFindingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeAddmDbFindingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAddmDbFindingsSortOrderEnum(val string) (SummarizeAddmDbFindingsSortOrderEnum, bool) {
	enum, ok := mappingSummarizeAddmDbFindingsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeAddmDbFindingsSortByEnum Enum with underlying type: string
type SummarizeAddmDbFindingsSortByEnum string

// Set of constants representing the allowable values for SummarizeAddmDbFindingsSortByEnum
const (
	SummarizeAddmDbFindingsSortByImpactoverallpercent    SummarizeAddmDbFindingsSortByEnum = "impactOverallPercent"
	SummarizeAddmDbFindingsSortByImpactmaxpercent        SummarizeAddmDbFindingsSortByEnum = "impactMaxPercent"
	SummarizeAddmDbFindingsSortByImpactavgactivesessions SummarizeAddmDbFindingsSortByEnum = "impactAvgActiveSessions"
	SummarizeAddmDbFindingsSortByFrequencycount          SummarizeAddmDbFindingsSortByEnum = "frequencyCount"
)

var mappingSummarizeAddmDbFindingsSortByEnum = map[string]SummarizeAddmDbFindingsSortByEnum{
	"impactOverallPercent":    SummarizeAddmDbFindingsSortByImpactoverallpercent,
	"impactMaxPercent":        SummarizeAddmDbFindingsSortByImpactmaxpercent,
	"impactAvgActiveSessions": SummarizeAddmDbFindingsSortByImpactavgactivesessions,
	"frequencyCount":          SummarizeAddmDbFindingsSortByFrequencycount,
}

var mappingSummarizeAddmDbFindingsSortByEnumLowerCase = map[string]SummarizeAddmDbFindingsSortByEnum{
	"impactoverallpercent":    SummarizeAddmDbFindingsSortByImpactoverallpercent,
	"impactmaxpercent":        SummarizeAddmDbFindingsSortByImpactmaxpercent,
	"impactavgactivesessions": SummarizeAddmDbFindingsSortByImpactavgactivesessions,
	"frequencycount":          SummarizeAddmDbFindingsSortByFrequencycount,
}

// GetSummarizeAddmDbFindingsSortByEnumValues Enumerates the set of values for SummarizeAddmDbFindingsSortByEnum
func GetSummarizeAddmDbFindingsSortByEnumValues() []SummarizeAddmDbFindingsSortByEnum {
	values := make([]SummarizeAddmDbFindingsSortByEnum, 0)
	for _, v := range mappingSummarizeAddmDbFindingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeAddmDbFindingsSortByEnumStringValues Enumerates the set of values in String for SummarizeAddmDbFindingsSortByEnum
func GetSummarizeAddmDbFindingsSortByEnumStringValues() []string {
	return []string{
		"impactOverallPercent",
		"impactMaxPercent",
		"impactAvgActiveSessions",
		"frequencyCount",
	}
}

// GetMappingSummarizeAddmDbFindingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeAddmDbFindingsSortByEnum(val string) (SummarizeAddmDbFindingsSortByEnum, bool) {
	enum, ok := mappingSummarizeAddmDbFindingsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
