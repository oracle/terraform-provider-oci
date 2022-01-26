// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListTargetsRequest wrapper for the ListTargets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/ListTargets.go.html to see an example of how to use ListTargetsRequest.
type ListTargetsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The field life cycle state. Only one state can be provided. Default value for state is active. If no value is specified state is active.
	LifecycleState ListTargetsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed
	// and all compartments and subcompartments in the tenancy are
	// returned depending on the the setting of `accessLevel`.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are `RESTRICTED` and `ACCESSIBLE`. Default is `RESTRICTED`.
	// Setting this to `ACCESSIBLE` returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment).
	// When set to `RESTRICTED` permissions are checked and no partial results are displayed.
	AccessLevel ListTargetsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTargetsResponse wrapper for the ListTargets operation
type ListTargetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetCollection instances
	TargetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetsLifecycleStateEnum Enum with underlying type: string
type ListTargetsLifecycleStateEnum string

// Set of constants representing the allowable values for ListTargetsLifecycleStateEnum
const (
	ListTargetsLifecycleStateCreating ListTargetsLifecycleStateEnum = "CREATING"
	ListTargetsLifecycleStateUpdating ListTargetsLifecycleStateEnum = "UPDATING"
	ListTargetsLifecycleStateActive   ListTargetsLifecycleStateEnum = "ACTIVE"
	ListTargetsLifecycleStateInactive ListTargetsLifecycleStateEnum = "INACTIVE"
	ListTargetsLifecycleStateDeleting ListTargetsLifecycleStateEnum = "DELETING"
	ListTargetsLifecycleStateDeleted  ListTargetsLifecycleStateEnum = "DELETED"
	ListTargetsLifecycleStateFailed   ListTargetsLifecycleStateEnum = "FAILED"
)

var mappingListTargetsLifecycleState = map[string]ListTargetsLifecycleStateEnum{
	"CREATING": ListTargetsLifecycleStateCreating,
	"UPDATING": ListTargetsLifecycleStateUpdating,
	"ACTIVE":   ListTargetsLifecycleStateActive,
	"INACTIVE": ListTargetsLifecycleStateInactive,
	"DELETING": ListTargetsLifecycleStateDeleting,
	"DELETED":  ListTargetsLifecycleStateDeleted,
	"FAILED":   ListTargetsLifecycleStateFailed,
}

// GetListTargetsLifecycleStateEnumValues Enumerates the set of values for ListTargetsLifecycleStateEnum
func GetListTargetsLifecycleStateEnumValues() []ListTargetsLifecycleStateEnum {
	values := make([]ListTargetsLifecycleStateEnum, 0)
	for _, v := range mappingListTargetsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListTargetsAccessLevelEnum Enum with underlying type: string
type ListTargetsAccessLevelEnum string

// Set of constants representing the allowable values for ListTargetsAccessLevelEnum
const (
	ListTargetsAccessLevelRestricted ListTargetsAccessLevelEnum = "RESTRICTED"
	ListTargetsAccessLevelAccessible ListTargetsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListTargetsAccessLevel = map[string]ListTargetsAccessLevelEnum{
	"RESTRICTED": ListTargetsAccessLevelRestricted,
	"ACCESSIBLE": ListTargetsAccessLevelAccessible,
}

// GetListTargetsAccessLevelEnumValues Enumerates the set of values for ListTargetsAccessLevelEnum
func GetListTargetsAccessLevelEnumValues() []ListTargetsAccessLevelEnum {
	values := make([]ListTargetsAccessLevelEnum, 0)
	for _, v := range mappingListTargetsAccessLevel {
		values = append(values, v)
	}
	return values
}

// ListTargetsSortOrderEnum Enum with underlying type: string
type ListTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListTargetsSortOrderEnum
const (
	ListTargetsSortOrderAsc  ListTargetsSortOrderEnum = "ASC"
	ListTargetsSortOrderDesc ListTargetsSortOrderEnum = "DESC"
)

var mappingListTargetsSortOrder = map[string]ListTargetsSortOrderEnum{
	"ASC":  ListTargetsSortOrderAsc,
	"DESC": ListTargetsSortOrderDesc,
}

// GetListTargetsSortOrderEnumValues Enumerates the set of values for ListTargetsSortOrderEnum
func GetListTargetsSortOrderEnumValues() []ListTargetsSortOrderEnum {
	values := make([]ListTargetsSortOrderEnum, 0)
	for _, v := range mappingListTargetsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListTargetsSortByEnum Enum with underlying type: string
type ListTargetsSortByEnum string

// Set of constants representing the allowable values for ListTargetsSortByEnum
const (
	ListTargetsSortByTimecreated ListTargetsSortByEnum = "timeCreated"
	ListTargetsSortByDisplayname ListTargetsSortByEnum = "displayName"
)

var mappingListTargetsSortBy = map[string]ListTargetsSortByEnum{
	"timeCreated": ListTargetsSortByTimecreated,
	"displayName": ListTargetsSortByDisplayname,
}

// GetListTargetsSortByEnumValues Enumerates the set of values for ListTargetsSortByEnum
func GetListTargetsSortByEnumValues() []ListTargetsSortByEnum {
	values := make([]ListTargetsSortByEnum, 0)
	for _, v := range mappingListTargetsSortBy {
		values = append(values, v)
	}
	return values
}
