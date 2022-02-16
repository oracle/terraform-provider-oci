// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_dataintegration "github.com/oracle/oci-go-sdk/v58/dataintegration"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_dataintegration.DataIntegrationClient", &OracleClient{InitClientFn: initDataintegrationDataIntegrationClient})
}

func initDataintegrationDataIntegrationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dataintegration.NewDataIntegrationClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DataIntegrationClient() *oci_dataintegration.DataIntegrationClient {
	return m.GetClient("oci_dataintegration.DataIntegrationClient").(*oci_dataintegration.DataIntegrationClient)
}
