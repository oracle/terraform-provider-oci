// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_database_management "github.com/oracle/oci-go-sdk/v43/databasemanagement"

	oci_common "github.com/oracle/oci-go-sdk/v43/common"
)

func init() {
	RegisterOracleClient("oci_database_management.DbManagementClient", &OracleClient{initClientFn: initDatabasemanagementDbManagementClient})
}

func initDatabasemanagementDbManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_database_management.NewDbManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) dbManagementClient() *oci_database_management.DbManagementClient {
	return m.GetClient("oci_database_management.DbManagementClient").(*oci_database_management.DbManagementClient)
}
