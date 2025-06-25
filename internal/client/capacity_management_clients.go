// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_capacity_management "github.com/oracle/oci-go-sdk/v65/capacitymanagement"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_capacity_management.CapacityManagementClient", &OracleClient{InitClientFn: initCapacitymanagementCapacityManagementClient})
	RegisterOracleClient("oci_capacity_management.DemandSignalClient", &OracleClient{InitClientFn: initCapacitymanagementDemandSignalClient})
	RegisterOracleClient("oci_capacity_management.InternalDemandSignalClient", &OracleClient{InitClientFn: initCapacitymanagementInternalDemandSignalClient})
}

func initCapacitymanagementCapacityManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_capacity_management.NewCapacityManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) CapacityManagementClient() *oci_capacity_management.CapacityManagementClient {
	return m.GetClient("oci_capacity_management.CapacityManagementClient").(*oci_capacity_management.CapacityManagementClient)
}

func initCapacitymanagementDemandSignalClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_capacity_management.NewDemandSignalClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DemandSignalClient() *oci_capacity_management.DemandSignalClient {
	return m.GetClient("oci_capacity_management.DemandSignalClient").(*oci_capacity_management.DemandSignalClient)
}

func initCapacitymanagementInternalDemandSignalClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_capacity_management.NewInternalDemandSignalClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) InternalDemandSignalClient() *oci_capacity_management.InternalDemandSignalClient {
	return m.GetClient("oci_capacity_management.InternalDemandSignalClient").(*oci_capacity_management.InternalDemandSignalClient)
}
