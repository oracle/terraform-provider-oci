// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
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

var mappingGetMirrorRecordMirrorRecordType = map[string]GetMirrorRecordMirrorRecordTypeEnum{
	"current":        GetMirrorRecordMirrorRecordTypeCurrent,
	"lastSuccessful": GetMirrorRecordMirrorRecordTypeLastsuccessful,
}

// GetGetMirrorRecordMirrorRecordTypeEnumValues Enumerates the set of values for GetMirrorRecordMirrorRecordTypeEnum
func GetGetMirrorRecordMirrorRecordTypeEnumValues() []GetMirrorRecordMirrorRecordTypeEnum {
	values := make([]GetMirrorRecordMirrorRecordTypeEnum, 0)
	for _, v := range mappingGetMirrorRecordMirrorRecordType {
		values = append(values, v)
	}
	return values
}
