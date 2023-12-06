// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTargetDatabasesRequest wrapper for the ListTargetDatabases operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListTargetDatabases.go.html to see an example of how to use ListTargetDatabasesRequest.
type ListTargetDatabasesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// A filter to return the target databases that are associated to the resource id passed in as a parameter value.
	AssociatedResourceId *string `mandatory:"false" contributesTo:"query" name:"associatedResourceId"`

	// A filter to return the target database that matches the specified OCID.
	TargetDatabaseId *string `mandatory:"false" contributesTo:"query" name:"targetDatabaseId"`

	// A filter to return only resources that match the specified display name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only target databases that match the specified lifecycle state.
	LifecycleState ListTargetDatabasesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only target databases that match the specified database type.
	DatabaseType ListTargetDatabasesDatabaseTypeEnum `mandatory:"false" contributesTo:"query" name:"databaseType" omitEmpty:"true"`

	// A filter to return only target databases that match the specified infrastructure type.
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

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTargetDatabasesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListTargetDatabasesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListTargetDatabasesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabasesDatabaseTypeEnum(string(request.DatabaseType)); !ok && request.DatabaseType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DatabaseType: %s. Supported values are: %s.", request.DatabaseType, strings.Join(GetListTargetDatabasesDatabaseTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabasesInfrastructureTypeEnum(string(request.InfrastructureType)); !ok && request.InfrastructureType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InfrastructureType: %s. Supported values are: %s.", request.InfrastructureType, strings.Join(GetListTargetDatabasesInfrastructureTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabasesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListTargetDatabasesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabasesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTargetDatabasesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTargetDatabasesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTargetDatabasesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingListTargetDatabasesLifecycleStateEnum = map[string]ListTargetDatabasesLifecycleStateEnum{
	"CREATING":        ListTargetDatabasesLifecycleStateCreating,
	"UPDATING":        ListTargetDatabasesLifecycleStateUpdating,
	"ACTIVE":          ListTargetDatabasesLifecycleStateActive,
	"INACTIVE":        ListTargetDatabasesLifecycleStateInactive,
	"DELETING":        ListTargetDatabasesLifecycleStateDeleting,
	"DELETED":         ListTargetDatabasesLifecycleStateDeleted,
	"NEEDS_ATTENTION": ListTargetDatabasesLifecycleStateNeedsAttention,
	"FAILED":          ListTargetDatabasesLifecycleStateFailed,
}

var mappingListTargetDatabasesLifecycleStateEnumLowerCase = map[string]ListTargetDatabasesLifecycleStateEnum{
	"creating":        ListTargetDatabasesLifecycleStateCreating,
	"updating":        ListTargetDatabasesLifecycleStateUpdating,
	"active":          ListTargetDatabasesLifecycleStateActive,
	"inactive":        ListTargetDatabasesLifecycleStateInactive,
	"deleting":        ListTargetDatabasesLifecycleStateDeleting,
	"deleted":         ListTargetDatabasesLifecycleStateDeleted,
	"needs_attention": ListTargetDatabasesLifecycleStateNeedsAttention,
	"failed":          ListTargetDatabasesLifecycleStateFailed,
}

// GetListTargetDatabasesLifecycleStateEnumValues Enumerates the set of values for ListTargetDatabasesLifecycleStateEnum
func GetListTargetDatabasesLifecycleStateEnumValues() []ListTargetDatabasesLifecycleStateEnum {
	values := make([]ListTargetDatabasesLifecycleStateEnum, 0)
	for _, v := range mappingListTargetDatabasesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabasesLifecycleStateEnumStringValues Enumerates the set of values in String for ListTargetDatabasesLifecycleStateEnum
func GetListTargetDatabasesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingListTargetDatabasesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabasesLifecycleStateEnum(val string) (ListTargetDatabasesLifecycleStateEnum, bool) {
	enum, ok := mappingListTargetDatabasesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabasesDatabaseTypeEnum Enum with underlying type: string
type ListTargetDatabasesDatabaseTypeEnum string

// Set of constants representing the allowable values for ListTargetDatabasesDatabaseTypeEnum
const (
	ListTargetDatabasesDatabaseTypeDatabaseCloudService ListTargetDatabasesDatabaseTypeEnum = "DATABASE_CLOUD_SERVICE"
	ListTargetDatabasesDatabaseTypeAutonomousDatabase   ListTargetDatabasesDatabaseTypeEnum = "AUTONOMOUS_DATABASE"
	ListTargetDatabasesDatabaseTypeInstalledDatabase    ListTargetDatabasesDatabaseTypeEnum = "INSTALLED_DATABASE"
)

var mappingListTargetDatabasesDatabaseTypeEnum = map[string]ListTargetDatabasesDatabaseTypeEnum{
	"DATABASE_CLOUD_SERVICE": ListTargetDatabasesDatabaseTypeDatabaseCloudService,
	"AUTONOMOUS_DATABASE":    ListTargetDatabasesDatabaseTypeAutonomousDatabase,
	"INSTALLED_DATABASE":     ListTargetDatabasesDatabaseTypeInstalledDatabase,
}

var mappingListTargetDatabasesDatabaseTypeEnumLowerCase = map[string]ListTargetDatabasesDatabaseTypeEnum{
	"database_cloud_service": ListTargetDatabasesDatabaseTypeDatabaseCloudService,
	"autonomous_database":    ListTargetDatabasesDatabaseTypeAutonomousDatabase,
	"installed_database":     ListTargetDatabasesDatabaseTypeInstalledDatabase,
}

// GetListTargetDatabasesDatabaseTypeEnumValues Enumerates the set of values for ListTargetDatabasesDatabaseTypeEnum
func GetListTargetDatabasesDatabaseTypeEnumValues() []ListTargetDatabasesDatabaseTypeEnum {
	values := make([]ListTargetDatabasesDatabaseTypeEnum, 0)
	for _, v := range mappingListTargetDatabasesDatabaseTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabasesDatabaseTypeEnumStringValues Enumerates the set of values in String for ListTargetDatabasesDatabaseTypeEnum
func GetListTargetDatabasesDatabaseTypeEnumStringValues() []string {
	return []string{
		"DATABASE_CLOUD_SERVICE",
		"AUTONOMOUS_DATABASE",
		"INSTALLED_DATABASE",
	}
}

// GetMappingListTargetDatabasesDatabaseTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabasesDatabaseTypeEnum(val string) (ListTargetDatabasesDatabaseTypeEnum, bool) {
	enum, ok := mappingListTargetDatabasesDatabaseTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
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

var mappingListTargetDatabasesInfrastructureTypeEnum = map[string]ListTargetDatabasesInfrastructureTypeEnum{
	"ORACLE_CLOUD":      ListTargetDatabasesInfrastructureTypeOracleCloud,
	"CLOUD_AT_CUSTOMER": ListTargetDatabasesInfrastructureTypeCloudAtCustomer,
	"ON_PREMISES":       ListTargetDatabasesInfrastructureTypeOnPremises,
	"NON_ORACLE_CLOUD":  ListTargetDatabasesInfrastructureTypeNonOracleCloud,
}

var mappingListTargetDatabasesInfrastructureTypeEnumLowerCase = map[string]ListTargetDatabasesInfrastructureTypeEnum{
	"oracle_cloud":      ListTargetDatabasesInfrastructureTypeOracleCloud,
	"cloud_at_customer": ListTargetDatabasesInfrastructureTypeCloudAtCustomer,
	"on_premises":       ListTargetDatabasesInfrastructureTypeOnPremises,
	"non_oracle_cloud":  ListTargetDatabasesInfrastructureTypeNonOracleCloud,
}

// GetListTargetDatabasesInfrastructureTypeEnumValues Enumerates the set of values for ListTargetDatabasesInfrastructureTypeEnum
func GetListTargetDatabasesInfrastructureTypeEnumValues() []ListTargetDatabasesInfrastructureTypeEnum {
	values := make([]ListTargetDatabasesInfrastructureTypeEnum, 0)
	for _, v := range mappingListTargetDatabasesInfrastructureTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabasesInfrastructureTypeEnumStringValues Enumerates the set of values in String for ListTargetDatabasesInfrastructureTypeEnum
func GetListTargetDatabasesInfrastructureTypeEnumStringValues() []string {
	return []string{
		"ORACLE_CLOUD",
		"CLOUD_AT_CUSTOMER",
		"ON_PREMISES",
		"NON_ORACLE_CLOUD",
	}
}

// GetMappingListTargetDatabasesInfrastructureTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabasesInfrastructureTypeEnum(val string) (ListTargetDatabasesInfrastructureTypeEnum, bool) {
	enum, ok := mappingListTargetDatabasesInfrastructureTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabasesAccessLevelEnum Enum with underlying type: string
type ListTargetDatabasesAccessLevelEnum string

// Set of constants representing the allowable values for ListTargetDatabasesAccessLevelEnum
const (
	ListTargetDatabasesAccessLevelRestricted ListTargetDatabasesAccessLevelEnum = "RESTRICTED"
	ListTargetDatabasesAccessLevelAccessible ListTargetDatabasesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListTargetDatabasesAccessLevelEnum = map[string]ListTargetDatabasesAccessLevelEnum{
	"RESTRICTED": ListTargetDatabasesAccessLevelRestricted,
	"ACCESSIBLE": ListTargetDatabasesAccessLevelAccessible,
}

var mappingListTargetDatabasesAccessLevelEnumLowerCase = map[string]ListTargetDatabasesAccessLevelEnum{
	"restricted": ListTargetDatabasesAccessLevelRestricted,
	"accessible": ListTargetDatabasesAccessLevelAccessible,
}

// GetListTargetDatabasesAccessLevelEnumValues Enumerates the set of values for ListTargetDatabasesAccessLevelEnum
func GetListTargetDatabasesAccessLevelEnumValues() []ListTargetDatabasesAccessLevelEnum {
	values := make([]ListTargetDatabasesAccessLevelEnum, 0)
	for _, v := range mappingListTargetDatabasesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabasesAccessLevelEnumStringValues Enumerates the set of values in String for ListTargetDatabasesAccessLevelEnum
func GetListTargetDatabasesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListTargetDatabasesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabasesAccessLevelEnum(val string) (ListTargetDatabasesAccessLevelEnum, bool) {
	enum, ok := mappingListTargetDatabasesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabasesSortOrderEnum Enum with underlying type: string
type ListTargetDatabasesSortOrderEnum string

// Set of constants representing the allowable values for ListTargetDatabasesSortOrderEnum
const (
	ListTargetDatabasesSortOrderAsc  ListTargetDatabasesSortOrderEnum = "ASC"
	ListTargetDatabasesSortOrderDesc ListTargetDatabasesSortOrderEnum = "DESC"
)

var mappingListTargetDatabasesSortOrderEnum = map[string]ListTargetDatabasesSortOrderEnum{
	"ASC":  ListTargetDatabasesSortOrderAsc,
	"DESC": ListTargetDatabasesSortOrderDesc,
}

var mappingListTargetDatabasesSortOrderEnumLowerCase = map[string]ListTargetDatabasesSortOrderEnum{
	"asc":  ListTargetDatabasesSortOrderAsc,
	"desc": ListTargetDatabasesSortOrderDesc,
}

// GetListTargetDatabasesSortOrderEnumValues Enumerates the set of values for ListTargetDatabasesSortOrderEnum
func GetListTargetDatabasesSortOrderEnumValues() []ListTargetDatabasesSortOrderEnum {
	values := make([]ListTargetDatabasesSortOrderEnum, 0)
	for _, v := range mappingListTargetDatabasesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabasesSortOrderEnumStringValues Enumerates the set of values in String for ListTargetDatabasesSortOrderEnum
func GetListTargetDatabasesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTargetDatabasesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabasesSortOrderEnum(val string) (ListTargetDatabasesSortOrderEnum, bool) {
	enum, ok := mappingListTargetDatabasesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTargetDatabasesSortByEnum Enum with underlying type: string
type ListTargetDatabasesSortByEnum string

// Set of constants representing the allowable values for ListTargetDatabasesSortByEnum
const (
	ListTargetDatabasesSortByTimecreated ListTargetDatabasesSortByEnum = "TIMECREATED"
	ListTargetDatabasesSortByDisplayname ListTargetDatabasesSortByEnum = "DISPLAYNAME"
)

var mappingListTargetDatabasesSortByEnum = map[string]ListTargetDatabasesSortByEnum{
	"TIMECREATED": ListTargetDatabasesSortByTimecreated,
	"DISPLAYNAME": ListTargetDatabasesSortByDisplayname,
}

var mappingListTargetDatabasesSortByEnumLowerCase = map[string]ListTargetDatabasesSortByEnum{
	"timecreated": ListTargetDatabasesSortByTimecreated,
	"displayname": ListTargetDatabasesSortByDisplayname,
}

// GetListTargetDatabasesSortByEnumValues Enumerates the set of values for ListTargetDatabasesSortByEnum
func GetListTargetDatabasesSortByEnumValues() []ListTargetDatabasesSortByEnum {
	values := make([]ListTargetDatabasesSortByEnum, 0)
	for _, v := range mappingListTargetDatabasesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTargetDatabasesSortByEnumStringValues Enumerates the set of values in String for ListTargetDatabasesSortByEnum
func GetListTargetDatabasesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListTargetDatabasesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTargetDatabasesSortByEnum(val string) (ListTargetDatabasesSortByEnum, bool) {
	enum, ok := mappingListTargetDatabasesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
