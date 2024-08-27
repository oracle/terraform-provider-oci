// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// DelegateAccessControlClient a client for DelegateAccessControl
type DelegateAccessControlClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewDelegateAccessControlClientWithConfigurationProvider Creates a new default DelegateAccessControl client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewDelegateAccessControlClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client DelegateAccessControlClient, err error) {
	if enabled := common.CheckForEnabledServices("delegateaccesscontrol"); !enabled {
		return client, fmt.Errorf("the Developer Tool configuration disabled this service, this behavior is controlled by OciSdkEnabledServicesMap variables. Please check if your local developer-tool-configuration.json file configured the service you're targeting or contact the cloud provider on the availability of this service")
	}
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newDelegateAccessControlClientFromBaseClient(baseClient, provider)
}

// NewDelegateAccessControlClientWithOboToken Creates a new default DelegateAccessControl client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewDelegateAccessControlClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client DelegateAccessControlClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newDelegateAccessControlClientFromBaseClient(baseClient, configProvider)
}

func newDelegateAccessControlClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client DelegateAccessControlClient, err error) {
	// DelegateAccessControl service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("DelegateAccessControl"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = DelegateAccessControlClient{BaseClient: baseClient}
	client.BasePath = "20230801"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *DelegateAccessControlClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("delegateaccesscontrol", "https://delegate-access-control.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *DelegateAccessControlClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *DelegateAccessControlClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ApproveDelegatedResourceAccessRequest Approves a Delegated Resource Access Request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ApproveDelegatedResourceAccessRequest.go.html to see an example of how to use ApproveDelegatedResourceAccessRequest API.
// A default retry strategy applies to this operation ApproveDelegatedResourceAccessRequest()
func (client DelegateAccessControlClient) ApproveDelegatedResourceAccessRequest(ctx context.Context, request ApproveDelegatedResourceAccessRequestRequest) (response ApproveDelegatedResourceAccessRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.approveDelegatedResourceAccessRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ApproveDelegatedResourceAccessRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ApproveDelegatedResourceAccessRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ApproveDelegatedResourceAccessRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ApproveDelegatedResourceAccessRequestResponse")
	}
	return
}

// approveDelegatedResourceAccessRequest implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) approveDelegatedResourceAccessRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/delegatedResourceAccessRequests/{delegatedResourceAccessRequestId}/actions/approve", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ApproveDelegatedResourceAccessRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ApproveDelegatedResourceAccessRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDelegationControlCompartment Moves the Delegation Control resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ChangeDelegationControlCompartment.go.html to see an example of how to use ChangeDelegationControlCompartment API.
// A default retry strategy applies to this operation ChangeDelegationControlCompartment()
func (client DelegateAccessControlClient) ChangeDelegationControlCompartment(ctx context.Context, request ChangeDelegationControlCompartmentRequest) (response ChangeDelegationControlCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeDelegationControlCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDelegationControlCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDelegationControlCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDelegationControlCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDelegationControlCompartmentResponse")
	}
	return
}

// changeDelegationControlCompartment implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) changeDelegationControlCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/delegationControls/{delegationControlId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDelegationControlCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ChangeDelegationControlCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeDelegationSubscriptionCompartment Moves the Delegation Subscription resource into a different compartment. When provided, 'If-Match' is checked against 'ETag' values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ChangeDelegationSubscriptionCompartment.go.html to see an example of how to use ChangeDelegationSubscriptionCompartment API.
// A default retry strategy applies to this operation ChangeDelegationSubscriptionCompartment()
func (client DelegateAccessControlClient) ChangeDelegationSubscriptionCompartment(ctx context.Context, request ChangeDelegationSubscriptionCompartmentRequest) (response ChangeDelegationSubscriptionCompartmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.changeDelegationSubscriptionCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeDelegationSubscriptionCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeDelegationSubscriptionCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeDelegationSubscriptionCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeDelegationSubscriptionCompartmentResponse")
	}
	return
}

// changeDelegationSubscriptionCompartment implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) changeDelegationSubscriptionCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/delegationSubscriptions/{delegationSubscriptionId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeDelegationSubscriptionCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ChangeDelegationSubscriptionCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDelegationControl Creates a Delegation Control.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/CreateDelegationControl.go.html to see an example of how to use CreateDelegationControl API.
// A default retry strategy applies to this operation CreateDelegationControl()
func (client DelegateAccessControlClient) CreateDelegationControl(ctx context.Context, request CreateDelegationControlRequest) (response CreateDelegationControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDelegationControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDelegationControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDelegationControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDelegationControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDelegationControlResponse")
	}
	return
}

