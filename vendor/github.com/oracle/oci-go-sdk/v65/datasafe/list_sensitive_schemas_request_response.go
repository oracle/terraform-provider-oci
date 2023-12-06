// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSensitiveSchemasRequest wrapper for the ListSensitiveSchemas operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveSchemas.go.html to see an example of how to use ListSensitiveSchemasRequest.
type ListSensitiveSchemasRequest struct {

	// The OCID of the sensitive data model.
	SensitiveDataModelId *string `mandatory:"true" contributesTo:"path" name:"sensitiveDataModelId"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSensitiveSchemasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder).
	// The default order is ascending.
	SortBy ListSensitiveSchemasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSensitiveSchemasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSensitiveSchemasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSensitiveSchemasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSensitiveSchemasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSensitiveSchemasRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSensitiveSchemasSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSensitiveSchemasSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveSchemasSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSensitiveSchemasSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSensitiveSchemasResponse wrapper for the ListSensitiveSchemas operation
type ListSensitiveSchemasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SensitiveSchemaCollection instances
	SensitiveSchemaCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSensitiveSchemasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSensitiveSchemasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSensitiveSchemasSortOrderEnum Enum with underlying type: string
type ListSensitiveSchemasSortOrderEnum string

// Set of constants representing the allowable values for ListSensitiveSchemasSortOrderEnum
const (
	ListSensitiveSchemasSortOrderAsc  ListSensitiveSchemasSortOrderEnum = "ASC"
	ListSensitiveSchemasSortOrderDesc ListSensitiveSchemasSortOrderEnum = "DESC"
)

var mappingListSensitiveSchemasSortOrderEnum = map[string]ListSensitiveSchemasSortOrderEnum{
	"ASC":  ListSensitiveSchemasSortOrderAsc,
	"DESC": ListSensitiveSchemasSortOrderDesc,
}

var mappingListSensitiveSchemasSortOrderEnumLowerCase = map[string]ListSensitiveSchemasSortOrderEnum{
	"asc":  ListSensitiveSchemasSortOrderAsc,
	"desc": ListSensitiveSchemasSortOrderDesc,
}

// GetListSensitiveSchemasSortOrderEnumValues Enumerates the set of values for ListSensitiveSchemasSortOrderEnum
func GetListSensitiveSchemasSortOrderEnumValues() []ListSensitiveSchemasSortOrderEnum {
	values := make([]ListSensitiveSchemasSortOrderEnum, 0)
	for _, v := range mappingListSensitiveSchemasSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveSchemasSortOrderEnumStringValues Enumerates the set of values in String for ListSensitiveSchemasSortOrderEnum
func GetListSensitiveSchemasSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSensitiveSchemasSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveSchemasSortOrderEnum(val string) (ListSensitiveSchemasSortOrderEnum, bool) {
	enum, ok := mappingListSensitiveSchemasSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveSchemasSortByEnum Enum with underlying type: string
type ListSensitiveSchemasSortByEnum string

// Set of constants representing the allowable values for ListSensitiveSchemasSortByEnum
const (
	ListSensitiveSchemasSortBySchemaname ListSensitiveSchemasSortByEnum = "schemaName"
)

var mappingListSensitiveSchemasSortByEnum = map[string]ListSensitiveSchemasSortByEnum{
	"schemaName": ListSensitiveSchemasSortBySchemaname,
}

var mappingListSensitiveSchemasSortByEnumLowerCase = map[string]ListSensitiveSchemasSortByEnum{
	"schemaname": ListSensitiveSchemasSortBySchemaname,
}

// GetListSensitiveSchemasSortByEnumValues Enumerates the set of values for ListSensitiveSchemasSortByEnum
func GetListSensitiveSchemasSortByEnumValues() []ListSensitiveSchemasSortByEnum {
	values := make([]ListSensitiveSchemasSortByEnum, 0)
	for _, v := range mappingListSensitiveSchemasSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveSchemasSortByEnumStringValues Enumerates the set of values in String for ListSensitiveSchemasSortByEnum
func GetListSensitiveSchemasSortByEnumStringValues() []string {
	return []string{
		"schemaName",
	}
}

// GetMappingListSensitiveSchemasSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveSchemasSortByEnum(val string) (ListSensitiveSchemasSortByEnum, bool) {
	enum, ok := mappingListSensitiveSchemasSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
