// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

//ManagementClient a client for Management
type ManagementClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewManagementClientWithConfigurationProvider Creates a new default Management client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewManagementClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ManagementClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newManagementClientFromBaseClient(baseClient, provider)
}

// NewManagementClientWithOboToken Creates a new default Management client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewManagementClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ManagementClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newManagementClientFromBaseClient(baseClient, configProvider)
}

func newManagementClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ManagementClient, err error) {
	// Management service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("Management"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ManagementClient{BaseClient: baseClient}
	client.BasePath = "20190506"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ManagementClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("oda", "https://digitalassistant-api.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ManagementClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	if client.Host == "" {
		return fmt.Errorf("Invalid region or Host. Endpoint cannot be constructed without endpointServiceName or serviceEndpointTemplate for a dotted region")
	}
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *ManagementClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ConfigureDigitalAssistantParameters This will store the provided parameters in the Digital Assistant instance and update any Digital Assistants with matching parameters.
func (client ManagementClient) ConfigureDigitalAssistantParameters(ctx context.Context, request ConfigureDigitalAssistantParametersRequest) (response ConfigureDigitalAssistantParametersResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.configureDigitalAssistantParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ConfigureDigitalAssistantParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ConfigureDigitalAssistantParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ConfigureDigitalAssistantParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ConfigureDigitalAssistantParametersResponse")
	}
	return
}

// configureDigitalAssistantParameters implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) configureDigitalAssistantParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/actions/configureDigitalAssistantParameters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ConfigureDigitalAssistantParametersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistantParameter/ConfigureDigitalAssistantParameters"
		err = common.PostProcessServiceError(err, "Management", "ConfigureDigitalAssistantParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAuthenticationProvider Creates a new Authentication Provider
func (client ManagementClient) CreateAuthenticationProvider(ctx context.Context, request CreateAuthenticationProviderRequest) (response CreateAuthenticationProviderResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createAuthenticationProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAuthenticationProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAuthenticationProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAuthenticationProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAuthenticationProviderResponse")
	}
	return
}

// createAuthenticationProvider implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) createAuthenticationProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/authenticationProviders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAuthenticationProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/AuthenticationProvider/CreateAuthenticationProvider"
		err = common.PostProcessServiceError(err, "Management", "CreateAuthenticationProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateChannel Creates a new Channel.
func (client ManagementClient) CreateChannel(ctx context.Context, request CreateChannelRequest) (response CreateChannelResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateChannelResponse")
	}
	return
}

// createChannel implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) createChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/channels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Channel/CreateChannel"
		err = common.PostProcessServiceError(err, "Management", "CreateChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &createchannelresult{})
	return response, err
}

// CreateDigitalAssistant Creates a new Digital Assistant.
func (client ManagementClient) CreateDigitalAssistant(ctx context.Context, request CreateDigitalAssistantRequest) (response CreateDigitalAssistantResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createDigitalAssistant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateDigitalAssistantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateDigitalAssistantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateDigitalAssistantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateDigitalAssistantResponse")
	}
	return
}

// createDigitalAssistant implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) createDigitalAssistant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/digitalAssistants", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateDigitalAssistantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistant/CreateDigitalAssistant"
		err = common.PostProcessServiceError(err, "Management", "CreateDigitalAssistant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSkill Creates a new Skill from scratch.
func (client ManagementClient) CreateSkill(ctx context.Context, request CreateSkillRequest) (response CreateSkillResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSkill, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSkillResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSkillResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSkillResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSkillResponse")
	}
	return
}

// createSkill implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) createSkill(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/skills", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSkillResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Skill/CreateSkill"
		err = common.PostProcessServiceError(err, "Management", "CreateSkill", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateSkillParameter Creates a new Skill Parameter.
func (client ManagementClient) CreateSkillParameter(ctx context.Context, request CreateSkillParameterRequest) (response CreateSkillParameterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createSkillParameter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateSkillParameterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateSkillParameterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateSkillParameterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateSkillParameterResponse")
	}
	return
}

// createSkillParameter implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) createSkillParameter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/skills/{skillId}/parameters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateSkillParameterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/SkillParameter/CreateSkillParameter"
		err = common.PostProcessServiceError(err, "Management", "CreateSkillParameter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateTranslator Creates a new Translator
func (client ManagementClient) CreateTranslator(ctx context.Context, request CreateTranslatorRequest) (response CreateTranslatorResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.createTranslator, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateTranslatorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateTranslatorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateTranslatorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateTranslatorResponse")
	}
	return
}

