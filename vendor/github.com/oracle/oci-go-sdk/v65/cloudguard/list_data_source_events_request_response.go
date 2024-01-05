// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDataSourceEventsRequest wrapper for the ListDataSourceEvents operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListDataSourceEvents.go.html to see an example of how to use ListDataSourceEventsRequest.
type ListDataSourceEventsRequest struct {

	// DataSource OCID
	DataSourceId *string `mandatory:"true" contributesTo:"path" name:"dataSourceId"`

	// A filter to return only resource their region matches the given region.
	Region *string `mandatory:"false" contributesTo:"query" name:"region"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDataSourceEventsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. If no value is specified timeCreated is default.
	SortBy ListDataSourceEventsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSourceEventsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSourceEventsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDataSourceEventsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSourceEventsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDataSourceEventsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListDataSourceEventsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDataSourceEventsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDataSourceEventsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDataSourceEventsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDataSourceEventsResponse wrapper for the ListDataSourceEvents operation
type ListDataSourceEventsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DataSourceEventCollection instances
	DataSourceEventCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataSourceEventsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSourceEventsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSourceEventsSortOrderEnum Enum with underlying type: string
type ListDataSourceEventsSortOrderEnum string

// Set of constants representing the allowable values for ListDataSourceEventsSortOrderEnum
const (
	ListDataSourceEventsSortOrderAsc  ListDataSourceEventsSortOrderEnum = "ASC"
	ListDataSourceEventsSortOrderDesc ListDataSourceEventsSortOrderEnum = "DESC"
)

var mappingListDataSourceEventsSortOrderEnum = map[string]ListDataSourceEventsSortOrderEnum{
	"ASC":  ListDataSourceEventsSortOrderAsc,
	"DESC": ListDataSourceEventsSortOrderDesc,
}

var mappingListDataSourceEventsSortOrderEnumLowerCase = map[string]ListDataSourceEventsSortOrderEnum{
	"asc":  ListDataSourceEventsSortOrderAsc,
	"desc": ListDataSourceEventsSortOrderDesc,
}

// GetListDataSourceEventsSortOrderEnumValues Enumerates the set of values for ListDataSourceEventsSortOrderEnum
func GetListDataSourceEventsSortOrderEnumValues() []ListDataSourceEventsSortOrderEnum {
	values := make([]ListDataSourceEventsSortOrderEnum, 0)
	for _, v := range mappingListDataSourceEventsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourceEventsSortOrderEnumStringValues Enumerates the set of values in String for ListDataSourceEventsSortOrderEnum
func GetListDataSourceEventsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDataSourceEventsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourceEventsSortOrderEnum(val string) (ListDataSourceEventsSortOrderEnum, bool) {
	enum, ok := mappingListDataSourceEventsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDataSourceEventsSortByEnum Enum with underlying type: string
type ListDataSourceEventsSortByEnum string

// Set of constants representing the allowable values for ListDataSourceEventsSortByEnum
const (
	ListDataSourceEventsSortByTimecreated ListDataSourceEventsSortByEnum = "timeCreated"
)

var mappingListDataSourceEventsSortByEnum = map[string]ListDataSourceEventsSortByEnum{
	"timeCreated": ListDataSourceEventsSortByTimecreated,
}

var mappingListDataSourceEventsSortByEnumLowerCase = map[string]ListDataSourceEventsSortByEnum{
	"timecreated": ListDataSourceEventsSortByTimecreated,
}

// GetListDataSourceEventsSortByEnumValues Enumerates the set of values for ListDataSourceEventsSortByEnum
func GetListDataSourceEventsSortByEnumValues() []ListDataSourceEventsSortByEnum {
	values := make([]ListDataSourceEventsSortByEnum, 0)
	for _, v := range mappingListDataSourceEventsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDataSourceEventsSortByEnumStringValues Enumerates the set of values in String for ListDataSourceEventsSortByEnum
func GetListDataSourceEventsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
	}
}

// GetMappingListDataSourceEventsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDataSourceEventsSortByEnum(val string) (ListDataSourceEventsSortByEnum, bool) {
	enum, ok := mappingListDataSourceEventsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
