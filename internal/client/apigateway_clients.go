// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_apigateway "github.com/oracle/oci-go-sdk/v65/apigateway"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_apigateway.ApiGatewayClient", &OracleClient{InitClientFn: initApigatewayApiGatewayClient})
	RegisterOracleClient("oci_apigateway.WorkRequestsClient", &OracleClient{InitClientFn: initApigatewayWorkRequestsClient})
	RegisterOracleClient("oci_apigateway.DeploymentClient", &OracleClient{InitClientFn: initApigatewayDeploymentClient})
	RegisterOracleClient("oci_apigateway.GatewayClient", &OracleClient{InitClientFn: initApigatewayGatewayClient})
	RegisterOracleClient("oci_apigateway.SubscribersClient", &OracleClient{InitClientFn: initApigatewaySubscribersClient})
	RegisterOracleClient("oci_apigateway.UsagePlansClient", &OracleClient{InitClientFn: initApigatewayUsagePlansClient})
}

func initApigatewayApiGatewayClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewApiGatewayClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ApiGatewayClient() *oci_apigateway.ApiGatewayClient {
	return m.GetClient("oci_apigateway.ApiGatewayClient").(*oci_apigateway.ApiGatewayClient)
}

func initApigatewayWorkRequestsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewWorkRequestsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ApigatewayWorkRequestsClient() *oci_apigateway.WorkRequestsClient {
	return m.GetClient("oci_apigateway.WorkRequestsClient").(*oci_apigateway.WorkRequestsClient)
}

func initApigatewayDeploymentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewDeploymentClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) DeploymentClient() *oci_apigateway.DeploymentClient {
	return m.GetClient("oci_apigateway.DeploymentClient").(*oci_apigateway.DeploymentClient)
}

func initApigatewayGatewayClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewGatewayClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) GatewayClient() *oci_apigateway.GatewayClient {
	return m.GetClient("oci_apigateway.GatewayClient").(*oci_apigateway.GatewayClient)
}

func initApigatewaySubscribersClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewSubscribersClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) SubscribersClient() *oci_apigateway.SubscribersClient {
	return m.GetClient("oci_apigateway.SubscribersClient").(*oci_apigateway.SubscribersClient)
}

func initApigatewayUsagePlansClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apigateway.NewUsagePlansClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) UsagePlansClient() *oci_apigateway.UsagePlansClient {
	return m.GetClient("oci_apigateway.UsagePlansClient").(*oci_apigateway.UsagePlansClient)
}
