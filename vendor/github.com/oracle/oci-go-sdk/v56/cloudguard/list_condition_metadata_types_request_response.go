// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListConditionMetadataTypesRequest wrapper for the ListConditionMetadataTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListConditionMetadataTypes.go.html to see an example of how to use ListConditionMetadataTypesRequest.
type ListConditionMetadataTypesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListConditionMetadataTypesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConditionMetadataTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListConditionMetadataTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConditionMetadataTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConditionMetadataTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConditionMetadataTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConditionMetadataTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListConditionMetadataTypesResponse wrapper for the ListConditionMetadataTypes operation
type ListConditionMetadataTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConditionMetadataTypeCollection instances
	ConditionMetadataTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListConditionMetadataTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConditionMetadataTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConditionMetadataTypesLifecycleStateEnum Enum with underlying type: string
type ListConditionMetadataTypesLifecycleStateEnum string

// Set of constants representing the allowable values for ListConditionMetadataTypesLifecycleStateEnum
const (
	ListConditionMetadataTypesLifecycleStateCreating ListConditionMetadataTypesLifecycleStateEnum = "CREATING"
	ListConditionMetadataTypesLifecycleStateUpdating ListConditionMetadataTypesLifecycleStateEnum = "UPDATING"
	ListConditionMetadataTypesLifecycleStateActive   ListConditionMetadataTypesLifecycleStateEnum = "ACTIVE"
	ListConditionMetadataTypesLifecycleStateInactive ListConditionMetadataTypesLifecycleStateEnum = "INACTIVE"
	ListConditionMetadataTypesLifecycleStateDeleting ListConditionMetadataTypesLifecycleStateEnum = "DELETING"
	ListConditionMetadataTypesLifecycleStateDeleted  ListConditionMetadataTypesLifecycleStateEnum = "DELETED"
	ListConditionMetadataTypesLifecycleStateFailed   ListConditionMetadataTypesLifecycleStateEnum = "FAILED"
)

var mappingListConditionMetadataTypesLifecycleState = map[string]ListConditionMetadataTypesLifecycleStateEnum{
	"CREATING": ListConditionMetadataTypesLifecycleStateCreating,
	"UPDATING": ListConditionMetadataTypesLifecycleStateUpdating,
	"ACTIVE":   ListConditionMetadataTypesLifecycleStateActive,
	"INACTIVE": ListConditionMetadataTypesLifecycleStateInactive,
	"DELETING": ListConditionMetadataTypesLifecycleStateDeleting,
	"DELETED":  ListConditionMetadataTypesLifecycleStateDeleted,
	"FAILED":   ListConditionMetadataTypesLifecycleStateFailed,
}

// GetListConditionMetadataTypesLifecycleStateEnumValues Enumerates the set of values for ListConditionMetadataTypesLifecycleStateEnum
func GetListConditionMetadataTypesLifecycleStateEnumValues() []ListConditionMetadataTypesLifecycleStateEnum {
	values := make([]ListConditionMetadataTypesLifecycleStateEnum, 0)
	for _, v := range mappingListConditionMetadataTypesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListConditionMetadataTypesSortOrderEnum Enum with underlying type: string
type ListConditionMetadataTypesSortOrderEnum string

// Set of constants representing the allowable values for ListConditionMetadataTypesSortOrderEnum
const (
	ListConditionMetadataTypesSortOrderAsc  ListConditionMetadataTypesSortOrderEnum = "ASC"
	ListConditionMetadataTypesSortOrderDesc ListConditionMetadataTypesSortOrderEnum = "DESC"
)

var mappingListConditionMetadataTypesSortOrder = map[string]ListConditionMetadataTypesSortOrderEnum{
	"ASC":  ListConditionMetadataTypesSortOrderAsc,
	"DESC": ListConditionMetadataTypesSortOrderDesc,
}

// GetListConditionMetadataTypesSortOrderEnumValues Enumerates the set of values for ListConditionMetadataTypesSortOrderEnum
func GetListConditionMetadataTypesSortOrderEnumValues() []ListConditionMetadataTypesSortOrderEnum {
	values := make([]ListConditionMetadataTypesSortOrderEnum, 0)
	for _, v := range mappingListConditionMetadataTypesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListConditionMetadataTypesSortByEnum Enum with underlying type: string
type ListConditionMetadataTypesSortByEnum string

// Set of constants representing the allowable values for ListConditionMetadataTypesSortByEnum
const (
	ListConditionMetadataTypesSortByTimecreated ListConditionMetadataTypesSortByEnum = "timeCreated"
	ListConditionMetadataTypesSortByDisplayname ListConditionMetadataTypesSortByEnum = "displayName"
)

var mappingListConditionMetadataTypesSortBy = map[string]ListConditionMetadataTypesSortByEnum{
	"timeCreated": ListConditionMetadataTypesSortByTimecreated,
	"displayName": ListConditionMetadataTypesSortByDisplayname,
}

// GetListConditionMetadataTypesSortByEnumValues Enumerates the set of values for ListConditionMetadataTypesSortByEnum
func GetListConditionMetadataTypesSortByEnumValues() []ListConditionMetadataTypesSortByEnum {
	values := make([]ListConditionMetadataTypesSortByEnum, 0)
	for _, v := range mappingListConditionMetadataTypesSortBy {
		values = append(values, v)
	}
	return values
}
