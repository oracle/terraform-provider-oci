// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListInternalOccmDemandSignalCatalogsRequest wrapper for the ListInternalOccmDemandSignalCatalogs operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignalCatalogs.go.html to see an example of how to use ListInternalOccmDemandSignalCatalogsRequest.
type ListInternalOccmDemandSignalCatalogsRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The customer group ocid by which we would filter the list.
	OccCustomerGroupId *string `mandatory:"true" contributesTo:"query" name:"occCustomerGroupId"`

	// A filter to return only the resources that match the entire display name. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInternalOccmDemandSignalCatalogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the response of List Demand catalog  API. Only one sort order may be provided. The default order for name is case sensitive alphabetical order.
	SortBy ListInternalOccmDemandSignalCatalogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalOccmDemandSignalCatalogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalOccmDemandSignalCatalogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalOccmDemandSignalCatalogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalOccmDemandSignalCatalogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalOccmDemandSignalCatalogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInternalOccmDemandSignalCatalogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalOccmDemandSignalCatalogsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccmDemandSignalCatalogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalOccmDemandSignalCatalogsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalOccmDemandSignalCatalogsResponse wrapper for the ListInternalOccmDemandSignalCatalogs operation
type ListInternalOccmDemandSignalCatalogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccmDemandSignalCatalogCollection instances
	OccmDemandSignalCatalogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInternalOccmDemandSignalCatalogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalOccmDemandSignalCatalogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalOccmDemandSignalCatalogsSortOrderEnum Enum with underlying type: string
type ListInternalOccmDemandSignalCatalogsSortOrderEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalCatalogsSortOrderEnum
const (
	ListInternalOccmDemandSignalCatalogsSortOrderAsc  ListInternalOccmDemandSignalCatalogsSortOrderEnum = "ASC"
	ListInternalOccmDemandSignalCatalogsSortOrderDesc ListInternalOccmDemandSignalCatalogsSortOrderEnum = "DESC"
)

var mappingListInternalOccmDemandSignalCatalogsSortOrderEnum = map[string]ListInternalOccmDemandSignalCatalogsSortOrderEnum{
	"ASC":  ListInternalOccmDemandSignalCatalogsSortOrderAsc,
	"DESC": ListInternalOccmDemandSignalCatalogsSortOrderDesc,
}

var mappingListInternalOccmDemandSignalCatalogsSortOrderEnumLowerCase = map[string]ListInternalOccmDemandSignalCatalogsSortOrderEnum{
	"asc":  ListInternalOccmDemandSignalCatalogsSortOrderAsc,
	"desc": ListInternalOccmDemandSignalCatalogsSortOrderDesc,
}

// GetListInternalOccmDemandSignalCatalogsSortOrderEnumValues Enumerates the set of values for ListInternalOccmDemandSignalCatalogsSortOrderEnum
func GetListInternalOccmDemandSignalCatalogsSortOrderEnumValues() []ListInternalOccmDemandSignalCatalogsSortOrderEnum {
	values := make([]ListInternalOccmDemandSignalCatalogsSortOrderEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalCatalogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalCatalogsSortOrderEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalCatalogsSortOrderEnum
func GetListInternalOccmDemandSignalCatalogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalOccmDemandSignalCatalogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalCatalogsSortOrderEnum(val string) (ListInternalOccmDemandSignalCatalogsSortOrderEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalCatalogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccmDemandSignalCatalogsSortByEnum Enum with underlying type: string
type ListInternalOccmDemandSignalCatalogsSortByEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalCatalogsSortByEnum
const (
	ListInternalOccmDemandSignalCatalogsSortByName ListInternalOccmDemandSignalCatalogsSortByEnum = "name"
)

var mappingListInternalOccmDemandSignalCatalogsSortByEnum = map[string]ListInternalOccmDemandSignalCatalogsSortByEnum{
	"name": ListInternalOccmDemandSignalCatalogsSortByName,
}

var mappingListInternalOccmDemandSignalCatalogsSortByEnumLowerCase = map[string]ListInternalOccmDemandSignalCatalogsSortByEnum{
	"name": ListInternalOccmDemandSignalCatalogsSortByName,
}

// GetListInternalOccmDemandSignalCatalogsSortByEnumValues Enumerates the set of values for ListInternalOccmDemandSignalCatalogsSortByEnum
func GetListInternalOccmDemandSignalCatalogsSortByEnumValues() []ListInternalOccmDemandSignalCatalogsSortByEnum {
	values := make([]ListInternalOccmDemandSignalCatalogsSortByEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalCatalogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalCatalogsSortByEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalCatalogsSortByEnum
func GetListInternalOccmDemandSignalCatalogsSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListInternalOccmDemandSignalCatalogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalCatalogsSortByEnum(val string) (ListInternalOccmDemandSignalCatalogsSortByEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalCatalogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
