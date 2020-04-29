// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAvailableSoftwareSourcesForManagedInstanceRequest wrapper for the ListAvailableSoftwareSourcesForManagedInstance operation
type ListAvailableSoftwareSourcesForManagedInstanceRequest struct {

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
	SortOrder ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListAvailableSoftwareSourcesForManagedInstanceSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAvailableSoftwareSourcesForManagedInstanceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAvailableSoftwareSourcesForManagedInstanceRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAvailableSoftwareSourcesForManagedInstanceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAvailableSoftwareSourcesForManagedInstanceResponse wrapper for the ListAvailableSoftwareSourcesForManagedInstance operation
type ListAvailableSoftwareSourcesForManagedInstanceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AvailableSoftwareSourceSummary instances
	Items []AvailableSoftwareSourceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `AvailableSoftwareSourceSummary`s. If
	// this header appears in the response, then this is a partial
	// list of `AvailableSoftwareSourceSummmary'`s for the managed instance.
	// Include this value as the `page` parameter in a subsequent
	// GET request to get the next batch of managed instances.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAvailableSoftwareSourcesForManagedInstanceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAvailableSoftwareSourcesForManagedInstanceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum Enum with underlying type: string
type ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum string

// Set of constants representing the allowable values for ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum
const (
	ListAvailableSoftwareSourcesForManagedInstanceSortOrderAsc  ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum = "ASC"
	ListAvailableSoftwareSourcesForManagedInstanceSortOrderDesc ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum = "DESC"
)

var mappingListAvailableSoftwareSourcesForManagedInstanceSortOrder = map[string]ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum{
	"ASC":  ListAvailableSoftwareSourcesForManagedInstanceSortOrderAsc,
	"DESC": ListAvailableSoftwareSourcesForManagedInstanceSortOrderDesc,
}

// GetListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumValues Enumerates the set of values for ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum
func GetListAvailableSoftwareSourcesForManagedInstanceSortOrderEnumValues() []ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum {
	values := make([]ListAvailableSoftwareSourcesForManagedInstanceSortOrderEnum, 0)
	for _, v := range mappingListAvailableSoftwareSourcesForManagedInstanceSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAvailableSoftwareSourcesForManagedInstanceSortByEnum Enum with underlying type: string
type ListAvailableSoftwareSourcesForManagedInstanceSortByEnum string

// Set of constants representing the allowable values for ListAvailableSoftwareSourcesForManagedInstanceSortByEnum
const (
	ListAvailableSoftwareSourcesForManagedInstanceSortByTimecreated ListAvailableSoftwareSourcesForManagedInstanceSortByEnum = "TIMECREATED"
	ListAvailableSoftwareSourcesForManagedInstanceSortByDisplayname ListAvailableSoftwareSourcesForManagedInstanceSortByEnum = "DISPLAYNAME"
)

var mappingListAvailableSoftwareSourcesForManagedInstanceSortBy = map[string]ListAvailableSoftwareSourcesForManagedInstanceSortByEnum{
	"TIMECREATED": ListAvailableSoftwareSourcesForManagedInstanceSortByTimecreated,
	"DISPLAYNAME": ListAvailableSoftwareSourcesForManagedInstanceSortByDisplayname,
}

// GetListAvailableSoftwareSourcesForManagedInstanceSortByEnumValues Enumerates the set of values for ListAvailableSoftwareSourcesForManagedInstanceSortByEnum
func GetListAvailableSoftwareSourcesForManagedInstanceSortByEnumValues() []ListAvailableSoftwareSourcesForManagedInstanceSortByEnum {
	values := make([]ListAvailableSoftwareSourcesForManagedInstanceSortByEnum, 0)
	for _, v := range mappingListAvailableSoftwareSourcesForManagedInstanceSortBy {
		values = append(values, v)
	}
	return values
}
