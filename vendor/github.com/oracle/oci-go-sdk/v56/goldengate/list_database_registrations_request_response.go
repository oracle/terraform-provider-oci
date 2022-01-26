// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListDatabaseRegistrationsRequest wrapper for the ListDatabaseRegistrations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListDatabaseRegistrations.go.html to see an example of how to use ListDatabaseRegistrationsRequest.
type ListDatabaseRegistrationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return only the resources that match the 'lifecycleState' given.
	LifecycleState ListDatabaseRegistrationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only the resources that match the entire 'displayName' given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListDatabaseRegistrationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is descending.  Default order for 'displayName' is ascending. If no value is specified timeCreated is the default.
	SortBy ListDatabaseRegistrationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListDatabaseRegistrationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListDatabaseRegistrationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListDatabaseRegistrationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListDatabaseRegistrationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListDatabaseRegistrationsResponse wrapper for the ListDatabaseRegistrations operation
type ListDatabaseRegistrationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of DatabaseRegistrationCollection instances
	DatabaseRegistrationCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListDatabaseRegistrationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListDatabaseRegistrationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListDatabaseRegistrationsLifecycleStateEnum Enum with underlying type: string
type ListDatabaseRegistrationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListDatabaseRegistrationsLifecycleStateEnum
const (
	ListDatabaseRegistrationsLifecycleStateCreating       ListDatabaseRegistrationsLifecycleStateEnum = "CREATING"
	ListDatabaseRegistrationsLifecycleStateUpdating       ListDatabaseRegistrationsLifecycleStateEnum = "UPDATING"
	ListDatabaseRegistrationsLifecycleStateActive         ListDatabaseRegistrationsLifecycleStateEnum = "ACTIVE"
	ListDatabaseRegistrationsLifecycleStateInactive       ListDatabaseRegistrationsLifecycleStateEnum = "INACTIVE"
	ListDatabaseRegistrationsLifecycleStateDeleting       ListDatabaseRegistrationsLifecycleStateEnum = "DELETING"
	ListDatabaseRegistrationsLifecycleStateDeleted        ListDatabaseRegistrationsLifecycleStateEnum = "DELETED"
	ListDatabaseRegistrationsLifecycleStateFailed         ListDatabaseRegistrationsLifecycleStateEnum = "FAILED"
	ListDatabaseRegistrationsLifecycleStateNeedsAttention ListDatabaseRegistrationsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListDatabaseRegistrationsLifecycleStateInProgress     ListDatabaseRegistrationsLifecycleStateEnum = "IN_PROGRESS"
	ListDatabaseRegistrationsLifecycleStateCanceling      ListDatabaseRegistrationsLifecycleStateEnum = "CANCELING"
	ListDatabaseRegistrationsLifecycleStateCanceled       ListDatabaseRegistrationsLifecycleStateEnum = "CANCELED"
	ListDatabaseRegistrationsLifecycleStateSucceeded      ListDatabaseRegistrationsLifecycleStateEnum = "SUCCEEDED"
)

var mappingListDatabaseRegistrationsLifecycleState = map[string]ListDatabaseRegistrationsLifecycleStateEnum{
	"CREATING":        ListDatabaseRegistrationsLifecycleStateCreating,
	"UPDATING":        ListDatabaseRegistrationsLifecycleStateUpdating,
	"ACTIVE":          ListDatabaseRegistrationsLifecycleStateActive,
	"INACTIVE":        ListDatabaseRegistrationsLifecycleStateInactive,
	"DELETING":        ListDatabaseRegistrationsLifecycleStateDeleting,
	"DELETED":         ListDatabaseRegistrationsLifecycleStateDeleted,
	"FAILED":          ListDatabaseRegistrationsLifecycleStateFailed,
	"NEEDS_ATTENTION": ListDatabaseRegistrationsLifecycleStateNeedsAttention,
	"IN_PROGRESS":     ListDatabaseRegistrationsLifecycleStateInProgress,
	"CANCELING":       ListDatabaseRegistrationsLifecycleStateCanceling,
	"CANCELED":        ListDatabaseRegistrationsLifecycleStateCanceled,
	"SUCCEEDED":       ListDatabaseRegistrationsLifecycleStateSucceeded,
}

// GetListDatabaseRegistrationsLifecycleStateEnumValues Enumerates the set of values for ListDatabaseRegistrationsLifecycleStateEnum
func GetListDatabaseRegistrationsLifecycleStateEnumValues() []ListDatabaseRegistrationsLifecycleStateEnum {
	values := make([]ListDatabaseRegistrationsLifecycleStateEnum, 0)
	for _, v := range mappingListDatabaseRegistrationsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListDatabaseRegistrationsSortOrderEnum Enum with underlying type: string
type ListDatabaseRegistrationsSortOrderEnum string

// Set of constants representing the allowable values for ListDatabaseRegistrationsSortOrderEnum
const (
	ListDatabaseRegistrationsSortOrderAsc  ListDatabaseRegistrationsSortOrderEnum = "ASC"
	ListDatabaseRegistrationsSortOrderDesc ListDatabaseRegistrationsSortOrderEnum = "DESC"
)

var mappingListDatabaseRegistrationsSortOrder = map[string]ListDatabaseRegistrationsSortOrderEnum{
	"ASC":  ListDatabaseRegistrationsSortOrderAsc,
	"DESC": ListDatabaseRegistrationsSortOrderDesc,
}

// GetListDatabaseRegistrationsSortOrderEnumValues Enumerates the set of values for ListDatabaseRegistrationsSortOrderEnum
func GetListDatabaseRegistrationsSortOrderEnumValues() []ListDatabaseRegistrationsSortOrderEnum {
	values := make([]ListDatabaseRegistrationsSortOrderEnum, 0)
	for _, v := range mappingListDatabaseRegistrationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListDatabaseRegistrationsSortByEnum Enum with underlying type: string
type ListDatabaseRegistrationsSortByEnum string

// Set of constants representing the allowable values for ListDatabaseRegistrationsSortByEnum
const (
	ListDatabaseRegistrationsSortByTimecreated ListDatabaseRegistrationsSortByEnum = "timeCreated"
	ListDatabaseRegistrationsSortByDisplayname ListDatabaseRegistrationsSortByEnum = "displayName"
)

var mappingListDatabaseRegistrationsSortBy = map[string]ListDatabaseRegistrationsSortByEnum{
	"timeCreated": ListDatabaseRegistrationsSortByTimecreated,
	"displayName": ListDatabaseRegistrationsSortByDisplayname,
}

// GetListDatabaseRegistrationsSortByEnumValues Enumerates the set of values for ListDatabaseRegistrationsSortByEnum
func GetListDatabaseRegistrationsSortByEnumValues() []ListDatabaseRegistrationsSortByEnum {
	values := make([]ListDatabaseRegistrationsSortByEnum, 0)
	for _, v := range mappingListDatabaseRegistrationsSortBy {
		values = append(values, v)
	}
	return values
}
