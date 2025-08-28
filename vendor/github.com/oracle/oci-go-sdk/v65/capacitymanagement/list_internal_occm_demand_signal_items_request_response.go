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

// ListInternalOccmDemandSignalItemsRequest wrapper for the ListInternalOccmDemandSignalItems operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignalItems.go.html to see an example of how to use ListInternalOccmDemandSignalItemsRequest.
type ListInternalOccmDemandSignalItemsRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The customer group ocid by which we would filter the list.
	OccCustomerGroupId *string `mandatory:"true" contributesTo:"query" name:"occCustomerGroupId"`

	// A query parameter to filter the list of demand signal items based on a demand signal id.
	OccmDemandSignalId *string `mandatory:"false" contributesTo:"query" name:"occmDemandSignalId"`

	// A query parameter to filter the list of demand signal details based on the resource name.
	ResourceName *string `mandatory:"false" contributesTo:"query" name:"resourceName"`

	// A query parameter to filter the list of demand signal details based on the namespace.
	DemandSignalNamespace ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum `mandatory:"false" contributesTo:"query" name:"demandSignalNamespace" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInternalOccmDemandSignalItemsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the response of List Demand Signal Details API. Only one sort order may be provided. The default order for resource name is case sensitive alphabetical order.
	SortBy ListInternalOccmDemandSignalItemsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalOccmDemandSignalItemsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalOccmDemandSignalItemsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalOccmDemandSignalItemsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalOccmDemandSignalItemsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalOccmDemandSignalItemsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum(string(request.DemandSignalNamespace)); !ok && request.DemandSignalNamespace != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DemandSignalNamespace: %s. Supported values are: %s.", request.DemandSignalNamespace, strings.Join(GetListInternalOccmDemandSignalItemsDemandSignalNamespaceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccmDemandSignalItemsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalOccmDemandSignalItemsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccmDemandSignalItemsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalOccmDemandSignalItemsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalOccmDemandSignalItemsResponse wrapper for the ListInternalOccmDemandSignalItems operation
type ListInternalOccmDemandSignalItemsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InternalOccmDemandSignalItemCollection instances
	InternalOccmDemandSignalItemCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInternalOccmDemandSignalItemsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalOccmDemandSignalItemsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum Enum with underlying type: string
type ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum
const (
	ListInternalOccmDemandSignalItemsDemandSignalNamespaceCompute ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum = "COMPUTE"
	ListInternalOccmDemandSignalItemsDemandSignalNamespaceNetwork ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum = "NETWORK"
	ListInternalOccmDemandSignalItemsDemandSignalNamespaceGpu     ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum = "GPU"
	ListInternalOccmDemandSignalItemsDemandSignalNamespaceStorage ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum = "STORAGE"
)

var mappingListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum = map[string]ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum{
	"COMPUTE": ListInternalOccmDemandSignalItemsDemandSignalNamespaceCompute,
	"NETWORK": ListInternalOccmDemandSignalItemsDemandSignalNamespaceNetwork,
	"GPU":     ListInternalOccmDemandSignalItemsDemandSignalNamespaceGpu,
	"STORAGE": ListInternalOccmDemandSignalItemsDemandSignalNamespaceStorage,
}

var mappingListInternalOccmDemandSignalItemsDemandSignalNamespaceEnumLowerCase = map[string]ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum{
	"compute": ListInternalOccmDemandSignalItemsDemandSignalNamespaceCompute,
	"network": ListInternalOccmDemandSignalItemsDemandSignalNamespaceNetwork,
	"gpu":     ListInternalOccmDemandSignalItemsDemandSignalNamespaceGpu,
	"storage": ListInternalOccmDemandSignalItemsDemandSignalNamespaceStorage,
}

// GetListInternalOccmDemandSignalItemsDemandSignalNamespaceEnumValues Enumerates the set of values for ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum
func GetListInternalOccmDemandSignalItemsDemandSignalNamespaceEnumValues() []ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum {
	values := make([]ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalItemsDemandSignalNamespaceEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum
func GetListInternalOccmDemandSignalItemsDemandSignalNamespaceEnumStringValues() []string {
	return []string{
		"COMPUTE",
		"NETWORK",
		"GPU",
		"STORAGE",
	}
}

// GetMappingListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum(val string) (ListInternalOccmDemandSignalItemsDemandSignalNamespaceEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalItemsDemandSignalNamespaceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccmDemandSignalItemsSortOrderEnum Enum with underlying type: string
type ListInternalOccmDemandSignalItemsSortOrderEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalItemsSortOrderEnum
const (
	ListInternalOccmDemandSignalItemsSortOrderAsc  ListInternalOccmDemandSignalItemsSortOrderEnum = "ASC"
	ListInternalOccmDemandSignalItemsSortOrderDesc ListInternalOccmDemandSignalItemsSortOrderEnum = "DESC"
)

var mappingListInternalOccmDemandSignalItemsSortOrderEnum = map[string]ListInternalOccmDemandSignalItemsSortOrderEnum{
	"ASC":  ListInternalOccmDemandSignalItemsSortOrderAsc,
	"DESC": ListInternalOccmDemandSignalItemsSortOrderDesc,
}

var mappingListInternalOccmDemandSignalItemsSortOrderEnumLowerCase = map[string]ListInternalOccmDemandSignalItemsSortOrderEnum{
	"asc":  ListInternalOccmDemandSignalItemsSortOrderAsc,
	"desc": ListInternalOccmDemandSignalItemsSortOrderDesc,
}

// GetListInternalOccmDemandSignalItemsSortOrderEnumValues Enumerates the set of values for ListInternalOccmDemandSignalItemsSortOrderEnum
func GetListInternalOccmDemandSignalItemsSortOrderEnumValues() []ListInternalOccmDemandSignalItemsSortOrderEnum {
	values := make([]ListInternalOccmDemandSignalItemsSortOrderEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalItemsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalItemsSortOrderEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalItemsSortOrderEnum
func GetListInternalOccmDemandSignalItemsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalOccmDemandSignalItemsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalItemsSortOrderEnum(val string) (ListInternalOccmDemandSignalItemsSortOrderEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalItemsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccmDemandSignalItemsSortByEnum Enum with underlying type: string
type ListInternalOccmDemandSignalItemsSortByEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalItemsSortByEnum
const (
	ListInternalOccmDemandSignalItemsSortByResourcename ListInternalOccmDemandSignalItemsSortByEnum = "resourceName"
)

var mappingListInternalOccmDemandSignalItemsSortByEnum = map[string]ListInternalOccmDemandSignalItemsSortByEnum{
	"resourceName": ListInternalOccmDemandSignalItemsSortByResourcename,
}

var mappingListInternalOccmDemandSignalItemsSortByEnumLowerCase = map[string]ListInternalOccmDemandSignalItemsSortByEnum{
	"resourcename": ListInternalOccmDemandSignalItemsSortByResourcename,
}

// GetListInternalOccmDemandSignalItemsSortByEnumValues Enumerates the set of values for ListInternalOccmDemandSignalItemsSortByEnum
func GetListInternalOccmDemandSignalItemsSortByEnumValues() []ListInternalOccmDemandSignalItemsSortByEnum {
	values := make([]ListInternalOccmDemandSignalItemsSortByEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalItemsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalItemsSortByEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalItemsSortByEnum
func GetListInternalOccmDemandSignalItemsSortByEnumStringValues() []string {
	return []string{
		"resourceName",
	}
}

// GetMappingListInternalOccmDemandSignalItemsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalItemsSortByEnum(val string) (ListInternalOccmDemandSignalItemsSortByEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalItemsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
