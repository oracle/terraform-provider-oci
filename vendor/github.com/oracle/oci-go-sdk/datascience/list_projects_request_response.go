// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListProjectsRequest wrapper for the ListProjects operation
type ListProjectsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by OCID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm). Must be an OCID of the correct type for the resource type.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListProjectsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// <b>Filter</b> results by the OCID (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the user who created the resource.
	CreatedBy *string `mandatory:"false" contributesTo:"query" name:"createdBy"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 1000 is the maximum.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListProjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, results are shown
	// in descending order. When you sort by `displayName`, results are
	// shown in ascending order. Sort order for `displayName` field is case sensitive.
	SortBy ListProjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProjectsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListProjectsResponse wrapper for the ListProjects operation
type ListProjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ProjectSummary instances
	Items []ProjectSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListProjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProjectsLifecycleStateEnum Enum with underlying type: string
type ListProjectsLifecycleStateEnum string

// Set of constants representing the allowable values for ListProjectsLifecycleStateEnum
const (
	ListProjectsLifecycleStateActive   ListProjectsLifecycleStateEnum = "ACTIVE"
	ListProjectsLifecycleStateDeleting ListProjectsLifecycleStateEnum = "DELETING"
	ListProjectsLifecycleStateDeleted  ListProjectsLifecycleStateEnum = "DELETED"
)

var mappingListProjectsLifecycleState = map[string]ListProjectsLifecycleStateEnum{
	"ACTIVE":   ListProjectsLifecycleStateActive,
	"DELETING": ListProjectsLifecycleStateDeleting,
	"DELETED":  ListProjectsLifecycleStateDeleted,
}

// GetListProjectsLifecycleStateEnumValues Enumerates the set of values for ListProjectsLifecycleStateEnum
func GetListProjectsLifecycleStateEnumValues() []ListProjectsLifecycleStateEnum {
	values := make([]ListProjectsLifecycleStateEnum, 0)
	for _, v := range mappingListProjectsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListProjectsSortOrderEnum Enum with underlying type: string
type ListProjectsSortOrderEnum string

// Set of constants representing the allowable values for ListProjectsSortOrderEnum
const (
	ListProjectsSortOrderAsc  ListProjectsSortOrderEnum = "ASC"
	ListProjectsSortOrderDesc ListProjectsSortOrderEnum = "DESC"
)

var mappingListProjectsSortOrder = map[string]ListProjectsSortOrderEnum{
	"ASC":  ListProjectsSortOrderAsc,
	"DESC": ListProjectsSortOrderDesc,
}

// GetListProjectsSortOrderEnumValues Enumerates the set of values for ListProjectsSortOrderEnum
func GetListProjectsSortOrderEnumValues() []ListProjectsSortOrderEnum {
	values := make([]ListProjectsSortOrderEnum, 0)
	for _, v := range mappingListProjectsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListProjectsSortByEnum Enum with underlying type: string
type ListProjectsSortByEnum string

// Set of constants representing the allowable values for ListProjectsSortByEnum
const (
	ListProjectsSortByTimecreated ListProjectsSortByEnum = "timeCreated"
	ListProjectsSortByDisplayname ListProjectsSortByEnum = "displayName"
)

var mappingListProjectsSortBy = map[string]ListProjectsSortByEnum{
	"timeCreated": ListProjectsSortByTimecreated,
	"displayName": ListProjectsSortByDisplayname,
}

// GetListProjectsSortByEnumValues Enumerates the set of values for ListProjectsSortByEnum
func GetListProjectsSortByEnumValues() []ListProjectsSortByEnum {
	values := make([]ListProjectsSortByEnum, 0)
	for _, v := range mappingListProjectsSortBy {
		values = append(values, v)
	}
	return values
}
