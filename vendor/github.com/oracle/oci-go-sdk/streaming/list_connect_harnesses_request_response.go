// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package streaming

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListConnectHarnessesRequest wrapper for the ListConnectHarnesses operation
type ListConnectHarnessesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the given ID exactly.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// A filter to return only resources that match the given name exactly.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return. The value must be between 1 and 50. The default is 10.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. You can provide no more than one sort order. By default, `TIMECREATED` sorts results in descending order and `NAME` sorts results in ascending order.
	SortBy ListConnectHarnessesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListConnectHarnessesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to only return resources that match the given lifecycle state. The state value is case-insensitive.
	LifecycleState ConnectHarnessSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListConnectHarnessesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListConnectHarnessesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListConnectHarnessesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListConnectHarnessesResponse wrapper for the ListConnectHarnesses operation
type ListConnectHarnessesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ConnectHarnessSummary instances
	Items []ConnectHarnessSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For list pagination. When this header appears in the response, previous pages of results exist. For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListConnectHarnessesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListConnectHarnessesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListConnectHarnessesSortByEnum Enum with underlying type: string
type ListConnectHarnessesSortByEnum string

// Set of constants representing the allowable values for ListConnectHarnessesSortByEnum
const (
	ListConnectHarnessesSortByName        ListConnectHarnessesSortByEnum = "NAME"
	ListConnectHarnessesSortByTimecreated ListConnectHarnessesSortByEnum = "TIMECREATED"
)

var mappingListConnectHarnessesSortBy = map[string]ListConnectHarnessesSortByEnum{
	"NAME":        ListConnectHarnessesSortByName,
	"TIMECREATED": ListConnectHarnessesSortByTimecreated,
}

// GetListConnectHarnessesSortByEnumValues Enumerates the set of values for ListConnectHarnessesSortByEnum
func GetListConnectHarnessesSortByEnumValues() []ListConnectHarnessesSortByEnum {
	values := make([]ListConnectHarnessesSortByEnum, 0)
	for _, v := range mappingListConnectHarnessesSortBy {
		values = append(values, v)
	}
	return values
}

// ListConnectHarnessesSortOrderEnum Enum with underlying type: string
type ListConnectHarnessesSortOrderEnum string

// Set of constants representing the allowable values for ListConnectHarnessesSortOrderEnum
const (
	ListConnectHarnessesSortOrderAsc  ListConnectHarnessesSortOrderEnum = "ASC"
	ListConnectHarnessesSortOrderDesc ListConnectHarnessesSortOrderEnum = "DESC"
)

var mappingListConnectHarnessesSortOrder = map[string]ListConnectHarnessesSortOrderEnum{
	"ASC":  ListConnectHarnessesSortOrderAsc,
	"DESC": ListConnectHarnessesSortOrderDesc,
}

// GetListConnectHarnessesSortOrderEnumValues Enumerates the set of values for ListConnectHarnessesSortOrderEnum
func GetListConnectHarnessesSortOrderEnumValues() []ListConnectHarnessesSortOrderEnum {
	values := make([]ListConnectHarnessesSortOrderEnum, 0)
	for _, v := range mappingListConnectHarnessesSortOrder {
		values = append(values, v)
	}
	return values
}
