// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Bridge API
//
// API for Oracle Cloud Bridge service.
//

package cloudbridge

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// OcbAgentSvcClient a client for OcbAgentSvc
type OcbAgentSvcClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewOcbAgentSvcClientWithConfigurationProvider Creates a new default OcbAgentSvc client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewOcbAgentSvcClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client OcbAgentSvcClient, err error) {
	if enabled := common.CheckForEnabledServices("cloudbridge"); !enabled {
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
	return newOcbAgentSvcClientFromBaseClient(baseClient, provider)
}

// NewOcbAgentSvcClientWithOboToken Creates a new default OcbAgentSvc client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewOcbAgentSvcClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client OcbAgentSvcClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newOcbAgentSvcClientFromBaseClient(baseClient, configProvider)
}

func newOcbAgentSvcClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client OcbAgentSvcClient, err error) {
	// OcbAgentSvc service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("OcbAgentSvc"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = OcbAgentSvcClient{BaseClient: baseClient}
	client.BasePath = "20220509"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *OcbAgentSvcClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("cloudbridge", "https://cloudbridge.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *OcbAgentSvcClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *OcbAgentSvcClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// AddAgentDependency Add a dependency to the environment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/AddAgentDependency.go.html to see an example of how to use AddAgentDependency API.
// A default retry strategy applies to this operation AddAgentDependency()
func (client OcbAgentSvcClient) AddAgentDependency(ctx context.Context, request AddAgentDependencyRequest) (response AddAgentDependencyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.addAgentDependency, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = AddAgentDependencyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = AddAgentDependencyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(AddAgentDependencyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into AddAgentDependencyResponse")
	}
	return
}

// addAgentDependency implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) addAgentDependency(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/environments/{environmentId}/actions/addAgentDependency", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response AddAgentDependencyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Environment/AddAgentDependency"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "AddAgentDependency", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAgentCompartment Moves an Agent resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ChangeAgentCompartment.go.html to see an example of how to use ChangeAgentCompartment API.
// A default retry strategy applies to this operation ChangeAgentCompartment()
func (client OcbAgentSvcClient) ChangeAgentCompartment(ctx context.Context, request ChangeAgentCompartmentRequest) (response ChangeAgentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAgentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAgentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAgentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAgentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAgentCompartmentResponse")
	}
	return
}

// changeAgentCompartment implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) changeAgentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/agents/{agentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAgentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Agent/ChangeAgentCompartment"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "ChangeAgentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeAgentDependencyCompartment Moves a AgentDependency resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ChangeAgentDependencyCompartment.go.html to see an example of how to use ChangeAgentDependencyCompartment API.
// A default retry strategy applies to this operation ChangeAgentDependencyCompartment()
func (client OcbAgentSvcClient) ChangeAgentDependencyCompartment(ctx context.Context, request ChangeAgentDependencyCompartmentRequest) (response ChangeAgentDependencyCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeAgentDependencyCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeAgentDependencyCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeAgentDependencyCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeAgentDependencyCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeAgentDependencyCompartmentResponse")
	}
	return
}

// changeAgentDependencyCompartment implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) changeAgentDependencyCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/agentDependencies/{agentDependencyId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeAgentDependencyCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AgentDependency/ChangeAgentDependencyCompartment"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "ChangeAgentDependencyCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ChangeEnvironmentCompartment Moves a source environment resource from one compartment identifier to another. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ChangeEnvironmentCompartment.go.html to see an example of how to use ChangeEnvironmentCompartment API.
// A default retry strategy applies to this operation ChangeEnvironmentCompartment()
func (client OcbAgentSvcClient) ChangeEnvironmentCompartment(ctx context.Context, request ChangeEnvironmentCompartmentRequest) (response ChangeEnvironmentCompartmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.changeEnvironmentCompartment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ChangeEnvironmentCompartmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ChangeEnvironmentCompartmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ChangeEnvironmentCompartmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ChangeEnvironmentCompartmentResponse")
	}
	return
}

