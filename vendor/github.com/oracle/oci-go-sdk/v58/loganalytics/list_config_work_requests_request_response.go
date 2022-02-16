// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListConfigWorkRequestsRequest wrapper for the ListConfigWorkRequests operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListConfigWorkRequests.go.html to see an example of how to use ListConfigWorkRequestsRequest.
type ListConfigWorkRequestsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListConfigWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned work requests
	SortBy ListConfigWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConfigWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConfigWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConfigWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConfigWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConfigWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListConfigWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConfigWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConfigWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConfigWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConfigWorkRequestsResponse wrapper for the ListConfigWorkRequests operation
type ListConfigWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsConfigWorkRequestCollection instances
	LogAnalyticsConfigWorkRequestCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListConfigWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConfigWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConfigWorkRequestsSortOrderEnum Enum with underlying type: string
type ListConfigWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListConfigWorkRequestsSortOrderEnum
const (
	ListConfigWorkRequestsSortOrderAsc  ListConfigWorkRequestsSortOrderEnum = "ASC"
	ListConfigWorkRequestsSortOrderDesc ListConfigWorkRequestsSortOrderEnum = "DESC"
)

var mappingListConfigWorkRequestsSortOrderEnum = map[string]ListConfigWorkRequestsSortOrderEnum{
	"ASC":  ListConfigWorkRequestsSortOrderAsc,
	"DESC": ListConfigWorkRequestsSortOrderDesc,
}

// GetListConfigWorkRequestsSortOrderEnumValues Enumerates the set of values for ListConfigWorkRequestsSortOrderEnum
func GetListConfigWorkRequestsSortOrderEnumValues() []ListConfigWorkRequestsSortOrderEnum {
	values := make([]ListConfigWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListConfigWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConfigWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListConfigWorkRequestsSortOrderEnum
func GetListConfigWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConfigWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConfigWorkRequestsSortOrderEnum(val string) (ListConfigWorkRequestsSortOrderEnum, bool) {
	mappingListConfigWorkRequestsSortOrderEnumIgnoreCase := make(map[string]ListConfigWorkRequestsSortOrderEnum)
	for k, v := range mappingListConfigWorkRequestsSortOrderEnum {
		mappingListConfigWorkRequestsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListConfigWorkRequestsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListConfigWorkRequestsSortByEnum Enum with underlying type: string
type ListConfigWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListConfigWorkRequestsSortByEnum
const (
	ListConfigWorkRequestsSortByTimeaccepted ListConfigWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListConfigWorkRequestsSortByEnum = map[string]ListConfigWorkRequestsSortByEnum{
	"timeAccepted": ListConfigWorkRequestsSortByTimeaccepted,
}

// GetListConfigWorkRequestsSortByEnumValues Enumerates the set of values for ListConfigWorkRequestsSortByEnum
func GetListConfigWorkRequestsSortByEnumValues() []ListConfigWorkRequestsSortByEnum {
	values := make([]ListConfigWorkRequestsSortByEnum, 0)
	for _, v := range mappingListConfigWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConfigWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListConfigWorkRequestsSortByEnum
func GetListConfigWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
	}
}

// GetMappingListConfigWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConfigWorkRequestsSortByEnum(val string) (ListConfigWorkRequestsSortByEnum, bool) {
	mappingListConfigWorkRequestsSortByEnumIgnoreCase := make(map[string]ListConfigWorkRequestsSortByEnum)
	for k, v := range mappingListConfigWorkRequestsSortByEnum {
		mappingListConfigWorkRequestsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListConfigWorkRequestsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
