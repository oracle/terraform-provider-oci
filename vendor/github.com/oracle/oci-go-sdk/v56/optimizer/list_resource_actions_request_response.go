// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListResourceActionsRequest wrapper for the ListResourceActions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListResourceActions.go.html to see an example of how to use ListResourceActionsRequest.
type ListResourceActionsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
	// Can only be set to true when performing ListCompartments on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"true" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The unique OCID associated with the recommendation.
	RecommendationId *string `mandatory:"true" contributesTo:"query" name:"recommendationId"`

	// Optional. A filter that returns results that match the name specified.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Optional. A filter that returns results that match the resource type specified.
	ResourceType *string `mandatory:"false" contributesTo:"query" name:"resourceType"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListResourceActionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListResourceActionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListResourceActionsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns recommendations that match the status specified.
	Status ListResourceActionsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceActionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceActionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceActionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceActionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListResourceActionsResponse wrapper for the ListResourceActions operation
type ListResourceActionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceActionCollection instances
	ResourceActionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListResourceActionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceActionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceActionsSortOrderEnum Enum with underlying type: string
type ListResourceActionsSortOrderEnum string

// Set of constants representing the allowable values for ListResourceActionsSortOrderEnum
const (
	ListResourceActionsSortOrderAsc  ListResourceActionsSortOrderEnum = "ASC"
	ListResourceActionsSortOrderDesc ListResourceActionsSortOrderEnum = "DESC"
)

var mappingListResourceActionsSortOrder = map[string]ListResourceActionsSortOrderEnum{
	"ASC":  ListResourceActionsSortOrderAsc,
	"DESC": ListResourceActionsSortOrderDesc,
}

// GetListResourceActionsSortOrderEnumValues Enumerates the set of values for ListResourceActionsSortOrderEnum
func GetListResourceActionsSortOrderEnumValues() []ListResourceActionsSortOrderEnum {
	values := make([]ListResourceActionsSortOrderEnum, 0)
	for _, v := range mappingListResourceActionsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListResourceActionsSortByEnum Enum with underlying type: string
type ListResourceActionsSortByEnum string

// Set of constants representing the allowable values for ListResourceActionsSortByEnum
const (
	ListResourceActionsSortByName        ListResourceActionsSortByEnum = "NAME"
	ListResourceActionsSortByTimecreated ListResourceActionsSortByEnum = "TIMECREATED"
)

var mappingListResourceActionsSortBy = map[string]ListResourceActionsSortByEnum{
	"NAME":        ListResourceActionsSortByName,
	"TIMECREATED": ListResourceActionsSortByTimecreated,
}

// GetListResourceActionsSortByEnumValues Enumerates the set of values for ListResourceActionsSortByEnum
func GetListResourceActionsSortByEnumValues() []ListResourceActionsSortByEnum {
	values := make([]ListResourceActionsSortByEnum, 0)
	for _, v := range mappingListResourceActionsSortBy {
		values = append(values, v)
	}
	return values
}

// ListResourceActionsLifecycleStateEnum Enum with underlying type: string
type ListResourceActionsLifecycleStateEnum string

// Set of constants representing the allowable values for ListResourceActionsLifecycleStateEnum
const (
	ListResourceActionsLifecycleStateActive    ListResourceActionsLifecycleStateEnum = "ACTIVE"
	ListResourceActionsLifecycleStateFailed    ListResourceActionsLifecycleStateEnum = "FAILED"
	ListResourceActionsLifecycleStateInactive  ListResourceActionsLifecycleStateEnum = "INACTIVE"
	ListResourceActionsLifecycleStateAttaching ListResourceActionsLifecycleStateEnum = "ATTACHING"
	ListResourceActionsLifecycleStateDetaching ListResourceActionsLifecycleStateEnum = "DETACHING"
	ListResourceActionsLifecycleStateDeleting  ListResourceActionsLifecycleStateEnum = "DELETING"
	ListResourceActionsLifecycleStateDeleted   ListResourceActionsLifecycleStateEnum = "DELETED"
	ListResourceActionsLifecycleStateUpdating  ListResourceActionsLifecycleStateEnum = "UPDATING"
	ListResourceActionsLifecycleStateCreating  ListResourceActionsLifecycleStateEnum = "CREATING"
)

var mappingListResourceActionsLifecycleState = map[string]ListResourceActionsLifecycleStateEnum{
	"ACTIVE":    ListResourceActionsLifecycleStateActive,
	"FAILED":    ListResourceActionsLifecycleStateFailed,
	"INACTIVE":  ListResourceActionsLifecycleStateInactive,
	"ATTACHING": ListResourceActionsLifecycleStateAttaching,
	"DETACHING": ListResourceActionsLifecycleStateDetaching,
	"DELETING":  ListResourceActionsLifecycleStateDeleting,
	"DELETED":   ListResourceActionsLifecycleStateDeleted,
	"UPDATING":  ListResourceActionsLifecycleStateUpdating,
	"CREATING":  ListResourceActionsLifecycleStateCreating,
}

// GetListResourceActionsLifecycleStateEnumValues Enumerates the set of values for ListResourceActionsLifecycleStateEnum
func GetListResourceActionsLifecycleStateEnumValues() []ListResourceActionsLifecycleStateEnum {
	values := make([]ListResourceActionsLifecycleStateEnum, 0)
	for _, v := range mappingListResourceActionsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListResourceActionsStatusEnum Enum with underlying type: string
type ListResourceActionsStatusEnum string

// Set of constants representing the allowable values for ListResourceActionsStatusEnum
const (
	ListResourceActionsStatusPending     ListResourceActionsStatusEnum = "PENDING"
	ListResourceActionsStatusDismissed   ListResourceActionsStatusEnum = "DISMISSED"
	ListResourceActionsStatusPostponed   ListResourceActionsStatusEnum = "POSTPONED"
	ListResourceActionsStatusImplemented ListResourceActionsStatusEnum = "IMPLEMENTED"
)

var mappingListResourceActionsStatus = map[string]ListResourceActionsStatusEnum{
	"PENDING":     ListResourceActionsStatusPending,
	"DISMISSED":   ListResourceActionsStatusDismissed,
	"POSTPONED":   ListResourceActionsStatusPostponed,
	"IMPLEMENTED": ListResourceActionsStatusImplemented,
}

// GetListResourceActionsStatusEnumValues Enumerates the set of values for ListResourceActionsStatusEnum
func GetListResourceActionsStatusEnumValues() []ListResourceActionsStatusEnum {
	values := make([]ListResourceActionsStatusEnum, 0)
	for _, v := range mappingListResourceActionsStatus {
		values = append(values, v)
	}
	return values
}
