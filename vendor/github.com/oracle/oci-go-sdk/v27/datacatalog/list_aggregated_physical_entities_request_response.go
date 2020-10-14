// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/v27/common"
	"net/http"
)

// ListAggregatedPhysicalEntitiesRequest wrapper for the ListAggregatedPhysicalEntities operation
type ListAggregatedPhysicalEntitiesRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Specifies the fields to return in an entity response.
	Fields []ListAggregatedPhysicalEntitiesFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

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
func (request ListAggregatedPhysicalEntitiesRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAggregatedPhysicalEntitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListAggregatedPhysicalEntitiesResponse wrapper for the ListAggregatedPhysicalEntities operation
type ListAggregatedPhysicalEntitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The EntityCollection instance
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
