// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package oci

import (
	oci_datascience "github.com/oracle/oci-go-sdk/v26/datascience"

	oci_common "github.com/oracle/oci-go-sdk/v26/common"
)

func init() {
	RegisterOracleClient("oci_datascience.DataScienceClient", &OracleClient{initClientFn: initDatascienceDataScienceClient})
}

func initDatascienceDataScienceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_datascience.NewDataScienceClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) dataScienceClient() *oci_datascience.DataScienceClient {
	return m.GetClient("oci_datascience.DataScienceClient").(*oci_datascience.DataScienceClient)
}
