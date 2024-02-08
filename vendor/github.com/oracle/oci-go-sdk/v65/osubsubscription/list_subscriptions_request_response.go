// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package osubsubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSubscriptionsRequest wrapper for the ListSubscriptions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/osubsubscription/ListSubscriptions.go.html to see an example of how to use ListSubscriptionsRequest.
type ListSubscriptionsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The Plan Number
	PlanNumber *string `mandatory:"false" contributesTo:"query" name:"planNumber"`

	// Line level Subscription Id
	SubscriptionId *string `mandatory:"false" contributesTo:"query" name:"subscriptionId"`

	// Buyer Email Id
	BuyerEmail *string `mandatory:"false" contributesTo:"query" name:"buyerEmail"`

	// Boolean value to decide whether commitment services will be shown
	IsCommitInfoRequired *bool `mandatory:"false" contributesTo:"query" name:"isCommitInfoRequired"`

	// The maximum number of items to return in a paginated "List" call. Default: (`50`)
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`).
	SortBy ListSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// This header is meant to be used only for internal purposes and will be ignored on any public request. The purpose of this header is
	// to help on Gateway to API calls identification.
	XOneGatewaySubscriptionId *string `mandatory:"false" contributesTo:"header" name:"x-one-gateway-subscription-id"`

	// The OCI home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc.
	XOneOriginRegion *string `mandatory:"false" contributesTo:"header" name:"x-one-origin-region"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSubscriptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSubscriptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSubscriptionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSubscriptionsResponse wrapper for the ListSubscriptions operation
type ListSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []SubscriptionSummary instances
	Items []SubscriptionSummary `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSubscriptionsSortOrderEnum Enum with underlying type: string
type ListSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListSubscriptionsSortOrderEnum
const (
	ListSubscriptionsSortOrderAsc  ListSubscriptionsSortOrderEnum = "ASC"
	ListSubscriptionsSortOrderDesc ListSubscriptionsSortOrderEnum = "DESC"
)

var mappingListSubscriptionsSortOrderEnum = map[string]ListSubscriptionsSortOrderEnum{
	"ASC":  ListSubscriptionsSortOrderAsc,
	"DESC": ListSubscriptionsSortOrderDesc,
}

var mappingListSubscriptionsSortOrderEnumLowerCase = map[string]ListSubscriptionsSortOrderEnum{
	"asc":  ListSubscriptionsSortOrderAsc,
	"desc": ListSubscriptionsSortOrderDesc,
}

// GetListSubscriptionsSortOrderEnumValues Enumerates the set of values for ListSubscriptionsSortOrderEnum
func GetListSubscriptionsSortOrderEnumValues() []ListSubscriptionsSortOrderEnum {
	values := make([]ListSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListSubscriptionsSortOrderEnum
func GetListSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSubscriptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscriptionsSortOrderEnum(val string) (ListSubscriptionsSortOrderEnum, bool) {
	enum, ok := mappingListSubscriptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSubscriptionsSortByEnum Enum with underlying type: string
type ListSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListSubscriptionsSortByEnum
const (
	ListSubscriptionsSortByTimecreated ListSubscriptionsSortByEnum = "TIMECREATED"
	ListSubscriptionsSortByTimestart   ListSubscriptionsSortByEnum = "TIMESTART"
)

var mappingListSubscriptionsSortByEnum = map[string]ListSubscriptionsSortByEnum{
	"TIMECREATED": ListSubscriptionsSortByTimecreated,
	"TIMESTART":   ListSubscriptionsSortByTimestart,
}

var mappingListSubscriptionsSortByEnumLowerCase = map[string]ListSubscriptionsSortByEnum{
	"timecreated": ListSubscriptionsSortByTimecreated,
	"timestart":   ListSubscriptionsSortByTimestart,
}

// GetListSubscriptionsSortByEnumValues Enumerates the set of values for ListSubscriptionsSortByEnum
func GetListSubscriptionsSortByEnumValues() []ListSubscriptionsSortByEnum {
	values := make([]ListSubscriptionsSortByEnum, 0)
	for _, v := range mappingListSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListSubscriptionsSortByEnum
func GetListSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"TIMESTART",
	}
}

// GetMappingListSubscriptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSubscriptionsSortByEnum(val string) (ListSubscriptionsSortByEnum, bool) {
	enum, ok := mappingListSubscriptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
