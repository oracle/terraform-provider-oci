// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetDataAssetTagRequest wrapper for the GetDataAssetTag operation
type GetDataAssetTagRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique tag key.
	TagKey *string `mandatory:"true" contributesTo:"path" name:"tagKey"`

	// Specifies the fields to return in a data asset tag response.
	Fields []GetDataAssetTagFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDataAssetTagRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDataAssetTagRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDataAssetTagRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetDataAssetTagResponse wrapper for the GetDataAssetTag operation
type GetDataAssetTagResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DataAssetTag instance
	DataAssetTag `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDataAssetTagResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDataAssetTagResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDataAssetTagFieldsEnum Enum with underlying type: string
type GetDataAssetTagFieldsEnum string

// Set of constants representing the allowable values for GetDataAssetTagFieldsEnum
const (
	GetDataAssetTagFieldsKey             GetDataAssetTagFieldsEnum = "key"
	GetDataAssetTagFieldsName            GetDataAssetTagFieldsEnum = "name"
	GetDataAssetTagFieldsTermkey         GetDataAssetTagFieldsEnum = "termKey"
	GetDataAssetTagFieldsTermpath        GetDataAssetTagFieldsEnum = "termPath"
	GetDataAssetTagFieldsTermdescription GetDataAssetTagFieldsEnum = "termDescription"
	GetDataAssetTagFieldsLifecyclestate  GetDataAssetTagFieldsEnum = "lifecycleState"
	GetDataAssetTagFieldsTimecreated     GetDataAssetTagFieldsEnum = "timeCreated"
	GetDataAssetTagFieldsCreatedbyid     GetDataAssetTagFieldsEnum = "createdById"
	GetDataAssetTagFieldsUri             GetDataAssetTagFieldsEnum = "uri"
	GetDataAssetTagFieldsDataassetkey    GetDataAssetTagFieldsEnum = "dataAssetKey"
)

var mappingGetDataAssetTagFields = map[string]GetDataAssetTagFieldsEnum{
	"key":             GetDataAssetTagFieldsKey,
	"name":            GetDataAssetTagFieldsName,
	"termKey":         GetDataAssetTagFieldsTermkey,
	"termPath":        GetDataAssetTagFieldsTermpath,
	"termDescription": GetDataAssetTagFieldsTermdescription,
	"lifecycleState":  GetDataAssetTagFieldsLifecyclestate,
	"timeCreated":     GetDataAssetTagFieldsTimecreated,
	"createdById":     GetDataAssetTagFieldsCreatedbyid,
	"uri":             GetDataAssetTagFieldsUri,
	"dataAssetKey":    GetDataAssetTagFieldsDataassetkey,
}

// GetGetDataAssetTagFieldsEnumValues Enumerates the set of values for GetDataAssetTagFieldsEnum
func GetGetDataAssetTagFieldsEnumValues() []GetDataAssetTagFieldsEnum {
	values := make([]GetDataAssetTagFieldsEnum, 0)
	for _, v := range mappingGetDataAssetTagFields {
		values = append(values, v)
	}
	return values
}
