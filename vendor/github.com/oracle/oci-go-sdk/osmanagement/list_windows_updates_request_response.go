// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListWindowsUpdatesRequest wrapper for the ListWindowsUpdates operation
type ListWindowsUpdatesRequest struct {

	// The ID of the compartment in which to list resources. This parameter is optional and in some cases may have no effect.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListWindowsUpdatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListWindowsUpdatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWindowsUpdatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWindowsUpdatesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWindowsUpdatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListWindowsUpdatesResponse wrapper for the ListWindowsUpdates operation
type ListWindowsUpdatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []WindowsUpdateSummary instances
	Items []WindowsUpdateSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of `WindowdsUpdateSummary`s. If this header
	// appears in the response, then this is a partial list of managed
	// instances. Include this value as the `page` parameter in a
	// subsequent GET request to get the next batch of managed
	// instances.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWindowsUpdatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWindowsUpdatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWindowsUpdatesSortOrderEnum Enum with underlying type: string
type ListWindowsUpdatesSortOrderEnum string

// Set of constants representing the allowable values for ListWindowsUpdatesSortOrderEnum
const (
	ListWindowsUpdatesSortOrderAsc  ListWindowsUpdatesSortOrderEnum = "ASC"
	ListWindowsUpdatesSortOrderDesc ListWindowsUpdatesSortOrderEnum = "DESC"
)

var mappingListWindowsUpdatesSortOrder = map[string]ListWindowsUpdatesSortOrderEnum{
	"ASC":  ListWindowsUpdatesSortOrderAsc,
	"DESC": ListWindowsUpdatesSortOrderDesc,
}

// GetListWindowsUpdatesSortOrderEnumValues Enumerates the set of values for ListWindowsUpdatesSortOrderEnum
func GetListWindowsUpdatesSortOrderEnumValues() []ListWindowsUpdatesSortOrderEnum {
	values := make([]ListWindowsUpdatesSortOrderEnum, 0)
	for _, v := range mappingListWindowsUpdatesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListWindowsUpdatesSortByEnum Enum with underlying type: string
type ListWindowsUpdatesSortByEnum string

// Set of constants representing the allowable values for ListWindowsUpdatesSortByEnum
const (
	ListWindowsUpdatesSortByTimecreated ListWindowsUpdatesSortByEnum = "TIMECREATED"
	ListWindowsUpdatesSortByDisplayname ListWindowsUpdatesSortByEnum = "DISPLAYNAME"
)

var mappingListWindowsUpdatesSortBy = map[string]ListWindowsUpdatesSortByEnum{
	"TIMECREATED": ListWindowsUpdatesSortByTimecreated,
	"DISPLAYNAME": ListWindowsUpdatesSortByDisplayname,
}

// GetListWindowsUpdatesSortByEnumValues Enumerates the set of values for ListWindowsUpdatesSortByEnum
func GetListWindowsUpdatesSortByEnumValues() []ListWindowsUpdatesSortByEnum {
	values := make([]ListWindowsUpdatesSortByEnum, 0)
	for _, v := range mappingListWindowsUpdatesSortBy {
		values = append(values, v)
	}
	return values
}
