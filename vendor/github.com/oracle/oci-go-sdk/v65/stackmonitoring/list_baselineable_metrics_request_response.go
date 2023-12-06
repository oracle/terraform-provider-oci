// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBaselineableMetricsRequest wrapper for the ListBaselineableMetrics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/stackmonitoring/ListBaselineableMetrics.go.html to see an example of how to use ListBaselineableMetricsRequest.
type ListBaselineableMetricsRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Resource Group
	ResourceGroup *string `mandatory:"false" contributesTo:"query" name:"resourceGroup"`

	// Metric Name
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return monitored resource types that has the matching namespace.
	MetricNamespace *string `mandatory:"false" contributesTo:"query" name:"metricNamespace"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the
	// previous "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The ID of the compartment in which data is listed.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Identifier for the metric
	BaselineableMetricId *string `mandatory:"false" contributesTo:"query" name:"baselineableMetricId"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListBaselineableMetricsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order is ascending.
	SortBy ListBaselineableMetricsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBaselineableMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBaselineableMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBaselineableMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBaselineableMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBaselineableMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBaselineableMetricsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBaselineableMetricsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBaselineableMetricsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBaselineableMetricsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBaselineableMetricsResponse wrapper for the ListBaselineableMetrics operation
type ListBaselineableMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of BaselineableMetricSummaryCollection instances
	BaselineableMetricSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// A decimal number representing the number of seconds the client should wait before polling this endpoint again.
	RetryAfter *int `presentIn:"header" name:"retry-after"`
}

func (response ListBaselineableMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBaselineableMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBaselineableMetricsSortOrderEnum Enum with underlying type: string
type ListBaselineableMetricsSortOrderEnum string

// Set of constants representing the allowable values for ListBaselineableMetricsSortOrderEnum
const (
	ListBaselineableMetricsSortOrderAsc  ListBaselineableMetricsSortOrderEnum = "ASC"
	ListBaselineableMetricsSortOrderDesc ListBaselineableMetricsSortOrderEnum = "DESC"
)

var mappingListBaselineableMetricsSortOrderEnum = map[string]ListBaselineableMetricsSortOrderEnum{
	"ASC":  ListBaselineableMetricsSortOrderAsc,
	"DESC": ListBaselineableMetricsSortOrderDesc,
}

var mappingListBaselineableMetricsSortOrderEnumLowerCase = map[string]ListBaselineableMetricsSortOrderEnum{
	"asc":  ListBaselineableMetricsSortOrderAsc,
	"desc": ListBaselineableMetricsSortOrderDesc,
}

// GetListBaselineableMetricsSortOrderEnumValues Enumerates the set of values for ListBaselineableMetricsSortOrderEnum
func GetListBaselineableMetricsSortOrderEnumValues() []ListBaselineableMetricsSortOrderEnum {
	values := make([]ListBaselineableMetricsSortOrderEnum, 0)
	for _, v := range mappingListBaselineableMetricsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBaselineableMetricsSortOrderEnumStringValues Enumerates the set of values in String for ListBaselineableMetricsSortOrderEnum
func GetListBaselineableMetricsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBaselineableMetricsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBaselineableMetricsSortOrderEnum(val string) (ListBaselineableMetricsSortOrderEnum, bool) {
	enum, ok := mappingListBaselineableMetricsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBaselineableMetricsSortByEnum Enum with underlying type: string
type ListBaselineableMetricsSortByEnum string

// Set of constants representing the allowable values for ListBaselineableMetricsSortByEnum
const (
	ListBaselineableMetricsSortByName          ListBaselineableMetricsSortByEnum = "name"
	ListBaselineableMetricsSortByNamespace     ListBaselineableMetricsSortByEnum = "namespace"
	ListBaselineableMetricsSortByResourcegroup ListBaselineableMetricsSortByEnum = "resourceGroup"
)

var mappingListBaselineableMetricsSortByEnum = map[string]ListBaselineableMetricsSortByEnum{
	"name":          ListBaselineableMetricsSortByName,
	"namespace":     ListBaselineableMetricsSortByNamespace,
	"resourceGroup": ListBaselineableMetricsSortByResourcegroup,
}

var mappingListBaselineableMetricsSortByEnumLowerCase = map[string]ListBaselineableMetricsSortByEnum{
	"name":          ListBaselineableMetricsSortByName,
	"namespace":     ListBaselineableMetricsSortByNamespace,
	"resourcegroup": ListBaselineableMetricsSortByResourcegroup,
}

// GetListBaselineableMetricsSortByEnumValues Enumerates the set of values for ListBaselineableMetricsSortByEnum
func GetListBaselineableMetricsSortByEnumValues() []ListBaselineableMetricsSortByEnum {
	values := make([]ListBaselineableMetricsSortByEnum, 0)
	for _, v := range mappingListBaselineableMetricsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBaselineableMetricsSortByEnumStringValues Enumerates the set of values in String for ListBaselineableMetricsSortByEnum
func GetListBaselineableMetricsSortByEnumStringValues() []string {
	return []string{
		"name",
		"namespace",
		"resourceGroup",
	}
}

// GetMappingListBaselineableMetricsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBaselineableMetricsSortByEnum(val string) (ListBaselineableMetricsSortByEnum, bool) {
	enum, ok := mappingListBaselineableMetricsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
