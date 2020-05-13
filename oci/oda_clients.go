// Copyright (c) 2020, Oracle and/or its affiliates. All rights reserved.

package oci

import (
	oci_oda "github.com/oracle/oci-go-sdk/oda"

	oci_common "github.com/oracle/oci-go-sdk/common"
)

func init() {
	RegisterOracleClient("oci_oda.OdaClient", &OracleClient{initClientFn: initOdaOdaClient})
}

func initOdaOdaClient(configProvider oci_common.ConfigurationProvider, configureClient ConfigureClient) (interface{}, error) {
	client, err := oci_oda.NewOdaClientWithConfigurationProvider(configProvider)
	if err != nil {
		return nil, err
	}
	err = configureClient(&client.BaseClient)
	if err != nil {
		return nil, err
	}
	return &client, nil
}

func (m *OracleClients) odaClient() *oci_oda.OdaClient {
	return m.GetClient("oci_oda.OdaClient").(*oci_oda.OdaClient)
}
