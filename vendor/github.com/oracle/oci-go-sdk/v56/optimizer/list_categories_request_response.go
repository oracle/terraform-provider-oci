// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListCategoriesRequest wrapper for the ListCategories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListCategories.go.html to see an example of how to use ListCategoriesRequest.
type ListCategoriesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.
	// Can only be set to true when performing ListCompartments on the tenancy (root compartment).
	CompartmentIdInSubtree *bool `mandatory:"true" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Optional. A filter that returns results that match the name specified.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListCategoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListCategoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListCategoriesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCategoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCategoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCategoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCategoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCategoriesResponse wrapper for the ListCategories operation
type ListCategoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CategoryCollection instances
	CategoryCollection `presentIn:"body"`

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

func (response ListCategoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCategoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCategoriesSortOrderEnum Enum with underlying type: string
type ListCategoriesSortOrderEnum string

// Set of constants representing the allowable values for ListCategoriesSortOrderEnum
const (
	ListCategoriesSortOrderAsc  ListCategoriesSortOrderEnum = "ASC"
	ListCategoriesSortOrderDesc ListCategoriesSortOrderEnum = "DESC"
)

var mappingListCategoriesSortOrder = map[string]ListCategoriesSortOrderEnum{
	"ASC":  ListCategoriesSortOrderAsc,
	"DESC": ListCategoriesSortOrderDesc,
}

// GetListCategoriesSortOrderEnumValues Enumerates the set of values for ListCategoriesSortOrderEnum
func GetListCategoriesSortOrderEnumValues() []ListCategoriesSortOrderEnum {
	values := make([]ListCategoriesSortOrderEnum, 0)
	for _, v := range mappingListCategoriesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListCategoriesSortByEnum Enum with underlying type: string
type ListCategoriesSortByEnum string

// Set of constants representing the allowable values for ListCategoriesSortByEnum
const (
	ListCategoriesSortByName        ListCategoriesSortByEnum = "NAME"
	ListCategoriesSortByTimecreated ListCategoriesSortByEnum = "TIMECREATED"
)

var mappingListCategoriesSortBy = map[string]ListCategoriesSortByEnum{
	"NAME":        ListCategoriesSortByName,
	"TIMECREATED": ListCategoriesSortByTimecreated,
}

// GetListCategoriesSortByEnumValues Enumerates the set of values for ListCategoriesSortByEnum
func GetListCategoriesSortByEnumValues() []ListCategoriesSortByEnum {
	values := make([]ListCategoriesSortByEnum, 0)
	for _, v := range mappingListCategoriesSortBy {
		values = append(values, v)
	}
	return values
}

// ListCategoriesLifecycleStateEnum Enum with underlying type: string
type ListCategoriesLifecycleStateEnum string

// Set of constants representing the allowable values for ListCategoriesLifecycleStateEnum
const (
	ListCategoriesLifecycleStateActive    ListCategoriesLifecycleStateEnum = "ACTIVE"
	ListCategoriesLifecycleStateFailed    ListCategoriesLifecycleStateEnum = "FAILED"
	ListCategoriesLifecycleStateInactive  ListCategoriesLifecycleStateEnum = "INACTIVE"
	ListCategoriesLifecycleStateAttaching ListCategoriesLifecycleStateEnum = "ATTACHING"
	ListCategoriesLifecycleStateDetaching ListCategoriesLifecycleStateEnum = "DETACHING"
	ListCategoriesLifecycleStateDeleting  ListCategoriesLifecycleStateEnum = "DELETING"
	ListCategoriesLifecycleStateDeleted   ListCategoriesLifecycleStateEnum = "DELETED"
	ListCategoriesLifecycleStateUpdating  ListCategoriesLifecycleStateEnum = "UPDATING"
	ListCategoriesLifecycleStateCreating  ListCategoriesLifecycleStateEnum = "CREATING"
)

var mappingListCategoriesLifecycleState = map[string]ListCategoriesLifecycleStateEnum{
	"ACTIVE":    ListCategoriesLifecycleStateActive,
	"FAILED":    ListCategoriesLifecycleStateFailed,
	"INACTIVE":  ListCategoriesLifecycleStateInactive,
	"ATTACHING": ListCategoriesLifecycleStateAttaching,
	"DETACHING": ListCategoriesLifecycleStateDetaching,
	"DELETING":  ListCategoriesLifecycleStateDeleting,
	"DELETED":   ListCategoriesLifecycleStateDeleted,
	"UPDATING":  ListCategoriesLifecycleStateUpdating,
	"CREATING":  ListCategoriesLifecycleStateCreating,
}

// GetListCategoriesLifecycleStateEnumValues Enumerates the set of values for ListCategoriesLifecycleStateEnum
func GetListCategoriesLifecycleStateEnumValues() []ListCategoriesLifecycleStateEnum {
	values := make([]ListCategoriesLifecycleStateEnum, 0)
	for _, v := range mappingListCategoriesLifecycleState {
		values = append(values, v)
	}
	return values
}
