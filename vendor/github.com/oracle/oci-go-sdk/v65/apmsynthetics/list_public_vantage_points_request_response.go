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

// ListPublicVantagePointsRequest wrapper for the ListPublicVantagePoints operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/apmsynthetics/ListPublicVantagePoints.go.html to see an example of how to use ListPublicVantagePointsRequest.
type ListPublicVantagePointsRequest struct {

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
	SortOrder ListPublicVantagePointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide one sort by (`sortBy`). Default order for displayName or name is ascending. The displayName or name
	// sort by is case insensitive.
	SortBy ListPublicVantagePointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

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

func (request ListPublicVantagePointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPublicVantagePointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPublicVantagePointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPublicVantagePointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPublicVantagePointsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPublicVantagePointsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPublicVantagePointsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPublicVantagePointsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPublicVantagePointsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPublicVantagePointsResponse wrapper for the ListPublicVantagePoints operation
type ListPublicVantagePointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PublicVantagePointCollection instances
	PublicVantagePointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPublicVantagePointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPublicVantagePointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPublicVantagePointsSortOrderEnum Enum with underlying type: string
type ListPublicVantagePointsSortOrderEnum string

// Set of constants representing the allowable values for ListPublicVantagePointsSortOrderEnum
const (
	ListPublicVantagePointsSortOrderAsc  ListPublicVantagePointsSortOrderEnum = "ASC"
	ListPublicVantagePointsSortOrderDesc ListPublicVantagePointsSortOrderEnum = "DESC"
)

var mappingListPublicVantagePointsSortOrderEnum = map[string]ListPublicVantagePointsSortOrderEnum{
	"ASC":  ListPublicVantagePointsSortOrderAsc,
	"DESC": ListPublicVantagePointsSortOrderDesc,
}

var mappingListPublicVantagePointsSortOrderEnumLowerCase = map[string]ListPublicVantagePointsSortOrderEnum{
	"asc":  ListPublicVantagePointsSortOrderAsc,
	"desc": ListPublicVantagePointsSortOrderDesc,
}

// GetListPublicVantagePointsSortOrderEnumValues Enumerates the set of values for ListPublicVantagePointsSortOrderEnum
func GetListPublicVantagePointsSortOrderEnumValues() []ListPublicVantagePointsSortOrderEnum {
	values := make([]ListPublicVantagePointsSortOrderEnum, 0)
	for _, v := range mappingListPublicVantagePointsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublicVantagePointsSortOrderEnumStringValues Enumerates the set of values in String for ListPublicVantagePointsSortOrderEnum
func GetListPublicVantagePointsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPublicVantagePointsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublicVantagePointsSortOrderEnum(val string) (ListPublicVantagePointsSortOrderEnum, bool) {
	enum, ok := mappingListPublicVantagePointsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPublicVantagePointsSortByEnum Enum with underlying type: string
type ListPublicVantagePointsSortByEnum string

// Set of constants representing the allowable values for ListPublicVantagePointsSortByEnum
const (
	ListPublicVantagePointsSortByName        ListPublicVantagePointsSortByEnum = "name"
	ListPublicVantagePointsSortByDisplayname ListPublicVantagePointsSortByEnum = "displayName"
)

var mappingListPublicVantagePointsSortByEnum = map[string]ListPublicVantagePointsSortByEnum{
	"name":        ListPublicVantagePointsSortByName,
	"displayName": ListPublicVantagePointsSortByDisplayname,
}

var mappingListPublicVantagePointsSortByEnumLowerCase = map[string]ListPublicVantagePointsSortByEnum{
	"name":        ListPublicVantagePointsSortByName,
	"displayname": ListPublicVantagePointsSortByDisplayname,
}

// GetListPublicVantagePointsSortByEnumValues Enumerates the set of values for ListPublicVantagePointsSortByEnum
func GetListPublicVantagePointsSortByEnumValues() []ListPublicVantagePointsSortByEnum {
	values := make([]ListPublicVantagePointsSortByEnum, 0)
	for _, v := range mappingListPublicVantagePointsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPublicVantagePointsSortByEnumStringValues Enumerates the set of values in String for ListPublicVantagePointsSortByEnum
func GetListPublicVantagePointsSortByEnumStringValues() []string {
	return []string{
		"name",
		"displayName",
	}
}

// GetMappingListPublicVantagePointsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPublicVantagePointsSortByEnum(val string) (ListPublicVantagePointsSortByEnum, bool) {
	enum, ok := mappingListPublicVantagePointsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
