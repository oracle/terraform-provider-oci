// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListSoftwareSourcesRequest wrapper for the ListSoftwareSources operation
type ListSoftwareSourcesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListSoftwareSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListSoftwareSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The current lifecycle state for the object.
	LifecycleState ListSoftwareSourcesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSoftwareSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSoftwareSourcesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSoftwareSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSoftwareSourcesResponse wrapper for the ListSoftwareSources operation
type ListSoftwareSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SoftwareSourceSummary instances
	Items []SoftwareSourceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `SoftwareSource`s. If this header
	// appears in the response, then this is a partial list of software
	// sources. Include this value as the `page` parameter in a
	// subsequent GET request to get the next batch of software
	// sources.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSoftwareSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSoftwareSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSoftwareSourcesSortOrderEnum Enum with underlying type: string
type ListSoftwareSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListSoftwareSourcesSortOrderEnum
const (
	ListSoftwareSourcesSortOrderAsc  ListSoftwareSourcesSortOrderEnum = "ASC"
	ListSoftwareSourcesSortOrderDesc ListSoftwareSourcesSortOrderEnum = "DESC"
)

var mappingListSoftwareSourcesSortOrder = map[string]ListSoftwareSourcesSortOrderEnum{
	"ASC":  ListSoftwareSourcesSortOrderAsc,
	"DESC": ListSoftwareSourcesSortOrderDesc,
}

// GetListSoftwareSourcesSortOrderEnumValues Enumerates the set of values for ListSoftwareSourcesSortOrderEnum
func GetListSoftwareSourcesSortOrderEnumValues() []ListSoftwareSourcesSortOrderEnum {
	values := make([]ListSoftwareSourcesSortOrderEnum, 0)
	for _, v := range mappingListSoftwareSourcesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSoftwareSourcesSortByEnum Enum with underlying type: string
type ListSoftwareSourcesSortByEnum string

// Set of constants representing the allowable values for ListSoftwareSourcesSortByEnum
const (
	ListSoftwareSourcesSortByTimecreated ListSoftwareSourcesSortByEnum = "TIMECREATED"
	ListSoftwareSourcesSortByDisplayname ListSoftwareSourcesSortByEnum = "DISPLAYNAME"
)

var mappingListSoftwareSourcesSortBy = map[string]ListSoftwareSourcesSortByEnum{
	"TIMECREATED": ListSoftwareSourcesSortByTimecreated,
	"DISPLAYNAME": ListSoftwareSourcesSortByDisplayname,
}

// GetListSoftwareSourcesSortByEnumValues Enumerates the set of values for ListSoftwareSourcesSortByEnum
func GetListSoftwareSourcesSortByEnumValues() []ListSoftwareSourcesSortByEnum {
	values := make([]ListSoftwareSourcesSortByEnum, 0)
	for _, v := range mappingListSoftwareSourcesSortBy {
		values = append(values, v)
	}
	return values
}

// ListSoftwareSourcesLifecycleStateEnum Enum with underlying type: string
type ListSoftwareSourcesLifecycleStateEnum string

// Set of constants representing the allowable values for ListSoftwareSourcesLifecycleStateEnum
const (
	ListSoftwareSourcesLifecycleStateCreating ListSoftwareSourcesLifecycleStateEnum = "CREATING"
	ListSoftwareSourcesLifecycleStateUpdating ListSoftwareSourcesLifecycleStateEnum = "UPDATING"
	ListSoftwareSourcesLifecycleStateActive   ListSoftwareSourcesLifecycleStateEnum = "ACTIVE"
	ListSoftwareSourcesLifecycleStateDeleting ListSoftwareSourcesLifecycleStateEnum = "DELETING"
	ListSoftwareSourcesLifecycleStateDeleted  ListSoftwareSourcesLifecycleStateEnum = "DELETED"
	ListSoftwareSourcesLifecycleStateFailed   ListSoftwareSourcesLifecycleStateEnum = "FAILED"
)

var mappingListSoftwareSourcesLifecycleState = map[string]ListSoftwareSourcesLifecycleStateEnum{
	"CREATING": ListSoftwareSourcesLifecycleStateCreating,
	"UPDATING": ListSoftwareSourcesLifecycleStateUpdating,
	"ACTIVE":   ListSoftwareSourcesLifecycleStateActive,
	"DELETING": ListSoftwareSourcesLifecycleStateDeleting,
	"DELETED":  ListSoftwareSourcesLifecycleStateDeleted,
	"FAILED":   ListSoftwareSourcesLifecycleStateFailed,
}

// GetListSoftwareSourcesLifecycleStateEnumValues Enumerates the set of values for ListSoftwareSourcesLifecycleStateEnum
func GetListSoftwareSourcesLifecycleStateEnumValues() []ListSoftwareSourcesLifecycleStateEnum {
	values := make([]ListSoftwareSourcesLifecycleStateEnum, 0)
	for _, v := range mappingListSoftwareSourcesLifecycleState {
		values = append(values, v)
	}
	return values
}
