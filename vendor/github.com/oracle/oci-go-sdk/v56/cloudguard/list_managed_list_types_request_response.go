// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListManagedListTypesRequest wrapper for the ListManagedListTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListManagedListTypes.go.html to see an example of how to use ListManagedListTypesRequest.
type ListManagedListTypesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListManagedListTypesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListManagedListTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListManagedListTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedListTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedListTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedListTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedListTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagedListTypesResponse wrapper for the ListManagedListTypes operation
type ListManagedListTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedListTypeCollection instances
	ManagedListTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedListTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedListTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedListTypesLifecycleStateEnum Enum with underlying type: string
type ListManagedListTypesLifecycleStateEnum string

// Set of constants representing the allowable values for ListManagedListTypesLifecycleStateEnum
const (
	ListManagedListTypesLifecycleStateCreating ListManagedListTypesLifecycleStateEnum = "CREATING"
	ListManagedListTypesLifecycleStateUpdating ListManagedListTypesLifecycleStateEnum = "UPDATING"
	ListManagedListTypesLifecycleStateActive   ListManagedListTypesLifecycleStateEnum = "ACTIVE"
	ListManagedListTypesLifecycleStateInactive ListManagedListTypesLifecycleStateEnum = "INACTIVE"
	ListManagedListTypesLifecycleStateDeleting ListManagedListTypesLifecycleStateEnum = "DELETING"
	ListManagedListTypesLifecycleStateDeleted  ListManagedListTypesLifecycleStateEnum = "DELETED"
	ListManagedListTypesLifecycleStateFailed   ListManagedListTypesLifecycleStateEnum = "FAILED"
)

var mappingListManagedListTypesLifecycleState = map[string]ListManagedListTypesLifecycleStateEnum{
	"CREATING": ListManagedListTypesLifecycleStateCreating,
	"UPDATING": ListManagedListTypesLifecycleStateUpdating,
	"ACTIVE":   ListManagedListTypesLifecycleStateActive,
	"INACTIVE": ListManagedListTypesLifecycleStateInactive,
	"DELETING": ListManagedListTypesLifecycleStateDeleting,
	"DELETED":  ListManagedListTypesLifecycleStateDeleted,
	"FAILED":   ListManagedListTypesLifecycleStateFailed,
}

// GetListManagedListTypesLifecycleStateEnumValues Enumerates the set of values for ListManagedListTypesLifecycleStateEnum
func GetListManagedListTypesLifecycleStateEnumValues() []ListManagedListTypesLifecycleStateEnum {
	values := make([]ListManagedListTypesLifecycleStateEnum, 0)
	for _, v := range mappingListManagedListTypesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListManagedListTypesSortOrderEnum Enum with underlying type: string
type ListManagedListTypesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedListTypesSortOrderEnum
const (
	ListManagedListTypesSortOrderAsc  ListManagedListTypesSortOrderEnum = "ASC"
	ListManagedListTypesSortOrderDesc ListManagedListTypesSortOrderEnum = "DESC"
)

var mappingListManagedListTypesSortOrder = map[string]ListManagedListTypesSortOrderEnum{
	"ASC":  ListManagedListTypesSortOrderAsc,
	"DESC": ListManagedListTypesSortOrderDesc,
}

// GetListManagedListTypesSortOrderEnumValues Enumerates the set of values for ListManagedListTypesSortOrderEnum
func GetListManagedListTypesSortOrderEnumValues() []ListManagedListTypesSortOrderEnum {
	values := make([]ListManagedListTypesSortOrderEnum, 0)
	for _, v := range mappingListManagedListTypesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagedListTypesSortByEnum Enum with underlying type: string
type ListManagedListTypesSortByEnum string

// Set of constants representing the allowable values for ListManagedListTypesSortByEnum
const (
	ListManagedListTypesSortByDisplayname ListManagedListTypesSortByEnum = "displayName"
	ListManagedListTypesSortByRisklevel   ListManagedListTypesSortByEnum = "riskLevel"
)

var mappingListManagedListTypesSortBy = map[string]ListManagedListTypesSortByEnum{
	"displayName": ListManagedListTypesSortByDisplayname,
	"riskLevel":   ListManagedListTypesSortByRisklevel,
}

// GetListManagedListTypesSortByEnumValues Enumerates the set of values for ListManagedListTypesSortByEnum
func GetListManagedListTypesSortByEnumValues() []ListManagedListTypesSortByEnum {
	values := make([]ListManagedListTypesSortByEnum, 0)
	for _, v := range mappingListManagedListTypesSortBy {
		values = append(values, v)
	}
	return values
}
