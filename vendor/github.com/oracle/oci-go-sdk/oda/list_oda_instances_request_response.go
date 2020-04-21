// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListOdaInstancesRequest wrapper for the ListOdaInstances operation
type ListOdaInstancesRequest struct {

	// List the Digital Assistant instances that belong to this compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// List only the information for the Digital Assistant instance with this user-friendly name. These names don't have to be unique and may change.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// List only the Digital Assistant instances that are in this lifecycle state.
	LifecycleState ListOdaInstancesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListOdaInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `TIMECREATED`.
	// The default sort order for `TIMECREATED` is descending, and the default sort order for `DISPLAYNAME` is ascending.
	SortBy ListOdaInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOdaInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOdaInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOdaInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOdaInstancesResponse wrapper for the ListOdaInstances operation
type ListOdaInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []OdaInstanceSummary instances
	Items []OdaInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you contact
	// Oracle about this request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOdaInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOdaInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOdaInstancesLifecycleStateEnum Enum with underlying type: string
type ListOdaInstancesLifecycleStateEnum string

// Set of constants representing the allowable values for ListOdaInstancesLifecycleStateEnum
const (
	ListOdaInstancesLifecycleStateCreating ListOdaInstancesLifecycleStateEnum = "CREATING"
	ListOdaInstancesLifecycleStateUpdating ListOdaInstancesLifecycleStateEnum = "UPDATING"
	ListOdaInstancesLifecycleStateActive   ListOdaInstancesLifecycleStateEnum = "ACTIVE"
	ListOdaInstancesLifecycleStateInactive ListOdaInstancesLifecycleStateEnum = "INACTIVE"
	ListOdaInstancesLifecycleStateDeleting ListOdaInstancesLifecycleStateEnum = "DELETING"
	ListOdaInstancesLifecycleStateDeleted  ListOdaInstancesLifecycleStateEnum = "DELETED"
	ListOdaInstancesLifecycleStateFailed   ListOdaInstancesLifecycleStateEnum = "FAILED"
)

var mappingListOdaInstancesLifecycleState = map[string]ListOdaInstancesLifecycleStateEnum{
	"CREATING": ListOdaInstancesLifecycleStateCreating,
	"UPDATING": ListOdaInstancesLifecycleStateUpdating,
	"ACTIVE":   ListOdaInstancesLifecycleStateActive,
	"INACTIVE": ListOdaInstancesLifecycleStateInactive,
	"DELETING": ListOdaInstancesLifecycleStateDeleting,
	"DELETED":  ListOdaInstancesLifecycleStateDeleted,
	"FAILED":   ListOdaInstancesLifecycleStateFailed,
}

// GetListOdaInstancesLifecycleStateEnumValues Enumerates the set of values for ListOdaInstancesLifecycleStateEnum
func GetListOdaInstancesLifecycleStateEnumValues() []ListOdaInstancesLifecycleStateEnum {
	values := make([]ListOdaInstancesLifecycleStateEnum, 0)
	for _, v := range mappingListOdaInstancesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListOdaInstancesSortOrderEnum Enum with underlying type: string
type ListOdaInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListOdaInstancesSortOrderEnum
const (
	ListOdaInstancesSortOrderAsc  ListOdaInstancesSortOrderEnum = "ASC"
	ListOdaInstancesSortOrderDesc ListOdaInstancesSortOrderEnum = "DESC"
)

var mappingListOdaInstancesSortOrder = map[string]ListOdaInstancesSortOrderEnum{
	"ASC":  ListOdaInstancesSortOrderAsc,
	"DESC": ListOdaInstancesSortOrderDesc,
}

// GetListOdaInstancesSortOrderEnumValues Enumerates the set of values for ListOdaInstancesSortOrderEnum
func GetListOdaInstancesSortOrderEnumValues() []ListOdaInstancesSortOrderEnum {
	values := make([]ListOdaInstancesSortOrderEnum, 0)
	for _, v := range mappingListOdaInstancesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListOdaInstancesSortByEnum Enum with underlying type: string
type ListOdaInstancesSortByEnum string

// Set of constants representing the allowable values for ListOdaInstancesSortByEnum
const (
	ListOdaInstancesSortByTimecreated ListOdaInstancesSortByEnum = "TIMECREATED"
	ListOdaInstancesSortByDisplayname ListOdaInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListOdaInstancesSortBy = map[string]ListOdaInstancesSortByEnum{
	"TIMECREATED": ListOdaInstancesSortByTimecreated,
	"DISPLAYNAME": ListOdaInstancesSortByDisplayname,
}

// GetListOdaInstancesSortByEnumValues Enumerates the set of values for ListOdaInstancesSortByEnum
func GetListOdaInstancesSortByEnumValues() []ListOdaInstancesSortByEnum {
	values := make([]ListOdaInstancesSortByEnum, 0)
	for _, v := range mappingListOdaInstancesSortBy {
		values = append(values, v)
	}
	return values
}
