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

// GetTermRequest wrapper for the GetTerm operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetTerm.go.html to see an example of how to use GetTermRequest.
type GetTermRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique glossary key.
	GlossaryKey *string `mandatory:"true" contributesTo:"path" name:"glossaryKey"`

	// Unique glossary term key.
	TermKey *string `mandatory:"true" contributesTo:"path" name:"termKey"`

	// Specifies the fields to return in a term response.
	Fields []GetTermFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetTermRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetTermRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetTermRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetTermRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetTermRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetTermFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetTermFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetTermResponse wrapper for the GetTerm operation
type GetTermResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Term instance
	Term `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetTermResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetTermResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetTermFieldsEnum Enum with underlying type: string
type GetTermFieldsEnum string

// Set of constants representing the allowable values for GetTermFieldsEnum
const (
	GetTermFieldsKey                       GetTermFieldsEnum = "key"
	GetTermFieldsDisplayname               GetTermFieldsEnum = "displayName"
	GetTermFieldsDescription               GetTermFieldsEnum = "description"
	GetTermFieldsGlossarykey               GetTermFieldsEnum = "glossaryKey"
	GetTermFieldsParenttermkey             GetTermFieldsEnum = "parentTermKey"
	GetTermFieldsIsallowedtohavechildterms GetTermFieldsEnum = "isAllowedToHaveChildTerms"
	GetTermFieldsPath                      GetTermFieldsEnum = "path"
	GetTermFieldsLifecyclestate            GetTermFieldsEnum = "lifecycleState"
	GetTermFieldsTimecreated               GetTermFieldsEnum = "timeCreated"
	GetTermFieldsTimeupdated               GetTermFieldsEnum = "timeUpdated"
	GetTermFieldsCreatedbyid               GetTermFieldsEnum = "createdById"
	GetTermFieldsUpdatedbyid               GetTermFieldsEnum = "updatedById"
	GetTermFieldsOwner                     GetTermFieldsEnum = "owner"
	GetTermFieldsWorkflowstatus            GetTermFieldsEnum = "workflowStatus"
	GetTermFieldsUri                       GetTermFieldsEnum = "uri"
	GetTermFieldsRelatedterms              GetTermFieldsEnum = "relatedTerms"
	GetTermFieldsAssociatedobjectcount     GetTermFieldsEnum = "associatedObjectCount"
	GetTermFieldsAssociatedobjects         GetTermFieldsEnum = "associatedObjects"
)

var mappingGetTermFieldsEnum = map[string]GetTermFieldsEnum{
	"key":                       GetTermFieldsKey,
	"displayName":               GetTermFieldsDisplayname,
	"description":               GetTermFieldsDescription,
	"glossaryKey":               GetTermFieldsGlossarykey,
	"parentTermKey":             GetTermFieldsParenttermkey,
	"isAllowedToHaveChildTerms": GetTermFieldsIsallowedtohavechildterms,
	"path":                      GetTermFieldsPath,
	"lifecycleState":            GetTermFieldsLifecyclestate,
	"timeCreated":               GetTermFieldsTimecreated,
	"timeUpdated":               GetTermFieldsTimeupdated,
	"createdById":               GetTermFieldsCreatedbyid,
	"updatedById":               GetTermFieldsUpdatedbyid,
	"owner":                     GetTermFieldsOwner,
	"workflowStatus":            GetTermFieldsWorkflowstatus,
	"uri":                       GetTermFieldsUri,
	"relatedTerms":              GetTermFieldsRelatedterms,
	"associatedObjectCount":     GetTermFieldsAssociatedobjectcount,
	"associatedObjects":         GetTermFieldsAssociatedobjects,
}

var mappingGetTermFieldsEnumLowerCase = map[string]GetTermFieldsEnum{
	"key":                       GetTermFieldsKey,
	"displayname":               GetTermFieldsDisplayname,
	"description":               GetTermFieldsDescription,
	"glossarykey":               GetTermFieldsGlossarykey,
	"parenttermkey":             GetTermFieldsParenttermkey,
	"isallowedtohavechildterms": GetTermFieldsIsallowedtohavechildterms,
	"path":                      GetTermFieldsPath,
	"lifecyclestate":            GetTermFieldsLifecyclestate,
	"timecreated":               GetTermFieldsTimecreated,
	"timeupdated":               GetTermFieldsTimeupdated,
	"createdbyid":               GetTermFieldsCreatedbyid,
	"updatedbyid":               GetTermFieldsUpdatedbyid,
	"owner":                     GetTermFieldsOwner,
	"workflowstatus":            GetTermFieldsWorkflowstatus,
	"uri":                       GetTermFieldsUri,
	"relatedterms":              GetTermFieldsRelatedterms,
	"associatedobjectcount":     GetTermFieldsAssociatedobjectcount,
	"associatedobjects":         GetTermFieldsAssociatedobjects,
}

// GetGetTermFieldsEnumValues Enumerates the set of values for GetTermFieldsEnum
func GetGetTermFieldsEnumValues() []GetTermFieldsEnum {
	values := make([]GetTermFieldsEnum, 0)
	for _, v := range mappingGetTermFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetTermFieldsEnumStringValues Enumerates the set of values in String for GetTermFieldsEnum
func GetGetTermFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"glossaryKey",
		"parentTermKey",
		"isAllowedToHaveChildTerms",
		"path",
		"lifecycleState",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"owner",
		"workflowStatus",
		"uri",
		"relatedTerms",
		"associatedObjectCount",
		"associatedObjects",
	}
}

// GetMappingGetTermFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetTermFieldsEnum(val string) (GetTermFieldsEnum, bool) {
	enum, ok := mappingGetTermFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
