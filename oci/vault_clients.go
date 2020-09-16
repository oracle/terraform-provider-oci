// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_vault "github.com/oracle/oci-go-sdk/v25/vault"

	oci_common "github.com/oracle/oci-go-sdk/v25/common"
)

func init() {
	RegisterOracleClient("oci_vault.VaultsClient", &OracleClient{initClientFn: initVaultVaultsClient})
}

func initVaultVaultsClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_vault.NewVaultsClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) vaultsClient() *oci_vault.VaultsClient {
	return m.GetClient("oci_vault.VaultsClient").(*oci_vault.VaultsClient)
}
