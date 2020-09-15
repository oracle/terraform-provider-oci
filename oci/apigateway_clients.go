// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_apigateway "github.com/oracle/oci-go-sdk/v25/apigateway"

	oci_common "github.com/oracle/oci-go-sdk/v25/common"
)

func init() {
	RegisterOracleClient("oci_apigateway.DeploymentClient", &OracleClient{initClientFn: initApigatewayDeploymentClient})
	RegisterOracleClient("oci_apigateway.GatewayClient", &OracleClient{initClientFn: initApigatewayGatewayClient})
}

func initApigatewayDeploymentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_apigateway.NewDeploymentClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) deploymentClient() *oci_apigateway.DeploymentClient {
	return m.GetClient("oci_apigateway.DeploymentClient").(*oci_apigateway.DeploymentClient)
}

func initApigatewayGatewayClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_apigateway.NewGatewayClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) gatewayClient() *oci_apigateway.GatewayClient {
	return m.GetClient("oci_apigateway.GatewayClient").(*oci_apigateway.GatewayClient)
}
