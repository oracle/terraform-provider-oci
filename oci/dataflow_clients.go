// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_dataflow "github.com/oracle/oci-go-sdk/dataflow"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_dataflow.DataFlowClient", &OracleClient{initClientFn: initDataflowDataFlowClient})
}

func initDataflowDataFlowClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_dataflow.NewDataFlowClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) dataFlowClient() *oci_dataflow.DataFlowClient {
	return m.GetClient("oci_dataflow.DataFlowClient").(*oci_dataflow.DataFlowClient)
}
