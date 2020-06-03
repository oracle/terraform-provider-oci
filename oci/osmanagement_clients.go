// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_osmanagement "github.com/oracle/oci-go-sdk/osmanagement"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_osmanagement.OsManagementClient", &OracleClient{initClientFn: initOsmanagementOsManagementClient})
}

func initOsmanagementOsManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_osmanagement.NewOsManagementClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) osManagementClient() *oci_osmanagement.OsManagementClient {
	return m.GetClient("oci_osmanagement.OsManagementClient").(*oci_osmanagement.OsManagementClient)
}
