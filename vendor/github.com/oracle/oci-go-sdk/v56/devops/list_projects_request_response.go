// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListProjectsRequest wrapper for the ListProjects operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListProjects.go.html to see an example of how to use ListProjectsRequest.
type ListProjectsRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier or OCID for listing a single resource by ID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only Projects that matches the given lifecycleState.
	LifecycleState ProjectLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListProjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for display name is ascending. If no value is specified, then the default time created value is considered.
	SortBy ListProjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListProjectsResponse wrapper for the ListProjects operation
type ListProjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProjectCollection instances
	ProjectCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProjectsSortOrderEnum Enum with underlying type: string
type ListProjectsSortOrderEnum string

// Set of constants representing the allowable values for ListProjectsSortOrderEnum
const (
	ListProjectsSortOrderAsc  ListProjectsSortOrderEnum = "ASC"
	ListProjectsSortOrderDesc ListProjectsSortOrderEnum = "DESC"
)

var mappingListProjectsSortOrder = map[string]ListProjectsSortOrderEnum{
	"ASC":  ListProjectsSortOrderAsc,
	"DESC": ListProjectsSortOrderDesc,
}

// GetListProjectsSortOrderEnumValues Enumerates the set of values for ListProjectsSortOrderEnum
func GetListProjectsSortOrderEnumValues() []ListProjectsSortOrderEnum {
	values := make([]ListProjectsSortOrderEnum, 0)
	for _, v := range mappingListProjectsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListProjectsSortByEnum Enum with underlying type: string
type ListProjectsSortByEnum string

// Set of constants representing the allowable values for ListProjectsSortByEnum
const (
	ListProjectsSortByTimecreated ListProjectsSortByEnum = "timeCreated"
	ListProjectsSortByDisplayname ListProjectsSortByEnum = "displayName"
)

var mappingListProjectsSortBy = map[string]ListProjectsSortByEnum{
	"timeCreated": ListProjectsSortByTimecreated,
	"displayName": ListProjectsSortByDisplayname,
}

// GetListProjectsSortByEnumValues Enumerates the set of values for ListProjectsSortByEnum
func GetListProjectsSortByEnumValues() []ListProjectsSortByEnum {
	values := make([]ListProjectsSortByEnum, 0)
	for _, v := range mappingListProjectsSortBy {
		values = append(values, v)
	}
	return values
}