// createDelegationControl implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) createDelegationControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/delegationControls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDelegationControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "CreateDelegationControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateDelegationSubscription Creates Delegation Subscription in Delegation Control.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/CreateDelegationSubscription.go.html to see an example of how to use CreateDelegationSubscription API.
// A default retry strategy applies to this operation CreateDelegationSubscription()
func (client DelegateAccessControlClient) CreateDelegationSubscription(ctx context.Context, request CreateDelegationSubscriptionRequest) (response CreateDelegationSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDelegationSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDelegationSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDelegationSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDelegationSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDelegationSubscriptionResponse")
	}
	return
}

// createDelegationSubscription implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) createDelegationSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/delegationSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDelegationSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "CreateDelegationSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDelegationControl Deletes a Delegation Control. You cannot delete a Delegation Control if it is assigned to govern any target resource currently or in the future.
// In that case, first, delete all of the current and future assignments before deleting the Delegation Control. A Delegation Control that was previously assigned to a target
// resource is marked as DELETED following a successful deletion. However, it is not completely deleted from the system. This is to ensure auditing information for the accesses
// done under the Delegation Control is preserved for future needs. The system purges the deleted Delegation Control only when all of the audit data associated with the
// Delegation Control are also deleted. Therefore, you cannot reuse the name of the deleted Delegation Control until the system purges the Delegation Control.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/DeleteDelegationControl.go.html to see an example of how to use DeleteDelegationControl API.
// A default retry strategy applies to this operation DeleteDelegationControl()
func (client DelegateAccessControlClient) DeleteDelegationControl(ctx context.Context, request DeleteDelegationControlRequest) (response DeleteDelegationControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDelegationControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDelegationControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDelegationControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDelegationControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDelegationControlResponse")
	}
	return
}

// deleteDelegationControl implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) deleteDelegationControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/delegationControls/{delegationControlId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDelegationControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "DeleteDelegationControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDelegationSubscription eletes an Delegation Subscription in Delegation Control.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/DeleteDelegationSubscription.go.html to see an example of how to use DeleteDelegationSubscription API.
// A default retry strategy applies to this operation DeleteDelegationSubscription()
func (client DelegateAccessControlClient) DeleteDelegationSubscription(ctx context.Context, request DeleteDelegationSubscriptionRequest) (response DeleteDelegationSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDelegationSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDelegationSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDelegationSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDelegationSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDelegationSubscriptionResponse")
	}
	return
}

// deleteDelegationSubscription implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) deleteDelegationSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/delegationSubscriptions/{delegationSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDelegationSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "DeleteDelegationSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDelegatedResourceAccessRequest Gets details of a Delegated Resource Access Request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/GetDelegatedResourceAccessRequest.go.html to see an example of how to use GetDelegatedResourceAccessRequest API.
// A default retry strategy applies to this operation GetDelegatedResourceAccessRequest()
func (client DelegateAccessControlClient) GetDelegatedResourceAccessRequest(ctx context.Context, request GetDelegatedResourceAccessRequestRequest) (response GetDelegatedResourceAccessRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDelegatedResourceAccessRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDelegatedResourceAccessRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDelegatedResourceAccessRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDelegatedResourceAccessRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDelegatedResourceAccessRequestResponse")
	}
	return
}

// getDelegatedResourceAccessRequest implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) getDelegatedResourceAccessRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegatedResourceAccessRequests/{delegatedResourceAccessRequestId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDelegatedResourceAccessRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "GetDelegatedResourceAccessRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDelegatedResourceAccessRequestAuditLogReport Gets the audit log report for the given Delegated Resource Access Request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/GetDelegatedResourceAccessRequestAuditLogReport.go.html to see an example of how to use GetDelegatedResourceAccessRequestAuditLogReport API.
// A default retry strategy applies to this operation GetDelegatedResourceAccessRequestAuditLogReport()
func (client DelegateAccessControlClient) GetDelegatedResourceAccessRequestAuditLogReport(ctx context.Context, request GetDelegatedResourceAccessRequestAuditLogReportRequest) (response GetDelegatedResourceAccessRequestAuditLogReportResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDelegatedResourceAccessRequestAuditLogReport, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDelegatedResourceAccessRequestAuditLogReportResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDelegatedResourceAccessRequestAuditLogReportResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDelegatedResourceAccessRequestAuditLogReportResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDelegatedResourceAccessRequestAuditLogReportResponse")
	}
	return
}

