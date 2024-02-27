// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Connector Hub API
//
// Use the Connector Hub API to transfer data between services in Oracle Cloud Infrastructure.
// For more information about Connector Hub, see
// the Connector Hub documentation (https://docs.cloud.oracle.com/iaas/Content/connector-hub/home.htm).
// Connector Hub is formerly known as Service Connector Hub.
//

package sch

import (
	"context"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/common/auth"
	"net/http"
)

// ConnectorPluginsClient a client for ConnectorPlugins
type ConnectorPluginsClient struct {
	common.BaseClient
	config *common.ConfigurationProvider
}

// NewConnectorPluginsClientWithConfigurationProvider Creates a new default ConnectorPlugins client with the given configuration provider.
// the configuration provider will be used for the default signer as well as reading the region
func NewConnectorPluginsClientWithConfigurationProvider(configProvider common.ConfigurationProvider) (client ConnectorPluginsClient, err error) {
	if enabled := common.CheckForEnabledServices("sch"); !enabled {
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
	return newConnectorPluginsClientFromBaseClient(baseClient, provider)
}

// NewConnectorPluginsClientWithOboToken Creates a new default ConnectorPlugins client with the given configuration provider.
// The obotoken will be added to default headers and signed; the configuration provider will be used for the signer
//
//	as well as reading the region
func NewConnectorPluginsClientWithOboToken(configProvider common.ConfigurationProvider, oboToken string) (client ConnectorPluginsClient, err error) {
	baseClient, err := common.NewClientWithOboToken(configProvider, oboToken)
	if err != nil {
		return client, err
	}

	return newConnectorPluginsClientFromBaseClient(baseClient, configProvider)
}

func newConnectorPluginsClientFromBaseClient(baseClient common.BaseClient, configProvider common.ConfigurationProvider) (client ConnectorPluginsClient, err error) {
	// ConnectorPlugins service default circuit breaker is enabled
	baseClient.Configuration.CircuitBreaker = common.NewCircuitBreaker(common.DefaultCircuitBreakerSettingWithServiceName("ConnectorPlugins"))
	common.ConfigCircuitBreakerFromEnvVar(&baseClient)
	common.ConfigCircuitBreakerFromGlobalVar(&baseClient)

	client = ConnectorPluginsClient{BaseClient: baseClient}
	client.BasePath = "20200909"
	err = client.setConfigurationProvider(configProvider)
	return
}

// SetRegion overrides the region of this client.
func (client *ConnectorPluginsClient) SetRegion(region string) {
	client.Host = common.StringToRegion(region).EndpointForTemplate("sch", "https://service-connector-hub.{region}.oci.{secondLevelDomain}")
}

// SetConfigurationProvider sets the configuration provider including the region, returns an error if is not valid
func (client *ConnectorPluginsClient) setConfigurationProvider(configProvider common.ConfigurationProvider) error {
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
func (client *ConnectorPluginsClient) ConfigurationProvider() *common.ConfigurationProvider {
	return client.config
}

// GetConnectorPlugin Gets the specified connector plugin configuration information.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/sch/GetConnectorPlugin.go.html to see an example of how to use GetConnectorPlugin API.
// A default retry strategy applies to this operation GetConnectorPlugin()
func (client ConnectorPluginsClient) GetConnectorPlugin(ctx context.Context, request GetConnectorPluginRequest) (response GetConnectorPluginResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.getConnectorPlugin, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = GetConnectorPluginResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = GetConnectorPluginResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(GetConnectorPluginResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into GetConnectorPluginResponse")
	}
	return
}

// getConnectorPlugin implements the OCIOperation interface (enables retrying operations)
func (client ConnectorPluginsClient) getConnectorPlugin(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connectorPlugins/{connectorPluginName}", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response GetConnectorPluginResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ConnectorPlugins", "GetConnectorPlugin", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponseWithPolymorphicBody(httpResponse, &response, &connectorplugin{})
	return response, err
}

// ListConnectorPlugins Lists connector plugins according to the specified filter.
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/sch/ListConnectorPlugins.go.html to see an example of how to use ListConnectorPlugins API.
// A default retry strategy applies to this operation ListConnectorPlugins()
func (client ConnectorPluginsClient) ListConnectorPlugins(ctx context.Context, request ListConnectorPluginsRequest) (response ListConnectorPluginsResponse, err error) {
	var ociResponse common.OCIResponse
	policy := common.DefaultRetryPolicy()
	if client.RetryPolicy() != nil {
		policy = *client.RetryPolicy()
	}
	if request.RetryPolicy() != nil {
		policy = *request.RetryPolicy()
	}
	ociResponse, err = common.Retry(ctx, request, client.listConnectorPlugins, policy)
	if err != nil {
		if ociResponse != nil {
			if httpResponse := ociResponse.HTTPResponse(); httpResponse != nil {
				opcRequestId := httpResponse.Header.Get("opc-request-id")
				response = ListConnectorPluginsResponse{RawResponse: httpResponse, OpcRequestId: &opcRequestId}
			} else {
				response = ListConnectorPluginsResponse{}
			}
		}
		return
	}
	if convertedResponse, ok := ociResponse.(ListConnectorPluginsResponse); ok {
		response = convertedResponse
	} else {
		err = fmt.Errorf("failed to convert OCIResponse into ListConnectorPluginsResponse")
	}
	return
}

// listConnectorPlugins implements the OCIOperation interface (enables retrying operations)
func (client ConnectorPluginsClient) listConnectorPlugins(ctx context.Context, request common.OCIRequest, binaryReqBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (common.OCIResponse, error) {

	httpRequest, err := request.HTTPRequest(http.MethodGet, "/connectorPlugins", binaryReqBody, extraHeaders)
	if err != nil {
		return nil, err
	}

	var response ListConnectorPluginsResponse
	var httpResponse *http.Response
	httpResponse, err = client.Call(ctx, &httpRequest)
	defer common.CloseBodyIfValid(httpResponse)
	response.RawResponse = httpResponse
	if err != nil {
		apiReferenceLink := ""
		err = common.PostProcessServiceError(err, "ConnectorPlugins", "ListConnectorPlugins", apiReferenceLink)
		return response, err
	}

	err = common.UnmarshalResponse(httpResponse, &response)
	return response, err
}
