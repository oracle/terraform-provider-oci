// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListColumnsRequest wrapper for the ListColumns operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListColumns.go.html to see an example of how to use ListColumnsRequest.
type ListColumnsRequest struct {

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

	// A filter to return only a specific column based on column name.
	ColumnName []string `contributesTo:"query" name:"columnName" collectionFormat:"multi"`

	// A filter to return only items related to specific datatype.
	Datatype []string `contributesTo:"query" name:"datatype" collectionFormat:"multi"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListColumnsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	SortBy ListColumnsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only items if schema name contains a specific string.
	SchemaNameContains *string `mandatory:"false" contributesTo:"query" name:"schemaNameContains"`

	// A filter to return only items if table name contains a specific string.
	TableNameContains *string `mandatory:"false" contributesTo:"query" name:"tableNameContains"`

	// A filter to return only items if column name contains a specific string.
	ColumnNameContains *string `mandatory:"false" contributesTo:"query" name:"columnNameContains"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListColumnsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListColumnsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListColumnsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListColumnsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListColumnsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListColumnsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListColumnsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListColumnsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListColumnsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListColumnsResponse wrapper for the ListColumns operation
type ListColumnsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ColumnSummary instances
	Items []ColumnSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListColumnsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListColumnsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListColumnsSortOrderEnum Enum with underlying type: string
type ListColumnsSortOrderEnum string

// Set of constants representing the allowable values for ListColumnsSortOrderEnum
const (
	ListColumnsSortOrderAsc  ListColumnsSortOrderEnum = "ASC"
	ListColumnsSortOrderDesc ListColumnsSortOrderEnum = "DESC"
)

var mappingListColumnsSortOrderEnum = map[string]ListColumnsSortOrderEnum{
	"ASC":  ListColumnsSortOrderAsc,
	"DESC": ListColumnsSortOrderDesc,
}

// GetListColumnsSortOrderEnumValues Enumerates the set of values for ListColumnsSortOrderEnum
func GetListColumnsSortOrderEnumValues() []ListColumnsSortOrderEnum {
	values := make([]ListColumnsSortOrderEnum, 0)
	for _, v := range mappingListColumnsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListColumnsSortOrderEnumStringValues Enumerates the set of values in String for ListColumnsSortOrderEnum
func GetListColumnsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListColumnsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListColumnsSortOrderEnum(val string) (ListColumnsSortOrderEnum, bool) {
	mappingListColumnsSortOrderEnumIgnoreCase := make(map[string]ListColumnsSortOrderEnum)
	for k, v := range mappingListColumnsSortOrderEnum {
		mappingListColumnsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListColumnsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListColumnsSortByEnum Enum with underlying type: string
type ListColumnsSortByEnum string

// Set of constants representing the allowable values for ListColumnsSortByEnum
const (
	ListColumnsSortBySchemaname ListColumnsSortByEnum = "SCHEMANAME"
	ListColumnsSortByTablename  ListColumnsSortByEnum = "TABLENAME"
	ListColumnsSortByColumnname ListColumnsSortByEnum = "COLUMNNAME"
	ListColumnsSortByDatatype   ListColumnsSortByEnum = "DATATYPE"
)

var mappingListColumnsSortByEnum = map[string]ListColumnsSortByEnum{
	"SCHEMANAME": ListColumnsSortBySchemaname,
	"TABLENAME":  ListColumnsSortByTablename,
	"COLUMNNAME": ListColumnsSortByColumnname,
	"DATATYPE":   ListColumnsSortByDatatype,
}

// GetListColumnsSortByEnumValues Enumerates the set of values for ListColumnsSortByEnum
func GetListColumnsSortByEnumValues() []ListColumnsSortByEnum {
	values := make([]ListColumnsSortByEnum, 0)
	for _, v := range mappingListColumnsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListColumnsSortByEnumStringValues Enumerates the set of values in String for ListColumnsSortByEnum
func GetListColumnsSortByEnumStringValues() []string {
	return []string{
		"SCHEMANAME",
		"TABLENAME",
		"COLUMNNAME",
		"DATATYPE",
	}
}

// GetMappingListColumnsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListColumnsSortByEnum(val string) (ListColumnsSortByEnum, bool) {
	mappingListColumnsSortByEnumIgnoreCase := make(map[string]ListColumnsSortByEnum)
	for k, v := range mappingListColumnsSortByEnum {
		mappingListColumnsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListColumnsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
