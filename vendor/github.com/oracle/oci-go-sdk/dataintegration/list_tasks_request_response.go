// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTasksRequest wrapper for the ListTasks operation
type ListTasksRequest struct {

	// DIS workspace id
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Unique key of the folder
	FolderId *string `mandatory:"false" contributesTo:"query" name:"folderId"`

	// This parameter allows users to specify which fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// This filter parameter can be used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// This filter parameter can be used to filter by the key of the object.
	Key []string `contributesTo:"query" name:"key" collectionFormat:"multi"`

	// This filter parameter can be used to filter by the identifier of the object.
	Identifier []string `contributesTo:"query" name:"identifier" collectionFormat:"multi"`

	// This filter parameter can be used to filter by the object type of the object. This parameter can be suffixed with an optional filter operator InSubtree. If this operator is not specified, then exact match is considered. <br><br><B>Examples:-</B><br> <ul> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=false</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK</B> returns all objects of type data loader task</li> <li><B>?type=DATA_LOADER_TASK&typeInSubtree=true</B> returns all objects of type data loader task</li> </ul>
	Type []string `contributesTo:"query" name:"type" collectionFormat:"multi"`

	// This parameter allows users to set the maximum number of items to return per page.  The value must be between 1 and 100 (inclusive).  Default value is 100.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// This parameter will control pagination.  Values for the parameter should come from the `opc-next-page` or `opc-prev-page` header in previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This parameter is used to control the sort order.  Supported values are `ASC` (ascending) and `DESC` (descending).
	SortOrder ListTasksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This parameter allows users to specify a sort field.  Supported sort fields are `name`, `identifier`, `timeCreated`, and `timeUpdated`.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListTasksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTasksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTasksRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTasksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTasksResponse wrapper for the ListTasks operation
type ListTasksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TaskSummaryCollection instances
	TaskSummaryCollection `presentIn:"body"`

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

func (response ListTasksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTasksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTasksSortOrderEnum Enum with underlying type: string
type ListTasksSortOrderEnum string

// Set of constants representing the allowable values for ListTasksSortOrderEnum
const (
	ListTasksSortOrderAsc  ListTasksSortOrderEnum = "ASC"
	ListTasksSortOrderDesc ListTasksSortOrderEnum = "DESC"
)

var mappingListTasksSortOrder = map[string]ListTasksSortOrderEnum{
	"ASC":  ListTasksSortOrderAsc,
	"DESC": ListTasksSortOrderDesc,
}

// GetListTasksSortOrderEnumValues Enumerates the set of values for ListTasksSortOrderEnum
func GetListTasksSortOrderEnumValues() []ListTasksSortOrderEnum {
	values := make([]ListTasksSortOrderEnum, 0)
	for _, v := range mappingListTasksSortOrder {
		values = append(values, v)
	}
	return values
}

// ListTasksSortByEnum Enum with underlying type: string
type ListTasksSortByEnum string

// Set of constants representing the allowable values for ListTasksSortByEnum
const (
	ListTasksSortByTimeCreated ListTasksSortByEnum = "TIME_CREATED"
	ListTasksSortByDisplayName ListTasksSortByEnum = "DISPLAY_NAME"
)

var mappingListTasksSortBy = map[string]ListTasksSortByEnum{
	"TIME_CREATED": ListTasksSortByTimeCreated,
	"DISPLAY_NAME": ListTasksSortByDisplayName,
}

// GetListTasksSortByEnumValues Enumerates the set of values for ListTasksSortByEnum
func GetListTasksSortByEnumValues() []ListTasksSortByEnum {
	values := make([]ListTasksSortByEnum, 0)
	for _, v := range mappingListTasksSortBy {
		values = append(values, v)
	}
	return values
}
