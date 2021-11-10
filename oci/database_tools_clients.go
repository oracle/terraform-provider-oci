// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_database_tools "github.com/oracle/oci-go-sdk/v51/databasetools"

	oci_common "github.com/oracle/oci-go-sdk/v51/common"
)

func init() {
	RegisterOracleClient("oci_database_tools.DatabaseToolsClient", &OracleClient{InitClientFn: initDatabasetoolsDatabaseToolsClient})
}

func initDatabasetoolsDatabaseToolsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_database_tools.NewDatabaseToolsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) databaseToolsClient() *oci_database_tools.DatabaseToolsClient {
	return m.GetClient("oci_database_tools.DatabaseToolsClient").(*oci_database_tools.DatabaseToolsClient)
}
