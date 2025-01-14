// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListBdsClusterVersionsRequest wrapper for the ListBdsClusterVersions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/ListBdsClusterVersions.go.html to see an example of how to use ListBdsClusterVersionsRequest.
type ListBdsClusterVersionsRequest struct {

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. If no value is specified bdsVersion is default.
	SortBy ListBdsClusterVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBdsClusterVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBdsClusterVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBdsClusterVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListBdsClusterVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBdsClusterVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListBdsClusterVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListBdsClusterVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListBdsClusterVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListBdsClusterVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListBdsClusterVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListBdsClusterVersionsResponse wrapper for the ListBdsClusterVersions operation
type ListBdsClusterVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BdsClusterVersionSummary instances
	Items []BdsClusterVersionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBdsClusterVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBdsClusterVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBdsClusterVersionsSortByEnum Enum with underlying type: string
type ListBdsClusterVersionsSortByEnum string

// Set of constants representing the allowable values for ListBdsClusterVersionsSortByEnum
const (
	ListBdsClusterVersionsSortByBdsversion ListBdsClusterVersionsSortByEnum = "bdsVersion"
)

var mappingListBdsClusterVersionsSortByEnum = map[string]ListBdsClusterVersionsSortByEnum{
	"bdsVersion": ListBdsClusterVersionsSortByBdsversion,
}

var mappingListBdsClusterVersionsSortByEnumLowerCase = map[string]ListBdsClusterVersionsSortByEnum{
	"bdsversion": ListBdsClusterVersionsSortByBdsversion,
}

// GetListBdsClusterVersionsSortByEnumValues Enumerates the set of values for ListBdsClusterVersionsSortByEnum
func GetListBdsClusterVersionsSortByEnumValues() []ListBdsClusterVersionsSortByEnum {
	values := make([]ListBdsClusterVersionsSortByEnum, 0)
	for _, v := range mappingListBdsClusterVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsClusterVersionsSortByEnumStringValues Enumerates the set of values in String for ListBdsClusterVersionsSortByEnum
func GetListBdsClusterVersionsSortByEnumStringValues() []string {
	return []string{
		"bdsVersion",
	}
}

// GetMappingListBdsClusterVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsClusterVersionsSortByEnum(val string) (ListBdsClusterVersionsSortByEnum, bool) {
	enum, ok := mappingListBdsClusterVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListBdsClusterVersionsSortOrderEnum Enum with underlying type: string
type ListBdsClusterVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListBdsClusterVersionsSortOrderEnum
const (
	ListBdsClusterVersionsSortOrderAsc  ListBdsClusterVersionsSortOrderEnum = "ASC"
	ListBdsClusterVersionsSortOrderDesc ListBdsClusterVersionsSortOrderEnum = "DESC"
)

var mappingListBdsClusterVersionsSortOrderEnum = map[string]ListBdsClusterVersionsSortOrderEnum{
	"ASC":  ListBdsClusterVersionsSortOrderAsc,
	"DESC": ListBdsClusterVersionsSortOrderDesc,
}

var mappingListBdsClusterVersionsSortOrderEnumLowerCase = map[string]ListBdsClusterVersionsSortOrderEnum{
	"asc":  ListBdsClusterVersionsSortOrderAsc,
	"desc": ListBdsClusterVersionsSortOrderDesc,
}

// GetListBdsClusterVersionsSortOrderEnumValues Enumerates the set of values for ListBdsClusterVersionsSortOrderEnum
func GetListBdsClusterVersionsSortOrderEnumValues() []ListBdsClusterVersionsSortOrderEnum {
	values := make([]ListBdsClusterVersionsSortOrderEnum, 0)
	for _, v := range mappingListBdsClusterVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListBdsClusterVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListBdsClusterVersionsSortOrderEnum
func GetListBdsClusterVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListBdsClusterVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListBdsClusterVersionsSortOrderEnum(val string) (ListBdsClusterVersionsSortOrderEnum, bool) {
	enum, ok := mappingListBdsClusterVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
