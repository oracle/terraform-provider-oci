// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_common "github.com/oracle/oci-go-sdk/v65/common"
	oci_opsi "github.com/oracle/oci-go-sdk/v65/opsi"
)

func init() {
	RegisterOracleClient("oci_opsi.OperationsInsightsClient", &OracleClient{InitClientFn: initOperationsinsightsOperationsInsightsClient})
}

func initOperationsinsightsOperationsInsightsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_opsi.NewOperationsInsightsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OperationsInsightsClient() *oci_opsi.OperationsInsightsClient {
	return m.GetClient("oci_opsi.OperationsInsightsClient").(*oci_opsi.OperationsInsightsClient)
}
