// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_dataflow "github.com/oracle/oci-go-sdk/v29/dataflow"

	oci_common "github.com/oracle/oci-go-sdk/v29/common"
)

func init() {
	RegisterOracleClient("oci_dataflow.DataFlowClient", &OracleClient{initClientFn: initDataflowDataFlowClient})
}

func initDataflowDataFlowClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_dataflow.NewDataFlowClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) dataFlowClient() *oci_dataflow.DataFlowClient {
	return m.GetClient("oci_dataflow.DataFlowClient").(*oci_dataflow.DataFlowClient)
}
