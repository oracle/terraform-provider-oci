// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListLogAnalyticsEntityTypesRequest wrapper for the ListLogAnalyticsEntityTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsEntityTypes.go.html to see an example of how to use ListLogAnalyticsEntityTypesRequest.
type ListLogAnalyticsEntityTypesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// A filter to return only log analytics entity types whose name matches the entire name given. The match is
	// case-insensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only log analytics entity types whose name or internalName contains name given. The match
	// is case-insensitive.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// A filter to return CLOUD or NON_CLOUD entity types.
	CloudType ListLogAnalyticsEntityTypesCloudTypeEnum `mandatory:"false" contributesTo:"query" name:"cloudType" omitEmpty:"true"`

	// A filter to return only those log analytics entities with the specified lifecycle state. The state
	// value is case-insensitive.
	LifecycleState ListLogAnalyticsEntityTypesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogAnalyticsEntityTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated and timeUpdated
	// is descending. Default order for name is ascending. If no value is specified timeCreated is default.
	SortBy ListLogAnalyticsEntityTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogAnalyticsEntityTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogAnalyticsEntityTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogAnalyticsEntityTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogAnalyticsEntityTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogAnalyticsEntityTypesResponse wrapper for the ListLogAnalyticsEntityTypes operation
type ListLogAnalyticsEntityTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsEntityTypeCollection instances
	LogAnalyticsEntityTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogAnalyticsEntityTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogAnalyticsEntityTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogAnalyticsEntityTypesCloudTypeEnum Enum with underlying type: string
type ListLogAnalyticsEntityTypesCloudTypeEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntityTypesCloudTypeEnum
const (
	ListLogAnalyticsEntityTypesCloudTypeCloud    ListLogAnalyticsEntityTypesCloudTypeEnum = "CLOUD"
	ListLogAnalyticsEntityTypesCloudTypeNonCloud ListLogAnalyticsEntityTypesCloudTypeEnum = "NON_CLOUD"
	ListLogAnalyticsEntityTypesCloudTypeAll      ListLogAnalyticsEntityTypesCloudTypeEnum = "ALL"
)

var mappingListLogAnalyticsEntityTypesCloudType = map[string]ListLogAnalyticsEntityTypesCloudTypeEnum{
	"CLOUD":     ListLogAnalyticsEntityTypesCloudTypeCloud,
	"NON_CLOUD": ListLogAnalyticsEntityTypesCloudTypeNonCloud,
	"ALL":       ListLogAnalyticsEntityTypesCloudTypeAll,
}

// GetListLogAnalyticsEntityTypesCloudTypeEnumValues Enumerates the set of values for ListLogAnalyticsEntityTypesCloudTypeEnum
func GetListLogAnalyticsEntityTypesCloudTypeEnumValues() []ListLogAnalyticsEntityTypesCloudTypeEnum {
	values := make([]ListLogAnalyticsEntityTypesCloudTypeEnum, 0)
	for _, v := range mappingListLogAnalyticsEntityTypesCloudType {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsEntityTypesLifecycleStateEnum Enum with underlying type: string
type ListLogAnalyticsEntityTypesLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntityTypesLifecycleStateEnum
const (
	ListLogAnalyticsEntityTypesLifecycleStateActive  ListLogAnalyticsEntityTypesLifecycleStateEnum = "ACTIVE"
	ListLogAnalyticsEntityTypesLifecycleStateDeleted ListLogAnalyticsEntityTypesLifecycleStateEnum = "DELETED"
)

var mappingListLogAnalyticsEntityTypesLifecycleState = map[string]ListLogAnalyticsEntityTypesLifecycleStateEnum{
	"ACTIVE":  ListLogAnalyticsEntityTypesLifecycleStateActive,
	"DELETED": ListLogAnalyticsEntityTypesLifecycleStateDeleted,
}

// GetListLogAnalyticsEntityTypesLifecycleStateEnumValues Enumerates the set of values for ListLogAnalyticsEntityTypesLifecycleStateEnum
func GetListLogAnalyticsEntityTypesLifecycleStateEnumValues() []ListLogAnalyticsEntityTypesLifecycleStateEnum {
	values := make([]ListLogAnalyticsEntityTypesLifecycleStateEnum, 0)
	for _, v := range mappingListLogAnalyticsEntityTypesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsEntityTypesSortOrderEnum Enum with underlying type: string
type ListLogAnalyticsEntityTypesSortOrderEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntityTypesSortOrderEnum
const (
	ListLogAnalyticsEntityTypesSortOrderAsc  ListLogAnalyticsEntityTypesSortOrderEnum = "ASC"
	ListLogAnalyticsEntityTypesSortOrderDesc ListLogAnalyticsEntityTypesSortOrderEnum = "DESC"
)

var mappingListLogAnalyticsEntityTypesSortOrder = map[string]ListLogAnalyticsEntityTypesSortOrderEnum{
	"ASC":  ListLogAnalyticsEntityTypesSortOrderAsc,
	"DESC": ListLogAnalyticsEntityTypesSortOrderDesc,
}

// GetListLogAnalyticsEntityTypesSortOrderEnumValues Enumerates the set of values for ListLogAnalyticsEntityTypesSortOrderEnum
func GetListLogAnalyticsEntityTypesSortOrderEnumValues() []ListLogAnalyticsEntityTypesSortOrderEnum {
	values := make([]ListLogAnalyticsEntityTypesSortOrderEnum, 0)
	for _, v := range mappingListLogAnalyticsEntityTypesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsEntityTypesSortByEnum Enum with underlying type: string
type ListLogAnalyticsEntityTypesSortByEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntityTypesSortByEnum
const (
	ListLogAnalyticsEntityTypesSortByTimecreated ListLogAnalyticsEntityTypesSortByEnum = "timeCreated"
	ListLogAnalyticsEntityTypesSortByTimeupdated ListLogAnalyticsEntityTypesSortByEnum = "timeUpdated"
	ListLogAnalyticsEntityTypesSortByName        ListLogAnalyticsEntityTypesSortByEnum = "name"
)

var mappingListLogAnalyticsEntityTypesSortBy = map[string]ListLogAnalyticsEntityTypesSortByEnum{
	"timeCreated": ListLogAnalyticsEntityTypesSortByTimecreated,
	"timeUpdated": ListLogAnalyticsEntityTypesSortByTimeupdated,
	"name":        ListLogAnalyticsEntityTypesSortByName,
}

// GetListLogAnalyticsEntityTypesSortByEnumValues Enumerates the set of values for ListLogAnalyticsEntityTypesSortByEnum
func GetListLogAnalyticsEntityTypesSortByEnumValues() []ListLogAnalyticsEntityTypesSortByEnum {
	values := make([]ListLogAnalyticsEntityTypesSortByEnum, 0)
	for _, v := range mappingListLogAnalyticsEntityTypesSortBy {
		values = append(values, v)
	}
	return values
}
