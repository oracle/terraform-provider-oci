// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWorkRequestErrorsRequest wrapper for the ListWorkRequestErrors operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListWorkRequestErrors.go.html to see an example of how to use ListWorkRequestErrorsRequest.
type ListWorkRequestErrorsRequest struct {

	// The ID of the asynchronous request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timestamp is descending.
	SortBy ListWorkRequestErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListWorkRequestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestErrorsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestErrorsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkRequestErrorsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkRequestErrorsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkRequestErrorsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestErrorsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkRequestErrorsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkRequestErrorsResponse wrapper for the ListWorkRequestErrors operation
type ListWorkRequestErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestErrorCollection instances
	WorkRequestErrorCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListWorkRequestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestErrorsSortByEnum Enum with underlying type: string
type ListWorkRequestErrorsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestErrorsSortByEnum
const (
	ListWorkRequestErrorsSortByTimestamp ListWorkRequestErrorsSortByEnum = "timestamp"
)

var mappingListWorkRequestErrorsSortByEnum = map[string]ListWorkRequestErrorsSortByEnum{
	"timestamp": ListWorkRequestErrorsSortByTimestamp,
}

var mappingListWorkRequestErrorsSortByEnumLowerCase = map[string]ListWorkRequestErrorsSortByEnum{
	"timestamp": ListWorkRequestErrorsSortByTimestamp,
}

// GetListWorkRequestErrorsSortByEnumValues Enumerates the set of values for ListWorkRequestErrorsSortByEnum
func GetListWorkRequestErrorsSortByEnumValues() []ListWorkRequestErrorsSortByEnum {
	values := make([]ListWorkRequestErrorsSortByEnum, 0)
	for _, v := range mappingListWorkRequestErrorsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestErrorsSortByEnumStringValues Enumerates the set of values in String for ListWorkRequestErrorsSortByEnum
func GetListWorkRequestErrorsSortByEnumStringValues() []string {
	return []string{
		"timestamp",
	}
}

// GetMappingListWorkRequestErrorsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestErrorsSortByEnum(val string) (ListWorkRequestErrorsSortByEnum, bool) {
	enum, ok := mappingListWorkRequestErrorsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestErrorsSortOrderEnum Enum with underlying type: string
type ListWorkRequestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestErrorsSortOrderEnum
const (
	ListWorkRequestErrorsSortOrderAsc  ListWorkRequestErrorsSortOrderEnum = "ASC"
	ListWorkRequestErrorsSortOrderDesc ListWorkRequestErrorsSortOrderEnum = "DESC"
)

var mappingListWorkRequestErrorsSortOrderEnum = map[string]ListWorkRequestErrorsSortOrderEnum{
	"ASC":  ListWorkRequestErrorsSortOrderAsc,
	"DESC": ListWorkRequestErrorsSortOrderDesc,
}

var mappingListWorkRequestErrorsSortOrderEnumLowerCase = map[string]ListWorkRequestErrorsSortOrderEnum{
	"asc":  ListWorkRequestErrorsSortOrderAsc,
	"desc": ListWorkRequestErrorsSortOrderDesc,
}

// GetListWorkRequestErrorsSortOrderEnumValues Enumerates the set of values for ListWorkRequestErrorsSortOrderEnum
func GetListWorkRequestErrorsSortOrderEnumValues() []ListWorkRequestErrorsSortOrderEnum {
	values := make([]ListWorkRequestErrorsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestErrorsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestErrorsSortOrderEnumStringValues Enumerates the set of values in String for ListWorkRequestErrorsSortOrderEnum
func GetListWorkRequestErrorsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkRequestErrorsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestErrorsSortOrderEnum(val string) (ListWorkRequestErrorsSortOrderEnum, bool) {
	enum, ok := mappingListWorkRequestErrorsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
