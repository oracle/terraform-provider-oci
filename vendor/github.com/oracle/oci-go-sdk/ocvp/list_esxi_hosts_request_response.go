// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package ocvp

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListEsxiHostsRequest wrapper for the ListEsxiHosts operation
type ListEsxiHostsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the SDDC.
	SddcId *string `mandatory:"false" contributesTo:"query" name:"sddcId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Compute instance.
	ComputeInstanceId *string `mandatory:"false" contributesTo:"query" name:"computeInstanceId"`

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
	SortOrder ListEsxiHostsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListEsxiHostsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request. If you need to contact Oracle about a particular
	// request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The lifecycle state of the resource.
	LifecycleState ListEsxiHostsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEsxiHostsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEsxiHostsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEsxiHostsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListEsxiHostsResponse wrapper for the ListEsxiHosts operation
type ListEsxiHostsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EsxiHostCollection instances
	EsxiHostCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEsxiHostsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEsxiHostsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEsxiHostsSortOrderEnum Enum with underlying type: string
type ListEsxiHostsSortOrderEnum string

// Set of constants representing the allowable values for ListEsxiHostsSortOrderEnum
const (
	ListEsxiHostsSortOrderAsc  ListEsxiHostsSortOrderEnum = "ASC"
	ListEsxiHostsSortOrderDesc ListEsxiHostsSortOrderEnum = "DESC"
)

var mappingListEsxiHostsSortOrder = map[string]ListEsxiHostsSortOrderEnum{
	"ASC":  ListEsxiHostsSortOrderAsc,
	"DESC": ListEsxiHostsSortOrderDesc,
}

// GetListEsxiHostsSortOrderEnumValues Enumerates the set of values for ListEsxiHostsSortOrderEnum
func GetListEsxiHostsSortOrderEnumValues() []ListEsxiHostsSortOrderEnum {
	values := make([]ListEsxiHostsSortOrderEnum, 0)
	for _, v := range mappingListEsxiHostsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListEsxiHostsSortByEnum Enum with underlying type: string
type ListEsxiHostsSortByEnum string

// Set of constants representing the allowable values for ListEsxiHostsSortByEnum
const (
	ListEsxiHostsSortByTimecreated ListEsxiHostsSortByEnum = "timeCreated"
	ListEsxiHostsSortByDisplayname ListEsxiHostsSortByEnum = "displayName"
)

var mappingListEsxiHostsSortBy = map[string]ListEsxiHostsSortByEnum{
	"timeCreated": ListEsxiHostsSortByTimecreated,
	"displayName": ListEsxiHostsSortByDisplayname,
}

// GetListEsxiHostsSortByEnumValues Enumerates the set of values for ListEsxiHostsSortByEnum
func GetListEsxiHostsSortByEnumValues() []ListEsxiHostsSortByEnum {
	values := make([]ListEsxiHostsSortByEnum, 0)
	for _, v := range mappingListEsxiHostsSortBy {
		values = append(values, v)
	}
	return values
}

// ListEsxiHostsLifecycleStateEnum Enum with underlying type: string
type ListEsxiHostsLifecycleStateEnum string

// Set of constants representing the allowable values for ListEsxiHostsLifecycleStateEnum
const (
	ListEsxiHostsLifecycleStateCreating ListEsxiHostsLifecycleStateEnum = "CREATING"
	ListEsxiHostsLifecycleStateUpdating ListEsxiHostsLifecycleStateEnum = "UPDATING"
	ListEsxiHostsLifecycleStateActive   ListEsxiHostsLifecycleStateEnum = "ACTIVE"
	ListEsxiHostsLifecycleStateDeleting ListEsxiHostsLifecycleStateEnum = "DELETING"
	ListEsxiHostsLifecycleStateDeleted  ListEsxiHostsLifecycleStateEnum = "DELETED"
	ListEsxiHostsLifecycleStateFailed   ListEsxiHostsLifecycleStateEnum = "FAILED"
)

var mappingListEsxiHostsLifecycleState = map[string]ListEsxiHostsLifecycleStateEnum{
	"CREATING": ListEsxiHostsLifecycleStateCreating,
	"UPDATING": ListEsxiHostsLifecycleStateUpdating,
	"ACTIVE":   ListEsxiHostsLifecycleStateActive,
	"DELETING": ListEsxiHostsLifecycleStateDeleting,
	"DELETED":  ListEsxiHostsLifecycleStateDeleted,
	"FAILED":   ListEsxiHostsLifecycleStateFailed,
}

// GetListEsxiHostsLifecycleStateEnumValues Enumerates the set of values for ListEsxiHostsLifecycleStateEnum
func GetListEsxiHostsLifecycleStateEnumValues() []ListEsxiHostsLifecycleStateEnum {
	values := make([]ListEsxiHostsLifecycleStateEnum, 0)
	for _, v := range mappingListEsxiHostsLifecycleState {
		values = append(values, v)
	}
	return values
}
