// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_integration "github.com/oracle/oci-go-sdk/v27/integration"

	oci_common "github.com/oracle/oci-go-sdk/v27/common"
)

func init() {
	RegisterOracleClient("oci_integration.IntegrationInstanceClient", &OracleClient{initClientFn: initIntegrationIntegrationInstanceClient})
}

func initIntegrationIntegrationInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_integration.NewIntegrationInstanceClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) integrationInstanceClient() *oci_integration.IntegrationInstanceClient {
	return m.GetClient("oci_integration.IntegrationInstanceClient").(*oci_integration.IntegrationInstanceClient)
}
