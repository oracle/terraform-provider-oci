// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apigateway

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListSdkLanguageTypesRequest wrapper for the ListSdkLanguageTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apigateway/ListSdkLanguageTypes.go.html to see an example of how to use ListSdkLanguageTypesRequest.
type ListSdkLanguageTypesRequest struct {

	// The ocid of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'. The default order depends on the sortBy value.
	SortOrder ListSdkLanguageTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	// Default order for `timeCreated` is descending. Default order for
	// `displayName` is ascending. The `displayName` sort order is case
	// sensitive.
	SortBy ListSdkLanguageTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request id for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSdkLanguageTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSdkLanguageTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSdkLanguageTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSdkLanguageTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListSdkLanguageTypesResponse wrapper for the ListSdkLanguageTypes operation
type ListSdkLanguageTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SdkLanguageTypeCollection instances
	SdkLanguageTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request
	// id.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain. For important details about how
	// pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response,
	// additional pages of results were seen previously. For important details
	// about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSdkLanguageTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSdkLanguageTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSdkLanguageTypesSortOrderEnum Enum with underlying type: string
type ListSdkLanguageTypesSortOrderEnum string

// Set of constants representing the allowable values for ListSdkLanguageTypesSortOrderEnum
const (
	ListSdkLanguageTypesSortOrderAsc  ListSdkLanguageTypesSortOrderEnum = "ASC"
	ListSdkLanguageTypesSortOrderDesc ListSdkLanguageTypesSortOrderEnum = "DESC"
)

var mappingListSdkLanguageTypesSortOrder = map[string]ListSdkLanguageTypesSortOrderEnum{
	"ASC":  ListSdkLanguageTypesSortOrderAsc,
	"DESC": ListSdkLanguageTypesSortOrderDesc,
}

// GetListSdkLanguageTypesSortOrderEnumValues Enumerates the set of values for ListSdkLanguageTypesSortOrderEnum
func GetListSdkLanguageTypesSortOrderEnumValues() []ListSdkLanguageTypesSortOrderEnum {
	values := make([]ListSdkLanguageTypesSortOrderEnum, 0)
	for _, v := range mappingListSdkLanguageTypesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListSdkLanguageTypesSortByEnum Enum with underlying type: string
type ListSdkLanguageTypesSortByEnum string

// Set of constants representing the allowable values for ListSdkLanguageTypesSortByEnum
const (
	ListSdkLanguageTypesSortByTimecreated ListSdkLanguageTypesSortByEnum = "timeCreated"
	ListSdkLanguageTypesSortByDisplayname ListSdkLanguageTypesSortByEnum = "displayName"
)

var mappingListSdkLanguageTypesSortBy = map[string]ListSdkLanguageTypesSortByEnum{
	"timeCreated": ListSdkLanguageTypesSortByTimecreated,
	"displayName": ListSdkLanguageTypesSortByDisplayname,
}

// GetListSdkLanguageTypesSortByEnumValues Enumerates the set of values for ListSdkLanguageTypesSortByEnum
func GetListSdkLanguageTypesSortByEnumValues() []ListSdkLanguageTypesSortByEnum {
	values := make([]ListSdkLanguageTypesSortByEnum, 0)
	for _, v := range mappingListSdkLanguageTypesSortBy {
		values = append(values, v)
	}
	return values
}
