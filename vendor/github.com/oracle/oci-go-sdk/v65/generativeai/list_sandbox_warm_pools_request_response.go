// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package generativeai

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSandboxWarmPoolsRequest wrapper for the ListSandboxWarmPools operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/generativeai/ListSandboxWarmPools.go.html to see an example of how to use ListSandboxWarmPoolsRequest.
type ListSandboxWarmPoolsRequest struct {

	// Owning Generative AI Project OCID. Resolve parent compartment for authorization.
	ProjectId *string `mandatory:"true" contributesTo:"path" name:"projectId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Exact name filter.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Filter by type.
	Type ListSandboxWarmPoolsTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// Filter by shape.
	Shape ListSandboxWarmPoolsShapeEnum `mandatory:"false" contributesTo:"query" name:"shape" omitEmpty:"true"`

	// Filter by sortBy.
	SortBy ListSandboxWarmPoolsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Filter by sortOrder.
	SortOrder ListSandboxWarmPoolsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSandboxWarmPoolsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSandboxWarmPoolsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSandboxWarmPoolsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSandboxWarmPoolsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSandboxWarmPoolsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSandboxWarmPoolsTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetListSandboxWarmPoolsTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSandboxWarmPoolsShapeEnum(string(request.Shape)); !ok && request.Shape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Shape: %s. Supported values are: %s.", request.Shape, strings.Join(GetListSandboxWarmPoolsShapeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSandboxWarmPoolsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSandboxWarmPoolsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSandboxWarmPoolsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSandboxWarmPoolsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSandboxWarmPoolsResponse wrapper for the ListSandboxWarmPools operation
type ListSandboxWarmPoolsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SandboxWarmPoolCollection instances
	SandboxWarmPoolCollection `presentIn:"body"`

	// Request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Opaque continuation token bound to Project and filters.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListSandboxWarmPoolsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSandboxWarmPoolsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSandboxWarmPoolsTypeEnum Enum with underlying type: string
type ListSandboxWarmPoolsTypeEnum string

// Set of constants representing the allowable values for ListSandboxWarmPoolsTypeEnum
const (
	ListSandboxWarmPoolsTypeManaged ListSandboxWarmPoolsTypeEnum = "MANAGED"
	ListSandboxWarmPoolsTypeCustom  ListSandboxWarmPoolsTypeEnum = "CUSTOM"
)

var mappingListSandboxWarmPoolsTypeEnum = map[string]ListSandboxWarmPoolsTypeEnum{
	"MANAGED": ListSandboxWarmPoolsTypeManaged,
	"CUSTOM":  ListSandboxWarmPoolsTypeCustom,
}

var mappingListSandboxWarmPoolsTypeEnumLowerCase = map[string]ListSandboxWarmPoolsTypeEnum{
	"managed": ListSandboxWarmPoolsTypeManaged,
	"custom":  ListSandboxWarmPoolsTypeCustom,
}

// GetListSandboxWarmPoolsTypeEnumValues Enumerates the set of values for ListSandboxWarmPoolsTypeEnum
func GetListSandboxWarmPoolsTypeEnumValues() []ListSandboxWarmPoolsTypeEnum {
	values := make([]ListSandboxWarmPoolsTypeEnum, 0)
	for _, v := range mappingListSandboxWarmPoolsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxWarmPoolsTypeEnumStringValues Enumerates the set of values in String for ListSandboxWarmPoolsTypeEnum
func GetListSandboxWarmPoolsTypeEnumStringValues() []string {
	return []string{
		"MANAGED",
		"CUSTOM",
	}
}

// GetMappingListSandboxWarmPoolsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxWarmPoolsTypeEnum(val string) (ListSandboxWarmPoolsTypeEnum, bool) {
	enum, ok := mappingListSandboxWarmPoolsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSandboxWarmPoolsShapeEnum Enum with underlying type: string
type ListSandboxWarmPoolsShapeEnum string

// Set of constants representing the allowable values for ListSandboxWarmPoolsShapeEnum
const (
	ListSandboxWarmPoolsShapeSmall  ListSandboxWarmPoolsShapeEnum = "SMALL"
	ListSandboxWarmPoolsShapeMedium ListSandboxWarmPoolsShapeEnum = "MEDIUM"
	ListSandboxWarmPoolsShapeLarge  ListSandboxWarmPoolsShapeEnum = "LARGE"
)

var mappingListSandboxWarmPoolsShapeEnum = map[string]ListSandboxWarmPoolsShapeEnum{
	"SMALL":  ListSandboxWarmPoolsShapeSmall,
	"MEDIUM": ListSandboxWarmPoolsShapeMedium,
	"LARGE":  ListSandboxWarmPoolsShapeLarge,
}

var mappingListSandboxWarmPoolsShapeEnumLowerCase = map[string]ListSandboxWarmPoolsShapeEnum{
	"small":  ListSandboxWarmPoolsShapeSmall,
	"medium": ListSandboxWarmPoolsShapeMedium,
	"large":  ListSandboxWarmPoolsShapeLarge,
}

// GetListSandboxWarmPoolsShapeEnumValues Enumerates the set of values for ListSandboxWarmPoolsShapeEnum
func GetListSandboxWarmPoolsShapeEnumValues() []ListSandboxWarmPoolsShapeEnum {
	values := make([]ListSandboxWarmPoolsShapeEnum, 0)
	for _, v := range mappingListSandboxWarmPoolsShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxWarmPoolsShapeEnumStringValues Enumerates the set of values in String for ListSandboxWarmPoolsShapeEnum
func GetListSandboxWarmPoolsShapeEnumStringValues() []string {
	return []string{
		"SMALL",
		"MEDIUM",
		"LARGE",
	}
}

// GetMappingListSandboxWarmPoolsShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxWarmPoolsShapeEnum(val string) (ListSandboxWarmPoolsShapeEnum, bool) {
	enum, ok := mappingListSandboxWarmPoolsShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSandboxWarmPoolsSortByEnum Enum with underlying type: string
type ListSandboxWarmPoolsSortByEnum string

// Set of constants representing the allowable values for ListSandboxWarmPoolsSortByEnum
const (
	ListSandboxWarmPoolsSortByTimecreated ListSandboxWarmPoolsSortByEnum = "timeCreated"
	ListSandboxWarmPoolsSortByName        ListSandboxWarmPoolsSortByEnum = "name"
)

var mappingListSandboxWarmPoolsSortByEnum = map[string]ListSandboxWarmPoolsSortByEnum{
	"timeCreated": ListSandboxWarmPoolsSortByTimecreated,
	"name":        ListSandboxWarmPoolsSortByName,
}

var mappingListSandboxWarmPoolsSortByEnumLowerCase = map[string]ListSandboxWarmPoolsSortByEnum{
	"timecreated": ListSandboxWarmPoolsSortByTimecreated,
	"name":        ListSandboxWarmPoolsSortByName,
}

// GetListSandboxWarmPoolsSortByEnumValues Enumerates the set of values for ListSandboxWarmPoolsSortByEnum
func GetListSandboxWarmPoolsSortByEnumValues() []ListSandboxWarmPoolsSortByEnum {
	values := make([]ListSandboxWarmPoolsSortByEnum, 0)
	for _, v := range mappingListSandboxWarmPoolsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxWarmPoolsSortByEnumStringValues Enumerates the set of values in String for ListSandboxWarmPoolsSortByEnum
func GetListSandboxWarmPoolsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListSandboxWarmPoolsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxWarmPoolsSortByEnum(val string) (ListSandboxWarmPoolsSortByEnum, bool) {
	enum, ok := mappingListSandboxWarmPoolsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSandboxWarmPoolsSortOrderEnum Enum with underlying type: string
type ListSandboxWarmPoolsSortOrderEnum string

// Set of constants representing the allowable values for ListSandboxWarmPoolsSortOrderEnum
const (
	ListSandboxWarmPoolsSortOrderAsc  ListSandboxWarmPoolsSortOrderEnum = "ASC"
	ListSandboxWarmPoolsSortOrderDesc ListSandboxWarmPoolsSortOrderEnum = "DESC"
)

var mappingListSandboxWarmPoolsSortOrderEnum = map[string]ListSandboxWarmPoolsSortOrderEnum{
	"ASC":  ListSandboxWarmPoolsSortOrderAsc,
	"DESC": ListSandboxWarmPoolsSortOrderDesc,
}

var mappingListSandboxWarmPoolsSortOrderEnumLowerCase = map[string]ListSandboxWarmPoolsSortOrderEnum{
	"asc":  ListSandboxWarmPoolsSortOrderAsc,
	"desc": ListSandboxWarmPoolsSortOrderDesc,
}

// GetListSandboxWarmPoolsSortOrderEnumValues Enumerates the set of values for ListSandboxWarmPoolsSortOrderEnum
func GetListSandboxWarmPoolsSortOrderEnumValues() []ListSandboxWarmPoolsSortOrderEnum {
	values := make([]ListSandboxWarmPoolsSortOrderEnum, 0)
	for _, v := range mappingListSandboxWarmPoolsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSandboxWarmPoolsSortOrderEnumStringValues Enumerates the set of values in String for ListSandboxWarmPoolsSortOrderEnum
func GetListSandboxWarmPoolsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSandboxWarmPoolsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSandboxWarmPoolsSortOrderEnum(val string) (ListSandboxWarmPoolsSortOrderEnum, bool) {
	enum, ok := mappingListSandboxWarmPoolsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