// createTranslator implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) createTranslator(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/translators", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateTranslatorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Translator/CreateTranslator"
		err = common.PostProcessServiceError(err, "Management", "CreateTranslator", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAuthenticationProvider Delete the specified Authentication Provider.
func (client ManagementClient) DeleteAuthenticationProvider(ctx context.Context, request DeleteAuthenticationProviderRequest) (response DeleteAuthenticationProviderResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAuthenticationProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAuthenticationProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAuthenticationProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAuthenticationProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAuthenticationProviderResponse")
	}
	return
}

// deleteAuthenticationProvider implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) deleteAuthenticationProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/odaInstances/{odaInstanceId}/authenticationProviders/{authenticationProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAuthenticationProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/AuthenticationProvider/DeleteAuthenticationProvider"
		err = common.PostProcessServiceError(err, "Management", "DeleteAuthenticationProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteChannel Delete the specified Channel.
func (client ManagementClient) DeleteChannel(ctx context.Context, request DeleteChannelRequest) (response DeleteChannelResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteChannelResponse")
	}
	return
}

// deleteChannel implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) deleteChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/odaInstances/{odaInstanceId}/channels/{channelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Channel/DeleteChannel"
		err = common.PostProcessServiceError(err, "Management", "DeleteChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteDigitalAssistant Delete the specified Digital Assistant.
func (client ManagementClient) DeleteDigitalAssistant(ctx context.Context, request DeleteDigitalAssistantRequest) (response DeleteDigitalAssistantResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteDigitalAssistant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteDigitalAssistantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteDigitalAssistantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteDigitalAssistantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteDigitalAssistantResponse")
	}
	return
}

// deleteDigitalAssistant implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) deleteDigitalAssistant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/odaInstances/{odaInstanceId}/digitalAssistants/{digitalAssistantId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteDigitalAssistantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistant/DeleteDigitalAssistant"
		err = common.PostProcessServiceError(err, "Management", "DeleteDigitalAssistant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSkill Delete the specified Skill.
func (client ManagementClient) DeleteSkill(ctx context.Context, request DeleteSkillRequest) (response DeleteSkillResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSkill, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSkillResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSkillResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSkillResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSkillResponse")
	}
	return
}

// deleteSkill implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) deleteSkill(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/odaInstances/{odaInstanceId}/skills/{skillId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSkillResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Skill/DeleteSkill"
		err = common.PostProcessServiceError(err, "Management", "DeleteSkill", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteSkillParameter Delete the specified Skill Parameter.
func (client ManagementClient) DeleteSkillParameter(ctx context.Context, request DeleteSkillParameterRequest) (response DeleteSkillParameterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteSkillParameter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteSkillParameterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteSkillParameterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteSkillParameterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteSkillParameterResponse")
	}
	return
}

// deleteSkillParameter implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) deleteSkillParameter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/odaInstances/{odaInstanceId}/skills/{skillId}/parameters/{parameterName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteSkillParameterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/SkillParameter/DeleteSkillParameter"
		err = common.PostProcessServiceError(err, "Management", "DeleteSkillParameter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteTranslator Delete the specified Translator.
func (client ManagementClient) DeleteTranslator(ctx context.Context, request DeleteTranslatorRequest) (response DeleteTranslatorResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteTranslator, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteTranslatorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteTranslatorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteTranslatorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteTranslatorResponse")
	}
	return
}

// deleteTranslator implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) deleteTranslator(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/odaInstances/{odaInstanceId}/translators/{translatorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteTranslatorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Translator/DeleteTranslator"
		err = common.PostProcessServiceError(err, "Management", "DeleteTranslator", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportDigitalAssistant Exports the specified Digital Assistant as an archive to Object Storage.
func (client ManagementClient) ExportDigitalAssistant(ctx context.Context, request ExportDigitalAssistantRequest) (response ExportDigitalAssistantResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.exportDigitalAssistant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportDigitalAssistantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportDigitalAssistantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportDigitalAssistantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportDigitalAssistantResponse")
	}
	return
}