// changeEnvironmentCompartment implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) changeEnvironmentCompartment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/environments/{environmentId}/actions/changeCompartment", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ChangeEnvironmentCompartmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Environment/ChangeEnvironmentCompartment"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "ChangeEnvironmentCompartment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAgent Creates an Agent.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/CreateAgent.go.html to see an example of how to use CreateAgent API.
// A default retry strategy applies to this operation CreateAgent()
func (client OcbAgentSvcClient) CreateAgent(ctx context.Context, request CreateAgentRequest) (response CreateAgentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAgent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAgentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAgentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAgentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAgentResponse")
	}
	return
}

// createAgent implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) createAgent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/agents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAgentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Agent/CreateAgent"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "CreateAgent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateAgentDependency Creates an AgentDependency.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/CreateAgentDependency.go.html to see an example of how to use CreateAgentDependency API.
// A default retry strategy applies to this operation CreateAgentDependency()
func (client OcbAgentSvcClient) CreateAgentDependency(ctx context.Context, request CreateAgentDependencyRequest) (response CreateAgentDependencyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createAgentDependency, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateAgentDependencyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateAgentDependencyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateAgentDependencyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateAgentDependencyResponse")
	}
	return
}

// createAgentDependency implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) createAgentDependency(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/agentDependencies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateAgentDependencyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AgentDependency/CreateAgentDependency"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "CreateAgentDependency", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// CreateEnvironment Creates a source environment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/CreateEnvironment.go.html to see an example of how to use CreateEnvironment API.
// A default retry strategy applies to this operation CreateEnvironment()
func (client OcbAgentSvcClient) CreateEnvironment(ctx context.Context, request CreateEnvironmentRequest) (response CreateEnvironmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.createEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = CreateEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = CreateEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(CreateEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into CreateEnvironmentResponse")
	}
	return
}

// createEnvironment implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) createEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/environments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response CreateEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Environment/CreateEnvironment"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "CreateEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAgent Deletes an Agent resource identified by an identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/DeleteAgent.go.html to see an example of how to use DeleteAgent API.
// A default retry strategy applies to this operation DeleteAgent()
func (client OcbAgentSvcClient) DeleteAgent(ctx context.Context, request DeleteAgentRequest) (response DeleteAgentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.deleteAgent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAgentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAgentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAgentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAgentResponse")
	}
	return
}

// deleteAgent implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) deleteAgent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/agents/{agentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAgentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Agent/DeleteAgent"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "DeleteAgent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteAgentDependency Deletes the AgentDependency resource based on an identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/DeleteAgentDependency.go.html to see an example of how to use DeleteAgentDependency API.
// A default retry strategy applies to this operation DeleteAgentDependency()
func (client OcbAgentSvcClient) DeleteAgentDependency(ctx context.Context, request DeleteAgentDependencyRequest) (response DeleteAgentDependencyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteAgentDependency, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteAgentDependencyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteAgentDependencyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteAgentDependencyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteAgentDependencyResponse")
	}
	return
}

// deleteAgentDependency implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) deleteAgentDependency(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/agentDependencies/{agentDependencyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteAgentDependencyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AgentDependency/DeleteAgentDependency"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "DeleteAgentDependency", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// DeleteEnvironment Deletes a the source environment resource identified by an identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/DeleteEnvironment.go.html to see an example of how to use DeleteEnvironment API.
// A default retry strategy applies to this operation DeleteEnvironment()
func (client OcbAgentSvcClient) DeleteEnvironment(ctx context.Context, request DeleteEnvironmentRequest) (response DeleteEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.deleteEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = DeleteEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = DeleteEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(DeleteEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into DeleteEnvironmentResponse")
	}
	return
}

// deleteEnvironment implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) deleteEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodDelete, "/environments/{environmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response DeleteEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Environment/DeleteEnvironment"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "DeleteEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAgent Gets an Agent by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/GetAgent.go.html to see an example of how to use GetAgent API.
// A default retry strategy applies to this operation GetAgent()
func (client OcbAgentSvcClient) GetAgent(ctx context.Context, request GetAgentRequest) (response GetAgentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAgent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAgentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAgentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAgentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAgentResponse")
	}
	return
}

