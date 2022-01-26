// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListTablespacesRequest wrapper for the ListTablespaces operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListTablespaces.go.html to see an example of how to use ListTablespacesRequest.
type ListTablespacesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListTablespacesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListTablespacesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTablespacesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTablespacesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTablespacesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTablespacesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTablespacesResponse wrapper for the ListTablespaces operation
type ListTablespacesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TablespaceCollection instances
	TablespaceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTablespacesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTablespacesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTablespacesSortByEnum Enum with underlying type: string
type ListTablespacesSortByEnum string

// Set of constants representing the allowable values for ListTablespacesSortByEnum
const (
	ListTablespacesSortByTimecreated ListTablespacesSortByEnum = "TIMECREATED"
	ListTablespacesSortByName        ListTablespacesSortByEnum = "NAME"
)

var mappingListTablespacesSortBy = map[string]ListTablespacesSortByEnum{
	"TIMECREATED": ListTablespacesSortByTimecreated,
	"NAME":        ListTablespacesSortByName,
}

// GetListTablespacesSortByEnumValues Enumerates the set of values for ListTablespacesSortByEnum
func GetListTablespacesSortByEnumValues() []ListTablespacesSortByEnum {
	values := make([]ListTablespacesSortByEnum, 0)
	for _, v := range mappingListTablespacesSortBy {
		values = append(values, v)
	}
	return values
}

// ListTablespacesSortOrderEnum Enum with underlying type: string
type ListTablespacesSortOrderEnum string

// Set of constants representing the allowable values for ListTablespacesSortOrderEnum
const (
	ListTablespacesSortOrderAsc  ListTablespacesSortOrderEnum = "ASC"
	ListTablespacesSortOrderDesc ListTablespacesSortOrderEnum = "DESC"
)

var mappingListTablespacesSortOrder = map[string]ListTablespacesSortOrderEnum{
	"ASC":  ListTablespacesSortOrderAsc,
	"DESC": ListTablespacesSortOrderDesc,
}

// GetListTablespacesSortOrderEnumValues Enumerates the set of values for ListTablespacesSortOrderEnum
func GetListTablespacesSortOrderEnumValues() []ListTablespacesSortOrderEnum {
	values := make([]ListTablespacesSortOrderEnum, 0)
	for _, v := range mappingListTablespacesSortOrder {
		values = append(values, v)
	}
	return values
}
