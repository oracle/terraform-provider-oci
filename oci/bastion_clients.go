// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_bastion "github.com/oracle/oci-go-sdk/v45/bastion"

	oci_common "github.com/oracle/oci-go-sdk/v45/common"
)

func init() {
	RegisterOracleClient("oci_bastion.BastionClient", &OracleClient{initClientFn: initBastionBastionClient})
}

func initBastionBastionClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_bastion.NewBastionClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) bastionClient() *oci_bastion.BastionClient {
	return m.GetClient("oci_bastion.BastionClient").(*oci_bastion.BastionClient)
}
