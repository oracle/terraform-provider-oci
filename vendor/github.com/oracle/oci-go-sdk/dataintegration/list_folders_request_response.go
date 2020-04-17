// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListFoldersRequest wrapper for the ListFolders operation
type ListFoldersRequest struct {

	// DIS workspace id
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This filter parameter can be used to filter by the project or the folder object.
	AggregatorKey *string `mandatory:"false" contributesTo:"query" name:"aggregatorKey"`

	// This parameter allows users to specify which fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// This filter parameter can be used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This filter parameter can be used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// This parameter will control pagination.  Values for the parameter should come from the `opc-next-page` or `opc-prev-page` header in previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This parameter allows users to set the maximum number of items to return per page.  The value must be between 1 and 100 (inclusive).  Default value is 100.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// This parameter is used to control the sort order.  Supported values are `ASC` (ascending) and `DESC` (descending).
	SortOrder ListFoldersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This parameter allows users to specify a sort field.  Supported sort fields are `name`, `identifier`, `timeCreated`, and `timeUpdated`.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListFoldersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFoldersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFoldersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFoldersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFoldersResponse wrapper for the ListFolders operation
type ListFoldersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FolderSummaryCollection instances
	FolderSummaryCollection `presentIn:"body"`

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

func (response ListFoldersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFoldersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFoldersSortOrderEnum Enum with underlying type: string
type ListFoldersSortOrderEnum string

// Set of constants representing the allowable values for ListFoldersSortOrderEnum
const (
	ListFoldersSortOrderAsc  ListFoldersSortOrderEnum = "ASC"
	ListFoldersSortOrderDesc ListFoldersSortOrderEnum = "DESC"
)

var mappingListFoldersSortOrder = map[string]ListFoldersSortOrderEnum{
	"ASC":  ListFoldersSortOrderAsc,
	"DESC": ListFoldersSortOrderDesc,
}

// GetListFoldersSortOrderEnumValues Enumerates the set of values for ListFoldersSortOrderEnum
func GetListFoldersSortOrderEnumValues() []ListFoldersSortOrderEnum {
	values := make([]ListFoldersSortOrderEnum, 0)
	for _, v := range mappingListFoldersSortOrder {
		values = append(values, v)
	}
	return values
}

// ListFoldersSortByEnum Enum with underlying type: string
type ListFoldersSortByEnum string

// Set of constants representing the allowable values for ListFoldersSortByEnum
const (
	ListFoldersSortByTimeCreated ListFoldersSortByEnum = "TIME_CREATED"
	ListFoldersSortByDisplayName ListFoldersSortByEnum = "DISPLAY_NAME"
)

var mappingListFoldersSortBy = map[string]ListFoldersSortByEnum{
	"TIME_CREATED": ListFoldersSortByTimeCreated,
	"DISPLAY_NAME": ListFoldersSortByDisplayName,
}

// GetListFoldersSortByEnumValues Enumerates the set of values for ListFoldersSortByEnum
func GetListFoldersSortByEnumValues() []ListFoldersSortByEnum {
	values := make([]ListFoldersSortByEnum, 0)
	for _, v := range mappingListFoldersSortBy {
		values = append(values, v)
	}
	return values
}
