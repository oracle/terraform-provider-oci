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

// ListSensitiveDataModelSensitiveTypesRequest wrapper for the ListSensitiveDataModelSensitiveTypes operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSensitiveDataModelSensitiveTypes.go.html to see an example of how to use ListSensitiveDataModelSensitiveTypesRequest.
type ListSensitiveDataModelSensitiveTypesRequest struct {

	// The OCID of the sensitive data model.
	SensitiveDataModelId *string `mandatory:"true" contributesTo:"path" name:"sensitiveDataModelId"`

	// A filter to return only items related to a specific sensitive type OCID.
	SensitiveTypeId *string `mandatory:"false" contributesTo:"query" name:"sensitiveTypeId"`

	// - The field to sort by. You can specify only one sorting parameter (sortorder).
	// The default order is descending.
	SortBy ListSensitiveDataModelSensitiveTypesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListSensitiveDataModelSensitiveTypesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSensitiveDataModelSensitiveTypesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSensitiveDataModelSensitiveTypesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSensitiveDataModelSensitiveTypesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSensitiveDataModelSensitiveTypesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSensitiveDataModelSensitiveTypesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSensitiveDataModelSensitiveTypesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSensitiveDataModelSensitiveTypesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSensitiveDataModelSensitiveTypesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSensitiveDataModelSensitiveTypesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSensitiveDataModelSensitiveTypesResponse wrapper for the ListSensitiveDataModelSensitiveTypes operation
type ListSensitiveDataModelSensitiveTypesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SensitiveDataModelSensitiveTypeCollection instances
	SensitiveDataModelSensitiveTypeCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSensitiveDataModelSensitiveTypesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSensitiveDataModelSensitiveTypesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSensitiveDataModelSensitiveTypesSortByEnum Enum with underlying type: string
type ListSensitiveDataModelSensitiveTypesSortByEnum string

// Set of constants representing the allowable values for ListSensitiveDataModelSensitiveTypesSortByEnum
const (
	ListSensitiveDataModelSensitiveTypesSortByCount ListSensitiveDataModelSensitiveTypesSortByEnum = "count"
)

var mappingListSensitiveDataModelSensitiveTypesSortByEnum = map[string]ListSensitiveDataModelSensitiveTypesSortByEnum{
	"count": ListSensitiveDataModelSensitiveTypesSortByCount,
}

var mappingListSensitiveDataModelSensitiveTypesSortByEnumLowerCase = map[string]ListSensitiveDataModelSensitiveTypesSortByEnum{
	"count": ListSensitiveDataModelSensitiveTypesSortByCount,
}

// GetListSensitiveDataModelSensitiveTypesSortByEnumValues Enumerates the set of values for ListSensitiveDataModelSensitiveTypesSortByEnum
func GetListSensitiveDataModelSensitiveTypesSortByEnumValues() []ListSensitiveDataModelSensitiveTypesSortByEnum {
	values := make([]ListSensitiveDataModelSensitiveTypesSortByEnum, 0)
	for _, v := range mappingListSensitiveDataModelSensitiveTypesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveDataModelSensitiveTypesSortByEnumStringValues Enumerates the set of values in String for ListSensitiveDataModelSensitiveTypesSortByEnum
func GetListSensitiveDataModelSensitiveTypesSortByEnumStringValues() []string {
	return []string{
		"count",
	}
}

// GetMappingListSensitiveDataModelSensitiveTypesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveDataModelSensitiveTypesSortByEnum(val string) (ListSensitiveDataModelSensitiveTypesSortByEnum, bool) {
	enum, ok := mappingListSensitiveDataModelSensitiveTypesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSensitiveDataModelSensitiveTypesSortOrderEnum Enum with underlying type: string
type ListSensitiveDataModelSensitiveTypesSortOrderEnum string

// Set of constants representing the allowable values for ListSensitiveDataModelSensitiveTypesSortOrderEnum
const (
	ListSensitiveDataModelSensitiveTypesSortOrderAsc  ListSensitiveDataModelSensitiveTypesSortOrderEnum = "ASC"
	ListSensitiveDataModelSensitiveTypesSortOrderDesc ListSensitiveDataModelSensitiveTypesSortOrderEnum = "DESC"
)

var mappingListSensitiveDataModelSensitiveTypesSortOrderEnum = map[string]ListSensitiveDataModelSensitiveTypesSortOrderEnum{
	"ASC":  ListSensitiveDataModelSensitiveTypesSortOrderAsc,
	"DESC": ListSensitiveDataModelSensitiveTypesSortOrderDesc,
}

var mappingListSensitiveDataModelSensitiveTypesSortOrderEnumLowerCase = map[string]ListSensitiveDataModelSensitiveTypesSortOrderEnum{
	"asc":  ListSensitiveDataModelSensitiveTypesSortOrderAsc,
	"desc": ListSensitiveDataModelSensitiveTypesSortOrderDesc,
}

// GetListSensitiveDataModelSensitiveTypesSortOrderEnumValues Enumerates the set of values for ListSensitiveDataModelSensitiveTypesSortOrderEnum
func GetListSensitiveDataModelSensitiveTypesSortOrderEnumValues() []ListSensitiveDataModelSensitiveTypesSortOrderEnum {
	values := make([]ListSensitiveDataModelSensitiveTypesSortOrderEnum, 0)
	for _, v := range mappingListSensitiveDataModelSensitiveTypesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSensitiveDataModelSensitiveTypesSortOrderEnumStringValues Enumerates the set of values in String for ListSensitiveDataModelSensitiveTypesSortOrderEnum
func GetListSensitiveDataModelSensitiveTypesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSensitiveDataModelSensitiveTypesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSensitiveDataModelSensitiveTypesSortOrderEnum(val string) (ListSensitiveDataModelSensitiveTypesSortOrderEnum, bool) {
	enum, ok := mappingListSensitiveDataModelSensitiveTypesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
