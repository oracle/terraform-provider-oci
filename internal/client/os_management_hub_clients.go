// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_os_management_hub.LifecycleEnvironmentClient", &OracleClient{InitClientFn: initOsmanagementhubLifecycleEnvironmentClient})
	RegisterOracleClient("oci_os_management_hub.ManagedInstanceGroupClient", &OracleClient{InitClientFn: initOsmanagementhubManagedInstanceGroupClient})
	RegisterOracleClient("oci_os_management_hub.ManagementStationClient", &OracleClient{InitClientFn: initOsmanagementhubManagementStationClient})
	RegisterOracleClient("oci_os_management_hub.OnboardingClient", &OracleClient{InitClientFn: initOsmanagementhubOnboardingClient})
	RegisterOracleClient("oci_os_management_hub.WorkRequestClient", &OracleClient{InitClientFn: initOsmanagementhubWorkRequestClient})
	RegisterOracleClient("oci_os_management_hub.SoftwareSourceClient", &OracleClient{InitClientFn: initOsmanagementhubSoftwareSourceClient})
}

func initOsmanagementhubLifecycleEnvironmentClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_os_management_hub.NewLifecycleEnvironmentClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) LifecycleEnvironmentClient() *oci_os_management_hub.LifecycleEnvironmentClient {
	return m.GetClient("oci_os_management_hub.LifecycleEnvironmentClient").(*oci_os_management_hub.LifecycleEnvironmentClient)
}

func initOsmanagementhubManagedInstanceGroupClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_os_management_hub.NewManagedInstanceGroupClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ManagedInstanceGroupClient() *oci_os_management_hub.ManagedInstanceGroupClient {
	return m.GetClient("oci_os_management_hub.ManagedInstanceGroupClient").(*oci_os_management_hub.ManagedInstanceGroupClient)
}

func initOsmanagementhubManagementStationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_os_management_hub.NewManagementStationClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ManagementStationClient() *oci_os_management_hub.ManagementStationClient {
	return m.GetClient("oci_os_management_hub.ManagementStationClient").(*oci_os_management_hub.ManagementStationClient)
}

func initOsmanagementhubOnboardingClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_os_management_hub.NewOnboardingClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OnboardingClient() *oci_os_management_hub.OnboardingClient {
	return m.GetClient("oci_os_management_hub.OnboardingClient").(*oci_os_management_hub.OnboardingClient)
}

func initOsmanagementhubWorkRequestClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_os_management_hub.NewWorkRequestClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OsManagementHubWorkRequestClient() *oci_os_management_hub.WorkRequestClient {
	return m.GetClient("oci_os_management_hub.WorkRequestClient").(*oci_os_management_hub.WorkRequestClient)
}

func initOsmanagementhubSoftwareSourceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_os_management_hub.NewSoftwareSourceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) SoftwareSourceClient() *oci_os_management_hub.SoftwareSourceClient {
	return m.GetClient("oci_os_management_hub.SoftwareSourceClient").(*oci_os_management_hub.SoftwareSourceClient)
}
