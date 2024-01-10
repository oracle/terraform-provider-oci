// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListWorkRequestsRequest wrapper for the ListWorkRequests operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudmigrations/ListWorkRequests.go.html to see an example of how to use ListWorkRequestsRequest.
type ListWorkRequestsRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The ID of the asynchronous work request.
	WorkRequestId *string `mandatory:"false" contributesTo:"query" name:"workRequestId"`

	// A filter to return only resources where the resource's lifecycle state matches the given operation status.
	Status ListWorkRequestsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// A filter to return only resources where the resource's lifecycle state matches the given operation type.
	OperationType ListWorkRequestsOperationTypeEnum `mandatory:"false" contributesTo:"query" name:"operationType" omitEmpty:"true"`

	// The ID of the resource affected by the work request.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of the previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListWorkRequestsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. The default order for 'timeAccepted' is descending.
	SortBy ListWorkRequestsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListWorkRequestsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListWorkRequestsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListWorkRequestsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListWorkRequestsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListWorkRequestsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListWorkRequestsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListWorkRequestsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsOperationTypeEnum(string(request.OperationType)); !ok && request.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", request.OperationType, strings.Join(GetListWorkRequestsOperationTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListWorkRequestsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListWorkRequestsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListWorkRequestsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListWorkRequestsResponse wrapper for the ListWorkRequests operation
type ListWorkRequestsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of WorkRequestSummaryCollection instances
	WorkRequestSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListWorkRequestsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListWorkRequestsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListWorkRequestsStatusEnum Enum with underlying type: string
type ListWorkRequestsStatusEnum string

// Set of constants representing the allowable values for ListWorkRequestsStatusEnum
const (
	ListWorkRequestsStatusAccepted       ListWorkRequestsStatusEnum = "ACCEPTED"
	ListWorkRequestsStatusInProgress     ListWorkRequestsStatusEnum = "IN_PROGRESS"
	ListWorkRequestsStatusWaiting        ListWorkRequestsStatusEnum = "WAITING"
	ListWorkRequestsStatusFailed         ListWorkRequestsStatusEnum = "FAILED"
	ListWorkRequestsStatusSucceeded      ListWorkRequestsStatusEnum = "SUCCEEDED"
	ListWorkRequestsStatusCanceling      ListWorkRequestsStatusEnum = "CANCELING"
	ListWorkRequestsStatusCanceled       ListWorkRequestsStatusEnum = "CANCELED"
	ListWorkRequestsStatusNeedsAttention ListWorkRequestsStatusEnum = "NEEDS_ATTENTION"
)

var mappingListWorkRequestsStatusEnum = map[string]ListWorkRequestsStatusEnum{
	"ACCEPTED":        ListWorkRequestsStatusAccepted,
	"IN_PROGRESS":     ListWorkRequestsStatusInProgress,
	"WAITING":         ListWorkRequestsStatusWaiting,
	"FAILED":          ListWorkRequestsStatusFailed,
	"SUCCEEDED":       ListWorkRequestsStatusSucceeded,
	"CANCELING":       ListWorkRequestsStatusCanceling,
	"CANCELED":        ListWorkRequestsStatusCanceled,
	"NEEDS_ATTENTION": ListWorkRequestsStatusNeedsAttention,
}

var mappingListWorkRequestsStatusEnumLowerCase = map[string]ListWorkRequestsStatusEnum{
	"accepted":        ListWorkRequestsStatusAccepted,
	"in_progress":     ListWorkRequestsStatusInProgress,
	"waiting":         ListWorkRequestsStatusWaiting,
	"failed":          ListWorkRequestsStatusFailed,
	"succeeded":       ListWorkRequestsStatusSucceeded,
	"canceling":       ListWorkRequestsStatusCanceling,
	"canceled":        ListWorkRequestsStatusCanceled,
	"needs_attention": ListWorkRequestsStatusNeedsAttention,
}

