// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetAttributeTagRequest wrapper for the GetAttributeTag operation
type GetAttributeTagRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Unique attribute key.
	AttributeKey *string `mandatory:"true" contributesTo:"path" name:"attributeKey"`

	// Unique tag key.
	TagKey *string `mandatory:"true" contributesTo:"path" name:"tagKey"`

	// Specifies the fields to return in an entity attribute tag response.
	Fields []GetAttributeTagFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAttributeTagRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAttributeTagRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAttributeTagRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetAttributeTagResponse wrapper for the GetAttributeTag operation
type GetAttributeTagResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The AttributeTag instance
	AttributeTag `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAttributeTagResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAttributeTagResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAttributeTagFieldsEnum Enum with underlying type: string
type GetAttributeTagFieldsEnum string

// Set of constants representing the allowable values for GetAttributeTagFieldsEnum
const (
	GetAttributeTagFieldsKey             GetAttributeTagFieldsEnum = "key"
	GetAttributeTagFieldsName            GetAttributeTagFieldsEnum = "name"
	GetAttributeTagFieldsTermkey         GetAttributeTagFieldsEnum = "termKey"
	GetAttributeTagFieldsTermpath        GetAttributeTagFieldsEnum = "termPath"
	GetAttributeTagFieldsTermdescription GetAttributeTagFieldsEnum = "termDescription"
	GetAttributeTagFieldsLifecyclestate  GetAttributeTagFieldsEnum = "lifecycleState"
	GetAttributeTagFieldsTimecreated     GetAttributeTagFieldsEnum = "timeCreated"
	GetAttributeTagFieldsCreatedbyid     GetAttributeTagFieldsEnum = "createdById"
	GetAttributeTagFieldsUri             GetAttributeTagFieldsEnum = "uri"
	GetAttributeTagFieldsAttributekey    GetAttributeTagFieldsEnum = "attributeKey"
)

var mappingGetAttributeTagFields = map[string]GetAttributeTagFieldsEnum{
	"key":             GetAttributeTagFieldsKey,
	"name":            GetAttributeTagFieldsName,
	"termKey":         GetAttributeTagFieldsTermkey,
	"termPath":        GetAttributeTagFieldsTermpath,
	"termDescription": GetAttributeTagFieldsTermdescription,
	"lifecycleState":  GetAttributeTagFieldsLifecyclestate,
	"timeCreated":     GetAttributeTagFieldsTimecreated,
	"createdById":     GetAttributeTagFieldsCreatedbyid,
	"uri":             GetAttributeTagFieldsUri,
	"attributeKey":    GetAttributeTagFieldsAttributekey,
}

// GetGetAttributeTagFieldsEnumValues Enumerates the set of values for GetAttributeTagFieldsEnum
func GetGetAttributeTagFieldsEnumValues() []GetAttributeTagFieldsEnum {
	values := make([]GetAttributeTagFieldsEnum, 0)
	for _, v := range mappingGetAttributeTagFields {
		values = append(values, v)
	}
	return values
}
