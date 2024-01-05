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

// ListFoldersRequest wrapper for the ListFolders operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListFolders.go.html to see an example of how to use ListFoldersRequest.
type ListFoldersRequest struct {

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

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// A filter to return only resources that match the specified lifecycle state. The value is case insensitive.
	LifecycleState ListFoldersLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Unique folder key.
	ParentFolderKey *string `mandatory:"false" contributesTo:"query" name:"parentFolderKey"`

	// Full path of the resource for resources that support paths.
	Path *string `mandatory:"false" contributesTo:"query" name:"path"`

	// Unique external identifier of this resource in the external source system.
	ExternalKey *string `mandatory:"false" contributesTo:"query" name:"externalKey"`

	// Time that the resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeCreated"`

	// Time that the resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" contributesTo:"query" name:"timeUpdated"`

	// OCID of the user who created the resource.
	CreatedById *string `mandatory:"false" contributesTo:"query" name:"createdById"`

	// OCID of the user who updated the resource.
	UpdatedById *string `mandatory:"false" contributesTo:"query" name:"updatedById"`

	// Harvest status of the harvestable resource as updated by the harvest process.
	HarvestStatus ListFoldersHarvestStatusEnum `mandatory:"false" contributesTo:"query" name:"harvestStatus" omitEmpty:"true"`

	// Key of the last harvest process to update this resource.
	LastJobKey *string `mandatory:"false" contributesTo:"query" name:"lastJobKey"`

	// Specifies the fields to return in a folder summary response.
	Fields []ListFoldersFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The key of the object type.
	TypeKey *string `mandatory:"false" contributesTo:"query" name:"typeKey"`

	// The field to sort by. Only one sort order may be provided. DISPLAYORBUSINESSNAME considers businessName of a given object if set, else its displayName is used.
	// Default sort order for TIMECREATED is descending and default sort order for DISPLAYNAME and DISPLAYORBUSINESSNAME is ascending. If no order is specified, TIMECREATED is the default.
	SortBy ListFoldersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListFoldersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFoldersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFoldersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFoldersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFoldersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFoldersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFoldersLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFoldersLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFoldersHarvestStatusEnum(string(request.HarvestStatus)); !ok && request.HarvestStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for HarvestStatus: %s. Supported values are: %s.", request.HarvestStatus, strings.Join(GetListFoldersHarvestStatusEnumStringValues(), ",")))
	}
	for _, val := range request.Fields {
		if _, ok := GetMappingListFoldersFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListFoldersFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListFoldersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFoldersSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFoldersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFoldersSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFoldersResponse wrapper for the ListFolders operation
type ListFoldersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FolderCollection instances
	FolderCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFoldersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFoldersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFoldersLifecycleStateEnum Enum with underlying type: string
type ListFoldersLifecycleStateEnum string

// Set of constants representing the allowable values for ListFoldersLifecycleStateEnum
const (
	ListFoldersLifecycleStateCreating ListFoldersLifecycleStateEnum = "CREATING"
	ListFoldersLifecycleStateActive   ListFoldersLifecycleStateEnum = "ACTIVE"
	ListFoldersLifecycleStateInactive ListFoldersLifecycleStateEnum = "INACTIVE"
	ListFoldersLifecycleStateUpdating ListFoldersLifecycleStateEnum = "UPDATING"
	ListFoldersLifecycleStateDeleting ListFoldersLifecycleStateEnum = "DELETING"
	ListFoldersLifecycleStateDeleted  ListFoldersLifecycleStateEnum = "DELETED"
	ListFoldersLifecycleStateFailed   ListFoldersLifecycleStateEnum = "FAILED"
	ListFoldersLifecycleStateMoving   ListFoldersLifecycleStateEnum = "MOVING"
)

var mappingListFoldersLifecycleStateEnum = map[string]ListFoldersLifecycleStateEnum{
	"CREATING": ListFoldersLifecycleStateCreating,
	"ACTIVE":   ListFoldersLifecycleStateActive,
	"INACTIVE": ListFoldersLifecycleStateInactive,
	"UPDATING": ListFoldersLifecycleStateUpdating,
	"DELETING": ListFoldersLifecycleStateDeleting,
	"DELETED":  ListFoldersLifecycleStateDeleted,
	"FAILED":   ListFoldersLifecycleStateFailed,
	"MOVING":   ListFoldersLifecycleStateMoving,
}

var mappingListFoldersLifecycleStateEnumLowerCase = map[string]ListFoldersLifecycleStateEnum{
	"creating": ListFoldersLifecycleStateCreating,
	"active":   ListFoldersLifecycleStateActive,
	"inactive": ListFoldersLifecycleStateInactive,
	"updating": ListFoldersLifecycleStateUpdating,
	"deleting": ListFoldersLifecycleStateDeleting,
	"deleted":  ListFoldersLifecycleStateDeleted,
	"failed":   ListFoldersLifecycleStateFailed,
	"moving":   ListFoldersLifecycleStateMoving,
}

// GetListFoldersLifecycleStateEnumValues Enumerates the set of values for ListFoldersLifecycleStateEnum
func GetListFoldersLifecycleStateEnumValues() []ListFoldersLifecycleStateEnum {
	values := make([]ListFoldersLifecycleStateEnum, 0)
	for _, v := range mappingListFoldersLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFoldersLifecycleStateEnumStringValues Enumerates the set of values in String for ListFoldersLifecycleStateEnum
func GetListFoldersLifecycleStateEnumStringValues() []string {
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

// GetMappingListFoldersLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFoldersLifecycleStateEnum(val string) (ListFoldersLifecycleStateEnum, bool) {
	enum, ok := mappingListFoldersLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFoldersHarvestStatusEnum Enum with underlying type: string
type ListFoldersHarvestStatusEnum string

// Set of constants representing the allowable values for ListFoldersHarvestStatusEnum
const (
	ListFoldersHarvestStatusComplete   ListFoldersHarvestStatusEnum = "COMPLETE"
	ListFoldersHarvestStatusError      ListFoldersHarvestStatusEnum = "ERROR"
	ListFoldersHarvestStatusInProgress ListFoldersHarvestStatusEnum = "IN_PROGRESS"
	ListFoldersHarvestStatusDeferred   ListFoldersHarvestStatusEnum = "DEFERRED"
)

var mappingListFoldersHarvestStatusEnum = map[string]ListFoldersHarvestStatusEnum{
	"COMPLETE":    ListFoldersHarvestStatusComplete,
	"ERROR":       ListFoldersHarvestStatusError,
	"IN_PROGRESS": ListFoldersHarvestStatusInProgress,
	"DEFERRED":    ListFoldersHarvestStatusDeferred,
}

var mappingListFoldersHarvestStatusEnumLowerCase = map[string]ListFoldersHarvestStatusEnum{
	"complete":    ListFoldersHarvestStatusComplete,
	"error":       ListFoldersHarvestStatusError,
	"in_progress": ListFoldersHarvestStatusInProgress,
	"deferred":    ListFoldersHarvestStatusDeferred,
}

// GetListFoldersHarvestStatusEnumValues Enumerates the set of values for ListFoldersHarvestStatusEnum
func GetListFoldersHarvestStatusEnumValues() []ListFoldersHarvestStatusEnum {
	values := make([]ListFoldersHarvestStatusEnum, 0)
	for _, v := range mappingListFoldersHarvestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListFoldersHarvestStatusEnumStringValues Enumerates the set of values in String for ListFoldersHarvestStatusEnum
func GetListFoldersHarvestStatusEnumStringValues() []string {
	return []string{
		"COMPLETE",
		"ERROR",
		"IN_PROGRESS",
		"DEFERRED",
	}
}

// GetMappingListFoldersHarvestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFoldersHarvestStatusEnum(val string) (ListFoldersHarvestStatusEnum, bool) {
	enum, ok := mappingListFoldersHarvestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFoldersFieldsEnum Enum with underlying type: string
type ListFoldersFieldsEnum string

// Set of constants representing the allowable values for ListFoldersFieldsEnum
const (
	ListFoldersFieldsKey             ListFoldersFieldsEnum = "key"
	ListFoldersFieldsDisplayname     ListFoldersFieldsEnum = "displayName"
	ListFoldersFieldsDescription     ListFoldersFieldsEnum = "description"
	ListFoldersFieldsParentfolderkey ListFoldersFieldsEnum = "parentFolderKey"
	ListFoldersFieldsPath            ListFoldersFieldsEnum = "path"
	ListFoldersFieldsDataassetkey    ListFoldersFieldsEnum = "dataAssetKey"
	ListFoldersFieldsExternalkey     ListFoldersFieldsEnum = "externalKey"
	ListFoldersFieldsTimeexternal    ListFoldersFieldsEnum = "timeExternal"
	ListFoldersFieldsTimecreated     ListFoldersFieldsEnum = "timeCreated"
	ListFoldersFieldsLifecyclestate  ListFoldersFieldsEnum = "lifecycleState"
	ListFoldersFieldsUri             ListFoldersFieldsEnum = "uri"
)

var mappingListFoldersFieldsEnum = map[string]ListFoldersFieldsEnum{
	"key":             ListFoldersFieldsKey,
	"displayName":     ListFoldersFieldsDisplayname,
	"description":     ListFoldersFieldsDescription,
	"parentFolderKey": ListFoldersFieldsParentfolderkey,
	"path":            ListFoldersFieldsPath,
	"dataAssetKey":    ListFoldersFieldsDataassetkey,
	"externalKey":     ListFoldersFieldsExternalkey,
	"timeExternal":    ListFoldersFieldsTimeexternal,
	"timeCreated":     ListFoldersFieldsTimecreated,
	"lifecycleState":  ListFoldersFieldsLifecyclestate,
	"uri":             ListFoldersFieldsUri,
}

var mappingListFoldersFieldsEnumLowerCase = map[string]ListFoldersFieldsEnum{
	"key":             ListFoldersFieldsKey,
	"displayname":     ListFoldersFieldsDisplayname,
	"description":     ListFoldersFieldsDescription,
	"parentfolderkey": ListFoldersFieldsParentfolderkey,
	"path":            ListFoldersFieldsPath,
	"dataassetkey":    ListFoldersFieldsDataassetkey,
	"externalkey":     ListFoldersFieldsExternalkey,
	"timeexternal":    ListFoldersFieldsTimeexternal,
	"timecreated":     ListFoldersFieldsTimecreated,
	"lifecyclestate":  ListFoldersFieldsLifecyclestate,
	"uri":             ListFoldersFieldsUri,
}

// GetListFoldersFieldsEnumValues Enumerates the set of values for ListFoldersFieldsEnum
func GetListFoldersFieldsEnumValues() []ListFoldersFieldsEnum {
	values := make([]ListFoldersFieldsEnum, 0)
	for _, v := range mappingListFoldersFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListFoldersFieldsEnumStringValues Enumerates the set of values in String for ListFoldersFieldsEnum
func GetListFoldersFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"parentFolderKey",
		"path",
		"dataAssetKey",
		"externalKey",
		"timeExternal",
		"timeCreated",
		"lifecycleState",
		"uri",
	}
}

// GetMappingListFoldersFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFoldersFieldsEnum(val string) (ListFoldersFieldsEnum, bool) {
	enum, ok := mappingListFoldersFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFoldersSortByEnum Enum with underlying type: string
type ListFoldersSortByEnum string

// Set of constants representing the allowable values for ListFoldersSortByEnum
const (
	ListFoldersSortByTimecreated           ListFoldersSortByEnum = "TIMECREATED"
	ListFoldersSortByDisplayname           ListFoldersSortByEnum = "DISPLAYNAME"
	ListFoldersSortByDisplayorbusinessname ListFoldersSortByEnum = "DISPLAYORBUSINESSNAME"
)

var mappingListFoldersSortByEnum = map[string]ListFoldersSortByEnum{
	"TIMECREATED":           ListFoldersSortByTimecreated,
	"DISPLAYNAME":           ListFoldersSortByDisplayname,
	"DISPLAYORBUSINESSNAME": ListFoldersSortByDisplayorbusinessname,
}

var mappingListFoldersSortByEnumLowerCase = map[string]ListFoldersSortByEnum{
	"timecreated":           ListFoldersSortByTimecreated,
	"displayname":           ListFoldersSortByDisplayname,
	"displayorbusinessname": ListFoldersSortByDisplayorbusinessname,
}

// GetListFoldersSortByEnumValues Enumerates the set of values for ListFoldersSortByEnum
func GetListFoldersSortByEnumValues() []ListFoldersSortByEnum {
	values := make([]ListFoldersSortByEnum, 0)
	for _, v := range mappingListFoldersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFoldersSortByEnumStringValues Enumerates the set of values in String for ListFoldersSortByEnum
func GetListFoldersSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
		"DISPLAYORBUSINESSNAME",
	}
}

// GetMappingListFoldersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFoldersSortByEnum(val string) (ListFoldersSortByEnum, bool) {
	enum, ok := mappingListFoldersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFoldersSortOrderEnum Enum with underlying type: string
type ListFoldersSortOrderEnum string

// Set of constants representing the allowable values for ListFoldersSortOrderEnum
const (
	ListFoldersSortOrderAsc  ListFoldersSortOrderEnum = "ASC"
	ListFoldersSortOrderDesc ListFoldersSortOrderEnum = "DESC"
)

var mappingListFoldersSortOrderEnum = map[string]ListFoldersSortOrderEnum{
	"ASC":  ListFoldersSortOrderAsc,
	"DESC": ListFoldersSortOrderDesc,
}

var mappingListFoldersSortOrderEnumLowerCase = map[string]ListFoldersSortOrderEnum{
	"asc":  ListFoldersSortOrderAsc,
	"desc": ListFoldersSortOrderDesc,
}

// GetListFoldersSortOrderEnumValues Enumerates the set of values for ListFoldersSortOrderEnum
func GetListFoldersSortOrderEnumValues() []ListFoldersSortOrderEnum {
	values := make([]ListFoldersSortOrderEnum, 0)
	for _, v := range mappingListFoldersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFoldersSortOrderEnumStringValues Enumerates the set of values in String for ListFoldersSortOrderEnum
func GetListFoldersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFoldersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFoldersSortOrderEnum(val string) (ListFoldersSortOrderEnum, bool) {
	enum, ok := mappingListFoldersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
