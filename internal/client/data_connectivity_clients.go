// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_data_connectivity "github.com/oracle/oci-go-sdk/v58/dataconnectivity"

	oci_common "github.com/oracle/oci-go-sdk/v58/common"
)

func init() {
	RegisterOracleClient("oci_data_connectivity.DataConnectivityManagementClient", &OracleClient{InitClientFn: initDataconnectivityDataConnectivityManagementClient})
	RegisterOracleClient("oci_data_connectivity.NetworkValidationClient", &OracleClient{InitClientFn: initDataconnectivityNetworkValidationClient})
}

func initDataconnectivityDataConnectivityManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_data_connectivity.NewDataConnectivityManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DataConnectivityManagementClient() *oci_data_connectivity.DataConnectivityManagementClient {
	return m.GetClient("oci_data_connectivity.DataConnectivityManagementClient").(*oci_data_connectivity.DataConnectivityManagementClient)
}

func initDataconnectivityNetworkValidationClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_data_connectivity.NewNetworkValidationClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) NetworkValidationClient() *oci_data_connectivity.NetworkValidationClient {
	return m.GetClient("oci_data_connectivity.NetworkValidationClient").(*oci_data_connectivity.NetworkValidationClient)
}
