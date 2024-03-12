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

// ListOccAvailabilityCatalogsInternalRequest wrapper for the ListOccAvailabilityCatalogsInternal operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccAvailabilityCatalogsInternal.go.html to see an example of how to use ListOccAvailabilityCatalogsInternalRequest.
type ListOccAvailabilityCatalogsInternalRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The namespace by which we would filter the list.
	Namespace ListOccAvailabilityCatalogsInternalNamespaceEnum `mandatory:"false" contributesTo:"query" name:"namespace" omitEmpty:"true"`

	// The OCID of the availability catalog to filter the list of availability catalogs.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only the resources that match the entire display name. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter the list of availability catalogs based on the catalog state.
	CatalogState OccAvailabilityCatalogCatalogStateEnum `mandatory:"false" contributesTo:"query" name:"catalogState" omitEmpty:"true"`

	// The customer group ocid by which we would filter the list.
	OccCustomerGroupId *string `mandatory:"false" contributesTo:"query" name:"occCustomerGroupId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOccAvailabilityCatalogsInternalSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for displayName is ascending. The default order for timeCreated is descending.
	SortBy ListOccAvailabilityCatalogsInternalSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccAvailabilityCatalogsInternalRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccAvailabilityCatalogsInternalRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccAvailabilityCatalogsInternalRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccAvailabilityCatalogsInternalRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccAvailabilityCatalogsInternalRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccAvailabilityCatalogsInternalNamespaceEnum(string(request.Namespace)); !ok && request.Namespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Namespace: %s. Supported values are: %s.", request.Namespace, strings.Join(GetListOccAvailabilityCatalogsInternalNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccAvailabilityCatalogCatalogStateEnum(string(request.CatalogState)); !ok && request.CatalogState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CatalogState: %s. Supported values are: %s.", request.CatalogState, strings.Join(GetOccAvailabilityCatalogCatalogStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccAvailabilityCatalogsInternalSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccAvailabilityCatalogsInternalSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccAvailabilityCatalogsInternalSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccAvailabilityCatalogsInternalSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccAvailabilityCatalogsInternalResponse wrapper for the ListOccAvailabilityCatalogsInternal operation
type ListOccAvailabilityCatalogsInternalResponse struct {

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

func (response ListOccAvailabilityCatalogsInternalResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccAvailabilityCatalogsInternalResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccAvailabilityCatalogsInternalNamespaceEnum Enum with underlying type: string
type ListOccAvailabilityCatalogsInternalNamespaceEnum string

// Set of constants representing the allowable values for ListOccAvailabilityCatalogsInternalNamespaceEnum
const (
	ListOccAvailabilityCatalogsInternalNamespaceCompute ListOccAvailabilityCatalogsInternalNamespaceEnum = "COMPUTE"
)

var mappingListOccAvailabilityCatalogsInternalNamespaceEnum = map[string]ListOccAvailabilityCatalogsInternalNamespaceEnum{
	"COMPUTE": ListOccAvailabilityCatalogsInternalNamespaceCompute,
}

var mappingListOccAvailabilityCatalogsInternalNamespaceEnumLowerCase = map[string]ListOccAvailabilityCatalogsInternalNamespaceEnum{
	"compute": ListOccAvailabilityCatalogsInternalNamespaceCompute,
}

// GetListOccAvailabilityCatalogsInternalNamespaceEnumValues Enumerates the set of values for ListOccAvailabilityCatalogsInternalNamespaceEnum
func GetListOccAvailabilityCatalogsInternalNamespaceEnumValues() []ListOccAvailabilityCatalogsInternalNamespaceEnum {
	values := make([]ListOccAvailabilityCatalogsInternalNamespaceEnum, 0)
	for _, v := range mappingListOccAvailabilityCatalogsInternalNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccAvailabilityCatalogsInternalNamespaceEnumStringValues Enumerates the set of values in String for ListOccAvailabilityCatalogsInternalNamespaceEnum
func GetListOccAvailabilityCatalogsInternalNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingListOccAvailabilityCatalogsInternalNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccAvailabilityCatalogsInternalNamespaceEnum(val string) (ListOccAvailabilityCatalogsInternalNamespaceEnum, bool) {
	enum, ok := mappingListOccAvailabilityCatalogsInternalNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccAvailabilityCatalogsInternalSortOrderEnum Enum with underlying type: string
type ListOccAvailabilityCatalogsInternalSortOrderEnum string

// Set of constants representing the allowable values for ListOccAvailabilityCatalogsInternalSortOrderEnum
const (
	ListOccAvailabilityCatalogsInternalSortOrderAsc  ListOccAvailabilityCatalogsInternalSortOrderEnum = "ASC"
	ListOccAvailabilityCatalogsInternalSortOrderDesc ListOccAvailabilityCatalogsInternalSortOrderEnum = "DESC"
)

var mappingListOccAvailabilityCatalogsInternalSortOrderEnum = map[string]ListOccAvailabilityCatalogsInternalSortOrderEnum{
	"ASC":  ListOccAvailabilityCatalogsInternalSortOrderAsc,
	"DESC": ListOccAvailabilityCatalogsInternalSortOrderDesc,
}

var mappingListOccAvailabilityCatalogsInternalSortOrderEnumLowerCase = map[string]ListOccAvailabilityCatalogsInternalSortOrderEnum{
	"asc":  ListOccAvailabilityCatalogsInternalSortOrderAsc,
	"desc": ListOccAvailabilityCatalogsInternalSortOrderDesc,
}

// GetListOccAvailabilityCatalogsInternalSortOrderEnumValues Enumerates the set of values for ListOccAvailabilityCatalogsInternalSortOrderEnum
func GetListOccAvailabilityCatalogsInternalSortOrderEnumValues() []ListOccAvailabilityCatalogsInternalSortOrderEnum {
	values := make([]ListOccAvailabilityCatalogsInternalSortOrderEnum, 0)
	for _, v := range mappingListOccAvailabilityCatalogsInternalSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccAvailabilityCatalogsInternalSortOrderEnumStringValues Enumerates the set of values in String for ListOccAvailabilityCatalogsInternalSortOrderEnum
func GetListOccAvailabilityCatalogsInternalSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccAvailabilityCatalogsInternalSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccAvailabilityCatalogsInternalSortOrderEnum(val string) (ListOccAvailabilityCatalogsInternalSortOrderEnum, bool) {
	enum, ok := mappingListOccAvailabilityCatalogsInternalSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccAvailabilityCatalogsInternalSortByEnum Enum with underlying type: string
type ListOccAvailabilityCatalogsInternalSortByEnum string

// Set of constants representing the allowable values for ListOccAvailabilityCatalogsInternalSortByEnum
const (
	ListOccAvailabilityCatalogsInternalSortByDisplayname ListOccAvailabilityCatalogsInternalSortByEnum = "displayName"
	ListOccAvailabilityCatalogsInternalSortByTimecreated ListOccAvailabilityCatalogsInternalSortByEnum = "timeCreated"
)

var mappingListOccAvailabilityCatalogsInternalSortByEnum = map[string]ListOccAvailabilityCatalogsInternalSortByEnum{
	"displayName": ListOccAvailabilityCatalogsInternalSortByDisplayname,
	"timeCreated": ListOccAvailabilityCatalogsInternalSortByTimecreated,
}

var mappingListOccAvailabilityCatalogsInternalSortByEnumLowerCase = map[string]ListOccAvailabilityCatalogsInternalSortByEnum{
	"displayname": ListOccAvailabilityCatalogsInternalSortByDisplayname,
	"timecreated": ListOccAvailabilityCatalogsInternalSortByTimecreated,
}

// GetListOccAvailabilityCatalogsInternalSortByEnumValues Enumerates the set of values for ListOccAvailabilityCatalogsInternalSortByEnum
func GetListOccAvailabilityCatalogsInternalSortByEnumValues() []ListOccAvailabilityCatalogsInternalSortByEnum {
	values := make([]ListOccAvailabilityCatalogsInternalSortByEnum, 0)
	for _, v := range mappingListOccAvailabilityCatalogsInternalSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccAvailabilityCatalogsInternalSortByEnumStringValues Enumerates the set of values in String for ListOccAvailabilityCatalogsInternalSortByEnum
func GetListOccAvailabilityCatalogsInternalSortByEnumStringValues() []string {
	return []string{
		"displayName",
		"timeCreated",
	}
}

// GetMappingListOccAvailabilityCatalogsInternalSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccAvailabilityCatalogsInternalSortByEnum(val string) (ListOccAvailabilityCatalogsInternalSortByEnum, bool) {
	enum, ok := mappingListOccAvailabilityCatalogsInternalSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
