// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOnPremiseVantagePointsRequest wrapper for the ListOnPremiseVantagePoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmsynthetics/ListOnPremiseVantagePoints.go.html to see an example of how to use ListOnPremiseVantagePointsRequest.
type ListOnPremiseVantagePointsRequest struct {

	// The APM domain ID the request is intended for.
	ApmDomainId *string `mandatory:"true" contributesTo:"query" name:"apmDomainId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The maximum number of results per page, or items to return in a paginated
	// "List" call. For information on how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). Default sort order is ascending.
	SortOrder ListOnPremiseVantagePointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order of displayName is ascending.
	// Default order of timeCreated and timeUpdated is descending.
	// The displayName sort by is case-sensitive.
	SortBy ListOnPremiseVantagePointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only the resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only the resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOnPremiseVantagePointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOnPremiseVantagePointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOnPremiseVantagePointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOnPremiseVantagePointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOnPremiseVantagePointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOnPremiseVantagePointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOnPremiseVantagePointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOnPremiseVantagePointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOnPremiseVantagePointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOnPremiseVantagePointsResponse wrapper for the ListOnPremiseVantagePoints operation
type ListOnPremiseVantagePointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OnPremiseVantagePointCollection instances
	OnPremiseVantagePointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOnPremiseVantagePointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOnPremiseVantagePointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOnPremiseVantagePointsSortOrderEnum Enum with underlying type: string
type ListOnPremiseVantagePointsSortOrderEnum string

// Set of constants representing the allowable values for ListOnPremiseVantagePointsSortOrderEnum
const (
	ListOnPremiseVantagePointsSortOrderAsc  ListOnPremiseVantagePointsSortOrderEnum = "ASC"
	ListOnPremiseVantagePointsSortOrderDesc ListOnPremiseVantagePointsSortOrderEnum = "DESC"
)

var mappingListOnPremiseVantagePointsSortOrderEnum = map[string]ListOnPremiseVantagePointsSortOrderEnum{
	"ASC":  ListOnPremiseVantagePointsSortOrderAsc,
	"DESC": ListOnPremiseVantagePointsSortOrderDesc,
}

var mappingListOnPremiseVantagePointsSortOrderEnumLowerCase = map[string]ListOnPremiseVantagePointsSortOrderEnum{
	"asc":  ListOnPremiseVantagePointsSortOrderAsc,
	"desc": ListOnPremiseVantagePointsSortOrderDesc,
}

// GetListOnPremiseVantagePointsSortOrderEnumValues Enumerates the set of values for ListOnPremiseVantagePointsSortOrderEnum
func GetListOnPremiseVantagePointsSortOrderEnumValues() []ListOnPremiseVantagePointsSortOrderEnum {
	values := make([]ListOnPremiseVantagePointsSortOrderEnum, 0)
	for _, v := range mappingListOnPremiseVantagePointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOnPremiseVantagePointsSortOrderEnumStringValues Enumerates the set of values in String for ListOnPremiseVantagePointsSortOrderEnum
func GetListOnPremiseVantagePointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOnPremiseVantagePointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOnPremiseVantagePointsSortOrderEnum(val string) (ListOnPremiseVantagePointsSortOrderEnum, bool) {
	enum, ok := mappingListOnPremiseVantagePointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOnPremiseVantagePointsSortByEnum Enum with underlying type: string
type ListOnPremiseVantagePointsSortByEnum string

// Set of constants representing the allowable values for ListOnPremiseVantagePointsSortByEnum
const (
	ListOnPremiseVantagePointsSortByDisplayname ListOnPremiseVantagePointsSortByEnum = "displayName"
	ListOnPremiseVantagePointsSortByName        ListOnPremiseVantagePointsSortByEnum = "name"
	ListOnPremiseVantagePointsSortByTimecreated ListOnPremiseVantagePointsSortByEnum = "timeCreated"
	ListOnPremiseVantagePointsSortByTimeupdated ListOnPremiseVantagePointsSortByEnum = "timeUpdated"
)

var mappingListOnPremiseVantagePointsSortByEnum = map[string]ListOnPremiseVantagePointsSortByEnum{
	"displayName": ListOnPremiseVantagePointsSortByDisplayname,
	"name":        ListOnPremiseVantagePointsSortByName,
	"timeCreated": ListOnPremiseVantagePointsSortByTimecreated,
	"timeUpdated": ListOnPremiseVantagePointsSortByTimeupdated,
}

var mappingListOnPremiseVantagePointsSortByEnumLowerCase = map[string]ListOnPremiseVantagePointsSortByEnum{
	"displayname": ListOnPremiseVantagePointsSortByDisplayname,
	"name":        ListOnPremiseVantagePointsSortByName,
	"timecreated": ListOnPremiseVantagePointsSortByTimecreated,
	"timeupdated": ListOnPremiseVantagePointsSortByTimeupdated,
}

// GetListOnPremiseVantagePointsSortByEnumValues Enumerates the set of values for ListOnPremiseVantagePointsSortByEnum
func GetListOnPremiseVantagePointsSortByEnumValues() []ListOnPremiseVantagePointsSortByEnum {
	values := make([]ListOnPremiseVantagePointsSortByEnum, 0)
	for _, v := range mappingListOnPremiseVantagePointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOnPremiseVantagePointsSortByEnumStringValues Enumerates the set of values in String for ListOnPremiseVantagePointsSortByEnum
func GetListOnPremiseVantagePointsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"name",
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListOnPremiseVantagePointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOnPremiseVantagePointsSortByEnum(val string) (ListOnPremiseVantagePointsSortByEnum, bool) {
	enum, ok := mappingListOnPremiseVantagePointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
