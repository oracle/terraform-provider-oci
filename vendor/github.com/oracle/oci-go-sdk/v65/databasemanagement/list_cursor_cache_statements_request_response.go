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

// ListCursorCacheStatementsRequest wrapper for the ListCursorCacheStatements operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListCursorCacheStatements.go.html to see an example of how to use ListCursorCacheStatementsRequest.
type ListCursorCacheStatementsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// A filter to return all the SQL plan baselines that match the SQL text. By default, the search
	// is case insensitive. To run an exact or case-sensitive search, double-quote the search string.
	// You may also use the '%' symbol as a wildcard.
	SqlText *string `mandatory:"false" contributesTo:"query" name:"sqlText"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The option to sort the SQL statement summary data.
	SortBy ListCursorCacheStatementsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListCursorCacheStatementsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCursorCacheStatementsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCursorCacheStatementsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCursorCacheStatementsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCursorCacheStatementsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCursorCacheStatementsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCursorCacheStatementsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCursorCacheStatementsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCursorCacheStatementsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCursorCacheStatementsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCursorCacheStatementsResponse wrapper for the ListCursorCacheStatements operation
type ListCursorCacheStatementsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CursorCacheStatementCollection instances
	CursorCacheStatementCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCursorCacheStatementsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCursorCacheStatementsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCursorCacheStatementsSortByEnum Enum with underlying type: string
type ListCursorCacheStatementsSortByEnum string

// Set of constants representing the allowable values for ListCursorCacheStatementsSortByEnum
const (
	ListCursorCacheStatementsSortBySqlid  ListCursorCacheStatementsSortByEnum = "sqlId"
	ListCursorCacheStatementsSortBySchema ListCursorCacheStatementsSortByEnum = "schema"
)

var mappingListCursorCacheStatementsSortByEnum = map[string]ListCursorCacheStatementsSortByEnum{
	"sqlId":  ListCursorCacheStatementsSortBySqlid,
	"schema": ListCursorCacheStatementsSortBySchema,
}

var mappingListCursorCacheStatementsSortByEnumLowerCase = map[string]ListCursorCacheStatementsSortByEnum{
	"sqlid":  ListCursorCacheStatementsSortBySqlid,
	"schema": ListCursorCacheStatementsSortBySchema,
}

// GetListCursorCacheStatementsSortByEnumValues Enumerates the set of values for ListCursorCacheStatementsSortByEnum
func GetListCursorCacheStatementsSortByEnumValues() []ListCursorCacheStatementsSortByEnum {
	values := make([]ListCursorCacheStatementsSortByEnum, 0)
	for _, v := range mappingListCursorCacheStatementsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCursorCacheStatementsSortByEnumStringValues Enumerates the set of values in String for ListCursorCacheStatementsSortByEnum
func GetListCursorCacheStatementsSortByEnumStringValues() []string {
	return []string{
		"sqlId",
		"schema",
	}
}

// GetMappingListCursorCacheStatementsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCursorCacheStatementsSortByEnum(val string) (ListCursorCacheStatementsSortByEnum, bool) {
	enum, ok := mappingListCursorCacheStatementsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCursorCacheStatementsSortOrderEnum Enum with underlying type: string
type ListCursorCacheStatementsSortOrderEnum string

// Set of constants representing the allowable values for ListCursorCacheStatementsSortOrderEnum
const (
	ListCursorCacheStatementsSortOrderAsc  ListCursorCacheStatementsSortOrderEnum = "ASC"
	ListCursorCacheStatementsSortOrderDesc ListCursorCacheStatementsSortOrderEnum = "DESC"
)

var mappingListCursorCacheStatementsSortOrderEnum = map[string]ListCursorCacheStatementsSortOrderEnum{
	"ASC":  ListCursorCacheStatementsSortOrderAsc,
	"DESC": ListCursorCacheStatementsSortOrderDesc,
}

var mappingListCursorCacheStatementsSortOrderEnumLowerCase = map[string]ListCursorCacheStatementsSortOrderEnum{
	"asc":  ListCursorCacheStatementsSortOrderAsc,
	"desc": ListCursorCacheStatementsSortOrderDesc,
}

// GetListCursorCacheStatementsSortOrderEnumValues Enumerates the set of values for ListCursorCacheStatementsSortOrderEnum
func GetListCursorCacheStatementsSortOrderEnumValues() []ListCursorCacheStatementsSortOrderEnum {
	values := make([]ListCursorCacheStatementsSortOrderEnum, 0)
	for _, v := range mappingListCursorCacheStatementsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCursorCacheStatementsSortOrderEnumStringValues Enumerates the set of values in String for ListCursorCacheStatementsSortOrderEnum
func GetListCursorCacheStatementsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCursorCacheStatementsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCursorCacheStatementsSortOrderEnum(val string) (ListCursorCacheStatementsSortOrderEnum, bool) {
	enum, ok := mappingListCursorCacheStatementsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
