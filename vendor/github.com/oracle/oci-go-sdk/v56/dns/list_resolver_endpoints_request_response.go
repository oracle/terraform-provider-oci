// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListResolverEndpointsRequest wrapper for the ListResolverEndpoints operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListResolverEndpoints.go.html to see an example of how to use ListResolverEndpointsRequest.
type ListResolverEndpointsRequest struct {

	// The OCID of the target resolver.
	ResolverId *string `mandatory:"true" contributesTo:"path" name:"resolverId"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The name of a resource.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a page of the collection.
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// The order to sort the resources.
	SortOrder ListResolverEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort resolver endpoints.
	SortBy ListResolverEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The state of a resource.
	LifecycleState ResolverEndpointSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListResolverEndpointsScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResolverEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResolverEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResolverEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResolverEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListResolverEndpointsResponse wrapper for the ListResolverEndpoints operation
type ListResolverEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ResolverEndpointSummary instances
	Items []ResolverEndpointSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResolverEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResolverEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResolverEndpointsSortOrderEnum Enum with underlying type: string
type ListResolverEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListResolverEndpointsSortOrderEnum
const (
	ListResolverEndpointsSortOrderAsc  ListResolverEndpointsSortOrderEnum = "ASC"
	ListResolverEndpointsSortOrderDesc ListResolverEndpointsSortOrderEnum = "DESC"
)

var mappingListResolverEndpointsSortOrder = map[string]ListResolverEndpointsSortOrderEnum{
	"ASC":  ListResolverEndpointsSortOrderAsc,
	"DESC": ListResolverEndpointsSortOrderDesc,
}

// GetListResolverEndpointsSortOrderEnumValues Enumerates the set of values for ListResolverEndpointsSortOrderEnum
func GetListResolverEndpointsSortOrderEnumValues() []ListResolverEndpointsSortOrderEnum {
	values := make([]ListResolverEndpointsSortOrderEnum, 0)
	for _, v := range mappingListResolverEndpointsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListResolverEndpointsSortByEnum Enum with underlying type: string
type ListResolverEndpointsSortByEnum string

// Set of constants representing the allowable values for ListResolverEndpointsSortByEnum
const (
	ListResolverEndpointsSortByName        ListResolverEndpointsSortByEnum = "name"
	ListResolverEndpointsSortByTimecreated ListResolverEndpointsSortByEnum = "timeCreated"
)

var mappingListResolverEndpointsSortBy = map[string]ListResolverEndpointsSortByEnum{
	"name":        ListResolverEndpointsSortByName,
	"timeCreated": ListResolverEndpointsSortByTimecreated,
}

// GetListResolverEndpointsSortByEnumValues Enumerates the set of values for ListResolverEndpointsSortByEnum
func GetListResolverEndpointsSortByEnumValues() []ListResolverEndpointsSortByEnum {
	values := make([]ListResolverEndpointsSortByEnum, 0)
	for _, v := range mappingListResolverEndpointsSortBy {
		values = append(values, v)
	}
	return values
}

// ListResolverEndpointsScopeEnum Enum with underlying type: string
type ListResolverEndpointsScopeEnum string

// Set of constants representing the allowable values for ListResolverEndpointsScopeEnum
const (
	ListResolverEndpointsScopeGlobal  ListResolverEndpointsScopeEnum = "GLOBAL"
	ListResolverEndpointsScopePrivate ListResolverEndpointsScopeEnum = "PRIVATE"
)

var mappingListResolverEndpointsScope = map[string]ListResolverEndpointsScopeEnum{
	"GLOBAL":  ListResolverEndpointsScopeGlobal,
	"PRIVATE": ListResolverEndpointsScopePrivate,
}

// GetListResolverEndpointsScopeEnumValues Enumerates the set of values for ListResolverEndpointsScopeEnum
func GetListResolverEndpointsScopeEnumValues() []ListResolverEndpointsScopeEnum {
	values := make([]ListResolverEndpointsScopeEnum, 0)
	for _, v := range mappingListResolverEndpointsScope {
		values = append(values, v)
	}
	return values
}
