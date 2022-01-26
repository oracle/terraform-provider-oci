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

//PluginconfigClient a client for Pluginconfig
type PluginconfigClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewPluginconfigClientWithConfigurationProvider Creates a new default Pluginconfig client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewPluginconfigClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client PluginconfigClient, err error) {
	provider, err := auth.GetGenericConfigurationProvider(configProvider)
	if err != nil {
		return client, err
	}
	baseClient, e := common.NewClientWithConfig(provider)
	if e != nil {
		return client, e
	}
	return newPluginconfigClientFromBaseClient(baseClient, provider)
}

// NewPluginconfigClientWithOboToken Creates a new default Pluginconfig client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//  as well as reading the region
func NewPluginconfigClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client PluginconfigClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newPluginconfigClientFromBaseClient(baseClient, configProvider)
}

func newPluginconfigClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client PluginconfigClient, err error) {
	// Pluginconfig service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSetting())
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = PluginconfigClient{BaseClient: baseClient}
	client.BasePath = "20180530"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *PluginconfigClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("computeinstanceagent", "https://iaas.{region}.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *PluginconfigClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *PluginconfigClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// ListInstanceagentAvailablePlugins The API to get the list of plugins that are available.
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/computeinstanceagent/ListInstanceagentAvailablePlugins.go.html to see an example of how to use ListInstanceagentAvailablePlugins API.
func (client PluginconfigClient) ListInstanceagentAvailablePlugins(ctx context.Context, request ListInstanceagentAvailablePluginsRequest) (response ListInstanceagentAvailablePluginsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.NoRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listInstanceagentAvailablePlugins, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListInstanceagentAvailablePluginsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListInstanceagentAvailablePluginsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListInstanceagentAvailablePluginsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListInstanceagentAvailablePluginsResponse")
	}
	return
}

// listInstanceagentAvailablePlugins implements the OCIOperation interface (enables retrying operations)
func (client PluginconfigClient) listInstanceagentAvailablePlugins(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/instanceagent/availablePlugins", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListInstanceagentAvailablePluginsResponse
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
