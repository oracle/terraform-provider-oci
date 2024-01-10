// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListNetworkFirewallPoliciesRequest wrapper for the ListNetworkFirewallPolicies operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListNetworkFirewallPolicies.go.html to see an example of how to use ListNetworkFirewallPoliciesRequest.
type ListNetworkFirewallPoliciesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Network Firewall Policy identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` or `opc-prev-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources with a lifecycleState matching the given value.
	LifecycleState ListNetworkFirewallPoliciesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListNetworkFirewallPoliciesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListNetworkFirewallPoliciesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkFirewallPoliciesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkFirewallPoliciesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkFirewallPoliciesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkFirewallPoliciesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNetworkFirewallPoliciesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNetworkFirewallPoliciesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListNetworkFirewallPoliciesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkFirewallPoliciesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNetworkFirewallPoliciesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkFirewallPoliciesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNetworkFirewallPoliciesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNetworkFirewallPoliciesResponse wrapper for the ListNetworkFirewallPolicies operation
type ListNetworkFirewallPoliciesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NetworkFirewallPolicySummaryCollection instances
	NetworkFirewallPolicySummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. This is to get the page counts overall.
	OpcPageCount *string `presentIn:"header" name:"opc-page-count"`

	// For pagination of a list of items. This provides the count of total items across pages.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListNetworkFirewallPoliciesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkFirewallPoliciesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkFirewallPoliciesLifecycleStateEnum Enum with underlying type: string
type ListNetworkFirewallPoliciesLifecycleStateEnum string

// Set of constants representing the allowable values for ListNetworkFirewallPoliciesLifecycleStateEnum
const (
	ListNetworkFirewallPoliciesLifecycleStateCreating       ListNetworkFirewallPoliciesLifecycleStateEnum = "CREATING"
	ListNetworkFirewallPoliciesLifecycleStateUpdating       ListNetworkFirewallPoliciesLifecycleStateEnum = "UPDATING"
	ListNetworkFirewallPoliciesLifecycleStateActive         ListNetworkFirewallPoliciesLifecycleStateEnum = "ACTIVE"
	ListNetworkFirewallPoliciesLifecycleStateDeleting       ListNetworkFirewallPoliciesLifecycleStateEnum = "DELETING"
	ListNetworkFirewallPoliciesLifecycleStateDeleted        ListNetworkFirewallPoliciesLifecycleStateEnum = "DELETED"
	ListNetworkFirewallPoliciesLifecycleStateFailed         ListNetworkFirewallPoliciesLifecycleStateEnum = "FAILED"
	ListNetworkFirewallPoliciesLifecycleStateNeedsAttention ListNetworkFirewallPoliciesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListNetworkFirewallPoliciesLifecycleStateAttaching      ListNetworkFirewallPoliciesLifecycleStateEnum = "ATTACHING"
	ListNetworkFirewallPoliciesLifecycleStateDetaching      ListNetworkFirewallPoliciesLifecycleStateEnum = "DETACHING"
)

var mappingListNetworkFirewallPoliciesLifecycleStateEnum = map[string]ListNetworkFirewallPoliciesLifecycleStateEnum{
	"CREATING":        ListNetworkFirewallPoliciesLifecycleStateCreating,
	"UPDATING":        ListNetworkFirewallPoliciesLifecycleStateUpdating,
	"ACTIVE":          ListNetworkFirewallPoliciesLifecycleStateActive,
	"DELETING":        ListNetworkFirewallPoliciesLifecycleStateDeleting,
	"DELETED":         ListNetworkFirewallPoliciesLifecycleStateDeleted,
	"FAILED":          ListNetworkFirewallPoliciesLifecycleStateFailed,
	"NEEDS_ATTENTION": ListNetworkFirewallPoliciesLifecycleStateNeedsAttention,
	"ATTACHING":       ListNetworkFirewallPoliciesLifecycleStateAttaching,
	"DETACHING":       ListNetworkFirewallPoliciesLifecycleStateDetaching,
}

var mappingListNetworkFirewallPoliciesLifecycleStateEnumLowerCase = map[string]ListNetworkFirewallPoliciesLifecycleStateEnum{
	"creating":        ListNetworkFirewallPoliciesLifecycleStateCreating,
	"updating":        ListNetworkFirewallPoliciesLifecycleStateUpdating,
	"active":          ListNetworkFirewallPoliciesLifecycleStateActive,
	"deleting":        ListNetworkFirewallPoliciesLifecycleStateDeleting,
	"deleted":         ListNetworkFirewallPoliciesLifecycleStateDeleted,
	"failed":          ListNetworkFirewallPoliciesLifecycleStateFailed,
	"needs_attention": ListNetworkFirewallPoliciesLifecycleStateNeedsAttention,
	"attaching":       ListNetworkFirewallPoliciesLifecycleStateAttaching,
	"detaching":       ListNetworkFirewallPoliciesLifecycleStateDetaching,
}

