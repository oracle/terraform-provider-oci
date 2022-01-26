// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_service_catalog "github.com/oracle/oci-go-sdk/v56/servicecatalog"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_service_catalog.ServiceCatalogClient", &OracleClient{InitClientFn: initServicecatalogServiceCatalogClient})
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

	if serviceClientOverrides.HostUrlOverride != "" {
		client.Host = serviceClientOverrides.HostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) ServiceCatalogClient() *oci_service_catalog.ServiceCatalogClient {
	return m.GetClient("oci_service_catalog.ServiceCatalogClient").(*oci_service_catalog.ServiceCatalogClient)
}
