// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListNetworkLoadBalancerHealthsRequest wrapper for the ListNetworkLoadBalancerHealths operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkloadbalancer/ListNetworkLoadBalancerHealths.go.html to see an example of how to use ListNetworkLoadBalancerHealthsRequest.
type ListNetworkLoadBalancerHealthsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancers to list.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The sort order to use, either 'asc' (ascending) or 'desc' (descending).
	SortOrder ListNetworkLoadBalancerHealthsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for timeCreated is descending.
	// The default order for displayName is ascending. If no value is specified, then timeCreated is the default.
	SortBy ListNetworkLoadBalancerHealthsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you must contact Oracle about a
	// particular request, then provide the request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page or items to return, in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from which to start retrieving results.
	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkLoadBalancerHealthsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkLoadBalancerHealthsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkLoadBalancerHealthsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkLoadBalancerHealthsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNetworkLoadBalancerHealthsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNetworkLoadBalancerHealthsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNetworkLoadBalancerHealthsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkLoadBalancerHealthsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNetworkLoadBalancerHealthsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNetworkLoadBalancerHealthsResponse wrapper for the ListNetworkLoadBalancerHealths operation
type ListNetworkLoadBalancerHealthsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NetworkLoadBalancerHealthCollection instances
	NetworkLoadBalancerHealthCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListNetworkLoadBalancerHealthsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkLoadBalancerHealthsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkLoadBalancerHealthsSortOrderEnum Enum with underlying type: string
type ListNetworkLoadBalancerHealthsSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkLoadBalancerHealthsSortOrderEnum
const (
	ListNetworkLoadBalancerHealthsSortOrderAsc  ListNetworkLoadBalancerHealthsSortOrderEnum = "ASC"
	ListNetworkLoadBalancerHealthsSortOrderDesc ListNetworkLoadBalancerHealthsSortOrderEnum = "DESC"
)

var mappingListNetworkLoadBalancerHealthsSortOrderEnum = map[string]ListNetworkLoadBalancerHealthsSortOrderEnum{
	"ASC":  ListNetworkLoadBalancerHealthsSortOrderAsc,
	"DESC": ListNetworkLoadBalancerHealthsSortOrderDesc,
}

// GetListNetworkLoadBalancerHealthsSortOrderEnumValues Enumerates the set of values for ListNetworkLoadBalancerHealthsSortOrderEnum
func GetListNetworkLoadBalancerHealthsSortOrderEnumValues() []ListNetworkLoadBalancerHealthsSortOrderEnum {
	values := make([]ListNetworkLoadBalancerHealthsSortOrderEnum, 0)
	for _, v := range mappingListNetworkLoadBalancerHealthsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkLoadBalancerHealthsSortOrderEnumStringValues Enumerates the set of values in String for ListNetworkLoadBalancerHealthsSortOrderEnum
func GetListNetworkLoadBalancerHealthsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNetworkLoadBalancerHealthsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkLoadBalancerHealthsSortOrderEnum(val string) (ListNetworkLoadBalancerHealthsSortOrderEnum, bool) {
	mappingListNetworkLoadBalancerHealthsSortOrderEnumIgnoreCase := make(map[string]ListNetworkLoadBalancerHealthsSortOrderEnum)
	for k, v := range mappingListNetworkLoadBalancerHealthsSortOrderEnum {
		mappingListNetworkLoadBalancerHealthsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListNetworkLoadBalancerHealthsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkLoadBalancerHealthsSortByEnum Enum with underlying type: string
type ListNetworkLoadBalancerHealthsSortByEnum string

// Set of constants representing the allowable values for ListNetworkLoadBalancerHealthsSortByEnum
const (
	ListNetworkLoadBalancerHealthsSortByTimecreated ListNetworkLoadBalancerHealthsSortByEnum = "timeCreated"
	ListNetworkLoadBalancerHealthsSortByDisplayname ListNetworkLoadBalancerHealthsSortByEnum = "displayName"
)

var mappingListNetworkLoadBalancerHealthsSortByEnum = map[string]ListNetworkLoadBalancerHealthsSortByEnum{
	"timeCreated": ListNetworkLoadBalancerHealthsSortByTimecreated,
	"displayName": ListNetworkLoadBalancerHealthsSortByDisplayname,
}

// GetListNetworkLoadBalancerHealthsSortByEnumValues Enumerates the set of values for ListNetworkLoadBalancerHealthsSortByEnum
func GetListNetworkLoadBalancerHealthsSortByEnumValues() []ListNetworkLoadBalancerHealthsSortByEnum {
	values := make([]ListNetworkLoadBalancerHealthsSortByEnum, 0)
	for _, v := range mappingListNetworkLoadBalancerHealthsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkLoadBalancerHealthsSortByEnumStringValues Enumerates the set of values in String for ListNetworkLoadBalancerHealthsSortByEnum
func GetListNetworkLoadBalancerHealthsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNetworkLoadBalancerHealthsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkLoadBalancerHealthsSortByEnum(val string) (ListNetworkLoadBalancerHealthsSortByEnum, bool) {
	mappingListNetworkLoadBalancerHealthsSortByEnumIgnoreCase := make(map[string]ListNetworkLoadBalancerHealthsSortByEnum)
	for k, v := range mappingListNetworkLoadBalancerHealthsSortByEnum {
		mappingListNetworkLoadBalancerHealthsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListNetworkLoadBalancerHealthsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
