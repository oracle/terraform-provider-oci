// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_oci_product_catalog "github.com/oracle/oci-go-sdk/v65/ociproductcatalog"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_oci_product_catalog.ProductClient", &OracleClient{InitClientFn: initOciproductcatalogProductClient})
	RegisterOracleClient("oci_oci_product_catalog.ProductAdminClient", &OracleClient{InitClientFn: initOciproductcatalogProductAdminClient})
	RegisterOracleClient("oci_oci_product_catalog.ProductInternalClient", &OracleClient{InitClientFn: initOciproductcatalogProductInternalClient})
}

func initOciproductcatalogProductClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_oci_product_catalog.NewProductClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ProductClient() *oci_oci_product_catalog.ProductClient {
	return m.GetClient("oci_oci_product_catalog.ProductClient").(*oci_oci_product_catalog.ProductClient)
}

func initOciproductcatalogProductAdminClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_oci_product_catalog.NewProductAdminClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ProductAdminClient() *oci_oci_product_catalog.ProductAdminClient {
	return m.GetClient("oci_oci_product_catalog.ProductAdminClient").(*oci_oci_product_catalog.ProductAdminClient)
}

func initOciproductcatalogProductInternalClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_oci_product_catalog.NewProductInternalClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) ProductInternalClient() *oci_oci_product_catalog.ProductInternalClient {
	return m.GetClient("oci_oci_product_catalog.ProductInternalClient").(*oci_oci_product_catalog.ProductInternalClient)
}
