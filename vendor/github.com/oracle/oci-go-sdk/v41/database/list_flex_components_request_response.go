// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v41/common"
	"net/http"
)

// ListFlexComponentsRequest wrapper for the ListFlexComponents operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListFlexComponents.go.html to see an example of how to use ListFlexComponentsRequest.
type ListFlexComponentsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name given. The match is not case sensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListFlexComponentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for NAME is ascending. The NAME sort order is case sensitive.
	SortBy ListFlexComponentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFlexComponentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFlexComponentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFlexComponentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFlexComponentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFlexComponentsResponse wrapper for the ListFlexComponents operation
type ListFlexComponentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FlexComponentCollection instances
	FlexComponentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFlexComponentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFlexComponentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFlexComponentsSortOrderEnum Enum with underlying type: string
type ListFlexComponentsSortOrderEnum string

// Set of constants representing the allowable values for ListFlexComponentsSortOrderEnum
const (
	ListFlexComponentsSortOrderAsc  ListFlexComponentsSortOrderEnum = "ASC"
	ListFlexComponentsSortOrderDesc ListFlexComponentsSortOrderEnum = "DESC"
)

var mappingListFlexComponentsSortOrder = map[string]ListFlexComponentsSortOrderEnum{
	"ASC":  ListFlexComponentsSortOrderAsc,
	"DESC": ListFlexComponentsSortOrderDesc,
}

// GetListFlexComponentsSortOrderEnumValues Enumerates the set of values for ListFlexComponentsSortOrderEnum
func GetListFlexComponentsSortOrderEnumValues() []ListFlexComponentsSortOrderEnum {
	values := make([]ListFlexComponentsSortOrderEnum, 0)
	for _, v := range mappingListFlexComponentsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListFlexComponentsSortByEnum Enum with underlying type: string
type ListFlexComponentsSortByEnum string

// Set of constants representing the allowable values for ListFlexComponentsSortByEnum
const (
	ListFlexComponentsSortByName ListFlexComponentsSortByEnum = "NAME"
)

var mappingListFlexComponentsSortBy = map[string]ListFlexComponentsSortByEnum{
	"NAME": ListFlexComponentsSortByName,
}

// GetListFlexComponentsSortByEnumValues Enumerates the set of values for ListFlexComponentsSortByEnum
func GetListFlexComponentsSortByEnumValues() []ListFlexComponentsSortByEnum {
	values := make([]ListFlexComponentsSortByEnum, 0)
	for _, v := range mappingListFlexComponentsSortBy {
		values = append(values, v)
	}
	return values
}
