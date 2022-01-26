// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package client

import (
	oci_datascience "github.com/oracle/oci-go-sdk/v56/datascience"

	oci_common "github.com/oracle/oci-go-sdk/v56/common"
)

func init() {
	RegisterOracleClient("oci_datascience.DataScienceClient", &OracleClient{InitClientFn: initDatascienceDataScienceClient})
}

func initDatascienceDataScienceClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient, serviceClientOverrides ServiceClientOverrides) (interface{}, error) {
	client, err := oci_datascience.NewDataScienceClientWithConfigurationProvider(configProvider)
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

func (m *OracleClients) DataScienceClient() *oci_datascience.DataScienceClient {
	return m.GetClient("oci_datascience.DataScienceClient").(*oci_datascience.DataScienceClient)
}
