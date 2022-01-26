// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListUserDefinedFunctionsRequest wrapper for the ListUserDefinedFunctions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListUserDefinedFunctions.go.html to see an example of how to use ListUserDefinedFunctionsRequest.
type ListUserDefinedFunctionsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique key of the FunctionLibrary.
	FunctionLibraryKey *string `mandatory:"false" contributesTo:"query" name:"functionLibraryKey"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListUserDefinedFunctionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListUserDefinedFunctionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUserDefinedFunctionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUserDefinedFunctionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUserDefinedFunctionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUserDefinedFunctionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListUserDefinedFunctionsResponse wrapper for the ListUserDefinedFunctions operation
type ListUserDefinedFunctionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UserDefinedFunctionSummaryCollection instances
	UserDefinedFunctionSummaryCollection `presentIn:"body"`

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

func (response ListUserDefinedFunctionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUserDefinedFunctionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUserDefinedFunctionsSortOrderEnum Enum with underlying type: string
type ListUserDefinedFunctionsSortOrderEnum string

// Set of constants representing the allowable values for ListUserDefinedFunctionsSortOrderEnum
const (
	ListUserDefinedFunctionsSortOrderAsc  ListUserDefinedFunctionsSortOrderEnum = "ASC"
	ListUserDefinedFunctionsSortOrderDesc ListUserDefinedFunctionsSortOrderEnum = "DESC"
)

var mappingListUserDefinedFunctionsSortOrder = map[string]ListUserDefinedFunctionsSortOrderEnum{
	"ASC":  ListUserDefinedFunctionsSortOrderAsc,
	"DESC": ListUserDefinedFunctionsSortOrderDesc,
}

// GetListUserDefinedFunctionsSortOrderEnumValues Enumerates the set of values for ListUserDefinedFunctionsSortOrderEnum
func GetListUserDefinedFunctionsSortOrderEnumValues() []ListUserDefinedFunctionsSortOrderEnum {
	values := make([]ListUserDefinedFunctionsSortOrderEnum, 0)
	for _, v := range mappingListUserDefinedFunctionsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListUserDefinedFunctionsSortByEnum Enum with underlying type: string
type ListUserDefinedFunctionsSortByEnum string

// Set of constants representing the allowable values for ListUserDefinedFunctionsSortByEnum
const (
	ListUserDefinedFunctionsSortByTimeCreated ListUserDefinedFunctionsSortByEnum = "TIME_CREATED"
	ListUserDefinedFunctionsSortByDisplayName ListUserDefinedFunctionsSortByEnum = "DISPLAY_NAME"
)

var mappingListUserDefinedFunctionsSortBy = map[string]ListUserDefinedFunctionsSortByEnum{
	"TIME_CREATED": ListUserDefinedFunctionsSortByTimeCreated,
	"DISPLAY_NAME": ListUserDefinedFunctionsSortByDisplayName,
}

// GetListUserDefinedFunctionsSortByEnumValues Enumerates the set of values for ListUserDefinedFunctionsSortByEnum
func GetListUserDefinedFunctionsSortByEnumValues() []ListUserDefinedFunctionsSortByEnum {
	values := make([]ListUserDefinedFunctionsSortByEnum, 0)
	for _, v := range mappingListUserDefinedFunctionsSortBy {
		values = append(values, v)
	}
	return values
}
