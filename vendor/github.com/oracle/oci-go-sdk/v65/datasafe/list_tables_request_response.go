// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTablesRequest wrapper for the ListTables operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTables.go.html to see an example of how to use ListTablesRequest.
type ListTablesRequest struct {

	// The OCID of the Data Safe target database.
	TargetDatabaseId *string `mandatory:"true" contributesTo:"path" name:"targetDatabaseId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return only items related to specific table name.
	TableName []string `contributesTo:"query" name:"tableName" collectionFormat:"multi"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTablesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	SortBy ListTablesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only items if table name contains a specific string.
	TableNameContains *string `mandatory:"false" contributesTo:"query" name:"tableNameContains"`

	// A filter to return only items if schema name contains a specific string.
	SchemaNameContains *string `mandatory:"false" contributesTo:"query" name:"schemaNameContains"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTablesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTablesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTablesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTablesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTablesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTablesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTablesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTablesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTablesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTablesResponse wrapper for the ListTables operation
type ListTablesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []TableSummary instances
	Items []TableSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTablesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTablesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTablesSortOrderEnum Enum with underlying type: string
type ListTablesSortOrderEnum string

// Set of constants representing the allowable values for ListTablesSortOrderEnum
const (
	ListTablesSortOrderAsc  ListTablesSortOrderEnum = "ASC"
	ListTablesSortOrderDesc ListTablesSortOrderEnum = "DESC"
)

var mappingListTablesSortOrderEnum = map[string]ListTablesSortOrderEnum{
	"ASC":  ListTablesSortOrderAsc,
	"DESC": ListTablesSortOrderDesc,
}

var mappingListTablesSortOrderEnumLowerCase = map[string]ListTablesSortOrderEnum{
	"asc":  ListTablesSortOrderAsc,
	"desc": ListTablesSortOrderDesc,
}

// GetListTablesSortOrderEnumValues Enumerates the set of values for ListTablesSortOrderEnum
func GetListTablesSortOrderEnumValues() []ListTablesSortOrderEnum {
	values := make([]ListTablesSortOrderEnum, 0)
	for _, v := range mappingListTablesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTablesSortOrderEnumStringValues Enumerates the set of values in String for ListTablesSortOrderEnum
func GetListTablesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTablesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTablesSortOrderEnum(val string) (ListTablesSortOrderEnum, bool) {
	enum, ok := mappingListTablesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTablesSortByEnum Enum with underlying type: string
type ListTablesSortByEnum string

// Set of constants representing the allowable values for ListTablesSortByEnum
const (
	ListTablesSortBySchemaname ListTablesSortByEnum = "SCHEMANAME"
	ListTablesSortByTablename  ListTablesSortByEnum = "TABLENAME"
)

var mappingListTablesSortByEnum = map[string]ListTablesSortByEnum{
	"SCHEMANAME": ListTablesSortBySchemaname,
	"TABLENAME":  ListTablesSortByTablename,
}

var mappingListTablesSortByEnumLowerCase = map[string]ListTablesSortByEnum{
	"schemaname": ListTablesSortBySchemaname,
	"tablename":  ListTablesSortByTablename,
}

// GetListTablesSortByEnumValues Enumerates the set of values for ListTablesSortByEnum
func GetListTablesSortByEnumValues() []ListTablesSortByEnum {
	values := make([]ListTablesSortByEnum, 0)
	for _, v := range mappingListTablesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTablesSortByEnumStringValues Enumerates the set of values in String for ListTablesSortByEnum
func GetListTablesSortByEnumStringValues() []string {
	return []string{
		"SCHEMANAME",
		"TABLENAME",
	}
}

// GetMappingListTablesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTablesSortByEnum(val string) (ListTablesSortByEnum, bool) {
	enum, ok := mappingListTablesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
