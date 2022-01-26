// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListObjectPrivilegesRequest wrapper for the ListObjectPrivileges operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListObjectPrivileges.go.html to see an example of how to use ListObjectPrivilegesRequest.
type ListObjectPrivilegesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The name of the user whose details are to be viewed.
	UserName *string `mandatory:"true" contributesTo:"path" name:"userName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘NAME’ is ascending. The ‘NAME’ sort order is case-sensitive.
	SortBy ListObjectPrivilegesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListObjectPrivilegesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListObjectPrivilegesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListObjectPrivilegesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListObjectPrivilegesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListObjectPrivilegesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListObjectPrivilegesResponse wrapper for the ListObjectPrivileges operation
type ListObjectPrivilegesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ObjectPrivilegeCollection instances
	ObjectPrivilegeCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListObjectPrivilegesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListObjectPrivilegesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListObjectPrivilegesSortByEnum Enum with underlying type: string
type ListObjectPrivilegesSortByEnum string

// Set of constants representing the allowable values for ListObjectPrivilegesSortByEnum
const (
	ListObjectPrivilegesSortByName ListObjectPrivilegesSortByEnum = "NAME"
)

var mappingListObjectPrivilegesSortBy = map[string]ListObjectPrivilegesSortByEnum{
	"NAME": ListObjectPrivilegesSortByName,
}

// GetListObjectPrivilegesSortByEnumValues Enumerates the set of values for ListObjectPrivilegesSortByEnum
func GetListObjectPrivilegesSortByEnumValues() []ListObjectPrivilegesSortByEnum {
	values := make([]ListObjectPrivilegesSortByEnum, 0)
	for _, v := range mappingListObjectPrivilegesSortBy {
		values = append(values, v)
	}
	return values
}

// ListObjectPrivilegesSortOrderEnum Enum with underlying type: string
type ListObjectPrivilegesSortOrderEnum string

// Set of constants representing the allowable values for ListObjectPrivilegesSortOrderEnum
const (
	ListObjectPrivilegesSortOrderAsc  ListObjectPrivilegesSortOrderEnum = "ASC"
	ListObjectPrivilegesSortOrderDesc ListObjectPrivilegesSortOrderEnum = "DESC"
)

var mappingListObjectPrivilegesSortOrder = map[string]ListObjectPrivilegesSortOrderEnum{
	"ASC":  ListObjectPrivilegesSortOrderAsc,
	"DESC": ListObjectPrivilegesSortOrderDesc,
}

// GetListObjectPrivilegesSortOrderEnumValues Enumerates the set of values for ListObjectPrivilegesSortOrderEnum
func GetListObjectPrivilegesSortOrderEnumValues() []ListObjectPrivilegesSortOrderEnum {
	values := make([]ListObjectPrivilegesSortOrderEnum, 0)
	for _, v := range mappingListObjectPrivilegesSortOrder {
		values = append(values, v)
	}
	return values
}
