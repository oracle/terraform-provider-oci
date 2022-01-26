// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDbManagementPrivateEndpointsRequest wrapper for the ListDbManagementPrivateEndpoints operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/ListDbManagementPrivateEndpoints.go.html to see an example of how to use ListDbManagementPrivateEndpointsRequest.
type ListDbManagementPrivateEndpointsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire name.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"false" contributesTo:"query" name:"vcnId"`

	// The option to filter Database Management private endpoints that can used for Oracle Databases in a cluster. This should be used along with the vcnId query parameter.
	IsCluster *bool `mandatory:"false" contributesTo:"query" name:"isCluster"`

	// The lifecycle state of a resource.
	LifecycleState ListDbManagementPrivateEndpointsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of records returned in the paginated response.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page from where the next set of paginated results
	// are retrieved. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The option to sort information in ascending (‘ASC’) or descending (‘DESC’) order. Ascending order is the default order.
	SortOrder ListDbManagementPrivateEndpointsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort information by. Only one sortOrder can be used. The default sort order
	// for ‘TIMECREATED’ is descending and the default sort order for ‘NAME’ is ascending.
	// The ‘NAME’ sort order is case-sensitive.
	SortBy ListDbManagementPrivateEndpointsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDbManagementPrivateEndpointsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDbManagementPrivateEndpointsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDbManagementPrivateEndpointsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDbManagementPrivateEndpointsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDbManagementPrivateEndpointsResponse wrapper for the ListDbManagementPrivateEndpoints operation
type ListDbManagementPrivateEndpointsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DbManagementPrivateEndpointCollection instances
	DbManagementPrivateEndpointCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDbManagementPrivateEndpointsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDbManagementPrivateEndpointsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDbManagementPrivateEndpointsLifecycleStateEnum Enum with underlying type: string
type ListDbManagementPrivateEndpointsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDbManagementPrivateEndpointsLifecycleStateEnum
const (
	ListDbManagementPrivateEndpointsLifecycleStateCreating ListDbManagementPrivateEndpointsLifecycleStateEnum = "CREATING"
	ListDbManagementPrivateEndpointsLifecycleStateUpdating ListDbManagementPrivateEndpointsLifecycleStateEnum = "UPDATING"
	ListDbManagementPrivateEndpointsLifecycleStateActive   ListDbManagementPrivateEndpointsLifecycleStateEnum = "ACTIVE"
	ListDbManagementPrivateEndpointsLifecycleStateDeleting ListDbManagementPrivateEndpointsLifecycleStateEnum = "DELETING"
	ListDbManagementPrivateEndpointsLifecycleStateDeleted  ListDbManagementPrivateEndpointsLifecycleStateEnum = "DELETED"
	ListDbManagementPrivateEndpointsLifecycleStateFailed   ListDbManagementPrivateEndpointsLifecycleStateEnum = "FAILED"
)

var mappingListDbManagementPrivateEndpointsLifecycleState = map[string]ListDbManagementPrivateEndpointsLifecycleStateEnum{
	"CREATING": ListDbManagementPrivateEndpointsLifecycleStateCreating,
	"UPDATING": ListDbManagementPrivateEndpointsLifecycleStateUpdating,
	"ACTIVE":   ListDbManagementPrivateEndpointsLifecycleStateActive,
	"DELETING": ListDbManagementPrivateEndpointsLifecycleStateDeleting,
	"DELETED":  ListDbManagementPrivateEndpointsLifecycleStateDeleted,
	"FAILED":   ListDbManagementPrivateEndpointsLifecycleStateFailed,
}

// GetListDbManagementPrivateEndpointsLifecycleStateEnumValues Enumerates the set of values for ListDbManagementPrivateEndpointsLifecycleStateEnum
func GetListDbManagementPrivateEndpointsLifecycleStateEnumValues() []ListDbManagementPrivateEndpointsLifecycleStateEnum {
	values := make([]ListDbManagementPrivateEndpointsLifecycleStateEnum, 0)
	for _, v := range mappingListDbManagementPrivateEndpointsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDbManagementPrivateEndpointsSortOrderEnum Enum with underlying type: string
type ListDbManagementPrivateEndpointsSortOrderEnum string

// Set of constants representing the allowable values for ListDbManagementPrivateEndpointsSortOrderEnum
const (
	ListDbManagementPrivateEndpointsSortOrderAsc  ListDbManagementPrivateEndpointsSortOrderEnum = "ASC"
	ListDbManagementPrivateEndpointsSortOrderDesc ListDbManagementPrivateEndpointsSortOrderEnum = "DESC"
)

var mappingListDbManagementPrivateEndpointsSortOrder = map[string]ListDbManagementPrivateEndpointsSortOrderEnum{
	"ASC":  ListDbManagementPrivateEndpointsSortOrderAsc,
	"DESC": ListDbManagementPrivateEndpointsSortOrderDesc,
}

// GetListDbManagementPrivateEndpointsSortOrderEnumValues Enumerates the set of values for ListDbManagementPrivateEndpointsSortOrderEnum
func GetListDbManagementPrivateEndpointsSortOrderEnumValues() []ListDbManagementPrivateEndpointsSortOrderEnum {
	values := make([]ListDbManagementPrivateEndpointsSortOrderEnum, 0)
	for _, v := range mappingListDbManagementPrivateEndpointsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDbManagementPrivateEndpointsSortByEnum Enum with underlying type: string
type ListDbManagementPrivateEndpointsSortByEnum string

// Set of constants representing the allowable values for ListDbManagementPrivateEndpointsSortByEnum
const (
	ListDbManagementPrivateEndpointsSortByTimecreated ListDbManagementPrivateEndpointsSortByEnum = "TIMECREATED"
	ListDbManagementPrivateEndpointsSortByName        ListDbManagementPrivateEndpointsSortByEnum = "NAME"
)

var mappingListDbManagementPrivateEndpointsSortBy = map[string]ListDbManagementPrivateEndpointsSortByEnum{
	"TIMECREATED": ListDbManagementPrivateEndpointsSortByTimecreated,
	"NAME":        ListDbManagementPrivateEndpointsSortByName,
}

// GetListDbManagementPrivateEndpointsSortByEnumValues Enumerates the set of values for ListDbManagementPrivateEndpointsSortByEnum
func GetListDbManagementPrivateEndpointsSortByEnumValues() []ListDbManagementPrivateEndpointsSortByEnum {
	values := make([]ListDbManagementPrivateEndpointsSortByEnum, 0)
	for _, v := range mappingListDbManagementPrivateEndpointsSortBy {
		values = append(values, v)
	}
	return values
}
