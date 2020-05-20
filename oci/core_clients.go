// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_core "github.com/oracle/oci-go-sdk/core"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_core.BlockstorageClient", &OracleClient{initClientFn: initCoreBlockstorageClient})
	RegisterOracleClient("oci_core.ComputeClient", &OracleClient{initClientFn: initCoreComputeClient})
	RegisterOracleClient("oci_core.ComputeManagementClient", &OracleClient{initClientFn: initCoreComputeManagementClient})
	RegisterOracleClient("oci_core.VirtualNetworkClient", &OracleClient{initClientFn: initCoreVirtualNetworkClient})
}

func initCoreBlockstorageClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_core.NewBlockstorageClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) blockstorageClient() *oci_core.BlockstorageClient {
	return m.GetClient("oci_core.BlockstorageClient").(*oci_core.BlockstorageClient)
}

func initCoreComputeClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_core.NewComputeClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) computeClient() *oci_core.ComputeClient {
	return m.GetClient("oci_core.ComputeClient").(*oci_core.ComputeClient)
}

func initCoreComputeManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_core.NewComputeManagementClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) computeManagementClient() *oci_core.ComputeManagementClient {
	return m.GetClient("oci_core.ComputeManagementClient").(*oci_core.ComputeManagementClient)
}

func initCoreVirtualNetworkClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_core.NewVirtualNetworkClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) virtualNetworkClient() *oci_core.VirtualNetworkClient {
	return m.GetClient("oci_core.VirtualNetworkClient").(*oci_core.VirtualNetworkClient)
}
