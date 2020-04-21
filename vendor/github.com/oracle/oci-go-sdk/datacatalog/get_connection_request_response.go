// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetConnectionRequest wrapper for the GetConnection operation
type GetConnectionRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique connection key.
	ConnectionKey *string `mandatory:"true" contributesTo:"path" name:"connectionKey"`

	// Specifies the fields to return in a connection response.
	Fields []GetConnectionFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetConnectionRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetConnectionRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetConnectionRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetConnectionResponse wrapper for the GetConnection operation
type GetConnectionResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Connection instance
	Connection `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetConnectionResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetConnectionResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetConnectionFieldsEnum Enum with underlying type: string
type GetConnectionFieldsEnum string

// Set of constants representing the allowable values for GetConnectionFieldsEnum
const (
	GetConnectionFieldsKey               GetConnectionFieldsEnum = "key"
	GetConnectionFieldsDisplayname       GetConnectionFieldsEnum = "displayName"
	GetConnectionFieldsDescription       GetConnectionFieldsEnum = "description"
	GetConnectionFieldsDataassetkey      GetConnectionFieldsEnum = "dataAssetKey"
	GetConnectionFieldsTypekey           GetConnectionFieldsEnum = "typeKey"
	GetConnectionFieldsTimecreated       GetConnectionFieldsEnum = "timeCreated"
	GetConnectionFieldsTimeupdated       GetConnectionFieldsEnum = "timeUpdated"
	GetConnectionFieldsCreatedbyid       GetConnectionFieldsEnum = "createdById"
	GetConnectionFieldsUpdatedbyid       GetConnectionFieldsEnum = "updatedById"
	GetConnectionFieldsProperties        GetConnectionFieldsEnum = "properties"
	GetConnectionFieldsExternalkey       GetConnectionFieldsEnum = "externalKey"
	GetConnectionFieldsTimestatusupdated GetConnectionFieldsEnum = "timeStatusUpdated"
	GetConnectionFieldsLifecyclestate    GetConnectionFieldsEnum = "lifecycleState"
	GetConnectionFieldsIsdefault         GetConnectionFieldsEnum = "isDefault"
	GetConnectionFieldsUri               GetConnectionFieldsEnum = "uri"
)

var mappingGetConnectionFields = map[string]GetConnectionFieldsEnum{
	"key":               GetConnectionFieldsKey,
	"displayName":       GetConnectionFieldsDisplayname,
	"description":       GetConnectionFieldsDescription,
	"dataAssetKey":      GetConnectionFieldsDataassetkey,
	"typeKey":           GetConnectionFieldsTypekey,
	"timeCreated":       GetConnectionFieldsTimecreated,
	"timeUpdated":       GetConnectionFieldsTimeupdated,
	"createdById":       GetConnectionFieldsCreatedbyid,
	"updatedById":       GetConnectionFieldsUpdatedbyid,
	"properties":        GetConnectionFieldsProperties,
	"externalKey":       GetConnectionFieldsExternalkey,
	"timeStatusUpdated": GetConnectionFieldsTimestatusupdated,
	"lifecycleState":    GetConnectionFieldsLifecyclestate,
	"isDefault":         GetConnectionFieldsIsdefault,
	"uri":               GetConnectionFieldsUri,
}

// GetGetConnectionFieldsEnumValues Enumerates the set of values for GetConnectionFieldsEnum
func GetGetConnectionFieldsEnumValues() []GetConnectionFieldsEnum {
	values := make([]GetConnectionFieldsEnum, 0)
	for _, v := range mappingGetConnectionFields {
		values = append(values, v)
	}
	return values
}
