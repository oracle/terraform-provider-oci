// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListNatGatewaysRequest wrapper for the ListNatGateways operation
type ListNatGatewaysRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID of the VCN.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Filter results by the specified lifecycle state. Must be a valid
	// state for the resource type. This feature is currently in preview and may change before public release. Do not use it for production workloads.
	LifecycleState ListNatGatewaysLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by Availability Domain if the scope of the resource type is within a
	// single Availability Domain. If you call one of these "List" operations without specifying
	// an Availability Domain, the resources are grouped by Availability Domain, then sorted.
	SortBy ListNatGatewaysSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListNatGatewaysSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListNatGatewaysRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListNatGatewaysRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListNatGatewaysRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListNatGatewaysResponse wrapper for the ListNatGateways operation
type ListNatGatewaysResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []NatGateway instances
	Items []NatGateway `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListNatGatewaysResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListNatGatewaysResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListNatGatewaysLifecycleStateEnum Enum with underlying type: string
type ListNatGatewaysLifecycleStateEnum string

// Set of constants representing the allowable values for ListNatGatewaysLifecycleState
const (
	ListNatGatewaysLifecycleStateProvisioning ListNatGatewaysLifecycleStateEnum = "PROVISIONING"
	ListNatGatewaysLifecycleStateAvailable    ListNatGatewaysLifecycleStateEnum = "AVAILABLE"
	ListNatGatewaysLifecycleStateTerminating  ListNatGatewaysLifecycleStateEnum = "TERMINATING"
	ListNatGatewaysLifecycleStateTerminated   ListNatGatewaysLifecycleStateEnum = "TERMINATED"
)

var mappingListNatGatewaysLifecycleState = map[string]ListNatGatewaysLifecycleStateEnum{
	"PROVISIONING": ListNatGatewaysLifecycleStateProvisioning,
	"AVAILABLE":    ListNatGatewaysLifecycleStateAvailable,
	"TERMINATING":  ListNatGatewaysLifecycleStateTerminating,
	"TERMINATED":   ListNatGatewaysLifecycleStateTerminated,
}

// GetListNatGatewaysLifecycleStateEnumValues Enumerates the set of values for ListNatGatewaysLifecycleState
func GetListNatGatewaysLifecycleStateEnumValues() []ListNatGatewaysLifecycleStateEnum {
	values := make([]ListNatGatewaysLifecycleStateEnum, 0)
	for _, v := range mappingListNatGatewaysLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListNatGatewaysSortByEnum Enum with underlying type: string
type ListNatGatewaysSortByEnum string

// Set of constants representing the allowable values for ListNatGatewaysSortBy
const (
	ListNatGatewaysSortByTimecreated ListNatGatewaysSortByEnum = "TIMECREATED"
	ListNatGatewaysSortByDisplayname ListNatGatewaysSortByEnum = "DISPLAYNAME"
)

var mappingListNatGatewaysSortBy = map[string]ListNatGatewaysSortByEnum{
	"TIMECREATED": ListNatGatewaysSortByTimecreated,
	"DISPLAYNAME": ListNatGatewaysSortByDisplayname,
}

// GetListNatGatewaysSortByEnumValues Enumerates the set of values for ListNatGatewaysSortBy
func GetListNatGatewaysSortByEnumValues() []ListNatGatewaysSortByEnum {
	values := make([]ListNatGatewaysSortByEnum, 0)
	for _, v := range mappingListNatGatewaysSortBy {
		values = append(values, v)
	}
	return values
}

// ListNatGatewaysSortOrderEnum Enum with underlying type: string
type ListNatGatewaysSortOrderEnum string

// Set of constants representing the allowable values for ListNatGatewaysSortOrder
const (
	ListNatGatewaysSortOrderAsc  ListNatGatewaysSortOrderEnum = "ASC"
	ListNatGatewaysSortOrderDesc ListNatGatewaysSortOrderEnum = "DESC"
)

var mappingListNatGatewaysSortOrder = map[string]ListNatGatewaysSortOrderEnum{
	"ASC":  ListNatGatewaysSortOrderAsc,
	"DESC": ListNatGatewaysSortOrderDesc,
}

// GetListNatGatewaysSortOrderEnumValues Enumerates the set of values for ListNatGatewaysSortOrder
func GetListNatGatewaysSortOrderEnumValues() []ListNatGatewaysSortOrderEnum {
	values := make([]ListNatGatewaysSortOrderEnum, 0)
	for _, v := range mappingListNatGatewaysSortOrder {
		values = append(values, v)
	}
	return values
}
