// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_opsi "github.com/oracle/oci-go-sdk/v44/opsi"

	oci_common "github.com/oracle/oci-go-sdk/v44/common"
)

func init() {
	RegisterOracleClient("oci_opsi.OperationsInsightsClient", &OracleClient{initClientFn: initOpsiOperationsInsightsClient})
}

func initOpsiOperationsInsightsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_opsi.NewOperationsInsightsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) operationsInsightsClient() *oci_opsi.OperationsInsightsClient {
	return m.GetClient("oci_opsi.OperationsInsightsClient").(*oci_opsi.OperationsInsightsClient)
}
