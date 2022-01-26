// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListOperationsInsightsWarehousesRequest wrapper for the ListOperationsInsightsWarehouses operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListOperationsInsightsWarehouses.go.html to see an example of how to use ListOperationsInsightsWarehousesRequest.
type ListOperationsInsightsWarehousesRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique Operations Insights Warehouse identifier
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// Lifecycle states
	LifecycleState []OperationsInsightsWarehouseLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListOperationsInsightsWarehousesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListOperationsInsightsWarehousesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOperationsInsightsWarehousesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOperationsInsightsWarehousesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOperationsInsightsWarehousesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOperationsInsightsWarehousesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOperationsInsightsWarehousesResponse wrapper for the ListOperationsInsightsWarehouses operation
type ListOperationsInsightsWarehousesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OperationsInsightsWarehouseSummaryCollection instances
	OperationsInsightsWarehouseSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOperationsInsightsWarehousesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOperationsInsightsWarehousesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOperationsInsightsWarehousesSortOrderEnum Enum with underlying type: string
type ListOperationsInsightsWarehousesSortOrderEnum string

// Set of constants representing the allowable values for ListOperationsInsightsWarehousesSortOrderEnum
const (
	ListOperationsInsightsWarehousesSortOrderAsc  ListOperationsInsightsWarehousesSortOrderEnum = "ASC"
	ListOperationsInsightsWarehousesSortOrderDesc ListOperationsInsightsWarehousesSortOrderEnum = "DESC"
)

var mappingListOperationsInsightsWarehousesSortOrder = map[string]ListOperationsInsightsWarehousesSortOrderEnum{
	"ASC":  ListOperationsInsightsWarehousesSortOrderAsc,
	"DESC": ListOperationsInsightsWarehousesSortOrderDesc,
}

// GetListOperationsInsightsWarehousesSortOrderEnumValues Enumerates the set of values for ListOperationsInsightsWarehousesSortOrderEnum
func GetListOperationsInsightsWarehousesSortOrderEnumValues() []ListOperationsInsightsWarehousesSortOrderEnum {
	values := make([]ListOperationsInsightsWarehousesSortOrderEnum, 0)
	for _, v := range mappingListOperationsInsightsWarehousesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListOperationsInsightsWarehousesSortByEnum Enum with underlying type: string
type ListOperationsInsightsWarehousesSortByEnum string

// Set of constants representing the allowable values for ListOperationsInsightsWarehousesSortByEnum
const (
	ListOperationsInsightsWarehousesSortByTimecreated ListOperationsInsightsWarehousesSortByEnum = "timeCreated"
	ListOperationsInsightsWarehousesSortByDisplayname ListOperationsInsightsWarehousesSortByEnum = "displayName"
)

var mappingListOperationsInsightsWarehousesSortBy = map[string]ListOperationsInsightsWarehousesSortByEnum{
	"timeCreated": ListOperationsInsightsWarehousesSortByTimecreated,
	"displayName": ListOperationsInsightsWarehousesSortByDisplayname,
}

// GetListOperationsInsightsWarehousesSortByEnumValues Enumerates the set of values for ListOperationsInsightsWarehousesSortByEnum
func GetListOperationsInsightsWarehousesSortByEnumValues() []ListOperationsInsightsWarehousesSortByEnum {
	values := make([]ListOperationsInsightsWarehousesSortByEnum, 0)
	for _, v := range mappingListOperationsInsightsWarehousesSortBy {
		values = append(values, v)
	}
	return values
}
