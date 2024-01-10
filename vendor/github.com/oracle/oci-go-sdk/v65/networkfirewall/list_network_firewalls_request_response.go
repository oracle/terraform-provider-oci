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

// ListNetworkFirewallsRequest wrapper for the ListNetworkFirewalls operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/networkfirewall/ListNetworkFirewalls.go.html to see an example of how to use ListNetworkFirewallsRequest.
type ListNetworkFirewallsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the entire networkFirewallPolicyId given.
	NetworkFirewallPolicyId *string `mandatory:"false" contributesTo:"query" name:"networkFirewallPolicyId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Network Firewall resource.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that are present within the specified availability domain.
	// To get a list of availability domains for a tenancy, use ListAvailabilityDomains operation.
	// Example: `kIdk:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` or `opc-prev-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources with a lifecycleState matching the given value.
	LifecycleState ListNetworkFirewallsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListNetworkFirewallsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListNetworkFirewallsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNetworkFirewallsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNetworkFirewallsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListNetworkFirewallsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNetworkFirewallsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListNetworkFirewallsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListNetworkFirewallsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListNetworkFirewallsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkFirewallsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListNetworkFirewallsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListNetworkFirewallsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListNetworkFirewallsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListNetworkFirewallsResponse wrapper for the ListNetworkFirewalls operation
type ListNetworkFirewallsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of NetworkFirewallCollection instances
	NetworkFirewallCollection `presentIn:"body"`

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

