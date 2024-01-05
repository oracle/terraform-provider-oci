// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_cloud_bridge "github.com/oracle/oci-go-sdk/v65/cloudbridge"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_cloud_bridge.CommonClient", &OracleClient{InitClientFn: initCloudbridgeCommonClient})
	RegisterOracleClient("oci_cloud_bridge.DiscoveryClient", &OracleClient{InitClientFn: initCloudbridgeDiscoveryClient})
	RegisterOracleClient("oci_cloud_bridge.InventoryClient", &OracleClient{InitClientFn: initCloudbridgeInventoryClient})
	RegisterOracleClient("oci_cloud_bridge.OcbAgentSvcClient", &OracleClient{InitClientFn: initCloudbridgeOcbAgentSvcClient})
}

func initCloudbridgeCommonClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_cloud_bridge.NewCommonClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) CommonClient() *oci_cloud_bridge.CommonClient {
	return m.GetClient("oci_cloud_bridge.CommonClient").(*oci_cloud_bridge.CommonClient)
}

func initCloudbridgeDiscoveryClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_cloud_bridge.NewDiscoveryClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DiscoveryClient() *oci_cloud_bridge.DiscoveryClient {
	return m.GetClient("oci_cloud_bridge.DiscoveryClient").(*oci_cloud_bridge.DiscoveryClient)
}

func initCloudbridgeInventoryClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_cloud_bridge.NewInventoryClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) InventoryClient() *oci_cloud_bridge.InventoryClient {
	return m.GetClient("oci_cloud_bridge.InventoryClient").(*oci_cloud_bridge.InventoryClient)
}

func initCloudbridgeOcbAgentSvcClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_cloud_bridge.NewOcbAgentSvcClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OcbAgentSvcClient() *oci_cloud_bridge.OcbAgentSvcClient {
	return m.GetClient("oci_cloud_bridge.OcbAgentSvcClient").(*oci_cloud_bridge.OcbAgentSvcClient)
}
