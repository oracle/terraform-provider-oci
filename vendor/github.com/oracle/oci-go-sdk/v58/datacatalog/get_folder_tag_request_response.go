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

// GetFolderTagRequest wrapper for the GetFolderTag operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetFolderTag.go.html to see an example of how to use GetFolderTagRequest.
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
func (request GetFolderTagRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetFolderTagRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetFolderTagRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetFolderTagRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetFolderTagFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetFolderTagFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingGetFolderTagFieldsEnum = map[string]GetFolderTagFieldsEnum{
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
	for _, v := range mappingGetFolderTagFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetFolderTagFieldsEnumStringValues Enumerates the set of values in String for GetFolderTagFieldsEnum
func GetGetFolderTagFieldsEnumStringValues() []string {
	return []string{
		"key",
		"name",
		"termKey",
		"termPath",
		"termDescription",
		"lifecycleState",
		"timeCreated",
		"createdById",
		"uri",
		"folderKey",
	}
}

// GetMappingGetFolderTagFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetFolderTagFieldsEnum(val string) (GetFolderTagFieldsEnum, bool) {
	mappingGetFolderTagFieldsEnumIgnoreCase := make(map[string]GetFolderTagFieldsEnum)
	for k, v := range mappingGetFolderTagFieldsEnum {
		mappingGetFolderTagFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetFolderTagFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
