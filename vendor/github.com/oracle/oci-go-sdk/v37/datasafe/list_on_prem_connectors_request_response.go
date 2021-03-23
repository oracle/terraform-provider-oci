// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v37/common"
	"net/http"
)

// ListOnPremConnectorsRequest wrapper for the ListOnPremConnectors operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListOnPremConnectors.go.html to see an example of how to use ListOnPremConnectorsRequest.
type ListOnPremConnectorsRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the on-premises connector that matches the specified id.
	OnPremConnectorId *string `mandatory:"false" contributesTo:"query" name:"onPremConnectorId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only on-premises connector resources that match the specified lifecycle state.
	OnPremConnectorLifecycleState ListOnPremConnectorsOnPremConnectorLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"onPremConnectorLifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListOnPremConnectorsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListOnPremConnectorsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOnPremConnectorsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOnPremConnectorsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOnPremConnectorsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListOnPremConnectorsResponse wrapper for the ListOnPremConnectors operation
type ListOnPremConnectorsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []OnPremConnectorSummary instances
	Items []OnPremConnectorSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListOnPremConnectorsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOnPremConnectorsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOnPremConnectorsOnPremConnectorLifecycleStateEnum Enum with underlying type: string
type ListOnPremConnectorsOnPremConnectorLifecycleStateEnum string

// Set of constants representing the allowable values for ListOnPremConnectorsOnPremConnectorLifecycleStateEnum
const (
	ListOnPremConnectorsOnPremConnectorLifecycleStateCreating ListOnPremConnectorsOnPremConnectorLifecycleStateEnum = "CREATING"
	ListOnPremConnectorsOnPremConnectorLifecycleStateUpdating ListOnPremConnectorsOnPremConnectorLifecycleStateEnum = "UPDATING"
	ListOnPremConnectorsOnPremConnectorLifecycleStateActive   ListOnPremConnectorsOnPremConnectorLifecycleStateEnum = "ACTIVE"
	ListOnPremConnectorsOnPremConnectorLifecycleStateInactive ListOnPremConnectorsOnPremConnectorLifecycleStateEnum = "INACTIVE"
	ListOnPremConnectorsOnPremConnectorLifecycleStateDeleting ListOnPremConnectorsOnPremConnectorLifecycleStateEnum = "DELETING"
	ListOnPremConnectorsOnPremConnectorLifecycleStateDeleted  ListOnPremConnectorsOnPremConnectorLifecycleStateEnum = "DELETED"
	ListOnPremConnectorsOnPremConnectorLifecycleStateFailed   ListOnPremConnectorsOnPremConnectorLifecycleStateEnum = "FAILED"
)

var mappingListOnPremConnectorsOnPremConnectorLifecycleState = map[string]ListOnPremConnectorsOnPremConnectorLifecycleStateEnum{
	"CREATING": ListOnPremConnectorsOnPremConnectorLifecycleStateCreating,
	"UPDATING": ListOnPremConnectorsOnPremConnectorLifecycleStateUpdating,
	"ACTIVE":   ListOnPremConnectorsOnPremConnectorLifecycleStateActive,
	"INACTIVE": ListOnPremConnectorsOnPremConnectorLifecycleStateInactive,
	"DELETING": ListOnPremConnectorsOnPremConnectorLifecycleStateDeleting,
	"DELETED":  ListOnPremConnectorsOnPremConnectorLifecycleStateDeleted,
	"FAILED":   ListOnPremConnectorsOnPremConnectorLifecycleStateFailed,
}

// GetListOnPremConnectorsOnPremConnectorLifecycleStateEnumValues Enumerates the set of values for ListOnPremConnectorsOnPremConnectorLifecycleStateEnum
func GetListOnPremConnectorsOnPremConnectorLifecycleStateEnumValues() []ListOnPremConnectorsOnPremConnectorLifecycleStateEnum {
	values := make([]ListOnPremConnectorsOnPremConnectorLifecycleStateEnum, 0)
	for _, v := range mappingListOnPremConnectorsOnPremConnectorLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListOnPremConnectorsSortOrderEnum Enum with underlying type: string
type ListOnPremConnectorsSortOrderEnum string

// Set of constants representing the allowable values for ListOnPremConnectorsSortOrderEnum
const (
	ListOnPremConnectorsSortOrderAsc  ListOnPremConnectorsSortOrderEnum = "ASC"
	ListOnPremConnectorsSortOrderDesc ListOnPremConnectorsSortOrderEnum = "DESC"
)

var mappingListOnPremConnectorsSortOrder = map[string]ListOnPremConnectorsSortOrderEnum{
	"ASC":  ListOnPremConnectorsSortOrderAsc,
	"DESC": ListOnPremConnectorsSortOrderDesc,
}

// GetListOnPremConnectorsSortOrderEnumValues Enumerates the set of values for ListOnPremConnectorsSortOrderEnum
func GetListOnPremConnectorsSortOrderEnumValues() []ListOnPremConnectorsSortOrderEnum {
	values := make([]ListOnPremConnectorsSortOrderEnum, 0)
	for _, v := range mappingListOnPremConnectorsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListOnPremConnectorsSortByEnum Enum with underlying type: string
type ListOnPremConnectorsSortByEnum string

// Set of constants representing the allowable values for ListOnPremConnectorsSortByEnum
const (
	ListOnPremConnectorsSortByTimecreated ListOnPremConnectorsSortByEnum = "TIMECREATED"
	ListOnPremConnectorsSortByDisplayname ListOnPremConnectorsSortByEnum = "DISPLAYNAME"
)

var mappingListOnPremConnectorsSortBy = map[string]ListOnPremConnectorsSortByEnum{
	"TIMECREATED": ListOnPremConnectorsSortByTimecreated,
	"DISPLAYNAME": ListOnPremConnectorsSortByDisplayname,
}

// GetListOnPremConnectorsSortByEnumValues Enumerates the set of values for ListOnPremConnectorsSortByEnum
func GetListOnPremConnectorsSortByEnumValues() []ListOnPremConnectorsSortByEnum {
	values := make([]ListOnPremConnectorsSortByEnum, 0)
	for _, v := range mappingListOnPremConnectorsSortBy {
		values = append(values, v)
	}
	return values
}