// exportDigitalAssistant implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) exportDigitalAssistant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/digitalAssistants/{digitalAssistantId}/actions/export", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportDigitalAssistantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Skill/ExportDigitalAssistant"
		err = common.PostProcessServiceError(err, "Management", "ExportDigitalAssistant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ExportSkill Exports the specified Skill as an archive to Object Storage.
func (client ManagementClient) ExportSkill(ctx context.Context, request ExportSkillRequest) (response ExportSkillResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.exportSkill, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ExportSkillResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ExportSkillResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ExportSkillResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ExportSkillResponse")
	}
	return
}

// exportSkill implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) exportSkill(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/skills/{skillId}/actions/export", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ExportSkillResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Skill/ExportSkill"
		err = common.PostProcessServiceError(err, "Management", "ExportSkill", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAuthenticationProvider Gets the specified Authentication Provider.
func (client ManagementClient) GetAuthenticationProvider(ctx context.Context, request GetAuthenticationProviderRequest) (response GetAuthenticationProviderResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAuthenticationProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAuthenticationProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAuthenticationProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAuthenticationProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAuthenticationProviderResponse")
	}
	return
}

// getAuthenticationProvider implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) getAuthenticationProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/authenticationProviders/{authenticationProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAuthenticationProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/AuthenticationProvider/GetAuthenticationProvider"
		err = common.PostProcessServiceError(err, "Management", "GetAuthenticationProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetChannel Gets the specified Channel.
func (client ManagementClient) GetChannel(ctx context.Context, request GetChannelRequest) (response GetChannelResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetChannelResponse")
	}
	return
}

// getChannel implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) getChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/channels/{channelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Channel/GetChannel"
		err = common.PostProcessServiceError(err, "Management", "GetChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &channel{})
	return response, err
}

// GetDigitalAssistant Gets the specified Digital Assistant.
func (client ManagementClient) GetDigitalAssistant(ctx context.Context, request GetDigitalAssistantRequest) (response GetDigitalAssistantResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDigitalAssistant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDigitalAssistantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDigitalAssistantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDigitalAssistantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDigitalAssistantResponse")
	}
	return
}

// getDigitalAssistant implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) getDigitalAssistant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/digitalAssistants/{digitalAssistantId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDigitalAssistantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistant/GetDigitalAssistant"
		err = common.PostProcessServiceError(err, "Management", "GetDigitalAssistant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetDigitalAssistantParameter Gets the specified Digital Assistant Parameter.
func (client ManagementClient) GetDigitalAssistantParameter(ctx context.Context, request GetDigitalAssistantParameterRequest) (response GetDigitalAssistantParameterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getDigitalAssistantParameter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetDigitalAssistantParameterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetDigitalAssistantParameterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetDigitalAssistantParameterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetDigitalAssistantParameterResponse")
	}
	return
}

// getDigitalAssistantParameter implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) getDigitalAssistantParameter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/digitalAssistants/{digitalAssistantId}/parameters/{parameterName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetDigitalAssistantParameterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistantParameter/GetDigitalAssistantParameter"
		err = common.PostProcessServiceError(err, "Management", "GetDigitalAssistantParameter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSkill Gets the specified Skill.
func (client ManagementClient) GetSkill(ctx context.Context, request GetSkillRequest) (response GetSkillResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSkill, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSkillResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSkillResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSkillResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSkillResponse")
	}
	return
}

// getSkill implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) getSkill(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/skills/{skillId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSkillResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Skill/GetSkill"
		err = common.PostProcessServiceError(err, "Management", "GetSkill", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetSkillParameter Gets the specified Skill Parameter.
func (client ManagementClient) GetSkillParameter(ctx context.Context, request GetSkillParameterRequest) (response GetSkillParameterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getSkillParameter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetSkillParameterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetSkillParameterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetSkillParameterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetSkillParameterResponse")
	}
	return
}

// getSkillParameter implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) getSkillParameter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/skills/{skillId}/parameters/{parameterName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetSkillParameterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/SkillParameter/GetSkillParameter"
		err = common.PostProcessServiceError(err, "Management", "GetSkillParameter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetTranslator Gets the specified Translator.
func (client ManagementClient) GetTranslator(ctx context.Context, request GetTranslatorRequest) (response GetTranslatorResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getTranslator, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetTranslatorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetTranslatorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetTranslatorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetTranslatorResponse")
	}
	return
}

