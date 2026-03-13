// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_distributed_database "github.com/oracle/oci-go-sdk/v65/distributeddatabase"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_distributed_database.DistributedAutonomousDbServiceClient", &OracleClient{InitClientFn: initDistributeddatabaseDistributedAutonomousDbServiceClient})
	RegisterOracleClient("oci_distributed_database.DistributedDbWorkRequestServiceClient", &OracleClient{InitClientFn: initDistributeddatabaseDistributedDbWorkRequestServiceClient})
	RegisterOracleClient("oci_distributed_database.DistributedDbPrivateEndpointServiceClient", &OracleClient{InitClientFn: initDistributeddatabaseDistributedDbPrivateEndpointServiceClient})
	RegisterOracleClient("oci_distributed_database.DistributedDbServiceClient", &OracleClient{InitClientFn: initDistributeddatabaseDistributedDbServiceClient})
}

func initDistributeddatabaseDistributedAutonomousDbServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_distributed_database.NewDistributedAutonomousDbServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DistributedAutonomousDbServiceClient() *oci_distributed_database.DistributedAutonomousDbServiceClient {
	return m.GetClient("oci_distributed_database.DistributedAutonomousDbServiceClient").(*oci_distributed_database.DistributedAutonomousDbServiceClient)
}

func initDistributeddatabaseDistributedDbWorkRequestServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_distributed_database.NewDistributedDbWorkRequestServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DistributedDatabaseDistributedDbWorkRequestServiceClient() *oci_distributed_database.DistributedDbWorkRequestServiceClient {
	return m.GetClient("oci_distributed_database.DistributedDbWorkRequestServiceClient").(*oci_distributed_database.DistributedDbWorkRequestServiceClient)
}

func initDistributeddatabaseDistributedDbPrivateEndpointServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_distributed_database.NewDistributedDbPrivateEndpointServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DistributedDbPrivateEndpointServiceClient() *oci_distributed_database.DistributedDbPrivateEndpointServiceClient {
	return m.GetClient("oci_distributed_database.DistributedDbPrivateEndpointServiceClient").(*oci_distributed_database.DistributedDbPrivateEndpointServiceClient)
}

func initDistributeddatabaseDistributedDbServiceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_distributed_database.NewDistributedDbServiceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DistributedDbServiceClient() *oci_distributed_database.DistributedDbServiceClient {
	return m.GetClient("oci_distributed_database.DistributedDbServiceClient").(*oci_distributed_database.DistributedDbServiceClient)
}
