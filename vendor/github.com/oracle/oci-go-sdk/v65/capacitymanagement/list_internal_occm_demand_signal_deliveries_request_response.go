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

// ListInternalOccmDemandSignalDeliveriesRequest wrapper for the ListInternalOccmDemandSignalDeliveries operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListInternalOccmDemandSignalDeliveries.go.html to see an example of how to use ListInternalOccmDemandSignalDeliveriesRequest.
type ListInternalOccmDemandSignalDeliveriesRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The customer group ocid by which we would filter the list.
	OccCustomerGroupId *string `mandatory:"true" contributesTo:"query" name:"occCustomerGroupId"`

	// A query parameter to filter the list of demand signals based on it's OCID.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A query parameter to filter the list of demand signal items based on it's OCID.
	OccmDemandSignalItemId *string `mandatory:"false" contributesTo:"query" name:"occmDemandSignalItemId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing. The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListInternalOccmDemandSignalDeliveriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the response of List Demand Signal Delivery API. Only one sort order may be provided. The default order for resource name is case sensitive alphabetical order.
	SortBy ListInternalOccmDemandSignalDeliveriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListInternalOccmDemandSignalDeliveriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListInternalOccmDemandSignalDeliveriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListInternalOccmDemandSignalDeliveriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListInternalOccmDemandSignalDeliveriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListInternalOccmDemandSignalDeliveriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListInternalOccmDemandSignalDeliveriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListInternalOccmDemandSignalDeliveriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListInternalOccmDemandSignalDeliveriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListInternalOccmDemandSignalDeliveriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListInternalOccmDemandSignalDeliveriesResponse wrapper for the ListInternalOccmDemandSignalDeliveries operation
type ListInternalOccmDemandSignalDeliveriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of InternalOccmDemandSignalDeliveryCollection instances
	InternalOccmDemandSignalDeliveryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListInternalOccmDemandSignalDeliveriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListInternalOccmDemandSignalDeliveriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListInternalOccmDemandSignalDeliveriesSortOrderEnum Enum with underlying type: string
type ListInternalOccmDemandSignalDeliveriesSortOrderEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalDeliveriesSortOrderEnum
const (
	ListInternalOccmDemandSignalDeliveriesSortOrderAsc  ListInternalOccmDemandSignalDeliveriesSortOrderEnum = "ASC"
	ListInternalOccmDemandSignalDeliveriesSortOrderDesc ListInternalOccmDemandSignalDeliveriesSortOrderEnum = "DESC"
)

var mappingListInternalOccmDemandSignalDeliveriesSortOrderEnum = map[string]ListInternalOccmDemandSignalDeliveriesSortOrderEnum{
	"ASC":  ListInternalOccmDemandSignalDeliveriesSortOrderAsc,
	"DESC": ListInternalOccmDemandSignalDeliveriesSortOrderDesc,
}

var mappingListInternalOccmDemandSignalDeliveriesSortOrderEnumLowerCase = map[string]ListInternalOccmDemandSignalDeliveriesSortOrderEnum{
	"asc":  ListInternalOccmDemandSignalDeliveriesSortOrderAsc,
	"desc": ListInternalOccmDemandSignalDeliveriesSortOrderDesc,
}

// GetListInternalOccmDemandSignalDeliveriesSortOrderEnumValues Enumerates the set of values for ListInternalOccmDemandSignalDeliveriesSortOrderEnum
func GetListInternalOccmDemandSignalDeliveriesSortOrderEnumValues() []ListInternalOccmDemandSignalDeliveriesSortOrderEnum {
	values := make([]ListInternalOccmDemandSignalDeliveriesSortOrderEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalDeliveriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalDeliveriesSortOrderEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalDeliveriesSortOrderEnum
func GetListInternalOccmDemandSignalDeliveriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListInternalOccmDemandSignalDeliveriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalDeliveriesSortOrderEnum(val string) (ListInternalOccmDemandSignalDeliveriesSortOrderEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalDeliveriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListInternalOccmDemandSignalDeliveriesSortByEnum Enum with underlying type: string
type ListInternalOccmDemandSignalDeliveriesSortByEnum string

// Set of constants representing the allowable values for ListInternalOccmDemandSignalDeliveriesSortByEnum
const (
	ListInternalOccmDemandSignalDeliveriesSortByResourcename ListInternalOccmDemandSignalDeliveriesSortByEnum = "resourceName"
)

var mappingListInternalOccmDemandSignalDeliveriesSortByEnum = map[string]ListInternalOccmDemandSignalDeliveriesSortByEnum{
	"resourceName": ListInternalOccmDemandSignalDeliveriesSortByResourcename,
}

var mappingListInternalOccmDemandSignalDeliveriesSortByEnumLowerCase = map[string]ListInternalOccmDemandSignalDeliveriesSortByEnum{
	"resourcename": ListInternalOccmDemandSignalDeliveriesSortByResourcename,
}

// GetListInternalOccmDemandSignalDeliveriesSortByEnumValues Enumerates the set of values for ListInternalOccmDemandSignalDeliveriesSortByEnum
func GetListInternalOccmDemandSignalDeliveriesSortByEnumValues() []ListInternalOccmDemandSignalDeliveriesSortByEnum {
	values := make([]ListInternalOccmDemandSignalDeliveriesSortByEnum, 0)
	for _, v := range mappingListInternalOccmDemandSignalDeliveriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListInternalOccmDemandSignalDeliveriesSortByEnumStringValues Enumerates the set of values in String for ListInternalOccmDemandSignalDeliveriesSortByEnum
func GetListInternalOccmDemandSignalDeliveriesSortByEnumStringValues() []string {
	return []string{
		"resourceName",
	}
}

// GetMappingListInternalOccmDemandSignalDeliveriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListInternalOccmDemandSignalDeliveriesSortByEnum(val string) (ListInternalOccmDemandSignalDeliveriesSortByEnum, bool) {
	enum, ok := mappingListInternalOccmDemandSignalDeliveriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
