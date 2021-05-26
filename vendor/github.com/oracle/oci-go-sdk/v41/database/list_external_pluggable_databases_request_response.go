// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// ListExternalPluggableDatabasesRequest wrapper for the ListExternalPluggableDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListExternalPluggableDatabases.go.html to see an example of how to use ListExternalPluggableDatabasesRequest.
type ListExternalPluggableDatabasesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The ExternalContainerDatabase OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	ExternalContainerDatabaseId *string `mandatory:"false" contributesTo:"query" name:"externalContainerDatabaseId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for TIMECREATED is descending.
	// Default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListExternalPluggableDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExternalPluggableDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the specified lifecycle state.
	LifecycleState ExternalDatabaseBaseLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalPluggableDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalPluggableDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalPluggableDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalPluggableDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListExternalPluggableDatabasesResponse wrapper for the ListExternalPluggableDatabases operation
type ListExternalPluggableDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExternalPluggableDatabaseSummary instances
	Items []ExternalPluggableDatabaseSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalPluggableDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalPluggableDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalPluggableDatabasesSortByEnum Enum with underlying type: string
type ListExternalPluggableDatabasesSortByEnum string

// Set of constants representing the allowable values for ListExternalPluggableDatabasesSortByEnum
const (
	ListExternalPluggableDatabasesSortByDisplayname ListExternalPluggableDatabasesSortByEnum = "DISPLAYNAME"
	ListExternalPluggableDatabasesSortByTimecreated ListExternalPluggableDatabasesSortByEnum = "TIMECREATED"
)

var mappingListExternalPluggableDatabasesSortBy = map[string]ListExternalPluggableDatabasesSortByEnum{
	"DISPLAYNAME": ListExternalPluggableDatabasesSortByDisplayname,
	"TIMECREATED": ListExternalPluggableDatabasesSortByTimecreated,
}

// GetListExternalPluggableDatabasesSortByEnumValues Enumerates the set of values for ListExternalPluggableDatabasesSortByEnum
func GetListExternalPluggableDatabasesSortByEnumValues() []ListExternalPluggableDatabasesSortByEnum {
	values := make([]ListExternalPluggableDatabasesSortByEnum, 0)
	for _, v := range mappingListExternalPluggableDatabasesSortBy {
		values = append(values, v)
	}
	return values
}

// ListExternalPluggableDatabasesSortOrderEnum Enum with underlying type: string
type ListExternalPluggableDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalPluggableDatabasesSortOrderEnum
const (
	ListExternalPluggableDatabasesSortOrderAsc  ListExternalPluggableDatabasesSortOrderEnum = "ASC"
	ListExternalPluggableDatabasesSortOrderDesc ListExternalPluggableDatabasesSortOrderEnum = "DESC"
)

var mappingListExternalPluggableDatabasesSortOrder = map[string]ListExternalPluggableDatabasesSortOrderEnum{
	"ASC":  ListExternalPluggableDatabasesSortOrderAsc,
	"DESC": ListExternalPluggableDatabasesSortOrderDesc,
}

// GetListExternalPluggableDatabasesSortOrderEnumValues Enumerates the set of values for ListExternalPluggableDatabasesSortOrderEnum
func GetListExternalPluggableDatabasesSortOrderEnumValues() []ListExternalPluggableDatabasesSortOrderEnum {
	values := make([]ListExternalPluggableDatabasesSortOrderEnum, 0)
	for _, v := range mappingListExternalPluggableDatabasesSortOrder {
		values = append(values, v)
	}
	return values
}
