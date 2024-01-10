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

// EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest wrapper for the EnableHighFrequencyAutomaticSpmEvolveAdvisorTask operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemanagement/EnableHighFrequencyAutomaticSpmEvolveAdvisorTask.go.html to see an example of how to use EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest.
type EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Managed Database.
	ManagedDatabaseId *string `mandatory:"true" contributesTo:"path" name:"managedDatabaseId"`

	// The details required to enable high frequency Automatic SPM Evolve Advisor task.
	EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskDetails `contributesTo:"body"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse wrapper for the EnableHighFrequencyAutomaticSpmEvolveAdvisorTask operation
type EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response EnableHighFrequencyAutomaticSpmEvolveAdvisorTaskResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