// getDelegatedResourceAccessRequestAuditLogReport implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) getDelegatedResourceAccessRequestAuditLogReport(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegatedResourceAccessRequests/{delegatedResourceAccessRequestId}/delegatedResourceAccessRequestAuditLogReport", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDelegatedResourceAccessRequestAuditLogReportResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "GetDelegatedResourceAccessRequestAuditLogReport", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDelegationControl Gets the Delegation Control associated with the specified Delegation Control ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/GetDelegationControl.go.html to see an example of how to use GetDelegationControl API.
// A default retry strategy applies to this operation GetDelegationControl()
func (client DelegateAccessControlClient) GetDelegationControl(ctx context.Context, request GetDelegationControlRequest) (response GetDelegationControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDelegationControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDelegationControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDelegationControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDelegationControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDelegationControlResponse")
	}
	return
}

// getDelegationControl implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) getDelegationControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegationControls/{delegationControlId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDelegationControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "GetDelegationControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDelegationSubscription Gets a DelegationSubscription by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/GetDelegationSubscription.go.html to see an example of how to use GetDelegationSubscription API.
// A default retry strategy applies to this operation GetDelegationSubscription()
func (client DelegateAccessControlClient) GetDelegationSubscription(ctx context.Context, request GetDelegationSubscriptionRequest) (response GetDelegationSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDelegationSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDelegationSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDelegationSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDelegationSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDelegationSubscriptionResponse")
	}
	return
}

// getDelegationSubscription implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) getDelegationSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegationSubscriptions/{delegationSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDelegationSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "GetDelegationSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetServiceProvider Gets a ServiceProvider by identifier
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/GetServiceProvider.go.html to see an example of how to use GetServiceProvider API.
// A default retry strategy applies to this operation GetServiceProvider()
func (client DelegateAccessControlClient) GetServiceProvider(ctx context.Context, request GetServiceProviderRequest) (response GetServiceProviderResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getServiceProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceProviderResponse")
	}
	return
}

// getServiceProvider implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) getServiceProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceProviders/{serviceProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetServiceProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "GetServiceProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetServiceProviderAction Gets the Service Provider Action associated with the specified Service Provider Action ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/GetServiceProviderAction.go.html to see an example of how to use GetServiceProviderAction API.
// A default retry strategy applies to this operation GetServiceProviderAction()
func (client DelegateAccessControlClient) GetServiceProviderAction(ctx context.Context, request GetServiceProviderActionRequest) (response GetServiceProviderActionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getServiceProviderAction, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetServiceProviderActionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetServiceProviderActionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetServiceProviderActionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetServiceProviderActionResponse")
	}
	return
}

// getServiceProviderAction implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) getServiceProviderAction(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceProviderActions/{serviceProviderActionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetServiceProviderActionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "GetServiceProviderAction", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDelegatedResourceAccessRequestHistories Returns a history of all status associated with the Delegated Resource Access RequestId.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListDelegatedResourceAccessRequestHistories.go.html to see an example of how to use ListDelegatedResourceAccessRequestHistories API.
// A default retry strategy applies to this operation ListDelegatedResourceAccessRequestHistories()
func (client DelegateAccessControlClient) ListDelegatedResourceAccessRequestHistories(ctx context.Context, request ListDelegatedResourceAccessRequestHistoriesRequest) (response ListDelegatedResourceAccessRequestHistoriesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDelegatedResourceAccessRequestHistories, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDelegatedResourceAccessRequestHistoriesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDelegatedResourceAccessRequestHistoriesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDelegatedResourceAccessRequestHistoriesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDelegatedResourceAccessRequestHistoriesResponse")
	}
	return
}

// listDelegatedResourceAccessRequestHistories implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) listDelegatedResourceAccessRequestHistories(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegatedResourceAccessRequests/{delegatedResourceAccessRequestId}/history", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDelegatedResourceAccessRequestHistoriesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ListDelegatedResourceAccessRequestHistories", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDelegatedResourceAccessRequests Lists all Delegated Resource Access Requests in the compartment. Note that only one of lifecycleState or requestStatus query parameter can be used.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListDelegatedResourceAccessRequests.go.html to see an example of how to use ListDelegatedResourceAccessRequests API.
// A default retry strategy applies to this operation ListDelegatedResourceAccessRequests()
func (client DelegateAccessControlClient) ListDelegatedResourceAccessRequests(ctx context.Context, request ListDelegatedResourceAccessRequestsRequest) (response ListDelegatedResourceAccessRequestsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDelegatedResourceAccessRequests, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDelegatedResourceAccessRequestsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDelegatedResourceAccessRequestsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDelegatedResourceAccessRequestsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDelegatedResourceAccessRequestsResponse")
	}
	return
}

