// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListPublishedObjectsRequest wrapper for the ListPublishedObjects operation
type ListPublishedObjectsRequest struct {

	// DIS workspace id
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// DIS application key
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// This parameter allows users to specify which fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// This filter parameter can be used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This filter parameter can be used to filter by the identifier of the published object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// This filter parameter can be used to filter by the object type of the object.
	// This parameter can be suffixed with an optional filter operator InSubtree.
	// For DIS APIs we will filter based on type Task.
	Type []string `contributesTo:"query" name:"type" collectionFormat:"multi"`

	// This is used in association with type parameter. If value is true,
	// then type all sub types of the given type parameter is considered.
	// If value is false, then sub types are not considered. Default is false.
	TypeInSubtree *string `mandatory:"false" contributesTo:"query" name:"typeInSubtree"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This parameter is used to control the sort order.  Supported values are `ASC` (ascending) and `DESC` (descending).
	SortOrder ListPublishedObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This parameter allows users to specify a sort field.  Supported sort fields are `name`, `identifier`, `timeCreated`, and `timeUpdated`.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListPublishedObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPublishedObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPublishedObjectsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPublishedObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPublishedObjectsResponse wrapper for the ListPublishedObjects operation
type ListPublishedObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PublishedObjectSummaryCollection instances
	PublishedObjectSummaryCollection `presentIn:"body"`

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

func (response ListPublishedObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPublishedObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPublishedObjectsSortOrderEnum Enum with underlying type: string
type ListPublishedObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListPublishedObjectsSortOrderEnum
const (
	ListPublishedObjectsSortOrderAsc  ListPublishedObjectsSortOrderEnum = "ASC"
	ListPublishedObjectsSortOrderDesc ListPublishedObjectsSortOrderEnum = "DESC"
)

var mappingListPublishedObjectsSortOrder = map[string]ListPublishedObjectsSortOrderEnum{
	"ASC":  ListPublishedObjectsSortOrderAsc,
	"DESC": ListPublishedObjectsSortOrderDesc,
}

// GetListPublishedObjectsSortOrderEnumValues Enumerates the set of values for ListPublishedObjectsSortOrderEnum
func GetListPublishedObjectsSortOrderEnumValues() []ListPublishedObjectsSortOrderEnum {
	values := make([]ListPublishedObjectsSortOrderEnum, 0)
	for _, v := range mappingListPublishedObjectsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListPublishedObjectsSortByEnum Enum with underlying type: string
type ListPublishedObjectsSortByEnum string

// Set of constants representing the allowable values for ListPublishedObjectsSortByEnum
const (
	ListPublishedObjectsSortByTimeCreated ListPublishedObjectsSortByEnum = "TIME_CREATED"
	ListPublishedObjectsSortByDisplayName ListPublishedObjectsSortByEnum = "DISPLAY_NAME"
)

var mappingListPublishedObjectsSortBy = map[string]ListPublishedObjectsSortByEnum{
	"TIME_CREATED": ListPublishedObjectsSortByTimeCreated,
	"DISPLAY_NAME": ListPublishedObjectsSortByDisplayName,
}

// GetListPublishedObjectsSortByEnumValues Enumerates the set of values for ListPublishedObjectsSortByEnum
func GetListPublishedObjectsSortByEnumValues() []ListPublishedObjectsSortByEnum {
	values := make([]ListPublishedObjectsSortByEnum, 0)
	for _, v := range mappingListPublishedObjectsSortBy {
		values = append(values, v)
	}
	return values
}
