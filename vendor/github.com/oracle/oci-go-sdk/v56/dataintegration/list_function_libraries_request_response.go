// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListFunctionLibrariesRequest wrapper for the ListFunctionLibraries operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListFunctionLibraries.go.html to see an example of how to use ListFunctionLibrariesRequest.
type ListFunctionLibrariesRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used to filter by the project or the folder object.
	AggregatorKey *string `mandatory:"false" contributesTo:"query" name:"aggregatorKey"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListFunctionLibrariesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListFunctionLibrariesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFunctionLibrariesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFunctionLibrariesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFunctionLibrariesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFunctionLibrariesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFunctionLibrariesResponse wrapper for the ListFunctionLibraries operation
type ListFunctionLibrariesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FunctionLibrarySummaryCollection instances
	FunctionLibrarySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListFunctionLibrariesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFunctionLibrariesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFunctionLibrariesSortOrderEnum Enum with underlying type: string
type ListFunctionLibrariesSortOrderEnum string

// Set of constants representing the allowable values for ListFunctionLibrariesSortOrderEnum
const (
	ListFunctionLibrariesSortOrderAsc  ListFunctionLibrariesSortOrderEnum = "ASC"
	ListFunctionLibrariesSortOrderDesc ListFunctionLibrariesSortOrderEnum = "DESC"
)

var mappingListFunctionLibrariesSortOrder = map[string]ListFunctionLibrariesSortOrderEnum{
	"ASC":  ListFunctionLibrariesSortOrderAsc,
	"DESC": ListFunctionLibrariesSortOrderDesc,
}

// GetListFunctionLibrariesSortOrderEnumValues Enumerates the set of values for ListFunctionLibrariesSortOrderEnum
func GetListFunctionLibrariesSortOrderEnumValues() []ListFunctionLibrariesSortOrderEnum {
	values := make([]ListFunctionLibrariesSortOrderEnum, 0)
	for _, v := range mappingListFunctionLibrariesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListFunctionLibrariesSortByEnum Enum with underlying type: string
type ListFunctionLibrariesSortByEnum string

// Set of constants representing the allowable values for ListFunctionLibrariesSortByEnum
const (
	ListFunctionLibrariesSortByTimeCreated ListFunctionLibrariesSortByEnum = "TIME_CREATED"
	ListFunctionLibrariesSortByDisplayName ListFunctionLibrariesSortByEnum = "DISPLAY_NAME"
)

var mappingListFunctionLibrariesSortBy = map[string]ListFunctionLibrariesSortByEnum{
	"TIME_CREATED": ListFunctionLibrariesSortByTimeCreated,
	"DISPLAY_NAME": ListFunctionLibrariesSortByDisplayName,
}

// GetListFunctionLibrariesSortByEnumValues Enumerates the set of values for ListFunctionLibrariesSortByEnum
func GetListFunctionLibrariesSortByEnumValues() []ListFunctionLibrariesSortByEnum {
	values := make([]ListFunctionLibrariesSortByEnum, 0)
	for _, v := range mappingListFunctionLibrariesSortBy {
		values = append(values, v)
	}
	return values
}
