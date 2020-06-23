// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListTaskRunLogsRequest wrapper for the ListTaskRunLogs operation
type ListTaskRunLogsRequest struct {

	// DIS workspace id
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// DIS application key
	ApplicationKey *string `mandatory:"true" contributesTo:"path" name:"applicationKey"`

	// DIS taskRun key
	TaskRunKey *string `mandatory:"true" contributesTo:"path" name:"taskRunKey"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This parameter will control pagination.  Values for the parameter should come from the `opc-next-page` or `opc-prev-page` header in previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This parameter allows users to set the maximum number of items to return per page.  The value must be between 1 and 100 (inclusive).  Default value is 100.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// This parameter is used to control the sort order.  Supported values are `ASC` (ascending) and `DESC` (descending).
	SortOrder ListTaskRunLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This parameter allows users to specify a sort field.  Supported sort fields are `name`, `identifier`, `timeCreated`, and `timeUpdated`.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListTaskRunLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTaskRunLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTaskRunLogsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTaskRunLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTaskRunLogsResponse wrapper for the ListTaskRunLogs operation
type ListTaskRunLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []TaskRunLogSummary instances
	Items []TaskRunLogSummary `presentIn:"body"`

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

func (response ListTaskRunLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTaskRunLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTaskRunLogsSortOrderEnum Enum with underlying type: string
type ListTaskRunLogsSortOrderEnum string

// Set of constants representing the allowable values for ListTaskRunLogsSortOrderEnum
const (
	ListTaskRunLogsSortOrderAsc  ListTaskRunLogsSortOrderEnum = "ASC"
	ListTaskRunLogsSortOrderDesc ListTaskRunLogsSortOrderEnum = "DESC"
)

var mappingListTaskRunLogsSortOrder = map[string]ListTaskRunLogsSortOrderEnum{
	"ASC":  ListTaskRunLogsSortOrderAsc,
	"DESC": ListTaskRunLogsSortOrderDesc,
}

// GetListTaskRunLogsSortOrderEnumValues Enumerates the set of values for ListTaskRunLogsSortOrderEnum
func GetListTaskRunLogsSortOrderEnumValues() []ListTaskRunLogsSortOrderEnum {
	values := make([]ListTaskRunLogsSortOrderEnum, 0)
	for _, v := range mappingListTaskRunLogsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListTaskRunLogsSortByEnum Enum with underlying type: string
type ListTaskRunLogsSortByEnum string

// Set of constants representing the allowable values for ListTaskRunLogsSortByEnum
const (
	ListTaskRunLogsSortByTimeCreated ListTaskRunLogsSortByEnum = "TIME_CREATED"
	ListTaskRunLogsSortByDisplayName ListTaskRunLogsSortByEnum = "DISPLAY_NAME"
)

var mappingListTaskRunLogsSortBy = map[string]ListTaskRunLogsSortByEnum{
	"TIME_CREATED": ListTaskRunLogsSortByTimeCreated,
	"DISPLAY_NAME": ListTaskRunLogsSortByDisplayName,
}

// GetListTaskRunLogsSortByEnumValues Enumerates the set of values for ListTaskRunLogsSortByEnum
func GetListTaskRunLogsSortByEnumValues() []ListTaskRunLogsSortByEnum {
	values := make([]ListTaskRunLogsSortByEnum, 0)
	for _, v := range mappingListTaskRunLogsSortBy {
		values = append(values, v)
	}
	return values
}
