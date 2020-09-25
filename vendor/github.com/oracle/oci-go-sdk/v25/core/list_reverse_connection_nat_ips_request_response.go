// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"net/http"
)

// ListReverseConnectionNatIpsRequest wrapper for the ListReverseConnectionNatIps operation
type ListReverseConnectionNatIpsRequest struct {

	// The private endpoint's OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	PrivateEndpointId *string `mandatory:"true" contributesTo:"path" name:"privateEndpointId"`

	// The reverse connection NAT IP address that corresponds to a customer's IP address.
	ReverseConnectionNatIp *string `mandatory:"false" contributesTo:"query" name:"reverseConnectionNatIp"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListReverseConnectionNatIpsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListReverseConnectionNatIpsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.
	LifecycleState ReverseConnectionNatIpLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListReverseConnectionNatIpsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListReverseConnectionNatIpsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListReverseConnectionNatIpsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListReverseConnectionNatIpsResponse wrapper for the ListReverseConnectionNatIps operation
type ListReverseConnectionNatIpsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ReverseConnectionNatIp instances
	Items []ReverseConnectionNatIp `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListReverseConnectionNatIpsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListReverseConnectionNatIpsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListReverseConnectionNatIpsSortByEnum Enum with underlying type: string
type ListReverseConnectionNatIpsSortByEnum string

// Set of constants representing the allowable values for ListReverseConnectionNatIpsSortByEnum
const (
	ListReverseConnectionNatIpsSortByTimecreated ListReverseConnectionNatIpsSortByEnum = "TIMECREATED"
	ListReverseConnectionNatIpsSortByDisplayname ListReverseConnectionNatIpsSortByEnum = "DISPLAYNAME"
)

var mappingListReverseConnectionNatIpsSortBy = map[string]ListReverseConnectionNatIpsSortByEnum{
	"TIMECREATED": ListReverseConnectionNatIpsSortByTimecreated,
	"DISPLAYNAME": ListReverseConnectionNatIpsSortByDisplayname,
}

// GetListReverseConnectionNatIpsSortByEnumValues Enumerates the set of values for ListReverseConnectionNatIpsSortByEnum
func GetListReverseConnectionNatIpsSortByEnumValues() []ListReverseConnectionNatIpsSortByEnum {
	values := make([]ListReverseConnectionNatIpsSortByEnum, 0)
	for _, v := range mappingListReverseConnectionNatIpsSortBy {
		values = append(values, v)
	}
	return values
}

// ListReverseConnectionNatIpsSortOrderEnum Enum with underlying type: string
type ListReverseConnectionNatIpsSortOrderEnum string

// Set of constants representing the allowable values for ListReverseConnectionNatIpsSortOrderEnum
const (
	ListReverseConnectionNatIpsSortOrderAsc  ListReverseConnectionNatIpsSortOrderEnum = "ASC"
	ListReverseConnectionNatIpsSortOrderDesc ListReverseConnectionNatIpsSortOrderEnum = "DESC"
)

var mappingListReverseConnectionNatIpsSortOrder = map[string]ListReverseConnectionNatIpsSortOrderEnum{
	"ASC":  ListReverseConnectionNatIpsSortOrderAsc,
	"DESC": ListReverseConnectionNatIpsSortOrderDesc,
}

// GetListReverseConnectionNatIpsSortOrderEnumValues Enumerates the set of values for ListReverseConnectionNatIpsSortOrderEnum
func GetListReverseConnectionNatIpsSortOrderEnumValues() []ListReverseConnectionNatIpsSortOrderEnum {
	values := make([]ListReverseConnectionNatIpsSortOrderEnum, 0)
	for _, v := range mappingListReverseConnectionNatIpsSortOrder {
		values = append(values, v)
	}
	return values
}
