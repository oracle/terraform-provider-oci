// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_marketplace "github.com/oracle/oci-go-sdk/marketplace"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_marketplace.MarketplaceClient", &OracleClient{initClientFn: initMarketplaceMarketplaceClient})
}

func initMarketplaceMarketplaceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_marketplace.NewMarketplaceClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) marketplaceClient() *oci_marketplace.MarketplaceClient {
	return m.GetClient("oci_marketplace.MarketplaceClient").(*oci_marketplace.MarketplaceClient)
}