// listDelegatedResourceAccessRequests implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) listDelegatedResourceAccessRequests(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegatedResourceAccessRequests", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDelegatedResourceAccessRequestsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ListDelegatedResourceAccessRequests", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDelegationControlResources Returns a list of resources associated with the Delegation Control.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListDelegationControlResources.go.html to see an example of how to use ListDelegationControlResources API.
// A default retry strategy applies to this operation ListDelegationControlResources()
func (client DelegateAccessControlClient) ListDelegationControlResources(ctx context.Context, request ListDelegationControlResourcesRequest) (response ListDelegationControlResourcesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDelegationControlResources, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDelegationControlResourcesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDelegationControlResourcesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDelegationControlResourcesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDelegationControlResourcesResponse")
	}
	return
}

// listDelegationControlResources implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) listDelegationControlResources(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegationControls/{delegationControlId}/resources", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDelegationControlResourcesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ListDelegationControlResources", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDelegationControls Lists the Delegation Controls in the compartment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListDelegationControls.go.html to see an example of how to use ListDelegationControls API.
// A default retry strategy applies to this operation ListDelegationControls()
func (client DelegateAccessControlClient) ListDelegationControls(ctx context.Context, request ListDelegationControlsRequest) (response ListDelegationControlsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDelegationControls, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDelegationControlsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDelegationControlsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDelegationControlsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDelegationControlsResponse")
	}
	return
}

// listDelegationControls implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) listDelegationControls(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegationControls", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDelegationControlsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ListDelegationControls", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDelegationSubscriptions Lists the Delegation Subscriptions in Delegation Control.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListDelegationSubscriptions.go.html to see an example of how to use ListDelegationSubscriptions API.
// A default retry strategy applies to this operation ListDelegationSubscriptions()
func (client DelegateAccessControlClient) ListDelegationSubscriptions(ctx context.Context, request ListDelegationSubscriptionsRequest) (response ListDelegationSubscriptionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDelegationSubscriptions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDelegationSubscriptionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDelegationSubscriptionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDelegationSubscriptionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDelegationSubscriptionsResponse")
	}
	return
}

// listDelegationSubscriptions implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) listDelegationSubscriptions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegationSubscriptions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDelegationSubscriptionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ListDelegationSubscriptions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServiceProviderActions Lists all the ServiceProviderActions available in the system.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListServiceProviderActions.go.html to see an example of how to use ListServiceProviderActions API.
// A default retry strategy applies to this operation ListServiceProviderActions()
func (client DelegateAccessControlClient) ListServiceProviderActions(ctx context.Context, request ListServiceProviderActionsRequest) (response ListServiceProviderActionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceProviderActions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceProviderActionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceProviderActionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceProviderActionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceProviderActionsResponse")
	}
	return
}

// listServiceProviderActions implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) listServiceProviderActions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceProviderActions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServiceProviderActionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ListServiceProviderActions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServiceProviderInteractions Lists the MoreInformation interaction between customer and support operators.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListServiceProviderInteractions.go.html to see an example of how to use ListServiceProviderInteractions API.
// A default retry strategy applies to this operation ListServiceProviderInteractions()
func (client DelegateAccessControlClient) ListServiceProviderInteractions(ctx context.Context, request ListServiceProviderInteractionsRequest) (response ListServiceProviderInteractionsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceProviderInteractions, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceProviderInteractionsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceProviderInteractionsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceProviderInteractionsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceProviderInteractionsResponse")
	}
	return
}

// listServiceProviderInteractions implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) listServiceProviderInteractions(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/delegatedResourceAccessRequests/{delegatedResourceAccessRequestId}/serviceProviderInteractions", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServiceProviderInteractionsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ListServiceProviderInteractions", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListServiceProviders Lists the Service Providers.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ListServiceProviders.go.html to see an example of how to use ListServiceProviders API.
// A default retry strategy applies to this operation ListServiceProviders()
func (client DelegateAccessControlClient) ListServiceProviders(ctx context.Context, request ListServiceProvidersRequest) (response ListServiceProvidersResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listServiceProviders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListServiceProvidersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListServiceProvidersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListServiceProvidersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListServiceProvidersResponse")
	}
	return
}