// GetListWorkRequestsStatusEnumValues Enumerates the set of values for ListWorkRequestsStatusEnum
func GetListWorkRequestsStatusEnumValues() []ListWorkRequestsStatusEnum {
	values := make([]ListWorkRequestsStatusEnum, 0)
	for _, v := range mappingListWorkRequestsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsStatusEnumStringValues Enumerates the set of values in String for ListWorkRequestsStatusEnum
func GetListWorkRequestsStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingListWorkRequestsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsStatusEnum(val string) (ListWorkRequestsStatusEnum, bool) {
	enum, ok := mappingListWorkRequestsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsOperationTypeEnum Enum with underlying type: string
type ListWorkRequestsOperationTypeEnum string

// Set of constants representing the allowable values for ListWorkRequestsOperationTypeEnum
const (
	ListWorkRequestsOperationTypeCreateMigration           ListWorkRequestsOperationTypeEnum = "CREATE_MIGRATION"
	ListWorkRequestsOperationTypeUpdateMigration           ListWorkRequestsOperationTypeEnum = "UPDATE_MIGRATION"
	ListWorkRequestsOperationTypeRefreshMigration          ListWorkRequestsOperationTypeEnum = "REFRESH_MIGRATION"
	ListWorkRequestsOperationTypeDeleteMigration           ListWorkRequestsOperationTypeEnum = "DELETE_MIGRATION"
	ListWorkRequestsOperationTypeMoveMigration             ListWorkRequestsOperationTypeEnum = "MOVE_MIGRATION"
	ListWorkRequestsOperationTypeStartAssetReplication     ListWorkRequestsOperationTypeEnum = "START_ASSET_REPLICATION"
	ListWorkRequestsOperationTypeStartMigrationReplication ListWorkRequestsOperationTypeEnum = "START_MIGRATION_REPLICATION"
	ListWorkRequestsOperationTypeCreateReplicationSchedule ListWorkRequestsOperationTypeEnum = "CREATE_REPLICATION_SCHEDULE"
	ListWorkRequestsOperationTypeUpdateReplicationSchedule ListWorkRequestsOperationTypeEnum = "UPDATE_REPLICATION_SCHEDULE"
	ListWorkRequestsOperationTypeDeleteReplicationSchedule ListWorkRequestsOperationTypeEnum = "DELETE_REPLICATION_SCHEDULE"
	ListWorkRequestsOperationTypeMoveReplicationSchedule   ListWorkRequestsOperationTypeEnum = "MOVE_REPLICATION_SCHEDULE"
	ListWorkRequestsOperationTypeCreateMigrationPlan       ListWorkRequestsOperationTypeEnum = "CREATE_MIGRATION_PLAN"
	ListWorkRequestsOperationTypeUpdateMigrationPlan       ListWorkRequestsOperationTypeEnum = "UPDATE_MIGRATION_PLAN"
	ListWorkRequestsOperationTypeDeleteMigrationPlan       ListWorkRequestsOperationTypeEnum = "DELETE_MIGRATION_PLAN"
	ListWorkRequestsOperationTypeMoveMigrationPlan         ListWorkRequestsOperationTypeEnum = "MOVE_MIGRATION_PLAN"
	ListWorkRequestsOperationTypeRefreshMigrationPlan      ListWorkRequestsOperationTypeEnum = "REFRESH_MIGRATION_PLAN"
	ListWorkRequestsOperationTypeExecuteMigrationPlan      ListWorkRequestsOperationTypeEnum = "EXECUTE_MIGRATION_PLAN"
	ListWorkRequestsOperationTypeRefreshMigrationAsset     ListWorkRequestsOperationTypeEnum = "REFRESH_MIGRATION_ASSET"
	ListWorkRequestsOperationTypeCreateMigrationAsset      ListWorkRequestsOperationTypeEnum = "CREATE_MIGRATION_ASSET"
	ListWorkRequestsOperationTypeDeleteMigrationAsset      ListWorkRequestsOperationTypeEnum = "DELETE_MIGRATION_ASSET"
	ListWorkRequestsOperationTypeCreateTargetAsset         ListWorkRequestsOperationTypeEnum = "CREATE_TARGET_ASSET"
	ListWorkRequestsOperationTypeUpdateTargetAsset         ListWorkRequestsOperationTypeEnum = "UPDATE_TARGET_ASSET"
	ListWorkRequestsOperationTypeDeleteTargetAsset         ListWorkRequestsOperationTypeEnum = "DELETE_TARGET_ASSET"
)

var mappingListWorkRequestsOperationTypeEnum = map[string]ListWorkRequestsOperationTypeEnum{
	"CREATE_MIGRATION":            ListWorkRequestsOperationTypeCreateMigration,
	"UPDATE_MIGRATION":            ListWorkRequestsOperationTypeUpdateMigration,
	"REFRESH_MIGRATION":           ListWorkRequestsOperationTypeRefreshMigration,
	"DELETE_MIGRATION":            ListWorkRequestsOperationTypeDeleteMigration,
	"MOVE_MIGRATION":              ListWorkRequestsOperationTypeMoveMigration,
	"START_ASSET_REPLICATION":     ListWorkRequestsOperationTypeStartAssetReplication,
	"START_MIGRATION_REPLICATION": ListWorkRequestsOperationTypeStartMigrationReplication,
	"CREATE_REPLICATION_SCHEDULE": ListWorkRequestsOperationTypeCreateReplicationSchedule,
	"UPDATE_REPLICATION_SCHEDULE": ListWorkRequestsOperationTypeUpdateReplicationSchedule,
	"DELETE_REPLICATION_SCHEDULE": ListWorkRequestsOperationTypeDeleteReplicationSchedule,
	"MOVE_REPLICATION_SCHEDULE":   ListWorkRequestsOperationTypeMoveReplicationSchedule,
	"CREATE_MIGRATION_PLAN":       ListWorkRequestsOperationTypeCreateMigrationPlan,
	"UPDATE_MIGRATION_PLAN":       ListWorkRequestsOperationTypeUpdateMigrationPlan,
	"DELETE_MIGRATION_PLAN":       ListWorkRequestsOperationTypeDeleteMigrationPlan,
	"MOVE_MIGRATION_PLAN":         ListWorkRequestsOperationTypeMoveMigrationPlan,
	"REFRESH_MIGRATION_PLAN":      ListWorkRequestsOperationTypeRefreshMigrationPlan,
	"EXECUTE_MIGRATION_PLAN":      ListWorkRequestsOperationTypeExecuteMigrationPlan,
	"REFRESH_MIGRATION_ASSET":     ListWorkRequestsOperationTypeRefreshMigrationAsset,
	"CREATE_MIGRATION_ASSET":      ListWorkRequestsOperationTypeCreateMigrationAsset,
	"DELETE_MIGRATION_ASSET":      ListWorkRequestsOperationTypeDeleteMigrationAsset,
	"CREATE_TARGET_ASSET":         ListWorkRequestsOperationTypeCreateTargetAsset,
	"UPDATE_TARGET_ASSET":         ListWorkRequestsOperationTypeUpdateTargetAsset,
	"DELETE_TARGET_ASSET":         ListWorkRequestsOperationTypeDeleteTargetAsset,
}

var mappingListWorkRequestsOperationTypeEnumLowerCase = map[string]ListWorkRequestsOperationTypeEnum{
	"create_migration":            ListWorkRequestsOperationTypeCreateMigration,
	"update_migration":            ListWorkRequestsOperationTypeUpdateMigration,
	"refresh_migration":           ListWorkRequestsOperationTypeRefreshMigration,
	"delete_migration":            ListWorkRequestsOperationTypeDeleteMigration,
	"move_migration":              ListWorkRequestsOperationTypeMoveMigration,
	"start_asset_replication":     ListWorkRequestsOperationTypeStartAssetReplication,
	"start_migration_replication": ListWorkRequestsOperationTypeStartMigrationReplication,
	"create_replication_schedule": ListWorkRequestsOperationTypeCreateReplicationSchedule,
	"update_replication_schedule": ListWorkRequestsOperationTypeUpdateReplicationSchedule,
	"delete_replication_schedule": ListWorkRequestsOperationTypeDeleteReplicationSchedule,
	"move_replication_schedule":   ListWorkRequestsOperationTypeMoveReplicationSchedule,
	"create_migration_plan":       ListWorkRequestsOperationTypeCreateMigrationPlan,
	"update_migration_plan":       ListWorkRequestsOperationTypeUpdateMigrationPlan,
	"delete_migration_plan":       ListWorkRequestsOperationTypeDeleteMigrationPlan,
	"move_migration_plan":         ListWorkRequestsOperationTypeMoveMigrationPlan,
	"refresh_migration_plan":      ListWorkRequestsOperationTypeRefreshMigrationPlan,
	"execute_migration_plan":      ListWorkRequestsOperationTypeExecuteMigrationPlan,
	"refresh_migration_asset":     ListWorkRequestsOperationTypeRefreshMigrationAsset,
	"create_migration_asset":      ListWorkRequestsOperationTypeCreateMigrationAsset,
	"delete_migration_asset":      ListWorkRequestsOperationTypeDeleteMigrationAsset,
	"create_target_asset":         ListWorkRequestsOperationTypeCreateTargetAsset,
	"update_target_asset":         ListWorkRequestsOperationTypeUpdateTargetAsset,
	"delete_target_asset":         ListWorkRequestsOperationTypeDeleteTargetAsset,
}

// GetListWorkRequestsOperationTypeEnumValues Enumerates the set of values for ListWorkRequestsOperationTypeEnum
func GetListWorkRequestsOperationTypeEnumValues() []ListWorkRequestsOperationTypeEnum {
	values := make([]ListWorkRequestsOperationTypeEnum, 0)
	for _, v := range mappingListWorkRequestsOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsOperationTypeEnumStringValues Enumerates the set of values in String for ListWorkRequestsOperationTypeEnum
func GetListWorkRequestsOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_MIGRATION",
		"UPDATE_MIGRATION",
		"REFRESH_MIGRATION",
		"DELETE_MIGRATION",
		"MOVE_MIGRATION",
		"START_ASSET_REPLICATION",
		"START_MIGRATION_REPLICATION",
		"CREATE_REPLICATION_SCHEDULE",
		"UPDATE_REPLICATION_SCHEDULE",
		"DELETE_REPLICATION_SCHEDULE",
		"MOVE_REPLICATION_SCHEDULE",
		"CREATE_MIGRATION_PLAN",
		"UPDATE_MIGRATION_PLAN",
		"DELETE_MIGRATION_PLAN",
		"MOVE_MIGRATION_PLAN",
		"REFRESH_MIGRATION_PLAN",
		"EXECUTE_MIGRATION_PLAN",
		"REFRESH_MIGRATION_ASSET",
		"CREATE_MIGRATION_ASSET",
		"DELETE_MIGRATION_ASSET",
		"CREATE_TARGET_ASSET",
		"UPDATE_TARGET_ASSET",
		"DELETE_TARGET_ASSET",
	}
}

// GetMappingListWorkRequestsOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsOperationTypeEnum(val string) (ListWorkRequestsOperationTypeEnum, bool) {
	enum, ok := mappingListWorkRequestsOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortOrderEnum Enum with underlying type: string
type ListWorkRequestsSortOrderEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortOrderEnum
const (
	ListWorkRequestsSortOrderAsc  ListWorkRequestsSortOrderEnum = "ASC"
	ListWorkRequestsSortOrderDesc ListWorkRequestsSortOrderEnum = "DESC"
)

var mappingListWorkRequestsSortOrderEnum = map[string]ListWorkRequestsSortOrderEnum{
	"ASC":  ListWorkRequestsSortOrderAsc,
	"DESC": ListWorkRequestsSortOrderDesc,
}

var mappingListWorkRequestsSortOrderEnumLowerCase = map[string]ListWorkRequestsSortOrderEnum{
	"asc":  ListWorkRequestsSortOrderAsc,
	"desc": ListWorkRequestsSortOrderDesc,
}

// GetListWorkRequestsSortOrderEnumValues Enumerates the set of values for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumValues() []ListWorkRequestsSortOrderEnum {
	values := make([]ListWorkRequestsSortOrderEnum, 0)
	for _, v := range mappingListWorkRequestsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortOrderEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortOrderEnum
func GetListWorkRequestsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListWorkRequestsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortOrderEnum(val string) (ListWorkRequestsSortOrderEnum, bool) {
	enum, ok := mappingListWorkRequestsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListWorkRequestsSortByEnum Enum with underlying type: string
type ListWorkRequestsSortByEnum string

// Set of constants representing the allowable values for ListWorkRequestsSortByEnum
const (
	ListWorkRequestsSortByTimeaccepted ListWorkRequestsSortByEnum = "timeAccepted"
)

var mappingListWorkRequestsSortByEnum = map[string]ListWorkRequestsSortByEnum{
	"timeAccepted": ListWorkRequestsSortByTimeaccepted,
}

var mappingListWorkRequestsSortByEnumLowerCase = map[string]ListWorkRequestsSortByEnum{
	"timeaccepted": ListWorkRequestsSortByTimeaccepted,
}

// GetListWorkRequestsSortByEnumValues Enumerates the set of values for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumValues() []ListWorkRequestsSortByEnum {
	values := make([]ListWorkRequestsSortByEnum, 0)
	for _, v := range mappingListWorkRequestsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListWorkRequestsSortByEnumStringValues Enumerates the set of values in String for ListWorkRequestsSortByEnum
func GetListWorkRequestsSortByEnumStringValues() []string {
	return []string{
		"timeAccepted",
	}
}

// GetMappingListWorkRequestsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListWorkRequestsSortByEnum(val string) (ListWorkRequestsSortByEnum, bool) {
	enum, ok := mappingListWorkRequestsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
