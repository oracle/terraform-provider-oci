// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_datacatalog "github.com/oracle/oci-go-sdk/v65/datacatalog"

	oci_common "github.com/oracle/oci-go-sdk/v65/common"
)

func init() {
	RegisterOracleClient("oci_datacatalog.DataCatalogClient", &OracleClient{InitClientFn: initDatacatalogDataCatalogClient})
}

func initDatacatalogDataCatalogClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_datacatalog.NewDataCatalogClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DataCatalogClient() *oci_datacatalog.DataCatalogClient {
	return m.GetClient("oci_datacatalog.DataCatalogClient").(*oci_datacatalog.DataCatalogClient)
}
