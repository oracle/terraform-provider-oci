// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_datacc "github.com/oracle/oci-go-sdk/v65/datacc"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_datacc.BaseinfraClient", &OracleClient{InitClientFn: initDataccBaseinfraClient})
}

func initDataccBaseinfraClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_datacc.NewBaseinfraClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) BaseinfraClient() *oci_datacc.BaseinfraClient {
	return m.GetClient("oci_datacc.BaseinfraClient").(*oci_datacc.BaseinfraClient)
}
