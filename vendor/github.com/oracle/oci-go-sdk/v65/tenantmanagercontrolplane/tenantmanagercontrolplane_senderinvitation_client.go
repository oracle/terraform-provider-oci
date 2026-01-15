// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// SenderInvitationClient a client for SenderInvitation
type SenderInvitationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewSenderInvitationClientWithConfigurationProvider Creates a new default SenderInvitation client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewSenderInvitationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client SenderInvitationClient, err error) {
	if enabled := common.CheckForEnabledServices("tenantmanagercontrolplane"); !enabled {
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
	return newSenderInvitationClientFromBaseClient(baseClient, provider)
}

// NewSenderInvitationClientWithOboToken Creates a new default SenderInvitation client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewSenderInvitationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client SenderInvitationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newSenderInvitationClientFromBaseClient(baseClient, configProvider)
}

func newSenderInvitationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client SenderInvitationClient, err error) {
	// SenderInvitation service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("SenderInvitation"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = SenderInvitationClient{BaseClient: baseClient}
	client.BasePath = "20230401"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *SenderInvitationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *SenderInvitationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *SenderInvitationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// CancelSenderInvitation Cancels a sender invitation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/CancelSenderInvitation.go.html to see an example of how to use CancelSenderInvitation API.
func (client SenderInvitationClient) CancelSenderInvitation(ctx context.Context, request CancelSenderInvitationRequest) (response CancelSenderInvitationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.cancelSenderInvitation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CancelSenderInvitationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CancelSenderInvitationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CancelSenderInvitationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CancelSenderInvitationResponse")
	}
	return
}

// cancelSenderInvitation implements the OCIOperation interface (enables retrying operations)
func (client SenderInvitationClient) cancelSenderInvitation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/senderInvitations/{senderInvitationId}/actions/cancel", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CancelSenderInvitationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SenderInvitation/CancelSenderInvitation"
		err = common.PostProcessServiceError(err, "SenderInvitation", "CancelSenderInvitation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSenderInvitation Creates a sender invitation and asynchronously sends the invitation to the recipient.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/CreateSenderInvitation.go.html to see an example of how to use CreateSenderInvitation API.
func (client SenderInvitationClient) CreateSenderInvitation(ctx context.Context, request CreateSenderInvitationRequest) (response CreateSenderInvitationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSenderInvitation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSenderInvitationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSenderInvitationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSenderInvitationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSenderInvitationResponse")
	}
	return
}

// createSenderInvitation implements the OCIOperation interface (enables retrying operations)
func (client SenderInvitationClient) createSenderInvitation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/senderInvitations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSenderInvitationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SenderInvitation/CreateSenderInvitation"
		err = common.PostProcessServiceError(err, "SenderInvitation", "CreateSenderInvitation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSenderInvitation Gets information about the sender invitation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetSenderInvitation.go.html to see an example of how to use GetSenderInvitation API.
func (client SenderInvitationClient) GetSenderInvitation(ctx context.Context, request GetSenderInvitationRequest) (response GetSenderInvitationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSenderInvitation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSenderInvitationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSenderInvitationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSenderInvitationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSenderInvitationResponse")
	}
	return
}

// getSenderInvitation implements the OCIOperation interface (enables retrying operations)
func (client SenderInvitationClient) getSenderInvitation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/senderInvitations/{senderInvitationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSenderInvitationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SenderInvitation/GetSenderInvitation"
		err = common.PostProcessServiceError(err, "SenderInvitation", "GetSenderInvitation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSenderInvitations Return a (paginated) list of sender invitations.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListSenderInvitations.go.html to see an example of how to use ListSenderInvitations API.
func (client SenderInvitationClient) ListSenderInvitations(ctx context.Context, request ListSenderInvitationsRequest) (response ListSenderInvitationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSenderInvitations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSenderInvitationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSenderInvitationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSenderInvitationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSenderInvitationsResponse")
	}
	return
}

// listSenderInvitations implements the OCIOperation interface (enables retrying operations)
func (client SenderInvitationClient) listSenderInvitations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/senderInvitations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSenderInvitationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SenderInvitation/ListSenderInvitations"
		err = common.PostProcessServiceError(err, "SenderInvitation", "ListSenderInvitations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSenderInvitation Updates the sender invitation.
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/UpdateSenderInvitation.go.html to see an example of how to use UpdateSenderInvitation API.
func (client SenderInvitationClient) UpdateSenderInvitation(ctx context.Context, request UpdateSenderInvitationRequest) (response UpdateSenderInvitationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSenderInvitation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSenderInvitationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSenderInvitationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSenderInvitationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSenderInvitationResponse")
	}
	return
}

// updateSenderInvitation implements the OCIOperation interface (enables retrying operations)
func (client SenderInvitationClient) updateSenderInvitation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/senderInvitations/{senderInvitationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSenderInvitationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/SenderInvitation/UpdateSenderInvitation"
		err = common.PostProcessServiceError(err, "SenderInvitation", "UpdateSenderInvitation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
