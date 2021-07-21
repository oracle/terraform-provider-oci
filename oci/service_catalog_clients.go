// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_service_catalog "github.com/oracle/oci-go-sdk/v45/servicecatalog"

	oci_common "github.com/oracle/oci-go-sdk/v45/common"
)

func init() {
	RegisterOracleClient("oci_service_catalog.ServiceCatalogClient", &OracleClient{initClientFn: initServicecatalogServiceCatalogClient})
}

func initServicecatalogServiceCatalogClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_service_catalog.NewServiceCatalogClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) serviceCatalogClient() *oci_service_catalog.ServiceCatalogClient {
	return m.GetClient("oci_service_catalog.ServiceCatalogClient").(*oci_service_catalog.ServiceCatalogClient)
}
