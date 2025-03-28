// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProductsRequest wrapper for the ListProducts operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListProducts.go.html to see an example of how to use ListProductsRequest.
type ListProductsRequest struct {

	// The OCID of the tenancy.
	TenancyId *string `mandatory:"true" contributesTo:"query" name:"tenancyId"`

	// The subscription ID for which rewards information is requested for.
	SubscriptionId *string `mandatory:"true" contributesTo:"path" name:"subscriptionId"`

	// The SPM Identifier for the usage period.
	UsagePeriodKey *string `mandatory:"true" contributesTo:"query" name:"usagePeriodKey"`

	// Unique, Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the 'opc-next-page' response header from the previous call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, which can be ascending (ASC) or descending (DESC).
	SortOrder ListProductsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Supports one sort order.
	SortBy ListProductsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The field to specify the type of product.
	Producttype ListProductsProducttypeEnum `mandatory:"false" contributesTo:"query" name:"producttype" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProductsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProductsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProductsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProductsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProductsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProductsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProductsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProductsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProductsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProductsProducttypeEnum(string(request.Producttype)); !ok && request.Producttype != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Producttype: %s. Supported values are: %s.", request.Producttype, strings.Join(GetListProductsProducttypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProductsResponse wrapper for the ListProducts operation
type ListProductsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ProductCollection instances
	ProductCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListProductsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProductsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProductsSortOrderEnum Enum with underlying type: string
type ListProductsSortOrderEnum string

// Set of constants representing the allowable values for ListProductsSortOrderEnum
const (
	ListProductsSortOrderAsc  ListProductsSortOrderEnum = "ASC"
	ListProductsSortOrderDesc ListProductsSortOrderEnum = "DESC"
)

var mappingListProductsSortOrderEnum = map[string]ListProductsSortOrderEnum{
	"ASC":  ListProductsSortOrderAsc,
	"DESC": ListProductsSortOrderDesc,
}

var mappingListProductsSortOrderEnumLowerCase = map[string]ListProductsSortOrderEnum{
	"asc":  ListProductsSortOrderAsc,
	"desc": ListProductsSortOrderDesc,
}

// GetListProductsSortOrderEnumValues Enumerates the set of values for ListProductsSortOrderEnum
func GetListProductsSortOrderEnumValues() []ListProductsSortOrderEnum {
	values := make([]ListProductsSortOrderEnum, 0)
	for _, v := range mappingListProductsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProductsSortOrderEnumStringValues Enumerates the set of values in String for ListProductsSortOrderEnum
func GetListProductsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProductsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProductsSortOrderEnum(val string) (ListProductsSortOrderEnum, bool) {
	enum, ok := mappingListProductsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProductsSortByEnum Enum with underlying type: string
type ListProductsSortByEnum string

// Set of constants representing the allowable values for ListProductsSortByEnum
const (
	ListProductsSortByTimecreated ListProductsSortByEnum = "TIMECREATED"
	ListProductsSortByTimestart   ListProductsSortByEnum = "TIMESTART"
)

var mappingListProductsSortByEnum = map[string]ListProductsSortByEnum{
	"TIMECREATED": ListProductsSortByTimecreated,
	"TIMESTART":   ListProductsSortByTimestart,
}

var mappingListProductsSortByEnumLowerCase = map[string]ListProductsSortByEnum{
	"timecreated": ListProductsSortByTimecreated,
	"timestart":   ListProductsSortByTimestart,
}

// GetListProductsSortByEnumValues Enumerates the set of values for ListProductsSortByEnum
func GetListProductsSortByEnumValues() []ListProductsSortByEnum {
	values := make([]ListProductsSortByEnum, 0)
	for _, v := range mappingListProductsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProductsSortByEnumStringValues Enumerates the set of values in String for ListProductsSortByEnum
func GetListProductsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"TIMESTART",
	}
}

// GetMappingListProductsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProductsSortByEnum(val string) (ListProductsSortByEnum, bool) {
	enum, ok := mappingListProductsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProductsProducttypeEnum Enum with underlying type: string
type ListProductsProducttypeEnum string

// Set of constants representing the allowable values for ListProductsProducttypeEnum
const (
	ListProductsProducttypeAll        ListProductsProducttypeEnum = "ALL"
	ListProductsProducttypeEligible   ListProductsProducttypeEnum = "ELIGIBLE"
	ListProductsProducttypeIneligible ListProductsProducttypeEnum = "INELIGIBLE"
)

var mappingListProductsProducttypeEnum = map[string]ListProductsProducttypeEnum{
	"ALL":        ListProductsProducttypeAll,
	"ELIGIBLE":   ListProductsProducttypeEligible,
	"INELIGIBLE": ListProductsProducttypeIneligible,
}

var mappingListProductsProducttypeEnumLowerCase = map[string]ListProductsProducttypeEnum{
	"all":        ListProductsProducttypeAll,
	"eligible":   ListProductsProducttypeEligible,
	"ineligible": ListProductsProducttypeIneligible,
}

// GetListProductsProducttypeEnumValues Enumerates the set of values for ListProductsProducttypeEnum
func GetListProductsProducttypeEnumValues() []ListProductsProducttypeEnum {
	values := make([]ListProductsProducttypeEnum, 0)
	for _, v := range mappingListProductsProducttypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListProductsProducttypeEnumStringValues Enumerates the set of values in String for ListProductsProducttypeEnum
func GetListProductsProducttypeEnumStringValues() []string {
	return []string{
		"ALL",
		"ELIGIBLE",
		"INELIGIBLE",
	}
}

// GetMappingListProductsProducttypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProductsProducttypeEnum(val string) (ListProductsProducttypeEnum, bool) {
	enum, ok := mappingListProductsProducttypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
