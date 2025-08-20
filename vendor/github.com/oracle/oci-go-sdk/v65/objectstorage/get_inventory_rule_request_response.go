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

// GetInventoryRuleRequest wrapper for the GetInventoryRule operation
type GetInventoryRuleRequest struct {

	// The Object Storage namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The ID of the inventory rule.
	InventoryRuleId *string `mandatory:"true" contributesTo:"path" name:"inventoryRuleId"`

	// The client request ID for tracing.
	OpcClientRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-client-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetInventoryRuleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetInventoryRuleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetInventoryRuleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// ReplaceMandatoryParamInPath replaces the mandatory parameter in the path with the value provided.
// Not all services are supporting this feature and this method will be a no-op for those services.
func (request GetInventoryRuleRequest) ReplaceMandatoryParamInPath(client *common.BaseClient, mandatoryParamMap map[string][]common.TemplateParamForPerRealmEndpoint) {
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
	if mandatoryParamMap["inventoryRuleId"] != nil {
		templateParam := mandatoryParamMap["inventoryRuleId"]
		for _, template := range templateParam {
			replacementParam := *request.InventoryRuleId
			if template.EndsWithDot {
				replacementParam = replacementParam + "."
			}
			client.Host = strings.Replace(client.Host, template.Template, replacementParam, -1)
		}
	}
}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetInventoryRuleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetInventoryRuleRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetInventoryRuleResponse wrapper for the GetInventoryRule operation
type GetInventoryRuleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The InventoryRule instance
	InventoryRule `presentIn:"body"`

	// Echoes back the value passed in the opc-client-request-id header, for use by clients when debugging.
	OpcClientRequestId *string `presentIn:"header" name:"opc-client-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular
	// request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The entity tag (ETag) for the inventory rule.
	Etag *string `presentIn:"header" name:"etag"`

	// The time the inventory rule was last modified, as described in RFC 2616 (https://tools.ietf.org/html/rfc2616#section-14.29)
	LastModified *common.SDKTime `presentIn:"header" name:"last-modified"`
}

func (response GetInventoryRuleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetInventoryRuleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
