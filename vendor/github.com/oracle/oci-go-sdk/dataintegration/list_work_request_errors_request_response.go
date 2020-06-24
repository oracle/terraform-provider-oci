// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListWorkRequestErrorsRequest wrapper for the ListWorkRequestErrors operation
type ListWorkRequestErrorsRequest struct {

	// The ID of the asynchronous request.
	WorkRequestId *string `mandatory:"true" contributesTo:"path" name:"workRequestId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This parameter will control pagination.  Values for the parameter should come from the `opc-next-page` or `opc-prev-page` header in previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This parameter allows users to set the maximum number of items to return per page.  The value must be between 1 and 100 (inclusive).  Default value is 100.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// This parameter is used to control the sort order.  Supported values are `ASC` (ascending) and `DESC` (descending).
	SortOrder ListWorkRequestErrorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This parameter allows users to specify a sort field.  Supported sort fields are `name`, `identifier`, `timeCreated`, and `timeUpdated`.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListWorkRequestErrorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestErrorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestErrorsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestErrorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkRequestErrorsResponse wrapper for the ListWorkRequestErrors operation
type ListWorkRequestErrorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkRequestError instances
	Items []WorkRequestError `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestErrorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestErrorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestErrorsSortOrderEnum Enum with underlying type: string
type ListWorkRequestErrorsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestErrorsSortOrderEnum
const (
	ListWorkRequestErrorsSortOrderAsc  ListWorkRequestErrorsSortOrderEnum = "ASC"
	ListWorkRequestErrorsSortOrderDesc ListWorkRequestErrorsSortOrderEnum = "DESC"
)

var mappingListWorkRequestErrorsSortOrder = map[string]ListWorkRequestErrorsSortOrderEnum{
	"ASC":  ListWorkRequestErrorsSortOrderAsc,
	"DESC": ListWorkRequestErrorsSortOrderDesc,
}

// GetListWorkRequestErrorsSortOrderEnumValues Enumerates the set of values for ListWorkRequestErrorsSortOrderEnum
func GetListWorkRequestErrorsSortOrderEnumValues() []ListWorkRequestErrorsSortOrderEnum {
	values := make([]ListWorkRequestErrorsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestErrorsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestErrorsSortByEnum Enum with underlying type: string
type ListWorkRequestErrorsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestErrorsSortByEnum
const (
	ListWorkRequestErrorsSortByTimeCreated ListWorkRequestErrorsSortByEnum = "TIME_CREATED"
	ListWorkRequestErrorsSortByDisplayName ListWorkRequestErrorsSortByEnum = "DISPLAY_NAME"
)

var mappingListWorkRequestErrorsSortBy = map[string]ListWorkRequestErrorsSortByEnum{
	"TIME_CREATED": ListWorkRequestErrorsSortByTimeCreated,
	"DISPLAY_NAME": ListWorkRequestErrorsSortByDisplayName,
}

// GetListWorkRequestErrorsSortByEnumValues Enumerates the set of values for ListWorkRequestErrorsSortByEnum
func GetListWorkRequestErrorsSortByEnumValues() []ListWorkRequestErrorsSortByEnum {
	values := make([]ListWorkRequestErrorsSortByEnum, 0)
	for _, v := range mappingListWorkRequestErrorsSortBy {
		values = append(values, v)
	}
	return values
}
