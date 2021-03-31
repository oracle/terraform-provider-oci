// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_blockchain "github.com/oracle/oci-go-sdk/v38/blockchain"

	oci_common "github.com/oracle/oci-go-sdk/v38/common"
)

func init() {
	RegisterOracleClient("oci_blockchain.BlockchainPlatformClient", &OracleClient{initClientFn: initBlockchainBlockchainPlatformClient})
}

func initBlockchainBlockchainPlatformClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_blockchain.NewBlockchainPlatformClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) blockchainPlatformClient() *oci_blockchain.BlockchainPlatformClient {
	return m.GetClient("oci_blockchain.BlockchainPlatformClient").(*oci_blockchain.BlockchainPlatformClient)
}
