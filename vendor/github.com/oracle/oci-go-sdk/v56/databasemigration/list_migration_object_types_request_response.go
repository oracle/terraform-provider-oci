// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListMigrationObjectTypesRequest wrapper for the ListMigrationObjectTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListMigrationObjectTypes.go.html to see an example of how to use ListMigrationObjectTypesRequest.
type ListMigrationObjectTypesRequest struct {

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for name is custom based on it's usage frequency. If no value is specified name is default.
	SortBy ListMigrationObjectTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMigrationObjectTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMigrationObjectTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMigrationObjectTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMigrationObjectTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMigrationObjectTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListMigrationObjectTypesResponse wrapper for the ListMigrationObjectTypes operation
type ListMigrationObjectTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MigrationObjectTypeSummaryCollection instances
	MigrationObjectTypeSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMigrationObjectTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMigrationObjectTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMigrationObjectTypesSortByEnum Enum with underlying type: string
type ListMigrationObjectTypesSortByEnum string

// Set of constants representing the allowable values for ListMigrationObjectTypesSortByEnum
const (
	ListMigrationObjectTypesSortByName ListMigrationObjectTypesSortByEnum = "name"
)

var mappingListMigrationObjectTypesSortBy = map[string]ListMigrationObjectTypesSortByEnum{
	"name": ListMigrationObjectTypesSortByName,
}

// GetListMigrationObjectTypesSortByEnumValues Enumerates the set of values for ListMigrationObjectTypesSortByEnum
func GetListMigrationObjectTypesSortByEnumValues() []ListMigrationObjectTypesSortByEnum {
	values := make([]ListMigrationObjectTypesSortByEnum, 0)
	for _, v := range mappingListMigrationObjectTypesSortBy {
		values = append(values, v)
	}
	return values
}

// ListMigrationObjectTypesSortOrderEnum Enum with underlying type: string
type ListMigrationObjectTypesSortOrderEnum string

// Set of constants representing the allowable values for ListMigrationObjectTypesSortOrderEnum
const (
	ListMigrationObjectTypesSortOrderAsc  ListMigrationObjectTypesSortOrderEnum = "ASC"
	ListMigrationObjectTypesSortOrderDesc ListMigrationObjectTypesSortOrderEnum = "DESC"
)

var mappingListMigrationObjectTypesSortOrder = map[string]ListMigrationObjectTypesSortOrderEnum{
	"ASC":  ListMigrationObjectTypesSortOrderAsc,
	"DESC": ListMigrationObjectTypesSortOrderDesc,
}

// GetListMigrationObjectTypesSortOrderEnumValues Enumerates the set of values for ListMigrationObjectTypesSortOrderEnum
func GetListMigrationObjectTypesSortOrderEnumValues() []ListMigrationObjectTypesSortOrderEnum {
	values := make([]ListMigrationObjectTypesSortOrderEnum, 0)
	for _, v := range mappingListMigrationObjectTypesSortOrder {
		values = append(values, v)
	}
	return values
}
