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

// ListPublicationsRequest wrapper for the ListPublications operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/marketplace/ListPublications.go.html to see an example of how to use ListPublicationsRequest.
type ListPublicationsRequest struct {

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The type of the listing.
	ListingType ListPublicationsListingTypeEnum `mandatory:"true" contributesTo:"query" name:"listingType" omitEmpty:"true"`

	// The name of the publication.
	Name []string `contributesTo:"query" name:"name" collectionFormat:"multi"`

	// The unique identifier for the publication.
	PublicationId *string `mandatory:"false" contributesTo:"query" name:"publicationId"`

	// The operating system of the listing.
	OperatingSystems []string `contributesTo:"query" name:"operatingSystems" collectionFormat:"multi"`

	// The field to use to sort listed results. You can only specify one field to sort by.
	// `TIMERELEASED` displays results in descending order by default.
	// You can change your preference by specifying a different sort order.
	SortBy ListPublicationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either `ASC` or `DESC`.
	SortOrder ListPublicationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPublicationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPublicationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPublicationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPublicationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPublicationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPublicationsListingTypeEnum(string(request.ListingType)); !ok && request.ListingType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ListingType: %s. Supported values are: %s.", request.ListingType, strings.Join(GetListPublicationsListingTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPublicationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPublicationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPublicationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPublicationsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPublicationsResponse wrapper for the ListPublications operation
type ListPublicationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []PublicationSummary instances
	Items []PublicationSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPublicationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPublicationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPublicationsListingTypeEnum Enum with underlying type: string
type ListPublicationsListingTypeEnum string

// Set of constants representing the allowable values for ListPublicationsListingTypeEnum
const (
	ListPublicationsListingTypeCommunity ListPublicationsListingTypeEnum = "COMMUNITY"
	ListPublicationsListingTypePartner   ListPublicationsListingTypeEnum = "PARTNER"
	ListPublicationsListingTypePrivate   ListPublicationsListingTypeEnum = "PRIVATE"
)

var mappingListPublicationsListingTypeEnum = map[string]ListPublicationsListingTypeEnum{
	"COMMUNITY": ListPublicationsListingTypeCommunity,
	"PARTNER":   ListPublicationsListingTypePartner,
	"PRIVATE":   ListPublicationsListingTypePrivate,
}

// GetListPublicationsListingTypeEnumValues Enumerates the set of values for ListPublicationsListingTypeEnum
func GetListPublicationsListingTypeEnumValues() []ListPublicationsListingTypeEnum {
	values := make([]ListPublicationsListingTypeEnum, 0)
	for _, v := range mappingListPublicationsListingTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublicationsListingTypeEnumStringValues Enumerates the set of values in String for ListPublicationsListingTypeEnum
func GetListPublicationsListingTypeEnumStringValues() []string {
	return []string{
		"COMMUNITY",
		"PARTNER",
		"PRIVATE",
	}
}

// GetMappingListPublicationsListingTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublicationsListingTypeEnum(val string) (ListPublicationsListingTypeEnum, bool) {
	mappingListPublicationsListingTypeEnumIgnoreCase := make(map[string]ListPublicationsListingTypeEnum)
	for k, v := range mappingListPublicationsListingTypeEnum {
		mappingListPublicationsListingTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPublicationsListingTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListPublicationsSortByEnum Enum with underlying type: string
type ListPublicationsSortByEnum string

// Set of constants representing the allowable values for ListPublicationsSortByEnum
const (
	ListPublicationsSortByTimereleased ListPublicationsSortByEnum = "TIMERELEASED"
)

var mappingListPublicationsSortByEnum = map[string]ListPublicationsSortByEnum{
	"TIMERELEASED": ListPublicationsSortByTimereleased,
}

// GetListPublicationsSortByEnumValues Enumerates the set of values for ListPublicationsSortByEnum
func GetListPublicationsSortByEnumValues() []ListPublicationsSortByEnum {
	values := make([]ListPublicationsSortByEnum, 0)
	for _, v := range mappingListPublicationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublicationsSortByEnumStringValues Enumerates the set of values in String for ListPublicationsSortByEnum
func GetListPublicationsSortByEnumStringValues() []string {
	return []string{
		"TIMERELEASED",
	}
}

// GetMappingListPublicationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublicationsSortByEnum(val string) (ListPublicationsSortByEnum, bool) {
	mappingListPublicationsSortByEnumIgnoreCase := make(map[string]ListPublicationsSortByEnum)
	for k, v := range mappingListPublicationsSortByEnum {
		mappingListPublicationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPublicationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListPublicationsSortOrderEnum Enum with underlying type: string
type ListPublicationsSortOrderEnum string

// Set of constants representing the allowable values for ListPublicationsSortOrderEnum
const (
	ListPublicationsSortOrderAsc  ListPublicationsSortOrderEnum = "ASC"
	ListPublicationsSortOrderDesc ListPublicationsSortOrderEnum = "DESC"
)

var mappingListPublicationsSortOrderEnum = map[string]ListPublicationsSortOrderEnum{
	"ASC":  ListPublicationsSortOrderAsc,
	"DESC": ListPublicationsSortOrderDesc,
}

// GetListPublicationsSortOrderEnumValues Enumerates the set of values for ListPublicationsSortOrderEnum
func GetListPublicationsSortOrderEnumValues() []ListPublicationsSortOrderEnum {
	values := make([]ListPublicationsSortOrderEnum, 0)
	for _, v := range mappingListPublicationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublicationsSortOrderEnumStringValues Enumerates the set of values in String for ListPublicationsSortOrderEnum
func GetListPublicationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPublicationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublicationsSortOrderEnum(val string) (ListPublicationsSortOrderEnum, bool) {
	mappingListPublicationsSortOrderEnumIgnoreCase := make(map[string]ListPublicationsSortOrderEnum)
	for k, v := range mappingListPublicationsSortOrderEnum {
		mappingListPublicationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListPublicationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
