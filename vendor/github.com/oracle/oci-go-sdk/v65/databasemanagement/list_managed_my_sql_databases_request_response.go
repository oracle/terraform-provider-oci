// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListManagedMySqlDatabasesRequest wrapper for the ListManagedMySqlDatabases operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListManagedMySqlDatabases.go.html to see an example of how to use ListManagedMySqlDatabasesRequest.
type ListManagedMySqlDatabasesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListManagedMySqlDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListManagedMySqlDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListManagedMySqlDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListManagedMySqlDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListManagedMySqlDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListManagedMySqlDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListManagedMySqlDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListManagedMySqlDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListManagedMySqlDatabasesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListManagedMySqlDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListManagedMySqlDatabasesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListManagedMySqlDatabasesResponse wrapper for the ListManagedMySqlDatabases operation
type ListManagedMySqlDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ManagedMySqlDatabaseCollection instances
	ManagedMySqlDatabaseCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListManagedMySqlDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListManagedMySqlDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListManagedMySqlDatabasesSortByEnum Enum with underlying type: string
type ListManagedMySqlDatabasesSortByEnum string

// Set of constants representing the allowable values for ListManagedMySqlDatabasesSortByEnum
const (
	ListManagedMySqlDatabasesSortByTimecreated ListManagedMySqlDatabasesSortByEnum = "TIMECREATED"
	ListManagedMySqlDatabasesSortByName        ListManagedMySqlDatabasesSortByEnum = "NAME"
)

var mappingListManagedMySqlDatabasesSortByEnum = map[string]ListManagedMySqlDatabasesSortByEnum{
	"TIMECREATED": ListManagedMySqlDatabasesSortByTimecreated,
	"NAME":        ListManagedMySqlDatabasesSortByName,
}

var mappingListManagedMySqlDatabasesSortByEnumLowerCase = map[string]ListManagedMySqlDatabasesSortByEnum{
	"timecreated": ListManagedMySqlDatabasesSortByTimecreated,
	"name":        ListManagedMySqlDatabasesSortByName,
}

// GetListManagedMySqlDatabasesSortByEnumValues Enumerates the set of values for ListManagedMySqlDatabasesSortByEnum
func GetListManagedMySqlDatabasesSortByEnumValues() []ListManagedMySqlDatabasesSortByEnum {
	values := make([]ListManagedMySqlDatabasesSortByEnum, 0)
	for _, v := range mappingListManagedMySqlDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedMySqlDatabasesSortByEnumStringValues Enumerates the set of values in String for ListManagedMySqlDatabasesSortByEnum
func GetListManagedMySqlDatabasesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"NAME",
	}
}

// GetMappingListManagedMySqlDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedMySqlDatabasesSortByEnum(val string) (ListManagedMySqlDatabasesSortByEnum, bool) {
	enum, ok := mappingListManagedMySqlDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListManagedMySqlDatabasesSortOrderEnum Enum with underlying type: string
type ListManagedMySqlDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListManagedMySqlDatabasesSortOrderEnum
const (
	ListManagedMySqlDatabasesSortOrderAsc  ListManagedMySqlDatabasesSortOrderEnum = "ASC"
	ListManagedMySqlDatabasesSortOrderDesc ListManagedMySqlDatabasesSortOrderEnum = "DESC"
)

var mappingListManagedMySqlDatabasesSortOrderEnum = map[string]ListManagedMySqlDatabasesSortOrderEnum{
	"ASC":  ListManagedMySqlDatabasesSortOrderAsc,
	"DESC": ListManagedMySqlDatabasesSortOrderDesc,
}

var mappingListManagedMySqlDatabasesSortOrderEnumLowerCase = map[string]ListManagedMySqlDatabasesSortOrderEnum{
	"asc":  ListManagedMySqlDatabasesSortOrderAsc,
	"desc": ListManagedMySqlDatabasesSortOrderDesc,
}

// GetListManagedMySqlDatabasesSortOrderEnumValues Enumerates the set of values for ListManagedMySqlDatabasesSortOrderEnum
func GetListManagedMySqlDatabasesSortOrderEnumValues() []ListManagedMySqlDatabasesSortOrderEnum {
	values := make([]ListManagedMySqlDatabasesSortOrderEnum, 0)
	for _, v := range mappingListManagedMySqlDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListManagedMySqlDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListManagedMySqlDatabasesSortOrderEnum
func GetListManagedMySqlDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListManagedMySqlDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListManagedMySqlDatabasesSortOrderEnum(val string) (ListManagedMySqlDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListManagedMySqlDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
