// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_integration "github.com/oracle/oci-go-sdk/v56/integration"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_integration.IntegrationInstanceClient", &OracleClient{InitClientFn: initIntegrationIntegrationInstanceClient})
}

func initIntegrationIntegrationInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_integration.NewIntegrationInstanceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) IntegrationInstanceClient() *oci_integration.IntegrationInstanceClient {
	return m.GetClient("oci_integration.IntegrationInstanceClient").(*oci_integration.IntegrationInstanceClient)
}
