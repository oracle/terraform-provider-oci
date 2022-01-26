// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package jms

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListFleetsRequest wrapper for the ListFleets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/jms/ListFleets.go.html to see an example of how to use ListFleetsRequest.
type ListFleetsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of the Fleet.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The state of the lifecycle.
	LifecycleState ListFleetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. The token is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order, either 'asc' or 'desc'.
	SortOrder ListFleetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort Fleets. Only one sort order may be provided.
	// Default order for _timeCreated_, _approximateJreCount_, _approximateInstallationCount_,
	// _approximateApplicationCount_ and _approximateManagedInstanceCount_  is **descending**.
	// Default order for _displayName_ is **ascending**.
	// If no value is specified _timeCreated_ is default.
	SortBy ListFleetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFleetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFleetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFleetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFleetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListFleetsResponse wrapper for the ListFleets operation
type ListFleetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FleetCollection instances
	FleetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain.
	// Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	// For important details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFleetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFleetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFleetsLifecycleStateEnum Enum with underlying type: string
type ListFleetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFleetsLifecycleStateEnum
const (
	ListFleetsLifecycleStateActive   ListFleetsLifecycleStateEnum = "ACTIVE"
	ListFleetsLifecycleStateCreating ListFleetsLifecycleStateEnum = "CREATING"
	ListFleetsLifecycleStateDeleted  ListFleetsLifecycleStateEnum = "DELETED"
	ListFleetsLifecycleStateDeleting ListFleetsLifecycleStateEnum = "DELETING"
	ListFleetsLifecycleStateFailed   ListFleetsLifecycleStateEnum = "FAILED"
	ListFleetsLifecycleStateUpdating ListFleetsLifecycleStateEnum = "UPDATING"
)

var mappingListFleetsLifecycleState = map[string]ListFleetsLifecycleStateEnum{
	"ACTIVE":   ListFleetsLifecycleStateActive,
	"CREATING": ListFleetsLifecycleStateCreating,
	"DELETED":  ListFleetsLifecycleStateDeleted,
	"DELETING": ListFleetsLifecycleStateDeleting,
	"FAILED":   ListFleetsLifecycleStateFailed,
	"UPDATING": ListFleetsLifecycleStateUpdating,
}

// GetListFleetsLifecycleStateEnumValues Enumerates the set of values for ListFleetsLifecycleStateEnum
func GetListFleetsLifecycleStateEnumValues() []ListFleetsLifecycleStateEnum {
	values := make([]ListFleetsLifecycleStateEnum, 0)
	for _, v := range mappingListFleetsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListFleetsSortOrderEnum Enum with underlying type: string
type ListFleetsSortOrderEnum string

// Set of constants representing the allowable values for ListFleetsSortOrderEnum
const (
	ListFleetsSortOrderAsc  ListFleetsSortOrderEnum = "ASC"
	ListFleetsSortOrderDesc ListFleetsSortOrderEnum = "DESC"
)

var mappingListFleetsSortOrder = map[string]ListFleetsSortOrderEnum{
	"ASC":  ListFleetsSortOrderAsc,
	"DESC": ListFleetsSortOrderDesc,
}

// GetListFleetsSortOrderEnumValues Enumerates the set of values for ListFleetsSortOrderEnum
func GetListFleetsSortOrderEnumValues() []ListFleetsSortOrderEnum {
	values := make([]ListFleetsSortOrderEnum, 0)
	for _, v := range mappingListFleetsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListFleetsSortByEnum Enum with underlying type: string
type ListFleetsSortByEnum string

// Set of constants representing the allowable values for ListFleetsSortByEnum
const (
	ListFleetsSortByDisplayname ListFleetsSortByEnum = "displayName"
	ListFleetsSortByTimecreated ListFleetsSortByEnum = "timeCreated"
)

var mappingListFleetsSortBy = map[string]ListFleetsSortByEnum{
	"displayName": ListFleetsSortByDisplayname,
	"timeCreated": ListFleetsSortByTimecreated,
}

// GetListFleetsSortByEnumValues Enumerates the set of values for ListFleetsSortByEnum
func GetListFleetsSortByEnumValues() []ListFleetsSortByEnum {
	values := make([]ListFleetsSortByEnum, 0)
	for _, v := range mappingListFleetsSortBy {
		values = append(values, v)
	}
	return values
}
