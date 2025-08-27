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

// ListOccmDemandSignalCatalogResourcesRequest wrapper for the ListOccmDemandSignalCatalogResources operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccmDemandSignalCatalogResources.go.html to see an example of how to use ListOccmDemandSignalCatalogResourcesRequest.
type ListOccmDemandSignalCatalogResourcesRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A query parameter to filter the list of demand signal catalog resource based on the resource name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// A query parameter to filter the list of demand signal catalog resources based on the namespace.
	DemandSignalNamespace ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum `mandatory:"false" contributesTo:"query" name:"demandSignalNamespace" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOccmDemandSignalCatalogResourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the response of List Demand catalog resources API. Only one sort order may be provided. The default order for name is case sensitive alphabetical order.
	SortBy ListOccmDemandSignalCatalogResourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccmDemandSignalCatalogResourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccmDemandSignalCatalogResourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccmDemandSignalCatalogResourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccmDemandSignalCatalogResourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccmDemandSignalCatalogResourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum(string(request.DemandSignalNamespace)); !ok && request.DemandSignalNamespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DemandSignalNamespace: %s. Supported values are: %s.", request.DemandSignalNamespace, strings.Join(GetListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccmDemandSignalCatalogResourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccmDemandSignalCatalogResourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccmDemandSignalCatalogResourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccmDemandSignalCatalogResourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccmDemandSignalCatalogResourcesResponse wrapper for the ListOccmDemandSignalCatalogResources operation
type ListOccmDemandSignalCatalogResourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccmDemandSignalCatalogResourceCollection instances
	OccmDemandSignalCatalogResourceCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccmDemandSignalCatalogResourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccmDemandSignalCatalogResourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum Enum with underlying type: string
type ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum
const (
	ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceCompute ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = "COMPUTE"
	ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceNetwork ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = "NETWORK"
	ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceGpu     ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = "GPU"
	ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceStorage ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = "STORAGE"
)

var mappingListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum = map[string]ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum{
	"COMPUTE": ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceCompute,
	"NETWORK": ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceNetwork,
	"GPU":     ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceGpu,
	"STORAGE": ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceStorage,
}

var mappingListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumLowerCase = map[string]ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum{
	"compute": ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceCompute,
	"network": ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceNetwork,
	"gpu":     ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceGpu,
	"storage": ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceStorage,
}

// GetListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumValues Enumerates the set of values for ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum
func GetListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumValues() []ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum {
	values := make([]ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum, 0)
	for _, v := range mappingListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum
func GetListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"NETWORK",
		"GPU",
		"STORAGE",
	}
}

// GetMappingListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum(val string) (ListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnum, bool) {
	enum, ok := mappingListOccmDemandSignalCatalogResourcesDemandSignalNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccmDemandSignalCatalogResourcesSortOrderEnum Enum with underlying type: string
type ListOccmDemandSignalCatalogResourcesSortOrderEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalCatalogResourcesSortOrderEnum
const (
	ListOccmDemandSignalCatalogResourcesSortOrderAsc  ListOccmDemandSignalCatalogResourcesSortOrderEnum = "ASC"
	ListOccmDemandSignalCatalogResourcesSortOrderDesc ListOccmDemandSignalCatalogResourcesSortOrderEnum = "DESC"
)

var mappingListOccmDemandSignalCatalogResourcesSortOrderEnum = map[string]ListOccmDemandSignalCatalogResourcesSortOrderEnum{
	"ASC":  ListOccmDemandSignalCatalogResourcesSortOrderAsc,
	"DESC": ListOccmDemandSignalCatalogResourcesSortOrderDesc,
}

var mappingListOccmDemandSignalCatalogResourcesSortOrderEnumLowerCase = map[string]ListOccmDemandSignalCatalogResourcesSortOrderEnum{
	"asc":  ListOccmDemandSignalCatalogResourcesSortOrderAsc,
	"desc": ListOccmDemandSignalCatalogResourcesSortOrderDesc,
}

// GetListOccmDemandSignalCatalogResourcesSortOrderEnumValues Enumerates the set of values for ListOccmDemandSignalCatalogResourcesSortOrderEnum
func GetListOccmDemandSignalCatalogResourcesSortOrderEnumValues() []ListOccmDemandSignalCatalogResourcesSortOrderEnum {
	values := make([]ListOccmDemandSignalCatalogResourcesSortOrderEnum, 0)
	for _, v := range mappingListOccmDemandSignalCatalogResourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalCatalogResourcesSortOrderEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalCatalogResourcesSortOrderEnum
func GetListOccmDemandSignalCatalogResourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccmDemandSignalCatalogResourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalCatalogResourcesSortOrderEnum(val string) (ListOccmDemandSignalCatalogResourcesSortOrderEnum, bool) {
	enum, ok := mappingListOccmDemandSignalCatalogResourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccmDemandSignalCatalogResourcesSortByEnum Enum with underlying type: string
type ListOccmDemandSignalCatalogResourcesSortByEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalCatalogResourcesSortByEnum
const (
	ListOccmDemandSignalCatalogResourcesSortByName ListOccmDemandSignalCatalogResourcesSortByEnum = "name"
)

var mappingListOccmDemandSignalCatalogResourcesSortByEnum = map[string]ListOccmDemandSignalCatalogResourcesSortByEnum{
	"name": ListOccmDemandSignalCatalogResourcesSortByName,
}

var mappingListOccmDemandSignalCatalogResourcesSortByEnumLowerCase = map[string]ListOccmDemandSignalCatalogResourcesSortByEnum{
	"name": ListOccmDemandSignalCatalogResourcesSortByName,
}

// GetListOccmDemandSignalCatalogResourcesSortByEnumValues Enumerates the set of values for ListOccmDemandSignalCatalogResourcesSortByEnum
func GetListOccmDemandSignalCatalogResourcesSortByEnumValues() []ListOccmDemandSignalCatalogResourcesSortByEnum {
	values := make([]ListOccmDemandSignalCatalogResourcesSortByEnum, 0)
	for _, v := range mappingListOccmDemandSignalCatalogResourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalCatalogResourcesSortByEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalCatalogResourcesSortByEnum
func GetListOccmDemandSignalCatalogResourcesSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingListOccmDemandSignalCatalogResourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalCatalogResourcesSortByEnum(val string) (ListOccmDemandSignalCatalogResourcesSortByEnum, bool) {
	enum, ok := mappingListOccmDemandSignalCatalogResourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
