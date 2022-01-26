// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package usage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListProductsRequest wrapper for the ListProducts operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/usage/ListProducts.go.html to see an example of how to use ListProductsRequest.
type ListProductsRequest struct {

	// The OCID of the tenancy.
	TenancyId *string `mandatory:"true" contributesTo:"query" name:"tenancyId"`

	// The subscriptionId for which rewards information is requested for.
	SubscriptionId *string `mandatory:"true" contributesTo:"path" name:"subscriptionId"`

	// The SPM Identifier for the usage period.
	UsagePeriodKey *string `mandatory:"true" contributesTo:"query" name:"usagePeriodKey"`

	// Unique, Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The value of the 'opc-next-page' response header from the previous call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, can be ascending (ASC) or descending (DESC).
	SortOrder ListProductsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by, supports one sort Order.
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

var mappingListProductsSortOrder = map[string]ListProductsSortOrderEnum{
	"ASC":  ListProductsSortOrderAsc,
	"DESC": ListProductsSortOrderDesc,
}

// GetListProductsSortOrderEnumValues Enumerates the set of values for ListProductsSortOrderEnum
func GetListProductsSortOrderEnumValues() []ListProductsSortOrderEnum {
	values := make([]ListProductsSortOrderEnum, 0)
	for _, v := range mappingListProductsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListProductsSortByEnum Enum with underlying type: string
type ListProductsSortByEnum string

// Set of constants representing the allowable values for ListProductsSortByEnum
const (
	ListProductsSortByTimecreated ListProductsSortByEnum = "TIMECREATED"
	ListProductsSortByTimestart   ListProductsSortByEnum = "TIMESTART"
)

var mappingListProductsSortBy = map[string]ListProductsSortByEnum{
	"TIMECREATED": ListProductsSortByTimecreated,
	"TIMESTART":   ListProductsSortByTimestart,
}

// GetListProductsSortByEnumValues Enumerates the set of values for ListProductsSortByEnum
func GetListProductsSortByEnumValues() []ListProductsSortByEnum {
	values := make([]ListProductsSortByEnum, 0)
	for _, v := range mappingListProductsSortBy {
		values = append(values, v)
	}
	return values
}

// ListProductsProducttypeEnum Enum with underlying type: string
type ListProductsProducttypeEnum string

// Set of constants representing the allowable values for ListProductsProducttypeEnum
const (
	ListProductsProducttypeAll        ListProductsProducttypeEnum = "ALL"
	ListProductsProducttypeEligible   ListProductsProducttypeEnum = "ELIGIBLE"
	ListProductsProducttypeIneligible ListProductsProducttypeEnum = "INELIGIBLE"
)

var mappingListProductsProducttype = map[string]ListProductsProducttypeEnum{
	"ALL":        ListProductsProducttypeAll,
	"ELIGIBLE":   ListProductsProducttypeEligible,
	"INELIGIBLE": ListProductsProducttypeIneligible,
}

// GetListProductsProducttypeEnumValues Enumerates the set of values for ListProductsProducttypeEnum
func GetListProductsProducttypeEnumValues() []ListProductsProducttypeEnum {
	values := make([]ListProductsProducttypeEnum, 0)
	for _, v := range mappingListProductsProducttype {
		values = append(values, v)
	}
	return values
}
