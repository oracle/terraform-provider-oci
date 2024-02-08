// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListConnectionValidationsRequest wrapper for the ListConnectionValidations operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dataintegration/ListConnectionValidations.go.html to see an example of how to use ListConnectionValidationsRequest.
type ListConnectionValidationsRequest struct {

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
	SortBy ListConnectionValidationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListConnectionValidationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConnectionValidationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectionValidationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListConnectionValidationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectionValidationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListConnectionValidationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListConnectionValidationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListConnectionValidationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListConnectionValidationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListConnectionValidationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListConnectionValidationsResponse wrapper for the ListConnectionValidations operation
type ListConnectionValidationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ConnectionValidationSummaryCollection instances
	ConnectionValidationSummaryCollection `presentIn:"body"`

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

func (response ListConnectionValidationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectionValidationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectionValidationsSortByEnum Enum with underlying type: string
type ListConnectionValidationsSortByEnum string

// Set of constants representing the allowable values for ListConnectionValidationsSortByEnum
const (
	ListConnectionValidationsSortByTimeCreated ListConnectionValidationsSortByEnum = "TIME_CREATED"
	ListConnectionValidationsSortByDisplayName ListConnectionValidationsSortByEnum = "DISPLAY_NAME"
	ListConnectionValidationsSortByTimeUpdated ListConnectionValidationsSortByEnum = "TIME_UPDATED"
)

var mappingListConnectionValidationsSortByEnum = map[string]ListConnectionValidationsSortByEnum{
	"TIME_CREATED": ListConnectionValidationsSortByTimeCreated,
	"DISPLAY_NAME": ListConnectionValidationsSortByDisplayName,
	"TIME_UPDATED": ListConnectionValidationsSortByTimeUpdated,
}

var mappingListConnectionValidationsSortByEnumLowerCase = map[string]ListConnectionValidationsSortByEnum{
	"time_created": ListConnectionValidationsSortByTimeCreated,
	"display_name": ListConnectionValidationsSortByDisplayName,
	"time_updated": ListConnectionValidationsSortByTimeUpdated,
}

// GetListConnectionValidationsSortByEnumValues Enumerates the set of values for ListConnectionValidationsSortByEnum
func GetListConnectionValidationsSortByEnumValues() []ListConnectionValidationsSortByEnum {
	values := make([]ListConnectionValidationsSortByEnum, 0)
	for _, v := range mappingListConnectionValidationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionValidationsSortByEnumStringValues Enumerates the set of values in String for ListConnectionValidationsSortByEnum
func GetListConnectionValidationsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DISPLAY_NAME",
		"TIME_UPDATED",
	}
}

// GetMappingListConnectionValidationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionValidationsSortByEnum(val string) (ListConnectionValidationsSortByEnum, bool) {
	enum, ok := mappingListConnectionValidationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListConnectionValidationsSortOrderEnum Enum with underlying type: string
type ListConnectionValidationsSortOrderEnum string

// Set of constants representing the allowable values for ListConnectionValidationsSortOrderEnum
const (
	ListConnectionValidationsSortOrderAsc  ListConnectionValidationsSortOrderEnum = "ASC"
	ListConnectionValidationsSortOrderDesc ListConnectionValidationsSortOrderEnum = "DESC"
)

var mappingListConnectionValidationsSortOrderEnum = map[string]ListConnectionValidationsSortOrderEnum{
	"ASC":  ListConnectionValidationsSortOrderAsc,
	"DESC": ListConnectionValidationsSortOrderDesc,
}

var mappingListConnectionValidationsSortOrderEnumLowerCase = map[string]ListConnectionValidationsSortOrderEnum{
	"asc":  ListConnectionValidationsSortOrderAsc,
	"desc": ListConnectionValidationsSortOrderDesc,
}

// GetListConnectionValidationsSortOrderEnumValues Enumerates the set of values for ListConnectionValidationsSortOrderEnum
func GetListConnectionValidationsSortOrderEnumValues() []ListConnectionValidationsSortOrderEnum {
	values := make([]ListConnectionValidationsSortOrderEnum, 0)
	for _, v := range mappingListConnectionValidationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListConnectionValidationsSortOrderEnumStringValues Enumerates the set of values in String for ListConnectionValidationsSortOrderEnum
func GetListConnectionValidationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListConnectionValidationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListConnectionValidationsSortOrderEnum(val string) (ListConnectionValidationsSortOrderEnum, bool) {
	enum, ok := mappingListConnectionValidationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
