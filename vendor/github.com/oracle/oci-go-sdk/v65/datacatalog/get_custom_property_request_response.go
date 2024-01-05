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

// GetCustomPropertyRequest wrapper for the GetCustomProperty operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetCustomProperty.go.html to see an example of how to use GetCustomPropertyRequest.
type GetCustomPropertyRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique namespace identifier.
	NamespaceId *string `mandatory:"true" contributesTo:"path" name:"namespaceId"`

	// Unique Custom Property key
	CustomPropertyKey *string `mandatory:"true" contributesTo:"path" name:"customPropertyKey"`

	// Specifies the fields to return in a custom property response.
	Fields []GetCustomPropertyFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetCustomPropertyRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetCustomPropertyRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetCustomPropertyRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetCustomPropertyRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetCustomPropertyRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetCustomPropertyFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetCustomPropertyFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetCustomPropertyResponse wrapper for the GetCustomProperty operation
type GetCustomPropertyResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The CustomProperty instance
	CustomProperty `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetCustomPropertyResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetCustomPropertyResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetCustomPropertyFieldsEnum Enum with underlying type: string
type GetCustomPropertyFieldsEnum string

// Set of constants representing the allowable values for GetCustomPropertyFieldsEnum
const (
	GetCustomPropertyFieldsKey            GetCustomPropertyFieldsEnum = "key"
	GetCustomPropertyFieldsDisplayname    GetCustomPropertyFieldsEnum = "displayName"
	GetCustomPropertyFieldsDescription    GetCustomPropertyFieldsEnum = "description"
	GetCustomPropertyFieldsDatatype       GetCustomPropertyFieldsEnum = "dataType"
	GetCustomPropertyFieldsNamespacename  GetCustomPropertyFieldsEnum = "namespaceName"
	GetCustomPropertyFieldsLifecyclestate GetCustomPropertyFieldsEnum = "lifecycleState"
	GetCustomPropertyFieldsTimecreated    GetCustomPropertyFieldsEnum = "timeCreated"
	GetCustomPropertyFieldsTimeupdated    GetCustomPropertyFieldsEnum = "timeUpdated"
	GetCustomPropertyFieldsCreatedbyid    GetCustomPropertyFieldsEnum = "createdById"
	GetCustomPropertyFieldsUpdatedbyid    GetCustomPropertyFieldsEnum = "updatedById"
	GetCustomPropertyFieldsProperties     GetCustomPropertyFieldsEnum = "properties"
)

var mappingGetCustomPropertyFieldsEnum = map[string]GetCustomPropertyFieldsEnum{
	"key":            GetCustomPropertyFieldsKey,
	"displayName":    GetCustomPropertyFieldsDisplayname,
	"description":    GetCustomPropertyFieldsDescription,
	"dataType":       GetCustomPropertyFieldsDatatype,
	"namespaceName":  GetCustomPropertyFieldsNamespacename,
	"lifecycleState": GetCustomPropertyFieldsLifecyclestate,
	"timeCreated":    GetCustomPropertyFieldsTimecreated,
	"timeUpdated":    GetCustomPropertyFieldsTimeupdated,
	"createdById":    GetCustomPropertyFieldsCreatedbyid,
	"updatedById":    GetCustomPropertyFieldsUpdatedbyid,
	"properties":     GetCustomPropertyFieldsProperties,
}

var mappingGetCustomPropertyFieldsEnumLowerCase = map[string]GetCustomPropertyFieldsEnum{
	"key":            GetCustomPropertyFieldsKey,
	"displayname":    GetCustomPropertyFieldsDisplayname,
	"description":    GetCustomPropertyFieldsDescription,
	"datatype":       GetCustomPropertyFieldsDatatype,
	"namespacename":  GetCustomPropertyFieldsNamespacename,
	"lifecyclestate": GetCustomPropertyFieldsLifecyclestate,
	"timecreated":    GetCustomPropertyFieldsTimecreated,
	"timeupdated":    GetCustomPropertyFieldsTimeupdated,
	"createdbyid":    GetCustomPropertyFieldsCreatedbyid,
	"updatedbyid":    GetCustomPropertyFieldsUpdatedbyid,
	"properties":     GetCustomPropertyFieldsProperties,
}

// GetGetCustomPropertyFieldsEnumValues Enumerates the set of values for GetCustomPropertyFieldsEnum
func GetGetCustomPropertyFieldsEnumValues() []GetCustomPropertyFieldsEnum {
	values := make([]GetCustomPropertyFieldsEnum, 0)
	for _, v := range mappingGetCustomPropertyFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetCustomPropertyFieldsEnumStringValues Enumerates the set of values in String for GetCustomPropertyFieldsEnum
func GetGetCustomPropertyFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"dataType",
		"namespaceName",
		"lifecycleState",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"properties",
	}
}

// GetMappingGetCustomPropertyFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetCustomPropertyFieldsEnum(val string) (GetCustomPropertyFieldsEnum, bool) {
	enum, ok := mappingGetCustomPropertyFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
