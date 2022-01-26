// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListRecommendationsRequest wrapper for the ListRecommendations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListRecommendations.go.html to see an example of how to use ListRecommendationsRequest.
type ListRecommendationsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
	// Can only be set to true when performing ListCompartments on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"true" contributesTo:"query" name:"compartmentIdInSubtree"`

	// The unique OCID associated with the category.
	CategoryId *string `mandatory:"true" contributesTo:"query" name:"categoryId"`

	// Optional. A filter that returns results that match the name specified.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListRecommendationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListRecommendationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListRecommendationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter that returns recommendations that match the status specified.
	Status ListRecommendationsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRecommendationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRecommendationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRecommendationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRecommendationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRecommendationsResponse wrapper for the ListRecommendations operation
type ListRecommendationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RecommendationCollection instances
	RecommendationCollection `presentIn:"body"`

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

func (response ListRecommendationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRecommendationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRecommendationsSortOrderEnum Enum with underlying type: string
type ListRecommendationsSortOrderEnum string

// Set of constants representing the allowable values for ListRecommendationsSortOrderEnum
const (
	ListRecommendationsSortOrderAsc  ListRecommendationsSortOrderEnum = "ASC"
	ListRecommendationsSortOrderDesc ListRecommendationsSortOrderEnum = "DESC"
)

var mappingListRecommendationsSortOrder = map[string]ListRecommendationsSortOrderEnum{
	"ASC":  ListRecommendationsSortOrderAsc,
	"DESC": ListRecommendationsSortOrderDesc,
}

// GetListRecommendationsSortOrderEnumValues Enumerates the set of values for ListRecommendationsSortOrderEnum
func GetListRecommendationsSortOrderEnumValues() []ListRecommendationsSortOrderEnum {
	values := make([]ListRecommendationsSortOrderEnum, 0)
	for _, v := range mappingListRecommendationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListRecommendationsSortByEnum Enum with underlying type: string
type ListRecommendationsSortByEnum string

// Set of constants representing the allowable values for ListRecommendationsSortByEnum
const (
	ListRecommendationsSortByName        ListRecommendationsSortByEnum = "NAME"
	ListRecommendationsSortByTimecreated ListRecommendationsSortByEnum = "TIMECREATED"
)

var mappingListRecommendationsSortBy = map[string]ListRecommendationsSortByEnum{
	"NAME":        ListRecommendationsSortByName,
	"TIMECREATED": ListRecommendationsSortByTimecreated,
}

// GetListRecommendationsSortByEnumValues Enumerates the set of values for ListRecommendationsSortByEnum
func GetListRecommendationsSortByEnumValues() []ListRecommendationsSortByEnum {
	values := make([]ListRecommendationsSortByEnum, 0)
	for _, v := range mappingListRecommendationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListRecommendationsLifecycleStateEnum Enum with underlying type: string
type ListRecommendationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListRecommendationsLifecycleStateEnum
const (
	ListRecommendationsLifecycleStateActive    ListRecommendationsLifecycleStateEnum = "ACTIVE"
	ListRecommendationsLifecycleStateFailed    ListRecommendationsLifecycleStateEnum = "FAILED"
	ListRecommendationsLifecycleStateInactive  ListRecommendationsLifecycleStateEnum = "INACTIVE"
	ListRecommendationsLifecycleStateAttaching ListRecommendationsLifecycleStateEnum = "ATTACHING"
	ListRecommendationsLifecycleStateDetaching ListRecommendationsLifecycleStateEnum = "DETACHING"
	ListRecommendationsLifecycleStateDeleting  ListRecommendationsLifecycleStateEnum = "DELETING"
	ListRecommendationsLifecycleStateDeleted   ListRecommendationsLifecycleStateEnum = "DELETED"
	ListRecommendationsLifecycleStateUpdating  ListRecommendationsLifecycleStateEnum = "UPDATING"
	ListRecommendationsLifecycleStateCreating  ListRecommendationsLifecycleStateEnum = "CREATING"
)

var mappingListRecommendationsLifecycleState = map[string]ListRecommendationsLifecycleStateEnum{
	"ACTIVE":    ListRecommendationsLifecycleStateActive,
	"FAILED":    ListRecommendationsLifecycleStateFailed,
	"INACTIVE":  ListRecommendationsLifecycleStateInactive,
	"ATTACHING": ListRecommendationsLifecycleStateAttaching,
	"DETACHING": ListRecommendationsLifecycleStateDetaching,
	"DELETING":  ListRecommendationsLifecycleStateDeleting,
	"DELETED":   ListRecommendationsLifecycleStateDeleted,
	"UPDATING":  ListRecommendationsLifecycleStateUpdating,
	"CREATING":  ListRecommendationsLifecycleStateCreating,
}

// GetListRecommendationsLifecycleStateEnumValues Enumerates the set of values for ListRecommendationsLifecycleStateEnum
func GetListRecommendationsLifecycleStateEnumValues() []ListRecommendationsLifecycleStateEnum {
	values := make([]ListRecommendationsLifecycleStateEnum, 0)
	for _, v := range mappingListRecommendationsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListRecommendationsStatusEnum Enum with underlying type: string
type ListRecommendationsStatusEnum string

// Set of constants representing the allowable values for ListRecommendationsStatusEnum
const (
	ListRecommendationsStatusPending     ListRecommendationsStatusEnum = "PENDING"
	ListRecommendationsStatusDismissed   ListRecommendationsStatusEnum = "DISMISSED"
	ListRecommendationsStatusPostponed   ListRecommendationsStatusEnum = "POSTPONED"
	ListRecommendationsStatusImplemented ListRecommendationsStatusEnum = "IMPLEMENTED"
)

var mappingListRecommendationsStatus = map[string]ListRecommendationsStatusEnum{
	"PENDING":     ListRecommendationsStatusPending,
	"DISMISSED":   ListRecommendationsStatusDismissed,
	"POSTPONED":   ListRecommendationsStatusPostponed,
	"IMPLEMENTED": ListRecommendationsStatusImplemented,
}

// GetListRecommendationsStatusEnumValues Enumerates the set of values for ListRecommendationsStatusEnum
func GetListRecommendationsStatusEnumValues() []ListRecommendationsStatusEnum {
	values := make([]ListRecommendationsStatusEnum, 0)
	for _, v := range mappingListRecommendationsStatus {
		values = append(values, v)
	}
	return values
}
