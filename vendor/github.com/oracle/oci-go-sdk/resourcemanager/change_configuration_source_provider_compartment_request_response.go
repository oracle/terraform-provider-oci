// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ChangeConfigurationSourceProviderCompartmentRequest wrapper for the ChangeConfigurationSourceProviderCompartment operation
type ChangeConfigurationSourceProviderCompartmentRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the configuration source provider.
	ConfigurationSourceProviderId *string `mandatory:"true" contributesTo:"path" name:"configurationSourceProviderId"`

	// Defines the properties of changeConfigurationSourceProviderCompartment operation.
	ChangeConfigurationSourceProviderCompartmentDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the `PUT` or `DELETE` call for a resource, set the `if-match`
	// parameter to the value of the etag from a previous `GET` or `POST` response for that resource.  The resource
	// will be updated or deleted only if the etag you provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of retrying the same action. Retry tokens expire after
	// 24 hours, but can be invalidated before then due to conflicting operations. For example,
	// if a resource has been deleted and purged from the system, then a retry of the original
	// creation request may be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ChangeConfigurationSourceProviderCompartmentRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ChangeConfigurationSourceProviderCompartmentRequest) HTTPRequest(method, path string) (http.Request, error) {
	return common.MakeDefaultHTTPRequestWithTaggedStruct(method, path, request)
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ChangeConfigurationSourceProviderCompartmentRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ChangeConfigurationSourceProviderCompartmentResponse wrapper for the ChangeConfigurationSourceProviderCompartment operation
type ChangeConfigurationSourceProviderCompartmentResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ChangeConfigurationSourceProviderCompartmentResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ChangeConfigurationSourceProviderCompartmentResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
