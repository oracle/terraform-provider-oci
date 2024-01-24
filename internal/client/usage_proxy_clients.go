// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_usage_proxy "github.com/oracle/oci-go-sdk/v65/usage"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_usage_proxy.ResourcesClient", &OracleClient{InitClientFn: initUsageResourcesClient})
	RegisterOracleClient("oci_usage_proxy.RewardsClient", &OracleClient{InitClientFn: initUsageRewardsClient})
	RegisterOracleClient("oci_usage_proxy.UsagelimitsClient", &OracleClient{InitClientFn: initUsageUsagelimitsClient})
}

func initUsageResourcesClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_usage_proxy.NewResourcesClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ResourcesClient() *oci_usage_proxy.ResourcesClient {
	return m.GetClient("oci_usage_proxy.ResourcesClient").(*oci_usage_proxy.ResourcesClient)
}

func initUsageRewardsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_usage_proxy.NewRewardsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) RewardsClient() *oci_usage_proxy.RewardsClient {
	return m.GetClient("oci_usage_proxy.RewardsClient").(*oci_usage_proxy.RewardsClient)
}

func initUsageUsagelimitsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_usage_proxy.NewUsagelimitsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) UsagelimitsClient() *oci_usage_proxy.UsagelimitsClient {
	return m.GetClient("oci_usage_proxy.UsagelimitsClient").(*oci_usage_proxy.UsagelimitsClient)
}
