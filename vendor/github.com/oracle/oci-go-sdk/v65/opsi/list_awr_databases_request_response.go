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

// ListAwrDatabasesRequest wrapper for the ListAwrDatabases operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrDatabases.go.html to see an example of how to use ListAwrDatabasesRequest.
type ListAwrDatabasesRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"path" name:"awrHubId"`

	// The optional single value query parameter to filter the entity name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The optional greater than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeGreaterThanOrEqualTo"`

	// The optional less than or equal to query parameter to filter the timestamp. The timestamp format to be followed is: YYYY-MM-DDTHH:MM:SSZ, example 2020-12-03T19:00:53Z
	TimeLessThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeLessThanOrEqualTo"`

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

	// The option to sort the AWR summary data.
	SortBy ListAwrDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAwrDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAwrDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAwrDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAwrDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAwrDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAwrDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAwrDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAwrDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAwrDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAwrDatabasesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAwrDatabasesResponse wrapper for the ListAwrDatabases operation
type ListAwrDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrDatabaseCollection instances
	AwrDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAwrDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAwrDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAwrDatabasesSortByEnum Enum with underlying type: string
type ListAwrDatabasesSortByEnum string

// Set of constants representing the allowable values for ListAwrDatabasesSortByEnum
const (
	ListAwrDatabasesSortByEndIntervalTime ListAwrDatabasesSortByEnum = "END_INTERVAL_TIME"
	ListAwrDatabasesSortByName            ListAwrDatabasesSortByEnum = "NAME"
)

var mappingListAwrDatabasesSortByEnum = map[string]ListAwrDatabasesSortByEnum{
	"END_INTERVAL_TIME": ListAwrDatabasesSortByEndIntervalTime,
	"NAME":              ListAwrDatabasesSortByName,
}

var mappingListAwrDatabasesSortByEnumLowerCase = map[string]ListAwrDatabasesSortByEnum{
	"end_interval_time": ListAwrDatabasesSortByEndIntervalTime,
	"name":              ListAwrDatabasesSortByName,
}

// GetListAwrDatabasesSortByEnumValues Enumerates the set of values for ListAwrDatabasesSortByEnum
func GetListAwrDatabasesSortByEnumValues() []ListAwrDatabasesSortByEnum {
	values := make([]ListAwrDatabasesSortByEnum, 0)
	for _, v := range mappingListAwrDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrDatabasesSortByEnumStringValues Enumerates the set of values in String for ListAwrDatabasesSortByEnum
func GetListAwrDatabasesSortByEnumStringValues() []string {
	return []string{
		"END_INTERVAL_TIME",
		"NAME",
	}
}

// GetMappingListAwrDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrDatabasesSortByEnum(val string) (ListAwrDatabasesSortByEnum, bool) {
	enum, ok := mappingListAwrDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAwrDatabasesSortOrderEnum Enum with underlying type: string
type ListAwrDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListAwrDatabasesSortOrderEnum
const (
	ListAwrDatabasesSortOrderAsc  ListAwrDatabasesSortOrderEnum = "ASC"
	ListAwrDatabasesSortOrderDesc ListAwrDatabasesSortOrderEnum = "DESC"
)

var mappingListAwrDatabasesSortOrderEnum = map[string]ListAwrDatabasesSortOrderEnum{
	"ASC":  ListAwrDatabasesSortOrderAsc,
	"DESC": ListAwrDatabasesSortOrderDesc,
}

var mappingListAwrDatabasesSortOrderEnumLowerCase = map[string]ListAwrDatabasesSortOrderEnum{
	"asc":  ListAwrDatabasesSortOrderAsc,
	"desc": ListAwrDatabasesSortOrderDesc,
}

// GetListAwrDatabasesSortOrderEnumValues Enumerates the set of values for ListAwrDatabasesSortOrderEnum
func GetListAwrDatabasesSortOrderEnumValues() []ListAwrDatabasesSortOrderEnum {
	values := make([]ListAwrDatabasesSortOrderEnum, 0)
	for _, v := range mappingListAwrDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListAwrDatabasesSortOrderEnum
func GetListAwrDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAwrDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrDatabasesSortOrderEnum(val string) (ListAwrDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListAwrDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
