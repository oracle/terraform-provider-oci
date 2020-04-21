// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDataAssetRequest wrapper for the GetDataAsset operation
type GetDataAssetRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Specifies the fields to return in a data asset response.
	Fields []GetDataAssetFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDataAssetRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDataAssetRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDataAssetRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDataAssetResponse wrapper for the GetDataAsset operation
type GetDataAssetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DataAsset instance
	DataAsset `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDataAssetResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDataAssetResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDataAssetFieldsEnum Enum with underlying type: string
type GetDataAssetFieldsEnum string

// Set of constants representing the allowable values for GetDataAssetFieldsEnum
const (
	GetDataAssetFieldsKey            GetDataAssetFieldsEnum = "key"
	GetDataAssetFieldsDisplayname    GetDataAssetFieldsEnum = "displayName"
	GetDataAssetFieldsDescription    GetDataAssetFieldsEnum = "description"
	GetDataAssetFieldsCatalogid      GetDataAssetFieldsEnum = "catalogId"
	GetDataAssetFieldsExternalkey    GetDataAssetFieldsEnum = "externalKey"
	GetDataAssetFieldsTypekey        GetDataAssetFieldsEnum = "typeKey"
	GetDataAssetFieldsLifecyclestate GetDataAssetFieldsEnum = "lifecycleState"
	GetDataAssetFieldsTimecreated    GetDataAssetFieldsEnum = "timeCreated"
	GetDataAssetFieldsTimeupdated    GetDataAssetFieldsEnum = "timeUpdated"
	GetDataAssetFieldsCreatedbyid    GetDataAssetFieldsEnum = "createdById"
	GetDataAssetFieldsUpdatedbyid    GetDataAssetFieldsEnum = "updatedById"
	GetDataAssetFieldsUri            GetDataAssetFieldsEnum = "uri"
	GetDataAssetFieldsProperties     GetDataAssetFieldsEnum = "properties"
)

var mappingGetDataAssetFields = map[string]GetDataAssetFieldsEnum{
	"key":            GetDataAssetFieldsKey,
	"displayName":    GetDataAssetFieldsDisplayname,
	"description":    GetDataAssetFieldsDescription,
	"catalogId":      GetDataAssetFieldsCatalogid,
	"externalKey":    GetDataAssetFieldsExternalkey,
	"typeKey":        GetDataAssetFieldsTypekey,
	"lifecycleState": GetDataAssetFieldsLifecyclestate,
	"timeCreated":    GetDataAssetFieldsTimecreated,
	"timeUpdated":    GetDataAssetFieldsTimeupdated,
	"createdById":    GetDataAssetFieldsCreatedbyid,
	"updatedById":    GetDataAssetFieldsUpdatedbyid,
	"uri":            GetDataAssetFieldsUri,
	"properties":     GetDataAssetFieldsProperties,
}

// GetGetDataAssetFieldsEnumValues Enumerates the set of values for GetDataAssetFieldsEnum
func GetGetDataAssetFieldsEnumValues() []GetDataAssetFieldsEnum {
	values := make([]GetDataAssetFieldsEnum, 0)
	for _, v := range mappingGetDataAssetFields {
		values = append(values, v)
	}
	return values
}
