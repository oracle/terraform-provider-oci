// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package cloudguard

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// UpdateTargetDetectorRecipeDetectorRuleRequest wrapper for the UpdateTargetDetectorRecipeDetectorRule operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudguard/UpdateTargetDetectorRecipeDetectorRule.go.html to see an example of how to use UpdateTargetDetectorRecipeDetectorRuleRequest.
type UpdateTargetDetectorRecipeDetectorRuleRequest struct {

	// OCID of target
	TargetId *string `mandatory:"true" contributesTo:"path" name:"targetId"`

	// OCID of TargetDetectorRecipe
	TargetDetectorRecipeId *string `mandatory:"true" contributesTo:"path" name:"targetDetectorRecipeId"`

	// The id of DetectorRule
	DetectorRuleId *string `mandatory:"true" contributesTo:"path" name:"detectorRuleId"`

	// The details to be updated for DetectorRule.
	UpdateTargetDetectorRecipeDetectorRuleDetails `contributesTo:"body"`

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

func (request UpdateTargetDetectorRecipeDetectorRuleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateTargetDetectorRecipeDetectorRuleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateTargetDetectorRecipeDetectorRuleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateTargetDetectorRecipeDetectorRuleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// UpdateTargetDetectorRecipeDetectorRuleResponse wrapper for the UpdateTargetDetectorRecipeDetectorRule operation
type UpdateTargetDetectorRecipeDetectorRuleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The TargetDetectorRecipeDetectorRule instance
	TargetDetectorRecipeDetectorRule `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateTargetDetectorRecipeDetectorRuleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateTargetDetectorRecipeDetectorRuleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
