// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetGlossaryRequest wrapper for the GetGlossary operation
type GetGlossaryRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique glossary key.
	GlossaryKey *string `mandatory:"true" contributesTo:"path" name:"glossaryKey"`

	// Specifies the fields to return in a glossary response.
	Fields []GetGlossaryFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetGlossaryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetGlossaryRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetGlossaryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetGlossaryResponse wrapper for the GetGlossary operation
type GetGlossaryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Glossary instance
	Glossary `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetGlossaryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetGlossaryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetGlossaryFieldsEnum Enum with underlying type: string
type GetGlossaryFieldsEnum string

// Set of constants representing the allowable values for GetGlossaryFieldsEnum
const (
	GetGlossaryFieldsKey            GetGlossaryFieldsEnum = "key"
	GetGlossaryFieldsDisplayname    GetGlossaryFieldsEnum = "displayName"
	GetGlossaryFieldsDescription    GetGlossaryFieldsEnum = "description"
	GetGlossaryFieldsCatalogid      GetGlossaryFieldsEnum = "catalogId"
	GetGlossaryFieldsLifecyclestate GetGlossaryFieldsEnum = "lifecycleState"
	GetGlossaryFieldsTimecreated    GetGlossaryFieldsEnum = "timeCreated"
	GetGlossaryFieldsTimeupdated    GetGlossaryFieldsEnum = "timeUpdated"
	GetGlossaryFieldsCreatedbyid    GetGlossaryFieldsEnum = "createdById"
	GetGlossaryFieldsUpdatedbyid    GetGlossaryFieldsEnum = "updatedById"
	GetGlossaryFieldsOwner          GetGlossaryFieldsEnum = "owner"
	GetGlossaryFieldsWorkflowstatus GetGlossaryFieldsEnum = "workflowStatus"
	GetGlossaryFieldsUri            GetGlossaryFieldsEnum = "uri"
)

var mappingGetGlossaryFields = map[string]GetGlossaryFieldsEnum{
	"key":            GetGlossaryFieldsKey,
	"displayName":    GetGlossaryFieldsDisplayname,
	"description":    GetGlossaryFieldsDescription,
	"catalogId":      GetGlossaryFieldsCatalogid,
	"lifecycleState": GetGlossaryFieldsLifecyclestate,
	"timeCreated":    GetGlossaryFieldsTimecreated,
	"timeUpdated":    GetGlossaryFieldsTimeupdated,
	"createdById":    GetGlossaryFieldsCreatedbyid,
	"updatedById":    GetGlossaryFieldsUpdatedbyid,
	"owner":          GetGlossaryFieldsOwner,
	"workflowStatus": GetGlossaryFieldsWorkflowstatus,
	"uri":            GetGlossaryFieldsUri,
}

// GetGetGlossaryFieldsEnumValues Enumerates the set of values for GetGlossaryFieldsEnum
func GetGetGlossaryFieldsEnumValues() []GetGlossaryFieldsEnum {
	values := make([]GetGlossaryFieldsEnum, 0)
	for _, v := range mappingGetGlossaryFields {
		values = append(values, v)
	}
	return values
}
