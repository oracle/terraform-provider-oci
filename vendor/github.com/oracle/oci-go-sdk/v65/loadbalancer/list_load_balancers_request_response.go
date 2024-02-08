// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListLoadBalancersRequest wrapper for the ListLoadBalancers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loadbalancer/ListLoadBalancers.go.html to see an example of how to use ListLoadBalancersRequest.
type ListLoadBalancersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the load balancers to list.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int64 `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `3`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The level of detail to return for each result. Can be `full` or `simple`.
	// Example: `full`
	Detail *string `mandatory:"false" contributesTo:"query" name:"detail"`

	// The field to sort by.  You can provide one sort order (`sortOrder`). Default order for TIMECREATED is descending.
	// Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListLoadBalancersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order is case sensitive.
	SortOrder ListLoadBalancersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given display name exactly.
	// Example: `example_load_balancer`
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state.
	// Example: `SUCCEEDED`
	LifecycleState LoadBalancerLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLoadBalancersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLoadBalancersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLoadBalancersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLoadBalancersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLoadBalancersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLoadBalancersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListLoadBalancersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListLoadBalancersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLoadBalancersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLoadBalancerLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetLoadBalancerLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLoadBalancersResponse wrapper for the ListLoadBalancers operation
type ListLoadBalancersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []LoadBalancer instances
	Items []LoadBalancer `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListLoadBalancersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLoadBalancersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLoadBalancersSortByEnum Enum with underlying type: string
type ListLoadBalancersSortByEnum string

// Set of constants representing the allowable values for ListLoadBalancersSortByEnum
const (
	ListLoadBalancersSortByTimecreated ListLoadBalancersSortByEnum = "TIMECREATED"
	ListLoadBalancersSortByDisplayname ListLoadBalancersSortByEnum = "DISPLAYNAME"
)

var mappingListLoadBalancersSortByEnum = map[string]ListLoadBalancersSortByEnum{
	"TIMECREATED": ListLoadBalancersSortByTimecreated,
	"DISPLAYNAME": ListLoadBalancersSortByDisplayname,
}

var mappingListLoadBalancersSortByEnumLowerCase = map[string]ListLoadBalancersSortByEnum{
	"timecreated": ListLoadBalancersSortByTimecreated,
	"displayname": ListLoadBalancersSortByDisplayname,
}

// GetListLoadBalancersSortByEnumValues Enumerates the set of values for ListLoadBalancersSortByEnum
func GetListLoadBalancersSortByEnumValues() []ListLoadBalancersSortByEnum {
	values := make([]ListLoadBalancersSortByEnum, 0)
	for _, v := range mappingListLoadBalancersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListLoadBalancersSortByEnumStringValues Enumerates the set of values in String for ListLoadBalancersSortByEnum
func GetListLoadBalancersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListLoadBalancersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLoadBalancersSortByEnum(val string) (ListLoadBalancersSortByEnum, bool) {
	enum, ok := mappingListLoadBalancersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListLoadBalancersSortOrderEnum Enum with underlying type: string
type ListLoadBalancersSortOrderEnum string

// Set of constants representing the allowable values for ListLoadBalancersSortOrderEnum
const (
	ListLoadBalancersSortOrderAsc  ListLoadBalancersSortOrderEnum = "ASC"
	ListLoadBalancersSortOrderDesc ListLoadBalancersSortOrderEnum = "DESC"
)

var mappingListLoadBalancersSortOrderEnum = map[string]ListLoadBalancersSortOrderEnum{
	"ASC":  ListLoadBalancersSortOrderAsc,
	"DESC": ListLoadBalancersSortOrderDesc,
}

var mappingListLoadBalancersSortOrderEnumLowerCase = map[string]ListLoadBalancersSortOrderEnum{
	"asc":  ListLoadBalancersSortOrderAsc,
	"desc": ListLoadBalancersSortOrderDesc,
}

// GetListLoadBalancersSortOrderEnumValues Enumerates the set of values for ListLoadBalancersSortOrderEnum
func GetListLoadBalancersSortOrderEnumValues() []ListLoadBalancersSortOrderEnum {
	values := make([]ListLoadBalancersSortOrderEnum, 0)
	for _, v := range mappingListLoadBalancersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLoadBalancersSortOrderEnumStringValues Enumerates the set of values in String for ListLoadBalancersSortOrderEnum
func GetListLoadBalancersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLoadBalancersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLoadBalancersSortOrderEnum(val string) (ListLoadBalancersSortOrderEnum, bool) {
	enum, ok := mappingListLoadBalancersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
