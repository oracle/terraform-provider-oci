// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package operatoraccesscontrol

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListOperatorControlAssignmentsRequest wrapper for the ListOperatorControlAssignments operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/operatoraccesscontrol/ListOperatorControlAssignments.go.html to see an example of how to use ListOperatorControlAssignmentsRequest.
type ListOperatorControlAssignmentsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return OperatorControl that match the given operatorControlName.
	OperatorControlName *string `mandatory:"false" contributesTo:"query" name:"operatorControlName"`

	// A filter to return only resources that match the given ResourceName.
	ResourceName *string `mandatory:"false" contributesTo:"query" name:"resourceName"`

	// A filter to return only lists of resources that match the entire given service type.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// A filter to return only resources whose lifecycleState matches the given OperatorControlAssignment lifecycleState.
	LifecycleState ListOperatorControlAssignmentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListOperatorControlAssignmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOperatorControlAssignmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOperatorControlAssignmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOperatorControlAssignmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOperatorControlAssignmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOperatorControlAssignmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOperatorControlAssignmentsResponse wrapper for the ListOperatorControlAssignments operation
type ListOperatorControlAssignmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OperatorControlAssignmentCollection instances
	OperatorControlAssignmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOperatorControlAssignmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOperatorControlAssignmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOperatorControlAssignmentsLifecycleStateEnum Enum with underlying type: string
type ListOperatorControlAssignmentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListOperatorControlAssignmentsLifecycleStateEnum
const (
	ListOperatorControlAssignmentsLifecycleStateCreated        ListOperatorControlAssignmentsLifecycleStateEnum = "CREATED"
	ListOperatorControlAssignmentsLifecycleStateApplied        ListOperatorControlAssignmentsLifecycleStateEnum = "APPLIED"
	ListOperatorControlAssignmentsLifecycleStateApplyfailed    ListOperatorControlAssignmentsLifecycleStateEnum = "APPLYFAILED"
	ListOperatorControlAssignmentsLifecycleStateUpdating       ListOperatorControlAssignmentsLifecycleStateEnum = "UPDATING"
	ListOperatorControlAssignmentsLifecycleStateDeleting       ListOperatorControlAssignmentsLifecycleStateEnum = "DELETING"
	ListOperatorControlAssignmentsLifecycleStateDeleted        ListOperatorControlAssignmentsLifecycleStateEnum = "DELETED"
	ListOperatorControlAssignmentsLifecycleStateDeletionfailed ListOperatorControlAssignmentsLifecycleStateEnum = "DELETIONFAILED"
)

var mappingListOperatorControlAssignmentsLifecycleState = map[string]ListOperatorControlAssignmentsLifecycleStateEnum{
	"CREATED":        ListOperatorControlAssignmentsLifecycleStateCreated,
	"APPLIED":        ListOperatorControlAssignmentsLifecycleStateApplied,
	"APPLYFAILED":    ListOperatorControlAssignmentsLifecycleStateApplyfailed,
	"UPDATING":       ListOperatorControlAssignmentsLifecycleStateUpdating,
	"DELETING":       ListOperatorControlAssignmentsLifecycleStateDeleting,
	"DELETED":        ListOperatorControlAssignmentsLifecycleStateDeleted,
	"DELETIONFAILED": ListOperatorControlAssignmentsLifecycleStateDeletionfailed,
}

// GetListOperatorControlAssignmentsLifecycleStateEnumValues Enumerates the set of values for ListOperatorControlAssignmentsLifecycleStateEnum
func GetListOperatorControlAssignmentsLifecycleStateEnumValues() []ListOperatorControlAssignmentsLifecycleStateEnum {
	values := make([]ListOperatorControlAssignmentsLifecycleStateEnum, 0)
	for _, v := range mappingListOperatorControlAssignmentsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListOperatorControlAssignmentsSortOrderEnum Enum with underlying type: string
type ListOperatorControlAssignmentsSortOrderEnum string

// Set of constants representing the allowable values for ListOperatorControlAssignmentsSortOrderEnum
const (
	ListOperatorControlAssignmentsSortOrderAsc  ListOperatorControlAssignmentsSortOrderEnum = "ASC"
	ListOperatorControlAssignmentsSortOrderDesc ListOperatorControlAssignmentsSortOrderEnum = "DESC"
)

var mappingListOperatorControlAssignmentsSortOrder = map[string]ListOperatorControlAssignmentsSortOrderEnum{
	"ASC":  ListOperatorControlAssignmentsSortOrderAsc,
	"DESC": ListOperatorControlAssignmentsSortOrderDesc,
}

// GetListOperatorControlAssignmentsSortOrderEnumValues Enumerates the set of values for ListOperatorControlAssignmentsSortOrderEnum
func GetListOperatorControlAssignmentsSortOrderEnumValues() []ListOperatorControlAssignmentsSortOrderEnum {
	values := make([]ListOperatorControlAssignmentsSortOrderEnum, 0)
	for _, v := range mappingListOperatorControlAssignmentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListOperatorControlAssignmentsSortByEnum Enum with underlying type: string
type ListOperatorControlAssignmentsSortByEnum string

// Set of constants representing the allowable values for ListOperatorControlAssignmentsSortByEnum
const (
	ListOperatorControlAssignmentsSortByTimecreated ListOperatorControlAssignmentsSortByEnum = "timeCreated"
	ListOperatorControlAssignmentsSortByDisplayname ListOperatorControlAssignmentsSortByEnum = "displayName"
)

var mappingListOperatorControlAssignmentsSortBy = map[string]ListOperatorControlAssignmentsSortByEnum{
	"timeCreated": ListOperatorControlAssignmentsSortByTimecreated,
	"displayName": ListOperatorControlAssignmentsSortByDisplayname,
}

// GetListOperatorControlAssignmentsSortByEnumValues Enumerates the set of values for ListOperatorControlAssignmentsSortByEnum
func GetListOperatorControlAssignmentsSortByEnumValues() []ListOperatorControlAssignmentsSortByEnum {
	values := make([]ListOperatorControlAssignmentsSortByEnum, 0)
	for _, v := range mappingListOperatorControlAssignmentsSortBy {
		values = append(values, v)
	}
	return values
}
