// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListAcceptedAgreementsRequest wrapper for the ListAcceptedAgreements operation
type ListAcceptedAgreementsRequest struct {

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The display name of the resource.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The unique identifier for the listing.
	ListingId *string `mandatory:"false" contributesTo:"query" name:"listingId"`

	// The version of the package. Package versions are unique within a listing.
	PackageVersion *string `mandatory:"false" contributesTo:"query" name:"packageVersion"`

	// The unique identifier for the accepted terms of use agreement.
	AcceptedAgreementId *string `mandatory:"false" contributesTo:"query" name:"acceptedAgreementId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to use to sort listed results. You can only specify one field to sort by.
	// `TIMEACCEPTED` displays results in descending order by default. You can change your
	// preference by specifying a different sort order.
	SortBy ListAcceptedAgreementsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either `ASC` or `DESC`.
	SortOrder ListAcceptedAgreementsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAcceptedAgreementsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAcceptedAgreementsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAcceptedAgreementsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAcceptedAgreementsResponse wrapper for the ListAcceptedAgreements operation
type ListAcceptedAgreementsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []AcceptedAgreementSummary instances
	Items []AcceptedAgreementSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAcceptedAgreementsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAcceptedAgreementsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAcceptedAgreementsSortByEnum Enum with underlying type: string
type ListAcceptedAgreementsSortByEnum string

// Set of constants representing the allowable values for ListAcceptedAgreementsSortByEnum
const (
	ListAcceptedAgreementsSortByTimeaccepted ListAcceptedAgreementsSortByEnum = "TIMEACCEPTED"
)

var mappingListAcceptedAgreementsSortBy = map[string]ListAcceptedAgreementsSortByEnum{
	"TIMEACCEPTED": ListAcceptedAgreementsSortByTimeaccepted,
}

// GetListAcceptedAgreementsSortByEnumValues Enumerates the set of values for ListAcceptedAgreementsSortByEnum
func GetListAcceptedAgreementsSortByEnumValues() []ListAcceptedAgreementsSortByEnum {
	values := make([]ListAcceptedAgreementsSortByEnum, 0)
	for _, v := range mappingListAcceptedAgreementsSortBy {
		values = append(values, v)
	}
	return values
}

// ListAcceptedAgreementsSortOrderEnum Enum with underlying type: string
type ListAcceptedAgreementsSortOrderEnum string

// Set of constants representing the allowable values for ListAcceptedAgreementsSortOrderEnum
const (
	ListAcceptedAgreementsSortOrderAsc  ListAcceptedAgreementsSortOrderEnum = "ASC"
	ListAcceptedAgreementsSortOrderDesc ListAcceptedAgreementsSortOrderEnum = "DESC"
)

var mappingListAcceptedAgreementsSortOrder = map[string]ListAcceptedAgreementsSortOrderEnum{
	"ASC":  ListAcceptedAgreementsSortOrderAsc,
	"DESC": ListAcceptedAgreementsSortOrderDesc,
}

// GetListAcceptedAgreementsSortOrderEnumValues Enumerates the set of values for ListAcceptedAgreementsSortOrderEnum
func GetListAcceptedAgreementsSortOrderEnumValues() []ListAcceptedAgreementsSortOrderEnum {
	values := make([]ListAcceptedAgreementsSortOrderEnum, 0)
	for _, v := range mappingListAcceptedAgreementsSortOrder {
		values = append(values, v)
	}
	return values
}
