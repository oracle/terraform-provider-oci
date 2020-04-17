// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
type ListWorkRequestsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Work Request status.
	WorkRequestStatus ListWorkRequestsWorkRequestStatusEnum `mandatory:"false" contributesTo:"query" name:"workRequestStatus" omitEmpty:"true"`

	// This parameter will control pagination.  Values for the parameter should come from the `opc-next-page` or `opc-prev-page` header in previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// This parameter allows users to set the maximum number of items to return per page.  The value must be between 1 and 100 (inclusive).  Default value is 100.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// This parameter is used to control the sort order.  Supported values are `ASC` (ascending) and `DESC` (descending).
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// This parameter allows users to specify a sort field.  Supported sort fields are `name`, `identifier`, `timeCreated`, and `timeUpdated`.  Default sort order is the descending order of `timeCreated` (most recently created objects at the top).  Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWorkRequestsResponse wrapper for the ListWorkRequests operation
type ListWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WorkRequestSummary instances
	Items []WorkRequestSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsWorkRequestStatusEnum Enum with underlying type: string
type ListWorkRequestsWorkRequestStatusEnum string

// Set of constants representing the allowable values for ListWorkRequestsWorkRequestStatusEnum
const (
	ListWorkRequestsWorkRequestStatusAccepted   ListWorkRequestsWorkRequestStatusEnum = "ACCEPTED"
	ListWorkRequestsWorkRequestStatusInProgress ListWorkRequestsWorkRequestStatusEnum = "IN_PROGRESS"
	ListWorkRequestsWorkRequestStatusFailed     ListWorkRequestsWorkRequestStatusEnum = "FAILED"
	ListWorkRequestsWorkRequestStatusSucceeded  ListWorkRequestsWorkRequestStatusEnum = "SUCCEEDED"
	ListWorkRequestsWorkRequestStatusCanceling  ListWorkRequestsWorkRequestStatusEnum = "CANCELING"
	ListWorkRequestsWorkRequestStatusCanceled   ListWorkRequestsWorkRequestStatusEnum = "CANCELED"
)

var mappingListWorkRequestsWorkRequestStatus = map[string]ListWorkRequestsWorkRequestStatusEnum{
	"ACCEPTED":    ListWorkRequestsWorkRequestStatusAccepted,
	"IN_PROGRESS": ListWorkRequestsWorkRequestStatusInProgress,
	"FAILED":      ListWorkRequestsWorkRequestStatusFailed,
	"SUCCEEDED":   ListWorkRequestsWorkRequestStatusSucceeded,
	"CANCELING":   ListWorkRequestsWorkRequestStatusCanceling,
	"CANCELED":    ListWorkRequestsWorkRequestStatusCanceled,
}

// GetListWorkRequestsWorkRequestStatusEnumValues Enumerates the set of values for ListWorkRequestsWorkRequestStatusEnum
func GetListWorkRequestsWorkRequestStatusEnumValues() []ListWorkRequestsWorkRequestStatusEnum {
	values := make([]ListWorkRequestsWorkRequestStatusEnum, 0)
	for _, v := range mappingListWorkRequestsWorkRequestStatus {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestsSortOrderEnum Enum with underlying type: string
type ListWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortOrderEnum
const (
	ListWorkRequestsSortOrderAsc  ListWorkRequestsSortOrderEnum = "ASC"
	ListWorkRequestsSortOrderDesc ListWorkRequestsSortOrderEnum = "DESC"
)

var mappingListWorkRequestsSortOrder = map[string]ListWorkRequestsSortOrderEnum{
	"ASC":  ListWorkRequestsSortOrderAsc,
	"DESC": ListWorkRequestsSortOrderDesc,
}

// GetListWorkRequestsSortOrderEnumValues Enumerates the set of values for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumValues() []ListWorkRequestsSortOrderEnum {
	values := make([]ListWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWorkRequestsSortByEnum Enum with underlying type: string
type ListWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortByEnum
const (
	ListWorkRequestsSortByTimeCreated ListWorkRequestsSortByEnum = "TIME_CREATED"
	ListWorkRequestsSortByDisplayName ListWorkRequestsSortByEnum = "DISPLAY_NAME"
)

var mappingListWorkRequestsSortBy = map[string]ListWorkRequestsSortByEnum{
	"TIME_CREATED": ListWorkRequestsSortByTimeCreated,
	"DISPLAY_NAME": ListWorkRequestsSortByDisplayName,
}

// GetListWorkRequestsSortByEnumValues Enumerates the set of values for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumValues() []ListWorkRequestsSortByEnum {
	values := make([]ListWorkRequestsSortByEnum, 0)
	for _, v := range mappingListWorkRequestsSortBy {
		values = append(values, v)
	}
	return values
}
