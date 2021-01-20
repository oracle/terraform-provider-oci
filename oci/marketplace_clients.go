// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_marketplace "github.com/oracle/oci-go-sdk/v33/marketplace"

	oci_common "github.com/oracle/oci-go-sdk/v33/common"
)

func init() {
	RegisterOracleClient("oci_marketplace.MarketplaceClient", &OracleClient{initClientFn: initMarketplaceMarketplaceClient})
}

func initMarketplaceMarketplaceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_marketplace.NewMarketplaceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) marketplaceClient() *oci_marketplace.MarketplaceClient {
	return m.GetClient("oci_marketplace.MarketplaceClient").(*oci_marketplace.MarketplaceClient)
}
