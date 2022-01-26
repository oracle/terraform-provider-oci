// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// UpdateLogAnalyticsObjectCollectionRuleRequest wrapper for the UpdateLogAnalyticsObjectCollectionRule operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/UpdateLogAnalyticsObjectCollectionRule.go.html to see an example of how to use UpdateLogAnalyticsObjectCollectionRuleRequest.
type UpdateLogAnalyticsObjectCollectionRuleRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The Logging Analytics Object Collection Rule OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	LogAnalyticsObjectCollectionRuleId *string `mandatory:"true" contributesTo:"path" name:"logAnalyticsObjectCollectionRuleId"`

	// The rule config to be updated.
	UpdateLogAnalyticsObjectCollectionRuleDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateLogAnalyticsObjectCollectionRuleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateLogAnalyticsObjectCollectionRuleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateLogAnalyticsObjectCollectionRuleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateLogAnalyticsObjectCollectionRuleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateLogAnalyticsObjectCollectionRuleResponse wrapper for the UpdateLogAnalyticsObjectCollectionRule operation
type UpdateLogAnalyticsObjectCollectionRuleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The LogAnalyticsObjectCollectionRule instance
	LogAnalyticsObjectCollectionRule `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateLogAnalyticsObjectCollectionRuleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateLogAnalyticsObjectCollectionRuleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
