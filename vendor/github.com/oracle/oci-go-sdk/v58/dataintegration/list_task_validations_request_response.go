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

// ListTaskValidationsRequest wrapper for the ListTaskValidations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListTaskValidations.go.html to see an example of how to use ListTaskValidationsRequest.
type ListTaskValidationsRequest struct {

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
	SortBy ListTaskValidationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListTaskValidationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTaskValidationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTaskValidationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTaskValidationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTaskValidationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTaskValidationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTaskValidationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTaskValidationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTaskValidationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTaskValidationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTaskValidationsResponse wrapper for the ListTaskValidations operation
type ListTaskValidationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TaskValidationSummaryCollection instances
	TaskValidationSummaryCollection `presentIn:"body"`

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

func (response ListTaskValidationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTaskValidationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTaskValidationsSortByEnum Enum with underlying type: string
type ListTaskValidationsSortByEnum string

// Set of constants representing the allowable values for ListTaskValidationsSortByEnum
const (
	ListTaskValidationsSortByTimeCreated ListTaskValidationsSortByEnum = "TIME_CREATED"
	ListTaskValidationsSortByDisplayName ListTaskValidationsSortByEnum = "DISPLAY_NAME"
)

var mappingListTaskValidationsSortByEnum = map[string]ListTaskValidationsSortByEnum{
	"TIME_CREATED": ListTaskValidationsSortByTimeCreated,
	"DISPLAY_NAME": ListTaskValidationsSortByDisplayName,
}

// GetListTaskValidationsSortByEnumValues Enumerates the set of values for ListTaskValidationsSortByEnum
func GetListTaskValidationsSortByEnumValues() []ListTaskValidationsSortByEnum {
	values := make([]ListTaskValidationsSortByEnum, 0)
	for _, v := range mappingListTaskValidationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskValidationsSortByEnumStringValues Enumerates the set of values in String for ListTaskValidationsSortByEnum
func GetListTaskValidationsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
	}
}

// GetMappingListTaskValidationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskValidationsSortByEnum(val string) (ListTaskValidationsSortByEnum, bool) {
	mappingListTaskValidationsSortByEnumIgnoreCase := make(map[string]ListTaskValidationsSortByEnum)
	for k, v := range mappingListTaskValidationsSortByEnum {
		mappingListTaskValidationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTaskValidationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListTaskValidationsSortOrderEnum Enum with underlying type: string
type ListTaskValidationsSortOrderEnum string

// Set of constants representing the allowable values for ListTaskValidationsSortOrderEnum
const (
	ListTaskValidationsSortOrderAsc  ListTaskValidationsSortOrderEnum = "ASC"
	ListTaskValidationsSortOrderDesc ListTaskValidationsSortOrderEnum = "DESC"
)

var mappingListTaskValidationsSortOrderEnum = map[string]ListTaskValidationsSortOrderEnum{
	"ASC":  ListTaskValidationsSortOrderAsc,
	"DESC": ListTaskValidationsSortOrderDesc,
}

// GetListTaskValidationsSortOrderEnumValues Enumerates the set of values for ListTaskValidationsSortOrderEnum
func GetListTaskValidationsSortOrderEnumValues() []ListTaskValidationsSortOrderEnum {
	values := make([]ListTaskValidationsSortOrderEnum, 0)
	for _, v := range mappingListTaskValidationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTaskValidationsSortOrderEnumStringValues Enumerates the set of values in String for ListTaskValidationsSortOrderEnum
func GetListTaskValidationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTaskValidationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTaskValidationsSortOrderEnum(val string) (ListTaskValidationsSortOrderEnum, bool) {
	mappingListTaskValidationsSortOrderEnumIgnoreCase := make(map[string]ListTaskValidationsSortOrderEnum)
	for k, v := range mappingListTaskValidationsSortOrderEnum {
		mappingListTaskValidationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListTaskValidationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