// getAgent implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) getAgent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/agents/{agentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAgentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Agent/GetAgent"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "GetAgent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetAgentDependency Gets an AgentDependency by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/GetAgentDependency.go.html to see an example of how to use GetAgentDependency API.
// A default retry strategy applies to this operation GetAgentDependency()
func (client OcbAgentSvcClient) GetAgentDependency(ctx context.Context, request GetAgentDependencyRequest) (response GetAgentDependencyResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getAgentDependency, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetAgentDependencyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetAgentDependencyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetAgentDependencyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetAgentDependencyResponse")
	}
	return
}

// getAgentDependency implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) getAgentDependency(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/agentDependencies/{agentDependencyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetAgentDependencyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AgentDependency/GetAgentDependency"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "GetAgentDependency", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetEnvironment Gets a source environment by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/GetEnvironment.go.html to see an example of how to use GetEnvironment API.
// A default retry strategy applies to this operation GetEnvironment()
func (client OcbAgentSvcClient) GetEnvironment(ctx context.Context, request GetEnvironmentRequest) (response GetEnvironmentResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetEnvironmentResponse")
	}
	return
}

// getEnvironment implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) getEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/environments/{environmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Environment/GetEnvironment"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "GetEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// GetPlugin Gets a plugin by identifier.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/GetPlugin.go.html to see an example of how to use GetPlugin API.
// A default retry strategy applies to this operation GetPlugin()
func (client OcbAgentSvcClient) GetPlugin(ctx context.Context, request GetPluginRequest) (response GetPluginResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getPlugin, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetPluginResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetPluginResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetPluginResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetPluginResponse")
	}
	return
}

// getPlugin implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) getPlugin(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/agents/{agentId}/plugins/{pluginName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetPluginResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Plugin/GetPlugin"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "GetPlugin", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAgentDependencies Returns a list of AgentDependencies such as AgentDependencyCollection.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListAgentDependencies.go.html to see an example of how to use ListAgentDependencies API.
// A default retry strategy applies to this operation ListAgentDependencies()
func (client OcbAgentSvcClient) ListAgentDependencies(ctx context.Context, request ListAgentDependenciesRequest) (response ListAgentDependenciesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAgentDependencies, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAgentDependenciesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAgentDependenciesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAgentDependenciesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAgentDependenciesResponse")
	}
	return
}

// listAgentDependencies implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) listAgentDependencies(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/agentDependencies", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAgentDependenciesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AgentDependencyCollection/ListAgentDependencies"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "ListAgentDependencies", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListAgents Returns a list of Agents.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListAgents.go.html to see an example of how to use ListAgents API.
// A default retry strategy applies to this operation ListAgents()
func (client OcbAgentSvcClient) ListAgents(ctx context.Context, request ListAgentsRequest) (response ListAgentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listAgents, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListAgentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListAgentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListAgentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListAgentsResponse")
	}
	return
}

// listAgents implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) listAgents(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/agents", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListAgentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AgentCollection/ListAgents"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "ListAgents", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListApplianceImages Returns a list of Appliance Images.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListApplianceImages.go.html to see an example of how to use ListApplianceImages API.
// A default retry strategy applies to this operation ListApplianceImages()
func (client OcbAgentSvcClient) ListApplianceImages(ctx context.Context, request ListApplianceImagesRequest) (response ListApplianceImagesResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listApplianceImages, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListApplianceImagesResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListApplianceImagesResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListApplianceImagesResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListApplianceImagesResponse")
	}
	return
}

// listApplianceImages implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) listApplianceImages(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/applianceImages", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListApplianceImagesResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/ApplianceImageCollection/ListApplianceImages"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "ListApplianceImages", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListEnvironments Returns a list of source environments.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/ListEnvironments.go.html to see an example of how to use ListEnvironments API.
// A default retry strategy applies to this operation ListEnvironments()
func (client OcbAgentSvcClient) ListEnvironments(ctx context.Context, request ListEnvironmentsRequest) (response ListEnvironmentsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listEnvironments, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListEnvironmentsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListEnvironmentsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListEnvironmentsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListEnvironmentsResponse")
	}
	return
}

