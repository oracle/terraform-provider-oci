// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package database

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListExadataInfrastructuresRequest wrapper for the ListExadataInfrastructures operation
type ListExadataInfrastructuresRequest struct {

	// The compartment OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The pagination token to continue listing from.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The field to sort by.  You can provide one sort order (`sortOrder`).  Default order for TIMECREATED is descending.  Default order for DISPLAYNAME is ascending. The DISPLAYNAME sort order is case sensitive.
	SortBy ListExadataInfrastructuresSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListExadataInfrastructuresSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// A filter to return only resources that match the given lifecycle state exactly.
	LifecycleState ExadataInfrastructureSummaryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExadataInfrastructuresRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExadataInfrastructuresRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExadataInfrastructuresRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListExadataInfrastructuresResponse wrapper for the ListExadataInfrastructures operation
type ListExadataInfrastructuresResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ExadataInfrastructureSummary instances
	Items []ExadataInfrastructureSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then there are additional items still to get. Include this value as the `page` parameter for the
	// subsequent GET request. For information about pagination, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExadataInfrastructuresResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExadataInfrastructuresResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExadataInfrastructuresSortByEnum Enum with underlying type: string
type ListExadataInfrastructuresSortByEnum string

// Set of constants representing the allowable values for ListExadataInfrastructuresSortByEnum
const (
	ListExadataInfrastructuresSortByTimecreated ListExadataInfrastructuresSortByEnum = "TIMECREATED"
	ListExadataInfrastructuresSortByDisplayname ListExadataInfrastructuresSortByEnum = "DISPLAYNAME"
)

var mappingListExadataInfrastructuresSortBy = map[string]ListExadataInfrastructuresSortByEnum{
	"TIMECREATED": ListExadataInfrastructuresSortByTimecreated,
	"DISPLAYNAME": ListExadataInfrastructuresSortByDisplayname,
}

// GetListExadataInfrastructuresSortByEnumValues Enumerates the set of values for ListExadataInfrastructuresSortByEnum
func GetListExadataInfrastructuresSortByEnumValues() []ListExadataInfrastructuresSortByEnum {
	values := make([]ListExadataInfrastructuresSortByEnum, 0)
	for _, v := range mappingListExadataInfrastructuresSortBy {
		values = append(values, v)
	}
	return values
}

// ListExadataInfrastructuresSortOrderEnum Enum with underlying type: string
type ListExadataInfrastructuresSortOrderEnum string

// Set of constants representing the allowable values for ListExadataInfrastructuresSortOrderEnum
const (
	ListExadataInfrastructuresSortOrderAsc  ListExadataInfrastructuresSortOrderEnum = "ASC"
	ListExadataInfrastructuresSortOrderDesc ListExadataInfrastructuresSortOrderEnum = "DESC"
)

var mappingListExadataInfrastructuresSortOrder = map[string]ListExadataInfrastructuresSortOrderEnum{
	"ASC":  ListExadataInfrastructuresSortOrderAsc,
	"DESC": ListExadataInfrastructuresSortOrderDesc,
}

// GetListExadataInfrastructuresSortOrderEnumValues Enumerates the set of values for ListExadataInfrastructuresSortOrderEnum
func GetListExadataInfrastructuresSortOrderEnumValues() []ListExadataInfrastructuresSortOrderEnum {
	values := make([]ListExadataInfrastructuresSortOrderEnum, 0)
	for _, v := range mappingListExadataInfrastructuresSortOrder {
		values = append(values, v)
	}
	return values
}
