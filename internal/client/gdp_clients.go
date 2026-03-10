// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_gdp "github.com/oracle/oci-go-sdk/v65/gdp"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_gdp.GuardedDataPipelineClient", &OracleClient{InitClientFn: initGdpGuardedDataPipelineClient})
}

func initGdpGuardedDataPipelineClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_gdp.NewGuardedDataPipelineClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) GuardedDataPipelineClient() *oci_gdp.GuardedDataPipelineClient {
	return m.GetClient("oci_gdp.GuardedDataPipelineClient").(*oci_gdp.GuardedDataPipelineClient)
}
