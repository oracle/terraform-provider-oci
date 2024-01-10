// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package functions

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPbfListingsRequest wrapper for the ListPbfListings operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/functions/ListPbfListings.go.html to see an example of how to use ListPbfListingsRequest.
type ListPbfListingsRequest struct {

	// unique PbfListing identifier
	PbfListingId *string `mandatory:"false" contributesTo:"query" name:"pbfListingId"`

	// A filter to return only resources that match the entire PBF name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A filter to return only resources that contain the supplied filter text in the PBF name given.
	NameContains *string `mandatory:"false" contributesTo:"query" name:"nameContains"`

	// A filter to return only resources that start with the supplied filter text in the PBF name given.
	NameStartsWith *string `mandatory:"false" contributesTo:"query" name:"nameStartsWith"`

	// A filter to return only resources that match the service trigger sources of a PBF.
	Trigger []string `contributesTo:"query" name:"trigger" collectionFormat:"multi"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState PbfListingLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return. 1 is the minimum, 50 is the maximum.
	// Default: 10
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token for a list query returned by a previous operation
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order.
	// * **ASC:** Ascending sort order.
	// * **DESC:** Descending sort order.
	SortOrder ListPbfListingsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for name is ascending.
	SortBy ListPbfListingsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPbfListingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPbfListingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPbfListingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPbfListingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPbfListingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPbfListingLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetPbfListingLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPbfListingsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPbfListingsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPbfListingsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPbfListingsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPbfListingsResponse wrapper for the ListPbfListings operation
type ListPbfListingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PbfListingsCollection instances
	PbfListingsCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPbfListingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPbfListingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPbfListingsSortOrderEnum Enum with underlying type: string
type ListPbfListingsSortOrderEnum string

// Set of constants representing the allowable values for ListPbfListingsSortOrderEnum
const (
	ListPbfListingsSortOrderAsc  ListPbfListingsSortOrderEnum = "ASC"
	ListPbfListingsSortOrderDesc ListPbfListingsSortOrderEnum = "DESC"
)

var mappingListPbfListingsSortOrderEnum = map[string]ListPbfListingsSortOrderEnum{
	"ASC":  ListPbfListingsSortOrderAsc,
	"DESC": ListPbfListingsSortOrderDesc,
}

var mappingListPbfListingsSortOrderEnumLowerCase = map[string]ListPbfListingsSortOrderEnum{
	"asc":  ListPbfListingsSortOrderAsc,
	"desc": ListPbfListingsSortOrderDesc,
}

// GetListPbfListingsSortOrderEnumValues Enumerates the set of values for ListPbfListingsSortOrderEnum
func GetListPbfListingsSortOrderEnumValues() []ListPbfListingsSortOrderEnum {
	values := make([]ListPbfListingsSortOrderEnum, 0)
	for _, v := range mappingListPbfListingsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPbfListingsSortOrderEnumStringValues Enumerates the set of values in String for ListPbfListingsSortOrderEnum
func GetListPbfListingsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPbfListingsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPbfListingsSortOrderEnum(val string) (ListPbfListingsSortOrderEnum, bool) {
	enum, ok := mappingListPbfListingsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPbfListingsSortByEnum Enum with underlying type: string
type ListPbfListingsSortByEnum string

// Set of constants representing the allowable values for ListPbfListingsSortByEnum
const (
	ListPbfListingsSortByTimecreated ListPbfListingsSortByEnum = "timeCreated"
	ListPbfListingsSortByName        ListPbfListingsSortByEnum = "name"
)

var mappingListPbfListingsSortByEnum = map[string]ListPbfListingsSortByEnum{
	"timeCreated": ListPbfListingsSortByTimecreated,
	"name":        ListPbfListingsSortByName,
}

var mappingListPbfListingsSortByEnumLowerCase = map[string]ListPbfListingsSortByEnum{
	"timecreated": ListPbfListingsSortByTimecreated,
	"name":        ListPbfListingsSortByName,
}

// GetListPbfListingsSortByEnumValues Enumerates the set of values for ListPbfListingsSortByEnum
func GetListPbfListingsSortByEnumValues() []ListPbfListingsSortByEnum {
	values := make([]ListPbfListingsSortByEnum, 0)
	for _, v := range mappingListPbfListingsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPbfListingsSortByEnumStringValues Enumerates the set of values in String for ListPbfListingsSortByEnum
func GetListPbfListingsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"name",
	}
}

// GetMappingListPbfListingsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPbfListingsSortByEnum(val string) (ListPbfListingsSortByEnum, bool) {
	enum, ok := mappingListPbfListingsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
