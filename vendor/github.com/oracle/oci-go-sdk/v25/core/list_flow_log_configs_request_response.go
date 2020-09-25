// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
	"net/http"
)

// ListFlowLogConfigsRequest wrapper for the ListFlowLogConfigs operation
type ListFlowLogConfigsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// For list pagination. The maximum number of results per page, or items to return in a paginated
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from the previous "List"
	// call. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given display name exactly.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	// TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	// sort order is case sensitive.
	// **Note:** In general, some "List" operations (for example, `ListInstances`) let you
	// optionally filter by availability domain if the scope of the resource type is within a
	// single availability domain. If you call one of these "List" operations without specifying
	// an availability domain, the resources are grouped by availability domain, then sorted.
	SortBy ListFlowLogConfigsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	// is case sensitive.
	SortOrder ListFlowLogConfigsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFlowLogConfigsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFlowLogConfigsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFlowLogConfigsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFlowLogConfigsResponse wrapper for the ListFlowLogConfigs operation
type ListFlowLogConfigsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FlowLogConfig instances
	Items []FlowLogConfig `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages
	// of results remain. For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListFlowLogConfigsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFlowLogConfigsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFlowLogConfigsSortByEnum Enum with underlying type: string
type ListFlowLogConfigsSortByEnum string

// Set of constants representing the allowable values for ListFlowLogConfigsSortByEnum
const (
	ListFlowLogConfigsSortByTimecreated ListFlowLogConfigsSortByEnum = "TIMECREATED"
	ListFlowLogConfigsSortByDisplayname ListFlowLogConfigsSortByEnum = "DISPLAYNAME"
)

var mappingListFlowLogConfigsSortBy = map[string]ListFlowLogConfigsSortByEnum{
	"TIMECREATED": ListFlowLogConfigsSortByTimecreated,
	"DISPLAYNAME": ListFlowLogConfigsSortByDisplayname,
}

// GetListFlowLogConfigsSortByEnumValues Enumerates the set of values for ListFlowLogConfigsSortByEnum
func GetListFlowLogConfigsSortByEnumValues() []ListFlowLogConfigsSortByEnum {
	values := make([]ListFlowLogConfigsSortByEnum, 0)
	for _, v := range mappingListFlowLogConfigsSortBy {
		values = append(values, v)
	}
	return values
}

// ListFlowLogConfigsSortOrderEnum Enum with underlying type: string
type ListFlowLogConfigsSortOrderEnum string

// Set of constants representing the allowable values for ListFlowLogConfigsSortOrderEnum
const (
	ListFlowLogConfigsSortOrderAsc  ListFlowLogConfigsSortOrderEnum = "ASC"
	ListFlowLogConfigsSortOrderDesc ListFlowLogConfigsSortOrderEnum = "DESC"
)

var mappingListFlowLogConfigsSortOrder = map[string]ListFlowLogConfigsSortOrderEnum{
	"ASC":  ListFlowLogConfigsSortOrderAsc,
	"DESC": ListFlowLogConfigsSortOrderDesc,
}

// GetListFlowLogConfigsSortOrderEnumValues Enumerates the set of values for ListFlowLogConfigsSortOrderEnum
func GetListFlowLogConfigsSortOrderEnumValues() []ListFlowLogConfigsSortOrderEnum {
	values := make([]ListFlowLogConfigsSortOrderEnum, 0)
	for _, v := range mappingListFlowLogConfigsSortOrder {
		values = append(values, v)
	}
	return values
}
