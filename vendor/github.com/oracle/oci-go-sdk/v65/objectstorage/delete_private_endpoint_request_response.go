// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// DeletePrivateEndpointRequest wrapper for the DeletePrivateEndpoint operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/objectstorage/DeletePrivateEndpoint.go.html to see an example of how to use DeletePrivateEndpointRequest.
type DeletePrivateEndpointRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The name of the private endpoint. Avoid entering confidential information.
	// Example: `my-new-pe-1`
	PeName *string `mandatory:"true" contributesTo:"path" name:"peName"`

	// The entity tag (ETag) to match with the ETag of an existing resource. If the specified ETag matches the ETag of
	// the existing resource, GET and HEAD requests will return the resource and PUT and POST requests will upload
	// the resource.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request DeletePrivateEndpointRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request DeletePrivateEndpointRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request DeletePrivateEndpointRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request DeletePrivateEndpointRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
	if mandatoryParamMap["namespaceName"] != nil {
		templateParam := mandatoryParamMap["namespaceName"]
		for _, template := range templateParam {
			replacementParam := *request.NamespaceName
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
	if mandatoryParamMap["peName"] != nil {
		templateParam := mandatoryParamMap["peName"]
		for _, template := range templateParam {
			replacementParam := *request.PeName
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request DeletePrivateEndpointRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request DeletePrivateEndpointRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeletePrivateEndpointResponse wrapper for the DeletePrivateEndpoint operation
type DeletePrivateEndpointResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the asynchronous request. If you need to contact Oracle about a
	// particular request, provide this request ID.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response DeletePrivateEndpointResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response DeletePrivateEndpointResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
