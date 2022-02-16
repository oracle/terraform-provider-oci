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

// GetFolderRequest wrapper for the GetFolder operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetFolder.go.html to see an example of how to use GetFolderRequest.
type GetFolderRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique folder key.
	FolderKey *string `mandatory:"true" contributesTo:"path" name:"folderKey"`

	// Indicates whether the list of objects and their relationships to this object will be provided in the response.
	IsIncludeObjectRelationships *bool `mandatory:"false" contributesTo:"query" name:"isIncludeObjectRelationships"`

	// Specifies the fields to return in a folder response.
	Fields []GetFolderFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetFolderRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetFolderRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetFolderRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetFolderRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetFolderRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetFolderFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetFolderFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetFolderResponse wrapper for the GetFolder operation
type GetFolderResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Folder instance
	Folder `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetFolderResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetFolderResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetFolderFieldsEnum Enum with underlying type: string
type GetFolderFieldsEnum string

// Set of constants representing the allowable values for GetFolderFieldsEnum
const (
	GetFolderFieldsKey             GetFolderFieldsEnum = "key"
	GetFolderFieldsDisplayname     GetFolderFieldsEnum = "displayName"
	GetFolderFieldsDescription     GetFolderFieldsEnum = "description"
	GetFolderFieldsParentfolderkey GetFolderFieldsEnum = "parentFolderKey"
	GetFolderFieldsPath            GetFolderFieldsEnum = "path"
	GetFolderFieldsDataassetkey    GetFolderFieldsEnum = "dataAssetKey"
	GetFolderFieldsProperties      GetFolderFieldsEnum = "properties"
	GetFolderFieldsExternalkey     GetFolderFieldsEnum = "externalKey"
	GetFolderFieldsTimecreated     GetFolderFieldsEnum = "timeCreated"
	GetFolderFieldsTimeupdated     GetFolderFieldsEnum = "timeUpdated"
	GetFolderFieldsCreatedbyid     GetFolderFieldsEnum = "createdById"
	GetFolderFieldsUpdatedbyid     GetFolderFieldsEnum = "updatedById"
	GetFolderFieldsTimeexternal    GetFolderFieldsEnum = "timeExternal"
	GetFolderFieldsLifecyclestate  GetFolderFieldsEnum = "lifecycleState"
	GetFolderFieldsHarveststatus   GetFolderFieldsEnum = "harvestStatus"
	GetFolderFieldsLastjobkey      GetFolderFieldsEnum = "lastJobKey"
	GetFolderFieldsUri             GetFolderFieldsEnum = "uri"
)

var mappingGetFolderFieldsEnum = map[string]GetFolderFieldsEnum{
	"key":             GetFolderFieldsKey,
	"displayName":     GetFolderFieldsDisplayname,
	"description":     GetFolderFieldsDescription,
	"parentFolderKey": GetFolderFieldsParentfolderkey,
	"path":            GetFolderFieldsPath,
	"dataAssetKey":    GetFolderFieldsDataassetkey,
	"properties":      GetFolderFieldsProperties,
	"externalKey":     GetFolderFieldsExternalkey,
	"timeCreated":     GetFolderFieldsTimecreated,
	"timeUpdated":     GetFolderFieldsTimeupdated,
	"createdById":     GetFolderFieldsCreatedbyid,
	"updatedById":     GetFolderFieldsUpdatedbyid,
	"timeExternal":    GetFolderFieldsTimeexternal,
	"lifecycleState":  GetFolderFieldsLifecyclestate,
	"harvestStatus":   GetFolderFieldsHarveststatus,
	"lastJobKey":      GetFolderFieldsLastjobkey,
	"uri":             GetFolderFieldsUri,
}

// GetGetFolderFieldsEnumValues Enumerates the set of values for GetFolderFieldsEnum
func GetGetFolderFieldsEnumValues() []GetFolderFieldsEnum {
	values := make([]GetFolderFieldsEnum, 0)
	for _, v := range mappingGetFolderFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetFolderFieldsEnumStringValues Enumerates the set of values in String for GetFolderFieldsEnum
func GetGetFolderFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"parentFolderKey",
		"path",
		"dataAssetKey",
		"properties",
		"externalKey",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"timeExternal",
		"lifecycleState",
		"harvestStatus",
		"lastJobKey",
		"uri",
	}
}

// GetMappingGetFolderFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetFolderFieldsEnum(val string) (GetFolderFieldsEnum, bool) {
	mappingGetFolderFieldsEnumIgnoreCase := make(map[string]GetFolderFieldsEnum)
	for k, v := range mappingGetFolderFieldsEnum {
		mappingGetFolderFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetFolderFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
