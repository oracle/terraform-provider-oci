// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListAggregatedPhysicalEntitiesRequest wrapper for the ListAggregatedPhysicalEntities operation
//
// See also
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

var mappingListAggregatedPhysicalEntitiesFields = map[string]ListAggregatedPhysicalEntitiesFieldsEnum{
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

// GetListAggregatedPhysicalEntitiesFieldsEnumValues Enumerates the set of values for ListAggregatedPhysicalEntitiesFieldsEnum
func GetListAggregatedPhysicalEntitiesFieldsEnumValues() []ListAggregatedPhysicalEntitiesFieldsEnum {
	values := make([]ListAggregatedPhysicalEntitiesFieldsEnum, 0)
	for _, v := range mappingListAggregatedPhysicalEntitiesFields {
		values = append(values, v)
	}
	return values
}

// ListAggregatedPhysicalEntitiesSortByEnum Enum with underlying type: string
type ListAggregatedPhysicalEntitiesSortByEnum string

// Set of constants representing the allowable values for ListAggregatedPhysicalEntitiesSortByEnum
const (
	ListAggregatedPhysicalEntitiesSortByTimecreated ListAggregatedPhysicalEntitiesSortByEnum = "TIMECREATED"
	ListAggregatedPhysicalEntitiesSortByDisplayname ListAggregatedPhysicalEntitiesSortByEnum = "DISPLAYNAME"
)

var mappingListAggregatedPhysicalEntitiesSortBy = map[string]ListAggregatedPhysicalEntitiesSortByEnum{
	"TIMECREATED": ListAggregatedPhysicalEntitiesSortByTimecreated,
	"DISPLAYNAME": ListAggregatedPhysicalEntitiesSortByDisplayname,
}

// GetListAggregatedPhysicalEntitiesSortByEnumValues Enumerates the set of values for ListAggregatedPhysicalEntitiesSortByEnum
func GetListAggregatedPhysicalEntitiesSortByEnumValues() []ListAggregatedPhysicalEntitiesSortByEnum {
	values := make([]ListAggregatedPhysicalEntitiesSortByEnum, 0)
	for _, v := range mappingListAggregatedPhysicalEntitiesSortBy {
		values = append(values, v)
	}
	return values
}

// ListAggregatedPhysicalEntitiesSortOrderEnum Enum with underlying type: string
type ListAggregatedPhysicalEntitiesSortOrderEnum string

// Set of constants representing the allowable values for ListAggregatedPhysicalEntitiesSortOrderEnum
const (
	ListAggregatedPhysicalEntitiesSortOrderAsc  ListAggregatedPhysicalEntitiesSortOrderEnum = "ASC"
	ListAggregatedPhysicalEntitiesSortOrderDesc ListAggregatedPhysicalEntitiesSortOrderEnum = "DESC"
)

var mappingListAggregatedPhysicalEntitiesSortOrder = map[string]ListAggregatedPhysicalEntitiesSortOrderEnum{
	"ASC":  ListAggregatedPhysicalEntitiesSortOrderAsc,
	"DESC": ListAggregatedPhysicalEntitiesSortOrderDesc,
}

// GetListAggregatedPhysicalEntitiesSortOrderEnumValues Enumerates the set of values for ListAggregatedPhysicalEntitiesSortOrderEnum
func GetListAggregatedPhysicalEntitiesSortOrderEnumValues() []ListAggregatedPhysicalEntitiesSortOrderEnum {
	values := make([]ListAggregatedPhysicalEntitiesSortOrderEnum, 0)
	for _, v := range mappingListAggregatedPhysicalEntitiesSortOrder {
		values = append(values, v)
	}
	return values
}
