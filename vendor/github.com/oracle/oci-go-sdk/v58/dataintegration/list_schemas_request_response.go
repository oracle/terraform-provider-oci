// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListSchemasRequest wrapper for the ListSchemas operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListSchemas.go.html to see an example of how to use ListSchemasRequest.
type ListSchemasRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// The connection key.
	ConnectionKey *string `mandatory:"true" contributesTo:"path" name:"connectionKey"`

	// Schema resource name used for retrieving schemas.
	SchemaResourceName *string `mandatory:"true" contributesTo:"query" name:"schemaResourceName"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListSchemasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListSchemasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Used to filter by the name of the object.
	NameList []string `contributesTo:"query" name:"nameList" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSchemasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSchemasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSchemasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSchemasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSchemasRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSchemasSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSchemasSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSchemasSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSchemasSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSchemasResponse wrapper for the ListSchemas operation
type ListSchemasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SchemaSummaryCollection instances
	SchemaSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Total items in the entire list.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListSchemasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSchemasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSchemasSortByEnum Enum with underlying type: string
type ListSchemasSortByEnum string

// Set of constants representing the allowable values for ListSchemasSortByEnum
const (
	ListSchemasSortByTimeCreated ListSchemasSortByEnum = "TIME_CREATED"
	ListSchemasSortByDisplayName ListSchemasSortByEnum = "DISPLAY_NAME"
)

var mappingListSchemasSortByEnum = map[string]ListSchemasSortByEnum{
	"TIME_CREATED": ListSchemasSortByTimeCreated,
	"DISPLAY_NAME": ListSchemasSortByDisplayName,
}

// GetListSchemasSortByEnumValues Enumerates the set of values for ListSchemasSortByEnum
func GetListSchemasSortByEnumValues() []ListSchemasSortByEnum {
	values := make([]ListSchemasSortByEnum, 0)
	for _, v := range mappingListSchemasSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchemasSortByEnumStringValues Enumerates the set of values in String for ListSchemasSortByEnum
func GetListSchemasSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListSchemasSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchemasSortByEnum(val string) (ListSchemasSortByEnum, bool) {
	mappingListSchemasSortByEnumIgnoreCase := make(map[string]ListSchemasSortByEnum)
	for k, v := range mappingListSchemasSortByEnum {
		mappingListSchemasSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSchemasSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListSchemasSortOrderEnum Enum with underlying type: string
type ListSchemasSortOrderEnum string

// Set of constants representing the allowable values for ListSchemasSortOrderEnum
const (
	ListSchemasSortOrderAsc  ListSchemasSortOrderEnum = "ASC"
	ListSchemasSortOrderDesc ListSchemasSortOrderEnum = "DESC"
)

var mappingListSchemasSortOrderEnum = map[string]ListSchemasSortOrderEnum{
	"ASC":  ListSchemasSortOrderAsc,
	"DESC": ListSchemasSortOrderDesc,
}

// GetListSchemasSortOrderEnumValues Enumerates the set of values for ListSchemasSortOrderEnum
func GetListSchemasSortOrderEnumValues() []ListSchemasSortOrderEnum {
	values := make([]ListSchemasSortOrderEnum, 0)
	for _, v := range mappingListSchemasSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSchemasSortOrderEnumStringValues Enumerates the set of values in String for ListSchemasSortOrderEnum
func GetListSchemasSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSchemasSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSchemasSortOrderEnum(val string) (ListSchemasSortOrderEnum, bool) {
	mappingListSchemasSortOrderEnumIgnoreCase := make(map[string]ListSchemasSortOrderEnum)
	for k, v := range mappingListSchemasSortOrderEnum {
		mappingListSchemasSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListSchemasSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
