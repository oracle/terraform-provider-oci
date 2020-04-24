// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListListingsRequest wrapper for the ListListings operation
type ListListingsRequest struct {

	// The name of the listing.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// The unique identifier for the listing.
	ListingId *string `mandatory:"false" contributesTo:"query" name:"listingId"`

	// Limit results to just this publisher.
	PublisherId *string `mandatory:"false" contributesTo:"query" name:"publisherId"`

	// A filter to return only packages that match the given package type exactly.
	PackageType *string `mandatory:"false" contributesTo:"query" name:"packageType"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to use to sort listed results. You can only specify one field to sort by.
	// `TIMERELEASED` displays results in descending order by default.
	// You can change your preference by specifying a different sort order.
	SortBy ListListingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either `ASC` or `DESC`.
	SortOrder ListListingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Name of the product category or categories. If you specify multiple categories, then Marketplace returns any listing with
	// one or more matching categories.
	Category []string `contributesTo:"query" name:"category" collectionFormat:"multi"`

	// Name of the pricing type. If multiple pricing types are provided, then any listing with
	// one or more matching pricing models will be returned.
	Pricing []ListListingsPricingEnum `contributesTo:"query" name:"pricing" omitEmpty:"true" collectionFormat:"multi"`

	// Indicates whether to show only featured listings. If this is set to `false` or is omitted, then all listings will be returned.
	IsFeatured *bool `mandatory:"false" contributesTo:"query" name:"isFeatured"`

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListListingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListListingsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListListingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListListingsResponse wrapper for the ListListings operation
type ListListingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ListingSummary instances
	Items []ListingSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListListingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListListingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListListingsSortByEnum Enum with underlying type: string
type ListListingsSortByEnum string

// Set of constants representing the allowable values for ListListingsSortByEnum
const (
	ListListingsSortByTimereleased ListListingsSortByEnum = "TIMERELEASED"
)

var mappingListListingsSortBy = map[string]ListListingsSortByEnum{
	"TIMERELEASED": ListListingsSortByTimereleased,
}

// GetListListingsSortByEnumValues Enumerates the set of values for ListListingsSortByEnum
func GetListListingsSortByEnumValues() []ListListingsSortByEnum {
	values := make([]ListListingsSortByEnum, 0)
	for _, v := range mappingListListingsSortBy {
		values = append(values, v)
	}
	return values
}

// ListListingsSortOrderEnum Enum with underlying type: string
type ListListingsSortOrderEnum string

// Set of constants representing the allowable values for ListListingsSortOrderEnum
const (
	ListListingsSortOrderAsc  ListListingsSortOrderEnum = "ASC"
	ListListingsSortOrderDesc ListListingsSortOrderEnum = "DESC"
)

var mappingListListingsSortOrder = map[string]ListListingsSortOrderEnum{
	"ASC":  ListListingsSortOrderAsc,
	"DESC": ListListingsSortOrderDesc,
}

// GetListListingsSortOrderEnumValues Enumerates the set of values for ListListingsSortOrderEnum
func GetListListingsSortOrderEnumValues() []ListListingsSortOrderEnum {
	values := make([]ListListingsSortOrderEnum, 0)
	for _, v := range mappingListListingsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListListingsPricingEnum Enum with underlying type: string
type ListListingsPricingEnum string

// Set of constants representing the allowable values for ListListingsPricingEnum
const (
	ListListingsPricingFree  ListListingsPricingEnum = "FREE"
	ListListingsPricingByol  ListListingsPricingEnum = "BYOL"
	ListListingsPricingPaygo ListListingsPricingEnum = "PAYGO"
)

var mappingListListingsPricing = map[string]ListListingsPricingEnum{
	"FREE":  ListListingsPricingFree,
	"BYOL":  ListListingsPricingByol,
	"PAYGO": ListListingsPricingPaygo,
}

// GetListListingsPricingEnumValues Enumerates the set of values for ListListingsPricingEnum
func GetListListingsPricingEnumValues() []ListListingsPricingEnum {
	values := make([]ListListingsPricingEnum, 0)
	for _, v := range mappingListListingsPricing {
		values = append(values, v)
	}
	return values
}