// listServiceProviders implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) listServiceProviders(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/serviceProviders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListServiceProvidersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ListServiceProviders", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RejectDelegatedResourceAccessRequest Rejects a Delegated Resource Access Request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/RejectDelegatedResourceAccessRequest.go.html to see an example of how to use RejectDelegatedResourceAccessRequest API.
// A default retry strategy applies to this operation RejectDelegatedResourceAccessRequest()
func (client DelegateAccessControlClient) RejectDelegatedResourceAccessRequest(ctx context.Context, request RejectDelegatedResourceAccessRequestRequest) (response RejectDelegatedResourceAccessRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.rejectDelegatedResourceAccessRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RejectDelegatedResourceAccessRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RejectDelegatedResourceAccessRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RejectDelegatedResourceAccessRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RejectDelegatedResourceAccessRequestResponse")
	}
	return
}

// rejectDelegatedResourceAccessRequest implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) rejectDelegatedResourceAccessRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/delegatedResourceAccessRequests/{delegatedResourceAccessRequestId}/actions/reject", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RejectDelegatedResourceAccessRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "RejectDelegatedResourceAccessRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RevokeDelegatedResourceAccessRequest Revokes an already approved Delegated Resource Access Request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/RevokeDelegatedResourceAccessRequest.go.html to see an example of how to use RevokeDelegatedResourceAccessRequest API.
// A default retry strategy applies to this operation RevokeDelegatedResourceAccessRequest()
func (client DelegateAccessControlClient) RevokeDelegatedResourceAccessRequest(ctx context.Context, request RevokeDelegatedResourceAccessRequestRequest) (response RevokeDelegatedResourceAccessRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.revokeDelegatedResourceAccessRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RevokeDelegatedResourceAccessRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RevokeDelegatedResourceAccessRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RevokeDelegatedResourceAccessRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RevokeDelegatedResourceAccessRequestResponse")
	}
	return
}

// revokeDelegatedResourceAccessRequest implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) revokeDelegatedResourceAccessRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/delegatedResourceAccessRequests/{delegatedResourceAccessRequestId}/actions/revoke", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RevokeDelegatedResourceAccessRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "RevokeDelegatedResourceAccessRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ServiceProviderInteractionRequest Posts query for additional information for the given Delegated Resource Access Request.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/ServiceProviderInteractionRequest.go.html to see an example of how to use ServiceProviderInteractionRequest API.
// A default retry strategy applies to this operation ServiceProviderInteractionRequest()
func (client DelegateAccessControlClient) ServiceProviderInteractionRequest(ctx context.Context, request ServiceProviderInteractionRequestRequest) (response ServiceProviderInteractionRequestResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.serviceProviderInteractionRequest, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ServiceProviderInteractionRequestResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ServiceProviderInteractionRequestResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ServiceProviderInteractionRequestResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ServiceProviderInteractionRequestResponse")
	}
	return
}

// serviceProviderInteractionRequest implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) serviceProviderInteractionRequest(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/delegatedResourceAccessRequests/{delegatedResourceAccessRequestId}/actions/serviceProviderInteractionRequest", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ServiceProviderInteractionRequestResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "ServiceProviderInteractionRequest", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDelegationControl Updates the existing DelegationControl for a given Delegation Control ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/UpdateDelegationControl.go.html to see an example of how to use UpdateDelegationControl API.
// A default retry strategy applies to this operation UpdateDelegationControl()
func (client DelegateAccessControlClient) UpdateDelegationControl(ctx context.Context, request UpdateDelegationControlRequest) (response UpdateDelegationControlResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDelegationControl, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDelegationControlResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDelegationControlResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDelegationControlResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDelegationControlResponse")
	}
	return
}

// updateDelegationControl implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) updateDelegationControl(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/delegationControls/{delegationControlId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDelegationControlResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "UpdateDelegationControl", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDelegationSubscription Updates the existing DelegationSubscription for a given Delegation Subscription ID.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/delegateaccesscontrol/UpdateDelegationSubscription.go.html to see an example of how to use UpdateDelegationSubscription API.
// A default retry strategy applies to this operation UpdateDelegationSubscription()
func (client DelegateAccessControlClient) UpdateDelegationSubscription(ctx context.Context, request UpdateDelegationSubscriptionRequest) (response UpdateDelegationSubscriptionResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDelegationSubscription, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDelegationSubscriptionResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDelegationSubscriptionResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDelegationSubscriptionResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDelegationSubscriptionResponse")
	}
	return
}

// updateDelegationSubscription implements the OCIOperation interface (enables retrying operations)
func (client DelegateAccessControlClient) updateDelegationSubscription(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/delegationSubscriptions/{delegationSubscriptionId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDelegationSubscriptionResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "DelegateAccessControl", "UpdateDelegationSubscription", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
