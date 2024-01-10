// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTimeAvailableForRefreshesRequest wrapper for the ListTimeAvailableForRefreshes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fusionapps/ListTimeAvailableForRefreshes.go.html to see an example of how to use ListTimeAvailableForRefreshesRequest.
type ListTimeAvailableForRefreshesRequest struct {

	// unique FusionEnvironment identifier
	FusionEnvironmentId *string `mandatory:"true" contributesTo:"path" name:"fusionEnvironmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTimeAvailableForRefreshesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListTimeAvailableForRefreshesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTimeAvailableForRefreshesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTimeAvailableForRefreshesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTimeAvailableForRefreshesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTimeAvailableForRefreshesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTimeAvailableForRefreshesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTimeAvailableForRefreshesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTimeAvailableForRefreshesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTimeAvailableForRefreshesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTimeAvailableForRefreshesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTimeAvailableForRefreshesResponse wrapper for the ListTimeAvailableForRefreshes operation
type ListTimeAvailableForRefreshesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TimeAvailableForRefreshCollection instances
	TimeAvailableForRefreshCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTimeAvailableForRefreshesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTimeAvailableForRefreshesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTimeAvailableForRefreshesSortOrderEnum Enum with underlying type: string
type ListTimeAvailableForRefreshesSortOrderEnum string

// Set of constants representing the allowable values for ListTimeAvailableForRefreshesSortOrderEnum
const (
	ListTimeAvailableForRefreshesSortOrderAsc  ListTimeAvailableForRefreshesSortOrderEnum = "ASC"
	ListTimeAvailableForRefreshesSortOrderDesc ListTimeAvailableForRefreshesSortOrderEnum = "DESC"
)

var mappingListTimeAvailableForRefreshesSortOrderEnum = map[string]ListTimeAvailableForRefreshesSortOrderEnum{
	"ASC":  ListTimeAvailableForRefreshesSortOrderAsc,
	"DESC": ListTimeAvailableForRefreshesSortOrderDesc,
}

var mappingListTimeAvailableForRefreshesSortOrderEnumLowerCase = map[string]ListTimeAvailableForRefreshesSortOrderEnum{
	"asc":  ListTimeAvailableForRefreshesSortOrderAsc,
	"desc": ListTimeAvailableForRefreshesSortOrderDesc,
}

// GetListTimeAvailableForRefreshesSortOrderEnumValues Enumerates the set of values for ListTimeAvailableForRefreshesSortOrderEnum
func GetListTimeAvailableForRefreshesSortOrderEnumValues() []ListTimeAvailableForRefreshesSortOrderEnum {
	values := make([]ListTimeAvailableForRefreshesSortOrderEnum, 0)
	for _, v := range mappingListTimeAvailableForRefreshesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTimeAvailableForRefreshesSortOrderEnumStringValues Enumerates the set of values in String for ListTimeAvailableForRefreshesSortOrderEnum
func GetListTimeAvailableForRefreshesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTimeAvailableForRefreshesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTimeAvailableForRefreshesSortOrderEnum(val string) (ListTimeAvailableForRefreshesSortOrderEnum, bool) {
	enum, ok := mappingListTimeAvailableForRefreshesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTimeAvailableForRefreshesSortByEnum Enum with underlying type: string
type ListTimeAvailableForRefreshesSortByEnum string

// Set of constants representing the allowable values for ListTimeAvailableForRefreshesSortByEnum
const (
	ListTimeAvailableForRefreshesSortByTimeCreated ListTimeAvailableForRefreshesSortByEnum = "TIME_CREATED"
	ListTimeAvailableForRefreshesSortByDisplayName ListTimeAvailableForRefreshesSortByEnum = "DISPLAY_NAME"
)

var mappingListTimeAvailableForRefreshesSortByEnum = map[string]ListTimeAvailableForRefreshesSortByEnum{
	"TIME_CREATED": ListTimeAvailableForRefreshesSortByTimeCreated,
	"DISPLAY_NAME": ListTimeAvailableForRefreshesSortByDisplayName,
}

var mappingListTimeAvailableForRefreshesSortByEnumLowerCase = map[string]ListTimeAvailableForRefreshesSortByEnum{
	"time_created": ListTimeAvailableForRefreshesSortByTimeCreated,
	"display_name": ListTimeAvailableForRefreshesSortByDisplayName,
}

// GetListTimeAvailableForRefreshesSortByEnumValues Enumerates the set of values for ListTimeAvailableForRefreshesSortByEnum
func GetListTimeAvailableForRefreshesSortByEnumValues() []ListTimeAvailableForRefreshesSortByEnum {
	values := make([]ListTimeAvailableForRefreshesSortByEnum, 0)
	for _, v := range mappingListTimeAvailableForRefreshesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTimeAvailableForRefreshesSortByEnumStringValues Enumerates the set of values in String for ListTimeAvailableForRefreshesSortByEnum
func GetListTimeAvailableForRefreshesSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListTimeAvailableForRefreshesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTimeAvailableForRefreshesSortByEnum(val string) (ListTimeAvailableForRefreshesSortByEnum, bool) {
	enum, ok := mappingListTimeAvailableForRefreshesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
