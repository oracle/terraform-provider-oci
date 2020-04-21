// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListBdsInstancesRequest wrapper for the ListBdsInstances operation
type ListBdsInstancesRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The state of the BDS instance.
	LifecycleState BdsInstanceLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListBdsInstancesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListBdsInstancesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListBdsInstancesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListBdsInstancesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListBdsInstancesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListBdsInstancesResponse wrapper for the ListBdsInstances operation
type ListBdsInstancesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []BdsInstanceSummary instances
	Items []BdsInstanceSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListBdsInstancesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListBdsInstancesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListBdsInstancesSortByEnum Enum with underlying type: string
type ListBdsInstancesSortByEnum string

// Set of constants representing the allowable values for ListBdsInstancesSortByEnum
const (
	ListBdsInstancesSortByTimecreated ListBdsInstancesSortByEnum = "timeCreated"
	ListBdsInstancesSortByDisplayname ListBdsInstancesSortByEnum = "displayName"
)

var mappingListBdsInstancesSortBy = map[string]ListBdsInstancesSortByEnum{
	"timeCreated": ListBdsInstancesSortByTimecreated,
	"displayName": ListBdsInstancesSortByDisplayname,
}

// GetListBdsInstancesSortByEnumValues Enumerates the set of values for ListBdsInstancesSortByEnum
func GetListBdsInstancesSortByEnumValues() []ListBdsInstancesSortByEnum {
	values := make([]ListBdsInstancesSortByEnum, 0)
	for _, v := range mappingListBdsInstancesSortBy {
		values = append(values, v)
	}
	return values
}

// ListBdsInstancesSortOrderEnum Enum with underlying type: string
type ListBdsInstancesSortOrderEnum string

// Set of constants representing the allowable values for ListBdsInstancesSortOrderEnum
const (
	ListBdsInstancesSortOrderAsc  ListBdsInstancesSortOrderEnum = "ASC"
	ListBdsInstancesSortOrderDesc ListBdsInstancesSortOrderEnum = "DESC"
)

var mappingListBdsInstancesSortOrder = map[string]ListBdsInstancesSortOrderEnum{
	"ASC":  ListBdsInstancesSortOrderAsc,
	"DESC": ListBdsInstancesSortOrderDesc,
}

// GetListBdsInstancesSortOrderEnumValues Enumerates the set of values for ListBdsInstancesSortOrderEnum
func GetListBdsInstancesSortOrderEnumValues() []ListBdsInstancesSortOrderEnum {
	values := make([]ListBdsInstancesSortOrderEnum, 0)
	for _, v := range mappingListBdsInstancesSortOrder {
		values = append(values, v)
	}
	return values
}
