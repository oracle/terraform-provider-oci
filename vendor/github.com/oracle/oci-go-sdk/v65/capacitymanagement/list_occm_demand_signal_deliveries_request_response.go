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

// ListOccmDemandSignalDeliveriesRequest wrapper for the ListOccmDemandSignalDeliveries operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/capacitymanagement/ListOccmDemandSignalDeliveries.go.html to see an example of how to use ListOccmDemandSignalDeliveriesRequest.
type ListOccmDemandSignalDeliveriesRequest struct {

	// The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

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
	SortOrder ListOccmDemandSignalDeliveriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort the response of List Demand Signal Delivery API. Only one sort order may be provided. The default order for resource name is case sensitive alphabetical order.
	SortBy ListOccmDemandSignalDeliveriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOccmDemandSignalDeliveriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOccmDemandSignalDeliveriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOccmDemandSignalDeliveriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOccmDemandSignalDeliveriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOccmDemandSignalDeliveriesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOccmDemandSignalDeliveriesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOccmDemandSignalDeliveriesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOccmDemandSignalDeliveriesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOccmDemandSignalDeliveriesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOccmDemandSignalDeliveriesResponse wrapper for the ListOccmDemandSignalDeliveries operation
type ListOccmDemandSignalDeliveriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OccmDemandSignalDeliveryCollection instances
	OccmDemandSignalDeliveryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOccmDemandSignalDeliveriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOccmDemandSignalDeliveriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOccmDemandSignalDeliveriesSortOrderEnum Enum with underlying type: string
type ListOccmDemandSignalDeliveriesSortOrderEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalDeliveriesSortOrderEnum
const (
	ListOccmDemandSignalDeliveriesSortOrderAsc  ListOccmDemandSignalDeliveriesSortOrderEnum = "ASC"
	ListOccmDemandSignalDeliveriesSortOrderDesc ListOccmDemandSignalDeliveriesSortOrderEnum = "DESC"
)

var mappingListOccmDemandSignalDeliveriesSortOrderEnum = map[string]ListOccmDemandSignalDeliveriesSortOrderEnum{
	"ASC":  ListOccmDemandSignalDeliveriesSortOrderAsc,
	"DESC": ListOccmDemandSignalDeliveriesSortOrderDesc,
}

var mappingListOccmDemandSignalDeliveriesSortOrderEnumLowerCase = map[string]ListOccmDemandSignalDeliveriesSortOrderEnum{
	"asc":  ListOccmDemandSignalDeliveriesSortOrderAsc,
	"desc": ListOccmDemandSignalDeliveriesSortOrderDesc,
}

// GetListOccmDemandSignalDeliveriesSortOrderEnumValues Enumerates the set of values for ListOccmDemandSignalDeliveriesSortOrderEnum
func GetListOccmDemandSignalDeliveriesSortOrderEnumValues() []ListOccmDemandSignalDeliveriesSortOrderEnum {
	values := make([]ListOccmDemandSignalDeliveriesSortOrderEnum, 0)
	for _, v := range mappingListOccmDemandSignalDeliveriesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalDeliveriesSortOrderEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalDeliveriesSortOrderEnum
func GetListOccmDemandSignalDeliveriesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOccmDemandSignalDeliveriesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalDeliveriesSortOrderEnum(val string) (ListOccmDemandSignalDeliveriesSortOrderEnum, bool) {
	enum, ok := mappingListOccmDemandSignalDeliveriesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOccmDemandSignalDeliveriesSortByEnum Enum with underlying type: string
type ListOccmDemandSignalDeliveriesSortByEnum string

// Set of constants representing the allowable values for ListOccmDemandSignalDeliveriesSortByEnum
const (
	ListOccmDemandSignalDeliveriesSortByResourcename ListOccmDemandSignalDeliveriesSortByEnum = "resourceName"
)

var mappingListOccmDemandSignalDeliveriesSortByEnum = map[string]ListOccmDemandSignalDeliveriesSortByEnum{
	"resourceName": ListOccmDemandSignalDeliveriesSortByResourcename,
}

var mappingListOccmDemandSignalDeliveriesSortByEnumLowerCase = map[string]ListOccmDemandSignalDeliveriesSortByEnum{
	"resourcename": ListOccmDemandSignalDeliveriesSortByResourcename,
}

// GetListOccmDemandSignalDeliveriesSortByEnumValues Enumerates the set of values for ListOccmDemandSignalDeliveriesSortByEnum
func GetListOccmDemandSignalDeliveriesSortByEnumValues() []ListOccmDemandSignalDeliveriesSortByEnum {
	values := make([]ListOccmDemandSignalDeliveriesSortByEnum, 0)
	for _, v := range mappingListOccmDemandSignalDeliveriesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOccmDemandSignalDeliveriesSortByEnumStringValues Enumerates the set of values in String for ListOccmDemandSignalDeliveriesSortByEnum
func GetListOccmDemandSignalDeliveriesSortByEnumStringValues() []string {
	return []string{
		"resourceName",
	}
}

// GetMappingListOccmDemandSignalDeliveriesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOccmDemandSignalDeliveriesSortByEnum(val string) (ListOccmDemandSignalDeliveriesSortByEnum, bool) {
	enum, ok := mappingListOccmDemandSignalDeliveriesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
