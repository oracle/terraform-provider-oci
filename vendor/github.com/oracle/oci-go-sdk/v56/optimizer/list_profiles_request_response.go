// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package optimizer

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListProfilesRequest wrapper for the ListProfiles operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/optimizer/ListProfiles.go.html to see an example of how to use ListProfilesRequest.
type ListProfilesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Optional. A filter that returns results that match the name specified.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListProfilesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending. Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListProfilesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter that returns results that match the lifecycle state specified.
	LifecycleState ListProfilesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProfilesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProfilesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProfilesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProfilesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListProfilesResponse wrapper for the ListProfiles operation
type ListProfilesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProfileCollection instances
	ProfileCollection `presentIn:"body"`

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

func (response ListProfilesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProfilesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProfilesSortOrderEnum Enum with underlying type: string
type ListProfilesSortOrderEnum string

// Set of constants representing the allowable values for ListProfilesSortOrderEnum
const (
	ListProfilesSortOrderAsc  ListProfilesSortOrderEnum = "ASC"
	ListProfilesSortOrderDesc ListProfilesSortOrderEnum = "DESC"
)

var mappingListProfilesSortOrder = map[string]ListProfilesSortOrderEnum{
	"ASC":  ListProfilesSortOrderAsc,
	"DESC": ListProfilesSortOrderDesc,
}

// GetListProfilesSortOrderEnumValues Enumerates the set of values for ListProfilesSortOrderEnum
func GetListProfilesSortOrderEnumValues() []ListProfilesSortOrderEnum {
	values := make([]ListProfilesSortOrderEnum, 0)
	for _, v := range mappingListProfilesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListProfilesSortByEnum Enum with underlying type: string
type ListProfilesSortByEnum string

// Set of constants representing the allowable values for ListProfilesSortByEnum
const (
	ListProfilesSortByName        ListProfilesSortByEnum = "NAME"
	ListProfilesSortByTimecreated ListProfilesSortByEnum = "TIMECREATED"
)

var mappingListProfilesSortBy = map[string]ListProfilesSortByEnum{
	"NAME":        ListProfilesSortByName,
	"TIMECREATED": ListProfilesSortByTimecreated,
}

// GetListProfilesSortByEnumValues Enumerates the set of values for ListProfilesSortByEnum
func GetListProfilesSortByEnumValues() []ListProfilesSortByEnum {
	values := make([]ListProfilesSortByEnum, 0)
	for _, v := range mappingListProfilesSortBy {
		values = append(values, v)
	}
	return values
}

// ListProfilesLifecycleStateEnum Enum with underlying type: string
type ListProfilesLifecycleStateEnum string

// Set of constants representing the allowable values for ListProfilesLifecycleStateEnum
const (
	ListProfilesLifecycleStateActive    ListProfilesLifecycleStateEnum = "ACTIVE"
	ListProfilesLifecycleStateFailed    ListProfilesLifecycleStateEnum = "FAILED"
	ListProfilesLifecycleStateInactive  ListProfilesLifecycleStateEnum = "INACTIVE"
	ListProfilesLifecycleStateAttaching ListProfilesLifecycleStateEnum = "ATTACHING"
	ListProfilesLifecycleStateDetaching ListProfilesLifecycleStateEnum = "DETACHING"
	ListProfilesLifecycleStateDeleting  ListProfilesLifecycleStateEnum = "DELETING"
	ListProfilesLifecycleStateDeleted   ListProfilesLifecycleStateEnum = "DELETED"
	ListProfilesLifecycleStateUpdating  ListProfilesLifecycleStateEnum = "UPDATING"
	ListProfilesLifecycleStateCreating  ListProfilesLifecycleStateEnum = "CREATING"
)

var mappingListProfilesLifecycleState = map[string]ListProfilesLifecycleStateEnum{
	"ACTIVE":    ListProfilesLifecycleStateActive,
	"FAILED":    ListProfilesLifecycleStateFailed,
	"INACTIVE":  ListProfilesLifecycleStateInactive,
	"ATTACHING": ListProfilesLifecycleStateAttaching,
	"DETACHING": ListProfilesLifecycleStateDetaching,
	"DELETING":  ListProfilesLifecycleStateDeleting,
	"DELETED":   ListProfilesLifecycleStateDeleted,
	"UPDATING":  ListProfilesLifecycleStateUpdating,
	"CREATING":  ListProfilesLifecycleStateCreating,
}

// GetListProfilesLifecycleStateEnumValues Enumerates the set of values for ListProfilesLifecycleStateEnum
func GetListProfilesLifecycleStateEnumValues() []ListProfilesLifecycleStateEnum {
	values := make([]ListProfilesLifecycleStateEnum, 0)
	for _, v := range mappingListProfilesLifecycleState {
		values = append(values, v)
	}
	return values
}
