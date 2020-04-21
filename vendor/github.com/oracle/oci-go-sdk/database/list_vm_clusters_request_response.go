// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListVmClustersRequest wrapper for the ListVmClusters operation
type ListVmClustersRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// If provided, filters the results for the given Exadata Infrastructure.
	ExadataInfrastructureId *string `mandatory:"false" contributesTo:"query" name:"exadataInfrastructureId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListVmClustersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListVmClustersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState VmClusterSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVmClustersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVmClustersRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVmClustersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListVmClustersResponse wrapper for the ListVmClusters operation
type ListVmClustersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []VmClusterSummary instances
	Items []VmClusterSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVmClustersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVmClustersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVmClustersSortByEnum Enum with underlying type: string
type ListVmClustersSortByEnum string

// Set of constants representing the allowable values for ListVmClustersSortByEnum
const (
	ListVmClustersSortByTimecreated ListVmClustersSortByEnum = "TIMECREATED"
	ListVmClustersSortByDisplayname ListVmClustersSortByEnum = "DISPLAYNAME"
)

var mappingListVmClustersSortBy = map[string]ListVmClustersSortByEnum{
	"TIMECREATED": ListVmClustersSortByTimecreated,
	"DISPLAYNAME": ListVmClustersSortByDisplayname,
}

// GetListVmClustersSortByEnumValues Enumerates the set of values for ListVmClustersSortByEnum
func GetListVmClustersSortByEnumValues() []ListVmClustersSortByEnum {
	values := make([]ListVmClustersSortByEnum, 0)
	for _, v := range mappingListVmClustersSortBy {
		values = append(values, v)
	}
	return values
}

// ListVmClustersSortOrderEnum Enum with underlying type: string
type ListVmClustersSortOrderEnum string

// Set of constants representing the allowable values for ListVmClustersSortOrderEnum
const (
	ListVmClustersSortOrderAsc  ListVmClustersSortOrderEnum = "ASC"
	ListVmClustersSortOrderDesc ListVmClustersSortOrderEnum = "DESC"
)

var mappingListVmClustersSortOrder = map[string]ListVmClustersSortOrderEnum{
	"ASC":  ListVmClustersSortOrderAsc,
	"DESC": ListVmClustersSortOrderDesc,
}

// GetListVmClustersSortOrderEnumValues Enumerates the set of values for ListVmClustersSortOrderEnum
func GetListVmClustersSortOrderEnumValues() []ListVmClustersSortOrderEnum {
	values := make([]ListVmClustersSortOrderEnum, 0)
	for _, v := range mappingListVmClustersSortOrder {
		values = append(values, v)
	}
	return values
}
