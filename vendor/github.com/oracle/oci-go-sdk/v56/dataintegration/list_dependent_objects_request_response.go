// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDependentObjectsRequest wrapper for the ListDependentObjects operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListDependentObjects.go.html to see an example of how to use ListDependentObjectsRequest.
type ListDependentObjectsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The application key.
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This parameter can be used to filter objects by the names that match partially or fully with the given value.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// Used to filter by the identifier of the published object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// Used to filter by the object type of the object.
	// It can be suffixed with an optional filter operator InSubtree.
	// For Data Integration APIs, a filter based on type Task is used.
	Type []string `contributesTo:"query" name:"type" collectionFormat:"multi"`

	// Used in association with type parameter. If value is true,
	// then type all sub types of the given type parameter is considered.
	// If value is false, then sub types are not considered. Default is false.
	TypeInSubtree *string `mandatory:"false" contributesTo:"query" name:"typeInSubtree"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListDependentObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListDependentObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDependentObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDependentObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDependentObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDependentObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDependentObjectsResponse wrapper for the ListDependentObjects operation
type ListDependentObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DependentObjectSummaryCollection instances
	DependentObjectSummaryCollection `presentIn:"body"`

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

func (response ListDependentObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDependentObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDependentObjectsSortOrderEnum Enum with underlying type: string
type ListDependentObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListDependentObjectsSortOrderEnum
const (
	ListDependentObjectsSortOrderAsc  ListDependentObjectsSortOrderEnum = "ASC"
	ListDependentObjectsSortOrderDesc ListDependentObjectsSortOrderEnum = "DESC"
)

var mappingListDependentObjectsSortOrder = map[string]ListDependentObjectsSortOrderEnum{
	"ASC":  ListDependentObjectsSortOrderAsc,
	"DESC": ListDependentObjectsSortOrderDesc,
}

// GetListDependentObjectsSortOrderEnumValues Enumerates the set of values for ListDependentObjectsSortOrderEnum
func GetListDependentObjectsSortOrderEnumValues() []ListDependentObjectsSortOrderEnum {
	values := make([]ListDependentObjectsSortOrderEnum, 0)
	for _, v := range mappingListDependentObjectsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDependentObjectsSortByEnum Enum with underlying type: string
type ListDependentObjectsSortByEnum string

// Set of constants representing the allowable values for ListDependentObjectsSortByEnum
const (
	ListDependentObjectsSortByTimeCreated ListDependentObjectsSortByEnum = "TIME_CREATED"
	ListDependentObjectsSortByDisplayName ListDependentObjectsSortByEnum = "DISPLAY_NAME"
)

var mappingListDependentObjectsSortBy = map[string]ListDependentObjectsSortByEnum{
	"TIME_CREATED": ListDependentObjectsSortByTimeCreated,
	"DISPLAY_NAME": ListDependentObjectsSortByDisplayName,
}

// GetListDependentObjectsSortByEnumValues Enumerates the set of values for ListDependentObjectsSortByEnum
func GetListDependentObjectsSortByEnumValues() []ListDependentObjectsSortByEnum {
	values := make([]ListDependentObjectsSortByEnum, 0)
	for _, v := range mappingListDependentObjectsSortBy {
		values = append(values, v)
	}
	return values
}
