// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// SummarizeExadataInsightResourceUsageRequest wrapper for the SummarizeExadataInsightResourceUsage operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceUsage.go.html to see an example of how to use SummarizeExadataInsightResourceUsageRequest.
type SummarizeExadataInsightResourceUsageRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Filter by resource.
	// Supported values are HOST , STORAGE_SERVER and DATABASE
	ResourceType *string `mandatory:"true" contributesTo:"query" name:"resourceType"`

	// Filter by resource metric.
	// Supported values are CPU , STORAGE, MEMORY, IO, IOPS, THROUGHPUT
	ResourceMetric *string `mandatory:"true" contributesTo:"query" name:"resourceMetric"`

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

	// Optional list of exadata insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	ExadataInsightId []string `contributesTo:"query" name:"exadataInsightId" collectionFormat:"multi"`

	// Filter by one or more Exadata types.
	// Possible value are DBMACHINE, EXACS, and EXACC.
	ExadataType []string `contributesTo:"query" name:"exadataType" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeExadataInsightResourceUsageSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The order in which resource usage summary records are listed
	SortBy SummarizeExadataInsightResourceUsageSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Percentile values of daily usage to be used for computing the aggregate resource usage.
	Percentile *int `mandatory:"false" contributesTo:"query" name:"percentile"`

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

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeExadataInsightResourceUsageRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeExadataInsightResourceUsageRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeExadataInsightResourceUsageRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeExadataInsightResourceUsageRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeExadataInsightResourceUsageRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceUsageSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeExadataInsightResourceUsageSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceUsageSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeExadataInsightResourceUsageSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceUsageResponse wrapper for the SummarizeExadataInsightResourceUsage operation
type SummarizeExadataInsightResourceUsageResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeExadataInsightResourceUsageCollection instances
	SummarizeExadataInsightResourceUsageCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeExadataInsightResourceUsageResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeExadataInsightResourceUsageResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeExadataInsightResourceUsageSortOrderEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageSortOrderEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageSortOrderEnum
const (
	SummarizeExadataInsightResourceUsageSortOrderAsc  SummarizeExadataInsightResourceUsageSortOrderEnum = "ASC"
	SummarizeExadataInsightResourceUsageSortOrderDesc SummarizeExadataInsightResourceUsageSortOrderEnum = "DESC"
)

var mappingSummarizeExadataInsightResourceUsageSortOrderEnum = map[string]SummarizeExadataInsightResourceUsageSortOrderEnum{
	"ASC":  SummarizeExadataInsightResourceUsageSortOrderAsc,
	"DESC": SummarizeExadataInsightResourceUsageSortOrderDesc,
}

// GetSummarizeExadataInsightResourceUsageSortOrderEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageSortOrderEnum
func GetSummarizeExadataInsightResourceUsageSortOrderEnumValues() []SummarizeExadataInsightResourceUsageSortOrderEnum {
	values := make([]SummarizeExadataInsightResourceUsageSortOrderEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUsageSortOrderEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUsageSortOrderEnum
func GetSummarizeExadataInsightResourceUsageSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeExadataInsightResourceUsageSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUsageSortOrderEnum(val string) (SummarizeExadataInsightResourceUsageSortOrderEnum, bool) {
	mappingSummarizeExadataInsightResourceUsageSortOrderEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceUsageSortOrderEnum)
	for k, v := range mappingSummarizeExadataInsightResourceUsageSortOrderEnum {
		mappingSummarizeExadataInsightResourceUsageSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceUsageSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceUsageSortByEnum Enum with underlying type: string
type SummarizeExadataInsightResourceUsageSortByEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceUsageSortByEnum
const (
	SummarizeExadataInsightResourceUsageSortByUtilizationpercent SummarizeExadataInsightResourceUsageSortByEnum = "utilizationPercent"
	SummarizeExadataInsightResourceUsageSortByUsage              SummarizeExadataInsightResourceUsageSortByEnum = "usage"
	SummarizeExadataInsightResourceUsageSortByCapacity           SummarizeExadataInsightResourceUsageSortByEnum = "capacity"
	SummarizeExadataInsightResourceUsageSortByUsagechangepercent SummarizeExadataInsightResourceUsageSortByEnum = "usageChangePercent"
)

var mappingSummarizeExadataInsightResourceUsageSortByEnum = map[string]SummarizeExadataInsightResourceUsageSortByEnum{
	"utilizationPercent": SummarizeExadataInsightResourceUsageSortByUtilizationpercent,
	"usage":              SummarizeExadataInsightResourceUsageSortByUsage,
	"capacity":           SummarizeExadataInsightResourceUsageSortByCapacity,
	"usageChangePercent": SummarizeExadataInsightResourceUsageSortByUsagechangepercent,
}

// GetSummarizeExadataInsightResourceUsageSortByEnumValues Enumerates the set of values for SummarizeExadataInsightResourceUsageSortByEnum
func GetSummarizeExadataInsightResourceUsageSortByEnumValues() []SummarizeExadataInsightResourceUsageSortByEnum {
	values := make([]SummarizeExadataInsightResourceUsageSortByEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceUsageSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceUsageSortByEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceUsageSortByEnum
func GetSummarizeExadataInsightResourceUsageSortByEnumStringValues() []string {
	return []string{
		"utilizationPercent",
		"usage",
		"capacity",
		"usageChangePercent",
	}
}

// GetMappingSummarizeExadataInsightResourceUsageSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceUsageSortByEnum(val string) (SummarizeExadataInsightResourceUsageSortByEnum, bool) {
	mappingSummarizeExadataInsightResourceUsageSortByEnumIgnoreCase := make(map[string]SummarizeExadataInsightResourceUsageSortByEnum)
	for k, v := range mappingSummarizeExadataInsightResourceUsageSortByEnum {
		mappingSummarizeExadataInsightResourceUsageSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSummarizeExadataInsightResourceUsageSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
