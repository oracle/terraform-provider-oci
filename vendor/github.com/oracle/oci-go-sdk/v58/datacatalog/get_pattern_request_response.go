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

// GetPatternRequest wrapper for the GetPattern operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datacatalog/GetPattern.go.html to see an example of how to use GetPatternRequest.
type GetPatternRequest struct {

	// Unique catalog identifier.
	CatalogId *string `mandatory:"true" contributesTo:"path" name:"catalogId"`

	// Unique pattern key.
	PatternKey *string `mandatory:"true" contributesTo:"path" name:"patternKey"`

	// Specifies the fields to return in a pattern response.
	Fields []GetPatternFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetPatternRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetPatternRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetPatternRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetPatternRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetPatternRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.Fields {
		if _, ok := GetMappingGetPatternFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Fields: %s. Supported values are: %s.", val, strings.Join(GetGetPatternFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetPatternResponse wrapper for the GetPattern operation
type GetPatternResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Pattern instance
	Pattern `presentIn:"body"`

	// For optimistic concurrency control. See ETags for Optimistic Concurrency Control (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#eleven).
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetPatternResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetPatternResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetPatternFieldsEnum Enum with underlying type: string
type GetPatternFieldsEnum string

// Set of constants representing the allowable values for GetPatternFieldsEnum
const (
	GetPatternFieldsKey            GetPatternFieldsEnum = "key"
	GetPatternFieldsDisplayname    GetPatternFieldsEnum = "displayName"
	GetPatternFieldsDescription    GetPatternFieldsEnum = "description"
	GetPatternFieldsCatalogid      GetPatternFieldsEnum = "catalogId"
	GetPatternFieldsExpression     GetPatternFieldsEnum = "expression"
	GetPatternFieldsLifecyclestate GetPatternFieldsEnum = "lifecycleState"
	GetPatternFieldsTimecreated    GetPatternFieldsEnum = "timeCreated"
	GetPatternFieldsTimeupdated    GetPatternFieldsEnum = "timeUpdated"
	GetPatternFieldsCreatedbyid    GetPatternFieldsEnum = "createdById"
	GetPatternFieldsUpdatedbyid    GetPatternFieldsEnum = "updatedById"
	GetPatternFieldsProperties     GetPatternFieldsEnum = "properties"
)

var mappingGetPatternFieldsEnum = map[string]GetPatternFieldsEnum{
	"key":            GetPatternFieldsKey,
	"displayName":    GetPatternFieldsDisplayname,
	"description":    GetPatternFieldsDescription,
	"catalogId":      GetPatternFieldsCatalogid,
	"expression":     GetPatternFieldsExpression,
	"lifecycleState": GetPatternFieldsLifecyclestate,
	"timeCreated":    GetPatternFieldsTimecreated,
	"timeUpdated":    GetPatternFieldsTimeupdated,
	"createdById":    GetPatternFieldsCreatedbyid,
	"updatedById":    GetPatternFieldsUpdatedbyid,
	"properties":     GetPatternFieldsProperties,
}

// GetGetPatternFieldsEnumValues Enumerates the set of values for GetPatternFieldsEnum
func GetGetPatternFieldsEnumValues() []GetPatternFieldsEnum {
	values := make([]GetPatternFieldsEnum, 0)
	for _, v := range mappingGetPatternFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetPatternFieldsEnumStringValues Enumerates the set of values in String for GetPatternFieldsEnum
func GetGetPatternFieldsEnumStringValues() []string {
	return []string{
		"key",
		"displayName",
		"description",
		"catalogId",
		"expression",
		"lifecycleState",
		"timeCreated",
		"timeUpdated",
		"createdById",
		"updatedById",
		"properties",
	}
}

// GetMappingGetPatternFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetPatternFieldsEnum(val string) (GetPatternFieldsEnum, bool) {
	mappingGetPatternFieldsEnumIgnoreCase := make(map[string]GetPatternFieldsEnum)
	for k, v := range mappingGetPatternFieldsEnum {
		mappingGetPatternFieldsEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetPatternFieldsEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
