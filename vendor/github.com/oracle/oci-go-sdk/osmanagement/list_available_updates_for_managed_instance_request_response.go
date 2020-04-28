// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAvailableUpdatesForManagedInstanceRequest wrapper for the ListAvailableUpdatesForManagedInstance operation
type ListAvailableUpdatesForManagedInstanceRequest struct {

	// OCID for the managed instance
	ManagedInstanceId *string `mandatory:"true" contributesTo:"path" name:"managedInstanceId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAvailableUpdatesForManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListAvailableUpdatesForManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableUpdatesForManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableUpdatesForManagedInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableUpdatesForManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAvailableUpdatesForManagedInstanceResponse wrapper for the ListAvailableUpdatesForManagedInstance operation
type ListAvailableUpdatesForManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AvailableUpdateSummary instances
	Items []AvailableUpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `UpdatablePackage`s. If this header
	// appears in the response, then this is a partial list of
	// `UpdatablePackage`s for the managed instance. Include this value
	// as the `page` parameter in a subsequent
	// GET request to get the next batch of managed instances.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableUpdatesForManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableUpdatesForManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableUpdatesForManagedInstanceSortOrderEnum Enum with underlying type: string
type ListAvailableUpdatesForManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableUpdatesForManagedInstanceSortOrderEnum
const (
	ListAvailableUpdatesForManagedInstanceSortOrderAsc  ListAvailableUpdatesForManagedInstanceSortOrderEnum = "ASC"
	ListAvailableUpdatesForManagedInstanceSortOrderDesc ListAvailableUpdatesForManagedInstanceSortOrderEnum = "DESC"
)

var mappingListAvailableUpdatesForManagedInstanceSortOrder = map[string]ListAvailableUpdatesForManagedInstanceSortOrderEnum{
	"ASC":  ListAvailableUpdatesForManagedInstanceSortOrderAsc,
	"DESC": ListAvailableUpdatesForManagedInstanceSortOrderDesc,
}

// GetListAvailableUpdatesForManagedInstanceSortOrderEnumValues Enumerates the set of values for ListAvailableUpdatesForManagedInstanceSortOrderEnum
func GetListAvailableUpdatesForManagedInstanceSortOrderEnumValues() []ListAvailableUpdatesForManagedInstanceSortOrderEnum {
	values := make([]ListAvailableUpdatesForManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListAvailableUpdatesForManagedInstanceSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAvailableUpdatesForManagedInstanceSortByEnum Enum with underlying type: string
type ListAvailableUpdatesForManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListAvailableUpdatesForManagedInstanceSortByEnum
const (
	ListAvailableUpdatesForManagedInstanceSortByTimecreated ListAvailableUpdatesForManagedInstanceSortByEnum = "TIMECREATED"
	ListAvailableUpdatesForManagedInstanceSortByDisplayname ListAvailableUpdatesForManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListAvailableUpdatesForManagedInstanceSortBy = map[string]ListAvailableUpdatesForManagedInstanceSortByEnum{
	"TIMECREATED": ListAvailableUpdatesForManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListAvailableUpdatesForManagedInstanceSortByDisplayname,
}

// GetListAvailableUpdatesForManagedInstanceSortByEnumValues Enumerates the set of values for ListAvailableUpdatesForManagedInstanceSortByEnum
func GetListAvailableUpdatesForManagedInstanceSortByEnumValues() []ListAvailableUpdatesForManagedInstanceSortByEnum {
	values := make([]ListAvailableUpdatesForManagedInstanceSortByEnum, 0)
	for _, v := range mappingListAvailableUpdatesForManagedInstanceSortBy {
		values = append(values, v)
	}
	return values
}
