// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_sch "github.com/oracle/oci-go-sdk/sch"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_sch.ServiceConnectorClient", &OracleClient{initClientFn: initSchServiceConnectorClient})
}

func initSchServiceConnectorClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_sch.NewServiceConnectorClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) serviceConnectorClient() *oci_sch.ServiceConnectorClient {
	return m.GetClient("oci_sch.ServiceConnectorClient").(*oci_sch.ServiceConnectorClient)
}
