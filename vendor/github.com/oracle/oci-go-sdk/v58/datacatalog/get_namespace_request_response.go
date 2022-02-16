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

// GetNamespaceRequest wrapper for the GetNamespace operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetNamespace.go.html to see an example of how to use GetNamespaceRequest.
type GetNamespaceRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique namespace identifier.
	NamespaceId *string `mandatory:"true" contributesTo:"path" name:"namespaceId"`

	// Specifies the fields to return in a namespace response.
	Fields []GetNamespaceFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetNamespaceRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetNamespaceRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetNamespaceRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetNamespaceRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetNamespaceRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetNamespaceFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetNamespaceFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetNamespaceResponse wrapper for the GetNamespace operation
type GetNamespaceResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Namespace instance
	Namespace `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetNamespaceResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetNamespaceResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetNamespaceFieldsEnum Enum with underlying type: string
type GetNamespaceFieldsEnum string

// Set of constants representing the allowable values for GetNamespaceFieldsEnum
const (
	GetNamespaceFieldsKey            GetNamespaceFieldsEnum = "key"
	GetNamespaceFieldsDisplayname    GetNamespaceFieldsEnum = "displayName"
	GetNamespaceFieldsDescription    GetNamespaceFieldsEnum = "description"
	GetNamespaceFieldsLifecyclestate GetNamespaceFieldsEnum = "lifecycleState"
	GetNamespaceFieldsTimecreated    GetNamespaceFieldsEnum = "timeCreated"
	GetNamespaceFieldsTimeupdated    GetNamespaceFieldsEnum = "timeUpdated"
	GetNamespaceFieldsCreatedbyid    GetNamespaceFieldsEnum = "createdById"
	GetNamespaceFieldsUpdatedbyid    GetNamespaceFieldsEnum = "updatedById"
	GetNamespaceFieldsProperties     GetNamespaceFieldsEnum = "properties"
)

var mappingGetNamespaceFieldsEnum = map[string]GetNamespaceFieldsEnum{
	"key":            GetNamespaceFieldsKey,
	"displayName":    GetNamespaceFieldsDisplayname,
	"description":    GetNamespaceFieldsDescription,
	"lifecycleState": GetNamespaceFieldsLifecyclestate,
	"timeCreated":    GetNamespaceFieldsTimecreated,
	"timeUpdated":    GetNamespaceFieldsTimeupdated,
	"createdById":    GetNamespaceFieldsCreatedbyid,
	"updatedById":    GetNamespaceFieldsUpdatedbyid,
	"properties":     GetNamespaceFieldsProperties,
}

// GetGetNamespaceFieldsEnumValues Enumerates the set of values for GetNamespaceFieldsEnum
func GetGetNamespaceFieldsEnumValues() []GetNamespaceFieldsEnum {
	values := make([]GetNamespaceFieldsEnum, 0)
	for _, v := range mappingGetNamespaceFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetNamespaceFieldsEnumStringValues Enumerates the set of values in String for GetNamespaceFieldsEnum
func GetGetNamespaceFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"lifecycleState",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"properties",
	}
}

// GetMappingGetNamespaceFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetNamespaceFieldsEnum(val string) (GetNamespaceFieldsEnum, bool) {
	mappingGetNamespaceFieldsEnumIgnoreCase := make(map[string]GetNamespaceFieldsEnum)
	for k, v := range mappingGetNamespaceFieldsEnum {
		mappingGetNamespaceFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetNamespaceFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
