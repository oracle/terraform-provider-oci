// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListSourceAssociationsRequest wrapper for the ListSourceAssociations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListSourceAssociations.go.html to see an example of how to use ListSourceAssociationsRequest.
type ListSourceAssociationsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The souce name used for filtering associations.
	SourceName *string `mandatory:"true" contributesTo:"query" name:"sourceName"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The entity OCID.
	EntityId *string `mandatory:"false" contributesTo:"query" name:"entityId"`

	// The life cycle state used for filtering.  Only associations with the specified
	// life cycle state will be returned.
	LifeCycleState ListSourceAssociationsLifeCycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifeCycleState" omitEmpty:"true"`

	// A flag indicating whether or not to return the total number of items returned.
	IsShowTotal *bool `mandatory:"false" contributesTo:"query" name:"isShowTotal"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSourceAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned associations
	SortBy ListSourceAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSourceAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSourceAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSourceAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSourceAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSourceAssociationsResponse wrapper for the ListSourceAssociations operation
type ListSourceAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsAssociationCollection instances
	LogAnalyticsAssociationCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the previous page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSourceAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSourceAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSourceAssociationsLifeCycleStateEnum Enum with underlying type: string
type ListSourceAssociationsLifeCycleStateEnum string

// Set of constants representing the allowable values for ListSourceAssociationsLifeCycleStateEnum
const (
	ListSourceAssociationsLifeCycleStateAll        ListSourceAssociationsLifeCycleStateEnum = "ALL"
	ListSourceAssociationsLifeCycleStateAccepted   ListSourceAssociationsLifeCycleStateEnum = "ACCEPTED"
	ListSourceAssociationsLifeCycleStateInProgress ListSourceAssociationsLifeCycleStateEnum = "IN_PROGRESS"
	ListSourceAssociationsLifeCycleStateSucceeded  ListSourceAssociationsLifeCycleStateEnum = "SUCCEEDED"
	ListSourceAssociationsLifeCycleStateFailed     ListSourceAssociationsLifeCycleStateEnum = "FAILED"
)

var mappingListSourceAssociationsLifeCycleState = map[string]ListSourceAssociationsLifeCycleStateEnum{
	"ALL":         ListSourceAssociationsLifeCycleStateAll,
	"ACCEPTED":    ListSourceAssociationsLifeCycleStateAccepted,
	"IN_PROGRESS": ListSourceAssociationsLifeCycleStateInProgress,
	"SUCCEEDED":   ListSourceAssociationsLifeCycleStateSucceeded,
	"FAILED":      ListSourceAssociationsLifeCycleStateFailed,
}

// GetListSourceAssociationsLifeCycleStateEnumValues Enumerates the set of values for ListSourceAssociationsLifeCycleStateEnum
func GetListSourceAssociationsLifeCycleStateEnumValues() []ListSourceAssociationsLifeCycleStateEnum {
	values := make([]ListSourceAssociationsLifeCycleStateEnum, 0)
	for _, v := range mappingListSourceAssociationsLifeCycleState {
		values = append(values, v)
	}
	return values
}

// ListSourceAssociationsSortOrderEnum Enum with underlying type: string
type ListSourceAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListSourceAssociationsSortOrderEnum
const (
	ListSourceAssociationsSortOrderAsc  ListSourceAssociationsSortOrderEnum = "ASC"
	ListSourceAssociationsSortOrderDesc ListSourceAssociationsSortOrderEnum = "DESC"
)

var mappingListSourceAssociationsSortOrder = map[string]ListSourceAssociationsSortOrderEnum{
	"ASC":  ListSourceAssociationsSortOrderAsc,
	"DESC": ListSourceAssociationsSortOrderDesc,
}

// GetListSourceAssociationsSortOrderEnumValues Enumerates the set of values for ListSourceAssociationsSortOrderEnum
func GetListSourceAssociationsSortOrderEnumValues() []ListSourceAssociationsSortOrderEnum {
	values := make([]ListSourceAssociationsSortOrderEnum, 0)
	for _, v := range mappingListSourceAssociationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSourceAssociationsSortByEnum Enum with underlying type: string
type ListSourceAssociationsSortByEnum string

// Set of constants representing the allowable values for ListSourceAssociationsSortByEnum
const (
	ListSourceAssociationsSortByEntityname        ListSourceAssociationsSortByEnum = "entityName"
	ListSourceAssociationsSortByTimelastattempted ListSourceAssociationsSortByEnum = "timeLastAttempted"
	ListSourceAssociationsSortByStatus            ListSourceAssociationsSortByEnum = "status"
)

var mappingListSourceAssociationsSortBy = map[string]ListSourceAssociationsSortByEnum{
	"entityName":        ListSourceAssociationsSortByEntityname,
	"timeLastAttempted": ListSourceAssociationsSortByTimelastattempted,
	"status":            ListSourceAssociationsSortByStatus,
}

// GetListSourceAssociationsSortByEnumValues Enumerates the set of values for ListSourceAssociationsSortByEnum
func GetListSourceAssociationsSortByEnumValues() []ListSourceAssociationsSortByEnum {
	values := make([]ListSourceAssociationsSortByEnum, 0)
	for _, v := range mappingListSourceAssociationsSortBy {
		values = append(values, v)
	}
	return values
}
