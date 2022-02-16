// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// GetMirrorRecordRequest wrapper for the GetMirrorRecord operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetMirrorRecord.go.html to see an example of how to use GetMirrorRecordRequest.
type GetMirrorRecordRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// The field of mirror record type. Only one mirror record type can be provided:
	// current - The current mirror record.
	// lastSuccessful - The last successful mirror record.
	MirrorRecordType GetMirrorRecordMirrorRecordTypeEnum `mandatory:"true" contributesTo:"path" name:"mirrorRecordType"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetMirrorRecordRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetMirrorRecordRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetMirrorRecordRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetMirrorRecordRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetMirrorRecordRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetMirrorRecordMirrorRecordTypeEnum(string(request.MirrorRecordType)); !ok && request.MirrorRecordType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MirrorRecordType: %s. Supported values are: %s.", request.MirrorRecordType, strings.Join(GetGetMirrorRecordMirrorRecordTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetMirrorRecordResponse wrapper for the GetMirrorRecord operation
type GetMirrorRecordResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The RepositoryMirrorRecord instance
	RepositoryMirrorRecord `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetMirrorRecordResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetMirrorRecordResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetMirrorRecordMirrorRecordTypeEnum Enum with underlying type: string
type GetMirrorRecordMirrorRecordTypeEnum string

// Set of constants representing the allowable values for GetMirrorRecordMirrorRecordTypeEnum
const (
	GetMirrorRecordMirrorRecordTypeCurrent        GetMirrorRecordMirrorRecordTypeEnum = "current"
	GetMirrorRecordMirrorRecordTypeLastsuccessful GetMirrorRecordMirrorRecordTypeEnum = "lastSuccessful"
)

var mappingGetMirrorRecordMirrorRecordTypeEnum = map[string]GetMirrorRecordMirrorRecordTypeEnum{
	"current":        GetMirrorRecordMirrorRecordTypeCurrent,
	"lastSuccessful": GetMirrorRecordMirrorRecordTypeLastsuccessful,
}

// GetGetMirrorRecordMirrorRecordTypeEnumValues Enumerates the set of values for GetMirrorRecordMirrorRecordTypeEnum
func GetGetMirrorRecordMirrorRecordTypeEnumValues() []GetMirrorRecordMirrorRecordTypeEnum {
	values := make([]GetMirrorRecordMirrorRecordTypeEnum, 0)
	for _, v := range mappingGetMirrorRecordMirrorRecordTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetMirrorRecordMirrorRecordTypeEnumStringValues Enumerates the set of values in String for GetMirrorRecordMirrorRecordTypeEnum
func GetGetMirrorRecordMirrorRecordTypeEnumStringValues() []string {
	return []string{
		"current",
		"lastSuccessful",
	}
}

// GetMappingGetMirrorRecordMirrorRecordTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetMirrorRecordMirrorRecordTypeEnum(val string) (GetMirrorRecordMirrorRecordTypeEnum, bool) {
	mappingGetMirrorRecordMirrorRecordTypeEnumIgnoreCase := make(map[string]GetMirrorRecordMirrorRecordTypeEnum)
	for k, v := range mappingGetMirrorRecordMirrorRecordTypeEnum {
		mappingGetMirrorRecordMirrorRecordTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingGetMirrorRecordMirrorRecordTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
