// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// UsersRequest wrapper for the Users operation
type UsersRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy UsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder UsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UsersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UsersResponse wrapper for the Users operation
type UsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of string instances
	Value *string `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response UsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UsersSortByEnum Enum with underlying type: string
type UsersSortByEnum string

// Set of constants representing the allowable values for UsersSortByEnum
const (
	UsersSortByTimecreated UsersSortByEnum = "TIMECREATED"
	UsersSortByDisplayname UsersSortByEnum = "DISPLAYNAME"
)

var mappingUsersSortBy = map[string]UsersSortByEnum{
	"TIMECREATED": UsersSortByTimecreated,
	"DISPLAYNAME": UsersSortByDisplayname,
}

// GetUsersSortByEnumValues Enumerates the set of values for UsersSortByEnum
func GetUsersSortByEnumValues() []UsersSortByEnum {
	values := make([]UsersSortByEnum, 0)
	for _, v := range mappingUsersSortBy {
		values = append(values, v)
	}
	return values
}

// UsersSortOrderEnum Enum with underlying type: string
type UsersSortOrderEnum string

// Set of constants representing the allowable values for UsersSortOrderEnum
const (
	UsersSortOrderAsc  UsersSortOrderEnum = "ASC"
	UsersSortOrderDesc UsersSortOrderEnum = "DESC"
)

var mappingUsersSortOrder = map[string]UsersSortOrderEnum{
	"ASC":  UsersSortOrderAsc,
	"DESC": UsersSortOrderDesc,
}

// GetUsersSortOrderEnumValues Enumerates the set of values for UsersSortOrderEnum
func GetUsersSortOrderEnumValues() []UsersSortOrderEnum {
	values := make([]UsersSortOrderEnum, 0)
	for _, v := range mappingUsersSortOrder {
		values = append(values, v)
	}
	return values
}
