// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetEntityRequest wrapper for the GetEntity operation
type GetEntityRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Specifies the fields to return in an entity response.
	Fields []GetEntityFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetEntityRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetEntityRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetEntityRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetEntityResponse wrapper for the GetEntity operation
type GetEntityResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Entity instance
	Entity `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetEntityResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetEntityResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetEntityFieldsEnum Enum with underlying type: string
type GetEntityFieldsEnum string

// Set of constants representing the allowable values for GetEntityFieldsEnum
const (
	GetEntityFieldsKey               GetEntityFieldsEnum = "key"
	GetEntityFieldsDisplayname       GetEntityFieldsEnum = "displayName"
	GetEntityFieldsDescription       GetEntityFieldsEnum = "description"
	GetEntityFieldsDataassetkey      GetEntityFieldsEnum = "dataAssetKey"
	GetEntityFieldsTimecreated       GetEntityFieldsEnum = "timeCreated"
	GetEntityFieldsTimeupdated       GetEntityFieldsEnum = "timeUpdated"
	GetEntityFieldsCreatedbyid       GetEntityFieldsEnum = "createdById"
	GetEntityFieldsUpdatedbyid       GetEntityFieldsEnum = "updatedById"
	GetEntityFieldsLifecyclestate    GetEntityFieldsEnum = "lifecycleState"
	GetEntityFieldsExternalkey       GetEntityFieldsEnum = "externalKey"
	GetEntityFieldsTimeexternal      GetEntityFieldsEnum = "timeExternal"
	GetEntityFieldsTimestatusupdated GetEntityFieldsEnum = "timeStatusUpdated"
	GetEntityFieldsIslogical         GetEntityFieldsEnum = "isLogical"
	GetEntityFieldsIspartition       GetEntityFieldsEnum = "isPartition"
	GetEntityFieldsFolderkey         GetEntityFieldsEnum = "folderKey"
	GetEntityFieldsFoldername        GetEntityFieldsEnum = "folderName"
	GetEntityFieldsTypekey           GetEntityFieldsEnum = "typeKey"
	GetEntityFieldsPath              GetEntityFieldsEnum = "path"
	GetEntityFieldsHarveststatus     GetEntityFieldsEnum = "harvestStatus"
	GetEntityFieldsLastjobkey        GetEntityFieldsEnum = "lastJobKey"
	GetEntityFieldsUri               GetEntityFieldsEnum = "uri"
	GetEntityFieldsProperties        GetEntityFieldsEnum = "properties"
)

var mappingGetEntityFields = map[string]GetEntityFieldsEnum{
	"key":               GetEntityFieldsKey,
	"displayName":       GetEntityFieldsDisplayname,
	"description":       GetEntityFieldsDescription,
	"dataAssetKey":      GetEntityFieldsDataassetkey,
	"timeCreated":       GetEntityFieldsTimecreated,
	"timeUpdated":       GetEntityFieldsTimeupdated,
	"createdById":       GetEntityFieldsCreatedbyid,
	"updatedById":       GetEntityFieldsUpdatedbyid,
	"lifecycleState":    GetEntityFieldsLifecyclestate,
	"externalKey":       GetEntityFieldsExternalkey,
	"timeExternal":      GetEntityFieldsTimeexternal,
	"timeStatusUpdated": GetEntityFieldsTimestatusupdated,
	"isLogical":         GetEntityFieldsIslogical,
	"isPartition":       GetEntityFieldsIspartition,
	"folderKey":         GetEntityFieldsFolderkey,
	"folderName":        GetEntityFieldsFoldername,
	"typeKey":           GetEntityFieldsTypekey,
	"path":              GetEntityFieldsPath,
	"harvestStatus":     GetEntityFieldsHarveststatus,
	"lastJobKey":        GetEntityFieldsLastjobkey,
	"uri":               GetEntityFieldsUri,
	"properties":        GetEntityFieldsProperties,
}

// GetGetEntityFieldsEnumValues Enumerates the set of values for GetEntityFieldsEnum
func GetGetEntityFieldsEnumValues() []GetEntityFieldsEnum {
	values := make([]GetEntityFieldsEnum, 0)
	for _, v := range mappingGetEntityFields {
		values = append(values, v)
	}
	return values
}
