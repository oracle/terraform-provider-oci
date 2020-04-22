// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetEntityTagRequest wrapper for the GetEntityTag operation
type GetEntityTagRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Unique tag key.
	TagKey *string `mandatory:"true" contributesTo:"path" name:"tagKey"`

	// Specifies the fields to return in an entity tag response.
	Fields []GetEntityTagFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetEntityTagRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetEntityTagRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetEntityTagRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetEntityTagResponse wrapper for the GetEntityTag operation
type GetEntityTagResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The EntityTag instance
	EntityTag `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetEntityTagResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetEntityTagResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetEntityTagFieldsEnum Enum with underlying type: string
type GetEntityTagFieldsEnum string

// Set of constants representing the allowable values for GetEntityTagFieldsEnum
const (
	GetEntityTagFieldsKey             GetEntityTagFieldsEnum = "key"
	GetEntityTagFieldsName            GetEntityTagFieldsEnum = "name"
	GetEntityTagFieldsTermkey         GetEntityTagFieldsEnum = "termKey"
	GetEntityTagFieldsTermpath        GetEntityTagFieldsEnum = "termPath"
	GetEntityTagFieldsTermdescription GetEntityTagFieldsEnum = "termDescription"
	GetEntityTagFieldsLifecyclestate  GetEntityTagFieldsEnum = "lifecycleState"
	GetEntityTagFieldsTimecreated     GetEntityTagFieldsEnum = "timeCreated"
	GetEntityTagFieldsCreatedbyid     GetEntityTagFieldsEnum = "createdById"
	GetEntityTagFieldsUri             GetEntityTagFieldsEnum = "uri"
	GetEntityTagFieldsEntitykey       GetEntityTagFieldsEnum = "entityKey"
)

var mappingGetEntityTagFields = map[string]GetEntityTagFieldsEnum{
	"key":             GetEntityTagFieldsKey,
	"name":            GetEntityTagFieldsName,
	"termKey":         GetEntityTagFieldsTermkey,
	"termPath":        GetEntityTagFieldsTermpath,
	"termDescription": GetEntityTagFieldsTermdescription,
	"lifecycleState":  GetEntityTagFieldsLifecyclestate,
	"timeCreated":     GetEntityTagFieldsTimecreated,
	"createdById":     GetEntityTagFieldsCreatedbyid,
	"uri":             GetEntityTagFieldsUri,
	"entityKey":       GetEntityTagFieldsEntitykey,
}

// GetGetEntityTagFieldsEnumValues Enumerates the set of values for GetEntityTagFieldsEnum
func GetGetEntityTagFieldsEnumValues() []GetEntityTagFieldsEnum {
	values := make([]GetEntityTagFieldsEnum, 0)
	for _, v := range mappingGetEntityTagFields {
		values = append(values, v)
	}
	return values
}
