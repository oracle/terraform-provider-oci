// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListRedemptionsRequest wrapper for the ListRedemptions operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListRedemptions.go.html to see an example of how to use ListRedemptionsRequest.
type ListRedemptionsRequest struct {

	// The OCID of the tenancy.
	TenancyId *string `mandatory:"true" contributesTo:"query" name:"tenancyId"`

	// The subscription ID for which rewards information is requested for.
	SubscriptionId *string `mandatory:"true" contributesTo:"path" name:"subscriptionId"`

	// The starting redeemed date filter for the redemption history.
	TimeRedeemedGreaterThanOrEqualTo *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeRedeemedGreaterThanOrEqualTo"`

	// The ending redeemed date filter for the redemption history.
	TimeRedeemedLessThan *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeRedeemedLessThan"`

	// Unique, Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the 'opc-next-page' response header from the previous call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, which can be ascending (ASC) or descending (DESC).
	SortOrder ListRedemptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to be used only for list redemptions API. Supports one sort order.
	SortBy ListRedemptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRedemptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRedemptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRedemptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRedemptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRedemptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRedemptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRedemptionsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRedemptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRedemptionsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRedemptionsResponse wrapper for the ListRedemptions operation
type ListRedemptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RedemptionCollection instances
	RedemptionCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRedemptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRedemptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRedemptionsSortOrderEnum Enum with underlying type: string
type ListRedemptionsSortOrderEnum string

// Set of constants representing the allowable values for ListRedemptionsSortOrderEnum
const (
	ListRedemptionsSortOrderAsc  ListRedemptionsSortOrderEnum = "ASC"
	ListRedemptionsSortOrderDesc ListRedemptionsSortOrderEnum = "DESC"
)

var mappingListRedemptionsSortOrderEnum = map[string]ListRedemptionsSortOrderEnum{
	"ASC":  ListRedemptionsSortOrderAsc,
	"DESC": ListRedemptionsSortOrderDesc,
}

var mappingListRedemptionsSortOrderEnumLowerCase = map[string]ListRedemptionsSortOrderEnum{
	"asc":  ListRedemptionsSortOrderAsc,
	"desc": ListRedemptionsSortOrderDesc,
}

// GetListRedemptionsSortOrderEnumValues Enumerates the set of values for ListRedemptionsSortOrderEnum
func GetListRedemptionsSortOrderEnumValues() []ListRedemptionsSortOrderEnum {
	values := make([]ListRedemptionsSortOrderEnum, 0)
	for _, v := range mappingListRedemptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRedemptionsSortOrderEnumStringValues Enumerates the set of values in String for ListRedemptionsSortOrderEnum
func GetListRedemptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRedemptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRedemptionsSortOrderEnum(val string) (ListRedemptionsSortOrderEnum, bool) {
	enum, ok := mappingListRedemptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRedemptionsSortByEnum Enum with underlying type: string
type ListRedemptionsSortByEnum string

// Set of constants representing the allowable values for ListRedemptionsSortByEnum
const (
	ListRedemptionsSortByTimeredeemed ListRedemptionsSortByEnum = "TIMEREDEEMED"
)

var mappingListRedemptionsSortByEnum = map[string]ListRedemptionsSortByEnum{
	"TIMEREDEEMED": ListRedemptionsSortByTimeredeemed,
}

var mappingListRedemptionsSortByEnumLowerCase = map[string]ListRedemptionsSortByEnum{
	"timeredeemed": ListRedemptionsSortByTimeredeemed,
}

// GetListRedemptionsSortByEnumValues Enumerates the set of values for ListRedemptionsSortByEnum
func GetListRedemptionsSortByEnumValues() []ListRedemptionsSortByEnum {
	values := make([]ListRedemptionsSortByEnum, 0)
	for _, v := range mappingListRedemptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRedemptionsSortByEnumStringValues Enumerates the set of values in String for ListRedemptionsSortByEnum
func GetListRedemptionsSortByEnumStringValues() []string {
	return []string{
		"TIMEREDEEMED",
	}
}

// GetMappingListRedemptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRedemptionsSortByEnum(val string) (ListRedemptionsSortByEnum, bool) {
	enum, ok := mappingListRedemptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
