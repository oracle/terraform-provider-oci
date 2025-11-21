// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_batch "github.com/oracle/oci-go-sdk/v65/batch"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_batch.BatchComputingClient", &OracleClient{InitClientFn: initBatchBatchComputingClient})
}

func initBatchBatchComputingClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_batch.NewBatchComputingClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) BatchComputingClient() *oci_batch.BatchComputingClient {
	return m.GetClient("oci_batch.BatchComputingClient").(*oci_batch.BatchComputingClient)
}
