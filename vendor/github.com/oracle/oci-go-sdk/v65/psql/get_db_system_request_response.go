// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package psql

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetDbSystemRequest wrapper for the GetDbSystem operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/psql/GetDbSystem.go.html to see an example of how to use GetDbSystemRequest.
type GetDbSystemRequest struct {

	// unique DbSystem identifier
	DbSystemId *string `mandatory:"true" contributesTo:"path" name:"dbSystemId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to exclude DB config  when this query param is set to OverrideDbConfig
	ExcludedFields []GetDbSystemExcludedFieldsEnum `contributesTo:"query" name:"excludedFields" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetDbSystemRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetDbSystemRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetDbSystemRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetDbSystemRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetDbSystemRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.ExcludedFields {
		if _, ok := GetMappingGetDbSystemExcludedFieldsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExcludedFields: %s. Supported values are: %s.", val, strings.Join(GetGetDbSystemExcludedFieldsEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetDbSystemResponse wrapper for the GetDbSystem operation
type GetDbSystemResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The DbSystem instance
	DbSystem `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetDbSystemResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetDbSystemResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetDbSystemExcludedFieldsEnum Enum with underlying type: string
type GetDbSystemExcludedFieldsEnum string

// Set of constants representing the allowable values for GetDbSystemExcludedFieldsEnum
const (
	GetDbSystemExcludedFieldsDbconfigurationparams GetDbSystemExcludedFieldsEnum = "dbConfigurationParams"
)

var mappingGetDbSystemExcludedFieldsEnum = map[string]GetDbSystemExcludedFieldsEnum{
	"dbConfigurationParams": GetDbSystemExcludedFieldsDbconfigurationparams,
}

var mappingGetDbSystemExcludedFieldsEnumLowerCase = map[string]GetDbSystemExcludedFieldsEnum{
	"dbconfigurationparams": GetDbSystemExcludedFieldsDbconfigurationparams,
}

// GetGetDbSystemExcludedFieldsEnumValues Enumerates the set of values for GetDbSystemExcludedFieldsEnum
func GetGetDbSystemExcludedFieldsEnumValues() []GetDbSystemExcludedFieldsEnum {
	values := make([]GetDbSystemExcludedFieldsEnum, 0)
	for _, v := range mappingGetDbSystemExcludedFieldsEnum {
		values = append(values, v)
	}
	return values
}

// GetGetDbSystemExcludedFieldsEnumStringValues Enumerates the set of values in String for GetDbSystemExcludedFieldsEnum
func GetGetDbSystemExcludedFieldsEnumStringValues() []string {
	return []string{
		"dbConfigurationParams",
	}
}

// GetMappingGetDbSystemExcludedFieldsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetDbSystemExcludedFieldsEnum(val string) (GetDbSystemExcludedFieldsEnum, bool) {
	enum, ok := mappingGetDbSystemExcludedFieldsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
