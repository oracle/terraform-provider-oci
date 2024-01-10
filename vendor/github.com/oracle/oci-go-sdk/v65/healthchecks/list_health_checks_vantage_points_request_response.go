// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package healthchecks

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListHealthChecksVantagePointsRequest wrapper for the ListHealthChecksVantagePoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/healthchecks/ListHealthChecksVantagePoints.go.html to see an example of how to use ListHealthChecksVantagePointsRequest.
type ListHealthChecksVantagePointsRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header
	// from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by when listing vantage points.
	SortBy ListHealthChecksVantagePointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Controls the sort order of results.
	SortOrder ListHealthChecksVantagePointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Filters results that exactly match the `name` field.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filters results that exactly match the `displayName` field.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListHealthChecksVantagePointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListHealthChecksVantagePointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListHealthChecksVantagePointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListHealthChecksVantagePointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListHealthChecksVantagePointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListHealthChecksVantagePointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListHealthChecksVantagePointsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListHealthChecksVantagePointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListHealthChecksVantagePointsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListHealthChecksVantagePointsResponse wrapper for the ListHealthChecksVantagePoints operation
type ListHealthChecksVantagePointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []HealthChecksVantagePointSummary instances
	Items []HealthChecksVantagePointSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if
	// this header appears in the response, then there may be additional
	// items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#List_Pagination).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListHealthChecksVantagePointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListHealthChecksVantagePointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListHealthChecksVantagePointsSortByEnum Enum with underlying type: string
type ListHealthChecksVantagePointsSortByEnum string

// Set of constants representing the allowable values for ListHealthChecksVantagePointsSortByEnum
const (
	ListHealthChecksVantagePointsSortByName        ListHealthChecksVantagePointsSortByEnum = "name"
	ListHealthChecksVantagePointsSortByDisplayname ListHealthChecksVantagePointsSortByEnum = "displayName"
)

var mappingListHealthChecksVantagePointsSortByEnum = map[string]ListHealthChecksVantagePointsSortByEnum{
	"name":        ListHealthChecksVantagePointsSortByName,
	"displayName": ListHealthChecksVantagePointsSortByDisplayname,
}

var mappingListHealthChecksVantagePointsSortByEnumLowerCase = map[string]ListHealthChecksVantagePointsSortByEnum{
	"name":        ListHealthChecksVantagePointsSortByName,
	"displayname": ListHealthChecksVantagePointsSortByDisplayname,
}

// GetListHealthChecksVantagePointsSortByEnumValues Enumerates the set of values for ListHealthChecksVantagePointsSortByEnum
func GetListHealthChecksVantagePointsSortByEnumValues() []ListHealthChecksVantagePointsSortByEnum {
	values := make([]ListHealthChecksVantagePointsSortByEnum, 0)
	for _, v := range mappingListHealthChecksVantagePointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListHealthChecksVantagePointsSortByEnumStringValues Enumerates the set of values in String for ListHealthChecksVantagePointsSortByEnum
func GetListHealthChecksVantagePointsSortByEnumStringValues() []string {
	return []string{
		"name",
		"displayName",
	}
}

// GetMappingListHealthChecksVantagePointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHealthChecksVantagePointsSortByEnum(val string) (ListHealthChecksVantagePointsSortByEnum, bool) {
	enum, ok := mappingListHealthChecksVantagePointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListHealthChecksVantagePointsSortOrderEnum Enum with underlying type: string
type ListHealthChecksVantagePointsSortOrderEnum string

// Set of constants representing the allowable values for ListHealthChecksVantagePointsSortOrderEnum
const (
	ListHealthChecksVantagePointsSortOrderAsc  ListHealthChecksVantagePointsSortOrderEnum = "ASC"
	ListHealthChecksVantagePointsSortOrderDesc ListHealthChecksVantagePointsSortOrderEnum = "DESC"
)

var mappingListHealthChecksVantagePointsSortOrderEnum = map[string]ListHealthChecksVantagePointsSortOrderEnum{
	"ASC":  ListHealthChecksVantagePointsSortOrderAsc,
	"DESC": ListHealthChecksVantagePointsSortOrderDesc,
}

var mappingListHealthChecksVantagePointsSortOrderEnumLowerCase = map[string]ListHealthChecksVantagePointsSortOrderEnum{
	"asc":  ListHealthChecksVantagePointsSortOrderAsc,
	"desc": ListHealthChecksVantagePointsSortOrderDesc,
}

// GetListHealthChecksVantagePointsSortOrderEnumValues Enumerates the set of values for ListHealthChecksVantagePointsSortOrderEnum
func GetListHealthChecksVantagePointsSortOrderEnumValues() []ListHealthChecksVantagePointsSortOrderEnum {
	values := make([]ListHealthChecksVantagePointsSortOrderEnum, 0)
	for _, v := range mappingListHealthChecksVantagePointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListHealthChecksVantagePointsSortOrderEnumStringValues Enumerates the set of values in String for ListHealthChecksVantagePointsSortOrderEnum
func GetListHealthChecksVantagePointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListHealthChecksVantagePointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListHealthChecksVantagePointsSortOrderEnum(val string) (ListHealthChecksVantagePointsSortOrderEnum, bool) {
	enum, ok := mappingListHealthChecksVantagePointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
