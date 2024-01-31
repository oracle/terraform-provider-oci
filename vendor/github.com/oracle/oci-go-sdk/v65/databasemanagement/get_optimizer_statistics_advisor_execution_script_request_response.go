// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetOptimizerStatisticsAdvisorExecutionScriptRequest wrapper for the GetOptimizerStatisticsAdvisorExecutionScript operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/GetOptimizerStatisticsAdvisorExecutionScript.go.html to see an example of how to use GetOptimizerStatisticsAdvisorExecutionScriptRequest.
type GetOptimizerStatisticsAdvisorExecutionScriptRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The name of the Optimizer Statistics Advisor execution.
	ExecutionName *string `mandatory:"true" contributesTo:"path" name:"executionName"`

	// The name of the optimizer statistics collection execution task.
	TaskName *string `mandatory:"true" contributesTo:"query" name:"taskName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The OCID of the Named Credential.
	OpcNamedCredentialId *string `mandatory:"false" contributesTo:"header" name:"opc-named-credential-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetOptimizerStatisticsAdvisorExecutionScriptRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetOptimizerStatisticsAdvisorExecutionScriptRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetOptimizerStatisticsAdvisorExecutionScriptRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetOptimizerStatisticsAdvisorExecutionScriptRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetOptimizerStatisticsAdvisorExecutionScriptRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetOptimizerStatisticsAdvisorExecutionScriptResponse wrapper for the GetOptimizerStatisticsAdvisorExecutionScript operation
type GetOptimizerStatisticsAdvisorExecutionScriptResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OptimizerStatisticsAdvisorExecutionScript instance
	OptimizerStatisticsAdvisorExecutionScript `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetOptimizerStatisticsAdvisorExecutionScriptResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetOptimizerStatisticsAdvisorExecutionScriptResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
