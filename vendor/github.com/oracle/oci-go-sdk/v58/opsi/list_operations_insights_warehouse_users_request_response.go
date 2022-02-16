// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListOperationsInsightsWarehouseUsersRequest wrapper for the ListOperationsInsightsWarehouseUsers operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOperationsInsightsWarehouseUsers.go.html to see an example of how to use ListOperationsInsightsWarehouseUsersRequest.
type ListOperationsInsightsWarehouseUsersRequest struct {

	// Unique Operations Insights Warehouse identifier
	OperationsInsightsWarehouseId *string `mandatory:"true" contributesTo:"query" name:"operationsInsightsWarehouseId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Operations Insights Warehouse User identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Lifecycle states
	LifecycleState []OperationsInsightsWarehouseUserLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOperationsInsightsWarehouseUsersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOperationsInsightsWarehouseUsersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOperationsInsightsWarehouseUsersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOperationsInsightsWarehouseUsersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOperationsInsightsWarehouseUsersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOperationsInsightsWarehouseUsersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOperationsInsightsWarehouseUsersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingOperationsInsightsWarehouseUserLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetOperationsInsightsWarehouseUserLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListOperationsInsightsWarehouseUsersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOperationsInsightsWarehouseUsersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOperationsInsightsWarehouseUsersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOperationsInsightsWarehouseUsersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOperationsInsightsWarehouseUsersResponse wrapper for the ListOperationsInsightsWarehouseUsers operation
type ListOperationsInsightsWarehouseUsersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OperationsInsightsWarehouseUserSummaryCollection instances
	OperationsInsightsWarehouseUserSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOperationsInsightsWarehouseUsersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOperationsInsightsWarehouseUsersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOperationsInsightsWarehouseUsersSortOrderEnum Enum with underlying type: string
type ListOperationsInsightsWarehouseUsersSortOrderEnum string

// Set of constants representing the allowable values for ListOperationsInsightsWarehouseUsersSortOrderEnum
const (
	ListOperationsInsightsWarehouseUsersSortOrderAsc  ListOperationsInsightsWarehouseUsersSortOrderEnum = "ASC"
	ListOperationsInsightsWarehouseUsersSortOrderDesc ListOperationsInsightsWarehouseUsersSortOrderEnum = "DESC"
)

var mappingListOperationsInsightsWarehouseUsersSortOrderEnum = map[string]ListOperationsInsightsWarehouseUsersSortOrderEnum{
	"ASC":  ListOperationsInsightsWarehouseUsersSortOrderAsc,
	"DESC": ListOperationsInsightsWarehouseUsersSortOrderDesc,
}

// GetListOperationsInsightsWarehouseUsersSortOrderEnumValues Enumerates the set of values for ListOperationsInsightsWarehouseUsersSortOrderEnum
func GetListOperationsInsightsWarehouseUsersSortOrderEnumValues() []ListOperationsInsightsWarehouseUsersSortOrderEnum {
	values := make([]ListOperationsInsightsWarehouseUsersSortOrderEnum, 0)
	for _, v := range mappingListOperationsInsightsWarehouseUsersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperationsInsightsWarehouseUsersSortOrderEnumStringValues Enumerates the set of values in String for ListOperationsInsightsWarehouseUsersSortOrderEnum
func GetListOperationsInsightsWarehouseUsersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOperationsInsightsWarehouseUsersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperationsInsightsWarehouseUsersSortOrderEnum(val string) (ListOperationsInsightsWarehouseUsersSortOrderEnum, bool) {
	mappingListOperationsInsightsWarehouseUsersSortOrderEnumIgnoreCase := make(map[string]ListOperationsInsightsWarehouseUsersSortOrderEnum)
	for k, v := range mappingListOperationsInsightsWarehouseUsersSortOrderEnum {
		mappingListOperationsInsightsWarehouseUsersSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListOperationsInsightsWarehouseUsersSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListOperationsInsightsWarehouseUsersSortByEnum Enum with underlying type: string
type ListOperationsInsightsWarehouseUsersSortByEnum string

// Set of constants representing the allowable values for ListOperationsInsightsWarehouseUsersSortByEnum
const (
	ListOperationsInsightsWarehouseUsersSortByTimecreated ListOperationsInsightsWarehouseUsersSortByEnum = "timeCreated"
	ListOperationsInsightsWarehouseUsersSortByDisplayname ListOperationsInsightsWarehouseUsersSortByEnum = "displayName"
)

var mappingListOperationsInsightsWarehouseUsersSortByEnum = map[string]ListOperationsInsightsWarehouseUsersSortByEnum{
	"timeCreated": ListOperationsInsightsWarehouseUsersSortByTimecreated,
	"displayName": ListOperationsInsightsWarehouseUsersSortByDisplayname,
}

// GetListOperationsInsightsWarehouseUsersSortByEnumValues Enumerates the set of values for ListOperationsInsightsWarehouseUsersSortByEnum
func GetListOperationsInsightsWarehouseUsersSortByEnumValues() []ListOperationsInsightsWarehouseUsersSortByEnum {
	values := make([]ListOperationsInsightsWarehouseUsersSortByEnum, 0)
	for _, v := range mappingListOperationsInsightsWarehouseUsersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOperationsInsightsWarehouseUsersSortByEnumStringValues Enumerates the set of values in String for ListOperationsInsightsWarehouseUsersSortByEnum
func GetListOperationsInsightsWarehouseUsersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListOperationsInsightsWarehouseUsersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOperationsInsightsWarehouseUsersSortByEnum(val string) (ListOperationsInsightsWarehouseUsersSortByEnum, bool) {
	mappingListOperationsInsightsWarehouseUsersSortByEnumIgnoreCase := make(map[string]ListOperationsInsightsWarehouseUsersSortByEnum)
	for k, v := range mappingListOperationsInsightsWarehouseUsersSortByEnum {
		mappingListOperationsInsightsWarehouseUsersSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListOperationsInsightsWarehouseUsersSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
