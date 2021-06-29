// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_database_migration "github.com/oracle/oci-go-sdk/v43/databasemigration"

	oci_common "github.com/oracle/oci-go-sdk/v43/common"
)

func init() {
	RegisterOracleClient("oci_database_migration.DatabaseMigrationClient", &OracleClient{initClientFn: initDatabasemigrationDatabaseMigrationClient})
}

func initDatabasemigrationDatabaseMigrationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_database_migration.NewDatabaseMigrationClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) databaseMigrationClient() *oci_database_migration.DatabaseMigrationClient {
	return m.GetClient("oci_database_migration.DatabaseMigrationClient").(*oci_database_migration.DatabaseMigrationClient)
}
