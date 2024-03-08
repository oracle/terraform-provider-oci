// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOccAvailabilityCatalogsRequest wrapper for the ListOccAvailabilityCatalogs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccAvailabilityCatalogs.go.html to see an example of how to use ListOccAvailabilityCatalogsRequest.
type ListOccAvailabilityCatalogsRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The namespace by which we would filter the list.
	Namespace ListOccAvailabilityCatalogsNamespaceEnum `mandatory:"false" contributesTo:"query" name:"namespace" omitEmpty:"true"`

	// The OCID of the availability catalog to filter the list of availability catalogs.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only the resources that match the entire display name. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter the list of availability catalogs based on the catalog state.
	CatalogState OccAvailabilityCatalogCatalogStateEnum `mandatory:"false" contributesTo:"query" name:"catalogState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOccAvailabilityCatalogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for displayName is ascending. The default order for timeCreated is descending.
	SortBy ListOccAvailabilityCatalogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccAvailabilityCatalogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccAvailabilityCatalogsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccAvailabilityCatalogsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccAvailabilityCatalogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccAvailabilityCatalogsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccAvailabilityCatalogsNamespaceEnum(string(request.Namespace)); !ok && request.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", request.Namespace, strings.Join(GetListOccAvailabilityCatalogsNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccAvailabilityCatalogCatalogStateEnum(string(request.CatalogState)); !ok && request.CatalogState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CatalogState: %s. Supported values are: %s.", request.CatalogState, strings.Join(GetOccAvailabilityCatalogCatalogStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccAvailabilityCatalogsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccAvailabilityCatalogsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccAvailabilityCatalogsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccAvailabilityCatalogsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccAvailabilityCatalogsResponse wrapper for the ListOccAvailabilityCatalogs operation
type ListOccAvailabilityCatalogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccAvailabilityCatalogCollection instances
	OccAvailabilityCatalogCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccAvailabilityCatalogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccAvailabilityCatalogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccAvailabilityCatalogsNamespaceEnum Enum with underlying type: string
type ListOccAvailabilityCatalogsNamespaceEnum string

// Set of constants representing the allowable values for ListOccAvailabilityCatalogsNamespaceEnum
const (
	ListOccAvailabilityCatalogsNamespaceCompute ListOccAvailabilityCatalogsNamespaceEnum = "COMPUTE"
)

var mappingListOccAvailabilityCatalogsNamespaceEnum = map[string]ListOccAvailabilityCatalogsNamespaceEnum{
	"COMPUTE": ListOccAvailabilityCatalogsNamespaceCompute,
}

var mappingListOccAvailabilityCatalogsNamespaceEnumLowerCase = map[string]ListOccAvailabilityCatalogsNamespaceEnum{
	"compute": ListOccAvailabilityCatalogsNamespaceCompute,
}

// GetListOccAvailabilityCatalogsNamespaceEnumValues Enumerates the set of values for ListOccAvailabilityCatalogsNamespaceEnum
func GetListOccAvailabilityCatalogsNamespaceEnumValues() []ListOccAvailabilityCatalogsNamespaceEnum {
	values := make([]ListOccAvailabilityCatalogsNamespaceEnum, 0)
	for _, v := range mappingListOccAvailabilityCatalogsNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccAvailabilityCatalogsNamespaceEnumStringValues Enumerates the set of values in String for ListOccAvailabilityCatalogsNamespaceEnum
func GetListOccAvailabilityCatalogsNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingListOccAvailabilityCatalogsNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccAvailabilityCatalogsNamespaceEnum(val string) (ListOccAvailabilityCatalogsNamespaceEnum, bool) {
	enum, ok := mappingListOccAvailabilityCatalogsNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccAvailabilityCatalogsSortOrderEnum Enum with underlying type: string
type ListOccAvailabilityCatalogsSortOrderEnum string

// Set of constants representing the allowable values for ListOccAvailabilityCatalogsSortOrderEnum
const (
	ListOccAvailabilityCatalogsSortOrderAsc  ListOccAvailabilityCatalogsSortOrderEnum = "ASC"
	ListOccAvailabilityCatalogsSortOrderDesc ListOccAvailabilityCatalogsSortOrderEnum = "DESC"
)

var mappingListOccAvailabilityCatalogsSortOrderEnum = map[string]ListOccAvailabilityCatalogsSortOrderEnum{
	"ASC":  ListOccAvailabilityCatalogsSortOrderAsc,
	"DESC": ListOccAvailabilityCatalogsSortOrderDesc,
}

var mappingListOccAvailabilityCatalogsSortOrderEnumLowerCase = map[string]ListOccAvailabilityCatalogsSortOrderEnum{
	"asc":  ListOccAvailabilityCatalogsSortOrderAsc,
	"desc": ListOccAvailabilityCatalogsSortOrderDesc,
}

// GetListOccAvailabilityCatalogsSortOrderEnumValues Enumerates the set of values for ListOccAvailabilityCatalogsSortOrderEnum
func GetListOccAvailabilityCatalogsSortOrderEnumValues() []ListOccAvailabilityCatalogsSortOrderEnum {
	values := make([]ListOccAvailabilityCatalogsSortOrderEnum, 0)
	for _, v := range mappingListOccAvailabilityCatalogsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccAvailabilityCatalogsSortOrderEnumStringValues Enumerates the set of values in String for ListOccAvailabilityCatalogsSortOrderEnum
func GetListOccAvailabilityCatalogsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccAvailabilityCatalogsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccAvailabilityCatalogsSortOrderEnum(val string) (ListOccAvailabilityCatalogsSortOrderEnum, bool) {
	enum, ok := mappingListOccAvailabilityCatalogsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccAvailabilityCatalogsSortByEnum Enum with underlying type: string
type ListOccAvailabilityCatalogsSortByEnum string

// Set of constants representing the allowable values for ListOccAvailabilityCatalogsSortByEnum
const (
	ListOccAvailabilityCatalogsSortByDisplayname ListOccAvailabilityCatalogsSortByEnum = "displayName"
	ListOccAvailabilityCatalogsSortByTimecreated ListOccAvailabilityCatalogsSortByEnum = "timeCreated"
)

var mappingListOccAvailabilityCatalogsSortByEnum = map[string]ListOccAvailabilityCatalogsSortByEnum{
	"displayName": ListOccAvailabilityCatalogsSortByDisplayname,
	"timeCreated": ListOccAvailabilityCatalogsSortByTimecreated,
}

var mappingListOccAvailabilityCatalogsSortByEnumLowerCase = map[string]ListOccAvailabilityCatalogsSortByEnum{
	"displayname": ListOccAvailabilityCatalogsSortByDisplayname,
	"timecreated": ListOccAvailabilityCatalogsSortByTimecreated,
}

// GetListOccAvailabilityCatalogsSortByEnumValues Enumerates the set of values for ListOccAvailabilityCatalogsSortByEnum
func GetListOccAvailabilityCatalogsSortByEnumValues() []ListOccAvailabilityCatalogsSortByEnum {
	values := make([]ListOccAvailabilityCatalogsSortByEnum, 0)
	for _, v := range mappingListOccAvailabilityCatalogsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccAvailabilityCatalogsSortByEnumStringValues Enumerates the set of values in String for ListOccAvailabilityCatalogsSortByEnum
func GetListOccAvailabilityCatalogsSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListOccAvailabilityCatalogsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccAvailabilityCatalogsSortByEnum(val string) (ListOccAvailabilityCatalogsSortByEnum, bool) {
	enum, ok := mappingListOccAvailabilityCatalogsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
