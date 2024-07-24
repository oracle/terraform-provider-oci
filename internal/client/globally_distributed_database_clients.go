// Copyright (c) 2017, 2023, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_globally_distributed_database "github.com/oracle/oci-go-sdk/v65/globallydistributeddatabase"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_globally_distributed_database.ShardedDatabaseServiceClient", &OracleClient{InitClientFn: initGloballydistributeddatabaseShardedDatabaseServiceClient})
}

func initGloballydistributeddatabaseShardedDatabaseServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_globally_distributed_database.NewShardedDatabaseServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ShardedDatabaseServiceClient() *oci_globally_distributed_database.ShardedDatabaseServiceClient {
	return m.GetClient("oci_globally_distributed_database.ShardedDatabaseServiceClient").(*oci_globally_distributed_database.ShardedDatabaseServiceClient)
}