// getTranslator implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) getTranslator(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/translators/{translatorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetTranslatorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Translator/GetTranslator"
		err = common.PostProcessServiceError(err, "Management", "GetTranslator", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ImportBot Import a Bot archive from Object Storage.
func (client ManagementClient) ImportBot(ctx context.Context, request ImportBotRequest) (response ImportBotResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.importBot, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ImportBotResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ImportBotResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ImportBotResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ImportBotResponse")
	}
	return
}

// importBot implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) importBot(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/actions/importBot", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ImportBotResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Bot/ImportBot"
		err = common.PostProcessServiceError(err, "Management", "ImportBot", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAuthenticationProviders Returns a page of Authentication Providers that belong to the specified Digital Assistant instance.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client ManagementClient) ListAuthenticationProviders(ctx context.Context, request ListAuthenticationProvidersRequest) (response ListAuthenticationProvidersResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAuthenticationProviders, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAuthenticationProvidersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAuthenticationProvidersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAuthenticationProvidersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAuthenticationProvidersResponse")
	}
	return
}

// listAuthenticationProviders implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) listAuthenticationProviders(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/authenticationProviders", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAuthenticationProvidersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/AuthenticationProvider/ListAuthenticationProviders"
		err = common.PostProcessServiceError(err, "Management", "ListAuthenticationProviders", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListChannels Returns a page of Channels that belong to the specified Digital Assistant instance.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client ManagementClient) ListChannels(ctx context.Context, request ListChannelsRequest) (response ListChannelsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listChannels, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListChannelsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListChannelsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListChannelsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListChannelsResponse")
	}
	return
}

// listChannels implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) listChannels(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/channels", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListChannelsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Channel/ListChannels"
		err = common.PostProcessServiceError(err, "Management", "ListChannels", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDigitalAssistantParameters Returns a page of Parameters that belong to the specified Digital Assistant.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client ManagementClient) ListDigitalAssistantParameters(ctx context.Context, request ListDigitalAssistantParametersRequest) (response ListDigitalAssistantParametersResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDigitalAssistantParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDigitalAssistantParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDigitalAssistantParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDigitalAssistantParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDigitalAssistantParametersResponse")
	}
	return
}

// listDigitalAssistantParameters implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) listDigitalAssistantParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/digitalAssistants/{digitalAssistantId}/parameters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDigitalAssistantParametersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistantParameter/ListDigitalAssistantParameters"
		err = common.PostProcessServiceError(err, "Management", "ListDigitalAssistantParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListDigitalAssistants Returns a page of Digital Assistants that belong to the specified Digital Assistant instance.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client ManagementClient) ListDigitalAssistants(ctx context.Context, request ListDigitalAssistantsRequest) (response ListDigitalAssistantsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listDigitalAssistants, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListDigitalAssistantsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListDigitalAssistantsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListDigitalAssistantsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListDigitalAssistantsResponse")
	}
	return
}

// listDigitalAssistants implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) listDigitalAssistants(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/digitalAssistants", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListDigitalAssistantsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistant/ListDigitalAssistants"
		err = common.PostProcessServiceError(err, "Management", "ListDigitalAssistants", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSkillParameters Returns a page of Skill Parameters that belong to the specified Skill.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client ManagementClient) ListSkillParameters(ctx context.Context, request ListSkillParametersRequest) (response ListSkillParametersResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSkillParameters, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSkillParametersResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSkillParametersResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSkillParametersResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSkillParametersResponse")
	}
	return
}

// listSkillParameters implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) listSkillParameters(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/skills/{skillId}/parameters", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSkillParametersResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/SkillParameter/ListSkillParameters"
		err = common.PostProcessServiceError(err, "Management", "ListSkillParameters", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListSkills Returns a page of Skills that belong to the specified Digital Assistant instance.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client ManagementClient) ListSkills(ctx context.Context, request ListSkillsRequest) (response ListSkillsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listSkills, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListSkillsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListSkillsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListSkillsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListSkillsResponse")
	}
	return
}

// listSkills implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) listSkills(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/skills", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListSkillsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Skill/ListSkills"
		err = common.PostProcessServiceError(err, "Management", "ListSkills", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListTranslators Returns a page of Translators that belong to the specified Digital Assistant instance.
// If the `opc-next-page` header appears in the response, then
// there are more items to retrieve. To get the next page in the subsequent
// GET request, include the header's value as the `page` query parameter.
func (client ManagementClient) ListTranslators(ctx context.Context, request ListTranslatorsRequest) (response ListTranslatorsResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listTranslators, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListTranslatorsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListTranslatorsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListTranslatorsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListTranslatorsResponse")
	}
	return
}

