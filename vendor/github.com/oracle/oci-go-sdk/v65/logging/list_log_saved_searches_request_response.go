// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package logging

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLogSavedSearchesRequest wrapper for the ListLogSavedSearches operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/logging/ListLogSavedSearches.go.html to see an example of how to use ListLogSavedSearchesRequest.
type ListLogSavedSearchesRequest struct {

	// Compartment OCID to list resources in. See compartmentIdInSubtree
	//      for nested compartments traversal.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// OCID of the LogSavedSearch.
	LogSavedSearchId *string `mandatory:"false" contributesTo:"query" name:"logSavedSearchId"`

	// Resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// For list pagination. The value of the `opc-next-page` or `opc-previous-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in a paginated "List" call.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by (one column only). Default sort order is
	// ascending exception of `timeCreated` and `timeLastModified` columns (descending).
	SortBy ListLogSavedSearchesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, whether 'asc' or 'desc'.
	SortOrder ListLogSavedSearchesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogSavedSearchesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogSavedSearchesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogSavedSearchesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogSavedSearchesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogSavedSearchesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLogSavedSearchesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLogSavedSearchesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLogSavedSearchesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogSavedSearchesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogSavedSearchesResponse wrapper for the ListLogSavedSearches operation
type ListLogSavedSearchesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogSavedSearchSummaryCollection instances
	LogSavedSearchSummaryCollection `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages
	// of results exist. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPreviousPage *string `presentIn:"header" name:"opc-previous-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLogSavedSearchesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogSavedSearchesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogSavedSearchesSortByEnum Enum with underlying type: string
type ListLogSavedSearchesSortByEnum string

// Set of constants representing the allowable values for ListLogSavedSearchesSortByEnum
const (
	ListLogSavedSearchesSortByTimecreated ListLogSavedSearchesSortByEnum = "timeCreated"
	ListLogSavedSearchesSortByDisplayname ListLogSavedSearchesSortByEnum = "displayName"
)

var mappingListLogSavedSearchesSortByEnum = map[string]ListLogSavedSearchesSortByEnum{
	"timeCreated": ListLogSavedSearchesSortByTimecreated,
	"displayName": ListLogSavedSearchesSortByDisplayname,
}

var mappingListLogSavedSearchesSortByEnumLowerCase = map[string]ListLogSavedSearchesSortByEnum{
	"timecreated": ListLogSavedSearchesSortByTimecreated,
	"displayname": ListLogSavedSearchesSortByDisplayname,
}

// GetListLogSavedSearchesSortByEnumValues Enumerates the set of values for ListLogSavedSearchesSortByEnum
func GetListLogSavedSearchesSortByEnumValues() []ListLogSavedSearchesSortByEnum {
	values := make([]ListLogSavedSearchesSortByEnum, 0)
	for _, v := range mappingListLogSavedSearchesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogSavedSearchesSortByEnumStringValues Enumerates the set of values in String for ListLogSavedSearchesSortByEnum
func GetListLogSavedSearchesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListLogSavedSearchesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogSavedSearchesSortByEnum(val string) (ListLogSavedSearchesSortByEnum, bool) {
	enum, ok := mappingListLogSavedSearchesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLogSavedSearchesSortOrderEnum Enum with underlying type: string
type ListLogSavedSearchesSortOrderEnum string

// Set of constants representing the allowable values for ListLogSavedSearchesSortOrderEnum
const (
	ListLogSavedSearchesSortOrderAsc  ListLogSavedSearchesSortOrderEnum = "ASC"
	ListLogSavedSearchesSortOrderDesc ListLogSavedSearchesSortOrderEnum = "DESC"
)

var mappingListLogSavedSearchesSortOrderEnum = map[string]ListLogSavedSearchesSortOrderEnum{
	"ASC":  ListLogSavedSearchesSortOrderAsc,
	"DESC": ListLogSavedSearchesSortOrderDesc,
}

var mappingListLogSavedSearchesSortOrderEnumLowerCase = map[string]ListLogSavedSearchesSortOrderEnum{
	"asc":  ListLogSavedSearchesSortOrderAsc,
	"desc": ListLogSavedSearchesSortOrderDesc,
}

// GetListLogSavedSearchesSortOrderEnumValues Enumerates the set of values for ListLogSavedSearchesSortOrderEnum
func GetListLogSavedSearchesSortOrderEnumValues() []ListLogSavedSearchesSortOrderEnum {
	values := make([]ListLogSavedSearchesSortOrderEnum, 0)
	for _, v := range mappingListLogSavedSearchesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogSavedSearchesSortOrderEnumStringValues Enumerates the set of values in String for ListLogSavedSearchesSortOrderEnum
func GetListLogSavedSearchesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogSavedSearchesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogSavedSearchesSortOrderEnum(val string) (ListLogSavedSearchesSortOrderEnum, bool) {
	enum, ok := mappingListLogSavedSearchesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
