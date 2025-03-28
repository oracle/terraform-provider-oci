// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListMlApplicationImplementationVersionsRequest wrapper for the ListMlApplicationImplementationVersions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplicationImplementationVersions.go.html to see an example of how to use ListMlApplicationImplementationVersionsRequest.
type ListMlApplicationImplementationVersionsRequest struct {

	// unique MlApplicationImplementation identifier
	MlApplicationImplementationId *string `mandatory:"true" contributesTo:"query" name:"mlApplicationImplementationId"`

	// A filter to return only resources matching the given lifecycleState.
	LifecycleState MlApplicationImplementationVersionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListMlApplicationImplementationVersionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListMlApplicationImplementationVersionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMlApplicationImplementationVersionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMlApplicationImplementationVersionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMlApplicationImplementationVersionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMlApplicationImplementationVersionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMlApplicationImplementationVersionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationImplementationVersionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMlApplicationImplementationVersionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationImplementationVersionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMlApplicationImplementationVersionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationImplementationVersionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMlApplicationImplementationVersionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMlApplicationImplementationVersionsResponse wrapper for the ListMlApplicationImplementationVersions operation
type ListMlApplicationImplementationVersionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MlApplicationImplementationVersionCollection instances
	MlApplicationImplementationVersionCollection `presentIn:"body"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMlApplicationImplementationVersionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMlApplicationImplementationVersionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMlApplicationImplementationVersionsSortOrderEnum Enum with underlying type: string
type ListMlApplicationImplementationVersionsSortOrderEnum string

// Set of constants representing the allowable values for ListMlApplicationImplementationVersionsSortOrderEnum
const (
	ListMlApplicationImplementationVersionsSortOrderAsc  ListMlApplicationImplementationVersionsSortOrderEnum = "ASC"
	ListMlApplicationImplementationVersionsSortOrderDesc ListMlApplicationImplementationVersionsSortOrderEnum = "DESC"
)

var mappingListMlApplicationImplementationVersionsSortOrderEnum = map[string]ListMlApplicationImplementationVersionsSortOrderEnum{
	"ASC":  ListMlApplicationImplementationVersionsSortOrderAsc,
	"DESC": ListMlApplicationImplementationVersionsSortOrderDesc,
}

var mappingListMlApplicationImplementationVersionsSortOrderEnumLowerCase = map[string]ListMlApplicationImplementationVersionsSortOrderEnum{
	"asc":  ListMlApplicationImplementationVersionsSortOrderAsc,
	"desc": ListMlApplicationImplementationVersionsSortOrderDesc,
}

// GetListMlApplicationImplementationVersionsSortOrderEnumValues Enumerates the set of values for ListMlApplicationImplementationVersionsSortOrderEnum
func GetListMlApplicationImplementationVersionsSortOrderEnumValues() []ListMlApplicationImplementationVersionsSortOrderEnum {
	values := make([]ListMlApplicationImplementationVersionsSortOrderEnum, 0)
	for _, v := range mappingListMlApplicationImplementationVersionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationImplementationVersionsSortOrderEnumStringValues Enumerates the set of values in String for ListMlApplicationImplementationVersionsSortOrderEnum
func GetListMlApplicationImplementationVersionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMlApplicationImplementationVersionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationImplementationVersionsSortOrderEnum(val string) (ListMlApplicationImplementationVersionsSortOrderEnum, bool) {
	enum, ok := mappingListMlApplicationImplementationVersionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMlApplicationImplementationVersionsSortByEnum Enum with underlying type: string
type ListMlApplicationImplementationVersionsSortByEnum string

// Set of constants representing the allowable values for ListMlApplicationImplementationVersionsSortByEnum
const (
	ListMlApplicationImplementationVersionsSortByTimecreated ListMlApplicationImplementationVersionsSortByEnum = "timeCreated"
	ListMlApplicationImplementationVersionsSortByName        ListMlApplicationImplementationVersionsSortByEnum = "name"
)

var mappingListMlApplicationImplementationVersionsSortByEnum = map[string]ListMlApplicationImplementationVersionsSortByEnum{
	"timeCreated": ListMlApplicationImplementationVersionsSortByTimecreated,
	"name":        ListMlApplicationImplementationVersionsSortByName,
}

var mappingListMlApplicationImplementationVersionsSortByEnumLowerCase = map[string]ListMlApplicationImplementationVersionsSortByEnum{
	"timecreated": ListMlApplicationImplementationVersionsSortByTimecreated,
	"name":        ListMlApplicationImplementationVersionsSortByName,
}

// GetListMlApplicationImplementationVersionsSortByEnumValues Enumerates the set of values for ListMlApplicationImplementationVersionsSortByEnum
func GetListMlApplicationImplementationVersionsSortByEnumValues() []ListMlApplicationImplementationVersionsSortByEnum {
	values := make([]ListMlApplicationImplementationVersionsSortByEnum, 0)
	for _, v := range mappingListMlApplicationImplementationVersionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationImplementationVersionsSortByEnumStringValues Enumerates the set of values in String for ListMlApplicationImplementationVersionsSortByEnum
func GetListMlApplicationImplementationVersionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListMlApplicationImplementationVersionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationImplementationVersionsSortByEnum(val string) (ListMlApplicationImplementationVersionsSortByEnum, bool) {
	enum, ok := mappingListMlApplicationImplementationVersionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
