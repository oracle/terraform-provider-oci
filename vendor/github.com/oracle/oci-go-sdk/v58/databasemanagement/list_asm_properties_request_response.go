// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListAsmPropertiesRequest wrapper for the ListAsmProperties operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListAsmProperties.go.html to see an example of how to use ListAsmPropertiesRequest.
type ListAsmPropertiesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListAsmPropertiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListAsmPropertiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAsmPropertiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAsmPropertiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAsmPropertiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAsmPropertiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAsmPropertiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListAsmPropertiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAsmPropertiesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAsmPropertiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAsmPropertiesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAsmPropertiesResponse wrapper for the ListAsmProperties operation
type ListAsmPropertiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AsmPropertyCollection instances
	AsmPropertyCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAsmPropertiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAsmPropertiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAsmPropertiesSortByEnum Enum with underlying type: string
type ListAsmPropertiesSortByEnum string

// Set of constants representing the allowable values for ListAsmPropertiesSortByEnum
const (
	ListAsmPropertiesSortByTimecreated ListAsmPropertiesSortByEnum = "TIMECREATED"
	ListAsmPropertiesSortByName        ListAsmPropertiesSortByEnum = "NAME"
)

var mappingListAsmPropertiesSortByEnum = map[string]ListAsmPropertiesSortByEnum{
	"TIMECREATED": ListAsmPropertiesSortByTimecreated,
	"NAME":        ListAsmPropertiesSortByName,
}

// GetListAsmPropertiesSortByEnumValues Enumerates the set of values for ListAsmPropertiesSortByEnum
func GetListAsmPropertiesSortByEnumValues() []ListAsmPropertiesSortByEnum {
	values := make([]ListAsmPropertiesSortByEnum, 0)
	for _, v := range mappingListAsmPropertiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAsmPropertiesSortByEnumStringValues Enumerates the set of values in String for ListAsmPropertiesSortByEnum
func GetListAsmPropertiesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListAsmPropertiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAsmPropertiesSortByEnum(val string) (ListAsmPropertiesSortByEnum, bool) {
	mappingListAsmPropertiesSortByEnumIgnoreCase := make(map[string]ListAsmPropertiesSortByEnum)
	for k, v := range mappingListAsmPropertiesSortByEnum {
		mappingListAsmPropertiesSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAsmPropertiesSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListAsmPropertiesSortOrderEnum Enum with underlying type: string
type ListAsmPropertiesSortOrderEnum string

// Set of constants representing the allowable values for ListAsmPropertiesSortOrderEnum
const (
	ListAsmPropertiesSortOrderAsc  ListAsmPropertiesSortOrderEnum = "ASC"
	ListAsmPropertiesSortOrderDesc ListAsmPropertiesSortOrderEnum = "DESC"
)

var mappingListAsmPropertiesSortOrderEnum = map[string]ListAsmPropertiesSortOrderEnum{
	"ASC":  ListAsmPropertiesSortOrderAsc,
	"DESC": ListAsmPropertiesSortOrderDesc,
}

// GetListAsmPropertiesSortOrderEnumValues Enumerates the set of values for ListAsmPropertiesSortOrderEnum
func GetListAsmPropertiesSortOrderEnumValues() []ListAsmPropertiesSortOrderEnum {
	values := make([]ListAsmPropertiesSortOrderEnum, 0)
	for _, v := range mappingListAsmPropertiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAsmPropertiesSortOrderEnumStringValues Enumerates the set of values in String for ListAsmPropertiesSortOrderEnum
func GetListAsmPropertiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAsmPropertiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAsmPropertiesSortOrderEnum(val string) (ListAsmPropertiesSortOrderEnum, bool) {
	mappingListAsmPropertiesSortOrderEnumIgnoreCase := make(map[string]ListAsmPropertiesSortOrderEnum)
	for k, v := range mappingListAsmPropertiesSortOrderEnum {
		mappingListAsmPropertiesSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListAsmPropertiesSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
