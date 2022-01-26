// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
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

var mappingListConfigWorkRequestsSortOrder = map[string]ListConfigWorkRequestsSortOrderEnum{
	"ASC":  ListConfigWorkRequestsSortOrderAsc,
	"DESC": ListConfigWorkRequestsSortOrderDesc,
}

// GetListConfigWorkRequestsSortOrderEnumValues Enumerates the set of values for ListConfigWorkRequestsSortOrderEnum
func GetListConfigWorkRequestsSortOrderEnumValues() []ListConfigWorkRequestsSortOrderEnum {
	values := make([]ListConfigWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListConfigWorkRequestsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListConfigWorkRequestsSortByEnum Enum with underlying type: string
type ListConfigWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListConfigWorkRequestsSortByEnum
const (
	ListConfigWorkRequestsSortByTimeaccepted ListConfigWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListConfigWorkRequestsSortBy = map[string]ListConfigWorkRequestsSortByEnum{
	"timeAccepted": ListConfigWorkRequestsSortByTimeaccepted,
}

// GetListConfigWorkRequestsSortByEnumValues Enumerates the set of values for ListConfigWorkRequestsSortByEnum
func GetListConfigWorkRequestsSortByEnumValues() []ListConfigWorkRequestsSortByEnum {
	values := make([]ListConfigWorkRequestsSortByEnum, 0)
	for _, v := range mappingListConfigWorkRequestsSortBy {
		values = append(values, v)
	}
	return values
}
