// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCccListingsRequest wrapper for the ListCccListings operation
type ListCccListingsRequest struct {

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose display name contains the substring.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The unique identifier for the listing.
	CccListingId *string `mandatory:"false" contributesTo:"query" name:"cccListingId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to
	// list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The client request OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListCccListingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCccListingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCccListingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCccListingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCccListingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCccListingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCccListingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCccListingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCccListingsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCccListingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCccListingsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCccListingsResponse wrapper for the ListCccListings operation
type ListCccListingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CccListingCollection instances
	CccListingCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListCccListingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCccListingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCccListingsSortByEnum Enum with underlying type: string
type ListCccListingsSortByEnum string

// Set of constants representing the allowable values for ListCccListingsSortByEnum
const (
	ListCccListingsSortByTimecreated ListCccListingsSortByEnum = "timeCreated"
	ListCccListingsSortByDisplayname ListCccListingsSortByEnum = "displayName"
)

var mappingListCccListingsSortByEnum = map[string]ListCccListingsSortByEnum{
	"timeCreated": ListCccListingsSortByTimecreated,
	"displayName": ListCccListingsSortByDisplayname,
}

var mappingListCccListingsSortByEnumLowerCase = map[string]ListCccListingsSortByEnum{
	"timecreated": ListCccListingsSortByTimecreated,
	"displayname": ListCccListingsSortByDisplayname,
}

// GetListCccListingsSortByEnumValues Enumerates the set of values for ListCccListingsSortByEnum
func GetListCccListingsSortByEnumValues() []ListCccListingsSortByEnum {
	values := make([]ListCccListingsSortByEnum, 0)
	for _, v := range mappingListCccListingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccListingsSortByEnumStringValues Enumerates the set of values in String for ListCccListingsSortByEnum
func GetListCccListingsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListCccListingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccListingsSortByEnum(val string) (ListCccListingsSortByEnum, bool) {
	enum, ok := mappingListCccListingsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCccListingsSortOrderEnum Enum with underlying type: string
type ListCccListingsSortOrderEnum string

// Set of constants representing the allowable values for ListCccListingsSortOrderEnum
const (
	ListCccListingsSortOrderAsc  ListCccListingsSortOrderEnum = "ASC"
	ListCccListingsSortOrderDesc ListCccListingsSortOrderEnum = "DESC"
)

var mappingListCccListingsSortOrderEnum = map[string]ListCccListingsSortOrderEnum{
	"ASC":  ListCccListingsSortOrderAsc,
	"DESC": ListCccListingsSortOrderDesc,
}

var mappingListCccListingsSortOrderEnumLowerCase = map[string]ListCccListingsSortOrderEnum{
	"asc":  ListCccListingsSortOrderAsc,
	"desc": ListCccListingsSortOrderDesc,
}

// GetListCccListingsSortOrderEnumValues Enumerates the set of values for ListCccListingsSortOrderEnum
func GetListCccListingsSortOrderEnumValues() []ListCccListingsSortOrderEnum {
	values := make([]ListCccListingsSortOrderEnum, 0)
	for _, v := range mappingListCccListingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCccListingsSortOrderEnumStringValues Enumerates the set of values in String for ListCccListingsSortOrderEnum
func GetListCccListingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCccListingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCccListingsSortOrderEnum(val string) (ListCccListingsSortOrderEnum, bool) {
	enum, ok := mappingListCccListingsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
