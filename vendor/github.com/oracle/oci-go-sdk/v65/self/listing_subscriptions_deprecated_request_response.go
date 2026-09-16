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

// ListingSubscriptionsDeprecatedRequest wrapper for the ListingSubscriptionsDeprecated operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/self/ListingSubscriptionsDeprecated.go.html to see an example of how to use ListingSubscriptionsDeprecatedRequest.
type ListingSubscriptionsDeprecatedRequest struct {

	// The unique identifier for the listing.
	ListingId *string `mandatory:"true" contributesTo:"query" name:"listingId"`

	// A filter to return only resources that match the given name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListingSubscriptionsDeprecatedSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListingSubscriptionsDeprecatedSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

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

func (request ListingSubscriptionsDeprecatedRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListingSubscriptionsDeprecatedRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListingSubscriptionsDeprecatedRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListingSubscriptionsDeprecatedRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListingSubscriptionsDeprecatedRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListingSubscriptionsDeprecatedSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListingSubscriptionsDeprecatedSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListingSubscriptionsDeprecatedSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListingSubscriptionsDeprecatedSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListingSubscriptionsDeprecatedResponse wrapper for the ListingSubscriptionsDeprecated operation
type ListingSubscriptionsDeprecatedResponse struct {

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

func (response ListingSubscriptionsDeprecatedResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListingSubscriptionsDeprecatedResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListingSubscriptionsDeprecatedSortByEnum Enum with underlying type: string
type ListingSubscriptionsDeprecatedSortByEnum string

// Set of constants representing the allowable values for ListingSubscriptionsDeprecatedSortByEnum
const (
	ListingSubscriptionsDeprecatedSortByTimecreated ListingSubscriptionsDeprecatedSortByEnum = "timeCreated"
	ListingSubscriptionsDeprecatedSortByDisplayname ListingSubscriptionsDeprecatedSortByEnum = "displayName"
	ListingSubscriptionsDeprecatedSortByProductid   ListingSubscriptionsDeprecatedSortByEnum = "productId"
)

var mappingListingSubscriptionsDeprecatedSortByEnum = map[string]ListingSubscriptionsDeprecatedSortByEnum{
	"timeCreated": ListingSubscriptionsDeprecatedSortByTimecreated,
	"displayName": ListingSubscriptionsDeprecatedSortByDisplayname,
	"productId":   ListingSubscriptionsDeprecatedSortByProductid,
}

var mappingListingSubscriptionsDeprecatedSortByEnumLowerCase = map[string]ListingSubscriptionsDeprecatedSortByEnum{
	"timecreated": ListingSubscriptionsDeprecatedSortByTimecreated,
	"displayname": ListingSubscriptionsDeprecatedSortByDisplayname,
	"productid":   ListingSubscriptionsDeprecatedSortByProductid,
}

// GetListingSubscriptionsDeprecatedSortByEnumValues Enumerates the set of values for ListingSubscriptionsDeprecatedSortByEnum
func GetListingSubscriptionsDeprecatedSortByEnumValues() []ListingSubscriptionsDeprecatedSortByEnum {
	values := make([]ListingSubscriptionsDeprecatedSortByEnum, 0)
	for _, v := range mappingListingSubscriptionsDeprecatedSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListingSubscriptionsDeprecatedSortByEnumStringValues Enumerates the set of values in String for ListingSubscriptionsDeprecatedSortByEnum
func GetListingSubscriptionsDeprecatedSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"productId",
	}
}

// GetMappingListingSubscriptionsDeprecatedSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingSubscriptionsDeprecatedSortByEnum(val string) (ListingSubscriptionsDeprecatedSortByEnum, bool) {
	enum, ok := mappingListingSubscriptionsDeprecatedSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListingSubscriptionsDeprecatedSortOrderEnum Enum with underlying type: string
type ListingSubscriptionsDeprecatedSortOrderEnum string

// Set of constants representing the allowable values for ListingSubscriptionsDeprecatedSortOrderEnum
const (
	ListingSubscriptionsDeprecatedSortOrderAsc  ListingSubscriptionsDeprecatedSortOrderEnum = "ASC"
	ListingSubscriptionsDeprecatedSortOrderDesc ListingSubscriptionsDeprecatedSortOrderEnum = "DESC"
)

var mappingListingSubscriptionsDeprecatedSortOrderEnum = map[string]ListingSubscriptionsDeprecatedSortOrderEnum{
	"ASC":  ListingSubscriptionsDeprecatedSortOrderAsc,
	"DESC": ListingSubscriptionsDeprecatedSortOrderDesc,
}

var mappingListingSubscriptionsDeprecatedSortOrderEnumLowerCase = map[string]ListingSubscriptionsDeprecatedSortOrderEnum{
	"asc":  ListingSubscriptionsDeprecatedSortOrderAsc,
	"desc": ListingSubscriptionsDeprecatedSortOrderDesc,
}

// GetListingSubscriptionsDeprecatedSortOrderEnumValues Enumerates the set of values for ListingSubscriptionsDeprecatedSortOrderEnum
func GetListingSubscriptionsDeprecatedSortOrderEnumValues() []ListingSubscriptionsDeprecatedSortOrderEnum {
	values := make([]ListingSubscriptionsDeprecatedSortOrderEnum, 0)
	for _, v := range mappingListingSubscriptionsDeprecatedSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListingSubscriptionsDeprecatedSortOrderEnumStringValues Enumerates the set of values in String for ListingSubscriptionsDeprecatedSortOrderEnum
func GetListingSubscriptionsDeprecatedSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListingSubscriptionsDeprecatedSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListingSubscriptionsDeprecatedSortOrderEnum(val string) (ListingSubscriptionsDeprecatedSortOrderEnum, bool) {
	enum, ok := mappingListingSubscriptionsDeprecatedSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
