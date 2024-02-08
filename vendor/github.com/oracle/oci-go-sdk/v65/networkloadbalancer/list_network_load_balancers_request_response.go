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

// ListNetworkLoadBalancersRequest wrapper for the ListNetworkLoadBalancers operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkloadbalancer/ListNetworkLoadBalancers.go.html to see an example of how to use ListNetworkLoadBalancersRequest.
type ListNetworkLoadBalancersRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the network load balancers to list.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given lifecycle state.
	LifecycleState ListNetworkLoadBalancersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// For list pagination. The maximum number of results per page or items to return, in a paginated "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from which to start retrieving results.
	// For list pagination. The value of the `opc-next-page` response header from the previous "List" call.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' (ascending) or 'desc' (descending).
	SortOrder ListNetworkLoadBalancersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. The default order for timeCreated is descending.
	// The default order for displayName is ascending. If no value is specified, then timeCreated is the default.
	SortBy ListNetworkLoadBalancersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you must contact Oracle about a
	// particular request, then provide the request identifier.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkLoadBalancersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkLoadBalancersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkLoadBalancersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkLoadBalancersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNetworkLoadBalancersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNetworkLoadBalancersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListNetworkLoadBalancersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkLoadBalancersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNetworkLoadBalancersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkLoadBalancersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNetworkLoadBalancersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNetworkLoadBalancersResponse wrapper for the ListNetworkLoadBalancers operation
type ListNetworkLoadBalancersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NetworkLoadBalancerCollection instances
	NetworkLoadBalancerCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you must contact
	// Oracle about a particular request, then provide the request identifier.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListNetworkLoadBalancersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkLoadBalancersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkLoadBalancersLifecycleStateEnum Enum with underlying type: string
type ListNetworkLoadBalancersLifecycleStateEnum string

// Set of constants representing the allowable values for ListNetworkLoadBalancersLifecycleStateEnum
const (
	ListNetworkLoadBalancersLifecycleStateCreating ListNetworkLoadBalancersLifecycleStateEnum = "CREATING"
	ListNetworkLoadBalancersLifecycleStateUpdating ListNetworkLoadBalancersLifecycleStateEnum = "UPDATING"
	ListNetworkLoadBalancersLifecycleStateActive   ListNetworkLoadBalancersLifecycleStateEnum = "ACTIVE"
	ListNetworkLoadBalancersLifecycleStateDeleting ListNetworkLoadBalancersLifecycleStateEnum = "DELETING"
	ListNetworkLoadBalancersLifecycleStateDeleted  ListNetworkLoadBalancersLifecycleStateEnum = "DELETED"
	ListNetworkLoadBalancersLifecycleStateFailed   ListNetworkLoadBalancersLifecycleStateEnum = "FAILED"
)

var mappingListNetworkLoadBalancersLifecycleStateEnum = map[string]ListNetworkLoadBalancersLifecycleStateEnum{
	"CREATING": ListNetworkLoadBalancersLifecycleStateCreating,
	"UPDATING": ListNetworkLoadBalancersLifecycleStateUpdating,
	"ACTIVE":   ListNetworkLoadBalancersLifecycleStateActive,
	"DELETING": ListNetworkLoadBalancersLifecycleStateDeleting,
	"DELETED":  ListNetworkLoadBalancersLifecycleStateDeleted,
	"FAILED":   ListNetworkLoadBalancersLifecycleStateFailed,
}

var mappingListNetworkLoadBalancersLifecycleStateEnumLowerCase = map[string]ListNetworkLoadBalancersLifecycleStateEnum{
	"creating": ListNetworkLoadBalancersLifecycleStateCreating,
	"updating": ListNetworkLoadBalancersLifecycleStateUpdating,
	"active":   ListNetworkLoadBalancersLifecycleStateActive,
	"deleting": ListNetworkLoadBalancersLifecycleStateDeleting,
	"deleted":  ListNetworkLoadBalancersLifecycleStateDeleted,
	"failed":   ListNetworkLoadBalancersLifecycleStateFailed,
}

