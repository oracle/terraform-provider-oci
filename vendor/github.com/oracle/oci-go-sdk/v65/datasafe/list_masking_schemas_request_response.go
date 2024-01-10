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

// ListMaskingSchemasRequest wrapper for the ListMaskingSchemas operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListMaskingSchemas.go.html to see an example of how to use ListMaskingSchemasRequest.
type ListMaskingSchemasRequest struct {

	// The OCID of the masking policy.
	MaskingPolicyId *string `mandatory:"true" contributesTo:"path" name:"maskingPolicyId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListMaskingSchemasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder).
	// The default order is ascending.
	SortBy ListMaskingSchemasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMaskingSchemasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMaskingSchemasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMaskingSchemasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMaskingSchemasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMaskingSchemasRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMaskingSchemasSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMaskingSchemasSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMaskingSchemasSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMaskingSchemasSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMaskingSchemasResponse wrapper for the ListMaskingSchemas operation
type ListMaskingSchemasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MaskingSchemaCollection instances
	MaskingSchemaCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListMaskingSchemasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMaskingSchemasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMaskingSchemasSortOrderEnum Enum with underlying type: string
type ListMaskingSchemasSortOrderEnum string

// Set of constants representing the allowable values for ListMaskingSchemasSortOrderEnum
const (
	ListMaskingSchemasSortOrderAsc  ListMaskingSchemasSortOrderEnum = "ASC"
	ListMaskingSchemasSortOrderDesc ListMaskingSchemasSortOrderEnum = "DESC"
)

var mappingListMaskingSchemasSortOrderEnum = map[string]ListMaskingSchemasSortOrderEnum{
	"ASC":  ListMaskingSchemasSortOrderAsc,
	"DESC": ListMaskingSchemasSortOrderDesc,
}

var mappingListMaskingSchemasSortOrderEnumLowerCase = map[string]ListMaskingSchemasSortOrderEnum{
	"asc":  ListMaskingSchemasSortOrderAsc,
	"desc": ListMaskingSchemasSortOrderDesc,
}

// GetListMaskingSchemasSortOrderEnumValues Enumerates the set of values for ListMaskingSchemasSortOrderEnum
func GetListMaskingSchemasSortOrderEnumValues() []ListMaskingSchemasSortOrderEnum {
	values := make([]ListMaskingSchemasSortOrderEnum, 0)
	for _, v := range mappingListMaskingSchemasSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingSchemasSortOrderEnumStringValues Enumerates the set of values in String for ListMaskingSchemasSortOrderEnum
func GetListMaskingSchemasSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMaskingSchemasSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingSchemasSortOrderEnum(val string) (ListMaskingSchemasSortOrderEnum, bool) {
	enum, ok := mappingListMaskingSchemasSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMaskingSchemasSortByEnum Enum with underlying type: string
type ListMaskingSchemasSortByEnum string

// Set of constants representing the allowable values for ListMaskingSchemasSortByEnum
const (
	ListMaskingSchemasSortBySchemaname ListMaskingSchemasSortByEnum = "schemaName"
)

var mappingListMaskingSchemasSortByEnum = map[string]ListMaskingSchemasSortByEnum{
	"schemaName": ListMaskingSchemasSortBySchemaname,
}

var mappingListMaskingSchemasSortByEnumLowerCase = map[string]ListMaskingSchemasSortByEnum{
	"schemaname": ListMaskingSchemasSortBySchemaname,
}

// GetListMaskingSchemasSortByEnumValues Enumerates the set of values for ListMaskingSchemasSortByEnum
func GetListMaskingSchemasSortByEnumValues() []ListMaskingSchemasSortByEnum {
	values := make([]ListMaskingSchemasSortByEnum, 0)
	for _, v := range mappingListMaskingSchemasSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMaskingSchemasSortByEnumStringValues Enumerates the set of values in String for ListMaskingSchemasSortByEnum
func GetListMaskingSchemasSortByEnumStringValues() []string {
	return []string{
		"schemaName",
	}
}

// GetMappingListMaskingSchemasSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMaskingSchemasSortByEnum(val string) (ListMaskingSchemasSortByEnum, bool) {
	enum, ok := mappingListMaskingSchemasSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
