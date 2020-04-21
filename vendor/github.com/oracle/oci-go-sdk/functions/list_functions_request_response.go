// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package functions

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListFunctionsRequest wrapper for the ListFunctions operation
type ListFunctionsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application to which this function belongs.
	ApplicationId *string `mandatory:"true" contributesTo:"query" name:"applicationId"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token for a list query returned by a previous operation
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only functions that match the lifecycle state in this parameter.
	// Example: `Creating`
	LifecycleState FunctionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only functions with display names that match the display name string. Matching is exact.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only functions with the specified OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Specifies sort order.
	// * **ASC:** Ascending sort order.
	// * **DESC:** Descending sort order.
	SortOrder ListFunctionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the attribute with which to sort the rules.
	// Default: `displayName`
	// * **timeCreated:** Sorts by timeCreated.
	// * **displayName:** Sorts by displayName.
	// * **id:** Sorts by id.
	SortBy ListFunctionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFunctionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFunctionsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFunctionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFunctionsResponse wrapper for the ListFunctions operation
type ListFunctionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FunctionSummary instances
	Items []FunctionSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of
	// results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListFunctionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFunctionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFunctionsSortOrderEnum Enum with underlying type: string
type ListFunctionsSortOrderEnum string

// Set of constants representing the allowable values for ListFunctionsSortOrderEnum
const (
	ListFunctionsSortOrderAsc  ListFunctionsSortOrderEnum = "ASC"
	ListFunctionsSortOrderDesc ListFunctionsSortOrderEnum = "DESC"
)

var mappingListFunctionsSortOrder = map[string]ListFunctionsSortOrderEnum{
	"ASC":  ListFunctionsSortOrderAsc,
	"DESC": ListFunctionsSortOrderDesc,
}

// GetListFunctionsSortOrderEnumValues Enumerates the set of values for ListFunctionsSortOrderEnum
func GetListFunctionsSortOrderEnumValues() []ListFunctionsSortOrderEnum {
	values := make([]ListFunctionsSortOrderEnum, 0)
	for _, v := range mappingListFunctionsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListFunctionsSortByEnum Enum with underlying type: string
type ListFunctionsSortByEnum string

// Set of constants representing the allowable values for ListFunctionsSortByEnum
const (
	ListFunctionsSortByTimecreated ListFunctionsSortByEnum = "timeCreated"
	ListFunctionsSortById          ListFunctionsSortByEnum = "id"
	ListFunctionsSortByDisplayname ListFunctionsSortByEnum = "displayName"
)

var mappingListFunctionsSortBy = map[string]ListFunctionsSortByEnum{
	"timeCreated": ListFunctionsSortByTimecreated,
	"id":          ListFunctionsSortById,
	"displayName": ListFunctionsSortByDisplayname,
}

// GetListFunctionsSortByEnumValues Enumerates the set of values for ListFunctionsSortByEnum
func GetListFunctionsSortByEnumValues() []ListFunctionsSortByEnum {
	values := make([]ListFunctionsSortByEnum, 0)
	for _, v := range mappingListFunctionsSortBy {
		values = append(values, v)
	}
	return values
}
