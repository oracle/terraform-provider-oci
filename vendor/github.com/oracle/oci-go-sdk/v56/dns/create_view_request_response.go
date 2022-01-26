// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// CreateViewRequest wrapper for the CreateView operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/CreateView.go.html to see an example of how to use CreateViewRequest.
type CreateViewRequest struct {

	// Details for creating a new view.
	CreateViewDetails `contributesTo:"body"`

	// A token that uniquely identifies a request so it can be retried in case
	// of a timeout or server error without risk of executing that same action
	// again. Retry tokens expire after 24 hours, but can be invalidated before
	// then due to conflicting operations (for example, if a resource has been
	// deleted and purged from the system, then a retry of the original creation
	// request may be rejected).
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope CreateViewScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request CreateViewRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request CreateViewRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request CreateViewRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request CreateViewRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// CreateViewResponse wrapper for the CreateView operation
type CreateViewResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The View instance
	View `presentIn:"body"`

	// The current version of the resource, ending with a
	// representation-specific suffix. This value may be used in If-Match
	// and If-None-Match headers for later requests of the same resource.
	Etag *string `presentIn:"header" name:"etag"`

	// The full URI of the resource related to the request.
	Location *string `presentIn:"header" name:"location"`

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response CreateViewResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response CreateViewResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// CreateViewScopeEnum Enum with underlying type: string
type CreateViewScopeEnum string

// Set of constants representing the allowable values for CreateViewScopeEnum
const (
	CreateViewScopeGlobal  CreateViewScopeEnum = "GLOBAL"
	CreateViewScopePrivate CreateViewScopeEnum = "PRIVATE"
)

var mappingCreateViewScope = map[string]CreateViewScopeEnum{
	"GLOBAL":  CreateViewScopeGlobal,
	"PRIVATE": CreateViewScopePrivate,
}

// GetCreateViewScopeEnumValues Enumerates the set of values for CreateViewScopeEnum
func GetCreateViewScopeEnumValues() []CreateViewScopeEnum {
	values := make([]CreateViewScopeEnum, 0)
	for _, v := range mappingCreateViewScope {
		values = append(values, v)
	}
	return values
}
