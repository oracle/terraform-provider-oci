// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Agent API
//
// API for the Oracle Cloud Agent software running on compute instances. Oracle Cloud Agent
// is a lightweight process that monitors and manages compute instances.
//

package computeinstanceagent

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v56/common"
	"github.com/oracle/oci-go-sdk/v56/common/auth"
	"net/http"
)

//PluginClient a client for Plugin
type PluginClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPluginClientWithConfigurationProvider Creates a new default Plugin client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPluginClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PluginClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newPluginClientFromBaseClient(baseClient, provider)
}

// NewPluginClientWithOboToken Creates a new default Plugin client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewPluginClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client PluginClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newPluginClientFromBaseClient(baseClient, configProvider)
}

func newPluginClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client PluginClient, err error) {
	// Plugin service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = PluginClient{BaseClient: baseClient}
	client.BasePath = "20180530"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PluginClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("computeinstanceagent", "https://iaas.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PluginClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
	if ok, err := common.IsConfigurationProviderValid(configProvider); !ok {
		return err
	}

	// Error has been checked already
	region, _ := configProvider.Region()
	client.SetRegion(region)
	client.config = &configProvider
	return nil
}

// ConfigurationProvider the ConfigurationProvider used in this client, or null if none set
func (client *PluginClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetInstanceAgentPlugin The API to get information for a plugin.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/GetInstanceAgentPlugin.go.html to see an example of how to use GetInstanceAgentPlugin API.
func (client PluginClient) GetInstanceAgentPlugin(ctx context.Context, request GetInstanceAgentPluginRequest) (response GetInstanceAgentPluginResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getInstanceAgentPlugin, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetInstanceAgentPluginResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetInstanceAgentPluginResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetInstanceAgentPluginResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetInstanceAgentPluginResponse")
	}
	return
}

// getInstanceAgentPlugin implements the OCIOperation interface (enables retrying operations)
func (client PluginClient) getInstanceAgentPlugin(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceagents/{instanceagentId}/plugins/{pluginName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetInstanceAgentPluginResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}

// ListInstanceAgentPlugins The API to get one or more plugin information.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/ListInstanceAgentPlugins.go.html to see an example of how to use ListInstanceAgentPlugins API.
func (client PluginClient) ListInstanceAgentPlugins(ctx context.Context, request ListInstanceAgentPluginsRequest) (response ListInstanceAgentPluginsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstanceAgentPlugins, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInstanceAgentPluginsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInstanceAgentPluginsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstanceAgentPluginsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstanceAgentPluginsResponse")
	}
	return
}

// listInstanceAgentPlugins implements the OCIOperation interface (enables retrying operations)
func (client PluginClient) listInstanceAgentPlugins(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceagents/{instanceagentId}/plugins", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInstanceAgentPluginsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
