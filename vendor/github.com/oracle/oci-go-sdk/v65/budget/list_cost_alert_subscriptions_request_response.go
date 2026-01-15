// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package budget

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCostAlertSubscriptionsRequest wrapper for the ListCostAlertSubscriptions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/budget/ListCostAlertSubscriptions.go.html to see an example of how to use ListCostAlertSubscriptionsRequest.
type ListCostAlertSubscriptionsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCostAlertSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. If not specified, the default is timeCreated.
	// The default sort order for timeCreated is DESC.
	// The default sort order for displayName is ASC in alphanumeric order.
	SortBy ListCostAlertSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique, non-changeable resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The current state of the cost alert subscription.
	LifecycleState CostAlertSubscriptionLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCostAlertSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCostAlertSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCostAlertSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCostAlertSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCostAlertSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCostAlertSubscriptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCostAlertSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCostAlertSubscriptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCostAlertSubscriptionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCostAlertSubscriptionLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCostAlertSubscriptionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCostAlertSubscriptionsResponse wrapper for the ListCostAlertSubscriptions operation
type ListCostAlertSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CostAlertSubscriptionCollection instances
	CostAlertSubscriptionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCostAlertSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCostAlertSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCostAlertSubscriptionsSortOrderEnum Enum with underlying type: string
type ListCostAlertSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListCostAlertSubscriptionsSortOrderEnum
const (
	ListCostAlertSubscriptionsSortOrderAsc  ListCostAlertSubscriptionsSortOrderEnum = "ASC"
	ListCostAlertSubscriptionsSortOrderDesc ListCostAlertSubscriptionsSortOrderEnum = "DESC"
)

var mappingListCostAlertSubscriptionsSortOrderEnum = map[string]ListCostAlertSubscriptionsSortOrderEnum{
	"ASC":  ListCostAlertSubscriptionsSortOrderAsc,
	"DESC": ListCostAlertSubscriptionsSortOrderDesc,
}

var mappingListCostAlertSubscriptionsSortOrderEnumLowerCase = map[string]ListCostAlertSubscriptionsSortOrderEnum{
	"asc":  ListCostAlertSubscriptionsSortOrderAsc,
	"desc": ListCostAlertSubscriptionsSortOrderDesc,
}

// GetListCostAlertSubscriptionsSortOrderEnumValues Enumerates the set of values for ListCostAlertSubscriptionsSortOrderEnum
func GetListCostAlertSubscriptionsSortOrderEnumValues() []ListCostAlertSubscriptionsSortOrderEnum {
	values := make([]ListCostAlertSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListCostAlertSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCostAlertSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListCostAlertSubscriptionsSortOrderEnum
func GetListCostAlertSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCostAlertSubscriptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCostAlertSubscriptionsSortOrderEnum(val string) (ListCostAlertSubscriptionsSortOrderEnum, bool) {
	enum, ok := mappingListCostAlertSubscriptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCostAlertSubscriptionsSortByEnum Enum with underlying type: string
type ListCostAlertSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListCostAlertSubscriptionsSortByEnum
const (
	ListCostAlertSubscriptionsSortByTimecreated ListCostAlertSubscriptionsSortByEnum = "timeCreated"
	ListCostAlertSubscriptionsSortByName        ListCostAlertSubscriptionsSortByEnum = "name"
	ListCostAlertSubscriptionsSortById          ListCostAlertSubscriptionsSortByEnum = "id"
)

var mappingListCostAlertSubscriptionsSortByEnum = map[string]ListCostAlertSubscriptionsSortByEnum{
	"timeCreated": ListCostAlertSubscriptionsSortByTimecreated,
	"name":        ListCostAlertSubscriptionsSortByName,
	"id":          ListCostAlertSubscriptionsSortById,
}

var mappingListCostAlertSubscriptionsSortByEnumLowerCase = map[string]ListCostAlertSubscriptionsSortByEnum{
	"timecreated": ListCostAlertSubscriptionsSortByTimecreated,
	"name":        ListCostAlertSubscriptionsSortByName,
	"id":          ListCostAlertSubscriptionsSortById,
}

// GetListCostAlertSubscriptionsSortByEnumValues Enumerates the set of values for ListCostAlertSubscriptionsSortByEnum
func GetListCostAlertSubscriptionsSortByEnumValues() []ListCostAlertSubscriptionsSortByEnum {
	values := make([]ListCostAlertSubscriptionsSortByEnum, 0)
	for _, v := range mappingListCostAlertSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCostAlertSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListCostAlertSubscriptionsSortByEnum
func GetListCostAlertSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
		"id",
	}
}

// GetMappingListCostAlertSubscriptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCostAlertSubscriptionsSortByEnum(val string) (ListCostAlertSubscriptionsSortByEnum, bool) {
	enum, ok := mappingListCostAlertSubscriptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
