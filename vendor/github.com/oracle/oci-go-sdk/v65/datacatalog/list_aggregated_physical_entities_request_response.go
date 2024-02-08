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

// ListAggregatedPhysicalEntitiesRequest wrapper for the ListAggregatedPhysicalEntities operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/ListAggregatedPhysicalEntities.go.html to see an example of how to use ListAggregatedPhysicalEntitiesRequest.
type ListAggregatedPhysicalEntitiesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Specifies the fields to return in an entity response.
	Fields []ListAggregatedPhysicalEntitiesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// A filter to return only resources that match display name pattern given. The match is not case sensitive.
	// For Example : /folders?displayNameContains=Cu.*
	// The above would match all folders with display name that starts with "Cu" or has the pattern "Cu" anywhere in between.
	DisplayNameContains *string `mandatory:"false" contributesTo:"query" name:"displayNameContains"`

	// The field to sort by. Only one sort order may be provided. Default order for TIMECREATED is descending. Default order for DISPLAYNAME is ascending. If no value is specified TIMECREATED is default.
	SortBy ListAggregatedPhysicalEntitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListAggregatedPhysicalEntitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Indicates whether the properties map will be provided in the response.
	IsIncludeProperties *bool `mandatory:"false" contributesTo:"query" name:"isIncludeProperties"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAggregatedPhysicalEntitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAggregatedPhysicalEntitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAggregatedPhysicalEntitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAggregatedPhysicalEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAggregatedPhysicalEntitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingListAggregatedPhysicalEntitiesFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetListAggregatedPhysicalEntitiesFieldsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListAggregatedPhysicalEntitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAggregatedPhysicalEntitiesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAggregatedPhysicalEntitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAggregatedPhysicalEntitiesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAggregatedPhysicalEntitiesResponse wrapper for the ListAggregatedPhysicalEntities operation
type ListAggregatedPhysicalEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of EntityCollection instances
	EntityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListAggregatedPhysicalEntitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAggregatedPhysicalEntitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAggregatedPhysicalEntitiesFieldsEnum Enum with underlying type: string
type ListAggregatedPhysicalEntitiesFieldsEnum string

// Set of constants representing the allowable values for ListAggregatedPhysicalEntitiesFieldsEnum
const (
	ListAggregatedPhysicalEntitiesFieldsKey               ListAggregatedPhysicalEntitiesFieldsEnum = "key"
	ListAggregatedPhysicalEntitiesFieldsDisplayname       ListAggregatedPhysicalEntitiesFieldsEnum = "displayName"
	ListAggregatedPhysicalEntitiesFieldsDescription       ListAggregatedPhysicalEntitiesFieldsEnum = "description"
	ListAggregatedPhysicalEntitiesFieldsDataassetkey      ListAggregatedPhysicalEntitiesFieldsEnum = "dataAssetKey"
	ListAggregatedPhysicalEntitiesFieldsTimecreated       ListAggregatedPhysicalEntitiesFieldsEnum = "timeCreated"
	ListAggregatedPhysicalEntitiesFieldsTimeupdated       ListAggregatedPhysicalEntitiesFieldsEnum = "timeUpdated"
	ListAggregatedPhysicalEntitiesFieldsCreatedbyid       ListAggregatedPhysicalEntitiesFieldsEnum = "createdById"
	ListAggregatedPhysicalEntitiesFieldsUpdatedbyid       ListAggregatedPhysicalEntitiesFieldsEnum = "updatedById"
	ListAggregatedPhysicalEntitiesFieldsLifecyclestate    ListAggregatedPhysicalEntitiesFieldsEnum = "lifecycleState"
	ListAggregatedPhysicalEntitiesFieldsExternalkey       ListAggregatedPhysicalEntitiesFieldsEnum = "externalKey"
	ListAggregatedPhysicalEntitiesFieldsTimeexternal      ListAggregatedPhysicalEntitiesFieldsEnum = "timeExternal"
	ListAggregatedPhysicalEntitiesFieldsTimestatusupdated ListAggregatedPhysicalEntitiesFieldsEnum = "timeStatusUpdated"
	ListAggregatedPhysicalEntitiesFieldsIslogical         ListAggregatedPhysicalEntitiesFieldsEnum = "isLogical"
	ListAggregatedPhysicalEntitiesFieldsIspartition       ListAggregatedPhysicalEntitiesFieldsEnum = "isPartition"
	ListAggregatedPhysicalEntitiesFieldsFolderkey         ListAggregatedPhysicalEntitiesFieldsEnum = "folderKey"
	ListAggregatedPhysicalEntitiesFieldsFoldername        ListAggregatedPhysicalEntitiesFieldsEnum = "folderName"
	ListAggregatedPhysicalEntitiesFieldsTypekey           ListAggregatedPhysicalEntitiesFieldsEnum = "typeKey"
	ListAggregatedPhysicalEntitiesFieldsPath              ListAggregatedPhysicalEntitiesFieldsEnum = "path"
	ListAggregatedPhysicalEntitiesFieldsHarveststatus     ListAggregatedPhysicalEntitiesFieldsEnum = "harvestStatus"
	ListAggregatedPhysicalEntitiesFieldsLastjobkey        ListAggregatedPhysicalEntitiesFieldsEnum = "lastJobKey"
	ListAggregatedPhysicalEntitiesFieldsUri               ListAggregatedPhysicalEntitiesFieldsEnum = "uri"
	ListAggregatedPhysicalEntitiesFieldsProperties        ListAggregatedPhysicalEntitiesFieldsEnum = "properties"
)

var mappingListAggregatedPhysicalEntitiesFieldsEnum = map[string]ListAggregatedPhysicalEntitiesFieldsEnum{
	"key":               ListAggregatedPhysicalEntitiesFieldsKey,
	"displayName":       ListAggregatedPhysicalEntitiesFieldsDisplayname,
	"description":       ListAggregatedPhysicalEntitiesFieldsDescription,
	"dataAssetKey":      ListAggregatedPhysicalEntitiesFieldsDataassetkey,
	"timeCreated":       ListAggregatedPhysicalEntitiesFieldsTimecreated,
	"timeUpdated":       ListAggregatedPhysicalEntitiesFieldsTimeupdated,
	"createdById":       ListAggregatedPhysicalEntitiesFieldsCreatedbyid,
	"updatedById":       ListAggregatedPhysicalEntitiesFieldsUpdatedbyid,
	"lifecycleState":    ListAggregatedPhysicalEntitiesFieldsLifecyclestate,
	"externalKey":       ListAggregatedPhysicalEntitiesFieldsExternalkey,
	"timeExternal":      ListAggregatedPhysicalEntitiesFieldsTimeexternal,
	"timeStatusUpdated": ListAggregatedPhysicalEntitiesFieldsTimestatusupdated,
	"isLogical":         ListAggregatedPhysicalEntitiesFieldsIslogical,
	"isPartition":       ListAggregatedPhysicalEntitiesFieldsIspartition,
	"folderKey":         ListAggregatedPhysicalEntitiesFieldsFolderkey,
	"folderName":        ListAggregatedPhysicalEntitiesFieldsFoldername,
	"typeKey":           ListAggregatedPhysicalEntitiesFieldsTypekey,
	"path":              ListAggregatedPhysicalEntitiesFieldsPath,
	"harvestStatus":     ListAggregatedPhysicalEntitiesFieldsHarveststatus,
	"lastJobKey":        ListAggregatedPhysicalEntitiesFieldsLastjobkey,
	"uri":               ListAggregatedPhysicalEntitiesFieldsUri,
	"properties":        ListAggregatedPhysicalEntitiesFieldsProperties,
}

var mappingListAggregatedPhysicalEntitiesFieldsEnumLowerCase = map[string]ListAggregatedPhysicalEntitiesFieldsEnum{
	"key":               ListAggregatedPhysicalEntitiesFieldsKey,
	"displayname":       ListAggregatedPhysicalEntitiesFieldsDisplayname,
	"description":       ListAggregatedPhysicalEntitiesFieldsDescription,
	"dataassetkey":      ListAggregatedPhysicalEntitiesFieldsDataassetkey,
	"timecreated":       ListAggregatedPhysicalEntitiesFieldsTimecreated,
	"timeupdated":       ListAggregatedPhysicalEntitiesFieldsTimeupdated,
	"createdbyid":       ListAggregatedPhysicalEntitiesFieldsCreatedbyid,
	"updatedbyid":       ListAggregatedPhysicalEntitiesFieldsUpdatedbyid,
	"lifecyclestate":    ListAggregatedPhysicalEntitiesFieldsLifecyclestate,
	"externalkey":       ListAggregatedPhysicalEntitiesFieldsExternalkey,
	"timeexternal":      ListAggregatedPhysicalEntitiesFieldsTimeexternal,
	"timestatusupdated": ListAggregatedPhysicalEntitiesFieldsTimestatusupdated,
	"islogical":         ListAggregatedPhysicalEntitiesFieldsIslogical,
	"ispartition":       ListAggregatedPhysicalEntitiesFieldsIspartition,
	"folderkey":         ListAggregatedPhysicalEntitiesFieldsFolderkey,
	"foldername":        ListAggregatedPhysicalEntitiesFieldsFoldername,
	"typekey":           ListAggregatedPhysicalEntitiesFieldsTypekey,
	"path":              ListAggregatedPhysicalEntitiesFieldsPath,
	"harveststatus":     ListAggregatedPhysicalEntitiesFieldsHarveststatus,
	"lastjobkey":        ListAggregatedPhysicalEntitiesFieldsLastjobkey,
	"uri":               ListAggregatedPhysicalEntitiesFieldsUri,
	"properties":        ListAggregatedPhysicalEntitiesFieldsProperties,
}

// GetListAggregatedPhysicalEntitiesFieldsEnumValues Enumerates the set of values for ListAggregatedPhysicalEntitiesFieldsEnum
func GetListAggregatedPhysicalEntitiesFieldsEnumValues() []ListAggregatedPhysicalEntitiesFieldsEnum {
	values := make([]ListAggregatedPhysicalEntitiesFieldsEnum, 0)
	for _, v := range mappingListAggregatedPhysicalEntitiesFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetListAggregatedPhysicalEntitiesFieldsEnumStringValues Enumerates the set of values in String for ListAggregatedPhysicalEntitiesFieldsEnum
func GetListAggregatedPhysicalEntitiesFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"dataAssetKey",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"lifecycleState",
		"externalKey",
		"timeExternal",
		"timeStatusUpdated",
		"isLogical",
		"isPartition",
		"folderKey",
		"folderName",
		"typeKey",
		"path",
		"harvestStatus",
		"lastJobKey",
		"uri",
		"properties",
	}
}

// GetMappingListAggregatedPhysicalEntitiesFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAggregatedPhysicalEntitiesFieldsEnum(val string) (ListAggregatedPhysicalEntitiesFieldsEnum, bool) {
	enum, ok := mappingListAggregatedPhysicalEntitiesFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAggregatedPhysicalEntitiesSortByEnum Enum with underlying type: string
type ListAggregatedPhysicalEntitiesSortByEnum string

// Set of constants representing the allowable values for ListAggregatedPhysicalEntitiesSortByEnum
const (
	ListAggregatedPhysicalEntitiesSortByTimecreated ListAggregatedPhysicalEntitiesSortByEnum = "TIMECREATED"
	ListAggregatedPhysicalEntitiesSortByDisplayname ListAggregatedPhysicalEntitiesSortByEnum = "DISPLAYNAME"
)

var mappingListAggregatedPhysicalEntitiesSortByEnum = map[string]ListAggregatedPhysicalEntitiesSortByEnum{
	"TIMECREATED": ListAggregatedPhysicalEntitiesSortByTimecreated,
	"DISPLAYNAME": ListAggregatedPhysicalEntitiesSortByDisplayname,
}

var mappingListAggregatedPhysicalEntitiesSortByEnumLowerCase = map[string]ListAggregatedPhysicalEntitiesSortByEnum{
	"timecreated": ListAggregatedPhysicalEntitiesSortByTimecreated,
	"displayname": ListAggregatedPhysicalEntitiesSortByDisplayname,
}

// GetListAggregatedPhysicalEntitiesSortByEnumValues Enumerates the set of values for ListAggregatedPhysicalEntitiesSortByEnum
func GetListAggregatedPhysicalEntitiesSortByEnumValues() []ListAggregatedPhysicalEntitiesSortByEnum {
	values := make([]ListAggregatedPhysicalEntitiesSortByEnum, 0)
	for _, v := range mappingListAggregatedPhysicalEntitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAggregatedPhysicalEntitiesSortByEnumStringValues Enumerates the set of values in String for ListAggregatedPhysicalEntitiesSortByEnum
func GetListAggregatedPhysicalEntitiesSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListAggregatedPhysicalEntitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAggregatedPhysicalEntitiesSortByEnum(val string) (ListAggregatedPhysicalEntitiesSortByEnum, bool) {
	enum, ok := mappingListAggregatedPhysicalEntitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAggregatedPhysicalEntitiesSortOrderEnum Enum with underlying type: string
type ListAggregatedPhysicalEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListAggregatedPhysicalEntitiesSortOrderEnum
const (
	ListAggregatedPhysicalEntitiesSortOrderAsc  ListAggregatedPhysicalEntitiesSortOrderEnum = "ASC"
	ListAggregatedPhysicalEntitiesSortOrderDesc ListAggregatedPhysicalEntitiesSortOrderEnum = "DESC"
)

var mappingListAggregatedPhysicalEntitiesSortOrderEnum = map[string]ListAggregatedPhysicalEntitiesSortOrderEnum{
	"ASC":  ListAggregatedPhysicalEntitiesSortOrderAsc,
	"DESC": ListAggregatedPhysicalEntitiesSortOrderDesc,
}

var mappingListAggregatedPhysicalEntitiesSortOrderEnumLowerCase = map[string]ListAggregatedPhysicalEntitiesSortOrderEnum{
	"asc":  ListAggregatedPhysicalEntitiesSortOrderAsc,
	"desc": ListAggregatedPhysicalEntitiesSortOrderDesc,
}

// GetListAggregatedPhysicalEntitiesSortOrderEnumValues Enumerates the set of values for ListAggregatedPhysicalEntitiesSortOrderEnum
func GetListAggregatedPhysicalEntitiesSortOrderEnumValues() []ListAggregatedPhysicalEntitiesSortOrderEnum {
	values := make([]ListAggregatedPhysicalEntitiesSortOrderEnum, 0)
	for _, v := range mappingListAggregatedPhysicalEntitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAggregatedPhysicalEntitiesSortOrderEnumStringValues Enumerates the set of values in String for ListAggregatedPhysicalEntitiesSortOrderEnum
func GetListAggregatedPhysicalEntitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAggregatedPhysicalEntitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAggregatedPhysicalEntitiesSortOrderEnum(val string) (ListAggregatedPhysicalEntitiesSortOrderEnum, bool) {
	enum, ok := mappingListAggregatedPhysicalEntitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
