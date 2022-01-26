// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListLogAnalyticsEntityTopologyRequest wrapper for the ListLogAnalyticsEntityTopology operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsEntityTopology.go.html to see an example of how to use ListLogAnalyticsEntityTopologyRequest.
type ListLogAnalyticsEntityTopologyRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The log analytics entity OCID.
	LogAnalyticsEntityId *string `mandatory:"true" contributesTo:"path" name:"logAnalyticsEntityId"`

	// A filter to return only those log analytics entities with the specified lifecycle state. The state
	// value is case-insensitive.
	LifecycleState ListLogAnalyticsEntityTopologyLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogAnalyticsEntityTopologySortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort entities by. Only one sort order may be provided. Default order for timeCreated and timeUpdated
	// is descending. Default order for entity name is ascending. If no value is specified timeCreated is default.
	SortBy ListLogAnalyticsEntityTopologySortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogAnalyticsEntityTopologyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogAnalyticsEntityTopologyRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogAnalyticsEntityTopologyRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogAnalyticsEntityTopologyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogAnalyticsEntityTopologyResponse wrapper for the ListLogAnalyticsEntityTopology operation
type ListLogAnalyticsEntityTopologyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsEntityTopologyCollection instances
	LogAnalyticsEntityTopologyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogAnalyticsEntityTopologyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogAnalyticsEntityTopologyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogAnalyticsEntityTopologyLifecycleStateEnum Enum with underlying type: string
type ListLogAnalyticsEntityTopologyLifecycleStateEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntityTopologyLifecycleStateEnum
const (
	ListLogAnalyticsEntityTopologyLifecycleStateActive  ListLogAnalyticsEntityTopologyLifecycleStateEnum = "ACTIVE"
	ListLogAnalyticsEntityTopologyLifecycleStateDeleted ListLogAnalyticsEntityTopologyLifecycleStateEnum = "DELETED"
)

var mappingListLogAnalyticsEntityTopologyLifecycleState = map[string]ListLogAnalyticsEntityTopologyLifecycleStateEnum{
	"ACTIVE":  ListLogAnalyticsEntityTopologyLifecycleStateActive,
	"DELETED": ListLogAnalyticsEntityTopologyLifecycleStateDeleted,
}

// GetListLogAnalyticsEntityTopologyLifecycleStateEnumValues Enumerates the set of values for ListLogAnalyticsEntityTopologyLifecycleStateEnum
func GetListLogAnalyticsEntityTopologyLifecycleStateEnumValues() []ListLogAnalyticsEntityTopologyLifecycleStateEnum {
	values := make([]ListLogAnalyticsEntityTopologyLifecycleStateEnum, 0)
	for _, v := range mappingListLogAnalyticsEntityTopologyLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsEntityTopologySortOrderEnum Enum with underlying type: string
type ListLogAnalyticsEntityTopologySortOrderEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntityTopologySortOrderEnum
const (
	ListLogAnalyticsEntityTopologySortOrderAsc  ListLogAnalyticsEntityTopologySortOrderEnum = "ASC"
	ListLogAnalyticsEntityTopologySortOrderDesc ListLogAnalyticsEntityTopologySortOrderEnum = "DESC"
)

var mappingListLogAnalyticsEntityTopologySortOrder = map[string]ListLogAnalyticsEntityTopologySortOrderEnum{
	"ASC":  ListLogAnalyticsEntityTopologySortOrderAsc,
	"DESC": ListLogAnalyticsEntityTopologySortOrderDesc,
}

// GetListLogAnalyticsEntityTopologySortOrderEnumValues Enumerates the set of values for ListLogAnalyticsEntityTopologySortOrderEnum
func GetListLogAnalyticsEntityTopologySortOrderEnumValues() []ListLogAnalyticsEntityTopologySortOrderEnum {
	values := make([]ListLogAnalyticsEntityTopologySortOrderEnum, 0)
	for _, v := range mappingListLogAnalyticsEntityTopologySortOrder {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsEntityTopologySortByEnum Enum with underlying type: string
type ListLogAnalyticsEntityTopologySortByEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEntityTopologySortByEnum
const (
	ListLogAnalyticsEntityTopologySortByTimecreated ListLogAnalyticsEntityTopologySortByEnum = "timeCreated"
	ListLogAnalyticsEntityTopologySortByTimeupdated ListLogAnalyticsEntityTopologySortByEnum = "timeUpdated"
	ListLogAnalyticsEntityTopologySortByName        ListLogAnalyticsEntityTopologySortByEnum = "name"
)

var mappingListLogAnalyticsEntityTopologySortBy = map[string]ListLogAnalyticsEntityTopologySortByEnum{
	"timeCreated": ListLogAnalyticsEntityTopologySortByTimecreated,
	"timeUpdated": ListLogAnalyticsEntityTopologySortByTimeupdated,
	"name":        ListLogAnalyticsEntityTopologySortByName,
}

// GetListLogAnalyticsEntityTopologySortByEnumValues Enumerates the set of values for ListLogAnalyticsEntityTopologySortByEnum
func GetListLogAnalyticsEntityTopologySortByEnumValues() []ListLogAnalyticsEntityTopologySortByEnum {
	values := make([]ListLogAnalyticsEntityTopologySortByEnum, 0)
	for _, v := range mappingListLogAnalyticsEntityTopologySortBy {
		values = append(values, v)
	}
	return values
}
