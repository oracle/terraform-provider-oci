// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNetworkLoadBalancersPoliciesRequest wrapper for the ListNetworkLoadBalancersPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkloadbalancer/ListNetworkLoadBalancersPolicies.go.html to see an example of how to use ListNetworkLoadBalancersPoliciesRequest.
type ListNetworkLoadBalancersPoliciesRequest struct {

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

	// The sort order to use, either 'asc' (ascending) or 'desc' (descending).
	SortOrder ListNetworkLoadBalancersPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for timeCreated is descending.
	// The default order for displayName is ascending. If no value is specified, then timeCreated is the default.
	SortBy ListNetworkLoadBalancersPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkLoadBalancersPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkLoadBalancersPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkLoadBalancersPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkLoadBalancersPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNetworkLoadBalancersPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNetworkLoadBalancersPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNetworkLoadBalancersPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkLoadBalancersPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNetworkLoadBalancersPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNetworkLoadBalancersPoliciesResponse wrapper for the ListNetworkLoadBalancersPolicies operation
type ListNetworkLoadBalancersPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NetworkLoadBalancersPolicyCollection instances
	NetworkLoadBalancersPolicyCollection `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListNetworkLoadBalancersPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkLoadBalancersPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkLoadBalancersPoliciesSortOrderEnum Enum with underlying type: string
type ListNetworkLoadBalancersPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkLoadBalancersPoliciesSortOrderEnum
const (
	ListNetworkLoadBalancersPoliciesSortOrderAsc  ListNetworkLoadBalancersPoliciesSortOrderEnum = "ASC"
	ListNetworkLoadBalancersPoliciesSortOrderDesc ListNetworkLoadBalancersPoliciesSortOrderEnum = "DESC"
)

var mappingListNetworkLoadBalancersPoliciesSortOrderEnum = map[string]ListNetworkLoadBalancersPoliciesSortOrderEnum{
	"ASC":  ListNetworkLoadBalancersPoliciesSortOrderAsc,
	"DESC": ListNetworkLoadBalancersPoliciesSortOrderDesc,
}

var mappingListNetworkLoadBalancersPoliciesSortOrderEnumLowerCase = map[string]ListNetworkLoadBalancersPoliciesSortOrderEnum{
	"asc":  ListNetworkLoadBalancersPoliciesSortOrderAsc,
	"desc": ListNetworkLoadBalancersPoliciesSortOrderDesc,
}

// GetListNetworkLoadBalancersPoliciesSortOrderEnumValues Enumerates the set of values for ListNetworkLoadBalancersPoliciesSortOrderEnum
func GetListNetworkLoadBalancersPoliciesSortOrderEnumValues() []ListNetworkLoadBalancersPoliciesSortOrderEnum {
	values := make([]ListNetworkLoadBalancersPoliciesSortOrderEnum, 0)
	for _, v := range mappingListNetworkLoadBalancersPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkLoadBalancersPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListNetworkLoadBalancersPoliciesSortOrderEnum
func GetListNetworkLoadBalancersPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNetworkLoadBalancersPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkLoadBalancersPoliciesSortOrderEnum(val string) (ListNetworkLoadBalancersPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListNetworkLoadBalancersPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkLoadBalancersPoliciesSortByEnum Enum with underlying type: string
type ListNetworkLoadBalancersPoliciesSortByEnum string

// Set of constants representing the allowable values for ListNetworkLoadBalancersPoliciesSortByEnum
const (
	ListNetworkLoadBalancersPoliciesSortByTimecreated ListNetworkLoadBalancersPoliciesSortByEnum = "timeCreated"
	ListNetworkLoadBalancersPoliciesSortByDisplayname ListNetworkLoadBalancersPoliciesSortByEnum = "displayName"
)

var mappingListNetworkLoadBalancersPoliciesSortByEnum = map[string]ListNetworkLoadBalancersPoliciesSortByEnum{
	"timeCreated": ListNetworkLoadBalancersPoliciesSortByTimecreated,
	"displayName": ListNetworkLoadBalancersPoliciesSortByDisplayname,
}

var mappingListNetworkLoadBalancersPoliciesSortByEnumLowerCase = map[string]ListNetworkLoadBalancersPoliciesSortByEnum{
	"timecreated": ListNetworkLoadBalancersPoliciesSortByTimecreated,
	"displayname": ListNetworkLoadBalancersPoliciesSortByDisplayname,
}

// GetListNetworkLoadBalancersPoliciesSortByEnumValues Enumerates the set of values for ListNetworkLoadBalancersPoliciesSortByEnum
func GetListNetworkLoadBalancersPoliciesSortByEnumValues() []ListNetworkLoadBalancersPoliciesSortByEnum {
	values := make([]ListNetworkLoadBalancersPoliciesSortByEnum, 0)
	for _, v := range mappingListNetworkLoadBalancersPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkLoadBalancersPoliciesSortByEnumStringValues Enumerates the set of values in String for ListNetworkLoadBalancersPoliciesSortByEnum
func GetListNetworkLoadBalancersPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNetworkLoadBalancersPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkLoadBalancersPoliciesSortByEnum(val string) (ListNetworkLoadBalancersPoliciesSortByEnum, bool) {
	enum, ok := mappingListNetworkLoadBalancersPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
