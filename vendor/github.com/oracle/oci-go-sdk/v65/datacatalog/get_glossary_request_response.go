// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetGlossaryRequest wrapper for the GetGlossary operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetGlossary.go.html to see an example of how to use GetGlossaryRequest.
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
func (request GetGlossaryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetGlossaryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetGlossaryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetGlossaryRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetGlossaryFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetGlossaryFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingGetGlossaryFieldsEnum = map[string]GetGlossaryFieldsEnum{
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

var mappingGetGlossaryFieldsEnumLowerCase = map[string]GetGlossaryFieldsEnum{
	"key":            GetGlossaryFieldsKey,
	"displayname":    GetGlossaryFieldsDisplayname,
	"description":    GetGlossaryFieldsDescription,
	"catalogid":      GetGlossaryFieldsCatalogid,
	"lifecyclestate": GetGlossaryFieldsLifecyclestate,
	"timecreated":    GetGlossaryFieldsTimecreated,
	"timeupdated":    GetGlossaryFieldsTimeupdated,
	"createdbyid":    GetGlossaryFieldsCreatedbyid,
	"updatedbyid":    GetGlossaryFieldsUpdatedbyid,
	"owner":          GetGlossaryFieldsOwner,
	"workflowstatus": GetGlossaryFieldsWorkflowstatus,
	"uri":            GetGlossaryFieldsUri,
}

// GetGetGlossaryFieldsEnumValues Enumerates the set of values for GetGlossaryFieldsEnum
func GetGetGlossaryFieldsEnumValues() []GetGlossaryFieldsEnum {
	values := make([]GetGlossaryFieldsEnum, 0)
	for _, v := range mappingGetGlossaryFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetGlossaryFieldsEnumStringValues Enumerates the set of values in String for GetGlossaryFieldsEnum
func GetGetGlossaryFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"catalogId",
		"lifecycleState",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"owner",
		"workflowStatus",
		"uri",
	}
}

// GetMappingGetGlossaryFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetGlossaryFieldsEnum(val string) (GetGlossaryFieldsEnum, bool) {
	enum, ok := mappingGetGlossaryFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
