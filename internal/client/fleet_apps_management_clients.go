// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_fleet_apps_management "github.com/oracle/oci-go-sdk/v65/fleetappsmanagement"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_fleet_apps_management.FleetAppsManagementClient", &OracleClient{InitClientFn: initFleetappsmanagementFleetAppsManagementClient})
	RegisterOracleClient("oci_fleet_apps_management.FleetAppsManagementAdminClient", &OracleClient{InitClientFn: initFleetappsmanagementFleetAppsManagementAdminClient})
	RegisterOracleClient("oci_fleet_apps_management.FleetAppsManagementMaintenanceWindowClient", &OracleClient{InitClientFn: initFleetappsmanagementFleetAppsManagementMaintenanceWindowClient})
	RegisterOracleClient("oci_fleet_apps_management.FleetAppsManagementOperationsClient", &OracleClient{InitClientFn: initFleetappsmanagementFleetAppsManagementOperationsClient})
	RegisterOracleClient("oci_fleet_apps_management.FleetAppsManagementRunbooksClient", &OracleClient{InitClientFn: initFleetappsmanagementFleetAppsManagementRunbooksClient})
}

func initFleetappsmanagementFleetAppsManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_fleet_apps_management.NewFleetAppsManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) FleetAppsManagementClient() *oci_fleet_apps_management.FleetAppsManagementClient {
	return m.GetClient("oci_fleet_apps_management.FleetAppsManagementClient").(*oci_fleet_apps_management.FleetAppsManagementClient)
}

func initFleetappsmanagementFleetAppsManagementAdminClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_fleet_apps_management.NewFleetAppsManagementAdminClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) FleetAppsManagementAdminClient() *oci_fleet_apps_management.FleetAppsManagementAdminClient {
	return m.GetClient("oci_fleet_apps_management.FleetAppsManagementAdminClient").(*oci_fleet_apps_management.FleetAppsManagementAdminClient)
}

func initFleetappsmanagementFleetAppsManagementMaintenanceWindowClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_fleet_apps_management.NewFleetAppsManagementMaintenanceWindowClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) FleetAppsManagementMaintenanceWindowClient() *oci_fleet_apps_management.FleetAppsManagementMaintenanceWindowClient {
	return m.GetClient("oci_fleet_apps_management.FleetAppsManagementMaintenanceWindowClient").(*oci_fleet_apps_management.FleetAppsManagementMaintenanceWindowClient)
}

func initFleetappsmanagementFleetAppsManagementOperationsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_fleet_apps_management.NewFleetAppsManagementOperationsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) FleetAppsManagementOperationsClient() *oci_fleet_apps_management.FleetAppsManagementOperationsClient {
	return m.GetClient("oci_fleet_apps_management.FleetAppsManagementOperationsClient").(*oci_fleet_apps_management.FleetAppsManagementOperationsClient)
}

func initFleetappsmanagementFleetAppsManagementRunbooksClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_fleet_apps_management.NewFleetAppsManagementRunbooksClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) FleetAppsManagementRunbooksClient() *oci_fleet_apps_management.FleetAppsManagementRunbooksClient {
	return m.GetClient("oci_fleet_apps_management.FleetAppsManagementRunbooksClient").(*oci_fleet_apps_management.FleetAppsManagementRunbooksClient)
}
