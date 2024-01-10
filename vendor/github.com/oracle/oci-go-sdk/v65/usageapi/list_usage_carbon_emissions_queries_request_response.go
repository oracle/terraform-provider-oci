// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListUsageCarbonEmissionsQueriesRequest wrapper for the ListUsageCarbonEmissionsQueries operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListUsageCarbonEmissionsQueries.go.html to see an example of how to use ListUsageCarbonEmissionsQueriesRequest.
type ListUsageCarbonEmissionsQueriesRequest struct {

	// The compartment ID in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximumimum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results.
	// This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. If not specified, the default is displayName.
	SortBy ListUsageCarbonEmissionsQueriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListUsageCarbonEmissionsQueriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUsageCarbonEmissionsQueriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUsageCarbonEmissionsQueriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUsageCarbonEmissionsQueriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUsageCarbonEmissionsQueriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUsageCarbonEmissionsQueriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUsageCarbonEmissionsQueriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUsageCarbonEmissionsQueriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUsageCarbonEmissionsQueriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUsageCarbonEmissionsQueriesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUsageCarbonEmissionsQueriesResponse wrapper for the ListUsageCarbonEmissionsQueries operation
type ListUsageCarbonEmissionsQueriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UsageCarbonEmissionsQueryCollection instances
	UsageCarbonEmissionsQueryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListUsageCarbonEmissionsQueriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUsageCarbonEmissionsQueriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUsageCarbonEmissionsQueriesSortByEnum Enum with underlying type: string
type ListUsageCarbonEmissionsQueriesSortByEnum string

// Set of constants representing the allowable values for ListUsageCarbonEmissionsQueriesSortByEnum
const (
	ListUsageCarbonEmissionsQueriesSortByDisplayname ListUsageCarbonEmissionsQueriesSortByEnum = "displayName"
)

var mappingListUsageCarbonEmissionsQueriesSortByEnum = map[string]ListUsageCarbonEmissionsQueriesSortByEnum{
	"displayName": ListUsageCarbonEmissionsQueriesSortByDisplayname,
}

var mappingListUsageCarbonEmissionsQueriesSortByEnumLowerCase = map[string]ListUsageCarbonEmissionsQueriesSortByEnum{
	"displayname": ListUsageCarbonEmissionsQueriesSortByDisplayname,
}

// GetListUsageCarbonEmissionsQueriesSortByEnumValues Enumerates the set of values for ListUsageCarbonEmissionsQueriesSortByEnum
func GetListUsageCarbonEmissionsQueriesSortByEnumValues() []ListUsageCarbonEmissionsQueriesSortByEnum {
	values := make([]ListUsageCarbonEmissionsQueriesSortByEnum, 0)
	for _, v := range mappingListUsageCarbonEmissionsQueriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsageCarbonEmissionsQueriesSortByEnumStringValues Enumerates the set of values in String for ListUsageCarbonEmissionsQueriesSortByEnum
func GetListUsageCarbonEmissionsQueriesSortByEnumStringValues() []string {
	return []string{
		"displayName",
	}
}

// GetMappingListUsageCarbonEmissionsQueriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsageCarbonEmissionsQueriesSortByEnum(val string) (ListUsageCarbonEmissionsQueriesSortByEnum, bool) {
	enum, ok := mappingListUsageCarbonEmissionsQueriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListUsageCarbonEmissionsQueriesSortOrderEnum Enum with underlying type: string
type ListUsageCarbonEmissionsQueriesSortOrderEnum string

// Set of constants representing the allowable values for ListUsageCarbonEmissionsQueriesSortOrderEnum
const (
	ListUsageCarbonEmissionsQueriesSortOrderAsc  ListUsageCarbonEmissionsQueriesSortOrderEnum = "ASC"
	ListUsageCarbonEmissionsQueriesSortOrderDesc ListUsageCarbonEmissionsQueriesSortOrderEnum = "DESC"
)

var mappingListUsageCarbonEmissionsQueriesSortOrderEnum = map[string]ListUsageCarbonEmissionsQueriesSortOrderEnum{
	"ASC":  ListUsageCarbonEmissionsQueriesSortOrderAsc,
	"DESC": ListUsageCarbonEmissionsQueriesSortOrderDesc,
}

var mappingListUsageCarbonEmissionsQueriesSortOrderEnumLowerCase = map[string]ListUsageCarbonEmissionsQueriesSortOrderEnum{
	"asc":  ListUsageCarbonEmissionsQueriesSortOrderAsc,
	"desc": ListUsageCarbonEmissionsQueriesSortOrderDesc,
}

// GetListUsageCarbonEmissionsQueriesSortOrderEnumValues Enumerates the set of values for ListUsageCarbonEmissionsQueriesSortOrderEnum
func GetListUsageCarbonEmissionsQueriesSortOrderEnumValues() []ListUsageCarbonEmissionsQueriesSortOrderEnum {
	values := make([]ListUsageCarbonEmissionsQueriesSortOrderEnum, 0)
	for _, v := range mappingListUsageCarbonEmissionsQueriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUsageCarbonEmissionsQueriesSortOrderEnumStringValues Enumerates the set of values in String for ListUsageCarbonEmissionsQueriesSortOrderEnum
func GetListUsageCarbonEmissionsQueriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUsageCarbonEmissionsQueriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUsageCarbonEmissionsQueriesSortOrderEnum(val string) (ListUsageCarbonEmissionsQueriesSortOrderEnum, bool) {
	enum, ok := mappingListUsageCarbonEmissionsQueriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