// listTranslators implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) listTranslators(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/odaInstances/{odaInstanceId}/translators", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListTranslatorsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Translator/ListTranslators"
		err = common.PostProcessServiceError(err, "Management", "ListTranslators", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishDigitalAssistant Publish a draft Digital Assistant.
// Once published the Digital Assistant cannot be modified.
func (client ManagementClient) PublishDigitalAssistant(ctx context.Context, request PublishDigitalAssistantRequest) (response PublishDigitalAssistantResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.publishDigitalAssistant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PublishDigitalAssistantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PublishDigitalAssistantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishDigitalAssistantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishDigitalAssistantResponse")
	}
	return
}

// publishDigitalAssistant implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) publishDigitalAssistant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/digitalAssistants/{digitalAssistantId}/actions/publish", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PublishDigitalAssistantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistant/PublishDigitalAssistant"
		err = common.PostProcessServiceError(err, "Management", "PublishDigitalAssistant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// PublishSkill Publish a draft Skill.
// Once published it cannot be modified.
func (client ManagementClient) PublishSkill(ctx context.Context, request PublishSkillRequest) (response PublishSkillResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.publishSkill, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = PublishSkillResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = PublishSkillResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(PublishSkillResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into PublishSkillResponse")
	}
	return
}

// publishSkill implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) publishSkill(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/skills/{skillId}/actions/publish", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response PublishSkillResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Skill/PublishSkill"
		err = common.PostProcessServiceError(err, "Management", "PublishSkill", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RotateChannelKeys This will generate new keys for any generated keys in the Channel (eg. secretKey, verifyToken).
// If a Channel has no generated keys then no changes will be made.
// Ensure that you take note of the newly generated keys in the response as they will not be returned again.
func (client ManagementClient) RotateChannelKeys(ctx context.Context, request RotateChannelKeysRequest) (response RotateChannelKeysResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.rotateChannelKeys, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RotateChannelKeysResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RotateChannelKeysResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RotateChannelKeysResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RotateChannelKeysResponse")
	}
	return
}

// rotateChannelKeys implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) rotateChannelKeys(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/channels/{channelId}/actions/rotateKeys", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RotateChannelKeysResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Channel/RotateChannelKeys"
		err = common.PostProcessServiceError(err, "Management", "RotateChannelKeys", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &createchannelresult{})
	return response, err
}

// StartChannel Starts a Channel so that it will begin accepting messages.
func (client ManagementClient) StartChannel(ctx context.Context, request StartChannelRequest) (response StartChannelResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.startChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StartChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StartChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StartChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StartChannelResponse")
	}
	return
}

// startChannel implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) startChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/channels/{channelId}/actions/start", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StartChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Channel/StartChannel"
		err = common.PostProcessServiceError(err, "Management", "StartChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &channel{})
	return response, err
}

// StopChannel Stops a Channel so that it will no longer accept messages.
func (client ManagementClient) StopChannel(ctx context.Context, request StopChannelRequest) (response StopChannelResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}

	if !(request.OpcRetryToken != nil && *request.OpcRetryToken != "") {
		request.OpcRetryToken = common.String(common.RetryToken())
	}

	ociResponse, err = common.Retry(ctx, request, client.stopChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = StopChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = StopChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(StopChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into StopChannelResponse")
	}
	return
}

// stopChannel implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) stopChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/odaInstances/{odaInstanceId}/channels/{channelId}/actions/stop", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response StopChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Channel/StopChannel"
		err = common.PostProcessServiceError(err, "Management", "StopChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &channel{})
	return response, err
}

// UpdateAuthenticationProvider Updates the specified Authentication Provider with the information in the request body.
func (client ManagementClient) UpdateAuthenticationProvider(ctx context.Context, request UpdateAuthenticationProviderRequest) (response UpdateAuthenticationProviderResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateAuthenticationProvider, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAuthenticationProviderResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAuthenticationProviderResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAuthenticationProviderResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAuthenticationProviderResponse")
	}
	return
}

