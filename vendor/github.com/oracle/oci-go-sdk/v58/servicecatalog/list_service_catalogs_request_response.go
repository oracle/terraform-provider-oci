// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicecatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListServiceCatalogsRequest wrapper for the ListServiceCatalogs operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/servicecatalog/ListServiceCatalogs.go.html to see an example of how to use ListServiceCatalogsRequest.
type ListServiceCatalogsRequest struct {

	// The unique identifier for the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The unique identifier for the service catalog.
	ServiceCatalogId *string `mandatory:"false" contributesTo:"query" name:"serviceCatalogId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// How many records to return. Specify a value greater than zero and less than or equal to 1000. The default is 30.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is `TIMECREATED`
	SortBy ListServiceCatalogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to apply, either `ASC` or `DESC`. Default is `ASC`.
	SortOrder ListServiceCatalogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Exact match name filter.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListServiceCatalogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListServiceCatalogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListServiceCatalogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListServiceCatalogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListServiceCatalogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListServiceCatalogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListServiceCatalogsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListServiceCatalogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListServiceCatalogsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListServiceCatalogsResponse wrapper for the ListServiceCatalogs operation
type ListServiceCatalogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ServiceCatalogCollection instances
	ServiceCatalogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListServiceCatalogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListServiceCatalogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListServiceCatalogsSortByEnum Enum with underlying type: string
type ListServiceCatalogsSortByEnum string

// Set of constants representing the allowable values for ListServiceCatalogsSortByEnum
const (
	ListServiceCatalogsSortByTimecreated ListServiceCatalogsSortByEnum = "TIMECREATED"
)

var mappingListServiceCatalogsSortByEnum = map[string]ListServiceCatalogsSortByEnum{
	"TIMECREATED": ListServiceCatalogsSortByTimecreated,
}

// GetListServiceCatalogsSortByEnumValues Enumerates the set of values for ListServiceCatalogsSortByEnum
func GetListServiceCatalogsSortByEnumValues() []ListServiceCatalogsSortByEnum {
	values := make([]ListServiceCatalogsSortByEnum, 0)
	for _, v := range mappingListServiceCatalogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceCatalogsSortByEnumStringValues Enumerates the set of values in String for ListServiceCatalogsSortByEnum
func GetListServiceCatalogsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
	}
}

// GetMappingListServiceCatalogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceCatalogsSortByEnum(val string) (ListServiceCatalogsSortByEnum, bool) {
	mappingListServiceCatalogsSortByEnumIgnoreCase := make(map[string]ListServiceCatalogsSortByEnum)
	for k, v := range mappingListServiceCatalogsSortByEnum {
		mappingListServiceCatalogsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListServiceCatalogsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListServiceCatalogsSortOrderEnum Enum with underlying type: string
type ListServiceCatalogsSortOrderEnum string

// Set of constants representing the allowable values for ListServiceCatalogsSortOrderEnum
const (
	ListServiceCatalogsSortOrderAsc  ListServiceCatalogsSortOrderEnum = "ASC"
	ListServiceCatalogsSortOrderDesc ListServiceCatalogsSortOrderEnum = "DESC"
)

var mappingListServiceCatalogsSortOrderEnum = map[string]ListServiceCatalogsSortOrderEnum{
	"ASC":  ListServiceCatalogsSortOrderAsc,
	"DESC": ListServiceCatalogsSortOrderDesc,
}

// GetListServiceCatalogsSortOrderEnumValues Enumerates the set of values for ListServiceCatalogsSortOrderEnum
func GetListServiceCatalogsSortOrderEnumValues() []ListServiceCatalogsSortOrderEnum {
	values := make([]ListServiceCatalogsSortOrderEnum, 0)
	for _, v := range mappingListServiceCatalogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListServiceCatalogsSortOrderEnumStringValues Enumerates the set of values in String for ListServiceCatalogsSortOrderEnum
func GetListServiceCatalogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListServiceCatalogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListServiceCatalogsSortOrderEnum(val string) (ListServiceCatalogsSortOrderEnum, bool) {
	mappingListServiceCatalogsSortOrderEnumIgnoreCase := make(map[string]ListServiceCatalogsSortOrderEnum)
	for k, v := range mappingListServiceCatalogsSortOrderEnum {
		mappingListServiceCatalogsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListServiceCatalogsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
