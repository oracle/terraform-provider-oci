// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListTargetDatabasesRequest wrapper for the ListTargetDatabases operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetDatabases.go.html to see an example of how to use ListTargetDatabasesRequest.
type ListTargetDatabasesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return the target database that matches the specified OCID.
	TargetDatabaseId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return the target databases that matches the current state of the target database.
	LifecycleState ListTargetDatabasesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return target databases that match the database type of the target database.
	DatabaseType ListTargetDatabasesDatabaseTypeEnum `mandatory:"false" contributesTo:"query" name:"databaseType" omitEmpty:"true"`

	// A filter to return target databases that match the infrastructure type of the target database.
	InfrastructureType ListTargetDatabasesInfrastructureTypeEnum `mandatory:"false" contributesTo:"query" name:"infrastructureType" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListTargetDatabasesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListTargetDatabasesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field used for sorting. Only one sorting order (sortOrder) can be specified.
	// The default order for TIMECREATED is descending. The default order for DISPLAYNAME is ascending.
	// The DISPLAYNAME sort order is case sensitive.
	SortBy ListTargetDatabasesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTargetDatabasesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTargetDatabasesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTargetDatabasesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTargetDatabasesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListTargetDatabasesResponse wrapper for the ListTargetDatabases operation
type ListTargetDatabasesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []TargetDatabaseSummary instances
	Items []TargetDatabaseSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListTargetDatabasesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTargetDatabasesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTargetDatabasesLifecycleStateEnum Enum with underlying type: string
type ListTargetDatabasesLifecycleStateEnum string

// Set of constants representing the allowable values for ListTargetDatabasesLifecycleStateEnum
const (
	ListTargetDatabasesLifecycleStateCreating       ListTargetDatabasesLifecycleStateEnum = "CREATING"
	ListTargetDatabasesLifecycleStateUpdating       ListTargetDatabasesLifecycleStateEnum = "UPDATING"
	ListTargetDatabasesLifecycleStateActive         ListTargetDatabasesLifecycleStateEnum = "ACTIVE"
	ListTargetDatabasesLifecycleStateInactive       ListTargetDatabasesLifecycleStateEnum = "INACTIVE"
	ListTargetDatabasesLifecycleStateDeleting       ListTargetDatabasesLifecycleStateEnum = "DELETING"
	ListTargetDatabasesLifecycleStateDeleted        ListTargetDatabasesLifecycleStateEnum = "DELETED"
	ListTargetDatabasesLifecycleStateNeedsAttention ListTargetDatabasesLifecycleStateEnum = "NEEDS_ATTENTION"
	ListTargetDatabasesLifecycleStateFailed         ListTargetDatabasesLifecycleStateEnum = "FAILED"
)

var mappingListTargetDatabasesLifecycleState = map[string]ListTargetDatabasesLifecycleStateEnum{
	"CREATING":        ListTargetDatabasesLifecycleStateCreating,
	"UPDATING":        ListTargetDatabasesLifecycleStateUpdating,
	"ACTIVE":          ListTargetDatabasesLifecycleStateActive,
	"INACTIVE":        ListTargetDatabasesLifecycleStateInactive,
	"DELETING":        ListTargetDatabasesLifecycleStateDeleting,
	"DELETED":         ListTargetDatabasesLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListTargetDatabasesLifecycleStateNeedsAttention,
	"FAILED":          ListTargetDatabasesLifecycleStateFailed,
}

// GetListTargetDatabasesLifecycleStateEnumValues Enumerates the set of values for ListTargetDatabasesLifecycleStateEnum
func GetListTargetDatabasesLifecycleStateEnumValues() []ListTargetDatabasesLifecycleStateEnum {
	values := make([]ListTargetDatabasesLifecycleStateEnum, 0)
	for _, v := range mappingListTargetDatabasesLifecycleState {
		values = append(values, v)
	}
	return values
}

// ListTargetDatabasesDatabaseTypeEnum Enum with underlying type: string
type ListTargetDatabasesDatabaseTypeEnum string

// Set of constants representing the allowable values for ListTargetDatabasesDatabaseTypeEnum
const (
	ListTargetDatabasesDatabaseTypeDatabaseCloudService ListTargetDatabasesDatabaseTypeEnum = "DATABASE_CLOUD_SERVICE"
	ListTargetDatabasesDatabaseTypeAutonomousDatabase   ListTargetDatabasesDatabaseTypeEnum = "AUTONOMOUS_DATABASE"
	ListTargetDatabasesDatabaseTypeInstalledDatabase    ListTargetDatabasesDatabaseTypeEnum = "INSTALLED_DATABASE"
)

var mappingListTargetDatabasesDatabaseType = map[string]ListTargetDatabasesDatabaseTypeEnum{
	"DATABASE_CLOUD_SERVICE": ListTargetDatabasesDatabaseTypeDatabaseCloudService,
	"AUTONOMOUS_DATABASE":    ListTargetDatabasesDatabaseTypeAutonomousDatabase,
	"INSTALLED_DATABASE":     ListTargetDatabasesDatabaseTypeInstalledDatabase,
}

// GetListTargetDatabasesDatabaseTypeEnumValues Enumerates the set of values for ListTargetDatabasesDatabaseTypeEnum
func GetListTargetDatabasesDatabaseTypeEnumValues() []ListTargetDatabasesDatabaseTypeEnum {
	values := make([]ListTargetDatabasesDatabaseTypeEnum, 0)
	for _, v := range mappingListTargetDatabasesDatabaseType {
		values = append(values, v)
	}
	return values
}

// ListTargetDatabasesInfrastructureTypeEnum Enum with underlying type: string
type ListTargetDatabasesInfrastructureTypeEnum string

// Set of constants representing the allowable values for ListTargetDatabasesInfrastructureTypeEnum
const (
	ListTargetDatabasesInfrastructureTypeOracleCloud     ListTargetDatabasesInfrastructureTypeEnum = "ORACLE_CLOUD"
	ListTargetDatabasesInfrastructureTypeCloudAtCustomer ListTargetDatabasesInfrastructureTypeEnum = "CLOUD_AT_CUSTOMER"
	ListTargetDatabasesInfrastructureTypeOnPremises      ListTargetDatabasesInfrastructureTypeEnum = "ON_PREMISES"
	ListTargetDatabasesInfrastructureTypeNonOracleCloud  ListTargetDatabasesInfrastructureTypeEnum = "NON_ORACLE_CLOUD"
)

var mappingListTargetDatabasesInfrastructureType = map[string]ListTargetDatabasesInfrastructureTypeEnum{
	"ORACLE_CLOUD":      ListTargetDatabasesInfrastructureTypeOracleCloud,
	"CLOUD_AT_CUSTOMER": ListTargetDatabasesInfrastructureTypeCloudAtCustomer,
	"ON_PREMISES":       ListTargetDatabasesInfrastructureTypeOnPremises,
	"NON_ORACLE_CLOUD":  ListTargetDatabasesInfrastructureTypeNonOracleCloud,
}

// GetListTargetDatabasesInfrastructureTypeEnumValues Enumerates the set of values for ListTargetDatabasesInfrastructureTypeEnum
func GetListTargetDatabasesInfrastructureTypeEnumValues() []ListTargetDatabasesInfrastructureTypeEnum {
	values := make([]ListTargetDatabasesInfrastructureTypeEnum, 0)
	for _, v := range mappingListTargetDatabasesInfrastructureType {
		values = append(values, v)
	}
	return values
}

// ListTargetDatabasesAccessLevelEnum Enum with underlying type: string
type ListTargetDatabasesAccessLevelEnum string

// Set of constants representing the allowable values for ListTargetDatabasesAccessLevelEnum
const (
	ListTargetDatabasesAccessLevelRestricted ListTargetDatabasesAccessLevelEnum = "RESTRICTED"
	ListTargetDatabasesAccessLevelAccessible ListTargetDatabasesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListTargetDatabasesAccessLevel = map[string]ListTargetDatabasesAccessLevelEnum{
	"RESTRICTED": ListTargetDatabasesAccessLevelRestricted,
	"ACCESSIBLE": ListTargetDatabasesAccessLevelAccessible,
}

// GetListTargetDatabasesAccessLevelEnumValues Enumerates the set of values for ListTargetDatabasesAccessLevelEnum
func GetListTargetDatabasesAccessLevelEnumValues() []ListTargetDatabasesAccessLevelEnum {
	values := make([]ListTargetDatabasesAccessLevelEnum, 0)
	for _, v := range mappingListTargetDatabasesAccessLevel {
		values = append(values, v)
	}
	return values
}

// ListTargetDatabasesSortOrderEnum Enum with underlying type: string
type ListTargetDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetDatabasesSortOrderEnum
const (
	ListTargetDatabasesSortOrderAsc  ListTargetDatabasesSortOrderEnum = "ASC"
	ListTargetDatabasesSortOrderDesc ListTargetDatabasesSortOrderEnum = "DESC"
)

var mappingListTargetDatabasesSortOrder = map[string]ListTargetDatabasesSortOrderEnum{
	"ASC":  ListTargetDatabasesSortOrderAsc,
	"DESC": ListTargetDatabasesSortOrderDesc,
}

// GetListTargetDatabasesSortOrderEnumValues Enumerates the set of values for ListTargetDatabasesSortOrderEnum
func GetListTargetDatabasesSortOrderEnumValues() []ListTargetDatabasesSortOrderEnum {
	values := make([]ListTargetDatabasesSortOrderEnum, 0)
	for _, v := range mappingListTargetDatabasesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListTargetDatabasesSortByEnum Enum with underlying type: string
type ListTargetDatabasesSortByEnum string

// Set of constants representing the allowable values for ListTargetDatabasesSortByEnum
const (
	ListTargetDatabasesSortByTimecreated ListTargetDatabasesSortByEnum = "TIMECREATED"
	ListTargetDatabasesSortByDisplayname ListTargetDatabasesSortByEnum = "DISPLAYNAME"
)

var mappingListTargetDatabasesSortBy = map[string]ListTargetDatabasesSortByEnum{
	"TIMECREATED": ListTargetDatabasesSortByTimecreated,
	"DISPLAYNAME": ListTargetDatabasesSortByDisplayname,
}

// GetListTargetDatabasesSortByEnumValues Enumerates the set of values for ListTargetDatabasesSortByEnum
func GetListTargetDatabasesSortByEnumValues() []ListTargetDatabasesSortByEnum {
	values := make([]ListTargetDatabasesSortByEnum, 0)
	for _, v := range mappingListTargetDatabasesSortBy {
		values = append(values, v)
	}
	return values
}
