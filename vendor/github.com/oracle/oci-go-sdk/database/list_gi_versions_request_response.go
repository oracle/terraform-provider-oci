// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListGiVersionsRequest wrapper for the ListGiVersions operation
type ListGiVersionsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListGiVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If provided, filters the results for the given shape.
	Shape *string `mandatory:"false" contributesTo:"query" name:"shape"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListGiVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListGiVersionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListGiVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListGiVersionsResponse wrapper for the ListGiVersions operation
type ListGiVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []GiVersionSummary instances
	Items []GiVersionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListGiVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListGiVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListGiVersionsSortOrderEnum Enum with underlying type: string
type ListGiVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListGiVersionsSortOrderEnum
const (
	ListGiVersionsSortOrderAsc  ListGiVersionsSortOrderEnum = "ASC"
	ListGiVersionsSortOrderDesc ListGiVersionsSortOrderEnum = "DESC"
)

var mappingListGiVersionsSortOrder = map[string]ListGiVersionsSortOrderEnum{
	"ASC":  ListGiVersionsSortOrderAsc,
	"DESC": ListGiVersionsSortOrderDesc,
}

// GetListGiVersionsSortOrderEnumValues Enumerates the set of values for ListGiVersionsSortOrderEnum
func GetListGiVersionsSortOrderEnumValues() []ListGiVersionsSortOrderEnum {
	values := make([]ListGiVersionsSortOrderEnum, 0)
	for _, v := range mappingListGiVersionsSortOrder {
		values = append(values, v)
	}
	return values
}
