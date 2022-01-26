// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListMigrationsRequest wrapper for the ListMigrations operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListMigrations.go.html to see an example of how to use ListMigrationsRequest.
type ListMigrationsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListMigrationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListMigrationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The lifecycle state of the Migration.
	LifecycleState ListMigrationsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The lifecycle detailed status of the Migration.
	LifecycleDetails ListMigrationsLifecycleDetailsEnum `mandatory:"false" contributesTo:"query" name:"lifecycleDetails" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListMigrationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListMigrationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListMigrationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListMigrationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListMigrationsResponse wrapper for the ListMigrations operation
type ListMigrationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of MigrationCollection instances
	MigrationCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListMigrationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListMigrationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListMigrationsSortByEnum Enum with underlying type: string
type ListMigrationsSortByEnum string

// Set of constants representing the allowable values for ListMigrationsSortByEnum
const (
	ListMigrationsSortByTimecreated ListMigrationsSortByEnum = "timeCreated"
	ListMigrationsSortByDisplayname ListMigrationsSortByEnum = "displayName"
)

var mappingListMigrationsSortBy = map[string]ListMigrationsSortByEnum{
	"timeCreated": ListMigrationsSortByTimecreated,
	"displayName": ListMigrationsSortByDisplayname,
}

// GetListMigrationsSortByEnumValues Enumerates the set of values for ListMigrationsSortByEnum
func GetListMigrationsSortByEnumValues() []ListMigrationsSortByEnum {
	values := make([]ListMigrationsSortByEnum, 0)
	for _, v := range mappingListMigrationsSortBy {
		values = append(values, v)
	}
	return values
}

// ListMigrationsSortOrderEnum Enum with underlying type: string
type ListMigrationsSortOrderEnum string

// Set of constants representing the allowable values for ListMigrationsSortOrderEnum
const (
	ListMigrationsSortOrderAsc  ListMigrationsSortOrderEnum = "ASC"
	ListMigrationsSortOrderDesc ListMigrationsSortOrderEnum = "DESC"
)

var mappingListMigrationsSortOrder = map[string]ListMigrationsSortOrderEnum{
	"ASC":  ListMigrationsSortOrderAsc,
	"DESC": ListMigrationsSortOrderDesc,
}

// GetListMigrationsSortOrderEnumValues Enumerates the set of values for ListMigrationsSortOrderEnum
func GetListMigrationsSortOrderEnumValues() []ListMigrationsSortOrderEnum {
	values := make([]ListMigrationsSortOrderEnum, 0)
	for _, v := range mappingListMigrationsSortOrder {
		values = append(values, v)
	}
	return values
}

// ListMigrationsLifecycleStateEnum Enum with underlying type: string
type ListMigrationsLifecycleStateEnum string

// Set of constants representing the allowable values for ListMigrationsLifecycleStateEnum
const (
	ListMigrationsLifecycleStateCreating       ListMigrationsLifecycleStateEnum = "CREATING"
	ListMigrationsLifecycleStateUpdating       ListMigrationsLifecycleStateEnum = "UPDATING"
	ListMigrationsLifecycleStateActive         ListMigrationsLifecycleStateEnum = "ACTIVE"
	ListMigrationsLifecycleStateInProgress     ListMigrationsLifecycleStateEnum = "IN_PROGRESS"
	ListMigrationsLifecycleStateAccepted       ListMigrationsLifecycleStateEnum = "ACCEPTED"
	ListMigrationsLifecycleStateSucceeded      ListMigrationsLifecycleStateEnum = "SUCCEEDED"
	ListMigrationsLifecycleStateCanceled       ListMigrationsLifecycleStateEnum = "CANCELED"
	ListMigrationsLifecycleStateWaiting        ListMigrationsLifecycleStateEnum = "WAITING"
	ListMigrationsLifecycleStateNeedsAttention ListMigrationsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListMigrationsLifecycleStateInactive       ListMigrationsLifecycleStateEnum = "INACTIVE"
	ListMigrationsLifecycleStateDeleting       ListMigrationsLifecycleStateEnum = "DELETING"
	ListMigrationsLifecycleStateDeleted        ListMigrationsLifecycleStateEnum = "DELETED"
	ListMigrationsLifecycleStateFailed         ListMigrationsLifecycleStateEnum = "FAILED"
)

var mappingListMigrationsLifecycleState = map[string]ListMigrationsLifecycleStateEnum{
	"CREATING":        ListMigrationsLifecycleStateCreating,
	"UPDATING":        ListMigrationsLifecycleStateUpdating,
	"ACTIVE":          ListMigrationsLifecycleStateActive,
	"IN_PROGRESS":     ListMigrationsLifecycleStateInProgress,
	"ACCEPTED":        ListMigrationsLifecycleStateAccepted,
	"SUCCEEDED":       ListMigrationsLifecycleStateSucceeded,
	"CANCELED":        ListMigrationsLifecycleStateCanceled,
	"WAITING":         ListMigrationsLifecycleStateWaiting,
	"NEEDS_ATTENTION": ListMigrationsLifecycleStateNeedsAttention,
	"INACTIVE":        ListMigrationsLifecycleStateInactive,
	"DELETING":        ListMigrationsLifecycleStateDeleting,
	"DELETED":         ListMigrationsLifecycleStateDeleted,
	"FAILED":          ListMigrationsLifecycleStateFailed,
}

// GetListMigrationsLifecycleStateEnumValues Enumerates the set of values for ListMigrationsLifecycleStateEnum
func GetListMigrationsLifecycleStateEnumValues() []ListMigrationsLifecycleStateEnum {
	values := make([]ListMigrationsLifecycleStateEnum, 0)
	for _, v := range mappingListMigrationsLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListMigrationsLifecycleDetailsEnum Enum with underlying type: string
type ListMigrationsLifecycleDetailsEnum string

// Set of constants representing the allowable values for ListMigrationsLifecycleDetailsEnum
const (
	ListMigrationsLifecycleDetailsReady      ListMigrationsLifecycleDetailsEnum = "READY"
	ListMigrationsLifecycleDetailsAborting   ListMigrationsLifecycleDetailsEnum = "ABORTING"
	ListMigrationsLifecycleDetailsValidating ListMigrationsLifecycleDetailsEnum = "VALIDATING"
	ListMigrationsLifecycleDetailsValidated  ListMigrationsLifecycleDetailsEnum = "VALIDATED"
	ListMigrationsLifecycleDetailsWaiting    ListMigrationsLifecycleDetailsEnum = "WAITING"
	ListMigrationsLifecycleDetailsMigrating  ListMigrationsLifecycleDetailsEnum = "MIGRATING"
	ListMigrationsLifecycleDetailsDone       ListMigrationsLifecycleDetailsEnum = "DONE"
)

var mappingListMigrationsLifecycleDetails = map[string]ListMigrationsLifecycleDetailsEnum{
	"READY":      ListMigrationsLifecycleDetailsReady,
	"ABORTING":   ListMigrationsLifecycleDetailsAborting,
	"VALIDATING": ListMigrationsLifecycleDetailsValidating,
	"VALIDATED":  ListMigrationsLifecycleDetailsValidated,
	"WAITING":    ListMigrationsLifecycleDetailsWaiting,
	"MIGRATING":  ListMigrationsLifecycleDetailsMigrating,
	"DONE":       ListMigrationsLifecycleDetailsDone,
}

// GetListMigrationsLifecycleDetailsEnumValues Enumerates the set of values for ListMigrationsLifecycleDetailsEnum
func GetListMigrationsLifecycleDetailsEnumValues() []ListMigrationsLifecycleDetailsEnum {
	values := make([]ListMigrationsLifecycleDetailsEnum, 0)
	for _, v := range mappingListMigrationsLifecycleDetails {
		values = append(values, v)
	}
	return values
}
