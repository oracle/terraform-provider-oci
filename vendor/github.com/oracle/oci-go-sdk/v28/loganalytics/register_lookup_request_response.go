// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v28/common"
	"io"
	"net/http"
)

// RegisterLookupRequest wrapper for the RegisterLookup operation
type RegisterLookupRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// type - possible values are Lookup or Dictionary
	Type RegisterLookupTypeEnum `mandatory:"true" contributesTo:"query" name:"type" omitEmpty:"true"`

	// file containing data for lookup creation
	RegisterLookupContentFileBody io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// A filter to return only log analytics entities whose name matches the entire name given. The match
	// is case-insensitive.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Description of the fields to get
	Description *string `mandatory:"false" contributesTo:"query" name:"description"`

	// Character Encoding
	CharEncoding *string `mandatory:"false" contributesTo:"query" name:"charEncoding"`

	// is hidden
	IsHidden *bool `mandatory:"false" contributesTo:"query" name:"isHidden"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request RegisterLookupRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RegisterLookupRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RegisterLookupRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// RegisterLookupResponse wrapper for the RegisterLookup operation
type RegisterLookupResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LogAnalyticsLookup instance
	LogAnalyticsLookup `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response RegisterLookupResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RegisterLookupResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RegisterLookupTypeEnum Enum with underlying type: string
type RegisterLookupTypeEnum string

// Set of constants representing the allowable values for RegisterLookupTypeEnum
const (
	RegisterLookupTypeLookup     RegisterLookupTypeEnum = "Lookup"
	RegisterLookupTypeDictionary RegisterLookupTypeEnum = "Dictionary"
)

var mappingRegisterLookupType = map[string]RegisterLookupTypeEnum{
	"Lookup":     RegisterLookupTypeLookup,
	"Dictionary": RegisterLookupTypeDictionary,
}

// GetRegisterLookupTypeEnumValues Enumerates the set of values for RegisterLookupTypeEnum
func GetRegisterLookupTypeEnumValues() []RegisterLookupTypeEnum {
	values := make([]RegisterLookupTypeEnum, 0)
	for _, v := range mappingRegisterLookupType {
		values = append(values, v)
	}
	return values
}
