// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osmanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListManagedInstancesRequest wrapper for the ListManagedInstances operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osmanagement/ListManagedInstances.go.html to see an example of how to use ListManagedInstancesRequest.
type ListManagedInstancesRequest struct {

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
	SortOrder ListManagedInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListManagedInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OS family for which to list resources.
	OsFamily ListManagedInstancesOsFamilyEnum `mandatory:"false" contributesTo:"query" name:"osFamily" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedInstancesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedInstancesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListManagedInstancesResponse wrapper for the ListManagedInstances operation
type ListManagedInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ManagedInstanceSummary instances
	Items []ManagedInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this
	// header appears in the response, then a partial list might have been
	// returned. Include this value as the `page` parameter for the subsequent
	// GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedInstancesSortOrderEnum Enum with underlying type: string
type ListManagedInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedInstancesSortOrderEnum
const (
	ListManagedInstancesSortOrderAsc  ListManagedInstancesSortOrderEnum = "ASC"
	ListManagedInstancesSortOrderDesc ListManagedInstancesSortOrderEnum = "DESC"
)

var mappingListManagedInstancesSortOrder = map[string]ListManagedInstancesSortOrderEnum{
	"ASC":  ListManagedInstancesSortOrderAsc,
	"DESC": ListManagedInstancesSortOrderDesc,
}

// GetListManagedInstancesSortOrderEnumValues Enumerates the set of values for ListManagedInstancesSortOrderEnum
func GetListManagedInstancesSortOrderEnumValues() []ListManagedInstancesSortOrderEnum {
	values := make([]ListManagedInstancesSortOrderEnum, 0)
	for _, v := range mappingListManagedInstancesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListManagedInstancesSortByEnum Enum with underlying type: string
type ListManagedInstancesSortByEnum string

// Set of constants representing the allowable values for ListManagedInstancesSortByEnum
const (
	ListManagedInstancesSortByTimecreated ListManagedInstancesSortByEnum = "TIMECREATED"
	ListManagedInstancesSortByDisplayname ListManagedInstancesSortByEnum = "DISPLAYNAME"
)

var mappingListManagedInstancesSortBy = map[string]ListManagedInstancesSortByEnum{
	"TIMECREATED": ListManagedInstancesSortByTimecreated,
	"DISPLAYNAME": ListManagedInstancesSortByDisplayname,
}

// GetListManagedInstancesSortByEnumValues Enumerates the set of values for ListManagedInstancesSortByEnum
func GetListManagedInstancesSortByEnumValues() []ListManagedInstancesSortByEnum {
	values := make([]ListManagedInstancesSortByEnum, 0)
	for _, v := range mappingListManagedInstancesSortBy {
		values = append(values, v)
	}
	return values
}

// ListManagedInstancesOsFamilyEnum Enum with underlying type: string
type ListManagedInstancesOsFamilyEnum string

// Set of constants representing the allowable values for ListManagedInstancesOsFamilyEnum
const (
	ListManagedInstancesOsFamilyLinux   ListManagedInstancesOsFamilyEnum = "LINUX"
	ListManagedInstancesOsFamilyWindows ListManagedInstancesOsFamilyEnum = "WINDOWS"
	ListManagedInstancesOsFamilyAll     ListManagedInstancesOsFamilyEnum = "ALL"
)

var mappingListManagedInstancesOsFamily = map[string]ListManagedInstancesOsFamilyEnum{
	"LINUX":   ListManagedInstancesOsFamilyLinux,
	"WINDOWS": ListManagedInstancesOsFamilyWindows,
	"ALL":     ListManagedInstancesOsFamilyAll,
}

// GetListManagedInstancesOsFamilyEnumValues Enumerates the set of values for ListManagedInstancesOsFamilyEnum
func GetListManagedInstancesOsFamilyEnumValues() []ListManagedInstancesOsFamilyEnum {
	values := make([]ListManagedInstancesOsFamilyEnum, 0)
	for _, v := range mappingListManagedInstancesOsFamily {
		values = append(values, v)
	}
	return values
}
