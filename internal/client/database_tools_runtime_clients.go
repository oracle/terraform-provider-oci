// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_database_tools_runtime "github.com/oracle/oci-go-sdk/v65/databasetoolsruntime"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_database_tools_runtime.DatabaseToolsRuntimeClient", &OracleClient{InitClientFn: initDatabasetoolsruntimeDatabaseToolsRuntimeClient})
}

func initDatabasetoolsruntimeDatabaseToolsRuntimeClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_database_tools_runtime.NewDatabaseToolsRuntimeClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DatabaseToolsRuntimeClient() *oci_database_tools_runtime.DatabaseToolsRuntimeClient {
	return m.GetClient("oci_database_tools_runtime.DatabaseToolsRuntimeClient").(*oci_database_tools_runtime.DatabaseToolsRuntimeClient)
}
