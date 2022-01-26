// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListPluggableDatabasesRequest wrapper for the ListPluggableDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListPluggableDatabases.go.html to see an example of how to use ListPluggableDatabasesRequest.
type ListPluggableDatabasesRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the database.
	DatabaseId *string `mandatory:"false" contributesTo:"query" name:"databaseId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for PDBNAME is ascending. The PDBNAME sort order is case sensitive.
	SortBy ListPluggableDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListPluggableDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState PluggableDatabaseSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only pluggable databases that match the entire name given. The match is not case sensitive.
	PdbName *string `mandatory:"false" contributesTo:"query" name:"pdbName"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPluggableDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPluggableDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPluggableDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPluggableDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListPluggableDatabasesResponse wrapper for the ListPluggableDatabases operation
type ListPluggableDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PluggableDatabaseSummary instances
	Items []PluggableDatabaseSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPluggableDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPluggableDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPluggableDatabasesSortByEnum Enum with underlying type: string
type ListPluggableDatabasesSortByEnum string

// Set of constants representing the allowable values for ListPluggableDatabasesSortByEnum
const (
	ListPluggableDatabasesSortByPdbname     ListPluggableDatabasesSortByEnum = "PDBNAME"
	ListPluggableDatabasesSortByTimecreated ListPluggableDatabasesSortByEnum = "TIMECREATED"
)

var mappingListPluggableDatabasesSortBy = map[string]ListPluggableDatabasesSortByEnum{
	"PDBNAME":     ListPluggableDatabasesSortByPdbname,
	"TIMECREATED": ListPluggableDatabasesSortByTimecreated,
}

// GetListPluggableDatabasesSortByEnumValues Enumerates the set of values for ListPluggableDatabasesSortByEnum
func GetListPluggableDatabasesSortByEnumValues() []ListPluggableDatabasesSortByEnum {
	values := make([]ListPluggableDatabasesSortByEnum, 0)
	for _, v := range mappingListPluggableDatabasesSortBy {
		values = append(values, v)
	}
	return values
}

// ListPluggableDatabasesSortOrderEnum Enum with underlying type: string
type ListPluggableDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListPluggableDatabasesSortOrderEnum
const (
	ListPluggableDatabasesSortOrderAsc  ListPluggableDatabasesSortOrderEnum = "ASC"
	ListPluggableDatabasesSortOrderDesc ListPluggableDatabasesSortOrderEnum = "DESC"
)

var mappingListPluggableDatabasesSortOrder = map[string]ListPluggableDatabasesSortOrderEnum{
	"ASC":  ListPluggableDatabasesSortOrderAsc,
	"DESC": ListPluggableDatabasesSortOrderDesc,
}

// GetListPluggableDatabasesSortOrderEnumValues Enumerates the set of values for ListPluggableDatabasesSortOrderEnum
func GetListPluggableDatabasesSortOrderEnumValues() []ListPluggableDatabasesSortOrderEnum {
	values := make([]ListPluggableDatabasesSortOrderEnum, 0)
	for _, v := range mappingListPluggableDatabasesSortOrder {
		values = append(values, v)
	}
	return values
}
