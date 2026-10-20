// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubsettingSchemasRequest wrapper for the ListSubsettingSchemas operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSubsettingSchemas.go.html to see an example of how to use ListSubsettingSchemasRequest.
type ListSubsettingSchemasRequest struct {

	// The OCID of the subsetting policy.
	SubsettingPolicyId *string `mandatory:"true" contributesTo:"path" name:"subsettingPolicyId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSubsettingSchemasSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sorting parameter (sortOrder).
	// The default order is ascending.
	SortBy ListSubsettingSchemasSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only items related to specific schema name.
	SchemaName []string `contributesTo:"query" name:"schemaName" collectionFormat:"multi"`

	// A filter to return the schemas which are derived. A schema is derived if it is related to a input schema.
	IsDerivedSchema *bool `mandatory:"false" contributesTo:"query" name:"isDerivedSchema"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubsettingSchemasRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubsettingSchemasRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubsettingSchemasRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubsettingSchemasRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubsettingSchemasRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubsettingSchemasSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubsettingSchemasSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubsettingSchemasSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubsettingSchemasSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubsettingSchemasResponse wrapper for the ListSubsettingSchemas operation
type ListSubsettingSchemasResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SubsettingSchemaCollection instances
	SubsettingSchemaCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSubsettingSchemasResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubsettingSchemasResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubsettingSchemasSortOrderEnum Enum with underlying type: string
type ListSubsettingSchemasSortOrderEnum string

// Set of constants representing the allowable values for ListSubsettingSchemasSortOrderEnum
const (
	ListSubsettingSchemasSortOrderAsc  ListSubsettingSchemasSortOrderEnum = "ASC"
	ListSubsettingSchemasSortOrderDesc ListSubsettingSchemasSortOrderEnum = "DESC"
)

var mappingListSubsettingSchemasSortOrderEnum = map[string]ListSubsettingSchemasSortOrderEnum{
	"ASC":  ListSubsettingSchemasSortOrderAsc,
	"DESC": ListSubsettingSchemasSortOrderDesc,
}

var mappingListSubsettingSchemasSortOrderEnumLowerCase = map[string]ListSubsettingSchemasSortOrderEnum{
	"asc":  ListSubsettingSchemasSortOrderAsc,
	"desc": ListSubsettingSchemasSortOrderDesc,
}

// GetListSubsettingSchemasSortOrderEnumValues Enumerates the set of values for ListSubsettingSchemasSortOrderEnum
func GetListSubsettingSchemasSortOrderEnumValues() []ListSubsettingSchemasSortOrderEnum {
	values := make([]ListSubsettingSchemasSortOrderEnum, 0)
	for _, v := range mappingListSubsettingSchemasSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingSchemasSortOrderEnumStringValues Enumerates the set of values in String for ListSubsettingSchemasSortOrderEnum
func GetListSubsettingSchemasSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubsettingSchemasSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingSchemasSortOrderEnum(val string) (ListSubsettingSchemasSortOrderEnum, bool) {
	enum, ok := mappingListSubsettingSchemasSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubsettingSchemasSortByEnum Enum with underlying type: string
type ListSubsettingSchemasSortByEnum string

// Set of constants representing the allowable values for ListSubsettingSchemasSortByEnum
const (
	ListSubsettingSchemasSortBySchemaname ListSubsettingSchemasSortByEnum = "schemaName"
)

var mappingListSubsettingSchemasSortByEnum = map[string]ListSubsettingSchemasSortByEnum{
	"schemaName": ListSubsettingSchemasSortBySchemaname,
}

var mappingListSubsettingSchemasSortByEnumLowerCase = map[string]ListSubsettingSchemasSortByEnum{
	"schemaname": ListSubsettingSchemasSortBySchemaname,
}

// GetListSubsettingSchemasSortByEnumValues Enumerates the set of values for ListSubsettingSchemasSortByEnum
func GetListSubsettingSchemasSortByEnumValues() []ListSubsettingSchemasSortByEnum {
	values := make([]ListSubsettingSchemasSortByEnum, 0)
	for _, v := range mappingListSubsettingSchemasSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubsettingSchemasSortByEnumStringValues Enumerates the set of values in String for ListSubsettingSchemasSortByEnum
func GetListSubsettingSchemasSortByEnumStringValues() []string {
	return []string{
		"schemaName",
	}
}

// GetMappingListSubsettingSchemasSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubsettingSchemasSortByEnum(val string) (ListSubsettingSchemasSortByEnum, bool) {
	enum, ok := mappingListSubsettingSchemasSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
