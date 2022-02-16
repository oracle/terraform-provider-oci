// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListListingsRequest wrapper for the ListListings operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListListings.go.html to see an example of how to use ListListingsRequest.
type ListListingsRequest struct {

	// The name of the listing.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// The unique identifier for the listing.
	ListingId *string `mandatory:"false" contributesTo:"query" name:"listingId"`

	// Image ID of the listing
	ImageId *string `mandatory:"false" contributesTo:"query" name:"imageId"`

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
	Pricing []PricingTypeEnumEnum `contributesTo:"query" name:"pricing" omitEmpty:"true" collectionFormat:"multi"`

	// Indicates whether to show only featured listings. If this is set to `false` or is omitted, then all listings will be returned.
	IsFeatured *bool `mandatory:"false" contributesTo:"query" name:"isFeatured"`

	// The type of the listing.
	ListingTypes []ListingTypeEnum `contributesTo:"query" name:"listingTypes" omitEmpty:"true" collectionFormat:"multi"`

	// The operating system of the listing.
	OperatingSystems []string `contributesTo:"query" name:"operatingSystems" collectionFormat:"multi"`

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
func (request ListListingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListListingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListListingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListListingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListListingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListListingsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListListingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListListingsSortOrderEnumStringValues(), ",")))
	}
	for _, val := range request.Pricing {
		if _, ok := GetMappingPricingTypeEnumEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Pricing: %s. Supported values are: %s.", val, strings.Join(GetPricingTypeEnumEnumStringValues(), ",")))
		}
	}

	for _, val := range request.ListingTypes {
		if _, ok := GetMappingListingTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingTypes: %s. Supported values are: %s.", val, strings.Join(GetListingTypeEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListListingsSortByEnum = map[string]ListListingsSortByEnum{
	"TIMERELEASED": ListListingsSortByTimereleased,
}

// GetListListingsSortByEnumValues Enumerates the set of values for ListListingsSortByEnum
func GetListListingsSortByEnumValues() []ListListingsSortByEnum {
	values := make([]ListListingsSortByEnum, 0)
	for _, v := range mappingListListingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingsSortByEnumStringValues Enumerates the set of values in String for ListListingsSortByEnum
func GetListListingsSortByEnumStringValues() []string {
	return []string{
		"TIMERELEASED",
	}
}

// GetMappingListListingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingsSortByEnum(val string) (ListListingsSortByEnum, bool) {
	mappingListListingsSortByEnumIgnoreCase := make(map[string]ListListingsSortByEnum)
	for k, v := range mappingListListingsSortByEnum {
		mappingListListingsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListListingsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListListingsSortOrderEnum Enum with underlying type: string
type ListListingsSortOrderEnum string

// Set of constants representing the allowable values for ListListingsSortOrderEnum
const (
	ListListingsSortOrderAsc  ListListingsSortOrderEnum = "ASC"
	ListListingsSortOrderDesc ListListingsSortOrderEnum = "DESC"
)

var mappingListListingsSortOrderEnum = map[string]ListListingsSortOrderEnum{
	"ASC":  ListListingsSortOrderAsc,
	"DESC": ListListingsSortOrderDesc,
}

// GetListListingsSortOrderEnumValues Enumerates the set of values for ListListingsSortOrderEnum
func GetListListingsSortOrderEnumValues() []ListListingsSortOrderEnum {
	values := make([]ListListingsSortOrderEnum, 0)
	for _, v := range mappingListListingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListListingsSortOrderEnumStringValues Enumerates the set of values in String for ListListingsSortOrderEnum
func GetListListingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListListingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListListingsSortOrderEnum(val string) (ListListingsSortOrderEnum, bool) {
	mappingListListingsSortOrderEnumIgnoreCase := make(map[string]ListListingsSortOrderEnum)
	for k, v := range mappingListListingsSortOrderEnum {
		mappingListListingsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListListingsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
