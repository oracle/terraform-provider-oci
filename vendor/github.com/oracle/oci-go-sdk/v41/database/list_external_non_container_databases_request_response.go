// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// ListExternalNonContainerDatabasesRequest wrapper for the ListExternalNonContainerDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExternalNonContainerDatabases.go.html to see an example of how to use ListExternalNonContainerDatabasesRequest.
type ListExternalNonContainerDatabasesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for TIMECREATED is descending.
	// Default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListExternalNonContainerDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExternalNonContainerDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle state.
	LifecycleState ExternalDatabaseBaseLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalNonContainerDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalNonContainerDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalNonContainerDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalNonContainerDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListExternalNonContainerDatabasesResponse wrapper for the ListExternalNonContainerDatabases operation
type ListExternalNonContainerDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExternalNonContainerDatabaseSummary instances
	Items []ExternalNonContainerDatabaseSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalNonContainerDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalNonContainerDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalNonContainerDatabasesSortByEnum Enum with underlying type: string
type ListExternalNonContainerDatabasesSortByEnum string

// Set of constants representing the allowable values for ListExternalNonContainerDatabasesSortByEnum
const (
	ListExternalNonContainerDatabasesSortByDisplayname ListExternalNonContainerDatabasesSortByEnum = "DISPLAYNAME"
	ListExternalNonContainerDatabasesSortByTimecreated ListExternalNonContainerDatabasesSortByEnum = "TIMECREATED"
)

var mappingListExternalNonContainerDatabasesSortBy = map[string]ListExternalNonContainerDatabasesSortByEnum{
	"DISPLAYNAME": ListExternalNonContainerDatabasesSortByDisplayname,
	"TIMECREATED": ListExternalNonContainerDatabasesSortByTimecreated,
}

// GetListExternalNonContainerDatabasesSortByEnumValues Enumerates the set of values for ListExternalNonContainerDatabasesSortByEnum
func GetListExternalNonContainerDatabasesSortByEnumValues() []ListExternalNonContainerDatabasesSortByEnum {
	values := make([]ListExternalNonContainerDatabasesSortByEnum, 0)
	for _, v := range mappingListExternalNonContainerDatabasesSortBy {
		values = append(values, v)
	}
	return values
}

// ListExternalNonContainerDatabasesSortOrderEnum Enum with underlying type: string
type ListExternalNonContainerDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalNonContainerDatabasesSortOrderEnum
const (
	ListExternalNonContainerDatabasesSortOrderAsc  ListExternalNonContainerDatabasesSortOrderEnum = "ASC"
	ListExternalNonContainerDatabasesSortOrderDesc ListExternalNonContainerDatabasesSortOrderEnum = "DESC"
)

var mappingListExternalNonContainerDatabasesSortOrder = map[string]ListExternalNonContainerDatabasesSortOrderEnum{
	"ASC":  ListExternalNonContainerDatabasesSortOrderAsc,
	"DESC": ListExternalNonContainerDatabasesSortOrderDesc,
}

// GetListExternalNonContainerDatabasesSortOrderEnumValues Enumerates the set of values for ListExternalNonContainerDatabasesSortOrderEnum
func GetListExternalNonContainerDatabasesSortOrderEnumValues() []ListExternalNonContainerDatabasesSortOrderEnum {
	values := make([]ListExternalNonContainerDatabasesSortOrderEnum, 0)
	for _, v := range mappingListExternalNonContainerDatabasesSortOrder {
		values = append(values, v)
	}
	return values
}
