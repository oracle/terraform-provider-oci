// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_adm "github.com/oracle/oci-go-sdk/v65/adm"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_adm.ApplicationDependencyManagementClient", &OracleClient{InitClientFn: initAdmApplicationDependencyManagementClient})
}

func initAdmApplicationDependencyManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_adm.NewApplicationDependencyManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ApplicationDependencyManagementClient() *oci_adm.ApplicationDependencyManagementClient {
	return m.GetClient("oci_adm.ApplicationDependencyManagementClient").(*oci_adm.ApplicationDependencyManagementClient)
}