// updateAuthenticationProvider implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) updateAuthenticationProvider(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/odaInstances/{odaInstanceId}/authenticationProviders/{authenticationProviderId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAuthenticationProviderResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/AuthenticationProvider/UpdateAuthenticationProvider"
		err = common.PostProcessServiceError(err, "Management", "UpdateAuthenticationProvider", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateChannel Updates the specified Channel with the information in the request body.
func (client ManagementClient) UpdateChannel(ctx context.Context, request UpdateChannelRequest) (response UpdateChannelResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateChannel, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateChannelResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateChannelResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateChannelResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateChannelResponse")
	}
	return
}

// updateChannel implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) updateChannel(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/odaInstances/{odaInstanceId}/channels/{channelId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateChannelResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Channel/UpdateChannel"
		err = common.PostProcessServiceError(err, "Management", "UpdateChannel", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &channel{})
	return response, err
}

// UpdateDigitalAssistant Updates the specified Digital Assistant with the information in the request body.
func (client ManagementClient) UpdateDigitalAssistant(ctx context.Context, request UpdateDigitalAssistantRequest) (response UpdateDigitalAssistantResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDigitalAssistant, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDigitalAssistantResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDigitalAssistantResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDigitalAssistantResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDigitalAssistantResponse")
	}
	return
}

// updateDigitalAssistant implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) updateDigitalAssistant(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/odaInstances/{odaInstanceId}/digitalAssistants/{digitalAssistantId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDigitalAssistantResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistant/UpdateDigitalAssistant"
		err = common.PostProcessServiceError(err, "Management", "UpdateDigitalAssistant", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateDigitalAssistantParameter Updates the specified Digital Assistant Parameter with the information in the request body.
func (client ManagementClient) UpdateDigitalAssistantParameter(ctx context.Context, request UpdateDigitalAssistantParameterRequest) (response UpdateDigitalAssistantParameterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateDigitalAssistantParameter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateDigitalAssistantParameterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateDigitalAssistantParameterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateDigitalAssistantParameterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateDigitalAssistantParameterResponse")
	}
	return
}

// updateDigitalAssistantParameter implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) updateDigitalAssistantParameter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/odaInstances/{odaInstanceId}/digitalAssistants/{digitalAssistantId}/parameters/{parameterName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateDigitalAssistantParameterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/DigitalAssistantParameter/UpdateDigitalAssistantParameter"
		err = common.PostProcessServiceError(err, "Management", "UpdateDigitalAssistantParameter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSkill Updates the specified Skill with the information in the request body.
func (client ManagementClient) UpdateSkill(ctx context.Context, request UpdateSkillRequest) (response UpdateSkillResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSkill, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSkillResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSkillResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSkillResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSkillResponse")
	}
	return
}

// updateSkill implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) updateSkill(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/odaInstances/{odaInstanceId}/skills/{skillId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSkillResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Skill/UpdateSkill"
		err = common.PostProcessServiceError(err, "Management", "UpdateSkill", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateSkillParameter Updates the specified Skill Parameter with the information in the request body.
func (client ManagementClient) UpdateSkillParameter(ctx context.Context, request UpdateSkillParameterRequest) (response UpdateSkillParameterResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateSkillParameter, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateSkillParameterResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateSkillParameterResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateSkillParameterResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateSkillParameterResponse")
	}
	return
}

// updateSkillParameter implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) updateSkillParameter(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/odaInstances/{odaInstanceId}/skills/{skillId}/parameters/{parameterName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateSkillParameterResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/SkillParameter/UpdateSkillParameter"
		err = common.PostProcessServiceError(err, "Management", "UpdateSkillParameter", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateTranslator Updates the specified Translator with the information in the request body.
func (client ManagementClient) UpdateTranslator(ctx context.Context, request UpdateTranslatorRequest) (response UpdateTranslatorResponse, err error) {
	var ociResponse common.OCIResponse
	var policy common.OCIRetry
	policy = common.NoRetryPolicyV2()
	if client.RetryPolicyV2() != nil {
		policy = client.RetryPolicyV2()
	}
	if request.RetryPolicy() != nil {
		policy = request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updateTranslator, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateTranslatorResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateTranslatorResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateTranslatorResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateTranslatorResponse")
	}
	return
}

// updateTranslator implements the OCIOperation interface (enables retrying operations)
func (client ManagementClient) updateTranslator(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/odaInstances/{odaInstanceId}/translators/{translatorId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateTranslatorResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/digital-assistant/20190506/Translator/UpdateTranslator"
		err = common.PostProcessServiceError(err, "Management", "UpdateTranslator", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
