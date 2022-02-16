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

// ListUserDefinedFunctionValidationsRequest wrapper for the ListUserDefinedFunctionValidations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListUserDefinedFunctionValidations.go.html to see an example of how to use ListUserDefinedFunctionValidationsRequest.
type ListUserDefinedFunctionValidationsRequest struct {

	// The workspace ID.
	WorkspaceId *string `mandatory:"true" contributesTo:"path" name:"workspaceId"`

	// Used to filter by the key of the object.
	Key *string `mandatory:"false" contributesTo:"query" name:"key"`

	// Used to filter by the name of the object.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Used to filter by the identifier of the object.
	Identifier *string `mandatory:"false" contributesTo:"query" name:"identifier"`

	// Specifies the fields to get for an object.
	Fields []string `contributesTo:"query" name:"fields" collectionFormat:"multi"`

	// For list pagination. The value for this parameter is the `opc-next-page` or the `opc-prev-page` response header from the previous `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sets the maximum number of results per page, or items to return in a paginated `List` call. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// Specifies the field to sort by. Accepts only one field. By default, when you sort by time fields, results are shown in descending order. All other fields default to ascending order. Sorting related parameters are ignored when parameter `query` is present (search operation and sorting order is by relevance score in descending order).
	SortBy ListUserDefinedFunctionValidationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListUserDefinedFunctionValidationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListUserDefinedFunctionValidationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListUserDefinedFunctionValidationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListUserDefinedFunctionValidationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListUserDefinedFunctionValidationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListUserDefinedFunctionValidationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListUserDefinedFunctionValidationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListUserDefinedFunctionValidationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListUserDefinedFunctionValidationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListUserDefinedFunctionValidationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListUserDefinedFunctionValidationsResponse wrapper for the ListUserDefinedFunctionValidations operation
type ListUserDefinedFunctionValidationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of UserDefinedFunctionValidationSummaryCollection instances
	UserDefinedFunctionValidationSummaryCollection `presentIn:"body"`

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

func (response ListUserDefinedFunctionValidationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListUserDefinedFunctionValidationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListUserDefinedFunctionValidationsSortByEnum Enum with underlying type: string
type ListUserDefinedFunctionValidationsSortByEnum string

// Set of constants representing the allowable values for ListUserDefinedFunctionValidationsSortByEnum
const (
	ListUserDefinedFunctionValidationsSortByTimeCreated ListUserDefinedFunctionValidationsSortByEnum = "TIME_CREATED"
	ListUserDefinedFunctionValidationsSortByDisplayName ListUserDefinedFunctionValidationsSortByEnum = "DISPLAY_NAME"
)

var mappingListUserDefinedFunctionValidationsSortByEnum = map[string]ListUserDefinedFunctionValidationsSortByEnum{
	"TIME_CREATED": ListUserDefinedFunctionValidationsSortByTimeCreated,
	"DISPLAY_NAME": ListUserDefinedFunctionValidationsSortByDisplayName,
}

// GetListUserDefinedFunctionValidationsSortByEnumValues Enumerates the set of values for ListUserDefinedFunctionValidationsSortByEnum
func GetListUserDefinedFunctionValidationsSortByEnumValues() []ListUserDefinedFunctionValidationsSortByEnum {
	values := make([]ListUserDefinedFunctionValidationsSortByEnum, 0)
	for _, v := range mappingListUserDefinedFunctionValidationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListUserDefinedFunctionValidationsSortByEnumStringValues Enumerates the set of values in String for ListUserDefinedFunctionValidationsSortByEnum
func GetListUserDefinedFunctionValidationsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListUserDefinedFunctionValidationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUserDefinedFunctionValidationsSortByEnum(val string) (ListUserDefinedFunctionValidationsSortByEnum, bool) {
	mappingListUserDefinedFunctionValidationsSortByEnumIgnoreCase := make(map[string]ListUserDefinedFunctionValidationsSortByEnum)
	for k, v := range mappingListUserDefinedFunctionValidationsSortByEnum {
		mappingListUserDefinedFunctionValidationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUserDefinedFunctionValidationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListUserDefinedFunctionValidationsSortOrderEnum Enum with underlying type: string
type ListUserDefinedFunctionValidationsSortOrderEnum string

// Set of constants representing the allowable values for ListUserDefinedFunctionValidationsSortOrderEnum
const (
	ListUserDefinedFunctionValidationsSortOrderAsc  ListUserDefinedFunctionValidationsSortOrderEnum = "ASC"
	ListUserDefinedFunctionValidationsSortOrderDesc ListUserDefinedFunctionValidationsSortOrderEnum = "DESC"
)

var mappingListUserDefinedFunctionValidationsSortOrderEnum = map[string]ListUserDefinedFunctionValidationsSortOrderEnum{
	"ASC":  ListUserDefinedFunctionValidationsSortOrderAsc,
	"DESC": ListUserDefinedFunctionValidationsSortOrderDesc,
}

// GetListUserDefinedFunctionValidationsSortOrderEnumValues Enumerates the set of values for ListUserDefinedFunctionValidationsSortOrderEnum
func GetListUserDefinedFunctionValidationsSortOrderEnumValues() []ListUserDefinedFunctionValidationsSortOrderEnum {
	values := make([]ListUserDefinedFunctionValidationsSortOrderEnum, 0)
	for _, v := range mappingListUserDefinedFunctionValidationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListUserDefinedFunctionValidationsSortOrderEnumStringValues Enumerates the set of values in String for ListUserDefinedFunctionValidationsSortOrderEnum
func GetListUserDefinedFunctionValidationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListUserDefinedFunctionValidationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListUserDefinedFunctionValidationsSortOrderEnum(val string) (ListUserDefinedFunctionValidationsSortOrderEnum, bool) {
	mappingListUserDefinedFunctionValidationsSortOrderEnumIgnoreCase := make(map[string]ListUserDefinedFunctionValidationsSortOrderEnum)
	for k, v := range mappingListUserDefinedFunctionValidationsSortOrderEnum {
		mappingListUserDefinedFunctionValidationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListUserDefinedFunctionValidationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
