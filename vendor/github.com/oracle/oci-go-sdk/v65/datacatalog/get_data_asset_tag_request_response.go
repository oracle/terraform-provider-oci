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

// GetDataAssetTagRequest wrapper for the GetDataAssetTag operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetDataAssetTag.go.html to see an example of how to use GetDataAssetTagRequest.
type GetDataAssetTagRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique data asset key.
	DataAssetKey *string `mandatory:"true" contributesTo:"path" name:"dataAssetKey"`

	// Unique tag key.
	TagKey *string `mandatory:"true" contributesTo:"path" name:"tagKey"`

	// Specifies the fields to return in a data asset tag response.
	Fields []GetDataAssetTagFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDataAssetTagRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDataAssetTagRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDataAssetTagRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDataAssetTagRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetDataAssetTagRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetDataAssetTagFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetDataAssetTagFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetDataAssetTagResponse wrapper for the GetDataAssetTag operation
type GetDataAssetTagResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DataAssetTag instance
	DataAssetTag `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDataAssetTagResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDataAssetTagResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDataAssetTagFieldsEnum Enum with underlying type: string
type GetDataAssetTagFieldsEnum string

// Set of constants representing the allowable values for GetDataAssetTagFieldsEnum
const (
	GetDataAssetTagFieldsKey             GetDataAssetTagFieldsEnum = "key"
	GetDataAssetTagFieldsName            GetDataAssetTagFieldsEnum = "name"
	GetDataAssetTagFieldsTermkey         GetDataAssetTagFieldsEnum = "termKey"
	GetDataAssetTagFieldsTermpath        GetDataAssetTagFieldsEnum = "termPath"
	GetDataAssetTagFieldsTermdescription GetDataAssetTagFieldsEnum = "termDescription"
	GetDataAssetTagFieldsLifecyclestate  GetDataAssetTagFieldsEnum = "lifecycleState"
	GetDataAssetTagFieldsTimecreated     GetDataAssetTagFieldsEnum = "timeCreated"
	GetDataAssetTagFieldsCreatedbyid     GetDataAssetTagFieldsEnum = "createdById"
	GetDataAssetTagFieldsUri             GetDataAssetTagFieldsEnum = "uri"
	GetDataAssetTagFieldsDataassetkey    GetDataAssetTagFieldsEnum = "dataAssetKey"
)

var mappingGetDataAssetTagFieldsEnum = map[string]GetDataAssetTagFieldsEnum{
	"key":             GetDataAssetTagFieldsKey,
	"name":            GetDataAssetTagFieldsName,
	"termKey":         GetDataAssetTagFieldsTermkey,
	"termPath":        GetDataAssetTagFieldsTermpath,
	"termDescription": GetDataAssetTagFieldsTermdescription,
	"lifecycleState":  GetDataAssetTagFieldsLifecyclestate,
	"timeCreated":     GetDataAssetTagFieldsTimecreated,
	"createdById":     GetDataAssetTagFieldsCreatedbyid,
	"uri":             GetDataAssetTagFieldsUri,
	"dataAssetKey":    GetDataAssetTagFieldsDataassetkey,
}

var mappingGetDataAssetTagFieldsEnumLowerCase = map[string]GetDataAssetTagFieldsEnum{
	"key":             GetDataAssetTagFieldsKey,
	"name":            GetDataAssetTagFieldsName,
	"termkey":         GetDataAssetTagFieldsTermkey,
	"termpath":        GetDataAssetTagFieldsTermpath,
	"termdescription": GetDataAssetTagFieldsTermdescription,
	"lifecyclestate":  GetDataAssetTagFieldsLifecyclestate,
	"timecreated":     GetDataAssetTagFieldsTimecreated,
	"createdbyid":     GetDataAssetTagFieldsCreatedbyid,
	"uri":             GetDataAssetTagFieldsUri,
	"dataassetkey":    GetDataAssetTagFieldsDataassetkey,
}

// GetGetDataAssetTagFieldsEnumValues Enumerates the set of values for GetDataAssetTagFieldsEnum
func GetGetDataAssetTagFieldsEnumValues() []GetDataAssetTagFieldsEnum {
	values := make([]GetDataAssetTagFieldsEnum, 0)
	for _, v := range mappingGetDataAssetTagFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDataAssetTagFieldsEnumStringValues Enumerates the set of values in String for GetDataAssetTagFieldsEnum
func GetGetDataAssetTagFieldsEnumStringValues() []string {
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
		"dataAssetKey",
	}
}

// GetMappingGetDataAssetTagFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDataAssetTagFieldsEnum(val string) (GetDataAssetTagFieldsEnum, bool) {
	enum, ok := mappingGetDataAssetTagFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
