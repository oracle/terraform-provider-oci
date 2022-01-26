// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListUsersRequest wrapper for the ListUsers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListUsers.go.html to see an example of how to use ListUsersRequest.
type ListUsersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListUsersResponse wrapper for the ListUsers operation
type ListUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UserCollection instances
	UserCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUsersSortByEnum Enum with underlying type: string
type ListUsersSortByEnum string

// Set of constants representing the allowable values for ListUsersSortByEnum
const (
	ListUsersSortByTimecreated ListUsersSortByEnum = "TIMECREATED"
	ListUsersSortByName        ListUsersSortByEnum = "NAME"
)

var mappingListUsersSortBy = map[string]ListUsersSortByEnum{
	"TIMECREATED": ListUsersSortByTimecreated,
	"NAME":        ListUsersSortByName,
}

// GetListUsersSortByEnumValues Enumerates the set of values for ListUsersSortByEnum
func GetListUsersSortByEnumValues() []ListUsersSortByEnum {
	values := make([]ListUsersSortByEnum, 0)
	for _, v := range mappingListUsersSortBy {
		values = append(values, v)
	}
	return values
}

// ListUsersSortOrderEnum Enum with underlying type: string
type ListUsersSortOrderEnum string

// Set of constants representing the allowable values for ListUsersSortOrderEnum
const (
	ListUsersSortOrderAsc  ListUsersSortOrderEnum = "ASC"
	ListUsersSortOrderDesc ListUsersSortOrderEnum = "DESC"
)

var mappingListUsersSortOrder = map[string]ListUsersSortOrderEnum{
	"ASC":  ListUsersSortOrderAsc,
	"DESC": ListUsersSortOrderDesc,
}

// GetListUsersSortOrderEnumValues Enumerates the set of values for ListUsersSortOrderEnum
func GetListUsersSortOrderEnumValues() []ListUsersSortOrderEnum {
	values := make([]ListUsersSortOrderEnum, 0)
	for _, v := range mappingListUsersSortOrder {
		values = append(values, v)
	}
	return values
}
