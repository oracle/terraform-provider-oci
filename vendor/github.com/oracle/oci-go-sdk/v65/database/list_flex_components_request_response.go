// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFlexComponentsRequest wrapper for the ListFlexComponents operation
//
// # See also
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
func (request ListFlexComponentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFlexComponentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFlexComponentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFlexComponentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFlexComponentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFlexComponentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFlexComponentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFlexComponentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListFlexComponentsSortOrderEnum = map[string]ListFlexComponentsSortOrderEnum{
	"ASC":  ListFlexComponentsSortOrderAsc,
	"DESC": ListFlexComponentsSortOrderDesc,
}

var mappingListFlexComponentsSortOrderEnumLowerCase = map[string]ListFlexComponentsSortOrderEnum{
	"asc":  ListFlexComponentsSortOrderAsc,
	"desc": ListFlexComponentsSortOrderDesc,
}

// GetListFlexComponentsSortOrderEnumValues Enumerates the set of values for ListFlexComponentsSortOrderEnum
func GetListFlexComponentsSortOrderEnumValues() []ListFlexComponentsSortOrderEnum {
	values := make([]ListFlexComponentsSortOrderEnum, 0)
	for _, v := range mappingListFlexComponentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFlexComponentsSortOrderEnumStringValues Enumerates the set of values in String for ListFlexComponentsSortOrderEnum
func GetListFlexComponentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFlexComponentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFlexComponentsSortOrderEnum(val string) (ListFlexComponentsSortOrderEnum, bool) {
	enum, ok := mappingListFlexComponentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFlexComponentsSortByEnum Enum with underlying type: string
type ListFlexComponentsSortByEnum string

// Set of constants representing the allowable values for ListFlexComponentsSortByEnum
const (
	ListFlexComponentsSortByName ListFlexComponentsSortByEnum = "NAME"
)

var mappingListFlexComponentsSortByEnum = map[string]ListFlexComponentsSortByEnum{
	"NAME": ListFlexComponentsSortByName,
}

var mappingListFlexComponentsSortByEnumLowerCase = map[string]ListFlexComponentsSortByEnum{
	"name": ListFlexComponentsSortByName,
}

// GetListFlexComponentsSortByEnumValues Enumerates the set of values for ListFlexComponentsSortByEnum
func GetListFlexComponentsSortByEnumValues() []ListFlexComponentsSortByEnum {
	values := make([]ListFlexComponentsSortByEnum, 0)
	for _, v := range mappingListFlexComponentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFlexComponentsSortByEnumStringValues Enumerates the set of values in String for ListFlexComponentsSortByEnum
func GetListFlexComponentsSortByEnumStringValues() []string {
	return []string{
		"NAME",
	}
}

// GetMappingListFlexComponentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFlexComponentsSortByEnum(val string) (ListFlexComponentsSortByEnum, bool) {
	enum, ok := mappingListFlexComponentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
