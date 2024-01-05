// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudbridge

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInventoriesRequest wrapper for the ListInventories operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListInventories.go.html to see an example of how to use ListInventoriesRequest.
type ListInventoriesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInventoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListInventoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return inventory if the lifecycleState matches the given lifecycleState.
	LifecycleState InventoryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInventoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInventoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInventoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInventoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInventoriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInventoriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInventoriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInventoriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInventoriesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInventoryLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetInventoryLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInventoriesResponse wrapper for the ListInventories operation
type ListInventoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InventoryCollection instances
	InventoryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInventoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInventoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInventoriesSortOrderEnum Enum with underlying type: string
type ListInventoriesSortOrderEnum string

// Set of constants representing the allowable values for ListInventoriesSortOrderEnum
const (
	ListInventoriesSortOrderAsc  ListInventoriesSortOrderEnum = "ASC"
	ListInventoriesSortOrderDesc ListInventoriesSortOrderEnum = "DESC"
)

var mappingListInventoriesSortOrderEnum = map[string]ListInventoriesSortOrderEnum{
	"ASC":  ListInventoriesSortOrderAsc,
	"DESC": ListInventoriesSortOrderDesc,
}

var mappingListInventoriesSortOrderEnumLowerCase = map[string]ListInventoriesSortOrderEnum{
	"asc":  ListInventoriesSortOrderAsc,
	"desc": ListInventoriesSortOrderDesc,
}

// GetListInventoriesSortOrderEnumValues Enumerates the set of values for ListInventoriesSortOrderEnum
func GetListInventoriesSortOrderEnumValues() []ListInventoriesSortOrderEnum {
	values := make([]ListInventoriesSortOrderEnum, 0)
	for _, v := range mappingListInventoriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInventoriesSortOrderEnumStringValues Enumerates the set of values in String for ListInventoriesSortOrderEnum
func GetListInventoriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInventoriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInventoriesSortOrderEnum(val string) (ListInventoriesSortOrderEnum, bool) {
	enum, ok := mappingListInventoriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInventoriesSortByEnum Enum with underlying type: string
type ListInventoriesSortByEnum string

// Set of constants representing the allowable values for ListInventoriesSortByEnum
const (
	ListInventoriesSortByTimecreated ListInventoriesSortByEnum = "timeCreated"
	ListInventoriesSortByTimeupdated ListInventoriesSortByEnum = "timeUpdated"
	ListInventoriesSortByDisplayname ListInventoriesSortByEnum = "displayName"
)

var mappingListInventoriesSortByEnum = map[string]ListInventoriesSortByEnum{
	"timeCreated": ListInventoriesSortByTimecreated,
	"timeUpdated": ListInventoriesSortByTimeupdated,
	"displayName": ListInventoriesSortByDisplayname,
}

var mappingListInventoriesSortByEnumLowerCase = map[string]ListInventoriesSortByEnum{
	"timecreated": ListInventoriesSortByTimecreated,
	"timeupdated": ListInventoriesSortByTimeupdated,
	"displayname": ListInventoriesSortByDisplayname,
}

// GetListInventoriesSortByEnumValues Enumerates the set of values for ListInventoriesSortByEnum
func GetListInventoriesSortByEnumValues() []ListInventoriesSortByEnum {
	values := make([]ListInventoriesSortByEnum, 0)
	for _, v := range mappingListInventoriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInventoriesSortByEnumStringValues Enumerates the set of values in String for ListInventoriesSortByEnum
func GetListInventoriesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
		"displayName",
	}
}

// GetMappingListInventoriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInventoriesSortByEnum(val string) (ListInventoriesSortByEnum, bool) {
	enum, ok := mappingListInventoriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
