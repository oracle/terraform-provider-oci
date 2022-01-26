// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package appmgmtcontrol

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListMonitoredInstancesRequest wrapper for the ListMonitoredInstances operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/appmgmtcontrol/ListMonitoredInstances.go.html to see an example of how to use ListMonitoredInstancesRequest.
type ListMonitoredInstancesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending ('ASC') or descending ('DESC').
	SortOrder ListMonitoredInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListMonitoredInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMonitoredInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMonitoredInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMonitoredInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMonitoredInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListMonitoredInstancesResponse wrapper for the ListMonitoredInstances operation
type ListMonitoredInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MonitoredInstanceCollection instances
	MonitoredInstanceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMonitoredInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMonitoredInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMonitoredInstancesSortOrderEnum Enum with underlying type: string
type ListMonitoredInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListMonitoredInstancesSortOrderEnum
const (
	ListMonitoredInstancesSortOrderAsc  ListMonitoredInstancesSortOrderEnum = "ASC"
	ListMonitoredInstancesSortOrderDesc ListMonitoredInstancesSortOrderEnum = "DESC"
)

var mappingListMonitoredInstancesSortOrder = map[string]ListMonitoredInstancesSortOrderEnum{
	"ASC":  ListMonitoredInstancesSortOrderAsc,
	"DESC": ListMonitoredInstancesSortOrderDesc,
}

// GetListMonitoredInstancesSortOrderEnumValues Enumerates the set of values for ListMonitoredInstancesSortOrderEnum
func GetListMonitoredInstancesSortOrderEnumValues() []ListMonitoredInstancesSortOrderEnum {
	values := make([]ListMonitoredInstancesSortOrderEnum, 0)
	for _, v := range mappingListMonitoredInstancesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListMonitoredInstancesSortByEnum Enum with underlying type: string
type ListMonitoredInstancesSortByEnum string

// Set of constants representing the allowable values for ListMonitoredInstancesSortByEnum
const (
	ListMonitoredInstancesSortByTimecreated ListMonitoredInstancesSortByEnum = "timeCreated"
	ListMonitoredInstancesSortByDisplayname ListMonitoredInstancesSortByEnum = "displayName"
)

var mappingListMonitoredInstancesSortBy = map[string]ListMonitoredInstancesSortByEnum{
	"timeCreated": ListMonitoredInstancesSortByTimecreated,
	"displayName": ListMonitoredInstancesSortByDisplayname,
}

// GetListMonitoredInstancesSortByEnumValues Enumerates the set of values for ListMonitoredInstancesSortByEnum
func GetListMonitoredInstancesSortByEnumValues() []ListMonitoredInstancesSortByEnum {
	values := make([]ListMonitoredInstancesSortByEnum, 0)
	for _, v := range mappingListMonitoredInstancesSortBy {
		values = append(values, v)
	}
	return values
}
