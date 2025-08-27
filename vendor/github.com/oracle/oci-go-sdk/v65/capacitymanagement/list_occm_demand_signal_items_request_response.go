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

// ListOccmDemandSignalItemsRequest wrapper for the ListOccmDemandSignalItems operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccmDemandSignalItems.go.html to see an example of how to use ListOccmDemandSignalItemsRequest.
type ListOccmDemandSignalItemsRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A query parameter to filter the list of demand signal items based on a demand signal id.
	OccmDemandSignalId *string `mandatory:"false" contributesTo:"query" name:"occmDemandSignalId"`

	// A query parameter to filter the list of demand signal details based on the resource name.
	ResourceName *string `mandatory:"false" contributesTo:"query" name:"resourceName"`

	// A query parameter to filter the list of demand signal details based on the namespace.
	DemandSignalNamespace ListOccmDemandSignalItemsDemandSignalNamespaceEnum `mandatory:"false" contributesTo:"query" name:"demandSignalNamespace" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListOccmDemandSignalItemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the response of List Demand Signal Details API. Only one sort order may be provided. The default order for resource name is case sensitive alphabetical order.
	SortBy ListOccmDemandSignalItemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccmDemandSignalItemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccmDemandSignalItemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccmDemandSignalItemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccmDemandSignalItemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccmDemandSignalItemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccmDemandSignalItemsDemandSignalNamespaceEnum(string(request.DemandSignalNamespace)); !ok && request.DemandSignalNamespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DemandSignalNamespace: %s. Supported values are: %s.", request.DemandSignalNamespace, strings.Join(GetListOccmDemandSignalItemsDemandSignalNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccmDemandSignalItemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccmDemandSignalItemsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccmDemandSignalItemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccmDemandSignalItemsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccmDemandSignalItemsResponse wrapper for the ListOccmDemandSignalItems operation
type ListOccmDemandSignalItemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccmDemandSignalItemCollection instances
	OccmDemandSignalItemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccmDemandSignalItemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccmDemandSignalItemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccmDemandSignalItemsDemandSignalNamespaceEnum Enum with underlying type: string
type ListOccmDemandSignalItemsDemandSignalNamespaceEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalItemsDemandSignalNamespaceEnum
const (
	ListOccmDemandSignalItemsDemandSignalNamespaceCompute ListOccmDemandSignalItemsDemandSignalNamespaceEnum = "COMPUTE"
	ListOccmDemandSignalItemsDemandSignalNamespaceNetwork ListOccmDemandSignalItemsDemandSignalNamespaceEnum = "NETWORK"
	ListOccmDemandSignalItemsDemandSignalNamespaceGpu     ListOccmDemandSignalItemsDemandSignalNamespaceEnum = "GPU"
	ListOccmDemandSignalItemsDemandSignalNamespaceStorage ListOccmDemandSignalItemsDemandSignalNamespaceEnum = "STORAGE"
)

var mappingListOccmDemandSignalItemsDemandSignalNamespaceEnum = map[string]ListOccmDemandSignalItemsDemandSignalNamespaceEnum{
	"COMPUTE": ListOccmDemandSignalItemsDemandSignalNamespaceCompute,
	"NETWORK": ListOccmDemandSignalItemsDemandSignalNamespaceNetwork,
	"GPU":     ListOccmDemandSignalItemsDemandSignalNamespaceGpu,
	"STORAGE": ListOccmDemandSignalItemsDemandSignalNamespaceStorage,
}

var mappingListOccmDemandSignalItemsDemandSignalNamespaceEnumLowerCase = map[string]ListOccmDemandSignalItemsDemandSignalNamespaceEnum{
	"compute": ListOccmDemandSignalItemsDemandSignalNamespaceCompute,
	"network": ListOccmDemandSignalItemsDemandSignalNamespaceNetwork,
	"gpu":     ListOccmDemandSignalItemsDemandSignalNamespaceGpu,
	"storage": ListOccmDemandSignalItemsDemandSignalNamespaceStorage,
}

// GetListOccmDemandSignalItemsDemandSignalNamespaceEnumValues Enumerates the set of values for ListOccmDemandSignalItemsDemandSignalNamespaceEnum
func GetListOccmDemandSignalItemsDemandSignalNamespaceEnumValues() []ListOccmDemandSignalItemsDemandSignalNamespaceEnum {
	values := make([]ListOccmDemandSignalItemsDemandSignalNamespaceEnum, 0)
	for _, v := range mappingListOccmDemandSignalItemsDemandSignalNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalItemsDemandSignalNamespaceEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalItemsDemandSignalNamespaceEnum
func GetListOccmDemandSignalItemsDemandSignalNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"NETWORK",
		"GPU",
		"STORAGE",
	}
}

// GetMappingListOccmDemandSignalItemsDemandSignalNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalItemsDemandSignalNamespaceEnum(val string) (ListOccmDemandSignalItemsDemandSignalNamespaceEnum, bool) {
	enum, ok := mappingListOccmDemandSignalItemsDemandSignalNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccmDemandSignalItemsSortOrderEnum Enum with underlying type: string
type ListOccmDemandSignalItemsSortOrderEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalItemsSortOrderEnum
const (
	ListOccmDemandSignalItemsSortOrderAsc  ListOccmDemandSignalItemsSortOrderEnum = "ASC"
	ListOccmDemandSignalItemsSortOrderDesc ListOccmDemandSignalItemsSortOrderEnum = "DESC"
)

var mappingListOccmDemandSignalItemsSortOrderEnum = map[string]ListOccmDemandSignalItemsSortOrderEnum{
	"ASC":  ListOccmDemandSignalItemsSortOrderAsc,
	"DESC": ListOccmDemandSignalItemsSortOrderDesc,
}

var mappingListOccmDemandSignalItemsSortOrderEnumLowerCase = map[string]ListOccmDemandSignalItemsSortOrderEnum{
	"asc":  ListOccmDemandSignalItemsSortOrderAsc,
	"desc": ListOccmDemandSignalItemsSortOrderDesc,
}

// GetListOccmDemandSignalItemsSortOrderEnumValues Enumerates the set of values for ListOccmDemandSignalItemsSortOrderEnum
func GetListOccmDemandSignalItemsSortOrderEnumValues() []ListOccmDemandSignalItemsSortOrderEnum {
	values := make([]ListOccmDemandSignalItemsSortOrderEnum, 0)
	for _, v := range mappingListOccmDemandSignalItemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalItemsSortOrderEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalItemsSortOrderEnum
func GetListOccmDemandSignalItemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccmDemandSignalItemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalItemsSortOrderEnum(val string) (ListOccmDemandSignalItemsSortOrderEnum, bool) {
	enum, ok := mappingListOccmDemandSignalItemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccmDemandSignalItemsSortByEnum Enum with underlying type: string
type ListOccmDemandSignalItemsSortByEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalItemsSortByEnum
const (
	ListOccmDemandSignalItemsSortByResourcename ListOccmDemandSignalItemsSortByEnum = "resourceName"
)

var mappingListOccmDemandSignalItemsSortByEnum = map[string]ListOccmDemandSignalItemsSortByEnum{
	"resourceName": ListOccmDemandSignalItemsSortByResourcename,
}

var mappingListOccmDemandSignalItemsSortByEnumLowerCase = map[string]ListOccmDemandSignalItemsSortByEnum{
	"resourcename": ListOccmDemandSignalItemsSortByResourcename,
}

// GetListOccmDemandSignalItemsSortByEnumValues Enumerates the set of values for ListOccmDemandSignalItemsSortByEnum
func GetListOccmDemandSignalItemsSortByEnumValues() []ListOccmDemandSignalItemsSortByEnum {
	values := make([]ListOccmDemandSignalItemsSortByEnum, 0)
	for _, v := range mappingListOccmDemandSignalItemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalItemsSortByEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalItemsSortByEnum
func GetListOccmDemandSignalItemsSortByEnumStringValues() []string {
	return []string{
		"resourceName",
	}
}

// GetMappingListOccmDemandSignalItemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalItemsSortByEnum(val string) (ListOccmDemandSignalItemsSortByEnum, bool) {
	enum, ok := mappingListOccmDemandSignalItemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
