// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_os_management_hub "github.com/oracle/oci-go-sdk/v65/osmanagementhub"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_os_management_hub.EventClient", &OracleClient{InitClientFn: initOsmanagementhubEventClient})
	RegisterOracleClient("oci_os_management_hub.LifecycleEnvironmentClient", &OracleClient{InitClientFn: initOsmanagementhubLifecycleEnvironmentClient})
	RegisterOracleClient("oci_os_management_hub.ManagedInstanceClient", &OracleClient{InitClientFn: initOsmanagementhubManagedInstanceClient})
	RegisterOracleClient("oci_os_management_hub.ManagedInstanceGroupClient", &OracleClient{InitClientFn: initOsmanagementhubManagedInstanceGroupClient})
	RegisterOracleClient("oci_os_management_hub.ManagementStationClient", &OracleClient{InitClientFn: initOsmanagementhubManagementStationClient})
	RegisterOracleClient("oci_os_management_hub.OnboardingClient", &OracleClient{InitClientFn: initOsmanagementhubOnboardingClient})
	RegisterOracleClient("oci_os_management_hub.WorkRequestClient", &OracleClient{InitClientFn: initOsmanagementhubWorkRequestClient})
	RegisterOracleClient("oci_os_management_hub.ScheduledJobClient", &OracleClient{InitClientFn: initOsmanagementhubScheduledJobClient})
	RegisterOracleClient("oci_os_management_hub.SoftwareSourceClient", &OracleClient{InitClientFn: initOsmanagementhubSoftwareSourceClient})
}

func initOsmanagementhubEventClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_os_management_hub.NewEventClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OsmhEventClient() *oci_os_management_hub.EventClient {
	return m.GetClient("oci_os_management_hub.EventClient").(*oci_os_management_hub.EventClient)
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

func initOsmanagementhubManagedInstanceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_os_management_hub.NewManagedInstanceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ManagedInstanceClient() *oci_os_management_hub.ManagedInstanceClient {
	return m.GetClient("oci_os_management_hub.ManagedInstanceClient").(*oci_os_management_hub.ManagedInstanceClient)
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

func initOsmanagementhubScheduledJobClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_os_management_hub.NewScheduledJobClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ScheduledJobClient() *oci_os_management_hub.ScheduledJobClient {
	return m.GetClient("oci_os_management_hub.ScheduledJobClient").(*oci_os_management_hub.ScheduledJobClient)
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
