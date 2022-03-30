// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataconnectivity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v64/common"
	"net/http"
	"strings"
)

// ListTypesRequest wrapper for the ListTypes operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataconnectivity/ListTypes.go.html to see an example of how to use ListTypesRequest.
type ListTypesRequest struct {

	// The registry Ocid.
	RegistryId *string `mandatory:"true" contributesTo:"path" name:"registryId"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Type of the object to filter the results with.
	Type *string `mandatory:"false" contributesTo:"query" name:"type"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTypesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTypesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTypesResponse wrapper for the ListTypes operation
type ListTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TypesSummaryCollection instances
	TypesSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Pagination token for the next page of objects.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Pagination token for the previous page of objects.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTypesSortByEnum Enum with underlying type: string
type ListTypesSortByEnum string

// Set of constants representing the allowable values for ListTypesSortByEnum
const (
	ListTypesSortById          ListTypesSortByEnum = "id"
	ListTypesSortByTimecreated ListTypesSortByEnum = "timeCreated"
	ListTypesSortByDisplayname ListTypesSortByEnum = "displayName"
)

var mappingListTypesSortByEnum = map[string]ListTypesSortByEnum{
	"id":          ListTypesSortById,
	"timeCreated": ListTypesSortByTimecreated,
	"displayName": ListTypesSortByDisplayname,
}

var mappingListTypesSortByEnumLowerCase = map[string]ListTypesSortByEnum{
	"id":          ListTypesSortById,
	"timecreated": ListTypesSortByTimecreated,
	"displayname": ListTypesSortByDisplayname,
}

// GetListTypesSortByEnumValues Enumerates the set of values for ListTypesSortByEnum
func GetListTypesSortByEnumValues() []ListTypesSortByEnum {
	values := make([]ListTypesSortByEnum, 0)
	for _, v := range mappingListTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTypesSortByEnumStringValues Enumerates the set of values in String for ListTypesSortByEnum
func GetListTypesSortByEnumStringValues() []string {
	return []string{
		"id",
		"timeCreated",
		"displayName",
	}
}

// GetMappingListTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTypesSortByEnum(val string) (ListTypesSortByEnum, bool) {
	enum, ok := mappingListTypesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTypesSortOrderEnum Enum with underlying type: string
type ListTypesSortOrderEnum string

// Set of constants representing the allowable values for ListTypesSortOrderEnum
const (
	ListTypesSortOrderAsc  ListTypesSortOrderEnum = "ASC"
	ListTypesSortOrderDesc ListTypesSortOrderEnum = "DESC"
)

var mappingListTypesSortOrderEnum = map[string]ListTypesSortOrderEnum{
	"ASC":  ListTypesSortOrderAsc,
	"DESC": ListTypesSortOrderDesc,
}

var mappingListTypesSortOrderEnumLowerCase = map[string]ListTypesSortOrderEnum{
	"asc":  ListTypesSortOrderAsc,
	"desc": ListTypesSortOrderDesc,
}

// GetListTypesSortOrderEnumValues Enumerates the set of values for ListTypesSortOrderEnum
func GetListTypesSortOrderEnumValues() []ListTypesSortOrderEnum {
	values := make([]ListTypesSortOrderEnum, 0)
	for _, v := range mappingListTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTypesSortOrderEnumStringValues Enumerates the set of values in String for ListTypesSortOrderEnum
func GetListTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTypesSortOrderEnum(val string) (ListTypesSortOrderEnum, bool) {
	enum, ok := mappingListTypesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
