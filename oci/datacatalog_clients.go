// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_datacatalog "github.com/oracle/oci-go-sdk/v31/datacatalog"

	oci_common "github.com/oracle/oci-go-sdk/v31/common"
)

func init() {
	RegisterOracleClient("oci_datacatalog.DataCatalogClient", &OracleClient{initClientFn: initDatacatalogDataCatalogClient})
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

	if serviceClientOverrides.hostUrlOverride != "" {
		client.Host = serviceClientOverrides.hostUrlOverride
	}
	return &client, nil
}

func (m *OracleClients) dataCatalogClient() *oci_datacatalog.DataCatalogClient {
	return m.GetClient("oci_datacatalog.DataCatalogClient").(*oci_datacatalog.DataCatalogClient)
}
