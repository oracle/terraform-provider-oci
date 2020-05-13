// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_data_safe "github.com/oracle/oci-go-sdk/datasafe"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_data_safe.DataSafeClient", &OracleClient{initClientFn: initDatasafeDataSafeClient})
}

func initDatasafeDataSafeClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_data_safe.NewDataSafeClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) dataSafeClient() *oci_data_safe.DataSafeClient {
	return m.GetClient("oci_data_safe.DataSafeClient").(*oci_data_safe.DataSafeClient)
}
