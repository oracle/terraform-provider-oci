// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// SummarizeExadataInsightResourceCapacityTrendAggregatedRequest wrapper for the SummarizeExadataInsightResourceCapacityTrendAggregated operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceCapacityTrendAggregated.go.html to see an example of how to use SummarizeExadataInsightResourceCapacityTrendAggregatedRequest.
type SummarizeExadataInsightResourceCapacityTrendAggregatedRequest struct {

	// Filter by resource.
	// Supported values are HOST , STORAGE_SERVER and DATABASE
	ResourceType *string `mandatory:"true" contributesTo:"query" name:"resourceType"`

	// Filter by resource metric.
	// Supported values are CPU , STORAGE, MEMORY, IO, IOPS, THROUGHPUT
	ResourceMetric *string `mandatory:"true" contributesTo:"query" name:"resourceMetric"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Specify time period in ISO 8601 format with respect to current time.
	// Default is last 30 days represented by P30D.
	// If timeInterval is specified, then timeIntervalStart and timeIntervalEnd will be ignored.
	// Examples  P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months), . Maximum value allowed is 25 months prior to current time (P25M).
	AnalysisTimeInterval *string `mandatory:"false" contributesTo:"query" name:"analysisTimeInterval"`

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

	// Optional list of exadata insight resource OCIDs (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more Exadata types.
	// Possible value are DBMACHINE, EXACS, and EXACC.
	ExadataType []string `contributesTo:"query" name:"exadataType" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sorts using end timestamp or capacity.
	SortBy SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeExadataInsightResourceCapacityTrendAggregatedRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeExadataInsightResourceCapacityTrendAggregatedRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeExadataInsightResourceCapacityTrendAggregatedRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeExadataInsightResourceCapacityTrendAggregatedRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeExadataInsightResourceCapacityTrendAggregatedRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceCapacityTrendAggregatedResponse wrapper for the SummarizeExadataInsightResourceCapacityTrendAggregated operation
type SummarizeExadataInsightResourceCapacityTrendAggregatedResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeExadataInsightResourceCapacityTrendAggregation instances
	SummarizeExadataInsightResourceCapacityTrendAggregation `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeExadataInsightResourceCapacityTrendAggregatedResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeExadataInsightResourceCapacityTrendAggregatedResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum Enum with underlying type: string
type SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum
const (
	SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderAsc  SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum = "ASC"
	SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderDesc SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum = "DESC"
)

var mappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum = map[string]SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum{
	"ASC":  SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderAsc,
	"DESC": SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderDesc,
}

var mappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnumLowerCase = map[string]SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum{
	"asc":  SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderAsc,
	"desc": SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderDesc,
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnumValues() []SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum(val string) (SummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum Enum with underlying type: string
type SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum
const (
	SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEndtimestamp SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum = "endTimestamp"
	SummarizeExadataInsightResourceCapacityTrendAggregatedSortByCapacity     SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum = "capacity"
)

var mappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum = map[string]SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum{
	"endTimestamp": SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEndtimestamp,
	"capacity":     SummarizeExadataInsightResourceCapacityTrendAggregatedSortByCapacity,
}

var mappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnumLowerCase = map[string]SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum{
	"endtimestamp": SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEndtimestamp,
	"capacity":     SummarizeExadataInsightResourceCapacityTrendAggregatedSortByCapacity,
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnumValues() []SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum
func GetSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnumStringValues() []string {
	return []string{
		"endTimestamp",
		"capacity",
	}
}

// GetMappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum(val string) (SummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceCapacityTrendAggregatedSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