// GetListNetworkFirewallPoliciesLifecycleStateEnumValues Enumerates the set of values for ListNetworkFirewallPoliciesLifecycleStateEnum
func GetListNetworkFirewallPoliciesLifecycleStateEnumValues() []ListNetworkFirewallPoliciesLifecycleStateEnum {
	values := make([]ListNetworkFirewallPoliciesLifecycleStateEnum, 0)
	for _, v := range mappingListNetworkFirewallPoliciesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkFirewallPoliciesLifecycleStateEnumStringValues Enumerates the set of values in String for ListNetworkFirewallPoliciesLifecycleStateEnum
func GetListNetworkFirewallPoliciesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
		"ATTACHING",
		"DETACHING",
	}
}

// GetMappingListNetworkFirewallPoliciesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkFirewallPoliciesLifecycleStateEnum(val string) (ListNetworkFirewallPoliciesLifecycleStateEnum, bool) {
	enum, ok := mappingListNetworkFirewallPoliciesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkFirewallPoliciesSortOrderEnum Enum with underlying type: string
type ListNetworkFirewallPoliciesSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkFirewallPoliciesSortOrderEnum
const (
	ListNetworkFirewallPoliciesSortOrderAsc  ListNetworkFirewallPoliciesSortOrderEnum = "ASC"
	ListNetworkFirewallPoliciesSortOrderDesc ListNetworkFirewallPoliciesSortOrderEnum = "DESC"
)

var mappingListNetworkFirewallPoliciesSortOrderEnum = map[string]ListNetworkFirewallPoliciesSortOrderEnum{
	"ASC":  ListNetworkFirewallPoliciesSortOrderAsc,
	"DESC": ListNetworkFirewallPoliciesSortOrderDesc,
}

var mappingListNetworkFirewallPoliciesSortOrderEnumLowerCase = map[string]ListNetworkFirewallPoliciesSortOrderEnum{
	"asc":  ListNetworkFirewallPoliciesSortOrderAsc,
	"desc": ListNetworkFirewallPoliciesSortOrderDesc,
}

// GetListNetworkFirewallPoliciesSortOrderEnumValues Enumerates the set of values for ListNetworkFirewallPoliciesSortOrderEnum
func GetListNetworkFirewallPoliciesSortOrderEnumValues() []ListNetworkFirewallPoliciesSortOrderEnum {
	values := make([]ListNetworkFirewallPoliciesSortOrderEnum, 0)
	for _, v := range mappingListNetworkFirewallPoliciesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkFirewallPoliciesSortOrderEnumStringValues Enumerates the set of values in String for ListNetworkFirewallPoliciesSortOrderEnum
func GetListNetworkFirewallPoliciesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNetworkFirewallPoliciesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkFirewallPoliciesSortOrderEnum(val string) (ListNetworkFirewallPoliciesSortOrderEnum, bool) {
	enum, ok := mappingListNetworkFirewallPoliciesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkFirewallPoliciesSortByEnum Enum with underlying type: string
type ListNetworkFirewallPoliciesSortByEnum string

// Set of constants representing the allowable values for ListNetworkFirewallPoliciesSortByEnum
const (
	ListNetworkFirewallPoliciesSortByTimecreated ListNetworkFirewallPoliciesSortByEnum = "timeCreated"
	ListNetworkFirewallPoliciesSortByDisplayname ListNetworkFirewallPoliciesSortByEnum = "displayName"
)

var mappingListNetworkFirewallPoliciesSortByEnum = map[string]ListNetworkFirewallPoliciesSortByEnum{
	"timeCreated": ListNetworkFirewallPoliciesSortByTimecreated,
	"displayName": ListNetworkFirewallPoliciesSortByDisplayname,
}

var mappingListNetworkFirewallPoliciesSortByEnumLowerCase = map[string]ListNetworkFirewallPoliciesSortByEnum{
	"timecreated": ListNetworkFirewallPoliciesSortByTimecreated,
	"displayname": ListNetworkFirewallPoliciesSortByDisplayname,
}

// GetListNetworkFirewallPoliciesSortByEnumValues Enumerates the set of values for ListNetworkFirewallPoliciesSortByEnum
func GetListNetworkFirewallPoliciesSortByEnumValues() []ListNetworkFirewallPoliciesSortByEnum {
	values := make([]ListNetworkFirewallPoliciesSortByEnum, 0)
	for _, v := range mappingListNetworkFirewallPoliciesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkFirewallPoliciesSortByEnumStringValues Enumerates the set of values in String for ListNetworkFirewallPoliciesSortByEnum
func GetListNetworkFirewallPoliciesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNetworkFirewallPoliciesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkFirewallPoliciesSortByEnum(val string) (ListNetworkFirewallPoliciesSortByEnum, bool) {
	enum, ok := mappingListNetworkFirewallPoliciesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
