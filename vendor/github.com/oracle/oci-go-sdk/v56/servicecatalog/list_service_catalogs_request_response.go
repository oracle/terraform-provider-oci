// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package servicecatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
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

var mappingListServiceCatalogsSortBy = map[string]ListServiceCatalogsSortByEnum{
	"TIMECREATED": ListServiceCatalogsSortByTimecreated,
}

// GetListServiceCatalogsSortByEnumValues Enumerates the set of values for ListServiceCatalogsSortByEnum
func GetListServiceCatalogsSortByEnumValues() []ListServiceCatalogsSortByEnum {
	values := make([]ListServiceCatalogsSortByEnum, 0)
	for _, v := range mappingListServiceCatalogsSortBy {
		values = append(values, v)
	}
	return values
}

// ListServiceCatalogsSortOrderEnum Enum with underlying type: string
type ListServiceCatalogsSortOrderEnum string

// Set of constants representing the allowable values for ListServiceCatalogsSortOrderEnum
const (
	ListServiceCatalogsSortOrderAsc  ListServiceCatalogsSortOrderEnum = "ASC"
	ListServiceCatalogsSortOrderDesc ListServiceCatalogsSortOrderEnum = "DESC"
)

var mappingListServiceCatalogsSortOrder = map[string]ListServiceCatalogsSortOrderEnum{
	"ASC":  ListServiceCatalogsSortOrderAsc,
	"DESC": ListServiceCatalogsSortOrderDesc,
}

// GetListServiceCatalogsSortOrderEnumValues Enumerates the set of values for ListServiceCatalogsSortOrderEnum
func GetListServiceCatalogsSortOrderEnumValues() []ListServiceCatalogsSortOrderEnum {
	values := make([]ListServiceCatalogsSortOrderEnum, 0)
	for _, v := range mappingListServiceCatalogsSortOrder {
		values = append(values, v)
	}
	return values
}
