// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetAttributeRequest wrapper for the GetAttribute operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetAttribute.go.html to see an example of how to use GetAttributeRequest.
type GetAttributeRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Unique attribute key.
	AttributeKey *string `mandatory:"true" contributesTo:"path" name:"attributeKey"`

	// Indicates whether the list of objects and their relationships to this object will be provided in the response.
	IsIncludeObjectRelationships *bool `mandatory:"false" contributesTo:"query" name:"isIncludeObjectRelationships"`

	// Specifies the fields to return in an entity attribute response.
	Fields []GetAttributeFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAttributeRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAttributeRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAttributeRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAttributeRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAttributeRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetAttributeFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetAttributeFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAttributeResponse wrapper for the GetAttribute operation
type GetAttributeResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Attribute instance
	Attribute `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetAttributeResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAttributeResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAttributeFieldsEnum Enum with underlying type: string
type GetAttributeFieldsEnum string

// Set of constants representing the allowable values for GetAttributeFieldsEnum
const (
	GetAttributeFieldsKey                        GetAttributeFieldsEnum = "key"
	GetAttributeFieldsDisplayname                GetAttributeFieldsEnum = "displayName"
	GetAttributeFieldsDescription                GetAttributeFieldsEnum = "description"
	GetAttributeFieldsEntitykey                  GetAttributeFieldsEnum = "entityKey"
	GetAttributeFieldsLifecyclestate             GetAttributeFieldsEnum = "lifecycleState"
	GetAttributeFieldsTimecreated                GetAttributeFieldsEnum = "timeCreated"
	GetAttributeFieldsTimeupdated                GetAttributeFieldsEnum = "timeUpdated"
	GetAttributeFieldsCreatedbyid                GetAttributeFieldsEnum = "createdById"
	GetAttributeFieldsUpdatedbyid                GetAttributeFieldsEnum = "updatedById"
	GetAttributeFieldsExternaldatatype           GetAttributeFieldsEnum = "externalDataType"
	GetAttributeFieldsExternalkey                GetAttributeFieldsEnum = "externalKey"
	GetAttributeFieldsIsincrementaldata          GetAttributeFieldsEnum = "isIncrementalData"
	GetAttributeFieldsIsnullable                 GetAttributeFieldsEnum = "isNullable"
	GetAttributeFieldsLength                     GetAttributeFieldsEnum = "length"
	GetAttributeFieldsPosition                   GetAttributeFieldsEnum = "position"
	GetAttributeFieldsPrecision                  GetAttributeFieldsEnum = "precision"
	GetAttributeFieldsScale                      GetAttributeFieldsEnum = "scale"
	GetAttributeFieldsTimeexternal               GetAttributeFieldsEnum = "timeExternal"
	GetAttributeFieldsUri                        GetAttributeFieldsEnum = "uri"
	GetAttributeFieldsProperties                 GetAttributeFieldsEnum = "properties"
	GetAttributeFieldsPath                       GetAttributeFieldsEnum = "path"
	GetAttributeFieldsMincollectioncount         GetAttributeFieldsEnum = "minCollectionCount"
	GetAttributeFieldsMaxcollectioncount         GetAttributeFieldsEnum = "maxCollectionCount"
	GetAttributeFieldsDatatypeentitykey          GetAttributeFieldsEnum = "datatypeEntityKey"
	GetAttributeFieldsExternaldatatypeentitykey  GetAttributeFieldsEnum = "externalDatatypeEntityKey"
	GetAttributeFieldsParentattributekey         GetAttributeFieldsEnum = "parentAttributeKey"
	GetAttributeFieldsExternalparentattributekey GetAttributeFieldsEnum = "externalParentAttributeKey"
	GetAttributeFieldsTypekey                    GetAttributeFieldsEnum = "typeKey"
)

var mappingGetAttributeFieldsEnum = map[string]GetAttributeFieldsEnum{
	"key":                        GetAttributeFieldsKey,
	"displayName":                GetAttributeFieldsDisplayname,
	"description":                GetAttributeFieldsDescription,
	"entityKey":                  GetAttributeFieldsEntitykey,
	"lifecycleState":             GetAttributeFieldsLifecyclestate,
	"timeCreated":                GetAttributeFieldsTimecreated,
	"timeUpdated":                GetAttributeFieldsTimeupdated,
	"createdById":                GetAttributeFieldsCreatedbyid,
	"updatedById":                GetAttributeFieldsUpdatedbyid,
	"externalDataType":           GetAttributeFieldsExternaldatatype,
	"externalKey":                GetAttributeFieldsExternalkey,
	"isIncrementalData":          GetAttributeFieldsIsincrementaldata,
	"isNullable":                 GetAttributeFieldsIsnullable,
	"length":                     GetAttributeFieldsLength,
	"position":                   GetAttributeFieldsPosition,
	"precision":                  GetAttributeFieldsPrecision,
	"scale":                      GetAttributeFieldsScale,
	"timeExternal":               GetAttributeFieldsTimeexternal,
	"uri":                        GetAttributeFieldsUri,
	"properties":                 GetAttributeFieldsProperties,
	"path":                       GetAttributeFieldsPath,
	"minCollectionCount":         GetAttributeFieldsMincollectioncount,
	"maxCollectionCount":         GetAttributeFieldsMaxcollectioncount,
	"datatypeEntityKey":          GetAttributeFieldsDatatypeentitykey,
	"externalDatatypeEntityKey":  GetAttributeFieldsExternaldatatypeentitykey,
	"parentAttributeKey":         GetAttributeFieldsParentattributekey,
	"externalParentAttributeKey": GetAttributeFieldsExternalparentattributekey,
	"typeKey":                    GetAttributeFieldsTypekey,
}

var mappingGetAttributeFieldsEnumLowerCase = map[string]GetAttributeFieldsEnum{
	"key":                        GetAttributeFieldsKey,
	"displayname":                GetAttributeFieldsDisplayname,
	"description":                GetAttributeFieldsDescription,
	"entitykey":                  GetAttributeFieldsEntitykey,
	"lifecyclestate":             GetAttributeFieldsLifecyclestate,
	"timecreated":                GetAttributeFieldsTimecreated,
	"timeupdated":                GetAttributeFieldsTimeupdated,
	"createdbyid":                GetAttributeFieldsCreatedbyid,
	"updatedbyid":                GetAttributeFieldsUpdatedbyid,
	"externaldatatype":           GetAttributeFieldsExternaldatatype,
	"externalkey":                GetAttributeFieldsExternalkey,
	"isincrementaldata":          GetAttributeFieldsIsincrementaldata,
	"isnullable":                 GetAttributeFieldsIsnullable,
	"length":                     GetAttributeFieldsLength,
	"position":                   GetAttributeFieldsPosition,
	"precision":                  GetAttributeFieldsPrecision,
	"scale":                      GetAttributeFieldsScale,
	"timeexternal":               GetAttributeFieldsTimeexternal,
	"uri":                        GetAttributeFieldsUri,
	"properties":                 GetAttributeFieldsProperties,
	"path":                       GetAttributeFieldsPath,
	"mincollectioncount":         GetAttributeFieldsMincollectioncount,
	"maxcollectioncount":         GetAttributeFieldsMaxcollectioncount,
	"datatypeentitykey":          GetAttributeFieldsDatatypeentitykey,
	"externaldatatypeentitykey":  GetAttributeFieldsExternaldatatypeentitykey,
	"parentattributekey":         GetAttributeFieldsParentattributekey,
	"externalparentattributekey": GetAttributeFieldsExternalparentattributekey,
	"typekey":                    GetAttributeFieldsTypekey,
}

// GetGetAttributeFieldsEnumValues Enumerates the set of values for GetAttributeFieldsEnum
func GetGetAttributeFieldsEnumValues() []GetAttributeFieldsEnum {
	values := make([]GetAttributeFieldsEnum, 0)
	for _, v := range mappingGetAttributeFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAttributeFieldsEnumStringValues Enumerates the set of values in String for GetAttributeFieldsEnum
func GetGetAttributeFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"entityKey",
		"lifecycleState",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"externalDataType",
		"externalKey",
		"isIncrementalData",
		"isNullable",
		"length",
		"position",
		"precision",
		"scale",
		"timeExternal",
		"uri",
		"properties",
		"path",
		"minCollectionCount",
		"maxCollectionCount",
		"datatypeEntityKey",
		"externalDatatypeEntityKey",
		"parentAttributeKey",
		"externalParentAttributeKey",
		"typeKey",
	}
}

// GetMappingGetAttributeFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAttributeFieldsEnum(val string) (GetAttributeFieldsEnum, bool) {
	enum, ok := mappingGetAttributeFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
