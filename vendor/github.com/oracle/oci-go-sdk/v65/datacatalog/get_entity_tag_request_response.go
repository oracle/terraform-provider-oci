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

// GetEntityTagRequest wrapper for the GetEntityTag operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetEntityTag.go.html to see an example of how to use GetEntityTagRequest.
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
func (request GetEntityTagRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetEntityTagRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetEntityTagRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetEntityTagRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetEntityTagFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetEntityTagFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingGetEntityTagFieldsEnum = map[string]GetEntityTagFieldsEnum{
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

var mappingGetEntityTagFieldsEnumLowerCase = map[string]GetEntityTagFieldsEnum{
	"key":             GetEntityTagFieldsKey,
	"name":            GetEntityTagFieldsName,
	"termkey":         GetEntityTagFieldsTermkey,
	"termpath":        GetEntityTagFieldsTermpath,
	"termdescription": GetEntityTagFieldsTermdescription,
	"lifecyclestate":  GetEntityTagFieldsLifecyclestate,
	"timecreated":     GetEntityTagFieldsTimecreated,
	"createdbyid":     GetEntityTagFieldsCreatedbyid,
	"uri":             GetEntityTagFieldsUri,
	"entitykey":       GetEntityTagFieldsEntitykey,
}

// GetGetEntityTagFieldsEnumValues Enumerates the set of values for GetEntityTagFieldsEnum
func GetGetEntityTagFieldsEnumValues() []GetEntityTagFieldsEnum {
	values := make([]GetEntityTagFieldsEnum, 0)
	for _, v := range mappingGetEntityTagFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetEntityTagFieldsEnumStringValues Enumerates the set of values in String for GetEntityTagFieldsEnum
func GetGetEntityTagFieldsEnumStringValues() []string {
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
		"entityKey",
	}
}

// GetMappingGetEntityTagFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetEntityTagFieldsEnum(val string) (GetEntityTagFieldsEnum, bool) {
	enum, ok := mappingGetEntityTagFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
