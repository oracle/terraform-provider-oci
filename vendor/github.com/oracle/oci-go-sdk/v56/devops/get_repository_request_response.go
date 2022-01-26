// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// GetRepositoryRequest wrapper for the GetRepository operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/GetRepository.go.html to see an example of how to use GetRepositoryRequest.
type GetRepositoryRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Fields parameter can contain multiple flags useful in deciding the API functionality.
	Fields []GetRepositoryFieldsEnum `contributesTo:"query" name:"fields" omitEmpty:"true" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetRepositoryRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetRepositoryRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetRepositoryRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetRepositoryRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// GetRepositoryResponse wrapper for the GetRepository operation
type GetRepositoryResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The Repository instance
	Repository `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetRepositoryResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetRepositoryResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetRepositoryFieldsEnum Enum with underlying type: string
type GetRepositoryFieldsEnum string

// Set of constants representing the allowable values for GetRepositoryFieldsEnum
const (
	GetRepositoryFieldsBranchcount GetRepositoryFieldsEnum = "branchCount"
	GetRepositoryFieldsCommitcount GetRepositoryFieldsEnum = "commitCount"
	GetRepositoryFieldsSizeinbytes GetRepositoryFieldsEnum = "sizeInBytes"
)

var mappingGetRepositoryFields = map[string]GetRepositoryFieldsEnum{
	"branchCount": GetRepositoryFieldsBranchcount,
	"commitCount": GetRepositoryFieldsCommitcount,
	"sizeInBytes": GetRepositoryFieldsSizeinbytes,
}

// GetGetRepositoryFieldsEnumValues Enumerates the set of values for GetRepositoryFieldsEnum
func GetGetRepositoryFieldsEnumValues() []GetRepositoryFieldsEnum {
	values := make([]GetRepositoryFieldsEnum, 0)
	for _, v := range mappingGetRepositoryFields {
		values = append(values, v)
	}
	return values
}
