// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usageapi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListQueriesRequest wrapper for the ListQueries operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usageapi/ListQueries.go.html to see an example of how to use ListQueriesRequest.
type ListQueriesRequest struct {

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
	SortBy ListQueriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListQueriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListQueriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListQueriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListQueriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListQueriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListQueriesResponse wrapper for the ListQueries operation
type ListQueriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of QueryCollection instances
	QueryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of Queries. If this header appears in the response, then this
	// is a partial list of Queries. Include this value as the `page` parameter in a subsequent
	// GET request, to get the next batch of Queries.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListQueriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListQueriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListQueriesSortByEnum Enum with underlying type: string
type ListQueriesSortByEnum string

// Set of constants representing the allowable values for ListQueriesSortByEnum
const (
	ListQueriesSortByDisplayname ListQueriesSortByEnum = "displayName"
)

var mappingListQueriesSortBy = map[string]ListQueriesSortByEnum{
	"displayName": ListQueriesSortByDisplayname,
}

// GetListQueriesSortByEnumValues Enumerates the set of values for ListQueriesSortByEnum
func GetListQueriesSortByEnumValues() []ListQueriesSortByEnum {
	values := make([]ListQueriesSortByEnum, 0)
	for _, v := range mappingListQueriesSortBy {
		values = append(values, v)
	}
	return values
}

// ListQueriesSortOrderEnum Enum with underlying type: string
type ListQueriesSortOrderEnum string

// Set of constants representing the allowable values for ListQueriesSortOrderEnum
const (
	ListQueriesSortOrderAsc  ListQueriesSortOrderEnum = "ASC"
	ListQueriesSortOrderDesc ListQueriesSortOrderEnum = "DESC"
)

var mappingListQueriesSortOrder = map[string]ListQueriesSortOrderEnum{
	"ASC":  ListQueriesSortOrderAsc,
	"DESC": ListQueriesSortOrderDesc,
}

// GetListQueriesSortOrderEnumValues Enumerates the set of values for ListQueriesSortOrderEnum
func GetListQueriesSortOrderEnumValues() []ListQueriesSortOrderEnum {
	values := make([]ListQueriesSortOrderEnum, 0)
	for _, v := range mappingListQueriesSortOrder {
		values = append(values, v)
	}
	return values
}
