// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListHistoricalMetricsRequest wrapper for the ListHistoricalMetrics operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListHistoricalMetrics.go.html to see an example of how to use ListHistoricalMetricsRequest.
type ListHistoricalMetricsRequest struct {

	// Unique asset identifier.
	AssetId *string `mandatory:"true" contributesTo:"path" name:"assetId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListHistoricalMetricsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListHistoricalMetricsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHistoricalMetricsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHistoricalMetricsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHistoricalMetricsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHistoricalMetricsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHistoricalMetricsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListHistoricalMetricsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHistoricalMetricsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHistoricalMetricsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHistoricalMetricsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHistoricalMetricsResponse wrapper for the ListHistoricalMetrics operation
type ListHistoricalMetricsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of HistoricalMetricCollection instances
	HistoricalMetricCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHistoricalMetricsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHistoricalMetricsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHistoricalMetricsSortOrderEnum Enum with underlying type: string
type ListHistoricalMetricsSortOrderEnum string

// Set of constants representing the allowable values for ListHistoricalMetricsSortOrderEnum
const (
	ListHistoricalMetricsSortOrderAsc  ListHistoricalMetricsSortOrderEnum = "ASC"
	ListHistoricalMetricsSortOrderDesc ListHistoricalMetricsSortOrderEnum = "DESC"
)

var mappingListHistoricalMetricsSortOrderEnum = map[string]ListHistoricalMetricsSortOrderEnum{
	"ASC":  ListHistoricalMetricsSortOrderAsc,
	"DESC": ListHistoricalMetricsSortOrderDesc,
}

var mappingListHistoricalMetricsSortOrderEnumLowerCase = map[string]ListHistoricalMetricsSortOrderEnum{
	"asc":  ListHistoricalMetricsSortOrderAsc,
	"desc": ListHistoricalMetricsSortOrderDesc,
}

// GetListHistoricalMetricsSortOrderEnumValues Enumerates the set of values for ListHistoricalMetricsSortOrderEnum
func GetListHistoricalMetricsSortOrderEnumValues() []ListHistoricalMetricsSortOrderEnum {
	values := make([]ListHistoricalMetricsSortOrderEnum, 0)
	for _, v := range mappingListHistoricalMetricsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHistoricalMetricsSortOrderEnumStringValues Enumerates the set of values in String for ListHistoricalMetricsSortOrderEnum
func GetListHistoricalMetricsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHistoricalMetricsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHistoricalMetricsSortOrderEnum(val string) (ListHistoricalMetricsSortOrderEnum, bool) {
	enum, ok := mappingListHistoricalMetricsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHistoricalMetricsSortByEnum Enum with underlying type: string
type ListHistoricalMetricsSortByEnum string

// Set of constants representing the allowable values for ListHistoricalMetricsSortByEnum
const (
	ListHistoricalMetricsSortByTimecreated ListHistoricalMetricsSortByEnum = "timeCreated"
	ListHistoricalMetricsSortByTimeupdated ListHistoricalMetricsSortByEnum = "timeUpdated"
	ListHistoricalMetricsSortByDisplayname ListHistoricalMetricsSortByEnum = "displayName"
)

var mappingListHistoricalMetricsSortByEnum = map[string]ListHistoricalMetricsSortByEnum{
	"timeCreated": ListHistoricalMetricsSortByTimecreated,
	"timeUpdated": ListHistoricalMetricsSortByTimeupdated,
	"displayName": ListHistoricalMetricsSortByDisplayname,
}

var mappingListHistoricalMetricsSortByEnumLowerCase = map[string]ListHistoricalMetricsSortByEnum{
	"timecreated": ListHistoricalMetricsSortByTimecreated,
	"timeupdated": ListHistoricalMetricsSortByTimeupdated,
	"displayname": ListHistoricalMetricsSortByDisplayname,
}

// GetListHistoricalMetricsSortByEnumValues Enumerates the set of values for ListHistoricalMetricsSortByEnum
func GetListHistoricalMetricsSortByEnumValues() []ListHistoricalMetricsSortByEnum {
	values := make([]ListHistoricalMetricsSortByEnum, 0)
	for _, v := range mappingListHistoricalMetricsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHistoricalMetricsSortByEnumStringValues Enumerates the set of values in String for ListHistoricalMetricsSortByEnum
func GetListHistoricalMetricsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListHistoricalMetricsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHistoricalMetricsSortByEnum(val string) (ListHistoricalMetricsSortByEnum, bool) {
	enum, ok := mappingListHistoricalMetricsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
