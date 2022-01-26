// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListResolversRequest wrapper for the ListResolvers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/ListResolvers.go.html to see an example of how to use ListResolversRequest.
type ListResolversRequest struct {

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
	SortOrder ListResolversSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field by which to sort resolvers.
	SortBy ListResolversSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The state of a resource.
	LifecycleState ResolverSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope ListResolversScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResolversRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResolversRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResolversRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResolversRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListResolversResponse wrapper for the ListResolvers operation
type ListResolversResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ResolverSummary instances
	Items []ResolverSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResolversResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResolversResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResolversSortOrderEnum Enum with underlying type: string
type ListResolversSortOrderEnum string

// Set of constants representing the allowable values for ListResolversSortOrderEnum
const (
	ListResolversSortOrderAsc  ListResolversSortOrderEnum = "ASC"
	ListResolversSortOrderDesc ListResolversSortOrderEnum = "DESC"
)

var mappingListResolversSortOrder = map[string]ListResolversSortOrderEnum{
	"ASC":  ListResolversSortOrderAsc,
	"DESC": ListResolversSortOrderDesc,
}

// GetListResolversSortOrderEnumValues Enumerates the set of values for ListResolversSortOrderEnum
func GetListResolversSortOrderEnumValues() []ListResolversSortOrderEnum {
	values := make([]ListResolversSortOrderEnum, 0)
	for _, v := range mappingListResolversSortOrder {
		values = append(values, v)
	}
	return values
}

// ListResolversSortByEnum Enum with underlying type: string
type ListResolversSortByEnum string

// Set of constants representing the allowable values for ListResolversSortByEnum
const (
	ListResolversSortByDisplayname ListResolversSortByEnum = "displayName"
	ListResolversSortByTimecreated ListResolversSortByEnum = "timeCreated"
)

var mappingListResolversSortBy = map[string]ListResolversSortByEnum{
	"displayName": ListResolversSortByDisplayname,
	"timeCreated": ListResolversSortByTimecreated,
}

// GetListResolversSortByEnumValues Enumerates the set of values for ListResolversSortByEnum
func GetListResolversSortByEnumValues() []ListResolversSortByEnum {
	values := make([]ListResolversSortByEnum, 0)
	for _, v := range mappingListResolversSortBy {
		values = append(values, v)
	}
	return values
}

// ListResolversScopeEnum Enum with underlying type: string
type ListResolversScopeEnum string

// Set of constants representing the allowable values for ListResolversScopeEnum
const (
	ListResolversScopeGlobal  ListResolversScopeEnum = "GLOBAL"
	ListResolversScopePrivate ListResolversScopeEnum = "PRIVATE"
)

var mappingListResolversScope = map[string]ListResolversScopeEnum{
	"GLOBAL":  ListResolversScopeGlobal,
	"PRIVATE": ListResolversScopePrivate,
}

// GetListResolversScopeEnumValues Enumerates the set of values for ListResolversScopeEnum
func GetListResolversScopeEnumValues() []ListResolversScopeEnum {
	values := make([]ListResolversScopeEnum, 0)
	for _, v := range mappingListResolversScope {
		values = append(values, v)
	}
	return values
}
