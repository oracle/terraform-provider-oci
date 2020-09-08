// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLogGroupsRequest wrapper for the ListLogGroups operation
type ListLogGroupsRequest struct {

	// Compartment OCID to list resources in. Please see compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Specifies whether or not nested compartments should be traversed. Defaults to false.
	IsCompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"isCompartmentIdInSubtree"`

	// Resource name
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogGroupsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListLogGroupsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogGroupsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogGroupsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogGroupsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogGroupsResponse wrapper for the ListLogGroups operation
type ListLogGroupsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LogGroupSummary instances
	Items []LogGroupSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLogGroupsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogGroupsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogGroupsSortByEnum Enum with underlying type: string
type ListLogGroupsSortByEnum string

// Set of constants representing the allowable values for ListLogGroupsSortByEnum
const (
	ListLogGroupsSortByTimecreated ListLogGroupsSortByEnum = "timeCreated"
	ListLogGroupsSortByDisplayname ListLogGroupsSortByEnum = "displayName"
)

var mappingListLogGroupsSortBy = map[string]ListLogGroupsSortByEnum{
	"timeCreated": ListLogGroupsSortByTimecreated,
	"displayName": ListLogGroupsSortByDisplayname,
}

// GetListLogGroupsSortByEnumValues Enumerates the set of values for ListLogGroupsSortByEnum
func GetListLogGroupsSortByEnumValues() []ListLogGroupsSortByEnum {
	values := make([]ListLogGroupsSortByEnum, 0)
	for _, v := range mappingListLogGroupsSortBy {
		values = append(values, v)
	}
	return values
}

// ListLogGroupsSortOrderEnum Enum with underlying type: string
type ListLogGroupsSortOrderEnum string

// Set of constants representing the allowable values for ListLogGroupsSortOrderEnum
const (
	ListLogGroupsSortOrderAsc  ListLogGroupsSortOrderEnum = "ASC"
	ListLogGroupsSortOrderDesc ListLogGroupsSortOrderEnum = "DESC"
)

var mappingListLogGroupsSortOrder = map[string]ListLogGroupsSortOrderEnum{
	"ASC":  ListLogGroupsSortOrderAsc,
	"DESC": ListLogGroupsSortOrderDesc,
}

// GetListLogGroupsSortOrderEnumValues Enumerates the set of values for ListLogGroupsSortOrderEnum
func GetListLogGroupsSortOrderEnumValues() []ListLogGroupsSortOrderEnum {
	values := make([]ListLogGroupsSortOrderEnum, 0)
	for _, v := range mappingListLogGroupsSortOrder {
		values = append(values, v)
	}
	return values
}