func (response ListNetworkFirewallsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNetworkFirewallsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNetworkFirewallsLifecycleStateEnum Enum with underlying type: string
type ListNetworkFirewallsLifecycleStateEnum string

// Set of constants representing the allowable values for ListNetworkFirewallsLifecycleStateEnum
const (
	ListNetworkFirewallsLifecycleStateCreating       ListNetworkFirewallsLifecycleStateEnum = "CREATING"
	ListNetworkFirewallsLifecycleStateUpdating       ListNetworkFirewallsLifecycleStateEnum = "UPDATING"
	ListNetworkFirewallsLifecycleStateActive         ListNetworkFirewallsLifecycleStateEnum = "ACTIVE"
	ListNetworkFirewallsLifecycleStateDeleting       ListNetworkFirewallsLifecycleStateEnum = "DELETING"
	ListNetworkFirewallsLifecycleStateDeleted        ListNetworkFirewallsLifecycleStateEnum = "DELETED"
	ListNetworkFirewallsLifecycleStateFailed         ListNetworkFirewallsLifecycleStateEnum = "FAILED"
	ListNetworkFirewallsLifecycleStateNeedsAttention ListNetworkFirewallsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListNetworkFirewallsLifecycleStateAttaching      ListNetworkFirewallsLifecycleStateEnum = "ATTACHING"
	ListNetworkFirewallsLifecycleStateDetaching      ListNetworkFirewallsLifecycleStateEnum = "DETACHING"
)

var mappingListNetworkFirewallsLifecycleStateEnum = map[string]ListNetworkFirewallsLifecycleStateEnum{
	"CREATING":        ListNetworkFirewallsLifecycleStateCreating,
	"UPDATING":        ListNetworkFirewallsLifecycleStateUpdating,
	"ACTIVE":          ListNetworkFirewallsLifecycleStateActive,
	"DELETING":        ListNetworkFirewallsLifecycleStateDeleting,
	"DELETED":         ListNetworkFirewallsLifecycleStateDeleted,
	"FAILED":          ListNetworkFirewallsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListNetworkFirewallsLifecycleStateNeedsAttention,
	"ATTACHING":       ListNetworkFirewallsLifecycleStateAttaching,
	"DETACHING":       ListNetworkFirewallsLifecycleStateDetaching,
}

var mappingListNetworkFirewallsLifecycleStateEnumLowerCase = map[string]ListNetworkFirewallsLifecycleStateEnum{
	"creating":        ListNetworkFirewallsLifecycleStateCreating,
	"updating":        ListNetworkFirewallsLifecycleStateUpdating,
	"active":          ListNetworkFirewallsLifecycleStateActive,
	"deleting":        ListNetworkFirewallsLifecycleStateDeleting,
	"deleted":         ListNetworkFirewallsLifecycleStateDeleted,
	"failed":          ListNetworkFirewallsLifecycleStateFailed,
	"needs_attention": ListNetworkFirewallsLifecycleStateNeedsAttention,
	"attaching":       ListNetworkFirewallsLifecycleStateAttaching,
	"detaching":       ListNetworkFirewallsLifecycleStateDetaching,
}

// GetListNetworkFirewallsLifecycleStateEnumValues Enumerates the set of values for ListNetworkFirewallsLifecycleStateEnum
func GetListNetworkFirewallsLifecycleStateEnumValues() []ListNetworkFirewallsLifecycleStateEnum {
	values := make([]ListNetworkFirewallsLifecycleStateEnum, 0)
	for _, v := range mappingListNetworkFirewallsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkFirewallsLifecycleStateEnumStringValues Enumerates the set of values in String for ListNetworkFirewallsLifecycleStateEnum
func GetListNetworkFirewallsLifecycleStateEnumStringValues() []string {
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

// GetMappingListNetworkFirewallsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkFirewallsLifecycleStateEnum(val string) (ListNetworkFirewallsLifecycleStateEnum, bool) {
	enum, ok := mappingListNetworkFirewallsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkFirewallsSortOrderEnum Enum with underlying type: string
type ListNetworkFirewallsSortOrderEnum string

// Set of constants representing the allowable values for ListNetworkFirewallsSortOrderEnum
const (
	ListNetworkFirewallsSortOrderAsc  ListNetworkFirewallsSortOrderEnum = "ASC"
	ListNetworkFirewallsSortOrderDesc ListNetworkFirewallsSortOrderEnum = "DESC"
)

var mappingListNetworkFirewallsSortOrderEnum = map[string]ListNetworkFirewallsSortOrderEnum{
	"ASC":  ListNetworkFirewallsSortOrderAsc,
	"DESC": ListNetworkFirewallsSortOrderDesc,
}

var mappingListNetworkFirewallsSortOrderEnumLowerCase = map[string]ListNetworkFirewallsSortOrderEnum{
	"asc":  ListNetworkFirewallsSortOrderAsc,
	"desc": ListNetworkFirewallsSortOrderDesc,
}

// GetListNetworkFirewallsSortOrderEnumValues Enumerates the set of values for ListNetworkFirewallsSortOrderEnum
func GetListNetworkFirewallsSortOrderEnumValues() []ListNetworkFirewallsSortOrderEnum {
	values := make([]ListNetworkFirewallsSortOrderEnum, 0)
	for _, v := range mappingListNetworkFirewallsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkFirewallsSortOrderEnumStringValues Enumerates the set of values in String for ListNetworkFirewallsSortOrderEnum
func GetListNetworkFirewallsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListNetworkFirewallsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkFirewallsSortOrderEnum(val string) (ListNetworkFirewallsSortOrderEnum, bool) {
	enum, ok := mappingListNetworkFirewallsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListNetworkFirewallsSortByEnum Enum with underlying type: string
type ListNetworkFirewallsSortByEnum string

// Set of constants representing the allowable values for ListNetworkFirewallsSortByEnum
const (
	ListNetworkFirewallsSortByTimecreated ListNetworkFirewallsSortByEnum = "timeCreated"
	ListNetworkFirewallsSortByDisplayname ListNetworkFirewallsSortByEnum = "displayName"
)

var mappingListNetworkFirewallsSortByEnum = map[string]ListNetworkFirewallsSortByEnum{
	"timeCreated": ListNetworkFirewallsSortByTimecreated,
	"displayName": ListNetworkFirewallsSortByDisplayname,
}

var mappingListNetworkFirewallsSortByEnumLowerCase = map[string]ListNetworkFirewallsSortByEnum{
	"timecreated": ListNetworkFirewallsSortByTimecreated,
	"displayname": ListNetworkFirewallsSortByDisplayname,
}

// GetListNetworkFirewallsSortByEnumValues Enumerates the set of values for ListNetworkFirewallsSortByEnum
func GetListNetworkFirewallsSortByEnumValues() []ListNetworkFirewallsSortByEnum {
	values := make([]ListNetworkFirewallsSortByEnum, 0)
	for _, v := range mappingListNetworkFirewallsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListNetworkFirewallsSortByEnumStringValues Enumerates the set of values in String for ListNetworkFirewallsSortByEnum
func GetListNetworkFirewallsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListNetworkFirewallsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListNetworkFirewallsSortByEnum(val string) (ListNetworkFirewallsSortByEnum, bool) {
	enum, ok := mappingListNetworkFirewallsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
