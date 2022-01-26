// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
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

var mappingListPublicationsListingType = map[string]ListPublicationsListingTypeEnum{
	"COMMUNITY": ListPublicationsListingTypeCommunity,
	"PARTNER":   ListPublicationsListingTypePartner,
	"PRIVATE":   ListPublicationsListingTypePrivate,
}

// GetListPublicationsListingTypeEnumValues Enumerates the set of values for ListPublicationsListingTypeEnum
func GetListPublicationsListingTypeEnumValues() []ListPublicationsListingTypeEnum {
	values := make([]ListPublicationsListingTypeEnum, 0)
	for _, v := range mappingListPublicationsListingType {
		values = append(values, v)
	}
	return values
}

// ListPublicationsSortByEnum Enum with underlying type: string
type ListPublicationsSortByEnum string

// Set of constants representing the allowable values for ListPublicationsSortByEnum
const (
	ListPublicationsSortByTimereleased ListPublicationsSortByEnum = "TIMERELEASED"
)

var mappingListPublicationsSortBy = map[string]ListPublicationsSortByEnum{
	"TIMERELEASED": ListPublicationsSortByTimereleased,
}

// GetListPublicationsSortByEnumValues Enumerates the set of values for ListPublicationsSortByEnum
func GetListPublicationsSortByEnumValues() []ListPublicationsSortByEnum {
	values := make([]ListPublicationsSortByEnum, 0)
	for _, v := range mappingListPublicationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListPublicationsSortOrderEnum Enum with underlying type: string
type ListPublicationsSortOrderEnum string

// Set of constants representing the allowable values for ListPublicationsSortOrderEnum
const (
	ListPublicationsSortOrderAsc  ListPublicationsSortOrderEnum = "ASC"
	ListPublicationsSortOrderDesc ListPublicationsSortOrderEnum = "DESC"
)

var mappingListPublicationsSortOrder = map[string]ListPublicationsSortOrderEnum{
	"ASC":  ListPublicationsSortOrderAsc,
	"DESC": ListPublicationsSortOrderDesc,
}

// GetListPublicationsSortOrderEnumValues Enumerates the set of values for ListPublicationsSortOrderEnum
func GetListPublicationsSortOrderEnumValues() []ListPublicationsSortOrderEnum {
	values := make([]ListPublicationsSortOrderEnum, 0)
	for _, v := range mappingListPublicationsSortOrder {
		values = append(values, v)
	}
	return values
}
