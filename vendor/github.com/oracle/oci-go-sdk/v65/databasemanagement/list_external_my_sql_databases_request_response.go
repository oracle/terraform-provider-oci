// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExternalMySqlDatabasesRequest wrapper for the ListExternalMySqlDatabases operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListExternalMySqlDatabases.go.html to see an example of how to use ListExternalMySqlDatabasesRequest.
type ListExternalMySqlDatabasesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The parameter to filter by MySQL Database System type.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListExternalMySqlDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListExternalMySqlDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExternalMySqlDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExternalMySqlDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExternalMySqlDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExternalMySqlDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExternalMySqlDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExternalMySqlDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExternalMySqlDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExternalMySqlDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExternalMySqlDatabasesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExternalMySqlDatabasesResponse wrapper for the ListExternalMySqlDatabases operation
type ListExternalMySqlDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExternalMySqlDatabaseCollection instances
	ExternalMySqlDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExternalMySqlDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExternalMySqlDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExternalMySqlDatabasesSortByEnum Enum with underlying type: string
type ListExternalMySqlDatabasesSortByEnum string

// Set of constants representing the allowable values for ListExternalMySqlDatabasesSortByEnum
const (
	ListExternalMySqlDatabasesSortByTimecreated ListExternalMySqlDatabasesSortByEnum = "TIMECREATED"
	ListExternalMySqlDatabasesSortByName        ListExternalMySqlDatabasesSortByEnum = "NAME"
)

var mappingListExternalMySqlDatabasesSortByEnum = map[string]ListExternalMySqlDatabasesSortByEnum{
	"TIMECREATED": ListExternalMySqlDatabasesSortByTimecreated,
	"NAME":        ListExternalMySqlDatabasesSortByName,
}

var mappingListExternalMySqlDatabasesSortByEnumLowerCase = map[string]ListExternalMySqlDatabasesSortByEnum{
	"timecreated": ListExternalMySqlDatabasesSortByTimecreated,
	"name":        ListExternalMySqlDatabasesSortByName,
}

// GetListExternalMySqlDatabasesSortByEnumValues Enumerates the set of values for ListExternalMySqlDatabasesSortByEnum
func GetListExternalMySqlDatabasesSortByEnumValues() []ListExternalMySqlDatabasesSortByEnum {
	values := make([]ListExternalMySqlDatabasesSortByEnum, 0)
	for _, v := range mappingListExternalMySqlDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalMySqlDatabasesSortByEnumStringValues Enumerates the set of values in String for ListExternalMySqlDatabasesSortByEnum
func GetListExternalMySqlDatabasesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListExternalMySqlDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalMySqlDatabasesSortByEnum(val string) (ListExternalMySqlDatabasesSortByEnum, bool) {
	enum, ok := mappingListExternalMySqlDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExternalMySqlDatabasesSortOrderEnum Enum with underlying type: string
type ListExternalMySqlDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListExternalMySqlDatabasesSortOrderEnum
const (
	ListExternalMySqlDatabasesSortOrderAsc  ListExternalMySqlDatabasesSortOrderEnum = "ASC"
	ListExternalMySqlDatabasesSortOrderDesc ListExternalMySqlDatabasesSortOrderEnum = "DESC"
)

var mappingListExternalMySqlDatabasesSortOrderEnum = map[string]ListExternalMySqlDatabasesSortOrderEnum{
	"ASC":  ListExternalMySqlDatabasesSortOrderAsc,
	"DESC": ListExternalMySqlDatabasesSortOrderDesc,
}

var mappingListExternalMySqlDatabasesSortOrderEnumLowerCase = map[string]ListExternalMySqlDatabasesSortOrderEnum{
	"asc":  ListExternalMySqlDatabasesSortOrderAsc,
	"desc": ListExternalMySqlDatabasesSortOrderDesc,
}

// GetListExternalMySqlDatabasesSortOrderEnumValues Enumerates the set of values for ListExternalMySqlDatabasesSortOrderEnum
func GetListExternalMySqlDatabasesSortOrderEnumValues() []ListExternalMySqlDatabasesSortOrderEnum {
	values := make([]ListExternalMySqlDatabasesSortOrderEnum, 0)
	for _, v := range mappingListExternalMySqlDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExternalMySqlDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListExternalMySqlDatabasesSortOrderEnum
func GetListExternalMySqlDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExternalMySqlDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExternalMySqlDatabasesSortOrderEnum(val string) (ListExternalMySqlDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListExternalMySqlDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
