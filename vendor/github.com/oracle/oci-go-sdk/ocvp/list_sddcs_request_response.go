// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocvp

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSddcsRequest wrapper for the ListSddcs operation
type ListSddcsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the availability domain that the Compute instances are running in.
	// Example: `Uocm:PHX-AD-1`
	ComputeAvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"computeAvailabilityDomain"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListSddcsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListSddcsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The lifecycle state of the resource.
	LifecycleState ListSddcsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSddcsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSddcsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSddcsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSddcsResponse wrapper for the ListSddcs operation
type ListSddcsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SddcCollection instances
	SddcCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSddcsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSddcsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSddcsSortOrderEnum Enum with underlying type: string
type ListSddcsSortOrderEnum string

// Set of constants representing the allowable values for ListSddcsSortOrderEnum
const (
	ListSddcsSortOrderAsc  ListSddcsSortOrderEnum = "ASC"
	ListSddcsSortOrderDesc ListSddcsSortOrderEnum = "DESC"
)

var mappingListSddcsSortOrder = map[string]ListSddcsSortOrderEnum{
	"ASC":  ListSddcsSortOrderAsc,
	"DESC": ListSddcsSortOrderDesc,
}

// GetListSddcsSortOrderEnumValues Enumerates the set of values for ListSddcsSortOrderEnum
func GetListSddcsSortOrderEnumValues() []ListSddcsSortOrderEnum {
	values := make([]ListSddcsSortOrderEnum, 0)
	for _, v := range mappingListSddcsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSddcsSortByEnum Enum with underlying type: string
type ListSddcsSortByEnum string

// Set of constants representing the allowable values for ListSddcsSortByEnum
const (
	ListSddcsSortByTimecreated ListSddcsSortByEnum = "timeCreated"
	ListSddcsSortByDisplayname ListSddcsSortByEnum = "displayName"
)

var mappingListSddcsSortBy = map[string]ListSddcsSortByEnum{
	"timeCreated": ListSddcsSortByTimecreated,
	"displayName": ListSddcsSortByDisplayname,
}

// GetListSddcsSortByEnumValues Enumerates the set of values for ListSddcsSortByEnum
func GetListSddcsSortByEnumValues() []ListSddcsSortByEnum {
	values := make([]ListSddcsSortByEnum, 0)
	for _, v := range mappingListSddcsSortBy {
		values = append(values, v)
	}
	return values
}

// ListSddcsLifecycleStateEnum Enum with underlying type: string
type ListSddcsLifecycleStateEnum string

// Set of constants representing the allowable values for ListSddcsLifecycleStateEnum
const (
	ListSddcsLifecycleStateCreating ListSddcsLifecycleStateEnum = "CREATING"
	ListSddcsLifecycleStateUpdating ListSddcsLifecycleStateEnum = "UPDATING"
	ListSddcsLifecycleStateActive   ListSddcsLifecycleStateEnum = "ACTIVE"
	ListSddcsLifecycleStateDeleting ListSddcsLifecycleStateEnum = "DELETING"
	ListSddcsLifecycleStateDeleted  ListSddcsLifecycleStateEnum = "DELETED"
	ListSddcsLifecycleStateFailed   ListSddcsLifecycleStateEnum = "FAILED"
)

var mappingListSddcsLifecycleState = map[string]ListSddcsLifecycleStateEnum{
	"CREATING": ListSddcsLifecycleStateCreating,
	"UPDATING": ListSddcsLifecycleStateUpdating,
	"ACTIVE":   ListSddcsLifecycleStateActive,
	"DELETING": ListSddcsLifecycleStateDeleting,
	"DELETED":  ListSddcsLifecycleStateDeleted,
	"FAILED":   ListSddcsLifecycleStateFailed,
}

// GetListSddcsLifecycleStateEnumValues Enumerates the set of values for ListSddcsLifecycleStateEnum
func GetListSddcsLifecycleStateEnumValues() []ListSddcsLifecycleStateEnum {
	values := make([]ListSddcsLifecycleStateEnum, 0)
	for _, v := range mappingListSddcsLifecycleState {
		values = append(values, v)
	}
	return values
}
