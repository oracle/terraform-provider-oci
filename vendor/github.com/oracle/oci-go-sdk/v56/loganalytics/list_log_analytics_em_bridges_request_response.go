// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListLogAnalyticsEmBridgesRequest wrapper for the ListLogAnalyticsEmBridges operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogAnalyticsEmBridges.go.html to see an example of how to use ListLogAnalyticsEmBridgesRequest.
type ListLogAnalyticsEmBridgesRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only log analytics enterprise manager bridge name whose name matches the entire name given. The match
	// is case-insensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only log analytics enterprise manager bridges matching all the lifecycle states specified for this parameter.
	LifecycleState []EmBridgeLifecycleStatesEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only log analytics enterprise manager bridges whose lifecycleDetails contains the specified string.
	LifecycleDetailsContains *string `mandatory:"false" contributesTo:"query" name:"lifecycleDetailsContains"`

	// Filter by the processing status of the latest upload from enterprise manager.
	ImportStatus []EmBridgeLatestImportProcessingStatusEnum `contributesTo:"query" name:"importStatus" omitEmpty:"true" collectionFormat:"multi"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogAnalyticsEmBridgesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort enterprise manager bridges by. Only one sort order may be provided. Default order for timeCreated and timeUpdated
	// is descending. Default order for enterprise manager name is ascending. If no value is specified timeCreated is default.
	SortBy ListLogAnalyticsEmBridgesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogAnalyticsEmBridgesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogAnalyticsEmBridgesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogAnalyticsEmBridgesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogAnalyticsEmBridgesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListLogAnalyticsEmBridgesResponse wrapper for the ListLogAnalyticsEmBridges operation
type ListLogAnalyticsEmBridgesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsEmBridgeCollection instances
	LogAnalyticsEmBridgeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogAnalyticsEmBridgesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogAnalyticsEmBridgesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogAnalyticsEmBridgesSortOrderEnum Enum with underlying type: string
type ListLogAnalyticsEmBridgesSortOrderEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEmBridgesSortOrderEnum
const (
	ListLogAnalyticsEmBridgesSortOrderAsc  ListLogAnalyticsEmBridgesSortOrderEnum = "ASC"
	ListLogAnalyticsEmBridgesSortOrderDesc ListLogAnalyticsEmBridgesSortOrderEnum = "DESC"
)

var mappingListLogAnalyticsEmBridgesSortOrder = map[string]ListLogAnalyticsEmBridgesSortOrderEnum{
	"ASC":  ListLogAnalyticsEmBridgesSortOrderAsc,
	"DESC": ListLogAnalyticsEmBridgesSortOrderDesc,
}

// GetListLogAnalyticsEmBridgesSortOrderEnumValues Enumerates the set of values for ListLogAnalyticsEmBridgesSortOrderEnum
func GetListLogAnalyticsEmBridgesSortOrderEnumValues() []ListLogAnalyticsEmBridgesSortOrderEnum {
	values := make([]ListLogAnalyticsEmBridgesSortOrderEnum, 0)
	for _, v := range mappingListLogAnalyticsEmBridgesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListLogAnalyticsEmBridgesSortByEnum Enum with underlying type: string
type ListLogAnalyticsEmBridgesSortByEnum string

// Set of constants representing the allowable values for ListLogAnalyticsEmBridgesSortByEnum
const (
	ListLogAnalyticsEmBridgesSortByTimecreated ListLogAnalyticsEmBridgesSortByEnum = "timeCreated"
	ListLogAnalyticsEmBridgesSortByTimeupdated ListLogAnalyticsEmBridgesSortByEnum = "timeUpdated"
	ListLogAnalyticsEmBridgesSortByDisplayname ListLogAnalyticsEmBridgesSortByEnum = "displayName"
)

var mappingListLogAnalyticsEmBridgesSortBy = map[string]ListLogAnalyticsEmBridgesSortByEnum{
	"timeCreated": ListLogAnalyticsEmBridgesSortByTimecreated,
	"timeUpdated": ListLogAnalyticsEmBridgesSortByTimeupdated,
	"displayName": ListLogAnalyticsEmBridgesSortByDisplayname,
}

// GetListLogAnalyticsEmBridgesSortByEnumValues Enumerates the set of values for ListLogAnalyticsEmBridgesSortByEnum
func GetListLogAnalyticsEmBridgesSortByEnumValues() []ListLogAnalyticsEmBridgesSortByEnum {
	values := make([]ListLogAnalyticsEmBridgesSortByEnum, 0)
	for _, v := range mappingListLogAnalyticsEmBridgesSortBy {
		values = append(values, v)
	}
	return values
}
