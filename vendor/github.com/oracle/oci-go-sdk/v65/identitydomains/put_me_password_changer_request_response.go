// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// PutMePasswordChangerRequest wrapper for the PutMePasswordChanger operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/identitydomains/PutMePasswordChanger.go.html to see an example of how to use PutMePasswordChangerRequest.
type PutMePasswordChangerRequest struct {

	// The Authorization field value consists of credentials containing the authentication information of the user agent for the realm of the resource being requested.
	Authorization *string `mandatory:"false" contributesTo:"header" name:"authorization"`

	// An endpoint-specific schema version number to use in the Request. Allowed version values are Earliest Version or Latest Version as specified in each REST API endpoint description, or any sequential number inbetween. All schema attributes/body parameters are a part of version 1. After version 1, any attributes added or deprecated will be tagged with the version that they were added to or deprecated in. If no version is provided, the latest schema version is returned.
	ResourceTypeSchemaVersion *string `mandatory:"false" contributesTo:"header" name:"resource_type_schema_version"`

	// Replace the current instance of MePasswordChanger with provided payload.
	// Before you specify an attribute-value in a request to replace a resource, please check the **'mutability'** property of that attribute in the resource-type schema below. Clicking on an attribute-row will expand that row to show the **SCIM++ Properties** of that attribute.
	// - Your request to create, update or replace a resource may specify in its payload a value for any attribute that is defined as *mutability:readWrite* or *mutability:writeOnly* or *mutability:immutable*:
	//   - The SCIM APIs to create a resource will ignore silently any value that you specify for an attribute that is defined as *mutability:readOnly*.
	//   - The SCIM APIs to update or replace a resource will fail with an error 400 Bad Request if you specify a value for an attribute that is defined as *mutability:readOnly*.
	//   - Similarly, the SCIM APIs to update or replace a resource will fail with an error 400 Bad Request if you specify any value for an attribute that is defined as *mutability:immutable* and that already has a value in the specified resource.
	// Also, before you use the query-parameter attributes to request specific attributes, please check the **'returned'** property of that attribute in the resource-type schema below:
	// - Your request to read a resource (or to search a resource-type) can specify as the value of attributes any attributes that are defined as *returned:default* or *returned:request* or *returned:always*:
	//   - If you request a specific set of attributes, the SCIM APIs to read a resource (or to search a resource-type) will return in each resource the set of attributes that you requested, as well as any attribute that is defined as *returned:always*.
	//   - If you do not request a specific set of attributes, the SCIM APIs to read a resource (or to search a resource-type) will return in each resource the the set of attributes defined as *returned:default*, as well as any attribute that is defined as *returned:always*.
	//   - The SCIM APIs to read a resource (or to search a resource-type) will ignore silently any request to return an attribute that is defined as *returned:never*.
	MePasswordChanger `contributesTo:"body"`

	// Used to make the request conditional on an ETag
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token you supply to uniquely identify the request and provide idempotency if the request is retried. Idempotency tokens expire after 24 hours.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request PutMePasswordChangerRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request PutMePasswordChangerRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request PutMePasswordChangerRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request PutMePasswordChangerRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request PutMePasswordChangerRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PutMePasswordChangerResponse wrapper for the PutMePasswordChanger operation
type PutMePasswordChangerResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The MePasswordChanger instance
	MePasswordChanger `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response PutMePasswordChangerResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response PutMePasswordChangerResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
