// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_cloud_migrations "github.com/oracle/oci-go-sdk/v65/cloudmigrations"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_cloud_migrations.MigrationClient", &OracleClient{InitClientFn: initCloudmigrationsMigrationClient})
}

func initCloudmigrationsMigrationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_cloud_migrations.NewMigrationClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) MigrationClient() *oci_cloud_migrations.MigrationClient {
	return m.GetClient("oci_cloud_migrations.MigrationClient").(*oci_cloud_migrations.MigrationClient)
}
