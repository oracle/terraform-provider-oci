// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_apiaccesscontrol "github.com/oracle/oci-go-sdk/v65/apiaccesscontrol"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_apiaccesscontrol.ApiMetadataClient", &OracleClient{InitClientFn: initApiaccesscontrolApiMetadataClient})
	RegisterOracleClient("oci_apiaccesscontrol.PrivilegedApiWorkRequestClient", &OracleClient{InitClientFn: initApiaccesscontrolPrivilegedApiWorkRequestClient})
	RegisterOracleClient("oci_apiaccesscontrol.PrivilegedApiControlClient", &OracleClient{InitClientFn: initApiaccesscontrolPrivilegedApiControlClient})
	RegisterOracleClient("oci_apiaccesscontrol.PrivilegedApiRequestsClient", &OracleClient{InitClientFn: initApiaccesscontrolPrivilegedApiRequestsClient})
}

func initApiaccesscontrolApiMetadataClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apiaccesscontrol.NewApiMetadataClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ApiMetadataClient() *oci_apiaccesscontrol.ApiMetadataClient {
	return m.GetClient("oci_apiaccesscontrol.ApiMetadataClient").(*oci_apiaccesscontrol.ApiMetadataClient)
}

func initApiaccesscontrolPrivilegedApiWorkRequestClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apiaccesscontrol.NewPrivilegedApiWorkRequestClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ApiaccesscontrolPrivilegedApiWorkRequestClient() *oci_apiaccesscontrol.PrivilegedApiWorkRequestClient {
	return m.GetClient("oci_apiaccesscontrol.PrivilegedApiWorkRequestClient").(*oci_apiaccesscontrol.PrivilegedApiWorkRequestClient)
}

func initApiaccesscontrolPrivilegedApiControlClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apiaccesscontrol.NewPrivilegedApiControlClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) PrivilegedApiControlClient() *oci_apiaccesscontrol.PrivilegedApiControlClient {
	return m.GetClient("oci_apiaccesscontrol.PrivilegedApiControlClient").(*oci_apiaccesscontrol.PrivilegedApiControlClient)
}

func initApiaccesscontrolPrivilegedApiRequestsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_apiaccesscontrol.NewPrivilegedApiRequestsClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) PrivilegedApiRequestsClient() *oci_apiaccesscontrol.PrivilegedApiRequestsClient {
	return m.GetClient("oci_apiaccesscontrol.PrivilegedApiRequestsClient").(*oci_apiaccesscontrol.PrivilegedApiRequestsClient)
}
