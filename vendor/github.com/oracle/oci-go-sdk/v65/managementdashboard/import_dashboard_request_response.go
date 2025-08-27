// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package managementdashboard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ImportDashboardRequest wrapper for the ImportDashboard operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/managementdashboard/ImportDashboard.go.html to see an example of how to use ImportDashboardRequest.
type ImportDashboardRequest struct {

	// JSON metadata for importing dashboards and their saved searches.
	ManagementDashboardImportDetails `contributesTo:"body"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// By default, if a resource with the same OCID exists in the target compartment, it is updated during the import process, otherwise, a new resource is created.
	// However, if this attribute is set to true, then during the import process if a resource with the same displayName exists in the compartment, then it is updated even if the OCIDs are different. This is useful when importing the same resource multiple times. If the compartment and displayName remain the same, the resource is only updated and multiple copies of a resource are not created.
	OverrideSameName *string `mandatory:"false" contributesTo:"query" name:"overrideSameName"`

	// If this attribute is set, the dashboard resources are created or updated in the compartment specified by OCID. If this attribute is not set, the compartment specified in the JSON metadata is used.
	OverrideDashboardCompartmentOcid *string `mandatory:"false" contributesTo:"query" name:"overrideDashboardCompartmentOcid"`

	// If this attribute is set, the saved search resources are created or updated in the compartment specified by OCID. If this attribute is not set, the compartment specified in the JSON metadata is used.
	OverrideSavedSearchCompartmentOcid *string `mandatory:"false" contributesTo:"query" name:"overrideSavedSearchCompartmentOcid"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ImportDashboardRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ImportDashboardRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ImportDashboardRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ImportDashboardRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ImportDashboardRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ImportDashboardResponse wrapper for the ImportDashboard operation
type ImportDashboardResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the asynchronous request. You can use this to query status of the asynchronous operation.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ImportDashboardResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ImportDashboardResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
