// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_osmanagement "github.com/oracle/oci-go-sdk/v48/osmanagement"

	oci_common "github.com/oracle/oci-go-sdk/v48/common"
)

func init() {
	RegisterOracleClient("oci_osmanagement.EventClient", &OracleClient{initClientFn: initOsmanagementEventClient})
	RegisterOracleClient("oci_osmanagement.OsManagementClient", &OracleClient{initClientFn: initOsmanagementOsManagementClient})
}

func initOsmanagementEventClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osmanagement.NewEventClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) eventClient() *oci_osmanagement.EventClient {
	return m.GetClient("oci_osmanagement.EventClient").(*oci_osmanagement.EventClient)
}

func initOsmanagementOsManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_osmanagement.NewOsManagementClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) osManagementClient() *oci_osmanagement.OsManagementClient {
	return m.GetClient("oci_osmanagement.OsManagementClient").(*oci_osmanagement.OsManagementClient)
}
