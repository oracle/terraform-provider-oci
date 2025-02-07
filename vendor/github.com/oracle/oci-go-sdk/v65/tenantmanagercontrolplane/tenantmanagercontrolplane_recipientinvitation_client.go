// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// RecipientInvitationClient a client for RecipientInvitation
type RecipientInvitationClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewRecipientInvitationClientWithConfigurationProvider Creates a new default RecipientInvitation client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewRecipientInvitationClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client RecipientInvitationClient, err error) {
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
	return newRecipientInvitationClientFromBaseClient(baseClient, provider)
}

// NewRecipientInvitationClientWithOboToken Creates a new default RecipientInvitation client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewRecipientInvitationClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client RecipientInvitationClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newRecipientInvitationClientFromBaseClient(baseClient, configProvider)
}

func newRecipientInvitationClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client RecipientInvitationClient, err error) {
	// RecipientInvitation service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("RecipientInvitation"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = RecipientInvitationClient{BaseClient: baseClient}
	client.BasePath = "20230401"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *RecipientInvitationClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("tenantmanagercontrolplane", "https://organizations.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *RecipientInvitationClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *RecipientInvitationClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AcceptRecipientInvitation Accepts a recipient invitation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/AcceptRecipientInvitation.go.html to see an example of how to use AcceptRecipientInvitation API.
func (client RecipientInvitationClient) AcceptRecipientInvitation(ctx context.Context, request AcceptRecipientInvitationRequest) (response AcceptRecipientInvitationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.acceptRecipientInvitation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AcceptRecipientInvitationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AcceptRecipientInvitationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AcceptRecipientInvitationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AcceptRecipientInvitationResponse")
	}
	return
}

// acceptRecipientInvitation implements the OCIOperation interface (enables retrying operations)
func (client RecipientInvitationClient) acceptRecipientInvitation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/recipientInvitations/{recipientInvitationId}/actions/accept", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AcceptRecipientInvitationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/RecipientInvitation/AcceptRecipientInvitation"
		err = common.PostProcessServiceError(err, "RecipientInvitation", "AcceptRecipientInvitation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetRecipientInvitation Gets information about the recipient invitation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/GetRecipientInvitation.go.html to see an example of how to use GetRecipientInvitation API.
func (client RecipientInvitationClient) GetRecipientInvitation(ctx context.Context, request GetRecipientInvitationRequest) (response GetRecipientInvitationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getRecipientInvitation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetRecipientInvitationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetRecipientInvitationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetRecipientInvitationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetRecipientInvitationResponse")
	}
	return
}

// getRecipientInvitation implements the OCIOperation interface (enables retrying operations)
func (client RecipientInvitationClient) getRecipientInvitation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recipientInvitations/{recipientInvitationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetRecipientInvitationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/RecipientInvitation/GetRecipientInvitation"
		err = common.PostProcessServiceError(err, "RecipientInvitation", "GetRecipientInvitation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// IgnoreRecipientInvitation Ignores a recipient invitation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/IgnoreRecipientInvitation.go.html to see an example of how to use IgnoreRecipientInvitation API.
func (client RecipientInvitationClient) IgnoreRecipientInvitation(ctx context.Context, request IgnoreRecipientInvitationRequest) (response IgnoreRecipientInvitationResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.ignoreRecipientInvitation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = IgnoreRecipientInvitationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = IgnoreRecipientInvitationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(IgnoreRecipientInvitationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into IgnoreRecipientInvitationResponse")
	}
	return
}

// ignoreRecipientInvitation implements the OCIOperation interface (enables retrying operations)
func (client RecipientInvitationClient) ignoreRecipientInvitation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/recipientInvitations/{recipientInvitationId}/actions/ignore", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response IgnoreRecipientInvitationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/RecipientInvitation/IgnoreRecipientInvitation"
		err = common.PostProcessServiceError(err, "RecipientInvitation", "IgnoreRecipientInvitation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListRecipientInvitations Return a (paginated) list of recipient invitations.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/ListRecipientInvitations.go.html to see an example of how to use ListRecipientInvitations API.
func (client RecipientInvitationClient) ListRecipientInvitations(ctx context.Context, request ListRecipientInvitationsRequest) (response ListRecipientInvitationsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listRecipientInvitations, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListRecipientInvitationsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListRecipientInvitationsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListRecipientInvitationsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListRecipientInvitationsResponse")
	}
	return
}

// listRecipientInvitations implements the OCIOperation interface (enables retrying operations)
func (client RecipientInvitationClient) listRecipientInvitations(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/recipientInvitations", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListRecipientInvitationsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/RecipientInvitation/ListRecipientInvitations"
		err = common.PostProcessServiceError(err, "RecipientInvitation", "ListRecipientInvitations", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateRecipientInvitation Updates the recipient invitation.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/tenantmanagercontrolplane/UpdateRecipientInvitation.go.html to see an example of how to use UpdateRecipientInvitation API.
func (client RecipientInvitationClient) UpdateRecipientInvitation(ctx context.Context, request UpdateRecipientInvitationRequest) (response UpdateRecipientInvitationResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateRecipientInvitation, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateRecipientInvitationResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateRecipientInvitationResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateRecipientInvitationResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateRecipientInvitationResponse")
	}
	return
}

// updateRecipientInvitation implements the OCIOperation interface (enables retrying operations)
func (client RecipientInvitationClient) updateRecipientInvitation(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/recipientInvitations/{recipientInvitationId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateRecipientInvitationResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/organizations/20230401/RecipientInvitation/UpdateRecipientInvitation"
		err = common.PostProcessServiceError(err, "RecipientInvitation", "UpdateRecipientInvitation", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
