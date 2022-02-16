// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// GetEntityRequest wrapper for the GetEntity operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetEntity.go.html to see an example of how to use GetEntityRequest.
type GetEntityRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique entity key.
	EntityKey *string `mandatory:"true" contributesTo:"path" name:"entityKey"`

	// Indicates whether the list of objects and their relationships to this object will be provided in the response.
	IsIncludeObjectRelationships *bool `mandatory:"false" contributesTo:"query" name:"isIncludeObjectRelationships"`

	// Specifies the fields to return in an entity response.
	Fields []GetEntityFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetEntityRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetEntityRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetEntityRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetEntityRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetEntityRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetEntityFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetEntityFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetEntityResponse wrapper for the GetEntity operation
type GetEntityResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Entity instance
	Entity `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetEntityResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetEntityResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetEntityFieldsEnum Enum with underlying type: string
type GetEntityFieldsEnum string

// Set of constants representing the allowable values for GetEntityFieldsEnum
const (
	GetEntityFieldsKey               GetEntityFieldsEnum = "key"
	GetEntityFieldsDisplayname       GetEntityFieldsEnum = "displayName"
	GetEntityFieldsDescription       GetEntityFieldsEnum = "description"
	GetEntityFieldsDataassetkey      GetEntityFieldsEnum = "dataAssetKey"
	GetEntityFieldsTimecreated       GetEntityFieldsEnum = "timeCreated"
	GetEntityFieldsTimeupdated       GetEntityFieldsEnum = "timeUpdated"
	GetEntityFieldsCreatedbyid       GetEntityFieldsEnum = "createdById"
	GetEntityFieldsUpdatedbyid       GetEntityFieldsEnum = "updatedById"
	GetEntityFieldsLifecyclestate    GetEntityFieldsEnum = "lifecycleState"
	GetEntityFieldsExternalkey       GetEntityFieldsEnum = "externalKey"
	GetEntityFieldsTimeexternal      GetEntityFieldsEnum = "timeExternal"
	GetEntityFieldsTimestatusupdated GetEntityFieldsEnum = "timeStatusUpdated"
	GetEntityFieldsIslogical         GetEntityFieldsEnum = "isLogical"
	GetEntityFieldsIspartition       GetEntityFieldsEnum = "isPartition"
	GetEntityFieldsFolderkey         GetEntityFieldsEnum = "folderKey"
	GetEntityFieldsFoldername        GetEntityFieldsEnum = "folderName"
	GetEntityFieldsTypekey           GetEntityFieldsEnum = "typeKey"
	GetEntityFieldsPath              GetEntityFieldsEnum = "path"
	GetEntityFieldsHarveststatus     GetEntityFieldsEnum = "harvestStatus"
	GetEntityFieldsLastjobkey        GetEntityFieldsEnum = "lastJobKey"
	GetEntityFieldsUri               GetEntityFieldsEnum = "uri"
	GetEntityFieldsProperties        GetEntityFieldsEnum = "properties"
)

var mappingGetEntityFieldsEnum = map[string]GetEntityFieldsEnum{
	"key":               GetEntityFieldsKey,
	"displayName":       GetEntityFieldsDisplayname,
	"description":       GetEntityFieldsDescription,
	"dataAssetKey":      GetEntityFieldsDataassetkey,
	"timeCreated":       GetEntityFieldsTimecreated,
	"timeUpdated":       GetEntityFieldsTimeupdated,
	"createdById":       GetEntityFieldsCreatedbyid,
	"updatedById":       GetEntityFieldsUpdatedbyid,
	"lifecycleState":    GetEntityFieldsLifecyclestate,
	"externalKey":       GetEntityFieldsExternalkey,
	"timeExternal":      GetEntityFieldsTimeexternal,
	"timeStatusUpdated": GetEntityFieldsTimestatusupdated,
	"isLogical":         GetEntityFieldsIslogical,
	"isPartition":       GetEntityFieldsIspartition,
	"folderKey":         GetEntityFieldsFolderkey,
	"folderName":        GetEntityFieldsFoldername,
	"typeKey":           GetEntityFieldsTypekey,
	"path":              GetEntityFieldsPath,
	"harvestStatus":     GetEntityFieldsHarveststatus,
	"lastJobKey":        GetEntityFieldsLastjobkey,
	"uri":               GetEntityFieldsUri,
	"properties":        GetEntityFieldsProperties,
}

// GetGetEntityFieldsEnumValues Enumerates the set of values for GetEntityFieldsEnum
func GetGetEntityFieldsEnumValues() []GetEntityFieldsEnum {
	values := make([]GetEntityFieldsEnum, 0)
	for _, v := range mappingGetEntityFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetEntityFieldsEnumStringValues Enumerates the set of values in String for GetEntityFieldsEnum
func GetGetEntityFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"dataAssetKey",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"lifecycleState",
		"externalKey",
		"timeExternal",
		"timeStatusUpdated",
		"isLogical",
		"isPartition",
		"folderKey",
		"folderName",
		"typeKey",
		"path",
		"harvestStatus",
		"lastJobKey",
		"uri",
		"properties",
	}
}

// GetMappingGetEntityFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetEntityFieldsEnum(val string) (GetEntityFieldsEnum, bool) {
	mappingGetEntityFieldsEnumIgnoreCase := make(map[string]GetEntityFieldsEnum)
	for k, v := range mappingGetEntityFieldsEnum {
		mappingGetEntityFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetEntityFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
