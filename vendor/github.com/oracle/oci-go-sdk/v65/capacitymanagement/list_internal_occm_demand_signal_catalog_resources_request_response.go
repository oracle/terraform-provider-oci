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

// ListInternalOccmDemandSignalCatalogResourcesRequest wrapper for the ListInternalOccmDemandSignalCatalogResources operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignalCatalogResources.go.html to see an example of how to use ListInternalOccmDemandSignalCatalogResourcesRequest.
type ListInternalOccmDemandSignalCatalogResourcesRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The customer group ocid by which we would filter the list.
	OccCustomerGroupId *string `mandatory:"true" contributesTo:"query" name:"occCustomerGroupId"`

	// The ocid of demand signal catalog id.
	OccmDemandSignalCatalogId *string `mandatory:"true" contributesTo:"query" name:"occmDemandSignalCatalogId"`

	// A query parameter to filter the list of demand signal catalog resource based on the resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A query parameter to filter the list of demand signal catalog resources based on the namespace.
	DemandSignalNamespace ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum `mandatory:"false" contributesTo:"query" name:"demandSignalNamespace" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the response of List Demand catalog resources API. Only one sort order may be provided. The default order for name is case sensitive alphabetical order.
	SortBy ListInternalOccmDemandSignalCatalogResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalOccmDemandSignalCatalogResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalOccmDemandSignalCatalogResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalOccmDemandSignalCatalogResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalOccmDemandSignalCatalogResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalOccmDemandSignalCatalogResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum(string(request.DemandSignalNamespace)); !ok && request.DemandSignalNamespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DemandSignalNamespace: %s. Supported values are: %s.", request.DemandSignalNamespace, strings.Join(GetListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccmDemandSignalCatalogResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalOccmDemandSignalCatalogResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccmDemandSignalCatalogResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalOccmDemandSignalCatalogResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalOccmDemandSignalCatalogResourcesResponse wrapper for the ListInternalOccmDemandSignalCatalogResources operation
type ListInternalOccmDemandSignalCatalogResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InternalOccmDemandSignalCatalogResourceCollection instances
	InternalOccmDemandSignalCatalogResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInternalOccmDemandSignalCatalogResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalOccmDemandSignalCatalogResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum Enum with underlying type: string
type ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum
const (
	ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceCompute ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = "COMPUTE"
	ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceNetwork ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = "NETWORK"
	ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceGpu     ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = "GPU"
	ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceStorage ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = "STORAGE"
)

var mappingListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = map[string]ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum{
	"COMPUTE": ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceCompute,
	"NETWORK": ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceNetwork,
	"GPU":     ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceGpu,
	"STORAGE": ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceStorage,
}

var mappingListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumLowerCase = map[string]ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum{
	"compute": ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceCompute,
	"network": ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceNetwork,
	"gpu":     ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceGpu,
	"storage": ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceStorage,
}

// GetListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumValues Enumerates the set of values for ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum
func GetListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumValues() []ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum {
	values := make([]ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum
func GetListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"NETWORK",
		"GPU",
		"STORAGE",
	}
}

// GetMappingListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum(val string) (ListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum Enum with underlying type: string
type ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum
const (
	ListInternalOccmDemandSignalCatalogResourcesSortOrderAsc  ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum = "ASC"
	ListInternalOccmDemandSignalCatalogResourcesSortOrderDesc ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum = "DESC"
)

var mappingListInternalOccmDemandSignalCatalogResourcesSortOrderEnum = map[string]ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum{
	"ASC":  ListInternalOccmDemandSignalCatalogResourcesSortOrderAsc,
	"DESC": ListInternalOccmDemandSignalCatalogResourcesSortOrderDesc,
}

var mappingListInternalOccmDemandSignalCatalogResourcesSortOrderEnumLowerCase = map[string]ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum{
	"asc":  ListInternalOccmDemandSignalCatalogResourcesSortOrderAsc,
	"desc": ListInternalOccmDemandSignalCatalogResourcesSortOrderDesc,
}

// GetListInternalOccmDemandSignalCatalogResourcesSortOrderEnumValues Enumerates the set of values for ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum
func GetListInternalOccmDemandSignalCatalogResourcesSortOrderEnumValues() []ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum {
	values := make([]ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalCatalogResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalCatalogResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum
func GetListInternalOccmDemandSignalCatalogResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalOccmDemandSignalCatalogResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalCatalogResourcesSortOrderEnum(val string) (ListInternalOccmDemandSignalCatalogResourcesSortOrderEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalCatalogResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccmDemandSignalCatalogResourcesSortByEnum Enum with underlying type: string
type ListInternalOccmDemandSignalCatalogResourcesSortByEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalCatalogResourcesSortByEnum
const (
	ListInternalOccmDemandSignalCatalogResourcesSortByName ListInternalOccmDemandSignalCatalogResourcesSortByEnum = "name"
)

var mappingListInternalOccmDemandSignalCatalogResourcesSortByEnum = map[string]ListInternalOccmDemandSignalCatalogResourcesSortByEnum{
	"name": ListInternalOccmDemandSignalCatalogResourcesSortByName,
}

var mappingListInternalOccmDemandSignalCatalogResourcesSortByEnumLowerCase = map[string]ListInternalOccmDemandSignalCatalogResourcesSortByEnum{
	"name": ListInternalOccmDemandSignalCatalogResourcesSortByName,
}

// GetListInternalOccmDemandSignalCatalogResourcesSortByEnumValues Enumerates the set of values for ListInternalOccmDemandSignalCatalogResourcesSortByEnum
func GetListInternalOccmDemandSignalCatalogResourcesSortByEnumValues() []ListInternalOccmDemandSignalCatalogResourcesSortByEnum {
	values := make([]ListInternalOccmDemandSignalCatalogResourcesSortByEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalCatalogResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalCatalogResourcesSortByEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalCatalogResourcesSortByEnum
func GetListInternalOccmDemandSignalCatalogResourcesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListInternalOccmDemandSignalCatalogResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalCatalogResourcesSortByEnum(val string) (ListInternalOccmDemandSignalCatalogResourcesSortByEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalCatalogResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
