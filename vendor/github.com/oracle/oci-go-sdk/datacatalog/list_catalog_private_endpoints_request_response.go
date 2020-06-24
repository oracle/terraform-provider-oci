// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListCatalogPrivateEndpointsRequest wrapper for the ListCatalogPrivateEndpoints operation
type ListCatalogPrivateEndpointsRequest struct {

	// The OCID of the compartment where you want to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListCatalogPrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCatalogPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListCatalogPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCatalogPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCatalogPrivateEndpointsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCatalogPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCatalogPrivateEndpointsResponse wrapper for the ListCatalogPrivateEndpoints operation
type ListCatalogPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CatalogPrivateEndpointSummary instances
	Items []CatalogPrivateEndpointSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCatalogPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCatalogPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCatalogPrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListCatalogPrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListCatalogPrivateEndpointsLifecycleStateEnum
const (
	ListCatalogPrivateEndpointsLifecycleStateCreating ListCatalogPrivateEndpointsLifecycleStateEnum = "CREATING"
	ListCatalogPrivateEndpointsLifecycleStateActive   ListCatalogPrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListCatalogPrivateEndpointsLifecycleStateInactive ListCatalogPrivateEndpointsLifecycleStateEnum = "INACTIVE"
	ListCatalogPrivateEndpointsLifecycleStateUpdating ListCatalogPrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListCatalogPrivateEndpointsLifecycleStateDeleting ListCatalogPrivateEndpointsLifecycleStateEnum = "DELETING"
	ListCatalogPrivateEndpointsLifecycleStateDeleted  ListCatalogPrivateEndpointsLifecycleStateEnum = "DELETED"
	ListCatalogPrivateEndpointsLifecycleStateFailed   ListCatalogPrivateEndpointsLifecycleStateEnum = "FAILED"
	ListCatalogPrivateEndpointsLifecycleStateMoving   ListCatalogPrivateEndpointsLifecycleStateEnum = "MOVING"
)

var mappingListCatalogPrivateEndpointsLifecycleState = map[string]ListCatalogPrivateEndpointsLifecycleStateEnum{
	"CREATING": ListCatalogPrivateEndpointsLifecycleStateCreating,
	"ACTIVE":   ListCatalogPrivateEndpointsLifecycleStateActive,
	"INACTIVE": ListCatalogPrivateEndpointsLifecycleStateInactive,
	"UPDATING": ListCatalogPrivateEndpointsLifecycleStateUpdating,
	"DELETING": ListCatalogPrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListCatalogPrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListCatalogPrivateEndpointsLifecycleStateFailed,
	"MOVING":   ListCatalogPrivateEndpointsLifecycleStateMoving,
}

// GetListCatalogPrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListCatalogPrivateEndpointsLifecycleStateEnum
func GetListCatalogPrivateEndpointsLifecycleStateEnumValues() []ListCatalogPrivateEndpointsLifecycleStateEnum {
	values := make([]ListCatalogPrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListCatalogPrivateEndpointsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListCatalogPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListCatalogPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListCatalogPrivateEndpointsSortOrderEnum
const (
	ListCatalogPrivateEndpointsSortOrderAsc  ListCatalogPrivateEndpointsSortOrderEnum = "ASC"
	ListCatalogPrivateEndpointsSortOrderDesc ListCatalogPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListCatalogPrivateEndpointsSortOrder = map[string]ListCatalogPrivateEndpointsSortOrderEnum{
	"ASC":  ListCatalogPrivateEndpointsSortOrderAsc,
	"DESC": ListCatalogPrivateEndpointsSortOrderDesc,
}

// GetListCatalogPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListCatalogPrivateEndpointsSortOrderEnum
func GetListCatalogPrivateEndpointsSortOrderEnumValues() []ListCatalogPrivateEndpointsSortOrderEnum {
	values := make([]ListCatalogPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListCatalogPrivateEndpointsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListCatalogPrivateEndpointsSortByEnum Enum with underlying type: string
type ListCatalogPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListCatalogPrivateEndpointsSortByEnum
const (
	ListCatalogPrivateEndpointsSortByTimecreated ListCatalogPrivateEndpointsSortByEnum = "TIMECREATED"
	ListCatalogPrivateEndpointsSortByDisplayname ListCatalogPrivateEndpointsSortByEnum = "DISPLAYNAME"
)

var mappingListCatalogPrivateEndpointsSortBy = map[string]ListCatalogPrivateEndpointsSortByEnum{
	"TIMECREATED": ListCatalogPrivateEndpointsSortByTimecreated,
	"DISPLAYNAME": ListCatalogPrivateEndpointsSortByDisplayname,
}

// GetListCatalogPrivateEndpointsSortByEnumValues Enumerates the set of values for ListCatalogPrivateEndpointsSortByEnum
func GetListCatalogPrivateEndpointsSortByEnumValues() []ListCatalogPrivateEndpointsSortByEnum {
	values := make([]ListCatalogPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListCatalogPrivateEndpointsSortBy {
		values = append(values, v)
	}
	return values
}
