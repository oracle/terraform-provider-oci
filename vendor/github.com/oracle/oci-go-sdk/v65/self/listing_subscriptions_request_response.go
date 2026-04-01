// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package self

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListingSubscriptionsRequest wrapper for the ListingSubscriptions operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/self/ListingSubscriptions.go.html to see an example of how to use ListingSubscriptionsRequest.
type ListingSubscriptionsRequest struct {

	// The unique identifier for the listing.
	ListingId *string `mandatory:"true" contributesTo:"query" name:"listingId"`

	// A filter to return only resources that match the given name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListingSubscriptionsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListingSubscriptionsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListingSubscriptionsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListingSubscriptionsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListingSubscriptionsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListingSubscriptionsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListingSubscriptionsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingSubscriptionsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListingSubscriptionsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingSubscriptionsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListingSubscriptionsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingSubscriptionsResponse wrapper for the ListingSubscriptions operation
type ListingSubscriptionsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ListingSubscriptionsCollection instances
	ListingSubscriptionsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListingSubscriptionsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListingSubscriptionsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListingSubscriptionsSortByEnum Enum with underlying type: string
type ListingSubscriptionsSortByEnum string

// Set of constants representing the allowable values for ListingSubscriptionsSortByEnum
const (
	ListingSubscriptionsSortByTimecreated ListingSubscriptionsSortByEnum = "timeCreated"
	ListingSubscriptionsSortByDisplayname ListingSubscriptionsSortByEnum = "displayName"
	ListingSubscriptionsSortBySelftokenid ListingSubscriptionsSortByEnum = "selfTokenId"
	ListingSubscriptionsSortByProductid   ListingSubscriptionsSortByEnum = "productId"
)

var mappingListingSubscriptionsSortByEnum = map[string]ListingSubscriptionsSortByEnum{
	"timeCreated": ListingSubscriptionsSortByTimecreated,
	"displayName": ListingSubscriptionsSortByDisplayname,
	"selfTokenId": ListingSubscriptionsSortBySelftokenid,
	"productId":   ListingSubscriptionsSortByProductid,
}

var mappingListingSubscriptionsSortByEnumLowerCase = map[string]ListingSubscriptionsSortByEnum{
	"timecreated": ListingSubscriptionsSortByTimecreated,
	"displayname": ListingSubscriptionsSortByDisplayname,
	"selftokenid": ListingSubscriptionsSortBySelftokenid,
	"productid":   ListingSubscriptionsSortByProductid,
}

// GetListingSubscriptionsSortByEnumValues Enumerates the set of values for ListingSubscriptionsSortByEnum
func GetListingSubscriptionsSortByEnumValues() []ListingSubscriptionsSortByEnum {
	values := make([]ListingSubscriptionsSortByEnum, 0)
	for _, v := range mappingListingSubscriptionsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListingSubscriptionsSortByEnumStringValues Enumerates the set of values in String for ListingSubscriptionsSortByEnum
func GetListingSubscriptionsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"selfTokenId",
		"productId",
	}
}

// GetMappingListingSubscriptionsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingSubscriptionsSortByEnum(val string) (ListingSubscriptionsSortByEnum, bool) {
	enum, ok := mappingListingSubscriptionsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListingSubscriptionsSortOrderEnum Enum with underlying type: string
type ListingSubscriptionsSortOrderEnum string

// Set of constants representing the allowable values for ListingSubscriptionsSortOrderEnum
const (
	ListingSubscriptionsSortOrderAsc  ListingSubscriptionsSortOrderEnum = "ASC"
	ListingSubscriptionsSortOrderDesc ListingSubscriptionsSortOrderEnum = "DESC"
)

var mappingListingSubscriptionsSortOrderEnum = map[string]ListingSubscriptionsSortOrderEnum{
	"ASC":  ListingSubscriptionsSortOrderAsc,
	"DESC": ListingSubscriptionsSortOrderDesc,
}

var mappingListingSubscriptionsSortOrderEnumLowerCase = map[string]ListingSubscriptionsSortOrderEnum{
	"asc":  ListingSubscriptionsSortOrderAsc,
	"desc": ListingSubscriptionsSortOrderDesc,
}

// GetListingSubscriptionsSortOrderEnumValues Enumerates the set of values for ListingSubscriptionsSortOrderEnum
func GetListingSubscriptionsSortOrderEnumValues() []ListingSubscriptionsSortOrderEnum {
	values := make([]ListingSubscriptionsSortOrderEnum, 0)
	for _, v := range mappingListingSubscriptionsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListingSubscriptionsSortOrderEnumStringValues Enumerates the set of values in String for ListingSubscriptionsSortOrderEnum
func GetListingSubscriptionsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListingSubscriptionsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingSubscriptionsSortOrderEnum(val string) (ListingSubscriptionsSortOrderEnum, bool) {
	enum, ok := mappingListingSubscriptionsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
