// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetFolderTagRequest wrapper for the GetFolderTag operation
type GetFolderTagRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique folder key.
	FolderKey *string `mandatory:"true" contributesTo:"path" name:"folderKey"`

	// Unique tag key.
	TagKey *string `mandatory:"true" contributesTo:"path" name:"tagKey"`

	// Specifies the fields to return in a folder tag response.
	Fields []GetFolderTagFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetFolderTagRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetFolderTagRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetFolderTagRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetFolderTagResponse wrapper for the GetFolderTag operation
type GetFolderTagResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The FolderTag instance
	FolderTag `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetFolderTagResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetFolderTagResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetFolderTagFieldsEnum Enum with underlying type: string
type GetFolderTagFieldsEnum string

// Set of constants representing the allowable values for GetFolderTagFieldsEnum
const (
	GetFolderTagFieldsKey             GetFolderTagFieldsEnum = "key"
	GetFolderTagFieldsName            GetFolderTagFieldsEnum = "name"
	GetFolderTagFieldsTermkey         GetFolderTagFieldsEnum = "termKey"
	GetFolderTagFieldsTermpath        GetFolderTagFieldsEnum = "termPath"
	GetFolderTagFieldsTermdescription GetFolderTagFieldsEnum = "termDescription"
	GetFolderTagFieldsLifecyclestate  GetFolderTagFieldsEnum = "lifecycleState"
	GetFolderTagFieldsTimecreated     GetFolderTagFieldsEnum = "timeCreated"
	GetFolderTagFieldsCreatedbyid     GetFolderTagFieldsEnum = "createdById"
	GetFolderTagFieldsUri             GetFolderTagFieldsEnum = "uri"
	GetFolderTagFieldsFolderkey       GetFolderTagFieldsEnum = "folderKey"
)

var mappingGetFolderTagFields = map[string]GetFolderTagFieldsEnum{
	"key":             GetFolderTagFieldsKey,
	"name":            GetFolderTagFieldsName,
	"termKey":         GetFolderTagFieldsTermkey,
	"termPath":        GetFolderTagFieldsTermpath,
	"termDescription": GetFolderTagFieldsTermdescription,
	"lifecycleState":  GetFolderTagFieldsLifecyclestate,
	"timeCreated":     GetFolderTagFieldsTimecreated,
	"createdById":     GetFolderTagFieldsCreatedbyid,
	"uri":             GetFolderTagFieldsUri,
	"folderKey":       GetFolderTagFieldsFolderkey,
}

// GetGetFolderTagFieldsEnumValues Enumerates the set of values for GetFolderTagFieldsEnum
func GetGetFolderTagFieldsEnumValues() []GetFolderTagFieldsEnum {
	values := make([]GetFolderTagFieldsEnum, 0)
	for _, v := range mappingGetFolderTagFields {
		values = append(values, v)
	}
	return values
}
