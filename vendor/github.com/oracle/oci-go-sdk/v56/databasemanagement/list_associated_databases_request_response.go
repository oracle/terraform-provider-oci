// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListAssociatedDatabasesRequest wrapper for the ListAssociatedDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAssociatedDatabases.go.html to see an example of how to use ListAssociatedDatabasesRequest.
type ListAssociatedDatabasesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint.
	DbManagementPrivateEndpointId *string `mandatory:"true" contributesTo:"path" name:"dbManagementPrivateEndpointId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListAssociatedDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The option to sort databases using a specific Database Management private endpoint.
	SortBy ListAssociatedDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAssociatedDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAssociatedDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAssociatedDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAssociatedDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAssociatedDatabasesResponse wrapper for the ListAssociatedDatabases operation
type ListAssociatedDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssociatedDatabaseCollection instances
	AssociatedDatabaseCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAssociatedDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAssociatedDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAssociatedDatabasesSortOrderEnum Enum with underlying type: string
type ListAssociatedDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListAssociatedDatabasesSortOrderEnum
const (
	ListAssociatedDatabasesSortOrderAsc  ListAssociatedDatabasesSortOrderEnum = "ASC"
	ListAssociatedDatabasesSortOrderDesc ListAssociatedDatabasesSortOrderEnum = "DESC"
)

var mappingListAssociatedDatabasesSortOrder = map[string]ListAssociatedDatabasesSortOrderEnum{
	"ASC":  ListAssociatedDatabasesSortOrderAsc,
	"DESC": ListAssociatedDatabasesSortOrderDesc,
}

// GetListAssociatedDatabasesSortOrderEnumValues Enumerates the set of values for ListAssociatedDatabasesSortOrderEnum
func GetListAssociatedDatabasesSortOrderEnumValues() []ListAssociatedDatabasesSortOrderEnum {
	values := make([]ListAssociatedDatabasesSortOrderEnum, 0)
	for _, v := range mappingListAssociatedDatabasesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListAssociatedDatabasesSortByEnum Enum with underlying type: string
type ListAssociatedDatabasesSortByEnum string

// Set of constants representing the allowable values for ListAssociatedDatabasesSortByEnum
const (
	ListAssociatedDatabasesSortByTimeregistered ListAssociatedDatabasesSortByEnum = "timeRegistered"
)

var mappingListAssociatedDatabasesSortBy = map[string]ListAssociatedDatabasesSortByEnum{
	"timeRegistered": ListAssociatedDatabasesSortByTimeregistered,
}

// GetListAssociatedDatabasesSortByEnumValues Enumerates the set of values for ListAssociatedDatabasesSortByEnum
func GetListAssociatedDatabasesSortByEnumValues() []ListAssociatedDatabasesSortByEnum {
	values := make([]ListAssociatedDatabasesSortByEnum, 0)
	for _, v := range mappingListAssociatedDatabasesSortBy {
		values = append(values, v)
	}
	return values
}
