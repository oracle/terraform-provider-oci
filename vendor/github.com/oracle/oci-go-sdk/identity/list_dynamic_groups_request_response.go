// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identity

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDynamicGroupsRequest wrapper for the ListDynamicGroups operation
type ListDynamicGroupsRequest struct {

	// The OCID of the compartment (remember that the tenancy is simply the root compartment).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A filter to only return resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for NAME is ascending. The NAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListDynamicGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The NAME sort order
	// is case sensitive.
	SortOrder ListDynamicGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState DynamicGroupLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDynamicGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDynamicGroupsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDynamicGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDynamicGroupsResponse wrapper for the ListDynamicGroups operation
type ListDynamicGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DynamicGroup instances
	Items []DynamicGroup `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDynamicGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDynamicGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDynamicGroupsSortByEnum Enum with underlying type: string
type ListDynamicGroupsSortByEnum string

// Set of constants representing the allowable values for ListDynamicGroupsSortByEnum
const (
	ListDynamicGroupsSortByTimecreated ListDynamicGroupsSortByEnum = "TIMECREATED"
	ListDynamicGroupsSortByName        ListDynamicGroupsSortByEnum = "NAME"
)

var mappingListDynamicGroupsSortBy = map[string]ListDynamicGroupsSortByEnum{
	"TIMECREATED": ListDynamicGroupsSortByTimecreated,
	"NAME":        ListDynamicGroupsSortByName,
}

// GetListDynamicGroupsSortByEnumValues Enumerates the set of values for ListDynamicGroupsSortByEnum
func GetListDynamicGroupsSortByEnumValues() []ListDynamicGroupsSortByEnum {
	values := make([]ListDynamicGroupsSortByEnum, 0)
	for _, v := range mappingListDynamicGroupsSortBy {
		values = append(values, v)
	}
	return values
}

// ListDynamicGroupsSortOrderEnum Enum with underlying type: string
type ListDynamicGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListDynamicGroupsSortOrderEnum
const (
	ListDynamicGroupsSortOrderAsc  ListDynamicGroupsSortOrderEnum = "ASC"
	ListDynamicGroupsSortOrderDesc ListDynamicGroupsSortOrderEnum = "DESC"
)

var mappingListDynamicGroupsSortOrder = map[string]ListDynamicGroupsSortOrderEnum{
	"ASC":  ListDynamicGroupsSortOrderAsc,
	"DESC": ListDynamicGroupsSortOrderDesc,
}

// GetListDynamicGroupsSortOrderEnumValues Enumerates the set of values for ListDynamicGroupsSortOrderEnum
func GetListDynamicGroupsSortOrderEnumValues() []ListDynamicGroupsSortOrderEnum {
	values := make([]ListDynamicGroupsSortOrderEnum, 0)
	for _, v := range mappingListDynamicGroupsSortOrder {
		values = append(values, v)
	}
	return values
}
