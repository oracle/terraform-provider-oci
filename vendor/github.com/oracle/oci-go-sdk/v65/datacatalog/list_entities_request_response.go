// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListEntitiesRequest wrapper for the ListEntities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListEntities.go.html to see an example of how to use ListEntitiesRequest.
type ListEntitiesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// A filter to return only resources that match the entire display name given. The match is not case sensitive.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the entire business name given. The match is not case sensitive.
	BusinessName *string `mandatory:"false" contributesTo:"query" name:"businessName"`

	// A filter to return only resources that match display name or business name pattern given. The match is not case sensitive.
	// For Example : /folders?displayOrBusinessNameContains=Cu.*
	// The above would match all folders with display name or business name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayOrBusinessNameContains *string `mandatory:"false" contributesTo:"query" name:"displayOrBusinessNameContains"`

	// The key of the object type.
	TypeKey *string `mandatory:"false" contributesTo:"query" name:"typeKey"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListEntitiesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Unique external identifier of this resource in the external source system.
	ExternalKey *string `mandatory:"false" contributesTo:"query" name:"externalKey"`

	// Unique pattern key.
	PatternKey *string `mandatory:"false" contributesTo:"query" name:"patternKey"`

	// Last modified timestamp of this object in the external system.
	TimeExternal *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeExternal"`

	// Time that the resource's status was last updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeStatusUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeStatusUpdated"`

	// Identifies if the object is a physical object (materialized) or virtual/logical object defined on other objects.
	IsLogical *bool `mandatory:"false" contributesTo:"query" name:"isLogical"`

	// Identifies if an object is a sub object (partition) of a physical or materialized parent object.
	IsPartition *bool `mandatory:"false" contributesTo:"query" name:"isPartition"`

	// Key of the associated folder.
	FolderKey *string `mandatory:"false" contributesTo:"query" name:"folderKey"`

	// Full path of the resource for resources that support paths.
	Path *string `mandatory:"false" contributesTo:"query" name:"path"`

	// Harvest status of the harvestable resource as updated by the harvest process.
	HarvestStatus ListEntitiesHarvestStatusEnum `mandatory:"false" contributesTo:"query" name:"harvestStatus" omitEmpty:"true"`

	// Key of the last harvest process to update this resource.
	LastJobKey *string `mandatory:"false" contributesTo:"query" name:"lastJobKey"`

	// Specifies the fields to return in an entity summary response.
	Fields []ListEntitiesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The field to sort by. Only one sort order may be provided. DISPLAYORBUSINESSNAME considers businessName of a given object if set, else its displayName is used.
	// Default sort order for TIMECREATED is descending and default sort order for DISPLAYNAME and DISPLAYORBUSINESSNAME is ascending. If no order is specified, TIMECREATED is the default.
	SortBy ListEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only process entities.
	IsProcess *bool `mandatory:"false" contributesTo:"query" name:"isProcess"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListEntitiesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListEntitiesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEntitiesHarvestStatusEnum(string(request.HarvestStatus)); !ok && request.HarvestStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HarvestStatus: %s. Supported values are: %s.", request.HarvestStatus, strings.Join(GetListEntitiesHarvestStatusEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListEntitiesFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListEntitiesFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListEntitiesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListEntitiesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListEntitiesResponse wrapper for the ListEntities operation
type ListEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EntityCollection instances
	EntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListEntitiesLifecycleStateEnum Enum with underlying type: string
type ListEntitiesLifecycleStateEnum string

// Set of constants representing the allowable values for ListEntitiesLifecycleStateEnum
const (
	ListEntitiesLifecycleStateCreating ListEntitiesLifecycleStateEnum = "CREATING"
	ListEntitiesLifecycleStateActive   ListEntitiesLifecycleStateEnum = "ACTIVE"
	ListEntitiesLifecycleStateInactive ListEntitiesLifecycleStateEnum = "INACTIVE"
	ListEntitiesLifecycleStateUpdating ListEntitiesLifecycleStateEnum = "UPDATING"
	ListEntitiesLifecycleStateDeleting ListEntitiesLifecycleStateEnum = "DELETING"
	ListEntitiesLifecycleStateDeleted  ListEntitiesLifecycleStateEnum = "DELETED"
	ListEntitiesLifecycleStateFailed   ListEntitiesLifecycleStateEnum = "FAILED"
	ListEntitiesLifecycleStateMoving   ListEntitiesLifecycleStateEnum = "MOVING"
)

var mappingListEntitiesLifecycleStateEnum = map[string]ListEntitiesLifecycleStateEnum{
	"CREATING": ListEntitiesLifecycleStateCreating,
	"ACTIVE":   ListEntitiesLifecycleStateActive,
	"INACTIVE": ListEntitiesLifecycleStateInactive,
	"UPDATING": ListEntitiesLifecycleStateUpdating,
	"DELETING": ListEntitiesLifecycleStateDeleting,
	"DELETED":  ListEntitiesLifecycleStateDeleted,
	"FAILED":   ListEntitiesLifecycleStateFailed,
	"MOVING":   ListEntitiesLifecycleStateMoving,
}

var mappingListEntitiesLifecycleStateEnumLowerCase = map[string]ListEntitiesLifecycleStateEnum{
	"creating": ListEntitiesLifecycleStateCreating,
	"active":   ListEntitiesLifecycleStateActive,
	"inactive": ListEntitiesLifecycleStateInactive,
	"updating": ListEntitiesLifecycleStateUpdating,
	"deleting": ListEntitiesLifecycleStateDeleting,
	"deleted":  ListEntitiesLifecycleStateDeleted,
	"failed":   ListEntitiesLifecycleStateFailed,
	"moving":   ListEntitiesLifecycleStateMoving,
}

// GetListEntitiesLifecycleStateEnumValues Enumerates the set of values for ListEntitiesLifecycleStateEnum
func GetListEntitiesLifecycleStateEnumValues() []ListEntitiesLifecycleStateEnum {
	values := make([]ListEntitiesLifecycleStateEnum, 0)
	for _, v := range mappingListEntitiesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitiesLifecycleStateEnumStringValues Enumerates the set of values in String for ListEntitiesLifecycleStateEnum
func GetListEntitiesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"MOVING",
	}
}

// GetMappingListEntitiesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitiesLifecycleStateEnum(val string) (ListEntitiesLifecycleStateEnum, bool) {
	enum, ok := mappingListEntitiesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntitiesHarvestStatusEnum Enum with underlying type: string
type ListEntitiesHarvestStatusEnum string

// Set of constants representing the allowable values for ListEntitiesHarvestStatusEnum
const (
	ListEntitiesHarvestStatusComplete   ListEntitiesHarvestStatusEnum = "COMPLETE"
	ListEntitiesHarvestStatusError      ListEntitiesHarvestStatusEnum = "ERROR"
	ListEntitiesHarvestStatusInProgress ListEntitiesHarvestStatusEnum = "IN_PROGRESS"
	ListEntitiesHarvestStatusDeferred   ListEntitiesHarvestStatusEnum = "DEFERRED"
)

var mappingListEntitiesHarvestStatusEnum = map[string]ListEntitiesHarvestStatusEnum{
	"COMPLETE":    ListEntitiesHarvestStatusComplete,
	"ERROR":       ListEntitiesHarvestStatusError,
	"IN_PROGRESS": ListEntitiesHarvestStatusInProgress,
	"DEFERRED":    ListEntitiesHarvestStatusDeferred,
}

var mappingListEntitiesHarvestStatusEnumLowerCase = map[string]ListEntitiesHarvestStatusEnum{
	"complete":    ListEntitiesHarvestStatusComplete,
	"error":       ListEntitiesHarvestStatusError,
	"in_progress": ListEntitiesHarvestStatusInProgress,
	"deferred":    ListEntitiesHarvestStatusDeferred,
}

// GetListEntitiesHarvestStatusEnumValues Enumerates the set of values for ListEntitiesHarvestStatusEnum
func GetListEntitiesHarvestStatusEnumValues() []ListEntitiesHarvestStatusEnum {
	values := make([]ListEntitiesHarvestStatusEnum, 0)
	for _, v := range mappingListEntitiesHarvestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitiesHarvestStatusEnumStringValues Enumerates the set of values in String for ListEntitiesHarvestStatusEnum
func GetListEntitiesHarvestStatusEnumStringValues() []string {
	return []string{
		"COMPLETE",
		"ERROR",
		"IN_PROGRESS",
		"DEFERRED",
	}
}

// GetMappingListEntitiesHarvestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitiesHarvestStatusEnum(val string) (ListEntitiesHarvestStatusEnum, bool) {
	enum, ok := mappingListEntitiesHarvestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntitiesFieldsEnum Enum with underlying type: string
type ListEntitiesFieldsEnum string

// Set of constants representing the allowable values for ListEntitiesFieldsEnum
const (
	ListEntitiesFieldsKey            ListEntitiesFieldsEnum = "key"
	ListEntitiesFieldsDisplayname    ListEntitiesFieldsEnum = "displayName"
	ListEntitiesFieldsDescription    ListEntitiesFieldsEnum = "description"
	ListEntitiesFieldsDataassetkey   ListEntitiesFieldsEnum = "dataAssetKey"
	ListEntitiesFieldsTimecreated    ListEntitiesFieldsEnum = "timeCreated"
	ListEntitiesFieldsTimeupdated    ListEntitiesFieldsEnum = "timeUpdated"
	ListEntitiesFieldsUpdatedbyid    ListEntitiesFieldsEnum = "updatedById"
	ListEntitiesFieldsLifecyclestate ListEntitiesFieldsEnum = "lifecycleState"
	ListEntitiesFieldsFolderkey      ListEntitiesFieldsEnum = "folderKey"
	ListEntitiesFieldsFoldername     ListEntitiesFieldsEnum = "folderName"
	ListEntitiesFieldsExternalkey    ListEntitiesFieldsEnum = "externalKey"
	ListEntitiesFieldsPath           ListEntitiesFieldsEnum = "path"
	ListEntitiesFieldsUri            ListEntitiesFieldsEnum = "uri"
)

var mappingListEntitiesFieldsEnum = map[string]ListEntitiesFieldsEnum{
	"key":            ListEntitiesFieldsKey,
	"displayName":    ListEntitiesFieldsDisplayname,
	"description":    ListEntitiesFieldsDescription,
	"dataAssetKey":   ListEntitiesFieldsDataassetkey,
	"timeCreated":    ListEntitiesFieldsTimecreated,
	"timeUpdated":    ListEntitiesFieldsTimeupdated,
	"updatedById":    ListEntitiesFieldsUpdatedbyid,
	"lifecycleState": ListEntitiesFieldsLifecyclestate,
	"folderKey":      ListEntitiesFieldsFolderkey,
	"folderName":     ListEntitiesFieldsFoldername,
	"externalKey":    ListEntitiesFieldsExternalkey,
	"path":           ListEntitiesFieldsPath,
	"uri":            ListEntitiesFieldsUri,
}

var mappingListEntitiesFieldsEnumLowerCase = map[string]ListEntitiesFieldsEnum{
	"key":            ListEntitiesFieldsKey,
	"displayname":    ListEntitiesFieldsDisplayname,
	"description":    ListEntitiesFieldsDescription,
	"dataassetkey":   ListEntitiesFieldsDataassetkey,
	"timecreated":    ListEntitiesFieldsTimecreated,
	"timeupdated":    ListEntitiesFieldsTimeupdated,
	"updatedbyid":    ListEntitiesFieldsUpdatedbyid,
	"lifecyclestate": ListEntitiesFieldsLifecyclestate,
	"folderkey":      ListEntitiesFieldsFolderkey,
	"foldername":     ListEntitiesFieldsFoldername,
	"externalkey":    ListEntitiesFieldsExternalkey,
	"path":           ListEntitiesFieldsPath,
	"uri":            ListEntitiesFieldsUri,
}

// GetListEntitiesFieldsEnumValues Enumerates the set of values for ListEntitiesFieldsEnum
func GetListEntitiesFieldsEnumValues() []ListEntitiesFieldsEnum {
	values := make([]ListEntitiesFieldsEnum, 0)
	for _, v := range mappingListEntitiesFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitiesFieldsEnumStringValues Enumerates the set of values in String for ListEntitiesFieldsEnum
func GetListEntitiesFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"dataAssetKey",
		"timeCreated",
		"timeUpdated",
		"updatedById",
		"lifecycleState",
		"folderKey",
		"folderName",
		"externalKey",
		"path",
		"uri",
	}
}

// GetMappingListEntitiesFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitiesFieldsEnum(val string) (ListEntitiesFieldsEnum, bool) {
	enum, ok := mappingListEntitiesFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntitiesSortByEnum Enum with underlying type: string
type ListEntitiesSortByEnum string

// Set of constants representing the allowable values for ListEntitiesSortByEnum
const (
	ListEntitiesSortByTimecreated           ListEntitiesSortByEnum = "TIMECREATED"
	ListEntitiesSortByDisplayname           ListEntitiesSortByEnum = "DISPLAYNAME"
	ListEntitiesSortByDisplayorbusinessname ListEntitiesSortByEnum = "DISPLAYORBUSINESSNAME"
)

var mappingListEntitiesSortByEnum = map[string]ListEntitiesSortByEnum{
	"TIMECREATED":           ListEntitiesSortByTimecreated,
	"DISPLAYNAME":           ListEntitiesSortByDisplayname,
	"DISPLAYORBUSINESSNAME": ListEntitiesSortByDisplayorbusinessname,
}

var mappingListEntitiesSortByEnumLowerCase = map[string]ListEntitiesSortByEnum{
	"timecreated":           ListEntitiesSortByTimecreated,
	"displayname":           ListEntitiesSortByDisplayname,
	"displayorbusinessname": ListEntitiesSortByDisplayorbusinessname,
}

// GetListEntitiesSortByEnumValues Enumerates the set of values for ListEntitiesSortByEnum
func GetListEntitiesSortByEnumValues() []ListEntitiesSortByEnum {
	values := make([]ListEntitiesSortByEnum, 0)
	for _, v := range mappingListEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitiesSortByEnumStringValues Enumerates the set of values in String for ListEntitiesSortByEnum
func GetListEntitiesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
		"DISPLAYORBUSINESSNAME",
	}
}

// GetMappingListEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitiesSortByEnum(val string) (ListEntitiesSortByEnum, bool) {
	enum, ok := mappingListEntitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListEntitiesSortOrderEnum Enum with underlying type: string
type ListEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListEntitiesSortOrderEnum
const (
	ListEntitiesSortOrderAsc  ListEntitiesSortOrderEnum = "ASC"
	ListEntitiesSortOrderDesc ListEntitiesSortOrderEnum = "DESC"
)

var mappingListEntitiesSortOrderEnum = map[string]ListEntitiesSortOrderEnum{
	"ASC":  ListEntitiesSortOrderAsc,
	"DESC": ListEntitiesSortOrderDesc,
}

var mappingListEntitiesSortOrderEnumLowerCase = map[string]ListEntitiesSortOrderEnum{
	"asc":  ListEntitiesSortOrderAsc,
	"desc": ListEntitiesSortOrderDesc,
}

// GetListEntitiesSortOrderEnumValues Enumerates the set of values for ListEntitiesSortOrderEnum
func GetListEntitiesSortOrderEnumValues() []ListEntitiesSortOrderEnum {
	values := make([]ListEntitiesSortOrderEnum, 0)
	for _, v := range mappingListEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListEntitiesSortOrderEnum
func GetListEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListEntitiesSortOrderEnum(val string) (ListEntitiesSortOrderEnum, bool) {
	enum, ok := mappingListEntitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
