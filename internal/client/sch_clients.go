// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_sch "github.com/oracle/oci-go-sdk/v65/sch"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_sch.ServiceConnectorClient", &OracleClient{InitClientFn: initSchServiceConnectorClient})
}

func initSchServiceConnectorClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_sch.NewServiceConnectorClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ServiceConnectorClient() *oci_sch.ServiceConnectorClient {
	return m.GetClient("oci_sch.ServiceConnectorClient").(*oci_sch.ServiceConnectorClient)
}
