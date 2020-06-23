// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListCatalogsRequest wrapper for the ListCatalogs operation
type ListCatalogsRequest struct {

	// The OCID of the compartment where you want to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListCatalogsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCatalogsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListCatalogsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCatalogsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCatalogsRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCatalogsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListCatalogsResponse wrapper for the ListCatalogs operation
type ListCatalogsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []CatalogSummary instances
	Items []CatalogSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCatalogsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCatalogsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCatalogsLifecycleStateEnum Enum with underlying type: string
type ListCatalogsLifecycleStateEnum string

// Set of constants representing the allowable values for ListCatalogsLifecycleStateEnum
const (
	ListCatalogsLifecycleStateCreating ListCatalogsLifecycleStateEnum = "CREATING"
	ListCatalogsLifecycleStateActive   ListCatalogsLifecycleStateEnum = "ACTIVE"
	ListCatalogsLifecycleStateInactive ListCatalogsLifecycleStateEnum = "INACTIVE"
	ListCatalogsLifecycleStateUpdating ListCatalogsLifecycleStateEnum = "UPDATING"
	ListCatalogsLifecycleStateDeleting ListCatalogsLifecycleStateEnum = "DELETING"
	ListCatalogsLifecycleStateDeleted  ListCatalogsLifecycleStateEnum = "DELETED"
	ListCatalogsLifecycleStateFailed   ListCatalogsLifecycleStateEnum = "FAILED"
	ListCatalogsLifecycleStateMoving   ListCatalogsLifecycleStateEnum = "MOVING"
)

var mappingListCatalogsLifecycleState = map[string]ListCatalogsLifecycleStateEnum{
	"CREATING": ListCatalogsLifecycleStateCreating,
	"ACTIVE":   ListCatalogsLifecycleStateActive,
	"INACTIVE": ListCatalogsLifecycleStateInactive,
	"UPDATING": ListCatalogsLifecycleStateUpdating,
	"DELETING": ListCatalogsLifecycleStateDeleting,
	"DELETED":  ListCatalogsLifecycleStateDeleted,
	"FAILED":   ListCatalogsLifecycleStateFailed,
	"MOVING":   ListCatalogsLifecycleStateMoving,
}

// GetListCatalogsLifecycleStateEnumValues Enumerates the set of values for ListCatalogsLifecycleStateEnum
func GetListCatalogsLifecycleStateEnumValues() []ListCatalogsLifecycleStateEnum {
	values := make([]ListCatalogsLifecycleStateEnum, 0)
	for _, v := range mappingListCatalogsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListCatalogsSortOrderEnum Enum with underlying type: string
type ListCatalogsSortOrderEnum string

// Set of constants representing the allowable values for ListCatalogsSortOrderEnum
const (
	ListCatalogsSortOrderAsc  ListCatalogsSortOrderEnum = "ASC"
	ListCatalogsSortOrderDesc ListCatalogsSortOrderEnum = "DESC"
)

var mappingListCatalogsSortOrder = map[string]ListCatalogsSortOrderEnum{
	"ASC":  ListCatalogsSortOrderAsc,
	"DESC": ListCatalogsSortOrderDesc,
}

// GetListCatalogsSortOrderEnumValues Enumerates the set of values for ListCatalogsSortOrderEnum
func GetListCatalogsSortOrderEnumValues() []ListCatalogsSortOrderEnum {
	values := make([]ListCatalogsSortOrderEnum, 0)
	for _, v := range mappingListCatalogsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListCatalogsSortByEnum Enum with underlying type: string
type ListCatalogsSortByEnum string

// Set of constants representing the allowable values for ListCatalogsSortByEnum
const (
	ListCatalogsSortByTimecreated ListCatalogsSortByEnum = "TIMECREATED"
	ListCatalogsSortByDisplayname ListCatalogsSortByEnum = "DISPLAYNAME"
)

var mappingListCatalogsSortBy = map[string]ListCatalogsSortByEnum{
	"TIMECREATED": ListCatalogsSortByTimecreated,
	"DISPLAYNAME": ListCatalogsSortByDisplayname,
}

// GetListCatalogsSortByEnumValues Enumerates the set of values for ListCatalogsSortByEnum
func GetListCatalogsSortByEnumValues() []ListCatalogsSortByEnum {
	values := make([]ListCatalogsSortByEnum, 0)
	for _, v := range mappingListCatalogsSortBy {
		values = append(values, v)
	}
	return values
}
