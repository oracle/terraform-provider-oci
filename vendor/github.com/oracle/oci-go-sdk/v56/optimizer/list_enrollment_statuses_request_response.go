// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListEnrollmentStatusesRequest wrapper for the ListEnrollmentStatuses operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListEnrollmentStatuses.go.html to see an example of how to use ListEnrollmentStatusesRequest.
type ListEnrollmentStatusesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListEnrollmentStatusesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListEnrollmentStatusesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListEnrollmentStatusesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns results that match the Cloud Advisor enrollment status specified.
	Status ListEnrollmentStatusesStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEnrollmentStatusesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEnrollmentStatusesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEnrollmentStatusesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEnrollmentStatusesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListEnrollmentStatusesResponse wrapper for the ListEnrollmentStatuses operation
type ListEnrollmentStatusesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EnrollmentStatusCollection instances
	EnrollmentStatusCollection `presentIn:"body"`

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

func (response ListEnrollmentStatusesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEnrollmentStatusesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEnrollmentStatusesSortOrderEnum Enum with underlying type: string
type ListEnrollmentStatusesSortOrderEnum string

// Set of constants representing the allowable values for ListEnrollmentStatusesSortOrderEnum
const (
	ListEnrollmentStatusesSortOrderAsc  ListEnrollmentStatusesSortOrderEnum = "ASC"
	ListEnrollmentStatusesSortOrderDesc ListEnrollmentStatusesSortOrderEnum = "DESC"
)

var mappingListEnrollmentStatusesSortOrder = map[string]ListEnrollmentStatusesSortOrderEnum{
	"ASC":  ListEnrollmentStatusesSortOrderAsc,
	"DESC": ListEnrollmentStatusesSortOrderDesc,
}

// GetListEnrollmentStatusesSortOrderEnumValues Enumerates the set of values for ListEnrollmentStatusesSortOrderEnum
func GetListEnrollmentStatusesSortOrderEnumValues() []ListEnrollmentStatusesSortOrderEnum {
	values := make([]ListEnrollmentStatusesSortOrderEnum, 0)
	for _, v := range mappingListEnrollmentStatusesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListEnrollmentStatusesSortByEnum Enum with underlying type: string
type ListEnrollmentStatusesSortByEnum string

// Set of constants representing the allowable values for ListEnrollmentStatusesSortByEnum
const (
	ListEnrollmentStatusesSortByName        ListEnrollmentStatusesSortByEnum = "NAME"
	ListEnrollmentStatusesSortByTimecreated ListEnrollmentStatusesSortByEnum = "TIMECREATED"
)

var mappingListEnrollmentStatusesSortBy = map[string]ListEnrollmentStatusesSortByEnum{
	"NAME":        ListEnrollmentStatusesSortByName,
	"TIMECREATED": ListEnrollmentStatusesSortByTimecreated,
}

// GetListEnrollmentStatusesSortByEnumValues Enumerates the set of values for ListEnrollmentStatusesSortByEnum
func GetListEnrollmentStatusesSortByEnumValues() []ListEnrollmentStatusesSortByEnum {
	values := make([]ListEnrollmentStatusesSortByEnum, 0)
	for _, v := range mappingListEnrollmentStatusesSortBy {
		values = append(values, v)
	}
	return values
}

// ListEnrollmentStatusesLifecycleStateEnum Enum with underlying type: string
type ListEnrollmentStatusesLifecycleStateEnum string

// Set of constants representing the allowable values for ListEnrollmentStatusesLifecycleStateEnum
const (
	ListEnrollmentStatusesLifecycleStateActive    ListEnrollmentStatusesLifecycleStateEnum = "ACTIVE"
	ListEnrollmentStatusesLifecycleStateFailed    ListEnrollmentStatusesLifecycleStateEnum = "FAILED"
	ListEnrollmentStatusesLifecycleStateInactive  ListEnrollmentStatusesLifecycleStateEnum = "INACTIVE"
	ListEnrollmentStatusesLifecycleStateAttaching ListEnrollmentStatusesLifecycleStateEnum = "ATTACHING"
	ListEnrollmentStatusesLifecycleStateDetaching ListEnrollmentStatusesLifecycleStateEnum = "DETACHING"
	ListEnrollmentStatusesLifecycleStateDeleting  ListEnrollmentStatusesLifecycleStateEnum = "DELETING"
	ListEnrollmentStatusesLifecycleStateDeleted   ListEnrollmentStatusesLifecycleStateEnum = "DELETED"
	ListEnrollmentStatusesLifecycleStateUpdating  ListEnrollmentStatusesLifecycleStateEnum = "UPDATING"
	ListEnrollmentStatusesLifecycleStateCreating  ListEnrollmentStatusesLifecycleStateEnum = "CREATING"
)

var mappingListEnrollmentStatusesLifecycleState = map[string]ListEnrollmentStatusesLifecycleStateEnum{
	"ACTIVE":    ListEnrollmentStatusesLifecycleStateActive,
	"FAILED":    ListEnrollmentStatusesLifecycleStateFailed,
	"INACTIVE":  ListEnrollmentStatusesLifecycleStateInactive,
	"ATTACHING": ListEnrollmentStatusesLifecycleStateAttaching,
	"DETACHING": ListEnrollmentStatusesLifecycleStateDetaching,
	"DELETING":  ListEnrollmentStatusesLifecycleStateDeleting,
	"DELETED":   ListEnrollmentStatusesLifecycleStateDeleted,
	"UPDATING":  ListEnrollmentStatusesLifecycleStateUpdating,
	"CREATING":  ListEnrollmentStatusesLifecycleStateCreating,
}

// GetListEnrollmentStatusesLifecycleStateEnumValues Enumerates the set of values for ListEnrollmentStatusesLifecycleStateEnum
func GetListEnrollmentStatusesLifecycleStateEnumValues() []ListEnrollmentStatusesLifecycleStateEnum {
	values := make([]ListEnrollmentStatusesLifecycleStateEnum, 0)
	for _, v := range mappingListEnrollmentStatusesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListEnrollmentStatusesStatusEnum Enum with underlying type: string
type ListEnrollmentStatusesStatusEnum string

// Set of constants representing the allowable values for ListEnrollmentStatusesStatusEnum
const (
	ListEnrollmentStatusesStatusActive   ListEnrollmentStatusesStatusEnum = "ACTIVE"
	ListEnrollmentStatusesStatusInactive ListEnrollmentStatusesStatusEnum = "INACTIVE"
)

var mappingListEnrollmentStatusesStatus = map[string]ListEnrollmentStatusesStatusEnum{
	"ACTIVE":   ListEnrollmentStatusesStatusActive,
	"INACTIVE": ListEnrollmentStatusesStatusInactive,
}

// GetListEnrollmentStatusesStatusEnumValues Enumerates the set of values for ListEnrollmentStatusesStatusEnum
func GetListEnrollmentStatusesStatusEnumValues() []ListEnrollmentStatusesStatusEnum {
	values := make([]ListEnrollmentStatusesStatusEnum, 0)
	for _, v := range mappingListEnrollmentStatusesStatus {
		values = append(values, v)
	}
	return values
}
