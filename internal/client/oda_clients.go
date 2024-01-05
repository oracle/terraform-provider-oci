// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_oda "github.com/oracle/oci-go-sdk/v65/oda"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_oda.ManagementClient", &OracleClient{InitClientFn: initOdaManagementClient})
	RegisterOracleClient("oci_oda.OdaClient", &OracleClient{InitClientFn: initOdaOdaClient})
}

func initOdaManagementClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_oda.NewManagementClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ManagementClient() *oci_oda.ManagementClient {
	return m.GetClient("oci_oda.ManagementClient").(*oci_oda.ManagementClient)
}

func initOdaOdaClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_oda.NewOdaClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) OdaClient() *oci_oda.OdaClient {
	return m.GetClient("oci_oda.OdaClient").(*oci_oda.OdaClient)
}