// listEnvironments implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) listEnvironments(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/environments", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListEnvironmentsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/EnvironmentCollection/ListEnvironments"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "ListEnvironments", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// RemoveAgentDependency Adds a dependency to the source environment. When provided, If-Match is checked against ETag values of the resource.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/RemoveAgentDependency.go.html to see an example of how to use RemoveAgentDependency API.
// A default retry strategy applies to this operation RemoveAgentDependency()
func (client OcbAgentSvcClient) RemoveAgentDependency(ctx context.Context, request RemoveAgentDependencyRequest) (response RemoveAgentDependencyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.removeAgentDependency, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = RemoveAgentDependencyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = RemoveAgentDependencyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(RemoveAgentDependencyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into RemoveAgentDependencyResponse")
	}
	return
}

// removeAgentDependency implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) removeAgentDependency(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPost, "/environments/{environmentId}/actions/removeAgentDependency", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response RemoveAgentDependencyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Environment/RemoveAgentDependency"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "RemoveAgentDependency", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAgent Updates the Agent.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/UpdateAgent.go.html to see an example of how to use UpdateAgent API.
// A default retry strategy applies to this operation UpdateAgent()
func (client OcbAgentSvcClient) UpdateAgent(ctx context.Context, request UpdateAgentRequest) (response UpdateAgentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateAgent, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAgentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAgentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAgentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAgentResponse")
	}
	return
}

// updateAgent implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) updateAgent(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/agents/{agentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAgentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Agent/UpdateAgent"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "UpdateAgent", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateAgentDependency Updates the AgentDependency.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/UpdateAgentDependency.go.html to see an example of how to use UpdateAgentDependency API.
// A default retry strategy applies to this operation UpdateAgentDependency()
func (client OcbAgentSvcClient) UpdateAgentDependency(ctx context.Context, request UpdateAgentDependencyRequest) (response UpdateAgentDependencyResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateAgentDependency, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateAgentDependencyResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateAgentDependencyResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateAgentDependencyResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateAgentDependencyResponse")
	}
	return
}

// updateAgentDependency implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) updateAgentDependency(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/agentDependencies/{agentDependencyId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateAgentDependencyResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/AgentDependency/UpdateAgentDependency"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "UpdateAgentDependency", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdateEnvironment Updates the source environment.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/UpdateEnvironment.go.html to see an example of how to use UpdateEnvironment API.
// A default retry strategy applies to this operation UpdateEnvironment()
func (client OcbAgentSvcClient) UpdateEnvironment(ctx context.Context, request UpdateEnvironmentRequest) (response UpdateEnvironmentResponse, err error) {
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

	ociResponse, err = common.Retry(ctx, request, client.updateEnvironment, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdateEnvironmentResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdateEnvironmentResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdateEnvironmentResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdateEnvironmentResponse")
	}
	return
}

// updateEnvironment implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) updateEnvironment(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/environments/{environmentId}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdateEnvironmentResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Environment/UpdateEnvironment"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "UpdateEnvironment", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// UpdatePlugin Updates the plugin.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/cloudbridge/UpdatePlugin.go.html to see an example of how to use UpdatePlugin API.
// A default retry strategy applies to this operation UpdatePlugin()
func (client OcbAgentSvcClient) UpdatePlugin(ctx context.Context, request UpdatePluginRequest) (response UpdatePluginResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.updatePlugin, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = UpdatePluginResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = UpdatePluginResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(UpdatePluginResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into UpdatePluginResponse")
	}
	return
}

// updatePlugin implements the OCIOperation interface (enables retrying operations)
func (client OcbAgentSvcClient) updatePlugin(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodPut, "/agents/{agentId}/plugins/{pluginName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response UpdatePluginResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := "https://docs.oracle.com/iaas/api/#/en/OCB/20220509/Plugin/UpdatePlugin"
		err = common.PostProcessServiceError(err, "OcbAgentSvc", "UpdatePlugin", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
