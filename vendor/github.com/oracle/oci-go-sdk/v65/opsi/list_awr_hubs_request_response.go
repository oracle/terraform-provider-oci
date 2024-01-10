// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAwrHubsRequest wrapper for the ListAwrHubs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrHubs.go.html to see an example of how to use ListAwrHubsRequest.
type ListAwrHubsRequest struct {

	// Unique Operations Insights Warehouse identifier
	OperationsInsightsWarehouseId *string `mandatory:"true" contributesTo:"query" name:"operationsInsightsWarehouseId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Awr Hub identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Lifecycle states
	LifecycleState []AwrHubLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

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
	SortOrder ListAwrHubsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAwrHubsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAwrHubsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAwrHubsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAwrHubsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAwrHubsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAwrHubsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.LifecycleState {
		if _, ok := GetMappingAwrHubLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetAwrHubLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListAwrHubsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAwrHubsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAwrHubsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAwrHubsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAwrHubsResponse wrapper for the ListAwrHubs operation
type ListAwrHubsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrHubSummaryCollection instances
	AwrHubSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAwrHubsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAwrHubsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAwrHubsSortOrderEnum Enum with underlying type: string
type ListAwrHubsSortOrderEnum string

// Set of constants representing the allowable values for ListAwrHubsSortOrderEnum
const (
	ListAwrHubsSortOrderAsc  ListAwrHubsSortOrderEnum = "ASC"
	ListAwrHubsSortOrderDesc ListAwrHubsSortOrderEnum = "DESC"
)

var mappingListAwrHubsSortOrderEnum = map[string]ListAwrHubsSortOrderEnum{
	"ASC":  ListAwrHubsSortOrderAsc,
	"DESC": ListAwrHubsSortOrderDesc,
}

var mappingListAwrHubsSortOrderEnumLowerCase = map[string]ListAwrHubsSortOrderEnum{
	"asc":  ListAwrHubsSortOrderAsc,
	"desc": ListAwrHubsSortOrderDesc,
}

// GetListAwrHubsSortOrderEnumValues Enumerates the set of values for ListAwrHubsSortOrderEnum
func GetListAwrHubsSortOrderEnumValues() []ListAwrHubsSortOrderEnum {
	values := make([]ListAwrHubsSortOrderEnum, 0)
	for _, v := range mappingListAwrHubsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrHubsSortOrderEnumStringValues Enumerates the set of values in String for ListAwrHubsSortOrderEnum
func GetListAwrHubsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAwrHubsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrHubsSortOrderEnum(val string) (ListAwrHubsSortOrderEnum, bool) {
	enum, ok := mappingListAwrHubsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAwrHubsSortByEnum Enum with underlying type: string
type ListAwrHubsSortByEnum string

// Set of constants representing the allowable values for ListAwrHubsSortByEnum
const (
	ListAwrHubsSortByTimecreated ListAwrHubsSortByEnum = "timeCreated"
	ListAwrHubsSortByDisplayname ListAwrHubsSortByEnum = "displayName"
)

var mappingListAwrHubsSortByEnum = map[string]ListAwrHubsSortByEnum{
	"timeCreated": ListAwrHubsSortByTimecreated,
	"displayName": ListAwrHubsSortByDisplayname,
}

var mappingListAwrHubsSortByEnumLowerCase = map[string]ListAwrHubsSortByEnum{
	"timecreated": ListAwrHubsSortByTimecreated,
	"displayname": ListAwrHubsSortByDisplayname,
}

// GetListAwrHubsSortByEnumValues Enumerates the set of values for ListAwrHubsSortByEnum
func GetListAwrHubsSortByEnumValues() []ListAwrHubsSortByEnum {
	values := make([]ListAwrHubsSortByEnum, 0)
	for _, v := range mappingListAwrHubsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrHubsSortByEnumStringValues Enumerates the set of values in String for ListAwrHubsSortByEnum
func GetListAwrHubsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAwrHubsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrHubsSortByEnum(val string) (ListAwrHubsSortByEnum, bool) {
	enum, ok := mappingListAwrHubsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
