// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListLogsRequest wrapper for the ListLogs operation
type ListLogsRequest struct {

	// OCID of a log group to work with.
	LogGroupId *string `mandatory:"true" contributesTo:"path" name:"logGroupId"`

	// The logType that the log object is for, custom or service.
	LogType ListLogsLogTypeEnum `mandatory:"false" contributesTo:"query" name:"logType" omitEmpty:"true"`

	// Service created the log object
	SourceService *string `mandatory:"false" contributesTo:"query" name:"sourceService"`

	// Log object resource
	SourceResource *string `mandatory:"false" contributesTo:"query" name:"sourceResource"`

	// Resource name
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Lifecycle state of the log object
	LifecycleState ListLogsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'
	SortOrder ListLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogsResponse wrapper for the ListLogs operation
type ListLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LogSummary instances
	Items []LogSummary `presentIn:"body"`

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

func (response ListLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogsLogTypeEnum Enum with underlying type: string
type ListLogsLogTypeEnum string

// Set of constants representing the allowable values for ListLogsLogTypeEnum
const (
	ListLogsLogTypeCustom  ListLogsLogTypeEnum = "CUSTOM"
	ListLogsLogTypeService ListLogsLogTypeEnum = "SERVICE"
)

var mappingListLogsLogType = map[string]ListLogsLogTypeEnum{
	"CUSTOM":  ListLogsLogTypeCustom,
	"SERVICE": ListLogsLogTypeService,
}

// GetListLogsLogTypeEnumValues Enumerates the set of values for ListLogsLogTypeEnum
func GetListLogsLogTypeEnumValues() []ListLogsLogTypeEnum {
	values := make([]ListLogsLogTypeEnum, 0)
	for _, v := range mappingListLogsLogType {
		values = append(values, v)
	}
	return values
}

// ListLogsLifecycleStateEnum Enum with underlying type: string
type ListLogsLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogsLifecycleStateEnum
const (
	ListLogsLifecycleStateCreating ListLogsLifecycleStateEnum = "CREATING"
	ListLogsLifecycleStateActive   ListLogsLifecycleStateEnum = "ACTIVE"
	ListLogsLifecycleStateUpdating ListLogsLifecycleStateEnum = "UPDATING"
	ListLogsLifecycleStateInactive ListLogsLifecycleStateEnum = "INACTIVE"
	ListLogsLifecycleStateDeleting ListLogsLifecycleStateEnum = "DELETING"
	ListLogsLifecycleStateFailed   ListLogsLifecycleStateEnum = "FAILED"
)

var mappingListLogsLifecycleState = map[string]ListLogsLifecycleStateEnum{
	"CREATING": ListLogsLifecycleStateCreating,
	"ACTIVE":   ListLogsLifecycleStateActive,
	"UPDATING": ListLogsLifecycleStateUpdating,
	"INACTIVE": ListLogsLifecycleStateInactive,
	"DELETING": ListLogsLifecycleStateDeleting,
	"FAILED":   ListLogsLifecycleStateFailed,
}

// GetListLogsLifecycleStateEnumValues Enumerates the set of values for ListLogsLifecycleStateEnum
func GetListLogsLifecycleStateEnumValues() []ListLogsLifecycleStateEnum {
	values := make([]ListLogsLifecycleStateEnum, 0)
	for _, v := range mappingListLogsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListLogsSortByEnum Enum with underlying type: string
type ListLogsSortByEnum string

// Set of constants representing the allowable values for ListLogsSortByEnum
const (
	ListLogsSortByTimecreated ListLogsSortByEnum = "timeCreated"
	ListLogsSortByDisplayname ListLogsSortByEnum = "displayName"
)

var mappingListLogsSortBy = map[string]ListLogsSortByEnum{
	"timeCreated": ListLogsSortByTimecreated,
	"displayName": ListLogsSortByDisplayname,
}

// GetListLogsSortByEnumValues Enumerates the set of values for ListLogsSortByEnum
func GetListLogsSortByEnumValues() []ListLogsSortByEnum {
	values := make([]ListLogsSortByEnum, 0)
	for _, v := range mappingListLogsSortBy {
		values = append(values, v)
	}
	return values
}

// ListLogsSortOrderEnum Enum with underlying type: string
type ListLogsSortOrderEnum string

// Set of constants representing the allowable values for ListLogsSortOrderEnum
const (
	ListLogsSortOrderAsc  ListLogsSortOrderEnum = "ASC"
	ListLogsSortOrderDesc ListLogsSortOrderEnum = "DESC"
)

var mappingListLogsSortOrder = map[string]ListLogsSortOrderEnum{
	"ASC":  ListLogsSortOrderAsc,
	"DESC": ListLogsSortOrderDesc,
}

// GetListLogsSortOrderEnumValues Enumerates the set of values for ListLogsSortOrderEnum
func GetListLogsSortOrderEnumValues() []ListLogsSortOrderEnum {
	values := make([]ListLogsSortOrderEnum, 0)
	for _, v := range mappingListLogsSortOrder {
		values = append(values, v)
	}
	return values
}
