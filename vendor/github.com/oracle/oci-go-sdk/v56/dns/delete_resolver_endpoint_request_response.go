// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dns

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// DeleteResolverEndpointRequest wrapper for the DeleteResolverEndpoint operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dns/DeleteResolverEndpoint.go.html to see an example of how to use DeleteResolverEndpointRequest.
type DeleteResolverEndpointRequest struct {

	// The OCID of the target resolver.
	ResolverId *string `mandatory:"true" contributesTo:"path" name:"resolverId"`

	// The name of the target resolver endpoint.
	ResolverEndpointName *string `mandatory:"true" contributesTo:"path" name:"resolverEndpointName"`

	// The `If-Match` header field makes the request method conditional on the
	// existence of at least one current representation of the target resource,
	// when the field-value is `*`, or having a current representation of the
	// target resource that has an entity-tag matching a member of the list of
	// entity-tags provided in the field-value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"If-Match"`

	// The `If-Unmodified-Since` header field makes the request method
	// conditional on the selected representation's last modification date being
	// earlier than or equal to the date provided in the field-value.  This
	// field accomplishes the same purpose as If-Match for cases where the user
	// agent does not have an entity-tag for the representation.
	IfUnmodifiedSince *string `mandatory:"false" contributesTo:"header" name:"If-Unmodified-Since"`

	// Unique Oracle-assigned identifier for the request. If you need
	// to contact Oracle about a particular request, please provide
	// the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Specifies to operate only on resources that have a matching DNS scope.
	Scope DeleteResolverEndpointScopeEnum `mandatory:"false" contributesTo:"query" name:"scope" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeleteResolverEndpointRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeleteResolverEndpointRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeleteResolverEndpointRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeleteResolverEndpointRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// DeleteResolverEndpointResponse wrapper for the DeleteResolverEndpoint operation
type DeleteResolverEndpointResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to
	// contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// Unique Oracle-assigned identifier for the asynchronous request.
	// You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`
}

func (response DeleteResolverEndpointResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeleteResolverEndpointResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// DeleteResolverEndpointScopeEnum Enum with underlying type: string
type DeleteResolverEndpointScopeEnum string

// Set of constants representing the allowable values for DeleteResolverEndpointScopeEnum
const (
	DeleteResolverEndpointScopeGlobal  DeleteResolverEndpointScopeEnum = "GLOBAL"
	DeleteResolverEndpointScopePrivate DeleteResolverEndpointScopeEnum = "PRIVATE"
)

var mappingDeleteResolverEndpointScope = map[string]DeleteResolverEndpointScopeEnum{
	"GLOBAL":  DeleteResolverEndpointScopeGlobal,
	"PRIVATE": DeleteResolverEndpointScopePrivate,
}

// GetDeleteResolverEndpointScopeEnumValues Enumerates the set of values for DeleteResolverEndpointScopeEnum
func GetDeleteResolverEndpointScopeEnumValues() []DeleteResolverEndpointScopeEnum {
	values := make([]DeleteResolverEndpointScopeEnum, 0)
	for _, v := range mappingDeleteResolverEndpointScope {
		values = append(values, v)
	}
	return values
}
