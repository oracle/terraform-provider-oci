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

// ListMlApplicationInstanceViewsRequest wrapper for the ListMlApplicationInstanceViews operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datascience/ListMlApplicationInstanceViews.go.html to see an example of how to use ListMlApplicationInstanceViewsRequest.
type ListMlApplicationInstanceViewsRequest struct {

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// unique MlApplication identifier
	MlApplicationId *string `mandatory:"false" contributesTo:"query" name:"mlApplicationId"`

	// unique MlApplicationImplementation identifier
	MlApplicationImplementationId *string `mandatory:"false" contributesTo:"query" name:"mlApplicationImplementationId"`

	// A filter to return only resources matching the given lifecycleState.
	LifecycleState MlApplicationInstanceViewLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

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
	SortOrder ListMlApplicationInstanceViewsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListMlApplicationInstanceViewsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMlApplicationInstanceViewsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMlApplicationInstanceViewsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMlApplicationInstanceViewsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMlApplicationInstanceViewsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMlApplicationInstanceViewsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMlApplicationInstanceViewLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetMlApplicationInstanceViewLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationInstanceViewsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMlApplicationInstanceViewsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMlApplicationInstanceViewsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMlApplicationInstanceViewsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListMlApplicationInstanceViewsResponse wrapper for the ListMlApplicationInstanceViews operation
type ListMlApplicationInstanceViewsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MlApplicationInstanceViewCollection instances
	MlApplicationInstanceViewCollection `presentIn:"body"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMlApplicationInstanceViewsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMlApplicationInstanceViewsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMlApplicationInstanceViewsSortOrderEnum Enum with underlying type: string
type ListMlApplicationInstanceViewsSortOrderEnum string

// Set of constants representing the allowable values for ListMlApplicationInstanceViewsSortOrderEnum
const (
	ListMlApplicationInstanceViewsSortOrderAsc  ListMlApplicationInstanceViewsSortOrderEnum = "ASC"
	ListMlApplicationInstanceViewsSortOrderDesc ListMlApplicationInstanceViewsSortOrderEnum = "DESC"
)

var mappingListMlApplicationInstanceViewsSortOrderEnum = map[string]ListMlApplicationInstanceViewsSortOrderEnum{
	"ASC":  ListMlApplicationInstanceViewsSortOrderAsc,
	"DESC": ListMlApplicationInstanceViewsSortOrderDesc,
}

var mappingListMlApplicationInstanceViewsSortOrderEnumLowerCase = map[string]ListMlApplicationInstanceViewsSortOrderEnum{
	"asc":  ListMlApplicationInstanceViewsSortOrderAsc,
	"desc": ListMlApplicationInstanceViewsSortOrderDesc,
}

// GetListMlApplicationInstanceViewsSortOrderEnumValues Enumerates the set of values for ListMlApplicationInstanceViewsSortOrderEnum
func GetListMlApplicationInstanceViewsSortOrderEnumValues() []ListMlApplicationInstanceViewsSortOrderEnum {
	values := make([]ListMlApplicationInstanceViewsSortOrderEnum, 0)
	for _, v := range mappingListMlApplicationInstanceViewsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationInstanceViewsSortOrderEnumStringValues Enumerates the set of values in String for ListMlApplicationInstanceViewsSortOrderEnum
func GetListMlApplicationInstanceViewsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMlApplicationInstanceViewsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationInstanceViewsSortOrderEnum(val string) (ListMlApplicationInstanceViewsSortOrderEnum, bool) {
	enum, ok := mappingListMlApplicationInstanceViewsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListMlApplicationInstanceViewsSortByEnum Enum with underlying type: string
type ListMlApplicationInstanceViewsSortByEnum string

// Set of constants representing the allowable values for ListMlApplicationInstanceViewsSortByEnum
const (
	ListMlApplicationInstanceViewsSortByTimecreated ListMlApplicationInstanceViewsSortByEnum = "timeCreated"
	ListMlApplicationInstanceViewsSortByName        ListMlApplicationInstanceViewsSortByEnum = "name"
)

var mappingListMlApplicationInstanceViewsSortByEnum = map[string]ListMlApplicationInstanceViewsSortByEnum{
	"timeCreated": ListMlApplicationInstanceViewsSortByTimecreated,
	"name":        ListMlApplicationInstanceViewsSortByName,
}

var mappingListMlApplicationInstanceViewsSortByEnumLowerCase = map[string]ListMlApplicationInstanceViewsSortByEnum{
	"timecreated": ListMlApplicationInstanceViewsSortByTimecreated,
	"name":        ListMlApplicationInstanceViewsSortByName,
}

// GetListMlApplicationInstanceViewsSortByEnumValues Enumerates the set of values for ListMlApplicationInstanceViewsSortByEnum
func GetListMlApplicationInstanceViewsSortByEnumValues() []ListMlApplicationInstanceViewsSortByEnum {
	values := make([]ListMlApplicationInstanceViewsSortByEnum, 0)
	for _, v := range mappingListMlApplicationInstanceViewsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMlApplicationInstanceViewsSortByEnumStringValues Enumerates the set of values in String for ListMlApplicationInstanceViewsSortByEnum
func GetListMlApplicationInstanceViewsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListMlApplicationInstanceViewsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMlApplicationInstanceViewsSortByEnum(val string) (ListMlApplicationInstanceViewsSortByEnum, bool) {
	enum, ok := mappingListMlApplicationInstanceViewsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
