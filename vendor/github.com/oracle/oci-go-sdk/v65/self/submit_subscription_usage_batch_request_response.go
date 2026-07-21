// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package self

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"io"
	"net/http"
	"strings"
)

// SubmitSubscriptionUsageBatchRequest wrapper for the SubmitSubscriptionUsageBatch operation
type SubmitSubscriptionUsageBatchRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment associated with the usage records request.
	CompartmentId *string `mandatory:"true" contributesTo:"header" name:"compartment-id"`

	// UTF-8 CSV file with no more than 10,000 usage records and a maximum size of 50 MB.
	// Required columns are `MarketplaceOfferId`, `Id`, `Amount`, `CurrencyCode`,
	// `UsageStartTime`, `UsageEndTime`, and `UsageDimensionName`.
	// Optional columns are `ConsumedQuantity`, `CustomerTenancyId`,
	// `BillingIdentifier`, `ProductSku`, `UnitOfMeasure`, `UnitPrice`,
	// `ContractDuration`, and `AdditionalMetadata`.
	// When provided in CSV, `AdditionalMetadata` must be a JSON array of `ExtendedMetadata`
	// objects. `MarketplaceOfferId` must be a subscription or private offer OCID.
	SubmitSubscriptionUsageBatchDetails io.ReadCloser `mandatory:"true" contributesTo:"body" encoding:"binary"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of running that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and removed from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request SubmitSubscriptionUsageBatchRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request SubmitSubscriptionUsageBatchRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {
	httpRequest, err := common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
	if err == nil && binaryRequestBody.Seekable() {
		common.UpdateRequestBinaryBody(&httpRequest, binaryRequestBody)
	}
	return httpRequest, err
}

// BinaryRequestBody implements the OCIRequest interface
func (request SubmitSubscriptionUsageBatchRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {
	rsc := common.NewOCIReadSeekCloser(request.SubmitSubscriptionUsageBatchDetails)
	if rsc.Seekable() {
		return rsc, true
	}
	return nil, true

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request SubmitSubscriptionUsageBatchRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request SubmitSubscriptionUsageBatchRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubmitSubscriptionUsageBatchResponse wrapper for the SubmitSubscriptionUsageBatch operation
type SubmitSubscriptionUsageBatchResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the asynchronous work request.
	// Use GetWorkRequest with this ID to track the status of the request.
	OpcWorkRequestId *string `presentIn:"header" name:"opc-work-request-id"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response SubmitSubscriptionUsageBatchResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response SubmitSubscriptionUsageBatchResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}
