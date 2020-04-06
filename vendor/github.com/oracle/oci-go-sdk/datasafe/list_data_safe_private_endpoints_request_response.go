// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListDataSafePrivateEndpointsRequest wrapper for the ListDataSafePrivateEndpoints operation
type ListDataSafePrivateEndpointsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the private endpoints that match the specified VCN OCID.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// A filter to return only resources that match the specified lifecycle state.
	LifecycleState ListDataSafePrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The beginning page from which the results start retrieving.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListDataSafePrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListDataSafePrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDataSafePrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDataSafePrivateEndpointsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDataSafePrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDataSafePrivateEndpointsResponse wrapper for the ListDataSafePrivateEndpoints operation
type ListDataSafePrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []DataSafePrivateEndpointSummary instances
	Items []DataSafePrivateEndpointSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDataSafePrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDataSafePrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDataSafePrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsLifecycleStateEnum
const (
	ListDataSafePrivateEndpointsLifecycleStateCreating ListDataSafePrivateEndpointsLifecycleStateEnum = "CREATING"
	ListDataSafePrivateEndpointsLifecycleStateUpdating ListDataSafePrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListDataSafePrivateEndpointsLifecycleStateActive   ListDataSafePrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListDataSafePrivateEndpointsLifecycleStateDeleting ListDataSafePrivateEndpointsLifecycleStateEnum = "DELETING"
	ListDataSafePrivateEndpointsLifecycleStateDeleted  ListDataSafePrivateEndpointsLifecycleStateEnum = "DELETED"
	ListDataSafePrivateEndpointsLifecycleStateFailed   ListDataSafePrivateEndpointsLifecycleStateEnum = "FAILED"
	ListDataSafePrivateEndpointsLifecycleStateNa       ListDataSafePrivateEndpointsLifecycleStateEnum = "NA"
)

var mappingListDataSafePrivateEndpointsLifecycleState = map[string]ListDataSafePrivateEndpointsLifecycleStateEnum{
	"CREATING": ListDataSafePrivateEndpointsLifecycleStateCreating,
	"UPDATING": ListDataSafePrivateEndpointsLifecycleStateUpdating,
	"ACTIVE":   ListDataSafePrivateEndpointsLifecycleStateActive,
	"DELETING": ListDataSafePrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListDataSafePrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListDataSafePrivateEndpointsLifecycleStateFailed,
	"NA":       ListDataSafePrivateEndpointsLifecycleStateNa,
}

// GetListDataSafePrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsLifecycleStateEnum
func GetListDataSafePrivateEndpointsLifecycleStateEnumValues() []ListDataSafePrivateEndpointsLifecycleStateEnum {
	values := make([]ListDataSafePrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDataSafePrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsSortOrderEnum
const (
	ListDataSafePrivateEndpointsSortOrderAsc  ListDataSafePrivateEndpointsSortOrderEnum = "ASC"
	ListDataSafePrivateEndpointsSortOrderDesc ListDataSafePrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListDataSafePrivateEndpointsSortOrder = map[string]ListDataSafePrivateEndpointsSortOrderEnum{
	"ASC":  ListDataSafePrivateEndpointsSortOrderAsc,
	"DESC": ListDataSafePrivateEndpointsSortOrderDesc,
}

// GetListDataSafePrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsSortOrderEnum
func GetListDataSafePrivateEndpointsSortOrderEnumValues() []ListDataSafePrivateEndpointsSortOrderEnum {
	values := make([]ListDataSafePrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDataSafePrivateEndpointsSortByEnum Enum with underlying type: string
type ListDataSafePrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListDataSafePrivateEndpointsSortByEnum
const (
	ListDataSafePrivateEndpointsSortByTimecreated ListDataSafePrivateEndpointsSortByEnum = "TIMECREATED"
	ListDataSafePrivateEndpointsSortByDisplayname ListDataSafePrivateEndpointsSortByEnum = "DISPLAYNAME"
)

var mappingListDataSafePrivateEndpointsSortBy = map[string]ListDataSafePrivateEndpointsSortByEnum{
	"TIMECREATED": ListDataSafePrivateEndpointsSortByTimecreated,
	"DISPLAYNAME": ListDataSafePrivateEndpointsSortByDisplayname,
}

// GetListDataSafePrivateEndpointsSortByEnumValues Enumerates the set of values for ListDataSafePrivateEndpointsSortByEnum
func GetListDataSafePrivateEndpointsSortByEnumValues() []ListDataSafePrivateEndpointsSortByEnum {
	values := make([]ListDataSafePrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListDataSafePrivateEndpointsSortBy {
		values = append(values, v)
	}
	return values
}
