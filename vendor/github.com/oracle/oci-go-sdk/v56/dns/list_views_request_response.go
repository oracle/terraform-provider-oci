// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListViewsRequest wrapper for the ListViews operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListViews.go.html to see an example of how to use ListViewsRequest.
type ListViewsRequest struct {

	// The OCID of the compartment the resource belongs to.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The displayName of a resource.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The OCID of a resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a page of the collection.
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The order to sort the resources.
	SortOrder ListViewsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort views.
	SortBy ListViewsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The state of a resource.
	LifecycleState ViewSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListViewsScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListViewsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListViewsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListViewsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListViewsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListViewsResponse wrapper for the ListViews operation
type ListViewsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ViewSummary instances
	Items []ViewSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListViewsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListViewsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListViewsSortOrderEnum Enum with underlying type: string
type ListViewsSortOrderEnum string

// Set of constants representing the allowable values for ListViewsSortOrderEnum
const (
	ListViewsSortOrderAsc  ListViewsSortOrderEnum = "ASC"
	ListViewsSortOrderDesc ListViewsSortOrderEnum = "DESC"
)

var mappingListViewsSortOrder = map[string]ListViewsSortOrderEnum{
	"ASC":  ListViewsSortOrderAsc,
	"DESC": ListViewsSortOrderDesc,
}

// GetListViewsSortOrderEnumValues Enumerates the set of values for ListViewsSortOrderEnum
func GetListViewsSortOrderEnumValues() []ListViewsSortOrderEnum {
	values := make([]ListViewsSortOrderEnum, 0)
	for _, v := range mappingListViewsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListViewsSortByEnum Enum with underlying type: string
type ListViewsSortByEnum string

// Set of constants representing the allowable values for ListViewsSortByEnum
const (
	ListViewsSortByDisplayname ListViewsSortByEnum = "displayName"
	ListViewsSortByTimecreated ListViewsSortByEnum = "timeCreated"
)

var mappingListViewsSortBy = map[string]ListViewsSortByEnum{
	"displayName": ListViewsSortByDisplayname,
	"timeCreated": ListViewsSortByTimecreated,
}

// GetListViewsSortByEnumValues Enumerates the set of values for ListViewsSortByEnum
func GetListViewsSortByEnumValues() []ListViewsSortByEnum {
	values := make([]ListViewsSortByEnum, 0)
	for _, v := range mappingListViewsSortBy {
		values = append(values, v)
	}
	return values
}

// ListViewsScopeEnum Enum with underlying type: string
type ListViewsScopeEnum string

// Set of constants representing the allowable values for ListViewsScopeEnum
const (
	ListViewsScopeGlobal  ListViewsScopeEnum = "GLOBAL"
	ListViewsScopePrivate ListViewsScopeEnum = "PRIVATE"
)

var mappingListViewsScope = map[string]ListViewsScopeEnum{
	"GLOBAL":  ListViewsScopeGlobal,
	"PRIVATE": ListViewsScopePrivate,
}

// GetListViewsScopeEnumValues Enumerates the set of values for ListViewsScopeEnum
func GetListViewsScopeEnumValues() []ListViewsScopeEnum {
	values := make([]ListViewsScopeEnum, 0)
	for _, v := range mappingListViewsScope {
		values = append(values, v)
	}
	return values
}
