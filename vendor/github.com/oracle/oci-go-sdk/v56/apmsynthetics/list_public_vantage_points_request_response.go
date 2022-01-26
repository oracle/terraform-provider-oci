// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListPublicVantagePointsRequest wrapper for the ListPublicVantagePoints operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmsynthetics/ListPublicVantagePoints.go.html to see an example of how to use ListPublicVantagePointsRequest.
type ListPublicVantagePointsRequest struct {

	// The APM domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). Default sort order is ascending.
	SortOrder ListPublicVantagePointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort by (`sortBy`). Default order for displayName or name is ascending. The displayName or name
	// sort by is case insensitive.
	SortBy ListPublicVantagePointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPublicVantagePointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPublicVantagePointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPublicVantagePointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPublicVantagePointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPublicVantagePointsResponse wrapper for the ListPublicVantagePoints operation
type ListPublicVantagePointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PublicVantagePointCollection instances
	PublicVantagePointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPublicVantagePointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPublicVantagePointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPublicVantagePointsSortOrderEnum Enum with underlying type: string
type ListPublicVantagePointsSortOrderEnum string

// Set of constants representing the allowable values for ListPublicVantagePointsSortOrderEnum
const (
	ListPublicVantagePointsSortOrderAsc  ListPublicVantagePointsSortOrderEnum = "ASC"
	ListPublicVantagePointsSortOrderDesc ListPublicVantagePointsSortOrderEnum = "DESC"
)

var mappingListPublicVantagePointsSortOrder = map[string]ListPublicVantagePointsSortOrderEnum{
	"ASC":  ListPublicVantagePointsSortOrderAsc,
	"DESC": ListPublicVantagePointsSortOrderDesc,
}

// GetListPublicVantagePointsSortOrderEnumValues Enumerates the set of values for ListPublicVantagePointsSortOrderEnum
func GetListPublicVantagePointsSortOrderEnumValues() []ListPublicVantagePointsSortOrderEnum {
	values := make([]ListPublicVantagePointsSortOrderEnum, 0)
	for _, v := range mappingListPublicVantagePointsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListPublicVantagePointsSortByEnum Enum with underlying type: string
type ListPublicVantagePointsSortByEnum string

// Set of constants representing the allowable values for ListPublicVantagePointsSortByEnum
const (
	ListPublicVantagePointsSortByName        ListPublicVantagePointsSortByEnum = "name"
	ListPublicVantagePointsSortByDisplayname ListPublicVantagePointsSortByEnum = "displayName"
)

var mappingListPublicVantagePointsSortBy = map[string]ListPublicVantagePointsSortByEnum{
	"name":        ListPublicVantagePointsSortByName,
	"displayName": ListPublicVantagePointsSortByDisplayname,
}

// GetListPublicVantagePointsSortByEnumValues Enumerates the set of values for ListPublicVantagePointsSortByEnum
func GetListPublicVantagePointsSortByEnumValues() []ListPublicVantagePointsSortByEnum {
	values := make([]ListPublicVantagePointsSortByEnum, 0)
	for _, v := range mappingListPublicVantagePointsSortBy {
		values = append(values, v)
	}
	return values
}
