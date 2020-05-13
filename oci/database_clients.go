// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_database "github.com/oracle/oci-go-sdk/database"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_database.DatabaseClient", &OracleClient{initClientFn: initDatabaseDatabaseClient})
}

func initDatabaseDatabaseClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_database.NewDatabaseClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) databaseClient() *oci_database.DatabaseClient {
	return m.GetClient("oci_database.DatabaseClient").(*oci_database.DatabaseClient)
}