// GetListNetworkLoadBalancersLifecycleStateEnumValues Enumerates the set of values for ListNetworkLoadBalancersLifecycleStateEnum
func GetListNetworkLoadBalancersLifecycleStateEnumValues() []ListNetworkLoadBalancersLifecycleStateEnum {
	values := make([]ListNetworkLoadBalancersLifecycleStateEnum, 0)
	for _, v := range mappingListNetworkLoadBalancersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkLoadBalancersLifecycleStateEnumStringValues Enumerates the set of values in String for ListNetworkLoadBalancersLifecycleStateEnum
func GetListNetworkLoadBalancersLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListNetworkLoadBalancersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkLoadBalancersLifecycleStateEnum(val string) (ListNetworkLoadBalancersLifecycleStateEnum, bool) {
	enum, ok := mappingListNetworkLoadBalancersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkLoadBalancersSortOrderEnum Enum with underlying type: string
type ListNetworkLoadBalancersSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkLoadBalancersSortOrderEnum
const (
	ListNetworkLoadBalancersSortOrderAsc  ListNetworkLoadBalancersSortOrderEnum = "ASC"
	ListNetworkLoadBalancersSortOrderDesc ListNetworkLoadBalancersSortOrderEnum = "DESC"
)

var mappingListNetworkLoadBalancersSortOrderEnum = map[string]ListNetworkLoadBalancersSortOrderEnum{
	"ASC":  ListNetworkLoadBalancersSortOrderAsc,
	"DESC": ListNetworkLoadBalancersSortOrderDesc,
}

var mappingListNetworkLoadBalancersSortOrderEnumLowerCase = map[string]ListNetworkLoadBalancersSortOrderEnum{
	"asc":  ListNetworkLoadBalancersSortOrderAsc,
	"desc": ListNetworkLoadBalancersSortOrderDesc,
}

// GetListNetworkLoadBalancersSortOrderEnumValues Enumerates the set of values for ListNetworkLoadBalancersSortOrderEnum
func GetListNetworkLoadBalancersSortOrderEnumValues() []ListNetworkLoadBalancersSortOrderEnum {
	values := make([]ListNetworkLoadBalancersSortOrderEnum, 0)
	for _, v := range mappingListNetworkLoadBalancersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkLoadBalancersSortOrderEnumStringValues Enumerates the set of values in String for ListNetworkLoadBalancersSortOrderEnum
func GetListNetworkLoadBalancersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNetworkLoadBalancersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkLoadBalancersSortOrderEnum(val string) (ListNetworkLoadBalancersSortOrderEnum, bool) {
	enum, ok := mappingListNetworkLoadBalancersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkLoadBalancersSortByEnum Enum with underlying type: string
type ListNetworkLoadBalancersSortByEnum string

// Set of constants representing the allowable values for ListNetworkLoadBalancersSortByEnum
const (
	ListNetworkLoadBalancersSortByTimecreated ListNetworkLoadBalancersSortByEnum = "timeCreated"
	ListNetworkLoadBalancersSortByDisplayname ListNetworkLoadBalancersSortByEnum = "displayName"
)

var mappingListNetworkLoadBalancersSortByEnum = map[string]ListNetworkLoadBalancersSortByEnum{
	"timeCreated": ListNetworkLoadBalancersSortByTimecreated,
	"displayName": ListNetworkLoadBalancersSortByDisplayname,
}

var mappingListNetworkLoadBalancersSortByEnumLowerCase = map[string]ListNetworkLoadBalancersSortByEnum{
	"timecreated": ListNetworkLoadBalancersSortByTimecreated,
	"displayname": ListNetworkLoadBalancersSortByDisplayname,
}

// GetListNetworkLoadBalancersSortByEnumValues Enumerates the set of values for ListNetworkLoadBalancersSortByEnum
func GetListNetworkLoadBalancersSortByEnumValues() []ListNetworkLoadBalancersSortByEnum {
	values := make([]ListNetworkLoadBalancersSortByEnum, 0)
	for _, v := range mappingListNetworkLoadBalancersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkLoadBalancersSortByEnumStringValues Enumerates the set of values in String for ListNetworkLoadBalancersSortByEnum
func GetListNetworkLoadBalancersSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNetworkLoadBalancersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkLoadBalancersSortByEnum(val string) (ListNetworkLoadBalancersSortByEnum, bool) {
	enum, ok := mappingListNetworkLoadBalancersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
