// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListResourceTypesRequest wrapper for the ListResourceTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListResourceTypes.go.html to see an example of how to use ListResourceTypesRequest.
type ListResourceTypesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListResourceTypesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListResourceTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for displayName is ascending. If no value is specified displayName is default.
	SortBy ListResourceTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListResourceTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListResourceTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListResourceTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListResourceTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListResourceTypesResponse wrapper for the ListResourceTypes operation
type ListResourceTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ResourceTypeCollection instances
	ResourceTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListResourceTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListResourceTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListResourceTypesLifecycleStateEnum Enum with underlying type: string
type ListResourceTypesLifecycleStateEnum string

// Set of constants representing the allowable values for ListResourceTypesLifecycleStateEnum
const (
	ListResourceTypesLifecycleStateCreating ListResourceTypesLifecycleStateEnum = "CREATING"
	ListResourceTypesLifecycleStateUpdating ListResourceTypesLifecycleStateEnum = "UPDATING"
	ListResourceTypesLifecycleStateActive   ListResourceTypesLifecycleStateEnum = "ACTIVE"
	ListResourceTypesLifecycleStateInactive ListResourceTypesLifecycleStateEnum = "INACTIVE"
	ListResourceTypesLifecycleStateDeleting ListResourceTypesLifecycleStateEnum = "DELETING"
	ListResourceTypesLifecycleStateDeleted  ListResourceTypesLifecycleStateEnum = "DELETED"
	ListResourceTypesLifecycleStateFailed   ListResourceTypesLifecycleStateEnum = "FAILED"
)

var mappingListResourceTypesLifecycleState = map[string]ListResourceTypesLifecycleStateEnum{
	"CREATING": ListResourceTypesLifecycleStateCreating,
	"UPDATING": ListResourceTypesLifecycleStateUpdating,
	"ACTIVE":   ListResourceTypesLifecycleStateActive,
	"INACTIVE": ListResourceTypesLifecycleStateInactive,
	"DELETING": ListResourceTypesLifecycleStateDeleting,
	"DELETED":  ListResourceTypesLifecycleStateDeleted,
	"FAILED":   ListResourceTypesLifecycleStateFailed,
}

// GetListResourceTypesLifecycleStateEnumValues Enumerates the set of values for ListResourceTypesLifecycleStateEnum
func GetListResourceTypesLifecycleStateEnumValues() []ListResourceTypesLifecycleStateEnum {
	values := make([]ListResourceTypesLifecycleStateEnum, 0)
	for _, v := range mappingListResourceTypesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListResourceTypesSortOrderEnum Enum with underlying type: string
type ListResourceTypesSortOrderEnum string

// Set of constants representing the allowable values for ListResourceTypesSortOrderEnum
const (
	ListResourceTypesSortOrderAsc  ListResourceTypesSortOrderEnum = "ASC"
	ListResourceTypesSortOrderDesc ListResourceTypesSortOrderEnum = "DESC"
)

var mappingListResourceTypesSortOrder = map[string]ListResourceTypesSortOrderEnum{
	"ASC":  ListResourceTypesSortOrderAsc,
	"DESC": ListResourceTypesSortOrderDesc,
}

// GetListResourceTypesSortOrderEnumValues Enumerates the set of values for ListResourceTypesSortOrderEnum
func GetListResourceTypesSortOrderEnumValues() []ListResourceTypesSortOrderEnum {
	values := make([]ListResourceTypesSortOrderEnum, 0)
	for _, v := range mappingListResourceTypesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListResourceTypesSortByEnum Enum with underlying type: string
type ListResourceTypesSortByEnum string

// Set of constants representing the allowable values for ListResourceTypesSortByEnum
const (
	ListResourceTypesSortByDisplayname ListResourceTypesSortByEnum = "displayName"
	ListResourceTypesSortByRisklevel   ListResourceTypesSortByEnum = "riskLevel"
)

var mappingListResourceTypesSortBy = map[string]ListResourceTypesSortByEnum{
	"displayName": ListResourceTypesSortByDisplayname,
	"riskLevel":   ListResourceTypesSortByRisklevel,
}

// GetListResourceTypesSortByEnumValues Enumerates the set of values for ListResourceTypesSortByEnum
func GetListResourceTypesSortByEnumValues() []ListResourceTypesSortByEnum {
	values := make([]ListResourceTypesSortByEnum, 0)
	for _, v := range mappingListResourceTypesSortBy {
		values = append(values, v)
	}
	return values
}
