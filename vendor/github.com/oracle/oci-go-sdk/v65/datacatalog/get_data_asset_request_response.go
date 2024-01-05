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

// GetDataAssetRequest wrapper for the GetDataAsset operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetDataAsset.go.html to see an example of how to use GetDataAssetRequest.
type GetDataAssetRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Specifies the fields to return in a data asset response.
	Fields []GetDataAssetFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDataAssetRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDataAssetRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDataAssetRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDataAssetRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetDataAssetRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetDataAssetFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetDataAssetFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetDataAssetResponse wrapper for the GetDataAsset operation
type GetDataAssetResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DataAsset instance
	DataAsset `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDataAssetResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDataAssetResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDataAssetFieldsEnum Enum with underlying type: string
type GetDataAssetFieldsEnum string

// Set of constants representing the allowable values for GetDataAssetFieldsEnum
const (
	GetDataAssetFieldsKey            GetDataAssetFieldsEnum = "key"
	GetDataAssetFieldsDisplayname    GetDataAssetFieldsEnum = "displayName"
	GetDataAssetFieldsDescription    GetDataAssetFieldsEnum = "description"
	GetDataAssetFieldsCatalogid      GetDataAssetFieldsEnum = "catalogId"
	GetDataAssetFieldsExternalkey    GetDataAssetFieldsEnum = "externalKey"
	GetDataAssetFieldsTypekey        GetDataAssetFieldsEnum = "typeKey"
	GetDataAssetFieldsLifecyclestate GetDataAssetFieldsEnum = "lifecycleState"
	GetDataAssetFieldsTimecreated    GetDataAssetFieldsEnum = "timeCreated"
	GetDataAssetFieldsTimeupdated    GetDataAssetFieldsEnum = "timeUpdated"
	GetDataAssetFieldsCreatedbyid    GetDataAssetFieldsEnum = "createdById"
	GetDataAssetFieldsUpdatedbyid    GetDataAssetFieldsEnum = "updatedById"
	GetDataAssetFieldsUri            GetDataAssetFieldsEnum = "uri"
	GetDataAssetFieldsProperties     GetDataAssetFieldsEnum = "properties"
)

var mappingGetDataAssetFieldsEnum = map[string]GetDataAssetFieldsEnum{
	"key":            GetDataAssetFieldsKey,
	"displayName":    GetDataAssetFieldsDisplayname,
	"description":    GetDataAssetFieldsDescription,
	"catalogId":      GetDataAssetFieldsCatalogid,
	"externalKey":    GetDataAssetFieldsExternalkey,
	"typeKey":        GetDataAssetFieldsTypekey,
	"lifecycleState": GetDataAssetFieldsLifecyclestate,
	"timeCreated":    GetDataAssetFieldsTimecreated,
	"timeUpdated":    GetDataAssetFieldsTimeupdated,
	"createdById":    GetDataAssetFieldsCreatedbyid,
	"updatedById":    GetDataAssetFieldsUpdatedbyid,
	"uri":            GetDataAssetFieldsUri,
	"properties":     GetDataAssetFieldsProperties,
}

var mappingGetDataAssetFieldsEnumLowerCase = map[string]GetDataAssetFieldsEnum{
	"key":            GetDataAssetFieldsKey,
	"displayname":    GetDataAssetFieldsDisplayname,
	"description":    GetDataAssetFieldsDescription,
	"catalogid":      GetDataAssetFieldsCatalogid,
	"externalkey":    GetDataAssetFieldsExternalkey,
	"typekey":        GetDataAssetFieldsTypekey,
	"lifecyclestate": GetDataAssetFieldsLifecyclestate,
	"timecreated":    GetDataAssetFieldsTimecreated,
	"timeupdated":    GetDataAssetFieldsTimeupdated,
	"createdbyid":    GetDataAssetFieldsCreatedbyid,
	"updatedbyid":    GetDataAssetFieldsUpdatedbyid,
	"uri":            GetDataAssetFieldsUri,
	"properties":     GetDataAssetFieldsProperties,
}

// GetGetDataAssetFieldsEnumValues Enumerates the set of values for GetDataAssetFieldsEnum
func GetGetDataAssetFieldsEnumValues() []GetDataAssetFieldsEnum {
	values := make([]GetDataAssetFieldsEnum, 0)
	for _, v := range mappingGetDataAssetFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDataAssetFieldsEnumStringValues Enumerates the set of values in String for GetDataAssetFieldsEnum
func GetGetDataAssetFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"catalogId",
		"externalKey",
		"typeKey",
		"lifecycleState",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"uri",
		"properties",
	}
}

// GetMappingGetDataAssetFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDataAssetFieldsEnum(val string) (GetDataAssetFieldsEnum, bool) {
	enum, ok := mappingGetDataAssetFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
