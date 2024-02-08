// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFunctionsRequest wrapper for the ListFunctions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListFunctions.go.html to see an example of how to use ListFunctionsRequest.
type ListFunctionsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the application to which this function belongs.
	ApplicationId *string `mandatory:"true" contributesTo:"query" name:"applicationId"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token for a list query returned by a previous operation
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only functions that match the lifecycle state in this parameter.
	// Example: `Creating`
	LifecycleState FunctionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only functions with display names that match the display name string. Matching is exact.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only functions with the specified OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Specifies sort order.
	// * **ASC:** Ascending sort order.
	// * **DESC:** Descending sort order.
	SortOrder ListFunctionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the attribute with which to sort the rules.
	// Default: `displayName`
	// * **timeCreated:** Sorts by timeCreated.
	// * **displayName:** Sorts by displayName.
	// * **id:** Sorts by id.
	SortBy ListFunctionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFunctionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFunctionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFunctionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFunctionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFunctionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFunctionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFunctionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFunctionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFunctionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFunctionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFunctionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFunctionsResponse wrapper for the ListFunctions operation
type ListFunctionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FunctionSummary instances
	Items []FunctionSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of
	// results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListFunctionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFunctionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFunctionsSortOrderEnum Enum with underlying type: string
type ListFunctionsSortOrderEnum string

// Set of constants representing the allowable values for ListFunctionsSortOrderEnum
const (
	ListFunctionsSortOrderAsc  ListFunctionsSortOrderEnum = "ASC"
	ListFunctionsSortOrderDesc ListFunctionsSortOrderEnum = "DESC"
)

var mappingListFunctionsSortOrderEnum = map[string]ListFunctionsSortOrderEnum{
	"ASC":  ListFunctionsSortOrderAsc,
	"DESC": ListFunctionsSortOrderDesc,
}

var mappingListFunctionsSortOrderEnumLowerCase = map[string]ListFunctionsSortOrderEnum{
	"asc":  ListFunctionsSortOrderAsc,
	"desc": ListFunctionsSortOrderDesc,
}

// GetListFunctionsSortOrderEnumValues Enumerates the set of values for ListFunctionsSortOrderEnum
func GetListFunctionsSortOrderEnumValues() []ListFunctionsSortOrderEnum {
	values := make([]ListFunctionsSortOrderEnum, 0)
	for _, v := range mappingListFunctionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFunctionsSortOrderEnumStringValues Enumerates the set of values in String for ListFunctionsSortOrderEnum
func GetListFunctionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFunctionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFunctionsSortOrderEnum(val string) (ListFunctionsSortOrderEnum, bool) {
	enum, ok := mappingListFunctionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFunctionsSortByEnum Enum with underlying type: string
type ListFunctionsSortByEnum string

// Set of constants representing the allowable values for ListFunctionsSortByEnum
const (
	ListFunctionsSortByTimecreated ListFunctionsSortByEnum = "timeCreated"
	ListFunctionsSortById          ListFunctionsSortByEnum = "id"
	ListFunctionsSortByDisplayname ListFunctionsSortByEnum = "displayName"
)

var mappingListFunctionsSortByEnum = map[string]ListFunctionsSortByEnum{
	"timeCreated": ListFunctionsSortByTimecreated,
	"id":          ListFunctionsSortById,
	"displayName": ListFunctionsSortByDisplayname,
}

var mappingListFunctionsSortByEnumLowerCase = map[string]ListFunctionsSortByEnum{
	"timecreated": ListFunctionsSortByTimecreated,
	"id":          ListFunctionsSortById,
	"displayname": ListFunctionsSortByDisplayname,
}

// GetListFunctionsSortByEnumValues Enumerates the set of values for ListFunctionsSortByEnum
func GetListFunctionsSortByEnumValues() []ListFunctionsSortByEnum {
	values := make([]ListFunctionsSortByEnum, 0)
	for _, v := range mappingListFunctionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFunctionsSortByEnumStringValues Enumerates the set of values in String for ListFunctionsSortByEnum
func GetListFunctionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"id",
		"displayName",
	}
}

// GetMappingListFunctionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFunctionsSortByEnum(val string) (ListFunctionsSortByEnum, bool) {
	enum, ok := mappingListFunctionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
