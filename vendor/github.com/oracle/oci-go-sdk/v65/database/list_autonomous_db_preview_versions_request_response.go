// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAutonomousDbPreviewVersionsRequest wrapper for the ListAutonomousDbPreviewVersions operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/database/ListAutonomousDbPreviewVersions.go.html to see an example of how to use ListAutonomousDbPreviewVersionsRequest.
type ListAutonomousDbPreviewVersionsRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for DBWORKLOAD is ascending.
	// **Note:** If you do not include the availability domain filter, the resources are grouped by availability domain, then sorted.
	SortBy ListAutonomousDbPreviewVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAutonomousDbPreviewVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAutonomousDbPreviewVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAutonomousDbPreviewVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAutonomousDbPreviewVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAutonomousDbPreviewVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAutonomousDbPreviewVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAutonomousDbPreviewVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAutonomousDbPreviewVersionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAutonomousDbPreviewVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAutonomousDbPreviewVersionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAutonomousDbPreviewVersionsResponse wrapper for the ListAutonomousDbPreviewVersions operation
type ListAutonomousDbPreviewVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AutonomousDbPreviewVersionSummary instances
	Items []AutonomousDbPreviewVersionSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAutonomousDbPreviewVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAutonomousDbPreviewVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAutonomousDbPreviewVersionsSortByEnum Enum with underlying type: string
type ListAutonomousDbPreviewVersionsSortByEnum string

// Set of constants representing the allowable values for ListAutonomousDbPreviewVersionsSortByEnum
const (
	ListAutonomousDbPreviewVersionsSortByDbworkload ListAutonomousDbPreviewVersionsSortByEnum = "DBWORKLOAD"
)

var mappingListAutonomousDbPreviewVersionsSortByEnum = map[string]ListAutonomousDbPreviewVersionsSortByEnum{
	"DBWORKLOAD": ListAutonomousDbPreviewVersionsSortByDbworkload,
}

var mappingListAutonomousDbPreviewVersionsSortByEnumLowerCase = map[string]ListAutonomousDbPreviewVersionsSortByEnum{
	"dbworkload": ListAutonomousDbPreviewVersionsSortByDbworkload,
}

// GetListAutonomousDbPreviewVersionsSortByEnumValues Enumerates the set of values for ListAutonomousDbPreviewVersionsSortByEnum
func GetListAutonomousDbPreviewVersionsSortByEnumValues() []ListAutonomousDbPreviewVersionsSortByEnum {
	values := make([]ListAutonomousDbPreviewVersionsSortByEnum, 0)
	for _, v := range mappingListAutonomousDbPreviewVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDbPreviewVersionsSortByEnumStringValues Enumerates the set of values in String for ListAutonomousDbPreviewVersionsSortByEnum
func GetListAutonomousDbPreviewVersionsSortByEnumStringValues() []string {
	return []string{
		"DBWORKLOAD",
	}
}

// GetMappingListAutonomousDbPreviewVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDbPreviewVersionsSortByEnum(val string) (ListAutonomousDbPreviewVersionsSortByEnum, bool) {
	enum, ok := mappingListAutonomousDbPreviewVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAutonomousDbPreviewVersionsSortOrderEnum Enum with underlying type: string
type ListAutonomousDbPreviewVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListAutonomousDbPreviewVersionsSortOrderEnum
const (
	ListAutonomousDbPreviewVersionsSortOrderAsc  ListAutonomousDbPreviewVersionsSortOrderEnum = "ASC"
	ListAutonomousDbPreviewVersionsSortOrderDesc ListAutonomousDbPreviewVersionsSortOrderEnum = "DESC"
)

var mappingListAutonomousDbPreviewVersionsSortOrderEnum = map[string]ListAutonomousDbPreviewVersionsSortOrderEnum{
	"ASC":  ListAutonomousDbPreviewVersionsSortOrderAsc,
	"DESC": ListAutonomousDbPreviewVersionsSortOrderDesc,
}

var mappingListAutonomousDbPreviewVersionsSortOrderEnumLowerCase = map[string]ListAutonomousDbPreviewVersionsSortOrderEnum{
	"asc":  ListAutonomousDbPreviewVersionsSortOrderAsc,
	"desc": ListAutonomousDbPreviewVersionsSortOrderDesc,
}

// GetListAutonomousDbPreviewVersionsSortOrderEnumValues Enumerates the set of values for ListAutonomousDbPreviewVersionsSortOrderEnum
func GetListAutonomousDbPreviewVersionsSortOrderEnumValues() []ListAutonomousDbPreviewVersionsSortOrderEnum {
	values := make([]ListAutonomousDbPreviewVersionsSortOrderEnum, 0)
	for _, v := range mappingListAutonomousDbPreviewVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAutonomousDbPreviewVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListAutonomousDbPreviewVersionsSortOrderEnum
func GetListAutonomousDbPreviewVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAutonomousDbPreviewVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAutonomousDbPreviewVersionsSortOrderEnum(val string) (ListAutonomousDbPreviewVersionsSortOrderEnum, bool) {
	enum, ok := mappingListAutonomousDbPreviewVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
