// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListJobLogsRequest wrapper for the ListJobLogs operation
type ListJobLogsRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique job key.
	JobKey *string `mandatory:"true" contributesTo:"path" name:"jobKey"`

	// The key of the job execution.
	JobExecutionKey *string `mandatory:"true" contributesTo:"path" name:"jobExecutionKey"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListJobLogsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Severity level for this Log.
	Severity *string `mandatory:"false" contributesTo:"query" name:"severity"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Specifies the fields to return in a job log summary response.
	Fields []ListJobLogsFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListJobLogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListJobLogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListJobLogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListJobLogsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListJobLogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListJobLogsResponse wrapper for the ListJobLogs operation
type ListJobLogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of JobLogCollection instances
	JobLogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListJobLogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListJobLogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListJobLogsLifecycleStateEnum Enum with underlying type: string
type ListJobLogsLifecycleStateEnum string

// Set of constants representing the allowable values for ListJobLogsLifecycleStateEnum
const (
	ListJobLogsLifecycleStateCreating ListJobLogsLifecycleStateEnum = "CREATING"
	ListJobLogsLifecycleStateActive   ListJobLogsLifecycleStateEnum = "ACTIVE"
	ListJobLogsLifecycleStateInactive ListJobLogsLifecycleStateEnum = "INACTIVE"
	ListJobLogsLifecycleStateUpdating ListJobLogsLifecycleStateEnum = "UPDATING"
	ListJobLogsLifecycleStateDeleting ListJobLogsLifecycleStateEnum = "DELETING"
	ListJobLogsLifecycleStateDeleted  ListJobLogsLifecycleStateEnum = "DELETED"
	ListJobLogsLifecycleStateFailed   ListJobLogsLifecycleStateEnum = "FAILED"
	ListJobLogsLifecycleStateMoving   ListJobLogsLifecycleStateEnum = "MOVING"
)

var mappingListJobLogsLifecycleState = map[string]ListJobLogsLifecycleStateEnum{
	"CREATING": ListJobLogsLifecycleStateCreating,
	"ACTIVE":   ListJobLogsLifecycleStateActive,
	"INACTIVE": ListJobLogsLifecycleStateInactive,
	"UPDATING": ListJobLogsLifecycleStateUpdating,
	"DELETING": ListJobLogsLifecycleStateDeleting,
	"DELETED":  ListJobLogsLifecycleStateDeleted,
	"FAILED":   ListJobLogsLifecycleStateFailed,
	"MOVING":   ListJobLogsLifecycleStateMoving,
}

// GetListJobLogsLifecycleStateEnumValues Enumerates the set of values for ListJobLogsLifecycleStateEnum
func GetListJobLogsLifecycleStateEnumValues() []ListJobLogsLifecycleStateEnum {
	values := make([]ListJobLogsLifecycleStateEnum, 0)
	for _, v := range mappingListJobLogsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListJobLogsFieldsEnum Enum with underlying type: string
type ListJobLogsFieldsEnum string

// Set of constants representing the allowable values for ListJobLogsFieldsEnum
const (
	ListJobLogsFieldsKey             ListJobLogsFieldsEnum = "key"
	ListJobLogsFieldsJobexecutionkey ListJobLogsFieldsEnum = "jobExecutionKey"
	ListJobLogsFieldsSeverity        ListJobLogsFieldsEnum = "severity"
	ListJobLogsFieldsTimecreated     ListJobLogsFieldsEnum = "timeCreated"
	ListJobLogsFieldsLogmessage      ListJobLogsFieldsEnum = "logMessage"
	ListJobLogsFieldsUri             ListJobLogsFieldsEnum = "uri"
)

var mappingListJobLogsFields = map[string]ListJobLogsFieldsEnum{
	"key":             ListJobLogsFieldsKey,
	"jobExecutionKey": ListJobLogsFieldsJobexecutionkey,
	"severity":        ListJobLogsFieldsSeverity,
	"timeCreated":     ListJobLogsFieldsTimecreated,
	"logMessage":      ListJobLogsFieldsLogmessage,
	"uri":             ListJobLogsFieldsUri,
}

// GetListJobLogsFieldsEnumValues Enumerates the set of values for ListJobLogsFieldsEnum
func GetListJobLogsFieldsEnumValues() []ListJobLogsFieldsEnum {
	values := make([]ListJobLogsFieldsEnum, 0)
	for _, v := range mappingListJobLogsFields {
		values = append(values, v)
	}
	return values
}

// ListJobLogsSortByEnum Enum with underlying type: string
type ListJobLogsSortByEnum string

// Set of constants representing the allowable values for ListJobLogsSortByEnum
const (
	ListJobLogsSortByTimecreated ListJobLogsSortByEnum = "TIMECREATED"
	ListJobLogsSortByDisplayname ListJobLogsSortByEnum = "DISPLAYNAME"
)

var mappingListJobLogsSortBy = map[string]ListJobLogsSortByEnum{
	"TIMECREATED": ListJobLogsSortByTimecreated,
	"DISPLAYNAME": ListJobLogsSortByDisplayname,
}

// GetListJobLogsSortByEnumValues Enumerates the set of values for ListJobLogsSortByEnum
func GetListJobLogsSortByEnumValues() []ListJobLogsSortByEnum {
	values := make([]ListJobLogsSortByEnum, 0)
	for _, v := range mappingListJobLogsSortBy {
		values = append(values, v)
	}
	return values
}

// ListJobLogsSortOrderEnum Enum with underlying type: string
type ListJobLogsSortOrderEnum string

// Set of constants representing the allowable values for ListJobLogsSortOrderEnum
const (
	ListJobLogsSortOrderAsc  ListJobLogsSortOrderEnum = "ASC"
	ListJobLogsSortOrderDesc ListJobLogsSortOrderEnum = "DESC"
)

var mappingListJobLogsSortOrder = map[string]ListJobLogsSortOrderEnum{
	"ASC":  ListJobLogsSortOrderAsc,
	"DESC": ListJobLogsSortOrderDesc,
}

// GetListJobLogsSortOrderEnumValues Enumerates the set of values for ListJobLogsSortOrderEnum
func GetListJobLogsSortOrderEnumValues() []ListJobLogsSortOrderEnum {
	values := make([]ListJobLogsSortOrderEnum, 0)
	for _, v := range mappingListJobLogsSortOrder {
		values = append(values, v)
	}
	return values
}
