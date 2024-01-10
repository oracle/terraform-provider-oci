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

// SummarizeExadataInsightResourceCapacityTrendRequest wrapper for the SummarizeExadataInsightResourceCapacityTrend operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/SummarizeExadataInsightResourceCapacityTrend.go.html to see an example of how to use SummarizeExadataInsightResourceCapacityTrendRequest.
type SummarizeExadataInsightResourceCapacityTrendRequest struct {

	// Filter by resource.
	// Supported values are HOST , STORAGE_SERVER and DATABASE
	ResourceType *string `mandatory:"true" contributesTo:"query" name:"resourceType"`

	// Filter by resource metric.
	// Supported values are CPU , STORAGE, MEMORY, IO, IOPS, THROUGHPUT
	ResourceMetric *string `mandatory:"true" contributesTo:"query" name:"resourceMetric"`

	// OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource.
	ExadataInsightId *string `mandatory:"true" contributesTo:"query" name:"exadataInsightId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
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

	// Optional list of database insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	DatabaseInsightId []string `contributesTo:"query" name:"databaseInsightId" collectionFormat:"multi"`

	// Optional list of host insight resource OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	HostInsightId []string `contributesTo:"query" name:"hostInsightId" collectionFormat:"multi"`

	// Optional storage server name on an exadata system.
	StorageServerName []string `contributesTo:"query" name:"storageServerName" collectionFormat:"multi"`

	// Filter by one or more Exadata types.
	// Possible value are DBMACHINE, EXACS, and EXACC.
	ExadataType []string `contributesTo:"query" name:"exadataType" collectionFormat:"multi"`

	// Filter by one or more cdb name.
	CdbName []string `contributesTo:"query" name:"cdbName" collectionFormat:"multi"`

	// Filter by hostname.
	HostName []string `contributesTo:"query" name:"hostName" collectionFormat:"multi"`

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

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder SummarizeExadataInsightResourceCapacityTrendSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The order in which resource capacity trend records are listed
	SortBy SummarizeExadataInsightResourceCapacityTrendSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SummarizeExadataInsightResourceCapacityTrendRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SummarizeExadataInsightResourceCapacityTrendRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request SummarizeExadataInsightResourceCapacityTrendRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SummarizeExadataInsightResourceCapacityTrendRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SummarizeExadataInsightResourceCapacityTrendRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSummarizeExadataInsightResourceCapacityTrendSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetSummarizeExadataInsightResourceCapacityTrendSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSummarizeExadataInsightResourceCapacityTrendSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetSummarizeExadataInsightResourceCapacityTrendSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SummarizeExadataInsightResourceCapacityTrendResponse wrapper for the SummarizeExadataInsightResourceCapacityTrend operation
type SummarizeExadataInsightResourceCapacityTrendResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SummarizeExadataInsightResourceCapacityTrendCollection instances
	SummarizeExadataInsightResourceCapacityTrendCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response SummarizeExadataInsightResourceCapacityTrendResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SummarizeExadataInsightResourceCapacityTrendResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// SummarizeExadataInsightResourceCapacityTrendSortOrderEnum Enum with underlying type: string
type SummarizeExadataInsightResourceCapacityTrendSortOrderEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceCapacityTrendSortOrderEnum
const (
	SummarizeExadataInsightResourceCapacityTrendSortOrderAsc  SummarizeExadataInsightResourceCapacityTrendSortOrderEnum = "ASC"
	SummarizeExadataInsightResourceCapacityTrendSortOrderDesc SummarizeExadataInsightResourceCapacityTrendSortOrderEnum = "DESC"
)

var mappingSummarizeExadataInsightResourceCapacityTrendSortOrderEnum = map[string]SummarizeExadataInsightResourceCapacityTrendSortOrderEnum{
	"ASC":  SummarizeExadataInsightResourceCapacityTrendSortOrderAsc,
	"DESC": SummarizeExadataInsightResourceCapacityTrendSortOrderDesc,
}

var mappingSummarizeExadataInsightResourceCapacityTrendSortOrderEnumLowerCase = map[string]SummarizeExadataInsightResourceCapacityTrendSortOrderEnum{
	"asc":  SummarizeExadataInsightResourceCapacityTrendSortOrderAsc,
	"desc": SummarizeExadataInsightResourceCapacityTrendSortOrderDesc,
}

// GetSummarizeExadataInsightResourceCapacityTrendSortOrderEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendSortOrderEnum
func GetSummarizeExadataInsightResourceCapacityTrendSortOrderEnumValues() []SummarizeExadataInsightResourceCapacityTrendSortOrderEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendSortOrderEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceCapacityTrendSortOrderEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceCapacityTrendSortOrderEnum
func GetSummarizeExadataInsightResourceCapacityTrendSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingSummarizeExadataInsightResourceCapacityTrendSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceCapacityTrendSortOrderEnum(val string) (SummarizeExadataInsightResourceCapacityTrendSortOrderEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceCapacityTrendSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SummarizeExadataInsightResourceCapacityTrendSortByEnum Enum with underlying type: string
type SummarizeExadataInsightResourceCapacityTrendSortByEnum string

// Set of constants representing the allowable values for SummarizeExadataInsightResourceCapacityTrendSortByEnum
const (
	SummarizeExadataInsightResourceCapacityTrendSortById   SummarizeExadataInsightResourceCapacityTrendSortByEnum = "id"
	SummarizeExadataInsightResourceCapacityTrendSortByName SummarizeExadataInsightResourceCapacityTrendSortByEnum = "name"
)

var mappingSummarizeExadataInsightResourceCapacityTrendSortByEnum = map[string]SummarizeExadataInsightResourceCapacityTrendSortByEnum{
	"id":   SummarizeExadataInsightResourceCapacityTrendSortById,
	"name": SummarizeExadataInsightResourceCapacityTrendSortByName,
}

var mappingSummarizeExadataInsightResourceCapacityTrendSortByEnumLowerCase = map[string]SummarizeExadataInsightResourceCapacityTrendSortByEnum{
	"id":   SummarizeExadataInsightResourceCapacityTrendSortById,
	"name": SummarizeExadataInsightResourceCapacityTrendSortByName,
}

// GetSummarizeExadataInsightResourceCapacityTrendSortByEnumValues Enumerates the set of values for SummarizeExadataInsightResourceCapacityTrendSortByEnum
func GetSummarizeExadataInsightResourceCapacityTrendSortByEnumValues() []SummarizeExadataInsightResourceCapacityTrendSortByEnum {
	values := make([]SummarizeExadataInsightResourceCapacityTrendSortByEnum, 0)
	for _, v := range mappingSummarizeExadataInsightResourceCapacityTrendSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSummarizeExadataInsightResourceCapacityTrendSortByEnumStringValues Enumerates the set of values in String for SummarizeExadataInsightResourceCapacityTrendSortByEnum
func GetSummarizeExadataInsightResourceCapacityTrendSortByEnumStringValues() []string {
	return []string{
		"id",
		"name",
	}
}

// GetMappingSummarizeExadataInsightResourceCapacityTrendSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSummarizeExadataInsightResourceCapacityTrendSortByEnum(val string) (SummarizeExadataInsightResourceCapacityTrendSortByEnum, bool) {
	enum, ok := mappingSummarizeExadataInsightResourceCapacityTrendSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
