// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetFolderRequest wrapper for the GetFolder operation
type GetFolderRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique folder key.
	FolderKey *string `mandatory:"true" contributesTo:"path" name:"folderKey"`

	// Specifies the fields to return in a folder response.
	Fields []GetFolderFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetFolderRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetFolderRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetFolderRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetFolderResponse wrapper for the GetFolder operation
type GetFolderResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Folder instance
	Folder `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetFolderResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetFolderResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetFolderFieldsEnum Enum with underlying type: string
type GetFolderFieldsEnum string

// Set of constants representing the allowable values for GetFolderFieldsEnum
const (
	GetFolderFieldsKey             GetFolderFieldsEnum = "key"
	GetFolderFieldsDisplayname     GetFolderFieldsEnum = "displayName"
	GetFolderFieldsDescription     GetFolderFieldsEnum = "description"
	GetFolderFieldsParentfolderkey GetFolderFieldsEnum = "parentFolderKey"
	GetFolderFieldsPath            GetFolderFieldsEnum = "path"
	GetFolderFieldsDataassetkey    GetFolderFieldsEnum = "dataAssetKey"
	GetFolderFieldsProperties      GetFolderFieldsEnum = "properties"
	GetFolderFieldsExternalkey     GetFolderFieldsEnum = "externalKey"
	GetFolderFieldsTimecreated     GetFolderFieldsEnum = "timeCreated"
	GetFolderFieldsTimeupdated     GetFolderFieldsEnum = "timeUpdated"
	GetFolderFieldsCreatedbyid     GetFolderFieldsEnum = "createdById"
	GetFolderFieldsUpdatedbyid     GetFolderFieldsEnum = "updatedById"
	GetFolderFieldsTimeexternal    GetFolderFieldsEnum = "timeExternal"
	GetFolderFieldsLifecyclestate  GetFolderFieldsEnum = "lifecycleState"
	GetFolderFieldsHarveststatus   GetFolderFieldsEnum = "harvestStatus"
	GetFolderFieldsLastjobkey      GetFolderFieldsEnum = "lastJobKey"
	GetFolderFieldsUri             GetFolderFieldsEnum = "uri"
)

var mappingGetFolderFields = map[string]GetFolderFieldsEnum{
	"key":             GetFolderFieldsKey,
	"displayName":     GetFolderFieldsDisplayname,
	"description":     GetFolderFieldsDescription,
	"parentFolderKey": GetFolderFieldsParentfolderkey,
	"path":            GetFolderFieldsPath,
	"dataAssetKey":    GetFolderFieldsDataassetkey,
	"properties":      GetFolderFieldsProperties,
	"externalKey":     GetFolderFieldsExternalkey,
	"timeCreated":     GetFolderFieldsTimecreated,
	"timeUpdated":     GetFolderFieldsTimeupdated,
	"createdById":     GetFolderFieldsCreatedbyid,
	"updatedById":     GetFolderFieldsUpdatedbyid,
	"timeExternal":    GetFolderFieldsTimeexternal,
	"lifecycleState":  GetFolderFieldsLifecyclestate,
	"harvestStatus":   GetFolderFieldsHarveststatus,
	"lastJobKey":      GetFolderFieldsLastjobkey,
	"uri":             GetFolderFieldsUri,
}

// GetGetFolderFieldsEnumValues Enumerates the set of values for GetFolderFieldsEnum
func GetGetFolderFieldsEnumValues() []GetFolderFieldsEnum {
	values := make([]GetFolderFieldsEnum, 0)
	for _, v := range mappingGetFolderFields {
		values = append(values, v)
	}
	return values
}
