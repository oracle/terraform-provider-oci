// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListMigrationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListMigrationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListMigrationsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListMigrationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListMigrationsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListMigrationsLifecycleDetailsEnum(string(request.LifecycleDetails)); !ok && request.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", request.LifecycleDetails, strings.Join(GetListMigrationsLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListMigrationsSortByEnum = map[string]ListMigrationsSortByEnum{
	"timeCreated": ListMigrationsSortByTimecreated,
	"displayName": ListMigrationsSortByDisplayname,
}

// GetListMigrationsSortByEnumValues Enumerates the set of values for ListMigrationsSortByEnum
func GetListMigrationsSortByEnumValues() []ListMigrationsSortByEnum {
	values := make([]ListMigrationsSortByEnum, 0)
	for _, v := range mappingListMigrationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationsSortByEnumStringValues Enumerates the set of values in String for ListMigrationsSortByEnum
func GetListMigrationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListMigrationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationsSortByEnum(val string) (ListMigrationsSortByEnum, bool) {
	mappingListMigrationsSortByEnumIgnoreCase := make(map[string]ListMigrationsSortByEnum)
	for k, v := range mappingListMigrationsSortByEnum {
		mappingListMigrationsSortByEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMigrationsSortByEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}

// ListMigrationsSortOrderEnum Enum with underlying type: string
type ListMigrationsSortOrderEnum string

// Set of constants representing the allowable values for ListMigrationsSortOrderEnum
const (
	ListMigrationsSortOrderAsc  ListMigrationsSortOrderEnum = "ASC"
	ListMigrationsSortOrderDesc ListMigrationsSortOrderEnum = "DESC"
)

var mappingListMigrationsSortOrderEnum = map[string]ListMigrationsSortOrderEnum{
	"ASC":  ListMigrationsSortOrderAsc,
	"DESC": ListMigrationsSortOrderDesc,
}

// GetListMigrationsSortOrderEnumValues Enumerates the set of values for ListMigrationsSortOrderEnum
func GetListMigrationsSortOrderEnumValues() []ListMigrationsSortOrderEnum {
	values := make([]ListMigrationsSortOrderEnum, 0)
	for _, v := range mappingListMigrationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationsSortOrderEnumStringValues Enumerates the set of values in String for ListMigrationsSortOrderEnum
func GetListMigrationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListMigrationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationsSortOrderEnum(val string) (ListMigrationsSortOrderEnum, bool) {
	mappingListMigrationsSortOrderEnumIgnoreCase := make(map[string]ListMigrationsSortOrderEnum)
	for k, v := range mappingListMigrationsSortOrderEnum {
		mappingListMigrationsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMigrationsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListMigrationsLifecycleStateEnum = map[string]ListMigrationsLifecycleStateEnum{
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
	for _, v := range mappingListMigrationsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationsLifecycleStateEnumStringValues Enumerates the set of values in String for ListMigrationsLifecycleStateEnum
func GetListMigrationsLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"IN_PROGRESS",
		"ACCEPTED",
		"SUCCEEDED",
		"CANCELED",
		"WAITING",
		"NEEDS_ATTENTION",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListMigrationsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationsLifecycleStateEnum(val string) (ListMigrationsLifecycleStateEnum, bool) {
	mappingListMigrationsLifecycleStateEnumIgnoreCase := make(map[string]ListMigrationsLifecycleStateEnum)
	for k, v := range mappingListMigrationsLifecycleStateEnum {
		mappingListMigrationsLifecycleStateEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMigrationsLifecycleStateEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListMigrationsLifecycleDetailsEnum = map[string]ListMigrationsLifecycleDetailsEnum{
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
	for _, v := range mappingListMigrationsLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetListMigrationsLifecycleDetailsEnumStringValues Enumerates the set of values in String for ListMigrationsLifecycleDetailsEnum
func GetListMigrationsLifecycleDetailsEnumStringValues() []string {
	return []string{
		"READY",
		"ABORTING",
		"VALIDATING",
		"VALIDATED",
		"WAITING",
		"MIGRATING",
		"DONE",
	}
}

// GetMappingListMigrationsLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListMigrationsLifecycleDetailsEnum(val string) (ListMigrationsLifecycleDetailsEnum, bool) {
	mappingListMigrationsLifecycleDetailsEnumIgnoreCase := make(map[string]ListMigrationsLifecycleDetailsEnum)
	for k, v := range mappingListMigrationsLifecycleDetailsEnum {
		mappingListMigrationsLifecycleDetailsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListMigrationsLifecycleDetailsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
