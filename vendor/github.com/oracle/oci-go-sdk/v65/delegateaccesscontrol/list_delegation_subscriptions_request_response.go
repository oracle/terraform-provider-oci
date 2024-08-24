// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListDelegationSubscriptionsRequest wrapper for the ListDelegationSubscriptions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListDelegationSubscriptions.go.html to see an example of how to use ListDelegationSubscriptionsRequest.
type ListDelegationSubscriptionsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only Delegation Subscription resources whose lifecycleState matches the given Delegation Subscription lifecycle state.
	LifecycleState DelegationSubscriptionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return Delegation Subscription resources that match the given display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDelegationSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified, timeCreated is default.
	SortBy ListDelegationSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDelegationSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDelegationSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDelegationSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDelegationSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListDelegationSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDelegationSubscriptionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDelegationSubscriptionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDelegationSubscriptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListDelegationSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListDelegationSubscriptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListDelegationSubscriptionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListDelegationSubscriptionsResponse wrapper for the ListDelegationSubscriptions operation
type ListDelegationSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DelegationSubscriptionSummaryCollection instances
	DelegationSubscriptionSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDelegationSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDelegationSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDelegationSubscriptionsSortOrderEnum Enum with underlying type: string
type ListDelegationSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListDelegationSubscriptionsSortOrderEnum
const (
	ListDelegationSubscriptionsSortOrderAsc  ListDelegationSubscriptionsSortOrderEnum = "ASC"
	ListDelegationSubscriptionsSortOrderDesc ListDelegationSubscriptionsSortOrderEnum = "DESC"
)

var mappingListDelegationSubscriptionsSortOrderEnum = map[string]ListDelegationSubscriptionsSortOrderEnum{
	"ASC":  ListDelegationSubscriptionsSortOrderAsc,
	"DESC": ListDelegationSubscriptionsSortOrderDesc,
}

var mappingListDelegationSubscriptionsSortOrderEnumLowerCase = map[string]ListDelegationSubscriptionsSortOrderEnum{
	"asc":  ListDelegationSubscriptionsSortOrderAsc,
	"desc": ListDelegationSubscriptionsSortOrderDesc,
}

// GetListDelegationSubscriptionsSortOrderEnumValues Enumerates the set of values for ListDelegationSubscriptionsSortOrderEnum
func GetListDelegationSubscriptionsSortOrderEnumValues() []ListDelegationSubscriptionsSortOrderEnum {
	values := make([]ListDelegationSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListDelegationSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegationSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListDelegationSubscriptionsSortOrderEnum
func GetListDelegationSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListDelegationSubscriptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegationSubscriptionsSortOrderEnum(val string) (ListDelegationSubscriptionsSortOrderEnum, bool) {
	enum, ok := mappingListDelegationSubscriptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListDelegationSubscriptionsSortByEnum Enum with underlying type: string
type ListDelegationSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListDelegationSubscriptionsSortByEnum
const (
	ListDelegationSubscriptionsSortByTimecreated ListDelegationSubscriptionsSortByEnum = "timeCreated"
	ListDelegationSubscriptionsSortByDisplayname ListDelegationSubscriptionsSortByEnum = "displayName"
)

var mappingListDelegationSubscriptionsSortByEnum = map[string]ListDelegationSubscriptionsSortByEnum{
	"timeCreated": ListDelegationSubscriptionsSortByTimecreated,
	"displayName": ListDelegationSubscriptionsSortByDisplayname,
}

var mappingListDelegationSubscriptionsSortByEnumLowerCase = map[string]ListDelegationSubscriptionsSortByEnum{
	"timecreated": ListDelegationSubscriptionsSortByTimecreated,
	"displayname": ListDelegationSubscriptionsSortByDisplayname,
}

// GetListDelegationSubscriptionsSortByEnumValues Enumerates the set of values for ListDelegationSubscriptionsSortByEnum
func GetListDelegationSubscriptionsSortByEnumValues() []ListDelegationSubscriptionsSortByEnum {
	values := make([]ListDelegationSubscriptionsSortByEnum, 0)
	for _, v := range mappingListDelegationSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListDelegationSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListDelegationSubscriptionsSortByEnum
func GetListDelegationSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListDelegationSubscriptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListDelegationSubscriptionsSortByEnum(val string) (ListDelegationSubscriptionsSortByEnum, bool) {
	enum, ok := mappingListDelegationSubscriptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
